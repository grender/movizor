package movizor

import (
	"encoding/json"
	"time"
)

func jsonNumberToTime(number json.Number) (time.Time, error) {
	val, err := number.Int64()
	if err != nil {
		return time.Unix(0, 0), err
	}
	return time.Unix(val, 0), nil
}

func jsonNumberToTimePointer(number *json.Number) (*time.Time, error) {
	if number == nil {
		return nil, nil
	}

	val, err := jsonNumberToTime(*number)
	if err != nil {
		return nil, err
	}
	return &val, nil
}

func jsonNumberToFloat32(number json.Number) (float32, error) {
	val, err := number.Float64()
	if err != nil {
		return 0.0, err
	}
	return float32(val), nil
}

func jsonNumberToFloat32Pointer(number *json.Number) (*float32, error) {
	if number == nil {
		return nil, nil
	}

	val, err := jsonNumberToFloat32(*number)
	if err != nil {
		return nil, err
	}
	return &val, nil
}
