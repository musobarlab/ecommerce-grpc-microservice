package main

import (
	"fmt"
	"os"

	configEnv "github.com/joho/godotenv"
)

func main() {
	//load environtment variables
	err := configEnv.Load(".env")
	if err != nil {
		fmt.Println(".env is not loaded properly")
		os.Exit(2)
	}

	grpcMain()
}
