package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/lloeki/bitcoinj-wallet-dump/base58"
	"github.com/lloeki/bitcoinj-wallet-dump/wallet"

	"github.com/golang/protobuf/proto"
)

func main() {
	w := &wallet.Wallet{}
	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal("loading error: ", err)
	}
	err = proto.Unmarshal(data, w)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}

	//fmt.Printf("%v", w)

	for i, key := range w.Key {
		fmt.Printf("%06d ", i)
		switch *key.Type {
		case wallet.Key_ORIGINAL:
			fmt.Printf("[original]\n")
		case wallet.Key_ENCRYPTED_SCRYPT_AES:
			fmt.Printf("[encrypted]\n")
		case wallet.Key_DETERMINISTIC_MNEMONIC:
			fmt.Printf("[deterministic mnemonic]\n")
			fmt.Printf("   Priv: '%s'\n", key.SecretBytes)
			if len(key.DeterministicSeed) > 0 {
				fmt.Printf("   Seed: %x\n", key.DeterministicSeed)
			}
			if key.EncryptedDeterministicSeed != nil {
				fmt.Printf("   EncSeed: %v\n", key.EncryptedDeterministicSeed)
			}
			continue
		case wallet.Key_DETERMINISTIC_KEY:
			fmt.Printf("[deterministic]\n")
		}
		if len(key.PublicKey) > 0 {
			fmt.Printf("   Publ: %s\n", base58.Check(base58.PublicHash(key.PublicKey), base58.BITCOIN_PUBKEY_HASH, true))
		}
		if len(key.SecretBytes) > 0 {
			fmt.Printf("   Priv: %s\n", base58.Check(key.SecretBytes, base58.BITCOIN_PRIVKEY, false))
			fmt.Printf("   Priv: %s\n", base58.Check(key.SecretBytes, base58.BITCOIN_PRIVKEY, true))
		}
		if key.EncryptedData != nil {
			fmt.Printf("   Enc: %v\n", key.EncryptedData)
		}
		if key.DeterministicKey != nil {
			fmt.Printf("       DetK %d chaincode: %x\n", i, key.DeterministicKey.ChainCode)
			for j, path := range key.DeterministicKey.Path {
				if path>>31 == 1 {
					// private key path
					fmt.Printf("       DetK %d path %d: [priv] %d\n", i, j, path&(2<<30-1))
				} else {
					// public key path
					fmt.Printf("       DetK %d path %d: [publ] %d\n", i, j, path)
				}
			}
		}

	}
}
