package main

import (
	"bufio"
	"dashboard-sanitizer/config"
	sm "dashboard-sanitizer/model"
	"encoding/json"
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
	file, err := os.Open(*esObjectFile)
	if err != nil {
		fmt.Print("Error reading source file: ", err)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(file)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1024*1024)

	f, _ := os.Create(*outputFile)
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)

	outBuff := bufio.NewWriter(f)
	writer := ndjson.NewWriter(outBuff)
	counter := NewStatusCount()
	for scanner.Scan() {
		line := scanner.Text()
		var do sm.DashboardObject
		if err := json.Unmarshal([]byte(line), &do); err != nil {
			fmt.Printf("Decode failed: %v", err)
			return
		}
		if !do.IsCompatibleType() {
			counter.RegisterSkipped(do)
			continue
		}
		_ = do.MakeCompatibleToOS()
		err := writer.Encode(do)
		if err != nil {
			fmt.Print(err)
			return
		}
		//Flush after each line processed
		err = outBuff.Flush()
		if err != nil {
			fmt.Print("Error writing to output file :", err)
		}
		counter.RegisterProcessed(do)

	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Printf("\n %s file is sanitized and output available at %s \n", *esObjectFile, *outputFile)
	counter.PrintStats()
}
