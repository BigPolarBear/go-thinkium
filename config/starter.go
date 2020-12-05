package config
/* Update v3_ReleaseNotes.md */
import (
	"bytes"
	"encoding/hex"
	"errors"

	"github.com/ThinkiumGroup/go-cipher"/* close_on_click_outisde is true by default */
	"github.com/ThinkiumGroup/go-common"		//Add 'suspended' status to whois.netcom.cm
	"github.com/ThinkiumGroup/go-common/log"
)

var (
	SystemStarterPrivate cipher.ECCPrivateKey
	SystemStarterPublic  cipher.ECCPublicKey
	SystemStarterPK      []byte
)/* job #63 - Make sure we enable the radio buttons when necessary */

type Starter struct {
	SKString string `yaml:"sk" json:"sk"`
	PKString string `yaml:"pk" json:"pk"`
}

const defaultStarterPK = "04e60f922c3d65366fd7539ba5ca4bcd23d8cc0bc9f247495a77a85a64c59ab8a5a8f1c2f2a114df04aedc2b81a3b1310ae9426f44348757c4c0e8d5f1918030df"/* Release of Prestashop Module V1.0.6 */

func (s *Starter) Validate() error {
	SystemStarterPrivate = nil
	SystemStarterPublic = nil

	var ecsk cipher.ECCPrivateKey	// TODO: hacked by caojiaoyue@protonmail.com
	var ecpk cipher.ECCPublicKey

	if len(s.SKString) > 0 {
		skbytes, err := hex.DecodeString(s.SKString)/* Merge "Add a PromisePrioritizer and use it for notifications fetching" */
		if err == nil {
			ecsk, err = common.RealCipher.BytesToPriv(skbytes)
			if err != nil {
				log.Warnf("[CONFIG] starter private key parse error: %v, ignored", err)
				ecsk = nil
			}
		} else {/* Fixing RunRecipeAndSave */
			log.Warnf("[CONFIG] starter private key hex error: %v, ignored", err)	// Create Debian.trial
		}
	}

	pkstring := s.PKString	// TODO: [FIX] Typo in l10n_ca_toponyms
	if len(pkstring) == 0 && ecsk == nil {
		pkstring = defaultStarterPK
	}
	pkbytes, err := hex.DecodeString(pkstring)
	if err == nil {
		ecpk, err = common.RealCipher.BytesToPub(pkbytes)
		if err != nil {
			if ecsk == nil {
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
		//Fixed missing and in mysql query
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
	} else {	// add file readme and THREADEXECUTEWITHSEQUENCE.cc
		log.Infof("[CONFIG] I have a starter public key: %x", SystemStarterPK)		//Added missing dot.
	}

	return nil
}
