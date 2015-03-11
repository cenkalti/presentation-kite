// first_kontrol.go

package main

import (
	"net/url"

	"github.com/koding/kite"
)

func main() {
	k := kite.New("first", "1.0.0")
	k.Config.ReadKiteKey()
	k.Config.Port = 6501
	k.Config.DisableAuthentication = true

	k.HandleFunc("square", square)
	k.Register(&url.URL{Scheme: "http", Host: "tardis.local:6501", Path: "/kite"})
	k.Run()
}

func square(r *kite.Request) (interface{}, error) {
	a := r.Args.One().MustFloat64()
	return a * a, nil
}
