# go-uniqwords

给一段文本去重单词，但和 `sort | uniq` 不一样的是：它保留第一次出现的顺序，不排序。想要词频也顺手给出来，加个 `-count` 就行。

词用空白切分，标点不算词的一部分（"well." 和 "well" 会被当成两个词），这点要留意。

## 装

```bash
go build -o uniqwords ./cmd/uniqwords
```

## 用

```bash
echo "a b a c b" | ./uniqwords
# a
# b
# c

echo "a b a c b" | ./uniqwords -count
# a 2
# b 2
# c 1
```

## 当库用

```go
import "uniqwords"

uniqwords.Unique("a b a c")        // ["a", "b", "c"]
uniqwords.Count("a b a c")         // [{a 2} {b 1} {c 1}]，按出现顺序
```

## License

MIT
