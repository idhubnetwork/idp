package resolver

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	did "resolver/contracts"
)

type resolver struct {
	client   *ethclient.Client
	contract *did.Did
	address  common.Address
	abi      abi.ABI
}

var urls map[string]string
var err error

func init() {
	urls = make(map[string]string, 5)
	urls["infuraMainnet"] = "https://mainnet.infura.io"
	urls["infuraRopsten"] = "https://ropsten.infura.io"
	urls["infuraRinkeby"] = "https://rinkeby.infura.io"
}

func NewResolver(net string, address string) (*resolver, error) {
	r := new(resolver)
	r.client, err = ethclient.Dial(urls[net])
	if err != nil {
		return nil, err
	}
	contractAddr := common.HexToAddress(address)
	r.address = contractAddr
	r.contract, err = did.NewDid(contractAddr, r.client)
	if err != nil {
		return nil, err
	}
	contractAbi, err := abi.JSON(strings.NewReader(string(did.DidABI)))
	if err != nil {
		return nil, err
	}
	r.abi = contractAbi
	return r, nil
}
