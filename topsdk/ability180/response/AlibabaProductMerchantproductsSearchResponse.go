package response

import (
)

type AlibabaProductMerchantproductsSearchResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        搜索结果
    */
    Products  string `json:"products,omitempty" `
}
