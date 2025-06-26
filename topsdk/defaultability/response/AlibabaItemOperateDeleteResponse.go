package response

import (
)

type AlibabaItemOperateDeleteResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        商品删除是否成功
    */
    Result  string `json:"result,omitempty" `
}
