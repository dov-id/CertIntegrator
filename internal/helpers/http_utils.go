package helpers

import (
	"bytes"
	"context"
	"io"
	"net/http"

	"github.com/dov-id/cert-integrator-svc/internal/data"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func MakeHttpRequest(ctx context.Context, params data.RequestParams) (*data.ResponseParams, error) {
	req, err := http.NewRequest(params.Method, params.Link, bytes.NewReader(params.Body))
	if err != nil {
		return nil, errors.Wrap(err, "couldn't create request")
	}

	ctx, cancel := context.WithTimeout(ctx, params.Timeout)
	defer cancel()
	req = req.WithContext(ctx)

	if params.Header != nil {
		for key, value := range params.Header {
			req.Header.Set(key, value)
		}
	}

	if params.Query != nil {
		q := req.URL.Query()
		for key, value := range params.Query {
			q.Add(key, value)
		}
		req.URL.RawQuery = q.Encode()
	}

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "error making http request")
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, errors.Wrap(err, "error reading response body")
	}
	clearBody := io.NopCloser(bytes.NewReader(body))

	return &data.ResponseParams{
		Body:       clearBody,
		Header:     response.Header,
		StatusCode: response.StatusCode,
	}, nil
}
