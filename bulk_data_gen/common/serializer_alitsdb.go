package common

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"sort"
	"strings"
	"sync"

	alitsdb_serialization "github.com/caict-benchmark/BDC-TS/alitsdb_serializaition"
	cmap "github.com/orcaman/concurrent-map"
)

var (
	sortedTagkeysCache = cmap.New()
)

//DateTimeStdFormat the standard string format for date time
const DateTimeStdFormat = "2006-01-02 15:04:05.000"

//SerieskeyDelimeter is the delimeter of series key
const SerieskeyDelimeter = ','

//KeyValuePairDelimeter is the delimeter of key value
const KeyValuePairDelimeter = '='

//SortedTagKeys saves the sorted tagkeys (alphabet order) and an array with original key indexes in the same order of sorted tagkeys
type SortedTagKeys struct {
	sort.StringSlice
	orgIndexs []int
}

func (stk *SortedTagKeys) Len() int {
	return stk.StringSlice.Len()
}

func (stk *SortedTagKeys) Swap(leftidx, rightidx int) {
	stk.StringSlice.Swap(leftidx, rightidx)
	stk.orgIndexs[leftidx], stk.orgIndexs[rightidx] = stk.orgIndexs[rightidx], stk.orgIndexs[leftidx]
}

func (stk *SortedTagKeys) Less(leftidx, rightidx int) bool {
	return (strings.Compare(stk.StringSlice[leftidx], stk.StringSlice[rightidx]) < 0)
}

func NewSortedTagKeys(tagkeys [][]byte) *SortedTagKeys {
	tagKeyStrings := make([]string, len(tagkeys))
	for i, bytes := range tagkeys {
		tagKeyStrings[i] = string(bytes)
	}
	ret := &SortedTagKeys{StringSlice: sort.StringSlice(tagKeyStrings), orgIndexs: make([]int, len(tagkeys))}
	for i := range ret.orgIndexs {
		ret.orgIndexs[i] = i // original index
	}
	sort.Sort(ret)
	return ret
}

type SerializerAliTSDBHttp struct {
}

type SerializerAliTSDB struct {
	lock sync.Mutex
	fields bool
}

func NewSerializerAliTSDBHttp() *SerializerAliTSDBHttp {
	return &SerializerAliTSDBHttp{}
}

func NewSerializerAliTSDB() *SerializerAliTSDB {
	serializerAliTSDB := &SerializerAliTSDB{}

	serializerAliTSDB.fields = false

	return serializerAliTSDB
}

// MultiFieldsJSONPoint defines the data structure of AliTSDB mput interface
type MultiFieldsJSONPoint struct {
	Metric    string             `json:"metric"`
	Timestamp int64              `json:"timestamp"`
	Tags      map[string]string  `json:"tags"`
	Fields    map[string]float64 `json:"fields"`
}

// SerializePoint writes JSON lines that looks like:
// { <metric>, <timestamp>, <fields>, <tags> }
//
func (m *SerializerAliTSDBHttp) SerializePoint(w io.Writer, p *Point) (err error) {

	encoder := json.NewEncoder(w)

	wp := MultiFieldsJSONPoint{}
	// Timestamps in AliTSDB must be millisecond precision:
	wp.Timestamp = p.Timestamp.UTC().UnixNano() / 1e6
	// sanity check
	{
		x := fmt.Sprintf("%d", wp.Timestamp)
		if len(x) != 13 {
			panic("serialized timestamp was not 13 digits")
		}
	}
	wp.Tags = make(map[string]string, len(p.TagKeys))
	for i := 0; i < len(p.TagKeys); i++ {
		// so many allocs..
		key := string(p.TagKeys[i])
		val := string(p.TagValues[i])
		wp.Tags[key] = val
	}
	// metric name
	wp.Metric = string(p.MeasurementName)

	// fields allocation
	wp.Fields = make(map[string]float64, len(p.FieldKeys))

	// for each Value, generate a new line in the output:
	for i := 0; i < len(p.FieldKeys); i++ {
		switch x := p.FieldValues[i].(type) {
		case int:
			wp.Fields[string(p.FieldKeys[i])] = float64(x)
		case int64:
			wp.Fields[string(p.FieldKeys[i])] = float64(x)
		case float32:
			wp.Fields[string(p.FieldKeys[i])] = float64(x)
		case float64:
			wp.Fields[string(p.FieldKeys[i])] = float64(x)
		default:
			panic("bad numeric value for AliTSDB serialization")
		}
	}
	err = encoder.Encode(wp)
	if err != nil {
		return err
	}

	return nil
}

