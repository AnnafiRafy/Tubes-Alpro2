package main

import "fmt"

const NMAX int = 1000

type tracking struct {
	client, projek, status, deadline string
	salary, idProyek                 int
}
type arrTrack [NMAX]tracking

var data arrTrack
var nProyek int

func main() {
	var pilihMenu int

	bacaData(&data, nProyek)
	menuUtama(&pilihMenu)
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
func menuEditProyek(data *arrTrack, n *int) {
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
			tambahProyek(data, n)
		}
		//else if pilMenu == 2
		//else if pilMenu == 3
	}
}

func bacaData(A *arrTrack, n int) {
	var i int

	fmt.Println()
	fmt.Println("-----------------------------------")
	fmt.Println("      ~~~ Selamat Datang ~~~	 	")
	fmt.Println("-----------------------------------")
	fmt.Println("Masukkan jumlah data yang ingin anda masukkan: ")
	fmt.Scan(&n)

	for i = 0; i < n; i++ {
		fmt.Printf("\nData ke - %d \n", i+1)
		fmt.Println("(Gunakan tanda '_' sebagai spasi jika Nama Client / Nama Proyek lebih dari 1 kata)")

		fmt.Print("ID Proyek : ")
		fmt.Scan(&A[i].idProyek)

		fmt.Print("Nama Client : ")
		fmt.Scan(&A[i].client)

		fmt.Print("Nama Proyek : ")
		fmt.Scan(&A[i].projek)

		fmt.Print("Status (Progress/Selesai/Pending): ")
		fmt.Scan(&A[i].status)

		fmt.Print("Deadline (Tanggal(DD)-Bulan(MM)): ")
		fmt.Scan(&A[i].deadline)

		fmt.Print("Gaji: ")
		fmt.Scan(&A[i].salary)

		fmt.Println()
	}
}

//menu lanjutan untuk menambahakn proyek
func tambahProyek(data *arrTrack, n *int) {
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
		fmt.Scan(&data[*n].idProyek)

		fmt.Print("Nama Client : ")
		fmt.Scan(&data[*n].client)

		fmt.Print("Nama Proyek : ")
		fmt.Scan(&data[*n].projek)

		fmt.Print("Status Proyek (Progress/Selesai/Pending) : ")
		fmt.Scan(&data[*n].status)

		fmt.Print("Gaji / Bayaran : ")
		fmt.Scan(&data[*n].salary)

		*n++
	}
	fmt.Println("Proyek telah ditambahkan")
}

//menu nomor 4
func menuList(data arrTrack, n int) {
	var pil string
	var idxGaji, idxDeadline int

	fmt.Println("a. Deadline terdekat")
	fmt.Println("b. Bayaran tertinggi")
	//fmt.Println("c. Data deadline terdekat dan yang terjauh")
	//fmt.Println("d. Data bayaran dari yang tertinggi ke terendah")
	fmt.Print("Pilihan: ")
	fmt.Scan(&pil)

	if pil == "a" {
		idxDeadline = deadline(data, n)
		if idxDeadline != -1 {
			fmt.Printf("\nDeadline Terdekat:\nClient: %s\nProyek: %s\nDeadline: %s\n",
				data[idxDeadline].client, data[idxDeadline].projek, data[idxDeadline].deadline)
		} else {
			fmt.Println("Data tidak tersedia.")
		}
	} else if pil == "b" {
		idxGaji = bayaranMax(data, n)
		if idxGaji != -1 {
			fmt.Printf("\nBayaran Tertinggi:\nClient: %s\nProyek: %s\nGaji: %d\n",
				data[idxGaji].client, data[idxGaji].projek, data[idxGaji].salary)
		} else {
			fmt.Println("Data tidak tersedia.")
		}
	} else {
		fmt.Println("Pilihan tidak valid.")
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
	var minIdx, i int
	if n == 0 {
		return -1
	}
	minIdx = 0
	for i = 1; i < n; i++ {
		if A[i].deadline < A[minIdx].deadline {
			minIdx = i
		}
	}
	return minIdx
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
			if A[j].deadline < A[minIdx].deadline {
				minIdx = j
			}
		}
		temp = A[i]
		A[i] = A[minIdx]
		A[minIdx] = temp
	}
}
