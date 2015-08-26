# type, struct, pointer

## type

Goではtypeを用いて既存の型を拡張した独自の型を定義できる。

関数に渡す引数の順番まで決めることが出来て安全!!

#### 例

```go
func ProcessTask(id,  priority int) {
}

id := 3 // int
priority := 5 // int
ProcessTask(id,  priority)

id := 3
priority := 5
ProcessTask(priority,  id) // 引数の型が合ってるのでコンパイルがとおってしまう
```

意味に応じた型を作れば( ・∀・)ｲｲ!!
そこでtypeですよ.

`type <型名> <型>`で宣言!

```go
type ID int
type Priority int

func ProcessTask(id ID, priority Priority)
```


## struct

構造体もデータ型の一種。
複数の型、関数、ポインターとかをまとめておける。
クラスみたいに使うことが多いらしい.

```go
type Task struct {
  ID int
  Detail string
  done bool
}
```

各フィールドのスコープは名前で決まる!!

大文字始まりならパブリック。
小文字始まりならパッケージ内のスコープ

```go
type Task struct {
	ID     int
	Detail string
	done   bool
}

func main() {
	var task Task = Task{
		ID:     1,
		Detail: "foo bar.",
		done:   false,
	}
	fmt.Println(task.ID)
	fmt.Println(task.Detail)
	fmt.Println(task.done)
}
```

#### 構造体の宣言方法あれこれ

順番に入れれば中の型は書かなくても大丈夫
```go
var task Task = Task{
  1,  "Hey! Hey!",  true,
}
```

構造体の生成時に値を明示的に指定しなかった場合は，
ゼロ値で初期化される。

```go
type Task struct {
	ID     int
	Detail string
	done   bool
}

func main() {
	var task Task = Task{
		ID:   1,
		done: false,
	}
	fmt.Println(task.ID)
	fmt.Println(task.Detail)
	fmt.Println(task.done)
}
```

結果
```
1

false
```

#### new()を使ってｲｯﾊﾟﾂで初期化する方法


```go
type Task struct {
    ID int
    Detail string
    done bool
}

func main() {
    var task *Task = new(Task)
    task.ID = 1
    task.Detail = "some text."
    fmt.Println(task.done) // false
}
```


## ポインタ型

構造体から値を生成するときに，構造体の名前の前に&を付けると，
変数には構造体の値ではなくアドレスが格納される。

```go
var kouzou Kouzou = Kouzou{} // Kouzou型
var kouzou *Kouzou = &Kouzou{} // Kouzouのポインタ型
```

`var 変数名 *型名 = *構造体`ってかんじ。

```go
func Finish(task *Task) {
	task.done = true
}

func main() {
	var task *Task = &Task{done: false}
	Finish(task)
	fmt.Println(task.done)
}
```

## コンストラクタ

構造体のコンストラクタは存在しないので

代わりにnewで始まる関数を定義してその中で

構造体を生成してそのポインタを返すのが慣習。

```go
type Task struct {
	ID     int
	Detail string
	done   bool
}

func NewTask(id int, detail string) *Task {
	task := &Task{
		ID:     id,
		Detail: detail,
	}
	return task
}

func main() {
	task := NewTask(1, "foo")
	fmt.Printf("%+v", task)
}
```

## メソッド

型にはメソッドを定義できる。

メソッドはそのメソッドを実行した対象の型をレシーバとして受け取って

メソッドの内部で使用できる。

例1
```go
type Task struct {
	ID     int
	Detail string
	done   bool
}

func NewTask(id int, detail string) *Task {
	task := &Task{
		ID:     id,
		Detail: detail,
	}
	return task
}

func (task Task) String() string {
	str := fmt.Sprintf("%d) %s", task.ID, task.Detail)
	return str
}

func main() {
	task := NewTask(1, "buy the milk")
	fmt.Printf("%s", task)
}
```

例1
```go
type Money struct {
	amount   uint
	currency string
}

func (this *Money) Format() string {
	return fmt.Sprintf("%d %s", this.amount, this.currency)
}

func main() {
	money := &Money{120, "yen"}
	log.Printf(money.Format())
}
```

例3
```go

type Money struct {
	amount   uint
	currency string
}

func (this *Money) Format() string {
	return fmt.Sprintf("%d %s", this.amount, this.currency)
}

func (this *Money) Add(that *Money) {
	this.amount += that.amount
}

func main() {
	money := &Money{120, "yen"}
	log.Printf(money.Format())
	money.Add(&Money{100, "yen"})
	log.Printf(money.Format())
}

```

例4

```go
type Money struct {
	amount   uint
	currency string
}

func (this *Money) ToEmpty() {
	this.amount = 0
}

func main() {
	money := &Money{240, "$"}
	money.ToEmpty()
	log.Printf("%+v", money)
}
```

