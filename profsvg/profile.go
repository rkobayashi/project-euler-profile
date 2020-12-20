package profsvg

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type profile struct {
	UserName string `xml:"username"`
	Country  string `xml:"country"`
	Language string `xml:"language"`
	Solved   string `xml:"solved"`
	Level    int    `xml:"level"`
}

func newProfile(buf []byte) (prof profile, err error) {
	err = xml.Unmarshal(buf, &prof)
	return
}

func getProfileFromServer(name string) (profile, error) {
	buf, err := getXML(name, http.Get)
	if err != nil {
		return profile{}, err
	}

	return newProfile(buf)
}

func getXML(name string, get func(url string) (*http.Response, error)) ([]byte, error) {
	url := fmt.Sprintf("https://projecteuler.net/profile/%s.xml", name)

	res, err := get(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("response error: status code = %v", res.StatusCode)
	}

	return ioutil.ReadAll(res.Body)
}
