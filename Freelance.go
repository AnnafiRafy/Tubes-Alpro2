package main

import "fmt"

const NMAX int = 1000

type tracking struct {
	client, projek, status string
	salary, idProyek       int
	//deadline
	hari, bulan int
}
type arrTrack [NMAX]tracking

var data arrTrack
var nProyek int

func main() {
	var pilihMenu int

	fmt.Println()
	fmt.Println("-----------------------------------")
	fmt.Println("      ~~~ Selamat Datang ~~~	 	")
	fmt.Println("-----------------------------------")
	fmt.Println("Masukkan jumlah data yang ingin anda masukkan: ")
	fmt.Scan(&nProyek)

	bacaData(&data, nProyek)
	menuUtama(&pilihMenu)
}

//procedure untuk scan dan menyimpan data di awal
func bacaData(A *arrTrack, n int) {
	var i int

	fmt.Println("(Gunakan tanda '_' sebagai spasi jika Nama Client / Nama Proyek lebih dari 1 kata)")
	for i = 0; i < n; i++ {
		fmt.Printf("\nData ke - %d \n", i+1)

		fmt.Print("ID Proyek : ")
		fmt.Scan(&A[i].idProyek)

		fmt.Print("Nama Client : ")
		fmt.Scan(&A[i].client)

		fmt.Print("Nama Proyek : ")
		fmt.Scan(&A[i].projek)

		fmt.Print("Status (Progress/Selesai/Pending): ")
		fmt.Scan(&A[i].status)

		fmt.Print("Deadline (Tanggal(DD) Bulan(MM)): ")
		fmt.Scan(&A[i].hari, &A[i].bulan)

		fmt.Print("Gaji: ")
		fmt.Scan(&A[i].salary)

		fmt.Println()
	}
}

//procedure pilihan menu utama
func menuUtama(pilMenu *int) {
	for *pilMenu != 6 {
		fmt.Println()
		fmt.Println("-------------------------------------------")
		fmt.Println(" Aplikasi Manajemen dan Tracking Freelance")
		fmt.Println("-------------------------------------------")
		fmt.Println("Silahakan pilih menu :")
		fmt.Println("1. Edit Proyek")
		fmt.Println("2. Perbaiki Status Proyek")
		fmt.Println("3. Cari Proyek")
		fmt.Println("4. Tampilkan list deadline / bayaran")
		fmt.Println("5. Tampilkan keseluruhan proyek")
		fmt.Println("6. Keluar")
		fmt.Println("Silahkan pilih : (1/2/3/4/5/6)?")

		fmt.Scan(pilMenu)
		if *pilMenu == 1 {
			menuEditProyek(&data, &nProyek)
		} else if *pilMenu == 4 {
			menuList(data, nProyek)
		} else if *pilMenu == 5 {
			tampilProyek(&data, nProyek)
		}
	}

	//else if *pilMenu == 2
	//editStatus(pilMenu)
	//else if *pilMenu == 3
	//cariProyek(pilMenu)
	//else if *pilMenu == 4
	//list(pilMenu)
	//else if *pilMenu == 5
	//show(pilMenu)
}

// pilihan 1
func menuEditProyek(A *arrTrack, n *int) {
	var pilMenu int

	for pilMenu != 4 {
		fmt.Println()
		fmt.Println("-------------------------------------------")
		fmt.Println("	Apa yang ingin anda lakukan?		")
		fmt.Println("-------------------------------------------")
		fmt.Println("Silahkan pilih : ")
		fmt.Println("1. Tambah Proyek")
		fmt.Println("2. Ubah Proyek")
		fmt.Println("3. Hapus Proyek")
		fmt.Println("4. Kembali")
		fmt.Println("Silahkan pilih : (1/2/3/4)?")

		fmt.Scan(&pilMenu)
		if pilMenu == 1 {
			tambahProyek(A, n)
		} else if pilMenu == 2 {
			ubahProyek(A, *n)
		} else if pilMenu == 3 {
			hapusProyek(A, n)
		}
		//else if pilMenu == 2
		//else if pilMenu == 3
	}
}

