package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	flagd "github.com/open-feature/go-sdk-contrib/providers/flagd/pkg"
	"github.com/open-feature/go-sdk/openfeature"
)

const FLAG_KEY = "instructlab_mode"

func main() {
	flagdPort := os.Getenv("FLAGD_PORT")
	if flagdPort == "" {
		log.Fatalf("FLAGD_PORT is not set")
	}
	// Initialize the Flagd provider
	fp, err := strconv.Atoi(flagdPort)
	if err != nil {
		log.Fatalf("Failed to convert FLAGD_PORT to int: %v", err)
	}
	provider := flagd.NewProvider(flagd.WithHost("localhost"), flagd.WithPort(uint16(fp)))

	// Set the Flagd provider as the default provider
	err = openfeature.SetProvider(provider)
	if err != nil {
		log.Fatalf("Failed to set the provider: %v", err)
	}

	// Create a client
	client := openfeature.NewClient("sample-app")

	// Define the handler function
	http.HandleFunc(
		"/", func(w http.ResponseWriter, r *http.Request) {
			ctx := context.Background()
			flagSet, err := client.BooleanValue(ctx, FLAG_KEY, false, openfeature.EvaluationContext{})
			if err != nil {
				http.Error(
					w, fmt.Sprintf("Failed to evaluate feature flag. Error: %v", err), http.StatusInternalServerError,
				)
				return
			}
			template := `<!DOCTYPE html>
<html>
<head>
	<title>Feature Flag Example</title>
</head>
<body>
	<h1>Feature Flag Example</h1>
	<p> Current suppoerted inference mode:</p>
	<h2>%s</h2>
</body>
</html>`

			if flagSet {
				fmt.Fprintln(w, fmt.Sprintf(template, "instructlab"))
			} else {
				fmt.Fprintln(w, fmt.Sprintf(template, "boring mode"))
			}
		},
	)

	// Start the web server
	log.Println("Starting server on :8888")
	log.Fatal(http.ListenAndServe(":8888", nil))
}
