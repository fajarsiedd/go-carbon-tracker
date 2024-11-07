package entities

type User struct {
	Base
	Name     string
	Email    string
	Password string
	Token    string
	Trips    []Trip
	Vehicles []Vehicle
}
