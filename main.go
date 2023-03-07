package main

import (
	"crypto/tls"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/go-ping/ping"
)

func getaddress(dns string) (string, error) {
	resp, err := http.Get(dns)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	log.Println(string(body))
	return string(body), nil
}

func main() {
	const trimwhite = " \n"
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	argLength := len(os.Args[1:])
	if argLength < 1 {
		log.Fatal("dns must be defined")
	}

	log.Println(os.Args[1])
	var address = os.Args[1]

	for {
		pinger, err := ping.NewPinger(address)
		if err != nil {
			panic(err)
		}

		dnsadd := strings.Trim(pinger.IPAddr().IP.String(), trimwhite)
		log.Println(dnsadd)

		myadd, err := getaddress("https://checkip.amazonaws.com/")

		if err == nil {
			myadd = strings.Trim(myadd, trimwhite)
			if myadd == dnsadd {
				log.Println("The ip address belongs to the dns")
			} else {
				log.Println("The ip address does not belong to the dns")
				log.Println("ALERT")
			}
		} else {
			log.Println("Error while quering ip address: " + err.Error())
		}

		time.Sleep(10 * time.Minute)
	}
}

/*
curl checkip.amazonaws.com
curl ifconfig.me
curl icanhazip.com
curl ipecho.net/plain
curl ifconfig.co
*/
