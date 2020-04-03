#!/bin/bash

set -ex

# cd to the root of the repo.
cd $(dirname $(dirname $0))

echo "Testing 0-base"
tilt ci --file 0-base/Tiltfile
tilt down --file 0-base/Tiltfile

echo "Testing 1-measured"
tilt ci --file 1-measured/Tiltfile
tilt down --file 1-measured/Tiltfile

echo "Testing 2-optimized"
tilt ci --file 2-optimized/Tiltfile
tilt down --file 2-optimized/Tiltfile

echo "Testing 3-recommended"
tilt ci --file 3-recommended/Tiltfile
tilt down --file 3-recommended/Tiltfile
