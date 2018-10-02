go string 用 for range 来遍历，用索引遍历和用值遍历，两者的类型都不同

```
str := "hello"
for i, v := range str {
    fmt.Println(reflect.TypeOf(str[i])) // uint8 即 byte
    fmt.Println(reflect.TypeOf(v)) // int32
```