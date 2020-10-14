#!/bin/bash

set -e

(
  cd flood
  go build .
)

(
  cd gtpsvr
  go build .
)
