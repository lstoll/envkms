# envkms

envkms is an implementation of [Tink's](https://developers.google.com/tink) [registry.KMSClient](https://pkg.go.dev/github.com/tink-crypto/tink-go/v2@v2.1.0/core/registry#KMSClient), that uses a pre-allocated keyset from an environment variable to perform KMS operations. It is designed as a substititute for a real KMS in dev, or for homelab use.

:warning: You should probably not use this. Use a Hardware/Cloud Provider KMS instead. :warning:

## Usage

The keyset should be stored in an environment variable called `ENVKMS_KEYSET`. A new keyset can be created with `go run ./cmd/envkmsutil generate-keyset`

See [e2e_test.go](e2e_test.go) for an example of its usage.
