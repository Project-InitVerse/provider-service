package util

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"providerService/src/config"
	"strings"
)

// AllEvent is scan event
type AllEvent struct {
	CanQuote        string
	UserCancelOrder string
	ChooseQuote     string
	OrderEnded      string
	ChallengeCreate string
	ChallengeEnd    string
}

// Init is initialize function
func (ae *AllEvent) Init() {
	ae.ChooseQuote = crypto.Keccak256Hash([]byte("ChooseQuote(address,uint256)")).String()
	ae.UserCancelOrder = crypto.Keccak256Hash([]byte("UserCancelOrder()")).String()
	ae.CanQuote = crypto.Keccak256Hash([]byte("CanQuote()")).String()
	ae.OrderEnded = crypto.Keccak256Hash([]byte("OrderEnded()")).String()
	ae.ChallengeCreate = crypto.Keccak256Hash([]byte("ChallengeCreate(address,uint256)")).String()
	ae.ChallengeEnd = crypto.Keccak256Hash([]byte("NeedChallenge()")).String()

}

// CheckEqual is function
func (ae *AllEvent) CheckEqual(a, b string) bool {
	if !strings.HasPrefix(a, "0x") {
		a = "0x" + a
	}
	if !strings.HasPrefix(b, "0x") {
		b = "0x" + b
	}
	if strings.ToLower(a) == strings.ToLower(b) {
		return true
	}
	return false
}

// GenerateOrderCall is function
func GenerateOrderCall(orderContract string) string {
	addr := common.HexToAddress(orderContract)
	methodID := "0x4bb21d44"

	return methodID + "000000000000000000000000" + addr.Hex()[2:]
}

// GetOrderFactory is function to order factory address
func GetOrderFactory(config *config.ProviderConfig) string {
	return config.OrderFactory
}

// GenerateOrderInfoCall is function
func GenerateOrderInfoCall() string {
	return "0x2500f01d"
}
