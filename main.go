package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"cloud.google.com/go/logging"
)

func main() {
	ctx := context.Background()

	parent := getEnvRequired("PIPELOG_GCP_PARENT")
	logName := getEnvRequired("PIPELOG_GCP_NAME")

	log.Printf("pipelog")
	log.Printf("mode: GCP")
	log.Printf("parent: %s", parent)
	log.Printf("log_name: %s", logName)

	client, err := logging.NewClient(ctx, parent)
	if err != nil {
		log.Fatalf("can not create log client; %v", err)
	}
	defer client.Close()

	logger := client.Logger(logName)

	r := bufio.NewReader(os.Stdin)
	for {
		l, _, err := r.ReadLine()
		if err != nil {
			return
		}
		fmt.Println(string(l))

		e := logging.Entry{
			Severity: logging.Info,
		}
		if json.Valid(l) {
			e.Payload = json.RawMessage(l)
		} else {
			e.Payload = string(l)
		}
		logger.Log(e)
	}
}

func getEnvRequired(name string) string {
	s := strings.TrimSpace(os.Getenv(name))
	if s == "" {
		log.Fatalf("env %s required", name)
	}
	return s
}
