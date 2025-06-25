package data

type OTPData struct {
	PhoneNumber string `json:"phoneNumber,omitempty" validation:"required"`
}

type VerifyData struct {
	User *OTPData `json:"user,omitempty" validation:"required"`
	Code string   `json:"code,omitempty" validation:"required"`
}
