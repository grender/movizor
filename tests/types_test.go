package tests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"movizor"
	"path/filepath"
	"testing"
)

const dataPath = "./test-data"

func TestErrorResponseUnmarshal(t *testing.T) {
	d, err := ioutil.ReadFile(filepath.Join(dataPath, "error_response.json"))
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	var errResp movizor.APIResponse
	err = json.Unmarshal(d, &errResp)

	if err != nil || errResp.Result != "error" {
		t.Fatalf("Input %s is not parsed to %T.\n\nError: %s", d, movizor.APIResponse{}, err)
	}
}

func TestResponseUnmarshal(t *testing.T) {
	var resp movizor.APIResponse

	d, err := ioutil.ReadFile(filepath.Join(dataPath, "success_response.json"))
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	err = json.Unmarshal(d, &resp)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	if resp.Result != "success" {
		t.Fatalf("Input %s is not parsed to %T.\n\nError: %s", d, movizor.APIResponse{}, err)
	}
}

func TestBalanceDataUnmarshal(t *testing.T) {
	d, err := ioutil.ReadFile(filepath.Join(dataPath, "balance.json"))
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	var bd movizor.Balance
	err = json.Unmarshal(d, &bd)

	if _, ok := bd.TariffPlans["mts"]; err != nil || !ok || bd.Balance != 476.50 || bd.Credit != 0.0 {
		t.Fatalf("Input %s is not parsed to %T.\n\nError: %s", d, movizor.Balance{}, err)
	}
}

func TestPositionUnmarshal(t *testing.T) {
	d1, err := ioutil.ReadFile(filepath.Join(dataPath, "pos_last1.json"))
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	var lp movizor.Position
	err = json.Unmarshal(d1, &lp)

	if err != nil || lp.ETAStatus != movizor.NoETAStatus {
		t.Fatalf("Input %s is not parsed to %T.\n\nError: %s", d1, movizor.Position{}, err)
	}

	d2, err := ioutil.ReadFile(filepath.Join(dataPath, "pos_last2.json"))
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	err = json.Unmarshal(d2, &lp)

	if err != nil || lp.ETAStatus != movizor.NoETAStatus {
		t.Fatalf("Input %s is not parsed to %T.\n\nError: %s", d2, movizor.Position{}, err)
	}

}

func TestPositionsUnmarshal(t *testing.T) {
	d, err := ioutil.ReadFile(filepath.Join(dataPath, "pos_list.json"))
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	var p movizor.Positions
	err = json.Unmarshal(d, &p)

	if err != nil {
		t.Fatalf("Input %s is not parsed to %T.\n\nError: %s", d, movizor.Positions{}, err)
	}

	if p[0].ETAStatus != movizor.NoETAStatus {
		t.Fatalf("Input %s is not parsed to %T.\n\nError: %s", d, movizor.Positions{}, err)
	}
}

func TestObjectsUnmarshal(t *testing.T) {
	d, err := ioutil.ReadFile(filepath.Join(dataPath, "object_list.json"))
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	var osl movizor.ObjectsWithStatus
	err = json.Unmarshal(d, &osl)

	if err != nil || osl[0].Phone != "79050005727" || osl[0].Status != movizor.StatusOff {
		t.Fatalf("Input %s is not parsed to %T.\n\nError: %s", d, movizor.ObjectPositions{}, err)
	}
}

func TestObjectPositionsUnmarshal(t *testing.T) {
	d, err := ioutil.ReadFile(filepath.Join(dataPath, "pos_objects.json"))
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	var po movizor.ObjectPositions
	err = json.Unmarshal(d, &po)

	if err != nil || po[0].Phone != "79630005272" || po[0].ETAStatus != movizor.NoETAStatus {
		fmt.Println(po[0].ETAStatus)
		t.Fatalf("Input %s is not parsed to %T.\n\nError: %s", d, movizor.ObjectPositions{}, err)
	}
}

func TestOperatorInfoUnmarshal(t *testing.T) {
	d, err := ioutil.ReadFile(filepath.Join(dataPath, "get_operator.json"))
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	var o movizor.OperatorInfo
	err = json.Unmarshal(d, &o)

	if err != nil ||
		!(o.Operator == movizor.OperatorMTS ||
			o.Operator == movizor.OperatorMegafon ||
			o.Operator == movizor.OperatorBeeline ||
			o.Operator == movizor.OperatorTele2) {
		t.Fatalf("Input %s is not parsed to %T.\n\nError: %s", d, movizor.ObjectPositions{}, err)
	}
}

