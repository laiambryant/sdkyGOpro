package ygoprodeck

import (
	"context"
	"net/url"

	"github.com/laiambryant/sdkygopro/client"
	"github.com/laiambryant/sdkygopro/endpoint"
	"github.com/laiambryant/sdkygopro/models"
	"github.com/laiambryant/sdkygopro/query"
)

// YGOProDeck is the top-level SDK client. Create one with [New] and call its
// methods to interact with the YGOProDeck API.
type YGOProDeck struct {
	// Client is the underlying HTTP client. It can be used to customise
	// transport behaviour directly, but in most cases the functional options
	// passed to [New] are sufficient.
	Client  *client.Client
	cards   *endpoint.Endpoint[models.CardResponse]
	sets    *endpoint.Endpoint[[]models.CardSet]
	setInfo *endpoint.Endpoint[[]models.CardSetInfo]
	archs   *endpoint.Endpoint[[]models.Archetype]
	dbVer   *endpoint.Endpoint[[]models.DBVersion]
}

// New creates a YGOProDeck SDK instance. The default base URL is
// https://db.ygoprodeck.com/api/v7. Pass [client.Option] values to override
// the base URL, user-agent, HTTP client, or enable response caching.
func New(opts ...client.Option) *YGOProDeck {
	c := client.NewHTTPClient(nil, opts...)
	return &YGOProDeck{
		Client:  c,
		cards:   endpoint.New[models.CardResponse](c),
		sets:    endpoint.New[[]models.CardSet](c),
		setInfo: endpoint.New[[]models.CardSetInfo](c),
		archs:   endpoint.New[[]models.Archetype](c),
		dbVer:   endpoint.New[[]models.DBVersion](c),
	}
}

// GetCards queries /cardinfo.php with the given query parameters and returns
// the matching cards together with optional pagination metadata. Pass nil to
// return all cards without filters.
func (y *YGOProDeck) GetCards(ctx context.Context, q *query.Query) (models.CardInfoResponse, error) {
	qs := ""
	if q != nil {
		qs = q.Build()
	}
	resp, err := y.cards.Fetch(ctx, "/cardinfo.php"+qs)
	if err != nil {
		return models.CardInfoResponse{}, err
	}
	return models.CardInfoResponse{Cards: resp.Data, Meta: resp.Meta}, nil
}

// GetRandomCard fetches a single random card from /randomcard.php.
func (y *YGOProDeck) GetRandomCard(ctx context.Context) (models.Card, error) {
	resp, err := y.cards.Fetch(ctx, "/randomcard.php")
	if err != nil {
		return models.Card{}, err
	}
	if len(resp.Data) == 0 {
		return models.Card{}, client.ErrNotFound
	}
	return resp.Data[0], nil
}

// GetCardSets returns all card sets from /cardsets.php.
func (y *YGOProDeck) GetCardSets(ctx context.Context) ([]models.CardSet, error) {
	return y.sets.Fetch(ctx, "/cardsets.php")
}

// GetCardSetInfo returns card membership details for the given set code from
// /cardsetsinfo.php. The set code is URL-encoded automatically.
func (y *YGOProDeck) GetCardSetInfo(ctx context.Context, setcode string) ([]models.CardSetInfo, error) {
	return y.setInfo.Fetch(ctx, "/cardsetsinfo.php?setcode="+url.QueryEscape(setcode))
}

// GetArchetypes returns all card archetypes from /archetypes.php.
func (y *YGOProDeck) GetArchetypes(ctx context.Context) ([]models.Archetype, error) {
	return y.archs.Fetch(ctx, "/archetypes.php")
}

// GetDBVersion returns the current database version from /checkDBVer.php.
func (y *YGOProDeck) GetDBVersion(ctx context.Context) (models.DBVersion, error) {
	versions, err := y.dbVer.Fetch(ctx, "/checkDBVer.php")
	if err != nil {
		return models.DBVersion{}, err
	}
	if len(versions) == 0 {
		return models.DBVersion{}, client.ErrNotFound
	}
	return versions[0], nil
}
