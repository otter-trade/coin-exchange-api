package i18n

const (
	// common msg
	Success                         = 200
	Failed                          = 5000
	UpdateFailed                    = 5001
	CreateFailed                    = 5002
	DeleteFailed                    = 5003
	TargetNotFound                  = 5004
	DatabaseError                   = 5005
	RedisError                      = 5006
	AlreadyInit                     = 5007
	InitRunning                     = 5008
	ConstraintError                 = 5009
	ValidationError                 = 5010
	NotSingularError                = 5011
	ApiRequestFailed                = 5012
	ServerError                     = 5013
	DbError                         = 5014
	VerificationCodeError           = 5015
	PasswordError                   = 5016
	AccountNotExist                 = 5017
	SignError                       = 5018
	NoData                          = 5019
	CopyError                       = 5020
	ParseParamsError                = 5021
	ParamError                      = 5022
	PublicParamError                = 5023
	Unauthorized                    = 5024
	IpNoAccess                      = 5025
	UserNoAccess                    = 5026
	InOfflineState                  = 5027
	NotBelongCurrentGateway         = 5028
	RequestParameterError           = 5029
	ParameterCannotEmpty            = 5030
	InternalSystemError             = 5031
	SystemError                     = 5032
	GatewayServiceError             = 5033
	InterfaceCurrentLimited         = 5034
	InterfaceTemporarilyUnavailable = 5035
	ServiceTemporarilyUnavailable   = 5036
	TooOften                        = 5037
	DisableCalls                    = 5038
	JsonParseError                  = 5039
)

var MapCodeMsg = map[int64]string{
	Success:                         "common.success",
	Failed:                          "common.failed",
	UpdateFailed:                    "common.updateFailed",
	CreateFailed:                    "common.createFailed",
	DeleteFailed:                    "common.deleteFailed",
	TargetNotFound:                  "common.targetNotExist",
	DatabaseError:                   "common.databaseError",
	RedisError:                      "common.redisError",
	AlreadyInit:                     "common.alreadyInit",
	InitRunning:                     "common.initializeIsRunning",
	ConstraintError:                 "common.constraintError",
	ValidationError:                 "common.validationError",
	NotSingularError:                "common.notSingularError",
	ApiRequestFailed:                "common.apiRequestFailed",
	ServerError:                     "common.serverError",
	DbError:                         "common.dbError",
	VerificationCodeError:           "common.verificationCodeError",
	PasswordError:                   "common.passwordError",
	AccountNotExist:                 "common.accountNotExist",
	SignError:                       "common.signError",
	NoData:                          "common.noData",
	CopyError:                       "common.copyError",
	ParseParamsError:                "common.parseParamsError",
	ParamError:                      "common.paramError",
	PublicParamError:                "common.publicParamError",
	Unauthorized:                    "common.unauthorized",
	IpNoAccess:                      "common.ipNoAccess",
	UserNoAccess:                    "common.userNoAccess",
	InOfflineState:                  "common.inOfflineState",
	NotBelongCurrentGateway:         "common.notBelongCurrentGateway",
	RequestParameterError:           "common.requestParameterError",
	ParameterCannotEmpty:            "common.parameterCannotEmpty",
	InternalSystemError:             "common.internalSystemError",
	SystemError:                     "common.systemError",
	GatewayServiceError:             "common.gatewayServiceError",
	InterfaceCurrentLimited:         "common.interfaceCurrentLimited",
	InterfaceTemporarilyUnavailable: "common.interfaceTemporarilyUnavailable",
	ServiceTemporarilyUnavailable:   "common.serviceTemporarilyUnavailable",
	TooOften:                        "common.tooOften",
	DisableCalls:                    "common.disableCalls",
	JsonParseError:                  "common.jsonParseError",
}

func CodeToMsg(code int64) string {
	if msg, ok := MapCodeMsg[code]; ok {
		return msg
	} else { // 没找到, 则返回默认的错误
		return MapCodeMsg[code]
	}
	return ""
}
