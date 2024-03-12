package data

import (
	"fmt"
	"strconv"
)

// For custom JSON encoding
type Runtime int32

// Implement a MarshalJSON() method on the Runtime type to satisfy the
// json.Marshaler interface.
func (self Runtime) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%d mins", self)
	// Wrap string in double quotes.
	quotedJSONValue := strconv.Quote(jsonValue)
	return []byte(quotedJSONValue), nil
}
