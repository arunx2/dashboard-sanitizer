package main

import (
	"flag"
	"fmt"
)

func main() {
	//generate command flags
	esObjectFile := flag.String("f", "", "The Elastic dashboard object file in ndjson")
	outputFile := flag.String("o", "os_dashboard_objects", "The OpenSearch compatible dashboard object file in ndjson")

	/*
		TODO:
		1. Read the source line by line
		2. skip if the `Type` is lens
		3. update the MigrationVersion
		4.
	*/
	fmt.Printf("%s file is sanitized and output available at %s", *esObjectFile, *outputFile)

}
