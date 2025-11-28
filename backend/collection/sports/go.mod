module app/spoDbase

go 1.24.5

replace app/updateS => ./update

require app/updateS v0.0.0-00010101000000-000000000000

require (
	app/sportsbook v0.0.0-00010101000000-000000000000 // indirect
	nyg/dataset v0.0.0-00010101000000-000000000000 // indirect
)

replace app/sportsbook => ./update/./books

replace nyg/dataset => ../dataset
