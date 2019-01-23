package tests

import (
	"oboz/movizor"
	"testing"
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

	if err != nil || b.Balance == 0.0 {
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

	o, err := api.GetObjectInfo("+7 963 654 5272")
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

func TestGetObjectPositions(t *testing.T) {
	api, err := movizor.NewMovizorAPI(project, token)
	if err != nil {
		t.Fatal(err)
	}
	api.IsDebug = testLogging

	op, err := api.GetObjectPositions()
	if err != nil {
		t.Fatal(err)
	}

	if v, err := op[0].Lat.Float64(); err != nil || v == 0.0 {
		t.Fatal("pos_objects action cannot be parsed")
	}
}

//func TestGetObjectPositionsChan(t *testing.T) {
//	api, err := movizor.NewMovizorAPI(project, token)
//	if err != nil {
//		t.Fatal(err)
//	}
//	//api.IsDebug = testLogging
//	pos, err := api.GetObjectPositionsChan(time.Second * 15)
//	for p := range pos {
//		if p.Phone.String() == "" {
//			continue
//		}
//		fmt.Printf("Object [%s] at lon: %s lat: %s", p.Phone.String(), p.Lon.String(), p.Lat.String())
//	}
//}

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

func TestGetEvents(t *testing.T) {
	api, err := movizor.NewMovizorAPI(project, token)
	if err != nil {
		t.Fatal(err)
	}
	api.IsDebug = testLogging

	e, err := api.GetEvents(movizor.ObjectEventsOptions{})
	if err != nil {
		t.Fatal(err)
	}

	if err != nil || !(e[0].Timestamp != 0) {
		t.Fatal("events action cannot be parsed")
	}
}

func TestDeleteEventsSubscription(t *testing.T) {
	api, err := movizor.NewMovizorAPI(project, token)
	if err != nil {
		t.Fatal(err)
	}
	api.IsDebug = testLogging

	resp, err := api.DeleteEventsSubscription(5027)
	if err != nil {
		t.Fatalf("events_subscribe_delete returned error: %s", resp.ErrorText)
	}
}

func TestGetEventSubscriptions(t *testing.T) {
	api, err := movizor.NewMovizorAPI(project, token)
	if err != nil {
		t.Fatal(err)
	}
	api.IsDebug = testLogging

	e, err := api.GetEventSubscriptions()
	if err != nil {
		t.Fatal(err)
	}

	if err != nil || !(e[0].Timestamp != 0) {
		t.Fatal("events_subscribe_list action cannot be parsed")
	}
}
