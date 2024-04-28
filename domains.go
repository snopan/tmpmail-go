package tmpmailgo

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var domains Domains

// SetDomains manually sets the domains to the provided list of strings
// there's no verification whether the provided domains are valid
// this is given so during development you can reduce calls to fetch
// domains if you already know what the valid ones are at the time
func SetDomains(d []string) {
	domains = d
}

// initDomains fetches domains from 1secmail if it's empty which
// is usually at the start. That's why methods that rely on domains
// should call this first always
func initDomains() error {
	if len(domains) > 0 {
		return nil
	}

	resp, err := http.Get(host1SecMail + "/?action=getDomainList")
	if err != nil {
		return fmt.Errorf("failed to fetch domains: %w", err)
	}

	var respDomains []string
	err = json.NewDecoder(resp.Body).Decode(&respDomains)
	if err != nil {
		return fmt.Errorf("failed to parse domains: %w", err)
	}
	if len(respDomains) == 0 {
		return fmt.Errorf("got empty domains, something wrong from 1sec mail: %w", err)
	}

	domains = Domains(respDomains)
	return nil
}
