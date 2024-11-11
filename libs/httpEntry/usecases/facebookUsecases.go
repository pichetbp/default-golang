package usecases

import (
	"errors"
	"net/http"
	"time"
)

var FacebookEntry FacebookEntryInterface

type FacebookEntryInterface interface {
}

type facebookEntry struct {
}

func NewFacebookEntry() FacebookEntryInterface {
	FacebookEntry = &facebookEntry{}
	return FacebookEntry
}

func (f *facebookEntry) Authentication() error {
	// add stirng from env
	apiUrl := "https://graph.facebook.com/{api-endpoint}&access_token={your-app_id}|{your-app_secret}"
	timeStart := time.Now()
	var err error
	defer func() {
		if err != nil {
		}
		timeEnd := time.Now()
		_ = timeEnd.Sub(timeStart)
	}()

	reqAuth, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		return err
	}

	q := reqAuth.URL.Query()
	q.Add("client_id", "your_client_id")
	q.Add("client_secret", "your_client_secret")
	q.Add("grant_type", "client_credentials")
	reqAuth.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err := client.Do(reqAuth)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("error")
	}

	// add response struct

	// add convert and return status

	return nil
}

func (f *facebookEntry) GetCampaign() error {
	return nil
}

func (f *facebookEntry) CreateCampaign() error {
	// add stirng from env
	apiUrl := "https://graph.facebook.com/v21.0/act_<AD_ACCOUNT_ID>/campaigns"
	timeStart := time.Now()
	var err error
	defer func() {
		if err != nil {
		}
		timeEnd := time.Now()
		_ = timeEnd.Sub(timeStart)
	}()

	reqAuth, err := http.NewRequest("POST", apiUrl, nil)
	if err != nil {
		return err
	}

	q := reqAuth.URL.Query()
	q.Add("client_id", "your_client_id")
	q.Add("client_secret", "your_client_secret")
	q.Add("grant_type", "client_credentials")
	reqAuth.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err := client.Do(reqAuth)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("error")
	}
	return nil
}

func (f *facebookEntry) UploadCustomAudience() error {
	// do multi thread call api
	return nil
}
