module nyg/validate

go 1.24.2

replace nyg/dataset => ../dataset

require (
	nyg/dataset v0.0.0-00010101000000-000000000000
	resty.dev/v3 v3.0.0-beta.3
)

require golang.org/x/net v0.33.0 // indirect
