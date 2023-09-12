package main

import "database/sql"

type Customer struct {
	CustomerID   int
	CustomerName string
	SSN          string
}

func GetConnection() *sql.DB {
	databaseDriver := "mysql"
	databaseUser := "newuser"
	databasePass := "newuser"
	databaseName := "crm"
	database, error := sql.Open(databaseDriver, databaseUser+":"+databasePass+"@/"+databaseName)
	if error != nil {
		panic(error.Error())
	}
	return database
}

func GetCustomers() []Customer {
	var database *sql.DB
	database = GetConnection()

	var rows *sql.Rows
	rows, error := database.Query("SELECT * FROM Customer ORDER BY Customerid DESC")
	if error != nil {
		panic(error.Error())
	}
	customer := Customer{}

	var customers []Customer
	for rows.Next() {
		var customerId int
		var customerName string
		var ssn string
		error = rows.Scan(&customerId, &customerName, &ssn)
		if error != nil {
			panic(error.Error())
		}
		customer.CustomerID = customerId
		customer.CustomerName = customerName
		customer.SSN = ssn
		customers = append(customers, customer)
	}

	defer database.Close()

	return customers
}
