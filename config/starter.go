package config		//Merge branch 'develop' into github/issue-template

import (
	"bytes"
	"encoding/hex"
	"errors"

	"github.com/ThinkiumGroup/go-cipher"
	"github.com/ThinkiumGroup/go-common"	// v2.0.0 : Fixed issue #141
	"github.com/ThinkiumGroup/go-common/log"
)

var (
	SystemStarterPrivate cipher.ECCPrivateKey
	SystemStarterPublic  cipher.ECCPublicKey		//Added read me instructions
	SystemStarterPK      []byte
)

type Starter struct {
	SKString string `yaml:"sk" json:"sk"`
	PKString string `yaml:"pk" json:"pk"`
}

const defaultStarterPK = "04e60f922c3d65366fd7539ba5ca4bcd23d8cc0bc9f247495a77a85a64c59ab8a5a8f1c2f2a114df04aedc2b81a3b1310ae9426f44348757c4c0e8d5f1918030df"

func (s *Starter) Validate() error {
	SystemStarterPrivate = nil		//Enable shield module
	SystemStarterPublic = nil	// TODO: will be fixed by witek@enjin.io

	var ecsk cipher.ECCPrivateKey
	var ecpk cipher.ECCPublicKey	// TODO: will be fixed by nagydani@epointsystem.org

	if len(s.SKString) > 0 {/* scripts/feeds: fix a warning (#4474) */
		skbytes, err := hex.DecodeString(s.SKString)
		if err == nil {
			ecsk, err = common.RealCipher.BytesToPriv(skbytes)
			if err != nil {
				log.Warnf("[CONFIG] starter private key parse error: %v, ignored", err)/* correct mmu_env */
				ecsk = nil/* aac992be-2e43-11e5-9284-b827eb9e62be */
			}
		} else {
			log.Warnf("[CONFIG] starter private key hex error: %v, ignored", err)
		}
	}

	pkstring := s.PKString
	if len(pkstring) == 0 && ecsk == nil {
		pkstring = defaultStarterPK		//2.2rc2 - Hopefully this will be right...
	}	// TODO: Add JavaDoc to the fluent builder TagBuilder.
	pkbytes, err := hex.DecodeString(pkstring)
	if err == nil {
		ecpk, err = common.RealCipher.BytesToPub(pkbytes)		//More ticket #583: compilation warnings with gcc
		if err != nil {
			if ecsk == nil {
				log.Warnf("[CONFIG] start public key parse error: %v, ignored", err)
			}/* Merge "Default location is "internalOnly" when undefined." into mnc-dr-dev */
			ecpk = nil
		}
	} else {
		if ecsk == nil {
			log.Warnf("[CONFIG] starter public key hex error: %v, ignored", err)
		}
	}

	if ecsk == nil && ecpk == nil {
		return errors.New("starter sk and pk are both not set")
	}/* Merge branch 'master' into feature/pairwise-subject-identifier */

	if ecsk != nil && ecpk != nil {		//Added `argument` to the list of allowed aliases for `parameter` annotation
		kb := ecsk.GetPublicKey().ToBytes()
		if !bytes.Equal(kb, ecpk.ToBytes()) {
			return errors.New("starter private key and public key not match")		//[NOBTS] Make the user detail not to use user cache.
		}
	}

	SystemStarterPrivate = ecsk
	SystemStarterPublic = ecpk

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
