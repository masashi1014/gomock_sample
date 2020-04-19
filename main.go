package main

import (
	"fmt"
	"github.com/masashi1014/gomock_sample/application"
)

func main() {
	var id int64
	fmt.Scan(&id)

	app := application.NewApplication()
	app.FindOrCreateUser(id)
}
