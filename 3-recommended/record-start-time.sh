#!/bin/sh

echo "package main

import \"time\"

var StartTime = time.Unix(0, $(date +%s%N))
" > start.go
