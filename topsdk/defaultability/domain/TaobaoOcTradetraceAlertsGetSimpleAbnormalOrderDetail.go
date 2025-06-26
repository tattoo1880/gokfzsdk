package domain


type TaobaoOcTradetraceAlertsGetSimpleAbnormalOrderDetail struct {
    /*
        订单ID     */
    Tid  *int64 `json:"tid,omitempty" `

}

func (s *TaobaoOcTradetraceAlertsGetSimpleAbnormalOrderDetail) SetTid(v int64) *TaobaoOcTradetraceAlertsGetSimpleAbnormalOrderDetail {
    s.Tid = &v
    return s
}
