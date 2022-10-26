# Copilot.cx Go SDK

[![Go Report Card](https://goreportcard.com/badge/github.com/GetWagz/go-copilot)](https://goreportcard.com/report/github.com/GetWagz/go-copilot)
[![GitHub license](https://img.shields.io/github/license/Naereen/StrapDown.js.svg)](https://github.com/GetWagz/go-copilot/blob/master/LICENSE)
[![Maintenance](https://img.shields.io/badge/Maintained%3F-yes-green.svg)](https://GitHub.com/Naereen/StrapDown.js/graphs/commit-activity)

This library is a wrapper around the Copilot.cx collect API.

## Usage

Usage is fairly straight-forward. The `init` function will read from the environment to try to configure the client. However, in some cases you may want to initialize the client programatically, so you may also call the `Setup` function directly.

## Environment Variables

* `COPILOT_CLIENT_ID` The client id for your Copilot instance
* `COPILOT_CLIENT_SECRET` The secret key for your instance
* `COPILOT_CLIENT_COLLECT_ENDPOINT` The collect endpoint
* `COPILOT_CLIENT_CONSENT_ENDPOINT` The consent endpoint, needed for GDPR systems

## Testing

Testing requires an actual account. We can do some initial error checking without credentials, but full testing requires the credentials of an instance to use. Copilot currently does not offer a way to completely remove and reset data so testing may appear minimal without credentials. This library is used in production.

## Other Libraries

We use the following additional tools in this library, and thank the maintainers and contributors of those libraries:

* [testify](https://github.com/stretchr/testify) - Makes our unit tests more readable and management

## Known Issues

* Testing needs to be expanded, ideally with an ability to use a testing sandbox or the ability to reset an environment

## Hiring

Are you interested in working on improving pet lives through innovative technologies and love Go, Typescript, Swift, or Java? Send an email to engineering@wagz.com and let's find out if we're a good match!

## Contributing

Pull Requests are welcome! See our `CONTRIBUTING.md` file for more information.
