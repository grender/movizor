package tests

import (
	"fmt"
	"movizor"
	"testing"
	"time"
)

const (
	project = "oboz"
	token   = "p1lb9h8xy5qe2ruf"
)

const testLogging = true

func TestNewMovizorAPIWithEndpoint(t *testing.T) {
	api, err := movizor.NewMovizorAPIWithEndpoint("https://movizor.ru/api", project, token)
	if err != nil {
		t.Fatal(err)
	}

	if api.Endpoint != "https://movizor.ru/api" || api.Project != project || api.Token != token {
		t.Fatal("type API cannot be set")
	}
}

func TestNewMovizorAPI(t *testing.T) {
	api, err := movizor.NewMovizorAPI(project, token)
	if err != nil {
		t.Fatal(err)
	}

	if api.Endpoint != "https://movizor.ru/api" || api.Project != project || api.Token != token {
		t.Fatal("type API cannot be set")
	}
}

func TestMakeRequestSuccess(t *testing.T) {
	api, err := movizor.NewMovizorAPI(project, token)
	if err != nil {
		t.Fatal(err)
	}
	api.IsDebug = testLogging

	r, err := api.MakeRequest("pos_objects", nil)
	if err != nil || r.Result != "success" {
		t.Fatal(err)
	}
}

func TestMakeRequestError(t *testing.T) {
	api, err := movizor.NewMovizorAPI("zobo", "xxx")
	if err != nil {
		t.Fatal(err)
	}
	api.IsDebug = testLogging

	r, err := api.MakeRequest("balance", nil)
	if err == nil || r.Result != "error" {
		t.Fatal(err)
	}
}

func TestGetBalance(t *testing.T) {
	api, err := movizor.NewMovizorAPI(project, token)
	if err != nil {
		t.Fatal(err)
	}

	b, err := api.GetBalance()
	if err != nil {
		t.Fatal(err)
	}

	if b.Balance == 0.0 {
		t.Fatal("balance action cannot be parsed")
	}
	// mts, megafon, beeline, tele2, eventsms, autoinform
	// fmt.Println(string(b.TariffPlans["eventsms"]))
}

func TestAddObject(t *testing.T) {
	api, err := movizor.NewMovizorAPI(project, token)
	if err != nil {
		t.Fatal(err)
	}
	api.IsDebug = testLogging

	r, err := api.AddObject("+7 915 454-67-77",
		&movizor.ObjectOptions{
			Title:  "Объект 8",
			Tariff: movizor.TariffManual,
			Tags:   []string{"test1", "test2"},
			Metadata: map[string]string{
				"test4": "val4",
				"test5": "val5",
				"test6": "val6",
			},
			Destinations: []movizor.DestinationOptions{{
				Text:         "test",
				Lon:          0.001,
				Lat:          1.000,
				ExpectedTime: time.Now(),
			},
			},
		})
	if err != nil && r.Result != "error" {
		t.Fatal("object_add action cannot be parsed")
	}
}

func TestGetObjectInfo(t *testing.T) {
	api, err := movizor.NewMovizorAPI(project, token)
	if err != nil {
		t.Fatal(err)
	}
	api.IsDebug = testLogging

	o, err := api.GetObjectInfo("+7 915 454 6777")
	if err != nil {
		t.Fatal(err)
	}

	if err != nil || !(o.Status == movizor.StatusOff ||
		o.Status == movizor.StatusOk ||
		o.Status == movizor.StatusRejected ||
		o.Status == movizor.StatusNew ||
		o.Status == movizor.StatusWaitOk ||
		o.Status == movizor.StatusWaitOff) {
		t.Fatal("object_get action cannot be parsed")
	}
}

func TestEditObject(t *testing.T) {
	api, err := movizor.NewMovizorAPI(project, token)
	if err != nil {
		t.Fatal(err)
	}
	api.IsDebug = testLogging

	r, err := api.EditObject("+7 915 454-67-77",
		&movizor.ObjectOptions{
			Title:  "Объект 1",
			Tariff: movizor.TariffManual,
			//DateOff: "23.12.2018 16:03:00",
		})
	if err != nil && r.Result != "error" {
		t.Fatal("object_edit action cannot be parsed")
	}
}

