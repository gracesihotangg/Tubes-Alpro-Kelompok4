package main

import "fmt"

type Twisata struct {
	nama, kategori string
	jarak          float64
	biaya          int
	fasilitas      DataFasilitas
}

type DataPariwisata struct {
	data      ArrWisata
	totalData int
}

type DataFasilitas struct {
	data      Arrfasilitas
	totalData int
}

type ArrWisata [999]Twisata

type Arrfasilitas [99]string

func inputTempatWisata(W *DataPariwisata) {
	var Namafasilitas string
	var wisata Twisata

	fmt.Println("Input Nama Tempat Wisata : ")
	fmt.Scan(&wisata.nama)
	for wisata.nama != "-" {
		fmt.Println("Input kategori : ")
		fmt.Scan(&wisata.kategori)
		fmt.Println("Input biaya masuk : ")
		fmt.Scan(&wisata.biaya)
		fmt.Println("Input jarak : ")
		fmt.Scan(&wisata.jarak)
		fmt.Println("")
		fmt.Println("Apa Saja Fasilitas yang ada disana ?")
		fmt.Scan(&Namafasilitas)
		for j := 0; Namafasilitas != "-"; j++ {
			wisata.fasilitas.data[wisata.fasilitas.totalData] = Namafasilitas
			wisata.fasilitas.totalData++
			fmt.Scan(&Namafasilitas)
		}
		W.data[W.totalData] = wisata
		W.totalData++
		wisata = Twisata{}
		fmt.Println("Input Nama Tempat Wisata : ")
		fmt.Scan(&wisata.nama)
		fmt.Println(" ")
	}
}

func cariIndeksTempatWisata(W DataPariwisata, nama string) int {
	for i := 0; i < W.totalData; i++ {
		if W.data[i].nama == nama {
			return i
		}
	}
	return -1
}

func ubahTempatWisata(W *DataPariwisata) {
	var nama string
	fmt.Print("Masukkan Nama Yang Ingin Diubah: ")
	fmt.Scan(&nama)
	var idx = cariIndeksTempatWisata(*W, nama)
	if idx == -1 {
		fmt.Println("Data Tidak Ditemukan")
	} else {
		var wisata = W.data[idx]
		var Namafasilitas string
		fmt.Print("Input Nama Tempat Wisata Baru: ")
		fmt.Scan(&wisata.nama)
		fmt.Println("Input kategori Baru : ")
		fmt.Scan(&wisata.kategori)
		fmt.Println("Input Biaya Masuk Baru : ")
		fmt.Scan(&wisata.biaya)
		fmt.Println("Input Jarak : ")
		fmt.Scan(&wisata.jarak)
		fmt.Println("")
		fmt.Println("Apa Saja Fasilitas yang ada disana ?")
		fmt.Scan(&Namafasilitas)
		wisata.fasilitas = DataFasilitas{}
		for j := 0; Namafasilitas != "-"; j++ {
			wisata.fasilitas.data[wisata.fasilitas.totalData] = Namafasilitas
			wisata.fasilitas.totalData++
			fmt.Scan(&Namafasilitas)
		}
		W.data[idx] = wisata
	}
}

func PrintTempatWisata(W DataPariwisata) {
	for i := 0; i < W.totalData; i++ {
		fmt.Println("Nama Tempat Pariwisata : ", W.data[i].nama)
		fmt.Println("Nama Kategori : ", W.data[i].kategori)
		fmt.Println("Jarak Tempuh  : ", W.data[i].jarak)
		fmt.Println("Biaya Masuk : ", W.data[i].biaya)
		fmt.Println("Fasilitas  : ")
		for j := 0; j < W.data[i].fasilitas.totalData; j++ {
			fmt.Print("\t", W.data[i].fasilitas.data[j])
			fmt.Println(" ")
		}
		fmt.Println(" ")
	}
}

func selectionSort(W *DataPariwisata, flag, status string) {
	for pass := 1; pass <= W.totalData-1; pass++ {
		idx := pass - 1
		for i := pass; i < W.totalData; i++ {
			switch flag {
			case "1":
				if status == "1" && W.data[idx].nama > W.data[i].nama {
					idx = i
				} else if status == "2" && W.data[idx].nama < W.data[i].nama {
					idx = i
				}
			case "2":
				if status == "1" && W.data[idx].kategori > W.data[i].kategori {
					idx = i
				} else if status == "2" && W.data[idx].kategori < W.data[i].kategori {
					idx = i
				}
			case "3":
				if status == "1" && W.data[idx].jarak > W.data[i].jarak {
					idx = i
				} else if status == "2" && W.data[idx].jarak < W.data[i].jarak {
					idx = i
				}
			case "4":
				if status == "1" && W.data[idx].biaya > W.data[i].biaya {
					idx = i
				} else if status == "2" && W.data[idx].biaya < W.data[i].biaya {
					idx = i
				}

			}
		}
		W.data[pass-1], W.data[idx] = W.data[idx], W.data[pass-1]
	}

}

