package util

const (
	LogErrDBConn              = "Failed Open Database | err : %v"
	LogErrDBConnClose         = "Failed close Database | err : %v"
	LogErrBeginTx             = "Failed Start Transaction | err : %v"
	LogErrRollback            = "failed rollback data | err : %v"
	LogErrCommit              = "failed commit data | err %v"
	LogErrPrepareContextClose = "failed close prepared context | err : %v"
	LogErrExecContext         = "failed exec context | err : %v"
)
