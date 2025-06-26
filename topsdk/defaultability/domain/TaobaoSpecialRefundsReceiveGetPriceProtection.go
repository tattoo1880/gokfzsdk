package domain

import (
        "topsdk/util"
    )

type TaobaoSpecialRefundsReceiveGetPriceProtection struct {
    /*
        商家出资金额。精确到2位小数;单位:元。如:200.07，表示:200元7分     */
    SellerRefundFee  *string `json:"seller_refund_fee,omitempty" `

    /*
        平台出资金额。精确到2位小数;单位:元。如:200.07，表示:200元7分     */
    PlatformRefundFee  *string `json:"platform_refund_fee,omitempty" `

    /*
        商家出资状态。1表示无需创建，2表示待创建，3表示进行中，4表示成功，5表示失败，6表示关闭。只有4是已成功出资。     */
    SellerRefundStatus  *int64 `json:"seller_refund_status,omitempty" `

    /*
        平台出资状态。1表示无需创建，2表示待创建，3表示进行中，4表示成功，5表示失败，6表示关闭。只有4是已成功出资。     */
    PlatformRefundStatus  *int64 `json:"platform_refund_status,omitempty" `

    /*
        价保场景。"pp"表示价保，"mgbp"表示买贵必赔，"pc"表示价差赔付。     */
    Scene  *string `json:"scene,omitempty" `

    /*
        价保单id，标识一次价保     */
    PriceProtectionId  *int64 `json:"price_protection_id,omitempty" `

    /*
        价保申请时间     */
    CreatedTime  *util.LocalTime `json:"created_time,omitempty" `

}

func (s *TaobaoSpecialRefundsReceiveGetPriceProtection) SetSellerRefundFee(v string) *TaobaoSpecialRefundsReceiveGetPriceProtection {
    s.SellerRefundFee = &v
    return s
}
func (s *TaobaoSpecialRefundsReceiveGetPriceProtection) SetPlatformRefundFee(v string) *TaobaoSpecialRefundsReceiveGetPriceProtection {
    s.PlatformRefundFee = &v
    return s
}
func (s *TaobaoSpecialRefundsReceiveGetPriceProtection) SetSellerRefundStatus(v int64) *TaobaoSpecialRefundsReceiveGetPriceProtection {
    s.SellerRefundStatus = &v
    return s
}
func (s *TaobaoSpecialRefundsReceiveGetPriceProtection) SetPlatformRefundStatus(v int64) *TaobaoSpecialRefundsReceiveGetPriceProtection {
    s.PlatformRefundStatus = &v
    return s
}
func (s *TaobaoSpecialRefundsReceiveGetPriceProtection) SetScene(v string) *TaobaoSpecialRefundsReceiveGetPriceProtection {
    s.Scene = &v
    return s
}
func (s *TaobaoSpecialRefundsReceiveGetPriceProtection) SetPriceProtectionId(v int64) *TaobaoSpecialRefundsReceiveGetPriceProtection {
    s.PriceProtectionId = &v
    return s
}
func (s *TaobaoSpecialRefundsReceiveGetPriceProtection) SetCreatedTime(v util.LocalTime) *TaobaoSpecialRefundsReceiveGetPriceProtection {
    s.CreatedTime = &v
    return s
}
