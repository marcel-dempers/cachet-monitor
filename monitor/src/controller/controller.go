package controller

import (
	 "fmt"
	 "net/http"
	 log "github.com/sirupsen/logrus"
	 "app/models"
	
	 "strconv"
	 "io/ioutil"
	 "encoding/json"
	 "time"
)

type Controller struct{
}

func (c Controller) StartMonitor(monitor models.Monitor, config models.Configuration) {
	fmt.Println("starting monitor: " + monitor.Name)
	
	monitor = Get_Cachet_Component_Statuses(monitor, config)
	
	switch monitor.Type {
		case "probe":
			fmt.Printf("Controller: starting probe type monitor %v \n", monitor.Name)
			go Controller_Probe_Start(monitor, config)
		default: fmt.Printf("Unsupported type %v \n", monitor.Type)
				 log.Fatal("exiting")
	}
}

//Gets latest component status from cachet
//Updates monitor status and returns latest monitor
func Get_Cachet_Component_Statuses(monitor models.Monitor, config models.Configuration ) models.Monitor {

	req, reqErr := http.NewRequest("GET", config.Cachet.Server + "/components/" + strconv.Itoa(monitor.Cachet.Componentid), nil)
	//req.Header.Add("content-type", "application/json")
	req.Header.Add("X-Cachet-Token", config.Cachet.Token)

	client := &http.Client{
		Timeout: time.Second * 200,
	}
	resp, reqErr := client.Do(req)
	
	if resp != nil {
		 
		body, readErr := ioutil.ReadAll(resp.Body)
		if readErr != nil {
			panic(readErr)
		}
		
		cachetComponent := models.CachetComponent{}

		jsonErr := json.Unmarshal(body, &cachetComponent)

		if jsonErr != nil {
			log.Fatal(jsonErr)
		}

		fmt.Printf("Cachet component: %v current status %v \n" , cachetComponent.Data.Id, cachetComponent.Data.Status)
		
		switch cachetComponent.Data.Status {
			case 1:
				monitor.Status = "Healthy"
			case 4:
				monitor.Status = "Failure"
		}

	}

	if reqErr != nil {
		fmt.Printf("Error message : %#v\n", reqErr.Error()) 
		panic(reqErr)
	} else {
		defer resp.Body.Close()
	}

	return monitor
}