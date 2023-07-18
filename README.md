# gomake

The Makefile parsing functionality in this project allows you to analyze and extract information from Makefiles. It provides functionality to parse Makefiles, build a dependency graph, check for circular dependencies, and run a specific target with its dependencies

## Features

- Take the file path and target that should be executed from the command
- Check if there is a circular dependency
- Check if there is a dependency not found
- Check if there is an invalid format
- Execute commands for the target and its dependencies

## Installation

- Clone the repository:

```sh
$ git clone https://github.com/codescalersinternships/gomake-Asmaa.git
```

- Change into the project directory:

```sh
$ cd gomake-Asmaa
```

- Build the project:

```sh
go build -o "bin/gomake" cmd/main.go
```

- Change into the project directory:

```sh
$ cd bin
```

- How to use:

  - if -f flag not found, program will search for file called "Makefile", if it's not exist, it will give error "file not found"
  - if -t flag not found, program will give error "no target found"

```sh
$ ./gomake -f [MakefilePath] -t [target]
```

## Makefile Example

```sh
build: test publish
	echo 'executing build'

test:
	echo 'executing test'

publish:
	echo 'executing publish'

```

## How to test

- Run the tests by running:

```sh
go test ./...
```

- If all tests pass, the output indicate that the tests have passed. if there is failure, the output will provide information about it.
