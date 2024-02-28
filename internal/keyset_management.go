package internal

import (
	"fmt"
	"log"

	envkmspb "github.com/lstoll/envkms/proto"
	"github.com/tink-crypto/tink-go/v2/aead"
	"github.com/tink-crypto/tink-go/v2/insecurecleartextkeyset"
	"github.com/tink-crypto/tink-go/v2/keyset"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func GenerateKeyset() (string, error) {
	handle, err := keyset.NewHandle(aead.AES256GCMKeyTemplate())
	if err != nil {
		log.Fatal(err)
	}

	ks := envkmspb.Keyset{
		KeyCreationDates: map[uint32]*timestamppb.Timestamp{
			handle.KeysetInfo().PrimaryKeyId: timestamppb.Now(),
		},
	}

	if err := insecurecleartextkeyset.Write(handle, &KeysetWriter{Into: &ks}); err != nil {
		return "", fmt.Errorf("writing keyset: %w", err)
	}

	s, err := MarshalKeyset(&ks)
	if err != nil {
		return "", fmt.Errorf("marshaling keyset: %w", err)
	}

	return s, err
}
