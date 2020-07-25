package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type DAILRDH9IDataForm struct {
	FormHead CommonFormHead
	FormData DAILRDH9I
}

type DAILRDH9ODataForm struct {
	FormHead CommonFormHead
	FormData DAILRDH9O
}

type DAILRDH9RequestForm struct {
	Form []DAILRDH9IDataForm
}

type DAILRDH9ResponseForm struct {
	Form []DAILRDH9ODataForm
}

type DAILRDH9I struct {
	CollTaskNo                   string  `json:"CollTaskNo"`
	CurrLoanRiskClsfCd           string  `json:"CurrLoanRiskClsfCd"`
	DlqcyProb                    float64 `json:"DlqcyProb"`
	RiskRsn                      string  `json:"RiskRsn"`
	SggestCollMgmtManrCd         string  `json:"SggestCollMgmtManrCd"`
	AfBnkLnModelExecDt           string  `json:"AfBnkLnModelExecDt"`
	CustName                     string  `json:"CustName"`
	CustNo                       string  `json:"CustNo"`
	IndvCrtfNo                   string  `json:"IndvCrtfNo"`
	IndvCrtfTypCd                string  `json:"IndvCrtfTypCd"`
	LoanDubilNo                  string  `json:"LoanDubilNo"`
	SggestCollExecWindowDays     int     `json:"SggestCollExecWindowDays"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
	PageNo                       int     `json:"PageNo"`
	PageRecCount                 int     `json:"PageRecCount"`
}

type DAILRDH9O struct {
	Records                      []DAILRDH9ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type DAILRDH9ORecords struct {
	CollTaskNo                   string  `json:"CollTaskNo"`
	CurrLoanRiskClsfCd           string  `json:"CurrLoanRiskClsfCd"`
	DlqcyProb                    float64 `json:"DlqcyProb"`
	RiskRsn                      string  `json:"RiskRsn"`
	SggestCollMgmtManrCd         string  `json:"SggestCollMgmtManrCd"`
	AfBnkLnModelExecDt           string  `json:"AfBnkLnModelExecDt"`
	CustName                     string  `json:"CustName"`
	CustNo                       string  `json:"CustNo"`
	IndvCrtfNo                   string  `json:"IndvCrtfNo"`
	IndvCrtfTypCd                string  `json:"IndvCrtfTypCd"`
	LoanDubilNo                  string  `json:"LoanDubilNo"`
	SggestCollExecWindowDays     int     `json:"SggestCollExecWindowDays"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
}

// @Desc Build request message
func (o *DAILRDH9RequestForm) PackRequest(dailrdh9I DAILRDH9I) (responseBody []byte, err error) {

	requestForm := DAILRDH9RequestForm{
		Form: []DAILRDH9IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRDH9I",
				},

				FormData: dailrdh9I,
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
func (o *DAILRDH9RequestForm) UnPackRequest(request []byte) (DAILRDH9I, error) {
	dailrdh9I := DAILRDH9I{}
	if err := json.Unmarshal(request, o); nil != err {
		return dailrdh9I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrdh9I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *DAILRDH9ResponseForm) PackResponse(dailrdh9O DAILRDH9O) (responseBody []byte, err error) {
	responseForm := DAILRDH9ResponseForm{
		Form: []DAILRDH9ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRDH9O",
				},
				FormData: dailrdh9O,
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
func (o *DAILRDH9ResponseForm) UnPackResponse(request []byte) (DAILRDH9O, error) {

	dailrdh9O := DAILRDH9O{}

	if err := json.Unmarshal(request, o); nil != err {
		return dailrdh9O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrdh9O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *DAILRDH9I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
