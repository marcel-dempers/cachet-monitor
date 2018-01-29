package cachet

import (
	"fmt"
	"net/http"
	"app/models"
	"strconv"
    "strings"
	"time"
)

type Cachet struct{
}


func (c Cachet) UpdateComponent(componentid int, status string, monitor models.Monitor, config models.Configuration ) models.Monitor {
	
	fmt.Printf("Updating Cachet component :  %v\n", monitor.Cachet.Componentid)
	monitor.Status = status

	var statusid string = "0"

	switch status {
		case "Failure":
			statusid = "4"
		case "Healthy":
			statusid = "1"
	}
	payload := strings.NewReader("{\n    \n    \"status\": " + statusid + " \n  \n}")
	req, reqErr := http.NewRequest("PUT", config.Cachet.Server + "/components/" + strconv.Itoa(componentid), payload)
	req.Header.Add("content-type", "application/json")
	req.Header.Add("X-Cachet-Token", config.Cachet.Token)

	client := &http.Client{
		Timeout: time.Second * 10,
	}
	resp, reqErr := client.Do(req)
	
	if resp != nil {
		fmt.Printf("%#v \n", resp)
	}
	if reqErr != nil {
		fmt.Printf("Error message : %#v\n", reqErr.Error()) 
		panic(reqErr)
	} else {
		fmt.Println("Updated component status success")
		defer resp.Body.Close()
	}
	
	fmt.Printf("Update Cachet component - status : %v component: %v \n", status, componentid)

	return monitor
}

func (c Cachet) CreateIncident(name string, message string, monitor models.Monitor , config models.Configuration ){

	fmt.Printf("Creating Cachet incident- name: %v componentid: %v message: %v \n", name,  monitor.Cachet.Componentid, message)
	
	payload := strings.NewReader("{\n\t\"name\" : \"" + name + "\",\n\t\"message\" : \"" + message + "\",\n\t\"status\" : 2,\n\t\"component_id\" : 1,\n\t\"component_status\": 4,\n\t\"visible\": 1\n}")
	req, reqErr := http.NewRequest("POST", config.Cachet.Server + "/incidents", payload)
	req.Header.Add("content-type", "application/json")
	req.Header.Add("X-Cachet-Token", config.Cachet.Token)

	client := &http.Client{
		Timeout: time.Second * 10,
	}
	resp, reqErr := client.Do(req)
	
	if resp != nil {
		fmt.Printf("%#v \n", resp)
	}
	if reqErr != nil {
		fmt.Printf("Error message : %#v\n", reqErr.Error()) 
		panic(reqErr)
	} else {
		fmt.Println("Created incident success")
		defer resp.Body.Close()
	}

}

func (c Cachet) UpdateIncident(incidentid int, status string, message string, monitor models.Monitor , config models.Configuration ){
	
		fmt.Printf("Updating Cachet incident- id: %v componentid: %v message: %v \n", strconv.Itoa(incidentid),  monitor.Cachet.Componentid, message)
		
		var statusid string = "0"
		
			switch status {
				case "Investigating":
					statusid = "1"
				case "Identified":
					statusid = "2"
				case "Watching":
					statusid = "3"
				case "Fixed":
					statusid = "4"
			}

		payload := strings.NewReader("{\n\t\"message\" : \"" + message + "\",\n\t\"status\" : " + statusid + "\n}")
		
		fmt.Println(payload)
		var url string =  config.Cachet.Server + "/incidents/" + strconv.Itoa(incidentid) + "/updates"
		req, reqErr := http.NewRequest("POST",url, payload)
		req.Header.Add("content-type", "application/json")
		req.Header.Add("X-Cachet-Token", config.Cachet.Token)
	
		client := &http.Client{
			Timeout: time.Second * 10,
		}
		resp, reqErr := client.Do(req)
		
		if resp != nil {
			fmt.Printf("%#v \n", resp)
		}
		if reqErr != nil {
			fmt.Printf("Error message : %#v\n", reqErr.Error()) 
			panic(reqErr)
		} else {
			fmt.Println("Update incident success")
			defer resp.Body.Close()
		}
	
	}