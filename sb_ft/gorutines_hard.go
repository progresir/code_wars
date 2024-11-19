package main

import (
	"fmt"
	"time"
)
// Вводные
//   c1 := make(chan int, 100)
//   c2 := make(chan int, 100)
//   for i := 0; i < 200; i++ {
//     if i%2 == 0 {
//       c1 <- i
//     } else {
//       c2 <- i
//     }
//   }
// }

// var a = func(i int) int {
//   time.Sleep(5 * time.Second)
//   return i
// }
// есть два канала, которые нужно объеденить в один по формуле a(<-c1)+a(<-c2) и результат вывести в порядке заполнения с1 и с2
// а - это какая-то функция, которая просто долго выполняется
// результат: 1 5 9 ..... 393 397
// 1. Сохранена последовательность
// 2. пункт 1 + выполнение происходит за время задержки в функции а 

func main() {

	c1 := make(chan int, 100)
	c2 := make(chan int, 100)
	c3 := sp(c1, c2)

	for i := 0; i < 200; i++ {
		if i%2 == 0 {
			c1 <- i
		} else {
			c2 <- i
		}

	}
	close(c1)
	close(c2)
	for r := range c3 {
		fmt.Println(r)
	}

}

var a = func(i int) int {
	time.Sleep(1 * time.Second)
	return i
}

func sp(c1, c2 chan int) chan int {
	type turn struct {
		nunber int
		isN    bool
	}
	mp1 := make([]turn, 100)
	mp2 := make([]turn, 100)
	c3 := make(chan int, 100)

	go func() {
		i := 0
		for f := range c1 {
			ii := i
			ff := f
			go func() {
				var t turn
				t.nunber = a(ff)
				t.isN = true
				mp1[ii] = t
			}()
			i++
		}
	}()

	go func() {
		i := 0
		for f := range c2 {
			ii := i
			ff := f
			go func() {
				var t turn
				t.nunber = a(ff)
				t.isN = true
				mp2[ii] = t
			}()
			i++
		}
	}()

	go func() {
		for j := 0; j < 100; {
			if mp1[j].isN {
				if mp2[j].isN {
					c3 <- mp1[j].nunber + mp2[j].nunber
					j++
				}
			}
		}
		close(c3)
	}()

	return c3
}
