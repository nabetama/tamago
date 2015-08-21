## Build

```go
$ go build hello.go
```

## Format
コードフォーマット

```go
$ go fmt hello.go
```

## Document

```go
$ godoc fmt
```

## Create Go Project

プロジェクトディレクトリの直下に
`bin`, `pkg`, `src`ってディレクトリを作っておく

```go
$ tree go-project
go-project
├── bin # go instするとここに格納される
├── pkg # 依存パッケージのオブジェクトファイル
└── src # コードはここ
```

そんで$GOPATHをそのディレクトリに通す

でも最近はホームディレクトリ直下の`$HOME/go`下に作ったりするのが定番らしい
このへんは要調査

## Create package.

プロジェクト下の

`go-project/src/gosample/gosample.go`に作る

##### gosampleパッケージ
```go
package gosample

var Message string = "hello, world."
```

##### mainパッケージ
`go-project/src/main/main.go`

```go
package main
import (
  "fmt"
  "gosample"
  )

func main() {
  fmt.Println(gosample.Message) // hello, world.
}
```

### 実行

```sh
$ cd $GOPATH/src/main
$ go run main.go
hello, world.
```

### ビルドしてみる

```sh
$ cd $GOPATH/src
$ go install
```

`go install`すると、ビルドされたパッケージが$GOPATH/bin/以下に配置される。

