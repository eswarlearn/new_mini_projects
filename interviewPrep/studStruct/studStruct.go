package main

import "fmt"

type Prods struct {
	Prod     string
	Quantity int64
}

func main() {
	sl := &[]Prods{{Prod: "first", Quantity: 2}}
	var ProdNm string

	for i := 0; i <= 2; i++ {
		fmt.Println("Enter the product")
		fmt.Scan(&ProdNm)
		*sl = append(*sl, Prods{Prod: ProdNm})

	}

	fmt.Println("your slice", *sl)
	fmt.Println("prod", Prods{})

}
