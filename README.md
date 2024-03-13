# LD-CLI

## Commands

`go run main.go setup` - runs the setup wizard to create a feature flag for a project and environment.

## Running Github Actions locally

To run the Github Actions locally, you need to install the `act` tool.
```bash
brew install act
```
You will also want to have the `gh` tool installed and authenticated.
```bash
brew install gh
gh auth login
```

Example of how to run the `release-please` action locally. 
`-s` flag is to pass in secrets (optional)
`-j` flag is to specify the job to run (optional)
`-W` flag is to specify the workflow file to run (optional)
```bash
act -s GITHUB_TOKEN="$(gh auth token)" -j release-please -W ./.github/workflows/release-please.yml
```