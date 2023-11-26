package postgres

import (
	"gorm.io/gorm"
)

type settingsRepo struct {
	db *gorm.DB
}

type Settings struct {
	gorm.Model

	Key   string
	Value string
}

func (db *Database) GetSettingsValue(key string) string {
	return ""
}

func (db *Database) SetSettingsValue(key string, value string) {

}
