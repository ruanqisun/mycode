package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type DAILR002IDataForm struct {
	FormHead CommonFormHead
	FormData DAILR002I
}

type DAILR002ODataForm struct {
	FormHead CommonFormHead
	FormData DAILR002O
}

type DAILR002RequestForm struct {
	Form []DAILR002IDataForm
}

type DAILR002ResponseForm struct {
	Form []DAILR002ODataForm
}

type DAILR002I struct {
	TxSn                           string  `json:"TxSn"`
	ExecDate                       string  `json:"ExecDate"`
	CustNm                         string  `json:"CustNm"`
	CustNo                         string  `json:"CustNo"`
	CrtfNo                         string  `json:"CrtfNo"`
	IndvCrtfTypCd                  string  `json:"IndvCrtfTypCd"`
	OprorEmpnbr                    string  `json:"OprorEmpnbr"`
	OprOrgNo                       string  `json:"OprOrgNo"`
	OrgNo                          string  `json:"OrgNo"`
	OrgNm                          string  `json:"OrgNm"`
	DubilNo                        string  `json:"DubilNo"`
	CurrLoanRiskClsfCd             string  `json:"CurrLoanRiskClsfCd"`
	DlqcyProb                      float64  `json:"DlqcyProb"`
	RiskRsn                        string  `json:"RiskRsn"`
	SggestMgmtManrCd               string  `json:"SggestMgmtManrCd"`
	AdjustCrdtQta                  float64  `json:"AdjustCrdtQta"`
	AdjustType                     string  `json:"AdjustType"`
	SggestExecWindow               string  `json:"SggestExecWindow"`
	WarnPushManrCd                 string  `json:"WarnPushManrCd"`
	TaskStartDt                    string  `json:"TaskStartDt"`
	TaskEndDt                      string  `json:"TaskEndDt"`
	CustNo1                        string  `json:"CustNo1"`
	CustName1                      string  `json:"CustName1"`
	IndvCrtfNo1                    string  `json:"IndvCrtfNo1"`
	IndvCrtfTypCd1                 string  `json:"IndvCrtfTypCd1"`
	LoanDubilNo1                   string  `json:"LoanDubilNo1"`
	AlrdyExecFlg                   string  `json:"AlrdyExecFlg"`
	AtExecWindowFlg                string  `json:"AtExecWindowFlg"`
	ExecTms                        int     `json:"ExecTms"`
	PageNo                         int     `json:"PageNo"`
	PageRecCount                   int     `json:"PageRecCount"`
}

type DAILR002O struct {
	Records                      []DAILR002ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type DAILR002ORecords struct {
	TxSn                           string  `json:"TxSn"`
	ExecDate                       string  `json:"ExecDate"`
	CustNm                         string  `json:"CustNm"`
	CustNo                         string  `json:"CustNo"`
	CrtfNo                         string  `json:"CrtfNo"`
	IndvCrtfTypCd                  string  `json:"IndvCrtfTypCd"`
	OprorEmpnbr                    string  `json:"OprorEmpnbr"`
	OprOrgNo                       string  `json:"OprOrgNo"`
	OrgNo                          string  `json:"OrgNo"`
	OrgNm                          string  `json:"OrgNm"`
	DubilNo                        string  `json:"DubilNo"`
	CurrLoanRiskClsfCd             string  `json:"CurrLoanRiskClsfCd"`
	DlqcyProb                      float64  `json:"DlqcyProb"`
	RiskRsn                        string  `json:"RiskRsn"`
	SggestMgmtManrCd               string  `json:"SggestMgmtManrCd"`
	AdjustCrdtQta                  float64  `json:"AdjustCrdtQta"`
	AdjustType                     string  `json:"AdjustType"`
	SggestExecWindow               string  `json:"SggestExecWindow"`
	WarnPushManrCd                 string  `json:"WarnPushManrCd"`
	TaskStartDt                    string  `json:"TaskStartDt"`
	TaskEndDt                      string  `json:"TaskEndDt"`
	CustNo1                        string  `json:"CustNo1"`
	CustName1                      string  `json:"CustName1"`
	IndvCrtfNo1                    string  `json:"IndvCrtfNo1"`
	IndvCrtfTypCd1                 string  `json:"IndvCrtfTypCd1"`
	LoanDubilNo1                   string  `json:"LoanDubilNo1"`
	AlrdyExecFlg                   string  `json:"AlrdyExecFlg"`
	AtExecWindowFlg                string  `json:"AtExecWindowFlg"`
	ExecTms                        int     `json:"ExecTms"`
}

// @Desc Build request message
func (o *DAILR002RequestForm) PackRequest(dailr002I DAILR002I) (responseBody []byte, err error) {

	requestForm := DAILR002RequestForm{
		Form: []DAILR002IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILR002I",
				},

				FormData: dailr002I,
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
func (o *DAILR002RequestForm) UnPackRequest(request []byte) (DAILR002I, error) {
	dailr002I := DAILR002I{}
	if err := json.Unmarshal(request, o); nil != err {
		return dailr002I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailr002I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *DAILR002ResponseForm) PackResponse(dailr002O DAILR002O) (responseBody []byte, err error) {
	responseForm := DAILR002ResponseForm{
		Form: []DAILR002ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILR002O",
				},
				FormData: dailr002O,
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
func (o *DAILR002ResponseForm) UnPackResponse(request []byte) (DAILR002O, error) {

	dailr002O := DAILR002O{}

	if err := json.Unmarshal(request, o); nil != err {
		return dailr002O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailr002O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *DAILR002I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
