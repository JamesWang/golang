package ds

type SimpleDataStore struct {
	userData map[string]string
}

func (sds SimpleDataStore) UserNameForId(userId string) (string, bool) {
	name, ok := sds.userData[userId]
	return name, ok
}

func NewSimpleDataStore() SimpleDataStore {
	return SimpleDataStore{
		userData: map[string]string{
			"1": "Fred",
			"2": "Mary",
			"3": "Pat",
		},
	}
}

type DataStore interface {
	UserNameForId(userId string) (string, bool)
}


