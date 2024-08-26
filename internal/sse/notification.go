package sse

import (
	"ahyalfan/golang_e_money/dto"
	"bufio"
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type notificationSse struct {
	hub *dto.Hub // disini kita masukan pointer, agar apa yg diubah akan mempengaruhi lainya
}

func NewNotificationSse(app *fiber.App, hub *dto.Hub, authMid fiber.Handler) {
	sse := notificationSse{hub: hub}

	app.Get("/api/sse/notifications", authMid, sse.Listen)
}

func (sse *notificationSse) Listen(ctx *fiber.Ctx) error {
	// c, cancel := context.WithCancel(ctx.Context())
	// defer cancel()

	ctx.Set("Content-Type", "text/event-stream") // syarat sse

	user := ctx.Locals("x-user").(dto.UserData)
	sse.hub.NotificationChannel[user.ID] = make(chan dto.NotificationData)

	// event: nama eventnya
	// data : datanya. tapi harus ada enter 2 kali, karena ini formar sse itu sendiri
	ctx.Context().SetBodyStreamWriter(func(w *bufio.Writer) {
		event := fmt.Sprintf("event: %s\n"+
			"data: \n\n", "initial")
		_, _ = fmt.Fprint(w, event)
		_ = w.Flush() // memastikan data dikirim

		for notification := range sse.hub.NotificationChannel[user.ID] {
			data, _ := json.Marshal(notification)
			event := fmt.Sprintf("event: %s\n"+
				"data: %s\n\n", "new_notification", string(data))
			_, _ = fmt.Fprint(w, event)
			_ = w.Flush() // memastikan data dikirim lagi
		}

	})

	return nil
}

// ctx.Context(): Mengakses konteks (context) dalam HTTP request untuk melakukan operasi lebih lanjut.
// SetBodyStreamWriter: Metode ini memungkinkan Anda untuk mengatur penulis (writer) untuk tubuh (body) respons HTTP.
//  Ini biasanya digunakan untuk menulis data secara bertahap ke klien, yang sangat berguna untuk SSE.

// event := fmt.Sprintf("event: %s\n" + "data: \n\n", "initial"): Menyusun pesan awal yang dikirimkan ke klien.
//  Ini adalah format standar untuk SSE, di mana event mendefinisikan jenis acara (event), dan data berisi payload pesan.
// _, _ = fmt.Fprint(w, event): Menulis pesan awal ke buffer writer (w).
// _ = w.Flush(): Memastikan bahwa data yang telah ditulis ke buffer benar-benar dikirim ke klien.
//  Ini penting karena data mungkin masih ada di buffer dan belum dikirim sampai Flush dipanggil.

// for notification := range sse.hub.NotificationChannel[user.ID]: Menggunakan loop untuk mendengarkan pesan yang masuk dari channel notifikasi yang terkait dengan user.ID. Selama channel ini aktif dan mengirimkan notifikasi, loop akan terus berjalan.
// data, _ := json.Marshal(notification): Mengubah notifikasi menjadi format JSON untuk dikirim ke klien.
// event := fmt.Sprintf("event: %s\n" + "data: %s\n\n", "new_notification", string(data)): Menyusun pesan SSE untuk notifikasi baru, di mana event adalah jenis acara (misalnya, "new_notification"), dan data berisi data JSON dari notifikasi.
// _, _ = fmt.Fprint(w, event): Menulis pesan notifikasi baru ke buffer writer.
// _ = w.Flush(): Memastikan bahwa data notifikasi dikirim ke klien segera setelah ditulis.
