package ngrok

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptrace"
	"strings"
)

type Debug struct {
	Stdout, Stderr io.Writer
	Verbose        bool
	Include        bool
	DryRun         bool
}

// implements the io.TeeReader logic, but closes the writer when the reader
// completes
type teeReaderCloser struct {
	r    io.Reader
	w    io.WriteCloser
	done <-chan struct{}
}

func (r teeReaderCloser) Read(p []byte) (n int, err error) {
	n, err = r.r.Read(p)
	if n > 0 {
		if n, err := r.w.Write(p[:n]); err != nil {
			return n, err
		}
	}
	if errors.Is(err, io.EOF) {
		r.w.Close()
		<-r.done
	}
	return n, err
}

func (d Debug) makeRequest(ctx context.Context, body io.Reader) (context.Context, io.Reader) {
	if (!d.Verbose && !d.DryRun) || d.Stderr == nil {
		return ctx, body
	}

	ctx = httptrace.WithClientTrace(ctx, &httptrace.ClientTrace{
		WroteHeaderField: func(key string, value []string) {
			fmt.Fprintf(d.Stderr, "> %s: %s\n", key, value)
		},
		WroteHeaders: func() {
			fmt.Fprintln(d.Stderr)
		},
	})

	if body != nil {
		ch := make(chan struct{})

		pr, pw := io.Pipe()
		body = teeReaderCloser{
			r:    body,
			w:    pw,
			done: ch,
		}

		go func() {
			defer close(ch)
			decoder := json.NewDecoder(pr)
			for decoder.More() {
				var i interface{}
				if err := decoder.Decode(&i); err == nil {
					enc := json.NewEncoder(d.Stderr)
					enc.SetIndent("", "  ")
					enc.Encode(i)
				}
			}
			fmt.Fprintln(d.Stderr)
		}()
	}

	return ctx, body
}

func (d Debug) printResponse(r *http.Response) {
	if d.Stderr != nil {
		// status
		if d.Verbose {
			fmt.Fprintf(d.Stderr, "Status: %s\n", r.Status)
		} else {
			fmt.Fprintln(d.Stderr, r.Status)
		}

		// headers
		if d.Verbose || d.Include {
			fmt.Fprintln(d.Stderr)
			for k, v := range r.Header {
				fmt.Fprintf(d.Stderr, "< %s: %v\n", k, v)
			}
			fmt.Fprintln(d.Stderr)
		}
	}

	if d.Stdout != nil && (r.StatusCode < 400 || d.Verbose) {
		body, _ := ioutil.ReadAll(r.Body)
		r.Body.Close()
		r.Body = ioutil.NopCloser(bytes.NewReader(body))

		var i interface{}
		if err := json.Unmarshal(body, &i); err == nil {
			enc := json.NewEncoder(d.Stdout)
			enc.SetIndent("", "  ")
			enc.Encode(i)
		} else {
			fmt.Fprint(d.Stdout, string(body))
		}
	}
}

type eofReader struct{}

func (eofReader) Read([]byte) (int, error) {
	return 0, io.EOF
}

func emptyHTTPResponse(req *http.Request) *http.Response {
	return &http.Response{
		Body:    ioutil.NopCloser(eofReader{}),
		Header:  http.Header{},
		Trailer: http.Header{},
		Request: req,
	}
}

func (d Debug) dryRunResponse(req *http.Request) (*http.Response, error) {
	if d.Verbose && d.Stderr != nil {
		fmt.Fprintf(d.Stderr, "> :authority: [%s]\n", req.URL.Hostname())
		fmt.Fprintf(d.Stderr, "> :method: [%s]\n", req.Method)
		fmt.Fprintf(d.Stderr, "> :path: [%s]\n", req.URL.Path)
		fmt.Fprintf(d.Stderr, "> :scheme: [%s]\n", req.URL.Scheme)
		for k, v := range req.Header {
			fmt.Fprintf(d.Stderr, "> %s: %s\n", strings.ToLower(k), v)
		}
		fmt.Fprintln(d.Stderr)
	}
	if req.Body != nil {
		// if Verbose this causes the request to be printed to stderr
		ioutil.ReadAll(req.Body)
	}
	return emptyHTTPResponse(req), nil
}
