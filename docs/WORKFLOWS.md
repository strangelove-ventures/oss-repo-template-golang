# Repository Workflows

This document outlines the necessary steps to correctly configure each of the repository's Github workflows. If a workflow does not apply to the repository, please delete the workflow configuration from the repository.

## Workflows

### Build And Test

File: [build-and-test.yml](../.github/workflows/build-and-test.yml)  

Description: This workflow ensures the project is tidy, buildable, and all go tests are passing before merging any changes into the main branch.  

Configuration Steps:

1. In the workflow, ensure the project's golang version matches what is configured under the `env` section.
1. If the project depends upon private repositories found in the `go.mod` file, another environment variable must be configured in the workflow. Set the `READ_PAT` environment variable to a personal access token with read access to all the dependent private repositories. Strangelove is already configured with a PAT labeled `ORG_READ_PAT` found in the organization's secrets. If this project belongs to Strangelove and will be accessing private repositories within the organization, no configuration is required.
1. If the project does not depend upon private repositories, remove the `READ_PAT` environment variable from the workflow and remove the step that configures the git config on lines `39`, `64`, `83`.

Notes:
- Please see [PAT Configuration](./PATs.md) for more information on how to configure a personal access token for private repository access.

---

### CodeQL Analysis

File: [codeql-analysis.yml](../.github/workflows/codeql-analysis.yml)

Description: This workflow runs a CodeQL analysis on the repository to identify any potential security vulnerabilities.

Configuration Steps:

1. Configure which languages to analyze by updating the `language` section under the `strategy` section. There are additional comments within the workflow stating which languages CodeQL supports.

---

### Interchain Tests

File: [interchain-tests.yml](../.github/workflows/interchain-tests.yml)

Description: This workflow runs this repository's interchain tests. Usually these tests will belong in their own `go` module named `interchaintest`.

Configuration Steps:

1. In the workflow, ensure the project's golang version matches what is configured under the `env` section.
1. In the project's base directory, open the [Makefile](../Makefile) and configure the `Interchain Tests` section. This section outlines the necessary targets to execute every interchain test. Ensure each target is prefixed with `ictest-`.
1. In the workflow, visit the `matrix` section on line `16` and configure the `test` array to match every `make` target previously defined in the last step. The array only needs to include the name of the target.
1. If the project depends upon private repositories found in the `go.mod` file, another environment variable must be configured in the workflow. Set the `READ_PAT` environment variable to a personal access token with read access to all the dependent private repositories. Strangelove is already configured with a PAT labeled `ORG_READ_PAT` found in the organization's secrets. If this project belongs to Strangelove and will be accessing private repositories within the organization, no configuration is required.
1. If the project does not depend upon private repositories, remove the `READ_PAT` environment variable from the workflow and remove the step that configures the git config on line `34`.

Notes:
- Please see [PAT Configuration](./PATs.md) for more information on how to configure a personal access token for private repository access.

---

### Lint

File: [lint.yml](../.github/workflows/lint.yml)

Description: This workflow runs linters on the repository to ensure code quality, consistency, and cleanliness.

Configuration Steps:

1. In the workflow, ensure the project's golang version matches what is configured under the `env` section.
1. In the project's base directory, configure the [.golangci.yml](../.golangci.yml) file to include/exclude linters and to make additional changes as needed.

Notes:
- The [Makefile](../Makefile) contains two targets, `lint` and `lint-fix`, which are used to run the linters locally and fix any issues. Please ensure the workflow and the Makefile targets are using the same version of `golangci-lint`.
- The [Makefile](../Makefile) also contains a target `gofumpt` which is used to format the code locally.

---

### Markdown Link Check

File: [markdown-link-check.yml](../.github/workflows/markdown-link-check.yml)

Description: This workflow checks all markdown files in the repository for broken links.

Configuration Steps: N/A

---

### Go Releaser

File: [release.yml](../.github/workflows/release.yml)

Description: This workflow automatically creates release drafts for the repository when a new tag is pushed to the `main` branch. (Note: release drafts must be manually published.)

Configuration:

1. In the workflow, ensure the project's golang version matches what is configured under the `env` section.
1. Ensure there is a secret labeled `SLACK_WEBHOOK` in the repository's secrets. This secret is used to send a message to the repository's Slack channel when a release is created. If the repository already belongs to the `Strangelove` organization, this secret should already be configured. If not, please see this [Slack Announcement Guide](https://goreleaser.com/customization/announce/slack/) for more information on how to configure the webhook & secret.
1. In the project's base directory, open the [.goreleaser.yaml](../.goreleaser.yaml) file and configure any settings *as needed*. This file is already pre-configured to work with `Strangelove`'s release process. For more information on how to configure this file, please see the [Goreleaser Documentation](https://goreleaser.com).

Notes:
- It is important that commits follow the [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) standard to ensure the release draft is created properly.
- Commits will be grouped by `Features`, `Bug Fixes`, and `Others`. To modify these groups, configure the `changelog` section in the [.goreleaser.yaml](../.goreleaser.yaml) file.

---

### Spell Check

File: [spell-check.yml](../.github/workflows/spell-check.yml)

Description: This workflow checks all files for spelling errors.

Configuration Steps:
1. No configuration steps are necessary for the workflow file.
1. In the project's base directory, open the [.codespellrc](../.codespellrc) file and configure any settings as needed. Sections exist for defining files to skip over, words to ignore, etc. See the [codespell Documentation](https://pypi.org/project/codespell/) under the `Using a config file` section for more information.

---

### Strangelove Project Management

File: [strangelove-project-management.yml](../.github/workflows/strangelove-project-management.yml)

Description: This workflow is used to automatically add issues to Strangelove's `Motherboard` project board.

Configuration Steps: N/A

Notes:
1. Remove this workflow if the repository is not part of the Strangelove organization.

---

### Pull Request Title Format

File: [title-format.yml](../.github/workflows/title-format.yml)

Description: This workflow enforces the [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) standard for pull request titles.

Configuration Steps: N/A

Notes:
- For more information about the workflow, visit [action-semantic-pull-request](https://github.com/amannn/action-semantic-pull-request/).