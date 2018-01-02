// Copyright © 2018 The Things Network Foundation, distributed under the MIT license (see LICENSE file)

package store

import (
	"github.com/TheThingsNetwork/ttn/pkg/identityserver/types"
	"github.com/TheThingsNetwork/ttn/pkg/ttnpb"
)

// UserFactory is a function that returns a types.User used to
// construct the results in read operations.
type UserFactory func() types.User

// UserStore is a store that holds Users.
type UserStore interface {
	// Create creates an user.
	Create(user types.User) error

	// GetByID finds the user by ID and retrieves it.
	GetByID(userID string, factory UserFactory) (types.User, error)

	// GetByEmail finds the user by email address and retrieves it.
	GetByEmail(email string, factory UserFactory) (types.User, error)

	// Update updates an user.
	Update(user types.User) error

	// TODO(gomezjdaniel#274): use sql 'ON DELETE CASCADE' when CockroachDB implements it.
	// Delete deletes an user.
	Delete(userID string) error

	// SaveValidationToken saves the validation token.
	SaveValidationToken(userID string, token *types.ValidationToken) error

	// GetValidationToken retrieves the validation token.
	GetValidationToken(token string) (string, *types.ValidationToken, error)

	// DeleteValidationToken deletes the validation token.
	DeleteValidationToken(token string) error

	// SaveAPIKey stores an API Key attached to an user.
	SaveAPIKey(userID string, key *ttnpb.APIKey) error

	// GetAPIKey retrieves an API key by value and the user ID.
	GetAPIKey(key string) (string, *ttnpb.APIKey, error)

	// GetAPIKeyByName retrieves an API key from an user.
	GetAPIKeyByName(userID, keyName string) (*ttnpb.APIKey, error)

	// UpdateAPIKeyRights updates the right of an API key.
	UpdateAPIKeyRights(userID, keyName string, rights []ttnpb.Right) error

	// ListAPIKey list all the API keys that an user has.
	ListAPIKeys(userID string) ([]*ttnpb.APIKey, error)

	// DeleteAPIKey deletes a given API key from an user.
	DeleteAPIKey(userID, keyName string) error

	// LoadAttributes loads all user attributes if the User is an Attributer.
	LoadAttributes(userID string, user types.User) error

	// StoreAttributes writes all of the user attributes if the User is an
	// Attributer and returns the written User in result.
	StoreAttributes(userID string, user, result types.User) error
}
