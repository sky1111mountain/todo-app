package myerrors

type ErrCode string

const (
	Unknown ErrCode = "U000"

	InsertFailed     ErrCode = "R001"
	NoRowsAffected   ErrCode = "R002"
	GetTaskFailed    ErrCode = "R003"
	ScanFailed       ErrCode = "R004"
	GetListfailed    ErrCode = "R005"
	BadColumn        ErrCode = "R006"
	TaskUpdateFailed ErrCode = "R007"
	TaskDeleteFailed ErrCode = "R008"
	UnChangeRows     ErrCode = "R009"
	NotFound         ErrCode = "R010"

	InsertDataFailed ErrCode = "S001"
	GetDataFailed    ErrCode = "S002"
	NAData           ErrCode = "S003"
	UpdateDataFailed ErrCode = "S004"
	DeleteDataFailed ErrCode = "S005"
	BadRequest       ErrCode = "S006"
	UnUpdatedTask    ErrCode = "S007"
	NoUpdateColumn   ErrCode = "S008"

	ReqBodyDecodeFailed  ErrCode = "C001"
	BadPath              ErrCode = "C002"
	BadQuery             ErrCode = "C003"
	ResponseEncodeFailed ErrCode = "C004"

	RequiredAuthorizationHeader ErrCode = "A001"
	CannotMakeValidator         ErrCode = "A002"
	Unauthorizated              ErrCode = "A003"
	NotMatchUserName            ErrCode = "A004"
	FailedValidate              ErrCode = "A005"
	FailedLoadEnv               ErrCode = "A006"
)
