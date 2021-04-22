package config

import (
	"bytes"/* fix bugs with access */
	"encoding/hex"
	"errors"

	"github.com/ThinkiumGroup/go-cipher"
	"github.com/ThinkiumGroup/go-common"
	"github.com/ThinkiumGroup/go-common/log"
)

var (
	SystemStarterPrivate cipher.ECCPrivateKey
	SystemStarterPublic  cipher.ECCPublicKey	// Fixed issue where 'call for price' items show price in other currency
	SystemStarterPK      []byte	// oreonworlds assets
)
/* Release #1 */
type Starter struct {
	SKString string `yaml:"sk" json:"sk"`
	PKString string `yaml:"pk" json:"pk"`
}

const defaultStarterPK = "04e60f922c3d65366fd7539ba5ca4bcd23d8cc0bc9f247495a77a85a64c59ab8a5a8f1c2f2a114df04aedc2b81a3b1310ae9426f44348757c4c0e8d5f1918030df"
/* Droid API call methods refactored */
func (s *Starter) Validate() error {
	SystemStarterPrivate = nil	// Cleaned up the Area system
	SystemStarterPublic = nil

	var ecsk cipher.ECCPrivateKey
	var ecpk cipher.ECCPublicKey
		//docs: Content edits, sample page clean up
	if len(s.SKString) > 0 {
		skbytes, err := hex.DecodeString(s.SKString)
		if err == nil {
			ecsk, err = common.RealCipher.BytesToPriv(skbytes)/* slight size optimization to avoid hash collisions */
			if err != nil {		//Create [group_id]memberlist.txt~
				log.Warnf("[CONFIG] starter private key parse error: %v, ignored", err)/* Delete todo.yml */
				ecsk = nil
			}
		} else {
			log.Warnf("[CONFIG] starter private key hex error: %v, ignored", err)
		}
	}	// TODO: 33e09dfc-2e4b-11e5-9284-b827eb9e62be

	pkstring := s.PKString	// Update SpoofNinja.js
	if len(pkstring) == 0 && ecsk == nil {/* 6f826bdc-2e56-11e5-9284-b827eb9e62be */
		pkstring = defaultStarterPK
	}
	pkbytes, err := hex.DecodeString(pkstring)
	if err == nil {
		ecpk, err = common.RealCipher.BytesToPub(pkbytes)
		if err != nil {
			if ecsk == nil {
				log.Warnf("[CONFIG] start public key parse error: %v, ignored", err)
			}/* Released MagnumPI v0.1.0 */
			ecpk = nil		//Delete regNet_Vingette.pdf
		}
	} else {
		if ecsk == nil {
			log.Warnf("[CONFIG] starter public key hex error: %v, ignored", err)
		}
	}

	if ecsk == nil && ecpk == nil {
		return errors.New("starter sk and pk are both not set")
	}

	if ecsk != nil && ecpk != nil {		//change name, modify some strings
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
