# Lightweight http framework for go.

Note: This is not full-featured and is not (yet) intended for public
consumption.

```go

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/stevenle/web"
)

func pingHandler(ctx *web.Context) {
	web.WriteResponseString(ctx, "pong\n")
}

func helloHandler(ctx *web.Context) {
	greeting := fmt.Sprintf("Hello, %v!\n", ctx.Params["name"])
	web.WriteResponseString(ctx, greeting)
}

func main() {
	router := web.NewRouter()
	router.HandleFunc("/ping", pingHandler)
	router.HandleFunc("/hello/:name", helloHandler)

	s := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	log.Fatal(s.ListenAndServe())
}

```
