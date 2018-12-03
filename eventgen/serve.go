package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/urfave/cli"
	commerce "github.wdf.sap.corp/team-pyrite/hackathon/proto/commerce"
)

var serveCommand = &cli.Command{
	Name:  "serve",
	Usage: "Serve web content",
	Flags: []cli.Flag{
		&cli.IntFlag{
			Name:  "http_port",
			Usage: "Port for http requests",
			Value: 8080,
		},
		&cli.IntFlag{
			Name:  "grpc_port",
			Usage: "Port for grpc requests",
			Value: 50051,
		},
		&cli.StringFlag{
			Name:  "home_dir",
			Usage: "Location of the filesystem root",
			Value: "/tmp/eventgen",
		},
	},
	Action: func(c *cli.Context) error {
		err := serve(c)
		if err != nil {
			return cli.NewExitError(fmt.Sprintf("%v", err), 1)
		}
		return nil
	},
}

func serve(c *cli.Context) error {

	// Create an http client.
	client := newClient()

	eventList := &commerce.EventList{}

	_, err := post(client, eventList)
	if err != nil {
		return err
	}

	return nil
}

func post(client *http.Client, eventList *commerce.EventList) (*http.Response, error) {
	return nil, fmt.Errorf("Not implemented")
}

func newClient() *http.Client {
	return &http.Client{
		Timeout: time.Second * 10,
	}
}
