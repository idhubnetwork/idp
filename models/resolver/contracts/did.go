// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package did

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// DidABI is the input ABI used to generate the binding from.
const DidABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"owners\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"authenticationChanged\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"},{\"name\":\"authenticationType\",\"type\":\"bytes32\"},{\"name\":\"authentication\",\"type\":\"bytes32\"},{\"name\":\"validity\",\"type\":\"uint256\"}],\"name\":\"addAuthentication\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"},{\"name\":\"sigV\",\"type\":\"uint8\"},{\"name\":\"sigR\",\"type\":\"bytes32\"},{\"name\":\"sigS\",\"type\":\"bytes32\"},{\"name\":\"authenticationType\",\"type\":\"bytes32\"},{\"name\":\"authentication\",\"type\":\"bytes32\"},{\"name\":\"validity\",\"type\":\"uint256\"}],\"name\":\"addAuthenticationSigned\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"publicKeyChanged\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"publicKeys\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"},{\"name\":\"sigV\",\"type\":\"uint8\"},{\"name\":\"sigR\",\"type\":\"bytes32\"},{\"name\":\"sigS\",\"type\":\"bytes32\"},{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"changeOwnerSigned\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"authentications\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"},{\"name\":\"sigV\",\"type\":\"uint8\"},{\"name\":\"sigR\",\"type\":\"bytes32\"},{\"name\":\"sigS\",\"type\":\"bytes32\"},{\"name\":\"publicKeyType\",\"type\":\"bytes32\"},{\"name\":\"publicKey\",\"type\":\"bytes32\"}],\"name\":\"revokePublicKeySigned\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"},{\"name\":\"authenticationType\",\"type\":\"bytes32\"},{\"name\":\"authentication\",\"type\":\"bytes32\"}],\"name\":\"revokeAuthentication\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"},{\"name\":\"sigV\",\"type\":\"uint8\"},{\"name\":\"sigR\",\"type\":\"bytes32\"},{\"name\":\"sigS\",\"type\":\"bytes32\"},{\"name\":\"publicKeyType\",\"type\":\"bytes32\"},{\"name\":\"publicKey\",\"type\":\"bytes32\"},{\"name\":\"validity\",\"type\":\"uint256\"}],\"name\":\"addPublicKeySigned\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"},{\"name\":\"publicKeyType\",\"type\":\"bytes32\"},{\"name\":\"publicKey\",\"type\":\"bytes32\"},{\"name\":\"validity\",\"type\":\"uint256\"}],\"name\":\"addPublicKey\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonce\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"},{\"name\":\"publicKeyType\",\"type\":\"bytes32\"},{\"name\":\"publicKey\",\"type\":\"bytes32\"}],\"name\":\"revokePublicKey\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"}],\"name\":\"identityOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"},{\"name\":\"name\",\"type\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"revokeAttribute\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"},{\"name\":\"sigV\",\"type\":\"uint8\"},{\"name\":\"sigR\",\"type\":\"bytes32\"},{\"name\":\"sigS\",\"type\":\"bytes32\"},{\"name\":\"name\",\"type\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes\"},{\"name\":\"validity\",\"type\":\"uint256\"}],\"name\":\"setAttributeSigned\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"},{\"name\":\"name\",\"type\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes\"},{\"name\":\"validity\",\"type\":\"uint256\"}],\"name\":\"setAttribute\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"attributeChanged\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"},{\"name\":\"sigV\",\"type\":\"uint8\"},{\"name\":\"sigR\",\"type\":\"bytes32\"},{\"name\":\"sigS\",\"type\":\"bytes32\"},{\"name\":\"name\",\"type\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"revokeAttributeSigned\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"},{\"name\":\"publicKeyType\",\"type\":\"bytes32\"},{\"name\":\"publicKey\",\"type\":\"bytes32\"}],\"name\":\"validPublicKey\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"},{\"name\":\"sigV\",\"type\":\"uint8\"},{\"name\":\"sigR\",\"type\":\"bytes32\"},{\"name\":\"sigS\",\"type\":\"bytes32\"},{\"name\":\"authenticationType\",\"type\":\"bytes32\"},{\"name\":\"authentication\",\"type\":\"bytes32\"}],\"name\":\"revokeAuthenticationSigned\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"},{\"name\":\"authenticationType\",\"type\":\"bytes32\"},{\"name\":\"authentication\",\"type\":\"bytes32\"}],\"name\":\"validAuthentication\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"identity\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"publicKeyType\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"publicKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"validTo\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"previousChange\",\"type\":\"uint256\"}],\"name\":\"DIDPublicKeyChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"identity\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"authenticationType\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"authentication\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"validTo\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"previousChange\",\"type\":\"uint256\"}],\"name\":\"DIDAuthenticationChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"identity\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"name\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"validTo\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"previousChange\",\"type\":\"uint256\"}],\"name\":\"DIDAttributeChanged\",\"type\":\"event\"}]"

// DidBin is the compiled bytecode used for deploying new contracts.
const DidBin = `608060405234801561001057600080fd5b50612f21806100206000396000f30060806040526004361061013e576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063022914a714610143578063048ed937146101c65780630824b3fa1461021d5780630e110989146102865780631a2bcf4c1461031857806320cacc321461036f578063240cf1fa146103e257806339fdc7281461046e5780633be52425146104e15780633ea1a65114610569578063456fcfeb146105c85780636866a6061461065a57806370ae92d2146106c3578063787768ef1461071a5780638733d4e8146107795780638dd83056146107fc578063c7b2864d146108cb578063ccbfa496146109cd578063d3c530d614610aa6578063dff0d6f414610afd578063e66bc6ba14610bf5578063f00d4b5d14610c6c578063f8156eb814610ccf578063ffb0254d14610d57575b600080fd5b34801561014f57600080fd5b50610184600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610dce565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156101d257600080fd5b50610207600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610e01565b6040518082815260200191505060405180910390f35b34801561022957600080fd5b50610284600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035600019169060200190929190803560001916906020019092919080359060200190929190505050610e19565b005b34801561029257600080fd5b50610316600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803560ff169060200190929190803560001916906020019092919080356000191690602001909291908035600019169060200190929190803560001916906020019092919080359060200190929190505050610e2c565b005b34801561032457600080fd5b50610359600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611064565b6040518082815260200191505060405180910390f35b34801561037b57600080fd5b506103cc600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035600019169060200190929190803560001916906020019092919050505061107c565b6040518082815260200191505060405180910390f35b3480156103ee57600080fd5b5061046c600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803560ff16906020019092919080356000191690602001909291908035600019169060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506110ae565b005b34801561047a57600080fd5b506104cb600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803560001916906020019092919080356000191690602001909291905050506112fd565b6040518082815260200191505060405180910390f35b3480156104ed57600080fd5b50610567600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803560ff169060200190929190803560001916906020019092919080356000191690602001909291908035600019169060200190929190803560001916906020019092919050505061132f565b005b34801561057557600080fd5b506105c6600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035600019169060200190929190803560001916906020019092919050505061155d565b005b3480156105d457600080fd5b50610658600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803560ff16906020019092919080356000191690602001909291908035600019169060200190929190803560001916906020019092919080356000191690602001909291908035906020019092919050505061156e565b005b34801561066657600080fd5b506106c1600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080356000191690602001909291908035600019169060200190929190803590602001909291905050506117a6565b005b3480156106cf57600080fd5b50610704600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506117b9565b6040518082815260200191505060405180910390f35b34801561072657600080fd5b50610777600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803560001916906020019092919080356000191690602001909291905050506117d1565b005b34801561078557600080fd5b506107ba600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506117e2565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561080857600080fd5b506108c9600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050611878565b005b3480156108d757600080fd5b506109cb600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803560ff16906020019092919080356000191690602001909291908035600019169060200190929190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001908201803590602001908080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050919291929080359060200190929190505050611889565b005b3480156109d957600080fd5b50610aa4600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001908201803590602001908080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050919291929080359060200190929190505050611b43565b005b348015610ab257600080fd5b50610ae7600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611b56565b6040518082815260200191505060405180910390f35b348015610b0957600080fd5b50610bf3600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803560ff16906020019092919080356000191690602001909291908035600019169060200190929190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050611b6e565b005b348015610c0157600080fd5b50610c52600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080356000191690602001909291908035600019169060200190929190505050611e1e565b604051808215151515815260200191505060405180910390f35b348015610c7857600080fd5b50610ccd600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611ebf565b005b348015610cdb57600080fd5b50610d55600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803560ff1690602001909291908035600019169060200190929190803560001916906020019092919080356000191690602001909291908035600019169060200190929190505050611ece565b005b348015610d6357600080fd5b50610db4600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803560001916906020019092919080356000191690602001909291905050506120fc565b604051808215151515815260200191505060405180910390f35b60006020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60046020528060005260406000206000915090505481565b610e26843385858561219d565b50505050565b600060197f01000000000000000000000000000000000000000000000000000000000000000260007f0100000000000000000000000000000000000000000000000000000000000000023060066000610e848d6117e2565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548b88888860405180897effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168152600101887effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191681526001018773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166c010000000000000000000000000281526014018681526020018573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166c01000000000000000000000000028152601401807f61646441757468656e7469636174696f6e00000000000000000000000000000081525060110184600019166000191681526020018360001916600019168152602001828152602001985050505050505050506040518091039020905061105a886110528a8a8a8a8761237a565b86868661219d565b5050505050505050565b60036020528060005260406000206000915090505481565b600160205282600052604060002060205281600052604060002060205280600052604060002060009250925050505481565b600060197f01000000000000000000000000000000000000000000000000000000000000000260007f01000000000000000000000000000000000000000000000000000000000000000230600660006111068b6117e2565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054898660405180877effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168152600101867effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191681526001018573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166c010000000000000000000000000281526014018481526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166c01000000000000000000000000028152601401807f6368616e67654f776e6572000000000000000000000000000000000000000000815250600b018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166c010000000000000000000000000281526014019650505050505050604051809103902090506112f5866112ef888888888761237a565b8461249e565b505050505050565b600260205282600052604060002060205281600052604060002060205280600052604060002060009250925050505481565b600060197f01000000000000000000000000000000000000000000000000000000000000000260007f01000000000000000000000000000000000000000000000000000000000000000230600660006113878c6117e2565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548a878760405180887effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168152600101877effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191681526001018673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166c010000000000000000000000000281526014018581526020018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166c01000000000000000000000000028152601401807f7265766f6b655075626c69634b65790000000000000000000000000000000000815250600f0183600019166000191681526020018260001916600019168152602001975050505050505050604051809103902090506115548761154d898989898761237a565b8585612566565b50505050505050565b6115698333848461273e565b505050565b600060197f01000000000000000000000000000000000000000000000000000000000000000260007f01000000000000000000000000000000000000000000000000000000000000000230600660006115c68d6117e2565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548b88888860405180897effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168152600101887effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191681526001018773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166c010000000000000000000000000281526014018681526020018573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166c01000000000000000000000000028152601401807f6164645075626c69634b65790000000000000000000000000000000000000000815250600c0184600019166000191681526020018360001916600019168152602001828152602001985050505050505050506040518091039020905061179c886117948a8a8a8a8761237a565b868686612916565b5050505050505050565b6117b38433858585612916565b50505050565b60066020528060005260406000206000915090505481565b6117dd83338484612566565b505050565b6000806000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905060008173ffffffffffffffffffffffffffffffffffffffff1614151561186e57809150611872565b8291505b50919050565b61188483338484612af3565b505050565b600060197f01000000000000000000000000000000000000000000000000000000000000000260007f01000000000000000000000000000000000000000000000000000000000000000230600660008c73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548b88888860405180897effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168152600101887effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191681526001018773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166c010000000000000000000000000281526014018681526020018573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166c01000000000000000000000000028152601401807f7365744174747269627574650000000000000000000000000000000000000000815250600c0184805190602001908083835b602083101515611a915780518252602082019150602081019050602083039250611a6c565b6001836020036101000a03801982511681845116808217855250505050505090500183805190602001908083835b602083101515611ae45780518252602082019150602081019050602083039250611abf565b6001836020036101000a0380198251168184511680821785525050505050509050018281526020019850505050505050505060405180910390209050611b3988611b318a8a8a8a8761237a565b868686612cf3565b5050505050505050565b611b508433858585612cf3565b50505050565b60056020528060005260406000206000915090505481565b600060197f01000000000000000000000000000000000000000000000000000000000000000260007f01000000000000000000000000000000000000000000000000000000000000000230600660008b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548a878760405180887effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168152600101877effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191681526001018673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166c010000000000000000000000000281526014018581526020018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166c01000000000000000000000000028152601401807f7265766f6b654174747269627574650000000000000000000000000000000000815250600f0183805190602001908083835b602083101515611d755780518252602082019150602081019050602083039250611d50565b6001836020036101000a03801982511681845116808217855250505050505090500182805190602001908083835b602083101515611dc85780518252602082019150602081019050602083039250611da3565b6001836020036101000a03801982511681845116808217855250505050505090500197505050505050505060405180910390209050611e1587611e0e898989898761237a565b8585612af3565b50505050505050565b600080600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600085604051808260001916600019168152602001915050604051809103902060001916600019168152602001908152602001600020600084600019166000191681526020019081526020016000205490504281119150509392505050565b611eca82338361249e565b5050565b600060197f01000000000000000000000000000000000000000000000000000000000000000260007f0100000000000000000000000000000000000000000000000000000000000000023060066000611f268c6117e2565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548a878760405180887effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168152600101877effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191681526001018673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166c010000000000000000000000000281526014018581526020018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166c01000000000000000000000000028152601401807f7265766f6b6541757468656e7469636174696f6e00000000000000000000000081525060140183600019166000191681526020018260001916600019168152602001975050505050505050604051809103902090506120f3876120ec898989898761237a565b858561273e565b50505050505050565b600080600260008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600085604051808260001916600019168152602001915050604051809103902060001916600019168152602001908152602001600020600084600019166000191681526020019081526020016000205490504281119150509392505050565b84846121a8826117e2565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415156121e157600080fd5b824201600260008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008760405180826000191660001916815260200191505060405180910390206000191660001916815260200190815260200160002060008660001916600019168152602001908152602001600020819055508673ffffffffffffffffffffffffffffffffffffffff167f6c1587d1ee8f4bdfb125ab048123336ebb2fb5f901a4561a8218e340ceffd2e08686864201600460008d73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054604051808560001916600019168152602001846000191660001916815260200183815260200182815260200194505050505060405180910390a243600460008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555050505050505050565b600080600183878787604051600081526020016040526040518085600019166000191681526020018460ff1660ff1681526020018360001916600019168152602001826000191660001916815260200194505050505060206040516020810390808403906000865af11580156123f4573d6000803e3d6000fd5b505050602060405103519050612409876117e2565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614151561244257600080fd5b600660008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600081548092919060010191905055508091505095945050505050565b82826124a9826117e2565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415156124e257600080fd5b826000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050505050565b8383612571826117e2565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415156125aa57600080fd5b42600160008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008660405180826000191660001916815260200191505060405180910390206000191660001916815260200190815260200160002060008560001916600019168152602001908152602001600020819055508573ffffffffffffffffffffffffffffffffffffffff167f213edff8255aa608dd36ca4cf7ed92d5480ddcd178865af1053f9d0bf5c77338858542600360008c73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054604051808560001916600019168152602001846000191660001916815260200183815260200182815260200194505050505060405180910390a243600360008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550505050505050565b8383612749826117e2565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614151561278257600080fd5b42600260008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008660405180826000191660001916815260200191505060405180910390206000191660001916815260200190815260200160002060008560001916600019168152602001908152602001600020819055508573ffffffffffffffffffffffffffffffffffffffff167f6c1587d1ee8f4bdfb125ab048123336ebb2fb5f901a4561a8218e340ceffd2e0858542600460008c73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054604051808560001916600019168152602001846000191660001916815260200183815260200182815260200194505050505060405180910390a243600460008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550505050505050565b8484612921826117e2565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614151561295a57600080fd5b824201600160008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008760405180826000191660001916815260200191505060405180910390206000191660001916815260200190815260200160002060008660001916600019168152602001908152602001600020819055508673ffffffffffffffffffffffffffffffffffffffff167f213edff8255aa608dd36ca4cf7ed92d5480ddcd178865af1053f9d0bf5c773388686864201600360008d73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054604051808560001916600019168152602001846000191660001916815260200183815260200182815260200194505050505060405180910390a243600360008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555050505050505050565b8383612afe826117e2565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141515612b3757600080fd5b8573ffffffffffffffffffffffffffffffffffffffff167f011b18dd995a3172c6dbe3b65ce383beec725369bb7cc6c16721013f9f993a7885856000600560008c73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054604051808060200180602001858152602001848152602001838103835287818151815260200191508051906020019080838360005b83811015612c03578082015181840152602081019050612be8565b50505050905090810190601f168015612c305780820380516001836020036101000a031916815260200191505b50838103825286818151815260200191508051906020019080838360005b83811015612c69578082015181840152602081019050612c4e565b50505050905090810190601f168015612c965780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390a243600560008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550505050505050565b8484612cfe826117e2565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141515612d3757600080fd5b8673ffffffffffffffffffffffffffffffffffffffff167f011b18dd995a3172c6dbe3b65ce383beec725369bb7cc6c16721013f9f993a788686864201600560008d73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054604051808060200180602001858152602001848152602001838103835287818151815260200191508051906020019080838360005b83811015612e04578082015181840152602081019050612de9565b50505050905090810190601f168015612e315780820380516001836020036101000a031916815260200191505b50838103825286818151815260200191508051906020019080838360005b83811015612e6a578082015181840152602081019050612e4f565b50505050905090810190601f168015612e975780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390a243600560008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550505050505050505600a165627a7a72305820b757b21c6e292d53669fb492e1b1f71b045393eeddc98a13c238a5202c2fd7b90029`

// DeployDid deploys a new Ethereum contract, binding an instance of Did to it.
func DeployDid(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Did, error) {
	parsed, err := abi.JSON(strings.NewReader(DidABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DidBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Did{DidCaller: DidCaller{contract: contract}, DidTransactor: DidTransactor{contract: contract}, DidFilterer: DidFilterer{contract: contract}}, nil
}

// Did is an auto generated Go binding around an Ethereum contract.
type Did struct {
	DidCaller     // Read-only binding to the contract
	DidTransactor // Write-only binding to the contract
	DidFilterer   // Log filterer for contract events
}

// DidCaller is an auto generated read-only Go binding around an Ethereum contract.
type DidCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DidTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DidTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DidFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DidFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DidSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DidSession struct {
	Contract     *Did              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DidCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DidCallerSession struct {
	Contract *DidCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// DidTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DidTransactorSession struct {
	Contract     *DidTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DidRaw is an auto generated low-level Go binding around an Ethereum contract.
type DidRaw struct {
	Contract *Did // Generic contract binding to access the raw methods on
}

// DidCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DidCallerRaw struct {
	Contract *DidCaller // Generic read-only contract binding to access the raw methods on
}

// DidTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DidTransactorRaw struct {
	Contract *DidTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDid creates a new instance of Did, bound to a specific deployed contract.
func NewDid(address common.Address, backend bind.ContractBackend) (*Did, error) {
	contract, err := bindDid(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Did{DidCaller: DidCaller{contract: contract}, DidTransactor: DidTransactor{contract: contract}, DidFilterer: DidFilterer{contract: contract}}, nil
}

// NewDidCaller creates a new read-only instance of Did, bound to a specific deployed contract.
func NewDidCaller(address common.Address, caller bind.ContractCaller) (*DidCaller, error) {
	contract, err := bindDid(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DidCaller{contract: contract}, nil
}

// NewDidTransactor creates a new write-only instance of Did, bound to a specific deployed contract.
func NewDidTransactor(address common.Address, transactor bind.ContractTransactor) (*DidTransactor, error) {
	contract, err := bindDid(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DidTransactor{contract: contract}, nil
}

// NewDidFilterer creates a new log filterer instance of Did, bound to a specific deployed contract.
func NewDidFilterer(address common.Address, filterer bind.ContractFilterer) (*DidFilterer, error) {
	contract, err := bindDid(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DidFilterer{contract: contract}, nil
}

// bindDid binds a generic wrapper to an already deployed contract.
func bindDid(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DidABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Did *DidRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Did.Contract.DidCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Did *DidRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Did.Contract.DidTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Did *DidRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Did.Contract.DidTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Did *DidCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Did.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Did *DidTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Did.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Did *DidTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Did.Contract.contract.Transact(opts, method, params...)
}

// AttributeChanged is a free data retrieval call binding the contract method 0xd3c530d6.
//
// Solidity: function attributeChanged( address) constant returns(uint256)
func (_Did *DidCaller) AttributeChanged(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Did.contract.Call(opts, out, "attributeChanged", arg0)
	return *ret0, err
}

// AttributeChanged is a free data retrieval call binding the contract method 0xd3c530d6.
//
// Solidity: function attributeChanged( address) constant returns(uint256)
func (_Did *DidSession) AttributeChanged(arg0 common.Address) (*big.Int, error) {
	return _Did.Contract.AttributeChanged(&_Did.CallOpts, arg0)
}

// AttributeChanged is a free data retrieval call binding the contract method 0xd3c530d6.
//
// Solidity: function attributeChanged( address) constant returns(uint256)
func (_Did *DidCallerSession) AttributeChanged(arg0 common.Address) (*big.Int, error) {
	return _Did.Contract.AttributeChanged(&_Did.CallOpts, arg0)
}

// AuthenticationChanged is a free data retrieval call binding the contract method 0x048ed937.
//
// Solidity: function authenticationChanged( address) constant returns(uint256)
func (_Did *DidCaller) AuthenticationChanged(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Did.contract.Call(opts, out, "authenticationChanged", arg0)
	return *ret0, err
}

// AuthenticationChanged is a free data retrieval call binding the contract method 0x048ed937.
//
// Solidity: function authenticationChanged( address) constant returns(uint256)
func (_Did *DidSession) AuthenticationChanged(arg0 common.Address) (*big.Int, error) {
	return _Did.Contract.AuthenticationChanged(&_Did.CallOpts, arg0)
}

// AuthenticationChanged is a free data retrieval call binding the contract method 0x048ed937.
//
// Solidity: function authenticationChanged( address) constant returns(uint256)
func (_Did *DidCallerSession) AuthenticationChanged(arg0 common.Address) (*big.Int, error) {
	return _Did.Contract.AuthenticationChanged(&_Did.CallOpts, arg0)
}

// Authentications is a free data retrieval call binding the contract method 0x39fdc728.
//
// Solidity: function authentications( address,  bytes32,  bytes32) constant returns(uint256)
func (_Did *DidCaller) Authentications(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte, arg2 [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Did.contract.Call(opts, out, "authentications", arg0, arg1, arg2)
	return *ret0, err
}

// Authentications is a free data retrieval call binding the contract method 0x39fdc728.
//
// Solidity: function authentications( address,  bytes32,  bytes32) constant returns(uint256)
func (_Did *DidSession) Authentications(arg0 common.Address, arg1 [32]byte, arg2 [32]byte) (*big.Int, error) {
	return _Did.Contract.Authentications(&_Did.CallOpts, arg0, arg1, arg2)
}

// Authentications is a free data retrieval call binding the contract method 0x39fdc728.
//
// Solidity: function authentications( address,  bytes32,  bytes32) constant returns(uint256)
func (_Did *DidCallerSession) Authentications(arg0 common.Address, arg1 [32]byte, arg2 [32]byte) (*big.Int, error) {
	return _Did.Contract.Authentications(&_Did.CallOpts, arg0, arg1, arg2)
}

// IdentityOwner is a free data retrieval call binding the contract method 0x8733d4e8.
//
// Solidity: function identityOwner(identity address) constant returns(address)
func (_Did *DidCaller) IdentityOwner(opts *bind.CallOpts, identity common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Did.contract.Call(opts, out, "identityOwner", identity)
	return *ret0, err
}

// IdentityOwner is a free data retrieval call binding the contract method 0x8733d4e8.
//
// Solidity: function identityOwner(identity address) constant returns(address)
func (_Did *DidSession) IdentityOwner(identity common.Address) (common.Address, error) {
	return _Did.Contract.IdentityOwner(&_Did.CallOpts, identity)
}

// IdentityOwner is a free data retrieval call binding the contract method 0x8733d4e8.
//
// Solidity: function identityOwner(identity address) constant returns(address)
func (_Did *DidCallerSession) IdentityOwner(identity common.Address) (common.Address, error) {
	return _Did.Contract.IdentityOwner(&_Did.CallOpts, identity)
}

// Nonce is a free data retrieval call binding the contract method 0x70ae92d2.
//
// Solidity: function nonce( address) constant returns(uint256)
func (_Did *DidCaller) Nonce(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Did.contract.Call(opts, out, "nonce", arg0)
	return *ret0, err
}

// Nonce is a free data retrieval call binding the contract method 0x70ae92d2.
//
// Solidity: function nonce( address) constant returns(uint256)
func (_Did *DidSession) Nonce(arg0 common.Address) (*big.Int, error) {
	return _Did.Contract.Nonce(&_Did.CallOpts, arg0)
}

// Nonce is a free data retrieval call binding the contract method 0x70ae92d2.
//
// Solidity: function nonce( address) constant returns(uint256)
func (_Did *DidCallerSession) Nonce(arg0 common.Address) (*big.Int, error) {
	return _Did.Contract.Nonce(&_Did.CallOpts, arg0)
}

// Owners is a free data retrieval call binding the contract method 0x022914a7.
//
// Solidity: function owners( address) constant returns(address)
func (_Did *DidCaller) Owners(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Did.contract.Call(opts, out, "owners", arg0)
	return *ret0, err
}

// Owners is a free data retrieval call binding the contract method 0x022914a7.
//
// Solidity: function owners( address) constant returns(address)
func (_Did *DidSession) Owners(arg0 common.Address) (common.Address, error) {
	return _Did.Contract.Owners(&_Did.CallOpts, arg0)
}

// Owners is a free data retrieval call binding the contract method 0x022914a7.
//
// Solidity: function owners( address) constant returns(address)
func (_Did *DidCallerSession) Owners(arg0 common.Address) (common.Address, error) {
	return _Did.Contract.Owners(&_Did.CallOpts, arg0)
}

// PublicKeyChanged is a free data retrieval call binding the contract method 0x1a2bcf4c.
//
// Solidity: function publicKeyChanged( address) constant returns(uint256)
func (_Did *DidCaller) PublicKeyChanged(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Did.contract.Call(opts, out, "publicKeyChanged", arg0)
	return *ret0, err
}

// PublicKeyChanged is a free data retrieval call binding the contract method 0x1a2bcf4c.
//
// Solidity: function publicKeyChanged( address) constant returns(uint256)
func (_Did *DidSession) PublicKeyChanged(arg0 common.Address) (*big.Int, error) {
	return _Did.Contract.PublicKeyChanged(&_Did.CallOpts, arg0)
}

// PublicKeyChanged is a free data retrieval call binding the contract method 0x1a2bcf4c.
//
// Solidity: function publicKeyChanged( address) constant returns(uint256)
func (_Did *DidCallerSession) PublicKeyChanged(arg0 common.Address) (*big.Int, error) {
	return _Did.Contract.PublicKeyChanged(&_Did.CallOpts, arg0)
}

// PublicKeys is a free data retrieval call binding the contract method 0x20cacc32.
//
// Solidity: function publicKeys( address,  bytes32,  bytes32) constant returns(uint256)
func (_Did *DidCaller) PublicKeys(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte, arg2 [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Did.contract.Call(opts, out, "publicKeys", arg0, arg1, arg2)
	return *ret0, err
}

// PublicKeys is a free data retrieval call binding the contract method 0x20cacc32.
//
// Solidity: function publicKeys( address,  bytes32,  bytes32) constant returns(uint256)
func (_Did *DidSession) PublicKeys(arg0 common.Address, arg1 [32]byte, arg2 [32]byte) (*big.Int, error) {
	return _Did.Contract.PublicKeys(&_Did.CallOpts, arg0, arg1, arg2)
}

// PublicKeys is a free data retrieval call binding the contract method 0x20cacc32.
//
// Solidity: function publicKeys( address,  bytes32,  bytes32) constant returns(uint256)
func (_Did *DidCallerSession) PublicKeys(arg0 common.Address, arg1 [32]byte, arg2 [32]byte) (*big.Int, error) {
	return _Did.Contract.PublicKeys(&_Did.CallOpts, arg0, arg1, arg2)
}

// ValidAuthentication is a free data retrieval call binding the contract method 0xffb0254d.
//
// Solidity: function validAuthentication(identity address, authenticationType bytes32, authentication bytes32) constant returns(bool)
func (_Did *DidCaller) ValidAuthentication(opts *bind.CallOpts, identity common.Address, authenticationType [32]byte, authentication [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Did.contract.Call(opts, out, "validAuthentication", identity, authenticationType, authentication)
	return *ret0, err
}

// ValidAuthentication is a free data retrieval call binding the contract method 0xffb0254d.
//
// Solidity: function validAuthentication(identity address, authenticationType bytes32, authentication bytes32) constant returns(bool)
func (_Did *DidSession) ValidAuthentication(identity common.Address, authenticationType [32]byte, authentication [32]byte) (bool, error) {
	return _Did.Contract.ValidAuthentication(&_Did.CallOpts, identity, authenticationType, authentication)
}

// ValidAuthentication is a free data retrieval call binding the contract method 0xffb0254d.
//
// Solidity: function validAuthentication(identity address, authenticationType bytes32, authentication bytes32) constant returns(bool)
func (_Did *DidCallerSession) ValidAuthentication(identity common.Address, authenticationType [32]byte, authentication [32]byte) (bool, error) {
	return _Did.Contract.ValidAuthentication(&_Did.CallOpts, identity, authenticationType, authentication)
}

// ValidPublicKey is a free data retrieval call binding the contract method 0xe66bc6ba.
//
// Solidity: function validPublicKey(identity address, publicKeyType bytes32, publicKey bytes32) constant returns(bool)
func (_Did *DidCaller) ValidPublicKey(opts *bind.CallOpts, identity common.Address, publicKeyType [32]byte, publicKey [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Did.contract.Call(opts, out, "validPublicKey", identity, publicKeyType, publicKey)
	return *ret0, err
}

// ValidPublicKey is a free data retrieval call binding the contract method 0xe66bc6ba.
//
// Solidity: function validPublicKey(identity address, publicKeyType bytes32, publicKey bytes32) constant returns(bool)
func (_Did *DidSession) ValidPublicKey(identity common.Address, publicKeyType [32]byte, publicKey [32]byte) (bool, error) {
	return _Did.Contract.ValidPublicKey(&_Did.CallOpts, identity, publicKeyType, publicKey)
}

// ValidPublicKey is a free data retrieval call binding the contract method 0xe66bc6ba.
//
// Solidity: function validPublicKey(identity address, publicKeyType bytes32, publicKey bytes32) constant returns(bool)
func (_Did *DidCallerSession) ValidPublicKey(identity common.Address, publicKeyType [32]byte, publicKey [32]byte) (bool, error) {
	return _Did.Contract.ValidPublicKey(&_Did.CallOpts, identity, publicKeyType, publicKey)
}

// AddAuthentication is a paid mutator transaction binding the contract method 0x0824b3fa.
//
// Solidity: function addAuthentication(identity address, authenticationType bytes32, authentication bytes32, validity uint256) returns()
func (_Did *DidTransactor) AddAuthentication(opts *bind.TransactOpts, identity common.Address, authenticationType [32]byte, authentication [32]byte, validity *big.Int) (*types.Transaction, error) {
	return _Did.contract.Transact(opts, "addAuthentication", identity, authenticationType, authentication, validity)
}

// AddAuthentication is a paid mutator transaction binding the contract method 0x0824b3fa.
//
// Solidity: function addAuthentication(identity address, authenticationType bytes32, authentication bytes32, validity uint256) returns()
func (_Did *DidSession) AddAuthentication(identity common.Address, authenticationType [32]byte, authentication [32]byte, validity *big.Int) (*types.Transaction, error) {
	return _Did.Contract.AddAuthentication(&_Did.TransactOpts, identity, authenticationType, authentication, validity)
}

// AddAuthentication is a paid mutator transaction binding the contract method 0x0824b3fa.
//
// Solidity: function addAuthentication(identity address, authenticationType bytes32, authentication bytes32, validity uint256) returns()
func (_Did *DidTransactorSession) AddAuthentication(identity common.Address, authenticationType [32]byte, authentication [32]byte, validity *big.Int) (*types.Transaction, error) {
	return _Did.Contract.AddAuthentication(&_Did.TransactOpts, identity, authenticationType, authentication, validity)
}

// AddAuthenticationSigned is a paid mutator transaction binding the contract method 0x0e110989.
//
// Solidity: function addAuthenticationSigned(identity address, sigV uint8, sigR bytes32, sigS bytes32, authenticationType bytes32, authentication bytes32, validity uint256) returns()
func (_Did *DidTransactor) AddAuthenticationSigned(opts *bind.TransactOpts, identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, authenticationType [32]byte, authentication [32]byte, validity *big.Int) (*types.Transaction, error) {
	return _Did.contract.Transact(opts, "addAuthenticationSigned", identity, sigV, sigR, sigS, authenticationType, authentication, validity)
}

// AddAuthenticationSigned is a paid mutator transaction binding the contract method 0x0e110989.
//
// Solidity: function addAuthenticationSigned(identity address, sigV uint8, sigR bytes32, sigS bytes32, authenticationType bytes32, authentication bytes32, validity uint256) returns()
func (_Did *DidSession) AddAuthenticationSigned(identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, authenticationType [32]byte, authentication [32]byte, validity *big.Int) (*types.Transaction, error) {
	return _Did.Contract.AddAuthenticationSigned(&_Did.TransactOpts, identity, sigV, sigR, sigS, authenticationType, authentication, validity)
}

// AddAuthenticationSigned is a paid mutator transaction binding the contract method 0x0e110989.
//
// Solidity: function addAuthenticationSigned(identity address, sigV uint8, sigR bytes32, sigS bytes32, authenticationType bytes32, authentication bytes32, validity uint256) returns()
func (_Did *DidTransactorSession) AddAuthenticationSigned(identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, authenticationType [32]byte, authentication [32]byte, validity *big.Int) (*types.Transaction, error) {
	return _Did.Contract.AddAuthenticationSigned(&_Did.TransactOpts, identity, sigV, sigR, sigS, authenticationType, authentication, validity)
}

// AddPublicKey is a paid mutator transaction binding the contract method 0x6866a606.
//
// Solidity: function addPublicKey(identity address, publicKeyType bytes32, publicKey bytes32, validity uint256) returns()
func (_Did *DidTransactor) AddPublicKey(opts *bind.TransactOpts, identity common.Address, publicKeyType [32]byte, publicKey [32]byte, validity *big.Int) (*types.Transaction, error) {
	return _Did.contract.Transact(opts, "addPublicKey", identity, publicKeyType, publicKey, validity)
}

// AddPublicKey is a paid mutator transaction binding the contract method 0x6866a606.
//
// Solidity: function addPublicKey(identity address, publicKeyType bytes32, publicKey bytes32, validity uint256) returns()
func (_Did *DidSession) AddPublicKey(identity common.Address, publicKeyType [32]byte, publicKey [32]byte, validity *big.Int) (*types.Transaction, error) {
	return _Did.Contract.AddPublicKey(&_Did.TransactOpts, identity, publicKeyType, publicKey, validity)
}

// AddPublicKey is a paid mutator transaction binding the contract method 0x6866a606.
//
// Solidity: function addPublicKey(identity address, publicKeyType bytes32, publicKey bytes32, validity uint256) returns()
func (_Did *DidTransactorSession) AddPublicKey(identity common.Address, publicKeyType [32]byte, publicKey [32]byte, validity *big.Int) (*types.Transaction, error) {
	return _Did.Contract.AddPublicKey(&_Did.TransactOpts, identity, publicKeyType, publicKey, validity)
}

// AddPublicKeySigned is a paid mutator transaction binding the contract method 0x456fcfeb.
//
// Solidity: function addPublicKeySigned(identity address, sigV uint8, sigR bytes32, sigS bytes32, publicKeyType bytes32, publicKey bytes32, validity uint256) returns()
func (_Did *DidTransactor) AddPublicKeySigned(opts *bind.TransactOpts, identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, publicKeyType [32]byte, publicKey [32]byte, validity *big.Int) (*types.Transaction, error) {
	return _Did.contract.Transact(opts, "addPublicKeySigned", identity, sigV, sigR, sigS, publicKeyType, publicKey, validity)
}

// AddPublicKeySigned is a paid mutator transaction binding the contract method 0x456fcfeb.
//
// Solidity: function addPublicKeySigned(identity address, sigV uint8, sigR bytes32, sigS bytes32, publicKeyType bytes32, publicKey bytes32, validity uint256) returns()
func (_Did *DidSession) AddPublicKeySigned(identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, publicKeyType [32]byte, publicKey [32]byte, validity *big.Int) (*types.Transaction, error) {
	return _Did.Contract.AddPublicKeySigned(&_Did.TransactOpts, identity, sigV, sigR, sigS, publicKeyType, publicKey, validity)
}

// AddPublicKeySigned is a paid mutator transaction binding the contract method 0x456fcfeb.
//
// Solidity: function addPublicKeySigned(identity address, sigV uint8, sigR bytes32, sigS bytes32, publicKeyType bytes32, publicKey bytes32, validity uint256) returns()
func (_Did *DidTransactorSession) AddPublicKeySigned(identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, publicKeyType [32]byte, publicKey [32]byte, validity *big.Int) (*types.Transaction, error) {
	return _Did.Contract.AddPublicKeySigned(&_Did.TransactOpts, identity, sigV, sigR, sigS, publicKeyType, publicKey, validity)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xf00d4b5d.
//
// Solidity: function changeOwner(identity address, newOwner address) returns()
func (_Did *DidTransactor) ChangeOwner(opts *bind.TransactOpts, identity common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _Did.contract.Transact(opts, "changeOwner", identity, newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xf00d4b5d.
//
// Solidity: function changeOwner(identity address, newOwner address) returns()
func (_Did *DidSession) ChangeOwner(identity common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _Did.Contract.ChangeOwner(&_Did.TransactOpts, identity, newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xf00d4b5d.
//
// Solidity: function changeOwner(identity address, newOwner address) returns()
func (_Did *DidTransactorSession) ChangeOwner(identity common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _Did.Contract.ChangeOwner(&_Did.TransactOpts, identity, newOwner)
}

// ChangeOwnerSigned is a paid mutator transaction binding the contract method 0x240cf1fa.
//
// Solidity: function changeOwnerSigned(identity address, sigV uint8, sigR bytes32, sigS bytes32, newOwner address) returns()
func (_Did *DidTransactor) ChangeOwnerSigned(opts *bind.TransactOpts, identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, newOwner common.Address) (*types.Transaction, error) {
	return _Did.contract.Transact(opts, "changeOwnerSigned", identity, sigV, sigR, sigS, newOwner)
}

// ChangeOwnerSigned is a paid mutator transaction binding the contract method 0x240cf1fa.
//
// Solidity: function changeOwnerSigned(identity address, sigV uint8, sigR bytes32, sigS bytes32, newOwner address) returns()
func (_Did *DidSession) ChangeOwnerSigned(identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, newOwner common.Address) (*types.Transaction, error) {
	return _Did.Contract.ChangeOwnerSigned(&_Did.TransactOpts, identity, sigV, sigR, sigS, newOwner)
}

// ChangeOwnerSigned is a paid mutator transaction binding the contract method 0x240cf1fa.
//
// Solidity: function changeOwnerSigned(identity address, sigV uint8, sigR bytes32, sigS bytes32, newOwner address) returns()
func (_Did *DidTransactorSession) ChangeOwnerSigned(identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, newOwner common.Address) (*types.Transaction, error) {
	return _Did.Contract.ChangeOwnerSigned(&_Did.TransactOpts, identity, sigV, sigR, sigS, newOwner)
}

// RevokeAttribute is a paid mutator transaction binding the contract method 0x8dd83056.
//
// Solidity: function revokeAttribute(identity address, name bytes, value bytes) returns()
func (_Did *DidTransactor) RevokeAttribute(opts *bind.TransactOpts, identity common.Address, name []byte, value []byte) (*types.Transaction, error) {
	return _Did.contract.Transact(opts, "revokeAttribute", identity, name, value)
}

// RevokeAttribute is a paid mutator transaction binding the contract method 0x8dd83056.
//
// Solidity: function revokeAttribute(identity address, name bytes, value bytes) returns()
func (_Did *DidSession) RevokeAttribute(identity common.Address, name []byte, value []byte) (*types.Transaction, error) {
	return _Did.Contract.RevokeAttribute(&_Did.TransactOpts, identity, name, value)
}

// RevokeAttribute is a paid mutator transaction binding the contract method 0x8dd83056.
//
// Solidity: function revokeAttribute(identity address, name bytes, value bytes) returns()
func (_Did *DidTransactorSession) RevokeAttribute(identity common.Address, name []byte, value []byte) (*types.Transaction, error) {
	return _Did.Contract.RevokeAttribute(&_Did.TransactOpts, identity, name, value)
}

// RevokeAttributeSigned is a paid mutator transaction binding the contract method 0xdff0d6f4.
//
// Solidity: function revokeAttributeSigned(identity address, sigV uint8, sigR bytes32, sigS bytes32, name bytes, value bytes) returns()
func (_Did *DidTransactor) RevokeAttributeSigned(opts *bind.TransactOpts, identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, name []byte, value []byte) (*types.Transaction, error) {
	return _Did.contract.Transact(opts, "revokeAttributeSigned", identity, sigV, sigR, sigS, name, value)
}

// RevokeAttributeSigned is a paid mutator transaction binding the contract method 0xdff0d6f4.
//
// Solidity: function revokeAttributeSigned(identity address, sigV uint8, sigR bytes32, sigS bytes32, name bytes, value bytes) returns()
func (_Did *DidSession) RevokeAttributeSigned(identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, name []byte, value []byte) (*types.Transaction, error) {
	return _Did.Contract.RevokeAttributeSigned(&_Did.TransactOpts, identity, sigV, sigR, sigS, name, value)
}

// RevokeAttributeSigned is a paid mutator transaction binding the contract method 0xdff0d6f4.
//
// Solidity: function revokeAttributeSigned(identity address, sigV uint8, sigR bytes32, sigS bytes32, name bytes, value bytes) returns()
func (_Did *DidTransactorSession) RevokeAttributeSigned(identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, name []byte, value []byte) (*types.Transaction, error) {
	return _Did.Contract.RevokeAttributeSigned(&_Did.TransactOpts, identity, sigV, sigR, sigS, name, value)
}

// RevokeAuthentication is a paid mutator transaction binding the contract method 0x3ea1a651.
//
// Solidity: function revokeAuthentication(identity address, authenticationType bytes32, authentication bytes32) returns()
func (_Did *DidTransactor) RevokeAuthentication(opts *bind.TransactOpts, identity common.Address, authenticationType [32]byte, authentication [32]byte) (*types.Transaction, error) {
	return _Did.contract.Transact(opts, "revokeAuthentication", identity, authenticationType, authentication)
}

// RevokeAuthentication is a paid mutator transaction binding the contract method 0x3ea1a651.
//
// Solidity: function revokeAuthentication(identity address, authenticationType bytes32, authentication bytes32) returns()
func (_Did *DidSession) RevokeAuthentication(identity common.Address, authenticationType [32]byte, authentication [32]byte) (*types.Transaction, error) {
	return _Did.Contract.RevokeAuthentication(&_Did.TransactOpts, identity, authenticationType, authentication)
}

// RevokeAuthentication is a paid mutator transaction binding the contract method 0x3ea1a651.
//
// Solidity: function revokeAuthentication(identity address, authenticationType bytes32, authentication bytes32) returns()
func (_Did *DidTransactorSession) RevokeAuthentication(identity common.Address, authenticationType [32]byte, authentication [32]byte) (*types.Transaction, error) {
	return _Did.Contract.RevokeAuthentication(&_Did.TransactOpts, identity, authenticationType, authentication)
}

// RevokeAuthenticationSigned is a paid mutator transaction binding the contract method 0xf8156eb8.
//
// Solidity: function revokeAuthenticationSigned(identity address, sigV uint8, sigR bytes32, sigS bytes32, authenticationType bytes32, authentication bytes32) returns()
func (_Did *DidTransactor) RevokeAuthenticationSigned(opts *bind.TransactOpts, identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, authenticationType [32]byte, authentication [32]byte) (*types.Transaction, error) {
	return _Did.contract.Transact(opts, "revokeAuthenticationSigned", identity, sigV, sigR, sigS, authenticationType, authentication)
}

// RevokeAuthenticationSigned is a paid mutator transaction binding the contract method 0xf8156eb8.
//
// Solidity: function revokeAuthenticationSigned(identity address, sigV uint8, sigR bytes32, sigS bytes32, authenticationType bytes32, authentication bytes32) returns()
func (_Did *DidSession) RevokeAuthenticationSigned(identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, authenticationType [32]byte, authentication [32]byte) (*types.Transaction, error) {
	return _Did.Contract.RevokeAuthenticationSigned(&_Did.TransactOpts, identity, sigV, sigR, sigS, authenticationType, authentication)
}

// RevokeAuthenticationSigned is a paid mutator transaction binding the contract method 0xf8156eb8.
//
// Solidity: function revokeAuthenticationSigned(identity address, sigV uint8, sigR bytes32, sigS bytes32, authenticationType bytes32, authentication bytes32) returns()
func (_Did *DidTransactorSession) RevokeAuthenticationSigned(identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, authenticationType [32]byte, authentication [32]byte) (*types.Transaction, error) {
	return _Did.Contract.RevokeAuthenticationSigned(&_Did.TransactOpts, identity, sigV, sigR, sigS, authenticationType, authentication)
}

// RevokePublicKey is a paid mutator transaction binding the contract method 0x787768ef.
//
// Solidity: function revokePublicKey(identity address, publicKeyType bytes32, publicKey bytes32) returns()
func (_Did *DidTransactor) RevokePublicKey(opts *bind.TransactOpts, identity common.Address, publicKeyType [32]byte, publicKey [32]byte) (*types.Transaction, error) {
	return _Did.contract.Transact(opts, "revokePublicKey", identity, publicKeyType, publicKey)
}

// RevokePublicKey is a paid mutator transaction binding the contract method 0x787768ef.
//
// Solidity: function revokePublicKey(identity address, publicKeyType bytes32, publicKey bytes32) returns()
func (_Did *DidSession) RevokePublicKey(identity common.Address, publicKeyType [32]byte, publicKey [32]byte) (*types.Transaction, error) {
	return _Did.Contract.RevokePublicKey(&_Did.TransactOpts, identity, publicKeyType, publicKey)
}

// RevokePublicKey is a paid mutator transaction binding the contract method 0x787768ef.
//
// Solidity: function revokePublicKey(identity address, publicKeyType bytes32, publicKey bytes32) returns()
func (_Did *DidTransactorSession) RevokePublicKey(identity common.Address, publicKeyType [32]byte, publicKey [32]byte) (*types.Transaction, error) {
	return _Did.Contract.RevokePublicKey(&_Did.TransactOpts, identity, publicKeyType, publicKey)
}

// RevokePublicKeySigned is a paid mutator transaction binding the contract method 0x3be52425.
//
// Solidity: function revokePublicKeySigned(identity address, sigV uint8, sigR bytes32, sigS bytes32, publicKeyType bytes32, publicKey bytes32) returns()
func (_Did *DidTransactor) RevokePublicKeySigned(opts *bind.TransactOpts, identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, publicKeyType [32]byte, publicKey [32]byte) (*types.Transaction, error) {
	return _Did.contract.Transact(opts, "revokePublicKeySigned", identity, sigV, sigR, sigS, publicKeyType, publicKey)
}

// RevokePublicKeySigned is a paid mutator transaction binding the contract method 0x3be52425.
//
// Solidity: function revokePublicKeySigned(identity address, sigV uint8, sigR bytes32, sigS bytes32, publicKeyType bytes32, publicKey bytes32) returns()
func (_Did *DidSession) RevokePublicKeySigned(identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, publicKeyType [32]byte, publicKey [32]byte) (*types.Transaction, error) {
	return _Did.Contract.RevokePublicKeySigned(&_Did.TransactOpts, identity, sigV, sigR, sigS, publicKeyType, publicKey)
}

// RevokePublicKeySigned is a paid mutator transaction binding the contract method 0x3be52425.
//
// Solidity: function revokePublicKeySigned(identity address, sigV uint8, sigR bytes32, sigS bytes32, publicKeyType bytes32, publicKey bytes32) returns()
func (_Did *DidTransactorSession) RevokePublicKeySigned(identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, publicKeyType [32]byte, publicKey [32]byte) (*types.Transaction, error) {
	return _Did.Contract.RevokePublicKeySigned(&_Did.TransactOpts, identity, sigV, sigR, sigS, publicKeyType, publicKey)
}

// SetAttribute is a paid mutator transaction binding the contract method 0xccbfa496.
//
// Solidity: function setAttribute(identity address, name bytes, value bytes, validity uint256) returns()
func (_Did *DidTransactor) SetAttribute(opts *bind.TransactOpts, identity common.Address, name []byte, value []byte, validity *big.Int) (*types.Transaction, error) {
	return _Did.contract.Transact(opts, "setAttribute", identity, name, value, validity)
}

// SetAttribute is a paid mutator transaction binding the contract method 0xccbfa496.
//
// Solidity: function setAttribute(identity address, name bytes, value bytes, validity uint256) returns()
func (_Did *DidSession) SetAttribute(identity common.Address, name []byte, value []byte, validity *big.Int) (*types.Transaction, error) {
	return _Did.Contract.SetAttribute(&_Did.TransactOpts, identity, name, value, validity)
}

// SetAttribute is a paid mutator transaction binding the contract method 0xccbfa496.
//
// Solidity: function setAttribute(identity address, name bytes, value bytes, validity uint256) returns()
func (_Did *DidTransactorSession) SetAttribute(identity common.Address, name []byte, value []byte, validity *big.Int) (*types.Transaction, error) {
	return _Did.Contract.SetAttribute(&_Did.TransactOpts, identity, name, value, validity)
}

// SetAttributeSigned is a paid mutator transaction binding the contract method 0xc7b2864d.
//
// Solidity: function setAttributeSigned(identity address, sigV uint8, sigR bytes32, sigS bytes32, name bytes, value bytes, validity uint256) returns()
func (_Did *DidTransactor) SetAttributeSigned(opts *bind.TransactOpts, identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, name []byte, value []byte, validity *big.Int) (*types.Transaction, error) {
	return _Did.contract.Transact(opts, "setAttributeSigned", identity, sigV, sigR, sigS, name, value, validity)
}

// SetAttributeSigned is a paid mutator transaction binding the contract method 0xc7b2864d.
//
// Solidity: function setAttributeSigned(identity address, sigV uint8, sigR bytes32, sigS bytes32, name bytes, value bytes, validity uint256) returns()
func (_Did *DidSession) SetAttributeSigned(identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, name []byte, value []byte, validity *big.Int) (*types.Transaction, error) {
	return _Did.Contract.SetAttributeSigned(&_Did.TransactOpts, identity, sigV, sigR, sigS, name, value, validity)
}

// SetAttributeSigned is a paid mutator transaction binding the contract method 0xc7b2864d.
//
// Solidity: function setAttributeSigned(identity address, sigV uint8, sigR bytes32, sigS bytes32, name bytes, value bytes, validity uint256) returns()
func (_Did *DidTransactorSession) SetAttributeSigned(identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, name []byte, value []byte, validity *big.Int) (*types.Transaction, error) {
	return _Did.Contract.SetAttributeSigned(&_Did.TransactOpts, identity, sigV, sigR, sigS, name, value, validity)
}

// DidDIDAttributeChangedIterator is returned from FilterDIDAttributeChanged and is used to iterate over the raw logs and unpacked data for DIDAttributeChanged events raised by the Did contract.
type DidDIDAttributeChangedIterator struct {
	Event *DidDIDAttributeChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DidDIDAttributeChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DidDIDAttributeChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DidDIDAttributeChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DidDIDAttributeChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DidDIDAttributeChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DidDIDAttributeChanged represents a DIDAttributeChanged event raised by the Did contract.
type DidDIDAttributeChanged struct {
	Identity       common.Address
	Name           []byte
	Value          []byte
	ValidTo        *big.Int
	PreviousChange *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDIDAttributeChanged is a free log retrieval operation binding the contract event 0x011b18dd995a3172c6dbe3b65ce383beec725369bb7cc6c16721013f9f993a78.
//
// Solidity: e DIDAttributeChanged(identity indexed address, name bytes, value bytes, validTo uint256, previousChange uint256)
func (_Did *DidFilterer) FilterDIDAttributeChanged(opts *bind.FilterOpts, identity []common.Address) (*DidDIDAttributeChangedIterator, error) {

	var identityRule []interface{}
	for _, identityItem := range identity {
		identityRule = append(identityRule, identityItem)
	}

	logs, sub, err := _Did.contract.FilterLogs(opts, "DIDAttributeChanged", identityRule)
	if err != nil {
		return nil, err
	}
	return &DidDIDAttributeChangedIterator{contract: _Did.contract, event: "DIDAttributeChanged", logs: logs, sub: sub}, nil
}

// WatchDIDAttributeChanged is a free log subscription operation binding the contract event 0x011b18dd995a3172c6dbe3b65ce383beec725369bb7cc6c16721013f9f993a78.
//
// Solidity: e DIDAttributeChanged(identity indexed address, name bytes, value bytes, validTo uint256, previousChange uint256)
func (_Did *DidFilterer) WatchDIDAttributeChanged(opts *bind.WatchOpts, sink chan<- *DidDIDAttributeChanged, identity []common.Address) (event.Subscription, error) {

	var identityRule []interface{}
	for _, identityItem := range identity {
		identityRule = append(identityRule, identityItem)
	}

	logs, sub, err := _Did.contract.WatchLogs(opts, "DIDAttributeChanged", identityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DidDIDAttributeChanged)
				if err := _Did.contract.UnpackLog(event, "DIDAttributeChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// DidDIDAuthenticationChangedIterator is returned from FilterDIDAuthenticationChanged and is used to iterate over the raw logs and unpacked data for DIDAuthenticationChanged events raised by the Did contract.
type DidDIDAuthenticationChangedIterator struct {
	Event *DidDIDAuthenticationChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DidDIDAuthenticationChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DidDIDAuthenticationChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DidDIDAuthenticationChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DidDIDAuthenticationChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DidDIDAuthenticationChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DidDIDAuthenticationChanged represents a DIDAuthenticationChanged event raised by the Did contract.
type DidDIDAuthenticationChanged struct {
	Identity           common.Address
	AuthenticationType [32]byte
	Authentication     [32]byte
	ValidTo            *big.Int
	PreviousChange     *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterDIDAuthenticationChanged is a free log retrieval operation binding the contract event 0x6c1587d1ee8f4bdfb125ab048123336ebb2fb5f901a4561a8218e340ceffd2e0.
//
// Solidity: e DIDAuthenticationChanged(identity indexed address, authenticationType bytes32, authentication bytes32, validTo uint256, previousChange uint256)
func (_Did *DidFilterer) FilterDIDAuthenticationChanged(opts *bind.FilterOpts, identity []common.Address) (*DidDIDAuthenticationChangedIterator, error) {

	var identityRule []interface{}
	for _, identityItem := range identity {
		identityRule = append(identityRule, identityItem)
	}

	logs, sub, err := _Did.contract.FilterLogs(opts, "DIDAuthenticationChanged", identityRule)
	if err != nil {
		return nil, err
	}
	return &DidDIDAuthenticationChangedIterator{contract: _Did.contract, event: "DIDAuthenticationChanged", logs: logs, sub: sub}, nil
}

// WatchDIDAuthenticationChanged is a free log subscription operation binding the contract event 0x6c1587d1ee8f4bdfb125ab048123336ebb2fb5f901a4561a8218e340ceffd2e0.
//
// Solidity: e DIDAuthenticationChanged(identity indexed address, authenticationType bytes32, authentication bytes32, validTo uint256, previousChange uint256)
func (_Did *DidFilterer) WatchDIDAuthenticationChanged(opts *bind.WatchOpts, sink chan<- *DidDIDAuthenticationChanged, identity []common.Address) (event.Subscription, error) {

	var identityRule []interface{}
	for _, identityItem := range identity {
		identityRule = append(identityRule, identityItem)
	}

	logs, sub, err := _Did.contract.WatchLogs(opts, "DIDAuthenticationChanged", identityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DidDIDAuthenticationChanged)
				if err := _Did.contract.UnpackLog(event, "DIDAuthenticationChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// DidDIDPublicKeyChangedIterator is returned from FilterDIDPublicKeyChanged and is used to iterate over the raw logs and unpacked data for DIDPublicKeyChanged events raised by the Did contract.
type DidDIDPublicKeyChangedIterator struct {
	Event *DidDIDPublicKeyChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DidDIDPublicKeyChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DidDIDPublicKeyChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DidDIDPublicKeyChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DidDIDPublicKeyChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DidDIDPublicKeyChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DidDIDPublicKeyChanged represents a DIDPublicKeyChanged event raised by the Did contract.
type DidDIDPublicKeyChanged struct {
	Identity       common.Address
	PublicKeyType  [32]byte
	PublicKey      [32]byte
	ValidTo        *big.Int
	PreviousChange *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDIDPublicKeyChanged is a free log retrieval operation binding the contract event 0x213edff8255aa608dd36ca4cf7ed92d5480ddcd178865af1053f9d0bf5c77338.
//
// Solidity: e DIDPublicKeyChanged(identity indexed address, publicKeyType bytes32, publicKey bytes32, validTo uint256, previousChange uint256)
func (_Did *DidFilterer) FilterDIDPublicKeyChanged(opts *bind.FilterOpts, identity []common.Address) (*DidDIDPublicKeyChangedIterator, error) {

	var identityRule []interface{}
	for _, identityItem := range identity {
		identityRule = append(identityRule, identityItem)
	}

	logs, sub, err := _Did.contract.FilterLogs(opts, "DIDPublicKeyChanged", identityRule)
	if err != nil {
		return nil, err
	}
	return &DidDIDPublicKeyChangedIterator{contract: _Did.contract, event: "DIDPublicKeyChanged", logs: logs, sub: sub}, nil
}

// WatchDIDPublicKeyChanged is a free log subscription operation binding the contract event 0x213edff8255aa608dd36ca4cf7ed92d5480ddcd178865af1053f9d0bf5c77338.
//
// Solidity: e DIDPublicKeyChanged(identity indexed address, publicKeyType bytes32, publicKey bytes32, validTo uint256, previousChange uint256)
func (_Did *DidFilterer) WatchDIDPublicKeyChanged(opts *bind.WatchOpts, sink chan<- *DidDIDPublicKeyChanged, identity []common.Address) (event.Subscription, error) {

	var identityRule []interface{}
	for _, identityItem := range identity {
		identityRule = append(identityRule, identityItem)
	}

	logs, sub, err := _Did.contract.WatchLogs(opts, "DIDPublicKeyChanged", identityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DidDIDPublicKeyChanged)
				if err := _Did.contract.UnpackLog(event, "DIDPublicKeyChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}
