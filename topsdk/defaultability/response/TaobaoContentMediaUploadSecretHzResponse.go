package response

import (
    "topsdk/defaultability/domain"
)

type TaobaoContentMediaUploadSecretHzResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        响应结果
    */
    GetStsUploadTokenResponse  domain.TaobaoContentMediaUploadSecretHzGetStsUploadTokenResponse `json:"get_sts_upload_token_response,omitempty" `
}
