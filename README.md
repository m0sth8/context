# context

[![Build Status](https://travis-ci.org/m0sth8/context.svg?branch=master)](https://travis-ci.org/m0sth8/context)
[![GoDoc](https://godoc.org/github.com/m0sth8/context?status.svg)](https://godoc.org/github.com/m0sth8/context)
[![Coverage Status](https://img.shields.io/coveralls/m0sth8/context.svg?flat=1)](https://coveralls.io/github/m0sth8/context)

Context helpers inspired by https://godoc.org/github.com/docker/distribution/context.

[Documentation](https://godoc.org/github.com/m0sth8/context)

Differences:
 1. Package github.com/satori/go.uuid used for uuid generation instead of github.com/docker/distribution/uuid
 2. Package github.com/apex/log used to log instead of github.com/Sirupsen/logrus
 3. gorilla mux extracted because I don't need it.