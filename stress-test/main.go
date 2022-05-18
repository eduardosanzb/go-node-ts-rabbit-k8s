package main

import (
	"fmt"
	"log"
	"os"

	pewpew "github.com/bengadbois/pewpew/lib"
)

func main() {
	url := os.Getenv("PRODUCER_URL")
	if len(url) == 0 {
		log.Fatal("you need to source the PRODUCER_URL")
	}
	stressCfg := pewpew.StressConfig{
		Count:       10000,
		Concurrency: 150,
		Verbose:     false,
		Targets: []pewpew.Target{{
			URL: url,
      Timeout: "1s",
      Method:  "POST",
      Body: `{
        "ts": "1530228282",
        "sender": "testy-test-service",
        "message": {
          "foo": "bar",
          "baz": "bang"
        },
        "sent-from-ip": "1.2.3.4",
        "priority": 2
      }`,
		}},
	}

	output := os.Stdout
	stats, err := pewpew.RunStress(stressCfg, output)
	if err != nil {
		fmt.Printf("pewpew stress failed:  %s", err.Error())
	}

  reqStats := pewpew.CreateRequestsStats(stats[0])
  fmt.Println(pewpew.CreateTextSummary(reqStats))
}
