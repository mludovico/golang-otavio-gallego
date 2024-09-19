package app

import (
	"log"
	"net"

	"github.com/urfave/cli"
)

func Generate() *cli.App {
	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com",
		},
	}

	app := cli.NewApp()
	app.Name = "cli_app"
	app.Usage = "This app is used to demonstrate how to create a CLI app using Go"

	app.Commands = []cli.Command{
		{
			Name:    "ip",
			Aliases: []string{"s"},
			Usage:   "Search for ip addresses on internet",
			Flags:   flags,
			Action:  searchIps,
		},
		{
			Name:   "hosting",
			Usage:  "Get the hosting service of a website",
			Flags:  flags,
			Action: searchHosting,
		},
	}

	return app
}

func searchIps(c *cli.Context) {
	host := c.String("host")
	println("Searching for ip addresses of host:", host)
	ip, error := net.LookupIP(host)
	if error != nil {
		log.Fatal(error)
	}
	for _, ip := range ip {
		println(ip.String())
	}
}

func searchHosting(c *cli.Context) {
	host := c.String("host")
	println("Searching for hosting service of host:", host)
	server, error := net.LookupNS(host)
	if error != nil {
		log.Fatal(error)
	}
	for _, server := range server {
		println(server.Host)
	}
}
