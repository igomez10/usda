package main

import (
	"encoding/xml"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"golang.org/x/net/html/charset"
)

func main() {
	numDays := flag.Int("days", 30, "Number of days to retrieve")
	flag.Parse()
	currentDate := time.Now()
	db, err := gorm.Open("mysql", "root:mypassword@tcp(0.0.0.0:3306)/dev?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Println(err)
		log.Fatal("Not possible to connect to database")
	}
	// db.Table("oranges").CreateTable(&reportEntry{})
	defer db.Close()

	channelHTTPResponses := make(chan *http.Response, *numDays)
	doneHTTPResponses := make(chan bool)
	channelReports := make(chan *report, *numDays*200)
	doneReports := make(chan bool)

	go func() {
		for {
			j, more := <-channelHTTPResponses
			if more {
				parseReport((*j).Body, channelReports)
			} else {
				println("COMPLETE HTTPRESPONSES CHANNEL")
				doneHTTPResponses <- true
				return
			}
		}
	}()

	go func() {
		for {
			j, more := <-channelReports
			if more {
				postReportToDatabase(j, db)
			} else {
				println("COMPLETE REPORTS CHANNEL")
				doneReports <- true
				return
			}
		}
	}()

	var wg sync.WaitGroup
	wg.Add(*numDays)

	for i := 0; i < *numDays; i++ {
		go func(i int) {
			date := currentDate.AddDate(0, 0, -i)
			var iterativeDate string = date.Format("01/02/2006")
			getReport("ORANGES", iterativeDate, channelHTTPResponses)
			wg.Done()
		}(i)
	}
	wg.Wait()

	close(channelHTTPResponses)
	<-doneHTTPResponses
	close(channelReports)
	<-doneReports
}

func postReportToDatabase(r *report, db *gorm.DB) {
	for i := 0; i < len(r.Items); i++ {
		r.Items[i].ID = r.Items[i].getHash()
		// fmt.Printf("%+v", r.Items[i])
		db.Create(&r.Items[i])
	}
}

func getReport(name string, date string, c chan *http.Response) {
	fmt.Println("Extracted date: ", date)
	var baseURL string = "https://www.marketnews.usda.gov/mnp/fv-report-top-filters?&navClass=FRUITS&commAbr=ORG&navType=byComm&repType=termPriceDaily&className=FRUITS&commName=%s&locName=&type=termPrice&repDate=%s&endDate=%s&format=xml&rebuild=false"
	r, err := http.Get(fmt.Sprintf(baseURL, name, date, date))
	// defer r.Body.Close()
	if err == nil {
		c <- r
	} else {
		msg := fmt.Sprintf("Error in getReport, not possible to retrive report, Error: %s", err)
		err = errors.New(msg)
		log.Println(err)
	}
}

func parseReport(input io.Reader, c chan *report) {
	var r report
	decoder := xml.NewDecoder(input)
	decoder.CharsetReader = charset.NewReaderLabel
	err := decoder.Decode(&r)
	if err != nil {
		msg := fmt.Sprintf("Error in parseReport: %s  ", err)
		err = errors.New(msg)
		log.Println(err)
	} else {
		c <- &r
	}
	log.Printf("Finished Parsing: %+v \n", input)
}
