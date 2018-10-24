package resolver

import (
	"fmt"
	"math/big"
	"os"
	"testing"
)

type ResolverTest struct {
	Owner                 string
	ContractAddress       string
	Identity              string
	AuthenticationChanged *big.Int
	PublicKey             string
	AttibuteChanged       *big.Int
}

var T = ResolverTest{
	"0x10F676D7326f3d62eE7480155bAF4daf833a980d",
	"0x1DbF8e4B47EA53a2b932850F7FEC8585C6A9c1d2",
	"0x49dBa8f906c745B0a82f4D21E02BAFD7Df1a0be4",
	big.NewInt(4203168),
	"0xa8d31fac615e26049614a3423eeb554bcada91024cf21059e806888c2cf4a756",
	big.NewInt(4203173),
}

func TestNewResolver(t *testing.T) {
	_, err := NewResolver("infuraRopsten", T.ContractAddress)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	fmt.Println("func NewResolver success")
	t.Log("func NewResolver success")
}

func TestGetAuthenticationChanged(t *testing.T) {
	r, _ := NewResolver("infuraRopsten", T.ContractAddress)
	block, err := r.GetAuthenticationChanged(T.Identity)
	if err != nil {
		t.Error(err)
		t.Fail()
	} else if T.AuthenticationChanged.Cmp(block) != 0 {
		t.Fail()
	} else {
		fmt.Println("func GetChanged success")
		t.Log("func GetChanged success")
	}
}

func TestIdentityOwner(t *testing.T) {
	r, _ := NewResolver("infuraRopsten", T.ContractAddress)
	addr, err := r.IdentityOwner(T.Identity)
	if err != nil {
		t.Error(err)
		t.Fail()
	} else if T.Owner != addr {
		t.Fail()
	} else {
		fmt.Println("func GetChanged success")
		t.Log("func GetChanged success")
	}
}

func TestValidPublicKey(t *testing.T) {
	r, _ := NewResolver("infuraRopsten", T.ContractAddress)
	_, err := r.ValidPublicKey(T.Identity, "veriKey", T.PublicKey)
	if err != nil {
		t.Error(err)
		t.Fail()
	} else {
		fmt.Println("func Valid success")
		t.Log("func Valid success")
	}
}

func TestEventAuthenticationChanged(t *testing.T) {
	r, _ := NewResolver("infuraRopsten", T.ContractAddress)
	_, err := r.EventAuthenticationChanged(T.AuthenticationChanged)
	if err != nil {
		t.Error(err)
		t.Fail()
	} else {
		fmt.Println("func AuthenticationChanged success")
		t.Log("func AuthenticationChanged success")
	}
}

func TestEventAttributeChanged(t *testing.T) {
	r, _ := NewResolver("infuraRopsten", T.ContractAddress)
	_, err := r.EventAttributeChanged(T.AttibuteChanged)
	if err != nil {
		t.Error(err)
		t.Fail()
	} else {
		fmt.Println("func AttributeChanged success")
		t.Log("func AttributeChanged success")
	}
}

func TestGetDIDLogs(t *testing.T) {
	r, _ := NewResolver("infuraRopsten", T.ContractAddress)
	_, err := r.GetDIDLogs(T.Identity)
	if err != nil {
		t.Error(err)
		t.Fail()
	} else {
		fmt.Println("func DIDLogs success")
		t.Log("func DIDLogs success")
	}
}

func TestGetDocument(t *testing.T) {
	r, _ := NewResolver("infuraRopsten", T.ContractAddress)
	doc, err := r.GetDocument(T.Identity)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	outputFile, outputError := os.OpenFile("did-doc.json", os.O_WRONLY|os.O_CREATE, 0666)
	if outputError != nil {
		fmt.Printf("An error occurred with file opening or creation\n")
		return
	}
	defer outputFile.Close()
	outputFile.WriteString(doc)
	fmt.Println("func DID Document success")
	t.Log("func DID Document success")
}
