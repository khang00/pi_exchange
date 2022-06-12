package core

import (
	"errors"
	"strings"
)

type PlaceHolderName = string
type PlaceHolderValue = string
type DataObject = map[PlaceHolderName]PlaceHolderValue

type Mail struct {
	From     string
	To       string
	Subject  string
	MIMEType string
	Body     string
}

type Template struct {
	From     string
	To       string
	Subject  string
	MIMEType string
	Body     Body
	OperationPlaceHolder
}

type Body struct {
	placeHoldersLookup map[string]int64
	Content            []Part
}

type PartType = string

const PartTypeHolder PartType = "holder"
const PartTypeContent PartType = "content"

type Part struct {
	PartType    PartType
	PartContent string
}

type PlaceHolder struct {
	start int
	end   int
	name  string
}

func NewTemplate(reader TemplateReader) (*Template, error) {
	from := reader.GetFrom()
	to := reader.GetTo()
	subject := reader.GetSubject()
	mimeType := reader.GetMIMEType()
	body := reader.GetBody()

	parts := make([]Part, 0)
	placeHolderLookup := make(map[string]int64, 0)

	for i := 0; i < len(body); {
		holderStartIndex := strings.Index(body[i:], "{{")
		if holderStartIndex != -1 {
			holderEndIndex := strings.Index(body[holderStartIndex:], "}}")

			textPart := Part{
				PartType:    PartTypeContent,
				PartContent: body[i : holderStartIndex-1],
			}
			parts = append(parts, textPart)

			placeHolderPart := Part{
				PartType:    PartTypeHolder,
				PartContent: body[holderStartIndex+2 : holderEndIndex-1],
			}
			parts = append(parts, placeHolderPart)

			i = holderEndIndex + 2
		} else {
			i = len(body)
		}
	}

	return &Template{
		From:     from,
		To:       to,
		Subject:  subject,
		MIMEType: mimeType,
		Body: Body{
			placeHoldersLookup: placeHolderLookup,
			Content:            parts,
		},
		OperationPlaceHolder: NewOperationPlaceHolder(),
	}, nil
}

func (t *Template) stream(reader DataReader, writer MailWriter, errHandler MailErrorHandler) error {
	for !reader.IsEmpty() {
		obj, err := reader.GetNextObject()
		if err != nil {
			return err
		}

		mail, err := t.parseEmail(obj)
		if err != nil {
			err = errHandler.handle(obj)
			if err != nil {
				return err
			}
		}

		err = writer.WriteNextMail(mail)
		if err != nil {
			return err
		}
	}

	return nil
}

func (t Template) parseEmail(obj DataObject) (Mail, error) {
	var body strings.Builder
	for _, part := range t.Body.Content {
		if part.PartType == PartTypeContent {
			body.WriteString(part.PartContent)
		} else if part.PartType == PartTypeHolder {
			value, ok := t.getValueForPlaceHolder(part.PartContent)
			if ok {
				body.WriteString(value)
			} else {
				placeHolderValue := obj[part.PartContent]
				body.WriteString(placeHolderValue)
			}

		} else {
			return Mail{}, errors.New("part have wrong type")
		}
	}

	return Mail{
		From:     t.From,
		To:       t.To,
		Subject:  t.Subject,
		MIMEType: t.MIMEType,
		Body:     body.String(),
	}, nil
}
