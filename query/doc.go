// Package query provides a fluent builder for constructing query strings
// accepted by the /cardinfo.php endpoint of the YGOProDeck API.
//
// Create a new [Query] with [New], chain the desired filter methods, then pass
// it to [ygoprodeck.YGOProDeck.GetCards]:
//
//	q := query.New().
//	    FuzzyName("Dragon").
//	    Attribute(enums.AttributeLight).
//	    Sort(enums.SortATK).
//	    Num(20).
//	    Offset(0)
//
//	resp, err := sdk.GetCards(ctx, q)
//
// Passing nil instead of a *Query returns all cards without filters.
//
// [Build] encodes all accumulated parameters as a URL query string (e.g.
// "?name=Dark+Magician&level=7"). Values are percent-encoded via
// [net/url.QueryEscape].
package query
