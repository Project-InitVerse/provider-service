package main

import (
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
)

func main() {
	privateKey, err := ecdsa.GenerateKey(secp256k1.S256(), rand.Reader)
	if err != nil {
		fmt.Println("Generate communication key fail ", err.Error())
	}
	fmt.Println("Communication Private key is [", hex.EncodeToString(crypto.FromECDSA(privateKey)), "]")
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		fmt.Println("Turn communication key fail ", err.Error())
		return
	}
	fmt.Println("Communication Public key is [", hex.EncodeToString(crypto.FromECDSAPub(publicKeyECDSA)), "]")
	fmt.Println("Please upload communication public key to the pool")
}
