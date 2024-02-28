package envkms

import (
	"fmt"
	"strings"

	"github.com/lstoll/envkms/internal"
	"github.com/tink-crypto/tink-go/v2/aead"
	"github.com/tink-crypto/tink-go/v2/core/registry"
	"github.com/tink-crypto/tink-go/v2/insecurecleartextkeyset"
	"github.com/tink-crypto/tink-go/v2/keyset"
	"github.com/tink-crypto/tink-go/v2/tink"
)

const envkmsPrefix = "envkms://"

var _ registry.KMSClient = (*KMS)(nil)

type KMS struct{}

func New() *KMS {
	return &KMS{}
}

// Supported true if this client does support keyURI
func (k *KMS) Supported(keyURI string) bool {
	return strings.HasPrefix(keyURI, envkmsPrefix)
}

// GetAEAD  gets an AEAD backend by keyURI.
func (k *KMS) GetAEAD(keyURI string) (tink.AEAD, error) {
	if !strings.HasPrefix(keyURI, envkmsPrefix) {
		return nil, fmt.Errorf("%s does not start with %s", keyURI, envkmsPrefix)
	}
	if len(internal.EnvContents) == 0 {
		return nil, fmt.Errorf("no key set in %s", internal.EnvVar)
	}

	// do a test read, to validate things work.
	_, err := getHandle()
	if err != nil {
		return nil, fmt.Errorf("handle not gettable: %w", err)
	}

	return &envAEAD{}, nil
}

// envAEAD reads and uses handles on the fly, to try and minimize the amount of
// time we have them in unprotected memory.
type envAEAD struct{}

func (*envAEAD) Encrypt(plaintext, associatedData []byte) ([]byte, error) {
	h, err := getHandle()
	if err != nil {
		return nil, err
	}

	p, err := aead.New(h)
	if err != nil {
		return nil, fmt.Errorf("creating aead from keyset: %w", err)
	}

	return p.Encrypt(plaintext, associatedData)
}

func (*envAEAD) Decrypt(ciphertext, associatedData []byte) ([]byte, error) {
	h, err := getHandle()
	if err != nil {
		return nil, err
	}

	p, err := aead.New(h)
	if err != nil {
		return nil, fmt.Errorf("creating aead from keyset: %w", err)
	}

	return p.Decrypt(ciphertext, associatedData)
}

func getHandle() (*keyset.Handle, error) {
	ks, err := internal.UnmarshalKeyset(string(internal.EnvContents))
	if err != nil {
		return nil, err
	}
	parsedHandle, err := insecurecleartextkeyset.Read(&internal.KeysetReader{From: ks})
	if err != nil {
		return nil, fmt.Errorf("reading keyset: %w", err)
	}
	return parsedHandle, nil
}
