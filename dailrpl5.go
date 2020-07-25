package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type DAILRPL5IDataForm struct {
	FormHead CommonFormHead
	FormData DAILRPL5I
}

type DAILRPL5ODataForm struct {
	FormHead CommonFormHead
	FormData DAILRPL5O
}

type DAILRPL5RequestForm struct {
	Form []DAILRPL5IDataForm
}

type DAILRPL5ResponseForm struct {
	Form []DAILRPL5ODataForm
}

type DAILRPL5I struct {
	LoanAcctNo                   string  `json:"LoanAcctNo"`
	AcctnDt                      string  `json:"AcctnDt"`
	Pridnum                      int     `json:"Pridnum"`
	LoanAcctiAcctNo              string  `json:"LoanAcctiAcctNo"`
	ChrgAcctNo                   string  `json:"ChrgAcctNo"`
	AcctNm                       string  `json:"AcctNm"`
	RepayDt                      string  `json:"RepayDt"`
	CurCd                        string  `json:"CurCd"`
	PlanRepayTotlAmt             float64 `json:"PlanRepayTotlAmt"`
	PlanRepayPrin                float64 `json:"PlanRepayPrin"`
	PlanRepayIntr                float64 `json:"PlanRepayIntr"`
	CoreChrgStusCd               string  `json:"CoreChrgStusCd"`
	LoanRepayStusCd              string  `json:"LoanRepayStusCd"`
	FailRsn		                 string  `json:"FailRsn"`
	PageNo                       int     `json:"PageNo"`
	PageRecCount                 int     `json:"PageRecCount"`
}

type DAILRPL5O struct {
	Records                      []DAILRPL5ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type DAILRPL5ORecords struct {
	LoanAcctNo                   string  `json:"LoanAcctNo"`
	AcctnDt                      string  `json:"AcctnDt"`
	Pridnum                      int     `json:"Pridnum"`
	LoanAcctiAcctNo              string  `json:"LoanAcctiAcctNo"`
	ChrgAcctNo                   string  `json:"ChrgAcctNo"`
	AcctNm                       string  `json:"AcctNm"`
	RepayDt                      string  `json:"RepayDt"`
	CurCd                        string  `json:"CurCd"`
	PlanRepayTotlAmt             float64 `json:"PlanRepayTotlAmt"`
	PlanRepayPrin                float64 `json:"PlanRepayPrin"`
	PlanRepayIntr                float64 `json:"PlanRepayIntr"`
	CoreChrgStusCd               string  `json:"CoreChrgStusCd"`
	LoanRepayStusCd              string  `json:"LoanRepayStusCd"`
	FailRsn		                 string  `json:"FailRsn"`
}

// @Desc Build request message
func (o *DAILRPL5RequestForm) PackRequest(dailrpl5I DAILRPL5I) (responseBody []byte, err error) {

	requestForm := DAILRPL5RequestForm{
		Form: []DAILRPL5IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRPL5I",
				},

				FormData: dailrpl5I,
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
func (o *DAILRPL5RequestForm) UnPackRequest(request []byte) (DAILRPL5I, error) {
	dailrpl5I := DAILRPL5I{}
	if err := json.Unmarshal(request, o); nil != err {
		return dailrpl5I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrpl5I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *DAILRPL5ResponseForm) PackResponse(dailrpl5O DAILRPL5O) (responseBody []byte, err error) {
	responseForm := DAILRPL5ResponseForm{
		Form: []DAILRPL5ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRPL5O",
				},
				FormData: dailrpl5O,
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
func (o *DAILRPL5ResponseForm) UnPackResponse(request []byte) (DAILRPL5O, error) {

	dailrpl5O := DAILRPL5O{}

	if err := json.Unmarshal(request, o); nil != err {
		return dailrpl5O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrpl5O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *DAILRPL5I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
