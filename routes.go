package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/gorilla/handlers"
)

func (s *server) Routes() {
	s.router.Handle("/version", handlers.LoggingHandler(os.Stdout, s.handleVersion()))
	s.router.Handle("/containerlist", handlers.LoggingHandler(os.Stdout, s.ListContainers()))
}

//handleVersion supports GET request and returns a json object of {"version":"<version>"}
func (s *server) handleVersion() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			resp := fmt.Sprintf("{\"version\":\"%s\"}", version)
			w.WriteHeader(http.StatusOK)
			w.Header().Add("Content-Type", "application/json")
			w.Write([]byte(resp))
		}
	}
}

//ListContainers returns the ID and image name of
//all containers on the docker host
func (s *server) ListContainers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			cli, err := client.NewClientWithOpts(client.WithVersion("1.37"))

			if err != nil {

				log.Println("Error creating new client.")
			}
			containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{All: true})

			if err != nil {
				log.Println("There was an error creating container list:", err)
			}

			for _, container := range containers {
				fmt.Printf("ID: %s, Image: %s, State: %s\n", container.ID[:10], container.Image, container.State)

			}
		}
	}
}
