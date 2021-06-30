package main

import (
	"github.com/newrelic/newrelic-client-go/newrelic"
	log "github.com/sirupsen/logrus"
	"os"
)

// Method to initialize the new relic client.
func initialize() *newrelic.NewRelic {

	newRelicConfig := make(map[string]string)
	if apiKey := os.Getenv("NEW_RELIC_API_KEY"); apiKey != "" {
		newRelicConfig["apiKey"] = os.Getenv("NEW_RELIC_API_KEY")
	}
	if accountIDs := os.Getenv("NEW_RELIC_ACCOUNT_ID"); accountIDs != "" {
		newRelicConfig["accountIDs"] = os.Getenv("NEW_RELIC_ACCOUNT_ID")
	}
	if region := os.Getenv("NEW_RELIC_REGION"); region != "" {
		newRelicConfig["region"] = os.Getenv("NEW_RELIC_REGION")
	}
	// Initialize the client
	client, err := newrelic.New(
		newrelic.ConfigPersonalAPIKey(newRelicConfig["apiKey"]),
		newrelic.ConfigRegion(newRelicConfig["region"]),
		newrelic.ConfigLogLevel("INFO"),
	)
	if err != nil {
		log.Fatal("Unable to initialize the new relic client: %v/n", err)
	}
	return client
}
