import time

contents = """package main

import "time"

var StartTime = time.Unix(0, %d)
""" % int(time.time() * 1000 * 1000 * 1000)

f = open("start.go", "w")
f.write(contents)
f.close()
