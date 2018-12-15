package tests

import (
	"encoding/json"
	"io/ioutil"
	"oboz/movizor"
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

	if _, ok := bd.TariffPlans["mts"]; err != nil || !ok {
		t.Fatalf("Input %s is not parsed to %T.\n\nError: %s", d, movizor.Balance{}, err)
	}
}

func TestObjectPositionsUnmarshal(t *testing.T) {
	d, err := ioutil.ReadFile(filepath.Join(dataPath, "pos_objects.json"))
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	var po movizor.ObjectPositions
	err = json.Unmarshal(d, &po)

	if err != nil || po[0].Phone != "79630005272" {
		t.Fatalf("Input %s is not parsed to %T.\n\nError: %s", d, movizor.ObjectPositions{}, err)
	}
}
