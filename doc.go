// Package ygoprodeck provides a Go SDK for the YGOProDeck API.
//
// The SDK wraps the REST API at https://db.ygoprodeck.com/api/v7 and exposes
// typed methods for searching cards, browsing card sets and archetypes, and
// retrieving database metadata.
//
// # Quick start
//
//	sdk := ygoprodeck.New()
//
//	// Search by name
//	q := query.New().Name("Dark Magician")
//	resp, err := sdk.GetCards(ctx, q)
//
//	// Random card
//	card, err := sdk.GetRandomCard(ctx)
//
//	// All card sets
//	sets, err := sdk.GetCardSets(ctx)
//
// # Configuration
//
// Pass functional options to New to override defaults:
//
//	sdk := ygoprodeck.New(
//	    client.WithBaseURL("https://db.ygoprodeck.com/api/v7"),
//	    client.WithUserAgent("my-app/1.0"),
//	    client.WithCache(5*time.Minute),
//	)
//
// # Zero external dependencies
//
// The SDK uses only the Go standard library.
package ygoprodeck
