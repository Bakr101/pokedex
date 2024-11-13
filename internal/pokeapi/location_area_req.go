package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)
func (c *Client) ListLocationAreas(pageURL *string) (LocationArea, error){
	endpoint := "/location-area"
	fullURL := baseUrl + endpoint
	if pageURL != nil{
		fullURL = *pageURL
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationArea{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 399 {
		return LocationArea{}, fmt.Errorf("bad status code: %v", res.StatusCode)
	}

	jsonData, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationArea{}, err
	}

	locationAreaResp := LocationArea{}
	if err := json.Unmarshal(jsonData, &locationAreaResp); err != nil{
		return LocationArea{}, fmt.Errorf("error unmarshalling json: %w", err)
	} 

	return locationAreaResp, nil
	
}