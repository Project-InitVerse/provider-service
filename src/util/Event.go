package util

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"providerService/src/config"
	"strings"
)

type AllEvent struct {
	CanQuote        string
	UserCancelOrder string
	ChooseQuote     string
}

func (self *AllEvent) Init() {
	self.ChooseQuote = crypto.Keccak256Hash([]byte("ChooseQuote((address,uint256,uint256,uint256),uint256)")).String()
	self.UserCancelOrder = crypto.Keccak256Hash([]byte("UserCancelOrder()")).String()
	self.CanQuote = crypto.Keccak256Hash([]byte("CanQuote()")).String()

}
func (self *AllEvent) CheckEqual(a, b string) bool {
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

func GenerateOrderCall(orderContract string) string {
	addr := common.HexToAddress(orderContract)
	method_id := "0x4bb21d44"

	return method_id + "000000000000000000000000" + addr.Hex()
}
func GetOrderFactory(config *config.ProviderConfig) string {
	return config.OrderFactory
}
func GenerateOrderInfoCall() string {
	return "0x2500f01d"
}
