# Contributing to st-server

First off, thanks for taking the time to contribute!

The following is a set of guidelines for contributing to st-server, which is hosted on GitHub. These are mostly
guidelines, not rules. Use your best judgment, and feel free to propose changes to this document in a pull request.

## How Can I Contribute?

### Reporting Bugs

This section guides you through submitting a bug report for st-server. Following these guidelines helps maintainers and
the community understand your report, reproduce the behavior, and find related reports.

Before creating bug reports, please check [this list](https://github.com/menezmethod/st-server/issues) as you might find
out that you don't need to create one. When you are creating a bug report, please include as many details as possible.
Fill out [the required template](ISSUE_TEMPLATE/bug_report.md), the information it asks for helps us resolve issues
faster.

### Suggesting Enhancements

This section guides you through submitting an enhancement suggestion for st-server, including completely new features
and minor improvements to existing functionality. Following these guidelines helps maintainers and the community
understand your suggestion and find related suggestions.

Before creating enhancement suggestions, please check [this list](https://github.com/menezmethod/st-server/issues) as
you might find out that you don't need to create one. When you are creating an enhancement suggestion, please include as
many details as possible. Fill in [the template](ISSUE_TEMPLATE/feature_request.md), including the steps that you
imagine you would take if the feature you're requesting existed.

### Pull Requests

The process described here has several goals:

- Maintain st-server's quality
- Fix problems that are important to users
- Engage the community in working toward the best possible st-server
- Enable a sustainable system for st-server's maintainers to review contributions

Please follow these steps to have your contribution considered by the maintainers:

1. Follow all instructions in [the template](PULL_REQUEST_TEMPLATE.md)
2. Follow the [styleguides](#styleguides)
3. After you submit your pull request, verify that
   all [status checks](https://help.github.com/articles/about-status-checks/) are passing

While the prerequisites above must be satisfied prior to having your pull request reviewed, the reviewer(s) may ask you
to complete additional design work, tests, or other changes before your pull request can be ultimately accepted.

## Styleguides

### Git Commit Messages

- Use the present tense ("Add feature" not "Added feature")
- Use the imperative mood ("Move cursor to..." not "Moves cursor to...")
- Limit the first line to 72 characters or less
- Reference issues and pull requests liberally after the first line
- Consider starting the commit message with an applicable emoji:
    - :art: `:art:` when improving the format/structure of the code
    - :racehorse: `:racehorse:` when improving performance
    - :memo: `:memo:` when writing docs
    - :penguin: `:penguin:` when fixing something on Linux
    - :apple: `:apple:` when fixing something on macOS
    - :checkered_flag: `:checkered_flag:` when fixing something on Windows

### Code Style

- Use `gofmt` for formatting
- Follow the [Effective Go](https://golang.org/doc/effective_go.html) guidelines for writing idiomatic Go code
- Add comments to your code where necessary

### Pull Request Process

- Include screenshots and animated GIFs in your pull request whenever possible.
- Follow the [GitHub Flow](https://guides.github.com/introduction/flow/).
- When you submit a pull request, it's expected to have a clear description of what the PR does, as well as passing all
  CI checks.

## Attribution

This Contributing guide is adapted from
the [Atom contributing guide](https://github.com/atom/atom/blob/master/CONTRIBUTING.md).
