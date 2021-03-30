package security

import (
	"errors"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"

	"time"
)

var otps = totp.ValidateOpts{
	Period:    120,
	Skew:      1,
	Digits:    otp.DigitsSix,
	Algorithm: otp.AlgorithmSHA512,
}

// GeneratePasscode ...
func GeneratePasscode(issuer string, account string) (passcode string, key *otp.Key, err error) {
	key, err = totp.Generate(totp.GenerateOpts{
		Issuer:      issuer,  //"Example.com",
		AccountName: account, //"alice@example.com",
		Period:      otps.Period,
	})

	secret := key.Secret()

	passcode, err = totp.GenerateCodeCustom(secret, time.Now(), otps)
	if err != nil {
		return "", nil, err
	}
	return passcode, key, err
}

// ValidatePasscode ...
func ValidatePasscode(passcode string, secret string) (valid bool, err error) {
	valid, err = totp.ValidateCustom(
		passcode,
		secret,
		time.Now(),
		otps,
	)
	if err != nil {
		return false, err
	}

	if valid {
		return true, nil
	}
	return false, errors.New("passcode is not valid")
}
