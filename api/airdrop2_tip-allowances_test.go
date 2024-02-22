package degenapi

import (
	"testing"

	"reflect"

	"sourcecode.social/reiver/go-opt"
)

func TestUnmarshalAirDrop2TipAllowances(t *testing.T) {

	tests := []struct{
		JSON string
		Expected []AirDrop2DailyTipAllowance
	}{
		{
			JSON:
				`[`+
				`]`,
			Expected: []AirDrop2DailyTipAllowance{},
		},
		{
			JSON:
				`[`+
					`{"fid":"269231","user_rank":"25789","avatar_url":"https://i.imgur.com/JAyw5BZ.jpg","display_name":"Charles","reactions_per_cast":"1","tip_allowance":"326"}`+
				`]`,
			Expected: []AirDrop2DailyTipAllowance{
				AirDrop2DailyTipAllowance{
					FID:              opt.Something("269231"),
					UserRank:         opt.Something("25789"),
					AvatarURL:        opt.Something("https://i.imgur.com/JAyw5BZ.jpg"),
					DisplayName:      opt.Something("Charles"),
					ReactionsPerCast: opt.Something("1"),
					TipAllowance:     opt.Something("326"),
				},
			},
		},
		{
			JSON:
				`[`+
					`{"fid":"269231","user_rank":"25789","avatar_url":"https://i.imgur.com/JAyw5BZ.jpg","display_name":"Charles","reactions_per_cast":"1","tip_allowance":"326"}`+
					`,`+
					`{"fid":"1595","user_rank":"1724","avatar_url":"https://i.seadn.io/gcs/files/8a7daa3ddf457228d0dc62f3c85a6e73.gif?w=500&auto=format","display_name":"chet","reactions_per_cast":"3","tip_allowance":"4381"}`+
				`]`,
			Expected: []AirDrop2DailyTipAllowance{
				AirDrop2DailyTipAllowance{
					FID:              opt.Something("269231"),
					UserRank:         opt.Something("25789"),
					AvatarURL:        opt.Something("https://i.imgur.com/JAyw5BZ.jpg"),
					DisplayName:      opt.Something("Charles"),
					ReactionsPerCast: opt.Something("1"),
					TipAllowance:     opt.Something("326"),
				},
				AirDrop2DailyTipAllowance{
					FID:              opt.Something("1595"),
					UserRank:         opt.Something("1724"),
					AvatarURL:        opt.Something("https://i.seadn.io/gcs/files/8a7daa3ddf457228d0dc62f3c85a6e73.gif?w=500&auto=format"),
					DisplayName:      opt.Something("chet"),
					ReactionsPerCast: opt.Something("3"),
					TipAllowance:     opt.Something("4381"),
				},
			},
		},
	}

	for testNumber, test := range tests {

		var p []byte = []byte(test.JSON)

		var actual []AirDrop2DailyTipAllowance

		err := unmarshalAirDrop2TipAllowances(&actual, p)

		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("JSON:\n%s", test.JSON)
			continue
		}

		{
			expected := test.Expected

			if !reflect.DeepEqual(expected, actual) {
				t.Errorf("For test #%d, the actual AirDrop1 tip-allowances is not what was expected.", testNumber)
				t.Logf("EXPECTED: %#v", expected)
				t.Logf("ACTUAL:   %#v", actual)
				t.Logf("JSON:\n%s", test.JSON)
				continue
			}
		}
	}
}
