[![Gopher Flow Control](https://s4.postimg.org/p61t5hs31/go-floc-logo.png)](https://postimg.org/image/uhgpq7e5l/)

# go-floc
**Flow Control** makes parallel programming easy.

[![GoDoc](https://godoc.org/gopkg.in/workanator/go-floc.v1?status.svg)](https://godoc.org/gopkg.in/workanator/go-floc.v1)
[![Build Status](https://travis-ci.org/workanator/go-floc.svg?branch=master)](https://travis-ci.org/workanator/go-floc)
[![Coverage Status](https://coveralls.io/repos/github/workanator/go-floc/badge.svg?branch=master)](https://coveralls.io/github/workanator/go-floc?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/workanator/go-floc)](https://goreportcard.com/report/github.com/workanator/go-floc)
[![License](https://img.shields.io/dub/l/vibe-d.svg)](https://github.com/workanator/go-floc/blob/master/LICENSE)

The goal of the project (floc) is to make the process of running goroutines in
parallel and synchronizing them easily.

## Installation and requirements

The package requires Go v1.8

To install the package use `go get gopkg.in/workanator/go-floc.v1`

## Features

- Easy to use functional interface.
- Simple parallelism and synchronization of jobs.
- As little overhead as possible, in comparison to direct use of goroutines
and sync primitives.
- Provide better control over execution with one entry point and one exit
point. That is achieved by allowing any job finish execution with `Cancel` or
`Complete`.

## Introduction

Floc introduces some terms which are widely used through the package.

### Flow

Flow is the overall process which can be controlled through `floc.Flow`. Flow
can be canceled or completed with any arbitrary data at any point of execution.
Flow has only one enter point and only one exit point.

```go
job := run.Sequence(do, something, here, ...)

// The enter point - Run the job
floc.Run(flow, state, update, job)

// The exit point - Check the result of the job.
result, data := flow.Result()
```

### State

State is an arbitrary data shared across all jobs in flow. Since `floc.State`
contains shared data it provides two locking methods, `Get()` for read-only
operations and `GetExclusive()` for read/write operations.

```go
// Read data
data, lock := state.Get()
container := data.(*MyContainer)

lock.Lock()
name := container.Name
date := container.Date
lock.Unlock()

// Write data
data, lock := state.GetExclusive()
container := data.(*MyContainer)

lock.Lock()
container.Counter = container.Counter + 1
lock.Unlock()
```

Floc does not restrict to use state locking methods, safe data read-write
operations can be done using, for example with `sync/atomic`. As well Floc does
not restrict to have data in state. State can contain say channels for
communication between jobs.

```go
type ChunkStream []byte

func WriteToDisk(flow floc.Flow, state floc.State, update floc.Update) {
  data, _ := state.Get()
  stream := data.(ChunkStream)

  file, _ := os.Create("/tmp/file")
  defer file.Close()

  for {
    select {
    case <-flow.Done():
      break
    case chunk := <-stream:
      file.Write(chunk)
    }
  }
}
```

### Update

Update is a function of prototype `floc.Update` which is responsible for
updating state. To identify what piece of state should be updated `key` is used
while `value` contains the data which should be written. It's up to the
implementation how to interpret `key` and `value`.

```go
type Dictionary map[string]interface{}

func UpdateMap(flow floc.Flow, state floc.State, key string, value interface{}) {
  data, lock := state.GetExclusive()
  m := data.(Dictionary)

  lock.Lock();
  defer lock.Unlock()

  m[key] = value
}
```

### Job

Job in Floc is a smallest piece of flow. The prototype of job function is
`floc.Job`. Each job has access to `floc.State` and `floc.Update`, so it can
read/write state data, and to `floc.Flow`, what allows finish flow with
`Cancel()` or `Complete()`.

`Cancel()` and `Complete()` methods of `floc.Flow` has permanent effect. So once
finished flow cannot be canceled or completed anymore.

```go
func ValidateContentLength(flow floc.Flow, state floc.State, update floc.Update) {
  data, _ := state.Get()
  request := data.(http.Request)

  if request.ContentLength > MaxContentLength {
    flow.Cancel(errors.New("content is too big"))
  }
}
```
