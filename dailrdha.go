package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type DAILRDHAIDataForm struct {
	FormHead CommonFormHead
	FormData DAILRDHAI
}

type DAILRDHAODataForm struct {
	FormHead CommonFormHead
	FormData DAILRDHAO
}

type DAILRDHARequestForm struct {
	Form []DAILRDHAIDataForm
}

type DAILRDHAResponseForm struct {
	Form []DAILRDHAODataForm
}

type DAILRDHAI struct {
	WarnTaskNo                   string  `json:"WarnTaskNo"`
	LoanDubilNo                  string  `json:"LoanDubilNo"`
	CustNo                       string  `json:"CustNo"`
	CustName                     string  `json:"CustName"`
	IndvCrtfTypCd                string  `json:"IndvCrtfTypCd"`
	IndvCrtfNo                   string  `json:"IndvCrtfNo"`
	RiskRsn                      string  `json:"RiskRsn"`
	SggestWarnSignalTypCd        string  `json:"SggestWarnSignalTypCd"`
	SggestExecWindowDays         int     `json:"SggestExecWindowDays"`
	AdjAftCrdtQta                float64 `json:"AdjAftCrdtQta"`
	AdjDrctCd                    string  `json:"AdjDrctCd"`
	InlnModelExecDt              string  `json:"InlnModelExecDt"`
	DlqcyProb                    float64 `json:"DlqcyProb"`
	CurrRiskGradeCd              string  `json:"CurrRiskGradeCd"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
	PageNo                       int     `json:"PageNo"`
	PageRecCount                 int     `json:"PageRecCount"`
}

type DAILRDHAO struct {
	Records                      []DAILRDHAORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type DAILRDHAORecords struct {
	WarnTaskNo                   string  `json:"WarnTaskNo"`
	LoanDubilNo                  string  `json:"LoanDubilNo"`
	CustNo                       string  `json:"CustNo"`
	CustName                     string  `json:"CustName"`
	IndvCrtfTypCd                string  `json:"IndvCrtfTypCd"`
	IndvCrtfNo                   string  `json:"IndvCrtfNo"`
	RiskRsn                      string  `json:"RiskRsn"`
	SggestWarnSignalTypCd        string  `json:"SggestWarnSignalTypCd"`
	SggestExecWindowDays         int     `json:"SggestExecWindowDays"`
	AdjAftCrdtQta                float64 `json:"AdjAftCrdtQta"`
	AdjDrctCd                    string  `json:"AdjDrctCd"`
	InlnModelExecDt              string  `json:"InlnModelExecDt"`
	DlqcyProb                    float64 `json:"DlqcyProb"`
	CurrRiskGradeCd              string  `json:"CurrRiskGradeCd"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
}

// @Desc Build request message
func (o *DAILRDHARequestForm) PackRequest(dailrdhaI DAILRDHAI) (responseBody []byte, err error) {

	requestForm := DAILRDHARequestForm{
		Form: []DAILRDHAIDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRDHAI",
				},

				FormData: dailrdhaI,
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
func (o *DAILRDHARequestForm) UnPackRequest(request []byte) (DAILRDHAI, error) {
	dailrdhaI := DAILRDHAI{}
	if err := json.Unmarshal(request, o); nil != err {
		return dailrdhaI, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrdhaI, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *DAILRDHAResponseForm) PackResponse(dailrdhaO DAILRDHAO) (responseBody []byte, err error) {
	responseForm := DAILRDHAResponseForm{
		Form: []DAILRDHAODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRDHAO",
				},
				FormData: dailrdhaO,
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
func (o *DAILRDHAResponseForm) UnPackResponse(request []byte) (DAILRDHAO, error) {

	dailrdhaO := DAILRDHAO{}

	if err := json.Unmarshal(request, o); nil != err {
		return dailrdhaO, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrdhaO, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *DAILRDHAI) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
