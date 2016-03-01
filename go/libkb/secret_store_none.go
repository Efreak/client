// Copyright 2015 Keybase, Inc. All rights reserved. Use of
// this source code is governed by the included BSD license.

// +build !darwin,!android

package libkb

import (
	"errors"
)

var secretStoreNoneMap = map[NormalizedUsername][]byte{}
var useMockSecrestStore bool

func HasSecretStore() bool {
	return useMockSecrestStore
}

func UseMockSecretStore(useMock bool) {
	useMockSecrestStore = useMock
	if !useMock {
		secretStoreNoneMap = map[NormalizedUsername][]byte{}
	}
}

func GetUsersWithStoredSecrets(c SecretStoreContext) (ret []string, err error) {
	for name, _ := range secretStoreNoneMap {
		ret = append(ret, string(name))
	}
	return
}

func GetTerminalPrompt() string {
	// TODO: Come up with specific prompts for other platforms.
	return "Store your key in the local secret store?"
}

// Used by tests that want to mock out the secret store.
type TestSecretStore struct {
	context     SecretStoreContext
	accountName NormalizedUsername
}

func NewSecretStore(context SecretStoreContext, accountName NormalizedUsername) SecretStore {
	if useMockSecrestStore {
		return TestSecretStore{context, accountName}
	}
	return nil
}

func (tss TestSecretStore) RetrieveSecret() (ret []byte, err error) {

	ret, ok := secretStoreNoneMap[tss.accountName]

	G.Log.Debug("| TestSecretStore::RetrieveSecret(%d)", len(ret))

	if !ok {
		return nil, errors.New("No secret to retrieve")
	}

	return
}

func (tss TestSecretStore) StoreSecret(secret []byte) error {
	G.Log.Debug("| TestSecretStore::StoreSecret(%d)", len(secret))

	secretStoreNoneMap[tss.accountName] = secret
	return nil
}

func (tss TestSecretStore) ClearSecret() error {
	G.Log.Debug("| TestSecretStore::ClearSecret()")

	delete(secretStoreNoneMap, tss.accountName)
	return nil
}
