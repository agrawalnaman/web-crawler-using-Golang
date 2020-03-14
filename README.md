
# Web Crawler using Go Language


## References:

    https://golang.org/
    https://golang.org/doc/
    https://github.com/jackdanger/collectlinks
    https://godoc.org/github.com/PuerkitoBio/goquery

## Downloads:

    https://golang.org/dl/
    https://golang.org/doc/install


## Install Go on Linux (STEP-1)
	1)Use apt package installer
```console
sudo add-apt-repository ppa:longsleep/golang-backports
sudo apt update
sudo apt install golang-go
```
	OR , Simply use Snap Installer
`sudo snap install --classic go`
## Configure Go Path (Step-2)

	1. Set $GOPATH environment variable points to Go Project directory.
	2. Add $GOPATH/bin to $PATH
	
	NOTE: Can set using Bash Alias (run to reset in each Go Project Root)
	set go path in linux using command : 	
```console
alias gopath='export GOPATH=$(pwd);export PATH=$PATH:$GOPATH/bin'
```
# If Golang is already installed on your system and Go path is configured then follow the steps below to clone the repo and run the script in Linux console: 
	
##  Installing 3rd party package (Required dependency) (Step-3)
```console
go get "github.com/jackdanger/collectlinks"
```
	
## Git Clone the repo (Step-4)
```console
cd go/src
git clone https://github.com/agrawalnaman/web-crawler-using-Golang.git
```
## Run crawl.go (Step-5)
command for running the script:`go run crawl.go <Base URL>`
```console
cd web-crawler-using-Golang/
go run crawl.go http://rescale.com/
```

# Testing
	Crawled some small websites and manually checked for loops and infinite call backs
```golang
go run crawl.go https://www.crawler-test.com/links/repeated_internal_links
```
	and other websites too .. for example
