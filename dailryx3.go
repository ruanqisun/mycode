package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type DAILRYX3IDataForm struct {
	FormHead CommonFormHead
	FormData DAILRYX3I
}

type DAILRYX3ODataForm struct {
	FormHead CommonFormHead
	FormData DAILRYX3O
}

type DAILRYX3RequestForm struct {
	Form []DAILRYX3IDataForm
}

type DAILRYX3ResponseForm struct {
	Form []DAILRYX3ODataForm
}

type DAILRYX3I struct {
	LoanAcctNo                     string  `json:"LoanAcctNo"`
	SeqNo                          int     `json:"SeqNo"`
	AcctnDt                        string  `json:"AcctnDt"`
	RepayDt                        string  `json:"RepayDt"`
	Pridnum                        int     `json:"Pridnum"`
	RepayWhenLoanStusCd            string  `json:"RepayWhenLoanStusCd"`
	RgtsintNo                      string  `json:"RgtsintNo"`
	CurrPeriodExecResultCd         string  `json:"CurrPeriodExecResultCd"`
	PrefrTypCd                     string  `json:"PrefrTypCd"`
	CurrPeriodNormlIntr            float64  `json:"CurrPeriodNormlIntr"`
	CurrPeriodPrefrIntr            float64  `json:"CurrPeriodPrefrIntr"`
	CurrPeriodPrefrAftIntr         float64  `json:"CurrPeriodPrefrAftIntr"`
	ActlPrefrIntr                  float64  `json:"ActlPrefrIntr"`
	ExecComnt                      string  `json:"ExecComnt"`
	PageNo                         int     `json:"PageNo"`
	PageRecCount                   int     `json:"PageRecCount"`
}

type DAILRYX3O struct {
	Records                      []DAILRYX3ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type DAILRYX3ORecords struct {
	LoanAcctNo                     string  `json:"LoanAcctNo"`
	SeqNo                          int     `json:"SeqNo"`
	AcctnDt                        string  `json:"AcctnDt"`
	RepayDt                        string  `json:"RepayDt"`
	Pridnum                        int     `json:"Pridnum"`
	RepayWhenLoanStusCd            string  `json:"RepayWhenLoanStusCd"`
	RgtsintNo                      string  `json:"RgtsintNo"`
	CurrPeriodExecResultCd         string  `json:"CurrPeriodExecResultCd"`
	PrefrTypCd                     string  `json:"PrefrTypCd"`
	CurrPeriodNormlIntr            float64  `json:"CurrPeriodNormlIntr"`
	CurrPeriodPrefrIntr            float64  `json:"CurrPeriodPrefrIntr"`
	CurrPeriodPrefrAftIntr         float64  `json:"CurrPeriodPrefrAftIntr"`
	ActlPrefrIntr                  float64  `json:"ActlPrefrIntr"`
	ExecComnt                      string  `json:"ExecComnt"`
}

// @Desc Build request message
func (o *DAILRYX3RequestForm) PackRequest(dailryx3I DAILRYX3I) (responseBody []byte, err error) {

	requestForm := DAILRYX3RequestForm{
		Form: []DAILRYX3IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRYX3I",
				},

				FormData: dailryx3I,
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
func (o *DAILRYX3RequestForm) UnPackRequest(request []byte) (DAILRYX3I, error) {
	dailryx3I := DAILRYX3I{}
	if err := json.Unmarshal(request, o); nil != err {
		return dailryx3I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailryx3I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *DAILRYX3ResponseForm) PackResponse(dailryx3O DAILRYX3O) (responseBody []byte, err error) {
	responseForm := DAILRYX3ResponseForm{
		Form: []DAILRYX3ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRYX3O",
				},
				FormData: dailryx3O,
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
func (o *DAILRYX3ResponseForm) UnPackResponse(request []byte) (DAILRYX3O, error) {

	dailryx3O := DAILRYX3O{}

	if err := json.Unmarshal(request, o); nil != err {
		return dailryx3O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailryx3O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *DAILRYX3I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
