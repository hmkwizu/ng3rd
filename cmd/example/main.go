package main

import (
	"log"

	"github.com/hmkwizu/ng3rd/fcm"

	"github.com/hmkwizu/ng3rd/sms"
)

func main() {

	log.Println("Example running..")

	//send SMS
	smsConfig := &sms.BLConfig{SMSSenderName: "YourSenderName", SMSUserName: "YourUserName", SMSPassword: "YourPassword"}
	smsConfig.SMSAPIKey = "YourAPIKey"
	smsConfig.SMSAPIURL = "API URL"
	sms.SendSMS(smsConfig, "Your code is 123456", "+XXXXXXXXXXXX")

	//send Push Notification
	fcm.SendPushMessage("API Key", []string{"Token"}, "My Title", "Hello world!", nil)

	log.Println("Example completed!")
}
