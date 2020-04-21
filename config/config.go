package config

import ("database/sql"
_"github.com/nakagami/firebirdsql"
)

func Koneksi() (db *sql.DB, err error) {

	db, err = sql.Open("firebirdsql", "SYSDBA:MASTERKEY@localhost:3050/C:/firebird_test/YOURDATABASENAME.TCS")

	return
}