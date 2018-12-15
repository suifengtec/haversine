# 计算地球上两点的距离

## 文档

[这里这里](https://godoc.org/github.com/suifengtec/haversine)

[这里这里](https://gowalker.org/github.com/suifengtec/haversine)

## 安装

```bash

go get github.com/suifengtec/haversine

```

## 使用

```go


package main

import (
    "fmt"
    distance "github.com/suifengtec/haversine"
)

func main() {

    RUC := distance.Point{Lon: 116.319469, Lat: 39.976778}
    DMC := distance.Point{Lon: 121.320982, Lat: 38.810314}
    m1, km1, mi1 := DMC.Distance(RUC)
    //448861.085595   448.861086      278.856094
    fmt.Printf("%v\t%v\t%v\n", m1, km1, mi1)
}
```

