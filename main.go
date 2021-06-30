package main

import (
	"./synthetics"
	"./terraform-exec"
	"fmt"

	"github.com/manifoldco/promptui"
)

func options() string{
	resources := []string{"Alerts & Policies", "Synthetics", "Dashboards"}
	var result string
	var err error

	prompt := promptui.Select{
		Label: "Select resource to be imported",
		Items: resources,
	}

	_, result, err = prompt.Run()

	if err != nil {
		fmt.Printf("Unable to select resource %v\n", err)
		return ""
	}

	return result
}

// Main method to use terraform-newrelic-importer
func main() {
	resource := options()
	client := initialize()
	terraform_exec.CheckVersion()
	switch resource {
	case "Alerts & Policies":
		fmt.Println("Alerts & Policies")
	case "Synthetics":
		monitorsIds := synthetics.GetAllMonitorsID(client)
		fmt.Println(monitorsIds)
	case "Dashboards":
		fmt.Println("Dashboards")

	}

}
