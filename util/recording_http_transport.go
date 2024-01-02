package util

import (
	"fmt"
	"github.com/yunarta/terraform-api-transport/transport"
	"os"
	"strings"
)

// Variable stores the count of how many times a method is called.
var count int

// RecordingHttpPayloadTransport struct implementing the PayloadTransport interface.
type RecordingHttpPayloadTransport struct {
	Transport transport.PayloadTransport
}

// Assert that RecordingHttpPayloadTransport struct implements the PayloadTransport interface.
var _ transport.PayloadTransport = &RecordingHttpPayloadTransport{}

// Send method for the RecordingHttpPayloadTransport struct. This sends the request and writes the response to a file.
func (r *RecordingHttpPayloadTransport) Send(request *transport.PayloadRequest) (*transport.PayloadResponse, error) {
	// Send the request using the underlying Transport's Send method.
	reply, err := r.Transport.Send(request)

	// Convert the URL to a safe format.
	safeUrl := strings.Replace(request.Url, "/", "-", -1)
	safeUrl = strings.Replace(safeUrl, "?", "-", -1)
	safeUrl = strings.Replace(safeUrl, "&", "-", -1)
	safeUrl = strings.Replace(safeUrl, "=", "-", -1)

	// If directory records exist, write the response to a file.
	if _, err := os.Stat("records"); !os.IsNotExist(err) {
		name := fmt.Sprintf("records/%d-%s.json", count, safeUrl)
		_ = os.WriteFile(name, []byte(reply.Body), 0644)

		count++
	}

	// Return the response and any error.
	return reply, err

}

// SendWithExpectedStatus method sends the request and expects a certain status. The response is written to a file.
func (r *RecordingHttpPayloadTransport) SendWithExpectedStatus(request *transport.PayloadRequest, expectedStatus ...int) (*transport.PayloadResponse, error) {
	// Send the request using the underlying Transport's SendWithExpectedStatus method.
	reply, err := r.Transport.SendWithExpectedStatus(request, expectedStatus...)

	// Convert the URL to a safe format.
	safeUrl := strings.Replace(request.Url, "/", "-", -1)
	safeUrl = strings.Replace(safeUrl, "?", "-", -1)
	safeUrl = strings.Replace(safeUrl, "&", "-", -1)
	safeUrl = strings.Replace(safeUrl, "=", "-", -1)

	// If directory records exist, write the response to a file.
	if _, err := os.Stat("records"); !os.IsNotExist(err) {
		name := fmt.Sprintf("records/%d-%s.json", count, safeUrl)
		_ = os.WriteFile(name, []byte(reply.Body), 0644)

		count++
	}
	// Return the response and any error.
	return reply, err
}
