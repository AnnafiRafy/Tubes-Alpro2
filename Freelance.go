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
	fmt.Println("Masukkan jumlah data yang ingin dimasukkan: ")
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
		} else if *pilMenu == 2 {
			perbaikiStatus(&data, nProyek)
		} else if *pilMenu == 3 {
			menuCariProyek(&data, nProyek)
		} else if *pilMenu == 4 {
			menuList(data, nProyek)
		} else if *pilMenu == 5 {
			tampilProyek(&data, nProyek)
		} else if *pilMenu == 6 {
			fmt.Println()
			fmt.Println("-------------------------------------------")
			fmt.Println("Terimakasih telah menggunakan aplikasi ini.")
			fmt.Println("-------------------------------------------")
		}
	}
}

// pilihan menu no 1
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

	for i = 0; i < n; i++ { // untuk mencari id proyek yang ingin diubah
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

//menu lanjutan dari menu 1, untuk menghapus data proyek menggunakan Id Proyek
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

	for i = 0; i < *n; i++ { //untuk mencari Id dari proyek yang ingin dihapus
		if A[i].idProyek == hapusId {
			found = i
		}
	}

	if found != -1 { //jika id yang dicari ketemu
		for i = found; i < *n-1; i++ {
			A[i] = A[i+1] //menghapus proyek dengan cara menimpa
		}
		*n = *n - 1
		fmt.Println("Proyek berhasil dihapus.")
	} else { //jika id tidak ditemukan
		fmt.Println()
		fmt.Println("Proyek dengan ID tersebut tidak ditemukan.")
	}
}

//pilihan menu no 2
func perbaikiStatus(A *arrTrack, n int) {
	var cariId, i int
	var found bool

	found = false

	fmt.Println()
	fmt.Println("--------------------------------------------")
	fmt.Println("          Perbaiki Status Proyek            ")
	fmt.Println("--------------------------------------------")
	fmt.Print("Masukkan ID Proyek yang ingin anda perbaiki : ")
	fmt.Scan(&cariId)

	//seq search
	for i = 0; i < n; i++ {
		if A[i].idProyek == cariId {
			found = true
			fmt.Println()
			fmt.Println("Proyek Ditemukan")
			fmt.Println("Nama client :", A[i].client)
			fmt.Println("Nama Proyek :", A[i].projek)
			fmt.Println("Status saat ini :", A[i].status)

			fmt.Print("Masukkan status baru (Progress/Selesai/Pending) : ")
			fmt.Scan(&A[i].status)

			fmt.Println()
			fmt.Println("Status telah diperbarui.")
		}
	}
	if !found {
		fmt.Println()
		fmt.Println("Proyek dengan ID tersebut tidak ditemukan.")
	}
}

//pilihan menu no 3
func menuCariProyek(A *arrTrack, n int) {
	var pilMenu int

	for pilMenu != 4 {
		fmt.Println()
		fmt.Println("-----------------------------------------")
		fmt.Println("              Menu Cari Proyek           ")
		fmt.Println("-----------------------------------------")
		fmt.Println("Pilih ingin mencari berdasarkan apa :")
		fmt.Println("1. ID Proyek")
		fmt.Println("2. Nama Klien")
		fmt.Println("3. Nama Proyek")
		fmt.Println("4. Kembali")
		fmt.Print("Pilih menu (1/2/3/4): ")
		fmt.Scan(&pilMenu)

		if pilMenu == 1 {
			berdasarkanId(*A, n)
		} else if pilMenu == 2 {
			berdasarkanKlien(*A, n)
		} else if pilMenu == 3 {
			berdasarkanProyek(*A, n)
		}
	}
}

