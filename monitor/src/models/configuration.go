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
	TimeoutInSec int32
}

type CachetItem struct {
	Componentid int
	Incidentid int
}

type Cachet struct {
	Server string
	Token string
}