```golang
go run crawl.go https://www.crawler-test.com/links/page_with_external_links
```
A good website for testing the crawler is 
```golang
go run crawl.go https://www.crawler-test.com/
```
Printed Timestamps after printing URLs of each page to show that multiple threads are simultaionsly parsing the URLs
# Output
### for test input `go run crawl.go http://rescale.com/`
<details><summary>Click here to view test run output snippet</summary>
<p>

	ubuntu@ip-10-0-0-136:~/go/src/web-crawler-using-Golang$ go run crawl.go http://rescale.com/
	[http://rescale.com/]
	Start time of crawl of this URL: 2020-03-13 22:07:37.810176103 +0000 UTC m=+0.000699272
	Stop time of crawl of this URL : 2020-03-13 22:07:38.578908454 +0000 UTC m=+0.769431501
	http://rescale.com/
		 http://rescale.com/
		 http://rescale.com/jp/
		 http://rescale.com/chs/
		 http://rescale.com/kr/
		 http://rescale.com/products/
		 http://rescale.com/products/enterprise/
		 http://rescale.com/products/advanced/
		 http://rescale.com/products/elements/
		 http://rescale.com/products/government/
		 http://rescale.com/products/universities/
		 http://rescale.com/products/developer/
		 http://rescale.com/features/
		 http://rescale.com/features/visualization/
		 http://rescale.com/security/
		 http://rescale.com/pricing/
		 http://rescale.com/infrastructure/
		 http://rescale.com/features/file-management/
		 http://rescale.com/features/admin-portal/
		 http://rescale.com/features/api/
		 http://rescale.com/features/workflow/
		 http://rescale.com/
		 http://rescale.com/aws/
		 http://rescale.com/azure/
		 http://rescale.com/gcp/
		 http://rescale.com/ibm/
		 http://rescale.com/ansys/
		 http://rescale.com/avl/
		 http://rescale.com/siemens/
		 http://rescale.com/solutions/
		 http://rescale.com/solutions/engineers-and-scientists/
		 http://rescale.com/solutions/cxos/
		 http://rescale.com/solutions/hpc-management/
		 http://rescale.com/solutions/academia/
		 http://rescale.com/solutions/startups/
		 http://rescale.com/solutions/aerospace/
		 http://rescale.com/solutions/automotive/
		 http://rescale.com/solutions/consumer-products/
		 http://rescale.com/solutions/eda-and-electronics/
		 http://rescale.com/solutions/oil-and-gas/
		 http://rescale.com/solutions/life-sciences/
		 http://rescale.com/solutions/autonomous-driving/
		 http://rescale.com/solutions/automation-api-cli/
		 http://rescale.com/solutions/cae-mdo/
		 http://rescale.com/solutions/digital-twin/
		 http://rescale.com/solutions/driver-assistance-adas/
		 http://rescale.com/solutions/disaster-recovery/
		 http://rescale.com/solutions/iot-big-data/
		 http://rescale.com/solutions/machine-learning/
		 http://rescale.com/solutions/spdm/
		 http://rescale.com/storage/
		 http://rescale.com/software/
		 http://rescale.com/partners/
		 http://rescale.com/partners-list/
		 https://resources.rescale.com/
		 https://docs.rescale.com/
		 https://resources.rescale.com/?wpv-resource-type=white-paper
		 http://rescale.com/about/
		 http://rescale.com/investors/
		 https://resources.rescale.com//blog
		 https://resources.rescale.com//news
		 https://resources.rescale.com//events
		 http://rescale.com/jobs/
		 http://rescale.com/legal/
		 http://rescale.com/signup/
		 http://rescale.com/login/
		 https://info.rescale.com/case-studies/nissan
		 https://info.rescale.com/white-papers/cloud-3.0-the-rise-of-big-compute
		 https://info.rescale.com/case-studies/dinex-reduces-time-to-market-of-exhaust-systems-by-25-percent
		 https://www.youtube.com/watch?v=05HfJ8dZJXE
		 https://info.rescale.com/case-studies/boom-supersonic
		 https://www.youtube.com/watch?v=umiGy7fe5zc
		 https://www.youtube.com/watch?v=h1nsUGuklHw
		 https://www.youtube.com/watch?v=tPaq3Hmeg5Y
		 https://resources.rescale.com/?wpv-resource-type=video
		 https://resources.rescale.com/resource/a3-project-vahana-rescale-power-personal-flight/
		 https://resources.rescale.com/boom-technology-leverages-rescale-platform-to-enable-a-rebirth-of-supersonic-passenger-travel/
		 https://resources.rescale.com/resource/the-need-for-speed-drives-nascars-richard-childress-racing-to-the-cloud/
		 https://support.rescale.com/customer/en/portal/articles/2778993-trek-bicycle-uses-rescale-to-run-cutting-edge-coupled-optimization-analysis
		 http://rescale.com/booking/
		 https://resources.rescale.com/events/
		 https://resources.rescale.com/news/
		 https://resources.rescale.com/rescale-enables-faster-time-to-market-for-nissan/
		 https://resources.rescale.com/announcements/rescale-announces-strategic-partnership-offering-with-siemens-plm/
		 https://resources.rescale.com/announcements/rescale-announces-innovations-to-accelerate-time-to-results/
		 https://resources.rescale.com/announcements/rescale-receives-2018-hpcwire-editors-choice-award-for-best-hpc-in-the-cloud-platform/
		 https://resources.rescale.com/blog
		 http://info.rescale.com/contact_sales
		 https://www.linkedin.com/company/rescale/
		 https://twitter.com/rescaleinc
		 https://www.facebook.com/rescaleinc/
	Start time of crawl of this URL: 2020-03-13 22:07:38.586394942 +0000 UTC m=+0.776918083
	Stop time of crawl of this URL : 2020-03-13 22:07:39.315722492 +0000 UTC m=+1.506245536
	http://rescale.com/products/
		 http://rescale.com/
		 http://rescale.com/products/
		 http://rescale.com/jp/products/
		 http://rescale.com/chs/products/
		 http://rescale.com/kr/products/
		 http://rescale.com/products/enterprise/
		 http://rescale.com/products/advanced/
		 http://rescale.com/products/elements/
		 http://rescale.com/products/government/
		 http://rescale.com/products/universities/
		 http://rescale.com/products/developer/
		 http://rescale.com/features/
		 http://rescale.com/features/visualization/
		 http://rescale.com/security/
		 http://rescale.com/pricing/
		 http://rescale.com/infrastructure/
		 http://rescale.com/features/file-management/
		 http://rescale.com/features/admin-portal/
		 http://rescale.com/features/api/
		 http://rescale.com/features/workflow/
		 http://rescale.com/products/
		 http://rescale.com/aws/
		 http://rescale.com/azure/
		 http://rescale.com/gcp/
		 http://rescale.com/ibm/
		 http://rescale.com/ansys/
		 http://rescale.com/avl/
		 http://rescale.com/siemens/
		 http://rescale.com/solutions/
		 http://rescale.com/solutions/engineers-and-scientists/
		 http://rescale.com/solutions/cxos/
		 http://rescale.com/solutions/hpc-management/
		 http://rescale.com/solutions/academia/
		 http://rescale.com/solutions/startups/
		 http://rescale.com/solutions/aerospace/
		 http://rescale.com/solutions/automotive/
		 http://rescale.com/solutions/consumer-products/
		 http://rescale.com/solutions/eda-and-electronics/
		 http://rescale.com/solutions/oil-and-gas/
		 http://rescale.com/solutions/life-sciences/
		 http://rescale.com/solutions/autonomous-driving/
		 http://rescale.com/solutions/automation-api-cli/
		 http://rescale.com/solutions/cae-mdo/
		 http://rescale.com/solutions/digital-twin/
		 http://rescale.com/solutions/driver-assistance-adas/
		 http://rescale.com/solutions/disaster-recovery/
		 http://rescale.com/solutions/iot-big-data/
		 http://rescale.com/solutions/machine-learning/
		 http://rescale.com/solutions/spdm/
		 http://rescale.com/storage/
		 http://rescale.com/software/
		 http://rescale.com/partners/
		 http://rescale.com/partners-list/
		 https://resources.rescale.com/
		 https://docs.rescale.com/
		 https://resources.rescale.com/?wpv-resource-type=white-paper
		 http://rescale.com/about/
		 http://rescale.com/investors/
		 https://resources.rescale.com//blog
		 https://resources.rescale.com//news
		 https://resources.rescale.com//events
		 http://rescale.com/jobs/
		 http://rescale.com/legal/
		 http://rescale.com/signup/
		 http://rescale.com/login/
		 http://rescale.com/products/pro/
		 http://rescale.com/booking/
		 http://info.rescale.com/contact_sales
		 https://resources.rescale.com/events/
		 https://resources.rescale.com/news/
		 https://resources.rescale.com/rescale-enables-faster-time-to-market-for-nissan/
		 https://resources.rescale.com/announcements/rescale-announces-strategic-partnership-offering-with-siemens-plm/
		 https://resources.rescale.com/announcements/rescale-announces-innovations-to-accelerate-time-to-results/
		 https://resources.rescale.com/announcements/rescale-receives-2018-hpcwire-editors-choice-award-for-best-hpc-in-the-cloud-platform/
		 https://resources.rescale.com/blog
		 https://www.linkedin.com/company/rescale/
		 https://twitter.com/rescaleinc
		 https://www.facebook.com/rescaleinc/
	Start time of crawl of this URL: 2020-03-13 22:07:38.585727075 +0000 UTC m=+0.776250248
	Stop time of crawl of this URL : 2020-03-13 22:07:39.326520789 +0000 UTC m=+1.517043764
	http://rescale.com/chs/
		 http://rescale.com/chs/
		 http://rescale.com/
		 http://rescale.com/jp/
		 http://rescale.com/kr/
		 http://rescale.com/chs/products/
		 http://rescale.com/chs/products/enterprise/
		 http://rescale.com/chs/products/advanced/
		 http://rescale.com/chs/products/elements/
		 http://rescale.com/chs/products/government/
		 http://rescale.com/chs/products/universities/
		 http://rescale.com/chs/products/developer/
		 http://rescale.com/chs/features/
		 http://rescale.com/chs/features/visualization/
		 http://rescale.com/chs/security/
		 http://rescale.com/chs/pricing/
		 http://rescale.com/chs/infrastructure/
		 http://rescale.com/chs/features/file-management/
		 http://rescale.com/chs/features/admin-portal/
		 http://rescale.com/chs/features/api/
		 http://rescale.com/chs/features/workflow/
		 http://rescale.com/chs/
		 http://rescale.com/chs/aws/
		 http://rescale.com/chs/azure/
		 http://rescale.com/chs/gcp/
		 http://rescale.com/chs/ibm/
		 http://rescale.com/chs/ansys/
		 http://rescale.com/chs/avl/
		 http://rescale.com/chs/siemens/
		 http://rescale.com/chs/solutions/
		 http://rescale.com/chs/solutions/engineers-and-scientists/
		 http://rescale.com/chs/solutions/cxos/
		 http://rescale.com/chs/solutions/hpc-management/
		 http://rescale.com/chs/solutions/academia/
		 http://rescale.com/chs/solutions/startups/
		 http://rescale.com/chs/solutions/aerospace/
		 http://rescale.com/chs/solutions/automotive/
		 http://rescale.com/chs/solutions/consumer-products/
		 http://rescale.com/chs/solutions/eda-and-electronics/
		 http://rescale.com/chs/solutions/oil-and-gas/
		 http://rescale.com/chs/solutions/life-sciences/
		 http://rescale.com/chs/solutions/autonomous-driving/
		 http://rescale.com/chs/solutions/automation-api-cli/
		 http://rescale.com/chs/solutions/cae-mdo/
		 http://rescale.com/chs/solutions/digital-twin/
		 http://rescale.com/chs/solutions/driver-assistance-adas/
		 http://rescale.com/chs/solutions/disaster-recovery/
		 http://rescale.com/chs/solutions/iot-big-data/
		 http://rescale.com/chs/solutions/machine-learning/
		 http://rescale.com/chs/solutions/spdm/
		 http://rescale.com/chs/storage/
		 http://rescale.com/chs/software/
		 http://rescale.com/chs/partners/
		 http://rescale.com/chs/partners-list/
		 https://resources.rescale.com/
		 https://docs.rescale.com/
		 https://resources.rescale.com/?wpv-resource-type=white-paper
		 http://rescale.com/chs/about/
		 http://rescale.com/chs/investors/
		 https://resources.rescale.com//blog
		 https://resources.rescale.com//news
		 https://resources.rescale.com//events
		 http://rescale.com/chs/jobs/
		 http://rescale.com/chs/legal/
		 http://rescale.com/signup/
		 http://rescale.com/login/
		 https://info.rescale.com/case-studies/nissan
		 https://info.rescale.com/white-papers/cloud-3.0-the-rise-of-big-compute
		 https://info.rescale.com/case-studies/dinex-reduces-time-to-market-of-exhaust-systems-by-25-percent
		 https://www.youtube.com/watch?v=05HfJ8dZJXE
		 https://info.rescale.com/case-studies/boom-supersonic
		 https://www.youtube.com/watch?v=umiGy7fe5zc
		 https://www.youtube.com/watch?v=h1nsUGuklHw
		 https://www.youtube.com/watch?v=tPaq3Hmeg5Y
		 https://resources.rescale.com/?wpv-resource-type=video
		 https://resources.rescale.com/resource/a3-project-vahana-rescale-power-personal-flight/
		 https://resources.rescale.com/boom-technology-leverages-rescale-platform-to-enable-a-rebirth-of-supersonic-passenger-travel/
		 https://resources.rescale.com/resource/the-need-for-speed-drives-nascars-richard-childress-racing-to-the-cloud/
		 https://support.rescale.com/customer/en/portal/articles/2778993-trek-bicycle-uses-rescale-to-run-cutting-edge-coupled-optimization-analysis
		 http://rescale.com/chs/booking/
		 https://resources.rescale.com/events/
		 https://resources.rescale.com/blog
		 http://info.rescale.com/contact_sales
		 https://www.linkedin.com/company/rescale/
		 https://twitter.com/rescaleinc
		 https://www.facebook.com/rescaleinc/
	Start time of crawl of this URL: 2020-03-13 22:07:38.5852948 +0000 UTC m=+0.775817962
	Stop time of crawl of this URL : 2020-03-13 22:07:39.337943349 +0000 UTC m=+1.528466402
	http://rescale.com/jp/
		 http://rescale.com/jp/
		 http://rescale.com/
		 http://rescale.com/chs/
		 http://rescale.com/kr/
		 http://rescale.com/jp/products/
		 http://rescale.com/jp/products/enterprise/
		 http://rescale.com/jp/products/advanced/
		 http://rescale.com/jp/products/elements/
		 http://rescale.com/jp/products/government/
		 http://rescale.com/jp/products/universities/
		 http://rescale.com/jp/products/developer/
		 http://rescale.com/jp/features/
		 http://rescale.com/jp/features/visualization/
		 http://rescale.com/jp/security/
		 http://rescale.com/jp/pricing/
		 http://rescale.com/jp/infrastructure/
		 http://rescale.com/jp/features/file-management/
		 http://rescale.com/jp/features/admin-portal/
		 http://rescale.com/jp/features/api/
		 http://rescale.com/jp/features/workflow/
		 http://rescale.com/jp/
		 http://rescale.com/jp/aws/
		 http://rescale.com/jp/azure/
		 http://rescale.com/jp/gcp/
		 http://rescale.com/jp/ibm/
		 http://rescale.com/jp/ansys/
		 http://rescale.com/jp/avl/
		 http://rescale.com/jp/siemens/
		 http://rescale.com/jp/solutions/
		 http://rescale.com/jp/solutions/engineers-and-scientists/
		 http://rescale.com/jp/solutions/cxos/
</p>
</details>
	


    
