package models

// DBVersion contains the current database version information as returned by
// /checkDBVer.php.
type DBVersion struct {
	// DatabaseVersion is the version string of the YGOProDeck database.
	DatabaseVersion string `json:"database_version"`
	// LastUpdate is a timestamp string indicating when the database was last
	// updated.
	LastUpdate string `json:"last_update"`
}
