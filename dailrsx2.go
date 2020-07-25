package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type DAILRSX2IDataForm struct {
	FormHead CommonFormHead
	FormData DAILRSX2I
}

type DAILRSX2ODataForm struct {
	FormHead CommonFormHead
	FormData DAILRSX2O
}

type DAILRSX2RequestForm struct {
	Form []DAILRSX2IDataForm
}

type DAILRSX2ResponseForm struct {
	Form []DAILRSX2ODataForm
}

type DAILRSX2I struct {
	CrdtAplySn                   string  `json:"CrdtAplySn"`
	CustNo                       string  `json:"CustNo"`
	MonIncomeAmt                 float64 `json:"MonIncomeAmt"`
	IncomeRemrk                  string  `json:"IncomeRemrk"`
	SocietyEnsrNo                string  `json:"SocietyEnsrNo"`
	AcmltnfundAcctNo             string  `json:"AcmltnfundAcctNo"`
	TaxpyerIdtfyNo               string  `json:"TaxpyerIdtfyNo"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type DAILRSX2O struct {
	Records                      []DAILRSX2ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type DAILRSX2ORecords struct {
	CrdtAplySn                   string  `json:"CrdtAplySn"`
	CustNo                       string  `json:"CustNo"`
	MonIncomeAmt                 float64 `json:"MonIncomeAmt"`
	IncomeRemrk                  string  `json:"IncomeRemrk"`
	SocietyEnsrNo                string  `json:"SocietyEnsrNo"`
	AcmltnfundAcctNo             string  `json:"AcmltnfundAcctNo"`
	TaxpyerIdtfyNo               string  `json:"TaxpyerIdtfyNo"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
}

// @Desc Build request message
func (o *DAILRSX2RequestForm) PackRequest(dailrsx2I DAILRSX2I) (responseBody []byte, err error) {

	requestForm := DAILRSX2RequestForm{
		Form: []DAILRSX2IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRSX2I",
				},

				FormData: dailrsx2I,
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
func (o *DAILRSX2RequestForm) UnPackRequest(request []byte) (DAILRSX2I, error) {
	dailrsx2I := DAILRSX2I{}
	if err := json.Unmarshal(request, o); nil != err {
		return dailrsx2I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrsx2I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *DAILRSX2ResponseForm) PackResponse(dailrsx2O DAILRSX2O) (responseBody []byte, err error) {
	responseForm := DAILRSX2ResponseForm{
		Form: []DAILRSX2ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRSX2O",
				},
				FormData: dailrsx2O,
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
func (o *DAILRSX2ResponseForm) UnPackResponse(request []byte) (DAILRSX2O, error) {

	dailrsx2O := DAILRSX2O{}

	if err := json.Unmarshal(request, o); nil != err {
		return dailrsx2O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrsx2O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *DAILRSX2I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
