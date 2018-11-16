package auth

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

func TestVerifyAuth(t *testing.T) {
	ok, err := verifyAuth(T.Msg, T.Id, T.Sig)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	fmt.Println(ok)
}
