package sqlite

import (
	"database/sql"

	"github.com/akshayjha21/Chat-App-in-GO/internal/config"
	// "golang.org/x/tools/go/analysis/passes/defers"
)

type Sqlite struct {
	Db *sql.DB
}

func New(cfg *config.Config)(*Sqlite,error){
	db,err:=sql.Open("sqlite",cfg.StoragePath)
	if err != nil {
		return nil,err
	}
	defer db.Close()
	db.Exec(`
	`)
}
