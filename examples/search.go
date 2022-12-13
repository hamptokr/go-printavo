package main

import (
	"fmt"
	"log"

	"github.com/hamptokr/go-printavo"
)

func main() {
	p, err := printavo.NewClient("test@example.com", "my-super-secret-token")
	if err != nil {
		log.Fatal(err)
	}

	opt := new(printavo.OrderSearchOptions)
	opt.Query = "123456"

	orders, _, err := p.OrdersService.Search(opt)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(orders.Data); i++ {
		o := orders.Data[i]
		fmt.Printf("[%d | %.2f]: %s\n", o.Id, o.OrderTotal, o.OrderNickname)
	}
}
