#!/bin/sh

echo "package start

import \"time\"

var StartTime = time.Unix(0, $(date +%s%N))
" > pkg/start/start.go
