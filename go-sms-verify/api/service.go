package api

import (
	"errors"

	twilio "github.com/twilio/twilio-go"
	verify "github.com/twilio/twilio-go/rest/verify/v2"
)

// Create a global Twilio client
var client *twilio.RestClient = twilio.NewRestClientWithParams(twilio.ClientParams{
	Username: envACCOUNTSID(),
	Password: envAUTHTOKEN(),
})

// Send OTP
func (app *Config) twilioSendOTP(phoneNumber string) (string, error) {
	params := &verify.CreateVerificationParams{}
	params.SetTo(phoneNumber)
	params.SetChannel("sms")

	resp, err := client.VerifyV2.CreateVerification(envSERVICEID(), params)
	if err != nil {
		return "", err
	}
	return *resp.Sid, nil
}

// Verify OTP
func (app *Config) twilioVerifyOTP(phoneNumber string, code string) error {
	params := &verify.CreateVerificationCheckParams{}
	params.SetTo(phoneNumber)
	params.SetCode(code)

	resp, err := client.VerifyV2.CreateVerificationCheck(envSERVICEID(), params)
	if err != nil {
		return err
	}
	if resp.Status != nil && *resp.Status != "approved" {
		return errors.New("not a valid code")
	}
	return nil
}
