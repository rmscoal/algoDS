package leetcode

// import (
// 	"encoding/base64"
// 	"fmt"
// 	"strings"
// )

// type Codec struct {
// 	baseTinyUrl string
// }

// func Constructor() Codec {
// 	return Codec{
// 		baseTinyUrl: "http://tinyurl.com/",
// 	}
// }

// // Encodes a URL to a shortened URL.
// func (this *Codec) encode(longUrl string) string {
// 	shortUrl := this.baseTinyUrl
// 	shortUrl += base64.URLEncoding.EncodeToString([]byte(longUrl))

// 	return shortUrl
// }

// // Decodes a shortened URL to its original URL.
// func (this *Codec) decode(shortUrl string) string {
// 	encodedUrl := strings.Split(shortUrl, this.baseTinyUrl)[1]
// 	originalUrl, err := base64.URLEncoding.DecodeString(encodedUrl)
// 	if err != nil {
// 		panic("Error")
// 	}

// 	return string(originalUrl)
// }

// func (D *EncodeAndDecodeURL) Solver(url string) {
// 	obj := Constructor()
// 	tinyURL := obj.encode(url)
// 	originalURL := obj.decode(tinyURL)

// 	fmt.Println("Tiny URL:", tinyURL)
// 	fmt.Println("Original URL:", originalURL)
// }

// // Funny solution 1
// // type Codec struct {
// // 	count int
// // }

// // func Constructor() Codec {
// // 	return Codec{ count: 0 }
// // }

// // // Encodes a URL to a shortened URL.
// // func (this *Codec) encode(longUrl string) string {
// // newStr := longUrl + " "
// // 	this.count++
// // 	return newStr
// // }

// // // Decodes a shortened URL to its original URL.
// // func (this *Codec) decode(shortUrl string) string {
// // 	return shortUrl[:len(shortUrl) - 1]
// // }

// // Funny solution 2
// // var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// // func RandStringRunes(n int) string {
// // 	b := make([]rune, n)
// // 	for i := range b {
// // 		b[i] = letterRunes[rand.Intn(len(letterRunes))]
// // 	}
// // 	return string(b)
// // }

// // type Codec struct {
// // 	longUrls  map[string]string
// // 	shortUrls map[string]string
// // }

// // func Constructor() Codec {
// // 	m := make(map[string]string)
// // 	m2 := make(map[string]string)
// // 	return Codec{longUrls: m, shortUrls: m2}
// // }

// // // Encodes a URL to a shortened URL.
// // func (this *Codec) encode(longUrl string) string {
// // 	random := RandStringRunes(7)
// // 	this.longUrls[longUrl] = random
// // 	this.shortUrls[random] = longUrl
// // 	return random
// // }

// // // Decodes a shortened URL to its original URL.
// // func (this *Codec) decode(shortUrl string) string {
// // 	return this.shortUrls[shortUrl]
// // }

// /**
//  * Your Codec object will be instantiated and called as such:
//  * obj := Constructor();
//  * url := obj.encode(longUrl);
//  * ans := obj.decode(url);
//  */

// // https://leetcode.com/discuss/interview-question/124658/Design-a-URL-Shortener-(-TinyURL-)-System/
