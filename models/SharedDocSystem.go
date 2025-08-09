package models

type SharedDocsSystem struct {
	users     map[string]*User
	documents Document
	dataDir   string
}

func NewSharedDocsSystem(users map[string]*User) *SharedDocsSystem {

	return &SharedDocsSystem{
		users: users,
		documents: Document{
			Name: "document.txt",
		},
	}
}
