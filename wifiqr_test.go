package wifiqr_test

import (
	"hash/fnv"
	"reflect"
	"testing"

	"github.com/reugn/wifiqr"
)

func TestCode(t *testing.T) {
	config := wifiqr.NewConfig("ssid1", "1234", "WPA2", false)
	qrCode, err := wifiqr.InitCode(config)
	if err != nil {
		t.Fatal(err)
	}
	assertEqual(t, hashCode(qrCode.ToString(false)), 550955445)
}

func hashCode(s string) int {
	h := fnv.New32a()
	h.Write([]byte(s))
	return int(h.Sum32())
}

func assertEqual(t *testing.T, a interface{}, b interface{}) {
	if !reflect.DeepEqual(a, b) {
		t.Fatalf("%v != %v", a, b)
	}
}
