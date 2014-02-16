# ntpclient
-------

ntpclient is a simple NTP client for go. It supports NTP version 4 and SNTP. Note that this client currently does not take network delay into account. This means the time that is received will not be accurate to the nanosecond.

## Installation
Make sure you have a working Go environment (go 1.1 is *required*). [See the install instructions](http://golang.org/doc/install.html).

To install ntpclient, run:
```
$ go get github.com/bt51/ntpclient
```

## Usage
Getting the current time from an ntp server is simple.

``` go
package main

import (
    "fmt"

    "github.com/bt51/ntpclient"
)

func main() {
    t, err := ntpclient.GetNetworkTime("0.pool.ntp.org", 123)
    if err != nil {
        fmt.Println(err)    
    }
    fmt.Println(t)
}
```
If the current time was returned from the ntp server, it will be converted to a time.Time type and returned. Note that the time will be in UTC so if you want it to be local, you will have to change the correct timezone.

License
-------

MIT
