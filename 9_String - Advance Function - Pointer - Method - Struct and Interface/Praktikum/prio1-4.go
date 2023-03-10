// SOAL
// Buatlah program di Golang untuk menemukan nilai maksimum serta minimum di antara 6 angka inputan. Gunakan multiple return fungsi, pointer untuk referencing maupun deferencing!

package main

import (
	"fmt"
)

func getMinMax(numbers ...*int) (min int, max int) {
	for _, angka := range numbers {
		if *angka > max {
			max = *angka
		}
		//conditional statement =>  untuk mengecek apakah nilai yang ditunjuk pointer angka lebih kecil dari nila variabel min atau min = 0.
		if *angka < min || min == 0 {
			min = *angka
		}
	}
	return min, max
}

func main() {
	var a1, a2, a3, a4, a5, a6, min, max int
	fmt.Scan(&a1)
	fmt.Scan(&a2)
	fmt.Scan(&a3)
	fmt.Scan(&a4)
	fmt.Scan(&a5)
	fmt.Scan(&a6)

	//memanggil func getMinMax  dengan paramater inputan berupa alamat dari variabel a1, a2, a3, a4, a5, a6. kemudian minyimpan nilai tsb ke variabel min dan max.
	min, max = getMinMax(&a1, &a2, &a3, &a4, &a5, &a6)

	//mencetak nilai yg sudah dikembalikan oleh func getMinMax yang disimpan di variabel min dan max.
	fmt.Printf("%d is the minimum number\n", min)
	fmt.Printf("%d is the maximum number\n", max)
}
