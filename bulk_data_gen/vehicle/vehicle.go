package vehicle

import (
	"fmt"
	. "github.com/influxdata/influxdb-comparisons/bulk_data_gen/common"
	"time"
)

var (
	// The duration of a log epoch.
	EpochDuration = 60 * time.Second

	// Tag fields common to all inside sensors:
	RoomTagKey = []byte("room_id")

	// Tag fields common to all inside sensors:
	SensorHomeTagKeys = [][]byte{
		[]byte("sensor_id"),
		[]byte("home_id"),
	}
)

// Mark 表的数量
const NVehicleSims = 1

// Type Host models a machine being monitored by Telegraf.
type Vehicle struct {
	SimulatedMeasurements []SimulatedMeasurement

	// These are all assigned once, at Host creation:
	Name         []byte
}

func NewHostMeasurements(start time.Time) []SimulatedMeasurement {
	sm := []SimulatedMeasurement{
		NewEntityMeasurement(start),
	}

	if len(sm) != NVehicleSims {
		panic("logic error: incorrect number of measurements")
	}
	return sm
}

func NewVehicle(i int, offset int, start time.Time) Vehicle {
	sm := NewHostMeasurements(start)

	h := Vehicle{
		// Tag Values that are static throughout the life of a Host:
		Name:               []byte(fmt.Sprintf("host_%d", i+offset)),
		SimulatedMeasurements: sm,
	}

	return h
}

// TickAll advances all Distributions of a Host.
func (v *Vehicle) TickAll(d time.Duration) {
	for i := range v.SimulatedMeasurements {
		v.SimulatedMeasurements[i].Tick(d)
	}
}

func (v *Vehicle) NumMeasurements() int {
	return len(v.SimulatedMeasurements)
}
