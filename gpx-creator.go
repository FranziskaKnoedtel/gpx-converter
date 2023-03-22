package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// parse arguments for trace file
	if len(os.Args) != 2 {
		fmt.Println("Usage: traceload <tracefile>")
		os.Exit(1)
	}
	tracefile := os.Args[1]

	// create gpx file
	file, err := os.Create("trace.gpx")
	defer file.Close()
	if err != nil {
		log.Fatalln("failed to open file", err)
	}

	var head = `<?xml version="1.0" encoding="UTF-8" standalone="no" ?>
<gpx version="1.1" creator="franzi/gpx-creator" >
 <trk>
  <trkseg>
`
	// write header to file
	if _, err := file.WriteString(head); err != nil {
		log.Fatalln("error writing head to file", err)
	}

	// open csv file
	csvFile, err := os.Open(tracefile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	tracedata, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for i, row := range tracedata {
		if i == 0 {
			continue
		}
		lat, _ := strconv.ParseFloat(row[1], 32)
		lon, _ := strconv.ParseFloat(row[2], 32)
		// eph, _ := strconv.ParseFloat(row[4], 32)
		line := "    <trkpt lat=\"" + strconv.FormatFloat(lat, 'f', 6, 32) + "\" lon=\"" + strconv.FormatFloat(lon, 'f', 6, 32) + "\">\n    </trkpt>\n"
		if _, err := file.WriteString(line); err != nil {
			log.Fatalln("error writing line to file", err)
		}
	}

	var tail = `   </trkseg>
  </trk>
</gpx>
`
	// write tail to file
	if _, err := file.WriteString(tail); err != nil {
		log.Fatalln("error writing tail to file", err)
	}
	file.Sync()
}
