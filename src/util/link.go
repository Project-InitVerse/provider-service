package util

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/bitly/go-simplejson"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type LinkClient struct {
	Url string
}

type HttpRequest struct {
	Id     string        `json:"id"`
	Method string        `json:"method"`
	Params []interface{} `json:"params"`
}

var substr = []string{"key_approvals_to_add", "true", "false"}

func notUseQuote(value string) bool {
	contain := false
	for _, subStr := range substr {
		if strings.Contains(value, subStr) {
			contain = true
			break
		}
	}
	return contain
}
func turnInterface(ina *[]interface{}) string {
	message := "["
	for _, value := range *ina {
		if reflect.TypeOf(value) == reflect.TypeOf(1) {
			if message == "[" {
				message = message + strconv.Itoa(value.(int))
			} else {
				message = message + "," + strconv.Itoa(value.(int))
			}
		} else if value == "true" {
			if message == "[" {
				message = message + value.(string)
			} else {
				message = message + "," + value.(string)
			}
		} else if reflect.TypeOf(value) == reflect.TypeOf(make(map[string]string)) {
			bak_str := "{"
			for k, v := range value.(map[string]string) {
				bak_str += "\"" + k + "\":\"" + v + "\","
			}
			bak_str = bak_str[:len(bak_str)-1]
			bak_str += "}"
			if message == "[" {

				message = message + bak_str
			} else {
				message = message + "," + bak_str
			}
		} else {
			if message == "[" {
				message = message + "\"" + value.(string) + "\""
			} else {
				message = message + ",\"" + value.(string) + "\""
			}
		}

	}
	message = message + "]"
	return message
}

func (client *LinkClient) SafeLinkHttpFunc(function string, params *[]interface{}) *simplejson.Json {
	sleepInterval := []int{5, 10, 20, 30, 40, 60, 120, 240, 480, 960, 1920, 3840}
	index := 0
	for {
		return_value := client.LinkHttpFunc(function, params)
		if return_value != nil {
			_, exist := return_value.CheckGet("result")
			if exist {
				return return_value
			}
		}
		{

			if index >= 12 {
				fmt.Println("http request is error,please wait to retry,current sleep time is ", time.Second*time.Duration(sleepInterval[11]), time.Now())
				time.Sleep(time.Second * time.Duration(sleepInterval[11]))
			} else {
				fmt.Println("http request is error,please wait to retry,current sleep time is ", time.Second*time.Duration(sleepInterval[index]), time.Now())
				time.Sleep(time.Second * time.Duration(sleepInterval[index]))
			}
		}
		index++
	}
}

func (client *LinkClient) UnSafeLinkHttpFunc(function string, params *[]interface{}) *simplejson.Json {
	sleepInterval := []int{5, 10, 20, 30, 40, 60, 120, 240, 480, 960, 1920, 3840}
	index := 0
	for {
		return_value := client.LinkHttpFunc(function, params)
		if return_value != nil {
			return return_value
		}
		{

			if index >= 12 {
				fmt.Println("http request is error,please wait to retry,current sleep time is ", time.Second*time.Duration(sleepInterval[11]), time.Now())
				time.Sleep(time.Second * time.Duration(sleepInterval[11]))
			} else {
				fmt.Println("http request is error,please wait to retry,current sleep time is ", time.Second*time.Duration(sleepInterval[index]), time.Now())
				time.Sleep(time.Second * time.Duration(sleepInterval[index]))
			}
		}
		index++
	}
}

func (client *LinkClient) LinkHttpFunc(function string, params *[]interface{}) *simplejson.Json {
	strParams := turnInterface(params)
	transport := http.Transport{
		DisableKeepAlives: true,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	clienta := http.Client{Transport: &transport}

	message := "{ \"jsonrpc\":\"2.0\",\"id\": 1, \"method\": \"" + function + "\", \"params\": " + strParams + "}"
	payload := strings.NewReader(message)
	//fmt.Println(payload)

	url := client.Url

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {

		fmt.Println(err.Error())
		return nil
	}
	req.Header.Add("content-type", "application/json")
	//if client.User != "" && client.PassWord != "" {
	//	encodeUser := base64.StdEncoding.EncodeToString([]byte((*client).User + ":" + (*client).PassWord))
	//	req.Header.Add("authorization", "Basic "+encodeUser)
	//}
	res, err := clienta.Do(req)
	if err != nil {
		fmt.Println("1")
		fmt.Println(err.Error())
		return nil
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if res.StatusCode != 200 {
		fmt.Println("error status code " + strconv.Itoa(res.StatusCode))
		return nil
	}
	if err != nil {

		fmt.Println(err.Error())
		return nil
	}
	//fmt.Println(string(body))
	js, err := simplejson.NewJson(body)
	if err != nil {

		fmt.Println(err.Error())
		return nil
	}
	//fmt.Println(js)
	return js
}
func (client *LinkClient) HttpRpcFunction(function string, param *[]interface{}) string {
	url := (*client).Url

	a := HttpRequest{"1", function, *param}
	fmt.Println(a)
	b, _ := json.Marshal(a)
	payload := strings.NewReader(string(b))
	//encodeUser := base64.StdEncoding.EncodeToString([]byte((*client).User + ":" + (*client).PassWord))
	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("content-type", "application/json")
	//req.Header.Add("authorization", "Basic "+encodeUser)
	req.Header.Add("cache-control", "no-cache")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	tempparam := make(map[string]interface{})
	json.Unmarshal([]byte(string(body)), &tempparam)
	fmt.Println(string(body))
	return string(body)
}
func (link_client *LinkClient) EthCall(address string, methodId string) string {
	call_map := make(map[string]string, 0)
	call_map["to"] = address
	call_map["data"] = methodId
	//call_data,_ := json.Marshal(call_map)
	call_param := make([]interface{}, 0, 2)
	call_param = append(call_param, call_map)
	call_param = append(call_param, "latest")
	call_return, exist := link_client.UnSafeLinkHttpFunc("eth_call", &call_param).CheckGet("result")

	if exist {
		call_res, err := call_return.String()
		//fmt.Println(call_res,err)
		if err != nil {
			return "0x"
		} else {
			return call_res
		}

	} else {
		return "0x"
	}
}
