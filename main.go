package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"

	"golang.org/x/net/html/charset"
)

func main() {
	currentDate := time.Now()
	var wg sync.WaitGroup
	wg.Add(30)
	for i := 0; i < 30; i++ {
		go func() {
			resp, err := getReport("ORANGES", currentDate.Format("01/02/2006"))
			if err != nil {
				log.Println(resp)
			}
			defer resp.Body.Close()
			aReport, err := parseReport(resp.Body)
			if err != nil {
				log.Println(err)
			}
			log.Printf("Number of registers: %d ", len(aReport.Items))
			log.Printf("Register 0 : %+v ", aReport.Items[0])
			currentDate = currentDate.AddDate(0, 0, -1)
			wg.Done()
		}()
	}
	wg.Wait()
}

func getReport(name, date string) (r *http.Response, err error) {
	var baseURL string = "https://www.marketnews.usda.gov/mnp/fv-report-top-filters?&navClass=FRUITS&commAbr=ORG&navType=byComm&repType=termPriceDaily&className=FRUITS&commName=%s&locName=&type=termPrice&repDate=%s&endDate=%s&format=xml&rebuild=false"
	r, err = http.Get(fmt.Sprintf(baseURL, name, date, date))
	return
}

func parseReport(input io.Reader) (r report, err error) {
	decoder := xml.NewDecoder(input)
	decoder.CharsetReader = charset.NewReaderLabel
	err = decoder.Decode(&r)
	return
}
