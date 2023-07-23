# tld-overlap

## small cli thing that checks a word / text for tld overlaps / domain hacks

# running it
first, you need to install the [go programming language](https://go.dev/)

then, clone this repository with your preferred method and run
```sh
go build main.go
```

when the build is done, you can run the cli

```sh
./main yourwordhere
```
enjoy

# updating the tld list
if there are any new tlds that are added, you can run the [make-tld-arr.py](make-tld-arr.py) script (included in the repository) which updates them and rebuilds the cli
