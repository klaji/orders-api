package main

import (
	"context"
	"fmt"
	"github.com/klaji/orders-api/application"
)

func main(){
	app := application.New()

	err := app.Start(context.TODO())
	if err != nil {
		fmt.Println("fail to start app:",err)
	}
}