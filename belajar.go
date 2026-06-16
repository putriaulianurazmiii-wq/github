package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ============================================================
// TUGAS BESAR - ALGORITMA PEMROGRAMAN 2026
// Aplikasi Pendaftaran Mahasiswa
// Kelompok 5 : Putri & Moreno
// Deskripsi  : Aplikasi untuk mengelola data calon mahasiswa
//              dan data jurusan di suatu perguruan tinggi.
// ============================================================

// ---------- Konstanta ----------
const MAKS_MHS = 100
const MAKS_JUR = 20

// ---------- Tipe Alias ----------
type NIM    = string
type KodeJur = string

// ---------- Tipe Bentukan ----------

// Jurusan menyimpan informasi satu jurusan di perguruan tinggi
type Jurusan struct {
	kode     KodeJur // kode unik jurusan, contoh: "TI", "SI"
	nama     string  // nama lengkap jurusan
	kuota    int     // batas maksimal mahasiswa diterima
	nilaiMin float64 // nilai minimum tes agar diterima
}

// Mahasiswa menyimpan informasi satu calon mahasiswa pendaftar
type Mahasiswa struct {
	nim     NIM     // nomor induk mahasiswa, bersifat unik
	nama    string  // nama lengkap mahasiswa
	asal    string  // asal sekolah mahasiswa
	jurusan KodeJur // kode jurusan yang dipilih
	nilai   float64 // nilai tes seleksi (0-100)
	status  string  // hasil seleksi: "Diterima" / "Ditolak" / ""
}

// ---------- Variabel Global (hanya array utama) ----------
var dataMhs [MAKS_MHS]Mahasiswa
var dataJur [MAKS_JUR]Jurusan
var jmlMhs  int
var jmlJur  int

// ---------- Reader global untuk input dengan spasi ----------
var sc = bufio.NewScanner(os.Stdin)

// ============================================================
// SUBPROGRAM INPUT
// I.S. : scanner siap membaca
// F.S. : mengembalikan nilai yang diinput user
// ============================================================

// Fungsi inputStr
// Spesifikasi : membaca satu baris string dari keyboard,
//               termasuk yang mengandung spasi
// Parameter   : prompt (string) - teks yang ditampilkan ke user
// Return      : string hasil input
func inputStr(prompt string) string {
	fmt.Print(prompt)
	sc.Scan()
	return strings.TrimSpace(sc.Text())
}

// Fungsi inputFloat
// Spesifikasi : membaca bilangan desimal dari keyboard
// Parameter   : prompt (string) - teks yang ditampilkan ke user
// Return      : float64 hasil input; jika tidak valid ulangi
func inputFloat(prompt string) float64 {
	var hasil float64
	valid := false
	for !valid {
		fmt.Print(prompt)
		_, err := fmt.Scan(&hasil)
		sc.Scan() // buang sisa baris
		if err == nil {
			valid = true
		} else {
			fmt.Println("  Input tidak valid, masukkan angka!")
		}
	}
	return hasil
}

// Fungsi inputInt
// Spesifikasi : membaca bilangan bulat dari keyboard
// Parameter   : prompt (string) - teks yang ditampilkan ke user
// Return      : int hasil input; jika tidak valid ulangi
func inputInt(prompt string) int {
	var hasil int
	valid := false
	for !valid {
		fmt.Print(prompt)
		_, err := fmt.Scan(&hasil)
		sc.Scan() // buang sisa baris
		if err == nil {
			valid = true
		} else {
			fmt.Println("  Input tidak valid, masukkan bilangan bulat!")
		}
	}
	return hasil
}

// ============================================================
// SUBPROGRAM PENCARIAN
// ============================================================

// Fungsi seqSearchMhs (Sequential Search)
// Spesifikasi : mencari mahasiswa berdasarkan NIM secara
//               sequential dari index 0 hingga jmlMhs-1
// Parameter   : nim (NIM) - NIM yang dicari
// Return      : index (int) jika ditemukan, -1 jika tidak ada
func seqSearchMhs(nim NIM) int {
	hasil := -1
	i     := 0
	for i < jmlMhs && hasil == -1 {
		if dataMhs[i].nim == nim {
			hasil = i
		}
		i++
	}
	return hasil
}

// Fungsi seqSearchJur (Sequential Search)
// Spesifikasi : mencari jurusan berdasarkan kode secara
//               sequential dari index 0 hingga jmlJur-1
// Parameter   : kode (KodeJur) - kode jurusan yang dicari
// Return      : index (int) jika ditemukan, -1 jika tidak ada
func seqSearchJur(kode KodeJur) int {
	hasil := -1
	i     := 0
	for i < jmlJur && hasil == -1 {
		if dataJur[i].kode == kode {
			hasil = i
		}
		i++
	}
	return hasil
}

