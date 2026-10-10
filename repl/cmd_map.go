package repl

import (
	"encoding/json"
	"fmt"
)

type Locations struct {
	Count    int    `json:"count"`
	Next     *string `json:"next"`
	Previous *string    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}


func printLocations(data []byte, conf *Config) error{
	var location Locations
	
	if err := json.Unmarshal(data, &location); err != nil {
		return err
	}
	
	conf.next = location.Next
	conf.previous = location.Previous
	
	for _, result := range location.Results {
		fmt.Printf(" - %s\n", result.Name)
	}
	return nil
}


func commandMap(conf *Config) error {
  if conf.next != nil {
    if data, exists := conf.cache.Get(*conf.next); exists {
		
			fmt.Println("Using cache")
		 	return  printLocations(data, conf)
		}
	}

	if conf.next != nil {				
		data, err := conf.client.GetData(*conf.next)
		
		if err != nil {
			return  err
		}

		conf.cache.Add(*conf.next, data)

		return printLocations(data, conf)
	}

	url := baseUrl + "location"
	data, err := conf.client.GetData(url)
		
	if err != nil {
		return  err
	}

	conf.cache.Add(url, data)

	return printLocations(data, conf)
}

func mapB(conf *Config) error {

	if conf.previous != nil {
	  if data, exists := conf.cache.Get(*conf.previous); exists {
			fmt.Println("Using cache")
			return printLocations(data, conf)
		}
	}

	if conf.previous != nil {
		data, err := conf.client.GetData(*conf.previous)
		
		if err != nil {
			return  err
		}

		conf.cache.Add(*conf.previous, data)

		return printLocations(data, conf)
	}
	
	fmt.Println("No previous page")
	return nil
}
