package models

type Contact struct {
	Name  string
	Email string
}

type Contacts = []Contact

func NewContact(name, email string) Contact {
	return Contact{Name: name, Email: email}
}

type FormData struct {
	Values map[string]string
	Errors map[string]string
}

func NewFormData() FormData {
	return FormData{Values: make(map[string]string), Errors: make(map[string]string)}
}

type Data struct {
	Contacts Contacts
}

func newData() Data {
	return Data{
		Contacts: []Contact{
			NewContact("Joe Doe", "joe@email.com"),
			NewContact("Clara Doe", "clara@email.com"),
		},
	}
}

func (d *Data) HasEmail(email string) bool {
	for _, contact := range d.Contacts {
		if contact.Email == email {
			return true
		}
	}
	return false
}

type Page struct {
	Data     Data
	FormData FormData
}

func NewPage() Page {
	return Page{Data: newData(), FormData: NewFormData()}
}
