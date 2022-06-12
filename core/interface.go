package core

type TemplateReader interface {
	GetFrom() string
	GetTo() string
	GetSubject() string
	GetMIMEType() string
	GetBody() string
}

type DataReader interface {
	IsEmpty() bool
	GetNextObject() (DataObject, error)
}

type MailWriter interface {
	WriteNextMail(mail Mail) error
}

type MailErrorHandler interface {
	handle(mail DataObject) error
}
