package vehicle

import (
	. "github.com/influxdata/influxdb-comparisons/bulk_data_gen/common"
	"math/rand"
	"time"
)

var (
	EntityByteString      = []byte("vehicle-entity")       // heap optimization
	EntityTotalByteString = []byte("vehicle-entity-total") // heap optimization
)

var (
	// Field keys for 'vehicle entity' points.
	EntityFieldKeys = [][]byte{
		[]byte("entity_status"),
		[]byte("entity_charge"),
		[]byte("entity_mode"),
		[]byte("entity_speed"),
		[]byte("entity_mileage"),
		[]byte("entity_voltage"),
		[]byte("entity_current_flows"),
		[]byte("entity_soc"),
		[]byte("entity_dc_dc"),
		[]byte("entity_gears"),
		[]byte("entity_resistance"),
		[]byte("entity_reserve"),
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
