package examples

import (
	"log"

	"github.com/rid/leasewebgo"
)

func main() {
	c, err := leasewebgo.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	ds, _, err := c.DedicatedServer.List(nil)
	if err != nil {
		log.Fatal(err)
	}
	for _, d := range ds {
		log.Println(d.Id, d.NetworkInterfaces.Public.Ip)
	}
}
