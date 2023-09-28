package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// This function reads the file and gives a string as output.
func readFile() string {
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}

// This function converts my string from above and makes it a slice of strings.
func getStringSlice(str string) []string {

	text := strings.Fields(str)
	return text

}

// I now range loop through slice of string until find "(cap)" when found go back to previous slice and capitalise the first letter of word. Return the new slice.
// Also dealt with capitalising the first letter of the specified words prior to command. Had to identify non closed bracket immediately after "cap", take the value given and work backwards to modify.
// I am left with a double space from the removal of "(cap)" which needs to be corrected later. I will use join joinSliceOfStrings to do this.
func titleCase(s []string) []string {
	for i, ch := range s {
		if ch == "(cap)" {
			s[i-1] = strings.Title(s[i-1])
			s[i] = ""

		} else if s[i] == "(cap," {

			num := ""
			for _, v := range s[i+1] {
				if v != 41 {
					num += string(v)
				}
			}
			val, err := strconv.Atoi(num)
			if err != nil {
				fmt.Printf("%v\n", err)
			}
			for j := i - val; j < i; j++ {
				s[j] = strings.Title(s[j])
			}
			s[i] = ""
			s[i+1] = ""
		}
	}
	return s
}

// range loop through slice of string until find "(up)" when found go back to previous slice and capitalise the entire word. Return the new slice.
// Also dealt with turning to upper case the specified words prior to command. Had to identify non closed bracket immediately after "up", take the value given and work backwards to modify.
// I am left with a double space from the removal of "(up)" which needs to be corrected later. I will use join joinSliceOfStrings to do this.
func toUpperCase(s []string) []string {
	for i, ch := range s {
		if ch == "(up)" {
			s[i-1] = strings.ToUpper(s[i-1])
			s[i] = ""

		} else if s[i] == "(up," {

			num := ""
			for _, v := range s[i+1] {
				if v != 41 {
					num += string(v)
				}
			}
			val, err := strconv.Atoi(num)
			if err != nil {
				fmt.Printf("%v\n", err)
			}
			for j := i - val; j < i; j++ {
				s[j] = strings.ToUpper(s[j])
			}
			s[i] = ""
			s[i+1] = ""
		}
	}
	return s
}

// I now range loop through slice of string until find "(low)" when found go back to previous slice and lower case the entire word. Return the new slice.
// Also dealt with turning to lower case the specified words prior to command. Had to identify non closed bracket immediately after "low", take the value given and work backwards to modify.
// I am left with a double space from the removal of "(low)" which needs to be corrected later. I will use join joinSliceOfStrings to do this.
func toLowerCase(s []string) []string {
	for i, ch := range s {
		if ch == "(low)" {
			s[i-1] = strings.ToLower(s[i-1])
			s[i] = ""

		} else if s[i] == "(low," {

			num := ""
			for _, v := range s[i+1] {
				if v != 41 {
					num += string(v)
				}
			}
			val, err := strconv.Atoi(num)
			if err != nil {
				fmt.Printf("%v\n", err)
			}
			for j := i - val; j < i; j++ {
				s[j] = strings.ToLower(s[j])
			}
			s[i] = ""
			s[i+1] = ""
		}
	}
	return s
}

// I range loop to find instances of "(hex)" and convert to decimal version of the hexadecimal value.
// I am left with a double space from the removal of "(hex)" which needs to be corrected later. I will use join joinSliceOfStrings to do this.
func hexTo(s []string) []string {
	for i, ch := range s {
		if ch == "(hex)" {

			hex_num := s[i-1]

			num, err := strconv.ParseInt(hex_num, 16, 32)
			if err != nil {
				panic(err)
			}

			stringNum := strconv.FormatInt(int64(num), 10)

			s[i-1] = stringNum
			s[i] = ""
		}
	}
	return s
}

// I range loop to find instances of "(bin)" and and convert to decimal version of the binary value.
// I am left with a double space from the removal of "(bin)" which needs to be corrected later. I will use join joinSliceOfStrings to do this.
func binTo(s []string) []string {
	for i, ch := range s {
		if ch == "(bin)" {

			bin_num := s[i-1]

			num, err := strconv.ParseInt(bin_num, 2, 32)
			if err != nil {
				panic(err)
			}

			stringNum := strconv.FormatInt(int64(num), 10)

			s[i-1] = stringNum
			s[i] = ""
		}
	}
	return s
}

// I originally wanted to remove double spaces but swapped to joining my slice of strings into one string which removed unecessary spaces.
func joinSliceOfStrings(s []string) string {
	res := strings.Join(s, " ")
	return res
}

// Deals with changing "a" to "an" in necessary occasions. Range loop through find "a" or "A", go to the next word and if that begins with a vowel or "h" then make the "a" and "an".
func changeA(s []string) []string {
	vowels := []string{"a", "e", "i", "o", "u", "h", "A", "E", "I", "O", "U", "H"}

	for i, word := range s {
		for _, letter := range vowels {
			if word == "a" && string(s[i+1][0]) == letter {
				s[i] = "an"
			} else if word == "A" && string(s[i+1][0]) == letter {
				s[i] = "An"
			}
		}
	}
	return s
}

// I need to range loop through until I find any instance of punctuation and then remove all spaces that occur prior. I then need to add a space after the punctuation.
func punctuations(s []string) []string {
	puncs := []string{",", ".", "!", "?", ":", ";"}

	// punc in the middle of a string connecting to word after
	for i, word := range s {
		for _, punc := range puncs {
			if string(word[0]) == punc && string(word[len(word)-1]) != punc {
				s[i-1] += punc
				s[i] = word[1:]
			}
		}
	}

	// punc at end of string
	for i, word := range s {
		for _, punc := range puncs {
			if (string(word[0]) == punc) && (s[len(s)-1] == s[i]) {
				s[i-1] += word
				s = s[:len(s)-1]
			}
		}
	}

	// punc in middle of string
	for i, word := range s {
		for _, punc := range puncs {
			if string(word[0]) == punc && string(word[len(word)-1]) == punc && s[i] != s[len(s)-1] {
				s[i-1] += word
				s = append(s[:i], s[i+1:]...)
			}
		}
	}

	// for apostrophe
	count := 0
	for i, word := range s {
		if word == "'" && count == 0 {
			count += 1
			s[i+1] = word + s[i+1]
			s = append(s[:i], s[i+1:]...)
		}

	}
	//  for second apostrophe
	for i, word := range s {
		if word == "'" {
			s[i-1] = s[i-1] + word
			s = append(s[:i], s[i+1:]...)
		}
	}
	return s
}

// writeToFile writes the completed string to a file called 'result.txt'
func writeToFile(s string) {
	err := os.WriteFile(os.Args[2], []byte(s), 0666)
	if err != nil {
		log.Fatalf("unable to write file: %v", err)
	}
}

func main() {
	s := readFile()
	sl := getStringSlice(s)
	sl = titleCase(sl)
	sl = toUpperCase(sl)
	sl = toLowerCase(sl)
	sl = hexTo(sl)
	sl = binTo(sl)
	s = joinSliceOfStrings(sl)
	sl = getStringSlice(s)
	sl = changeA(sl)
	sl = punctuations(sl)
	s = joinSliceOfStrings(sl)
	writeToFile(s)
	fmt.Println(s)

}
