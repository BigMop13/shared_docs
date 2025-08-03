package models

type SharedDocsSystem struct {
	users     map[string]*User
	documents Document
	dataDir   string
}

func NewSharedDocsSystem(users map[string]*User) *SharedDocsSystem {

	return &SharedDocsSystem{
		users: make(map[string]*User),
		documents: Document{
			ID: "document.txt",
		},
		dataDir: dataDir,
	}
}
