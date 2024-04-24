# gobrew

Count all programs written in Go and distributed via Homebrew.

## commands

```sh
go mod tidy && go build -o gobrew main.go && ./gobrew
```

## tasks

- [x] get all Homebrew Core formulas
- [x] save core_formulas as a file onto the disk
- [x] get each package JSON file
- [x] get count of packages which are written/built in Go language
