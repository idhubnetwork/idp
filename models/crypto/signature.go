package crypto

import (
	"errors"
	"fmt"
	"idp/models/resolver"
	"log"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func signHash(data []byte) []byte {
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(data), data)
	log.Println(msg)
	return crypto.Keccak256([]byte(msg))
}

// EcRecover returns the address for the Account that was used to create the signature.
// func EcRecover(msg, sig string) (addr string, err error) {
func EcRecover(msg, sigHex string) (addr string, err error) {
	defer func() {
		if r := recover(); r != nil {
			addr = common.Address{}.String()
			err = r.(error)
		}
	}()

	data := []byte(msg)
	hash := signHash(data)

	sig := hexutil.MustDecode(sigHex)
	// https://github.com/ethereum/go-ethereum/blob/55599ee95d4151a2502465e0afc7c47bd1acba77/internal/ethapi/api.go#L442
	if sig[64] != 27 && sig[64] != 28 {
		return common.Address{}.String(), errors.New("nvalid Ethereum signature (V is not 27 or 28)")
	}
	sig[64] -= 27

	sigPiblicKey, err := crypto.SigToPub(hash, sig)

	if err != nil {
		fmt.Println("ECRECOVER" + err.Error())
		panic(err)
	}

	commAddr := crypto.PubkeyToAddress(*sigPiblicKey)

	return commAddr.String(), nil
}

func SigPublicKey(msg, sigHex string) (publickey string, err error) {
	data := []byte(msg)
	hash := signHash(data)

	sig := hexutil.MustDecode(sigHex)

	// https://github.com/ethereum/go-ethereum/blob/55599ee95d4151a2502465e0afc7c47bd1acba77/internal/ethapi/api.go#L442
	if sig[64] != 27 && sig[64] != 28 {
		return common.Address{}.String(), errors.New("nvalid Ethereum signature (V is not 27 or 28)")
	}
	sig[64] -= 27

	sigPiblicKey, err := crypto.SigToPub(hash, sig)

	if err != nil {
		fmt.Println("ECRECOVER" + err.Error())
		return "", err
	}
	publickeyBytes := crypto.FromECDSAPub(sigPiblicKey)
	publickey = hexutil.Encode(publickeyBytes)

	return publickey, nil
}

func VerifyAuth(msg, id, sig string) (info string, err error) {
	r, err := resolver.NewResolver("infuraRopsten", "0x1DbF8e4B47EA53a2b932850F7FEC8585C6A9c1d2")
	owner, err := r.IdentityOwner(id)

	if err != nil {
		log.Println(err)
		return "", err
	}

	publickey, err := SigPublicKey(msg, sig)
	if err != nil {
		log.Println(err)
		return "", err
	}

	ok, err := r.ValidAuthentication(id, "sigAuth", publickey)

	addr, err := EcRecover(msg, sig)
	if err != nil {
		log.Println(err)
		return "", err
	}

	if strings.ToLower(addr) == strings.ToLower(id) {
		return "sign by self", nil
	} else if strings.ToLower(owner) == strings.ToLower(addr) {
		return "sign by did owner", nil
	} else if ok {
		return "sign with authKey", nil
	}
	return "", errors.New("verify failed")
}
