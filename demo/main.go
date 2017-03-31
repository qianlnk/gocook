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

func newInTradeNo(args ...interface{}) interface{} {
	var res string

	for _, arg := range args {
		res += fmt.Sprintf("%v", arg)
	}

	return res + newRandomString(5)
}

func newRandomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	rs := make([]string, length)
	for start := 0; start < length; start++ {
		rs = append(rs, strconv.Itoa(rand.Intn(10)))
	}
	return strings.Join(rs, "")
}
