package main

import (
	"fmt"
	"maya/assig2/domain"
	"maya/assig2/mapstore"
)

type CustomerController struct {
	store domain.CustomerStore
}

func (c *CustomerController) Add(cust domain.Customer) {
	err := c.store.Create(cust)
	if err != nil {
		fmt.Println("error", err)
		return
	}
	fmt.Println("created")
}

func (c CustomerController) Update(ID string, cust domain.Customer) {
	err := c.store.Update(ID, cust)
	if err != nil {
		fmt.Println("error", err)
		return
	}
	fmt.Println("Updated")
}

func (c CustomerController) GETID(ID string) {
	result, err := c.store.GetByID(ID)
	if err != nil {
		fmt.Println("error", err)
		return
	}
	fmt.Println(result)
}

func (c CustomerController) Delete(ID string) {
	err := c.store.Delete(ID)
	if err != nil {
		fmt.Println("error", err)
		return
	}
	fmt.Println("deleted")
}
func (c CustomerController) GetAll() {
	result, err := c.store.GetAll()
	if err != nil {
		fmt.Println("error", err)
		return
	}
	fmt.Println(result)
}

func main() {
	controller := CustomerController{
		store: mapstore.NewMapStore(),
	}
	customer := domain.Customer{
		ID:   "1",
		Name: "Raj Das",
	}
	controller.Add(customer)

	newCustomer := domain.Customer{
		ID:   "1",
		Name: "Raj Das",
	}
	controller.GETID("1")

	controller.Update("1", newCustomer)
	controller.Delete("1")
	controller.GetAll()

}
