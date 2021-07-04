package main

import (
	"github.com/smith-30/go-dmmf/domain"
)

func main() {
	e1, _ := domain.NewEmailAddress("test___1@gmail.com")
	c1 := domain.Customer{
		Email: domain.NewUnverifiedEmail(e1),
	}
	e2, _ := domain.NewEmailAddress("test___2@gmail.com")
	c2 := domain.Customer{
		Email: domain.NewVerifiedEmail(e2),
	}
	PasswordReset(*c1.Email.(*domain.VerifiedEmail))
	PasswordReset(*c2.Email.(*domain.VerifiedEmail))
}

func PasswordReset(v domain.VerifiedEmail) {
}
