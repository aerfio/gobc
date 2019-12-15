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
curl -Lo gobc "https://github.com/aerfio/gobc/releases/download/$(curl -s https://api.github.com/repos/aerfio/gobc/releases/latest | grep tag_name | cut -d '"' -f 4)/gobc_$(uname)_amd64" \
&& chmod +x ./gobc \
&& sudo mv gobc /usr/local/bin
```


## Usage

Just type `gobc` to list local and remote branches. `gobc rm` to remove those excess ones, `gobc completion {SHELL}` to output completion.  
