package split_csv

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

// 拆分
func Split(pathRecord string, pathSetting string) error {
	records, err := readCsv(pathRecord)
	if err != nil {
		return err
	}

	settings, err := readCsv(pathSetting)
	if err != nil {
		return err
	}

	pathBase := filepath.Dir(pathRecord)

	if len(records) < count(settings) {
		return errors.New("数量不够，请检查文件！")
	}

	for _, s := range settings {
		nRecords, err := strconv.Atoi(s[0])
		if err != nil {
			return err
		}

		nFile, err := strconv.Atoi(s[1])
		if err != nil {
			return err
		}

		for i := 1; i <= nFile; i++ {
			// 创建文件
			path := fmt.Sprintf("output-%d-%d.csv", nRecords, i)
			path = filepath.Join(pathBase, path)

			err = writeCsv(path, records[:nRecords])
			if err != nil {
				return err
			}
			// 从新设置文件记录
			records = records[nRecords:]
		}
	}

	// 生成新的记录文件
	return writeCsv(pathRecord, records)
}

// 写入 CSV 文件
func writeCsv(path string, data [][]string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	w := csv.NewWriter(f)
	// TODO Added the "email" to the first line.
	w.Write([]string{"emails"})
	return w.WriteAll(data)
}

// 读取 CSV 文件
func readCsv(path string) ([][]string, error) {
	f, err := os.Open(path)

	if err != nil {
		return [][]string{}, err
	}

	defer f.Close()

	r := csv.NewReader(f)

	records, err := r.ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return records[1:], nil
}

// count setting
func count(settings [][]string) int {
	count := 0
	for _, s := range settings {
		nRecords, err := strconv.Atoi(s[0])
		if err != nil {
			return 0
		}

		nFile, err := strconv.Atoi(s[1])
		if err != nil {
			return 0
		}

		count = count + nRecords*nFile
	}

	return count
}
