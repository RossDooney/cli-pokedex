package pokeapi

import (
	"net/http"
	"fmt"
	"os"
)

func (c *Client) ListLocationAreas() (LocationAreaResp, error){
	endpoint := "/location-area"
	fullURL := baseURL + endpoint

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil{
		return LocationAreaResp{}, err
	}
	
	resp, err := c.httpClient.Do(req)

	if err != nil{
		return LocationAreaResp{}, err
	}

	defer resp.body.Close()

	if resp.StatusCode > 399{
		return LocationAreaResp{}, fmt.Println("Bad Status code %v", resp.StatusCode)
	}

	io.ReadAll(resp.body)

}