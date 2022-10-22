package repositories

import (
	"bytes"
	bolt "go.etcd.io/bbolt"
)

func (r *Repository) Autocomplete(word string) (map[string]string, error) {
	db, err := bolt.Open(r.AutocompleteDB, 0640, nil)
	if err != nil {
		return map[string]string{}, err
	}
	defer db.Close()
	wordByteLower := bytes.ToLower([]byte(word))
	spaceWordByteLower := bytes.ToLower([]byte(" " + word))
	list := make(map[string]string, 0)
	err = db.View(func(tx *bolt.Tx) error {
		c := tx.Cursor()
		first, _ := c.First()
		b := tx.Bucket(first)
		cur := b.Cursor()
		for k, v := cur.First(); k != nil; k, v = cur.Next() {
			if bytes.HasPrefix(bytes.ToLower(k), wordByteLower) {
				list[string(k)] = string(v)
				if len(list) >= 10 {
					break
				}
			}
		}
		for k, v := cur.First(); k != nil; k, v = cur.Next() {
			if bytes.Contains(bytes.ToLower(k), spaceWordByteLower) {
				list[string(k)] = string(v)
				if len(list) >= 10 {
					break
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
