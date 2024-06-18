package formatter

type Status string

var (
	Success             Status = "00"
	CacheError          Status = "HV01"
	DatabaseError       Status = "HV02"
	InvalidRequest      Status = "HV03"
	DataNotFound        Status = "HV04"
	InternalServerError Status = "HV05"
	DataConflict        Status = "HV06"
	Unauthorized        Status = "HV07"
)

func (s Status) String() string {
	return string(s)
}
