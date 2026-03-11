package models

// Archetype represents a single Yu-Gi-Oh! card archetype as returned by
// /archetypes.php.
type Archetype struct {
	// ArchetypeName is the archetype name (e.g. "Blue-Eyes").
	ArchetypeName string `json:"archetype_name"`
}
