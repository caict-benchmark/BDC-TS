package common

import (
	"encoding/json"
	"io"
)

type SerializerBceTSDB struct {
}

func NewSerializerBceTSDB() *SerializerBceTSDB {
	return &SerializerBceTSDB{}
}

// This function writes JSON lines that looks like:
// { <metric>, <timestamp>, <value>, <tags> }
//
// For example:
// {"datapoints":[{"metric":"cpu","field":"usage_user","timestamp":1451606400000,"tags":{"arch":"x86",
// "datacenter":"us-west-2b","hostname":"host_0","os":"Ubuntu16.04LTS","rack":"43","region":"us-west-2",
// "service":"13","service_environment":"production","service_version":"0","team":"LON"},

func (m *SerializerBceTSDB) SerializePoint(w io.Writer, p *Point) (err error) {
    type valuesPoint struct {
        Metric string            `json:"metric"`
        Field  string            `json:"field"`
        Tags   map[string]string `json:"tags"`
        Values []interface{}        `json:"values"`
    }
    type dataPoints struct {
        PointArray []interface{}      `json:"datapoints"`
    }

	metricBase := string(p.MeasurementName) // will be re-used
	encoder := json.NewEncoder(w)

    vp := valuesPoint{}
	dp := dataPoints{}
	
	dp.PointArray = make([]interface{}, 0)

    // for each Value, generate a new line in the output:
    for i := 0; i < len(p.FieldKeys); i++ {
        vp.Metric = metricBase
        vp.Field = string(p.FieldKeys[i])
        for j := 0; j < int(1); j++ {
            tv := make([]interface{}, 2)
            tv[0] = p.Timestamp.UTC().UnixNano()/1e6 + int64(j)
            switch x := p.FieldValues[i].(type) {
            case int:
                tv[1] = float64(x)
            case int64:
                tv[1] = float64(x)
            case float32:
                tv[1] = float64(x)
            case float64:
                tv[1] = float64(x)
            default:
                panic("bad numeric value for BceTSDB serialization")
            }
            vp.Values = append(vp.Values, tv)
        }
        vp.Tags = make(map[string]string, len(p.TagKeys))
        for i := 0; i < len(p.TagKeys); i++ {
            // so many allocs..
            key := string(p.TagKeys[i])
            val := string(p.TagValues[i])
            vp.Tags[key] = val
        }
        dp.PointArray = append(dp.PointArray, vp)
        vp = valuesPoint{}
    }
    err = encoder.Encode(dp)
    if err != nil {
        return err
    }

	return nil
}

func (s *SerializerBceTSDB) SerializeSize(w io.Writer, points int64, values int64) error {
	return nil
}
