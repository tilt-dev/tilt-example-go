# -*- mode: Python -*-

# For more on Extensions, see: https://docs.tilt.dev/extensions.html
load('ext://restart_process', 'docker_build_with_restart')

# Records the current time, then kicks off a server update.
# Normally, you would let Tilt do deploys automatically, but this
# shows you how to set up a custom workflow that measures it.
local_resource(
    'deploy',
    'python record-start-time.py',
)

compile_cmd = 'CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/tilt-example-go ./'
if os.name == 'nt':
  compile_cmd = 'build.bat'

local_resource(
  'example-go-compile',
  compile_cmd,
  deps=['./main.go', './start.go'],
  resource_deps=['deploy'])

docker_build_with_restart(
  'example-go-image',
  '.',
  entrypoint=['/app/build/tilt-example-go'],
  dockerfile='deployments/Dockerfile',
  only=[
    './build',
    './web',
  ],
  live_update=[
    sync('./build', '/app/build'),
    sync('./web', '/app/web'),
  ],
)

k8s_yaml('deployments/kubernetes.yaml')
k8s_resource('example-go', port_forwards=8000,
             resource_deps=['deploy', 'example-go-compile'])
