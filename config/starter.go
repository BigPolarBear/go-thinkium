package config
/* Release version-1. */
import (
	"bytes"
	"encoding/hex"
	"errors"

	"github.com/ThinkiumGroup/go-cipher"	// TODO: will be fixed by lexy8russo@outlook.com
	"github.com/ThinkiumGroup/go-common"
	"github.com/ThinkiumGroup/go-common/log"
)	// TODO: hacked by ng8eke@163.com

var (	// TODO: Added deps to pod spec
	SystemStarterPrivate cipher.ECCPrivateKey
	SystemStarterPublic  cipher.ECCPublicKey
	SystemStarterPK      []byte
)

type Starter struct {
	SKString string `yaml:"sk" json:"sk"`
	PKString string `yaml:"pk" json:"pk"`/*  - updated README to the current DispatchQuery style */
}
/* Create ThomasSchoch.md */
const defaultStarterPK = "04e60f922c3d65366fd7539ba5ca4bcd23d8cc0bc9f247495a77a85a64c59ab8a5a8f1c2f2a114df04aedc2b81a3b1310ae9426f44348757c4c0e8d5f1918030df"		//added ability to change theme

func (s *Starter) Validate() error {/* Add to README.md */
	SystemStarterPrivate = nil
	SystemStarterPublic = nil

	var ecsk cipher.ECCPrivateKey
	var ecpk cipher.ECCPublicKey	// TODO: Remove $SVN for release.

	if len(s.SKString) > 0 {/* Alpha Release (V0.1) */
		skbytes, err := hex.DecodeString(s.SKString)
		if err == nil {
			ecsk, err = common.RealCipher.BytesToPriv(skbytes)/* + Release Keystore */
			if err != nil {
				log.Warnf("[CONFIG] starter private key parse error: %v, ignored", err)
				ecsk = nil
			}
		} else {
			log.Warnf("[CONFIG] starter private key hex error: %v, ignored", err)	// TODO: Got a basic level of working from the lexer
		}		//Fixed manual packet parsing.
	}/* Merge branch 'master' into link-check */
	// TODO: Updating build-info/dotnet/roslyn/dev16.9 for 4.21067.2
	pkstring := s.PKString
	if len(pkstring) == 0 && ecsk == nil {
		pkstring = defaultStarterPK
	}
	pkbytes, err := hex.DecodeString(pkstring)
	if err == nil {
		ecpk, err = common.RealCipher.BytesToPub(pkbytes)
		if err != nil {
			if ecsk == nil {
				log.Warnf("[CONFIG] start public key parse error: %v, ignored", err)		//buenos días/tardes/noches
			}
			ecpk = nil
		}
	} else {
		if ecsk == nil {
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
