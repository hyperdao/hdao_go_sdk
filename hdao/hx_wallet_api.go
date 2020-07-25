package hdao

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"

	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// construct rpc request msg to hx mainchain
type HXWalletApi struct {
	Name         string
	Rpc_user     string
	Rpc_password string
	Rpc_url      string
}

type HxPayLoad struct {
	Id     int
	Method string
	Params []interface{}
}

func NewHXWalletApi(name string, rpc_user string, rpc_password string, rpc_url string) *HXWalletApi {
	r := HXWalletApi{
		Name:         name,
		Rpc_user:     rpc_user,
		Rpc_url:      rpc_url,
		Rpc_password: rpc_password,
	}
	return &r
}

func (walletApi *HXWalletApi) Rpc_request(method string, args []interface{}) (string, error) {
	var result = ""
	argstr, err := json.Marshal(args)
	payload := fmt.Sprintf("{\r\n \"id\": 1,\r\n \"method\": \"%s\",\r\n \"params\": %s\r\n}", method, argstr)

	//fmt.Printf("palaoad:%s\n", payload)
	request, err := http.NewRequest("POST", walletApi.Rpc_url, strings.NewReader(payload))
	if err != nil {
		return result, err
	}

	auth_str := walletApi.Rpc_user + ":" + walletApi.Rpc_password

	basic_auth := base64.StdEncoding.EncodeToString([]byte(auth_str))
	request.Header.Set("Content-Type", "application/json") //add request header
	request.Header.Add("cache-control", "no-cache")
	request.Header.Add("User-Agent", "Web Client")
	request.Header.Add("Authorization", basic_auth)
	resp, err := http.DefaultClient.Do(request.WithContext(context.TODO())) //send request
	defer resp.Body.Close()
	if err != nil {
		fmt.Println(err.Error())
		return result, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return result, err
	}

	var m map[string]interface{}
	err = json.Unmarshal(body, &m)
	if err != nil {
		return result, err
	}
	r, ok := m["result"]
	if !ok {
		errmsg := "no result in response"
		errorinfo, ok := m["error"]
		if ok {
			errormap, ok := errorinfo.(map[string]interface{})
			if ok {
				ermsg, _ := errormap["message"]
				errmsg, _ = ermsg.(string)
			}
		}
		return result, errors.New(errmsg)
	}
	result, ok = r.(string)
	if !ok {
		return result, errors.New("result is not string")
	}
	return result, nil
}
