// this file SHALL be generated

package cmd

import (
	"github.com/spf13/cobra"
	"ldcli/cmd/resource"
	"net/http"
	"time"
)

func addAllResourceCmds(rootCmd *cobra.Command) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	g_teamsCmd := resource.NewResourceCmd(rootCmd, "teams")

	resource.NewOperationCmd(
		g_teamsCmd.Cmd,
		client,
		"getTeam",
		"/teams/{teamKey}",
		"get",
		"Get team",
		"Fetch a team by key.\n\n### Expanding the teams response\nLaunchDarkly supports four fields for expanding the \"Get team\" response. By default, these fields are **not** included in the response.\n\nTo expand the response, append the `expand` query parameter and add a comma-separated list with any of the following fields:\n\n* `members` includes the total count of members that belong to the team.\n* `roles` includes a paginated list of the custom roles that you have assigned to the team.\n* `projects` includes a paginated list of the projects that the team has any write access to.\n* `maintainers` includes a paginated list of the maintainers that you have assigned to the team.\n\nFor example, `expand=members,roles` includes the `members` and `roles` fields in the response.\n",
		[]string{"team-key"},
	)
}
