package bcetsdb

import "time"
import bulkQuerygen "github.com/caict-benchmark/BDC-TS/bulk_query_gen"

type BceTSDBVehicleRealTime struct {
	BceTSDBVehicle
}

func NewBceTSDBVehicleRealTime(_ bulkQuerygen.DatabaseConfig, queriesFullRange bulkQuerygen.TimeInterval, queryInterval time.Duration, scaleVar int) bulkQuerygen.QueryGenerator {
	underlying := NewBceTSDBVehicle(queriesFullRange, scaleVar, queryInterval).(*BceTSDBVehicle)
	return &BceTSDBVehicleRealTime{
		BceTSDBVehicle: *underlying,
	}
}

func (d *BceTSDBVehicleRealTime) Dispatch(i int) bulkQuerygen.Query {
	q := bulkQuerygen.NewHTTPQuery() // from pool
	d.RealTimeQueries(q)
	return q
}
