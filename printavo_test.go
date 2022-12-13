package printavo

import "testing"

func TestNewClient(t *testing.T) {
	c, err := NewClient("test@example.com", "1234abcd")

	if err != nil {
		t.Fatalf("Failed to create the client: %v", err)
	}

	expectedBaseURL := baseURL + apiVersionPath

	if c.BaseURL().String() != expectedBaseURL {
		t.Errorf("NewClient BaseURL is %s, expected %s", c.BaseURL().String(), expectedBaseURL)
	}

	if c.UserAgent != userAgent {
		t.Errorf("NewClient UserAgent is %s, expected %s", c.UserAgent, userAgent)
	}
}
