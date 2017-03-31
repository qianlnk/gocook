package gocook

import (
	"sync"
)

const cookers = 100

//CookMethod customer must tell chef how to cook
type CookMethod func(flavouring ...interface{}) interface{}

//messboy transfer the meal
type messboy chan interface{}

type chef struct {
	mealName   string
	cookMethod CookMethod
	flavouring []interface{}
}
type order struct {
	mu    sync.Mutex
	meals map[string][]messboy
	chef  chan chef
}

var (
	orders *order
	once   sync.Once
)

func init() {
	once.Do(func() {
		orders = &order{
			meals: make(map[string][]messboy),
			chef:  make(chan chef),
		}
		for i := 0; i < cookers; i++ {
			go orders.cook()
		}
	})
}

func (o *order) cook() {
	for c := range o.chef {
		res := c.cookMethod(c.flavouring...)
		o.mu.Lock()
		for i := range o.meals[c.mealName] {
			o.meals[c.mealName][i] <- res
		}
		delete(o.meals, c.mealName)
		o.mu.Unlock()
	}
}

//Meal store it's messboy
type Meal struct {
	messboy messboy
}

//NewMeal an new meal, name is unique meal name
func NewMeal(name string, method CookMethod, flavouring ...interface{}) *Meal {
	orders.mu.Lock()
	newMessboy := make(messboy)
	_, ok := orders.meals[name]
	orders.meals[name] = append(orders.meals[name], newMessboy)
	orders.mu.Unlock()

	if !ok {
		orders.chef <- chef{
			mealName:   name,
			cookMethod: method,
			flavouring: flavouring,
		}
	}

	return &Meal{
		messboy: newMessboy,
	}
}

//Get meal
func (c *Meal) Get() interface{} {
	defer close(c.messboy)

	return <-c.messboy
}
