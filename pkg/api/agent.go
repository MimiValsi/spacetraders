package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/MimiValsi/spacetraders/pkg/model"
)


func (c *Client) Register(symbol, faction string) error {
	url := url.URL{Path: "register"}
	uri := c.BaseURI.ResolveReference(&url)

	data := strings.NewReader(fmt.Sprintf("{\"symbol\": \"%s\", \"faction\": \"%s\"}", symbol, faction))

	req, err := http.NewRequest("POST", uri.String(), data)
	if err != nil {
		return err
	}

	req.Header = *c.Header
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	agent := model.AgentRegister{}
	if err = decoder.Decode(&agent); err != nil {
		return err
	}

	fmt.Printf("Register: %+v\n", agent)

	return nil
}
