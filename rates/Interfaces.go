package rates

type RequestMaker interface {
	makeGet(url string) Coins
}
