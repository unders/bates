# Bates
A tool for testing APIs and Web UIs.

## TODOs
* Fix so that one can run e2e test from commandline with a visible browser
* Fix Jenkins deploy
* Add chromedp for e3e tests
* Think about contract tests... do we need them?

### Why?

#### Testing
* Test APIs using Go's test library
* End-to-End tests through the web browser using Puppeteer and Jest.

#### Deployment
* Deploy Jenkins to remote machine
* Deploy the tests wrapped in a docker image to a Docker registry

### Installation

#### 1. Manually install these programs:
* [Nvm](https://github.com/creationix/nvm)
* Install Node: `nvm install v10.2.1`
* [Go](https://golang.org/doc/install)
* [Docker](https://docs.docker.com/install/)

#### 2. Then install the rest with this command:
```
make install
```

### Troubleshoot
Run command:
```
nvm use v10.2.1
```

#### Doc
* https://github.com/GoogleChrome/puppeteer/blob/master/docs/troubleshooting.md
* https://peter.sh/experiments/chromium-command-line-switches/


### Doc

#### API Testing
* [Go testing](https://golang.org/pkg/testing/)

#### Pact
Consumer Driven Contracts
* [Pact Go](https://github.com/pact-foundation/pact-go/)
* [Pact Docs](https://docs.pact.io/)
* [Consumer driven contracts](https://martinfowler.com/articles/consumerDrivenContracts.html)

#### Headless Chrome
* [headless chrome](https://developers.google.com/web/updates/2017/04/headless-chrome)

#### Puppeteer
Acceptance Tests
* [google-chrome-unstable](https://peter.sh/experiments/chromium-command-line-switches/)
* [jest-environment-puppeteer](https://github.com/smooth-code/jest-puppeteer/tree/master/packages/jest-environment-puppeteer)
* [Docs](https://jasmine.github.io/2.1/introduction.html)
* [Getting started with Puppeteer](https://developers.google.com/web/tools/puppeteer/)
* [Using Puppeteer in TypeScript](https://www.lewuathe.com/using-puppeteer-in-typescript.htm)
* [jest Puppeteer](https://github.com/smooth-code/jest-puppeteer)
* [expect Puppeteer](https://github.com/smooth-code/jest-puppeteer/blob/master/packages/expect-puppeteer/README.md#api)
* [Puppeteer API](https://github.com/GoogleChrome/puppeteer/blob/master/docs/api.md)
* [https://ropig.com/blog/end-end-tests-dont-suck-puppeteer/](https://ropig.com/blog/end-end-tests-dont-suck-puppeteer/)
* [testing-your-react-app-with-puppeteer-and-jest](https://blog.bitsrc.io/testing-your-react-app-with-puppeteer-and-jest-c72b3dfcde59)

#### Chromedp
* [chromedp](https://github.com/chromedp/examples)

#### Cypress
Puppeteer alternative
* https://www.cypress.io/