func TestDeleteObject(t *testing.T) {
	api, err := movizor.NewMovizorAPI(project, token)
	if err != nil {
		t.Fatal(err)
	}
	api.IsDebug = testLogging

	r, err := api.DeleteObject("+7 968 062-15-56")
	if err != nil {
		t.Fatal(err)
	}

	if err != nil && r.Result != "error" {
		t.Fatal("object_delete action cannot be parsed")
	}
}

func TestReactivateObject(t *testing.T) {
	api, err := movizor.NewMovizorAPI(project, token)
	if err != nil {
		t.Fatal(err)
	}
	api.IsDebug = testLogging

	r, err := api.ReactivateObject("+7 963 654 5272")
	if err != nil {
		t.Fatal(err)
	}

	if err != nil && r.Result != "error" {
		t.Fatal("object_reactivate action cannot be parsed")
	}
}

func TestCancelTariffChangeObject(t *testing.T) {
	api, err := movizor.NewMovizorAPI(project, token)
	if err != nil {
		t.Fatal(err)
	}
	api.IsDebug = testLogging

	r, err := api.CancelTariffChangeObject("79063520695")
	if err != nil {
		t.Fatal(err)
	}

	if err != nil && r.Result != "error" {
		t.Fatal("object_cancel_tariff action cannot be parsed")
	}
}

func TestGetObjectLastPosition(t *testing.T) {
	api, err := movizor.NewMovizorAPI(project, token)
	if err != nil {
		t.Fatal(err)
	}
	api.IsDebug = testLogging

	_, err = api.GetObjectLastPosition("79154546777")
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetObjectPositions(t *testing.T) {
	api, err := movizor.NewMovizorAPI(project, token)
	if err != nil {
		t.Fatal(err)
	}
	api.IsDebug = testLogging

	_, err = api.GetObjectPositions("79154546777", nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetObjectsPositions(t *testing.T) {
	api, err := movizor.NewMovizorAPI(project, token)
	if err != nil {
		t.Fatal(err)
	}
	api.IsDebug = testLogging

	op, err := api.GetObjectsPositions()
	if err != nil {
		t.Fatal(err)
	}

	if v, err := op[0].Lat.Float64(); err != nil || v == 0.0 {
		t.Fatal("pos_objects action cannot be parsed")
	}
}

func TestRequestPositions(t *testing.T) {
	api, err := movizor.NewMovizorAPI(project, token)
	if err != nil {
		t.Fatal(err)
	}
	api.IsDebug = testLogging

	id, err := api.RequestPosition("79154546777")
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < 10; i++ {
		p, err := api.GetRequestedPosition(id)
		if err != nil {
			time.Sleep(time.Minute)
			continue
		}
		fmt.Println(p)
		break
	}
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetObjects(t *testing.T) {
	api, err := movizor.NewMovizorAPI(project, token)
	if err != nil {
		t.Fatal(err)
	}
	api.IsDebug = testLogging

	o, err := api.GetObjects()
	if err != nil {
		t.Fatal(err)
	}

	if err != nil || !(o[0].Status == movizor.StatusOk || o[0].Status == movizor.StatusOff) {
		t.Fatal("object_list action cannot be parsed")
	}
}

func TestGetOperatorInfo(t *testing.T) {
	api, err := movizor.NewMovizorAPI(project, token)
	if err != nil {
		t.Fatal(err)
	}
	api.IsDebug = testLogging

	o, err := api.GetOperatorInfo("79858393293")
	if err != nil {
		t.Fatal(err)
	}

	if err != nil || !(o.Operator == movizor.OperatorMTS ||
		o.Operator == movizor.OperatorMegafon ||
		o.Operator == movizor.OperatorBeeline ||
		o.Operator == movizor.OperatorTele2) {
		t.Fatal("get_operator action cannot be parsed")
	}
}
