package repositories

import (
	"bytes"
	bolt "go.etcd.io/bbolt"
	"github.com/tenderdb/tenderdb/domain"
)

func (r *Repository) GetChartOnPrefix(okpd string) ([]domain.Dataset, error) {
	db, err := bolt.Open(r.MinDB, 0400, &bolt.Options{ReadOnly: true})
	if err != nil {
		return []domain.Dataset{}, err
	}
	defer db.Close()
	var charts []domain.Dataset
	//Create array  with size of number of okpds
	err = db.View(func(tx *bolt.Tx) error {

		curBucket := tx.Cursor()
		prefix := []byte(okpd)

		for kb, vb := curBucket.Seek(prefix); kb != nil && vb == nil && bytes.HasPrefix(kb, prefix); kb, vb = curBucket.Next() {
			yearMap := make(map[string]string)
			curRecord := tx.Bucket(kb).Cursor()
			for kr, vr := curRecord.First(); kr != nil; kr, vr = curRecord.Next() {
				yearMap[string(kr)] = string(vr)
			}
			charts = append(charts, domain.Dataset{Name: string(kb), Data: yearMap})

		}

		return nil
	})
	if err != nil {
		return []domain.Dataset{}, err
	}

	return charts, nil

}
