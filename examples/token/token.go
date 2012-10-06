// token.go - roscoe token example

package main

import (
    "flag"
    "fmt"
    "log"

    "roscoe/client"
    "roscoe/osclib"
)


// debug flag
var debug = flag.Bool("x", false, "Enable debug mode")


func main() {
    help := flag.Bool("help", false, "Show usage")
    verbose := flag.Bool("v", false, "Show token details")

	flag.Usage = func() {
		fmt.Printf("Usage:\n")
		flag.PrintDefaults()
	}
	flag.Parse()
//    client.Debug = debug

    if *help == true {
        flag.Usage()
        return
    }

    // Get auth values from the environment
    var creds osclib.Creds
    c, err := client.NewClient(creds)
    if err != nil {
        log.Fatal(err)
    }

    osclib.GetVersions(c.Auth)

    if *verbose == true {
        fmt.Printf("Token.Id=%s\n", c.Token.Id)
        fmt.Printf("Token.Expires=%s\n", c.Token.Expires)
        fmt.Printf("Tenant.Id=%s\n", c.Token.Tenant.Id)
        fmt.Printf("Tenant.Name=%s\n", c.Token.Tenant.Name)
    } else {
        fmt.Printf("%s %s\n", c.Token.Id, c.Token.Expires)
    }
}
