package db

import (
    "gopkg.in/pg.v4"
)

var (
    pgOpts = &pg.Options{
        Addr: "localhost:5432",
        User: "arnestaphorsius",
        Password: "password",
        Database: "postgis",
    }
    _this Datastore
)

type Datastore struct {
    db *pg.DB
}

func init() {
    db := pg.Connect(pgOpts)

    _this.db = db
}

func Connection() *pg.DB {
    return _this.db
}