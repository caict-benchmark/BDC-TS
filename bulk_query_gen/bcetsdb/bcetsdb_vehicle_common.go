package bcetsdb

import (
	"bytes"
	"fmt"
	bulkQuerygen "github.com/caict-benchmark/BDC-TS/bulk_query_gen"
	"text/template"
	"time"
	"io"
)

var (
	vehicleRealTimeQuery *template.Template
)

func init() {
	vehicleRealTimeQuery = template.Must(template.New("vehicleRealTimeQuery").Parse(rawVehicleRealTimeQuery))
}

type BceTSDBVehicle struct {
	bulkQuerygen.CommonParams
	bulkQuerygen.TimeWindow
}

func NewBceTSDBVehicle(interval bulkQuerygen.TimeInterval, scaleVar int, duration time.Duration) bulkQuerygen.QueryGenerator {
	return &BceTSDBVehicle{
		CommonParams: *bulkQuerygen.NewCommonParams(interval, scaleVar),
		TimeWindow:   bulkQuerygen.TimeWindow{interval.Start, time.Second},
	}
}

// Dispatch fulfills the QueryGenerator interface.
func (d *BceTSDBVehicle) Dispatch(i int) bulkQuerygen.Query {
	q := bulkQuerygen.NewHTTPQuery() // from pool
	return q
}

func (d *BceTSDBVehicle) RealTimeQueries(q bulkQuerygen.Query) {
	// hard code vin, because I don't know how to change it.
	d.realTimeQueries(q.(*bulkQuerygen.HTTPQuery), time.Second, "LSVNV2182E2100001")
}

func (d *BceTSDBVehicle) realTimeQueries(qi bulkQuerygen.Query, timeRange time.Duration, vin string) {
	var interval bulkQuerygen.TimeInterval
	if bulkQuerygen.TimeWindowShift > 0 {
		interval = d.TimeWindow.SlidingWindow(&d.AllInterval)
	} else {
		interval = d.AllInterval.RandWindow(d.Duration)
	}

	body := new(bytes.Buffer)
	mustExecuteTemplate(vehicleRealTimeQuery, body, VehicleRealTimeParams{
		Start: interval.StartUnix(),
		End:   interval.EndUnix(),
		Size:  30000,
	})

	humanLabel := []byte(fmt.Sprintf("Elastic real time query, rand %s by 1m", timeRange))
	q := qi.(*bulkQuerygen.HTTPQuery)
	q.HumanLabel = humanLabel
	q.HumanDescription = []byte(fmt.Sprintf("%s: %s", humanLabel, interval.StartString()))
	q.Method = []byte("POST")

	q.Path = []byte("/vehicle/_search")
	q.Body = body.Bytes()
}

type VehicleRealTimeParams struct {
	//Vin        string
	Start, End int64
	Size       int
}

const rawVehicleRealTimeQuery = `
{
  "query": {
    "bool": {
      "must": [
        {
          "range": {
            "timestamp": {
              "gte": "{{.Start}}",
              "lte": "{{.End}}"
            }
          }
        }
      ]
    }
  },
  "size": {{.Size}}
}
`

func mustExecuteTemplate(t *template.Template, w io.Writer, params interface{}) {
	err := t.Execute(w, params)
	if err != nil {
		panic(fmt.Sprintf("logic error in executing template: %s", err))
	}
}