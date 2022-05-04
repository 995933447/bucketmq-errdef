package errdef

var (
	FileCorruptionErr = NewErr(ErrCodeFileCorruption, "file corruption.")
	FileIsNotDirErr = NewErr(ErrCodeIsNotDir, "file is not dir.")
	FileIsNotRegularFileErr = NewErr(ErrCodeIsNotRegularFile, "file is not regular file.")
	FileSeqNotFoundErr = NewErr(ErrCodeFileSeqNotFound, "file sequence not found.")
	MadeStructNotByNewFuncErr = NewErr(ErrCodeMadeStructNotByNewFunc, "target struct must to make by new(New)Func.")
)
