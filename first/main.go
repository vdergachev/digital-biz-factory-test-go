package main

// На вход подаются два неупорядоченных массива любой длины.
// Необходимо написать функцию, которая возвращает пересечение массивов
func main() {

	r := intersection([]int{1, 5, 89, 3, 97, 332, 1, 1, 0}, []int{1, 1, 9, 4, 8, 4, 2, 1})
	for _, x := range r {
		println(x)
	}
}

func intersection(a []int, b []int) []int {
	var rez []int
	ma := make(map[int]int)

	if len(b) < len(a) {
		t := a
		a = b
		b = t
	}

	for _, i := range a {
		ma[i]++
	}

	for _, i := range b {
		if ma[i] > 0 {
			rez = append(rez, i)
			ma[i]--
		}
	}

	return rez
}
