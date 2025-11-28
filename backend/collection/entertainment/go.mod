module app/entDbase

go 1.24.4

replace nyg/dataset => ../dataset

replace app/updateE => ./update

require app/updateE v0.0.0-00010101000000-000000000000

require (
	app/entertainmentbook v0.0.0-00010101000000-000000000000 // indirect
	nyg/dataset v0.0.0-00010101000000-000000000000 // indirect
)

replace app/book => ./update/./books

replace app/entertainmentbook => ./update/./books