//menu lanjutan dari menu 3, mencari proyek berdasarkan Id
func berdasarkanId(A arrTrack, n int) {
	var cariId, left, right, mid int
	var found bool

	found = false
	left = 0
	right = n - 1

	fmt.Print("Masukkan ID Proyek yang ingin dicari: ")
	fmt.Scan(&cariId)

	urutkanId(&A, n)

	//binary search
	for left <= right && !found {
		mid = (left + right) / 2
		if A[mid].idProyek == cariId {
			found = true
		} else if A[mid].idProyek < cariId {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	if found {
		fmt.Println()
		fmt.Println("Proyek dengan ID :", A[mid].idProyek)
		fmt.Println("------------------------------")
		fmt.Println("Nama Client  :", A[mid].client)
		fmt.Println("Nama Proyek  :", A[mid].projek)
		fmt.Println("Status Proyek:", A[mid].status)
		fmt.Println("Gaji         :", A[mid].salary)
		fmt.Println("Deadline     :", A[mid].hari, "-", A[mid].bulan)
		fmt.Println("------------------------------")
	} else {
		fmt.Println()
		fmt.Println("Proyek dengan ID tersebut tidak ditemukan.")
	}
}

//procedure untuk mengurutkan id secara ascending, menggunakan insertion sort
func urutkanId(A *arrTrack, n int) {
	var i, pass int
	var temp tracking

	pass = 1

	for pass < n {
		i = pass
		temp = A[pass]
		for i > 0 && A[i-1].idProyek > temp.idProyek {
			A[i] = A[i-1]
			i--
		}
		A[i] = temp
		pass++
	}
}

//menu lanjutan dari menu 3, mencari proyek berdasarkan nama klien
func berdasarkanKlien(A arrTrack, n int) {
	var cariKlien string
	var i int
	var found bool

	found = false

	fmt.Println("(Gunakan tanda '_' sebagai spasi jika Nama Client lebih dari 1 kata)")
	fmt.Print("Masukkan nama klien yang ingin dicari: ")
	fmt.Scan(&cariKlien)

	//seq search
	for i = 0; i < n; i++ {
		if A[i].client == cariKlien {
			found = true
			fmt.Println()
			fmt.Println("Proyek dengan Nama Klien :", A[i].client)
			fmt.Println("------------------------------")
			fmt.Println("ID Proyek    :", A[i].idProyek)
			fmt.Println("Nama Proyek  :", A[i].projek)
			fmt.Println("Status Proyek:", A[i].status)
			fmt.Println("Gaji         :", A[i].salary)
			fmt.Println("Deadline     :", A[i].hari, "-", A[i].bulan)
			fmt.Println("------------------------------")
		}
	}

	if !found {
		fmt.Println()
		fmt.Println("Proyek dengan nama klien tersebut tidak ditemukan.")
	}
}

//menu lanjutan dari menu 3, mencari proyek berdasarkan nama proyek
func berdasarkanProyek(A arrTrack, n int) {
	var cariProyek string
	var i int
	var found bool

	found = false

	fmt.Println("(Gunakan tanda '_' sebagai spasi jika Nama Proyek lebih dari 1 kata)")
	fmt.Print("Masukkan nama proyek yang ingin dicari: ")
	fmt.Scan(&cariProyek)

	//seq search
	for i = 0; i < n; i++ {
		if A[i].projek == cariProyek {
			found = true
			fmt.Println()
			fmt.Println("Proyek dengan Nama Proyek :", A[i].projek)
			fmt.Println("------------------------------")
			fmt.Println("ID Proyek    :", A[i].idProyek)
			fmt.Println("Nama Klien   :", A[i].client)
			fmt.Println("Status Proyek:", A[i].status)
			fmt.Println("Gaji         :", A[i].salary)
			fmt.Println("Deadline     :", A[i].hari, "-", A[i].bulan)
			fmt.Println("------------------------------")
		}
	}

	if !found {
		fmt.Println()
		fmt.Println("Proyek dengan nama proyek tersebut tidak ditemukan.")
	}
}

//menu nomor 4
func menuList(A arrTrack, n int) {
	var pil string
	var idxGaji, idxDeadline, i int

	fmt.Println()
	fmt.Println("--------------------------------------------")
	fmt.Println("      Apa yang ingin anda tampilkan?	 	")
	fmt.Println("--------------------------------------------")
	fmt.Println()
	fmt.Println("a. Deadline terdekat")
	fmt.Println("b. Bayaran tertinggi")
	fmt.Println("c. Data deadline terdekat - terjauh")
	fmt.Println("d. Data deadline terjauh - terdekat")
	fmt.Println("e. Data bayaran terendah - tertinggi")
	fmt.Println("f. Data bayaran tertinggi - terendah")
	fmt.Print("Pilihan: ")
	fmt.Scan(&pil)

	if pil == "a" {
		idxDeadline = deadline(A, n)
		if idxDeadline != -1 {
			fmt.Printf("\nDeadline Terdekat:\nID Proyek: %d\nClient: %s\nProyek: %s\nDeadline: %02d-%02d\n",
				A[idxDeadline].idProyek, A[idxDeadline].client, A[idxDeadline].projek, A[idxDeadline].hari, A[idxDeadline].bulan)
		} else {
			fmt.Println("Tidak ada data deadline.")
		}
	}

	if pil == "b" {
		idxGaji = bayaranMax(A, n)
		if idxGaji != -1 {
			fmt.Printf("\nBayaran Tertinggi:\nID Proyek: %d\nClient: %s\nProyek: %s\nGaji: %d\n",
				A[idxGaji].idProyek, A[idxGaji].client, A[idxGaji].projek, A[idxGaji].salary)
		} else {
			fmt.Println("Data tidak tersedia.")
		}
	}

	if pil == "c" {
		deadlineAsc(&A, n)
		if n > 0 {
			fmt.Println("\nData deadline urut terdekat - terjauh:")
			for i = 0; i < n; i++ {
				fmt.Println("------------------------------")
				fmt.Printf("Data ke - %d", i+1)
				fmt.Printf("\nID Proyek: %d\nClient: %s\nProyek: %s\nDeadline: %02d-%02d\nGaji: %d\n",
					A[i].idProyek, A[i].client, A[i].projek, A[i].hari, A[i].bulan, A[i].salary)
				fmt.Println("------------------------------")
			}
		} else {
			fmt.Println("Tidak ada data.")
		}
	}

	if pil == "d" {
		deadlineDesc(&A, n)
		if n > 0 {
			fmt.Println("\nData deadline urut terjauh - terdekat:")
			for i = 0; i < n; i++ {
				fmt.Println("------------------------------")
				fmt.Printf("Data ke - %d", i+1)
				fmt.Printf("\nID Proyek: %d\nClient: %s\nProyek: %s\nDeadline: %02d-%02d\nGaji: %d\n",
					A[i].idProyek, A[i].client, A[i].projek, A[i].hari, A[i].bulan, A[i].salary)
				fmt.Println("------------------------------")
			}
		} else {
			fmt.Println("Tidak ada data.")
		}
	}

	if pil == "e" {
		bayarAsc(&A, n)
		if n > 0 {
			fmt.Println("\nData bayaran urut terendah - tertinggi:")
			for i = 0; i < n; i++ {
				fmt.Println("------------------------------")
				fmt.Printf("Data ke - %d", i+1)
				fmt.Printf("\nID Proyek: %d\nClient: %s\nProyek: %s\nDeadline: %02d-%02d\nGaji: %d\n",
					A[i].idProyek, A[i].client, A[i].projek, A[i].hari, A[i].bulan, A[i].salary)
				fmt.Println("------------------------------")
			}
		} else {
			fmt.Println("Tidak ada data.")
		}
	}

	if pil == "f" {
		bayarDesc(&A, n)
		if n > 0 {
			fmt.Println("\nData bayaran urut tertinggi - terendah:")
			for i = 0; i < n; i++ {
				fmt.Println("------------------------------")
				fmt.Printf("Data ke - %d", i+1)
				fmt.Printf("\nID Proyek: %d\nClient: %s\nProyek: %s\nDeadline: %02d-%02d\nGaji: %d\n",
					A[i].idProyek, A[i].client, A[i].projek, A[i].hari, A[i].bulan, A[i].salary)
				fmt.Println("------------------------------")
			}
		} else {
			fmt.Println("Tidak ada data.")
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

// Insertion berdasarkan bayaran terendah
func bayarAsc(A *arrTrack, n int) {
	var pass, i int
	var temp tracking

	pass = 1
	for pass <= n-1 {
		i = pass
		temp = A[pass]

		for i > 0 && temp.salary < A[i-1].salary {
			A[i] = A[i-1]
			i = i - 1
		}

		A[i] = temp
		pass = pass + 1
	}
}

// Insertion berdasarkan bayaran tertinggi
func bayarDesc(A *arrTrack, n int) {
	var pass, i int
	var temp tracking

	pass = 1
	for pass <= n-1 {
		i = pass
		temp = A[pass]

		for i > 0 && temp.salary > A[i-1].salary {
			A[i] = A[i-1]
			i = i - 1
		}

		A[i] = temp
		pass = pass + 1
	}
}

// Sorting berdasarkan deadline terdekat
func deadlineAsc(A *arrTrack, n int) {
	var pass, i, idx int
	var temp tracking

	for pass = 1; pass < n; pass++ {
		idx = pass - 1
		for i = pass; i < n; i++ {
			if A[i].bulan < A[idx].bulan ||
				(A[i].bulan == A[idx].bulan && A[i].hari < A[idx].hari) {
				idx = i
			}
		}
		temp = A[pass-1]
		A[pass-1] = A[idx]
		A[idx] = temp
	}
}

// Sorting berdasarkan deadline terjauh
func deadlineDesc(A *arrTrack, n int) {
	var pass, i, idx int
	var temp tracking

	for pass = 1; pass < n; pass++ {
		idx = pass - 1
		for i = pass; i < n; i++ {
			if A[i].bulan > A[idx].bulan ||
				(A[i].bulan == A[idx].bulan && A[i].hari > A[idx].hari) {
				idx = i
			}
		}
		temp = A[pass-1]
		A[pass-1] = A[idx]
		A[idx] = temp
	}
}

//menu 5
func tampilProyek(A *arrTrack, n int) {
	var i int

	if n == 0 {
		fmt.Println("Data kosong")
		return
	}

	for i = 0; i < n; i++ {
		fmt.Println()
		fmt.Println("------------------------------")
		fmt.Println("Data ke-", i+1)
		fmt.Println("ID Proyek:", A[i].idProyek)
		fmt.Println("Client   :", A[i].client)
		fmt.Println("Proyek   :", A[i].projek)
		fmt.Println("Status   :", A[i].status)
		fmt.Println("Deadline :", A[i].hari, "-", A[i].bulan)
		fmt.Println("Gaji     :", A[i].salary)
		fmt.Println("------------------------------")
	}
}
