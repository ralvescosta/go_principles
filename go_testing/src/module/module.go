package module

import (
	database "study/src/database"
)

type impl struct {
	db *database.IDatabase
}

func (impl *impl) DoSomething(id uint16) *SomethingStruct {
	batata := (*impl.db).GetTable()

	return &SomethingStruct{
		id:   id,
		name: batata,
	}
}

// Module ...
func Module(db database.IDatabase) IModule {
	return &impl{&db}
}
