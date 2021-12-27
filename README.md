# config
Go configuration,你能够链式的访问的你任何节点的配置信息

example.yaml
```yaml
steve:
  Hacker: true
  hobbies:
    - skateboarding
    - 
    - snowboarding
    - go
  clothing:
    jacket: leather
    trousers: denim
  age: 35
  eyes : brown
  beard: true
```

example.go
```go
package main

import "github.com/holbos-deng/rdc"

func main() {
    conf := rdc.New("test.yaml")

    conf.Get("steve.age").Value()     // 35
    conf.Get("steve.hobbies").Value() // [skateboarding snowboarding go ]
    steve := conf.Get("steve")
    steve.Get("clothing").Get("jacket").Value() // leather
    steve.Get("clothing").Get("jacket").Key()   // steve.clothing.jacket
}
```