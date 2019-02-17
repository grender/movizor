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

	_, err = api.GetEvents(movizor.ObjectEventsOptions{})
	if err != nil {
		t.Fatal(err)
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

	_, err = api.GetEventSubscriptions()
	if err != nil {
		t.Fatal(err)
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
