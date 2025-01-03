package goubus

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// UbusResponseCode represent the status code from JSON-RPC Call
//
//go:generate stringer -type=UbusResponseCode
type UbusResponseCode int

// Represents enum ubus_msg_status from https://git.openwrt.org/?p=project/ubus.git;a=blob;f=ubusmsg.h;h=398b126b6dc01833937749a110181ea0debb1476;hb=HEAD
const (
	UbusStatusOK               UbusResponseCode = 0
	UbusStatusInvalidCommand   UbusResponseCode = 1
	UbusStatusInvalidArgument  UbusResponseCode = 2
	UbusStatusMethodNotFound   UbusResponseCode = 3
	UbusStatusNotFound         UbusResponseCode = 4
	UbusStatusNoData           UbusResponseCode = 5
	UbusStatusPermissionDenied UbusResponseCode = 6
	UbusStatusTimeout          UbusResponseCode = 7
	UbusStatusNotSupported     UbusResponseCode = 8
	UbusStatusUnknownError     UbusResponseCode = 9
	UbusStatusConnectionFailed UbusResponseCode = 10
	UbusStatusLast             UbusResponseCode = 11
)

// Ubus represents information to JSON-RPC Interaction with router
type Ubus struct {
	Username string
	Password string
	URL      string
	AuthData *UbusAuthData
}

// UbusResponse represents a response from JSON-RPC
type UbusResponse struct {
	JSONRPC          string
	ID               int
	Error            UbusResponseError
	Result           interface{}
	UbusResponseCode UbusResponseCode
}

type UbusResponseError struct {
	Code    int
	Message string
}

type UbusExec struct {
	Code   int
	Stdout string
}

// LoginCheck check if login RPC Session id has expired
func (u *Ubus) LoginCheck() error {
	var err error
	for i := 0; i < 3; i++ {
		if u.AuthData == nil || u.AuthData.ExpireTime.Before(time.Now().Add(time.Minute*3)) {
			_, err = u.AuthLogin()
			if err == nil {
				return nil
			}
			time.Sleep(1 * time.Second)
		} else {
			return nil
		}
	}
	return err
}

// Call do a call to Json-RPC to get/set information
func (u *Ubus) Call(jsonStr []byte) (UbusResponse, error) {
	req, err := http.NewRequest("POST", u.URL, bytes.NewBuffer(jsonStr))
	if err != nil {
		return UbusResponse{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return UbusResponse{}, err
	}
	defer resp.Body.Close()
	if resp.Status != "200 OK" {
		return UbusResponse{}, fmt.Errorf("error %s on (%s)", resp.Status, u.URL)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	result := UbusResponse{}
	json.Unmarshal([]byte(body), &result)
	//Function Error
	if result.Error.Code != 0 {
		if strings.Compare(result.Error.Message, "Access denied") == 0 {
			return UbusResponse{}, errors.New("access denied for this instance, read https://openwrt.org/docs/techref/ubus#acls ")
		}
		return UbusResponse{}, errors.New(result.Error.Message)
	}
	//Workaround cause response code not contempled by unmarshal function
	result.UbusResponseCode = UbusResponseCode(result.Result.([]interface{})[0].(float64))
	//Workaround to get UbusData cause the structure of this array has a problem with unmarshal
	if result.UbusResponseCode == UbusStatusOK {
		return result, nil
	}
	return UbusResponse{}, fmt.Errorf("Ubus Status Failed: %s", result.UbusResponseCode)
}
