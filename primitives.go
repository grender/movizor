package movizor

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"
	"unicode"
)

// Object - это номер подключаемого абонента в формате MSISDN
// (например, 79210010203).
// Возможно так же передавать номер при добавлении в систему в
// следующих форматах:( +7 (921) 001-02-03; 8-921-001-02-03).
type Object string

// Stringer returns clean format of cell number.
// Casting string(Object) gives Original value.
// fmt.Println(v), fmt.Printf("%s",Object), fmt.Printf("%v",Object) return clean format.
func (o Object) String() string {
	f := func(c rune) rune {
		if unicode.IsNumber(c) {
			return c
		}
		return '&'
	}
	num := strings.Replace(strings.Map(f, string(o)), "&", "", -1)

	for i := 0; i < len(num); i++ {
		if num[i:i+1] == "9" {
			num = num[i:]
			break
		}
	}

	if len(num) >= 10 {
		num := num[len(num)-10:]
		return "7" + num
	}

	return ""
}

func (o Object) values() (url.Values, error) {
	p := o.String()
	if p == "" {
		return url.Values{}, fmt.Errorf("invalid format of phone number: %s (MSISDN: %s) , should be 79XXXXXXXXX", string(o), o.String())
	}

	return url.Values{"phone": {p}}, nil
}

// Coordinate - гео-координата.
// ToDo: сделать проверку на допустимые значения и возможно заменить на float64
type Coordinate float32

// Float32 возвращает гео-координату в виде float32.
func (c Coordinate) Float32() float32 {
	return float32(c)
}

// String возвращает гео-координату в виде строки формата "%.8f"
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

// Time - временная метка.
type Time time.Time

// Time возвращает временную метку в виде time.Time.
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

// Int - определение типа int для unmarshaling json с возможным
// значением json null.
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
