package ewsutil

import "github.com/hurtuh/ews"

// SendEmail helper method to send Message
func SendEmail(c ews.Client, to []string, subject, body string) error {

	m := ews.Message{
		ItemClass: "IPM.Note",
		Subject:   subject,
		Body: ews.Body{
			BodyType: "Text",
			Body:     []byte(body),
		},
		Sender: ews.OneMailbox{
			Mailbox: ews.Mailbox{
				EmailAddress: c.GetUsername(),
			},
		},
	}
	mb := make([]ews.Mailbox, len(to))
	for i, addr := range to {
		mb[i].EmailAddress = addr
	}
	m.ToRecipients.Mailbox = append(m.ToRecipients.Mailbox, mb...)

	return ews.CreateMessageItem(c, m)
}

func SendEmailWithAttachments(c ews.Client, to []string, subject, body string, fileName string, content string) error {

	m := ews.Message{
		ItemClass: "IPM.Note",
		Subject:   subject,
		Body: ews.Body{
			BodyType: "Text",
			Body:     []byte(body),
		},
		Attachments: ews.Attachments{[]ews.FileAttachment{{
			Name:    fileName,
			Content: content,
		},
		},
		},
		Sender: ews.OneMailbox{
			Mailbox: ews.Mailbox{
				EmailAddress: c.GetUsername(),
			},
		},
	}
	mb := make([]ews.Mailbox, len(to))
	for i, addr := range to {
		mb[i].EmailAddress = addr
	}
	m.ToRecipients.Mailbox = append(m.ToRecipients.Mailbox, mb...)

	return ews.CreateMessageItem(c, m)
}
