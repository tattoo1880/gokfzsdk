package response

import (
)

type AlibabaProductMerchantproductPublishGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        发布规则
    */
    Product  string `json:"product,omitempty" `
}
