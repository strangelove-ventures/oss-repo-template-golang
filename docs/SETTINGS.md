# Repository Settings

These are the default settings that should be configured in each repo, changes can be made as needed.
The settings described here should serve as a baseline, it is more important that there is some setting in place rather 
than the exact settings being used across every repo.

To ensure that the repo is in compliance with standard OSS community expectations, you should make use of the GitHub
[community standards and guidelines page](https://github.com/strangelove-ventures/oss-repo-template-golang/community) 
and ensure that everything is in the green. You may notice that the `Pull request template` section is not green, this 
seems to be an issue with the way GitHub is checking for the presence of the file.

After a repo's settings are configured this file can be safely deleted from the repo since it is not needed for the
project to run.

## General

### Default Branch
- Default branch should be `main`

### Features
- Enable `Issues` 
- Enable `Projects` 

`Wikis` and `Discussions` can be enabled if needed.

### Pull Requests
- Disable allow merge commits
- Enable squash merging with default commit message set to `Pull request title`
- Disable allow rebase merging
- Enable always suggest updating pull request branches
- Enable allow auto-merge
- Enable automatically delete head branches

## Collaborators and teams  

Every repo should have a team or set of teams configured in the `Manage Access` section. Ideally there should be
two teams where one team is given `Admin` access and the other team is given either `Maintain` or `Write` access.

The team responsible for code review should also be added to the [CODEOWNERS](../.github/CODEOWNERS) file so that
they are automatically added as reviewers when a PR is opened. Read more on code owners [here](https://help.github.com/articles/about-codeowners/).

If it is unclear which teams should be configured as codeowners or have admin access, please reach out in Slack in the
`#help-github` channel.

## Moderation options

### Interaction limits

Default settings should be alright here and can also be configured across the entire organization 
in the [organization settings](https://github.com/organizations/strangelove-ventures/settings/interaction_limits).

### Code review limits

Default settings should be alright here.

### Reported content

Enable `Prior contributors and collaborators` to ensure that users who have previously contributed to the repo are able to
report content.

## Branches

A new branch protection rule should be created for the `main` branch. This rule should have the following settings:

- Enable require a pull request before merging
- Enable require approvals with at least 1 approval needed before merging
- Enable require review from Code Owners
- Enable require status checks to pass before merging with the following CI actions being required:
  - Lint PR titles to validate they follow conventional commit format
  - golangci-lint should be passing
  - All unit, integration, and e2e tests are passing
  - Codebase is compiling from the branch used in the PR
- The following CI actions should run on every PR but should not be required status checks before merging:
  - Markdown link checker
  - CodeQL analysis
  - Spell checker
- Enable require branches to be up-to-date before merging
- Enable do not allow bypassing the above settings

Additional branch protection rules should be created for branches that are currently being maintained, used for releases,
and/or other important branches that should not be merged into without proper review and restrictions in place. 

## Tags  

Default settings should be alright here. 

## Rules  

### Rulesets

Default settings should be alright here.

### Insights 

Default settings should be alright here.

## Actions

### General

#### Actions permissions

- Allow all actions and reusable workflows

#### Artifact and log retention 

- Default value of 90 days should be alright here

#### Fork pull request workflows from outside collaborators

- Require approval for first-time contributors

#### Workflow permissions

- Read and write permissions
- Allow GitHub Actions to create and approve pull requests

### Runners

Default settings should be alright here.

## Webhooks  

Default settings should be alright here.

## Copilot  

Default settings should be alright here.

## Environments  

Default settings should be alright here.

## Codespaces  

Default settings should be alright here.

## Pages  

Default settings should be alright here.

## Custom Properties 

Default settings should be alright here.

## Code security and analysis

- Enable Private vulnerability reporting
- Enable dependency graph

It is important to ensure that at least two maintainers are subscribed to receive alerts for security vulnerabilities.
To ensure this is the case, you need to ensure that you are watching the repo and that you enable `Custom` events for
`Security alerts`. After that you will need to navigate to your personal GitHub account settings and ensure that
you have `Notifications` configured to receive alerts for `Participating`, `@mentions`, and `custom` such that you 
will receive notifications via email.

To read more about how to ensure you are receiving notifications for security alerts please review the comment found in
[this issue](https://github.com/strangelove-ventures/oss-repo-template-golang/pull/31#issuecomment-2142932841). If you
are unsure about any of this or need further guidance please reach out in the `#help-github` channel in Slack.

### Dependabot

- Enable Dependabot alerts
- Enable Dependabot security updates
- Enable Dependabot version updates

### Code scanning

- Enable CodeQL analysis

### Secret scanning

- Enable secret scanning to receive alerts on GitHub for detected secrets, keys, or other tokens.

## Deploy keys  

Default settings should be alright here and can be adjusted as needed.

## Secrets and variables

Default settings should be alright here and can be adjusted as needed.



