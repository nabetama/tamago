# interface
typeは型を宣言する

## 宣言

```go

type Foo struct {
  フィールドリスト
}

// interface
type Foo interface {
  メソッドリスト
}
```

### メソッドとは

レシーバ名とメソッド名、関数の本体を指定して定義する。

このばあいはpがレシーバ。

```go
// func (レシーバ名, レシーバの型) メソッド名 戻り値の型
func (p *Person) String() string {
  return fmt.Sprintf("%s %s (%d)",  p.FirstName,  p.LastName,  p.Age)
}
```

上のメソッドは下記の構造体へのポインタ型のメソッドとして定義されている

```go
type Person struct {
  FirstName string
  LastName  string
  Age       int
}
```

構造体型にメソッドを設けることが可能になる。

また、インタフェース型でなければ，どんな型にでもメソッドを定義することができる。

```go
type Hex int

func (h Hex) sonomama() Hex {
	return h
}

func main() {
	var h Hex = 42
	fmt.Println(h.sonomama())
}
```

### interfaceを使う


```go
type TypeName interface {
    // メソッドリスト
    Method1()
    Method2()
}
```

インターフェース型宣言時に指定したメソッドリストを全て実装することで

interfaceを実装することができる。Javaみたいに明示的にimplementsとかする必要なし。


```go

package main

import (
	"fmt"
	"strconv"
)

type Car interface {
	run(int) string
	stop()
}

type MyCar struct {
	name  string
	speed int
}

func (u *MyCar) run(speed int) string {
	u.speed = speed
	return strconv.Itoa(u.speed) + "Kmで走行"
}

func (u *MyCar) stop() {
	u.speed = 0
}

func main() {
	car := &MyCar{name: "nabetama", speed: 10}
	car.run(20)
	fmt.Println(car.speed)
	car.stop()
	fmt.Println(car.speed)
}
```

### 空のインターフェース

空のインターフェースにはどんな型でも代入できる。

例1
```go
var x interface{}
num := 0
str := "hello"

x = num // intもおｋ
x = str // stringもおｋ
```

例2
```go
type Any interface{}

func Do(any Any) {
	fmt.Println(any)
}

func main() {
	Do("a")
	Do(123)
}
```

```go
func Do(any interface{}) {
	fmt.Println(any)
}
```

って書き方でもOK。

## 型の埋め込み

クラスが無いから継承が無いので他の型を埋め込むっていうらしい。


```go
type User struct {
	FirstName string
	LastName  string
}

func (u *User) FullName() string {
	fullname := fmt.Sprintf("%s %s", u.FirstName, u.LastName)
	return fullname
}

func NewUser(firstName, lastName string) *User {
	return &User{
		FirstName: firstName,
		LastName:  lastName,
	}
}

type Task struct {
	ID     int
	Detail string
	done   bool
	*User  // Userを埋め込む
}

func NewTask(id int, detail, firstName, lastName string) *Task {
	return &Task{
		ID:     id,
		Detail: detail,
		done:   false,
		User:   NewUser(firstName, lastName),
	}
}

func main() {
	task := NewTask(1, "buy the book", "tama", "nabe")
	fmt.Println(task.FirstName)
	fmt.Println(task.LastName)
	fmt.Println(task.User.FullName())
	fmt.Println(task.FullName())
}
```
