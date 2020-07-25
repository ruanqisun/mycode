package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type DAILRSX0IDataForm struct {
	FormHead CommonFormHead
	FormData DAILRSX0I
}

type DAILRSX0ODataForm struct {
	FormHead CommonFormHead
	FormData DAILRSX0O
}

type DAILRSX0RequestForm struct {
	Form []DAILRSX0IDataForm
}

type DAILRSX0ResponseForm struct {
	Form []DAILRSX0ODataForm
}

type DAILRSX0I struct {
	ArtgclInvstgTaskNo           string  `json:"ArtgclInvstgTaskNo"`
	SeqNo                        int     `json:"SeqNo"`
	ReacCustTypCd                string  `json:"ReacCustTypCd"`
	ReacCustTms                  int     `json:"ReacCustTms"`
	KeprcdStusCd                 string  `json:"KeprcdStusCd"`
	PutawyGatherResultCd         string  `json:"PutawyGatherResultCd"`
	CotactTelNo                  string  `json:"CotactTelNo"`
	CurrAddr                     string  `json:"CurrAddr"`
	Lgtd                         float64 `json:"Lgtd"`
	LgtdDrctCd                   string  `json:"LgtdDrctCd"`
	Lttd                         float64 `json:"Lttd"`
	LttdDrctCd                   string  `json:"LttdDrctCd"`
	RsdnceCmprmntCd              string  `json:"RsdnceCmprmntCd"`
	RsdnceDtlAddr                string  `json:"RsdnceDtlAddr"`
	RsdnceAddr                   string  `json:"RsdnceAddr"`
	CotactCmprmntCd              string  `json:"CotactCmprmntCd"`
	CmunicDtlAddr                string  `json:"CmunicDtlAddr"`
	CmunicAddr                   string  `json:"CmunicAddr"`
	WechatNo                     string  `json:"WechatNo"`
	QqNo                         string  `json:"QqNo"`
	IndusTypCd                   string  `json:"IndusTypCd"`
	CareerTypCd                  string  `json:"CareerTypCd"`
	WorkUnitNm                   string  `json:"WorkUnitNm"`
	DutyCd                       string  `json:"DutyCd"`
	AlrdyVrfyMonIncomeAmt        float64 `json:"AlrdyVrfyMonIncomeAmt"`
	MonIncomeVrfyMthdCd          string  `json:"MonIncomeVrfyMthdCd"`
	AlrdyVrfyAsstAmt             float64 `json:"AlrdyVrfyAsstAmt"`
	AsstVrfyMthdCd               string  `json:"AsstVrfyMthdCd"`
	CmtdInfo                     string  `json:"CmtdInfo"`
	SponsorCustMgrEmpnbr         string  `json:"SponsorCustMgrEmpnbr"`
	SponsorOrgNo                 string  `json:"SponsorOrgNo"`
	CorgCustMgrEmpnbr            string  `json:"CorgCustMgrEmpnbr"`
	CorgOrgNo                    string  `json:"CorgOrgNo"`
	DlwthDt                      string  `json:"DlwthDt"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
	PageNo                       int     `json:"PageNo"`
	PageRecCount                 int     `json:"PageRecCount"`
}

type DAILRSX0O struct {
	Records                      []DAILRSX0ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type DAILRSX0ORecords struct {
	ArtgclInvstgTaskNo           string  `json:"ArtgclInvstgTaskNo"`
	SeqNo                        int     `json:"SeqNo"`
	ReacCustTypCd                string  `json:"ReacCustTypCd"`
	ReacCustTms                  int     `json:"ReacCustTms"`
	KeprcdStusCd                 string  `json:"KeprcdStusCd"`
	PutawyGatherResultCd         string  `json:"PutawyGatherResultCd"`
	CotactTelNo                  string  `json:"CotactTelNo"`
	CurrAddr                     string  `json:"CurrAddr"`
	Lgtd                         float64 `json:"Lgtd"`
	LgtdDrctCd                   string  `json:"LgtdDrctCd"`
	Lttd                         float64 `json:"Lttd"`
	LttdDrctCd                   string  `json:"LttdDrctCd"`
	RsdnceCmprmntCd              string  `json:"RsdnceCmprmntCd"`
	RsdnceDtlAddr                string  `json:"RsdnceDtlAddr"`
	RsdnceAddr                   string  `json:"RsdnceAddr"`
	CotactCmprmntCd              string  `json:"CotactCmprmntCd"`
	CmunicDtlAddr                string  `json:"CmunicDtlAddr"`
	CmunicAddr                   string  `json:"CmunicAddr"`
	WechatNo                     string  `json:"WechatNo"`
	QqNo                         string  `json:"QqNo"`
	IndusTypCd                   string  `json:"IndusTypCd"`
	CareerTypCd                  string  `json:"CareerTypCd"`
	WorkUnitNm                   string  `json:"WorkUnitNm"`
	DutyCd                       string  `json:"DutyCd"`
	AlrdyVrfyMonIncomeAmt        float64 `json:"AlrdyVrfyMonIncomeAmt"`
	MonIncomeVrfyMthdCd          string  `json:"MonIncomeVrfyMthdCd"`
	AlrdyVrfyAsstAmt             float64 `json:"AlrdyVrfyAsstAmt"`
	AsstVrfyMthdCd               string  `json:"AsstVrfyMthdCd"`
	CmtdInfo                     string  `json:"CmtdInfo"`
	SponsorCustMgrEmpnbr         string  `json:"SponsorCustMgrEmpnbr"`
	SponsorOrgNo                 string  `json:"SponsorOrgNo"`
	CorgCustMgrEmpnbr            string  `json:"CorgCustMgrEmpnbr"`
	CorgOrgNo                    string  `json:"CorgOrgNo"`
	DlwthDt                      string  `json:"DlwthDt"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
}

// @Desc Build request message
func (o *DAILRSX0RequestForm) PackRequest(dailrsx0I DAILRSX0I) (responseBody []byte, err error) {

	requestForm := DAILRSX0RequestForm{
		Form: []DAILRSX0IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRSX0I",
				},

				FormData: dailrsx0I,
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
func (o *DAILRSX0RequestForm) UnPackRequest(request []byte) (DAILRSX0I, error) {
	dailrsx0I := DAILRSX0I{}
	if err := json.Unmarshal(request, o); nil != err {
		return dailrsx0I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrsx0I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *DAILRSX0ResponseForm) PackResponse(dailrsx0O DAILRSX0O) (responseBody []byte, err error) {
	responseForm := DAILRSX0ResponseForm{
		Form: []DAILRSX0ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRSX0O",
				},
				FormData: dailrsx0O,
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
func (o *DAILRSX0ResponseForm) UnPackResponse(request []byte) (DAILRSX0O, error) {

	dailrsx0O := DAILRSX0O{}

	if err := json.Unmarshal(request, o); nil != err {
		return dailrsx0O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrsx0O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *DAILRSX0I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
