package main

import (
	"bufio"
	"fmt"
	"github.com/gocolly/colly/v2"
	"os"
	"strconv"
	"time"
)

func main() {
	url := "https://www.meteorologia.gov.py/nivel-rio/vermas_convencional.php?code=2000086218&page=%s"
	collector := colly.NewCollector()
	file, err := os.Open("output.txt")
	if err != nil {
		return
	}
	defer file.Close()

	// whenever the collector is about to make a new request
	collector.OnRequest(func(r *colly.Request) {
		// print the url of that request
		fmt.Println("Visiting", r.URL)
	})
	collector.OnResponse(func(r *colly.Response) {
		fmt.Println("Got a response from", r.Request.URL)
	})
	collector.OnError(func(r *colly.Response, e error) {
		fmt.Println("Blimey, an error occurred!:", e)
	})
	//collector.OnHTML("#theDataTable", func(e *colly.HTMLElement) {
	//	fmt.Println(e)
	//})
	collector.OnHTML("tbody tr", func(e *colly.HTMLElement) {
		text := e.Text
		filePath := "./output.txt"
		fmt.Println("First column of a table row:", text)

		// Open file in append mode, create it if it doesn't exist
		file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			return
		}
		defer file.Close()

		// Create a writer for the file
		writer := bufio.NewWriter(file)

		// Write the new line to the file
		_, err = writer.WriteString(text + "\n")
		if err != nil {
			return
		}

		// Flush the writer buffer
		err = writer.Flush()
		if err != nil {
			return
		}

	})

	for i := 1; i < 1325; i++ {
		collector.Visit(fmt.Sprintf(url, strconv.Itoa(i)))
		time.Sleep(time.Second * 3)
	}

}
