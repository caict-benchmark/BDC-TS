package main

import (
	"context"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	alitsdb_serialization "github.com/caict-benchmark/BDC-TS/alitsdb_serializaition"
	cmap "github.com/orcaman/concurrent-map"

	"github.com/caict-benchmark/BDC-TS/bulk_data_gen/common"

	"google.golang.org/grpc"
)

// MputAttemptsLimit indicates the max attemps which the mput can tries
const MputAttemptsLimit = 4

var (
	fieldNameCache = cmap.New()
)

type RpcWriter struct {
	c          WriterConfig
	url        string
	close 	   bool
	pointsChan chan *alitsdb_serialization.MputRequest
}

var logcount = 0

// WriteLineProtocol returns the latency in nanoseconds and any error received while sending the data over RPC,
// or it returns a new error if the RPC response isn't as expected.
func (w *RpcWriter) WriteLineProtocol(client *Client, req *alitsdb_serialization.MputRequest) (latencyNs int64, err error) {
	start := time.Now()

	if doLoad {
		retries := MputAttemptsLimit

		for retries > 0 {
			last := time.Now()
			//TODO: send the write request
			ctx, cel := context.WithTimeout(context.Background(), time.Second*120)
			defer cel()
			resp, err := client.client.Mput(ctx, req)
			now := time.Now()
			dur := now.Sub(last).Milliseconds()
			if dur > 4000 {
				logcount++
				if logcount % 1000 == 0 {
					log.Printf("Timeout request mput points(%d): %dms %s\n", len(req.Points), dur, client.url)
				}
			}

			if err == nil {
				if !resp.Ret {
					log.Println("[WARN] mput request succeeded but retval is false")
				}
				// request succeeded so no need to retry
				break
			} else {
				log.Printf("Error request mput interface(%d: %d): %s %s\n", len(req.Fnames), len(req.Points), client.url, err.Error())
				retries--

				// wait a while
				time.Sleep(time.Duration((MputAttemptsLimit-retries)*10) * time.Second)
				// then start to retry
				client.close()
				if client.init() != nil {
					/* init failed */
					log.Println("[WARN] MultiFieldsPutServiceClient initialization failed")
					retries = 0
				}
			}
		}

		// it means all attempts failed when the retries decreased to zero
		if retries == 0 {
			log.Fatalf("[Fatal]Error caused all retry attempts failed")
		}
	}

	lat := time.Since(start).Nanoseconds()

	return lat, err
}

// NewRPCWriter returns a new RPCWriter from the supplied WriterConfig.
func NewRPCWriter(c WriterConfig) LineProtocolWriter {
	writer := &RpcWriter{
		c:   c,
		url: fmt.Sprintf("%s:%d", c.Host, c.Port),
	}

	client := newClient(writer.url)
	err := client.init()
	if err != nil {
		log.Fatalf("Error connecting: %s\n", err.Error())
	}

	client.close()

	writer.pointsChan = make(chan *alitsdb_serialization.MputRequest, batchSize*100)
	writer.close = false

	return writer
}

func (w *RpcWriter) Close() {
	close(w.pointsChan)
	w.close = true
}

func (w *RpcWriter) PutPoint(point *alitsdb_serialization.MputRequest) {
	if debug {
		size := len(w.pointsChan)
		if size <= batchSize && size > 0 {
			fmt.Printf("size: %d\n", size)
		}
	}
	w.pointsChan <- point
}

type Client struct {
	url    string
	conn   *grpc.ClientConn
	client alitsdb_serialization.MultiFieldsPutServiceClient
}

func newClient(url string) *Client {
	return &Client{url: url}
}

func (c *Client) close() {
	c.conn.Close()
}

func (c *Client) init() error {
	retries := 0
	for {
		conn, err := grpc.Dial(c.url, grpc.WithInsecure())
		if err != nil {
			log.Printf("Error connecting: %s\n", err.Error())
			retries++
			if retries > 3 {
				return err
			}
			time.Sleep(time.Duration(retries*1) * time.Second)
		}

		c.conn = conn
		c.client = alitsdb_serialization.NewMultiFieldsPutServiceClient(c.conn)
		return nil
	}
}

var requestPool = sync.Pool{
	New: func() interface{} {
		return new(alitsdb_serialization.MputRequest)
	},
}

// ProcessBatches read the data from input stream and write by batch
func (w *RpcWriter) ProcessBatches(doLoad bool, bufPool *sync.Pool, wg *sync.WaitGroup, backoff time.Duration, backingOffChan chan bool) {
	client := newClient(w.url)
	if client.init() != nil {
		return
	}

	buff := make([]*alitsdb_serialization.MputRequest, 0, batchSize)
	var n int

	for {
		var basePoint *alitsdb_serialization.MputRequest
		timeout := false
		tick := time.NewTicker(time.Second)

		select {
		case basePoint = <-w.pointsChan:
			if basePoint != nil {
				buff = append(buff, basePoint)
				n++
			}
		case <-tick.C:
			timeout = true
		}

		tick.Stop()

		if n > 0 && (n >= batchSize || timeout) {
			var err error
			for {
				req := requestPool.Get().(*alitsdb_serialization.MputRequest)
				//req := new(alitsdb_serialization.MputRequest)
				req.Points = make([]*alitsdb_serialization.MputPoint, len(buff))
				for i, p := range buff {
					req.Points[i] = p.Points[0]
					req.Fnames = p.Fnames
				}
				_, err = w.WriteLineProtocol(client, req)
				req.Reset()
				requestPool.Put(req)

				for _, p := range buff {
					tasksGroup.Done()
					p.Reset()
					pointPool.Put(p)
				}

				backingOffChan <- false
				break
			}
			if err != nil {
				log.Fatalf("Error writing: %s\n", err.Error())
			}

			n = 0
			buff = nil
			buff = make([]*alitsdb_serialization.MputRequest, 0, batchSize)
		}

		if w.close {
			break
		}
	}

	wg.Done()
}

func getMetric(mp *alitsdb_serialization.MultifieldPoint) string {
	firstDeli := strings.IndexByte(mp.GetSerieskey(), byte(common.SerieskeyDelimeter))
	var metric string
	if firstDeli < 0 {
		//not found
		metric = mp.Serieskey
	} else {
		metric = mp.Serieskey[:firstDeli]
	}

	return metric
}
