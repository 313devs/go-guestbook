package main

import (
	"context"
	"fmt"

	"github.com/313devs/go-guestbook/application"
)

func main() {
	app := application.New()
	err := app.Start(context.TODO())
	if err != nil {
		fmt.Printf("failed to start application: %v\n", err)
	}
}

