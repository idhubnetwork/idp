package crypto

import (
	"fmt"
	"testing"
)

type AuthTest struct {
	Id  string
	Msg string
	Sig string
}

var T = AuthTest{
	"0x49dBa8f906c745B0a82f4D21E02BAFD7Df1a0be4",
	"IDHub: 0x29F20242051AccDA50D52a7E272A5F23237e4696 @ 2018-10-25 14:46:44.384195779 +0800 CST m=+30.668203170",
	"0x6da7f9726b43c7b1158d93e14bfcf11caf134d3a78e256b80ee65c9531ae02c81d7222227a84d746484b93d8b492ab12f4ab78b790088ea01675f4750b59228c1b",
}

/*
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
*/
func TestSigPublicKey(t *testing.T) {
	key, err := SigPublicKey(T.Msg, T.Sig)
	if err != nil {
		fmt.Println(err)
		t.Error(err)
		t.Fail()
	}
	fmt.Println(key)
}

func TestEcrecover(t *testing.T) {
	addr, err := EcRecover(T.Msg, T.Sig)
	if err != nil {
		fmt.Println(err)
		t.Error(err)
		t.Fail()
	}
	fmt.Println(addr)
}

func TestVerifyAuth(t *testing.T) {
	ok, err := verifyAuth(T.Msg, T.Id, T.Sig)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	fmt.Println(ok)
}
