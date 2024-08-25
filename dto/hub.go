package dto

type Hub struct {
	// disini kita buat penghubungnya
	// sehingga tiap user ada channelnya sendiri
	NotificationChannel map[int64]chan NotificationData
}
