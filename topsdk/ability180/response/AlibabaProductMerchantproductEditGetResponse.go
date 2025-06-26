package response

import (
)

type AlibabaProductMerchantproductEditGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        编辑规则
    */
    Product  string `json:"product,omitempty" `
}
