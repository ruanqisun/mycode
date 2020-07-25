package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type IL8R0003IDataForm struct {
	FormHead CommonFormHead
	FormData IL8R0003I
}

type IL8R0003ODataForm struct {
	FormHead CommonFormHead
	FormData IL8R0003O
}

type IL8R0003RequestForm struct {
	Form []IL8R0003IDataForm
}

type IL8R0003ResponseForm struct {
	Form []IL8R0003ODataForm
}

type IL8R0003I struct {
	TxSn                           string  `json:"TxSn"`
	ExecDate                       string  `json:"ExecDate"`
	CustNm                         string  `json:"CustNm"`
	CustNo                         string  `json:"CustNo"`
	CrtfNo                         string  `json:"CrtfNo"`
	IndvCrtfTypCd                  string  `json:"IndvCrtfTypCd"`
	OrgNo                          string  `json:"OrgNo"`
	DubilNo                        string  `json:"DubilNo"`
	CurrLoanRiskClsfCd             string  `json:"CurrLoanRiskClsfCd"`
	DlqcyProb                      float64  `json:"DlqcyProb"`
	RiskRsn                        string  `json:"RiskRsn"`
	SggestMgmtManrCd               string  `json:"SggestMgmtManrCd"`
	AdjustCrdtQta                  float64  `json:"AdjustCrdtQta"`
	SggestExecWindow               string  `json:"SggestExecWindow"`
	TaskStartDt                    string  `json:"TaskStartDt"`
	TaskEndDt                      string  `json:"TaskEndDt"`
	CustNo1                        string  `json:"CustNo1"`
	CustName1                      string  `json:"CustName1"`
	IndvCrtfNo1                    string  `json:"IndvCrtfNo1"`
	IndvCrtfTypCd1                 string  `json:"IndvCrtfTypCd1"`
	LoanDubilNo1                   string  `json:"LoanDubilNo1"`
	AlrdyExecFlg                   string  `json:"AlrdyExecFlg"`
	AtExecWindowFlg                string  `json:"AtExecWindowFlg"`
	ExecTms                        int  `json:"ExecTms"`
	ReachTms                       int  `json:"ReachTms"`
	PageNo                         int     `json:"PageNo"`
	PageRecCount                   int     `json:"PageRecCount"`
}

type IL8R0003O struct {
	Records                      []IL8R0003ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type IL8R0003ORecords struct {
	TxSn                           string  `json:"TxSn"`
	ExecDate                       string  `json:"ExecDate"`
	CustNm                         string  `json:"CustNm"`
	CustNo                         string  `json:"CustNo"`
	CrtfNo                         string  `json:"CrtfNo"`
	IndvCrtfTypCd                  string  `json:"IndvCrtfTypCd"`
	OrgNo                          string  `json:"OrgNo"`
	DubilNo                        string  `json:"DubilNo"`
	CurrLoanRiskClsfCd             string  `json:"CurrLoanRiskClsfCd"`
	DlqcyProb                      float64  `json:"DlqcyProb"`
	RiskRsn                        string  `json:"RiskRsn"`
	SggestMgmtManrCd               string  `json:"SggestMgmtManrCd"`
	AdjustCrdtQta                  float64  `json:"AdjustCrdtQta"`
	SggestExecWindow               string  `json:"SggestExecWindow"`
	TaskStartDt                    string  `json:"TaskStartDt"`
	TaskEndDt                      string  `json:"TaskEndDt"`
	CustNo1                        string  `json:"CustNo1"`
	CustName1                      string  `json:"CustName1"`
	IndvCrtfNo1                    string  `json:"IndvCrtfNo1"`
	IndvCrtfTypCd1                 string  `json:"IndvCrtfTypCd1"`
	LoanDubilNo1                   string  `json:"LoanDubilNo1"`
	AlrdyExecFlg                   string  `json:"AlrdyExecFlg"`
	AtExecWindowFlg                string  `json:"AtExecWindowFlg"`
	ExecTms                        int  `json:"ExecTms"`
	ReachTms                       int  `json:"ReachTms"`
}

// @Desc Build request message
func (o *IL8R0003RequestForm) PackRequest(il8r0003I IL8R0003I) (responseBody []byte, err error) {

	requestForm := IL8R0003RequestForm{
		Form: []IL8R0003IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "IL8R0003I",
				},

				FormData: il8r0003I,
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
func (o *IL8R0003RequestForm) UnPackRequest(request []byte) (IL8R0003I, error) {
	il8r0003I := IL8R0003I{}
	if err := json.Unmarshal(request, o); nil != err {
		return il8r0003I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return il8r0003I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *IL8R0003ResponseForm) PackResponse(il8r0003O IL8R0003O) (responseBody []byte, err error) {
	responseForm := IL8R0003ResponseForm{
		Form: []IL8R0003ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "IL8R0003O",
				},
				FormData: il8r0003O,
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
func (o *IL8R0003ResponseForm) UnPackResponse(request []byte) (IL8R0003O, error) {

	il8r0003O := IL8R0003O{}

	if err := json.Unmarshal(request, o); nil != err {
		return il8r0003O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return il8r0003O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *IL8R0003I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
