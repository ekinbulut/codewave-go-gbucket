package main

import (
	"fmt"
	"os"
)

type App struct {
	Name    string
	Version string
}

func (a *App) Run() {

	// display app info
	fmt.Printf("%s version %s\n", a.Name, a.Version)

	// get environment variable of aws_access_key_id
	accessKeyId := os.Getenv("AWS_ACCESS_KEY_ID")

	// get environment variable of aws_secret_access_key
	secretAccessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")

	// print access key id and secret access key
	fmt.Println(accessKeyId)
	fmt.Println(secretAccessKey)

}

// read aws credentials from ~/.aws/credentials
// [default]
// aws_access_key_id = AKIAJZQZQZQZQZQZQZQ

// read flags from command line
// gbucket-cli -v
// gbucket-cli -h
// gbucket-cli -help
// gbucket-cli -version
// gbucket-cli -access-key-id AKIAJZQZQZQZQZQZQZQ -secret-access-key AKIAJZQZQZQZQZQZQZQ
// gbucket-cli -access-key-id AKIAJZQZQZQZQZQZQZQ -secret-access-key AKIAJZQZQZQZQZQZQZQ -region us-east-1
// gbucket-cli -access-key-id AKIAJZQZQZQZQZQZQZQ -secret-access-key AKIAJZQZQZQZQZQZQZQ -region us-east-1 -bucket-name gbucket-cli-test
// gbucket-cli -access-key-id AKIAJZQZQZQZQZQZQZQ -secret-access-key AKIAJZQZQZQZQZQZQZQ -region us-east-1 -bucket-name gbucket-cli-test -file-name test.txt
// gbucket-cli -access-key-id AKIAJZQZQZQZQZQZQZQ -secret-access-key AKIAJZQZQZQZQZQZQZQ -region us-east-1 -bucket-name gbucket-cli-test -file-name test.txt -file-content test

func main() {

	// create new app
	app := &App{
		Name:    "gbucket-cli",
		Version: "0.0.1",
	}

	// run app
	app.Run()
}
