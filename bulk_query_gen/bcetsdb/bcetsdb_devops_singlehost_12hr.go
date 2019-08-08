package bcetsdb

import "time"
import bulkQuerygen "github.com/caict-benchmark/BDC-TS/bulk_query_gen"

type BceTSDBDevopsSingleHost12hr struct {
    BceTSDBDevops
}

func NewBceTSDBDevopsSingleHost12hr(_ bulkQuerygen.DatabaseConfig, queriesFullRange bulkQuerygen.TimeInterval, queryInterval time.Duration, scaleVar int) bulkQuerygen.QueryGenerator {
    underlying := newBceTSDBDevopsCommon(queriesFullRange, scaleVar).(*BceTSDBDevops)
    return &BceTSDBDevopsSingleHost12hr{
        BceTSDBDevops: *underlying,
    }
}

func (d *BceTSDBDevopsSingleHost12hr) Dispatch(i int) bulkQuerygen.Query {
    q := bulkQuerygen.NewHTTPQuery() // from pool
    d.MaxCPUUsage12HoursByMinuteOneHost(q)
    return q
}

// func NewBceTSDBDevopsSingleHost12hr(dbConfig bulkQuerygen.DatabaseConfig, start, end time.Time) QueryGenerator {
//     underlying := newBceTSDBDevopsCommon(dbConfig, start, end).(*BceTSDBDevops)
//     return &BceTSDBDevopsSingleHost12hr{
//         BceTSDBDevops: *underlying,
//     }
// }

// func (d *BceTSDBDevopsSingleHost12hr) Dispatch(i, scaleVar int) Query {
//     q := NewBceTSDBQuery() // from pool
//     d.MaxCPUUsage12HoursByMinuteOneHost(q, scaleVar)
//     return q
// }
