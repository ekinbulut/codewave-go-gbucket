package main

import (
	"flag"
	"fmt"
	"os"
)

type App struct {
	Name    string
	Version string
}

func (a *App) Run() {

	parseFlags()

	if help {
		a.ShowHelp()
		os.Exit(0)
	}

	if version {
		a.ShowVersion()
		os.Exit(0)
	}

	a.Validate()

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

func (a *App) Validate() {
	// validate app
	if accessKeyId == "" {
		fmt.Println("access-key-id is required")
		os.Exit(1)
	}
	if secretAccessKey == "" {
		fmt.Println("secret-access-key is required")
		os.Exit(1)
	}
	if region == "" {
		fmt.Println("region is required")
		os.Exit(1)
	}
	if bucketName == "" {
		fmt.Println("bucket-name is required")
		os.Exit(1)
	}
	if fileName == "" {
		fmt.Println("file-name is required")
		os.Exit(1)
	}
	if fileContent == "" {
		fmt.Println("file-content is required")
		os.Exit(1)
	}
}

func (a *App) ShowHelp() {
	fmt.Printf("%s version %s\n", a.Name, a.Version)
	fmt.Println("Usage:")
	fmt.Println("  main [options]")
	fmt.Println("")
	fmt.Println("Options:")
	fmt.Println("  -h, --help")
	fmt.Println("    Show this help message and exit.")
	fmt.Println("")
}

func (a *App) ShowVersion() {
	fmt.Printf("%s version %s\n", a.Name, a.Version)
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

var (
	accessKeyId     string
	secretAccessKey string
	region          string
	bucketName      string
	fileName        string
	fileContent     string
	version         bool
	help            bool
)

func parseFlags() {

	flag.StringVar(&accessKeyId, "access-key-id", "", "AWS access key id")
	flag.StringVar(&secretAccessKey, "secret-access-key", "", "AWS secret access key")
	flag.StringVar(&region, "region", "", "AWS region")
	flag.StringVar(&bucketName, "bucket-name", "", "AWS bucket name")
	flag.StringVar(&fileName, "file-name", "", "AWS file name")
	flag.StringVar(&fileContent, "file-content", "", "AWS file content")
	flag.BoolVar(&version, "version", false, "show version")
	flag.BoolVar(&help, "help", false, "show help")
	flag.Parse()
}

func main() {

	// create new app
	app := &App{
		Name:    "gbucket-cli",
		Version: "0.0.1",
	}

	// run app
	app.Run()
}
