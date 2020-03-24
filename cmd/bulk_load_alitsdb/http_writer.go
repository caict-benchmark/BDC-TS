package main

import (
	"bytes"
	"fmt"
	"log"
	"sync"
	"time"

	alitsdb_serialization "github.com/caict-benchmark/BDC-TS/alitsdb_serializaition"

	"github.com/valyala/fasthttp"
)

var (
	BackoffError      error  = fmt.Errorf("backpressure is needed")
	backoffMagicWords []byte = []byte("engine: cache maximum memory size exceeded")
)

// LineProtocolWriter is the interface used to write AliTSDB bulk data.
type LineProtocolWriter interface {
	// WriteLineProtocol writes the given byte slice containing bulk data
	// to an implementation-specific remote server.
	// Returns the latency, in nanoseconds, of executing the write against the remote server and applicable errors.
	// Implementers must return errors returned by the underlying transport but are free to return
	// other, context-specific errors.

	//WriteLineProtocol([]byte) (latencyNs int64, err error)
	ProcessBatches(doLoad bool, bufPool *sync.Pool, wg *sync.WaitGroup, backoff time.Duration, backingOffChan chan bool)
	PutPoint(point *alitsdb_serialization.MputRequest)
}

// HTTPWriterConfig is the configuration used to create an HTTPWriter.
type WriterConfig struct {
	// URL of the host, in form "http://example.com:8086"
	Host string
	Port int
}

// HTTPWriter is a Writer that writes to an AliTSDB HTTP server.
type HTTPWriter struct {
	client fasthttp.Client

	c   WriterConfig
	url []byte
}

// NewHTTPWriter returns a new HTTPWriter from the supplied HTTPWriterConfig.
func NewHTTPWriter(c WriterConfig) LineProtocolWriter {
	return &HTTPWriter{
		client: fasthttp.Client{
			Name: "bulk_load_alitsdb",
		},

		c:   c,
		url: []byte(fmt.Sprintf("http://%s:%d/api/mput", c.Host, c.Port)),
	}
}

var (
	post                  = []byte("POST")
	applicationJsonHeader = []byte("application/json")
)

func (w *HTTPWriter) PutPoint(point *alitsdb_serialization.MputRequest) {
}

// WriteLineProtocol writes the given byte slice to the HTTP server described in the Writer's HTTPWriterConfig.
// It returns the latency in nanoseconds and any error received while sending the data over HTTP,
// or it returns a new error if the HTTP response isn't as expected.
func (w *HTTPWriter) WriteLineProtocol(body []byte) (int64, error) {
	req := fasthttp.AcquireRequest()
	req.Header.SetContentTypeBytes(applicationJsonHeader)
	req.Header.Set("Content-Encoding", "gzip")
	req.Header.SetMethodBytes(post)
	req.Header.SetRequestURIBytes(w.url)
	req.SetBody(body)

	resp := fasthttp.AcquireResponse()
	start := time.Now()
	err := w.client.Do(req, resp)
	lat := time.Since(start).Nanoseconds()
	if err == nil {
		sc := resp.StatusCode()
		//if sc == 500 && backpressurePred(resp.Body()) {
		//	err = BackoffError
		if sc != fasthttp.StatusNoContent && sc != fasthttp.StatusOK {
			if sc == 500 && backpressurePred(resp.Body()) {
				err = BackoffError
			} else {
				err = fmt.Errorf("Invalid write response (status %d): %s", sc, resp.Body())
			}
		}
	}

	fasthttp.ReleaseResponse(resp)
	fasthttp.ReleaseRequest(req)

	return lat, err
}

// ProcessBatches read the data from input stream and write by batch
func (w *HTTPWriter) ProcessBatches(doLoad bool, bufPool *sync.Pool, wg *sync.WaitGroup, backoff time.Duration, backingOffChan chan bool) {
	for batch := range batchChan {
		// Write the batch: try until backoff is not needed.
		if doLoad {
			var err error
			for {
				_, err = w.WriteLineProtocol(batch.Bytes())
				if err == BackoffError {
					backingOffChan <- true
					time.Sleep(backoff)
				} else {
					backingOffChan <- false
					break
				}
			}
			if err != nil {
				log.Fatalf("Error writing: %s\n", err.Error())
			}
		}

		// Return the batch buffer to the pool.
		batch.Reset()
		bufPool.Put(batch)
	}
	wg.Done()
}

func backpressurePred(body []byte) bool {
	return bytes.Contains(body, backoffMagicWords)
}