func cariDataTempatWisata(W DataPariwisata, flag string) int {
	var kiri, tengah, kanan int
	kanan = W.totalData - 1
	switch flag {
	case "1":
		var search string
		fmt.Print("Masukkan nama tempat: ")
		fmt.Scan(&search)
		for kiri <= kanan {
			tengah = (kiri + kanan) / 2
			if W.data[tengah].nama == search {
				return tengah
			} else if W.data[tengah].nama > search {
				kanan = tengah - 1
			} else {
				kiri = tengah + 1
			}
		}
	case "2":
		var search string
		fmt.Print("Masukkan kategori: ")
		fmt.Scan(&search)
		for kiri <= kanan {
			tengah = (kiri + kanan) / 2
			if W.data[tengah].kategori == search {
				return tengah
			} else if W.data[tengah].kategori > search {
				kanan = tengah - 1
			} else {
				kiri = tengah + 1
			}
		}
	case "3":
		var search float64
		fmt.Print("Masukkan jarak: ")
		fmt.Scan(&search)
		for kiri <= kanan {
			tengah = (kiri + kanan) / 2
			if W.data[tengah].jarak == search {
				return tengah
			} else if W.data[tengah].jarak > search {
				kanan = tengah - 1
			} else {
				kiri = tengah + 1
			}
		}
	case "4":
		var search int
		fmt.Print("Masukkan biaya: ")
		fmt.Scan(&search)
		for kiri <= kanan {
			tengah = (kiri + kanan) / 2
			if W.data[tengah].biaya == search {
				return tengah
			} else if W.data[tengah].biaya > search {
				kanan = tengah - 1
			} else {
				kiri = tengah + 1
			}
		}
	}

	return -1
}

func HapusDataTempatWisata(W *DataPariwisata) {
	var nama string
	fmt.Print("Masukkan nama yang ingin dihapus")
	fmt.Scan(&nama)
	var idx int
	idx = cariIndeksTempatWisata(*W, nama)
	if idx == -1 {
		fmt.Println("Data Tidak Ditemukan")
	} else {
		W.totalData--
		for i := idx; i < W.totalData; i++ {
			W.data[i] = W.data[i+1]
		}
	}
}

func MenuAdmin(W *DataPariwisata) {
	var pilihMenu, flag, status string
	for {
		fmt.Println("=== Aplikasi Pariwisata ===")
		fmt.Println("1. Input Data Pariwisata ")
		fmt.Println("2. Tampilkan Keseluruhan Data")
		fmt.Println("3. Mencari Tempat Pariwisata")
		fmt.Println("4. Ubah Data Tempat Pariwisata")
		fmt.Println("5. Hapus Data Tempat Pariwisata")
		fmt.Println("0. Kembali")
		fmt.Println("==========================")

		for {
			fmt.Print("Pilih menu: ")
			fmt.Scan(&pilihMenu)
			if pilihMenu == "1" || pilihMenu == "2" || pilihMenu == "3" || pilihMenu == "4" || pilihMenu == "5" || pilihMenu == "0" {
				break
			}
		}
		switch pilihMenu {
		case "1":
			inputTempatWisata(W)
		case "2":
			opsiSortKategori()
			fmt.Scan(&flag)
			opsiSort()
			fmt.Scan(&status)
			var temp = *W
			selectionSort(&temp, flag, status)
			PrintTempatWisata(temp)
		case "3": 
			var temp = *W
			opsiCari()
			fmt.Scan(&flag)
			selectionSort(&temp, flag, "1")
			var idx = cariDataTempatWisata(temp, flag)
			if idx == -1 {
				fmt.Println("Data tidak Ditemukan")
			} else {
				printSatuData(temp.data[idx])
			}
		case "4":
			ubahTempatWisata(W)
		case "5":
			HapusDataTempatWisata(W)
		case "0":
			return
		}
	}
}

