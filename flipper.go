package main

import (
	"log"
	"sync"

	"github.com/koding/kite"
	"github.com/koding/kite/protocol"
)

func main() {
	k := kite.New("flipper", "1.0.0")
	k.Config.ReadKiteKey()

	kites, _ := k.GetKites(&protocol.KontrolQuery{
		Username:    k.Config.Username,
		Environment: k.Config.Environment,
		Name:        "coin",
	})

	log.Println("There are", len(kites), "running coin kite")

	results := make(chan string)

	var wg sync.WaitGroup
	wg.Add(len(kites))

	for _, k := range kites {
		go func(client *kite.Client) {
			defer wg.Done()

			err := client.Dial()
			if err != nil {
				log.Print(err)
				return
			}
			defer client.Close()

			response, err := client.Tell("flip")
			if err != nil {
				log.Print(err)
				return
			}
			s, err := response.String()
			if err != nil {
				log.Print(err)
				return
			}
			results <- s
		}(k)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	var heads, tails int
	for result := range results {
		switch result {
		case "heads":
			heads++
		case "tails":
			tails++
		}
	}
	log.Println("Total heads:", heads, "tails:", tails)
}
