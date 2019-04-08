package tests

import (
	"encoding/json"
	"flatten_unflatten_test/flatten_unflatten/converting"
	"reflect"
	"testing"
)

func TestUnflatten(t *testing.T) {
	cases := []struct {
		test   string
		expect string
	}{
		{
			`{"1":"3","5.67.4":"ksjdghf","5.8":"jhsdgflksdmf"}`,
			`{"1":"3","5":{"67":{"4":"ksjdghf"},"8":"jhsdgflksdmf"}}`,
		},
		{
			`{"test.value":12.453}`,
			`{"test":{"value":12.453}}`,
		},
		{
			`{"test.0":"value","test.1":12.453}`,
			`{"test":{"0":"value","1":12.453}}`,
		},
	}

	for i, test := range cases {
		var m interface{}
		err := json.Unmarshal([]byte(test.test), &m)
		if err != nil {
			t.Errorf("%d: failed to unmarshal test: %v", i+1, err)
		}

		got, err := converting.Unflatten(m.(map[string]interface{}))
		if err != nil {
			t.Errorf("%d: failed to flatten: %v", i+1, err)
		}

		result, err := json.Marshal(&got)
		if err != nil {
			t.Errorf("failed marshaling: %v", err)
		}
		if !reflect.DeepEqual(string(result), test.expect) {
			t.Errorf("%d: mismatch, got: %v want: %v", i+1, string(result), test.expect)
		}
	}
}
