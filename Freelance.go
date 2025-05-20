package main

import "fmt"

const NMAX int = 1000

type tracking struct {
	client, projek, status string
	salary, idProyek       int
}
type arrTrack [NMAX]tracking

var data arrTrack
var nProyek int

func main() {
	var pilihMenu int

	menuUtama(&pilihMenu)
}

//procedure pilihan menu utama
func menuUtama(pilMenu *int) {
	for *pilMenu != 6 {
		fmt.Println()
		fmt.Println("   ---------- Selamat Datang ----------")
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

func tambahProyek(data *arrTrack, n *int) {
	var jumlah, i int

	fmt.Println()
	fmt.Println("-------------------------------------------------------------------")
	fmt.Println("	Silahkan isi data proyek yang ingin ditambahkan 		")
	fmt.Println("-------------------------------------------------------------------")
	fmt.Println("(Gunakan tanda '_' sebagai spasi jika Nama Client / Nama Proyek lebih dari 1 kata)")
	fmt.Println("Jumlah proyek yang ingin ditambahkan : ")
	fmt.Scan(&jumlah)
	if jumlah+*n > NMAX {
		fmt.Println("Jumlah proyek melebihi batas maksimal")
		return
	}

	for i = 0; i < jumlah && *n < NMAX; i++ {
		fmt.Printf("\n Proyek ke - %d \n", i+1)
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
