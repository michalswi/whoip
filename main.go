package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/michalswi/whoipis/whoip"
)

func main() {
	var result string
	if len(os.Args) < 2 {
		log.Fatalf("Missing arg with URL, e.g. `whoip google.com`")
	} else {
		tldResult := tld(os.Args[1])
		switch tldResult {
		// https://www.nirsoft.net/whois-servers.txt
		case "com":
			result = whoip.WHOip(os.Args[1], "whois.verisign-grs.com")
		case "pl":
			result = whoip.WHOip(os.Args[1], "whois.dns.pl")
		}
	}
	fmt.Println(result)
}

func tld(inURL string) string {
	urlToParse := fmt.Sprintf("https://%s", inURL)
	u, err := url.Parse(urlToParse)
	if err != nil {
		log.Fatalf("Error parsing URL: %v", err)
	}
	parts := strings.Split(u.Hostname(), ".")
	tld := parts[len(parts)-1]
	if tld == "net" {
		tld = "com"
	}
	return tld
}
