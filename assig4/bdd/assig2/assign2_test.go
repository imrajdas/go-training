package main

import (
	"fmt"
	"maya/assig2/domain"
	"maya/assig2/mapstore"
	"reflect"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type CustomerController struct {
	store domain.CustomerStore
}

var controller = CustomerController{
	store: mapstore.NewMapStore(),
}

func TestControllers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Controllers Suite")
}

var _ = Describe("BDD on employee app", func() {

	Context("Employee add testing", func() {

		It("should be return nil", func() {

			customer := domain.Customer{
				ID:   "1",
				Name: "Raj Das",
			}

			err := controller.store.Create(customer)
			if err != nil {
				fmt.Println("error", err)
				return
			}
			fmt.Println(reflect.TypeOf(err))
			Expect(err).To(BeNil())

		})
	})

	Context("Employee update testing", func() {

		It("should be return nil", func() {

			customer := domain.Customer{
				ID:   "2",
				Name: "Deep",
			}

			err := controller.store.Update("1", customer)
			if err != nil {
				fmt.Println("error", err)
				return
			}

			Expect(err).To(BeNil())

		})
	})

	Context("Employee get testing", func() {

		It("should be return error nil", func() {

			_, err := controller.store.GetByID("1")
			Expect(err).To(BeNil())

		})
	})

	Context("Employee get all testing", func() {

		It("should be return nill", func() {

			_, err := controller.store.GetAll()

			Expect(err).To(BeNil())

		})
	})

	Context("Employee delete testing", func() {

		It("should be return nil", func() {

			err := controller.store.Delete("2")
			Expect(err).To(BeNil())

		})
	})

})
