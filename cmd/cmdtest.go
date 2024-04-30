package cmd

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"

	"ldcli/cmd/cliflags"
	"ldcli/internal/analytics"
)

var StubbedSuccessResponse = `{
	"key": "test-key",
	"name": "test-name"
}`

func CallCmd(
	t *testing.T,
	clients APIClients,
	tracker analytics.Tracker,
	args []string,
) ([]byte, error) {
	rootCmd, err := NewRootCommand(
		tracker,
		clients,
		"test",
		false,
	)
	require.NoError(t, err)
	b := bytes.NewBufferString("")
	rootCmd.SetOut(b)
	rootCmd.SetArgs(args)

	err = rootCmd.Execute()
	if err != nil {
		tracker.SendCommandCompletedEvent(
			viper.GetString(cliflags.AccessTokenFlag),
			viper.GetString(cliflags.BaseURIFlag),
			viper.GetBool(cliflags.AnalyticsOptOut),
			analytics.ERROR,
		)
		return nil, err
	}

	tracker.SendCommandCompletedEvent(
		viper.GetString(cliflags.AccessTokenFlag),
		viper.GetString(cliflags.BaseURIFlag),
		viper.GetBool(cliflags.AnalyticsOptOut),
		analytics.SUCCESS,
	)

	out, err := io.ReadAll(b)
	require.NoError(t, err)

	return out, nil
}

// SetupTestEnvVars sets up and tears down tests for checking that environment variables are set.
func SetupTestEnvVars(_ *testing.T) func(t *testing.T) {
	os.Setenv("LD_ACCESS_TOKEN", "testAccessToken")
	os.Setenv("LD_BASE_URI", "http://test.com")

	return func(t *testing.T) {
		os.Unsetenv("LD_ACCESS_TOKEN")
		os.Unsetenv("LD_BASE_URI")
	}
}

func ExtraErrorHelp(cmdName string, cmdAction string) string {
	out := ".\n\nUse `ldcli config --set access-token <value>` to configure the value to persist across CLI commands."
	out += fmt.Sprintf("\n\nSee `ldcli %s %s --help` for supported flags and usage.", cmdName, cmdAction)

	return out
}
