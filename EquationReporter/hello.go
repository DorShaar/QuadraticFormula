package equationreporter

import (
	"fmt"
	"net/http"
)

type HttpHandler struct{}

// Implement `ServeHTTP` method on `HttpHandler` struct
func (h HttpHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	data := []byte("Hello World!")
	res.Write(data)
}

func main() {
	fmt.Println("Equation Reporter started")

	handler := HttpHandler{}

	http.ListenAndServe(":9000", handler)
}
