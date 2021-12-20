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

	emptyRow := make([]string, 100+3+3)

	for i := 0; i < len(emptyRow); i++ {
		emptyRow[i] = "."
	}

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) > 0 {
			if !isInputImage {
				image_enhancement_alg = append(image_enhancement_alg, strings.Split(line, "")...)
			} else {
				if len(input_image) == 0 {
					input_image = append(input_image, emptyRow)
					input_image = append(input_image, emptyRow)
					input_image = append(input_image, emptyRow)
				}

				var row []string
				row = append(row, []string{".", ".", "."}...)
				row = append(row, strings.Split(line, "")...)
				row = append(row, []string{".", ".", "."}...)

				//input_image = append(input_image, []string{".", ".", ".", strings.Split(line, "")..., ".", ".", "."})
				input_image = append(input_image, row)
			}
		} else {
			isInputImage = true
		}
	}

	input_image = append(input_image, emptyRow)
	input_image = append(input_image, emptyRow)
	input_image = append(input_image, emptyRow)

	file.Close()

	fmt.Println(image_enhancement_alg, len(image_enhancement_alg))
	//PrintImage(input_image)

	input_image = FindPixelNeighbours(input_image, image_enhancement_alg, true)
	fmt.Println()
	//PrintImage(input_image)
	input_image = FindPixelNeighbours(input_image, image_enhancement_alg, false)
	fmt.Println()
	PrintImage(input_image)
}

func FindPixelNeighbours(input_image [][]string, image_enhancement_alg []string, allEmpty bool) [][]string {
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
					if allEmpty {
						all_pixels = append(all_pixels, ".")
					} else {
						all_pixels = append(all_pixels, "#")
					}
				}

				all_pixels = append(all_pixels, input_image[y-1][x])

				if x+1 < len(input_image[y]) {
					all_pixels = append(all_pixels, input_image[y-1][x+1])
				} else {
					if allEmpty {
						all_pixels = append(all_pixels, ".")
					} else {
						all_pixels = append(all_pixels, "#")
					}
				}
			} else {
				if allEmpty {
					all_pixels = append(all_pixels, []string{".", ".", "."}...)
				} else {
					all_pixels = append(all_pixels, []string{"#", "#", "#"}...)
				}
			}

			// this row
			if x-1 >= 0 {
				all_pixels = append(all_pixels, input_image[y][x-1])
			} else {
				if allEmpty {
					all_pixels = append(all_pixels, ".")
				} else {
					all_pixels = append(all_pixels, "#")
				}
			}

			all_pixels = append(all_pixels, input_image[y][x])

			if x+1 < len(input_image[y]) {
				all_pixels = append(all_pixels, input_image[y][x+1])
			} else {
				if allEmpty {
					all_pixels = append(all_pixels, ".")
				} else {
					all_pixels = append(all_pixels, "#")
				}
			}

			// next row
			if y+1 < len(input_image) {
				if x-1 >= 0 {
					all_pixels = append(all_pixels, input_image[y+1][x-1])
				} else {
					if allEmpty {
						all_pixels = append(all_pixels, ".")
					} else {
						all_pixels = append(all_pixels, "#")
					}
				}

				all_pixels = append(all_pixels, input_image[y+1][x])

				if x+1 < len(input_image[y]) {
					all_pixels = append(all_pixels, input_image[y+1][x+1])
				} else {
					if allEmpty {
						all_pixels = append(all_pixels, ".")
					} else {
						all_pixels = append(all_pixels, "#")
					}
				}
			} else {
				if allEmpty {
					all_pixels = append(all_pixels, []string{".", ".", "."}...)
				} else {
					all_pixels = append(all_pixels, []string{"#", "#", "#"}...)
				}
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
			fmt.Print(input_image[y][x] + "")

			if input_image[y][x] == "#" {
				lit_counter++
			}
		}

		fmt.Println()
	}

	fmt.Println(lit_counter)
}
