//go:build gen_resources
// +build gen_resources

package main

import (
	"log"

	"ldcli/cmd/resources"
)

func main() {
	data, err := resources.GetTemplateData("../ld-teams-openapi.json")
	if err != nil {
		log.Println(err)
	}
	log.Println(data.Resources["Teams"].Description)
}
