package oauth

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/robdimsdale/wundergo"
)

// Notes gets all tasks for all lists.
func (c oauthClient) Notes() ([]wundergo.Note, error) {
	lists, err := c.Lists()
	if err != nil {
		return nil, err
	}

	listCount := len(lists)
	c.logger.Debug(
		"tasks",
		map[string]interface{}{"listCount": listCount},
	)

	notesChan := make(chan []wundergo.Note, listCount)
	idErrChan := make(chan idErr, listCount)
	for _, l := range lists {
		go func(list wundergo.List) {
			c.logger.Debug(
				"notes - getting notes for list",
				map[string]interface{}{"listID": list.ID},
			)
			notes, err := c.NotesForListID(list.ID)
			idErrChan <- idErr{idType: "list", id: list.ID, err: err}
			notesChan <- notes
		}(l)
	}

	e := multiIDErr{}
	for i := 0; i < listCount; i++ {
		idErr := <-idErrChan
		if idErr.err != nil {
			c.logger.Debug(
				"notes - error received getting notes for list",
				map[string]interface{}{"listID": idErr.id, "err": err},
			)
			e.addError(idErr)
		}
	}

	totalNotes := []wundergo.Note{}
	for i := 0; i < listCount; i++ {
		notes := <-notesChan
		totalNotes = append(totalNotes, notes...)
	}

	if len(e.errors()) > 0 {
		return totalNotes, e
	}

	return totalNotes, nil
}

// NotesForListID returns Notes for the provided listID.
func (c oauthClient) NotesForListID(listID uint) ([]wundergo.Note, error) {
	if listID == 0 {
		return nil, errors.New("listID must be > 0")
	}

	url := fmt.Sprintf(
		"%s/notes?list_id=%d",
		c.apiURL,
		listID,
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

	notes := []wundergo.Note{}
	err = json.NewDecoder(resp.Body).Decode(&notes)
	if err != nil {
		return nil, err
	}
	return notes, nil
}

// NotesForTaskID returns Notes for the provided taskID.
func (c oauthClient) NotesForTaskID(taskID uint) ([]wundergo.Note, error) {
	if taskID == 0 {
		return nil, errors.New("taskID must be > 0")
	}

	url := fmt.Sprintf(
		"%s/notes?task_id=%d",
		c.apiURL,
		taskID,
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

	notes := []wundergo.Note{}
	err = json.NewDecoder(resp.Body).Decode(&notes)
	if err != nil {
		return nil, err
	}
	return notes, nil
}

// Note returns the Note for the corresponding noteID.
func (c oauthClient) Note(noteID uint) (wundergo.Note, error) {
	if noteID == 0 {
		return wundergo.Note{}, errors.New("noteID must be > 0")
	}

	url := fmt.Sprintf(
		"%s/notes/%d",
		c.apiURL,
		noteID,
	)

	req, err := c.newGetRequest(url)
	if err != nil {
		return wundergo.Note{}, err
	}

	resp, err := c.do(req)
	if err != nil {
		return wundergo.Note{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return wundergo.Note{}, fmt.Errorf("Unexpected response code %d - expected %d", resp.StatusCode, http.StatusOK)
	}

	note := wundergo.Note{}
	err = json.NewDecoder(resp.Body).Decode(&note)
	if err != nil {
		return wundergo.Note{}, err
	}
	return note, nil
}

// CreateNote creates a note with the provided content associated with the
// Task for the corresponding taskID.
func (c oauthClient) CreateNote(content string, taskID uint) (wundergo.Note, error) {
	if taskID == 0 {
		return wundergo.Note{}, errors.New("taskID must be > 0")
	}

	body := []byte(fmt.Sprintf(`{"content":"%s","task_id":%d}`, content, taskID))

	url := fmt.Sprintf("%s/notes", c.apiURL)

	req, err := c.newPostRequest(url, body)
	if err != nil {
		return wundergo.Note{}, err
	}

	resp, err := c.do(req)
	if err != nil {
		return wundergo.Note{}, err
	}

	if resp.StatusCode != http.StatusCreated {
		return wundergo.Note{}, fmt.Errorf("Unexpected response code %d - expected %d", resp.StatusCode, http.StatusCreated)
	}

	note := wundergo.Note{}
	err = json.NewDecoder(resp.Body).Decode(&note)
	if err != nil {
		return wundergo.Note{}, err
	}
	return note, nil
}

// UpdateNote updates the provided Note.
// Notes cannot be moved between tasks; note.TaskID is ignored
func (c oauthClient) UpdateNote(note wundergo.Note) (wundergo.Note, error) {
	body, err := json.Marshal(note)
	if err != nil {
		return wundergo.Note{}, err
	}

	url := fmt.Sprintf(
		"%s/notes/%d",
		c.apiURL,
		note.ID,
	)

	req, err := c.newPatchRequest(url, body)
	if err != nil {
		return wundergo.Note{}, err
	}

	resp, err := c.do(req)
	if err != nil {
		return wundergo.Note{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return wundergo.Note{}, fmt.Errorf("Unexpected response code %d - expected %d", resp.StatusCode, http.StatusOK)
	}

	returnedNote := wundergo.Note{}
	err = json.NewDecoder(resp.Body).Decode(&returnedNote)
	if err != nil {
		return wundergo.Note{}, err
	}
	return returnedNote, nil
}

// DeleteNote deletes the provided note.
func (c oauthClient) DeleteNote(note wundergo.Note) error {
	url := fmt.Sprintf(
		"%s/notes/%d?revision=%d",
		c.apiURL,
		note.ID,
		note.Revision,
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
