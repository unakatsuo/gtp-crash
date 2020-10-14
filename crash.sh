#!/bin/bash

set -e

sysctl -w kernel.panic_print=10


./flood/flood &

sleep 1

./gtpsvr/gtpsvr
## kernel will crash.
