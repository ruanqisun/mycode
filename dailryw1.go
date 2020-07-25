package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type DAILRYW1IDataForm struct {
	FormHead CommonFormHead
	FormData DAILRYW1I
}

type DAILRYW1ODataForm struct {
	FormHead CommonFormHead
	FormData DAILRYW1O
}

type DAILRYW1RequestForm struct {
	Form []DAILRYW1IDataForm
}

type DAILRYW1ResponseForm struct {
	Form []DAILRYW1ODataForm
}

type DAILRYW1I struct {
	RiskClsfNo                   string  `json:"RiskClsfNo"`
	EfftDt                       string  `json:"EfftDt"`
	LoanProdtClsfCd              string  `json:"LoanProdtClsfCd"`
	LoanGuarManrCd               string  `json:"LoanGuarManrCd"`
	LoanRiskClsfCd               string  `json:"LoanRiskClsfCd"`
	LoanRiskClsfDescr            string  `json:"LoanRiskClsfDescr"`
	OvdueAmtTypCd                string  `json:"OvdueAmtTypCd"`
	OvdueDaysTypCd               string  `json:"OvdueDaysTypCd"`
	ExecCycCd                    string  `json:"ExecCycCd"`
	BatClsfDtCd                  string  `json:"BatClsfDtCd"`
	OvdueFloorDays               int     `json:"OvdueFloorDays"`
	OvdueCeilDays                int     `json:"OvdueCeilDays"`
	RiskClsfOvdueDecdManrCd      string  `json:"RiskClsfOvdueDecdManrCd"`
	CrtTelrNo                    string  `json:"CrtTelrNo"`
	CrtDt                        string  `json:"CrtDt"`
	CrtTm                        string  `json:"CrtTm"`
	CrtOrgNo                     string  `json:"CrtOrgNo"`
	RiskClsfValidFlg             string  `json:"RiskClsfValidFlg"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
	PageNo                       int     `json:"PageNo"`
	PageRecCount                 int     `json:"PageRecCount"`
}

type DAILRYW1O struct {
	Records                      []DAILRYW1ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type DAILRYW1ORecords struct {
	RiskClsfNo                   string  `json:"RiskClsfNo"`
	EfftDt                       string  `json:"EfftDt"`
	LoanProdtClsfCd              string  `json:"LoanProdtClsfCd"`
	LoanGuarManrCd               string  `json:"LoanGuarManrCd"`
	LoanRiskClsfCd               string  `json:"LoanRiskClsfCd"`
	LoanRiskClsfDescr            string  `json:"LoanRiskClsfDescr"`
	OvdueAmtTypCd                string  `json:"OvdueAmtTypCd"`
	OvdueDaysTypCd               string  `json:"OvdueDaysTypCd"`
	ExecCycCd                    string  `json:"ExecCycCd"`
	BatClsfDtCd                  string  `json:"BatClsfDtCd"`
	OvdueFloorDays               int     `json:"OvdueFloorDays"`
	OvdueCeilDays                int     `json:"OvdueCeilDays"`
	RiskClsfOvdueDecdManrCd      string  `json:"RiskClsfOvdueDecdManrCd"`
	CrtTelrNo                    string  `json:"CrtTelrNo"`
	CrtDt                        string  `json:"CrtDt"`
	CrtTm                        string  `json:"CrtTm"`
	CrtOrgNo                     string  `json:"CrtOrgNo"`
	RiskClsfValidFlg             string  `json:"RiskClsfValidFlg"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
}

// @Desc Build request message
func (o *DAILRYW1RequestForm) PackRequest(dailryw1I DAILRYW1I) (responseBody []byte, err error) {

	requestForm := DAILRYW1RequestForm{
		Form: []DAILRYW1IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRYW1I",
				},

				FormData: dailryw1I,
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
func (o *DAILRYW1RequestForm) UnPackRequest(request []byte) (DAILRYW1I, error) {
	dailryw1I := DAILRYW1I{}
	if err := json.Unmarshal(request, o); nil != err {
		return dailryw1I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailryw1I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *DAILRYW1ResponseForm) PackResponse(dailryw1O DAILRYW1O) (responseBody []byte, err error) {
	responseForm := DAILRYW1ResponseForm{
		Form: []DAILRYW1ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRYW1O",
				},
				FormData: dailryw1O,
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
func (o *DAILRYW1ResponseForm) UnPackResponse(request []byte) (DAILRYW1O, error) {

	dailryw1O := DAILRYW1O{}

	if err := json.Unmarshal(request, o); nil != err {
		return dailryw1O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailryw1O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *DAILRYW1I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