//menu lanjutan dari menu 1 untuk menambahakn proyek
func tambahProyek(A *arrTrack, n *int) {
	var jumlah, i int

	fmt.Println()
	fmt.Println("-------------------------------------------------------------------")
	fmt.Println("	Silahkan isi data proyek yang ingin ditambahkan 		")
	fmt.Println("-------------------------------------------------------------------")
	fmt.Println("(Gunakan tanda '_' sebagai spasi jika Nama Client / Nama Proyek lebih dari 1 kata)")
	fmt.Println("Jumlah proyek yang ingin ditambahkan : ")
	fmt.Scan(&jumlah)
	//agar tidak melebihi batas array
	if jumlah+*n > NMAX {
		fmt.Println("Jumlah proyek melebihi batas maksimal")
		return
	}

	for i = 0; i < jumlah && *n < NMAX; i++ {
		fmt.Printf("\nProyek ke - %d \n", i+1)
		fmt.Print("ID Proyek : ")
		fmt.Scan(&A[*n].idProyek)

		fmt.Print("Nama Client : ")
		fmt.Scan(&A[*n].client)

		fmt.Print("Nama Proyek : ")
		fmt.Scan(&A[*n].projek)

		fmt.Print("Status Proyek (Progress/Selesai/Pending) : ")
		fmt.Scan(&A[*n].status)

		fmt.Print("Deadline (Tanggal(DD) Bulan(MM)) : ")
		fmt.Scan(&A[*n].hari, &A[*n].bulan)

		fmt.Print("Gaji / Bayaran : ")
		fmt.Scan(&A[*n].salary)

		*n++
	}
	fmt.Println("Proyek telah ditambahkan")
}

//menu lanjutan dari menu 1, untuk mengubah data proyek menggunakan Id Proyek
func ubahProyek(A *arrTrack, n int) {
	var i, cariId int
	var found int
	found = -1

	fmt.Println()
	fmt.Println("---------------------------------")
	fmt.Println("	Ubah Data Proyek		")
	fmt.Println("---------------------------------")
	fmt.Print("Masukkan ID Proyek yang ingin anda ubah : ")
	fmt.Scan(&cariId)

	for i = 0; i < n; i++ {
		// jika id yang dicari ketemu
		if A[i].idProyek == cariId {
			found = i
		}
	}
	//jika id yang dicari ketemu
	if found != -1 {
		fmt.Printf("\nData proyek dengan ID %d telah dipilih. \n", cariId)
		fmt.Println()
		fmt.Println("Silahkan masukan data baru : ")
		fmt.Println("(Gunakan tanda '_' sebagai spasi jika Nama Client / Nama Proyek lebih dari 1 kata)")

		fmt.Print("Nama Client (baru) : ")
		fmt.Scan(&A[found].client)

		fmt.Print("Nama Proyek (baru) : ")
		fmt.Scan(&A[found].projek)

		fmt.Print("Status baru (Progress/Selesai/Pending) : ")
		fmt.Scan(&A[found].status)

		fmt.Print("Deadline baru (Tanggal(DD)-Bulan(MM)) : ")
		fmt.Scan(&A[found].hari, A[found].bulan)

		fmt.Print("Gaji (baru) : ")
		fmt.Scan(&A[found].salary)

		fmt.Println()
		fmt.Println("Data proyek berhasil diperbarui.")
	} else { // jika id yang dicari tidak ketemu
		fmt.Println()
		fmt.Println("Proyek dengan ID tersebut tidak ditemukan.")
	}
}

func hapusProyek(A *arrTrack, n *int) {
	var hapusId, i int
	var found int
	found = -1

	fmt.Println()
	fmt.Println("---------------------------------")
	fmt.Println("	Hapus Data Proyek		")
	fmt.Println("---------------------------------")
	fmt.Print("Masukkan ID Proyek yang ingin anda hapus : ")
	fmt.Scan(&hapusId)

	for i = 0; i < *n; i++ {
		if A[i].idProyek == hapusId {
			found = i
		}
	}

	if found != -1 {
		for i = found; i < *n-1; i++ {
			A[i] = A[i+1]
		}
		*n = *n - 1
		fmt.Println("Proyek berhasil dihapus.")
	} else {
		fmt.Println()
		fmt.Println("Proyek dengan ID tersebut tidak ditemukan.")
	}
}

