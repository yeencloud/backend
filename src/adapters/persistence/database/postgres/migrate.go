package postgres

import "github.com/davecgh/go-spew/spew"

func (db *Database) Migrate() {
	err := db.engine.AutoMigrate(Settings{})
	if err != nil {
		spew.Dump(err)
		return
	}
}
