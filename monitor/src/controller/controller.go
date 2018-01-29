package controller

import (
	 "fmt"
	 log "github.com/sirupsen/logrus"
	 "app/models"
)

type Controller struct{
}

func (c Controller) StartMonitor(monitor models.Monitor, config models.Configuration) {
	fmt.Println("starting monitor: " + monitor.Name)

	switch monitor.Type {
		case "probe":
			fmt.Printf("Controller: starting probe type monitor %v \n", monitor.Name)
			go Controller_Probe_Start(monitor, config)
		default: fmt.Printf("Unsupported type %v \n", monitor.Type)
				 log.Fatal("exiting")
	}
	
}