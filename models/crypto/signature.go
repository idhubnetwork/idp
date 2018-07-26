package crypto

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/signer/core"
)

// EcRecover returns the address for the Account that was used to create the signature.
// func EcRecover(msg, sig string) (addr string, err error) {
func EcRecover(msg, sig string) (addr string, err error) {
	defer func() {
		if r := recover(); r != nil {
			addr = common.Address{}.String()
			err = r.(error)
		}
	}()

	personal := &core.SignerAPI{}

	commAddr, err := personal.EcRecover(nil, []byte(msg), hexutil.MustDecode(sig))

	if err != nil {
		panic(err)
	}

	return commAddr.String(), nil
}
