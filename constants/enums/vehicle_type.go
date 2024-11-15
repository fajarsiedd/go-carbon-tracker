package enums

import "database/sql/driver"

type VehicleType string

const (
	MOBIL VehicleType = "MOBIL"
	BUS   VehicleType = "BUS"
	TRUK  VehicleType = "TRUK"
	MOTOR VehicleType = "MOTOR"
)

func (vt *VehicleType) Scan(value interface{}) error {
	*vt = VehicleType(value.([]byte))
	return nil
}

func (vt VehicleType) Value() (driver.Value, error) {
	return string(vt), nil
}
