package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type DAILRDK3IDataForm struct {
	FormHead CommonFormHead
	FormData DAILRDK3I
}

type DAILRDK3ODataForm struct {
	FormHead CommonFormHead
	FormData DAILRDK3O
}

type DAILRDK3RequestForm struct {
	Form []DAILRDK3IDataForm
}

type DAILRDK3ResponseForm struct {
	Form []DAILRDK3ODataForm
}

type DAILRDK3I struct {
	LoanDubilNo                  string  `json:"LoanDubilNo"`
	RecalDt                      string  `json:"RecalDt"`
	SeqNo                        int     `json:"SeqNo"`
	CustNo                       string  `json:"CustNo"`
	RemainPrin                   float64 `json:"RemainPrin"`
	ActlstPrin                   float64 `json:"ActlstPrin"`
	ActlstIntr                   float64 `json:"ActlstIntr"`
	Pridnum                      int     `json:"Pridnum"`
	KeprcdStusCd                 string  `json:"KeprcdStusCd"`
	CurrPeriodRecalFlg           string  `json:"CurrPeriodRecalFlg"`
	CurrPeriodIntacrBgnDt        string  `json:"CurrPeriodIntacrBgnDt"`
	CurrPeriodIntacrEndDt        string  `json:"CurrPeriodIntacrEndDt"`
	RepayDt                      string  `json:"RepayDt"`
	PlanRepayTotlAmt             float64 `json:"PlanRepayTotlAmt"`
	PlanRepayPrin                float64 `json:"PlanRepayPrin"`
	PlanRepayIntr                float64 `json:"PlanRepayIntr"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type DAILRDK3O struct {
	Records                      []DAILRDK3ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type DAILRDK3ORecords struct {
	LoanDubilNo                  string  `json:"LoanDubilNo"`
	RecalDt                      string  `json:"RecalDt"`
	SeqNo                        int     `json:"SeqNo"`
	CustNo                       string  `json:"CustNo"`
	RemainPrin                   float64 `json:"RemainPrin"`
	ActlstPrin                   float64 `json:"ActlstPrin"`
	ActlstIntr                   float64 `json:"ActlstIntr"`
	Pridnum                      int     `json:"Pridnum"`
	KeprcdStusCd                 string  `json:"KeprcdStusCd"`
	CurrPeriodRecalFlg           string  `json:"CurrPeriodRecalFlg"`
	CurrPeriodIntacrBgnDt        string  `json:"CurrPeriodIntacrBgnDt"`
	CurrPeriodIntacrEndDt        string  `json:"CurrPeriodIntacrEndDt"`
	RepayDt                      string  `json:"RepayDt"`
	PlanRepayTotlAmt             float64 `json:"PlanRepayTotlAmt"`
	PlanRepayPrin                float64 `json:"PlanRepayPrin"`
	PlanRepayIntr                float64 `json:"PlanRepayIntr"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
}

// @Desc Build request message
func (o *DAILRDK3RequestForm) PackRequest(dailrdk3I DAILRDK3I) (responseBody []byte, err error) {

	requestForm := DAILRDK3RequestForm{
		Form: []DAILRDK3IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRDK3I",
				},
				FormData: dailrdk3I,
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
func (o *DAILRDK3RequestForm) UnPackRequest(request []byte) (DAILRDK3I, error) {
	dailrdk3I := DAILRDK3I{}
	if err := json.Unmarshal(request, o); nil != err {
		return dailrdk3I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrdk3I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *DAILRDK3ResponseForm) PackResponse(dailrdk3O DAILRDK3O) (responseBody []byte, err error) {
	responseForm := DAILRDK3ResponseForm{
		Form: []DAILRDK3ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRDK3O",
				},
				FormData: dailrdk3O,
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
func (o *DAILRDK3ResponseForm) UnPackResponse(request []byte) (DAILRDK3O, error) {

	dailrdk3O := DAILRDK3O{}

	if err := json.Unmarshal(request, o); nil != err {
		return dailrdk3O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrdk3O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *DAILRDK3I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
