package main

import (
	"fmt"
	"net/http"
	"os"
	log "github.com/sirupsen/logrus"
	"github.com/julienschmidt/httprouter"
	"app/models"
	controller "app/controller"
	"io/ioutil"
	"github.com/ghodss/yaml"
	cachet "app/cachet"
	  
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	if os.Getenv("ENVIRONMENT") == "Production" {
		log.SetLevel(log.WarnLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}
}

func Option(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
}

func GetConfiguration() models.Configuration{

	file, e := ioutil.ReadFile("/app/Configs/config.yaml")
    if e != nil {
        fmt.Printf("Read config file error: %v\n", e)
        os.Exit(1)
	}

	var config models.Configuration
	err := yaml.Unmarshal(file, &config)
	
	if err != nil {
        fmt.Printf("Problem in config file: %v\n", err.Error())
        os.Exit(1)
	}

	config.Cachet.Server = os.Getenv("CACHET_SERVER_URL")
	config.Cachet.Token = os.Getenv("CACHET_TOKEN")

	return config
}

func main() {

	config := GetConfiguration()
	
	var controller controller.Controller
	for _, monitor := range config.Monitors {

		var test cachet.Cachet
		test.CreateIncident("test from monitor2", "this is a test2", monitor, config)
		controller.StartMonitor(monitor, config)
	}

	router := httprouter.New()
	fmt.Println("Running...")
	log.Fatal(http.ListenAndServe(":10010", router))

}