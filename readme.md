# Tetang Repository Ini
Repository ini merupakan cloning dari proyek [qiscust](https://bitbucket.org/frchandra/qiscust) dengan bahasa golang dan framework gin.
Aplikasi dalam repository ini adalah service yang mengimplementasikan Custom Agent Allocator untuk layanan Qiscus Omnichannel.
Service ini dapat menampung pesan yang masuk ke Qiscus Omnichannel kemudian dapat mengalokasikan agen yang cocok kedalam ruang obrolan untuk menjawab pesan tersebut.
Layanan ini dapat menerapkan aturan alokasi agen yang dapat dimodifikasi sesuai dengan persyaratan tertentu. Dalam kasus ini, agen diatur supaya 
maksimal hanya dapat melayani dua pesan (berada dalam dua ruang obrolan). Apabila jumlah pesan pasuk melebihi ketersediaan agen, maka pesan tersebut akan
dimasukan kedalam antrean sampai agen selesai melayani pesan-pesan sebelumnya. <br><br>

Untuk memicu proses alokasi agen, aplikasi ini harus menerima notifikasi masukknya pesan baru dari Qiscus Omnichannel 
yang ditujukan ke aplikasi ini meelalui endpoint ```POST http://[baseurl]/api/v1/new_request```. Endpoint ini akan menerima data detail pesan baru yang masuk. Setelah itu aplikasi 
ini akan mencarikan agen yang tersedia untuk melayani pesan masuk tersebut. Apabila tidak ada agen yang dapat melayani pesan tersebut, maka pesan akan dimasukkan 
kedalam antrean di database sampai terdapat agen yang tersedia. Apabila terdapat agen  yang tersedia, maka agen tersebut akan dialokasikan ke ruang obrolan 
tertentu untuk menjawab pesan tersebut.<br><br>

Pesan yang berada dalam antrean akan diproses ketika agen telah selesai melayani pesan-pesan sebelumnya. Ketika agen telah selesai menjawab pesan, maka agen
akan menutup ruang obrolam. Aksi ini akan memicu Qiscus Omnichannel mengirimkan pesan ke aplikasi ini yang memberitahu jika terdapat agen yang selesai melayani pesan.
Lalu Qiscus Omnichannel akan mengirimkan pesan ke aplikasi ini melalui endpoint ```POST http://[baseurl]/api/v1/close+_request```. Notifikasi dari Qiscus Omnichannel
akan memicu aplikasi mengambil pesan selanjutnya dari antrean untuk dialokasikan ke agen yang sesuai.

# Tech Stack
- Go Language v1.19.3
- PostgreSQL v15
- Gin framework v1.8.2
- GORM v1.24.3

<br>
chandra herdiputra, Januari 2022




