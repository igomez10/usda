package main

import (
	"encoding/xml"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
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

	channelHTTPResponses := make(chan *http.Response, 100)
	doneHTTPResponses := make(chan bool)
	channelReports := make(chan *report, 20000)
	doneReports := make(chan bool)
	channelDates := make(chan string, 30)
	doneChannelDates := make(chan bool)

	/** This go routine listens on the chanel channelHTTPResponse and invokes the parseReport function with the item in the channelHTTPResponses `j`  and `channelReports`
	**/
	go func() {
		for {
			j, more := <-channelHTTPResponses
			if more {
				go parseReport(*j, channelReports)
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
				go postReportToDatabase(j, db)
			} else {
				println("COMPLETE REPORTS CHANNEL")
				doneReports <- true
				return
			}
		}
	}()

	go func() {
		for {
			j, more := <-channelDates
			if more {
				go getReport("ORANGES", j, channelHTTPResponses)
			} else {
				println("COMPLETE CHANNEL DATES")
				doneChannelDates <- true
				return
			}
		}
	}()

	for i := 0; i < *numDays; i++ {
		date := currentDate.AddDate(0, 0, -i)
		if date.Weekday() != 0 && date.Weekday() != 6 {
			var iterativeDate string = date.Format("01/02/2006")
			go func() {
				channelDates <- iterativeDate
			}()
		}
	}

	<-doneChannelDates
	close(channelDates)
	<-doneHTTPResponses
	close(channelHTTPResponses)
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

/**
getReport makes the http request with the provided fruit name `name` on the date `date` and sends the http response to the channel channelHTTPResponses `c`
**/
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

/**
parseReport receives the http response and parses it with an xml parser into a struct report r
After this, it sends the remaining struct r into the channel `c` which is `channelReports`
**/
func parseReport(input http.Response, c chan *report) {
	defer input.Body.Close()
	var r report
	decoder := xml.NewDecoder(input.Body)
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
