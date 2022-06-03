package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

/* Converting hexadecimal number to decimal number
Hexadecimal contains 16 digits (0-F) with a maximum range of storage capacity of 64 bits */
func HexToDecimal(s string) string {
	num, _ := strconv.ParseInt(s, 16, 64)
	return fmt.Sprint(num)
}

/* Converting binary number to decimal number
Binary is based on base 2 as it only contains zero or one digits, also known as bits
Maximum range of storage capacity is 64 bits */
func BinToDecimal(s string) string {
	value, _ := strconv.ParseInt(s, 2, 64)
	return fmt.Sprint(value)
}

// Converting strings in lower case to upper case and returns the value
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// Converting strings in upper case to lower case and returns to value
func ToLower(s string) string {
	return strings.ToLower(s)
}

// Converting the first letter of a string to capital
func Capitalize(s string) string {
	a := []rune(s)
	if a[0] >= 'a' && a[0] <= 'z' {
		a[0] = rune(a[0] - 32)
	}
	return string(a)
}

/* Iterates through the range of string arrays to move punctuation marks next to previous string without spaces
checks if the string has said rune e.g. (';')
then assigns fixed elements to a preceeding function which will then find the range minus 1 within the fixed elements */
func punctuation(strarr []string) []string {
	var newarr []string
	for _, str := range strarr {
		srune := []rune(str)
		if srune[0] == ';' || srune[0] == ':' || srune[0] == '!' || srune[0] == '?' || srune[0] == '.' || srune[0] == ',' {
			if len(srune) == 1 {
				newarr[len(newarr)-1] += string(srune[0])
			} else if len(srune) > 1 {
				newarr[len(newarr)-1] += string(srune[0])
				newarr = append(newarr, string(srune[1:]))
			}
		} else {
			newarr = append(newarr, str)
		}
	}
	return newarr
}

//
func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

// Iterates through the length of string to find quotation marks and removes the following space before the next word
func Quotes(puncEdit []string) []string {
	openQuote := true

	for i := 0; i < len(puncEdit); i++ {
		if puncEdit[i] == "'" {
			if openQuote {
				puncEdit[i+1] = "'" + puncEdit[i+1]
				puncEdit = remove(puncEdit, i)
				i--
				openQuote = false
			} else if !openQuote {
				puncEdit[i-1] = puncEdit[i-1] + "'"
				puncEdit = remove(puncEdit, i)
				i--
			}
		}
	}
	return puncEdit
}

// Iterates through the length of string arrays to change a to an if the next word begins with a vowel or h
func aToAn(strarr []string) []string {
	for index, str := range strarr {
		srune := []rune(strarr[index])
		if index+1 < len(strarr) {
			firstChar := rune(strarr[index+1][0])
			if str == "a" || str == "A" {
				if firstChar == 'a' || firstChar == 'e' || firstChar == 'i' || firstChar == 'o' || firstChar == 'u' || firstChar == 'h' {
					srune = append(srune, 'n')
					strarr[index] = string(srune)
				}
			}
		}
	}
	return strarr
}

// Used to compare two strings and return that that value as an integer
func Compare(a, b string) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return 1
}

/* Used to return a slice of string s with all leading and trailing unicode points contained removed
Then converts the given string s in a given base and bit size and returns it to an integer */
func trimAtoi(s string) int {
	srune := []rune(s)
	n := 0
	for _, rune := range string(srune) {
		y := 0
		if rune >= '0' && rune <= '9' {
			y = int(rune - '0')
			n = n*10 + y
		}
	}
	return n
}

/* Opens the first argument of a file, returns error if not found then closes
Also declaring and assigning values to variables of the string arrays */
func main() {
	var strarr []string
	var finalSlice []string
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	/* Used to read the file before passing on the data, quickly and efficiently
	Iterates through the range of fixed elements with an initial boolean of false
	checks if the string has said word e.g. (hex) and looks for the length of fixed elements minus 1
	then assigns fixed elements to a preceeding function which will then find the range minus 1 within the fixed elements */
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanWords)
	for fileScanner.Scan() {
		strarr = append(strarr, fileScanner.Text())
	}
	if err := fileScanner.Err(); err != nil {
		log.Fatal(err)
	}

	for i, str := range strarr {
		added := false
		if str == "(hex)" {
			finalSlice[len(finalSlice)-1] = HexToDecimal(strarr[i-1])
			added = true
		}
		if str == "(bin)" {
			finalSlice[len(finalSlice)-1] = BinToDecimal(strarr[i-1])
			added = true
		}
		if str == "(up)" {
			finalSlice[len(finalSlice)-1] = ToUpper(strarr[i-1])
			added = true
		}
		if str == "(low)" {
			finalSlice[len(finalSlice)-1] = ToLower(strarr[i-1])
			added = true
		}
		if str == "(cap)" {
			finalSlice[len(finalSlice)-1] = Capitalize(strarr[i-1])
			added = true
		}
		if !added {
			finalSlice = append(finalSlice, str)
		}
	}

	d := finalSlice

	/* iterates through the range of d which is the finalSlice and checks through the compare function
	if said "(low) etc" exists then check the full range and store the print value when the said word is present
	I think... */
	for i, val := range d {
		if Compare(val, "(low,") == 0 {
			numb := d[i+1]
			num := trimAtoi(numb)
			for j := 1; j <= num; j++ {
				d[i-j] = ToLower(d[i-j])
			}
			d = remove(d, i)
			d = remove(d, i)
		}
		if Compare(val, "(cap,") == 0 {
			numb := d[i+1]
			num := trimAtoi(numb)
			for j := 1; j <= num; j++ {
				d[i-j] = Capitalize(d[i-j])
			}
			d = remove(d, i)
			d = remove(d, i)

		}
		if Compare(val, "(up,") == 0 {
			numb := d[i+1]
			num := trimAtoi(numb)
			for j := 1; j <= num; j++ {
				d[i-j] = ToUpper(d[i-j])
			}
			d = remove(d, i)
			d = remove(d, i)
		}

	}
	d = punctuation(d)
	d = Quotes(d)
	d = aToAn(d)

	conv := strings.Join(d, " ")

	f, err := os.Create(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	io.WriteString(f, conv)
	f.Sync()
}
