## Glans

simple POC of mysql proxy, dengan studi kasus aplikasi php (ci3, wp, etc) yang bertindak sebagai : 

- connection pool, karena php itu state-less, open connection, dan close, tidak cocok ketika terjadi lonjakan traffic
- smart cache (soon), ini khusus untuk table yang read-intensive, jadi cache nya level di proxy tapi tetep bisa di validate
- load balancer (soon), anda punya lebih dari 2 instance, gaskan, master-svaled oke


## Kenapa nggak pakai vites aja bang?

bang, ini buat tugas laporan magang gue bang, ini aja di kerjain cuman 4 harian, ya kan biar keren aja, aslinya mah cuman implement mysql protocol itu aja pakai library, aowkowkowkwokwokwo