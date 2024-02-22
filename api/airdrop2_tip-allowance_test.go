package degenapi

import (
	"testing"

	"sourcecode.social/reiver/go-opt"
)

func TestUnmarshalAirDrop2TipAllowance(t *testing.T) {

	tests := []struct{
		JSON string
		Expected AirDrop2DailyTipAllowance
	}{
		{
			JSON: "[]",
			Expected: AirDrop2DailyTipAllowance{},
		},



		{
			JSON:
				`[`+
				`{`+
				`"snapshot_date":"2024-02-21T00:00:00.000Z"`+
				`}`+
				`]`,
			Expected: AirDrop2DailyTipAllowance{
				SnapshotDate: opt.Something("2024-02-21T00:00:00.000Z"),
			},
		},
		{
			JSON:
				`[`+
				`{`+
				`"snapshot_date":"2024-02-21T00:00:00.000Z"`+
				`,`+
				`"user_rank":"25789"`+
				`}`+
				`]`,
			Expected: AirDrop2DailyTipAllowance{
				SnapshotDate: opt.Something("2024-02-21T00:00:00.000Z"),
				UserRank: opt.Something("25789"),
			},
		},
		{
			JSON:
				`[`+
				`{`+
				`"snapshot_date":"2024-02-21T00:00:00.000Z"`+
				`,`+
				`"user_rank":"25789"`+
				`,`+
				`"wallet_address":"0xfecdab9876543210123456789abcdefedcba9876"`+
				`}`+
				`]`,
			Expected: AirDrop2DailyTipAllowance{
				SnapshotDate: opt.Something("2024-02-21T00:00:00.000Z"),
				UserRank: opt.Something("25789"),
				WalletAddress: opt.Something("0xfecdab9876543210123456789abcdefedcba9876"),
			},
		},
		{
			JSON:
				`[`+
				`{`+
				`"snapshot_date":"2024-02-21T00:00:00.000Z"`+
				`,`+
				`"user_rank":"25789"`+
				`,`+
				`"wallet_address":"0xfecdab9876543210123456789abcdefedcba9876"`+
				`,`+
				`"avatar_url":"https://example.com/bN604M3.jpeg"`+
				`}`+
				`]`,
			Expected: AirDrop2DailyTipAllowance{
				SnapshotDate: opt.Something("2024-02-21T00:00:00.000Z"),
				UserRank: opt.Something("25789"),
				WalletAddress: opt.Something("0xfecdab9876543210123456789abcdefedcba9876"),
				AvatarURL: opt.Something("https://example.com/bN604M3.jpeg"),
			},
		},
		{
			JSON:
				`[`+
				`{`+
				`"snapshot_date":"2024-02-21T00:00:00.000Z"`+
				`,`+
				`"user_rank":"25789"`+
				`,`+
				`"wallet_address":"0xfecdab9876543210123456789abcdefedcba9876"`+
				`,`+
				`"avatar_url":"https://example.com/bN604M3.jpeg"`+
				`,`+
				`"display_name":"Joe Blow"`+
				`}`+
				`]`,
			Expected: AirDrop2DailyTipAllowance{
				SnapshotDate: opt.Something("2024-02-21T00:00:00.000Z"),
				UserRank: opt.Something("25789"),
				WalletAddress: opt.Something("0xfecdab9876543210123456789abcdefedcba9876"),
				AvatarURL: opt.Something("https://example.com/bN604M3.jpeg"),
				DisplayName: opt.Something("Joe Blow"),
			},
		},
		{
			JSON:
				`[`+
				`{`+
				`"snapshot_date":"2024-02-21T00:00:00.000Z"`+
				`,`+
				`"user_rank":"25789"`+
				`,`+
				`"wallet_address":"0xfecdab9876543210123456789abcdefedcba9876"`+
				`,`+
				`"avatar_url":"https://example.com/bN604M3.jpeg"`+
				`,`+
				`"display_name":"Joe Blow"`+
				`,`+
				`"tip_allowance":"326"`+
				`}`+
				`]`,
			Expected: AirDrop2DailyTipAllowance{
				SnapshotDate: opt.Something("2024-02-21T00:00:00.000Z"),
				UserRank: opt.Something("25789"),
				WalletAddress: opt.Something("0xfecdab9876543210123456789abcdefedcba9876"),
				AvatarURL: opt.Something("https://example.com/bN604M3.jpeg"),
				DisplayName: opt.Something("Joe Blow"),
				TipAllowance: opt.Something("326"),
			},
		},
		{
			JSON:
				`[`+
				`{`+
				`"snapshot_date":"2024-02-21T00:00:00.000Z"`+
				`,`+
				`"user_rank":"25789"`+
				`,`+
				`"wallet_address":"0xfecdab9876543210123456789abcdefedcba9876"`+
				`,`+
				`"avatar_url":"https://example.com/bN604M3.jpeg"`+
				`,`+
				`"display_name":"Joe Blow"`+
				`,`+
				`"tip_allowance":"326"`+
				`,`+
				`"remaining_allowance":"26"`+
				`}`+
				`]`,
			Expected: AirDrop2DailyTipAllowance{
				SnapshotDate: opt.Something("2024-02-21T00:00:00.000Z"),
				UserRank: opt.Something("25789"),
				WalletAddress: opt.Something("0xfecdab9876543210123456789abcdefedcba9876"),
				AvatarURL: opt.Something("https://example.com/bN604M3.jpeg"),
				DisplayName: opt.Something("Joe Blow"),
				TipAllowance: opt.Something("326"),
				RemainingAllowance: opt.Something("26"),
			},
		},
		{
			JSON:
				`[`+
				`{`+
				`"fid":"269231"`+
				`,`+
				`"snapshot_date":"2024-02-21T00:00:00.000Z"`+
				`,`+
				`"user_rank":"25789"`+
				`,`+
				`"wallet_address":"0xfecdab9876543210123456789abcdefedcba9876"`+
				`,`+
				`"avatar_url":"https://example.com/bN604M3.jpeg"`+
				`,`+
				`"display_name":"Joe Blow"`+
				`,`+
				`"tip_allowance":"326"`+
				`,`+
				`"remaining_allowance":"26"`+
				`}`+
				`]`,
			Expected: AirDrop2DailyTipAllowance{
				FID: opt.Something("269231"),
				SnapshotDate: opt.Something("2024-02-21T00:00:00.000Z"),
				UserRank: opt.Something("25789"),
				WalletAddress: opt.Something("0xfecdab9876543210123456789abcdefedcba9876"),
				AvatarURL: opt.Something("https://example.com/bN604M3.jpeg"),
				DisplayName: opt.Something("Joe Blow"),
				TipAllowance: opt.Something("326"),
				RemainingAllowance: opt.Something("26"),
			},
		},
		{
			JSON:
				`[`+
				`{`+
				`"fid":"269231"`+
				`,`+
				`"snapshot_date":"2024-02-21T00:00:00.000Z"`+
				`,`+
				`"user_rank":"25789"`+
				`,`+
				`"wallet_address":"0xfecdab9876543210123456789abcdefedcba9876"`+
				`,`+
				`"avatar_url":"https://example.com/bN604M3.jpeg"`+
				`,`+
				`"display_name":"Joe Blow"`+
				`,`+
				`"reactions_per_cast":"1"`+
				`,`+
				`"tip_allowance":"326"`+
				`,`+
				`"remaining_allowance":"26"`+
				`}`+
				`]`,
			Expected: AirDrop2DailyTipAllowance{
				FID: opt.Something("269231"),
				SnapshotDate: opt.Something("2024-02-21T00:00:00.000Z"),
				UserRank: opt.Something("25789"),
				WalletAddress: opt.Something("0xfecdab9876543210123456789abcdefedcba9876"),
				AvatarURL: opt.Something("https://example.com/bN604M3.jpeg"),
				DisplayName: opt.Something("Joe Blow"),
				ReactionsPerCast: opt.Something("1"),
				TipAllowance: opt.Something("326"),
				RemainingAllowance: opt.Something("26"),
			},
		},
	}

	for testNumber, test := range tests {

		var p []byte = []byte(test.JSON)

		var actual AirDrop2DailyTipAllowance

		err := unmarshalAirDrop2TipAllowance(&actual, p)

		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one." , testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("JSON:\n%s", test.JSON)
			return
		}

		{
			expected := test.Expected

			if expected != actual {
				t.Errorf("For test #%d, the actual airdrop2 daily-tip-allowance is not what was expected.", testNumber)
				t.Logf("EXPECTED: %v", expected)
				t.Logf("ACTUAL:   %v", actual)
				t.Logf("JSON:\n%s", test.JSON)
				continue
			}
		}
	}
}
