<div align="center">
  <a href="#">
    <img src="assets/mooncake_without_name_logo.svg" width="150px" height="150px" />
  </a>

  <h1>Mooncake</h1>
  <p>A simple way to generate mocks for multiple purposes</p>

</div>

[![Go Report Card](https://goreportcard.com/badge/github.com/GuilhermeCaruso/mooncake)](https://goreportcard.com/report/github.com/GuilhermeCaruso/mooncake) [![Build Status](https://app.travis-ci.com/GuilhermeCaruso/mooncake.svg?branch=main)](https://app.travis-ci.com/GuilhermeCaruso/mooncake) [![codecov](https://codecov.io/gh/GuilhermeCaruso/mooncake/branch/main/graph/badge.svg?token=MctCNBxckn)](https://codecov.io/gh/GuilhermeCaruso/mooncake) ![GitHub](https://img.shields.io/badge/golang%20->=1.18-blue.svg) [![GoDoc](https://godoc.org/github.com/GuilhermeCaruso/mooncake?status.svg)](https://godoc.org/github.com/GuilhermeCaruso/mooncake) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT) 

**Table of Contents**

- [What is Mooncake](#what-is-mooncake)
- [Development Status](#development-status)
- [Getting Started](#getting-start)
  - [Installation](#installation)
  - [Mooncake Configuration File](#mooncake-configuration-file)
  - [How to generate](#how-to-generate)
- [License](#license)


## What is Mooncake

Mooncake is a simple way to generate mocks for multiple purposes. 

It was designed to be uncomplicated and simples, focused on development agility while reducing bureaucracy.

Compatible with different types of interfaces such as:

- Default interfaces

```go
type Simple interface{
  MyMethod()
}
```

- Nested interfaces

```go
type Nested interface{
  Simple
}
```

- Generic Interfaces

```go
type Generic[T,Z any] interface{
  MyCustomMethod(T) (T,Z)
}
```
- Generic Nested Interfaces

```go
type NestGeneric[T,Z any] interface{
  Generic[T,Z]
}
```

## Development Status

This project is under development. Therefore, some features may contain minor instabilities, in addition to the possibility of new features being added periodically.


## Getting Start

To start using `mooncake` you need to follow the steps below

### Installation

To install the mooncake generator (`moongen`) run:

```sh

go install github.com/GuilhermeCaruso/mooncake/moongen

```

### Mooncake Configuration File

Once you have decided to use mooncake in your project you will need to create a configuration file

The file must be in the yaml extension. His name doesn't matter, however we recommend it to be mooncake

- Create `mooncake.yaml` file

Once created the following template must be used

```yaml
mocks:
  package: #package
  path: #path
  files:
    - #files
  output: #output
  prefix: #prefix
```

| Field | Definition | Example |
|-|-|-| 
| package | package name of files created | mocks |
| path | path for the interfaces directory | interfaces/ |
| files | list of interface files to be mocked | - |
| output | path to the directory of the generated files| mocks/ |
| prefix | optional value to be added as prefix on generated files | generated |

### How to generate

Once the configuration file is done, to generate the files, run:

```
moongen --file <path_to_config_file>
```


## License

MIT licensed. See the LICENSE file for details.