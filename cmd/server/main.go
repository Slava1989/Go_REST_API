package main

import "fmt"

// MARK: responsible for instantiation and startup of app
func Run() error {
	fmt.Println("startin up our app")
	return nil
}

func main() {
	fmt.Println("Go Rest API")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
