package elasticsearch

import (
	"bytes"
	"fmt"
	bulkQuerygen "github.com/caict-benchmark/BDC-TS/bulk_query_gen"
	"text/template"
	"time"
)

var (
	vehicleRealTimeQuery *template.Template
)

func init() {
	vehicleRealTimeQuery = template.Must(template.New("vehicleRealTimeQuery").Parse(rawVehicleRealTimeQuery))
}

// ElasticSearchDevops produces ES-specific queries for the devops use case.
type ElasticSearchVehicle struct {
	bulkQuerygen.CommonParams
	bulkQuerygen.TimeWindow
}

// NewElasticSearchDevops makes an ElasticSearchDevops object ready to generate Queries.
func NewElasticSearchVehicle(interval bulkQuerygen.TimeInterval, scaleVar int, duration time.Duration) bulkQuerygen.QueryGenerator {
	return &ElasticSearchVehicle{
		CommonParams: *bulkQuerygen.NewCommonParams(interval, scaleVar),
		TimeWindow:   bulkQuerygen.TimeWindow{interval.Start, time.Second},
	}
}

// Dispatch fulfills the QueryGenerator interface.
func (d *ElasticSearchVehicle) Dispatch(i int) bulkQuerygen.Query {
	q := bulkQuerygen.NewHTTPQuery() // from pool
	//bulkQuerygen.DevopsDispatchAll(d, i, q, d.ScaleVar)
	return q
}

func (d *ElasticSearchVehicle) RealTimeQueries(q bulkQuerygen.Query) {
	// hard code vin, because I don't know how to change it.
	d.realTimeQueries(q.(*bulkQuerygen.HTTPQuery), time.Second, "LSVNV2182E2100001")
}

func (d *ElasticSearchVehicle) realTimeQueries(qi bulkQuerygen.Query, timeRange time.Duration, vin string) {
	//interval := d.AllInterval.RandWindow(timeRange)
	//interval := d.AllInterval.addTimeWindow(timeRange, d)
	var interval bulkQuerygen.TimeInterval
	if bulkQuerygen.TimeWindowShift > 0 {
		interval = d.TimeWindow.SlidingWindow(&d.AllInterval)
	} else {
		interval = d.AllInterval.RandWindow(d.Duration)
	}

	body := new(bytes.Buffer)
	mustExecuteTemplate(vehicleRealTimeQuery, body, VehicleRealTimeParams{
		Start: interval.StartUnix() * 1000,
		End:   interval.EndUnix() * 1000,
		Size:  20000,
		//Vin:   vin,
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
      "filter": [
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
  "docvalue_fields": ["timestamp", "value4"],
  "size": {{.Size}}
}
`
