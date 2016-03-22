package main  // import "github.com/RobertPierceCB/hal-test-server"

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
)

func halResponse(c web.C, w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")

	testResponse, _ := ioutil.ReadFile("fixtures/tunnelResponse.json")

	fmt.Fprintln(w, string(testResponse))
}

func main() {
	goji.Get("/*", halResponse)
	goji.Serve()
}
