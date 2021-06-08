package main

import (
	"fmt"
	"os"
	"time"

	githubactions "github.com/sethvargo/go-githubactions"
)

func main() {
	fmt.Printf("Hello, %s! I learned your name by directly accessing the environment variable. \n", os.Getenv("INPUT_FIRST"))
	fmt.Printf("::set-output name=one::%s \n", time.Now())

	fmt.Printf("Hello, %s! I learned your name from go-githubactions. \n", githubactions.GetInput("second"))
	githubactions.SetOutput("two", time.Now().String())
}