//menu nomor 4
func menuList(A arrTrack, n int) {
	var pil string
	var idxGaji, idxDeadline int

	fmt.Println()
	fmt.Println("--------------------------------------------")
	fmt.Println("      Apa yang ingin anda tampilkan?	 	")
	fmt.Println("--------------------------------------------")
	fmt.Println()
	fmt.Println("a. Deadline terdekat")
	fmt.Println("b. Bayaran tertinggi")
	//fmt.Println("c. Data deadline terdekat dan yang terjauh")
	//fmt.Println("d. Data bayaran dari yang tertinggi ke terendah")
	fmt.Print("Pilihan: ")
	fmt.Scan(&pil)

	if pil == "a" {
		idxDeadline = deadline(A, n)
		if idxDeadline != -1 {
			fmt.Printf("\nDeadline Terdekat:\nClient: %s\nProyek: %s\nDeadline: %02d-%02d\n",
				A[idxDeadline].client, A[idxDeadline].projek, A[idxDeadline].hari, A[idxDeadline].bulan)
		} else {
			fmt.Println("Tidak ada data deadline.")
		}
	}

	if pil == "b" {
		idxGaji = bayaranMax(A, n)
		if idxGaji != -1 {
			fmt.Printf("\nBayaran Tertinggi:\nClient: %s\nProyek: %s\nGaji: %d\n",
				A[idxGaji].client, A[idxGaji].projek, A[idxGaji].salary)
		} else {
			fmt.Println("Data tidak tersedia.")
		}
	}
}

//sequential search
func bayaranMax(A arrTrack, n int) int {
	var maxIdx, i int
	if n == 0 {
		return -1
	}
	maxIdx = 0
	for i = 1; i < n; i++ {
		if A[i].salary > A[maxIdx].salary {
			maxIdx = i
		}
	}
	return maxIdx
}

//sequential search
func deadline(A arrTrack, n int) int {
	var terdekatIdx int
	var i int

	for i = 1; i < n; i++ {
		if A[i].bulan < A[terdekatIdx].bulan ||
			(A[i].bulan == A[terdekatIdx].bulan && A[i].hari < A[terdekatIdx].hari) {
			terdekatIdx = i
		}
	}
	return terdekatIdx
}

// Sorting berdasarkan bayaran tertinggi ke bawah
func bayarTerurut(A *arrTrack, n int) {
	var i, j, maxIdx int
	var temp tracking
	for i = 0; i < n-1; i++ {
		maxIdx = i
		for j = i + 1; j < n; j++ {
			if A[j].salary > A[maxIdx].salary {
				maxIdx = j
			}
		}
		temp = A[i]
		A[i] = A[maxIdx]
		A[maxIdx] = temp
	}
}

// Sorting berdasarkan deadline terdekat
func deadlineUrut(A *arrTrack, n int) {
	var i, j, minIdx int
	var temp tracking
	for i = 0; i < n-1; i++ {
		minIdx = i
		for j = i + 1; j < n; j++ {
			if A[j].hari < A[minIdx].hari {
				minIdx = j
			}
		}
		temp = A[i]
		A[i] = A[minIdx]
		A[minIdx] = temp
	}
}

func tampilProyek(A *arrTrack, n int) {
	var i int

	if n == 0 {
		fmt.Println("Data kosong")
		return
	}

	for i = 0; i < n; i++ {
		fmt.Println("Data ke-", i+1)
		fmt.Println("Client   :", A[i].client)
		fmt.Println("Projek   :", A[i].projek)
		fmt.Println("Status   :", A[i].status)
		fmt.Println("Gaji     :", A[i].salary)
		fmt.Println("ID Proyek:", A[i].idProyek)
		fmt.Println("Deadline :", A[i].hari, "-", A[i].bulan)
		fmt.Println("------------------------------")
	}
}
