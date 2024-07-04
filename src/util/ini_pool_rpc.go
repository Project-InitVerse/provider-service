package util

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type QuoteBidOrder struct {
	ProviderAddr string `json:"provider_addr"`
	OrderAddr    string `json:"order_addr"`
	CpuPrice     string `json:"cpu_price"`
	MemoryPrice  string `json:"memory_price"`
	StoragePrice string `json:"storage_price"`
	SignMsg      string `json:"sign_msg,omitempty"`
}

type SubmitURI struct {
	ProviderAddr string `json:"provider_addr"`
	OrderAddr    string `json:"order_addr"`
	NewUrl       string `json:"new_url"`
	SignMsg      string `json:"sign_msg,omitempty"`
}

type PayBill struct {
	ProviderAddr string `json:"provider_addr"`
	OrderAddr    string `json:"order_addr"`
	SignMsg      string `json:"sign_msg,omitempty"`
}

type UpdateResource struct {
	ProviderAddr string `json:"provider_addr"`
	Cpu          string `json:"cpu"`
	Memory       string `json:"memory"`
	Storage      string `json:"storage"`
	SignMsg      string `json:"sign_msg,omitempty"`
}

type EndChallenge struct {
	ProviderAddr string `json:"provider_addr"`
	SignMsg      string `json:"sign_msg,omitempty"`
}
type RespResult struct {
	TrxID string `json:"result"`
}
type RespError struct {
	ErrMsg string `json:"error"`
}

func (ini *EndChallenge) Sign(prev *ecdsa.PrivateKey) {
	ini.SignMsg = ""
	hashBytes, err := json.Marshal(ini)
	if err != nil {
		fmt.Println("sign end challenge fail")
		return
	}
	hash := sha256.Sum256(hashBytes)
	sig, err := ecdsa.SignASN1(rand.Reader, prev, hash[:])
	if err != nil {
		fmt.Println("sign end challenge fail")
		return
	}
	ini.SignMsg = hex.EncodeToString(sig)
}

func (ini *EndChallenge) Send(baseURI string) (bool, string) {
	url := baseURI + "/endChallenge"
	cli := &http.Client{
		Timeout: 5 * time.Second,
	}
	jsonData, err := json.Marshal(ini)
	if err != nil {
		fmt.Println("Error marshaling quote:", err)
		return false, ""
	}

	body := bytes.NewBuffer(jsonData)

	method := http.MethodPost
	req, _ := http.NewRequest(method, url, body)
	req.Header.Set("Content-Type", "application/json")
	resp, err := cli.Do(req)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	var result RespResult
	var errResult RespError
	data, _ := io.ReadAll(resp.Body)
	err = json.Unmarshal(data, &result)
	if err != nil {
		err = json.Unmarshal(data, &errResult)
		if err != nil {
			fmt.Println("Error marshaling quote:", err)
			return false, ""
		}
		return false, errResult.ErrMsg
	}
	return true, result.TrxID
}

func (ini *PayBill) Sign(prev *ecdsa.PrivateKey) {
	ini.SignMsg = ""
	hashBytes, err := json.Marshal(ini)
	if err != nil {
		fmt.Println("sign end challenge fail")
		return
	}
	hash := sha256.Sum256(hashBytes)
	sig, err := ecdsa.SignASN1(rand.Reader, prev, hash[:])
	if err != nil {
		fmt.Println("sign end challenge fail")
		return
	}
	ini.SignMsg = hex.EncodeToString(sig)
}

func (ini *PayBill) Send(baseURI string) (bool, string) {
	url := baseURI + "/payBill"
	cli := &http.Client{
		Timeout: 5 * time.Second,
	}
	jsonData, err := json.Marshal(ini)
	if err != nil {
		fmt.Println("Error marshaling quote:", err)
		return false, ""
	}

	body := bytes.NewBuffer(jsonData)

	method := http.MethodPost
	req, _ := http.NewRequest(method, url, body)
	req.Header.Set("Content-Type", "application/json")
	resp, err := cli.Do(req)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	var result RespResult
	var errResult RespError
	data, _ := io.ReadAll(resp.Body)
	err = json.Unmarshal(data, &result)
	if err != nil {
		err = json.Unmarshal(data, &errResult)
		if err != nil {
			fmt.Println("Error marshaling quote:", err)
			return false, ""
		}
		return false, errResult.ErrMsg
	}
	return true, result.TrxID
}

