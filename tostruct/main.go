package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

var RAW_TYPES = [...]string{"CHAR", "BIN"}
var RAW_COL_SPLITTER = "\t"
var TAB_SIZE = 2

// Handle parent line position
type FormattedLine struct {
	vLeadingLevel int
	vName         string
	vIsField      bool
	vType         string
	vSize         int
	vQuantity     int
}

func (v *FormattedLine) String() string {
	return fmt.Sprintf("%v, %-50v, %5v, %5v, %2v, %2v", v.vLeadingLevel, v.vName, v.vIsField, v.vType, v.vSize, v.vQuantity)
}

// GetGoVariableName
func GetGoName(name string) string {
	// Remove extra space
	cleanName := removeExtraSpaces(name)
	// Replace space with _
	cleanName = strings.ReplaceAll(cleanName, " ", "_")
	// To lowercase
	cleanName = strings.ToLower(cleanName)
	// Remove all chars, keep [a-zA-Z0-9_]
	pattern := regexp.MustCompile("[^a-z0-9_]+")
	cleanName = pattern.ReplaceAllString(cleanName, "")
	return cleanName
}

func NextTabPos(str string) int {
	tabPos := strings.Index(str, RAW_COL_SPLITTER)
	if tabPos < 1 {
		tabPos = len(str)
	}
	return tabPos
}

func LeadingLevel(s string) int {
	re := regexp.MustCompile(fmt.Sprintf("^%s+", RAW_COL_SPLITTER))
	match := re.FindString(s)
	return len(match)
}

func TabString(nOfTabs int) string {
	s := ""
	for i := 0; i < nOfTabs; i++ {
		s += RAW_COL_SPLITTER
	}
	return s
}

func FullVariableInfo(str string) (FormattedLine, error) {
	var err error
	fLine := FormattedLine{}
	fLine.vIsField = false
	// get leading lvel
	fLine.vLeadingLevel = LeadingLevel(str)
	str = strings.TrimLeft(str, RAW_COL_SPLITTER)
	// get name
	splitPos := NextTabPos(str)
	fLine.vName = GetGoName(str[0:splitPos])
	if splitPos == len(str) {
		return fLine, nil
	}
	// get type
	str = str[splitPos+1:]
	splitPos = NextTabPos(str)
	fLine.vType = str[0:splitPos]
	// get isField
	fLine.vIsField = true
	if splitPos == len(str) {
		return fLine, nil
	}
	// get size
	str = str[splitPos+1:]
	splitPos = NextTabPos(str)
	fLine.vSize, _ = strconv.Atoi(str[0:splitPos])
	if splitPos == len(str) {
		return fLine, nil
	}
	// get quality
	str = str[splitPos+1:]
	fLine.vQuantity, err = strconv.Atoi(str[0:])
	return fLine, err
}

func ParseVariable(filepath string, nOfCols int) ([]FormattedLine, error) {
	nOfTabs := nOfCols - 1
	var formattedLines = []FormattedLine{}
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// Handle each line
	currentLine := 0
	for scanner.Scan() {
		aLine := scanner.Text()
		leadingLv := LeadingLevel(aLine)
		aLine = aLine[leadingLv:]
		// Struct case
		for strings.Count(aLine, RAW_COL_SPLITTER) > nOfTabs {
			fLine := FormattedLine{}
			fLine.vIsField = false
			fLine.vLeadingLevel = leadingLv
			fLine.vName = aLine[0:strings.Index(aLine, RAW_COL_SPLITTER)]
			aLine = aLine[strings.Index(aLine, RAW_COL_SPLITTER)+1:]
			formattedLines = append(formattedLines, fLine)
			leadingLv += 1
		}
		// Field case
		fLine, err := FullVariableInfo(fmt.Sprintf("%s%s", TabString(leadingLv), aLine))
		if err != nil {
			// Do warning
			fmt.Printf("[Warning] parse problem at %v, detail: %v\n", currentLine, err)
		}
		formattedLines = append(formattedLines, fLine)
		currentLine++
	}

	if err := scanner.Err(); err != nil {
		return formattedLines, err
	}
	return formattedLines, nil
}

// Removes all spaces from a string except for single spaces between words
func removeExtraSpaces(str string) string {
	var b strings.Builder
	var last rune

	for _, r := range str {
		if unicode.IsSpace(r) {
			if !unicode.IsSpace(last) {
				b.WriteRune(' ')
			}
		} else {
			b.WriteRune(r)
		}
		last = r
	}

	return strings.TrimSpace(b.String())
}

func main() {
	filepath := `D:\TrongTran\Job\Reactjs\golang\tostruct\config.yaml`
	data, _ := ParseVariable(filepath, 3)
	fmt.Println("Data has loaded as: [")
	for id, aLine := range data {
		fmt.Printf("%3v: %v\n", id+1, aLine.String())
	}
	fmt.Println("]")
}
