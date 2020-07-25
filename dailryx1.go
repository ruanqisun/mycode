package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type DAILRYX1IDataForm struct {
	FormHead CommonFormHead
	FormData DAILRYX1I
}

type DAILRYX1ODataForm struct {
	FormHead CommonFormHead
	FormData DAILRYX1O
}

type DAILRYX1RequestForm struct {
	Form []DAILRYX1IDataForm
}

type DAILRYX1ResponseForm struct {
	Form []DAILRYX1ODataForm
}

type DAILRYX1I struct {
	RgtsintClsfNo                 string  `json:"RgtsintClsfNo"`
	KeprcdValidFlg                string  `json:"KeprcdValidFlg"`
	RgtsintValidPrd               int     `json:"RgtsintValidPrd"`
	PrefrTypCd                    string  `json:"PrefrTypCd"`
	IntrtDiscntRatio              float64 `json:"IntrtDiscntRatio"`
	IntrMitgtAmt                  float64 `json:"IntrMitgtAmt"`
	CashDctblCupnAmt              float64 `json:"CashDctblCupnAmt"`
	AccmValidPrd                  int     `json:"AccmValidPrd"`
	LoanAccmTms                   int     `json:"LoanAccmTms"`
	LoanAccmAmt                   float64 `json:"LoanAccmAmt"`
	LoanSglAmt                    float64 `json:"LoanSglAmt"`
	InfrTotlDays                  int     `json:"InfrTotlDays"`
	InfrTotlPridnum               int     `json:"InfrTotlPridnum"`
	InfrPridnum                   int     `json:"InfrPridnum"`
	PageNo                        int     `json:"PageNo"`
	PageRecCount                  int     `json:"PageRecCount"`
}

type DAILRYX1O struct {
	Records                      []DAILRYX1ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type DAILRYX1ORecords struct {
	RgtsintClsfNo                 string  `json:"RgtsintClsfNo"`
	KeprcdValidFlg                string  `json:"KeprcdValidFlg"`
	RgtsintValidPrd               int     `json:"RgtsintValidPrd"`
	PrefrTypCd                    string  `json:"PrefrTypCd"`
	IntrtDiscntRatio              float64 `json:"IntrtDiscntRatio"`
	IntrMitgtAmt                  float64 `json:"IntrMitgtAmt"`
	CashDctblCupnAmt              float64 `json:"CashDctblCupnAmt"`
	AccmValidPrd                  int     `json:"AccmValidPrd"`
	LoanAccmTms                   int     `json:"LoanAccmTms"`
	LoanAccmAmt                   float64 `json:"LoanAccmAmt"`
	LoanSglAmt                    float64 `json:"LoanSglAmt"`
	InfrTotlDays                  int     `json:"InfrTotlDays"`
	InfrTotlPridnum               int     `json:"InfrTotlPridnum"`
	InfrPridnum                   int     `json:"InfrPridnum"`
}

// @Desc Build request message
func (o *DAILRYX1RequestForm) PackRequest(dailryx1I DAILRYX1I) (responseBody []byte, err error) {

	requestForm := DAILRYX1RequestForm{
		Form: []DAILRYX1IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRYX1I",
				},

				FormData: dailryx1I,
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
func (o *DAILRYX1RequestForm) UnPackRequest(request []byte) (DAILRYX1I, error) {
	dailryx1I := DAILRYX1I{}
	if err := json.Unmarshal(request, o); nil != err {
		return dailryx1I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailryx1I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *DAILRYX1ResponseForm) PackResponse(dailryx1O DAILRYX1O) (responseBody []byte, err error) {
	responseForm := DAILRYX1ResponseForm{
		Form: []DAILRYX1ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRYX1O",
				},
				FormData: dailryx1O,
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
func (o *DAILRYX1ResponseForm) UnPackResponse(request []byte) (DAILRYX1O, error) {

	dailryx1O := DAILRYX1O{}

	if err := json.Unmarshal(request, o); nil != err {
		return dailryx1O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailryx1O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *DAILRYX1I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
