# tagext
<p align="center">
<a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/license-MIT-_red.svg"></a>
<a href="https://github.com/kawakatz/tagext/issues"><img src="https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat"></a>
<a href="https://goreportcard.com/badge/github.com/kawakatz/tagext"><img src="https://goreportcard.com/badge/github.com/kawakatz/tagext"></a>
<a href="https://www.codefactor.io/repository/github/kawakatz/tagext/badge"><img src="https://www.codefactor.io/repository/github/kawakatz/tagext/badge"></a>
<a href="https://twitter.com/kawakatz"><img src="https://img.shields.io/twitter/follow/kawakatz.svg?logo=twitter"></a>
</p>

tagext extracts the necessary tags and attributes.
## Installation
```sh
➜  ~ go install -v github.com/kawakatz/tagext@latest
```

## Usage
- Extract Anchor Tags
```sh
➜  ~ curl -s https://github.com/ | tagext a
```
- Extract Anchor Tags as HTML
```sh
➜  ~ curl -s https://github.com/ | tagext a -html
```
- Extract href attributes of Anchor tags
```sh
➜  ~ curl -s https://github.com/ | tagext a href
```
