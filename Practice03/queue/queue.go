package main

import "fmt"

type Queue []*Order

type Order struct {
	priority     int
	quantity     int
	product      string
	customerName string
}

func (o *Order) New(priority int, quantity int, product string, customerName string) {
	o.priority = priority
	o.quantity = quantity
	o.product = product
	o.customerName = customerName
}

func (q *Queue) Add(order *Order) {
	if len(*q) == 0 {
		*q = append(*q, order)
	} else {
		appended := false
		for i, addedOrder := range *q {
			if order.priority > addedOrder.priority {
				*q = append((*q)[:i], append(Queue{order}, (*q)[i:]...)...)
				appended = true
				break
			}
		}
		if !appended {
			*q = append(*q, order)
		}
	}
}

func main() {
	var queue Queue
	queue = make(Queue, 0)
	var order1 *Order = &Order{}
	var priority1 int = 2
	var quantity1 int = 20
	var product1 string = "Computer"
	var customerName1 string = "Greg White"
	order1.New(priority1, quantity1, product1, customerName1)
	var order2 *Order = &Order{}
	var priority2 int = 1
	var quantity2 int = 10
	var product2 string = "Monitor"
	var customerName2 string = "John Smith"
	order2.New(priority2, quantity2, product2, customerName2)
	queue.Add(order1)

	queue.Add(order2)
	var i int
	for i = 0; i < len(queue); i++ {
		fmt.Println(queue[i])
	}
}
