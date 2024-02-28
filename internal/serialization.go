package internal

import (
	"encoding/base64"
	"fmt"

	envkmspb "github.com/lstoll/envkms/proto"
	"github.com/tink-crypto/tink-go/v2/keyset"
	tinkpb "github.com/tink-crypto/tink-go/v2/proto/tink_go_proto"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

var ksEncoding = base64.StdEncoding

// UnmarshalKeyset unmarshals the serialized keyset
func UnmarshalKeyset(serialized string) (*envkmspb.Keyset, error) {
	b, err := ksEncoding.DecodeString(serialized)
	if err != nil {
		return nil, fmt.Errorf("decoding serialized value: %v", err)
	}

	ks := envkmspb.Keyset{}
	if err := proto.Unmarshal(b, &ks); err != nil {
		return nil, fmt.Errorf("unmarshaling keyset: %w", err)
	}

	return &ks, nil
}

// MarshalKeyset seralizes the given keyset
func MarshalKeyset(ks *envkmspb.Keyset) (string, error) {
	b, err := proto.Marshal(ks)
	if err != nil {
		return "", fmt.Errorf("marshaling keyset: %w", err)
	}
	return ksEncoding.EncodeToString(b), nil
}

var _ keyset.Reader = (*KeysetReader)(nil)

type KeysetReader struct {
	From *envkmspb.Keyset
}

func (k *KeysetReader) Read() (*tinkpb.Keyset, error) {
	tks := tinkpb.Keyset{}
	if err := k.From.TinkKeyset.UnmarshalTo(&tks); err != nil {
		return nil, fmt.Errorf("unmarshaling tink keyset: %w", err)
	}
	return &tks, nil
}

func (k *KeysetReader) ReadEncrypted() (*tinkpb.EncryptedKeyset, error) {
	panic("not supported")
}

var _ keyset.Writer = (*KeysetWriter)(nil)

type KeysetWriter struct {
	Into *envkmspb.Keyset
}

func (k *KeysetWriter) Write(keyset *tinkpb.Keyset) error {
	a, err := anypb.New(keyset)
	if err != nil {
		return fmt.Errorf("marshaling keyset: %w", err)
	}
	k.Into.TinkKeyset = a
	return nil
}

func (k *KeysetWriter) WriteEncrypted(keyset *tinkpb.EncryptedKeyset) error {
	panic("not supported")
}
