package repositories

import (
	"bufio"
	"bytes"
	"fmt"
	bolt "go.etcd.io/bbolt"
)

func (r *Repository) CsvDownload(okpd, region, innCustomer, innSupplier string) ([]byte, error) {
	writebuffer := bytes.NewBuffer(nil)
	writer := bufio.NewWriter(writebuffer)
	if _, err := writer.WriteString("Год,ФЗ,Регион,ИНН поставщика,ИНН заказчика,Номер контракта,Номер лота,Сумма лота,Комментарий\n"); err != nil {
		return []byte(""), err
	}
	regionBytes := []byte(region)
	innCustomerBytes := []byte(innCustomer)
	innSupplierBytes := []byte(innSupplier)
	var comment string

	for _, year := range r.Years {
		for _, fz := range r.FZs {
			db, err := bolt.Open(fmt.Sprintf(r.MaxDB, fz, year), 0400, &bolt.Options{ReadOnly: true})

			if err != nil {
				return []byte(""), err
			}
			err = db.View(func(tx *bolt.Tx) error {
				curBucket := tx.Cursor()

				for kb, vb := curBucket.Seek([]byte(okpd)); kb != nil && vb == nil && bytes.HasPrefix(kb, []byte(okpd)); kb, vb = curBucket.Next() {
					curRecord := tx.Bucket(kb).Cursor()

					for kr, vr := curRecord.First(); kr != nil; kr, vr = curRecord.Next() {
						if !bytes.Contains(kr[0:2], regionBytes) {
							continue
						}
						if !bytes.Contains(kr[3:13], innCustomerBytes) {
							continue
						}
						if !bytes.Contains(kr[14:26], innSupplierBytes) {
							continue
						}
						if !bytes.Contains(kr[14:24], innSupplierBytes) {
							continue
						}

						if string(vr) == "0" {
							comment = "Сумма лота не выделена"
						} else {
							comment = "-"
						}

						if _, err := writer.WriteString(year + "," + fz + "," + string(kr) + "," + string(vr) + "," + comment + "\n"); err != nil {
							return err
						}
					}
				}
				return nil
			})
			if err != nil {
				return []byte(""), err
			}

			db.Close()
		}
	}
	writer.Flush()
	return writebuffer.Bytes(), nil
}
