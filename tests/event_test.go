package tests

import (
	"movizor"
	"testing"
)

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

	if v, err := e[0].Timestamp.Float64(); err != nil || v == 0.0 {
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

	if v, err := e[0].Timestamp.Float64(); err != nil || v == 0.0 {
		t.Fatal("events_subscribe_list action cannot be parsed")
	}
}

func TestSubscribeEvent(t *testing.T) {
	api, err := movizor.NewMovizorAPI(project, token)
	if err != nil {
		t.Fatal(err)
	}
	seo := movizor.SubscribeEventOptions{
		Objects: []movizor.Object{"+7 915 454-67-77"},
		Event:   movizor.RejectEvent,
	}
	err = seo.SetEMailNotification("nikolay.demidov@gmail.com")
	if err != nil {
		t.Fatal(err)
	}
	resp, err := api.SubscribeEvent(seo)
	if err != nil {
		t.Fatalf("events_subscribe_add returned error: %s", resp.ErrorText)
	}
}

func TestClearUnusedSubscriptions(t *testing.T) {
	api, err := movizor.NewMovizorAPI(project, token)
	if err != nil {
		t.Fatal(err)
	}
	api.IsDebug = testLogging

	err = api.ClearUnusedSubscriptions()
	if err != nil {
		t.Fatal(err)
	}
}
