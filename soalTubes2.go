package main

import "fmt"

type User struct {
	ID      int
	Nama    string
	Skor    int
	Tanggal string
}

type TabUser [100]User //menyimpan 100 user
var data TabUser

type TabInt [100]int //menyimpan 100 ID agar bisa di-sort atau dicari
var arrID TabInt

var jumlahData int //total data user yang tersimpan

const Qtotal int = 9

type TabPertanyaan [Qtotal]string

func main() {
	var totalSkor, inputID, pilihan int
	var nama, tanggal string
	var pertanyaan TabPertanyaan
	var u User

	fmt.Println("=== PENILAIAN KESEHATAN MENTAL: SELF-ASSESSMENT DEPRESI ===")
	fmt.Println("Masukkan nama Anda: ")
	fmt.Scan(&nama)

	fmt.Printf("Masukkan ID anda: ")
	fmt.Scan(&inputID)

	fmt.Print("Masukkan tanggal pengisian (dd-mm-yyyy): ")
	fmt.Scan(&tanggal)

	inisialisasiPertanyaan(&pertanyaan)

	totalSkor = pertanyaanSkor(pertanyaan)

	fmt.Printf("\nHalo, %s (ID: %03d)\n", nama, inputID)
	fmt.Println("Total skor Anda adalah: ", totalSkor)

	hasilSkor(totalSkor)
	u.ID = inputID
	u.Nama = nama
	u.Skor = totalSkor
	u.Tanggal = tanggal

	data[jumlahData] = u
	arrID[jumlahData] = inputID
	jumlahData++

	for {
		fmt.Println("\nMenu Lanjutan:")
		fmt.Println("1. Ubah Data")
		fmt.Println("2. Hapus Data")
		fmt.Println("3. Cari Data (Sequential Search)")
		fmt.Println("4. Cari Data (Binary Search)")
		fmt.Println("5. Urutkan Berdasarkan Skor (Selection Sort)")
		fmt.Println("6. Urutkan Berdasarkan Tanggal (Insertion Sort)")
		fmt.Println("7. Tampilkan 5 Data Terakhir")
		fmt.Println("8. Tampilkan Rata-rata Skor")
		fmt.Println("9. Tambah Data")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			ubahData(&data, jumlahData)
		case 2:
			hapusData(&data, &arrID, &jumlahData)
		case 3:
			fmt.Print("Masukkan ID yang ingin dicari (Sequential): ")
			fmt.Scan(&inputID)
			idx := sequentialSearch(data, jumlahData, inputID)
			if idx != -1 {
				fmt.Printf("Data ditemukan: %s (Skor: %d)\n", data[idx].Nama, data[idx].Skor)
			} else {
				fmt.Println("Data tidak ditemukan.")
			}
		case 4:
			urutkanData(&data, &arrID, jumlahData) // pastikan data sudah terurut
			fmt.Print("Masukkan ID yang ingin dicari (Binary): ")
			fmt.Scan(&inputID)
			idx := binarySearch(data, jumlahData, inputID)
			if idx != -1 {
				fmt.Printf("Data ditemukan: %s (Skor: %d)\n", data[idx].Nama, data[idx].Skor)
			} else {
				fmt.Println("Data tidak ditemukan.")
			}
		case 5:
			selectionSortSkor(&data, jumlahData)
			fmt.Println("Data berhasil diurutkan berdasarkan skor.")
			tampilan(&data, jumlahData)
		case 6:
			insertionSortTanggal(&data, jumlahData)
			fmt.Println("Data berhasil diurutkan berdasarkan tanggal.")
			tampilan(&data, jumlahData)
		case 7:
			tampilkan5Terakhir(data, jumlahData)

		case 8:
			var hariIni string
			fmt.Print("Masukkan tanggal hari ini (dd-mm-yyyy): ")
			fmt.Scan(&hariIni)
			avg := rataRataBulanTerakhir(data, jumlahData, hariIni)
			if avg == 0 {
				fmt.Println("Tidak ada data dalam 30 hari terakhir.")
			} else {
				fmt.Printf("Rata-rata skor dalam 30 hari terakhir: %.2f\n", avg)
			}
		case 9:
			tambahData(&data, &arrID, &jumlahData)

		case 0:
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}

}

