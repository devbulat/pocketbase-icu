package main

import (
    "database/sql"

    "github.com/mattn/go-sqlite3"
    "github.com/pocketbase/dbx"
)

func init() {
    sql.Register("pb_sqlite3", &sqlite3.SQLiteDriver{
        ConnectHook: func(conn *sqlite3.SQLiteConn) error {
            _, err := conn.Exec(`
                PRAGMA busy_timeout = 10000;
                PRAGMA journal_mode = WAL;
                PRAGMA journal_size_limit = 200000000;
                PRAGMA synchronous = NORMAL;
                PRAGMA foreign_keys = ON;
                PRAGMA temp_store = MEMORY;
                PRAGMA cache_size = -32000;
            `, nil)

            return err
        },
    })

    dbx.BuilderFuncMap["pb_sqlite3"] = dbx.BuilderFuncMap["sqlite3"]
}

func icuDBConnect(dbPath string) (*dbx.DB, error) {
    return dbx.Open("pb_sqlite3", dbPath)
}
