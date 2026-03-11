// Package enums defines named string types for the enumerated values accepted
// by the YGOProDeck API.
//
// Using these constants instead of bare strings makes queries self-documenting
// and protects against typos:
//
//	q := query.New().
//	    Attribute(enums.AttributeDark).
//	    Format(enums.FormatTCG).
//	    BanlistFilter(enums.BanlistTCG).
//	    Sort(enums.SortATK)
package enums
