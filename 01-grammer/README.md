# 文法

## hello, world
```go

package main

import (
»-"fmt"
)

func main() {
»-fmt.Println("hello, world")
}
```

### import
パッケージを取り込む

```go
import (
»-"fmt"
)
```

#### import option
importにはオプションがあって、

```go
import (
    f "fmt" // fmtをfで別名にする. python の asと一緒
    _ "github.com/foo/bar"  // 使わないよってことを明示する
    . "strings" // パッケージ名を省略できる
```

別名以外は使ったら可読性下がりそう

### 組み込みの型


- uint8
    - 8ビット符号なし整数
- uint16
    - 16ビット符号なし整数
- uint32
    - 32ビット符号なし整数
- uint64
    - 64ビット符号なし整数
- int8
    - 8ビット符号あり整数
- int16
    - 16ビット符号あり整数
- int32
    - 32ビット符号あり整数
- int64
    - 64ビット符号あり整数
- float32
    - 32ビット浮動小数
- float64
    - 64ビット浮動小数
- complex64
    - 64ビット複素数
- complex128
    - 128ビット複素数
- byte
    - uint8のエイリアス
- rune
    - Unicodeのコードポイント
- uint
    - 32か64ビットの符号なし整数
- int
    - 32か64ビットの符号あり整数
- uintptr
    - ポインタ値用符号なし整数
- error
    - エラーを表わすインタフェース

## 変数

`var 変数名 型名 = <代入するやーつ>`

```go
var message string = "some text."
```

### 一度に複数宣言して代入

横に並べる
```go
var foo, bar, baz = "foo", "bar", "baz"
```

縦に並べる(しっくりこない)
```go
var (
one string = "one"
two = "two"
three = "three"
four = "four"
```

### 関数内部では型推論が行われる

```go
func foo() {
  // var x = "some text."
  x := "some text."
  ...
}
```

### 定数

変数宣言の`var`を`const`に変えると定数になる
定数への再代入はコンパイル エラーになる

```go
func main() {
  const x int = 0
  x = 1 // error cannot assign to x
}
```

### ゼロ値

変数宣言した時に明示的な初期化を行わなかった場合は、
それぞれの型のデフォルト値で暗黙の初期化が行われる。

```go
func main() {
  var i int //  0で初期化される
  fmt.Println(i)  // 0
}
```

### if文

if文はカッコ不要

```go
func main() {
  a, b := 10, 100
  if a > b {
    fmt.Println("a is larger than b")
  } else if a < b {
    fmt.Println("a is smaller than b")
  } else {
    fmt.Println("a equals b")
  }
}
```

- 三項演算子は無し
- カッコつけるとコンパイルエラー

### for文

こっちもカッコ必要なし
while, do/while的なやつらも全てfor文のみで行う

```go
func main() {
  for i := 0; i < 10; i++ {
    fmt.Println(i)
  }
}
```

while置き換え

```go
n := 0
for n < 10 {
  doSomething()
  n++
}

```

#### 無限ループ

for文の条件部を省略すればおｋ

```go
for {
  doSomething()
}
```

### break/continue

```go
func main() {
  n := 0
  for {
    n++
    if n > 10 {
      break // ループから抜ける
    }
    if n%2 == 0 {
      continue // 偶数はスキップ(この時点で次の走査に移る)
    }
    fmt.Println(n) // 最終的に奇数のみ表示される
  }
}
```

#### switch

```go
func main() {
  n : = 10
  switch n {
  case 15:
    fmt.Println("FizzBuzz")
  case 5, 10: // 複数条件を指定できる
    fmt.Println("Buzz")
  case 3, 6, 9:
    fmt.Println("Fizz")
  default:  //  条件に一致しない時
    fmt.Println(n)
  }
}
```

### fallthrough

1つのcase式が処理された後も次のcaseを実行したいって時に使う

```go
func main() {
    n := 3
    switch n {
    case 3:
        n = n - 1
        fallthrough
    case 2:
        n = n - 1
        fallthrough
    case 1:
        n = n - 1
        fmt.Println(n) // 0
    }
}
```

## 関数

`func`で宣言する

#### 引数・戻り値がない関数

```go
func foo() {
  fmt.Println("foo")
}
```

#### 引数がある
引数と型をせんげんする

```go
func sum(i, j int) {  // func foo(i int, j int) と同じ
  fmt.Println(i + j)
}
```

#### 戻り値がある
戻り値の型を宣言する


```go
func sum(i, j int) int {
  return i + j
}
```

#### 複数の戻り値を返す関数

pythonと一緒で複数の戻り値を返すことが出来る

```go
func swap(i, j int) (int, int) {
  return j, i
}
```

```go
a, b = swap(1, 2)
a = swap(1, 2)  // error multiple-value swap(in single-value context)
a, _ = swap(1, 2)  // 無視したければ明示的に書く必要がある
```

#### エラーを返す関数

複数の戻り値を取れることを利用して、内部で発生したエラーについては

`実行結果, エラー情報` 

って感じで返す関数を作るのが作法っぽい。

関数の処理に成功した場合はエラーはnilにし, 
異常があった場合はエラーだけに値が入り、もう一方はゼロ値を返すのが作法っぽい.

##### os.Open()の例

```go
func main() {
  file, err := os.Open("hello.go")
  if err != nil {
    // エラー処理
    // returnなどで処理を別の場所に抜ける
  }
  // fileを用いた処理
}
```

