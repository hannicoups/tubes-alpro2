package main

import "fmt"

func main() {
	var totalSkor int
	var nama string

	fmt.Println("=== PENILAIAN KESEHATAN MENTAL: SELF-ASSESSMENT DEPRESI ===")
	fmt.Print("Masukkan nama Anda: ")
	fmt.Scanln(&nama)

	fmt.Println("\nJawablah pertanyaan berikut dengan skor (0–3):")
	fmt.Println("0 = Tidak Pernah, 1 = Kadang-kadang, 2 = Sering, 3 = Hampir Setiap Hari")

	pertanyaan := [9]string{
		"a. Kurang berminat atau bergairah dalam melakukan apapun",
		"b. Merasa murung, sedih, atau putus asa",
		"c. Sulit tidur/mudah terbangun, atau terlalu banyak tidur",
		"d. Merasa lelah atau kurang bertenaga",
		"e. Kurang percaya diri atau merasa gagal",
		"f. Kurang nafsu makan atau terlalu banyak makan",
		"g. Sulit berkonsentrasi pada sesuatu, misalnya membaca koran atau menonton televisi",
		"h. Bergerak atau berbicara sangat lambat sehingga orang lain memperhatikannya. Atau sebaliknya; merasa resah atau gelisah sehingga Anda lebih sering bergerak dari biasanya.",
		"i. Merasa lebih baik mati atau ingin melukai diri sendiri dengan cara apapun.",
	}

	totalSkor = pertanyaanSkor(pertanyaan)
	fmt.Println("Total skor Anda andalah: ", totalSkor)

	hasilSkor(totalSkor)
}

func pertanyaanSkor(pertanyaan [9]string) int {
	var i int
	var total, skor int
	total = 0

	for i = 0; i < 9; i++ {
		fmt.Println(pertanyaan[i])
		fmt.Print("Skor Anda (0-3): ")
		fmt.Scan(&skor)
		total += skor
	}
	return total
}

func hasilSkor(total int) {
	fmt.Println("\nHasil Penilaian: ")

	if total <= 4 {
		fmt.Println("Tidak ada gejala depresi.")
	} else if total <= 9 {
		fmt.Println("Gejala depresi ringan. Disarankan psikoedukasi jika gejala memburuk.")
	} else if total <= 14 {
		fmt.Println("Depresi ringan. Observasi 1 bulan dan pertimbangkan terapi atau antidepresan.")
	} else if total <= 19 {
		fmt.Println("Depresi sedang. Disarankan antidepresan atau psikoterapi.")
	} else {
		fmt.Println("Depresi berat. Diperlukan antidepresan dan psikoterapi intensif.")
	}

	fmt.Println("\nTerima kasih telah melakukan self-assessment, semoga Anda sehat selalu.")

}


