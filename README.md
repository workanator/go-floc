[![Gopher Floc Control](https://s12.postimg.org/69tzd2ckt/go-floc-logo.png)](https://postimg.org/image/8eece5e7d/)

# go-floc
Floc: Orchestrate goroutines with ease.

[![GoDoc](https://godoc.org/gopkg.in/workanator/go-floc.v1?status.svg)](https://godoc.org/gopkg.in/workanator/go-floc.v1)
[![Build Status](https://travis-ci.org/workanator/go-floc.svg?branch=master)](https://travis-ci.org/workanator/go-floc)
[![Coverage Status](https://coveralls.io/repos/github/workanator/go-floc/badge.svg?branch=master)](https://coveralls.io/github/workanator/go-floc?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/workanator/go-floc)](https://goreportcard.com/report/github.com/workanator/go-floc)
[![License](https://img.shields.io/dub/l/vibe-d.svg)](https://github.com/workanator/go-floc/blob/master/LICENSE)

The goal of the project is to make the process of running goroutines in parallel
and synchronizing them easy.

## Installation and requirements

The package requires Go v1.8

To install the package use `go get gopkg.in/workanator/go-floc.v1`

## Documentation and examples

Please refer [Godoc reference](https://godoc.org/gopkg.in/workanator/go-floc.v1)
of the package for more details.

Some examples are available at the Godoc reference. Additional examples can
be found in [go-floc-showcase](https://github.com/workanator/go-floc-showcase).

## Community

- IRC: `#floc` on `irc.freenode.net`

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
// Design the job
job := run.Sequence(do, something, here, ...)

// The enter point - Run the job
floc.Run(flow, state, update, job)

// The exit point - Check the result of the job.
result, data := flow.Result()
```

### State

State is an arbitrary data shared across all jobs in flow. Since `floc.State`
contains shared data it provides methods which return data alongside with
read-only and/or read/write lockers. Returned lockers are not locked and
the caller is responsible for obtaining and releasing locks.

```go
// Read data
data, locker := state.DataWithReadLocker()

locker.Lock()
container := data.(*MyContainer)
name := container.Name
date := container.Date
locker.Unlock()

// Write data
data, locker := state.DataWithWriteLocker()

locker.Lock()
container := data.(*MyContainer)
container.Counter = container.Counter + 1
locker.Unlock()
```

Floc does not restrict to use state locking methods, safe data read-write
operations can be done using for example `sync/atomic`. As well Floc does
not restrict to have data in state. State can contain say channels for
communication between jobs.

```go
type ChunkStream chan []byte

func WriteToDisk(flow floc.Flow, state floc.State, update floc.Update) {
  stream := state.Data().(ChunkStream)

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
  data, locker := state.DataWithWriteLocker()

  locker.Lock();
  defer locker.Unlock()

  m := data.(Dictionary)
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
  request := state.Data().(http.Request)

  // Cancel the flow with error if request body size is too big
  if request.ContentLength > MaxContentLength {
    flow.Cancel(errors.New("content is too big"))
  }
}
```

## Example

Lets have some fun and write a simple example which calculates some statistics
on text given. The example designed so it does not require locking because each
part of the `Statistics` struct is accessible only by one job at a moment.

```go
const Text = `Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed
  do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim
  veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo
  consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum
  dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident,
  sunt in culpa qui officia deserunt mollit anim id est laborum.`

var sanitizeWordRe = regexp.MustCompile(`\W`)

type Statistics struct {
  Words      []string
  Characters int
  Occurrence map[string]int
}

// Split to words and sanitize them
SplitToWords := func(flow floc.Flow, state floc.State, update floc.Update) {
  statistics := state.Data().(*Statistics)

  statistics.Words = strings.Split(Text, " ")
  for i, word := range statistics.Words {
    statistics.Words[i] = sanitizeWordRe.ReplaceAllString(word, "")
  }
}

// Count and sum the number of characters in the each word
CountCharacters := func(flow floc.Flow, state floc.State, update floc.Update) {
  statistics := state.Data().(*Statistics)

  for _, word := range statistics.Words {
    statistics.Characters += len(word)
  }
}

// Count the number of unique words
CountUniqueWords := func(flow floc.Flow, state floc.State, update floc.Update) {
  statistics := state.Data().(*Statistics)

  statistics.Occurrence = make(map[string]int)
  for _, word := range statistics.Words {
    statistics.Occurrence[word] = statistics.Occurrence[word] + 1
  }
}

// Print result
PrintResult := func(flow floc.Flow, state floc.State, update floc.Update) {
  statistics := state.Data().(*Statistics)

  fmt.Printf("Words Total       : %d\n", len(statistics.Words))
  fmt.Printf("Unique Word Count : %d\n", len(statistics.Occurrence))
  fmt.Printf("Character Count   : %d\n", statistics.Characters)
}

// Design the job and run it
job := run.Sequence(
  SplitToWords,
  run.Parallel(
    CountCharacters,
    CountUniqueWords,
  ),
  PrintResult,
)

floc.Run(
  floc.NewFlow(),
  floc.NewState(new(Statistics)),
  nil,
  job,
)

// Output:
// Words Total       : 64
// Unique Word Count : 60
// Character Count   : 370
```

## Contributing

Please found information about contributing in [CONTRIBUTING.md](https://github.com/workanator/go-floc/blob/master/CONTRIBUTING.md).
