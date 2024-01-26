package main

import (
	"fmt"
	"log"
	"net"
	"net/url"
	"os"
	"strings"

	"github.com/michalswi/whoisgo/whois"
)

func main() {
	var result string
	if len(os.Args) < 2 {
		log.Fatalf("Missing arg with URL, e.g. `whois google.com`")
	} else {
		tldResult := tld(os.Args[1])
		switch tldResult {
		// https://www.nirsoft.net/whois-servers.txt
		case "com":
			result = whois.WHOis(os.Args[1], "whois.verisign-grs.com")
		case "pl":
			result = whois.WHOis(os.Args[1], "whois.dns.pl")
		case "rawIP":
			result = whois.WHOis(os.Args[1], "whois.ripe.net")
		}
	}
	fmt.Println(result)
}

func tld(record string) string {
	if net.ParseIP(record) == nil {
		urlToParse := fmt.Sprintf("https://%s", record)
		u, err := url.Parse(urlToParse)
		if err != nil {
			log.Fatalf("Error parsing URL: %v", err)
		}
		fmt.Println(u.Hostname())

		parts := strings.Split(u.Hostname(), ".")
		tld := parts[len(parts)-1]
		if tld == "net" {
			tld = "com"
		}
		return tld
	}
	return "rawIP"
}
