package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type DAILRDK2IDataForm struct {
	FormHead CommonFormHead
	FormData DAILRDK2I
}

type DAILRDK2ODataForm struct {
	FormHead CommonFormHead
	FormData DAILRDK2O
}

type DAILRDK2RequestForm struct {
	Form []DAILRDK2IDataForm
}

type DAILRDK2ResponseForm struct {
	Form []DAILRDK2ODataForm
}

type DAILRDK2I struct {
	LoanDubilNo                  string  `json:"LoanDubilNo"`
	SeqNo                        int     `json:"SeqNo"`
	RprincPridnum                int     `json:"RprincPridnum"`
	CustNo                       string  `json:"CustNo"`
	SpDay                        string  `json:"SpDay"`
	IntStlDay                    string  `json:"IntStlDay"`
	RepayPlanValidFlg            string  `json:"RepayPlanValidFlg"`
	PlanStartDt                  string  `json:"PlanStartDt"`
	PlanMatrDt                   string  `json:"PlanMatrDt"`
	RepayManrCd                  string  `json:"RepayManrCd"`
	IntStlCycCd                  string  `json:"IntStlCycCd"`
	IntStlCycQty                 int     `json:"IntStlCycQty"`
	RpyintPridnum                int     `json:"RpyintPridnum"`
	EqamtEvrpridRepayAmt         float64 `json:"EqamtEvrpridRepayAmt"`
	SpCycCd                      string  `json:"SpCycCd"`
	SpCycQty                     int     `json:"SpCycQty"`
	CycRatioRprincRatio          float64 `json:"CycRatioRprincRatio"`
	CycFixFrhdRprincAmt          float64 `json:"CycFixFrhdRprincAmt"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type DAILRDK2O struct {
	Records                      []DAILRDK2ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type DAILRDK2ORecords struct {
	LoanDubilNo                  string  `json:"LoanDubilNo"`
	SeqNo                        int     `json:"SeqNo"`
	RprincPridnum                int     `json:"RprincPridnum"`
	CustNo                       string  `json:"CustNo"`
	SpDay                        string  `json:"SpDay"`
	IntStlDay                    string  `json:"IntStlDay"`
	RepayPlanValidFlg            string  `json:"RepayPlanValidFlg"`
	PlanStartDt                  string  `json:"PlanStartDt"`
	PlanMatrDt                   string  `json:"PlanMatrDt"`
	RepayManrCd                  string  `json:"RepayManrCd"`
	IntStlCycCd                  string  `json:"IntStlCycCd"`
	IntStlCycQty                 int     `json:"IntStlCycQty"`
	RpyintPridnum                int     `json:"RpyintPridnum"`
	EqamtEvrpridRepayAmt         float64 `json:"EqamtEvrpridRepayAmt"`
	SpCycCd                      string  `json:"SpCycCd"`
	SpCycQty                     int     `json:"SpCycQty"`
	CycRatioRprincRatio          float64 `json:"CycRatioRprincRatio"`
	CycFixFrhdRprincAmt          float64 `json:"CycFixFrhdRprincAmt"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
}

// @Desc Build request message
func (o *DAILRDK2RequestForm) PackRequest(dailrdk2I DAILRDK2I) (responseBody []byte, err error) {

	requestForm := DAILRDK2RequestForm{
		Form: []DAILRDK2IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRDK2I",
				},
				FormData: dailrdk2I,
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
func (o *DAILRDK2RequestForm) UnPackRequest(request []byte) (DAILRDK2I, error) {
	dailrdk2I := DAILRDK2I{}
	if err := json.Unmarshal(request, o); nil != err {
		return dailrdk2I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrdk2I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *DAILRDK2ResponseForm) PackResponse(dailrdk2O DAILRDK2O) (responseBody []byte, err error) {
	responseForm := DAILRDK2ResponseForm{
		Form: []DAILRDK2ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRDK2O",
				},
				FormData: dailrdk2O,
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
func (o *DAILRDK2ResponseForm) UnPackResponse(request []byte) (DAILRDK2O, error) {

	dailrdk2O := DAILRDK2O{}

	if err := json.Unmarshal(request, o); nil != err {
		return dailrdk2O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrdk2O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *DAILRDK2I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
