package main

import (
	"bufio"
	"context"
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

	logger := client.Logger(logName).StandardLogger(logging.Info)

	r := bufio.NewReader(os.Stdin)
	for {
		l, _, err := r.ReadLine()
		if err != nil {
			return
		}
		fmt.Println(string(l))
		logger.Println(string(l))
	}
}

func getEnvRequired(name string) string {
	s := strings.TrimSpace(os.Getenv(name))
	if s == "" {
		log.Fatalf("env %s required", name)
	}
	return s
}
