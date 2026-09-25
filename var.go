package main

import "fmt"

func showVar() {
	var s string = "hello"
	var a, b = 1, 2
	fmt.Println(s, a, b)

	var d int
	var e bool
	fmt.Println(d, e)

	var i, j, k int
	var f, g, h = 1, false, "good"
	fmt.Println(i, j, k, f, g, h)

	// Short variable declaration
	i1 := 100
	var x float64 = 100
	p, q := 2, 3
	fmt.Println(i1, x, p, q)

	// constants value
	const pi = 3.14
	fmt.Println(pi)

	// iota
	type Weekday int
	const (
		Monday Weekday = iota
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		Sunday
	)
	fmt.Println(Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday)
	// 有新的const关键字的时候才重新计数
	const (
		a2 = 1
		b2
		c2
		d2 = "ha"
		e2
		f2 = 100
		g2
		h2 = iota
		i2
	)
	fmt.Println(a2, b2, c2, d2, e2, f2, g2, h2, i2)
	// string
	s2 := "hello"
	fmt.Println(len(s2))
	fmt.Println(s2[0], s2[3])
	fmt.Println(s2[:4])
	fmt.Println(s2[2:])
	fmt.Println(s2[:])
	fmt.Println(s2[:] + " openai")
	//条件和循环
	var grade string = "B"
	var marks int = 90
	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 70:
		grade = "C"
	default:
		grade = "D"
	}
	fmt.Println(grade)
	//乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d ", j, i, i*j)
		}
		fmt.Println()
	}
	//1-100素数
	var a3, b3 int
	for a3 = 2; a3 <= 100; a3++ {
		for b3 = 2; b3 <= (a3 / b3); b3++ {
			if a3%b3 == 0 {
				break
			}
		}
		if b3 > (a3 / b3) {
			fmt.Printf("%d ", a3)
		}
	}

	//fallthrough
	switch {
	case false:
		fmt.Println("1、case 条件语句为 false")
		fallthrough
	case true:
		fmt.Println("2、case 条件语句为 true")
		fallthrough
	case false:
		fmt.Println("3、case 条件语句为 false")
		fallthrough
	case true:
		fmt.Println("4、case 条件语句为 true")
	case false:
		fmt.Println("5、case 条件语句为 false")
		fallthrough
	default:
		fmt.Println("6、默认 case")
	}
}
