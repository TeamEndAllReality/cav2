package cav2

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
)

const (
	//APIEndpoint URL of API
	APIEndpoint = "https://addons-ecs.forgesvc.net/api/v2/"
)

var (
	//EMPTY Empty req data
	EMPTY []byte
)

//GetAddon gets Addon with addonID
func GetAddon(addonID string) (*Addon, error) {
	response, err := GetHTTPResponse("GET", getAPI()+"addon/"+addonID, EMPTY)
	if err != nil {
		return nil, err
	}

	var addonResponse *Addon
	err = json.NewDecoder(response.Body).Decode(&addonResponse)
	if err != nil {
		return nil, err
	}
	return addonResponse, nil
}

//GetAddons Gets addons with []AddonID
func GetAddons(addons []int) ([]*Addon, error) {
	jsonPayload, _ := json.Marshal(addons)
	response, err := GetHTTPResponse("POST", getAPI()+"addon", jsonPayload)
	if err != nil {
		return nil, err
	}
	var addonResponse []*Addon
	err = json.NewDecoder(response.Body).Decode(&addonResponse)
	if err != nil {
		return nil, err
	}
	return addonResponse, nil
}

//GetAddonFile Gets an addons file with AddonID and fileId
func GetAddonFile(addon int, fileID int) (*File, error) {
	response, err := GetHTTPResponse("GET", getAPI()+"addon/"+strconv.Itoa(addon)+"/file/"+strconv.Itoa(fileID), EMPTY)
	if err != nil {
		return nil, err
	}
	var addonResponse *File
	err = json.NewDecoder(response.Body).Decode(&addonResponse)
	if err != nil {
		return nil, err
	}
	return addonResponse, nil
}

//GetAddonFiles Gets files for addonID
func GetAddonFiles(addonID string) ([]*File, error) {
	response, err := GetHTTPResponse("GET", getAPI()+"addon/"+addonID+"/files", EMPTY)
	if err != nil {
		return nil, err
	}

	var addonResponse []*File
	err = json.NewDecoder(response.Body).Decode(&addonResponse)
	if err != nil {
		return nil, err
	}
	return addonResponse, nil
}

//GetHashMatches Gets addon Fingerprints
func GetHashMatches(addons []int) (*FingerprintList, error) {
	jsonPayload, _ := json.Marshal(addons)
	response, err := GetHTTPResponse("POST", APIEndpoint+"fingerprint", jsonPayload)
	if err != nil {
		return nil, err
	}
	var addonResponse *FingerprintList
	err = json.NewDecoder(response.Body).Decode(&addonResponse)
	if err != nil {
		return nil, err
	}
	return addonResponse, nil
}

//GetAddonDatabaseTimestamp Gets last update of addon DB
func GetAddonDatabaseTimestamp() (string, error) {
	response, err := GetHTTPResponse("GET", getAPI()+"addon/timestamp", EMPTY)
	if err != nil {
		return "", err
	}

	bodyBytes, _ := ioutil.ReadAll(response.Body)
	bodyString := string(bodyBytes)

	return bodyString, nil
}

//Search searches for addons with query
func Search(query string) ([]*Addon, error) {
	//TODO expand this to allow more things to be searched for
	//"api/addon/search?gameId={0}&sectionId={1}&categoryId={2}&gameVersion={3}&index={4}&pageSize={5}&searchFilter={6}&sort={7}&sortDescending={8}"
	seachPayload := "?gameId=432&sectionId=-1&categoryId=-1&index=0&pageSize=1000&sort=TotalDownloads&sortDescending=true&searchFilter=" + query
	response, err := GetHTTPResponse("GET", getAPI()+"addon/search"+seachPayload, EMPTY)
	if err != nil {
		return nil, err
	}
	var addonResponse []*Addon
	err = json.NewDecoder(response.Body).Decode(&addonResponse)
	if err != nil {
		return nil, err
	}
	return addonResponse, nil
}

//GetAllAddons Uses the search to find as many addons as possible, not perfect but should get most of them.
//This is a very expensive call, takes a while and uses a lot of api requests
func GetAllAddons() ([]*Addon, error) {
	allAddons := make([]*Addon, 0)

	lastFind := 1000
	index := 0
	for lastFind == 1000 {
		if index > 100 {
			fmt.Println("Breaking out index too big")
			break
		}
		seachPayload := "?gameId=432&sectionId=6&categoryId=0&index=" + strconv.Itoa(len(allAddons)) + "&searchFilter=&pageSize=1000&sort=5"
		response, err := GetHTTPResponse("GET", APIEndpoint+"addon/search"+seachPayload, EMPTY)
		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}
		var addonResponse []*Addon
		err = json.NewDecoder(response.Body).Decode(&addonResponse)
		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}

		lastFind = len(addonResponse)
		index++
		allAddons = append(allAddons, addonResponse...)
	}

	return allAddons, nil
}

func getAPI() string {
	return APIEndpoint
}
