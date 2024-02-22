package degenapi

import (
	"testing"

	"sourcecode.social/reiver/go-opt"
)

func TestUnmarshalAirDrop2Points(t *testing.T) {

	tests := []struct{
		JSON string
		Expected AirDrop2Points
	}{
		{
			JSON:
				`[`+
				`]`,
			Expected: AirDrop2Points{},
		},
		{
			JSON:
				`[`+
				`{"display_name":"Joe Blow","points":"25100"}`+
				`]`,
			Expected: AirDrop2Points{
				DisplayName: opt.Something("Joe Blow"),
				Points:      opt.Something("25100"),
			},
		},
	}

	for testNumber, test := range tests {

		var p []byte = []byte(test.JSON)

		var actual AirDrop2Points

		err := unmarshalAirDrop2Points(&actual, p)

		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("JSON:\n%s", test.JSON)
			continue
		}

		{
			expected := test.Expected

			if expected != actual {
				t.Errorf("For test #%d, the actual AirDrop2 points value is not what was expected.", testNumber)
				t.Logf("EXPECTED: %#v", expected)
				t.Logf("ACTUAL:   %#v", actual)
				t.Logf("JSON:\n%s", test.JSON)
				continue
			}
		}
	}
}
