package data

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrInvalidRuntimeFormat = errors.New("invalid runtime format")

type Runtime int32

// MarshalJSON returns the string representation of runtime of a movie
// in minutes surrounded by double quotes.
//
// Example: if the runtime is 102, output returned is: "102 mins", nil
func (r Runtime) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%d mins", r)
	quotedJSONValue := strconv.Quote(jsonValue)
	return []byte(quotedJSONValue), nil
}

// UnmarshalJSON converts the user provided runtime value to int32.
// If an error occurs during unquoting the value, or splitting of
// the data an ErrInvalidRuntimeFormat is returned else nil is returned.
//
// Examples:
//
// curl -d '{"title": "Moana", "runtime": "107 minutes"}' localhost:4000/v1/movies
// returns in invalid runtime format
//
// curl -d '{"title": "Moana", "runtime": 107 }' localhost:4000/v1/movies returns
// in invalid runtime format
//
// curl -d '{"title": "Moana", "runtime": "107 mins"}' localhost:4000/v1/movies returns
// {Title:Moana Year:0 Runtime:107 Genres:[]}
func (r *Runtime) UnmarshalJSON(jsonValue []byte) error {
	unquotedJSONValue, err := strconv.Unquote(string(jsonValue))
	if err != nil {
		return ErrInvalidRuntimeFormat
	}

	parts := strings.Split(unquotedJSONValue, " ")
	if len(parts) != 2 || parts[1] != "mins" {
		return ErrInvalidRuntimeFormat
	}

	i, err := strconv.ParseInt(parts[0], 10, 32)
	if err != nil {
		return ErrInvalidRuntimeFormat
	}
	*r = Runtime(i)
	return nil
}