#### 自作のエラーを作る

`errors`パッケージを使う

```go
package main

import (
    "errors"
    "fmt"
    "log"
)

func div(i, j int) (int, error) {
  if j == 0 {
    // 自作のエラーを返す
    return 0, errors.New("divied by zero")
  }
  return i / j, nil
}

func main() {
  n, err := div(10, 0)
  if err != nil {
    // エラーを出力しプログラムを終了する
    log.Fatal(err)
  }
  fmt.Println(n)
}
```

複数の値を返す場合、errorオブジェクトをいちばん最後に返すのが作法なので

自分で作る場合もそうやって作るのがいい。

#### 名前つき戻り値

戻り値にあらかじめ名前をつけて返すことができる。

`func div(i, j int) (result int, err error)`

resultって変数と, errって変数を返すってこと。

名前付き戻り値を宣言すると関数内でその変数がゼロ値で初期化された状態になる。

returnの後に明示的に戻り値を指定しなくてもreturnされた時点で、

名前付き戻り値が勝手に返される。

```go
func div(i, j int) (result int, err error) {
  if j == 0 {
    err = errors.New("Divide by zero.")
    return
  }
  result = i / j
  return
}
```

関数宣言の1行を見れば戻り値の型と内容が読み取れるのは良さそう。


#### 関数リテラル

無名関数のこと。
即時実行する場合は宣言後に`()`で実行する

```go
func main() {
  func(i, j int) {
    fmt.Println(i + j)
  }(2, 4)
}
```

##### 関数オブジェクトを変数へ代入するとき

```go
var sum func(x, y int) (result int) = func(x, y int) (result int) {
	result = x + y
	return
}

func main() {
	fmt.Println(sum(2, 4))  // 6
}
```


## 配列

- 固定長
- 可変長配列は「スライス」って言って別にある

`var arr = [4]string`で宣言

```go
// 宣言してから代入
var arr1 [4]string
arr1[0] = "zero"
arr1[1] = "one"
arr1[2] = "two"
arr1[3] = "three"
fmt.Println(arr[0]) // zero


// 宣言と同時に代入1
var arr := [4]string{"zero", "one", "two", "three"}

// 宣言と同時に代入2
// [...]で必要な配列の長さを暗黙的に指定することが出来る
var arr := [...]string{"zero", "one", "two", "three"}
```

#### 配列は値渡し
値渡しなのでコピーされる.

```go
func fn(arr [4]string) {
	arr[0] = "x"
	fmt.Println(arr)
}

func main() {
	arr := [4]string{"a", "b", "c", "d"}
	fn(arr)           // [x b c d]
	fmt.Println(arr)  // [a b c d]
}
```


## スライス

可変長配列
基本、配列じゃなくてこのスライスを使ったほうがいい

`var slc = []string{"a", "b", "c", "d"}`


#### range

var arr [4]string

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var arr [4]string
	for i, _ := range arr {
		arr[i] = strconv.Itoa(i)
	}
	for i, s := range arr {
		fmt.Println(i, s)
	}
}
```

```
0 0
1 1
2 2
3 3
```


`for i, _ := range arr {`んところ、`for i, s := range arr {`ってやると

`_`つけて明示的に無視してやるとエラーにならない。


#### 値をスライスしても切り出す

```
s := [int{0,  1,  2,  3,  4,  5}
fmt.Println(s[2:4])      // [2 3]
fmt.Println(s[0:len(s)]) // [0 1 2 3 4 5]
fmt.Println(s[:3])       // [0 1 2]
fmt.Println(s[3:])       // [3 4 5]
fmt.Println(s[:])        // [0 1 2 3 4 5]]
```

#### 可変長引数

関数への可変長引数は


```
package main

import (
	"fmt"
)

func sum(nums ...int) (result int) {
	for _, n := range nums {
		result += n
	}
	return
}

func main() {
	fmt.Println(sum(12, 3, 4, 5))
}
```

## マップ

pythonでいうところのdict.

key-valurでデータを保存する。

#### 宣言と初期化
```go
var books[int]string = map[int]string{}

books[1] = "ワングレン"
books[2] = "ベイツール"
```


同時に

```go
func main() {
	books := map[int]string{
		1: "ワングレン",
		2: "ベイツール",
	}

	for k, v := range books {
		fmt.Println(k, v)
	}
}
```

なおmapは順番が保証されてない

#### mapからデータ削除

```go
books := map[int]string{
  1: "ワングレン",
  2: "ベイツール",
}

delete(books, 1)  // map[2:ベイツール]

```

## defer

ファイル開いたあとは閉じる必要がある。

defer宣言の後に処理を書くと、

その処理は関数の最後で必ず実行されるようになる。

```go

func main() {
  file, err := os.Open("./error.go")
  if err != nil {
    // 例外処理
  }
  // 関数を抜ける前に必ず実行する
  defer file.Close()
  // 以下正常処理を記述
}
```

## パニック

python のIndexError, ZeroDivisionError的なやつを起こすと

例外じゃなくてパニックというものが発生する。

```go
func main() {

	defer func() {
		err := recover()
		if err != nil {
			// runtime error: index out of range
			fmt.Println("recover!")
		}
	}()

	a := []int{1, 2, 3}
	fmt.Println(a[10]) // パニックが発生
}
```

パニックは組み込み関数のpanic()を使って自分で起こすことが出来る。


