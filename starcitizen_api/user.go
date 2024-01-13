package starcitizen_api

import (
	"github.com/imroc/req/v3"
	"party_room_server/build_conf"
)

type UserData struct {
	Data struct {
		Affiliation []struct {
			Image string `json:"image"`
			Name  string `json:"name"`
			Rank  string `json:"rank"`
			Sid   string `json:"sid"`
			Stars int    `json:"stars"`
		} `json:"affiliation"`
		Organization struct {
			Image string `json:"image"`
			Name  string `json:"name"`
			Rank  string `json:"rank"`
			Sid   string `json:"sid"`
			Stars int    `json:"stars"`
		} `json:"organization"`
		Profile struct {
			Badge      string   `json:"badge"`
			BadgeImage string   `json:"badge_image"`
			Display    string   `json:"display"`
			Enlisted   string   `json:"enlisted"`
			Fluency    []string `json:"fluency"`
			Handle     string   `json:"handle"`
			Id         string   `json:"id"`
			Image      string   `json:"image"`
			Location   struct {
				Country string `json:"country"`
				Region  string `json:"region"`
			} `json:"location"`
			Page struct {
				Title string `json:"title"`
				Url   string `json:"url"`
			} `json:"page"`
			Website string `json:"website"`
		} `json:"profile"`
	} `json:"data"`
	Message string `json:"message"`
	Source  string `json:"source"`
	Success int    `json:"success"`
}

func GetUserData(userName string) (*UserData, error) {
	resp := req.MustGet("https://api.starcitizen-api.com/" + build_conf.StarCitizenWikiUserApiKey + "/v1/auto/user/" + userName)
	var userData UserData
	return &userData, resp.UnmarshalJson(&userData)
}
