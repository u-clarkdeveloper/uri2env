# uri2env

uri2env is a tool to convert uri to env and back.

For example:

```shell
uri2env -v -p "MY_PREFIX" "http://userA:passwordA@localhost:8080/path/stuff/here?query=yes&testing=1#some_fragment"
```
Would output
```shell
Input: [http://userA:passwordA@localhost:8080/path/stuff/here?query=yes&testing=1#my_fragment]
Prefix: MY_PREFIX
MY_PREFIX_SCHEME=http
MY_PREFIX_HOST=localhost
MY_PREFIX_PORT=8080
MY_PREFIX_USER=userA
MY_PREFIX_PASSWORD=passwordA
MY_PREFIX_PATH=/path/stuff/here
MY_PREFIX_QUERY=yes
MY_PREFIX_TESTING=1
MY_PREFIX_FRAGMENT=my_fragment
```

### help output
```shell
Usage:
  uri2env [OPTIONS] [COMMANDS] [URI] [flags]

Flags:
  -h, --help            help for uri2env
  -p, --prefix string   Prefix for environment variables
  -v, --verbose         Verbose output: include input and prefix in output
      --version         version for uri2env
```

## Run Dev Version
```shell
go run main.go -v -p MY_PREFIX "http://userA:passwordA@localhost:8080/path/stuff/here?query=yes&testing=1#my_fragment"
```

## Build
```shell
go build -o ./dist/uri2env
```

## Run Build Version

```shell
cd dist
```
```
./uri2env -v -p MY_PREFIX "http://userA:passwordA@localhost:8080/path/stuff/here?query=yes&testing=1#my_fragment"
```

## Create .env file 

### Linux

```shell
./uri2env -v -p MY_PREFIX "http://userA:passwordA@localhost:8080/path/stuff/here?query=yes&testing=1#my_fragment" > .env
```

### Windows

```shell
uri2env.exe -v -p MY_PREFIX "http://userA:passwordA@localhost:8080/path/stuff/here?query=yes&testing=1#my_fragment" > .env
```