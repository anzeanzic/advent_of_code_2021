package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

type Beacon struct {
	x int
	y int
	z int
}
type ScannerReport struct {
	index  int
	report []Beacon
}

func main() {
	ReadFile()
}

func ReadFile() {
	file, err := os.Open("input_test.txt")

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	var ScannerReports []ScannerReport
	var scanner_report ScannerReport
	scanner_ndx := -1

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) > 0 {
			if string(line[0:3]) == "---" {
				if len(scanner_report.report) > 0 {
					ScannerReports = append(ScannerReports, scanner_report)
				}

				scanner_ndx++
				scanner_report.index = scanner_ndx
				scanner_report.report = make([]Beacon, 0)
			} else {
				line_arr := strings.Split(line, ",")
				x, x_err := strconv.Atoi(line_arr[0])
				y, y_err := strconv.Atoi(line_arr[1])
				//z, z_err := strconv.Atoi(line_arr[2])

				if x_err != nil {
					panic(x_err)
				}
				if y_err != nil {
					panic(y_err)
				}
				/*if z_err != nil {
					panic(z_err)
				}*/

				scanner_report.report = append(scanner_report.report, Beacon{x: x, y: y, z: 0})
			}
		}
	}

	if len(scanner_report.report) > 0 {
		ScannerReports = append(ScannerReports, scanner_report)
	}

	file.Close()
	fmt.Println(ScannerReports)

	CalculateDistanceBetweenBeacons(ScannerReports)
}

type ScannerDistMap struct {
	index    int
	dist_map map[Beacon][]float64
}

func CalculateDistanceBetweenBeacons(ScannerReports []ScannerReport) {
	var scannerDistMaps []ScannerDistMap

	for i := 0; i < len(ScannerReports); i++ {
		scanner_report := ScannerReports[i]

		dist_map := make(map[Beacon][]float64, 0)

		// build a distance map for every beacon of a scanner
		for i, beacon1 := range scanner_report.report {
			for j, beacon2 := range scanner_report.report {
				if i != j {
					dist := CalculateDistance(beacon1, beacon2)

					if _, ok := dist_map[beacon1]; !ok {
						dist_map[beacon1] = make([]float64, 0)
					}

					dist_map[beacon1] = append(dist_map[beacon1], dist)
				}
			}

			sort.Float64s(dist_map[beacon1])
		}

		scannerDistMaps = append(scannerDistMaps, ScannerDistMap{index: i, dist_map: dist_map})
	}

	fmt.Println(scannerDistMaps)

	CompareDistances(ScannerReports, scannerDistMaps)
}

func CalculateDistance(beacon1 Beacon, beacon2 Beacon) float64 {
	return math.Sqrt(math.Pow(float64(beacon2.x)-float64(beacon1.x), 2) + math.Pow(float64(beacon2.y)-float64(beacon1.y), 2))
}

func CompareDistances(ScannerReports []ScannerReport, scannerDistMaps []ScannerDistMap) {
	for i, scannerDistMap1 := range scannerDistMaps {
		for j, scannerDistMap2 := range scannerDistMaps {
			if i != j {
				//for j, beacon2 := range scanner_report.report {
				for beacon1, dist1 := range scannerDistMap1.dist_map {
					for beacon2, dist2 := range scannerDistMap2.dist_map {
						if reflect.DeepEqual(dist1, dist2) {
							fmt.Println(beacon1, beacon2)
						}
					}
				}
			}
		}
	}
}
