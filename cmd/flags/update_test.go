package flags_test

import (
	"ldcli/cmd"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"ldcli/internal/errors"
	"ldcli/internal/flags"
)

func TestUpdate(t *testing.T) {
	errorHelp := ". See `ldcli flags update --help` for supported flags and usage."
	mockArgs := []interface{}{
		"testAccessToken",
		"http://test.com",
		"test-proj-key",
		"test-key",
		[]flags.UpdateInput{
			{
				Op:    "replace",
				Path:  "/name",
				Value: "new-name",
			},
		},
	}
	t.Run("with valid flags calls projects API", func(t *testing.T) {
		client := flags.MockClient{}
		client.
			On("Update", mockArgs...).
			Return([]byte(cmd.ValidResponse), nil)
		args := []string{
			"flags", "update",
			"--access-token", "testAccessToken",
			"--base-uri", "http://test.com",
			"-d", `[{"op": "replace", "path": "/name", "value": "new-name"}]`,
			"--flag", "test-key",
			"--project", "test-proj-key",
		}

		output, err := cmd.CallCmd(t, nil, &client, nil, nil, args)

		require.NoError(t, err)
		assert.JSONEq(t, `{"valid": true}`, string(output))
	})

	t.Run("with an error response is an error", func(t *testing.T) {
		client := flags.MockClient{}
		client.
			On("Update", mockArgs...).
			Return([]byte(`{}`), errors.NewError("An error"))
		args := []string{
			"flags", "update",
			"--access-token", "testAccessToken",
			"--base-uri", "http://test.com",
			"-d", `[{"op": "replace", "path": "/name", "value": "new-name"}]`,
			"--flag", "test-key",
			"--project", "test-proj-key",
		}

		_, err := cmd.CallCmd(t, nil, &client, nil, nil, args)

		require.EqualError(t, err, "An error")
	})

	t.Run("with missing required flags is an error", func(t *testing.T) {
		args := []string{
			"flags", "update",
		}

		_, err := cmd.CallCmd(t, nil, &flags.MockClient{}, nil, nil, args)

		assert.EqualError(t, err, `required flag(s) "access-token", "data", "flag", "project" not set`+errorHelp)
	})

	t.Run("with invalid base-uri is an error", func(t *testing.T) {
		args := []string{
			"flags", "update",
			"--access-token", "testAccessToken",
			"--base-uri", "invalid",
			"-d", `{"key": "test-key", "name": "test-name"}`,
			"--project", "test-proj-key",
		}

		_, err := cmd.CallCmd(t, nil, &flags.MockClient{}, nil, nil, args)

		assert.EqualError(t, err, "base-uri is invalid"+errorHelp)
	})
}

func TestToggle(t *testing.T) {
	errorHelp := ". See `ldcli flags toggle-on --help` for supported flags and usage."
	mockArgs := []interface{}{
		"testAccessToken",
		"http://test.com",
		"test-proj-key",
		"test-flag-key",
		[]flags.UpdateInput{
			{
				Op:    "replace",
				Path:  "/environments/test-env-key/on",
				Value: true,
			},
		},
	}
	t.Run("with valid flags calls projects API", func(t *testing.T) {
		client := flags.MockClient{}
		client.
			On("Update", mockArgs...).
			Return([]byte(cmd.ValidResponse), nil)
		args := []string{
			"flags", "toggle-on",
			"--access-token", "testAccessToken",
			"--base-uri", "http://test.com",
			"--flag", "test-flag-key",
			"--project", "test-proj-key",
			"--environment", "test-env-key",
		}

		output, err := cmd.CallCmd(t, nil, &client, nil, nil, args)

		require.NoError(t, err)
		assert.JSONEq(t, `{"valid": true}`, string(output))
	})

	t.Run("with an error response is an error", func(t *testing.T) {
		client := flags.MockClient{}
		client.
			On("Update", mockArgs...).
			Return([]byte(`{}`), errors.NewError("An error"))
		args := []string{
			"flags", "toggle-on",
			"--access-token", "testAccessToken",
			"--base-uri", "http://test.com",
			"--flag", "test-flag-key",
			"--project", "test-proj-key",
			"--environment", "test-env-key",
		}

		_, err := cmd.CallCmd(t, nil, &client, nil, nil, args)

		require.EqualError(t, err, "An error")
	})

	t.Run("with missing required flags is an error", func(t *testing.T) {
		args := []string{
			"flags", "toggle-on",
		}

		_, err := cmd.CallCmd(t, nil, &flags.MockClient{}, nil, nil, args)

		assert.EqualError(t, err, `required flag(s) "access-token", "environment", "flag", "project" not set`+errorHelp)
	})

	t.Run("with invalid base-uri is an error", func(t *testing.T) {
		args := []string{
			"flags", "toggle-on",
			"--access-token", "testAccessToken",
			"--base-uri", "invalid",
			"--flag", "test-flag-key",
			"--project", "test-proj-key",
			"--environment", "test-env-key",
		}

		_, err := cmd.CallCmd(t, nil, &flags.MockClient{}, nil, nil, args)

		assert.EqualError(t, err, "base-uri is invalid"+errorHelp)
	})
}
