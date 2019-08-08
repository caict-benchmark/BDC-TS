package bcetsdb

import "time"
import bulkQuerygen "github.com/caict-benchmark/BDC-TS/bulk_query_gen"

type BceTSDBDevops8Hosts struct {
    BceTSDBDevops
}

func NewBceTSDBDevops8Hosts(_ bulkQuerygen.DatabaseConfig, queriesFullRange bulkQuerygen.TimeInterval, queryInterval time.Duration, scaleVar int) bulkQuerygen.QueryGenerator {
    underlying := newBceTSDBDevopsCommon(queriesFullRange, scaleVar).(*BceTSDBDevops)
    return &BceTSDBDevops8Hosts{
        BceTSDBDevops: *underlying,
    }
}

func (d *BceTSDBDevops8Hosts) Dispatch(i int) bulkQuerygen.Query {
    q := bulkQuerygen.NewHTTPQuery() // from pool
    d.MaxCPUUsageHourByMinuteEightHosts(q)
    return q
}
