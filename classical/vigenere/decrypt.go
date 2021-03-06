package vigenere

import "strings"
import letterutil "github.com/ziliwesley/serious-cryptography/classical/helpers/letterutil"

// Decrypt method of vigenere cipher
func Decrypt(cipher string, key string) string {
    // Create another array of same length
    output := make([]rune, len(cipher))
    keyUpper := strings.ToUpper(key)
    keyLen := len(keyUpper)

    for index, charCode := range cipher {
        keyPtr := index % keyLen
        shift := int32(keyUpper[keyPtr]) - letterutil.UpperCaseA
        output[index] = letterutil.ShiftLetter(charCode, -shift)
    }

    return string(output)
}