package main

import (
	// digunakan untuk encoding (marshaling) dan decoding (unmarshaling) data dalam format JSON.
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	//untuk mengubah huruf besar dan kecil.
	"strings"
	//Paket time di Golang menyediakan fungsi-fungsi untuk pengolahan waktu dan penanganan waktu.
	"time"
)

// Struktur data untuk menyimpan informasi tentang event
type Event struct {
	Nama    string `json:"nama"`
	Tanggal string `json:"tanggal"`
	Lokasi  string `json:"lokasi"`
}

// Struktur data untuk menyimpan informasi tentang cosplayer
type Cosplayer struct {
	Nama     string `json:"nama"`
	Karakter string `json:"karakter"`
	Anime    string `json:"anime"`
}

// Struktur data untuk menyimpan informasi tentang partisipasi cosplayer dalam event
type CosplayerEvent struct {
	Cosplayer
	EventNama string `json:"eventNama"`
}

// Struktur data untuk menyimpan keseluruhan database aplikasi ke Json
type Database struct {
	Events     []Event          `json:"events"`
	Cosplayers []Cosplayer      `json:"cosplayers"`
	CosEvents  []CosplayerEvent `json:"cosplayerEvents"`
}

// MENAMPILKAN SEBUAH MENU SETELAH RUNNING FILE main.go
func showMenu() {
	fmt.Println("┌────────────────────────────────────────────────┐")
	fmt.Println("|      Program Registrasi Event Jejepangan       │")
	fmt.Println("├────────────────────────────────────────────────┤")
	fmt.Println("|   PILIHAN MENU                                 |")
	fmt.Println("│   1. TambahEvent                               │")
	fmt.Println("│   2. Cari Event                                │")
	fmt.Println("│   3. Hapus Event                               |")
	fmt.Println("|   4. Tambah Cosplayer                          |")      
	fmt.Println("|   5. Cari Cosplayer                            |")
	fmt.Println("|   6. Hapus Cosplayer                           |")
	fmt.Println("|   7. Tambahkan Cosplayer dalam Event           |")
	fmt.Println("|   8. Cari Cosplayer dalam Event                |")
	fmt.Println("|   9. Hapus Cosplayer dalam Event               |")
	fmt.Println("|  10. Event yang akan diadakan 7 hari mendatang |")
	fmt.Println("|  11. Keluar                                    │")
	fmt.Println("|                                                |")
	fmt.Println("└────────────────────────────────────────────────┘")
}

