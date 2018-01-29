package models

type Configuration struct {
	Monitors []Monitor
	Cachet Cachet
}

type Monitor struct {
	Name string
	Type string
	Url string
	Method string
	Frequency int
	ExpectedResponseCode int
	Maxfailures int
	Cachet CachetItem
	Status string
}

type CachetItem struct {
	Componentid int
}

type Cachet struct {
	Server string
	Token string
}