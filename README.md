## Reverse Slice Element Order

Buatlah fungsi di Golang yang menerima sebuah slice of integer dan membalik urutan elemennya secara in-place (tanpa membuat slice tambahan).

### Contoh:

Input: []int{1, 2, 3, 4, 5}
Output: []int{5, 4, 3, 2, 1}

### Solusi:
1. Definisi Fungsi:
Fungsi reverseSlice menerima parameter berupa slice of integer ([]int).

2. Mekanisme Pembalikan:
Menggunakan loop dengan dua pointer (i dan j):
- i dimulai dari indeks 0 (awal slice).
- j dimulai dari indeks terakhir (len(nums)-1).
- Pada tiap iterasi, elemen di indeks i dan j ditukar, lalu i dinaikkan dan j diturunkan hingga i tidak lagi lebih kecil dari j.

3. In-place Operation:
Karena proses penukaran dilakukan langsung pada slice yang diterima, tidak ada slice baru yang dibuat—ini adalah operasi in-place.

4. Fungsi Main:
Fungsi main() mendemonstrasikan cara kerja fungsi dengan:
- Mencetak slice awal.
- Memanggil reverseSlice untuk membalik urutan elemen.
- Mencetak slice setelah dibalik.
