package check

import (
	"fmt"

	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

type Response struct {
	Title      string
	Errors     []*Response
	Infos      []*Response
	hasErrors  bool
	Details    string
	LevelError string
}

func NewResponse(title string) *Response {
	return &Response{
		Title: title,
	}
}
func NewResponseFromMetadata(meta core.Metadata) *Response {
	title := fmt.Sprintf("%v %v", meta.Name, meta.Namespace)
	return NewResponse(title)
}

func (r *Response) SummarizeErrors(depth int) string {
	msg := ""
	if r.hasErrors {
		indent := spaces(depth)
		msg = fmt.Sprintf("%v%v\n", indent, r.Title)
		if r.LevelError != "" {
			msg += fmt.Sprintf("%v- %v\n", indent, r.LevelError)
		}
		for _, subResp := range r.Errors {
			msg += subResp.SummarizeErrors(depth + 1)
		}
	}
	return msg
}
func (r *Response) SummarizeInfo(depth int) string {
	msg := ""
	indent := spaces(depth)
	result := "ok"
	if r.hasErrors {
		result = "error"
	}
	msg = fmt.Sprintf("%v%v - %v\n", indent, r.Title, result)
	if r.Details != "" {
		msg += fmt.Sprintf("%v- %v\n", indent, r.Details)
	}
	for _, subResp := range r.Infos {
		msg += subResp.SummarizeInfo(depth + 1)
	}
	return msg
}

func (r *Response) String() string {
	// when calling the String method we treat our self as the root
	depth := 0
	info := r.SummarizeInfo(depth)
	str := fmt.Sprintf("info:\n%v\n", info)
	if r.hasErrors {
		errs := r.SummarizeErrors(depth)
		str += fmt.Sprintf("errors:\n%v\n", errs)
	}
	return str
}

func (r *Response) AddResponse(resp *Response) {
	if resp.LevelError != "" {
		r.Errors = append(r.Errors, resp)
		r.hasErrors = true
	} else {
		r.Infos = append(r.Infos, resp)
	}
}
func (r *Response) Check(name string, err error) {
	subResp := NewResponse(name)
	if err != nil {
		subResp.LevelError = err.Error()
	}
	r.AddResponse(subResp)
}

func spaces(n int) string {
	s := ""
	// limit max indentation to 30
	for i := 0; i < n && i < 30; i++ {
		s += " "
	}
	return s
}
