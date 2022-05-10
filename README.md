# gh
go helper

## Install
`go get github.com/arnoluo/gh`

## Usage

### common
```go
import (
    "github.com/arnoluo/gh"
)

// ...
if gh.InArray(1, []int{1, 2, 3}) {
    fmt.Println("yes")
} else {
    fmt.Println("no")
}
```

### int
```go
s123 := gh.Int.Str(123)
i331 := gh.Int.If(1 == 2, 332, 331)
```
