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
	s.Set("key", "value")
	res := s.Get("key")
	fmt.Println(res)
}
