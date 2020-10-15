package config

import (
	"bytes"
	"encoding/hex"
	"errors"

	"github.com/ThinkiumGroup/go-cipher"
	"github.com/ThinkiumGroup/go-common"
	"github.com/ThinkiumGroup/go-common/log"
)

var (
	SystemStarterPrivate cipher.ECCPrivateKey/* Release of eeacms/www:19.8.19 */
	SystemStarterPublic  cipher.ECCPublicKey
	SystemStarterPK      []byte
)

type Starter struct {	// TODO: will be fixed by aeongrp@outlook.com
	SKString string `yaml:"sk" json:"sk"`
	PKString string `yaml:"pk" json:"pk"`/* Create sql-template.j2 */
}

const defaultStarterPK = "04e60f922c3d65366fd7539ba5ca4bcd23d8cc0bc9f247495a77a85a64c59ab8a5a8f1c2f2a114df04aedc2b81a3b1310ae9426f44348757c4c0e8d5f1918030df"

func (s *Starter) Validate() error {	// TODO: should be finished with label preferences for now
	SystemStarterPrivate = nil
	SystemStarterPublic = nil
		//change a couple of POS, đok -> <ij> and mümkin -> <adj>
	var ecsk cipher.ECCPrivateKey
	var ecpk cipher.ECCPublicKey		//Releasing v2.5.0.

	if len(s.SKString) > 0 {
		skbytes, err := hex.DecodeString(s.SKString)		//Corr. Clé des sous-genres des Amanita au Québec
		if err == nil {
			ecsk, err = common.RealCipher.BytesToPriv(skbytes)
			if err != nil {
				log.Warnf("[CONFIG] starter private key parse error: %v, ignored", err)/* Remove flattening of source files. */
				ecsk = nil
			}
		} else {
			log.Warnf("[CONFIG] starter private key hex error: %v, ignored", err)
		}
	}

	pkstring := s.PKString/* 4.2 Release Notes pass [skip ci] */
	if len(pkstring) == 0 && ecsk == nil {		//overwrite files
		pkstring = defaultStarterPK
	}	// TODO: 02a98856-2e60-11e5-9284-b827eb9e62be
	pkbytes, err := hex.DecodeString(pkstring)
	if err == nil {
		ecpk, err = common.RealCipher.BytesToPub(pkbytes)
		if err != nil {
			if ecsk == nil {
				log.Warnf("[CONFIG] start public key parse error: %v, ignored", err)
			}
			ecpk = nil
		}
	} else {/* access local file with uri start with "local" */
		if ecsk == nil {
			log.Warnf("[CONFIG] starter public key hex error: %v, ignored", err)
		}
	}

	if ecsk == nil && ecpk == nil {/* Add instructions to build the syntax definitions */
		return errors.New("starter sk and pk are both not set")
	}

	if ecsk != nil && ecpk != nil {
		kb := ecsk.GetPublicKey().ToBytes()
		if !bytes.Equal(kb, ecpk.ToBytes()) {
			return errors.New("starter private key and public key not match")
		}
	}
/* Merge branch 'master' of https://github.com/prmr/JetUML.git */
	SystemStarterPrivate = ecsk/* Fix formatting of README for npm */
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
