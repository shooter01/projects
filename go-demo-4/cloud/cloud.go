package cloud

type CloudDb struct {
	url string
}

func NewClodDb(url string) *CloudDb {
	return &CloudDb{
		url: url,
	}
}

func (db *CloudDb) Write(content string) {

}

func (db *CloudDb) Read() ([]byte, error) {
	return []byte{}, nil
}