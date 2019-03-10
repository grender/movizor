package movizor

import (
	"net/url"
	"reflect"
	"testing"
)

//func Test_Test (t *testing.T) {
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
			o:    Object("9123456787"),
			want: "79123456787",
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

//
//func TestCoordinate_Float32(t *testing.T) {
//	tests := []struct {
//		name string
//		c    Coordinate
//		want float32
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := tt.c.Float32(); got != tt.want {
//				t.Errorf("Coordinate.Float32() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestCoordinate_String(t *testing.T) {
//	tests := []struct {
//		name string
//		c    Coordinate
//		want string
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := tt.c.String(); got != tt.want {
//				t.Errorf("Coordinate.String() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestCoordinate_UnmarshalJSON(t *testing.T) {
//	type args struct {
//		data []byte
//	}
//	tests := []struct {
//		name    string
//		c       *Coordinate
//		args    args
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if err := tt.c.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
//				t.Errorf("Coordinate.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}
//
//func TestTime_Time(t *testing.T) {
//	tests := []struct {
//		name string
//		t    Time
//		want time.Time
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			t := Time{}
//			if got := t.Time(); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("Time.Time() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestTime_UnmarshalJSON(t *testing.T) {
//	type args struct {
//		data []byte
//	}
//	tests := []struct {
//		name    string
//		t       *Time
//		args    args
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			t := &Time{}
//			if err := t.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
//				t.Errorf("Time.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}
//
//func TestInt_Int(t *testing.T) {
//	tests := []struct {
//		name string
//		i    Int
//		want int
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := tt.i.Int(); got != tt.want {
//				t.Errorf("Int.Int() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestInt_String(t *testing.T) {
//	tests := []struct {
//		name string
//		i    Int
//		want string
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := tt.i.String(); got != tt.want {
//				t.Errorf("Int.String() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestInt_UnmarshalJSON(t *testing.T) {
//	type args struct {
//		data []byte
//	}
//	tests := []struct {
//		name    string
//		i       *Int
//		args    args
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if err := tt.i.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
//				t.Errorf("Int.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}
//
//func Test_jsonNumberToFloat32(t *testing.T) {
//	type args struct {
//		number json.Number
//	}
//	tests := []struct {
//		name    string
//		args    args
//		want    float32
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			got, err := jsonNumberToFloat32(tt.args.number)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("jsonNumberToFloat32() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if got != tt.want {
//				t.Errorf("jsonNumberToFloat32() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
