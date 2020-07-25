package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type DAILRPL2IDataForm struct {
	FormHead CommonFormHead
	FormData DAILRPL2I
}

type DAILRPL2ODataForm struct {
	FormHead CommonFormHead
	FormData DAILRPL2O
}

type DAILRPL2RequestForm struct {
	Form []DAILRPL2IDataForm
}

type DAILRPL2ResponseForm struct {
	Form []DAILRPL2ODataForm
}

type DAILRPL2I struct {
	LoanAcctNo                    string  `json:"LoanAcctNo"`
	AcctnDt                       string  `json:"AcctnDt"`
	PridnumSeqNo                  int     `json:"PridnumSeqNo"`
	AcctiAcctNo                   string  `json:"AcctiAcctNo"`
	RepayDt                       string  `json:"RepayDt"`
	CurCd                         string  `json:"CurCd"`
	ActlRepayTotlAmt              float64 `json:"ActlRepayTotlAmt"`
	ActlRepayPrin                 float64 `json:"ActlRepayPrin"`
	ActlRepayIntr                 float64 `json:"ActlRepayIntr"`
	ActlRepayOvdueIntr            float64 `json:"ActlRepayOvdueIntr"`
	ActlRepayCmpdAmt              float64 `json:"ActlRepayCmpdAmt"`
	CurrPeriodUnStillTotlAmt      float64 `json:"CurrPeriodUnStillTotlAmt"`
	CurrPeriodUnStillPrin         float64 `json:"CurrPeriodUnStillPrin"`
	CurrPeriodUnStillIntr         float64 `json:"CurrPeriodUnStillIntr"`
	CurrPeriodUnStillOvdueIntr    float64 `json:"CurrPeriodUnStillOvdueIntr"`
	CurrPeriodUnStillCmpdAmt      float64 `json:"CurrPeriodUnStillCmpdAmt"`
	CoreChrgStusCd                string  `json:"CoreChrgStusCd"`
	LoanRepayStusCd               string  `json:"LoanRepayStusCd"`
	PageNo                        int     `json:"PageNo"`
	PageRecCount                  int     `json:"PageRecCount"`
}

type DAILRPL2O struct {
	Records                      []DAILRPL2ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type DAILRPL2ORecords struct {
	LoanAcctNo                    string  `json:"LoanAcctNo"`
	AcctnDt                       string  `json:"AcctnDt"`
	PridnumSeqNo                  int     `json:"PridnumSeqNo"`
	AcctiAcctNo                   string  `json:"AcctiAcctNo"`
	RepayDt                       string  `json:"RepayDt"`
	CurCd                         string  `json:"CurCd"`
	ActlRepayTotlAmt              float64 `json:"ActlRepayTotlAmt"`
	ActlRepayPrin                 float64 `json:"ActlRepayPrin"`
	ActlRepayIntr                 float64 `json:"ActlRepayIntr"`
	ActlRepayOvdueIntr            float64 `json:"ActlRepayOvdueIntr"`
	ActlRepayCmpdAmt              float64 `json:"ActlRepayCmpdAmt"`
	CurrPeriodUnStillTotlAmt      float64 `json:"CurrPeriodUnStillTotlAmt"`
	CurrPeriodUnStillPrin         float64 `json:"CurrPeriodUnStillPrin"`
	CurrPeriodUnStillIntr         float64 `json:"CurrPeriodUnStillIntr"`
	CurrPeriodUnStillOvdueIntr    float64 `json:"CurrPeriodUnStillOvdueIntr"`
	CurrPeriodUnStillCmpdAmt      float64 `json:"CurrPeriodUnStillCmpdAmt"`
	CoreChrgStusCd                string  `json:"CoreChrgStusCd"`
	LoanRepayStusCd               string  `json:"LoanRepayStusCd"`
}

// @Desc Build request message
func (o *DAILRPL2RequestForm) PackRequest(dailrpl2I DAILRPL2I) (responseBody []byte, err error) {

	requestForm := DAILRPL2RequestForm{
		Form: []DAILRPL2IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRPL2I",
				},

				FormData: dailrpl2I,
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
func (o *DAILRPL2RequestForm) UnPackRequest(request []byte) (DAILRPL2I, error) {
	dailrpl2I := DAILRPL2I{}
	if err := json.Unmarshal(request, o); nil != err {
		return dailrpl2I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrpl2I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *DAILRPL2ResponseForm) PackResponse(dailrpl2O DAILRPL2O) (responseBody []byte, err error) {
	responseForm := DAILRPL2ResponseForm{
		Form: []DAILRPL2ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRPL2O",
				},
				FormData: dailrpl2O,
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
func (o *DAILRPL2ResponseForm) UnPackResponse(request []byte) (DAILRPL2O, error) {

	dailrpl2O := DAILRPL2O{}

	if err := json.Unmarshal(request, o); nil != err {
		return dailrpl2O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrpl2O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *DAILRPL2I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
