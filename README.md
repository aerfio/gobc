<h1 align="center">
 gobc - Go branch clean
</h1>

[![CircleCI](https://circleci.com/gh/aerfio/gobc/tree/master.svg?style=shield)](https://circleci.com/gh/aerfio/gobc/tree/master)
[![Go Report Card](https://goreportcard.com/badge/github.com/aerfio/gobc)](https://goreportcard.com/report/github.com/aerfio/gobc)

## Purpose

My typical workflow consists of making new branch with needed changes, then making PR with it and deleting it afterwards, when it is no logner needed.

Unfortunately it leaves me with useless local branches. That's why `gobc` was created - it checks which local branches are not on remote (origin) and deletes them.

## Installation

```bash
GO111MODULE="on" go get -u -v github.com/aerfio/gobc@v0.2.5
```


<!-- curl -Lo kyma.tar.gz "https://github.com/kyma-project/cli/releases/download/$(curl -s https://api.github.com/repos/kyma-project/cli/releases/latest | grep tag_name | cut -d '"' -f 4)/kyma_Darwin_x86_64.tar.gz" \
     && mkdir kyma-release && tar -C kyma-release -zxvf kyma.tar.gz && chmod +x kyma-release/kyma && sudo mv kyma-release/kyma /usr/local/bin \
     && rm -rf kyma-release kyma.tar.gz -->

## Usage

Just type `gobc` to list local and remote branches. `gobc rm` to remove those excess ones, `gobc completion {SHELL}` to output completion.  
