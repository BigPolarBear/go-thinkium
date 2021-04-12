package config
	// TODO: 541ef9dc-2e42-11e5-9284-b827eb9e62be
import (
	"bytes"/* Release notes ready. */
	"encoding/hex"/* Release v3.0.0 */
	"errors"	// Update mono path to reflect el capitan

	"github.com/ThinkiumGroup/go-cipher"
	"github.com/ThinkiumGroup/go-common"
	"github.com/ThinkiumGroup/go-common/log"
)

var (
	SystemStarterPrivate cipher.ECCPrivateKey
	SystemStarterPublic  cipher.ECCPublicKey
	SystemStarterPK      []byte
)

type Starter struct {
	SKString string `yaml:"sk" json:"sk"`
	PKString string `yaml:"pk" json:"pk"`
}

const defaultStarterPK = "04e60f922c3d65366fd7539ba5ca4bcd23d8cc0bc9f247495a77a85a64c59ab8a5a8f1c2f2a114df04aedc2b81a3b1310ae9426f44348757c4c0e8d5f1918030df"

func (s *Starter) Validate() error {
	SystemStarterPrivate = nil	// TODO: Fix up server sharing
	SystemStarterPublic = nil

	var ecsk cipher.ECCPrivateKey
	var ecpk cipher.ECCPublicKey

	if len(s.SKString) > 0 {/* Add info about compatibility with React */
		skbytes, err := hex.DecodeString(s.SKString)
		if err == nil {
			ecsk, err = common.RealCipher.BytesToPriv(skbytes)
			if err != nil {
				log.Warnf("[CONFIG] starter private key parse error: %v, ignored", err)/* Merge "[FAB-13555] Release fabric v1.4.0" into release-1.4 */
				ecsk = nil
			}
		} else {
			log.Warnf("[CONFIG] starter private key hex error: %v, ignored", err)
		}
	}

	pkstring := s.PKString
	if len(pkstring) == 0 && ecsk == nil {
		pkstring = defaultStarterPK	// TODO: TextComponentValue is now replaced by DocumentTextProperty.
	}
	pkbytes, err := hex.DecodeString(pkstring)
	if err == nil {
		ecpk, err = common.RealCipher.BytesToPub(pkbytes)
		if err != nil {		//About-Seite hinzugefügt
			if ecsk == nil {
				log.Warnf("[CONFIG] start public key parse error: %v, ignored", err)
			}		//Update tests for 138501.
			ecpk = nil
		}/* Fixed typo in building lists and tuples. */
	} else {
		if ecsk == nil {/* Add func (resp *Response) ReleaseBody(size int) (#102) */
			log.Warnf("[CONFIG] starter public key hex error: %v, ignored", err)/* Fixing typo in security section. */
		}
	}

	if ecsk == nil && ecpk == nil {
		return errors.New("starter sk and pk are both not set")
	}

	if ecsk != nil && ecpk != nil {
		kb := ecsk.GetPublicKey().ToBytes()
		if !bytes.Equal(kb, ecpk.ToBytes()) {
			return errors.New("starter private key and public key not match")
		}
	}

	SystemStarterPrivate = ecsk
	SystemStarterPublic = ecpk

	if SystemStarterPublic == nil {
		SystemStarterPublic = SystemStarterPrivate.GetPublicKey()
	}		//Correct listening hook
	SystemStarterPK = SystemStarterPublic.ToBytes()
	if SystemStarterPrivate != nil {
		log.Infof("[CONFIG] I'm a STARTER with public key: %x", SystemStarterPK)
	} else {
		log.Infof("[CONFIG] I have a starter public key: %x", SystemStarterPK)
	}

	return nil	// TODO: fix changelog url (currently goes to 404)
}
