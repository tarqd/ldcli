package members_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"ldcli/cmd"
	"ldcli/internal/errors"
	"ldcli/internal/members"
)

func TestCreate(t *testing.T) {
	mockArgs := []interface{}{
		"testAccessToken",
		"http://test.com",
		"testemail@test.com",
		"writer",
	}
	t.Run("with valid flags calls members API", func(t *testing.T) {
		client := members.MockClient{}
		client.
			On("Create", mockArgs...).
			Return([]byte(cmd.ValidResponse), nil)
		args := []string{
			"members",
			"create",
			"-t",
			"testAccessToken",
			"-u",
			"http://test.com",
			"-d",
			`{"email": "testemail@test.com", "role": "writer"}`,
		}

		output, err := cmd.CallCmd(t, nil, &client, nil, args)

		require.NoError(t, err)
		assert.JSONEq(t, `{"valid": true}`, string(output))
	})

	t.Run("with an error response is an error", func(t *testing.T) {
		client := members.MockClient{}
		client.
			On("Create", mockArgs...).
			Return([]byte(`{}`), errors.NewError("An error"))
		args := []string{
			"members",
			"create",
			"-t",
			"testAccessToken",
			"-u",
			"http://test.com",
			"-d",
			`{"email": "testemail@test.com", "role": "writer"}`,
		}

		_, err := cmd.CallCmd(t, nil, &client, nil, args)

		require.EqualError(t, err, "An error")
	})

	t.Run("with missing required flags is an error", func(t *testing.T) {
		args := []string{
			"members",
			"create",
		}

		_, err := cmd.CallCmd(t, nil, &members.MockClient{}, nil, args)

		assert.EqualError(t, err, `required flag(s) "accessToken", "data" not set`)
	})

	t.Run("with invalid baseUri is an error", func(t *testing.T) {
		args := []string{
			"projects",
			"create",
			"--baseUri", "invalid",
		}

		_, err := cmd.CallCmd(t, nil, &members.MockClient{}, nil, args)

		assert.EqualError(t, err, "baseUri is invalid")
	})
}