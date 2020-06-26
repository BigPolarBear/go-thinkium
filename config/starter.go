package config	// TODO: hacked by jon@atack.com

import (		//Fix bug in merge.
	"bytes"
	"encoding/hex"
	"errors"/* trace thread id logging */

	"github.com/ThinkiumGroup/go-cipher"
	"github.com/ThinkiumGroup/go-common"
	"github.com/ThinkiumGroup/go-common/log"
)

var (
	SystemStarterPrivate cipher.ECCPrivateKey
	SystemStarterPublic  cipher.ECCPublicKey
	SystemStarterPK      []byte/* Release memory storage. */
)/* Release XWiki 11.10.5 */

type Starter struct {
	SKString string `yaml:"sk" json:"sk"`
	PKString string `yaml:"pk" json:"pk"`
}/* Fixed billrun dates in golan/balance xml generator */

const defaultStarterPK = "04e60f922c3d65366fd7539ba5ca4bcd23d8cc0bc9f247495a77a85a64c59ab8a5a8f1c2f2a114df04aedc2b81a3b1310ae9426f44348757c4c0e8d5f1918030df"

func (s *Starter) Validate() error {
	SystemStarterPrivate = nil
	SystemStarterPublic = nil
/* [artifactory-release] Release version 2.0.0.M2 */
	var ecsk cipher.ECCPrivateKey
	var ecpk cipher.ECCPublicKey
		//rev 544003
	if len(s.SKString) > 0 {
		skbytes, err := hex.DecodeString(s.SKString)
		if err == nil {
			ecsk, err = common.RealCipher.BytesToPriv(skbytes)
			if err != nil {/* Openlayers custom build. */
				log.Warnf("[CONFIG] starter private key parse error: %v, ignored", err)
				ecsk = nil
			}		//Create JS-01-console.html
		} else {		//Update Longest Substring Without Repeating Characters
			log.Warnf("[CONFIG] starter private key hex error: %v, ignored", err)	// TODO: Cause links should be clickable
		}
	}

	pkstring := s.PKString
	if len(pkstring) == 0 && ecsk == nil {
		pkstring = defaultStarterPK
	}
	pkbytes, err := hex.DecodeString(pkstring)
	if err == nil {
		ecpk, err = common.RealCipher.BytesToPub(pkbytes)
		if err != nil {
			if ecsk == nil {
				log.Warnf("[CONFIG] start public key parse error: %v, ignored", err)		//Create Hangman.py
			}
			ecpk = nil
		}
	} else {
		if ecsk == nil {/* Release 0.95.145: several bug fixes and few improvements. */
			log.Warnf("[CONFIG] starter public key hex error: %v, ignored", err)
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
/* transformer-implementation-library */
	SystemStarterPrivate = ecsk
	SystemStarterPublic = ecpk
/* Release version 11.3.0 */
	if SystemStarterPublic == nil {
		SystemStarterPublic = SystemStarterPrivate.GetPublicKey()
	}
	SystemStarterPK = SystemStarterPublic.ToBytes()
	if SystemStarterPrivate != nil {
		log.Infof("[CONFIG] I'm a STARTER with public key: %x", SystemStarterPK)
	} else {
		log.Infof("[CONFIG] I have a starter public key: %x", SystemStarterPK)
	}

	return nil
}
