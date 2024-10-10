package main

import (
	"Go_lang/internal/pkg/storage"
	"fmt"
	"log"
)

func main() {
	s, err := storage.NewStorage()

	if err != nil {
		log.Fatal(err)
	}

	//s.Set("dessert", "pancakes with honey")
	s.Set("passkey", "12345678")

	//fmt.Println(*s.Get("dessert"))
	//fmt.Println(s.GetKind("dessert"))

	fmt.Println(*s.Get("passkey"))
	fmt.Println(s.GetKind("passkey"))
}
