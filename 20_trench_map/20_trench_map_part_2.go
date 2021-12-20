package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	ReadFile()
}

func ReadFile() {
	file, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	var image_enhancement_alg []string
	var input_image [][]string
	isInputImage := false

	extensionNum := 60
	emptyRow := make([]string, 100+extensionNum+extensionNum)

	for i := 0; i < len(emptyRow); i++ {
		emptyRow[i] = "."
	}

	emptyCells := make([]string, 0)

	for i := 0; i < extensionNum; i++ {
		emptyCells = append(emptyCells, ".")
	}

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) > 0 {
			if !isInputImage {
				image_enhancement_alg = append(image_enhancement_alg, strings.Split(line, "")...)
			} else {
				if len(input_image) == 0 {
					for i := 0; i < extensionNum; i++ {
						input_image = append(input_image, emptyRow)
					}
				}

				var row []string

				row = append(row, emptyCells...)
				row = append(row, strings.Split(line, "")...)
				row = append(row, emptyCells...)

				input_image = append(input_image, row)
			}
		} else {
			isInputImage = true
		}
	}

	for i := 0; i < extensionNum; i++ {
		input_image = append(input_image, emptyRow)
	}

	file.Close()

	fmt.Println(image_enhancement_alg, len(image_enhancement_alg))
	//PrintImage(input_image)
	input_image = FindPixelNeighbours(input_image, image_enhancement_alg, ".")

	for i := 0; i < 49; i++ {
		infinite_pixel := image_enhancement_alg[0]

		if image_enhancement_alg[0] == "#" {
			if i%2 == 1 {
				infinite_pixel = image_enhancement_alg[511]
			}
		}

		input_image = FindPixelNeighbours(input_image, image_enhancement_alg, infinite_pixel)
	}

	fmt.Println()
	PrintImage(input_image)
}

func FindPixelNeighbours(input_image [][]string, image_enhancement_alg []string, infinite_pixel string) [][]string {
	new_input_image := make([][]string, len(input_image))

	for i := 0; i < len(new_input_image); i++ {
		new_input_image[i] = make([]string, len(input_image[i]))
	}

	for y := 0; y < len(input_image); y++ {
		for x := 0; x < len(input_image[y]); x++ {
			all_pixels := make([]string, 0)

			// previous row
			if y-1 >= 0 {
				if x-1 >= 0 {
					all_pixels = append(all_pixels, input_image[y-1][x-1])
				} else {
					all_pixels = append(all_pixels, infinite_pixel)
				}

				all_pixels = append(all_pixels, input_image[y-1][x])

				if x+1 < len(input_image[y]) {
					all_pixels = append(all_pixels, input_image[y-1][x+1])
				} else {
					all_pixels = append(all_pixels, infinite_pixel)
				}
			} else {
				all_pixels = append(all_pixels, []string{infinite_pixel, infinite_pixel, infinite_pixel}...)
			}

			// this row
			if x-1 >= 0 {
				all_pixels = append(all_pixels, input_image[y][x-1])
			} else {
				all_pixels = append(all_pixels, infinite_pixel)
			}

			all_pixels = append(all_pixels, input_image[y][x])

			if x+1 < len(input_image[y]) {
				all_pixels = append(all_pixels, input_image[y][x+1])
			} else {
				all_pixels = append(all_pixels, infinite_pixel)
			}

			// next row
			if y+1 < len(input_image) {
				if x-1 >= 0 {
					all_pixels = append(all_pixels, input_image[y+1][x-1])
				} else {
					all_pixels = append(all_pixels, infinite_pixel)
				}

				all_pixels = append(all_pixels, input_image[y+1][x])

				if x+1 < len(input_image[y]) {
					all_pixels = append(all_pixels, input_image[y+1][x+1])
				} else {
					all_pixels = append(all_pixels, infinite_pixel)
				}
			} else {
				all_pixels = append(all_pixels, []string{infinite_pixel, infinite_pixel, infinite_pixel}...)
			}

			// create binary
			binary := ""

			for i := 0; i < len(all_pixels); i++ {
				if all_pixels[i] == "#" {
					binary += "1"
				} else {
					binary += "0"
				}
			}

			binary_num, err := strconv.ParseInt(binary, 2, 32)

			if err != nil {
				panic(err)
			}

			// write the new sign from the image enhancement algorithm
			new_input_image[y][x] = image_enhancement_alg[binary_num]
		}
	}

	return new_input_image
}

func PrintImage(input_image [][]string) {
	lit_counter := 0

	for y := 0; y < len(input_image); y++ {
		for x := 0; x < len(input_image[y]); x++ {
			//fmt.Print(input_image[y][x] + "")

			if input_image[y][x] == "#" {
				lit_counter++
			}
		}

		//fmt.Println()
	}

	fmt.Println(lit_counter)
}
