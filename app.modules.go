package main

import (
	"app/products"
	"app/users"
	"net/http"
)

func Bootstrap(port string) {
	users.UsersModule()
	products.ProductsModule()
	http.ListenAndServe(port, nil)
}
