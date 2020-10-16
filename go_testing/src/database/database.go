package database

type impl struct{}

func (*impl) GetTable() string {
	return "First Table"
}

// Database ...
func Database() IDatabase {
	return &impl{}
}
