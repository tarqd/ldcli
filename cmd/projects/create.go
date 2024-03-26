package projects

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"ldcli/internal/projects"
)

func NewCreateCmd(client projects.Client) (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:     "create",
		Short:   "Create a new project",
		Long:    "Create a new project",
		PreRunE: validate,
		RunE:    runCreate(client),
	}

	cmd.Flags().StringP("data", "d", "", "Input data in JSON")
	err := cmd.MarkFlagRequired("data")
	if err != nil {
		return nil, err
	}
	err = viper.BindPFlag("data", cmd.Flags().Lookup("data"))
	if err != nil {
		return nil, err
	}

	return cmd, nil
}

type inputData struct {
	Name string `json:"name"`
	Key  string `json:"key"`
}

func runCreate(client projects.Client) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		var data inputData
		// TODO: why does viper.GetString("data") not work?
		err := json.Unmarshal([]byte(cmd.Flags().Lookup("data").Value.String()), &data)
		if err != nil {
			return err
		}

		response, err := client.Create(
			context.Background(),
			viper.GetString("accessToken"),
			viper.GetString("baseUri"),
			data.Name,
			data.Key,
		)
		if err != nil {
			return err
		}

		fmt.Fprintf(cmd.OutOrStdout(), string(response)+"\n")

		return nil
	}
}