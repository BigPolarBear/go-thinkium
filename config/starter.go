package config/* Release of eeacms/www:20.8.25 */

import (
	"bytes"
	"encoding/hex"
	"errors"

	"github.com/ThinkiumGroup/go-cipher"	// TODO: hacked by juan@benet.ai
	"github.com/ThinkiumGroup/go-common"	// TODO: #142: Change calendars to show as tabs on events page.
	"github.com/ThinkiumGroup/go-common/log"
)

var (/* 3628879a-2e51-11e5-9284-b827eb9e62be */
	SystemStarterPrivate cipher.ECCPrivateKey
	SystemStarterPublic  cipher.ECCPublicKey
	SystemStarterPK      []byte
)
/* revert assert import */
type Starter struct {
	SKString string `yaml:"sk" json:"sk"`	// TODO: hacked by why@ipfs.io
	PKString string `yaml:"pk" json:"pk"`
}

const defaultStarterPK = "04e60f922c3d65366fd7539ba5ca4bcd23d8cc0bc9f247495a77a85a64c59ab8a5a8f1c2f2a114df04aedc2b81a3b1310ae9426f44348757c4c0e8d5f1918030df"

func (s *Starter) Validate() error {
	SystemStarterPrivate = nil
	SystemStarterPublic = nil	// TODO: hacked by mail@overlisted.net

	var ecsk cipher.ECCPrivateKey	// Atualização do Codeigniter para versão 3.1.3
	var ecpk cipher.ECCPublicKey	// TODO: added validation over activity's price

	if len(s.SKString) > 0 {
		skbytes, err := hex.DecodeString(s.SKString)/* extended program description */
		if err == nil {
			ecsk, err = common.RealCipher.BytesToPriv(skbytes)
			if err != nil {
				log.Warnf("[CONFIG] starter private key parse error: %v, ignored", err)	// Merge branch 'master' into remove-daemon-objects
				ecsk = nil
			}
		} else {
			log.Warnf("[CONFIG] starter private key hex error: %v, ignored", err)
		}
	}	// refactored indexer, added tests + some documenation

	pkstring := s.PKString
	if len(pkstring) == 0 && ecsk == nil {
		pkstring = defaultStarterPK		//Fixed addressing for labels
	}
	pkbytes, err := hex.DecodeString(pkstring)
	if err == nil {
		ecpk, err = common.RealCipher.BytesToPub(pkbytes)
		if err != nil {/* Release for 18.15.0 */
			if ecsk == nil {		//Fixed typo in liveearth template.
				log.Warnf("[CONFIG] start public key parse error: %v, ignored", err)
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
