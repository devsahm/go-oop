package main

import (
	"fmt"
	"phptogo-oop/model"
)

// interface in go

type programmer interface {
	speak() string
}

type phpProgrammer struct {
}

type javaProgrammer struct {
}

type contactable interface {
	openingPhrase() string
}

type account struct {
	balance int
}

type adultAccount struct {
	account
}

type minorAccount struct {
	account
	limit int
}

func main() {

	//Adult Account
	adult := adultAccount{
		account: account{balance: 2000},
	}

	// adult.account.withdraw(1000)

	// adult.account.sendStatment()

	fmt.Println(eligible(adult.account))

	//Minor Account
	minor := minorAccount{
		account: account{balance: 3000},
		limit:   1000,
	}

	// minor.withdraw(2000)

	// minor.sendStatment()

	greetings := "Thank you for being our customer!"

	messageAdult := sendEmail(adult, greetings)

	messageMinor := sendEmail(minor, greetings)

	fmt.Println(messageAdult)

	fmt.Println(messageMinor)

	//-----------------------------------------------------------------------------

	programmers := []programmer{javaProgrammer{}, &phpProgrammer{}}

	for _, v := range programmers {
		fmt.Println(v.speak())
	}

	ip, _ := model.NewIp("829.920.23.9")

	fmt.Println(model.GetParts(ip))

}

func sendEmail(c contactable, message string) string {

	message = c.openingPhrase() + " " + message

	return message
}

func eligible(a account) bool {

	return a.balance >= 900

}

func (a *account) withdraw(amount int) {
	a.balance -= amount
}

func (a *account) sendStatment() {
	fmt.Println("The balance is:", a.balance)
}

func (a *minorAccount) withdraw(amount int) {

	if a.balance-amount < a.limit {
		fmt.Println("Account limit is exceeded")
	} else {
		a.account.balance -= amount
	}
}

// To implement the contactable interface to the adultAccount and minorAccount, do this

func (a adultAccount) openingPhrase() string {

	return "This message is in regards to your account."
}

func (a minorAccount) openingPhrase() string {

	return "This message is in regards to the minor account."
}

// This implement the programmer interface by using one of it method
func (p *phpProgrammer) speak() string {
	return "PHP is not dead"
}

func (j javaProgrammer) speak() string {
	return "OOP for life"
}

// func (i ipAddress) network() string {
// 	return strings.Join(i.parts[0:3], ".")

// }
