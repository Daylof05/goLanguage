module project

go 1.21.6

replace project/filesmanagement => ../filesmanagement

replace project/foldersmanagement => ../foldersmanagement

replace project/sql => ../sql

require (
	project/filesmanagement v0.0.0-00010101000000-000000000000
	project/foldersmanagement v0.0.0-00010101000000-000000000000
	project/sql v0.0.0-00010101000000-000000000000
)

require github.com/go-sql-driver/mysql v1.7.1 // indirect
