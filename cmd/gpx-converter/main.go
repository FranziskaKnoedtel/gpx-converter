package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// parse arguments for trace file
	if len(os.Args) != 2 {
		fmt.Println("Usage: traceload <tracefile>")
		os.Exit(1)
	}
	tracefile := os.Args[1]
	gpxfile := strings.TrimSuffix(tracefile, ".csv") + ".gpx"

	// create gpx file
	file, err := os.Create(gpxfile)
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

	var idxLat int = -1
	var idxLon int = -1

	for i, row := range tracedata {
		if i == 0 {
			for j, col := range row {
				if col == "lat" {
					idxLat = j
				}
				if col == "lon" {
					idxLon = j
				}
			}

			if idxLat == -1 || idxLon == -1 {
				log.Panic("error parsing csv file: could not find lat/lon column")
			}
			continue
		}
		lat, _ := strconv.ParseFloat(row[idxLat], 32)
		lon, _ := strconv.ParseFloat(row[idxLon], 32)
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
