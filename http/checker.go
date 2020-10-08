package http

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	"wait4it/model"
)

func (h *HttpCheck) BuildContext(cx model.CheckContext) {
	h.Url = cx.Host
	h.Status = cx.HttpConf.StatusCode
	if len(cx.HttpConf.Text) > 0 {
		h.Text = cx.HttpConf.Text
	}
}

func (h *HttpCheck) Validate() error {
	if !h.validateUrl() {
		return errors.New("invalid URL provided")
	}

	if !h.validateStatusCode() {
		return errors.New("invalid status code provided")
	}

	return nil
}

func (h *HttpCheck) Check() (bool, bool, error) {
	resp, err := http.Get(h.Url)

	if err != nil {
		return false, true, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, true, err
	}

	if resp.StatusCode != h.Status {
		return false, false, errors.New("invalid status code")
	}

	if len(h.Text) > 0 {
		if !strings.Contains(string(body), h.Text) {
			return false, false, errors.New("can't find substring in response")
		}
	}

	return true, false, nil
}
