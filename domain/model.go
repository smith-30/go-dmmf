package domain

import (
	"net/mail"
)

type EmailAddress string

func NewEmailAddress(v string) (EmailAddress, error) {
	re := EmailAddress(v)
	if err := re.Validate(); err != nil {
		return re, err
	}
	return re, nil
}

func (a EmailAddress) String() string {
	return string(a)
}

func (a EmailAddress) Validate() error {
	_, err := mail.ParseAddress(a.String())
	if err != nil {
		return err
	}
	return nil
}

type CustomerEmail interface {
	Verified() bool
	String() EmailAddress
}

type UnverifiedEmail EmailAddress

func NewUnverifiedEmail(v EmailAddress) CustomerEmail {
	_v := UnverifiedEmail(v)
	return &_v
}

func (a *UnverifiedEmail) Verified() bool {
	return false
}

func (a *UnverifiedEmail) String() EmailAddress {
	return EmailAddress(*a)
}

type VerifiedEmail EmailAddress

type Customer struct {
	Email CustomerEmail
}

func NewCustomer(e EmailAddress) Customer {
	return Customer{
		Email: NewUnverifiedEmail(e),
	}
}
