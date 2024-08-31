1.  Bikin authentication tanpa jwt
package tambahannya bigcache dan go/x/crypto/bcrypt
<!-- tapi disarankan pakai jwt jika sudah bisa integrasinya -->
2.  Bikin inti dari app e wallet
<!-- tappi disarankan diperbaiki lagi karena masih banyak yg perlu dilakukan -->
3.  menganti bigcache dengan redis
4.  bikin sebuah notifikasi secara realtime
<!-- menggunkan sse -->
5.  integration top up dengan midtrans
<!-- package tambahan go-midtrans -->
6.  implemntasi mobile pin
7.  mendeteksi perubahan lokasi login
    <!-- ini untuk pengecekan user ini pakai vpn atau tidak, karena jika dia pakai ip jakarta, kemudian  jadi ip londong maka tidak bisa
        kecuali waktu yg ditempuh dari jarak tersebut masih masuk akal -->

        <!-- dan sebeanrnya masih banyak lagi jika mau implemntasikan factorized
        misal biometrik atau registrasi device id, yg mana membuat 1 akun hanya bisa pakai 1 device -->

8.  bikin fitur log terpusat meggunkan logrus, elasticsearch ,kibana, filebeat
    <!-- install file beat dan konfiguration file log yg ingin di eksekusi -->
    <!-- intall elasticseacrch dan kibana, disini saya coba pakai docker -->
    <!-- install dependesi golang untuk  integrasi log dan elastic search-->
    <!-- go get github.com/sirupsen/logrus go.elastic.co/ecslogrus -->

<!-- masuk ke kibana :5601 -->
