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