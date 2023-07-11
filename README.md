## ymirblog

> lorem ipsum abracadabra, make some a noise..

[![forthebadge](https://forthebadge.com/images/badges/made-with-go.svg)](https://forthebadge.com) [![forthebadge](https://forthebadge.com/images/badges/built-with-love.svg)](https://forthebadge.com)

### Prerequisites

- Go 1.19+
- Docker (Developed with version 20.+)
- [Taskfile](https://taskfile.dev/) Task runner or Build tool.
- [golangci](https://golangci-lint.run/usage/install/) golang linter.
- [entgo](https://entgo.io/) ORM adapter with sql engine, mysql, postgresql and sqlite.
- For Changelog using [git-chglog](https://github.com/git-chglog/git-chglog)

#### Create Config File

Create an env file development by copying the given example.

```
$ cp .env.defaultexample .env
```

### Developer Guide

#### General Rules

* We use conventional commits to deal with git commits: https://www.conventionalcommits.org
    - Use `feat: commit message` to do git commit related to feature.
    - Use `refactor: commit message` to do git commit related to code refactorings.
    - Use `fix: commit message` to do git commit related to bugfix.
    - Use `test: commit message` to do git commit related to test files.
    - Use `docs: commit message` to do git commit related to documentations (including README.md files).
    - Use `style: commit message` to do git commit related to code style.

* Use git-chglog https://github.com/git-chglog/git-chglog to generate changelog (CHANGELOG.md) before merging to release
  branch.

#### Branching Strategy

* Keep your branch strategy simple. Build your strategy from these three concepts:
    - Use feature branches for all new features and bug fixes.
    - Merge feature branches into the main branch using pull requests.
    - Keep a high quality, up-to-date main branch.

#### Use feature branches for your work

Develop your features and fix bugs in feature branches based off your main branch. These branches are also known as
topic branches. Feature branches isolate work in progress from the completed work in the main branch. Git branches are
inexpensive to create and maintain. Even small fixes and changes should have their own feature branch.

<p align="left"><img src="./featurebranching.png" width="360"></p>

#### Name your feature branches by convention

* Use a consistent naming convention for your feature branches to identify the work done in the branch. You can also
  include other information in the branch name, such as who created the branch.

* Some suggestions for naming your feature branches:
    - users/username/description
    - users/username/workitem
    - bugfix/description
    - feature/feature-name
    - feature/feature-area/feature-name
    - hotfix/description

#### Use release branches

* Create a release branch from the main branch when you get close to your release or other milestone, such as the end of
  a sprint. Give this branch a clear name associating it with the release, for example release/20.
* Create branches to fix bugs from the release branch and merge them back into the release branch in a pull request

<p align="left"><img src="./releasebranching_release.png" width="360"></p>

### Installing and Development

### Development run

```bash
$ task dev
```

### Running the tests

#### Unit Testing

```bash
$ task test
```

#### Code Linter

```bash
$ task lint
```

### Deployment

### Built With

* [chi](https://github.com/go-chi/chi) lightweight, idiomatic and composable router for building Go HTTP services.
* [ent.](https://entgo.io/) An entity framework for Go,
* [openTelemetry](https://opentelemetry.io/) High-quality, ubiquitous, and portable telemetry to enable effective
  observability.
* [zerolog](https://github.com/rs/zerolog) Zero Allocation JSON Logger.

### Authors

* *Nanang Suryadi* - Creator and Initial work - [@Suryakencana07](https://github.com/suryakencana007)

See also the list of [maintainer](MAINTAINER) who participated in this project.

### Contributing
See also the guide to [contributing](CONTRIBUTING).

### License

This project under the Apache License - see the [LICENSE](LICENSE) file for details.