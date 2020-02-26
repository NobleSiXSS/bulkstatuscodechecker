# bulkstatuscodechecker
A bulk status code checker built in go that accepts input from stdin and supports multithreading

## Usage:
`cat urls.txt | bulkstatuscodechecker`
(by default the program will use 20 threads)

  #### Concurrency:
  Use the flag -t to specify the number of threads that the program will use.
  
  Example: `cat urls.txt | bulkstatuscodechecker -t 30` will use 30 threads and will be faster

## Installation: 
`go get -u github.com/NobleSiXSS/bulkstatuscodechecker`

## Example:

#### urls.txt
```
https://google.com
https://facebook.com
https://github.com
```
`cat urls.txt | bulkstatuscodechecker`

#### Output
```
200 https://github.com
301 https://google.com
301 https://facebook.com
```
