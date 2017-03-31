# gocook

gocook 提供一个工作池，按照传递进来的`CookMethod`，对唯一名称的meal进行加工并返回加工结果。主要功能是在加工过程中，如果同一个meal名称进来，那么返回的加工结果跟之前的一样。防止在高并发的情况下加工出不一样的结果。

## 应用场景

外部订单号同内部订单号一一对应

## 安装

```shell
go get github.com/qianlnk/gocook
```

## 使用

```go
package main

import (
    "fmt"
    "math/rand"
    "strconv"
    "strings"
    "time"

    "github.com/qianlnk/gocook"
)

func main() {
    mchID := "10001"
    outTradeNo := "20170331100001"
    meal := gocook.NewMeal(mchID+outTradeNo, newInTradeNo, mchID, outTradeNo)
    inTradeNo := meal.Get()
    fmt.Println(inTradeNo)
}

//CookMethod
func newInTradeNo(args ...interface{}) interface{} {
    //TODO 查库 如果有则返回
    var res string

    for _, arg := range args {
        res += fmt.Sprintf("%v", arg)
    }

    res +＝ newRandomString(5)
    //TODO 入库
    return res
}

func newRandomString(length int) string {
    rand.Seed(time.Now().UnixNano())
    rs := make([]string, length)
    for start := 0; start < length; start++ {
        rs = append(rs, strconv.Itoa(rand.Intn(10)))
    }

    return strings.Join(rs, "")
}

```