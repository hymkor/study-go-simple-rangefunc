# Go 1.23 の "iter" パッケージって冗長すぎない？（→ 関数仕様をより簡単にする試み）

Go 1.23 で、簡単な rangefunc を書いてみたのだが、戻り値の型が無駄に複雑で過剰ラッピングに思われた。

[f.go](f.go)

```f.go
package main

import (
    "iter"
)

type seq struct {
    values []int
}

func newSeq(v ...int) *seq {
    return &seq{values: v}
}

func (s *seq) Each() iter.Seq[int] {
    return func(callback func(int) bool) {
        for _, v := range s.values {
            if !callback(v) {
                break
            }
        }
    }
}

func main() {
    s := newSeq(1, 3, 5, 7, 9)
    for v := range s.Each() {
        println(v)
    }
}
```

`C:> go run f.go`

```go run f.go|
1
3
5
7
9
```

そこで "iter" に頼らず、rangefunc 関数を書いてみよう。そして、呼び出しも `range s.Each()` ではなく、`range s.Each` と少し簡単にしてみる。

[h.go](h.go)

```h.go
package main

type seq struct {
    values []int
}

func newSeq(v ...int) *seq {
    return &seq{values: v}
}

func (s *seq) Each(callback func(int) bool) {
    for _, v := range s.values {
        if !callback(v) {
            break
        }
    }
}

func main() {
    s := newSeq(1, 3, 5, 7, 9)
    for v := range s.Each {
        println(v)
    }
}
```

そして、この体裁だと、"iter" を import していないので、実は Go 1.23 以前のパッケージからも呼び出すことができる

[i.go](i.go)

```i.go
package main

type seq struct {
    values []int
}

func newSeq(v ...int) *seq {
    return &seq{values: v}
}

func (s *seq) Each(callback func(int) bool) {
    for _, v := range s.values {
        if !callback(v) {
            break
        }
    }
}

func main() {
    s := newSeq(1, 3, 5, 7, 9)
    s.Each(func(v int) bool {
        println(v)
        return true
    })
}
```

#### まとめ

rangefunc 使う場合でも、なるべく `import "iter"` なしで書いた方が、古い Go からも利用できてよいかもしれない。
