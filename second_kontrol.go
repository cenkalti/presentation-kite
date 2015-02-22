// second_kontrol.go

package main

import (
	"fmt"

	"github.com/koding/kite"
	"github.com/koding/kite/protocol"
)

func main() {
	k := kite.New("second", "1.0.0")
	k.Config.ReadKiteKey()

	kites, _ := k.GetKites(&protocol.KontrolQuery{
		Username:    k.Config.Username,
		Environment: k.Config.Environment,
		Name:        "first",
	})

	client := kites[0]
	client.Dial()

	response, _ := client.Tell("square", 4)
	fmt.Println(response.MustFloat64())
}
