# quickUrlChecker

### Is it any good?
For day to day use, to quickly check a newline seperated list of urls, 
or to test the urls via cmdline, its going to do what is needed. 

### What does it do for now?
Tells you if the website is up or not by sharing the http status code.
Welcome to change it into a full fledged curl replacement.

### How to build ?
```
go build main.go -o urlchecker
```

### How to invoke
```
Takes 2 params filepath and url.
--filepath : Path to the file that has the newline seperated list of urls.
--url comma seperated list of urls directly enetered in the command line, no spaces allowed.
```

### Test
```
time ./urlchecker --filepath tester/urls.txt
[200] https://github.com/ is up
[404] https://www.coursera.org/" is up
[200] https://stackoverflow.com/ is up
[200] https://golang.org/ is up
[200] https://www.linkedin.com/ is up
[200] http://medium.com/ is up
[200] https://www.udemy.com/ is up
[200] https://wesionary.team/ is up
Website null is down!

real    0m3.270s
user    0m0.119s
sys     0m0.076s

```
With cmdline options
```
time ./urlchecker --url https://www.google.com,https://o2.de,https://local.de
Website https://local.de is down!
[200] https://www.google.com is up
[200] https://o2.de is up

real    0m0.548s
user    0m0.083s
sys     0m0.045s

```