package module

import (
	"context"
	"errors"
	"fmt"
	"github.com/serlip06/ujicobapackage/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoConnect(dbname string) (db *mongo.Database) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
	}
	return client.Database(dbname)
}

func InsertOneDoc(db string, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertPresensi(db *mongo.Database, col string, long float64, lat float64, lokasi string, phonenumber string, checkin string, biodata model.Karyawan) (insertedID primitive.ObjectID, err error) {
	presensi := bson.M{
		"longitude":    long,
		"latitude":     lat,
		"location":     lokasi,
		"phone_number": phonenumber,
		"datetime":     primitive.NewDateTimeFromTime(time.Now().UTC()),
		"checkin":      checkin,
		"biodata":      biodata,
	}
	result, err := db.Collection(col).InsertOne(context.Background(), presensi)
	if err != nil {
		fmt.Printf("InsertPresensi: %v\n", err)
		return
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}
func GetKaryawanFromPhoneNumber(phone_number string, db *mongo.Database, col string) (staf model.Presensi, errs error) {
	karyawan := db.Collection(col)
	filter := bson.M{"phone_number": phone_number}
	err := karyawan.FindOne(context.TODO(), filter).Decode(&staf)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return staf, fmt.Errorf("no data found for phone number %s", phone_number)
		}
		return staf, fmt.Errorf("error retrieving data for phone number %s: %s", phone_number, err.Error())
	}
	return staf, nil
}

func GetAllPresensi(db *mongo.Database, col string) (data []model.Presensi) {
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

func GetPresensiFromID(_id primitive.ObjectID, db *mongo.Database, col string) (presensi model.Presensi, errs error) {
	karyawan := db.Collection(col)
	filter := bson.M{"_id": _id}
	err := karyawan.FindOne(context.TODO(), filter).Decode(&presensi)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return presensi, fmt.Errorf("no data found for ID %s", _id)
		}
		return presensi, fmt.Errorf("error retrieving data for ID %s: %s", _id, err.Error())
	}
	return presensi, nil
}