package script

import (
	"github.com/Galaxy-Railway/GGargantua/internal/gargantua/domain/script/vm"
	"testing"
)

func TestScript(t *testing.T) {
	v := vm.NewVM()
	err := v.SetCacheGetter()
	if err != nil {
		t.Fatalf("err should be nil, got: %v", err)
	}
	err = v.SetCacheSetter()
	if err != nil {
		t.Fatalf("err should be nil, got: %v", err)
	}

	type JsTestStruct struct {
		Id     int           `json:"id"`
		Name   string        `json:"name"`
		Father *JsTestStruct `json:"father"`
	}
	testData := JsTestStruct{
		Id:   0,
		Name: "json",
		Father: &JsTestStruct{
			Id:     1,
			Name:   "javascript",
			Father: nil,
		},
	}

	err = v.SetResponse(testData)
	if err != nil {
		t.Fatalf("err should be nil, got: %v", err)
	}
	_, err = v.RunScript(`
id = response.father.id
setCache("id", "int64", id)
`)
	if err != nil {
		t.Fatalf("err should be nil, got: %v", err)
	}
	_, err = v.RunScript(`
cId = getCache("id")
answer = cId+1
`)
	if err != nil {
		t.Fatalf("err should be nil, got: %v", err)
	}
	value, err := v.Ot.Get("answer")
	if err != nil {
		t.Fatalf("err should be nil, got: %v", err)
	}
	answer, err := value.ToInteger()
	if err != nil {
		t.Fatalf("err should be nil, got: %v", err)
	}

	t.Logf("got value: %v", answer)
}
