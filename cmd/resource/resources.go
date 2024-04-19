package resource

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"ldcli/cmd/validators"
	"log"
	"net/http"
	"regexp"
	"strings"
)

// Base encapsulates the required information needed to make requests to the API
type Base struct {
	APIBaseURL string
}

func (oc *OperationCmd) InitFlags() {
	for _, p := range oc.URLParams {
		p = strings.ReplaceAll(p, "{", "")
		p = strings.ReplaceAll(p, "}", "")
		oc.Cmd.Flags().String(p, "", "")
	}
}

func NewOperationCmd(parentCmd *cobra.Command, client *http.Client, name, path, httpVerb, shortDescription, longDescription string, pathParams []string) *OperationCmd {
	urlParams := extractURLParams(path)
	httpVerb = strings.ToUpper(httpVerb)
	operationCmd := &OperationCmd{
		Base:      &Base{APIBaseURL: "http://localhost"},
		Client:    client,
		Name:      name,
		HTTPVerb:  httpVerb,
		Path:      path,
		URLParams: urlParams,
	}
	cmd := &cobra.Command{
		Use:   name,
		Short: shortDescription,
		Long:  longDescription,
		RunE:  operationCmd.runOperationCmd,
		Args:  validators.Validate(), // TBD on these validators?
	}

	operationCmd.Cmd = cmd
	// TODO: add flags
	operationCmd.InitFlags()

	parentCmd.AddCommand(cmd)

	return operationCmd
}

func extractURLParams(path string) []string {
	re := regexp.MustCompile(`{\w+}`)
	return re.FindAllString(path, -1)
}

func formatURL(path string, urlParams []string) string {
	s := make([]interface{}, len(urlParams))
	for i, v := range urlParams {
		s[i] = v
	}

	re := regexp.MustCompile(`{\w+}`)
	format := re.ReplaceAllString(path, "%s")

	return fmt.Sprintf(format, s...)
}

func NewResourceCmd(parentCmd *cobra.Command, resourceName string) *ResourceCmd {
	description := "CRUD operations for " + resourceName
	cmd := &cobra.Command{
		Use:   resourceName,
		Short: description,
		Long:  description,
	}

	parentCmd.AddCommand(cmd)

	return &ResourceCmd{
		Cmd:           cmd,
		Name:          resourceName,
		OperationCmds: make(map[string]*OperationCmd),
	}
}

// ResourceCmd represents top-level resource commands. Resource commands
// are containers for operation commands.
//
// Example of resources: `projects`, `environments`, `flags`, `teams`
type ResourceCmd struct { //nolint:revive
	Cmd           *cobra.Command
	Name          string
	OperationCmds map[string]*OperationCmd
}

// OperationCmd represents operation commands. Operation commands are nested
// under resource commands and represent a specific API operation for that
// resource.
//
// Examples of operations: `get`, `post`, `delete`, `patch` (standard CRUD methods),
type OperationCmd struct {
	*Base
	Cmd       *cobra.Command
	Client    *http.Client
	Name      string
	HTTPVerb  string
	Path      string
	URLParams []string

	stringFlags map[string]*string
	arrayFlags  map[string]*[]string

	data string
}

func (oc *OperationCmd) runOperationCmd(cmd *cobra.Command, args []string) error {
	url := "/teams/some-key"
	req, _ := http.NewRequest(oc.HTTPVerb, oc.APIBaseURL+url, bytes.NewBuffer(nil))
	req.Header.Add("Authorization", "api-e9e08468-f91e-4e45-a7db-34cc6008ad9d") // TODO: add token here
	req.Header.Add("Content-type", "application/json")

	res, err := oc.Client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	var response interface{}
	body, err := io.ReadAll(res.Body)

	err = json.Unmarshal(body, &response)

	//fmt.Fprintf(cmd.OutOrStdout(), string(response))
	log.Println(response)
	return nil
}
