package main

import "fmt"

type pepeschnele struct {
	speed    int
	charisma int
	wisdom   int
}

func newpepeschnele(speed, charisma, wisdom int) *pepeschnele {
	return &pepeschnele{speed: speed, charisma: charisma, wisdom: wisdom}
}

func (p *pepeschnele) getrating() int {
	return (p.speed * 2) + (p.charisma * 3) + p.wisdom
}

func (p *pepeschnele) String() string {
	return fmt.Sprintf("Пепе Шнеле [Скорость: %d, Харизма: %d, Мудрость: %d] | Рейтинг: %d",
		p.speed, p.charisma, p.wisdom, p.getrating())
}

func main() {
	pepe_1 := newpepeschnele(10, 11, 12)
	pepe_2 := newpepeschnele(13, 14, 15)

	fmt.Println(pepe_1)
	fmt.Println(pepe_2)

	if pepe_1.getrating() > pepe_2.getrating() {
		fmt.Println("Первый выше")
	} else if pepe_1.getrating() < pepe_2.getrating() {
		fmt.Println("Второй выше")
	} else {
		fmt.Println("Одинаково")
	}
}
