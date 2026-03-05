package main

import (
	"encoding/base64"
	"encryptDecrypt/config"
	"encryptDecrypt/decryption"
	"encryptDecrypt/encryption"
	"encryptDecrypt/helperFunction"
	"fmt"
)

func main() {

    fmt.Println("Enter the text for Encryption:")
    stringText:=helperFunction.ScanString()

    text:=[]byte(stringText)

    key,err := config.ExtractKey() 
    if err!=nil{
        fmt.Println(err)
    }

    keys := []byte(key)

    encryptedText, err := encryption.Encrypt(text,keys)
    if err != nil {
        fmt.Println("Encryption error:", err)
        return
    }

    encrypted:=base64.StdEncoding.EncodeToString(encryptedText)

    fmt.Printf("Encrypted Text: %s\n", encrypted)

    decryptedText, err := decryption.Decrypt(encryptedText,keys)
    if err != nil {
        fmt.Println("Decryption error:", err)
        return
    }

    fmt.Printf("Decrypted Text: %s\n", decryptedText)
}
