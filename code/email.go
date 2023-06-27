package mailtrap

import "errors"

// Creates a new customer object.
type EmailParams struct {
	From *Params
	To []*Params
	Cc []*Params
	Bcc []*Params
	Attachments *AttachmentParams
	Headers *HeaderParams
	CustomVariables *CustomerVariableParam
	Subject *string
	Text *string
	Html *string
	Category *string

}


type Params struct {
	Email *string
	Name *string
}

type AttachmentParams struct {
	Content *string
	Type *string
	Filename *string
	Disposition *string
	ContentID *string
}

type HeaderParams struct{
	//
}

type CustomerVariableParam struct{
	//
}

type Response struct {
	Success *bool
	MessageIds []*string
}

func (mg *MailtrapImpl) Send(params *EmailParams) (mes string, err error) {

	if mg.apiKey == "" {
		err = errors.New("you must provide a valid api-key before calling Send()")
		return
	}

	return "test", nil;

}