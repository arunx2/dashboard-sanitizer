package main

import (
	"bufio"
	"bytes"
	"dashboard-sanitizer/config"
	sm "dashboard-sanitizer/model"
	"flag"
	"fmt"
	"github.com/olivere/ndjson"
	"os"
)

func main() {
	//generate command flags
	esObjectFile := flag.String("source", "", "The Elastic dashboard object file in ndjson.")
	outputFile := flag.String("output", "os_dashboard_objects.ndjson", "The OpenSearch compatible dashboard object file in ndjson.")
	versionFlag := flag.Bool("version", false, "Prints the version number")
	flag.Parse()

	if *versionFlag {
		fmt.Print(config.Version)
		os.Exit(0)
	}
	if *esObjectFile == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	b, _ := os.ReadFile(*esObjectFile)
	reader := ndjson.NewReader(bytes.NewReader(b))
	f, _ := os.Create(*outputFile)
	defer f.Close()

	writer := ndjson.NewWriter(bufio.NewWriter(f))
	for reader.Next() {
		var do sm.DashboardObject
		if err := reader.Decode(&do); err != nil {
			fmt.Fprintf(os.Stderr, "Decode failed: %v", err)
			return
		}
		if do.Type == "lens" || do.Type == "" {
			continue
		}
		_ = do.MakeCompatibleToOS()
		err := writer.Encode(do)
		if err != nil {
			return
		}

	}
	fmt.Printf("%s file is sanitized and output available at %s", *esObjectFile, *outputFile)

}
