package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
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

		log.Println(pinger.IPAddr())

		myadd, err := getaddress("https://checkip.amazonaws.com/")

		if err == nil {
			if myadd == pinger.IPAddr().String() {
				log.Println("The ip address belongs to the dns")
			} else {
				log.Println("The ip address does not belong to the dns")
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
