package util

import (
	"github.com/yunarta/terraform-api-transport/transport"
	"os"
	"testing"
)

type fakePayloadTransport struct{}

func (f *fakePayloadTransport) Send(request *transport.PayloadRequest) (*transport.PayloadResponse, error) {
	return &transport.PayloadResponse{StatusCode: 200, Body: "{}"}, nil
}

func (f *fakePayloadTransport) SendWithExpectedStatus(request *transport.PayloadRequest, expectedStatus ...int) (*transport.PayloadResponse, error) {
	return &transport.PayloadResponse{StatusCode: 200, Body: "{}"}, nil
}

func TestRecordingHttpPayloadTransport_Send(t *testing.T) {
	r := &RecordingHttpPayloadTransport{Transport: &fakePayloadTransport{}}
	dir := "records"
	_ = os.Mkdir(dir, 0777)

	req := &transport.PayloadRequest{Url: "test/url"}
	_, _ = r.Send(req)

	files, _ := os.ReadDir(dir)
	if len(files) != 1 {
		t.Errorf("Expected file not created")
	}
	_ = os.RemoveAll(dir)
}

func TestRecordingHttpPayloadTransport_SendWithExpectedStatus(t *testing.T) {
	r := &RecordingHttpPayloadTransport{Transport: &fakePayloadTransport{}}
	dir := "records"
	_ = os.Mkdir(dir, 0777)

	req := &transport.PayloadRequest{Url: "test/url"}
	_, _ = r.SendWithExpectedStatus(req, 200)

	files, _ := os.ReadDir("records")
	if len(files) != 1 {
		t.Errorf("Expected file not created")
	}
	_ = os.RemoveAll("records")
}
