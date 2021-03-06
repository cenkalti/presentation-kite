// second.go

package main

import (
	"fmt"

	"github.com/koding/kite"
)

func main() {
	k := kite.New("second", "1.0.0")

	client := k.NewClient("http://localhost:6500/kite")
	client.Dial()

	response, _ := client.Tell("square", 4)
	fmt.Println(response.MustFloat64())
}
