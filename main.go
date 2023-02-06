package main

import (
	"context"
	"fmt"

	"cloud.google.com/go/logging"
)

func main() {
	ctx := context.Background()
	client, err := logging.NewClient(ctx, "gcs-fuse-test")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	logger := client.Logger("prince-log")
	logger.Log(logging.Entry{Payload: "Prince Cloud Logging setup!"})

	err = client.Close()
	if err != nil {
		fmt.Println("Error in closing the logger-client: ", err)
	}
}
