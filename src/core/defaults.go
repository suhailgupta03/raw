package core

const (
	// Disk Defaults
	DefaultDirectoryPermission = 0777
	DefaultFilePermission      = 0777

	// Disk Errors
	RecordAlreadyExists = "record already exists"

	// Writer Defaults
	PrimaryKeyConstraintViolated = "duplicate primary key not allowed"

	// Writer Errors
	PrimaryKeyTypeExtractionError = "primary key must be one of string,int,float64"
)
