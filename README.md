
# Web Crawler using Go Language


## References:

    https://golang.org/
    https://golang.org/doc/
    https://github.com/jackdanger/collectlinks
    https://godoc.org/github.com/PuerkitoBio/goquery

## Downloads:

    https://golang.org/dl/
    https://golang.org/doc/install


## Install on Linux
	1)Use apt package installer
	
	sudo add-apt-repository ppa:longsleep/golang-backports
	sudo apt update
	sudo apt install golang-go

	OR , Simply use Snap Installer
	
	2)
	sudo snap install --classic go
## Configure & Test Go

	1. Set $GOPATH environment variable points to Go Project directory.
	2. Add $GOPATH/bin to $PATH
	
	NOTE: Can set using Bash Alias (run to reset in each Go Project Root)
	
	alias gopath='export GOPATH=$(pwd);export PATH=$PATH:$GOPATH/bin'
	
##  Installing 3rd party package (REQUIRED)
 	go get "github.com/jackdanger/collectlinks"
	
## Run crawl.go
	1)download the src folder where your go project directory
	2)open terminal inside the src/crawl/
	3)Run command :
	go run crawl.go http://rescale.com/
	


    
