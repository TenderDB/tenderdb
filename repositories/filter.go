package repositories

import (
	"bytes"
	bolt "go.etcd.io/bbolt"
	"log"
	"fmt"
	"strconv"
	"strings"
	"github.com/tenderdb/tenderdb/domain"
)

type CodeYearPrice struct {
	Code  string
	Year  string
	Price string
}

func initYearsMap(years []string) map[string]string {
	yearsMap := make(map[string]string)
	for _, year := range years {
		yearsMap[year] = "0"
	}
	return yearsMap
}

func (r *Repository) GetChartOnPrefixFilter(okpd, region, innCustomer, innSupplier string) ([]domain.Dataset, error) {

	regionBytes := []byte(region)
	innCustomerBytes := []byte(innCustomer)
	innSupplierBytes := []byte(innSupplier)
	
	codeYearPrice := make([]CodeYearPrice, 0, 8)

	for _, year := range r.Years {
		for _, fz := range r.FZs {

			db, err := bolt.Open(fmt.Sprintf(r.MaxDB, fz, year), 0400, &bolt.Options{ReadOnly: true})
			if err != nil {
				log.Println(err)
				return []domain.Dataset{}, nil
			}
			//Create array  with size of number of okpds
			err = db.View(func(tx *bolt.Tx) error {
				curBucket := tx.Cursor()

				for kb, vb := curBucket.Seek([]byte(okpd)); kb != nil && vb == nil && bytes.HasPrefix(kb, []byte(okpd)); kb, vb = curBucket.Next() {
					curRecord := tx.Bucket(kb).Cursor()
					sum := 0
					inc := 0

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

						inc, _ = strconv.Atoi(strings.Split(string(vr), ".")[0])

						sum = sum + inc
					}
					if sum == 0 {
						continue
					}

					codeYearPrice = append(codeYearPrice, CodeYearPrice{Code: string(kb) + "(" + fz + ")", Year: year, Price: strconv.Itoa(sum)})

				}
				return nil
			})
			if err != nil {
				log.Println(err)
				return []domain.Dataset{}, err
			}

			db.Close()

		}
	}
	codeYearPriceMap := make(map[string]map[string]string)

	for _, record := range codeYearPrice {
		if len(codeYearPriceMap[record.Code]) == 0 {
			codeYearPriceMap[record.Code] = make(map[string]string)
			codeYearPriceMap[record.Code] = initYearsMap(r.Years)
		}
		codeYearPriceMap[record.Code][record.Year] = record.Price
	}

	var charts []domain.Dataset

	for code, yearPriceMap := range codeYearPriceMap {
		charts = append(charts, domain.Dataset{Name: code, Data: yearPriceMap})
	}

	return charts, nil
}
