1️⃣ fasthttp (⚡ Paling Cepat)
https://github.com/valyala/fasthttp
Tidak kompatibel langsung dengan net/http
Lebih sulit digunakan untuk proyek besar
Sangat cepat untuk kebutuhan high-performance API
2️⃣ chi (🚀 Ringan & Fleksibel)
https://github.com/go-chi/chi

Mirip gorilla/mux tetapi lebih cepat
Bisa digunakan sebagai drop-in replacement untuk http.NewServeMux()
3️⃣ gorilla/mux (🔧 Powerful, tapi Lebih Lambat dari chi)
https://github.com/gorilla/mux
Middleware support yang bagus
Banyak fitur seperti regex routing
Sedikit lebih lambat dari chi
💡 Kesimpulan:

Performansi terbaik? → Gunakan fasthttp
Butuh fleksibilitas dengan performa bagus? → Gunakan chi
Butuh fitur lanjutan dan lebih familiar? → Gunakan gorilla/mux

🔥 1. Performa
Framework           | Benchmark     (req/sec, lebih tinggi lebih baik)
net/http (ServeMux)	| ~75k req/sec ✅
chi	                | ~90k req/sec 🚀

👉 chi lebih cepat karena optimasi internal dan penggunaan radix tree untuk routing.

🏗️ 2. Routing & Middleware
Aspek	                           | net/http(ServeMux) |	chi
Static Routes	                   |✅ Ya	           |✅ Ya    
Dynamic Routes (/users/{id})       |	❌ Tidak        |✅ Ya
Middleware Support	               |❌ Terbatas (manual)|✅ Ya (built-in)
Grouping Routes	                   |❌ Tidak	✅ Ya      |
Kesimpulan: Pilih yang Mana?
Pakai net/http jika:
✅ API sederhana tanpa banyak route
✅ Tidak butuh middleware atau fitur lanjutan
✅ Mau dependensi minimal (hanya bawaan Go)

Pakai chi jika:
✅ Perlu dynamic routing (/users/{id})
✅ Ingin middleware support bawaan
✅ API skala menengah/besar dengan modularitas tinggi

🔥 Kesimpulan: Kenapa chi Lebih Cepat?
Faktor	                    | http.NewServeMux()|	chi
Routing Algorithm	        | Linear Search (lebih lambat)	|Radix Tree (lebih cepat)
Middleware Handling	Manual  | (lebih lambat)	Built-in    |chaining (lebih cepat)
Memory Efficiency	        | Lebih boros dalam high-load	|Zero-allocation routing

🔥 Perbandingan chi vs gorilla/mux
Kedua router ini populer untuk Go, tapi mana yang lebih baik?

Aspek	                |chi 🚀	            | gorilla/mux 🦍
Performance	Lebih cepat | (Radix Tree) ✅   | Lebih lambat (Regex-based) ❌
Routing Method	        |Radix Tree ✅	   |Regex Parsing ❌
Dynamic Routes	        |✅ Ya	           |✅ Ya
Middleware Support	    |✅ Built-in	       |❌ Manual atau third-party
Grouping Routes     	|✅ Ya	           |❌ Tidak
Community & Maintenance |✅ Aktif & modern  |❌ Tidak lagi aktif (archived)

1️⃣ Performa: chi Lebih Cepat dari gorilla/mux
chi menggunakan Radix Tree → lebih cepat dalam pencocokan routes.
gorilla/mux menggunakan Regex-based routing → lebih fleksibel, tapi lebih lambat.
🔹 Benchmark (req/sec)

Framework   |	Performance
chi	        | ~90k req/sec 🚀
gorilla/mux	| ~40k req/sec ❌ (lebih lambat)
✅ Jika API butuh performa tinggi → chi lebih baik`.

2️⃣ Middleware Support: chi Menang
chi punya middleware bawaan (Logger, CORS, Recoverer, dll.).
gorilla/mux tidak punya middleware, harus pakai third-party.

🔥 Kesimpulan: Pilih Mana?
Gunakan chi jika:	                          | Gunakan gorilla/mux jika:
🚀 Performa tinggi & ringan	🤔                | Sudah terlanjur pakai gorilla/mux
🔥 Butuh middleware bawaan	📌                |Butuh regex routing yang kompleks
✅ Ingin kode yang lebih clean & maintainable |	❌ Tidak masalah dengan library yang tidak aktif
