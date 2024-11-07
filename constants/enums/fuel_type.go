package enums

import "database/sql/driver"

type FuelType string

const (
	PREMIUM FuelType = "PREMIUM"
	SOLAR   FuelType = "SOLAR"
)

func (vt *FuelType) Scan(value interface{}) error {
	*vt = FuelType(value.([]byte))
	return nil
}

func (vt FuelType) Value() (driver.Value, error) {
	return string(vt), nil
}
