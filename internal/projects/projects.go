package projects

import (
	"context"
	"encoding/json"
	"ldcli/internal/errors"

	ldapi "github.com/launchdarkly/api-client-go/v14"
)

type Client interface {
	Create(ctx context.Context, name string, key string) ([]byte, error)
	List(ctx context.Context) ([]byte, error)
}

type ProjectsClient struct {
	client *ldapi.APIClient
}

func NewClient(accessToken string, baseURI string) ProjectsClient {
	config := ldapi.NewConfiguration()
	config.AddDefaultHeader("Authorization", accessToken)
	config.Servers[0].URL = baseURI
	client := ldapi.NewAPIClient(config)

	return ProjectsClient{
		client: client,
	}
}

func (c ProjectsClient) Create(ctx context.Context, name string, key string) ([]byte, error) {
	projectPost := ldapi.NewProjectPost(name, key)
	project, _, err := c.client.ProjectsApi.PostProject(ctx).ProjectPost(*projectPost).Execute()
	if err != nil {
		return nil, errors.NewAPIError(err)
	}
	projectJSON, err := json.Marshal(project)
	if err != nil {
		return nil, err
	}

	return projectJSON, nil
}

func (c ProjectsClient) List(ctx context.Context) ([]byte, error) {
	projects, _, err := c.client.ProjectsApi.
		GetProjects(ctx).
		Limit(2).
		Execute()
	if err != nil {
		return nil, errors.NewAPIError(err)
	}

	projectsJSON, err := json.Marshal(projects)
	if err != nil {
		return nil, err
	}

	return projectsJSON, nil
}
