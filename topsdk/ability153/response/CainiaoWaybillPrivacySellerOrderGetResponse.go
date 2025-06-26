package response

import (
    "topsdk/ability153/domain"
)

type CainiaoWaybillPrivacySellerOrderGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        错误列表
    */
    ErrorCodeList  []string `json:"error_code_list,omitempty" `
    /*
        第一个错误
    */
    OneErrorInfo  string `json:"one_error_info,omitempty" `
    /*
        是否失败
    */
    Failure  bool `json:"failure,omitempty" `
    /*
        返回值
    */
    ResponseList  []domain.CainiaoWaybillPrivacySellerOrderGetModule `json:"response_list,omitempty" `
    /*
        错误信息
    */
    ErrorInfoList  []string `json:"error_info_list,omitempty" `
    /*
        objectId
    */
    ObjectId  string `json:"object_id,omitempty" `
}
