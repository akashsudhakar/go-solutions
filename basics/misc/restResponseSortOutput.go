package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
 * Complete the 'topArticles' function below.
 *
 * The function is expected to return a STRING_ARRAY.
 * The function accepts INTEGER limit as parameter.
 * base url for copy/paste:
 * https://jsonmock.hackerrank.com/api/articles?page=<pageNumber>
 */

func topArticles(limit int32) []string {
	// Make rest call get json response
	var titleDataValue []titleData

	titleDataValue = makeRestCall(1, titleDataValue)

	// Sort
	//sort.Sort(sortData(titleDataValue))

	sort.Slice(titleDataValue, func(i, j int) bool {
		if titleDataValue[i].count != titleDataValue[j].count {
			return titleDataValue[i].count > titleDataValue[j].count
		}
		return titleDataValue[i].title > titleDataValue[j].title
	})

	// Return top limit elements
	var out []string
	for i := 0; i < int(limit) && i < len(titleDataValue); i++ {
		out = append(out, titleDataValue[i].title)
	}
	return out
}

func makeRestCall(pageCount int, titleDataValue []titleData) []titleData {
	// Make rest call get json response
	var url string = "https://jsonmock.hackerrank.com/api/articles?page=" + fmt.Sprintf("%v", pageCount)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error while making rest call %v", err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	//bodyString := string(bodyBytes)
	//fmt.Println("API Response as String:\n" + bodyString)

	var todoStruct Response
	json.Unmarshal(bodyBytes, &todoStruct)
	//fmt.Printf("API Response as struct %+v\n", todoStruct)

	//fmt.Println(todoStruct)

	var numOfRestCalls int = todoStruct.Total_pages

	for _, dataValue := range todoStruct.Data {
		if len(dataValue.Title) > 0 {
			titleDataValue = append(titleDataValue, titleData{dataValue.Title, dataValue.Num_comments})
		} else if len(dataValue.Story_title) > 0 {
			titleDataValue = append(titleDataValue, titleData{dataValue.Story_title, dataValue.Num_comments})
		}
	}

	if pageCount == 1 {
		for j := 1; j <= numOfRestCalls; j++ {
			titleDataValue = makeRestCall(j+1, titleDataValue)
			fmt.Println(len(titleDataValue))
		}
	}

	return titleDataValue
}

type Response struct {
	Page        int    `json:"page"`
	Per_page    int    `json:"per_page"`
	Total       int    `json:"total"`
	Total_pages int    `json:"total_pages"`
	Data        []Data `json:"data"`
}

type Data struct {
	Title        string `json:"title"`
	Url          string `json:"url"`
	Author       string `json:"author"`
	Num_comments int    `json:"num_comments"`
	Story_id     string `json:"story_id"`
	Story_title  string `json:"story_title"`
	Story_url    string `json:"story_url"`
	Parent_id    string `json:"parent_id"`
	Created_at   string `json:"created_at"`
}

type titleData struct {
	title string
	count int
}

type sortData []titleData

func (s sortData) Len() int {
	return len(s)
}

func (s sortData) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortData) Less(i, j int) bool {
	if s[i].count != s[j].count {
		return s[i].count > s[j].count
	}
	return s[i].title > s[j].title
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	limitTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	limit := int32(limitTemp)

	result := topArticles(limit)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%s", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
