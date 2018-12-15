package tests

import (
	"oboz/movizor"
	"testing"
)

const (
	Project = "oboz"
	Token   = "p1lb9h8xy5qe2ruf"
)

func TestNewMovizorAPIWithEndpoint(t *testing.T) {
	api, err := movizor.NewMovizorAPIWithEndpoint("https://movizor.ru/api", Project, Token)
	if err != nil {
		t.Fatal(err)
	}

	if api.Endpoint != "https://movizor.ru/api" || api.Project != Project || api.Token != Token {
		t.Fatal("type API cannot be set")
	}
}

func TestNewMovizorAPI(t *testing.T) {
	api, err := movizor.NewMovizorAPI(Project, Token)
	if err != nil {
		t.Fatal(err)
	}

	if api.Endpoint != "https://movizor.ru/api" || api.Project != Project || api.Token != Token {
		t.Fatal("type API cannot be set")
	}
}

func TestMakeRequestSuccess(t *testing.T) {
	api, err := movizor.NewMovizorAPI(Project, Token)
	if err != nil {
		t.Fatal(err)
	}
	r, err := api.MakeRequest("pos_objects", nil)
	if err != nil || r.Result != "success" {
		t.Fatal(err)
	}
	//fmt.Println(string(r.Data))
}

func TestMakeRequestError(t *testing.T) {
	api, err := movizor.NewMovizorAPI("zobo", "xxx")
	if err != nil {
		t.Fatal(err)
	}
	r, err := api.MakeRequest("balance", nil)
	if err == nil || r.Result != "error" {
		t.Fatal(err)
	}
}

func TestGetBalance(t *testing.T) {
	api, err := movizor.NewMovizorAPI(Project, Token)
	if err != nil {
		t.Fatal(err)
	}

	b, err := api.GetBalance()
	if err != nil {
		t.Fatal(err)
	}

	if f, err := b.Balance.Float64(); err != nil || f == 0.0 {
		t.Fatal("balance action cannot be parsed")
	}
	// mts, megafon, beeline, tele2, eventsms, autoinform
	// fmt.Println(string(b.TariffPlans["eventsms"]))
}

func TestGetObjectPosition(t *testing.T) {
	api, err := movizor.NewMovizorAPI(Project, Token)
	if err != nil {
		t.Fatal(err)
	}

	op, err := api.GetObjectPositions()
	if err != nil {
		t.Fatal(err)
	}

	if v, err := op[0].Lat.Float64(); err != nil || v == 0.0 {
		t.Fatal("pos_objects action cannot be parsed")
	}
}
