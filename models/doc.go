// Package models contains the data types returned by the YGOProDeck API.
//
// Each type corresponds directly to the JSON structure documented at
// https://ygoprodeck.com/api-guide/. Optional fields that may be absent in the
// API response are represented as pointer types so that callers can distinguish
// between a zero value and an absent value.
//
// Top-level response types:
//
//   - [Card] — a single card from /cardinfo.php or /randomcard.php
//   - [CardSet] — a card set entry from /cardsets.php
//   - [CardSetInfo] — per-card set membership from /cardsetsinfo.php
//   - [Archetype] — an archetype name from /archetypes.php
//   - [DBVersion] — database version from /checkDBVer.php
//   - [CardInfoResponse] — cards plus optional pagination [Meta]
package models
