package main

import (
    "fmt"
    "strconv"
    "strings"
    "bytes"
)

type UriParts struct{
    uri string 
    hostname string 
    pathname string 
    port int
}

func filterNewLines(s string) string {
    return strings.Map(func(r rune) rune {
        switch r {
        case 0x000A, 0x000B, 0x000C, 0x000D, 0x0085, 0x2028, 0x2029:
            return -1
        default:
            return r
        }
    }, s)
}

func (uri *UriParts) parseURI() bool{
    if !strings.HasPrefix(uri.uri, "http://") && !strings.HasPrefix(uri.uri, "https://"){
        fmt.Println(uri.uri)
        uri.hostname = "";
        return false;
    }
    var testsplit =  strings.Split(uri.uri, "/")
    if strings.Contains(testsplit[2],":") {
        var hostsplit = strings.Split(testsplit[2],":")
        testsplit[2] = hostsplit[0]
        hostsplit[1] = string(bytes.Trim([]byte(filterNewLines(hostsplit[1])), "\x00"))
        port, err := strconv.Atoi(hostsplit[1])
        if err != nil{
            fmt.Println(err)
            return false
        }
        uri.port = port
    } else {
        if strings.HasPrefix(uri.uri, "https://"){
            uri.port = 443
        } else {
            uri.port = 80
        }
    }
    uri.hostname = testsplit[2]
    var paths = strings.SplitAfterN(uri.uri, "/", 4)
    if len(paths) >= 4 {
        uri.pathname = string(bytes.Trim([]byte(filterNewLines(paths[3])), "\x00"))
    } else {
        uri.pathname = ""
    }
    return true
}

func main() {
    
}
