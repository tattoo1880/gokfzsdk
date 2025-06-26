package response

import (
)

type AlibabaProductMerchantproductPublishResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        产品信息
    */
    ProductInfo  string `json:"product_info,omitempty" `
}
