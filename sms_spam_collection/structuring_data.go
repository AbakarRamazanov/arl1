package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"regexp"
	"strings"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Structuring_data start")
	fmt.Println("Use file: " + os.Args[1])
	csvFile, _ := os.Open(os.Args[1])
	readerCSV := csv.NewReader(bufio.NewReader(csvFile))
	lines, _ := readerCSV.ReadAll()
	texts := make(map[string] []string)
	//lines = lines[1:]
	for _, line := range lines {
		texts[line[0]] = append(texts[line[0]], line[1])
	}
	insertToDB(texts)
	}

func insertToDB(texts map[string] []string){
	r, _ := regexp.Compile("[^A-za-z]")
	classesInString := make([]string, 0)
	countDocumentClassesInString := make(map[string] int)
	countWordInClass := make(map[string] map[string] int)
	for nameClass, lines := range texts {
		classesInString = append(classesInString, nameClass)
		countDocumentClassesInString[nameClass] = len(lines)
		countWordInClass[nameClass] = make(map[string] int)
		for _, line := range lines {
			words := strings.Split(r.ReplaceAllString(line, " "), " ")
			for _, word := range words {
				countWordInClass[nameClass][strings.ToLower(word)]++
			}
		}
	}
	for i := 0; i < len(countWordInClass); i++ {
		if _, ok := countWordInClass[classesInString[i]][""]; ok {
			delete(countWordInClass[classesInString[i]],"")
		}
	}

	classesCsv, _ := os.Create("classes.csv")
	writerClasses := csv.NewWriter(classesCsv)
	writerClasses.Write(classesInString)
	writerClasses.Flush()
	classesCsv.Close()

	documentsCsv, _ := os.Create("documentsInClasses.csv")
	writerDocument := csv.NewWriter(documentsCsv)
	for key, value := range countDocumentClassesInString {
		line := []string{key, strconv.Itoa(value)}
		writerDocument.Write(line)
		writerDocument.Flush()
	}
	documentsCsv.Close()

	wordsCsv, _ := os.Create("wordsInClasses.csv")
	writerWord := csv.NewWriter(wordsCsv)
	for class, words := range countWordInClass {
		for word, count := range words {
			line := []string{class, word, strconv.Itoa(count)}
			writerWord.Write(line)
			writerWord.Flush()
		}
	}
	wordsCsv.Close()
	//(map[string] map[string] int)
}