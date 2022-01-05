package users

import (
	"net/http"
)

func UsersModule() {
	http.HandleFunc("/users", UsersController)
}
