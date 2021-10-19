package hashdb

import (
	"container/list"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type HashDB struct {
	conn *sql.DB
}

type Software struct {
	HostType  string
	Path      string
	Algorithm int64
	Data      []byte
}

func ConnectDB(datasource string) (db *HashDB, err error) {
	conn, err := sql.Open("sqlite3", datasource)
	if err != nil {
		return nil, err
	}
	db = new(HashDB)
	db.conn = conn

	if err = db.InitDB(); err != nil {
		db.Close()
		return nil, err
	}

	return db, err
}

func (v *HashDB) Close() {
	v.conn.Close()
}

func (v *HashDB) InitDB() error {
	query := "CREATE TABLE IF NOT EXISTS `hashtable` ("
	query += "`hosttype` TEXT, "
	query += "`path` TEXT, "
	query += "`algorithm` INTEGER, "
	query += "`data` BLOB, "
	query += "PRIMARY KEY(`hosttype`, `path`)"
	query += ")"
	_, err := v.conn.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func (v *HashDB) InsertUpdateData(hostType, path string, algorithm int64, data []byte) (int64, error) {
	query := "INSERT OR REPLACE INTO `hashtable` (`hosttype`,`path`, `algorithm`, `data`) VALUES (?,?,?,?)"
	result, err := v.conn.Exec(query, hostType, path, algorithm, data)
	if err != nil {
		return -1, err
	}

	return result.LastInsertId()
}

func (v *HashDB) InsertIgnoreData(hostType, path string, algorithm int64, data []byte) (int64, error) {
	query := "INSERT OR IGNORE INTO `hashtable` (`hosttype`,`path`, `algorithm`, `data`) VALUES (?,?,?,?)"
	result, err := v.conn.Exec(query, hostType, path, algorithm, data)
	if err != nil {
		return -1, err
	}

	return result.LastInsertId()
}

func (v *HashDB) GetData() (*list.List, error) {
	// 데이터 조회
	query := "SELECT `hosttype`,`path`, `algorithm`, `data` FROM `hashtable`"

	rows, err := v.conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results := list.New()
	for rows.Next() {
		row := new(Software)
		err := rows.Scan(&row.HostType, &row.Path, &row.Algorithm, &row.Data)
		if err != nil {
			return nil, err
		}
		results.PushBack(row)
	}

	return results, nil
}

func (v *HashDB) Count() (int64, error) {
	// 데이터 조회
	query := "SELECT count(0) FROM `hashtable`"

	var rows *sql.Rows = nil
	var err error
	if rows, err = v.conn.Query(query); err != nil {
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

func (v *HashDB) DeleteData(hosttype, path string) (int64, error) {
	query := "DELETE FROM `hashtable` WHERE `hosttype` = ? and `path` = ?"
	result, err := v.conn.Exec(query, hosttype, path)
	if err != nil {
		return -1, err
	}

	return result.RowsAffected()
}