// Fungsi binSearchNIM (Binary Search)
// Spesifikasi : mencari mahasiswa berdasarkan NIM pada array
//               yang sudah terurut ascending berdasarkan NIM.
//               Harus dipanggil setelah array diurutkan by NIM.
// Parameter   : arr ([MAKS_MHS]Mahasiswa) - array yang sudah terurut
//               n   (int)                 - jumlah elemen aktif
//               target (NIM)             - NIM yang dicari
// Return      : index (int) jika ditemukan, -1 jika tidak ada
func binSearchNIM(arr [MAKS_MHS]Mahasiswa, n int, target NIM) int {
	kiri   := 0
	kanan  := n - 1
	hasil  := -1
	for kiri <= kanan && hasil == -1 {
		tengah := (kiri + kanan) / 2
		if arr[tengah].nim == target {
			hasil = tengah
		} else if arr[tengah].nim < target {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}
	return hasil
}

// Fungsi binSearchNama (Binary Search)
// Spesifikasi : mencari mahasiswa berdasarkan nama pada array
//               yang sudah terurut ascending berdasarkan nama.
// Parameter   : arr ([MAKS_MHS]Mahasiswa) - array yang sudah terurut
//               n   (int)                 - jumlah elemen aktif
//               target (string)           - nama yang dicari
// Return      : index (int) jika ditemukan, -1 jika tidak ada
func binSearchNama(arr [MAKS_MHS]Mahasiswa, n int, target string) int {
	kiri   := 0
	kanan  := n - 1
	hasil  := -1
	for kiri <= kanan && hasil == -1 {
		tengah := (kiri + kanan) / 2
		if arr[tengah].nama == target {
			hasil = tengah
		} else if arr[tengah].nama < target {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}
	return hasil
}

// ============================================================
// SUBPROGRAM PENGURUTAN
// ============================================================

// Prosedur selSortNilai (Selection Sort)
// Spesifikasi : mengurutkan array mahasiswa berdasarkan nilai tes
//               menggunakan algoritma Selection Sort.
//               Setiap iterasi mencari nilai ekstrem (min/max)
//               dari sisa array lalu ditukar ke posisi i.
// Parameter   : arr (*[MAKS_MHS]Mahasiswa) - array yang diurutkan (pass by pointer)
//               n   (int)                  - jumlah elemen aktif
//               asc (bool)                 - true=ascending, false=descending
func selSortNilai(arr *[MAKS_MHS]Mahasiswa, n int, asc bool) {
	for i := 0; i < n-1; i++ {
		posEkstrem := i
		for j := i + 1; j < n; j++ {
			if asc && arr[j].nilai < arr[posEkstrem].nilai {
				posEkstrem = j
			} else if !asc && arr[j].nilai > arr[posEkstrem].nilai {
				posEkstrem = j
			}
		}
		arr[i], arr[posEkstrem] = arr[posEkstrem], arr[i]
	}
}

// Prosedur insSortNama (Insertion Sort)
// Spesifikasi : mengurutkan array mahasiswa berdasarkan nama
//               menggunakan algoritma Insertion Sort.
//               Setiap elemen disisipkan ke posisi yang tepat
//               di bagian array yang sudah terurut.
// Parameter   : arr (*[MAKS_MHS]Mahasiswa) - array yang diurutkan (pass by pointer)
//               n   (int)                  - jumlah elemen aktif
//               asc (bool)                 - true=ascending, false=descending
func insSortNama(arr *[MAKS_MHS]Mahasiswa, n int, asc bool) {
	for i := 1; i < n; i++ {
		temp   := arr[i]
		j      := i - 1
		geser  := true
		for j >= 0 && geser {
			if asc && arr[j].nama > temp.nama {
				arr[j+1] = arr[j]
				j--
			} else if !asc && arr[j].nama < temp.nama {
				arr[j+1] = arr[j]
				j--
			} else {
				geser = false
			}
		}
		arr[j+1] = temp
	}
}

// Prosedur insSortJurusan (Insertion Sort)
// Spesifikasi : mengurutkan array mahasiswa berdasarkan kode jurusan
//               menggunakan algoritma Insertion Sort.
// Parameter   : arr (*[MAKS_MHS]Mahasiswa) - array yang diurutkan (pass by pointer)
//               n   (int)                  - jumlah elemen aktif
//               asc (bool)                 - true=ascending, false=descending
func insSortJurusan(arr *[MAKS_MHS]Mahasiswa, n int, asc bool) {
	for i := 1; i < n; i++ {
		temp  := arr[i]
		j     := i - 1
		geser := true
		for j >= 0 && geser {
			if asc && arr[j].jurusan > temp.jurusan {
				arr[j+1] = arr[j]
				j--
			} else if !asc && arr[j].jurusan < temp.jurusan {
				arr[j+1] = arr[j]
				j--
			} else {
				geser = false
			}
		}
		arr[j+1] = temp
	}
}

// Prosedur insSortNIM (Insertion Sort)
// Spesifikasi : mengurutkan array mahasiswa berdasarkan NIM
//               menggunakan algoritma Insertion Sort.
//               Digunakan sebagai persiapan sebelum Binary Search by NIM.
// Parameter   : arr (*[MAKS_MHS]Mahasiswa) - array yang diurutkan (pass by pointer)
//               n   (int)                  - jumlah elemen aktif
//               asc (bool)                 - true=ascending, false=descending
func insSortNIM(arr *[MAKS_MHS]Mahasiswa, n int, asc bool) {
	for i := 1; i < n; i++ {
		temp  := arr[i]
		j     := i - 1
		geser := true
		for j >= 0 && geser {
			if asc && arr[j].nim > temp.nim {
				arr[j+1] = arr[j]
				j--
			} else if !asc && arr[j].nim < temp.nim {
				arr[j+1] = arr[j]
				j--
			} else {
				geser = false
			}
		}
		arr[j+1] = temp
	}
}

// ============================================================
// SUBPROGRAM PEMBANTU
// ============================================================

// Prosedur garis
// Spesifikasi : mencetak garis pemisah horizontal
func garis() {
	fmt.Println("------------------------------------------------------------")
}

// Fungsi getNamaJur
// Spesifikasi : mengambil nama jurusan berdasarkan kode
// Parameter   : kode (KodeJur) - kode jurusan yang dicari
// Return      : nama jurusan (string), "?" jika tidak ditemukan
func getNamaJur(kode KodeJur) string {
	i := seqSearchJur(kode)
	if i != -1 {
		return dataJur[i].nama
	}
	return "?"
}

// Fungsi hitungDiterima
// Spesifikasi : menghitung jumlah mahasiswa yang sudah diterima
//               pada jurusan tertentu (untuk cek kuota)
// Parameter   : kode (KodeJur) - kode jurusan yang dicek
// Return      : jumlah mahasiswa diterima (int)
func hitungDiterima(kode KodeJur) int {
	total := 0
	for i := 0; i < jmlMhs; i++ {
		if dataMhs[i].jurusan == kode && dataMhs[i].status == "Diterima" {
			total++
		}
	}
	return total
}

// Prosedur updateStatus
// Spesifikasi : memperbarui status diterima/ditolak mahasiswa
//               berdasarkan nilai tes dan nilai minimum jurusan.
//               Juga mempertimbangkan kuota jurusan.
// Parameter   : i (int) - index mahasiswa di dataMhs
func updateStatus(i int) {
	j := seqSearchJur(dataMhs[i].jurusan)
	if j != -1 && dataMhs[i].nilai > -1 {
		if dataMhs[i].nilai >= dataJur[j].nilaiMin {
			// cek kuota: jika sudah diterima sebelumnya tidak dihitung ulang
			sudahDiterima := dataMhs[i].status == "Diterima"
			terisi := hitungDiterima(dataMhs[i].jurusan)
			if sudahDiterima || terisi < dataJur[j].kuota {
				dataMhs[i].status = "Diterima"
			} else {
				dataMhs[i].status = "Ditolak" // lulus nilai tapi kuota penuh
			}
		} else {
			dataMhs[i].status = "Ditolak"
		}
	}
}

// Prosedur headerMhs
// Spesifikasi : mencetak baris header tabel data mahasiswa
func headerMhs() {
	garis()
	fmt.Printf("  %-5s %-12s %-22s %-8s %-8s %-10s\n",
		"No", "NIM", "NAMA", "JURUSAN", "NILAI", "STATUS")
	garis()
}

// Prosedur barisMhs
// Spesifikasi : mencetak satu baris data mahasiswa dalam format tabel
// Parameter   : no (int)        - nomor urut tampil
//               m  (Mahasiswa)  - data mahasiswa yang dicetak
func barisMhs(no int, m Mahasiswa) {
	st := m.status
	if st == "" {
		st = "Belum ada"
	}
	fmt.Printf("  %-5d %-12s %-22s %-8s %-8.2f %-10s\n",
		no, m.nim, m.nama, m.jurusan, m.nilai, st)
}

// Prosedur salinArray
// Spesifikasi : menyalin isi dataMhs ke array baru agar data
//               asli tidak berubah saat proses pengurutan
// Return      : salinan array ([MAKS_MHS]Mahasiswa) dan jumlahnya (int)
func salinArray() ([MAKS_MHS]Mahasiswa, int) {
	var temp [MAKS_MHS]Mahasiswa
	for i := 0; i < jmlMhs; i++ {
		temp[i] = dataMhs[i]
	}
	return temp, jmlMhs
}

// Fungsi pilihanAsc
// Spesifikasi : menampilkan menu pilihan urutan dan membaca pilihan user
// Return      : true jika ascending, false jika descending
func pilihanAsc() bool {
	fmt.Println("  1. Ascending  (A-Z / Kecil ke Besar)")
	fmt.Println("  2. Descending (Z-A / Besar ke Kecil)")
	p := inputInt("  Pilih urutan: ")
	return p == 1
}

// ============================================================
// MODUL JURUSAN
// ============================================================

// Prosedur tambahJurusan
// Spesifikasi : menambahkan satu data jurusan baru ke dataJur.
//               Kode jurusan harus unik.
func tambahJurusan() {
	fmt.Println("\n=== TAMBAH JURUSAN ===")
	if jmlJur >= MAKS_JUR {
		fmt.Println("  Data jurusan sudah penuh!")
		return
	}
	kode := inputStr("  Kode Jurusan  : ")
	if seqSearchJur(kode) != -1 {
		fmt.Println("  Kode jurusan sudah ada!")
		return
	}
	nama  := inputStr("  Nama Jurusan  : ")
	kuota := inputInt("  Kuota         : ")
	min   := inputFloat("  Nilai Minimum : ")

	dataJur[jmlJur] = Jurusan{kode, nama, kuota, min}
	jmlJur++
	fmt.Println("  Jurusan berhasil ditambahkan!")
}

// Prosedur editJurusan
// Spesifikasi : mengubah data jurusan yang sudah ada.
//               Input kosong (Enter) berarti field tidak diubah.
//               Jika nilai minimum diubah, status semua mahasiswa
//               di jurusan tersebut akan diperbarui otomatis.
func editJurusan() {
	fmt.Println("\n=== EDIT JURUSAN ===")
	kode := inputStr("  Kode jurusan yang diedit: ")
	i    := seqSearchJur(kode)
	if i == -1 {
		fmt.Println("  Jurusan tidak ditemukan!")
		return
	}
	fmt.Printf("  Data saat ini: [%s] %s | Kuota: %d | Nilai Min: %.2f\n",
		dataJur[i].kode, dataJur[i].nama, dataJur[i].kuota, dataJur[i].nilaiMin)
	fmt.Println("  (Tekan Enter untuk melewati field)")

	nama := inputStr("  Nama baru     : ")
	if nama != "" {
		dataJur[i].nama = nama
	}
	kuotaStr := inputStr("  Kuota baru    : ")
	if kuotaStr != "" {
		kuota := 0
		fmt.Sscan(kuotaStr, &kuota)
		if kuota > 0 {
			dataJur[i].kuota = kuota
		}
	}
	minStr := inputStr("  Nilai min baru: ")
	if minStr != "" {
		var min float64
		fmt.Sscan(minStr, &min)
		if min > 0 {
			dataJur[i].nilaiMin = min
			for k := 0; k < jmlMhs; k++ {
				if dataMhs[k].jurusan == kode {
					updateStatus(k)
				}
			}
			fmt.Println("  Status mahasiswa di jurusan ini diperbarui otomatis.")
		}
	}
	fmt.Println("  Jurusan berhasil diperbarui!")
}

// Prosedur hapusJurusan
// Spesifikasi : menghapus data jurusan dari dataJur.
//               Tidak bisa dihapus jika masih ada mahasiswa
//               yang mendaftar ke jurusan tersebut.
func hapusJurusan() {
	fmt.Println("\n=== HAPUS JURUSAN ===")
	kode := inputStr("  Kode jurusan yang dihapus: ")
	i    := seqSearchJur(kode)
	if i == -1 {
		fmt.Println("  Jurusan tidak ditemukan!")
		return
	}
	ada := false
	for k := 0; k < jmlMhs; k++ {
		if dataMhs[k].jurusan == kode {
			ada = true
		}
	}
	if ada {
		fmt.Println("  Tidak bisa dihapus, masih ada mahasiswa yang mendaftar di jurusan ini!")
		return
	}
	fmt.Printf("  Hapus jurusan [%s] %s? (y/n): ", dataJur[i].kode, dataJur[i].nama)
	konfirm := inputStr("")
	if konfirm == "y" {
		for k := i; k < jmlJur-1; k++ {
			dataJur[k] = dataJur[k+1]
		}
		jmlJur--
		fmt.Println("  Jurusan berhasil dihapus!")
	} else {
		fmt.Println("  Penghapusan dibatalkan.")
	}
}

// Prosedur tampilSemuaJurusan
// Spesifikasi : menampilkan seluruh data jurusan dalam format tabel,
//               dilengkapi info jumlah pendaftar dan jumlah diterima.
func tampilSemuaJurusan() {
	fmt.Println("\n=== DAFTAR JURUSAN ===")
	if jmlJur == 0 {
		fmt.Println("  Belum ada data jurusan.")
		return
	}
	garis()
	fmt.Printf("  %-8s %-25s %-8s %-10s %-10s %-9s\n",
		"KODE", "NAMA", "KUOTA", "NILAI MIN", "PENDAFTAR", "DITERIMA")
	garis()
	for i := 0; i < jmlJur; i++ {
		j := dataJur[i]
		pendaftar := 0
		for k := 0; k < jmlMhs; k++ {
			if dataMhs[k].jurusan == j.kode {
				pendaftar++
			}
		}
		diterima := hitungDiterima(j.kode)
		fmt.Printf("  %-8s %-25s %-8d %-10.2f %-10d %-9d\n",
			j.kode, j.nama, j.kuota, j.nilaiMin, pendaftar, diterima)
	}
	garis()
}

// ============================================================
// MODUL MAHASISWA
// ============================================================

// Prosedur tambahMhs
// Spesifikasi : mendaftarkan satu calon mahasiswa baru.
//               NIM harus unik. Jurusan harus sudah ada.
//               Nilai tes diisi terpisah melalui menu Nilai Tes.
func tambahMhs() {
	fmt.Println("\n=== TAMBAH MAHASISWA ===")
	if jmlMhs >= MAKS_MHS {
		fmt.Println("  Data mahasiswa sudah penuh!")
		return
	}
	if jmlJur == 0 {
		fmt.Println("  Tambah jurusan terlebih dahulu!")
		return
	}
	nim := inputStr("  NIM           : ")
	if seqSearchMhs(nim) != -1 {
		fmt.Println("  NIM sudah terdaftar!")
		return
	}
	nama  := inputStr("  Nama Lengkap  : ")
	asal  := inputStr("  Asal Sekolah  : ")
	tampilSemuaJurusan()
	kodeJ := inputStr("  Kode Jurusan  : ")
	if seqSearchJur(kodeJ) == -1 {
		fmt.Println("  Kode jurusan tidak ditemukan!")
		return
	}
	dataMhs[jmlMhs] = Mahasiswa{nim, nama, asal, kodeJ, 0, ""}
	jmlMhs++
	fmt.Println("  Mahasiswa berhasil didaftarkan!")
}

// Prosedur editMhs
// Spesifikasi : mengubah data mahasiswa yang sudah terdaftar.
//               Input kosong (Enter) berarti field tidak diubah.
//               Jika jurusan diubah, status akan diperbarui otomatis.
func editMhs() {
	fmt.Println("\n=== EDIT MAHASISWA ===")
	nim := inputStr("  NIM mahasiswa yang diedit: ")
	i   := seqSearchMhs(nim)
	if i == -1 {
		fmt.Println("  Mahasiswa tidak ditemukan!")
		return
	}
	fmt.Printf("  Data: [%s] %s | Asal: %s | Jurusan: %s\n",
		dataMhs[i].nim, dataMhs[i].nama, dataMhs[i].asal, dataMhs[i].jurusan)
	fmt.Println("  (Tekan Enter untuk melewati field)")

	nama := inputStr("  Nama baru       : ")
	if nama != "" {
		dataMhs[i].nama = nama
	}
	asal := inputStr("  Asal sekolah baru: ")
	if asal != "" {
		dataMhs[i].asal = asal
	}
	tampilSemuaJurusan()
	kodeJ := inputStr("  Kode jurusan baru: ")
	if kodeJ != "" {
		if seqSearchJur(kodeJ) == -1 {
			fmt.Println("  Kode jurusan tidak ada, jurusan tidak diubah.")
		} else {
			dataMhs[i].jurusan = kodeJ
			dataMhs[i].status  = ""
			updateStatus(i)
		}
	}
	fmt.Println("  Data mahasiswa berhasil diperbarui!")
}

// Prosedur hapusMhs
// Spesifikasi : menghapus data mahasiswa dari dataMhs.
//               Data digeser ke kiri untuk menutup celah.
func hapusMhs() {
	fmt.Println("\n=== HAPUS MAHASISWA ===")
	nim := inputStr("  NIM mahasiswa yang dihapus: ")
	i   := seqSearchMhs(nim)
	if i == -1 {
		fmt.Println("  Mahasiswa tidak ditemukan!")
		return
	}
	fmt.Printf("  Hapus [%s] %s - Jurusan %s? (y/n): ",
		dataMhs[i].nim, dataMhs[i].nama, dataMhs[i].jurusan)
	konfirm := inputStr("")
	if konfirm == "y" {
		for k := i; k < jmlMhs-1; k++ {
			dataMhs[k] = dataMhs[k+1]
		}
		jmlMhs--
		fmt.Println("  Mahasiswa berhasil dihapus!")
	} else {
		fmt.Println("  Penghapusan dibatalkan.")
	}
}

// Prosedur tampilSemuaMhs
// Spesifikasi : menampilkan seluruh data mahasiswa dalam format tabel
func tampilSemuaMhs() {
	fmt.Println("\n=== DAFTAR SEMUA MAHASISWA ===")
	if jmlMhs == 0 {
		fmt.Println("  Belum ada data mahasiswa.")
		return
	}
	headerMhs()
	for i := 0; i < jmlMhs; i++ {
		barisMhs(i+1, dataMhs[i])
	}
	garis()
	fmt.Printf("  Total: %d mahasiswa\n", jmlMhs)
}

// Prosedur tampilMhsPerJurusan
// Spesifikasi : menampilkan semua mahasiswa yang mendaftar
//               ke jurusan tertentu (berdasarkan kode jurusan)
func tampilMhsPerJurusan() {
	fmt.Println("\n=== MAHASISWA PER JURUSAN ===")
	tampilSemuaJurusan()
	kode := inputStr("  Masukkan kode jurusan: ")
	if seqSearchJur(kode) == -1 {
		fmt.Println("  Jurusan tidak ditemukan!")
		return
	}
	fmt.Printf("\n  Jurusan: %s (%s)\n", getNamaJur(kode), kode)
	headerMhs()
	no    := 1
	for i := 0; i < jmlMhs; i++ {
		if dataMhs[i].jurusan == kode {
			barisMhs(no, dataMhs[i])
			no++
		}
	}
	garis()
	fmt.Printf("  Total pendaftar: %d | Diterima: %d\n", no-1, hitungDiterima(kode))
}

// Prosedur tampilDiterima
// Spesifikasi : menampilkan semua mahasiswa dengan status "Diterima"
func tampilDiterima() {
	fmt.Println("\n=== MAHASISWA DITERIMA ===")
	headerMhs()
	no    := 1
	for i := 0; i < jmlMhs; i++ {
		if dataMhs[i].status == "Diterima" {
			barisMhs(no, dataMhs[i])
			no++
		}
	}
	garis()
	fmt.Printf("  Total diterima: %d mahasiswa\n", no-1)
}

// Prosedur tampilDitolak
// Spesifikasi : menampilkan semua mahasiswa dengan status "Ditolak"
func tampilDitolak() {
	fmt.Println("\n=== MAHASISWA DITOLAK ===")
	headerMhs()
	no    := 1
	for i := 0; i < jmlMhs; i++ {
		if dataMhs[i].status == "Ditolak" {
			barisMhs(no, dataMhs[i])
			no++
		}
	}
	garis()
	fmt.Printf("  Total ditolak: %d mahasiswa\n", no-1)
}

// ============================================================
// MODUL NILAI TES
// ============================================================

// Prosedur tambahNilai
// Spesifikasi : menambahkan atau mengubah nilai tes mahasiswa.
//               Setelah nilai diisi, status diterima/ditolak
//               akan diperbarui otomatis berdasarkan nilai minimum
//               jurusan dan kuota yang tersedia.
func tambahNilai() {
	fmt.Println("\n=== TAMBAH / UBAH NILAI TES ===")
	nim := inputStr("  NIM mahasiswa: ")
	i   := seqSearchMhs(nim)
	if i == -1 {
		fmt.Println("  Mahasiswa tidak ditemukan!")
		return
	}
	fmt.Printf("  Nama     : %s\n", dataMhs[i].nama)
	fmt.Printf("  Jurusan  : %s (%s)\n", getNamaJur(dataMhs[i].jurusan), dataMhs[i].jurusan)
	fmt.Printf("  Nilai saat ini: %.2f | Status: %s\n", dataMhs[i].nilai, dataMhs[i].status)

	nilai := inputFloat("  Nilai tes baru (0-100): ")
	if nilai < 0 || nilai > 100 {
		fmt.Println("  Nilai harus antara 0 sampai 100!")
		return
	}
	dataMhs[i].nilai = nilai
	updateStatus(i)
	fmt.Printf("  Nilai berhasil disimpan! Status: %s\n", dataMhs[i].status)
}

// Prosedur hapusNilai
// Spesifikasi : menghapus nilai tes mahasiswa dan mereset statusnya.
//               Status akan dikembalikan menjadi kosong ("").
func hapusNilai() {
	fmt.Println("\n=== HAPUS NILAI TES ===")
	nim := inputStr("  NIM mahasiswa: ")
	i   := seqSearchMhs(nim)
	if i == -1 {
		fmt.Println("  Mahasiswa tidak ditemukan!")
		return
	}
	fmt.Printf("  Nama: %s | Nilai: %.2f | Status: %s\n",
		dataMhs[i].nama, dataMhs[i].nilai, dataMhs[i].status)
	fmt.Print("  Yakin hapus nilai? (y/n): ")
	konfirm := inputStr("")
	if konfirm == "y" {
		dataMhs[i].nilai  = 0
		dataMhs[i].status = ""
		fmt.Println("  Nilai berhasil dihapus, status direset.")
	} else {
		fmt.Println("  Penghapusan dibatalkan.")
	}
}

// ============================================================
// MODUL TAMPIL TERURUT & PENCARIAN
// ============================================================

// Prosedur tampilUrutNilai
// Spesifikasi : menampilkan data mahasiswa terurut berdasarkan
//               nilai tes menggunakan Selection Sort.
//               User memilih urutan ascending atau descending.
func tampilUrutNilai() {
	fmt.Println("\n=== URUT BERDASARKAN NILAI TES (Selection Sort) ===")
	if jmlMhs == 0 {
		fmt.Println("  Belum ada data mahasiswa.")
		return
	}
	asc       := pilihanAsc()
	temp, n   := salinArray()
	selSortNilai(&temp, n, asc)
	arah := "Ascending"
	if !asc {
		arah = "Descending"
	}
	fmt.Printf("  Urutan: %s\n", arah)
	headerMhs()
	for i := 0; i < n; i++ {
		barisMhs(i+1, temp[i])
	}
	garis()
}

// Prosedur tampilUrutNama
// Spesifikasi : menampilkan data mahasiswa terurut berdasarkan
//               nama menggunakan Insertion Sort.
//               Jika ascending, tersedia fitur Binary Search by nama.
func tampilUrutNama() {
	fmt.Println("\n=== URUT BERDASARKAN NAMA (Insertion Sort) ===")
	if jmlMhs == 0 {
		fmt.Println("  Belum ada data mahasiswa.")
		return
	}
	asc       := pilihanAsc()
	temp, n   := salinArray()
	insSortNama(&temp, n, asc)
	arah := "Ascending"
	if !asc {
		arah = "Descending"
	}
	fmt.Printf("  Urutan: %s\n", arah)

	if asc {
		fmt.Println("\n  [Fitur Binary Search tersedia karena data sudah terurut ascending]")
		cari := inputStr("  Cari nama (Enter untuk skip): ")
		if cari != "" {
			idx := binSearchNama(temp, n, cari)
			if idx != -1 {
				fmt.Printf("  Ditemukan! Berikut datanya:\n")
				headerMhs()
				barisMhs(idx+1, temp[idx])
				garis()
			} else {
				fmt.Println("  Nama tidak ditemukan.")
			}
		}
	}

	headerMhs()
	for i := 0; i < n; i++ {
		barisMhs(i+1, temp[i])
	}
	garis()
}

// Prosedur tampilUrutJurusan
// Spesifikasi : menampilkan data mahasiswa terurut berdasarkan
//               kode jurusan menggunakan Insertion Sort.
func tampilUrutJurusan() {
	fmt.Println("\n=== URUT BERDASARKAN JURUSAN (Insertion Sort) ===")
	if jmlMhs == 0 {
		fmt.Println("  Belum ada data mahasiswa.")
		return
	}
	asc       := pilihanAsc()
	temp, n   := salinArray()
	insSortJurusan(&temp, n, asc)
	arah := "Ascending"
	if !asc {
		arah = "Descending"
	}
	fmt.Printf("  Urutan: %s\n", arah)
	headerMhs()
	for i := 0; i < n; i++ {
		barisMhs(i+1, temp[i])
	}
	garis()
}

// Prosedur tampilUrutNIM
// Spesifikasi : menampilkan data mahasiswa terurut berdasarkan NIM
//               menggunakan Insertion Sort.
//               Tersedia fitur Binary Search by NIM setelah terurut.
func tampilUrutNIM() {
	fmt.Println("\n=== URUT BERDASARKAN NIM (Insertion Sort) ===")
	if jmlMhs == 0 {
		fmt.Println("  Belum ada data mahasiswa.")
		return
	}
	asc       := pilihanAsc()
	temp, n   := salinArray()
	insSortNIM(&temp, n, asc)
	arah := "Ascending"
	if !asc {
		arah = "Descending"
	}
	fmt.Printf("  Urutan: %s\n", arah)

	if asc {
		fmt.Println("\n  [Fitur Binary Search tersedia karena data sudah terurut ascending]")
		cari := inputStr("  Cari NIM (Enter untuk skip): ")
		if cari != "" {
			idx := binSearchNIM(temp, n, cari)
			if idx != -1 {
				fmt.Println("  Ditemukan!")
				headerMhs()
				barisMhs(idx+1, temp[idx])
				garis()
			} else {
				fmt.Println("  NIM tidak ditemukan.")
			}
		}
	}

	headerMhs()
	for i := 0; i < n; i++ {
		barisMhs(i+1, temp[i])
	}
	garis()
}

// ============================================================
// MENU
// ============================================================

func menuJurusan() {
	selesai := false
	for !selesai {
		fmt.Println("\n--- MENU JURUSAN ---")
		fmt.Println("  1. Tambah Jurusan")
		fmt.Println("  2. Edit Jurusan")
		fmt.Println("  3. Hapus Jurusan")
		fmt.Println("  4. Lihat Semua Jurusan")
		fmt.Println("  0. Kembali")
		p := inputInt("  Pilih: ")
		if p == 1 {
			tambahJurusan()
		} else if p == 2 {
			editJurusan()
		} else if p == 3 {
			hapusJurusan()
		} else if p == 4 {
			tampilSemuaJurusan()
		} else if p == 0 {
			selesai = true
		} else {
			fmt.Println("  Pilihan tidak valid!")
		}
	}
}

func menuMahasiswa() {
	selesai := false
	for !selesai {
		fmt.Println("\n--- MENU MAHASISWA ---")
		fmt.Println("  1. Tambah Mahasiswa")
		fmt.Println("  2. Edit Mahasiswa")
		fmt.Println("  3. Hapus Mahasiswa")
		fmt.Println("  4. Lihat Semua Mahasiswa")
		fmt.Println("  5. Lihat Mahasiswa per Jurusan")
		fmt.Println("  6. Lihat Mahasiswa Diterima")
		fmt.Println("  7. Lihat Mahasiswa Ditolak")
		fmt.Println("  0. Kembali")
		p := inputInt("  Pilih: ")
		if p == 1 {
			tambahMhs()
		} else if p == 2 {
			editMhs()
		} else if p == 3 {
			hapusMhs()
		} else if p == 4 {
			tampilSemuaMhs()
		} else if p == 5 {
			tampilMhsPerJurusan()
		} else if p == 6 {
			tampilDiterima()
		} else if p == 7 {
			tampilDitolak()
		} else if p == 0 {
			selesai = true
		} else {
			fmt.Println("  Pilihan tidak valid!")
		}
	}
}

func menuNilai() {
	selesai := false
	for !selesai {
		fmt.Println("\n--- MENU NILAI TES ---")
		fmt.Println("  1. Tambah / Ubah Nilai Tes")
		fmt.Println("  2. Hapus Nilai Tes")
		fmt.Println("  0. Kembali")
		p := inputInt("  Pilih: ")
		if p == 1 {
			tambahNilai()
		} else if p == 2 {
			hapusNilai()
		} else if p == 0 {
			selesai = true
		} else {
			fmt.Println("  Pilihan tidak valid!")
		}
	}
}

func menuTampilUrut() {
	selesai := false
	for !selesai {
		fmt.Println("\n--- MENU TAMPIL TERURUT ---")
		fmt.Println("  1. Urut Nilai Tes   (Selection Sort)")
		fmt.Println("  2. Urut Nama        (Insertion Sort + Binary Search)")
		fmt.Println("  3. Urut Jurusan     (Insertion Sort)")
		fmt.Println("  4. Urut NIM         (Insertion Sort + Binary Search)")
		fmt.Println("  0. Kembali")
		p := inputInt("  Pilih: ")
		if p == 1 {
			tampilUrutNilai()
		} else if p == 2 {
			tampilUrutNama()
		} else if p == 3 {
			tampilUrutJurusan()
		} else if p == 4 {
			tampilUrutNIM()
		} else if p == 0 {
			selesai = true
		} else {
			fmt.Println("  Pilihan tidak valid!")
		}
	}
}

// ============================================================
// MAIN
// ============================================================

func main() {
	fmt.Println("============================================================")
	fmt.Println("        APLIKASI PENDAFTARAN MAHASISWA")
	fmt.Println("        Tugas Besar Algoritma Pemrograman 2026")
	fmt.Println("        Kelompok 5 : Putri & Moreno")
	fmt.Println("============================================================")

	selesai := false
	for !selesai {
		fmt.Println("\n========== MENU UTAMA ==========")
		fmt.Println("  1. Manajemen Jurusan")
		fmt.Println("  2. Manajemen Mahasiswa")
		fmt.Println("  3. Manajemen Nilai Tes")
		fmt.Println("  4. Tampil Data Terurut & Cari")
		fmt.Println("  0. Keluar")
		p := inputInt("  Pilih: ")
		if p == 1 {
			menuJurusan()
		} else if p == 2 {
			menuMahasiswa()
		} else if p == 3 {
			menuNilai()
		} else if p == 4 {
			menuTampilUrut()
		} else if p == 0 {
			fmt.Println("\n  Terima kasih, sampai jumpa!")
			selesai = true
		} else {
			fmt.Println("  Pilihan tidak valid!")
		}
	}
}