#### s.encode('utf-8') in Go:
- by default, go already encode to utf-8
- just use `[]byte(someString)`
- 안녕하세요 = [236 149 136 235 133 149 237 149 152 236 132 184 236 154 148]
- in python, you need to list(map(int)) to it, shorter in go


#### ord() & chr() in go
- **ord()**: rune('a') -> 97
- **chr()**: string(byte(97)) -> 'a'


#### token length bytes vs string
* if string unicode, its always gonna be bigger than normal string, because they *expands*
* e.g.:
```go
// HELLO is a var of complex string
inString := HELLO // len 533
inByte := []byte(HELLO) // len: 616
```
