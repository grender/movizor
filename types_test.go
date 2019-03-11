package movizor

import (
	"io/ioutil"
	"path/filepath"
	"testing"
)

const dataPath = "./test-data"

//func TestTariff_UnmarshalJSON(t *testing.T) {
//	type fields struct {
//		AbonentPayment float64
//		RequestCost    float64
//		TariffTitle    string
//	}
//	type args struct {
//		data []byte
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			t := &Tariff{
//				AbonentPayment: tt.fields.AbonentPayment,
//				RequestCost:    tt.fields.RequestCost,
//				TariffTitle:    tt.fields.TariffTitle,
//			}
//			if err := t.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
//				t.Errorf("Tariff.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}

func TestBalance_UnmarshalJSON(t *testing.T) {
	type fields struct {
		Balance         float64
		Credit          float64
		ContractType    string
		OperatorTariffs map[Operator]map[TariffType]Tariff
		ServiceTariffs  map[Service][]Tariff
	}

	tests := []struct {
		name     string
		fields   fields
		filename string
		wantErr  bool
	}{
		{
			name:     "balance_good",
			fields:   fields{},
			filename: "balance.json",
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d, err := ioutil.ReadFile(filepath.Join(dataPath, tt.filename))
			if err != nil {
				t.Errorf("err: %s", err)
			}
			b := &Balance{
				Balance:         tt.fields.Balance,
				Credit:          tt.fields.Credit,
				ContractType:    tt.fields.ContractType,
				OperatorTariffs: tt.fields.OperatorTariffs,
				ServiceTariffs:  tt.fields.ServiceTariffs,
			}
			if err := b.UnmarshalJSON(d); (err != nil) != tt.wantErr {
				t.Errorf("Balance.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestObjectInfo_UnmarshalJSON(t *testing.T) {
	type fields struct {
		Phone                 Object
		Status                Status
		Confirmed             bool
		Title                 string
		Tariff                TariffType
		TariffNew             *TariffType
		LastTimestamp         Time
		AtRequest             bool
		CurrentCoordinates    CurrentCoordinates
		CoordinatesAttributes CoordinatesAttributes
		OnParking             *bool
		Destination           []Destination
		OfflineTime           Time
		PosError              bool
		TimestampOff          Time
		TimestampAdd          Time
		Metadata              map[string]string
	}

	tests := []struct {
		name     string
		fields   fields
		filename string
		wantErr  bool
	}{
		{
			name:     "object_get1",
			fields:   fields{},
			filename: "object_get1.json",
			wantErr:  false,
		},
		{
			name:     "object_get2",
			fields:   fields{},
			filename: "object_get2.json",
			wantErr:  false,
		},
		{
			name:     "object_get3",
			fields:   fields{},
			filename: "object_get3.json",
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d, err := ioutil.ReadFile(filepath.Join(dataPath, tt.filename))
			if err != nil {
				t.Errorf("err: %s", err)
			}

			oi := &ObjectInfo{
				Phone:                 tt.fields.Phone,
				Status:                tt.fields.Status,
				Confirmed:             tt.fields.Confirmed,
				Title:                 tt.fields.Title,
				Tariff:                tt.fields.Tariff,
				TariffNew:             tt.fields.TariffNew,
				LastTimestamp:         tt.fields.LastTimestamp,
				AtRequest:             tt.fields.AtRequest,
				CurrentCoordinates:    tt.fields.CurrentCoordinates,
				CoordinatesAttributes: tt.fields.CoordinatesAttributes,
				OnParking:             tt.fields.OnParking,
				Destination:           tt.fields.Destination,
				OfflineTime:           tt.fields.OfflineTime,
				PosError:              tt.fields.PosError,
				TimestampOff:          tt.fields.TimestampOff,
				TimestampAdd:          tt.fields.TimestampAdd,
				Metadata:              tt.fields.Metadata,
			}
			if err := oi.UnmarshalJSON(d); (err != nil) != tt.wantErr {
				t.Errorf("ObjectInfo.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

//func TestObjectsWithStatus_Len(t *testing.T) {
//	tests := []struct {
//		name string
//		os   ObjectsWithStatus
//		want int
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := tt.os.Len(); got != tt.want {
//				t.Errorf("ObjectsWithStatus.Len() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestObjectsWithStatus_Swap(t *testing.T) {
//	type args struct {
//		i int
//		j int
//	}
//	tests := []struct {
//		name string
//		os   ObjectsWithStatus
//		args args
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			tt.os.Swap(tt.args.i, tt.args.j)
//		})
//	}
//}
//
//func TestObjectsWithStatus_Less(t *testing.T) {
//	type args struct {
//		i int
//		j int
//	}
//	tests := []struct {
//		name string
//		os   ObjectsWithStatus
//		args args
//		want bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := tt.os.Less(tt.args.i, tt.args.j); got != tt.want {
//				t.Errorf("ObjectsWithStatus.Less() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestObjectsWithStatus_IsObjectIn(t *testing.T) {
//	type args struct {
//		o Object
//	}
//	tests := []struct {
//		name string
//		os   ObjectsWithStatus
//		args args
//		want bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := tt.os.IsObjectIn(tt.args.o); got != tt.want {
//				t.Errorf("ObjectsWithStatus.IsObjectIn() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestPositionRequest_values(t *testing.T) {
//	type fields struct {
//		RequestID int64
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		want   url.Values
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			pr := PositionRequest{
//				RequestID: tt.fields.RequestID,
//			}
//			if got := pr.values(); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("PositionRequest.values() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

func TestObjectEvent_UnmarshalJSON(t *testing.T) {
	type fields struct {
		EventID   int64
		Timestamp Time
		Phone     Object
		Event     EventType
	}

	tests := []struct {
		name     string
		fields   fields
		filename string
		wantErr  bool
	}{
		{
			name:     "event1",
			fields:   fields{},
			filename: "event1.json",
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d, err := ioutil.ReadFile(filepath.Join(dataPath, tt.filename))
			if err != nil {
				t.Errorf("err: %s", err)
			}

			oe := &ObjectEvent{
				EventID:   tt.fields.EventID,
				Timestamp: tt.fields.Timestamp,
				Phone:     tt.fields.Phone,
				Event:     tt.fields.Event,
			}
			if err := oe.UnmarshalJSON(d); (err != nil) != tt.wantErr {
				t.Errorf("ObjectEvent.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSubscribedEvent_UnmarshalJSON(t *testing.T) {
	type fields struct {
		SubscriptionID         int64
		IsAllObjectsSubscribed bool
		ObjectsSubscribed      []Object
		Timestamp              Time
		Event                  EventType
		Phone                  Object
		EMail                  string
		IsTelegram             bool
	}
	tests := []struct {
		name     string
		fields   fields
		filename string
		wantErr  bool
	}{
		{
			name:     "event_subscription1",
			fields:   fields{},
			filename: "event_subscription1.json",
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d, err := ioutil.ReadFile(filepath.Join(dataPath, tt.filename))
			if err != nil {
				t.Errorf("err: %s", err)
			}
			se := &SubscribedEvent{
				SubscriptionID:         tt.fields.SubscriptionID,
				IsAllObjectsSubscribed: tt.fields.IsAllObjectsSubscribed,
				ObjectsSubscribed:      tt.fields.ObjectsSubscribed,
				Timestamp:              tt.fields.Timestamp,
				Event:                  tt.fields.Event,
				Phone:                  tt.fields.Phone,
				EMail:                  tt.fields.EMail,
				IsTelegram:             tt.fields.IsTelegram,
			}
			if err := se.UnmarshalJSON(d); (err != nil) != tt.wantErr {
				t.Errorf("SubscribedEvent.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

//func TestSubscribedEvent_MakeOptions(t *testing.T) {
//	type fields struct {
//		SubscriptionID         int64
//		IsAllObjectsSubscribed bool
//		ObjectsSubscribed      []Object
//		Timestamp              Time
//		Event                  EventType
//		Phone                  Object
//		EMail                  string
//		IsTelegram             bool
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		wantSeo SubscribeEventOptions
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			se := SubscribedEvent{
//				SubscriptionID:         tt.fields.SubscriptionID,
//				IsAllObjectsSubscribed: tt.fields.IsAllObjectsSubscribed,
//				ObjectsSubscribed:      tt.fields.ObjectsSubscribed,
//				Timestamp:              tt.fields.Timestamp,
//				Event:                  tt.fields.Event,
//				Phone:                  tt.fields.Phone,
//				EMail:                  tt.fields.EMail,
//				IsTelegram:             tt.fields.IsTelegram,
//			}
//			gotSeo, err := se.MakeOptions()
//			if (err != nil) != tt.wantErr {
//				t.Errorf("SubscribedEvent.MakeOptions() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(gotSeo, tt.wantSeo) {
//				t.Errorf("SubscribedEvent.MakeOptions() = %v, want %v", gotSeo, tt.wantSeo)
//			}
//		})
//	}
//}
