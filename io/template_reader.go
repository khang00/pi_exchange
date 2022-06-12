package io

import (
	"encoding/json"
	"os"
)

type JsonTemplate struct {
	From     string `json:"from"`
	To       string `json:"to"`
	Subject  string `json:"subject"`
	MIMEType string `json:"MIMEType"`
	Body     string `json:"body"`
}

type JsonTemplateReader struct {
	template JsonTemplate
}

func NewJsonTemplateReader(file string) (JsonTemplateReader, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return JsonTemplateReader{}, nil
	}

	var template JsonTemplate
	if err = json.Unmarshal(data, &template); err != nil {
		return JsonTemplateReader{}, err
	}

	return JsonTemplateReader{
		template: template,
	}, nil
}
func (r *JsonTemplateReader) GetFrom() string {
	return r.template.From
}

func (r *JsonTemplateReader) GetTo() string {
	return r.template.To
}

func (r *JsonTemplateReader) GetSubject() string {
	return r.template.Subject
}

func (r *JsonTemplateReader) GetMIMEType() string {
	return r.template.MIMEType
}

func (r *JsonTemplateReader) GetBody() string {
	return r.template.Body
}