func tampilan(D *TabUser, N int) {
	var i int
	fmt.Println("\nData setelah diurutkan berdasarkan ID:")
	fmt.Println("ID\tNama\tSkor")
	for i = 0; i < N; i++ {
		fmt.Printf("%03d\t%s\t%d\n", D[i].ID, D[i].Nama, D[i].Skor)
	}
}

func inisialisasiPertanyaan(p *TabPertanyaan) {
	*p = TabPertanyaan{ //Mengisi (atau mengganti) isi dari TabPertanyaan yang ditunjuk oleh pointer p dengan data baru.
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
}

func pertanyaanSkor(A TabPertanyaan) int {
	var i int
	var total, skor int

	fmt.Println("\nJawablah pertanyaan berikut dengan skor (0–3):")
	fmt.Println("0 = Tidak Pernah, 1 = Kadang-kadang, 2 = Sering, 3 = Hampir Setiap Hari")

	total = 0

	for i = 0; i < Qtotal; i++ {
		fmt.Println(A[i])
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

func urutkanData(D *TabUser, ID *TabInt, N int) {
	var i, j int
	for i = 0; i < N-1; i++ {
		for j = i + 1; j < N; j++ {
			if ID[i] > ID[j] {
				ID[i], ID[j] = ID[j], ID[i]
				D[i], D[j] = D[j], D[i]
			}
		}
	}
}

func ubahData(D *TabUser, N int) {
	var cariID, i int
	var tanggalBaru string
	var ditemukan bool = false
	var idxDitemukan int = -1
	var pertanyaan TabPertanyaan

	fmt.Print("Masukkan ID yang ingin diubah: ")
	fmt.Scan(&cariID)

	i = 0
	for i < N {
		if D[i].ID == cariID {
			ditemukan = true
			idxDitemukan = i
		}
		i++
	}

	if ditemukan {
		fmt.Printf("Data ditemukan: %s (Skor lama: %d, Tanggal lama: %s)\n", D[idxDitemukan].Nama, D[idxDitemukan].Skor, D[idxDitemukan].Tanggal)

		fmt.Println("\nSilakan isi kembali kuisioner:")
		inisialisasiPertanyaan(&pertanyaan)
		skorBaru := pertanyaanSkor(pertanyaan)

		fmt.Print("Masukkan tanggal baru (dd-mm-yyyy): ")
		fmt.Scan(&tanggalBaru)
		D[idxDitemukan].Skor = skorBaru
		D[idxDitemukan].Tanggal = tanggalBaru
		fmt.Println("Data berhasil diubah.")
		hasilSkor(skorBaru)
	} else {
		fmt.Println("Data dengan ID tersebut tidak ditemukan.")
	}
}

func hapusData(D *TabUser, ID *TabInt, N *int) {
	var cariID, i, j, indeks int
	var ditemukan bool = false

	fmt.Print("Masukkan ID yang ingin dihapus: ")
	fmt.Scan(&cariID)

	indeks = -1
	for i = 0; i < *N; i++ {
		if D[i].ID == cariID {
			ditemukan = true
			indeks = i
		}
	}

	if ditemukan {
		for j = indeks; j < *N-1; j++ {
			(*D)[j] = (*D)[j+1]
			(*ID)[j] = (*ID)[j+1]
		}
		*N-- 
		fmt.Println("Data berhasil dihapus.")
	} else {
		fmt.Println("ID tidak ditemukan.")
	}
}

func sequentialSearch(D TabUser, N int, cariID int) int {
	for i := 0; i < N; i++ {
		if D[i].ID == cariID {
			return i
		}
	}
	return -1
}

func binarySearch(D TabUser, N int, cariID int) int {
	var left, right, mid int
	left = 0
	right = N - 1

	for left <= right {
		mid = (left + right) / 2
		if D[mid].ID == cariID {
			return mid
		} else if D[mid].ID < cariID {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func selectionSortSkor(D *TabUser, N int) {
	var pass, idx, i int
	var temp User

	for pass = 0; pass < N-1; pass++ {
		idx = pass
		for i = pass + 1; i < N; i++ {
			if (*D)[i].Skor < (*D)[idx].Skor {
				idx = i
			}
		}
		temp = (*D)[pass]
		(*D)[pass] = (*D)[idx]
		(*D)[idx] = temp
	}
}

func insertionSortTanggal(D *TabUser, N int) {
	var pass int
	var temp User
	var j int

	for pass = 1; pass < N; pass++ {
		temp = (*D)[pass]
		j = pass - 1

		for j >= 0 && tanggalToInt((*D)[j].Tanggal) > tanggalToInt(temp.Tanggal) {
			(*D)[j+1] = (*D)[j]
			j--
		}
		(*D)[j+1] = temp
	}
}

func tampilkan5Terakhir(D TabUser, N int) {
	fmt.Println("\n=== 5 Hasil Terakhir Self-Assessment ===")

	if N == 0 {
		fmt.Println("Belum ada data.")
	} else {
		start := N - 5
		if start < 0 {
			start = 0
		}

		fmt.Println("ID\tNama\tSkor")
		for i := N - 1; i >= start; i-- {
			fmt.Printf("%03d\t%s\t%d\n", D[i].ID, D[i].Nama, D[i].Skor)
		}
	}
}

func tanggalToInt(tgl string) int {
	var dd, mm, yyyy int
	// tgl format "dd-mm-yyyy"
	// kita ambil yyyy, mm, dd lalu gabungkan jadi int yyyymmdd
	dd = (int(tgl[0]-'0')*10 + int(tgl[1]-'0')) //27 -> 2 = 50 - 48 = 2 * 10 = 20
	mm = (int(tgl[3]-'0')*10 + int(tgl[4]-'0'))
	yyyy = (int(tgl[6]-'0')*1000 + int(tgl[7]-'0')*100 + int(tgl[8]-'0')*10 + int(tgl[9]-'0'))
	return yyyy*10000 + mm*100 + dd
}

func rataRataBulanTerakhir(D TabUser, N int, hariIni string) float64 {
	var sum, count int
	hariIniInt := tanggalToInt(hariIni)

	// Ambil data yang tanggalnya dalam 30 hari terakhir
	for i := 0; i < N; i++ {
		tglInt := tanggalToInt(D[i].Tanggal)
		if hariIniInt-tglInt >= 0 && hariIniInt-tglInt <= 30 {
			sum += D[i].Skor
			count++
		}
	}

	if count == 0 {
		return 0
	}
	return float64(sum) / float64(count)
}

func tambahData(D *TabUser, ID *TabInt, N *int) {
	var u User
    var nama, tanggal string
    var inputID int
    var pertanyaan TabPertanyaan
	var duplikat int = 0

    fmt.Println()

    if *N >= len(*D) {
        fmt.Println("Data sudah penuh, tidak bisa tambah lagi.")
    } else {

        fmt.Print("Masukkan ID baru: ")
        fmt.Scan(&inputID)

        // Cek apakah ID sudah ada agar tidak duplikat
        for i := 0; i < *N; i++ {
            if (*D)[i].ID == inputID {
                duplikat++
            }
        }

        if duplikat > 0 {
            fmt.Println("ID sudah ada, tidak bisa tambah data.")
        } else {
            fmt.Print("Masukkan nama: ")
            fmt.Scan(&nama)

            fmt.Print("Masukkan tanggal (dd-mm-yyyy): ")
            fmt.Scan(&tanggal)

            // Inisialisasi pertanyaan
            inisialisasiPertanyaan(&pertanyaan)

            // Minta user isi skor dengan pertanyaan
            skor := pertanyaanSkor(pertanyaan)

            u.ID = inputID
            u.Nama = nama
            u.Skor = skor
            u.Tanggal = tanggal

            (*D)[*N] = u
            (*ID)[*N] = inputID
            (*N)++

            fmt.Println("Data berhasil ditambahkan.")
        }
    }
}
