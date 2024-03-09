package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageUrl *string) (LocationAreaRespone, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint
	if pageUrl != nil {
		fullURL = *pageUrl
	}
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreaRespone{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaRespone{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode > 399 {
		return LocationAreaRespone{}, fmt.Errorf("Bad status code: %v", resp.StatusCode)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaRespone{}, err
	}
	locationAreasResp := LocationAreaRespone{}
	err = json.Unmarshal(data, &locationAreasResp)
	if err != nil {
		return LocationAreaRespone{}, err
	}
	return locationAreasResp, nil
}