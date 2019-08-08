package bcetsdb

import (
    "time"
    "bytes"
    "fmt"
    "math/rand"
    "text/template"
    "strings"
    "regexp"
    // "strconv"
    bulkQuerygen "github.com/caict-benchmark/BDC-TS/bulk_query_gen"
)

type BceTSDBDevops struct {
    // FieldVar string
    // AllInterval bulkQuerygen.TimeInterval
    bulkQuerygen.CommonParams
}

// func newBceTSDBDevopsCommon(dbConfig bulkQuerygen.DatabaseConfig, start, end time.Time) bulkQuerygen.QueryGenerator {
//     if !start.Before(end) {
//         panic("bad time order")
//     }

//     return &BceTSDBDevops{
//         // FieldVar: dbConfig["field-var"],
//         AllInterval: NewTimeInterval(start, end),
//     }
// }

func newBceTSDBDevopsCommon(queriesFullRange bulkQuerygen.TimeInterval, scaleVar int) bulkQuerygen.QueryGenerator {
	return &BceTSDBDevops{
		CommonParams: *bulkQuerygen.NewCommonParams(queriesFullRange, scaleVar),
	}
}

// func (d *BceTSDBDevops) Dispatch(i, scaleVar int) bulkQuerygen.Query {
//     q := NewBceTSDBQuery() // from pool
//     devopsDispatchAll(d, i, q, scaleVar)
//     return q
// }

// Dispatch fulfills the QueryGenerator interface.
func (d *BceTSDBDevops) Dispatch(i int) bulkQuerygen.Query {
	q := bulkQuerygen.NewHTTPQuery() // from pool
	bulkQuerygen.DevopsDispatchAll(d, i, q, d.ScaleVar)
	return q
}

// func (d *BceTSDBDevops) MaxCPUUsageHourByMinuteOneHost(q bulkQuerygen.Query, scaleVar int) {
//     d.maxCPUUsageHourByMinuteNHosts(q.(*BceTSDBQuery), scaleVar, 1, time.Hour)
// }

// func (d *BceTSDBDevops) MaxCPUUsageHourByMinuteTwoHosts(q bulkQuerygen.Query, scaleVar int) {
//     d.maxCPUUsageHourByMinuteNHosts(q.(*BceTSDBQuery), scaleVar, 2, time.Hour)
// }

// func (d *BceTSDBDevops) MaxCPUUsageHourByMinuteFourHosts(q bulkQuerygen.Query, scaleVar int) {
//     d.maxCPUUsageHourByMinuteNHosts(q.(*BceTSDBQuery), scaleVar, 4, time.Hour)
// }

// func (d *BceTSDBDevops) MaxCPUUsageHourByMinuteEightHosts(q bulkQuerygen.Query, scaleVar int) {
//     d.maxCPUUsageHourByMinuteNHosts(q.(*BceTSDBQuery), scaleVar, 8, time.Hour)
// }

// func (d *BceTSDBDevops) MaxCPUUsageHourByMinuteSixteenHosts(q bulkQuerygen.Query, scaleVar int) {
//     d.maxCPUUsageHourByMinuteNHosts(q.(*BceTSDBQuery), scaleVar, 16, time.Hour)
// }

// func (d *BceTSDBDevops) MaxCPUUsageHourByMinuteThirtyTwoHosts(q bulkQuerygen.Query, scaleVar int) {
//     d.maxCPUUsageHourByMinuteNHosts(q.(*BceTSDBQuery), scaleVar, 32, time.Hour)
// }

// func (d *BceTSDBDevops) MaxCPUUsageHourByMinute1000Hosts(q bulkQuerygen.Query, scaleVar int) {
//     d.maxCPUUsageHourByMinuteNHosts(q.(*BceTSDBQuery), scaleVar, 1000, time.Hour)
// }

// func (d *BceTSDBDevops) MaxCPUUsage12HoursByMinuteOneHost(q bulkQuerygen.Query, scaleVar int) {
//     d.maxCPUUsageHourByMinuteNHosts(q.(*BceTSDBQuery), scaleVar, 1, 12*time.Hour)
// }

func (d *BceTSDBDevops) MaxCPUUsageHourByMinuteOneHost(q bulkQuerygen.Query) {
    d.maxCPUUsageHourByMinuteNHosts(q.(*bulkQuerygen.HTTPQuery), 1, time.Hour)
}

func (d *BceTSDBDevops) MaxCPUUsageHourByMinuteTwoHosts(q bulkQuerygen.Query) {
    d.maxCPUUsageHourByMinuteNHosts(q.(*bulkQuerygen.HTTPQuery), 2, time.Hour)
}

func (d *BceTSDBDevops) MaxCPUUsageHourByMinuteFourHosts(q bulkQuerygen.Query) {
    d.maxCPUUsageHourByMinuteNHosts(q.(*bulkQuerygen.HTTPQuery), 4, time.Hour)
}

func (d *BceTSDBDevops) MaxCPUUsageHourByMinuteEightHosts(q bulkQuerygen.Query) {
    d.maxCPUUsageHourByMinuteNHosts(q.(*bulkQuerygen.HTTPQuery), 8, time.Hour)
}

func (d *BceTSDBDevops) MaxCPUUsageHourByMinuteSixteenHosts(q bulkQuerygen.Query) {
    d.maxCPUUsageHourByMinuteNHosts(q.(*bulkQuerygen.HTTPQuery), 16, time.Hour)
}

