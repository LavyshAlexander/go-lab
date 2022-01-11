package main

import "fmt"

func main() {
	var DeploymentOptions = [...]string{"R-pi", "AWS", "GCP", "Azure"}

	for index, option := range DeploymentOptions {
		fmt.Println(index, option)
	}
}
