// resolver.go
package main

import (
	"fmt"
	"sync"
)

type Resolver struct {
	service *Service
	once    sync.Once
}

func (r *Resolver) ResolveService() *Service {
	r.once.Do(func() {
		fmt.Println("Inicializando o serviço...")
		r.service = &Service{}
	})
	return r.service
}
