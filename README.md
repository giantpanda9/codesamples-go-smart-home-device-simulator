# codesamples-go-smart-home-device-simulator
Smart device source code for the https://github.com/giantpanda9/codesamples-nodejs-ts-smart-home-simulator no necessity to run it from here, unless you want to rebuild it, if you run the goSmart device from here, the Node.js part should exchange signals via post requests (it is an emulator after all) with this pseudo and the binary that Node.js system executes will be stopped, because port 8081 already busy.
# Installation
1) go1.18.linux-amd64.tar.gz from https://go.dev/
2) cd /path/to/downloaded/go1.18.linux-amd64.tar.gz
3) sudo tar -C /usr/local -xzf go1.18.linux-amd64.tar.gz
3-1) [optional] do sudo rm -rf /usr/local/go if you have other versions of Go installed - not needed for the first time installation
4) sudo nano ~/.profile
5) add to the end of file the following lines:
export GOPATH=$HOME/work
export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
6) cd /path/to/the/folder/of/the/cloned/project
7) go mod init goSmart - may not be needed, but should not hurt
8) go install 'github.com/gorilla/mux@latest' - may not be needed, but should not hurt
9) go run goSmart.go
OR 
10) go build goSmart.go
