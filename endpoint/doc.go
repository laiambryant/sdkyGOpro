// Package endpoint provides the generic HTTP endpoint used internally by the
// YGOProDeck SDK.
//
// [Endpoint] is a thin wrapper around [client.Client] that fetches a path and
// unmarshals the JSON response body into a typed value. Each SDK method
// constructs the path (including any query string) and delegates to
// [Endpoint.Fetch].
//
// Callers working directly with the SDK should use the methods on
// [ygoprodeck.YGOProDeck] rather than constructing Endpoint values themselves.
package endpoint
