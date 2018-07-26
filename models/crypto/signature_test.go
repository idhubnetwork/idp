package crypto

import (
	"testing"
)

func TestEcRecover(t *testing.T) {
	type args struct {
		msg string
		sig string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			"sign from geth",
			args{
				"哈哈",
				"0xd0b825e7b1f9feab8086d2f2dc3bd93653c51c83efdacd2472c75affab88d9a32c365e3c5532a12339a50c683c163aaf2715a9acc832af1aba377956c0adf7251b",
			},
			"0x5B69df1b73d8217DC683B4D30014EA629C454C82",
			false,
		}, {
			"sign from metamask",
			args{
				"idhub 花瓣",
				"0x8c8d683301856180bc65e4749428d75d46ed7328b479009b93255923a452206178ef3632c1a7c11f784ba6e19f0ed99851473822b55884756f2072e4a6c2478f1b",
			},
			"0x73E0d4e2F01A33dF9Bd4a62Ecb11110db51E4975",
			false,
		}, {
			"empty string",
			args{
				"",
				"0x",
			},
			"0x0000000000000000000000000000000000000000",
			true,
		}, {
			"when panic",
			args{
				"",
				"",
			},
			"0x0000000000000000000000000000000000000000",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := EcRecover(tt.args.msg, tt.args.sig)
			if (err != nil) != tt.wantErr {
				t.Errorf("EcRecover() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("EcRecover() = %v, want %v", got, tt.want)
			}
		})
	}
}