func printSatuData(W Twisata) {
	fmt.Println("Nama Tempat Pariwisata : ", W.nama)
	fmt.Println("Nama Kategori : ", W.kategori)
	fmt.Println("Jarak Tempuh  : ", W.jarak)
	fmt.Println("Biaya Masuk : ", W.biaya)
	fmt.Println("Fasilitas  : ")
	for j := 0; j < W.fasilitas.totalData; j++ {
		fmt.Print("\t", W.fasilitas.data[j])
		fmt.Println(" ")
	}
}

func opsiCari() {
	fmt.Println("Silahkan Pilih Opsi Dibawah ini ")
	fmt.Println("1. Berdasarkan Nama")
	fmt.Println("2. Berdasarkan Kategori")
	fmt.Println("3. Berdasarkan Jarak")
	fmt.Println("4. Berdasarkan Baya")
	fmt.Println("Pilih Menu : ")
}

func opsi3() {
	fmt.Println("Ingin Mencari Berdasarkan Apa?")
	fmt.Println("1. Berdasarkan Kategori")
	fmt.Println("2. Berdasarkan Biaya")
	fmt.Println("3. Berdasarkan Jarak")
	fmt.Println("4. Kembali ")
	fmt.Println("Pilih menu: ")
}

func opsiSortKategori() {
	fmt.Println("Silahkan Pilih Opsi Dibawah ini ")
	fmt.Println("1. Urutkan Data Berdasarkan Nama")
	fmt.Println("2. Urutkan Data Berdasarkan Kategori")
	fmt.Println("3. Urutkan Data Berdasarkan Jarak")
	fmt.Println("4. Urutkan Data Berdasarkan Baya")
	fmt.Println("5. Kembali")
	fmt.Println("Pilih Menu : ")
}

func opsiSort() {
	fmt.Println(" Silahkan Pilih Opsi Dibawah Ini ")
	fmt.Println("1. ascending ")
	fmt.Println("2. Descanding ")
	fmt.Println(" Pilih  Menu ")
}

func insertionSort(W *DataPariwisata, flag, status string) {
	for pass := 1; pass < W.totalData; pass++ {
		var i int = pass
		var pariwisata Twisata = W.data[i]
		switch flag {
		case "1":
			if status == "1" {
				for i > 0 && W.data[i-1].nama > pariwisata.nama {
					W.data[i] = W.data[i-1]
					i--
				}
			} else {
				for i > 0 && W.data[i-1].nama < pariwisata.nama {
					W.data[i] = W.data[i-1]
					i--
				}
			}
		case "2":
			if status == "1" {
				for i > 0 && (W.data[i-1].kategori > pariwisata.kategori) {
					W.data[i] = W.data[i-1]
					i--
				}
			} else {
				for i > 0 && W.data[i-1].kategori < pariwisata.kategori {
					W.data[i] = W.data[i-1]
					i--
				}
			}
		case "3":
			if status == "1" {
				for i > 0 && (W.data[i-1].jarak) > (pariwisata.jarak) {
					W.data[i] = W.data[i-1]
					i--
				}
			} else {
				for i > 0 && W.data[i-1].jarak < pariwisata.jarak {
					W.data[i] = W.data[i-1]
					i--
				}
			}
		case "4":
			if status == "1" {
				for i > 0 && (W.data[i-1].biaya) > (pariwisata.biaya) {
					W.data[i] = W.data[i-1]
					i--
				}
			} else {
				for i > 0 && W.data[i-1].biaya < pariwisata.biaya {
					W.data[i] = W.data[i-1]
					i--
				}
			}
		}
		W.data[i] = pariwisata
	}
}

func MenuPengunjung(W DataPariwisata) {
	var pilihanMenu string
	for {

		fmt.Println("=== Aplikasi Pengguna ===")
		fmt.Println("1. Tampilkan tempat Pariwisata")
		fmt.Scan(&pilihanMenu)
		switch pilihanMenu {
		case "1":
			var flag, status string
			var temp DataPariwisata = W
			opsiSortKategori()
			fmt.Scan(&flag)
			opsiSort()
			fmt.Scan(&status)
			insertionSort(&temp, flag, status)
			PrintTempatWisata(temp)
		case "0":
			return
		}
	}
}

func main() {
	var W DataPariwisata
	var pilihMenu string
	for {
		fmt.Printf("===== Menu =====\n1. Admin\n2. Pengunjung\n0. Keluar\n=====================\n")
		for true {
			fmt.Print("Pilih menu: ")
			fmt.Scan(&pilihMenu)
			if pilihMenu == "0" || pilihMenu == "1" || pilihMenu == "2" {
				break
			}
		}
		switch pilihMenu {
		case "1":
			MenuAdmin(&W)
		case "2":
			MenuPengunjung(W)
		}
	}
}
