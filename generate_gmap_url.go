// address = 東京都中央区銀座 5−4−3 対鶴館 B2
// apiKey = (API Key)
// apiSecret = (Secret Key)

package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"net/url"
	"os"
)

func signUrl(urlToSign, googleApiKey, googleApiSecret string) (string, error) {
    // Decode the secret key from Base64
    decodedKey, err := base64.StdEncoding.DecodeString(googleApiSecret)
    if err != nil {
        return "", err
    }

    // Sign using HMAC-SHA1
    hmacSha1 := hmac.New(sha1.New, decodedKey)
    hmacSha1.Write([]byte(urlToSign))

    // Base64 encode the signature
    rawSignature := hmacSha1.Sum(nil)
    signature := base64.StdEncoding.EncodeToString(rawSignature)

    // Build the signed URL
    signedUrl := fmt.Sprintf("https://maps.googleapis.com%s&signature=%s", urlToSign, url.QueryEscape(signature))

    return signedUrl, nil
}

func main() {
    // Retrieve the API key and secret key from environment variables
    apiKey := os.Getenv("GOOGLE_API_KEY")
    apiSecret := os.Getenv("GOOGLE_API_SECRET") // Base64 encoded secret key

    // URL encode the address
    address := "東京都中央区銀座5−4−3 対鶴館 B2"
    encodedAddress := url.QueryEscape(address)

    // URL to be signed
    urlToSign := "/maps/api/staticmap?center=" + encodedAddress + "&zoom=18&size=400x400&format=png&maptype=roadmap&markers=color:red%7Clabel:H%7C" + encodedAddress + "&key=" + apiKey

    // Sign the URL
    signedUrl, err := signUrl(urlToSign, apiKey, apiSecret)
    if err != nil {
        fmt.Println("Error signing URL:", err)
        return
    }

    fmt.Println("Signed URL:", signedUrl)
}

