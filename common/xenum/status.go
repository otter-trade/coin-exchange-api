package xenum

const (
	ShowStatus   = 1
	HideStatus   = 2
	DeleteStatus = 3
)

const (
	//1，默认值，2，提交审核，4，审核未通过，8，审核成功
	DefaultStatus = 1 //默认值
	AuditStatus   = 2 //提交审核
	FailStatus    = 4 //审核未通过
	SuccessStatus = 8 //审核成功
)
