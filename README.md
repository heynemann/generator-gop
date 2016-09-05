# gop

[![Build Status](https://travis-ci.org/heynemann/generator-gop.svg?branch=master)](https://travis-ci.org/heynemann/generator-gop)
[![Coverage Status](https://coveralls.io/repos/github/heynemann/generator-gop/badge.svg?branch=master)](https://coveralls.io/github/heynemann/generator-gop?branch=master)

Gop generator is an yeoman generator that will get you up and running with your go package in seconds.

## Features

- [x] Thorough suite of unit tests for generator code;
- [x] Comprehensive Makefile to build, test, analyze and document your Go Package;
- [x] [Glide](https://glide.sh/) package manager;
- [x] [Ginkgo](https://onsi.github.io/ginkgo/) and [Gomega](http://onsi.github.io/gomega/) test suite with Coverage support;
- [x] Generated code has 100% code coverage in [Ginkgo](https://onsi.github.io/ginkgo/) test suite;
- [x] [SemVer](http://semver.org/) compatible metadata versioning;
- [x] Cross-compiling of package to Darwin and Linux, for both x86_64 and i386 architectures;
- [x] [Cobra](https://github.com/spf13/cobra) commands;
- [x] Current version command;
- [ ] Configuration system using [Viper](https://github.com/spf13/viper);
- [ ] Structured logging using [Zap](https://github.com/uber-go/zap);
- [ ] [Docker](https://www.docker.com/) container to run API;
- [ ] [Docker](https://www.docker.com/) container to run API in Dev Mode with [Docker Compose](https://docs.docker.com/compose/);
- [ ] Benchmark of API routes;
- [ ] Static Analysis of syntax, comments and duplicated code;
- [ ] Allow easy creation of new commands with `yo gop:command`;
- [ ] Allow easy creation of migrations with `yo gop:migration`;
- [ ] Allow easy creation of API handlers with `yo gop:handler`;
- [ ] Easy version handling with `yo gop:major`, `yo gop:minor`, `yo gop:patch`, `yo gop:beta` and `yo gop:rc 1`;
- [ ] [Sphinx](http://www.sphinx-doc.org/en/stable/) documentation;
- [ ] [Travis CI](https://travis-ci.org/) ready.

### Supported Services

- [ ] [Iris](http://iris-go.com/)-powered HTTP API;
- [ ] [Redis](http://redis.io/) Service using [redis](https://github.com/go-redis/redis) as the library;
- [ ] [MongoDB](https://www.mongodb.com/) Database using [mgo](https://labix.org/mgo) as the ORM;
- [ ] [Nats](https://nats.io/) PubSub using [nats](https://github.com/nats-io/nats) as the library;
- [ ] [ElasticSearch](https://www.elastic.co/) Service using [elastic](https://github.com/olivere/elastic) as the library;
- [ ] [PostgreSQL](https://www.postgresql.org/) Database using [PG](https://github.com/go-pg/pg) as the ORM;
- [ ] [MySQL](https://www.mysql.com/) Database using [GORM](https://github.com/jinzhu/gorm) as the ORM.

## Getting started

Make sure you have the latest version of Yeoman:

```
$ npm install -g yo
```

To install generator-gop from npm, run:

```
$ npm install -g generator-gop
```

## Usage

Go to your package directory (must be inside `$GOPATH`) and run:

```
$ yo gop
```

## Contributing

Just fork, go to a feature branch, make your changes, run tests and open a pull request. Rinse and repeat!

### Installing locally

To run the generator locally, in order to test changes:

```
$ npm link
```

After linking you'll be able to run `yo gop` commands.

### Running tests

Just run `npm test` and wait. The tests take a while due to glide installing the packages and go compiling the binary.
