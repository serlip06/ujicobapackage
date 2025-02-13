package module

import (
	"context"
	//"errors"
	"fmt"
	"github.com/serlip06/ujicobapackage/model"
	"time"
 	//"net/http"
	"go.mongodb.org/mongo-driver/bson"
	//"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	//"encoding/json"
   // "golang.org/x/crypto/bcrypt"
)


func MongoConnectdb(dbname string) (db *mongo.Database) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
	}
	return client.Database(dbname)
}

// Handler untuk menerima input registrasi
func SignupHandler(req model.SignupRequest, db *mongo.Database) (string, error) {
	// Membuat objek registrasi pengguna tanpa hash password
	pengguna := model.Pengguna{
		Username:  req.Username,
		Password:  req.Password, // Menyimpan password dalam bentuk plaintext
		CreatedAt: time.Now(),
	}

	// Simpan data pengguna ke dalam koleksi "penggunas" (user resmi)
	err := SavePengguna(pengguna, db)
	if err != nil {
		return "", fmt.Errorf("failed to save user: %v", err)
	}

	// Return sukses setelah registrasi
	return "Registration successful", nil
}

// Fungsi untuk menyimpan pengguna ke dalam koleksi "penggunas"
func SavePengguna(pengguna model.Pengguna, db *mongo.Database) error {
	collection := db.Collection("penggunas")
	_, err := collection.InsertOne(context.Background(), pengguna)
	return err
}

// Fungsi untuk mencari pengguna berdasarkan username
func FindPenggunaByUsername(username string, db *mongo.Database) (model.Pengguna, error) {
	collection := db.Collection("penggunas")
	var pengguna model.Pengguna
	err := collection.FindOne(context.Background(), bson.M{"username": username}).Decode(&pengguna)
	if err != nil {
		return model.Pengguna{}, fmt.Errorf("user not found: %v", err)
	}
	return pengguna, nil
}

// Fungsi untuk menangani login pengguna
func SigninHandler(req model.SigninRequest, db *mongo.Database) (model.Pengguna, string, error) {
	// Mencari pengguna berdasarkan username
	pengguna, err := FindPenggunaByUsername(req.Username, db)
	if err != nil {
		return model.Pengguna{}, "", fmt.Errorf("login failed: %v", err)
	}

	// Verifikasi password plaintext
	if pengguna.Password != req.Password {
		return model.Pengguna{}, "", fmt.Errorf("invalid password")
	}

	// Return pengguna dan pesan sukses jika login berhasil
	return pengguna, "Login successful", nil
}

func GetAllPengguna(db *mongo.Database, col string) (data []model.Pengguna) {
	karyawan := db.Collection(col)
	filter := bson.M{}
	cursor, err := karyawan.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetALLData :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}













































































// Contoh menyimpan data registrasi

//untuk resminya namanya savePendingRegistration namacolection-nya = pending_registrations
// func SaveUnverifiedUsers(registration model.UnverifiedUsers, db*mongo.Database) error {
//     collection := db.Collection("unverified_users")
//     _, err := collection.InsertOne(context.Background(), registration)
//     return err
// }

// Contoh memindahkan data dari pending ke users

// untuk resminya pake nama approveRegistration 
// inget collection pending_registrations(resmi) ganti sama unverified_users(ujicoba)
// // collection users(resmi)  pengguna(ujicoba)
// func ConfirmRegistration(id string, db *mongo.Database) (model.UnverifiedUsers, model.Pengguna, error) {
// 	// function yang dipake untuk mindahil data progress ke colekcion pengguna 
// 	collectionUnverifiedusers := db.Collection("unverified_users")
// 	collectionPengguna := db.Collection("penggunas")

// 	objID, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		return model.UnverifiedUsers{}, model.Pengguna{}, err
// 	}

// 	var unverifiedUser model.UnverifiedUsers
// 	err = collectionUnverifiedusers.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&unverifiedUser)
// 	if err != nil {
// 		return model.UnverifiedUsers{}, model.Pengguna{}, err
// 	}

// 	pengguna := model.Pengguna{
// 		Username:  unverifiedUser.Username,
// 		Password:  unverifiedUser.Password,
// 		Role:      unverifiedUser.Role,
// 		CreatedAt: time.Now(),
// 	}

// 	_, err = collectionPengguna.InsertOne(context.Background(), pengguna)
// 	if err != nil {
// 		return model.UnverifiedUsers{}, model.Pengguna{}, err
// 	}

// 	_, err = collectionUnverifiedusers.DeleteOne(context.Background(), bson.M{"_id": objID})
// 	return unverifiedUser, pengguna, err
// }

// //handler registernya aja 


// // Handler untuk menerima input registrasi
// func SignupHandler(req model.SignupRequest, db *mongo.Database) (string, error) {
// 	// Proses hash password
// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
// 	if err != nil {
// 		return "", fmt.Errorf("failed to hash password: %v", err)
// 	}

// 	// Membuat objek registrasi
// 	registration := model.UnverifiedUsers{
// 		Username:    req.Username,
// 		Password:    string(hashedPassword),
// 		Role:        req.Role,
// 		SubmittedAt: time.Now(),
// 	}

// 	// Panggil fungsi untuk menyimpan data pengguna
// 	err = SaveUnverifiedUsers(registration, db)
// 	if err != nil {
// 		return "", fmt.Errorf("failed to save registration: %v", err)
// 	}

// 	// Return success message
// 	return "Registration submitted, waiting for admin approval", nil
// }
