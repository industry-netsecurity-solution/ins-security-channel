package firewall

import (
	"container/list"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"sync"
)

const (
	WB_BLACK = 0
	WB_WHITE = 1
	TYPE_IP  = 0
	TYPE_MAC = 1
)

type FirewallDB struct {
	conn     *sql.DB
	locker   *sync.Mutex
}

type FirewallRule struct {
	Id          int64
	WhiteBlack  int
	AddressType int
	Address   string
}

func ConnectDB(datasource string) (db *FirewallDB, err error) {
	conn, err := sql.Open("sqlite3", datasource)
	if err != nil {
		return nil, err
	}
	db = new(FirewallDB)
	db.conn = conn
	db.locker = &sync.Mutex{}

	if err = db.AutoVacuum(); err != nil {
		db.Close()
		return nil, err
	}

	if err = db.Reduce(); err != nil {
		db.Close()
		return nil, err
	}

	if err = db.InitDB(); err != nil {
		db.Close()
		return nil, err
	}

	return db, err
}

func (v *FirewallDB) Close() {
	v.conn.Close()
}

func (v *FirewallDB) Lock() {
	v.locker.Lock()
}

func (v *FirewallDB) Unlock() {
	v.locker.Unlock()
}

func (v *FirewallDB) InitDB() error {
	query := "CREATE TABLE IF NOT EXISTS `hosttable` ("
	query += "`id` INTEGER PRIMARY KEY AUTOINCREMENT, "
	query += "`whiteblack` INTEGER, "
	query += "`addresstype` INTEGER, "
	query += "`address` TEXT"
	query += ")"
	if _, err := v.conn.Exec(query); err != nil {
		return err
	}

	if _, err := v.createUniqueIndex(); err != nil {
		return err
	}

	return nil
}

func (v *FirewallDB) AutoVacuum() error {

	query := "PRAGMA auto_vacuum=1"
	_, err := v.conn.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func (v *FirewallDB) Reduce() error {

	query := "VACUUM"
	_, err := v.conn.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func (v *FirewallDB) ResetAutoincrement() (int64, error) {
	query := "UPDATE SQLITE_SEQUENCE SET `seq` = (SELECT MAX(`id`) FROM `hosttable`) WHERE `name` = 'hosttable'"
	result, err := v.conn.Exec(query)
	if err != nil {
		return -1, err
	}

	return result.RowsAffected()
}

func (v *FirewallDB) createUniqueIndex() (int64, error) {
	query := "CREATE UNIQUE INDEX IF NOT EXISTS `hosttable_index` ON `hosttable` (`whiteblack`,`addresstype`, `address`)"
	result, err := v.conn.Exec(query)
	if err != nil {
		return -1, err
	}

	return result.RowsAffected()
}


func (v *FirewallDB) InsertData(whiteblack, addresstype int, address string) (int64, error) {
	query := "INSERT INTO `hosttable` (`whiteblack`,`addresstype`, `address`) VALUES (?,?,?)"
	result, err := v.conn.Exec(query, whiteblack, addresstype, address)
	if err != nil {
		return -1, err
	}

	return result.LastInsertId()
}

func (v *FirewallDB) InsertIgnoreData(whiteblack, addresstype int, address string) (int64, error) {
	query := "INSERT OR IGNORE INTO `hosttable` (`whiteblack`,`addresstype`, `address`) VALUES (?,?,?)"
	result, err := v.conn.Exec(query, whiteblack, addresstype, address)
	if err != nil {
		return -1, err
	}

	fmt.Print(result.RowsAffected())

	return result.LastInsertId()
}

func (v *FirewallDB) GetAllHosts() (*list.List, error) {
	// 데이터 조회
	query := "SELECT `id`,`whiteblack`,`addresstype`, `address` FROM `hosttable`"

	rows, err := v.conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results := list.New()
	for rows.Next() {
		row := new(FirewallRule)
		err := rows.Scan(&row.Id, &row.WhiteBlack, &row.AddressType, &row.Address)
		if err != nil {
			return nil, err
		}
		results.PushBack(row)
	}

	return results, nil
}

func (v *FirewallDB) GetHosts(whiteblack, addresstype int) (*list.List, error) {
	// 데이터 조회
	query := "SELECT `id`,`whiteblack`,`addresstype`, `address` FROM `hosttable` WHERE `whiteblack` = ? AND `addresstype` = ?"

	rows, err := v.conn.Query(query, whiteblack, addresstype)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results := list.New()
	for rows.Next() {
		row := new(FirewallRule)
		err := rows.Scan(&row.Id, &row.WhiteBlack, &row.AddressType, &row.Address)
		if err != nil {
			return nil, err
		}
		results.PushBack(row)
	}

	return results, nil
}

func (v *FirewallDB) GetBlackHosts(addresstype int) (*list.List, error) {
	// 데이터 조회
	query := "SELECT `id`,`whiteblack`,`addresstype`, `address` FROM `hosttable` WHERE `whiteblack` = ? AND `addresstype` = ?"

	rows, err := v.conn.Query(query, WB_BLACK, addresstype)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results := list.New()
	for rows.Next() {
		row := new(FirewallRule)
		err := rows.Scan(&row.Id, &row.WhiteBlack, &row.AddressType, &row.Address)
		if err != nil {
			return nil, err
		}
		results.PushBack(row)
	}

	return results, nil
}

func (v *FirewallDB) GetWhiteHosts(addresstype int) (*list.List, error) {
	// 데이터 조회
	query := "SELECT `id`,`whiteblack`,`addresstype`, `address` FROM `hosttable` WHERE `whiteblack` = ? AND `addresstype` = ?"

	rows, err := v.conn.Query(query, WB_WHITE, addresstype)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results := list.New()
	for rows.Next() {
		row := new(FirewallRule)
		err := rows.Scan(&row.Id, &row.WhiteBlack, &row.AddressType, &row.Address)
		if err != nil {
			return nil, err
		}
		results.PushBack(row)
	}

	return results, nil
}

func (v *FirewallDB) Count() (int64, error) {
	// 데이터 조회
	query := "SELECT count(0) FROM `hosttable`"

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

func (v *FirewallDB) DeleteData(whiteblack, addresstype int, address string) (int64, error) {
	query := "DELETE FROM `hosttable` WHERE `whiteblack` = ? AND `addresstype` = ? AND `address` = ?"
	result, err := v.conn.Exec(query, whiteblack, addresstype, address)
	if err != nil {
		return -1, err
	}

	return result.RowsAffected()
}
