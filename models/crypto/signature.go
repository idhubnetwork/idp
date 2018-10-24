package crypto

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
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

	data := []byte(msg)
	hash := crypto.Keccak256Hash(data)
	sigPiblicKey, err := crypto.Ecrecover(hash.Bytes(), hexutil.MustDecode(sig))
	pubkey, err := crypto.UnmarshalPubkey(sigPiblicKey)

	if err != nil {
		panic(err)
	}

	commAddr := crypto.PubkeyToAddress(*pubkey)

	return commAddr.String(), nil
}

func SigPublicKey(msg, sig string) (publickey string, err error) {
	data := []byte(msg)
	hash := crypto.Keccak256Hash(data)
	sigPiblicKey, err := crypto.Ecrecover(hash.Bytes(), hexutil.MustDecode(sig))
	if err != nil {
		return "", err
	}
	publickey = hexutil.Encode(sigPiblicKey)
	return publickey, nil
}
