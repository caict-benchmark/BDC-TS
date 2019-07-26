package elasticsearch

import "time"
import bulkQuerygen "github.com/caict-benchmark/BDC-TS/bulk_query_gen"

// ElasticSearchDevopsSingleHost produces ES-specific queries for the devops single-host case.
type ElasticSearchVehicleRealTime struct {
	ElasticSearchVehicle
}

func NewElasticSearchVehicleRealTime(_ bulkQuerygen.DatabaseConfig, queriesFullRange bulkQuerygen.TimeInterval, queryInterval time.Duration, scaleVar int) bulkQuerygen.QueryGenerator {
	underlying := NewElasticSearchVehicle(queriesFullRange, scaleVar, queryInterval).(*ElasticSearchVehicle)
	return &ElasticSearchVehicleRealTime{
		ElasticSearchVehicle: *underlying,
	}
}

func (d *ElasticSearchVehicleRealTime) Dispatch(i int) bulkQuerygen.Query {
	q := bulkQuerygen.NewHTTPQuery() // from pool
	d.RealTimeQueries(q)
	return q
}
