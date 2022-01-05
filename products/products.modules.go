package products

import (
	"net/http"
)

func ProductsModule() {
	http.HandleFunc("/products", ProductsController)
}
