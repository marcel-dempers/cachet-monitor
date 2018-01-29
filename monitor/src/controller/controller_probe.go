package controller

import (
	 
	 "fmt"
	 "net/http"
	 "app/models"
	 "time"
	 "net"
	 "strings"
	 
	 "math/rand"
	 cachet "app/cachet"
)

func Controller_Probe_Start(monitor models.Monitor, config models.Configuration){
	fmt.Println("Starting monitor : " + monitor.Name + " ...")
	fmt.Println(monitor)
	var failureCount int = 0
	var recoveryCount int = 0
	var recoveryReached int = 3

	var cachet cachet.Cachet
	for {
        
		req, reqErr := http.NewRequest(monitor.Method, monitor.Url, nil)
		fmt.Printf("Connecting %v ... \n" , monitor.Name)
		client := &http.Client{
			Timeout: time.Second * time.Duration(rand.Int31n(monitor.TimeoutInSec)),
		}
		resp, reqErr := client.Do(req)

		if reqErr != nil {
			if reqErr, ok := reqErr.(net.Error); ok && reqErr.Timeout() {
				fmt.Printf("Monitor %v timeout ... \n" , monitor.Name, reqErr)
			} else if reqErr, ok := reqErr.(net.Error); ok && reqErr.Temporary() {
				fmt.Printf("Monitor %v temporary failure ... \n" , monitor.Name, reqErr)
			} else if strings.Contains(reqErr.Error(), "getsockopt") || strings.Contains(reqErr.Error(), "dial tcp: lookup") { 
				fmt.Printf("Monitor %v network failure ... \n" , monitor.Name, reqErr)
			} else {
				fmt.Printf("Error message : %#v\n", reqErr.Error()) 
				panic(reqErr)
			}
		} else {
			defer resp.Body.Close()
		}
		
		//failure handlers
		
		//failure making request - getsockops errors
		if resp == nil && reqErr != nil && (strings.Contains(reqErr.Error(), "getsockopt") || strings.Contains(reqErr.Error(), "dial tcp: lookup")) {
			failureCount++
			fmt.Printf("Monitor failure making request (getsockopt): %v [ %v / %v ] \n" , monitor.Name, failureCount , monitor.Maxfailures)
			
			if failureCount == monitor.Maxfailures {
				fmt.Printf("Monitor: %v max failures reached %v / %v\n", monitor.Name, failureCount ,monitor.Maxfailures)
				monitor = cachet.UpdateComponent(monitor.Cachet.Componentid, "Failure", monitor, config)
				fmt.Printf("Monitor: %v is now in a status: %v \n" , monitor.Name, monitor.Status)
			}

		//failure making request - timeout
		} else if reqErr, ok := reqErr.(net.Error); ok && reqErr.Timeout() && reqErr != nil  {
			failureCount++
			fmt.Printf("Monitor failure making request (timeout): %v [ %v / %v ] \n" , monitor.Name, failureCount , monitor.Maxfailures)
			
			if failureCount == monitor.Maxfailures {
				fmt.Printf("Monitor: %v max failures reached %v / %v\n", monitor.Name, failureCount ,monitor.Maxfailures)
				monitor = cachet.UpdateComponent(monitor.Cachet.Componentid, "Failure", monitor, config)
				fmt.Printf("Monitor: %v is now in a status: %v \n" , monitor.Name, monitor.Status)
			}
		//failure by status code from endpoint!	
		} else if resp != nil && resp.StatusCode != monitor.ExpectedResponseCode && monitor.Status != "Failure" {
			failureCount++
			fmt.Printf("Monitor failure detected: %v | Response Status: %v [ %v / %v ] \n" , monitor.Name, resp.Status, failureCount , monitor.Maxfailures)
			
			if failureCount == monitor.Maxfailures {
				fmt.Printf("Monitor: %v max failures reached %v / %v\n", monitor.Name, failureCount ,monitor.Maxfailures)
				monitor = cachet.UpdateComponent(monitor.Cachet.Componentid, "Failure", monitor, config)
				fmt.Printf("Monitor: %v is now in a status: %v \n" , monitor.Name, monitor.Status)
			}
		//ongoing failure - response code
		} else if resp != nil && resp.StatusCode != monitor.ExpectedResponseCode && monitor.Status == "Failure" {
			//fmt.Printf("Monitor: %v remains in a status: %v \n" , monitor.Name, monitor.Status)
		//potential recovery
		} else if resp != nil && resp.StatusCode == monitor.ExpectedResponseCode && monitor.Status == "Failure" {

			if recoveryCount == recoveryReached {
				fmt.Printf("Monitor: %v reached reached recovery status \n", monitor.Name)
				monitor = cachet.UpdateComponent(monitor.Cachet.Componentid, "Healthy", monitor, config)
				fmt.Printf("Monitor: %v is now in a status: %v \n" , monitor.Name, monitor.Status)
			}

			fmt.Printf("Monitor: %v completed a success probe and in potential recovery. Current status: %v [ %v / %v ] \n" , monitor.Name, monitor.Status , recoveryCount, recoveryReached)
			recoveryCount++
		} else {
			//reset failure on every loop - because consequetive failure needed
			failureCount = 0
			//reset recover on every loop - because consequetive successes needed
			recoveryCount = 0
		}

		time.Sleep(time.Duration(monitor.Frequency)*time.Second)

	}

	fmt.Println("Monitor : " + monitor.Name + " is done!")
}