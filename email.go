package tmpmailgo

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"golang.org/x/exp/slices"
)

// NewEmail generates and returns a random email
func NewEmail() (Email, error) {
	if err := initDomains(); err != nil {
		return "", fmt.Errorf("failed to init domains: %w", err)
	}

	username := randomString(emailUsernameLength)
	domain := randomValue(domains)
	email := fmt.Sprintf("%s@%s", username, domain)

	return Email(email), nil
}

// CustomEmail verifies the provided email is valid before returnning that email
func CustomEmail(email string) (Email, error) {
	if err := initDomains(); err != nil {
		return "", fmt.Errorf("failed to init domains: %w", err)
	}

	emailParts := strings.Split(email, "@")
	if len(emailParts) != 2 {
		return "", errorInvalidEmail
	}

	domain := emailParts[1]
	if !slices.Contains(domains, domain) {
		return "", errorInvalidDomain
	}

	return Email(email), nil
}

// GetInbox returns all the messages in the inbox currently
func (e Email) GetInbox() ([]MessageSummary, error) {
	var messages []MessageSummary

	resp, err := http.Get(fmt.Sprintf("%s/?action=getMessages&domain=%s", host1SecMail, e))
	if err != nil {
		return messages, fmt.Errorf("failed to fetch inbox: %w", err)
	}

	err = json.NewDecoder(resp.Body).Decode(&messages)
	if err != nil {
		return messages, fmt.Errorf("failed to parse inbox messages: %w", err)
	}

	return messages, nil
}

// ReadMessage returns the full detail of a message when the message id is provided
func (e Email) ReadMessage(id int) (Message, error) {
	var message Message

	resp, err := http.Get(fmt.Sprintf("%s/?action=readMessage&domain=%s&id=%d", host1SecMail, e, id))
	if err != nil {
		return message, fmt.Errorf("failed to fetch message: %w", err)
	}

	err = json.NewDecoder(resp.Body).Decode(&message)
	if err != nil {
		return message, fmt.Errorf("failed to parse message: %w", err)
	}

	return message, nil
}

// randomString generates a random string with the provided length
func randomString(length int) string {
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	b := make([]byte, length+2)
	r.Read(b)
	return fmt.Sprintf("%x", b)[2 : length+2]
}

// randomValue gets a random value from a slice
func randomValue[T any](array []T) T {
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	return array[r.Intn(len(array))]
}
