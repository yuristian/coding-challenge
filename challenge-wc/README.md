# ccwc — Build Your Own `wc` Tool

Implementasi ulang command line tool Unix `wc` (word, line, character, and byte count) menggunakan Go, sebagai bagian dari [Coding Challenges #1](https://codingchallenges.fyi/challenges/challenge-wc).

## Fitur

- `-c` — menghitung jumlah **byte** dalam file
- `-l` — menghitung jumlah **baris** (line)
- `-w` — menghitung jumlah **kata** (word)
- `-m` — menghitung jumlah **karakter** (character, UTF-8 aware)
- Tanpa flag — default menampilkan `-l -w -c` sekaligus (sesuai perilaku `wc` asli)
- Mendukung input dari **file** maupun **stdin** (pipe)

## Instalasi

Pastikan [Go](https://go.dev/dl/) sudah terinstal (Go 1.20+ disarankan).

```bash
git clone <url-repo-kamu>
cd challenge-wc
go build -o ccwc
```

## Penggunaan

```bash
./ccwc [flag] [nama_file]
```

### Contoh

Menghitung jumlah byte:

```bash
./ccwc -c test.txt
  342190 test.txt
```

Menghitung jumlah baris:

```bash
./ccwc -l test.txt
    7145 test.txt
```

Menghitung jumlah kata:

```bash
./ccwc -w test.txt
   58164 test.txt
```

Menghitung jumlah karakter:

```bash
./ccwc -m test.txt
  339292 test.txt
```

Tanpa flag (default: lines, words, bytes):

```bash
./ccwc test.txt
    7145   58164  342190 test.txt
```

Membaca dari standard input (pipe):

```bash
cat test.txt | ./ccwc -l
    7145
```

Kombinasi beberapa flag:

```bash
./ccwc -l -w test.txt
    7145   58164 test.txt
```

> Catatan: urutan angka pada output selalu mengikuti urutan tetap `lines → chars → bytes → words`, mengikuti perilaku `wc` asli, terlepas dari urutan flag yang diketik di command line.

## Menjalankan tanpa build (development)

```bash
go run main.go -l test.txt
```

## Struktur Kode

```
.
├── main.go     # entry point: parsing flag CLI & pemilihan sumber input (file/stdin)
└── test.txt    # file contoh untuk testing (unduh dari soal challenge)
```

Logika inti penghitungan (`count`) menerima parameter bertipe `io.Reader`, sehingga fungsi yang sama dapat digunakan baik untuk file maupun stdin tanpa duplikasi logika.

```go
type Count struct {
    Lines int
    Words int
    Chars int
    Bytes int
}

func count(r io.Reader) Count
```

## Testing Manual

Unduh file uji `test.txt` dari [tautan soal challenge](https://www.dropbox.com/scl/fi/d4zs6aoq6hr3oew2b6a9v/test.txt?rlkey=20c9d257pxd5emjjzd1gcbn03&dl=0), lalu bandingkan hasil `ccwc` dengan `wc` bawaan sistem (Linux/macOS):

| Flag | Expected |
| ---- | -------- |
| `-c` | 342190   |
| `-l` | 7145     |
| `-w` | 58164    |
| `-m` | 339292   |

## Known Limitations / Improvement Selanjutnya

- Padding angka pada output menggunakan lebar tetap (fixed-width), belum sepenuhnya identik dengan `wc` asli yang menyesuaikan lebar kolom secara dinamis berdasarkan angka terbesar.
- Belum mendukung banyak file sekaligus dalam satu pemanggilan (`wc -l file1.txt file2.txt`).
- Urutan output angka mengikuti urutan tetap (bukan urutan flag di command line), sesuai perilaku `wc` asli.

## Referensi

- [Coding Challenges — Build Your Own wc Tool](https://codingchallenges.fyi/challenges/challenge-wc)
- [Go `flag` package documentation](https://pkg.go.dev/flag)
- [Go `io` package documentation](https://pkg.go.dev/io)
