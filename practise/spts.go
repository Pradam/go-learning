package main

import (
	"fmt"
	"strconv"
	"strings"
)

func (p IPAddr) String() string {
	str := make([]string, len(p))
	for i := range p {

		str[i] = strconv.Itoa(int(p[i]))
	}
	return fmt.Sprintf("%v", strings.Join(str, "."))
}

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
