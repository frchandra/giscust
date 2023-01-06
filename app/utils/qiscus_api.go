package utils

import (
	"bitbucket.org/frchandra/giscust/app/validations"
	"bitbucket.org/frchandra/giscust/config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type QiscusApi struct {
}

func GetAllAgentsByDivision() (validations.AgentListResponse, error) {
	config := config.GetAppConfig()

	url := config.QiscusUrl + "/api/v2/admin/agents/by_division?division_ids[]=125015"
	method := "GET"

	var body []byte
	var agents validations.AgentListResponse

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return agents, err
	}
	req.Header.Add("Authorization", config.QiscusAuthnToken)
	req.Header.Add("Qiscus-App-Id", config.QiscusAppId)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return agents, err
	}
	defer res.Body.Close()

	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return agents, err
	}

	err = json.Unmarshal(body, &agents)
	if err != nil {
		fmt.Println(err)
	}

	//return getAllMyAgents(agents), err
	return agents, err
}

func AssignAgentToRoom(agentId int, roomId string) []byte {
	config := config.GetAppConfig()
	url := config.QiscusUrl + "/api/v1/admin/service/assign_agent"
	method := "POST"

	var params string = "room_id=" + roomId + "&agent_id=" + strconv.Itoa(agentId) + "&max_agent=1"
	payload := strings.NewReader(params)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Qiscus-App-Id", config.QiscusAppId)
	req.Header.Add("Qiscus-Secret-Key", config.QiscusSecretKey)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fmt.Println(string(body))
	return body
}
