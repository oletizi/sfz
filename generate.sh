#!/usr/bin/env bash

# NOTE: You must have antlr installed to run this.
antlr -o ./pkg/parser -Dlanguage=Go ./Sfz.g4