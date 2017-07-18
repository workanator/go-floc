# go-floc
**Flow Control** makes parallel programming easy.

[![GoDoc](https://godoc.org/gopkg.in/workanator/go-floc.v1?status.svg)](https://godoc.org/gopkg.in/workanator/go-floc.v1)
[![Build Status](https://travis-ci.org/workanator/go-floc.svg?branch=master)](https://travis-ci.org/workanator/go-floc)
[![Coverage Status](https://coveralls.io/repos/github/workanator/go-floc/badge.svg?branch=master)](https://coveralls.io/github/workanator/go-floc?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/workanator/go-floc)](https://goreportcard.com/report/github.com/workanator/go-floc)
[![License](https://img.shields.io/dub/l/vibe-d.svg)](https://github.com/workanator/go-floc/blob/master/LICENSE)

The goal of the project (floc) is to make the process of running goroutines in
parallel and synchronizing them easily. Floc achieves that through functional
paradigm where high-order functions are used for building the result execution
flow.

## Features

- Easy to use functional interface.
- Simple parallelism and synchronization of jobs.
- As little overhead as possible, in comparison to direct use of goroutines
and sync primitives.
- Provide better control over execution with one entry point and one exit
point. That is achieved by allowing any job finish execution with `Cancel` or
`Complete`.

## Terms

Floc introduces and operates those terms.
- **Job** is the smaller piece of the overall work.
- **Execution Flow** is the overall work expressed through jobs.
- **State** is the data shared between jobs through the execution flow.

## Installation

To install the package use `go get gopkg.in/workanator/go-floc.v1`
