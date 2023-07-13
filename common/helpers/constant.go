package helpers

import "errors"

const DEFAULT int64 = 1    // 默认值
const WAIT int64 = 2       // 等待状态
const REMOVE int64 = 3     // 删除状态
const FAILURE int64 = 4    // 失败状态
const PROCESS int64 = 6    // 处理中状态
const SUCCEED int64 = 8    // 成功状态
const FINISH int64 = 9     // 最后状态
const CANCEL int64 = 10    // 取消状态
const REFUNDING int64 = 11 // 退款中状态
const REFUND int64 = 12    // 已退款状态
const FORBIDDEN int64 = 16 // 禁用状态
const RECOMMEND int64 = 32 // 推荐状态
const ISTOP int64 = 64     // 置顶状态
const PAYIN int64 = 1      // 收入
const PAYOUT int64 = 2     // 支出

const (
	SubAction = "sub" // 订阅

	UnsubAction = "unsub" // 取消订阅
)

// ws连接错误
var (
	CONNECTION_IS_CLOSED = errors.New("CONNECTION_IS_CLOSED")
)
