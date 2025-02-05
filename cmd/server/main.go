package main

import (
	"github.com/ivofreitas/device-api/internal/api"
)

func main() {
	api.NewServer().Run()
}
