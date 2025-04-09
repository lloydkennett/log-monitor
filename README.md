# Log Monitor
This project was built and tested using Go 1.21.6 on Ubuntu 22.04.5.

## Test
`go test .`

## Build
`go build .`

## Run
`./log-monitor`  

This defaults to read from `logs.log` and write results to `output.csv`  
You can provide command line arguments to specify input and output file  

`./log-monitor logs.log`  
`./log-monitor logs.log output.csv`

## Improvements
Possible improvements with more time:
- More validation on input and job matching such as checking description matches alongside the pid
- More test cases
- Introduce concurrency with goroutines to improve performance.
- Periodically check the input file for new updates