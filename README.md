# config
Go configuration,你能够链式的访问的你任何节点的配置信息

example.yaml
```yaml
steve:
  Hacker: true
  hobbies:
    - skateboarding
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

import (
	"fmt"

	"github.com/holbos-deng/dc"
)

func main() {
	conf := dc.New("example.yaml")

	v1 := conf.Get("steve.age").Value()     // 35
	v2 := conf.Get("steve.hobbies").Value() // [skateboarding snowboarding go ]
	steve := conf.Get("steve")
	v3 := steve.Get("clothing").Get("jacket").Value() // leather
	v4 := steve.Get("clothing").Get("jacket").Key()   // steve.clothing.jacket

	fmt.Println(v1)
	fmt.Println(v2)
	fmt.Println(v3)
	fmt.Println(v4)
}

```