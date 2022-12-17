package main

import (
	"bufio"
	"bytes"
	sm "dashboard-sanitizer/model"
	"flag"
	"fmt"
	"github.com/olivere/ndjson"
	"os"
	"strings"
)

func main() {
	//generate command flags
	esObjectFile := flag.String("source", "", "The Elastic dashboard object file in ndjson.")
	outputFile := flag.String("output", "os_dashboard_objects.ndjson", "The OpenSearch compatible dashboard object file in ndjson.")

	/*
		TODO:
		1. Read the source line by line
		2. skip if the `Type` is lens
		3. update the MigrationVersion
		4.
	*/
	flag.Parse()

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
		switch do.Type {
		case "dashboard":
			//TODO: check if the version is greater than 7.9.3, leave the value as is if it is less than this.
			do.MigrationVersion.Dashboard = "7.9.3"
			//fix some visualization references name
			//var temp []sm.References
			for i, _ := range do.References {
				//if do.References[i].Type == "visualization" {
				do.References[i].Name = getNormalizedVizName(do.References[i].Name)
				//temp = append(temp, do.References[i])
				//}
			}
			//do.References = temp
		case "visualization":
			do.MigrationVersion.Visualization = "7.9.3"
			break
		case "index-pattern":
			do.MigrationVersion.IndexPattern = "7.6.0"
			break
		}
		writer.Encode(do)

	}
	fmt.Printf("%s file is sanitized and output available at %s", *esObjectFile, *outputFile)

}

func getNormalizedVizName(s string) string {
	if idx := strings.Index(s, ":"); idx != -1 {
		return s[idx+1:]
	}
	return s
}
