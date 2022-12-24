# sfz-parser
![CI](https://github.com/oletizi/sfz/actions/workflows/go.yml/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/oletizi/sfz-parser)](https://goreportcard.com/report/github.com/oletizi/sfz-parser)
![Last commit](https://img.shields.io/github/last-commit/oletizi/sfz)

Basic ANTLR parser for the [sfz file format](https://sfzformat.com/). It currently generates Java and Golang libraries.

The grammar is here: [Sfz.g4](Sfz.g4).

Look at the tests for usage.

# Build and test
    
    %> make
    %> make test
    
# Get it in Go

    import github.com/oletizi/sfz-parser/go


