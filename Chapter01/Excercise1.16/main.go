package main

import "fmt"

const GlobalLimit = 100
const MaxCacheSize int = 10 * GlobalLimit

const (
	CacheKeyBook = "book_"
	CacheKeyCD   = "cd_"
)

var cache map[string]string

func cacheGet(key string) string {
	return cache[key]
}

func cacheSet(key, val string) {
	if len(cache)+1 >= MaxCacheSize {
		return
	}
	cache[key] = val
}

func getBook(isbn string) string {
	return cacheGet(CacheKeyBook + isbn)
}

func setBook(isbn string, name string) {
	cacheSet(CacheKeyBook+isbn, name)
}

func getCD(sku string) string {
	return cacheGet(CacheKeyCD + sku)
}

func setCD(sku string, title string) {
	cacheSet(CacheKeyCD+sku, title)
}

func main() {
	// init cache
	cache = make(map[string]string)
	setBook("1234-5678", "Get Ready To Go")
	setCD("1234-5678", "Get Ready To Go Audio Book")
	fmt.Println("Book :", getBook("1234-5678"))
	fmt.Println("CD :", getCD("1234-5678"))

}
