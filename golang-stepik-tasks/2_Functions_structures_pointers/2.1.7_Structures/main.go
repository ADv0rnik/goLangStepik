/*
В рамках этого урока мы постарались представить себе уже привычные нам переменные и функции, как объекты из реальной жизни. Чтобы закрепить результат мы предлагаем вам небольшую творческую задачу.
Вам необходимо реализовать структуру со свойствами-полями On, Ammo и Power, с типами bool, int, int соответственно. У этой структуры должны быть методы: Shoot и RideBike, которые не принимают аргументов, но возвращают значение bool.
Если значение On == false, то оба метода вернут false.
Делать Shoot можно только при наличии Ammo (тогда Ammo уменьшается на единицу, а метод возвращает true), если его нет, то метод вернет false. Метод RideBike работает также, но только зависит от свойства Power.
*/

package main

import "fmt"

type Foo struct {
	On          bool
	Ammo, Power int
}

func (f *Foo) Shoot() bool {
	if f.On == false {
		return false
	} else {
		if f.Ammo > 0 {
			f.Ammo--
			return true
		} else {
			return false
		}
	}
}

func (f *Foo) RideBike() bool {
	if f.On == false {
		return false
	} else {
		if f.Power > 0 {
			f.Power--
			return true
		} else {
			return false
		}
	}
}

func main() {
	foo := Foo{true, 3, 3}
	testStruct := &foo

	fmt.Println(testStruct.Ammo)
	fmt.Println(testStruct.Shoot())
	fmt.Println(testStruct.Ammo)
}
