/*@ NAMAN AGRAWAL
Description : It is a script which crawls a web page and puts all the unique URLs into a queue which is then crawled to fetch
all the url on that page. It is multithreaded i.e. goroutines are used to speed up the crawling and one slow page will not 
block the crawling process. Used a mutex lock to print Urls of one thread at a time.
output:  the page being crawled is printed without a tab space and all urls on that page are printed with a tab seperation.
printed start and stop time of each page being crawled to check they are being crawled concurrently.
*/
package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"github.com/jackdanger/collectlinks"
	"net/http"
	"net/url"
	"os"
	"sync"
	"time"
)

var mu = &sync.Mutex{}

func main() {
	flag.Parse()
	// Using Flag Module to get the command line arguments from user.
	args := flag.Args()
	fmt.Println(args)
	if len(args) < 1 {
		fmt.Println("Please specify base URL to crawl ")
		os.Exit(1)
	}
	 _, err := url.ParseRequestURI(args[0])
	 if err != nil {
	 	fmt.Println("Please enter a valid URL to crawl")
		os.Exit(2)
	}


	queue := make(chan string)
	filteredQueue := make(chan string)

	go func() { queue <- args[0] }()
	go filterQueue(queue, filteredQueue)

	// introduce a bool channel to synchronize execution of concurrently running crawlers
	done := make(chan bool)

	// pull from the filtered queue, add to the unfiltered queue
	// Creating 5 threads which will be running parallelly
	for i := 0; i < 5; i++ {
		go func() {
			for uri := range filteredQueue {
				addToQueue(uri, queue)
			}
			done <- true
		}()

	}
	<-done
}

func filterQueue(in chan string, out chan string) {
	//putting only unseen URLs in the queue i.e. filtering to get only unique URLs.
	// Avoiding looping to same URLs.
	var seen = make(map[string]bool)
	for val := range in {
		if !seen[val] {
			seen[val] = true
			out <- val
		}
	}
}


//Adding URLs found on a page which are unique and need to be crawled.
func addToQueue(uri string, queue chan string) {

	start := time.Now()

	//Avoiding and Ignorig URLs which require SSL certificates and other validations.
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	//Opening Client connection
	client := http.Client{Transport: transport}
	resp, err := client.Get(uri)
	if err != nil {
		return
	}
	//Closing the connection
	defer resp.Body.Close()

	//Extracting all links on given page using a custom public module package from Github
	links := collectlinks.All(resp.Body)
	foundUrls := []string{}
	for _, link := range links {
		//Fixing the URLs with relative or broken links before adding to the queue.
		absolute := cleanUrl(link, uri)
		foundUrls = append(foundUrls, absolute)       
		if uri != "" {
			go func() { queue <- absolute }()
		}
	}
	stop := time.Now()
	display(uri,foundUrls,start,stop)
}

//Display all the URLs found on a page and print them on console together.
//Multiple Threads want to print on console at the same time so we use Mutex to lock this function to be accessed by one thread at a time.
func display(uri string,found []string,start time.Time,stop time.Time){
 	mu.Lock()
	fmt.Println("Start time of crawl of this URL:",start)
	fmt.Println("Stop time of crawl of this URL :",stop)
	fmt.Println(uri)

	//Filtering Phone numbers and Email ids which are missed out by url.Parse() and only printing valid URLs.
	for _,str := range found{
		str, err := url.Parse(str)
		if err== nil{
			if str.Scheme =="http" || str.Scheme == "https"{
				fmt.Println("\t", str)
			}
		}
	}
	mu.Unlock()
}



//Check if broken URL or relative link and resolve it using default Golang Library
func cleanUrl(href, base string) string {
	uri, err := url.Parse(href)
	if err != nil {
		return ""
	}
	baseUrl, err := url.Parse(base)
	if err != nil {
		return ""
	}
	uri = baseUrl.ResolveReference(uri)
	return uri.String()
}
