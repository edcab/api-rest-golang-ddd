package main

import rest "api-rest-golang-ddd/infrastructure"

func main() {

	api := rest.NewAPI()

	api.Start()

}
