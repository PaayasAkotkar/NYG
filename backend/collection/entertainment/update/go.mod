module app/updateE

go 1.24.4

replace nyg/dataset => ../../dataset

require (
	app/entertainmentbook v0.0.0-00010101000000-000000000000
	nyg/dataset v0.0.0-00010101000000-000000000000
)

replace app/updateS => ./update

replace app/book => ./books

replace app/entertainmentbook => ./books