func TestObjectInfoUnmarshal(t *testing.T) {
	d, err := ioutil.ReadFile(filepath.Join(dataPath, "object_get1.json"))
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	var oi movizor.ObjectInfo
	err = json.Unmarshal(d, &oi)
	// ToDo: Проверить все виды статусов, а также парсинг ETAStatus
	if err != nil ||
		!(oi.Status == movizor.StatusOff || oi.Status == movizor.StatusOk) {
		t.Fatalf("Input %s is not parsed to %T.\n\nError: %s", d, movizor.ObjectInfo{}, err)
	}
}

func TestObjectString(t *testing.T) {
	v := movizor.Object("+7 (456) 765-43 57")

	if v.String() != "74567654357" {
		t.Fatalf("Object: %s cannot be clean properly", v)
	}
}

func TestObjectEventsUnmarshal(t *testing.T) {
	d, err := ioutil.ReadFile(filepath.Join(dataPath, "events.json"))
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	var e movizor.ObjectEvents
	err = json.Unmarshal(d, &e)

	if err != nil {
		t.Fatalf("Input %s is not parsed to %T.\n\nError: %s", d, movizor.ObjectEvents{}, err)
	}

	for _, v := range e {
		if !(v.Event == movizor.AddEvent ||
			v.Event == movizor.AutoOffEvent ||
			v.Event == movizor.OffEvent ||
			v.Event == movizor.RequestOkEvent ||
			v.Event == movizor.RequestErrorEvent ||
			v.Event == movizor.ConfirmEvent ||
			v.Event == movizor.RejectEvent ||
			v.Event == movizor.RequestObjectOfflineEvent ||
			v.Event == movizor.RequestObjectInRoamingEvent ||
			v.Event == movizor.ReactivateEvent ||
			v.Event == movizor.ChangeTariffEvent ||
			v.Event == movizor.InTimeEvent ||
			v.Event == movizor.LateEvent ||
			v.Event == movizor.FinishedEvent ||
			v.Event == movizor.CallToDriverEvent ||
			v.Event == movizor.NoConfirmationEvent ||
			v.Event == movizor.ObjectLimitedEvent ||
			v.Event == movizor.OnRouteEvent ||
			v.Event == movizor.ReturnRouteEvent ||
			v.Event == movizor.LeftRouteEvent ||
			v.Event == movizor.NotRouteEvent ||
			v.Event == movizor.OnParkingEvent ||
			v.Event == movizor.OffParkingEvent ||
			v.Event == movizor.MStopEvent ||
			v.Event == movizor.MStartEvent) {
			t.Fatalf("Something wen wrong with %T", movizor.ObjectEvents{})
		}
	}
}

func TestSubscribedEventsUnmarshal(t *testing.T) {
	d, err := ioutil.ReadFile(filepath.Join(dataPath, "events_subscribe_list.json"))
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	var se movizor.SubscribedEvents
	err = json.Unmarshal(d, &se)
	if err != nil {
		t.Fatalf("Input %s is not parsed to %T.\n\nError: %s", d, movizor.SubscribedEvents{}, err)
	}

	if v, err := se[0].Timestamp.Int64(); err != nil ||
		!(v == 1548084632 || se[0].Event != movizor.RejectEvent) {
		t.Fatalf("Input %s is not parsed to %T.\n\nError: %s", d, movizor.SubscribedEvents{}, err)
	}
}

func TestObjectInfo_UnmarshalJSON(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		file    string
		args    args
		wantErr bool
	}{
		{
			name:    "object_get1",
			file:    "object_get1.json",
			wantErr: false,
		},
		{
			name:    "object_get2",
			file:    "object_get2.json",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oi := &movizor.ObjectInfo{}
			var err error

			tt.args.data, err = ioutil.ReadFile(filepath.Join(dataPath, tt.file))
			if err != nil {
				t.Fatalf("err: %s", err)
			}
			if err := oi.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("ObjectInfo.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}

			fmt.Println(oi.Metadata)
		})
	}
}
