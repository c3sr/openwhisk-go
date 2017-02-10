Blackbox Actions
================

1. Download and install the OpenWhisk CLI
2. Build image
3. Push image
4. Test out action with CLI

The script `buildAndPush.sh` is provided for your convenience. The following command sequence
runs the included example Docker action container using OpenWhisk.

```

# build/push, argument is your docker hub user name and a valid docker image name
./buildAndPush <dockerhub username>/echo-go

# create docker action
wsk action create echo-go --docker <dockerhub username>/echo-go

# invoke created action
wsk action invoke echo-go --blocking
wsk action invoke echo-go --blocking --param key1 value1 --param key2 value2
```

The executable file must be located in the `/action` folder.
The name of the executable must be `/action/exec` and can be any file with executable permissions.
This sample action runs `example.go` by copying and building the source inside the container
as `/action/exec`.
