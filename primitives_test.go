package movizor

import (
	"net/url"
	"reflect"
	"testing"
	"time"
)

//func Test_Test(t *testing.T) {
//}

func TestObject_String(t *testing.T) {
	tests := []struct {
		name string
		o    Object
		want string
	}{
		{
			name: "nice format 1",
			o:    Object("7912-345-67-87"),
			want: "79123456787",
		},
		{
			name: "nice format 2",
			o:    Object("912-345-67-87"),
			want: "79123456787",
		},
		{
			name: "nice format 3",
			o:    Object("8(912)345-67-87"),
			want: "79123456787",
		},
		{
			name: "nice format 4",
			o:    Object("+7912-345-67-87"),
			want: "79123456787",
		},
		{
			name: "nice format 5",
			o:    Object("79898069996"),
			want: "79898069996",
		},
		{
			name: "not nice format 1",
			o:    Object("+7912-345-67-8"),
			want: "",
		},
		{
			name: "not nice format 2",
			o:    Object("912345678"),
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.String(); got != tt.want {
				t.Errorf("Object.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestObject_values(t *testing.T) {
	tests := []struct {
		name    string
		o       Object
		want    url.Values
		wantErr bool
	}{
		{
			name: "nice format 1",
			o:    Object("7912-345-67-87"),
			want: url.Values{
				"phone": {"79123456787"},
			},
			wantErr: false,
		},
		{
			name: "nice format 2",
			o:    Object("912-345-67-87"),
			want: url.Values{
				"phone": {"79123456787"},
			},
			wantErr: false,
		},
		{
			name: "nice format 3",
			o:    Object("8(912)345-67-87"),
			want: url.Values{
				"phone": {"79123456787"},
			},
			wantErr: false,
		},
		{
			name: "nice format 4",
			o:    Object("+7912-345-67-87"),
			want: url.Values{
				"phone": {"79123456787"},
			},
			wantErr: false,
		},
		{
			name: "nice format 5",
			o:    Object("9123456787"),
			want: url.Values{
				"phone": {"79123456787"},
			},
			wantErr: false,
		},
		{
			name:    "not nice format 1",
			o:       Object("+7912-345-67-8"),
			want:    url.Values{},
			wantErr: true,
		},
		{
			name:    "not nice format 2",
			o:       Object("912345678"),
			want:    url.Values{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.o.values()
			if (err != nil) != tt.wantErr {
				t.Errorf("Object.values() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Object.values() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCoordinate_Float32(t *testing.T) {
	tests := []struct {
		name string
		c    Coordinate
		want float32
	}{
		{
			name: "nice format 1",
			c:    Coordinate(0.0),
			want: 0.0,
		},
		{
			name: "nice format 2",
			c:    Coordinate(180.0),
			want: 180.0,
		},
		{
			name: "nice format 3",
			c:    Coordinate(-180.0),
			want: -180.0,
		},
		{
			name: "nice format 4",
			c:    Coordinate(179.99999999),
			want: 179.99999999,
		},
		{
			name: "nice format 5",
			c:    Coordinate(-0.00000001),
			want: -0.00000001,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Float32(); got != tt.want {
				t.Errorf("Coordinate.Float32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCoordinate_String(t *testing.T) {
	tests := []struct {
		name string
		c    Coordinate
		want string
	}{
		{
			name: "nice format 1",
			c:    Coordinate(0.0),
			want: "0.00000000",
		},
		{
			name: "nice format 2",
			c:    Coordinate(180.0),
			want: "180.00000000",
		},
		{
			name: "nice format 3",
			c:    Coordinate(-180.0),
			want: "-180.00000000",
		},
		//{
		//	name: "nice format 4",
		//	c:    Coordinate(179.99999999),
		//	want: "180.00000000",
		//},
		{
			name: "nice format 5",
			c:    Coordinate(-0.00000001),
			want: "-0.00000001",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.String(); got != tt.want {
				t.Errorf("Coordinate.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCoordinate_UnmarshalJSON(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		c       *Coordinate
		args    args
		wantErr bool
	}{
		{
			name: "nice format 1",
			c:    new(Coordinate),
			args: args{
				data: []byte("0.00000000"),
			},
			wantErr: false,
		},
		{
			name: "nice format 2",
			c:    new(Coordinate),
			args: args{
				data: []byte("180.00000000"),
			},
			wantErr: false,
		},
		{
			name: "nice format 3",
			c:    new(Coordinate),
			args: args{
				data: []byte("-180.00000000"),
			},
			wantErr: false,
		},
		{
			name: "nice format 4",
			c:    new(Coordinate),
			args: args{
				data: []byte(`"180.00000000"`),
			},
			wantErr: false,
		},
		{
			name: "not nice format 1",
			c:    new(Coordinate),
			args: args{
				data: []byte("180,00000000"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Coordinate.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTime_Time(t *testing.T) {
	tim := time.Now()
	tests := []struct {
		name string
		t    Time
		want time.Time
	}{
		{
			name: "nice format 1",
			t:    Time(tim),
			want: tim,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//ti := Time{}
			if got := tt.t.Time(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Time.Time() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTime_UnmarshalJSON(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		t       *Time
		args    args
		wantErr bool
	}{
		{
			name: "nice format 1",
			args: args{
				data: []byte("1548075614"),
			},
			wantErr: false,
		},
		{
			name: "nice format 2",
			args: args{
				data: []byte(`"1548075614"`),
			},
			wantErr: false,
		},
		{
			name: "null",
			args: args{
				data: []byte("null"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.t = &Time{}
			if err := tt.t.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Time.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestInt_Int(t *testing.T) {
	const UintSize = 32 << (^uint(0) >> 32 & 1)

	const (
		MaxInt = 1<<(UintSize-1) - 1 // 1<<31 - 1 or 1<<63 - 1
		MinInt = -MaxInt - 1         // -1 << 31 or -1 << 63
	)

	tests := []struct {
		name string
		i    Int
		want int
	}{
		{
			name: "zero",
			i:    0,
			want: 0,
		},
		{
			name: "max",
			i:    MaxInt,
			want: MaxInt,
		},
		{
			name: "min",
			i:    MinInt,
			want: MinInt,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.Int(); got != tt.want {
				t.Errorf("Int.Int() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_String(t *testing.T) {
	tests := []struct {
		name string
		i    Int
		want string
	}{
		{
			name: "zero",
			i:    0,
			want: "0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.String(); got != tt.want {
				t.Errorf("Int.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_UnmarshalJSON(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		i       *Int
		args    args
		wantErr bool
	}{
		{
			name: "zero",
			i:    new(Int),
			args: args{
				data: []byte("0"),
			},
			wantErr: false,
		},
		{
			name: "quoted zero",
			i:    new(Int),
			args: args{
				data: []byte(`"0"`),
			},
			wantErr: false,
		},
		{
			name: "null",
			i:    new(Int),
			args: args{
				data: []byte("null"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.i.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Int.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
