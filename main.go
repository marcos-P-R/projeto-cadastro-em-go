package main

import (
	"net/http"
	"projeto-atelie-di/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
