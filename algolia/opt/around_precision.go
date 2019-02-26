// Code generated by go generate. DO NOT EDIT.

package opt

import "encoding/json"

type AroundPrecisionOption struct {
	value int
}

func AroundPrecision(v int) AroundPrecisionOption {
	return AroundPrecisionOption{v}
}

func (o AroundPrecisionOption) Get() int {
	return o.value
}

func (o AroundPrecisionOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *AroundPrecisionOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = 0
		return nil
	}
	return json.Unmarshal(data, &o.value)
}
