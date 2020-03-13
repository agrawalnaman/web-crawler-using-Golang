
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
alias gopath='export GOPATH=$(pwd);export PATH=$PATH:$GOPATH/bin'`
```
### If go already installed follow the steps below to clone the repo and run the script : 
	
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
```console
cd web-crawler-using-Golang/
go run crawl.go http://rescale.com/
```


	


    
