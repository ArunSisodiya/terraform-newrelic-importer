package synthetics

import (
	"github.com/newrelic/newrelic-client-go/newrelic"
	"github.com/newrelic/newrelic-client-go/pkg/synthetics"
	log "github.com/sirupsen/logrus"
)

// GetAllMonitors method to get all the monitors
func GetAllMonitors(client *newrelic.NewRelic) []*synthetics.Monitor {

	monitors, err := client.Synthetics.ListMonitors()
	if err != nil {
		log.Fatal("error listing monitors:", err)
	}

	return monitors
}

// GetAllMonitorsID method to get all the monitors IDs
func GetAllMonitorsID(client *newrelic.NewRelic) []string {

	var monitorIds []string
	monitors := GetAllMonitors(client)

	for _, monitor := range monitors {
		monitorIds = append(monitorIds, monitor.ID)
	}

	return monitorIds

}