func (d *BceTSDBDevops) MaxCPUUsageHourByMinuteThirtyTwoHosts(q bulkQuerygen.Query) {
    d.maxCPUUsageHourByMinuteNHosts(q.(*bulkQuerygen.HTTPQuery), 32, time.Hour)
}

func (d *BceTSDBDevops) MaxCPUUsageHourByMinute1000Hosts(q bulkQuerygen.Query) {
    d.maxCPUUsageHourByMinuteNHosts(q.(*bulkQuerygen.HTTPQuery), 1000, time.Hour)
}

func (d *BceTSDBDevops) MaxCPUUsage12HoursByMinuteOneHost(q bulkQuerygen.Query) {
    d.maxCPUUsageHourByMinuteNHosts(q.(*bulkQuerygen.HTTPQuery), 1, 12*time.Hour)
}

func (d *BceTSDBDevops) maxCPUUsageHourByMinuteNHosts(qi bulkQuerygen.Query, nhosts int, timeRange time.Duration) {
    interval := d.AllInterval.RandWindow(timeRange)
    // nn := rand.Perm(scaleVar)[:nhosts]
    nn := rand.Perm(1)[:nhosts]
    hostnames := []string{}
    for _, n := range nn {
        hostnames = append(hostnames, fmt.Sprintf("\"%d\"", n))
    }
    combinedHostnameClause := strings.Join(hostnames, ",")

    // fieldVar,errs := strconv.Atoi(d.FieldVar)
    fieldVar := 1 
    // if errs != nil {
    //     panic(errs)
    // }
    fields := []string{}
    cpuFields := []string{"usage_user", "usage_system","usage_idle", "usage_nice","usage_iowait", "usage_irq","usage_softirq", "usage_steal","usage_guest", "usage_guest_nice"}
    for n := 0; n < fieldVar; n++ {
        fields = append(fields, fmt.Sprintf("\"%s\"", cpuFields[n]))
    }
    combinedFieldsClause := strings.Join(fields, ",")

    samplingInterval := "\"1 minutes\""

    startTimestamp := interval.StartUnixNano() / 1e6
    endTimestamp := interval.EndUnixNano() / 1e6

    var tmplString = `
{
    "queries": [{
        "metric": "cpu",
        "fields": [{{.CombinedFieldsClause}}],
        "filters": {
            "start": {{.StartTimestamp}},
            "end": {{.EndTimestamp}},
            "tags": {
                "time_serial_tag": [{{.CombinedHostnameClause}}]
            }
        },
        "limit": 200000,
        "aggregators:": [{
            "name": "Max",
            "sampling": {{.SamplingInterval}}
        }]
    }]
}
`
    var re = regexp.MustCompile("\\n")
    tmplString = re.ReplaceAllString(tmplString, "")
    re = regexp.MustCompile(" ")
    tmplString = re.ReplaceAllString(tmplString, "")
    tmpl := template.Must(template.New("tmpl").Parse(tmplString))
    bodyWriter := new(bytes.Buffer)

    arg := struct {
        StartTimestamp, EndTimestamp int64
        CombinedHostnameClause       string
        CombinedFieldsClause         string
        SamplingInterval             string
    }{
        startTimestamp,
        endTimestamp,
        combinedHostnameClause,
        combinedFieldsClause,
        samplingInterval,
    }
    err := tmpl.Execute(bodyWriter, arg)
    if err != nil {
        panic(err)
    }

    q := qi.(*bulkQuerygen.HTTPQuery)
    q.Body = bodyWriter.Bytes()
    // q.Body = bodyWriter.String()
}

func (d *BceTSDBDevops) MeanCPUUsageDayByHourAllHostsGroupbyHost(qi bulkQuerygen.Query) {
    interval := d.AllInterval.RandWindow(24 * time.Hour)
    startTimestamp := interval.StartUnixNano() / 1e6
    endTimestamp := interval.EndUnixNano() / 1e6
    samplingInterval := "\"1 hours\""

    var tmplString = `
{
    "queries": [{
        "metric": "cpu",
        "field": "user_usage",
        "filters": {
            "start": {{.StartTimestamp}},
            "end": {{.EndTimestamp}}
        },
        "groupBy": [{
            "name": "Tag",
            "tags": ["hostname"]
        }],
        "limit": 200000,
        "aggregators:": [{
            "name": "Avg",
            "sampling": {{.SamplingInterval}}
        }]
    }]
}`
    var re = regexp.MustCompile("\\n")
    tmplString = re.ReplaceAllString(tmplString, "")
    re = regexp.MustCompile(" ")
    tmplString = re.ReplaceAllString(tmplString, "")
    tmpl := template.Must(template.New("tmpl").Parse(tmplString))
    bodyWriter := new(bytes.Buffer)

    arg := struct {
        StartTimestamp, EndTimestamp int64
        SamplingInterval             string
    }{
        startTimestamp,
        endTimestamp,
        samplingInterval,
    }
    err := tmpl.Execute(bodyWriter, arg)
    if err != nil {
        panic("logic error")
    }

    // q := qi.(*BceTSDBQuery)
    // q.Body = bodyWriter.String()
    q := qi.(*bulkQuerygen.HTTPQuery)
    q.Body = bodyWriter.Bytes()
}
