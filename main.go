package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

type FollowerData struct {
	Href      string `json:"href"`
	Value     string `json:"value"`
	Timestamp int    `json:"timestamp"`
}

type StringListData struct {
	StringListData []FollowerData `json:"string_list_data"`
}

func main() {
	followersfile := "data/followers_1.json"
	followingfile := "data/following.json"

	followerData, err := readJSONFile(followersfile)
	if err != nil {
		panic(err)
	}

	followingData, err := readJSONFile(followingfile)
	if err != nil {
		panic(err)
	}

	followerListData := getFollowData(followerData)
	followingListData := getFollowData(followingData)

	FormatTable(followingListData, followerListData)
}

func readJSONFile(filename string) ([]StringListData, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data []StringListData

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func getFollowData(stringListData []StringListData) []FollowerData {
	var followerData []FollowerData

	for _, item := range stringListData {
		followerData = append(followerData, item.StringListData...)
	}

	return followerData
}

func FormatTable(followings, followers []FollowerData) {
	followerStatus := make(map[string]struct{})

	for _, follower := range followers {
		followerStatus[follower.Value] = struct{}{}
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Username", "Following Me?"})

	for _, following := range followings {
		_, isFollower := followerStatus[following.Value]
		followerStatusStr := fmt.Sprintf("%v", isFollower)
		
		table.Append([]string{following.Value, followerStatusStr})
	}
}