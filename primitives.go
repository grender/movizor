package movizor

import (
	"encoding/json"
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"time"
)

// Номер подключаемого абонента в формате MSISDN (например, 79210010203).
// Возможно так же передавать номер при добавлении в систему в следующих форматах:( +7 (921) 001-02-03; 8-921-001-02-03).
type Object string

// Stringer returns clean format of cell number.
// Casting string(Object) gives Original value.
// fmt.Println(v), fmt.Printf("%s",Object), fmt.Printf("%v",Object) return clean format.
func (o Object) String() string {
	// ToDo: Переписать на что-то более надежное
	return regexp.MustCompile("[^0-9]").ReplaceAllString(string(o), "")
}

func (o Object) values() url.Values {
	return url.Values{"phone": {o.String()}}
}

type Coordinate float32

func (c Coordinate) Float32() float32 {
	return float32(c)
}

func (c Coordinate) String() string {
	return fmt.Sprintf("%.8f", c.Float32())
}

func (c *Coordinate) UnmarshalJSON(data []byte) (err error) {
	var num json.Number
	err = json.Unmarshal(data, &num)
	if err != nil {
		return
	}

	var val float32
	val, err = jsonNumberToFloat32(num)
	if err != nil {
		return
	}

	*c = Coordinate(val)
	return nil
}

type Time time.Time

func (t Time) Time() time.Time {
	return time.Time(t)
}

func (t *Time) UnmarshalJSON(data []byte) error {
	var num json.Number
	err := json.Unmarshal(data, &num)
	if err != nil {
		return err
	}

	val, err := num.Int64()
	if err != nil {
		return err
	}

	*t = Time(time.Unix(val, 0))
	return nil
}

type Int int

func (i Int) Int() int {
	return int(i)
}

func (i Int) String() string {
	return strconv.FormatInt(int64(i), 10)
}

func (i *Int) UnmarshalJSON(data []byte) error {
	var num json.Number
	err := json.Unmarshal(data, &num)
	if err != nil {
		return err
	}

	val, err := num.Int64()
	if err != nil {
		return err
	}

	*i = Int(val)
	return nil
}

func jsonNumberToFloat32(number json.Number) (float32, error) {
	val, err := number.Float64()
	if err != nil {
		return 0.0, err
	}
	return float32(val), nil
}
