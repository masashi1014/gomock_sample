Before Start
```bash
$ go get github.com/golang/mock/gomock
$ go get github.com/golang/mock/mockgen

```

Create Mock file
```bash
$ mockgen -source repository/user.go -destination repository/mocks/user.go
```

Execute Test
```bash
$ go test ./... -v
```