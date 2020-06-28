# proxycheck
Check proxy is operating correctly and the allow list is as expected by running parallel GET requests on multiple endpoints.
Log proxy URL for verbosity.

## Usage
```
$ go build
$ HTTPS_PROXY=my_local_proxy:80 ./proxycheck urls.txt
2020/06/28 22:32:47 Using URLs from urls.txt.
2020/06/28 22:32:47 WARNING: Proxy URL not found.
https://bbc.co.uk S (200)
https://google.com S (200)
https://amazon.co.uk S (200)
https://golang.org S (200)
https://amazon.com S (200)

Summary:
Success: 5
Error:   0
```
