package core

type TemplateReader interface {
	GetFrom() (string, error)
	GetTo() (string, error)
	GetSubject() (string, error)
	GetMIMEType() (string, error)
	GetBody() (Body, error)
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
