## case

**String 與 fmt.Sprintf 循環引用問題**

下面的代碼是有問題的，請說明原因。

```go
type People struct {
	Name string
}

func (p *People) String() string {
	return fmt.Sprintf("%v", p)
}

func main() {
 	p := &People{}
    println(p.String())
}
```

## solution

出現循環調用，然後棧溢出。

在 People 的 String 方法中使用了 %v 打印自身，於是會調用自身的 String 方法，從而循環調用。

查看 fmt 包的文檔，可知：

The default format for %v is: // %v 的默認格式爲：

```go
bool:                    %t
int, int8 etc.:          %d
uint, uint8 etc.:        %d, %#x if printed with %#v
float32, complex64, etc: %g
string:                  %s
chan:                    %p
pointer:                 %p
```

Stringer is implemented by any value that has a String method, which defines the “native” format for that value. The String method is used to print values passed as an operand to any format that accepts a string or to an unformatted printer such as Print. // Stringer包含了一個 String 方法，可以被任何一個值實現，用來定義值的「原始」格式。String 方法被用來打印值，值會作爲操作數傳遞到任何一個接受 string 的 format 中，或者一個不帶 format 的打印器，如 Print。

```go
type Stringer interface {
    String() string
}
```

### f1

String 方法中禁止使用 fmt 包的任何帶 format 的方法。直接使用 StringBuilder 進行字符串拼接。代碼見 [main.go](main.go)。