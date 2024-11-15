package constants

import "errors"

var (
	// QUERY PARAMS
	Err_INVALID_START_DATE_FORMAT error = errors.New("invalid 'start_date' format. must be in 'yyyy-mm-dd' format")
	Err_INVALID_END_DATE_FORMAT   error = errors.New("invalid 'end_date' format. must be in 'yyyy-mm-dd' format")

	// AUTH
	Err_INVALID_TOKEN    error = errors.New("token is invalid")
	Err_INVALID_EMAIL    error = errors.New("email is invalid")
	Err_INVALID_PASSWORD error = errors.New("password is invalid")

	// GOOGLE MAP
	Err_GET_DISTANCE_FAILED    error = errors.New("unable to calculate distance")
	Err_GET_COORDINATES_FAILED error = errors.New("no results found for the address")

	// HASH PASSWORD
	Err_INVALID_HASH         error = errors.New("the encoded hash is not in the correct format")
	Err_INCOMPATIBLE_VERSION error = errors.New("incompatible version of argon2")
)
