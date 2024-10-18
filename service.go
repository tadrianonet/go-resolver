package main

import "fmt"

type Service struct{}

func (s *Service) MakeRequest() {
	fmt.Println("Fazendo uma requisição!")
}
