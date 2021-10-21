package cachedb

import (
	"container/list"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)


type CacheDB struct {
	conn         *sql.DB
	tables        int
}

type Data struct {
	Id            int64
	Data        []byte
}

func ConnectDB(datasource string) (*CacheDB, error) {
	conn, err := sql.Open("sqlite3", datasource)
	if err != nil {
		return nil, err
	}
	db := new(CacheDB)
	db.conn = conn

	if err = db.AutoVacuum(); err != nil {
		db.Close()
		return nil, err
	}

	if err = db.Reduce(); err != nil {
		db.Close()
		return nil, err
	}

	return db, err
}

func (v *CacheDB) Close() {
	v.conn.Close()
}

func (v *CacheDB) InitDB(num int) error {
	for i := 0; i < num; i++ {
		tname := fmt.Sprintf("T%02d", i)

		query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS `%s` (`id` INTEGER, `data` BLOB NOT NULL, PRIMARY KEY (`id`))", tname)
		_, err := v.conn.Exec(query)
		if err != nil {
			return err
		}

		v.tables = i+1
	}

	return nil
}

func (v *CacheDB) AutoVacuum() error {

	query := "PRAGMA auto_vacuum=1"
	_, err := v.conn.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func (v *CacheDB) Reduce() error {

	query := "VACUUM"
	_, err := v.conn.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func getInsertQuery(db *CacheDB, table int) (string, error) {
	if db.tables <= table {
		return "", errors.New("Not support table.")
	}

	tname := fmt.Sprintf("T%02d", table)
	query := fmt.Sprintf("INSERT INTO `%s` (`id`,`data`) VALUES (?,?)", tname)

	return query, nil
}

func (v *CacheDB) InsertData(table int, id int64, data []byte) (int64, error) {
	if v.tables <= table {
		return -1, errors.New("Not support table.")
	}

	query, err := getInsertQuery(v, table)
	if err != nil {
		return -1, err
	}
	result, err := v.conn.Exec(query, id, data)
	if err != nil {
		return -1, err
	}

	return result.LastInsertId()
}

func getNextDataQuery(db *CacheDB, table int) (string, error) {
	if db.tables <= table {
		return "", errors.New("Not support table.")
	}

	tname := fmt.Sprintf("T%02d", table)
	query := fmt.Sprintf("SELECT id, data FROM `%s` WHERE id = (SELECT min(id) FROM `%s` WHERE ? < id)", tname, tname)

	return query, nil
}

func (v *CacheDB) GetNextData(table int, curr int64) (int64, []byte, error){
	// 데이터 조회
	query, err := getNextDataQuery(v, table)
	if err != nil {
		return -1, nil, err
	}

	rows, err := v.conn.Query(query, curr)
	if err != nil {
		return -1, nil, err
	}
	defer rows.Close()

	if rows.Next() {
		var id int64
		var data []byte
		err := rows.Scan(&id, &data)
		if err != nil {
			return -1, nil, err
		}

		return id, data, nil
	}

	return -1, nil, nil
}

func getNextLimitDataQuery(db *CacheDB, table int, limit int64) (string, error) {
	if db.tables <= table {
		return "", errors.New("Not support table.")
	}

	tname := fmt.Sprintf("T%02d", table)
	if limit <= 0 {
		query := fmt.Sprintf("SELECT id, data FROM `%s` WHERE ? < id ORDER BY id ASC", tname)
		return query, nil
	}
	query := fmt.Sprintf("SELECT id, data FROM `%s` WHERE ? < id ORDER BY id ASC LIMIT %d", tname, limit)

	return query, nil
}

func (v *CacheDB) GetNextLimitData(table int, curr int64, limit int64) (*list.List, error){
	// 데이터 조회
	query, err := getNextLimitDataQuery(v, table, limit)
	if err != nil {
		return nil, err
	}

	rows, err := v.conn.Query(query, curr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results := list.New()
	for rows.Next() {
		row := new(Data)
		err := rows.Scan(&row.Id, &row.Data)
		if err != nil {
			return nil, err
		}
		results.PushBack(row)
	}

	return results, nil
}


func getCountQuery(db *CacheDB, table int, min, max int64) (string, error) {
	if db.tables <= table {
		return "", errors.New("Not support table.")
	}

	tname := fmt.Sprintf("T%02d", table)

	if min < 0 && 0 <= max {
		query := fmt.Sprintf("SELECT count(0) FROM `%s` WHERE id <= ?", tname)
		return query, nil
	} else if 0 <= min && max < 0 {
		query := fmt.Sprintf("SELECT count(0) FROM `%s` WHERE ? <= id", tname)
		return query, nil
	} else if 0 <= min && 0 <= max {
		query := fmt.Sprintf("SELECT count(0) FROM `%s` WHERE ? <= id and id <= ?", tname)
		return query, nil
	}

	query := fmt.Sprintf("SELECT count(0) FROM `%s`", tname)
	return query, nil
}

func (v *CacheDB) Count(table int, min, max int64) (int64, error){
	// 데이터 조회
	query, err := getCountQuery(v, table, min, max)
	if err != nil {
		return -1, err
	}

	var rows *sql.Rows = nil
	if min < 0 && 0 <= max {
		rows, err = v.conn.Query(query, max)
	} else if 0 <= min && max < 0 {
		rows, err = v.conn.Query(query, min)
	} else if 0 <= min && 0 <= max {
		rows, err = v.conn.Query(query, min, max)
	} else {
		rows, err = v.conn.Query(query)
	}

	if err != nil {
		return -1, err
	}
	defer rows.Close()

	if rows.Next() {
		var count int64
		err := rows.Scan(&count)
		if err != nil {
			return -1, err
		}

		return count, nil
	}

	return 0, nil
}

func getDeleteQuery(db *CacheDB, table int) (string, error) {
	if db.tables <= table {
		return "", errors.New("Not support table.")
	}

	tname := fmt.Sprintf("T%02d", table)
	query := fmt.Sprintf("DELETE FROM `%s` WHERE id = ?", tname)

	return query, nil
}

func (v *CacheDB) DeleteData(table int, id int64) (int64, error) {
	query, err := getDeleteQuery(v, table)
	if err != nil {
		return -1, err
	}
	result, err := v.conn.Exec(query, id)
	if err != nil {
		return -1, err
	}

	return result.RowsAffected()
}