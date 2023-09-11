package main

import (
	"fmt"
)

type Account struct {
	id          string
	accountType string
}

func (a *Account) create(accountType string) *Account {
	fmt.Println("account creation with type")
	a.accountType = accountType
	return a
}

func (a *Account) getById(id string) *Account {
	fmt.Println("getting account by Id")
	return a
}

func (a *Account) deleteById(id string) {
	fmt.Println("delete account by Id")
}

type Customer struct {
	name string
	id   int
}

func (c *Customer) create(name string) *Customer {
	fmt.Println("creating customer")
	c.name = name
	return c
}

type Transaction struct {
	id            string
	amount        float32
	srcAccountId  string
	destAccountId string
}

func (t *Transaction) create(srcAccountId, destAccountId string, amount float32) *Transaction {
	fmt.Println("creating customer")
	t.srcAccountId = srcAccountId
	t.destAccountId = destAccountId
	t.amount = amount
	return t
}

type BranchManagerFacade struct {
	account     *Account
	customer    *Customer
	transaction *Transaction
}

func NewBranchManagerFacade() *BranchManagerFacade {
	return &BranchManagerFacade{&Account{}, &Customer{}, &Transaction{}}
}

func (facade BranchManagerFacade) createCustomerAccount(customerName string, accountType string) (*Customer, *Account) {
	customer := facade.customer.create(customerName)
	account := facade.account.create(accountType)
	return customer, account
}

func (facade BranchManagerFacade) createTransaction(srcAccountId, destAccountId string, amount float32) *Transaction {
	transaction := facade.transaction.create(srcAccountId, destAccountId, amount)
	return transaction
}

func main() {
	var facade = NewBranchManagerFacade()
	var customer *Customer
	var account *Account
	customer, account = facade.createCustomerAccount("Thomas Smith",
		"Savings")
	fmt.Println(customer.name)
	fmt.Println(account.accountType)
	var transaction = facade.createTransaction("21456", "87345", 1000)
	fmt.Println(transaction.amount)
}
