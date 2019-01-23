package email_validator

import "testing"

var (
	goodEmail = "dimon6407@gmail.com"
	badFormatEmail = "qwe;qwe;qwe:qwe-@qwer)).sdf"
	badDomainEmail = "qwerty@aaaaa.rg"
	badRequestEmail = "qwertyvbnfg@gmail.com"
)

func TestValidateGoodEmail(t *testing.T) {
	v, err := NewEmailValidator(nil)
	if err != nil {
		t.Fatal(err)
	}

	err = v.CheckEmail(goodEmail)
	if err != nil {
		t.Fatal(err)
	}
}

func TestValidateBadFormatEmail(t *testing.T) {
	v, err := NewEmailValidator(nil)
	if err != nil {
		t.Fatal(err)
	}

	err = v.CheckEmail(badFormatEmail)
	if err == nil || err != ErrorFormat {
		t.Fatal(err)
	}
}

func TestValidateBadDomainEmail(t *testing.T) {
	v, err := NewEmailValidator(nil)
	if err != nil {
		t.Fatal(err)
	}

	err = v.CheckEmail(badDomainEmail)
	if err == nil || err != ErrorDomain {
		t.Fatal(err)
	}
}

func TestValidateBadRequestEmail(t *testing.T) {
	v, err := NewEmailValidator(nil)
	if err != nil {
		t.Fatal(err)
	}

	err = v.CheckEmail(badRequestEmail)
	if err == nil || err != ErrorRequest {
		t.Fatal(err)
	}
}