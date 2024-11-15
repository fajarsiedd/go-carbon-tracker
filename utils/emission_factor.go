package utils

import (
	"go-carbon-tracker/constants"
	"go-carbon-tracker/constants/enums"
)

func GetCO2Factor(fuelType enums.FuelType) float32 {
	switch fuelType {
	case enums.PREMIUM:
		return constants.CO2_FACTOR_FOR_PREMIUM_FUEL
	case enums.SOLAR:
		return constants.CO2_FACTOR_FOR_SOLAR_FUEL
	default:
		return 0
	}
}

func GetEnergyConsume(vehicleType enums.VehicleType) float32 {
	switch vehicleType {
	case enums.MOTOR:
		return constants.ENERGY_CONSUME_FOR_MOTOR
	case enums.MOBIL:
		return constants.ENERGY_CONSUME_FOR_MOBIL
	case enums.BUS:
		return constants.ENERGY_CONSUME_FOR_BUS
	case enums.TRUK:
		return constants.ENERGY_CONSUME_FOR_TRUK
	default:
		return 0
	}
}
