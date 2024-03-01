#### s.encode('utf-8') in Go:

- just use `[]byte(someString)`
- `[]byte()` longer than `[]rune()`, this makes sense because thats the reason by **bytes** are longer than **runes** especially in unicode, unicode joins multiple bytes to form a single character (**bytes** are normalize in 1...256), while **runes** can go beyond just 256 (1 each)
- 안녕하세요 = [236 149 136 235 133 149 237 149 152 236 132 184 236 154 148]
- in python, you need to list(map(int)) to it, shorter in go
- based on the video:

  1. find length of string

  - py: `len(text)`
  - go: `len([]rune(text)`

  2. find length of bytes

  - py: `len(text.encode('utf-8'))`
  - go: `len([]byte(text)`

  3. example 1:

  - py: `[ord(x)for x in "안녕하세요 👋 (hello in Korean!)"]`
  - go: `[]rune()`

  4. example 2:

  - py: `list("안녕하세요 👋 (hello in Korean!)".encode("utf-8"))`
  - go: `[]byte()`

#### ord() & chr() in go

- **ord()**: rune('a') -> 97
- **chr()**: string(byte(97)) -> 'a'

#### token length bytes vs string

- if string unicode, its always gonna be bigger than normal string, because they _expands_
- e.g.:

```go
// TEXT is a var of complex string
inString := TEXT // len 533
inByte := []byte(TEXT) // len: 616
inByte := []rune(TEXT) // len: 533
```
