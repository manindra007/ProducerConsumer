package queue

import "fmt"
type Product struct {
	Message string
}

var q []Product
var mut bool = true

func Push(msg Product) error {
	for {
		if mut == true {
			mut = false
			q = append(q, msg)
			mut = true
			fmt.Println("New Product added, ",len(q)," product available.")
			return nil
		}
	}
}
func Pop() (Product,error) {
	for {
		if len(q) == 0 {
			return Product{},fmt.Errorf("no product found")
		}
		if mut == true {
			mut = false
			var msg Product = q[0]
			q = q[1:]
			fmt.Println("A Product Consumed, ",len(q)," product available.")
			mut = true
			return msg,nil
		}
	}
}
