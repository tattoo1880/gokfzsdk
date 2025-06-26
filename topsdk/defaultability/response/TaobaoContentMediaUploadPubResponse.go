package response

import (
    "topsdk/defaultability/domain"
)

type TaobaoContentMediaUploadPubResponse struct {

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
    PubOssFileToMultiMediaResponse  domain.TaobaoContentMediaUploadPubPubOssFileToMultiMediaResponse `json:"pub_oss_file_to_multi_media_response,omitempty" `
}
