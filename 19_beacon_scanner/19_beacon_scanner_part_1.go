package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Beacon struct {
	x int
	y int
	z int
}
type BeaconPair struct {
	beacon_1 Beacon
	beacon_2 Beacon
}
type ScannerReport struct {
	index  int
	report []Beacon
}
type ScannerDistMap struct {
	index    int
	dist_map map[Beacon][]float64
}
type ScannerPair struct {
	scanner1_index      int
	scanner2_index      int
	scanner1_coordinate Beacon
	scanner2_coordinate Beacon
}

var ScannerReports []ScannerReport
var WholeMap map[Beacon]int
var FoundScanners []bool
var equalsMap map[Beacon]Beacon
var transformedCoordinatesMap map[Beacon]int

func main() {
	ReadFile()
}

func ReadFile() {
	file, err := os.Open("input_test.txt")

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
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
				z, z_err := strconv.Atoi(line_arr[2])

				if x_err != nil {
					panic(x_err)
				}
				if y_err != nil {
					panic(y_err)
				}
				if z_err != nil {
					panic(z_err)
				}

				scanner_report.report = append(scanner_report.report, Beacon{x: x, y: y, z: z})
			}
		}
	}

	if len(scanner_report.report) > 0 {
		ScannerReports = append(ScannerReports, scanner_report)
	}

	equalsMap = make(map[Beacon]Beacon, 0)
	transformedCoordinatesMap = make(map[Beacon]int)

	file.Close()
	//fmt.Println(ScannerReports)

	FoundScanners = make([]bool, len(ScannerReports))
	FoundScanners[0] = true

	orientations := [][]int{
		{1, 1, 1},
		{-1, 1, 1},
		{1, -1, 1},
		{1, 1, -1},
		{-1, -1, 1},
		{-1, 1, -1},
		{1, -1, -1},
		{-1, -1, -1},
	}
	orders := [][]int{
		{0, 1, 2},
		{0, 2, 1},
		{1, 0, 2},
		{1, 2, 0},
		{2, 0, 1},
		{2, 1, 0},
	}

	//for i := 0; i < len(orientations); i++ {
	//for j := 0; j < len(orders); j++ {
	fmt.Println(orientations[0], orders[0])
	CalculateDistanceBetweenBeacons(ScannerReports, orientations[0], orders[0])
	//}
	//}
}

