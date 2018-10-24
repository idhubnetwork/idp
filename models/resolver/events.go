package resolver

import (
	"context"
	"math/big"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
)

type LogPublicKeyChanged struct {
	Identity       common.Address
	PublicKeyType  [32]byte
	PublicKey      [32]byte
	ValidTo        *big.Int
	PreviousChange *big.Int
}

type LogAuthenticationChanged struct {
	Identity           common.Address
	AuthenticationType [32]byte
	Authentication     [32]byte
	ValidTo            *big.Int
	PreviousChange     *big.Int
}

type LogAttributeChanged struct {
	Identity       common.Address
	Name           []byte
	Value          []byte
	ValidTo        *big.Int
	PreviousChange *big.Int
}

func (r *resolver) EventPublicKeyChanged(blockNumber *big.Int) ([]LogPublicKeyChanged, error) {
	query := ethereum.FilterQuery{
		FromBlock: blockNumber,
		ToBlock:   blockNumber,
		Addresses: []common.Address{
			r.address,
		},
	}
	logs, err := r.client.FilterLogs(context.Background(), query)
	if err != nil {
		return nil, err
	}
	var logPublicKeyChangeds = make([]LogPublicKeyChanged, 0)
	for _, vLog := range logs {
		var logPublicKeyChanged LogPublicKeyChanged
		err := r.abi.Unpack(&logPublicKeyChanged, "DIDPublicKeyChanged", vLog.Data)
		if err != nil {
			return nil, err
		}
		logPublicKeyChanged.Identity = common.HexToAddress(vLog.Topics[1].Hex())
		logPublicKeyChangeds = append(logPublicKeyChangeds, logPublicKeyChanged)
	}
	return logPublicKeyChangeds, nil
}

func (r *resolver) EventAuthenticationChanged(blockNumber *big.Int) ([]LogAuthenticationChanged, error) {
	query := ethereum.FilterQuery{
		FromBlock: blockNumber,
		ToBlock:   blockNumber,
		Addresses: []common.Address{
			r.address,
		},
	}
	logs, err := r.client.FilterLogs(context.Background(), query)
	if err != nil {
		return nil, err
	}
	var logAuthenticationChangeds = make([]LogAuthenticationChanged, 0)
	for _, vLog := range logs {
		var logAuthenticationChanged LogAuthenticationChanged
		err := r.abi.Unpack(&logAuthenticationChanged, "DIDAuthenticationChanged", vLog.Data)
		if err != nil {
			return nil, err
		}
		logAuthenticationChanged.Identity = common.HexToAddress(vLog.Topics[1].Hex())
		logAuthenticationChangeds = append(logAuthenticationChangeds, logAuthenticationChanged)
	}
	return logAuthenticationChangeds, nil
}

func (r *resolver) EventAttributeChanged(blockNumber *big.Int) ([]LogAttributeChanged, error) {
	query := ethereum.FilterQuery{
		FromBlock: blockNumber,
		ToBlock:   blockNumber,
		Addresses: []common.Address{
			r.address,
		},
	}
	logs, err := r.client.FilterLogs(context.Background(), query)
	if err != nil {
		return nil, err
	}
	var logAttributeChangeds = make([]LogAttributeChanged, 0)
	for _, vLog := range logs {
		var logAttributeChanged LogAttributeChanged
		err := r.abi.Unpack(&logAttributeChanged, "DIDAttributeChanged", vLog.Data)
		if err != nil {
			return nil, err
		}
		logAttributeChanged.Identity = common.HexToAddress(vLog.Topics[1].Hex())
		logAttributeChangeds = append(logAttributeChangeds, logAttributeChanged)
	}
	return logAttributeChangeds, nil
}
