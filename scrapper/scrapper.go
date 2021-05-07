package scrapper

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type extractedJob struct {
	id 		 string
	title	 string
	location string
	salary   string
	summary	 string
}

// Scrape Indeed by a term and location
func Scrape(term string, location string) {
	var baseURL string = strings.Replace("https://jp.indeed.com/jobs?q=" + term + "&l=" + location + "&limit=50", " ", "+", -1)
	fmt.Println(baseURL)
	var jobs []extractedJob
	c := make(chan []extractedJob)
	totalPages :=getPages(baseURL)

	for i:=0; i<totalPages; i++{
		go getPage(i, baseURL, c)
	}
	for i :=0; i<totalPages; i++ {
		extractedJobs := <- c // get extractedJobs as the message of channel
		jobs = append(jobs, extractedJobs...) // ... : get the content and merge (리스트의 리스트가 아닌, 리스트 하나로 merge
	}

	writeJobs(jobs)
	fmt.Println("Done.", "Extracted. No. of jobs :", len(jobs))
}

func getPage(page int, url string, mainC chan <-[]extractedJob) {
	var jobs []extractedJob

	c := make(chan extractedJob)
	pageURL := url + "&start=" + strconv.Itoa(page*50)
	fmt.Println("Requesting", pageURL)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close() // res.Body를 실행 후 종료 -> prevent memory leaks
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find(".jobsearch-SerpJobCard")
	searchCards.Each(func(i int, card *goquery.Selection) {
		go extractJob(card, c)
	})
	for i:=0; i < searchCards.Length(); i++ {
		job := <-c
		jobs = append(jobs, job)
	}

	mainC <- jobs
}

func extractJob(card *goquery.Selection, c chan <- extractedJob) {
	id, _ := card.Attr("data-jk")
	title := CleanString(card.Find(".title>a").Text())
	location := CleanString(card.Find(".sjcl").Text())
	salary := CleanString(card.Find(".salaryText").Text())
	summary := CleanString(card.Find(".summary").Text())

	c <- extractedJob{
		id: 	  id, 
		title:	  title, 
		location: location, 
		salary:   salary, 
		summary:  summary}
}

func getPages(url string) int {
	pages := 0
	res, err := http.Get(url)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close() // res.Body를 실행 후 종료 -> prevent memory leaks
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	// find pagination div and ~ to each divs
	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length()
	})
	return pages
}

func writeJobs(jobs []extractedJob) {
	// create jobs.csv using package os
	file, err := os.Create("jobs.csv")
	checkErr(err)
	// encoding
	utf8bom := []byte{0xEF, 0xBB, 0xBF}
	file.Write(utf8bom)

	// Create a writer which write on the file (Not standard writer)
	w := csv.NewWriter(file)
	defer w.Flush() // defer : 함수 안에서 마지막에 실행. Flush: clean the writer

	headers := []string{"Link", "Title", "Location", "Salary", "Summary"}

	wErr := w.Write(headers)
	checkErr(wErr)

	for _, job := range jobs {
		jobSlice := []string{"https://jp.indeed.com/viewjob?jk=" + job.id, job.title, job.location, job.salary, job.summary}
		jwErr := w.Write(jobSlice)
		checkErr(jwErr)
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status:", res.StatusCode)
	}
}

// CleanString cleans a string
func CleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}