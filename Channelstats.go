package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type ChannelStats struct {
	Subscribers    int    `json:"subscribers"`
	ChannelName    string `json:"channelName"`
	MinutesWatched int    `json:"minutesWatched"`
	Views          int    `json:"views"`
}

func GetChannelStats() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		//w.Write([]byte("response!"))
		yt := ChannelStats{
			Subscribers:    5,
			ChannelName:    "my channel",
			MinutesWatched: 50,
			Views:          100,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(yt); err != nil {
			panic(err)
		}
	}

}
