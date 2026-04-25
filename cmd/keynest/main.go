package main

import (
	"fmt"

	"github.com/light-devteam/keynest/internal/config"
)

func main() {
	config := config.MustLoad()
	fmt.Println(config)
}
