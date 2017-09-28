[![Gopher Floc Control](https://s12.postimg.org/69tzd2ckt/go-floc-logo.png)](https://postimg.org/image/8eece5e7d/)

# go-floc
Floc: Orchestrate goroutines with ease.

[![GoDoc](https://godoc.org/gopkg.in/workanator/go-floc.v2?status.svg)](https://godoc.org/gopkg.in/workanator/go-floc.v2)
[![Build Status](https://travis-ci.org/workanator/go-floc.svg?branch=v2)](https://travis-ci.org/workanator/go-floc)
[![Coverage Status](https://coveralls.io/repos/github/workanator/go-floc/badge.svg?branch=v2)](https://coveralls.io/github/workanator/go-floc?branch=v2)
[![Go Report Card](https://goreportcard.com/badge/gopkg.in/workanator/go-floc.v2)](https://goreportcard.com/report/gopkg.in/workanator/go-floc.v2)
[![Join the chat at https://gitter.im/go-floc/Lobby](https://badges.gitter.im/go-floc/Lobby.svg)](https://gitter.im/go-floc/Lobby?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)
[![License](https://img.shields.io/dub/l/vibe-d.svg)](https://github.com/workanator/go-floc/blob/master/LICENSE)

The goal of the project is to make the process of running goroutines in parallel
and synchronizing them easy.

## Installation and requirements

The package requires Go v1.8 or later.

To install the package use `go get gopkg.in/workanator/go-floc.v2`

## Documentation and examples

Please refer [Godoc reference](https://godoc.org/gopkg.in/workanator/go-floc.v2)
of the package for more details.

Some examples are available at the Godoc reference. Additional examples can
be found in [go-floc-showcase](https://github.com/workanator/go-floc-showcase?branch=v2).

## Features

- Easy to use functional interface.
- Simple parallelism and synchronization of jobs.
- As little overhead as possible, in comparison to direct use of goroutines
and sync primitives.
- Provide better control over execution with one entry point and one exit
point.

## Introduction

Floc introduces some terms which are widely used through the package.

### Flow

Flow is the overall process which can be controlled through `floc.Flow`. Flow
can be canceled or completed with any arbitrary data at any point of execution.
Flow has only one enter point and only one exit point.

```go
// Design the job
flow := run.Sequence(do, something, here, ...)

// The enter point: Run the job
result, data, err := floc.Run(flow)

// The exit point: Check the result of the job.
if err != nil {
	// Handle the error
} else if result.IsCompleted() {
	// Handle the success
} else {
	// Handle other cases
}
```

### Job

Job in Floc is a smallest piece of flow. The prototype of job function is
`floc.Job`. Each job can read/write data with `floc.Context` and control
the flow with `floc.Control`.

`Cancel()`, `Complete()`, `Fail()` methods of `floc.Flow` has permanent effect.
Once finished flow cannot be canceled or completed anymore. Calling `Fail` and
returning error from job is almost equal.

```go
func ValidateContentLength(ctx floc.Context, ctrl floc.Control) error {
  request := ctx.Value("request").(http.Request)

  // Cancel the flow with error if request body size is too big
  if request.ContentLength > MaxContentLength {
    return errors.New("content is too big")
  }
}
```

## Example

Lets have some fun and write a simple example which calculates some statistics
on text given. The example designed so it does not require locking because each
part of the `Statistics` structure is accessible only by one job at a moment.

```go
const Text = `Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed
  do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim
  veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo
  consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum
  dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident,
  sunt in culpa qui officia deserunt mollit anim id est laborum.`
  
const keyStatistics = 1

var sanitizeWordRe = regexp.MustCompile(`\W`)

type Statistics struct {
  Words      []string
  Characters int
  Occurrence map[string]int
}

// Split to words and sanitize them
SplitToWords := func(ctx floc.Context, ctrl floc.Control) error {
  statistics := ctx.Value(keyStatistics).(*Statistics)

  statistics.Words = strings.Split(Text, " ")
  for i, word := range statistics.Words {
    statistics.Words[i] = sanitizeWordRe.ReplaceAllString(word, "")
  }
}

// Count and sum the number of characters in the each word
CountCharacters := func(ctx floc.Context, ctrl floc.Control) error {
  statistics := ctx.Value(keyStatistics).(*Statistics)

  for _, word := range statistics.Words {
    statistics.Characters += len(word)
  }
}

// Count the number of unique words
CountUniqueWords := func(ctx floc.Context, ctrl floc.Control) error {
  statistics := ctx.Value(keyStatistics).(*Statistics)

  statistics.Occurrence = make(map[string]int)
  for _, word := range statistics.Words {
    statistics.Occurrence[word] = statistics.Occurrence[word] + 1
  }
}

// Print result
PrintResult := func(ctx floc.Context, ctrl floc.Control) error {
  statistics := ctx.Value(keyStatistics).(*Statistics)

  fmt.Printf("Words Total       : %d\n", len(statistics.Words))
  fmt.Printf("Unique Word Count : %d\n", len(statistics.Occurrence))
  fmt.Printf("Character Count   : %d\n", statistics.Characters)
}

// Design the flow and run it
flow := run.Sequence(
  SplitToWords,
  run.Parallel(
    CountCharacters,
    CountUniqueWords,
  ),
  PrintResult,
)

ctx := floc.NewContext()
ctx.AddValue(keyStatistics, new(Statistics))

ctrl := floc.NewControl(ctx)

_, _, err := floc.RunWith(ctx, ctrl, flow)
if err != nil {
	panic(err)
}

// Output:
// Words Total       : 64
// Unique Word Count : 60
// Character Count   : 370
```

## Contributing

Please found information about contributing in
[CONTRIBUTING.md](https://github.com/workanator/go-floc/blob/master/CONTRIBUTING.md).

