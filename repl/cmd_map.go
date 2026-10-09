package repl

import "fmt"

type Locations struct {
	Count    int    `json:"count"`
	Next     *string `json:"next"`
	Previous *string    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}



func commandMap(conf *config) error {
	var location Locations

	if conf.next != nil {
		err := fetch(*conf.next, &location)
		
		if err != nil {
			return  err
		}
		conf.next = location.Next
		conf.previous = location.Previous

		for _, result := range location.Results {
			
			fmt.Printf(" - %s\n", result.Name)
		}
		return nil
	}

	err := fetch(baseUrl + "location", &location)

	if err != nil {
		return err
	}
  conf.next = location.Next
	conf.previous = location.Previous
	for _, result := range location.Results {

		fmt.Printf(" - %s\n", result.Name)
	}


	return nil
}

func mapB(conf *config) error {
	var location Locations
	if conf.previous != nil {
		err := fetch(*conf.previous, &location)
		
		if err != nil {
			return  err
		}
		conf.next = location.Next
		conf.previous = location.Previous

		for _, result := range location.Results {
			fmt.Printf(" - %s\n", result.Name)
		}
		return nil
	}
	
	fmt.Println("No previous page")
	return nil
}
