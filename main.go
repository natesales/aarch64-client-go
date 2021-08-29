package aarch64

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

// API Server
const server = "https://console.aarch64.com/api"

type Client struct {
	APIKey string

	//reuse http client across requests
	client *http.Client
}

type APIMeta struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type APIResponse struct {
	Data interface{} `json:"data"`
	Meta APIMeta     `json:"meta"`
}

type VM struct {
	Id      string `json:"_id"`
	Project string `json:"project"`
	PoP     string `json:"pop"`
	Host    uint   `json:"host"`
	Index   uint   `json:"index"`

	Hostname string `json:"hostname"`
	VCPUs    uint   `json:"vcpus"`
	Memory   uint   `json:"memory"`
	Disk     uint   `json:"disk"`
	OS       string `json:"os"`
	Prefix   string `json:"prefix"`
	Gateway  string `json:"gateway"`
	Address  string `json:"address"`

	Password string `json:"password"`

	PhonedHome bool `json:"phoned_home"`
}

type Project struct {
	Id    string   `json:"_id"`
	Name  string   `json:"name"`
	Users []string `json:"users"`
	VMs   []VM     `json:"vms"`
}

type ProjectsResponse struct {
	Meta     APIMeta   `json:"meta"`
	Projects []Project `json:"data"`
}

func NewClient(APIKey string) Client {
	return Client{APIKey: APIKey, client: &http.Client{}}
}

func (c Client) req(method string, endpoint string, body interface{}, output interface{}) error {
	var _body io.Reader

	// If there is data in the body
	if body != nil {
		// Marshal body interface as JSON
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return err
		}
		_body = bytes.NewBuffer(jsonBody)
	}

	req, err := http.NewRequest(method, server+endpoint, _body)
	if err != nil {
		return err
	}

	// Set request headers
	req.Header.Set("Content-Type", "application/json")
	if c.APIKey != "" {
		req.Header.Set("Authorization", c.APIKey)
	}

	// Send the request
	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}

	if err := json.NewDecoder(resp.Body).Decode(output); err != nil {
		return err
	}

	return nil // nil error
}

func (c Client) Projects() (ProjectsResponse, error) {
	var resp ProjectsResponse
	if err := c.req("GET", "/projects", nil, &resp); err != nil {
		return ProjectsResponse{}, err
	}

	return resp, nil
}

func (c Client) CreateProject(name string) (APIResponse, error) {
	var resp APIResponse
	if err := c.req("POST", "/project", map[string]string{"name": name}, &resp); err != nil {
		return APIResponse{}, err
	}

	return resp, nil
}

func (c Client) AddUser(projectId string, email string) (APIResponse, error) {
	var resp APIResponse
	if err := c.req("POST", "/vms/adduser", map[string]string{
		"project": projectId,
		"email":   email,
	}, &resp); err != nil {
		return APIResponse{}, err
	}

	return resp, nil
}

func (c Client) CreateVM(hostname string, pop string, projectId string, plan string, os string) (APIResponse, error) {
	var resp APIResponse
	if err := c.req("POST", "/vms/create", map[string]string{
		"hostname": hostname,
		"pop":      pop,
		"project":  projectId,
		"plan":     plan,
		"os":       os,
	}, &resp); err != nil {
		return APIResponse{}, err
	}

	return resp, nil
}

func (c Client) DeleteVM(vm string) (APIResponse, error) {
	var resp APIResponse
	if err := c.req("DELETE", "/vms/delete", map[string]string{"vm": vm}, &resp); err != nil {
		return APIResponse{}, err
	}

	return resp, nil
}

func (c Client) SignUp(email string, password string) (APIResponse, error) {
	var resp APIResponse
	if err := c.req("POST", "/auth/signup", map[string]string{"email": email, "password": password}, &resp); err != nil {
		return APIResponse{}, err
	}

	return resp, nil
}

// Currently the aarch64 api only sends the Api key through the set-cookie headers.
func (c Client) Login(email string, password string) (APIResponse, error) {
	var resp APIResponse
	if err := c.req("POST", "/auth/login", map[string]string{"email": email, "password": password}, &resp); err != nil {
		return APIResponse{}, err
	}
	return resp, nil
}
