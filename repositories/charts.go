package repositories

import (
	bolt "go.etcd.io/bbolt"
	"log"
	"errors"
)

func (r *Repository) CreateChartsUser(email string) error {
	db, err := bolt.Open(r.CartsDB, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = db.Update(func(tx *bolt.Tx) error {

		_, err := tx.CreateBucketIfNotExists([]byte(email))
		if err != nil {
			return err
		}
		return nil

	})
	if err != nil {
		return err
	}

	return nil
}
func (r *Repository) UpdateChart(email, title, link string) error {
	db, err := bolt.Open(r.CartsDB, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(email))
		if b == nil {
			return errors.New("Error db user id matching")
		}
		err = b.Put([]byte(title), []byte(link))
		if err != nil {
			return err
		}

		return nil

	})
	if err != nil {
		return err
	}
	return nil
}
func (r *Repository) DeleteChart(email, title string) error {
	db, err := bolt.Open(r.CartsDB, 0600, nil)
	if err != nil {
		return err
	}
	defer db.Close()
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(email))
		if b == nil {
			return errors.New("Error db user id matching")
		}
		k := b.Get([]byte(title))
		if k == nil {
			return errors.New("Error title matching")
		}
		err := b.Delete([]byte(title))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
func (r *Repository) ReadCharts(email string) (map[string]string, error) {
	db, err := bolt.Open(r.CartsDB, 0400, &bolt.Options{ReadOnly: true})
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rec := make(map[string]string)

	err = db.View(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte(email))
		if b == nil {
			return errors.New("Error db user id matching")
		}
		cur := tx.Bucket([]byte(email)).Cursor()
		for k, v := cur.First(); k != nil; k, v = cur.Next() {
			rec[string(k)] = string(v)
		}
		return nil
	})
	if err != nil {
		return map[string]string{}, err
	}

	return rec, nil
}
