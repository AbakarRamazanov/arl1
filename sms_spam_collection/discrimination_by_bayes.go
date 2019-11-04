package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Discrimination_by_Bayes start")
	fmt.Println("Use file: " + os.Args[1])
	textFile, _ := os.Open(os.Args[1])
	textByte, _ := ioutil.ReadAll(textFile)
	text := string(textByte)
	fmt.Println("Discrimination text: " + text)
	classesInString,
	countDocumentClasses,
	countWordInClass := loadFromDB(
		"classes.csv",
		"documentsInClasses.csv",
		"wordsInClasses.csv")
	discriminationTextByBayes(text, classesInString, countDocumentClasses, countWordInClass)
}


func discriminationTextByBayes(text string, classesInString []string, countDocumentClasses map[string] int, countWordInClass map[string] map[string] int){
	allDocuments := 0
	for _, count := range countDocumentClasses {
		allDocuments += count
	}

	logDocumentsCof := make(map[string]float64)
	for className, count := range countDocumentClasses {
		logDocumentsCof[className] = math.Log10(float64(count)/float64(allDocuments))
	}


	countAllWordsInClass := make(map[string]int)
	allWordList := make([]string, 0)
	for className, countWord := range countWordInClass {
		for word, count := range countWord {
			countAllWordsInClass[className] += count
			if !contains(allWordList,word) {
				allWordList = append(allWordList, word)
			}
		}
	}
	countAllWords := len(allWordList)

	r, _ := regexp.Compile("[^A-za-z]")
	words := strings.Split(r.ReplaceAllString(text, " "), " ")
	textWord := make(map[string]int)
	for _, word := range words {
		textWord[strings.ToLower(word)]++
	}

	var max float64
	flag := true
	maxClass := ""

	for _, className := range classesInString {
		var currentSum float64
		for word, count := range textWord {
			currentSum += math.Log10( float64(count) * (
				float64(countWordInClass[className][word]+1)/
					float64(countAllWordsInClass[className]+countAllWords) ) )
		}
		currentSum += logDocumentsCof[className]
		//currentSum *= -1
		if flag {
			max = currentSum
			maxClass = className
			flag = false
		}
		if currentSum > max {
			max = currentSum
			maxClass = string(className)
		}
	}
	fmt.Println("It's " + maxClass)
}

func loadFromDB(classesNameCSV,
	documentsNameCSV,
	wordsNameCSV string) (classesInString []string,
						countDocumentClasses map[string] int,
						countWordInClass map[string] map[string] int) {

	classesCSV, _ := os.Open(classesNameCSV)
	readerClassesCSV := csv.NewReader(bufio.NewReader(classesCSV))
	classesInString, _ = readerClassesCSV.Read()

	documentsCSV, _ := os.Open(documentsNameCSV)
	readerDocumentsCSV := csv.NewReader(bufio.NewReader(documentsCSV))
	documentsInString, _ := readerDocumentsCSV.ReadAll()
	countDocumentClasses = make(map[string] int)
	for _, line := range documentsInString {
		countDocumentClasses[line[0]], _ = strconv.Atoi(line[1])
	}

	wordsCSV, _ := os.Open(wordsNameCSV)
	readerWordsCSV := csv.NewReader(bufio.NewReader(wordsCSV))
	wordsInString, _ := readerWordsCSV.ReadAll()
	countWordInClass = make(map[string] map[string] int)
	for _, line := range wordsInString {
		if _, ok := countWordInClass[line[0]]; !ok {
			countWordInClass[line[0]] = make(map[string] int)
		}
		countWordInClass[line[0]][line[1]], _ = strconv.Atoi(line[2])
	}
	return
}

func contains(stringList []string, newItem string) bool {
	for _, n := range stringList {
		if newItem == n {
			return true
		}
	}
	return false
}