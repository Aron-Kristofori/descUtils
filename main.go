package main

import (
	"fmt"
	"os"
	"bufio"
	s "strings"
)

//Error handler for file manipulation
func check(e error) {
	if e != nil {
		panic(e)
	}
}

//Function that strips extra spaces inside of a string
func strip_whitespace(text string) string {
	return s.Join(s.Fields(text), " ")
}

func contains(word string, keywords []string) bool {
	for i := 0; i < len(keywords); i++ {
		if word == keywords[i] {
			return true
		}
	}
	return false
}

func prepend(slice []string, word string) []string {
	temp := slice
	slice = nil
	slice = append(slice, word)
	for i := 0; i < len(temp); i++ {
		slice = append(slice, temp[i])
	}
	return slice
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 || args[0] == "-h" || args[0] == "--help" {
		fmt.Print("descUtils [<flags>...]\n\nFlags:\n-a=\"text\": append text\n-p=\"text\": prepend text\n-o=\"[PATH]\": path to output file\n-i=\"[PATH]\": " +
			"path to input\n-h,--help: help menu")
		return
	}
	//Input file search
	var inPath string = ""
	for i := 0; i < len(args); i++ {
		if args[i][0:3] == "-i=" {
			inPath = args[i][3:]
		}
	}
	if len(inPath) == 0 {
		fmt.Println("No input file detected! Please specify an input file using the -i=\"[PATH]\" flag")
		return
	}
	//Output file search
	var outPath string = ""
	for i := 0; i < len(args); i++ {
		if args[i][0:3] == "-o=" {
			outPath = args[i][3:]
		}
	}
	if len(outPath) == 0 {
		outPath = "out.txt"
	}
	//String search
	var modifier string = ""
	var mode string
	for i := 0; i < len(args); i++ {
		if args[i][0:3] == "-a=" || args[i][0:3] == "-p=" {
			modifier = args[i][3:]
			mode = args[i][1:2]
		}
	}
	if len(modifier) == 0 {
		fmt.Println("No modifier string has been specified! Please specify a string to modify the descirption using the -a or -p flag.")
		return
	}
	//Reading the file line by line
	file, err := os.Open(inPath)
	check(err)
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, strip_whitespace(scanner.Text()))
	}
	keywords := []string{"admin", "down", "up"}
	f, errr := os.Create(outPath)
	check(errr)
	for i := 0; i <  len(lines); i++ {
		words := s.Split(lines[i], " ")
		var desc []string
		var index int = len(words)-1
		for !contains(words[index], keywords) {
			desc = prepend(desc, words[index])
			index -= 1
		}
		if len(desc) > 0 {
			var description string = ""
			switch mode {
			case "a":
				for i := 0; i < len(desc); i++ {
					if i == len(desc)-1 {
						description += desc[i] + modifier
					} else {
						description += desc[i] + " "
					}
				}
				f.WriteString("int" + words[0] + "\n")
				f.WriteString("description " + description + "\n")
				case "p":
					description += modifier 
					for i := 0; i < len(desc); i++ {
						if i == len(desc)-1 {
							description += desc[i]
						} else {
							description += desc[i] + " "
						}
					}
					f.WriteString("int" + words[0] + "\n")
					f.WriteString("description " + description + "\n")
			}
		}
	}	
	check(scanner.Err())
}