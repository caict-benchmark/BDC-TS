package bcetsdb

import "time"
import bulkQuerygen "github.com/caict-benchmark/BDC-TS/bulk_query_gen"

type BceTSDBDevopsSingleHost struct {
    BceTSDBDevops
}

func NewBceTSDBDevopsSingleHost(_ bulkQuerygen.DatabaseConfig, queriesFullRange bulkQuerygen.TimeInterval, queryInterval time.Duration, scaleVar int) bulkQuerygen.QueryGenerator {
    underlying := newBceTSDBDevopsCommon(queriesFullRange, scaleVar).(*BceTSDBDevops)
    return &BceTSDBDevopsSingleHost12hr{
        BceTSDBDevops: *underlying,
    }
}

func (d *BceTSDBDevopsSingleHost) Dispatch(i int) bulkQuerygen.Query {
    q := bulkQuerygen.NewHTTPQuery() // from pool
    d.MaxCPUUsageHourByMinuteOneHost(q)
    return q
}

// func NewBceTSDBDevopsSingleHost(dbConfig bulkQuerygen.DatabaseConfig, start, end time.Time) QueryGenerator {
//     underlying := newBceTSDBDevopsCommon(dbConfig, start, end).(*BceTSDBDevops)
//     return &BceTSDBDevopsSingleHost{
//         BceTSDBDevops: *underlying,
//     }
// }

// func (d *BceTSDBDevopsSingleHost) Dispatch(i, scaleVar int) Query {
//     q := NewBceTSDBQuery() // from pool
//     d.MaxCPUUsageHourByMinuteOneHost(q, scaleVar)
//     return q
// }
