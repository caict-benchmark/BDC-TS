package vehicle

import (
	. "github.com/caict-benchmark/BDC-TS/bulk_data_gen/common"
	"time"
)

// Type IotSimulatorConfig is used to create a IotSimulator.
type VehicleSimulatorConfig struct {
	Start time.Time
	End   time.Time

	VehicleCount int64
	VehicleOffset int64
}

func (d *VehicleSimulatorConfig) ToSimulator() *VehicleSimulator {
	vehicleInfos := make([]Vehicle, d.VehicleCount)
	var measNum int64

	for i := 0; i < len(vehicleInfos); i++ {
		//vehicleInfos[i] = NewSmartHome(i, int(d.SmartHomeOffset), d.Start)
		vehicleInfos[i] = NewVehicle(i, int(d.VehicleOffset), d.Start)
		measNum += int64(vehicleInfos[i].NumMeasurements())
	}

	epochs := d.End.Sub(d.Start).Nanoseconds() / EpochDuration.Nanoseconds()
	maxPoints := epochs * measNum
	dg := &VehicleSimulator{
		madePoints: 0,
		madeValues: 0,
		maxPoints:  maxPoints,

		currentVehicleIndex: 0,
		vehicles:            vehicleInfos,

		timestampNow:   d.Start,
		timestampStart: d.Start,
		timestampEnd:   d.End,
	}

	return dg
}

// A IotSimulator generates data similar to telemetry from Telegraf.
// It fulfills the Simulator interface.
type VehicleSimulator struct {
	madePoints    int64
	maxPoints     int64
	madeValues    int64

	simulatedMeasurementIndex int

	currentVehicleIndex int
	vehicles            []Vehicle

	timestampNow   time.Time
	timestampStart time.Time
	timestampEnd   time.Time
}

func (g *VehicleSimulator) SeenPoints() int64 {
	return g.madePoints
}

func (g *VehicleSimulator) SeenValues() int64 {
	return g.madeValues
}

func (g *VehicleSimulator) Total() int64 {
	return g.maxPoints
}

func (g *VehicleSimulator) Finished() bool {
	return g.madePoints >= g.maxPoints
}

// Next advances a Point to the next state in the generator.
func (v *VehicleSimulator) Next(p *Point) {
	// switch to the next metric if needed
	if v.currentVehicleIndex == len(v.vehicles) {
		v.currentVehicleIndex = 0
		v.simulatedMeasurementIndex++
	}

	if v.simulatedMeasurementIndex == NVehicleSims {
		v.simulatedMeasurementIndex = 0

		for i := 0; i < len(v.vehicles); i++ {
			v.vehicles[i].TickAll(EpochDuration)
		}
	}

	vehicle := &v.vehicles[v.currentVehicleIndex]

	//// Populate host-specific tags:
	//p.AppendTag(MachineTagKeys[0], host.Name)
	//p.AppendTag(MachineTagKeys[1], host.Region)
	//p.AppendTag(MachineTagKeys[2], host.Datacenter)
	//p.AppendTag(MachineTagKeys[3], host.Rack)
	//p.AppendTag(MachineTagKeys[4], host.OS)
	//p.AppendTag(MachineTagKeys[5], host.Arch)
	//p.AppendTag(MachineTagKeys[6], host.Team)
	//p.AppendTag(MachineTagKeys[7], host.Service)
	//p.AppendTag(MachineTagKeys[8], host.ServiceVersion)
	//p.AppendTag(MachineTagKeys[9], host.ServiceEnvironment)

	// Populate measurement-specific tags and fields:
	vehicle.SimulatedMeasurements[v.simulatedMeasurementIndex].ToPoint(p)

	v.madePoints++
	v.currentVehicleIndex++
	v.madeValues += int64(len(p.FieldValues))

	return
}
