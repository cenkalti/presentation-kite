// coin.go
//
// run with:
//     $ parallel -j0 ./coin ::: {1..100}

package main

import (
	"crypto/rand"
	"net/url"
	"strconv"

	"github.com/koding/kite"
)

func main() {
	k := kite.New("coin", "1.0.0")
	k.Config.ReadKiteKey()
	k.Config.Port = 0 // auto-assign
	k.Config.DisableAuthentication = true

	k.HandleFunc("flip", flip)
	go k.Run()
	<-k.ServerReadyNotify()
	k.RegisterForever(&url.URL{Scheme: "http", Host: "127.0.0.1:" + strconv.Itoa(k.Port()), Path: "/kite"})
	<-k.ServerCloseNotify()
}

func flip(r *kite.Request) (interface{}, error) {
	b := make([]byte, 1)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	switch b[0] % 2 {
	case 0:
		return "heads", nil
	case 1:
		return "tails", nil
	default:
		panic("unicorns")
	}
}
