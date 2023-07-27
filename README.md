# uri2env

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