# alfred-jira-workflow

Alfred workflow to view the JIRA resources.

TBD: This workflow doesn't fix the spec, yet...

# HowToDev

## Environment Variables

TBD: Add environment variables requirements

| Environment Variable | Description | Purpose |
|:---|:---|:---|
| GO111MODULE | `on` | Dev/Build |
| alfred_workflow_bundleid | `jp.altab.app.alfred.workflow.jira` | aw-go |
| alfred_workflow_cache | `` `pwd`/tmp/cache`` | aw-go |
| alfred_workflow_data | `` `pwd`/tmp/data`` | aw-go |
| alfred_workflow_version | `1` | aw-go |
| JIRA_API_TOKEN | Your JIRA API Token | JIRA auth |
| JIRA_EMAIL | Your email | JIRA auth |
| JIRA_BASE_URL | Your JIRA endpoint | JIRA auth |

# HowToUse

## CLI Usage

```
$ jira {search}
$ jira MJP-1234

$ jira issue {search}
$ jira issue #1234

$ jira board {search}
$ jira board {name} issue {search}
$ jira board {name} issue #1234

$ jira project {project} issue {search}
$ jira project {project} issue #1234
```

# License

See LICENSE file.
