package email

import (
	"bytes"
	_ "embed"
	"fmt"
	"html/template"
)

//go:embed templates/otp.html
var otpTemplateHTML string

var otpTmpl = template.Must(template.New("otp").Parse(otpTemplateHTML))

type otpEmailData struct {
	OTP   string
	Email string
}

func RenderOTPEmail(otp, email string) (string, error) {
	var buf bytes.Buffer
	if err := otpTmpl.Execute(&buf, otpEmailData{OTP: otp, Email: email}); err != nil {
		return "", fmt.Errorf("render otp email: %w", err)
	}
	return buf.String(), nil
}
