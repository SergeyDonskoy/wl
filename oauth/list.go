package oauth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/robdimsdale/wundergo"
)

// Lists returns all lists the client has permission to access.
func (c oauthClient) Lists() ([]wundergo.List, error) {
	url := fmt.Sprintf(
		"%s/lists",
		c.apiURL,
	)

	req, err := c.newGetRequest(url)
	if err != nil {
		return nil, err
	}

	resp, err := c.do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Unexpected response code %d - expected %d", resp.StatusCode, http.StatusOK)
	}

	lists := []wundergo.List{}
	err = json.NewDecoder(resp.Body).Decode(&lists)
	if err != nil {
		return lists, err
	}
	return lists, nil
}

// List returns the list for the corresponding listID.
func (c oauthClient) List(listID uint) (wundergo.List, error) {
	url := fmt.Sprintf(
		"%s/lists/%d",
		c.apiURL,
		listID,
	)

	req, err := c.newGetRequest(url)
	if err != nil {
		return wundergo.List{}, err
	}

	resp, err := c.do(req)
	if err != nil {
		return wundergo.List{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return wundergo.List{}, fmt.Errorf("Unexpected response code %d - expected %d", resp.StatusCode, http.StatusOK)
	}

	list := wundergo.List{}
	err = json.NewDecoder(resp.Body).Decode(&list)
	if err != nil {
		return wundergo.List{}, err
	}
	return list, nil
}

// CreateList creates a list with the provided title.
func (c oauthClient) CreateList(title string) (wundergo.List, error) {
	if title == "" {
		return wundergo.List{}, fmt.Errorf("title must be non-empty")
	}

	url := fmt.Sprintf("%s/lists", c.apiURL)
	body := []byte(fmt.Sprintf(`{"title":"%s"}`, title))

	req, err := c.newPostRequest(url, body)
	if err != nil {
		return wundergo.List{}, err
	}

	resp, err := c.do(req)
	if err != nil {
		return wundergo.List{}, err
	}

	if resp.StatusCode != http.StatusCreated {
		return wundergo.List{}, fmt.Errorf("Unexpected response code %d - expected %d", resp.StatusCode, http.StatusCreated)
	}

	list := wundergo.List{}
	err = json.NewDecoder(resp.Body).Decode(&list)
	if err != nil {
		return wundergo.List{}, err
	}
	return list, nil
}

// UpdateList updates the provided List.
func (c oauthClient) UpdateList(list wundergo.List) (wundergo.List, error) {
	body, err := json.Marshal(list)
	if err != nil {
		return wundergo.List{}, err
	}

	url := fmt.Sprintf(
		"%s/lists/%d",
		c.apiURL,
		list.ID,
	)

	req, err := c.newPatchRequest(url, body)
	if err != nil {
		return wundergo.List{}, err
	}

	resp, err := c.do(req)
	if err != nil {
		return wundergo.List{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return wundergo.List{}, fmt.Errorf("Unexpected response code %d - expected %d", resp.StatusCode, http.StatusOK)
	}

	returnedList := wundergo.List{}
	err = json.NewDecoder(resp.Body).Decode(&returnedList)
	if err != nil {
		return wundergo.List{}, err
	}
	return returnedList, nil
}

// DeleteList deletes the provided list.
func (c oauthClient) DeleteList(list wundergo.List) error {
	url := fmt.Sprintf(
		"%s/lists/%d?revision=%d",
		c.apiURL,
		list.ID,
		list.Revision,
	)

	req, err := c.newDeleteRequest(url)
	if err != nil {
		return err
	}

	resp, err := c.do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("Unexpected response code %d - expected %d", resp.StatusCode, http.StatusNoContent)
	}

	return nil
}

// DeleteAllLists gets a list of all lists via Lists() and deletes them
// via DeleteList(listID)
// It will not attempt to delete the inbox
func (c oauthClient) DeleteAllLists() error {
	lists, err := c.Lists()
	if err != nil {
		return err
	}

	listCount := len(lists)
	c.logger.Debug("delete-all-lists", map[string]interface{}{"listCount": listCount})
	idErrChan := make(chan idErr, listCount)
	for _, l := range lists {
		go func(list wundergo.List) {
			c.logger.Debug("delete-all-lists - deleting list", map[string]interface{}{"listID": list.ID})
			var err error
			if list.ListType == "inbox" {
				err = nil
			} else {
				err = c.DeleteList(list)
			}
			idErrChan <- idErr{idType: "list", id: list.ID, err: err}
		}(l)
	}

	e := multiIDErr{}
	for i := 0; i < len(lists); i++ {
		idErr := <-idErrChan
		if idErr.err != nil {
			c.logger.Debug("delete-all-lists - error received", map[string]interface{}{"id": idErr.id, "err": err})
			e.addError(idErr)
		}
	}

	if len(e.errors()) > 0 {
		return e
	}

	return nil
}
