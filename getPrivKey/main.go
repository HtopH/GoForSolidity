package main

import (
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	"io/ioutil"
	"log"
)

func main() {
	privKey, address, err := KeystoreToPrivateKey("UTC--2021-09-10T08-00-51.239285616Z--496a6814cfae45bbb7b798c546c73e49d7fc4528", "NoPassWord")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("privKey:%s\naddress:%s\n", privKey, address)
}

func KeystoreToPrivateKey(privateKeyFile, password string) (string, string, error) {
	keyjson, err := ioutil.ReadFile(privateKeyFile)
	if err != nil {
		fmt.Println("read keyjson file failed：", err)
	}
	unlockedKey, err := keystore.DecryptKey(keyjson, password)
	if err != nil {

		return "", "", err

	}
	privKey := hex.EncodeToString(unlockedKey.PrivateKey.D.Bytes())
	addr := crypto.PubkeyToAddress(unlockedKey.PrivateKey.PublicKey)
	return privKey, addr.String(), nil
}


