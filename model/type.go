package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Karyawan struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama         string             `bson:"nama,omitempty" json:"nama,omitempty"`
	Phone_number string             `bson:"phone_number,omitempty" json:"phone_number,omitempty"`
	Jabatan      string             `bson:"jabatan,omitempty" json:"jabatan,omitempty"`
	Jam_kerja    []JamKerja         `bson:"jam_kerja,omitempty" json:"jam_kerja,omitempty"`
	Hari_kerja   []string           `bson:"hari_kerja,omitempty" json:"hari_kerja,omitempty"`
}
//[] adalah array sedangkan jamkerja berasal dari struk dibawah ini 
type JamKerja struct {
	Durasi     int      `bson:"durasi,omitempty" json:"durasi,omitempty"`
	Jam_masuk  string   `bson:"jam_masuk,omitempty" json:"jam_masuk,omitempty"`
	Jam_keluar string   `bson:"jam_keluar,omitempty" json:"jam_keluar,omitempty"`
	Gmt        int      `bson:"gmt,omitempty" json:"gmt,omitempty"`
	Hari       []string `bson:"hari,omitempty" json:"hari,omitempty"`
	Shift      int      `bson:"shift,omitempty" json:"shift,omitempty"`
	Piket_tim  string   `bson:"piket_tim,omitempty" json:"piket_tim,omitempty"`
}

type Presensi struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Longitude    float64            `bson:"longitude,omitempty" json:"longitude,omitempty"`
	Latitude     float64            `bson:"latitude,omitempty" json:"latitude,omitempty"`
	Location     string             `bson:"location,omitempty" json:"location,omitempty"`
	Phone_number string             `bson:"phone_number,omitempty" json:"phone_number,omitempty"`
	Datetime     primitive.DateTime `bson:"datetime,omitempty" json:"datetime,omitempty"`
	Checkin      string             `bson:"checkin,omitempty" json:"checkin,omitempty"`
	Biodata      Karyawan           `bson:"biodata,omitempty" json:"biodata,omitempty"`
}
//karyawan berasal dari tabel karyawan di atas 
//moggo db bentuk satabasenya objek dengan tipe code json, struct itu berhubungan satu sama lain

//untuk uji coba login user (kalo resminya PendingRegistration)
type UnverifiedUsers struct {
    ID          string    `bson:"_id,omitempty" json:"id,omitempty"` // ID unik dari MongoDB
    Username    string    `bson:"username" json:"username"`          // Username pengguna
    Password    string    `bson:"password" json:"password"`          // Password dalam bentuk hash
    Role        string    `bson:"role" json:"role"`                  // Peran pengguna (customer, kasir, operator)
    SubmittedAt time.Time `bson:"submitted_at" json:"submitted_at"`  // Waktu registrasi
}

//untuk pengguna (kalo resminya User)
type Pengguna struct {
    ID        string    `bson:"_id,omitempty" json:"id,omitempty"` // ID unik dari MongoDB
    Username  string    `bson:"username" json:"username"`          // Username pengguna
    Password  string    `bson:"password" json:"password"`          // Password dalam bentuk hash
    Role      string    `bson:"role" json:"role"`                  // Peran pengguna (admin, customer, kasir, operator)
    CreatedAt time.Time `bson:"created_at" json:"created_at"`      // Waktu pembuatan akun
}
// permintaan untuk registrasi 
// RegisterRequest (resmi) SignupRequest(ujicoba)
type SignupRequest struct {
    Username string `json:"username"`
    Password string `json:"password"`
    Role     string `json:"role"`
}

// permintaan untuk login (kalo di dbresmi namaya :LoginRequest )
type SigninRequest struct {
    Username string `json:"username"` // Username pengguna
    Password string `json:"password"` // Password pengguna
}

//respon sis sitemnya kalo ada yang request mau daftar (kalo resminya namanya Response)
type AccessResponse struct {
    Status  string `json:"status"`  // Status operasi (success, error)
    Message string `json:"message"` // Pesan deskripsi
    // Data    any    `json:"data"`    // Data tambahan (opsional)
}

// untuk admin 
type Admin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}