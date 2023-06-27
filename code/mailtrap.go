package mailtrap

import "net/http"


const (
	// Base Url the library uses to contact mailgun. Use SetAPIBase() to override
	APIBase	= "https://send.api.mailtrap.io/api"
)


// NewMailtrap creates a new client instance.
func NewMailtrap(apiKey string) *MailtrapImpl {
	return &MailtrapImpl{
		apiBase: APIBase,
		apiKey:  apiKey,
		client:  http.DefaultClient,
	}
}

// MailgunImpl bundles data needed by a large number of methods in order to interact with the Mailgun API.
// Colloquially, we refer to instances of this structure as "clients."
type MailtrapImpl struct {
	apiBase            string
	domain             string
	apiKey             string
	client             *http.Client
	baseURL            string
	overrideHeaders    map[string]string
	capturedCurlOutput string
}