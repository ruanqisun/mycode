package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type DAILRDHBIDataForm struct {
	FormHead CommonFormHead
	FormData DAILRDHBI
}

type DAILRDHBODataForm struct {
	FormHead CommonFormHead
	FormData DAILRDHBO
}

type DAILRDHBRequestForm struct {
	Form []DAILRDHBIDataForm
}

type DAILRDHBResponseForm struct {
	Form []DAILRDHBODataForm
}

type DAILRDHBI struct {
	TaskDlwthResultNo            string  `json:"TaskDlwthResultNo"`
	WarnTaskNo                   string  `json:"WarnTaskNo"`
	SeqNo                        int     `json:"SeqNo"`
	CustNo                       string  `json:"CustNo"`
	HndlrEmpnbr                  string  `json:"HndlrEmpnbr"`
	ReacCustManrCd               string  `json:"ReacCustManrCd"`
	KeprcdStusCd                 string  `json:"KeprcdStusCd"`
	CorgDlwthResultCd            string  `json:"CorgDlwthResultCd"`
	CollTms                      int     `json:"CollTms"`
	WarnSignalDlwthResultCd      string  `json:"WarnSignalDlwthResultCd"`
	WarnDlwthRemrk               string  `json:"WarnDlwthRemrk"`
	CotactTelNo                  string  `json:"CotactTelNo"`
	CurrAddr                     string  `json:"CurrAddr"`
	Lgtd                         float64 `json:"Lgtd"`
	LgtdDrctCd                   string  `json:"LgtdDrctCd"`
	Lttd                         float64 `json:"Lttd"`
	LttdDrctCd                   string  `json:"LttdDrctCd"`
	UpdAftRsdnceScAdcmCd         string  `json:"UpdAftRsdnceScAdcmCd"`
	UpdAftRsdnceDtlAddr          string  `json:"UpdAftRsdnceDtlAddr"`
	UpdAftRsdnceAddr             string  `json:"UpdAftRsdnceAddr"`
	UpdAftCmunicAddrScAdcmCd     string  `json:"UpdAftCmunicAddrScAdcmCd"`
	UpdAftCmunicDtlAddr          string  `json:"UpdAftCmunicDtlAddr"`
	UpdAftCmunicAddr             string  `json:"UpdAftCmunicAddr"`
	UpdGrndPostCd                string  `json:"UpdGrndPostCd"`
	OneslfCmtdInfo               string  `json:"OneslfCmtdInfo"`
	MrgncyConterName             string  `json:"MrgncyConterName"`
	PtyRelKindCd                 string  `json:"PtyRelKindCd"`
	MrgncyConterTelNo            string  `json:"MrgncyConterTelNo"`
	SponsorEmpnbr                string  `json:"SponsorEmpnbr"`
	SponsorOrgNo                 string  `json:"SponsorOrgNo"`
	CorgEmpnbr                   string  `json:"CorgEmpnbr"`
	CorgOrgNo                    string  `json:"CorgOrgNo"`
	DlwthDt                      string  `json:"DlwthDt"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
	PageNo                       int     `json:"PageNo"`
	PageRecCount                 int     `json:"PageRecCount"`
}

type DAILRDHBO struct {
	Records                      []DAILRDHBORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type DAILRDHBORecords struct {
	TaskDlwthResultNo            string  `json:"TaskDlwthResultNo"`
	WarnTaskNo                   string  `json:"WarnTaskNo"`
	SeqNo                        int     `json:"SeqNo"`
	CustNo                       string  `json:"CustNo"`
	HndlrEmpnbr                  string  `json:"HndlrEmpnbr"`
	ReacCustManrCd               string  `json:"ReacCustManrCd"`
	KeprcdStusCd                 string  `json:"KeprcdStusCd"`
	CorgDlwthResultCd            string  `json:"CorgDlwthResultCd"`
	CollTms                      int     `json:"CollTms"`
	WarnSignalDlwthResultCd      string  `json:"WarnSignalDlwthResultCd"`
	WarnDlwthRemrk               string  `json:"WarnDlwthRemrk"`
	CotactTelNo                  string  `json:"CotactTelNo"`
	CurrAddr                     string  `json:"CurrAddr"`
	Lgtd                         float64 `json:"Lgtd"`
	LgtdDrctCd                   string  `json:"LgtdDrctCd"`
	Lttd                         float64 `json:"Lttd"`
	LttdDrctCd                   string  `json:"LttdDrctCd"`
	UpdAftRsdnceScAdcmCd         string  `json:"UpdAftRsdnceScAdcmCd"`
	UpdAftRsdnceDtlAddr          string  `json:"UpdAftRsdnceDtlAddr"`
	UpdAftRsdnceAddr             string  `json:"UpdAftRsdnceAddr"`
	UpdAftCmunicAddrScAdcmCd     string  `json:"UpdAftCmunicAddrScAdcmCd"`
	UpdAftCmunicDtlAddr          string  `json:"UpdAftCmunicDtlAddr"`
	UpdAftCmunicAddr             string  `json:"UpdAftCmunicAddr"`
	UpdGrndPostCd                string  `json:"UpdGrndPostCd"`
	OneslfCmtdInfo               string  `json:"OneslfCmtdInfo"`
	MrgncyConterName             string  `json:"MrgncyConterName"`
	PtyRelKindCd                 string  `json:"PtyRelKindCd"`
	MrgncyConterTelNo            string  `json:"MrgncyConterTelNo"`
	SponsorEmpnbr                string  `json:"SponsorEmpnbr"`
	SponsorOrgNo                 string  `json:"SponsorOrgNo"`
	CorgEmpnbr                   string  `json:"CorgEmpnbr"`
	CorgOrgNo                    string  `json:"CorgOrgNo"`
	DlwthDt                      string  `json:"DlwthDt"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
}

// @Desc Build request message
func (o *DAILRDHBRequestForm) PackRequest(dailrdhbI DAILRDHBI) (responseBody []byte, err error) {

	requestForm := DAILRDHBRequestForm{
		Form: []DAILRDHBIDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRDHBI",
				},

				FormData: dailrdhbI,
			},
		},
	}

	responseBody, err = json.Marshal(requestForm)
	if err != nil {
		return nil, errors.Wrap(err, 0, constant.REQPACKERR)
	}

	return responseBody, nil
}

// @Desc Parsing request message
func (o *DAILRDHBRequestForm) UnPackRequest(request []byte) (DAILRDHBI, error) {
	dailrdhbI := DAILRDHBI{}
	if err := json.Unmarshal(request, o); nil != err {
		return dailrdhbI, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrdhbI, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *DAILRDHBResponseForm) PackResponse(dailrdhbO DAILRDHBO) (responseBody []byte, err error) {
	responseForm := DAILRDHBResponseForm{
		Form: []DAILRDHBODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRDHBO",
				},
				FormData: dailrdhbO,
			},
		},
	}

	responseBody, err = json.Marshal(responseForm)
	if err != nil {
		return nil, errors.Wrap(err, 0, constant.RSPPACKERR)
	}

	return responseBody, nil
}

// @Desc Parsing response message
func (o *DAILRDHBResponseForm) UnPackResponse(request []byte) (DAILRDHBO, error) {

	dailrdhbO := DAILRDHBO{}

	if err := json.Unmarshal(request, o); nil != err {
		return dailrdhbO, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrdhbO, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *DAILRDHBI) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
