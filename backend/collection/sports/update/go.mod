module app/updateS

go 1.24.4

require (
	app/sportsbook v0.0.0-00010101000000-000000000000
	nyg/dataset v0.0.0-00010101000000-000000000000
)

replace app/dataset => ../dataset

replace nyg/dataset => ../../dataset

replace app/book => ./books

replace app/sports => ./books

replace app/sportsbook => ./books
