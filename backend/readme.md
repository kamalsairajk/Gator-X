# Backend 

Commands and steps to run the project are as follows:

Firstly setup the environment:
## Setup environment

- Install [go version 1.17](https://go.dev/dl/go1.17.6.windows-amd64.msi)
- Add GOPATH env variable - this location can be anywhere, it'll store all the dependencies installed
- Add %USERPROFILE%\go\bin to the path env variable

After opening the webapp directory in the terminal, execute the commands in the order as follows:

### `go mod tidy`

This command installs the dependencies that are in used inside the project

### `go run main.go`

This command is used to start the webserver in debug mode and logs all requests and responses sent and received.

## File Structure

```webapp
|-views
|--reviewsview.go
|--placeview.go
|-model
|--base.go
|-main.go
```
