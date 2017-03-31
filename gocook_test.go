package gocook

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"
)

func TestCustomer(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 1100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println("res", NewMeal(strconv.Itoa(i), howToCook, strconv.Itoa(i), i).Get())
		}(i)

		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println("res", NewMeal(strconv.Itoa(i), howToCook, strconv.Itoa(i), strconv.Itoa(i)).Get())
		}(i)
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println("res", NewMeal(strconv.Itoa(i), howToCook, strconv.Itoa(i), strconv.Itoa(i)).Get())
		}(i)

	}
	wg.Wait()
}

func howToCook(args ...interface{}) interface{} {
	var res string
	for _, a := range args {
		res += fmt.Sprintf("%v", a)
	}
	time.Sleep(time.Second)
	return res + "qianlnk" + newRandomString(5)
}

func newRandomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	rs := make([]string, length)
	for start := 0; start < length; start++ {
		rs = append(rs, strconv.Itoa(rand.Intn(10)))
	}
	return strings.Join(rs, "")
}
