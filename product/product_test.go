package product_test

import (
	"fmt"
	"project6/product"
)

func ExampleMonthlyPayment() {
	p := product.Product{
		ID:     7291,
		Name:   "Xiaomi Redmi Note 14",
		Price:  2499000,
		Months: 12,
	}
	fmt.Println(product.MonthlyPayment(p))
	// Output: 208250
}
