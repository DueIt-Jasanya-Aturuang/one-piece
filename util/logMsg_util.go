package util

const (
	LogErrDBConn              = "Failed Open Database | err : %v"
	LogErrDBConnClose         = "Failed close Database | err : %v"
	LogErrBeginTx             = "Failed Start Transaction | err : %v"
	LogErrRollback            = "failed rollback data | err : %v"
	LogInfoRollback           = "rollback data | err : %v"
	LogErrCommit              = "failed commit data | err %v"
	LogErrPrepareContextClose = "failed close prepared context | err : %v"
	LogErrPrepareContext      = "failed open prepared context | err : %v"
	LogErrExecContext         = "failed exec context | err : %v"
	LogErrQueryRowContextScan = "failed scan data | err : %v"
	LogErrQueryRows           = "failed Rows data | err : %v"
	LogErrQueryRowsScan       = "failed scan data | err : %v"
)

const (
	LogErrOpenFile  = "failed open file | file : %v | err : %v"
	LogErrCloseFile = "failed close file | file : %v | err : %v"
)

const (
	LogErrPutObjectMinio   = "failed put object minio | err : %v"
	LogErrDelObjectMinio   = "failed delete object minio | err : %v"
	LogInfoFileUploadMinio = "file upload info | msg : %v"
)
