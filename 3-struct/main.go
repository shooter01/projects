package main

type bin struct {
	id        string
	private   bool
	createdAt string
	name      string
}

func newBin(id string, private bool, createdAt string, name string) *bin {
	return &bin{
		id:        id,
		private:   private,
		createdAt: createdAt,
		name:      name,
	}
}
