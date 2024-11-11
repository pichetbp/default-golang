package models

import "time"

// url description : https://developers.facebook.com/docs/facebook-login/guides/access-tokens/
type FacebookAuthenticationReq struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type FacebookAuthenticationRes struct {
}

type AdLabel struct {
	ID          string    `json:"id"`
	CreatedTime time.Time `json:"created_time"`
	Name        string    `json:"name"`
	UpdatedTime time.Time `json:"updated_time"`
}

// url description : https://developers.facebook.com/docs/marketing-api/reference/ad-campaign-group
type FacebookGetCampaignReq struct {
	ID        string    `json:"id"`
	AccountId string    `json:"account_id"`
	AdBatch   []AdLabel `json:"ad_batch"`
	Name      string    `json:"name"`
}

type FacebookGetCampaignRes struct {
}
