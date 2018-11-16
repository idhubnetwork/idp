package resolver

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type DIDLog struct {
	DidPublicKeyLogs      []LogPublicKeyChanged
	DidAuthenticationLogs []LogAuthenticationChanged
	DidAttributeLogs      []LogAttributeChanged
}

type DIDDocument struct {
	Context        string              `json:"@context"`
	Id             string              `json:"id"`
	PublicKey      []DIDPublicKey      `json:"publicKey"`
	Authentication []DIDAuthentication `json:"authentication"`
	Service        []DIDAttribute      `json:"service"`
}

type DIDPublicKey struct {
	Id           string `json:"id"`
	Type         string `json:"type"`
	Owner        string `json:"owner"`
	PublicKeyHex string `json:"publicKeyHex"`
}

type DIDAuthentication struct {
	Type      string `json:"type"`
	PublicKey string `json:"publicKey"`
}

type DIDAttribute struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func (r *resolver) GetDIDLogs(address string) (DIDLog, error) {
	identity := common.HexToAddress(address)
	publicKeyChanged, err := r.GetPublicKeyChanged(address)
	if err != nil {
		return DIDLog{}, err
	}
	authenticationChanged, err := r.GetAuthenticationChanged(address)
	if err != nil {
		return DIDLog{}, err
	}
	attributeChanged, err := r.GetAttributeChanged(address)
	if err != nil {
		return DIDLog{}, err
	}
	var DidPublicKeyLogs = make([]LogPublicKeyChanged, 0)
	for publicKeyChanged.Sign() != 0 {
		logPublicKeyChangeds, err := r.EventPublicKeyChanged(publicKeyChanged)
		if err != nil {
			return DIDLog{}, err
		}
		for _, logV := range logPublicKeyChangeds {
			if identity == logV.Identity {
				DidPublicKeyLogs = append(DidPublicKeyLogs, logV)
				publicKeyChanged = logV.PreviousChange
			} else {
				continue
			}
		}
	}
	var DidAuthenticationLogs = make([]LogAuthenticationChanged, 0)
	for authenticationChanged.Sign() != 0 {
		logAuthenticationChangeds, err := r.EventAuthenticationChanged(authenticationChanged)
		if err != nil {
			return DIDLog{}, err
		}
		for _, logV := range logAuthenticationChangeds {
			if identity == logV.Identity {
				DidAuthenticationLogs = append(DidAuthenticationLogs, logV)
				authenticationChanged = logV.PreviousChange
			} else {
				continue
			}
		}
	}
	var DidAttributeLogs = make([]LogAttributeChanged, 0)
	for attributeChanged.Sign() != 0 {
		logAttributeChangeds, err := r.EventAttributeChanged(attributeChanged)
		if err != nil {
			return DIDLog{}, err
		}
		for _, logV := range logAttributeChangeds {
			if identity == logV.Identity {
				DidAttributeLogs = append(DidAttributeLogs, logV)
				attributeChanged = logV.PreviousChange
			} else {
				continue
			}
		}
	}
	return DIDLog{DidPublicKeyLogs, DidAuthenticationLogs, DidAttributeLogs}, nil
}

func (r *resolver) getDIDPublicKeys(address string, did DIDLog) ([]DIDPublicKey, error) {
	owner, err := r.IdentityOwner(address)
	if err != nil {
		return nil, err
	}
	DIDPublicKeys := make([]DIDPublicKey, 0)
	DIDPublicKeys = append(DIDPublicKeys, DIDPublicKey{
		"did:idhub:" + address + "#owner",
		"Secp256k1VerificationKey2018",
		"did:idhub:" + address,
		owner})
	for i, logV := range did.DidPublicKeyLogs {
		ok, err := r.ValidPublicKey(address, "veriKey",
			hexutil.Encode(logV.PublicKey[:]))
		if err != nil {
			return nil, err
		} else if ok {
			DIDPublicKeys = append(DIDPublicKeys, DIDPublicKey{
				"did:idhub:" + address + "#" + strconv.Itoa(i+1),
				"Secp256k1VerificationKey2018",
				"did:idhub:" + address,
				hexutil.Encode(logV.PublicKey[:])})
		}
	}
	return DIDPublicKeys, nil
}

func (r *resolver) getDIDAuthentications(address string, did DIDLog) ([]DIDAuthentication, error) {
	DIDPublicKeys, err := r.getDIDPublicKeys(address, did)
	if err != nil {
		return nil, err
	}
	DIDAuthentications := make([]DIDAuthentication, 0)
	DIDAuthentications = append(DIDAuthentications, DIDAuthentication{
		"Secp256k1SignatureAuthentication2018",
		"did:idhub:" + address + "#owner"})
	for _, logV := range did.DidAuthenticationLogs {
		ok, err := r.ValidAuthentication(address, "sigAuth",
			hexutil.Encode(logV.Authentication[:]))
		if err != nil {
			return nil, err
		} else if ok {
			for _, v := range DIDPublicKeys {
				if hexutil.Encode(logV.Authentication[:]) ==
					v.PublicKeyHex {
					DIDAuthentications = append(DIDAuthentications,
						DIDAuthentication{
							"Secp256k1SignatureAuthentication2018",
							v.Id})
				}
			}
		}
	}
	return DIDAuthentications, nil
}

func (r *resolver) getDIDAttributes(address string, did DIDLog) []DIDAttribute {
	DIDAttributes := make([]DIDAttribute, 0)
	for _, logV := range did.DidAttributeLogs {
		DIDAttributes = append(DIDAttributes, DIDAttribute{
			string(logV.Name[:]),
			string(logV.Value[:])})
	}
	return DIDAttributes
}

func (r *resolver) getDIDDocument(address string) (DIDDocument, error) {
	var document DIDDocument
	document.Context = "https://w3id.org/did/v1"
	document.Id = "did:idhub:" + address
	did, err := r.GetDIDLogs(address)
	if err != nil {
		return DIDDocument{}, err
	}
	document.PublicKey, err = r.getDIDPublicKeys(address, did)
	if err != nil {
		return DIDDocument{}, err
	}
	document.Authentication, err = r.getDIDAuthentications(address, did)
	if err != nil {
		return DIDDocument{}, err
	}
	document.Service = r.getDIDAttributes(address, did)
	return document, nil
}

func (r *resolver) GetDocument(address string) (string, error) {
	document, err := r.getDIDDocument(address)
	if err != nil {
		return "", err
	}
	data, err := json.Marshal(document)
	if err != nil {
		return "", err
	}
	fmt.Println(document)
	fmt.Println(string(data))
	return string(data), nil
}
