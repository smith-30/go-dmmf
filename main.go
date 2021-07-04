package main

import (
	"encoding/json"
	"fmt"

	"github.com/smith-30/go-dmmf/domain"
)

func main() {
	c := domain.NewCustomer("test___@gmail.com")
	bss, _ := json.MarshalIndent(c, "", "	")
	fmt.Printf("%v\n", string(bss))
}
