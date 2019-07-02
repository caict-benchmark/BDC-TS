package vehicle

import (
	. "github.com/caict-benchmark/BDC-TS/bulk_data_gen/common"
	"math/rand"
	"time"
)

var (
	EntityByteString      = []byte("vehicle")       // heap optimization
	EntityTotalByteString = []byte("vehicle-total") // heap optimization
)

var (
	// Field keys for 'vehicle entity' points.
	EntityFieldKeys = [][]byte{
		[]byte("VIN"),
		[]byte("value1"),
		[]byte("value2"),
		[]byte("value3"),
		[]byte("value4"),
		[]byte("value5"),
		[]byte("value6"),
		[]byte("value7"),
		[]byte("value8"),
		[]byte("value9"),
		[]byte("value10"),
		[]byte("value11"),
		[]byte("value12"),
		[]byte("value13"),
		[]byte("value14"),
		[]byte("value15"),
		[]byte("value16"),
		[]byte("value17"),
		[]byte("value18"),
		[]byte("value19"),
		[]byte("value20"),
		[]byte("value21"),
		[]byte("value22"),
		[]byte("value23"),
		[]byte("value24"),
		[]byte("value25"),
		[]byte("value26"),
		[]byte("value27"),
		[]byte("value28"),
		[]byte("value29"),
		[]byte("value30"),
		[]byte("value31"),
		[]byte("value32"),
		[]byte("value33"),
		[]byte("value34"),
		[]byte("value35"),
		[]byte("value36"),
		[]byte("value37"),
		[]byte("value38"),
		[]byte("value39"),
		[]byte("value40"),
		[]byte("value41"),
		[]byte("value42"),
		[]byte("value43"),
		[]byte("value44"),
		[]byte("value45"),
		[]byte("value46"),
		[]byte("value47"),
		[]byte("value48"),
		[]byte("value49"),
		[]byte("value50"),
		[]byte("value51"),
		[]byte("value52"),
		[]byte("value53"),
		[]byte("value54"),
		[]byte("value55"),
		[]byte("value56"),
		[]byte("value57"),
		[]byte("value58"),
		[]byte("value59"),
		[]byte("value60"),
	}
)

type EntityMeasurement struct {
	timestamp     time.Time
	distributions []Distribution
}

func NewEntityMeasurement(start time.Time) *EntityMeasurement {
	distributions := make([]Distribution, len(EntityFieldKeys))
	for i := range distributions {
		distributions[i] = &ClampedRandomWalkDistribution{
			State: rand.Float64() * 100.0,
			Min:   0.0,
			Max:   100.0,
			Step: &NormalDistribution{
				Mean:   0.0,
				StdDev: 1.0,
			},
		}
	}
	return &EntityMeasurement{
		timestamp:     start,
		distributions: distributions,
	}
}

func (m *EntityMeasurement) Tick(d time.Duration) {
	m.timestamp = m.timestamp.Add(d)
	for i := range m.distributions {
		m.distributions[i].Advance()
	}
}

func (m *EntityMeasurement) ToPoint(p *Point) bool {
	p.SetMeasurementName(EntityByteString)
	p.SetTimestamp(&m.timestamp)

	for i := range m.distributions {
		p.AppendField(EntityFieldKeys[i], m.distributions[i].Get())
	}
	return true
}
