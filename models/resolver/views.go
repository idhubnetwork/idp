package resolver

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

func (r *resolver) GetPublicKeyChanged(address string) (*big.Int, error) {
	instance := r.contract
	identity := common.HexToAddress(address)
	publicKeyChanged, err := instance.PublicKeyChanged(nil, identity)
	if err != nil {
		return nil, err

	}
	return publicKeyChanged, nil
}

func (r *resolver) GetAuthenticationChanged(address string) (*big.Int, error) {
	instance := r.contract
	identity := common.HexToAddress(address)
	authenticationChanged, err := instance.AuthenticationChanged(nil, identity)
	if err != nil {
		return nil, err
	}
	return authenticationChanged, nil
}

func (r *resolver) GetAttributeChanged(address string) (*big.Int, error) {
	instance := r.contract
	identity := common.HexToAddress(address)
	attributeChanged, err := instance.AttributeChanged(nil, identity)
	if err != nil {
		return nil, err
	}
	return attributeChanged, nil
}

func (r *resolver) IdentityOwner(address string) (string, error) {
	instance := r.contract
	identity := common.HexToAddress(address)
	owner, err := instance.IdentityOwner(nil, identity)
	if err != nil {
		return "", err
	}
	return owner.Hex(), nil
}

func (r *resolver) ValidPublicKey(address, keyType, key string) (bool, error) {
	instance := r.contract
	identity := common.HexToAddress(address)
	publickKeyType := [32]byte{}
	copy(publickKeyType[:], keyType)
	publickKey := [32]byte{}
	hexKey, err := hexutil.Decode(key)
	if err != nil {
		return false, err
	}
	copy(publickKey[:], hexKey)
	ok, err := instance.ValidPublicKey(nil, identity, publickKeyType,
		publickKey)
	if err != nil {
		return false, err
	}
	return ok, nil
}

func (r *resolver) ValidAuthentication(address, keyType, key string) (bool, error) {
	instance := r.contract
	identity := common.HexToAddress(address)
	authenticationType := [32]byte{}
	copy(authenticationType[:], keyType)
	authentication := [32]byte{}
	hexKey, err := hexutil.Decode(key)
	if err != nil {
		return false, err
	}
	copy(authentication[:], hexKey)
	ok, err := instance.ValidAuthentication(nil, identity, authenticationType,
		authentication)
	if err != nil {
		return false, err
	}
	return ok, nil
}
