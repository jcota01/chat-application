module server

go 1.22rc1

replace communicate => ../communicate

require (
	communicate v0.0.0-00010101000000-000000000000
	github.com/google/uuid v1.5.0
)
