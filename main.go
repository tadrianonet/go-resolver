package main

func main() {
	resolver := &Resolver{}

	service := resolver.ResolveService()
	service.MakeRequest()

	serviceAgain := resolver.ResolveService()
	serviceAgain.MakeRequest()
}