func (s *SerializerAliTSDBHttp) SerializeSize(w io.Writer, points int64, values int64) error {
	//return serializeSizeInText(w, points, values)
	return nil
}

type task struct {
	point *Point
	w io.Writer
}

func (m *SerializerAliTSDB) handleTask(w io.Writer, p *Point) {
	var mp alitsdb_serialization.MputRequest
	var wp alitsdb_serialization.MputPoint
	mp.Points = make([]*alitsdb_serialization.MputPoint, 1)
	mp.Points[0] = &wp
	//wp.Reset()

	// Timestamps in AliTSDB must be millisecond precision:
	wp.Timestamp = p.Timestamp.UTC().UnixNano() / 1e6
	// sanity check
	{
		x := fmt.Sprintf("%d", wp.Timestamp)
		if len(x) != 13 {
			panic("serialized timestamp was not 13 digits")
		}
	}

	// series key allocation
	var sortedKeys *SortedTagKeys
	if sortedTagkeysCache.Has(string(p.MeasurementName)) {
		tmp, ok := sortedTagkeysCache.Get(string(p.MeasurementName))
		if !ok {
			log.Fatalf("measurement \"%s\" lost in the concurrent access\n", string(p.MeasurementName))
		}

		sortedKeys, ok = tmp.(*SortedTagKeys)
		if !ok {
			log.Fatalf("the value retrieved from sortedTagkeysCache is not the expected type")
		}
	} else {
		sortedKeys = NewSortedTagKeys(p.TagKeys)
		sortedTagkeysCache.SetIfAbsent(string(p.MeasurementName), sortedKeys)
	}

	var serieskeyBuf bytes.Buffer
	serieskeyBuf.Write(p.MeasurementName)
	for i := 0; i < len(sortedKeys.StringSlice); i++ {
		// append the ",""
		serieskeyBuf.WriteByte(byte(SerieskeyDelimeter))
		serieskeyBuf.WriteString(sortedKeys.StringSlice[i])
		serieskeyBuf.WriteByte(byte(KeyValuePairDelimeter))
		serieskeyBuf.Write(p.TagValues[sortedKeys.orgIndexs[i]])
	}
	wp.Serieskey = serieskeyBuf.String()

	// fields allocation
	mp.Fnames = make([]string, len(p.FieldKeys))
	wp.Fvalues = make([]float64, len(p.FieldKeys))

	// for each Value, generate a new line in the output:
	for i := 0; i < len(p.FieldKeys); i++ {
		mp.Fnames[i] = string(p.FieldKeys[i])
		switch x := p.FieldValues[i].(type) {
		case int:
			wp.Fvalues[i] = float64(x)
		case int64:
			wp.Fvalues[i] = float64(x)
		case float32:
			wp.Fvalues[i] = float64(x)
		case float64:
			wp.Fvalues[i] = float64(x)
		default:
			panic("bad numeric value for AliTSDB serialization")
		}
	}

	// write to the out stream
	out, err := mp.Marshal()
	if err != nil {
		log.Fatal(err)
	}
	s := uint64(len(out))
	binary.Write(w, binary.LittleEndian, s)
	w.Write(out)
}

func (m *SerializerAliTSDB) SerializePoint(w io.Writer, p *Point) (err error) {
	m.handleTask(w, p)
	return nil
}

func (m *SerializerAliTSDB) SerializeSize(w io.Writer, points int64, values int64) error {
	//return serializeSizeInText(w, points, values)
	return nil
}
