package test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestWtf(t *testing.T) {
	type tmp struct {
		T time.Time `json:"time"`
	}
	v := tmp{T: time.Now()}
	str, err := json.Marshal(v)
	fmt.Println(string(str), err)
}
