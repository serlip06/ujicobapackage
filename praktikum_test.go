package _714220023

import (
	"fmt"
	"testing"
	"context"
	"github.com/serlip06/ujicobapackage/model"
	"github.com/serlip06/ujicobapackage/module"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
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
		Nama:         "Ruud Gullit",
		Phone_number: "628456456222222",
		Jabatan:      "Football Player",
		Jam_kerja:    []model.JamKerja{jamKerja1, jamKerja2},
		Hari_kerja:   []string{"Senin", "Selasa"},
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
	id := "665991fb37646aa6f1c8a892"
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

func TestDeletePresensiByID(t *testing.T) {
	id := "6412ce78686d9e9ba557cf8a" // ID data yang ingin dihapus
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("error converting id to ObjectID: %v", err)
	}

	err = module.DeletePresensiByID(objectID, module.MongoConn, "presensi")
	if err != nil {
		t.Fatalf("error calling DeletePresensiByID: %v", err)
	}

	// Verifikasi bahwa data telah dihapus dengan melakukan pengecekan menggunakan GetPresensiFromID
	_, err = module.GetPresensiFromID(objectID, module.MongoConn, "presensi")
	if err == nil {
		t.Fatalf("expected data to be deleted, but it still exists")
	}
}

func TestSignupHandler(t *testing.T) {
	// Setup test database
	db := module.MongoConnectdb("tesdb2024")

	// Test case input
	req := model.SignupRequest{
		Username: "testuser",
		Password: "testpassword",
		Role:     "customer",
	}

	// Call the signupHandler function
	message, err := module.SignupHandler(req, db)

	// Test if there were no errors
	if err != nil {
		t.Errorf("Error in signupHandler: %v", err)
	}

	// Test if the message is as expected
	expectedMessage := "Registration submitted, waiting for admin approval"
	if message != expectedMessage {
		t.Errorf("Expected message: %s, got: %s", expectedMessage, message)
	}

	// Verify if the user was inserted into the database
	collection := db.Collection("unverified_users")
	var result model.UnverifiedUsers
	err = collection.FindOne(context.TODO(), bson.M{"username": req.Username}).Decode(&result)
	if err != nil {
		t.Fatalf("Failed to find user in the database: %v", err)
	}

	// Check if the user data is correct
	if result.Username != req.Username {
		t.Errorf("Expected username: %s, got: %s", req.Username, result.Username)
	}

	// Verify if the password is correctly hashed
	err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(req.Password))
	if err != nil {
		t.Fatalf("Password hash mismatch: %v", err)
	}

	// Print confirmation message
	fmt.Printf("User %s successfully registered and saved to the database.\n", req.Username)
}
