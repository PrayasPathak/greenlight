package data

import (
	"fmt"
	"strconv"
)

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
