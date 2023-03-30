package prometheus

import (
	"encoding/json"
	"strconv"
	"time"
)

// Value - A single value from a prometheus query
type Value struct {
	Timestamp time.Time
	Value     float64
}

// UnmarshalJSON - Unmarshal the json response from prometheus
func (v *Value) UnmarshalJSON(b []byte) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()
	var intermediateValue []interface{}
	if err := json.Unmarshal(b, &intermediateValue); err != nil {
		return err
	}

	value, err := strconv.ParseFloat(intermediateValue[1].(string), 64)
	if err != nil {
		return err
	}

	timestamp := intermediateValue[0].(float64)
	seconds := int64(timestamp)
	nanoseconds := int64((timestamp - float64(seconds)) * 1e9)

	v.Timestamp = time.Unix(seconds, nanoseconds)
	v.Value = value

	return nil
}
