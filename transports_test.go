// Code generated for API Clients. DO NOT EDIT.

package ngrok_test

import (
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/http/httputil"
)

type debugTransport struct {
	rt  http.RoundTripper
	out io.Writer
}

func (t *debugTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// dump request
	rtid := rand.Int()
	fmt.Fprintf(t.out, "\n************** ROUND TRIP %x ********************\n", rtid)
	reqBytes, _ := httputil.DumpRequestOut(req, true)
	t.out.Write(reqBytes)
	fmt.Fprintln(t.out, "")

	// execute round trip
	resp, err := t.rt.RoundTrip(req)

	// dump response
	if err != nil {
		fmt.Fprintf(t.out, "RoundTrip error: %v\n", err)
	} else {
		respBytes, _ := httputil.DumpResponse(resp, true)
		t.out.Write(respBytes)
		fmt.Fprintln(t.out, "")
	}
	fmt.Fprintf(t.out, "\n************** END ROUND TRIP %x ********************\n", rtid)

	return resp, err
}

type mockTransport struct {
	mockResponse      []*http.Response
	mockResponseError []error
}

func (s *mockTransport) SetResponse(statusCode int, body string) {
	s.mockResponse = append(s.mockResponse, &http.Response{
		Status:     http.StatusText(statusCode),
		StatusCode: statusCode,
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		Header: http.Header{
			"Content-Length": []string{fmt.Sprintf("%d", len(body))},
			"Connection":     []string{"close"},
		},
		ContentLength: int64(len(body)),
		Body:          io.NopCloser(bytes.NewReader([]byte(body))),
	})
	s.mockResponseError = append(s.mockResponseError, nil)
}

func (s *mockTransport) SetResponseError(err error) {
	s.mockResponse = append(s.mockResponse, nil)
	s.mockResponseError = append(s.mockResponseError, err)
}

func (t *mockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	resp := t.mockResponse[0]
	respErr := t.mockResponseError[0]
	t.mockResponse = t.mockResponse[1:]
	t.mockResponseError = t.mockResponseError[1:]
	return resp, respErr
}