func CalculateDistanceBetweenBeacons(ScannerReports []ScannerReport, orientation []int, order []int) {
	var scannerDistMaps []ScannerDistMap

	for i := 0; i < len(ScannerReports); i++ {
		scanner_report := ScannerReports[i]
		dist_map := make(map[Beacon][]float64, 0)

		// build a distance map for every beacon of a scanner
		for j, beacon1 := range scanner_report.report {
			for z, beacon2 := range scanner_report.report {
				if i != 0 {
					beacon1 = ConsiderCoordinateSwap(ConsiderOrientation(beacon1, orientation), order)
					beacon2 = ConsiderCoordinateSwap(ConsiderOrientation(beacon2, orientation), order)
				}

				if j != z {
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

	//fmt.Println(scannerDistMaps)

	//CompareDistances(ScannerReports, scannerDistMaps)
	fmt.Println(scannerDistMaps[0].dist_map)
	fmt.Println(scannerDistMaps[1].dist_map)
}

func ConsiderOrientation(beacon Beacon, orientation []int) Beacon {
	beacon.x = beacon.x * orientation[0]
	beacon.y = beacon.y * orientation[1]
	beacon.z = beacon.z * orientation[2]

	return beacon
}

func ConsiderCoordinateSwap(beacon Beacon, order []int) Beacon { // 1 0 2
	init_arr := []int{beacon.x, beacon.y, beacon.z}
	new_arr := make([]int, len(init_arr))

	for i := 0; i < len(order); i++ {
		new_arr[i] = init_arr[order[i]]
	}

	beacon.x = new_arr[0]
	beacon.y = new_arr[1]
	beacon.z = new_arr[2]

	return beacon
}

func CalculateDistance(beacon1 Beacon, beacon2 Beacon) (int, int, int) {
	diff1 := beacon2.x - beacon1.x
	diff2 := beacon2.y - beacon1.y
	diff3 := beacon2.z - beacon1.z

	return diff1, diff2, diff3
}

func CompareDistances(ScannerReports []ScannerReport, scannerDistMaps []ScannerDistMap) {
	count_equals := 0
	ScannerPairs := make([]ScannerPair, 0)

	// go through all scanners and compare them to other scanners
	for i, scannerDistMap1 := range scannerDistMaps {
		// go through the same scanners just don't compare two same scanners
		for j, scannerDistMap2 := range scannerDistMaps {
			if i != j {
				semi_count := 0
				SameBeacons := make([]BeaconPair, 0)

				// go through a hashmap of beacons and distances of a scanner
				for beacon1, dist1 := range scannerDistMap1.dist_map {
					// go through a hashmap of beacons and distances of a scanner to be compared to
					for beacon2, dist2 := range scannerDistMap2.dist_map {
						//fmt.Println("beacon2", dist2)

						/*if reflect.DeepEqual(dist1, dist2) {
							fmt.Println(beacon1, beacon2, dist1, dist2)
							count_equals++
						}*/
						equals := FindOverlappingBeacons(dist1, dist2)
						if equals > 0 {
							if equals >= 11 {
								if b1_val, ok := equalsMap[beacon1]; !ok {
									if b2_val, ok := equalsMap[beacon2]; !ok {
										//fmt.Println("scanner", i, "scanner", j, "equals", beacon1, beacon2, equals)
										equalsMap[beacon2] = beacon1
										transformedCoordinatesMap[beacon1] = 1
										SameBeacons = append(SameBeacons, BeaconPair{beacon_1: beacon1, beacon_2: beacon2})
										//count_equals++
										semi_count++
									} else {
										if b2_val != beacon1 {
											//fmt.Println("scanner", i, "scanner", j, "equals", beacon1, beacon2, equals)
											equalsMap[beacon2] = beacon1
											SameBeacons = append(SameBeacons, BeaconPair{beacon_1: beacon1, beacon_2: beacon2})
											//count_equals++
											semi_count++
										}
									}
								} else {
									if b1_val != beacon2 {
										if b2_val, ok := equalsMap[beacon2]; !ok {
											//fmt.Println("scanner", i, "scanner", j, "equals", beacon1, beacon2, equals)
											equalsMap[beacon2] = beacon1
											SameBeacons = append(SameBeacons, BeaconPair{beacon_1: beacon1, beacon_2: beacon2})
											//count_equals++
											semi_count++
										} else {
											if b2_val != beacon1 {
												//fmt.Println("scanner", i, "scanner", j, "equals", beacon1, beacon2, equals)
												equalsMap[beacon2] = beacon1
												SameBeacons = append(SameBeacons, BeaconPair{beacon_1: beacon1, beacon_2: beacon2})
												//count_equals++
												semi_count++
											}
										}
									}
								}
							}
						}
					}
				}

				if semi_count > 0 {
					fmt.Println("semi count", semi_count, "scanner", i, "scanner", j)
				}

				if semi_count >= 12 {
					count_equals += semi_count
					ScannerPairs = append(ScannerPairs, ScannerPair{scanner1_index: i, scanner2_index: j, scanner1_coordinate: SameBeacons[0].beacon_1, scanner2_coordinate: SameBeacons[0].beacon_2})
					fmt.Println(SameBeacons)
					//ChangeCoordinatesRelativelyToScanner0(ScannerReports[j])
				}
			}
		}
	}

	if count_equals > 0 {
		fmt.Println("Counted equals:", count_equals)
		fmt.Println(ScannerPairs)
		ConvertAllCoordinates(ScannerPairs)
	}
}

// 986
// 446 too low

func FindOverlappingBeacons(dist1 []float64, dist2 []float64) int {
	count_equals := 0

	for i := 0; i < len(dist1); i++ {
		for j := 0; j < len(dist2); j++ {
			if dist1[i] == dist2[j] {
				count_equals++
			}
		}
	}

	return count_equals
}

func ChangeCoordinatesRelativelyToScanner0(scanner_report ScannerReport) {
	fmt.Println("------")
	scanner := Beacon{x: 0, y: 0, z: 0}

	fmt.Println(scanner_report.report)

	for i, beacon := range scanner_report.report {
		beacon2, ok := equalsMap[beacon]

		fmt.Println(beacon, ok)

		if ok {
			scanner.x = beacon2.x + beacon.x
			scanner.y = beacon2.y - beacon.y
			scanner.z = beacon2.z + beacon.z

			scanner_report.report[i] = beacon2
		} else {
			scanner_report.report[i] = ConvertCoordinates(scanner, beacon)
			fmt.Println(beacon, ok, scanner, scanner_report.report[i])
		}
	}

	fmt.Println(scanner_report.report)
	fmt.Println("scanner", scanner)
}

func ConvertCoordinates(scanner Beacon, beacon Beacon) Beacon {
	beacon.x = beacon.x + scanner.x
	beacon.y = beacon.y + scanner.y
	beacon.z = beacon.z + scanner.z

	return beacon
}

func ConvertAllCoordinates(ScannerPairs []ScannerPair) {
	scanners := make(map[int]Beacon, len(FoundScanners))

	for {
		allScannersFound := false
		trueCounter := 0

		// check if all scanners are found
		for i := 0; i < len(FoundScanners); i++ {
			if FoundScanners[i] {
				trueCounter++
			}
		}

		allScannersFound = len(FoundScanners) == trueCounter
		fmt.Println(trueCounter, FoundScanners)

		if allScannersFound {
			//fmt.Println("breaking")
			break
		} else {
			// go through all scanner pairs and check if we can figure them out
			for _, scanner_pair := range ScannerPairs {
				if !(FoundScanners[scanner_pair.scanner1_index] && FoundScanners[scanner_pair.scanner2_index]) {
					//fmt.Println(scanner_pair.scanner1_index, scanner_pair.scanner2_index)

					// if we find a scanner that we know, we can calculate the other one
					if FoundScanners[scanner_pair.scanner1_index] {
						scanner_rel_0 := GetScannerCoordinatesRelativelyTo0(scanner_pair.scanner2_coordinate, scanner_pair.scanner1_coordinate)
						//fmt.Println(scanner_rel_0)

						FoundScanners[scanner_pair.scanner2_index] = true
						scanners[scanner_pair.scanner2_index] = scanner_rel_0
						//fmt.Println("Found", scanner_pair.scanner2_index)
					} else if FoundScanners[scanner_pair.scanner2_index] {
						scanner_rel_0 := GetScannerCoordinatesRelativelyTo0(scanner_pair.scanner2_coordinate, scanner_pair.scanner1_coordinate)
						//fmt.Println(scanner_rel_0)

						FoundScanners[scanner_pair.scanner1_index] = true
						scanners[scanner_pair.scanner1_index] = scanner_rel_0
						//fmt.Println("Found", scanner_pair.scanner1_index)
					}
				}
			}
		}
	}

	//fmt.Println(ScannerPairs)
	//fmt.Println(scanners)
	//fmt.Println(CoordinateRelativeTo0(Beacon{68, -1246, -43}, Beacon{x: 567, y: -361, z: 727}))
	scanner_rel_0 := GetScannerCoordinatesRelativelyTo0(Beacon{x: -258, y: -428, z: 682}, Beacon{x: 553, y: 889, z: -390})
	fmt.Println(CoordinateRelativeTo0(Beacon{68, -1246, -43}, Beacon{x: -759, y: -112, z: -32}))
	fmt.Println(scanner_rel_0)

	/*for i := 0; i < len(ScannerReports[1].report); i++ {
		fmt.Println(ScannerReports[1].report[i], CoordinateRelativeTo0(Beacon{68, -1246, -43}, ScannerReports[1].report[i]))
	}*/
}

func GetScannerCoordinatesRelativelyTo0(beacon1 Beacon, beacon2 Beacon) Beacon {
	scanner := Beacon{x: 0, y: 0, z: 0}

	scanner.x = beacon2.x + beacon1.x
	scanner.y = beacon2.y - beacon1.y
	scanner.z = beacon2.z + beacon1.z

	// {703 -491 -529} {-575 615 604} -> (128, 1106, 75)
	// {515 917 -361} {-743 427 -804}
	// 686,422,578   -618,-824,-621   ->  68, -1246, -43
	// 605,423,415   -537,-823,-458   ->  68, -1246, -43

	return scanner

}

func CoordinateRelativeTo0(scanner Beacon, beacon Beacon) Beacon {
	relative_beacon := Beacon{x: 0, y: 0, z: 0}

	relative_beacon.x = scanner.x - beacon.x
	relative_beacon.y = scanner.y + beacon.y
	relative_beacon.z = scanner.z - beacon.z
	/*
		relative_beacon.x = beacon.x + scanner.x
		relative_beacon.y = beacon.y + scanner.y
		relative_beacon.z = beacon.z + scanner.z
	*/

	return relative_beacon
}
