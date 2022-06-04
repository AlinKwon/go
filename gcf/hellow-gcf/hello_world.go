package helloworld

import (
	"fmt"
	"net/http"
)

// HelloGet is on HTTP Cloud Function.
func HelloGet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}
