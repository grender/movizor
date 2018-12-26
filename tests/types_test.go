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

	if err != nil || po[0].Phone != "79630005272" {
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
	d, err := ioutil.ReadFile(filepath.Join(dataPath, "object_get.json"))
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	var oi movizor.ObjectInfo
	err = json.Unmarshal(d, &oi)

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
