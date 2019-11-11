# Mongo Go Driver Sandbox

Quick POC implementation to debug to debug go driver behavior. Specifically, the current program is meant to show 
how different versions of the driver result in different behavior when executing the "listCommands" command 
against an auth-enabled mongod.

## Dependencies
- Mongodb v3.6.x
- Go 1.13+

## Usage
- `go build -o main`
- `./main -hostname <mongodb hostname> -port <mongodb port>`
- NOTE: default hostname/port will be `<os.Hostname()>:27500`

## Repro steps
- Start a mongod with version `3.6.x` with auth enabled
- Set the go driver version to `1.0.1` by running `go get go.mongodb.org/mongo-driver@v1.0.1`
- Build and run the POC against your mongod (see above)
- Observe a large printout of all the available commands on that mongod instance
- Set the go driver version to `1.1.3` by running `go get go.mongodb.org/mongo-driver@v1.1.3`
- Rebuild and rerun the POC against your mongod
- Observe the "no users authenticated error"
