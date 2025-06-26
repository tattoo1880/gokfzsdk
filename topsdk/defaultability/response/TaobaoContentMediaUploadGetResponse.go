package response

import (
    "topsdk/defaultability/domain"
)

type TaobaoContentMediaUploadGetResponse struct {

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
    GetMediaUploadResultResponse  domain.TaobaoContentMediaUploadGetGetMediaUploadResultResponse `json:"get_media_upload_result_response,omitempty" `
}
