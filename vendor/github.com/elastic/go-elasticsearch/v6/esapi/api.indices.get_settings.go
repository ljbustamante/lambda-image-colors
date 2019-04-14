// Code generated from specification version 6.7.2: DO NOT EDIT

package esapi

import (
	"context"
	"strconv"
	"strings"
	"time"
)

func newIndicesGetSettingsFunc(t Transport) IndicesGetSettings {
	return func(o ...func(*IndicesGetSettingsRequest)) (*Response, error) {
		var r = IndicesGetSettingsRequest{}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

// IndicesGetSettings returns settings for one or more indices.
//
// See full documentation at http://www.elastic.co/guide/en/elasticsearch/reference/master/indices-get-settings.html.
//
type IndicesGetSettings func(o ...func(*IndicesGetSettingsRequest)) (*Response, error)

// IndicesGetSettingsRequest configures the Indices  Get Settings API request.
//
type IndicesGetSettingsRequest struct {
	Index []string

	Name []string

	AllowNoIndices    *bool
	ExpandWildcards   string
	FlatSettings      *bool
	IgnoreUnavailable *bool
	IncludeDefaults   *bool
	Local             *bool
	MasterTimeout     time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r IndicesGetSettingsRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "GET"

	path.Grow(1 + len(strings.Join(r.Index, ",")) + 1 + len("_settings") + 1 + len(strings.Join(r.Name, ",")))
	if len(r.Index) > 0 {
		path.WriteString("/")
		path.WriteString(strings.Join(r.Index, ","))
	}
	path.WriteString("/")
	path.WriteString("_settings")
	if len(r.Name) > 0 {
		path.WriteString("/")
		path.WriteString(strings.Join(r.Name, ","))
	}

	params = make(map[string]string)

	if r.AllowNoIndices != nil {
		params["allow_no_indices"] = strconv.FormatBool(*r.AllowNoIndices)
	}

	if r.ExpandWildcards != "" {
		params["expand_wildcards"] = r.ExpandWildcards
	}

	if r.FlatSettings != nil {
		params["flat_settings"] = strconv.FormatBool(*r.FlatSettings)
	}

	if r.IgnoreUnavailable != nil {
		params["ignore_unavailable"] = strconv.FormatBool(*r.IgnoreUnavailable)
	}

	if r.IncludeDefaults != nil {
		params["include_defaults"] = strconv.FormatBool(*r.IncludeDefaults)
	}

	if r.Local != nil {
		params["local"] = strconv.FormatBool(*r.Local)
	}

	if r.MasterTimeout != 0 {
		params["master_timeout"] = formatDuration(r.MasterTimeout)
	}

	if r.Pretty {
		params["pretty"] = "true"
	}

	if r.Human {
		params["human"] = "true"
	}

	if r.ErrorTrace {
		params["error_trace"] = "true"
	}

	if len(r.FilterPath) > 0 {
		params["filter_path"] = strings.Join(r.FilterPath, ",")
	}

	req, _ := newRequest(method, path.String(), nil)

	if len(params) > 0 {
		q := req.URL.Query()
		for k, v := range params {
			q.Set(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	if ctx != nil {
		req = req.WithContext(ctx)
	}

	res, err := transport.Perform(req)
	if err != nil {
		return nil, err
	}

	response := Response{
		StatusCode: res.StatusCode,
		Body:       res.Body,
		Header:     res.Header,
	}

	return &response, nil
}

// WithContext sets the request context.
//
func (f IndicesGetSettings) WithContext(v context.Context) func(*IndicesGetSettingsRequest) {
	return func(r *IndicesGetSettingsRequest) {
		r.ctx = v
	}
}

// WithIndex - a list of index names; use _all to perform the operation on all indices.
//
func (f IndicesGetSettings) WithIndex(v ...string) func(*IndicesGetSettingsRequest) {
	return func(r *IndicesGetSettingsRequest) {
		r.Index = v
	}
}

// WithName - the name of the settings that should be included.
//
func (f IndicesGetSettings) WithName(v ...string) func(*IndicesGetSettingsRequest) {
	return func(r *IndicesGetSettingsRequest) {
		r.Name = v
	}
}

// WithAllowNoIndices - whether to ignore if a wildcard indices expression resolves into no concrete indices. (this includes `_all` string or when no indices have been specified).
//
func (f IndicesGetSettings) WithAllowNoIndices(v bool) func(*IndicesGetSettingsRequest) {
	return func(r *IndicesGetSettingsRequest) {
		r.AllowNoIndices = &v
	}
}

// WithExpandWildcards - whether to expand wildcard expression to concrete indices that are open, closed or both..
//
func (f IndicesGetSettings) WithExpandWildcards(v string) func(*IndicesGetSettingsRequest) {
	return func(r *IndicesGetSettingsRequest) {
		r.ExpandWildcards = v
	}
}

// WithFlatSettings - return settings in flat format (default: false).
//
func (f IndicesGetSettings) WithFlatSettings(v bool) func(*IndicesGetSettingsRequest) {
	return func(r *IndicesGetSettingsRequest) {
		r.FlatSettings = &v
	}
}

// WithIgnoreUnavailable - whether specified concrete indices should be ignored when unavailable (missing or closed).
//
func (f IndicesGetSettings) WithIgnoreUnavailable(v bool) func(*IndicesGetSettingsRequest) {
	return func(r *IndicesGetSettingsRequest) {
		r.IgnoreUnavailable = &v
	}
}

// WithIncludeDefaults - whether to return all default setting for each of the indices..
//
func (f IndicesGetSettings) WithIncludeDefaults(v bool) func(*IndicesGetSettingsRequest) {
	return func(r *IndicesGetSettingsRequest) {
		r.IncludeDefaults = &v
	}
}

// WithLocal - return local information, do not retrieve the state from master node (default: false).
//
func (f IndicesGetSettings) WithLocal(v bool) func(*IndicesGetSettingsRequest) {
	return func(r *IndicesGetSettingsRequest) {
		r.Local = &v
	}
}

// WithMasterTimeout - specify timeout for connection to master.
//
func (f IndicesGetSettings) WithMasterTimeout(v time.Duration) func(*IndicesGetSettingsRequest) {
	return func(r *IndicesGetSettingsRequest) {
		r.MasterTimeout = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f IndicesGetSettings) WithPretty() func(*IndicesGetSettingsRequest) {
	return func(r *IndicesGetSettingsRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f IndicesGetSettings) WithHuman() func(*IndicesGetSettingsRequest) {
	return func(r *IndicesGetSettingsRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f IndicesGetSettings) WithErrorTrace() func(*IndicesGetSettingsRequest) {
	return func(r *IndicesGetSettingsRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f IndicesGetSettings) WithFilterPath(v ...string) func(*IndicesGetSettingsRequest) {
	return func(r *IndicesGetSettingsRequest) {
		r.FilterPath = v
	}
}