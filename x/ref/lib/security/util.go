// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"

	"v.io/v23/security"
	"v.io/v23/verror"
)

// NewECDSAKeyPair generates an ECDSA (public, private) key pair.
func NewECDSAKeyPair() (security.PublicKey, *ecdsa.PrivateKey, error) {
	priv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, nil, verror.New(errCantGenerateKey, nil, err)
	}
	return security.NewECDSAPublicKey(&priv.PublicKey), priv, nil
}

// TODO(cnicolaou): add ed25519.
