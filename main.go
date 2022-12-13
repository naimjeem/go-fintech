package main

import (
	"fmt"

	"github.com/naimjeem/go-fintech/pkg/utils/migrations"
)

func main() {
	migrations.Migrate()
	fmt.Println("Server is running on port")

}
