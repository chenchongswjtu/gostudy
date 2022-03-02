package main

import (
	"crypto/sha256"
	"fmt"

	"github.com/tjfoc/gmsm/sm2"
)

func handleErr(e error, why string) {
	if e != nil {
		fmt.Println(why, e)
	}
}

func main() {
	privateKey, e := sm2.GenerateKey()
	handleErr(e, "get priv key error ")

	publicKey := privateKey.PublicKey
	hash := sha256.Sum256([]byte("wek"))
	r, s, err := sm2.Sign(privateKey, hash[:])
	handleErr(err, "sign")
	verify := sm2.Verify(&publicKey, hash[:], r, s)
	fmt.Println(verify)

}
