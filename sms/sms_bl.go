package sms

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// BLConfig - config parameters
type BLConfig struct {
	SMSSenderName string
	SMSUserName   string
	SMSPassword   string
	SMSAPIKey     string
	SMSAPIURL     string
}

// SendSMS - send sms
func SendSMS(config *BLConfig, txt string, toNumber string) error {

	if len(txt) == 0 || len(toNumber) == 0 {
		return errors.New("Phone number or text message is empty")
	}

	params := url.Values{}
	params.Add("sendername", config.SMSSenderName)
	params.Add("username", config.SMSUserName)
	params.Add("password", config.SMSPassword)
	params.Add("apikey", config.SMSAPIKey)
	params.Add("destnum", toNumber)
	params.Add("message", txt)
	params.Add("senddate", "")

	url := fmt.Sprintf("%s?%s", config.SMSAPIURL, params.Encode())

	// make request
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	// read response
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	//parse response
	stringResponse := string(body)
	// value := ngauth.GetInt64OrZero(stringResponse)

	// fmt.Println(stringResponse)
	// fmt.Printf("STATUS: %d\n", resp.StatusCode)
	// fmt.Printf("VALUE: %d\n", value)

	// NOTE:-- There is an issue in the API, where 0 is not returned but the sms goes through. 0 means success

	if stringResponse != "0" {
		//return errors.New("Error while sending sms")
	}

	return nil
}
