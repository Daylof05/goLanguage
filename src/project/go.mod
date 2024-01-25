module project

go 1.21.6

replace project/filesmanagement => ../filesmanagement

replace project/foldersmanagement => ../foldersmanagement

require (
	project/filesmanagement v0.0.0-00010101000000-000000000000
	project/foldersmanagement v0.0.0-00010101000000-000000000000
)
