package repositories

import (
	"bytes"
	bolt "go.etcd.io/bbolt"
	"log"
	"time"
	"errors"
)

type User struct {
	Number int    `json:"number"`
	Visit  int    `json:"visit"`
	Stamp  string `json:"stamp"`
}

func (r *Repository) CreateUser(email string) error {
	db, err := bolt.Open(r.UsersDB, 0600, nil)
	if err != nil {
		return err
	}
	defer db.Close()
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(email))
		if err != nil {
			log.Fatal(err)
		}
		return nil
	})
	return err
}
func (r *Repository) CheckUserAction(email string, action string) (string, error) {
	db, err := bolt.Open(r.UsersDB, 0400, &bolt.Options{})
	if err != nil {
		return "", err
	}
	defer db.Close()
	//reference time in RFC3339
	refTime := []byte("2006-01-02T15:04:05Z")
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(email))
		if b == nil {
			return errors.New("Error db user id matching")
		}
		cur := b.Cursor()
		for k, v := cur.First(); k != nil && v != nil; k, v = cur.Next() {
			if (bytes.Compare(k, refTime) > 0) && bytes.Equal(v, []byte(action)) {
				refTime = k
			}
		}
		return nil
	})
	return string(refTime), nil
}

func (r *Repository) UpdateUser(email string, action string) error {
	db, err := bolt.Open(r.UsersDB, 0600, &bolt.Options{})
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(email))
		if b == nil {
			return errors.New("Error db user id matching")
		}
		err = b.Put([]byte(time.Now().Format(time.RFC3339)), []byte(action))
		if err != nil {
			return err
		}
		return nil
	})

	return err
}
