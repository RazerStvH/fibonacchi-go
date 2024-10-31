# fibonacchi-go
A simple program for calculating an array of Fibonacci numbers and writing them to a file.

## Build
```bash
$ git clone https://github.com/RazerStvH/fibonacchi-go.git
$ cd fibonacchi-go
$ go build fibonacchi-go.go
```
If you want to install
```bash
$ sudo cp fibonachhi-go /usr/bin
```
## Usage
```bash
Usage of ./fibonacchi-go:
  -n int
    	The number up to which to calculate the Fibonacci sequence (default 10)
  -o string
    	The name of the output file to record the result (default "output.txt")
```
## Warning
Do not specify too large numbers because the output file may be too large.
