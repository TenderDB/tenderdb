package repositories

import (
	"bytes"
	bolt "go.etcd.io/bbolt"
	"github.com/tenderdb/tenderdb/valid"
)

func (r *Repository)GetList(okpd string) (map[string]string, error) {
	list := make(map[string]string, 0)
	var err error
	lengthOkpd := len(okpd)

	if lengthOkpd == 0 {
		return r.GetMapABC()
	}
	if lengthOkpd == 1 && valid.ABC(okpd) {
		return r.GetMapA(okpd)
	}

	if lengthOkpd > 1 && valid.Okpd(okpd) {

		list, err = r.GetMapPrefix(lengthOkpd, okpd)
		if err != nil {
			return map[string]string{}, err
		}
		return list, nil
	}
	return map[string]string{}, nil
}

func (r *Repository)GetValue(code string) (string, error) {
	db, err := bolt.Open(r.TreeDB, 0400, &bolt.Options{ReadOnly: true})
	if err != nil {
		return "", err
	}
	defer db.Close()
	var resp []byte
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("my"))
		resp = b.Get([]byte(code))
		return nil
	})
	if err != nil {
		return "", err
	}

	return string(resp), nil
}

func (r *Repository)GetMapPrefix(lengthOkpd int, prefixStr string) (map[string]string, error) {
	db, err := bolt.Open(r.TreeDB, 0400, &bolt.Options{ReadOnly: true})
	if err != nil {
		return map[string]string{}, err
	}
	defer db.Close()
	prefix := []byte(prefixStr)
	matrix := []int{0, 0, 5, 5, 5, 7, 7, 8, 12, 12, 12, 12, 12}
	length := matrix[lengthOkpd]
	list := make(map[string]string, 0)
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("my"))
		cur := b.Cursor()
		for k, v := cur.Seek(prefix); bytes.HasPrefix(k, prefix); k, v = cur.Next() {
			if len(k) == length {
				list[string(k)] = string(v)
			}

		}
		if len(list) <= 1 {
			list = map[string]string{}
			length = 12
			for k, v := cur.Seek(prefix); bytes.HasPrefix(k, prefix); k, v = cur.Next() {
				if len(k) == length {
					list[string(k)] = string(v)
				}

			}

		}
		return nil
	})
	if err != nil {
		return map[string]string{}, err
	}

	return list, nil
}
func (r *Repository)GetMapABC() (map[string]string, error) {
	db, err := bolt.Open(r.TreeDB, 0400, &bolt.Options{ReadOnly: true})
	if err != nil {
		return map[string]string{}, err
	}
	defer db.Close()
	list := make(map[string]string, 0)
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("my"))
		letters := "ABCDEFGHIGKLMNOPQRSTU"
		for _, v := range letters {
			list[string(v)] = string(b.Get([]byte(string(v))))
		}

		return nil
	})
	if err != nil {
		return map[string]string{}, err
	}

	return list, nil
}
func (r *Repository)GetMapA(class string) (map[string]string, error) {
	db, err := bolt.Open(r.TreeDB, 0400, &bolt.Options{ReadOnly: true})
	if err != nil {
		return map[string]string{}, err
	}
	defer db.Close()
	list := make(map[string]string, 0)
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(class))
		cur := b.Cursor()
		for k, v := cur.First(); v != nil; k, v = cur.Next() {
			list[string(k)] = string(v)
		}

		return nil
	})
	if err != nil {
		return map[string]string{}, err
	}

	return list, nil
}
