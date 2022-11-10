package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"runtime/debug"

	"github.com/oscerai/gqlgen/_examples/todo"
	"github.com/oscerai/gqlgen/graphql/handler"
	"github.com/oscerai/gqlgen/graphql/playground"
)

func main() {
	srv := handler.NewDefaultServer(todo.NewExecutableSchema(todo.New()))
	srv.SetRecoverFunc(func(ctx context.Context, err interface{}) (userMessage error) {
		// send this panic somewhere
		log.Print(err)
		debug.PrintStack()
		return errors.New("user message on panic")
	})

	http.Handle("/", playground.Handler("Todo", "/query"))
	http.Handle("/query", srv)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
