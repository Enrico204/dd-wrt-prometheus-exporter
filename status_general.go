package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (e *WRTExporter) getStatusGeneral() (InfoLive, error) {
	ret := InfoLive{}

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/Info.live.htm", e.routerURL), nil)
	if err != nil {
		return ret, err
	}

	if e.username != "" {
		req.SetBasicAuth(e.username, e.password)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return ret, err
	} else if resp.StatusCode != http.StatusOK {
		return ret, errors.New("unexpected status: " + resp.Status)
	}

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ret, err
	}
	_ = resp.Body.Close()

	ret.Scan(string(content))

	return ret, nil
}
