package main

import (
	"fmt"

	"github.com/katmai1/go-p2p-keymanager/keymanager"
)

func main() {
	clave := keymanager.Newkey()
	fmt.Println(clave.Public_string)
}
