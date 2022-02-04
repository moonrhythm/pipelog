package main

import (
	"context"
	"io"
	"log"
	"os"
	"strings"

	"cloud.google.com/go/logging"
)

func main() {
	ctx := context.Background()

	parent := getEnvRequired("PIPELOG_GCP_PARENT")
	logName := getEnvRequired("PIPELOG_GCP_NAME")

	client, err := logging.NewClient(ctx, parent)
	if err != nil {
		log.Fatalf("can not create log client; %v", err)
	}
	defer client.Close()

	logger := client.Logger(logName).StandardLogger(logging.Info)

	go io.Copy(os.Stdout, os.Stdin)
	io.Copy(logger.Writer(), os.Stdin)
}

func getEnvRequired(name string) string {
	s := strings.TrimSpace(os.Getenv(name))
	if s == "" {
		log.Fatalf("env %s required", name)
	}
	return s
}
