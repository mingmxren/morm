package morm

type PersistentObject interface {
	TableName() string
}

type PersistentObjectWithPrimaryKey interface {
	PersistentObject
	PrimaryKey() []interface{}
}
