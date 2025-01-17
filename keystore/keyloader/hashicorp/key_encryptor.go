package hashicorp

import (
	"flag"

	"github.com/cossacklabs/acra/keystore"
	baseKMS "github.com/cossacklabs/acra/keystore/kms/base"
	keystoreV2 "github.com/cossacklabs/acra/keystore/v2/keystore"
	"github.com/cossacklabs/acra/keystore/v2/keystore/crypto"
	log "github.com/sirupsen/logrus"
)

// KeyEncryptorFabric implementation of keyloader.KeyEncryptorFabric for `vault_master_key` strategy
type KeyEncryptorFabric struct{}

// NewKeyEncryptor fabric of keystore.KeyEncryptor for for `vault_master_key` strategy
func (k KeyEncryptorFabric) NewKeyEncryptor(flags *flag.FlagSet, prefix string) (keystore.KeyEncryptor, error) {
	loader, err := NewMasterKeyLoader(flags, prefix)
	if err != nil {
		return nil, err
	}

	key, err := loader.LoadMasterKey()
	if err != nil {
		return nil, err
	}
	return keystore.NewSCellKeyEncryptor(key)
}

// NewKeyEncryptorSuite fabric of crypto.KeyStoreSuite for `vault_master_key` strategy
func (k KeyEncryptorFabric) NewKeyEncryptorSuite(flags *flag.FlagSet, prefix string) (*crypto.KeyStoreSuite, error) {
	loader, err := NewMasterKeyLoader(flags, prefix)
	if err != nil {
		return nil, err
	}

	encryption, signature, err := loader.LoadMasterKeys()
	if err != nil {
		log.WithError(err).Errorln("Cannot load master key")
		return nil, err
	}
	return keystoreV2.NewSCellSuite(encryption, signature)
}

// RegisterCLIParameters empty implementation of KeyEncryptorFabric interface
func (k KeyEncryptorFabric) RegisterCLIParameters(flags *flag.FlagSet, prefix, description string) {
	RegisterCLIParametersWithFlagSet(flags, prefix, description)
}

// GetKeyMapper return KeyMapper for `vault_master_key` strategy
func (k KeyEncryptorFabric) GetKeyMapper() baseKMS.KeyMapper {
	panic("No KeyMapper for vault_master_key strategy")
}
