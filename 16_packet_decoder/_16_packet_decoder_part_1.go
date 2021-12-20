package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type part_rule struct {
	adjacent string
	inserted string
}

var sum int

func main() {
	ReadFile()
}

func ReadFile() {
	file, err := os.Open("input_test.txt")

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	var str string

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.ReplaceAll(line, " ", "")
		str = line
	}

	file.Close()

	sum = 0
	fmt.Println(str)

	binary := HexToBin(str)
	fmt.Println("init binary", binary)
	ParseString(binary)
}

func ParseString(binary string) {
	if len(binary) == 0 {
		return
	}

	fmt.Println("------------")

	ver, str_t := GetVersionAndType(binary)

	sum += int(ver)
	remainder := binary[6:]

	// literal or operator depending on the packet type
	if str_t == 4 {
		fmt.Println("literal")
		ParseTheLiteral(remainder)
	} else {
		fmt.Println("Operator")
		ParseTheOperator(remainder)
	}
}

func GetVersionAndType(binary string) (int, int) {
	str_version := binary[0:3]
	ver, _ := strconv.ParseInt(str_version, 2, 32)
	str_type := binary[3:6]
	str_t, _ := strconv.ParseInt(str_type, 2, 32)
	fmt.Println("ver", ver, "type", str_t)

	return int(ver), int(str_t)
}

func HexToBin(Hexadecimal string) string {
	var BCD = [16]string{
		"0000", "0001", "0010", "0011", "0100", "0101", "0110", "0111",
		"1000", "1001", "1010", "1011", "1100", "1101", "1110", "1111"}

	var Rst string

	for _, rune := range Hexadecimal {
		num, err := strconv.ParseUint(string(rune), 16, 32)

		if err != nil {
			fmt.Printf("%s", err)
		}

		Rst = Rst + BCD[num]
	}

	return Rst
}

func BinToHex(binary string) int {
	var BCD = [16]string{
		"0000", "0001", "0010", "0011", "0100", "0101", "0110", "0111",
		"1000", "1001", "1010", "1011", "1100", "1101", "1110", "1111"}

	for i := 0; i < len(BCD); i++ {
		if BCD[i] == binary {
			return i
		}
	}

	return -1
}

func ParseTheLiteral(remainder string) int {
	isFirstZero := false
	number_str := ""
	temp := ""
	last_i := -1

	for i, rune := range remainder {
		temp += string(rune)

		if i%5 == 4 {
			if string(temp[0]) == "1" || !isFirstZero && string(temp[0]) == "0" {
				temp = temp[1:]
				//num := BinToHex(temp)
				//fmt.Println(temp, num)
				number_str += temp
				temp = ""
				last_i = i
			}
		}
	}

	num, _ := strconv.ParseInt(number_str, 2, 32)
	fmt.Println(number_str, num)

	return last_i
}

func ParseTheOperator(remainder string) {
	packets_arr := make([]string, 0)
	lengthTypeID, _ := strconv.Atoi(string(remainder[0]))

	if lengthTypeID == 0 {
		// 15
		length_in_bits := remainder[1:16]
		length_num, _ := strconv.ParseInt(length_in_bits, 2, 32)
		packets_binary_str := remainder[16:]
		subpackets := ""

		if int(length_num) <= len(packets_binary_str) {
			subpackets = packets_binary_str[0:length_num]
		} else {
			subpackets = packets_binary_str
		}

		fmt.Println("subpackets", subpackets)

		for {
			_, str_t := GetVersionAndType(subpackets)

			if str_t == 4 {
				last_i := ParseTheLiteral(subpackets[6:])

				if last_i != -1 {
					packets_arr = append(packets_arr, subpackets[0:6+last_i])

					if 6+last_i != len(subpackets) {
						subpackets = subpackets[6+last_i:]
					}
				}
			}
			break
		}
	} else {
		// 11
		number_of_subpackets_binary_str := remainder[1:12]
		number_of_subpackets, _ := strconv.ParseInt(number_of_subpackets_binary_str, 2, 32)
		packets_binary_str := remainder[12:]
		packets_str_arr := make([]string, 0)

		// convert to
		temp := ""

		for _, rune := range packets_binary_str {
			temp += string(rune)

			if len(temp)%int(number_of_subpackets) == 0 {
				packets_str_arr = append(packets_str_arr, temp)
				temp = ""
			}
		}

		// reverse the array
		packets_str_arr = reverse(packets_str_arr)

		// check if there are leading zeros added to the array
		for _, str := range packets_str_arr {
			num, _ := strconv.Atoi(str)

			if num == 0 {
				packets_str_arr = packets_str_arr[1:]
			} else {
				break
			}
		}

		// reverse the array back
		packets_str_arr = reverse(packets_str_arr)

		// build the binary string back together
		packets_str := ""

		for _, str := range packets_str_arr {
			packets_str += str
		}

		packet_len := len(packets_str) / int(number_of_subpackets)
		fmt.Println("Packet len:", packet_len, "Num of subpackets:", number_of_subpackets)

		// break the binary into packets
		temp = ""

		for i, rune := range packets_str {
			if i != 0 && i%packet_len == 0 {
				packets_arr = append(packets_arr, temp)
				temp = ""
			}

			temp += string(rune)
		}

		packets_arr = append(packets_arr, temp)
	}

	//previous := math.Pow(2, math.Floor(math.Log(float64(length_num))/math.Log(2)))
	//diff := int(length_num) - int(previous)
	//fmt.Println(length, "Len:", length_num, "prevous:", previous, "ID:", lengthTypeID)

	fmt.Println("packets", packets_arr)

	if len(packets_arr) > 0 {
		for i := 0; i < len(packets_arr); i++ {
			ParseString(packets_arr[i])
		}
	}

	fmt.Println("sum", sum)
}

func reverse(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	return s
}
