# bulkstatuscodechecker
A bulk status code checker built in go that accepts input from stdin and supports multithreading

####Usage:
cat urls.txt | bulkstatuscodechecker
(by default the program will use 20 threads)

  Concurrency:
  Use the flag -t to specify the number of threads that the program will use.
  Example: cat urls.txt | bulkstatuscodechecker -t 30 will use 30 threads

####Installation: 
go get github.com/NobleSiXSS/bulkstatuscodechecker
