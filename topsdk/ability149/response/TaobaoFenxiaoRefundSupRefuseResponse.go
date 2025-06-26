package response

import (
)

type TaobaoFenxiaoRefundSupRefuseResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        执行失败的错误码
    */
    BizFailedCode  string `json:"biz_failed_code,omitempty" `
    /*
        执行失败的错误原因
    */
    BizFailedMsg  string `json:"biz_failed_msg,omitempty" `
    /*
        执行是否成功
    */
    BizSuccess  bool `json:"biz_success,omitempty" `
}