// KODE UNTUK MEMILIH PILIHAN YANG TERSEDIA SESUAI FUNGSI MULAI DARI 1-11
func main() {
	app := loadAppData()

	for {
		showMenu()
		var choice int
		fmt.Print("Pilih Menu [1-11]: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			tambahEvent(&app)
		case 2:
			cariEvent(&app)
		case 3:
			hapusEvent(&app)
		case 4:
			tambahCosplayer(&app)
		case 5:
			cariCosplayer(&app)
		case 6:
			hapusCosplayer(&app)
		case 7:
			tambahCosplayerKeEvent(&app)
		case 8:
			cariCosplayerDalamEvent(&app)
		case 9:
			hapusCosplayerDalamEvent(&app)
		case 10:
			tampilkanEventMendatang(&app)
		case 11:
			fmt.Println("Program selesai. Sampai jumpa!")
			os.Exit(0)
		default:
			// TERJADI JIKA TIDAK MEMILIH ANGKA YANG DI SEDIAKAN
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}
}

//1) Fungsi untuk menambahkan event baru ke dalam database
func tambahEvent(app *Database) {
	var event Event
	fmt.Print("Nama Event: ")
	fmt.Scanln(&event.Nama)
	fmt.Print("Tanggal Event (YYYY-MM-DD): ")
	fmt.Scanln(&event.Tanggal)
	fmt.Print("Lokasi Event: ")
	fmt.Scanln(&event.Lokasi)
    
	app.Events = append(app.Events, event)
	saveAppData(*app)//FUNGSI UNTUK MENAMBAHKAN KR DATABASE
	fmt.Println("")
	fmt.Println("Event berhasil ditambahkan!")
	fmt.Println("")
    
}

//2) Fungsi untuk mencari event berdasarkan nama
func cariEvent(app *Database) {
	var eventName string
	fmt.Print("Masukkan nama event: ")
	fmt.Println("")
	fmt.Scanln(&eventName)

	for _, event := range app.Events {
		// STRINGS BERGUNA UNTUK MENCARI EVEN DENGAN HURUP KECIL MAUPUN BESAR
		if strings.Contains(strings.ToLower(event.Nama), strings.ToLower(eventName)){
			fmt.Println("Event ditemukan!")
			fmt.Printf("Nama Event: %s\nTanggal Event: %s\nLokasi Event: %s\n", event.Nama, event.Tanggal, event.Lokasi)
			return
		}
	}

	fmt.Println("Event tidak ditemukan.")
}

// 3)Fungsi untuk menghapus event dari database
func hapusEvent(app *Database) {
	fmt.Println("Daftar Event:")
	for i, event := range app.Events {
		fmt.Printf("%d. %s\n", i+1, event.Nama)//%d untuk int ,%s untuk string
	}

	var choice int
	fmt.Print("Pilih Event yang akan dihapus [1-", len(app.Events), "]: ")
	fmt.Scanln(&choice)

	if choice < 1 || choice > len(app.Events) {
		fmt.Println("Pilihan tidak valid.")
		return
	}

	eventName := app.Events[choice-1].Nama
	app.Events = append(app.Events[:choice-1], app.Events[choice:]...)
	saveAppData(*app)
	fmt.Printf("Event %s berhasil dihapus!\n", eventName)
}

// 4)Fungsi untuk menambahkan cosplayer baru ke dalam database
func tambahCosplayer(app *Database) {
	var cosplayer Cosplayer

	fmt.Print("Nama Cosplayer: ")
	fmt.Scanln(&cosplayer.Nama)
	fmt.Print("Nama Karakter: ")
	fmt.Scanln(&cosplayer.Karakter)
	fmt.Print("Nama Anime: ")
	fmt.Scanln(&cosplayer.Anime)

	app.Cosplayers = append(app.Cosplayers, cosplayer)
	saveAppData(*app)
	fmt.Println("")
	fmt.Println("Cosplayer berhasil ditambahkan!")
	fmt.Println("")
}

//5) Fungsi untuk mencari cosplayer berdasarkan nama
func cariCosplayer(app *Database) {
	var cosplayerName string
	fmt.Print("Masukkan nama cosplayer: ")
	fmt.Scanln(&cosplayerName)

	for _, cosplayer := range app.Cosplayers {
			if strings.Contains(strings.ToLower(cosplayer.Nama), strings.ToLower(cosplayerName)){
			fmt.Println("Cosplayer ditemukan!")
			fmt.Printf("Nama Cosplayer: %s\nNama Karakter: %s\nNama Anime: %s\n", cosplayer.Nama, cosplayer.Karakter, cosplayer.Anime)
			return
		}
	}

	fmt.Println("Cosplayer tidak ditemukan.")
}

//6) Fungsi untuk menghapus cosplayer dari database
func hapusCosplayer(app *Database) {
	fmt.Println("Daftar Cosplayer:")
	for i, cosplayer := range app.Cosplayers {
		fmt.Printf("%d. %s\n", i+1, cosplayer.Nama)
	}

	var choice int
	fmt.Print("Pilih Cosplayer yang akan dihapus [1-", len(app.Cosplayers), "]: ")
	fmt.Scanln(&choice)

	if choice < 1 || choice > len(app.Cosplayers) {
		fmt.Println("Pilihan tidak valid.")
		return
	}

	cosplayerName := app.Cosplayers[choice-1].Nama
	app.Cosplayers = append(app.Cosplayers[:choice-1], app.Cosplayers[choice:]...)
	saveAppData(*app)
	fmt.Printf("Cosplayer %s berhasil dihapus!\n", cosplayerName)
}

// 7)Fungsi untuk menambahkan cosplayer ke dalam suatu event
func tambahCosplayerKeEvent(app *Database) {
	fmt.Println("Daftar Cosplayer:")
	for i, cosplayer := range app.Cosplayers {
		fmt.Printf("%d. %s\n", i+1, cosplayer.Nama)
	}

	var cosplayerChoice int
	fmt.Print("Pilih Cosplayer yang akan ditambahkan ke event [1-", len(app.Cosplayers), "]: ")
	fmt.Scanln(&cosplayerChoice)

	if cosplayerChoice < 1 || cosplayerChoice > len(app.Cosplayers) {
		fmt.Println("Pilihan tidak valid.")
		return
	}

	fmt.Println("Daftar Event:")
	for i, event := range app.Events {
		fmt.Printf("%d. %s\n", i+1, event.Nama)
	}

	var eventChoice int
	fmt.Print("Pilih Event yang akan ditambahkan cosplayer [1-", len(app.Events), "]: ")
	fmt.Scanln(&eventChoice)

	if eventChoice < 1 || eventChoice > len(app.Events) {
		fmt.Println("Pilihan tidak valid.")
		return
	}

	cosplayerEvent := CosplayerEvent{
		Cosplayer: app.Cosplayers[cosplayerChoice-1],
		EventNama: app.Events[eventChoice-1].Nama,
	}

	app.CosEvents = append(app.CosEvents, cosplayerEvent)
	saveAppData(*app)
	fmt.Printf("Cosplayer %s berhasil ditambahkan ke event %s!\n", cosplayerEvent.Nama, cosplayerEvent.EventNama)
}

// 8)Fungsi untuk mencari cosplayer dalam suatu event
func cariCosplayerDalamEvent(app *Database) {
	fmt.Println("Daftar Cosplayer:")
	for i, cosplayer := range app.Cosplayers {
		fmt.Printf("%d. %s\n", i+1, cosplayer.Nama)
	}

	var cosplayerChoice int
	fmt.Print("Pilih Cosplayer yang akan dicari dalam event [1-", len(app.Cosplayers), "]: ")
	fmt.Scanln(&cosplayerChoice)

	if cosplayerChoice < 1 || cosplayerChoice > len(app.Cosplayers) {
		fmt.Println("Pilihan tidak valid.")
		return
	}

	cosplayerName := app.Cosplayers[cosplayerChoice-1].Nama

	fmt.Println("Daftar Event:")
	for i, event := range app.CosEvents {
		if event.Cosplayer.Nama == cosplayerName {
			fmt.Printf("%d. %s\n", i+1, event.EventNama)
		}
	}

	fmt.Print("Pilih Event yang ingin dilihat cosplayernya [1-", len(app.CosEvents), "]: ")
	var eventChoice int
	fmt.Scanln(&eventChoice)

	if eventChoice < 1 || eventChoice > len(app.CosEvents) {
		fmt.Println("Pilihan tidak valid.")
		return
	}

	cosplayerEvent := app.CosEvents[eventChoice-1]
	fmt.Printf("Cosplayer %s ditemukan dalam event %s!\n", cosplayerEvent.Cosplayer.Nama, cosplayerEvent.EventNama)
	fmt.Printf("Nama Cosplayer: %s\nNama Karakter: %s\nNama Anime: %s\n", cosplayerEvent.Cosplayer.Nama, cosplayerEvent.Cosplayer.Karakter, cosplayerEvent.Cosplayer.Anime)
}

// 9)Fungsi untuk menghapus cosplayer dari suatu event
func hapusCosplayerDalamEvent(app *Database) {
	fmt.Println("Daftar Cosplayer:")
	for i, cosplayer := range app.Cosplayers {
		fmt.Printf("%d. %s\n", i+1, cosplayer.Nama)
	}

	var cosplayerChoice int
	fmt.Print("Pilih Cosplayer yang akan dihapus dalam event [1-", len(app.Cosplayers), "]: ")
	fmt.Scanln(&cosplayerChoice)

	if cosplayerChoice < 1 || cosplayerChoice > len(app.Cosplayers) {
		fmt.Println("Pilihan tidak valid.")
		return
	}

	cosplayerName := app.Cosplayers[cosplayerChoice-1].Nama

	fmt.Println("Daftar Event:")
	for i, event := range app.CosEvents {
		if event.Cosplayer.Nama == cosplayerName {
			fmt.Printf("%d. %s\n", i+1, event.EventNama)
		}
	}

	fmt.Print("Pilih Event yang ingin dihapus cosplayernya [1-", len(app.CosEvents), "]: ")
	var eventChoice int
	fmt.Scanln(&eventChoice)

	if eventChoice < 1 || eventChoice > len(app.CosEvents) {
		fmt.Println("Pilihan tidak valid.")
		return
	}

	cosplayerEvent := app.CosEvents[eventChoice-1]
	app.CosEvents = append(app.CosEvents[:eventChoice-1], app.CosEvents[eventChoice:]...)
	saveAppData(*app)
	fmt.Printf("Cosplayer %s berhasil dihapus dari event %s!\n", cosplayerEvent.Cosplayer.Nama, cosplayerEvent.EventNama)
}


// 10)Fungsi untuk menampilkan event yang akan diadakan dalam 7 hari mendatang
func tampilkanEventMendatang(app *Database) {
	fmt.Println("Event yang akan diadakan dalam 7 hari mendatang adalah:")

	// Waktu sekarang
	now := time.Now()
    //Untuk Menandai Event yang akan datang
	var adaEvent bool

	for _, event := range app.Events {
		eventDate, err := time.Parse("2004-12-25", event.Tanggal)
		if err == nil {
			// Hitung selisih waktu antara sekarang dan tanggal event
			duration := eventDate.Sub(now)

			// Tampilkan event yang akan diadakan dalam 7 hari mendatang
			if duration > 0 && duration.Hours() <= 7*24 {
				fmt.Printf("%s - %s\n", event.Nama, eventDate.Format("2004-12-25"))
				adaEvent =true
			}
		}
	}
	// Jika Tidak Ada Event Dalam 7 Hari Tampilkan Ini
	if !adaEvent {
		fmt.Println("Tidak Ada Event Dalam 7 Hari Mendatang")
	}
}

// FUNGSI UNTUK MENYIMPAN DATA APLIKASI KE FILE JSON
func saveAppData(app  Database) {
	data, err := json.MarshalIndent(app, "", "  ")
	if err != nil {
		fmt.Println("Gagal menyimpan data aplikasi:", err)
		return
	}

	err = ioutil.WriteFile("JEJEPANGAN.json", data, 0644)
	if err != nil {
		fmt.Println("Gagal menyimpan data aplikasi:", err)
	}
}

// FUNGSI UNTUK MEMUAT DATA APLIKASI DARI FILE JSON
func loadAppData()  Database {
	data, err := ioutil.ReadFile("JEJEPANGAN.json")
	if err != nil {
		return  Database{}
	}

	var app  Database
	err = json.Unmarshal(data, &app)
	if err != nil {
		fmt.Println("Gagal memuat data aplikasi:", err)
		return  Database{}
	}

	return app
}




