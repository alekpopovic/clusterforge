package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "version" {
		fmt.Println("clusterforge dev")
		return
	}

	fmt.Println("ClusterForge CLI")
	fmt.Println("Terraform/OpenTofu logic stays visible in generated files.")
}
