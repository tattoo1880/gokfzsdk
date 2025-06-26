package response

import (
    "topsdk/ability316/domain"
)

type TaobaoWirelessPictureCheckResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        检查结果
    */
    CheckResults  []domain.TaobaoWirelessPictureCheckCheckpoints `json:"check_results,omitempty" `
    /*
        综合结果建议。建议用户执行的操作，取值范围： pass：文本正常； review：需要人工审核； block：文本违规，可以直接删除或者做限制处理
    */
    Suggestion  string `json:"suggestion,omitempty" `
}
