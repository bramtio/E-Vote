package main

import "fmt"

const MAX = 100

type Kandidat struct {
	Nomor    int
	Nama     string
	VisiMisi string
	Suara    int
}

var data [MAX]Kandidat
var n int

func tambahKandidat() {
	if n >= MAX {
		fmt.Println("Data penuh!")
		return
	}

	fmt.Print("Nomor Urut : ")
	fmt.Scan(&data[n].Nomor)

	fmt.Print("Nama Kandidat : ")
	fmt.Scan(&data[n].Nama)

	fmt.Print("Visi Misi : ")
	fmt.Scan(&data[n].VisiMisi)

	data[n].Suara = 0
	n++

	fmt.Println("Data berhasil ditambahkan!")
}

func insertionSortNomor() {
	var temp Kandidat

	for i := 1; i < n; i++ {
		temp = data[i]
		j := i - 1

		for j >= 0 && data[j].Nomor > temp.Nomor {
			data[j+1] = data[j]
			j--
		}

		data[j+1] = temp
	}
}

func selectionSortSuara() {
	for i := 0; i < n-1; i++ {
		max := i

		for j := i + 1; j < n; j++ {
			if data[j].Suara > data[max].Suara {
				max = j
			}
		}

		data[i], data[max] = data[max], data[i]
	}
}

func tampilkanKandidat() {
	if n == 0 {
		fmt.Println("Belum ada data kandidat.")
		return
	}

	var pilih int

	fmt.Println("\n=== TAMPILKAN DATA ===")
	fmt.Println("1. Berdasarkan Nomor Urut")
	fmt.Println("2. Berdasarkan Suara Terbanyak")
	fmt.Print("Pilihan : ")
	fmt.Scan(&pilih)

	if pilih == 1 {
		insertionSortNomor()
	} else if pilih == 2 {
		selectionSortSuara()
	} else {
		fmt.Println("Pilihan tidak valid.")
		return
	}

	fmt.Println("\n===== DATA KANDIDAT =====")

	for i := 0; i < n; i++ {
		fmt.Println("Nomor Urut :", data[i].Nomor)
		fmt.Println("Nama       :", data[i].Nama)
		fmt.Println("Visi Misi  :", data[i].VisiMisi)
		fmt.Println("Suara      :", data[i].Suara)
		fmt.Println("-------------------------")
	}
}

func ubahKandidat() {
	var nomor int

	fmt.Print("Masukkan nomor kandidat yang akan diubah : ")
	fmt.Scan(&nomor)

	for i := 0; i < n; i++ {
		if data[i].Nomor == nomor {

			fmt.Print("Nama Baru : ")
			fmt.Scan(&data[i].Nama)

			fmt.Print("Visi Misi Baru : ")
			fmt.Scan(&data[i].VisiMisi)

			fmt.Println("Data berhasil diubah!")
			return
		}
	}

	fmt.Println("Data tidak ditemukan.")
}

func hapusKandidat() {
	var nomor int

	fmt.Print("Masukkan nomor kandidat yang akan dihapus : ")
	fmt.Scan(&nomor)

	for i := 0; i < n; i++ {
		if data[i].Nomor == nomor {

			for j := i; j < n-1; j++ {
				data[j] = data[j+1]
			}

			n--
			fmt.Println("Data berhasil dihapus!")
			return
		}
	}

	fmt.Println("Data tidak ditemukan.")
}

func voting() {
	var nomor int

	fmt.Print("Masukkan nomor kandidat yang dipilih : ")
	fmt.Scan(&nomor)

	for i := 0; i < n; i++ {
		if data[i].Nomor == nomor {
			data[i].Suara++
			fmt.Println("Voting berhasil!")
			return
		}
	}

	fmt.Println("Kandidat tidak ditemukan.")
}

func cariNomorUrut() {
	if n == 0 {
		fmt.Println("Belum ada data kandidat.")
		return
	}

	insertionSortNomor()

	var nomor int

	fmt.Print("Masukkan nomor urut : ")
	fmt.Scan(&nomor)

	low := 0
	high := n - 1

	for low <= high {

		mid := (low + high) / 2

		if data[mid].Nomor == nomor {

			fmt.Println("\n=== DATA DITEMUKAN ===")
			fmt.Println("Nomor     :", data[mid].Nomor)
			fmt.Println("Nama      :", data[mid].Nama)
			fmt.Println("Visi Misi :", data[mid].VisiMisi)
			fmt.Println("Suara     :", data[mid].Suara)
			return
		}

		if nomor < data[mid].Nomor {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	fmt.Println("Data tidak ditemukan.")
}

func cariSuaraTerbanyak() {
	if n == 0 {
		fmt.Println("Belum ada data kandidat.")
		return
	}

	max := 0

	for i := 1; i < n; i++ {
		if data[i].Suara > data[max].Suara {
			max = i
		}
	}

	fmt.Println("\n=== KANDIDAT DENGAN SUARA TERBANYAK ===")
	fmt.Println("Nomor     :", data[max].Nomor)
	fmt.Println("Nama      :", data[max].Nama)
	fmt.Println("Visi Misi :", data[max].VisiMisi)
	fmt.Println("Suara     :", data[max].Suara)
}

func searchKandidat() {
	var pilih int

	fmt.Println("\n=== SEARCH KANDIDAT ===")
	fmt.Println("1. Cari Berdasarkan Nomor Urut")
	fmt.Println("2. Cari Kandidat Dengan Suara Terbanyak")
	fmt.Print("Pilihan : ")
	fmt.Scan(&pilih)

	switch pilih {
	case 1:
		cariNomorUrut()
	case 2:
		cariSuaraTerbanyak()
	default:
		fmt.Println("Pilihan tidak valid.")
	}
}

func statistik() {
	total := 0

	for i := 0; i < n; i++ {
		total += data[i].Suara
	}

	fmt.Println("\n===== STATISTIK =====")

	for i := 0; i < n; i++ {

		persen := 0.0

		if total > 0 {
			persen = float64(data[i].Suara) * 100 / float64(total)
		}

		fmt.Printf("%s : %d suara (%.2f%%)\n",
			data[i].Nama,
			data[i].Suara,
			persen)
	}

	fmt.Println("Total Suara :", total)
}

func main() {
	var pilih int

	for {
		fmt.Println("\n===== E-VOTING =====")
		fmt.Println("1. Tambah Kandidat")
		fmt.Println("2. Tampilkan Kandidat")
		fmt.Println("3. Ubah Kandidat")
		fmt.Println("4. Hapus Kandidat")
		fmt.Println("5. Voting")
		fmt.Println("6. Search Kandidat")
		fmt.Println("7. Statistik")
		fmt.Println("0. Keluar")

		fmt.Print("Pilih Menu : ")
		fmt.Scan(&pilih)

		switch pilih {

		case 1:
			tambahKandidat()

		case 2:
			tampilkanKandidat()

		case 3:
			ubahKandidat()

		case 4:
			hapusKandidat()

		case 5:
			voting()

		case 6:
			searchKandidat()

		case 7:
			statistik()

		case 0:
			fmt.Println("Program selesai.")
			return

		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}
