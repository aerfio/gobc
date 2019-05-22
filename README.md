<h1 align="center">
 gobc - Go branch clean
</h1>

[![CircleCI](https://circleci.com/gh/aerfio/gobc/tree/master.svg?style=shield)](https://circleci.com/gh/aerfio/gobc/tree/master)
[![Go Report Card](https://goreportcard.com/badge/github.com/aerfio/gobc)](https://goreportcard.com/report/github.com/aerfio/gobc)

## Purpose

My typical workflow consists of making new branch with needed changes, then making PR with it and deleting it afterwards, when it is no logner needed.

Unfortunately it leaves me with useless local branches. That's why `gobc` was created - it checks which local branches are not on remote (origin) and deletes them. 

## Usage

Just type `gobc` to delete branches that are not on remote (origin). You'll encounter pretty prompt which you can use (`tab` + `space`) to select appropriate items. Use `-a` flag to delete _all_ excess branches. Use `-l` flag to list local and origin branches, without deleting them.