func (ini *UpdateResource) Sign(prev *ecdsa.PrivateKey) {
	ini.SignMsg = ""
	hashBytes, err := json.Marshal(ini)
	if err != nil {
		fmt.Println("sign end challenge fail")
		return
	}
	hash := sha256.Sum256(hashBytes)
	sig, err := ecdsa.SignASN1(rand.Reader, prev, hash[:])
	if err != nil {
		fmt.Println("sign end challenge fail")
		return
	}
	ini.SignMsg = hex.EncodeToString(sig)
}

func (ini *UpdateResource) Send(baseURI string) (bool, string) {
	url := baseURI + "/updateResource"
	cli := &http.Client{
		Timeout: 5 * time.Second,
	}
	jsonData, err := json.Marshal(ini)
	if err != nil {
		fmt.Println("Error marshaling quote:", err)
		return false, ""
	}

	body := bytes.NewBuffer(jsonData)

	method := http.MethodPost
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		fmt.Println(err.Error())
		return false, ""
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := cli.Do(req)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	var result RespResult
	var errResult RespError
	data, _ := io.ReadAll(resp.Body)
	fmt.Println(string(data))
	err = json.Unmarshal(data, &result)
	if err != nil {
		err = json.Unmarshal(data, &errResult)
		if err != nil {
			fmt.Println("Error marshaling quote:", err)
			return false, ""
		}
		return false, errResult.ErrMsg
	}
	return true, result.TrxID
}

func (ini *SubmitURI) Sign(prev *ecdsa.PrivateKey) {
	ini.SignMsg = ""
	hashBytes, err := json.Marshal(ini)
	if err != nil {
		fmt.Println("sign end challenge fail")
		return
	}
	hash := sha256.Sum256(hashBytes)
	sig, err := ecdsa.SignASN1(rand.Reader, prev, hash[:])
	if err != nil {
		fmt.Println("sign end challenge fail")
		return
	}
	ini.SignMsg = hex.EncodeToString(sig)
}

func (ini *SubmitURI) Send(baseURI string) (bool, string) {
	url := baseURI + "/submitURI"
	cli := &http.Client{
		Timeout: 5 * time.Second,
	}
	jsonData, err := json.Marshal(ini)
	if err != nil {
		fmt.Println("Error marshaling quote:", err)
		return false, ""
	}

	body := bytes.NewBuffer(jsonData)

	method := http.MethodPost
	req, _ := http.NewRequest(method, url, body)
	req.Header.Set("Content-Type", "application/json")
	resp, err := cli.Do(req)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	var result RespResult
	var errResult RespError
	data, _ := io.ReadAll(resp.Body)
	err = json.Unmarshal(data, &result)
	if err != nil {
		err = json.Unmarshal(data, &errResult)
		if err != nil {
			fmt.Println("Error marshaling quote:", err)
			return false, ""
		}
		return false, errResult.ErrMsg
	}
	return true, result.TrxID
}

func (ini *QuoteBidOrder) Sign(prev *ecdsa.PrivateKey) {
	ini.SignMsg = ""
	hashBytes, err := json.Marshal(ini)
	if err != nil {
		fmt.Println("sign end challenge fail")
		return
	}
	hash := sha256.Sum256(hashBytes)
	sig, err := ecdsa.SignASN1(rand.Reader, prev, hash[:])
	if err != nil {
		fmt.Println("sign end challenge fail")
		return
	}
	ini.SignMsg = hex.EncodeToString(sig)
}

func (ini *QuoteBidOrder) Send(baseURI string) (bool, string) {
	url := baseURI + "/quoteBidOrder"
	cli := &http.Client{
		Timeout: 5 * time.Second,
	}
	jsonData, err := json.Marshal(ini)
	if err != nil {
		fmt.Println("Error marshaling quote:", err)
		return false, ""
	}

	body := bytes.NewBuffer(jsonData)

	method := http.MethodPost
	req, _ := http.NewRequest(method, url, body)
	req.Header.Set("Content-Type", "application/json")
	resp, err := cli.Do(req)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	var result RespResult
	var errResult RespError
	data, _ := io.ReadAll(resp.Body)
	err = json.Unmarshal(data, &result)
	if err != nil {
		err = json.Unmarshal(data, &errResult)
		if err != nil {
			fmt.Println("Error marshaling quote:", err)
			return false, ""
		}
		return false, errResult.ErrMsg
	}
	return true, result.TrxID
}
