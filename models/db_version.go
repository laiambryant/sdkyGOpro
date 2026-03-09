package models

type DBVersion struct {
	DatabaseVersion string `json:"database_version"`
	LastUpdate      string `json:"last_update"`
}
