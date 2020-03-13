
# Web Crawler using Go Language


## References:

    https://golang.org/
    https://golang.org/doc/
    https://github.com/jackdanger/collectlinks
    https://godoc.org/github.com/PuerkitoBio/goquery

## Downloads:

    https://golang.org/dl/
    https://golang.org/doc/install


## Install on Linux (STEP-1)
	1)Use apt package installer
```bash
sudo add-apt-repository ppa:longsleep/golang-backports
sudo apt update
sudo apt install golang-go
```
	OR , Simply use Snap Installer
`sudo snap install --classic go`
## Configure Go Path (STEP-2)

	1. Set $GOPATH environment variable points to Go Project directory.
	2. Add $GOPATH/bin to $PATH
	
	NOTE: Can set using Bash Alias (run to reset in each Go Project Root)
	set go path in linux using command : 
	
`alias gopath='export GOPATH=$(pwd);export PATH=$PATH:$GOPATH/bin'`
	
##  Installing 3rd party package (REQUIRED) (STEP-3)
`go get "github.com/jackdanger/collectlinks"`
	
## Git Clone the repo
```bash
cd go/src
git clone https://github.com/agrawalnaman/web-crawler-using-Golang.git
```
## Run crawl.go
```bash
cd web-crawler-using-Golang/
go run crawl.go http://rescale.com/
```


	2)open terminal inside the src/crawl/
	3)Run command :
	go run crawl.go http://rescale.com/
	


    
