package _714220023

import (
	"fmt"
	"testing"

	"github.com/indrariksa/cobapakcage/model"
	"github.com/indrariksa/cobapakcage/module"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestInsertPresensi(t *testing.T) {
	var jamKerja1 = model.JamKerja{
		Durasi:     8,
		Jam_masuk:  "08:00",
		Jam_keluar: "16:00",
		Gmt:        7,
		Hari:       []string{"Senin", "Rabu", "Kamis"},
		Shift:      1,
		Piket_tim:  "Piket A",
	}
	var jamKerja2 = model.JamKerja{
		Durasi:     8,
		Jam_masuk:  "09:00",
		Jam_keluar: "17:00",
		Gmt:        7,
		Hari:       []string{"Sabtu"},
		Shift:      2,
		Piket_tim:  "",
	}

	long := 98.345345
	lat := 123.561651
	lokasi := "Amsterdam"
	phonenumber := "6811110023231"
	checkin := "masuk"
	biodata := model.Karyawan{
		Nama:        "Ruud Gullit",
		Phone_number: "628456456222222",
		Jabatan:     "Football Player",
		Jam_kerja:   []model.JamKerja{jamKerja1, jamKerja2},
		Hari_kerja:  []string{"Senin", "Selasa"},
	}
	insertedID, err := module.InsertPresensi(module.MongoConn, "presensi", long, lat, lokasi, phonenumber, checkin, biodata)
	if err != nil {
		t.Errorf("Error inserting data: %v", err)
	}
	fmt.Printf("Data berhasil disimpan dengan id %s", insertedID.Hex())
}

func TestGetKaryawanFromPhoneNumber(t *testing.T) {
	phonenumber := "68122221814"
	biodata, err := module.GetKaryawanFromPhoneNumber(phonenumber, module.MongoConn, "presensi")
	if err != nil {
		t.Fatalf("error calling GetKaryawanFromPhoneNumber: %v", err)
	}
	fmt.Println(biodata)
}

func TestGetPresensiFromID(t *testing.T) {
	id := "641898e76dd7bd217d69762a"
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("error converting id to ObjectID: %v", err)
	}
	biodata, err := module.GetPresensiFromID(objectID, module.MongoConn, "presensi")
	if err != nil {
		t.Fatalf("error calling GetPresensiFromID: %v", err)
	}
	fmt.Println(biodata)
}

func TestGetAll(t *testing.T) {
	data := module.GetAllPresensi(module.MongoConn, "presensi")
	fmt.Println(data)
}