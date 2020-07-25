package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type DAILRDK8IDataForm struct {
	FormHead CommonFormHead
	FormData DAILRDK8I
}

type DAILRDK8ODataForm struct {
	FormHead CommonFormHead
	FormData DAILRDK8O
}

type DAILRDK8RequestForm struct {
	Form []DAILRDK8IDataForm
}

type DAILRDK8ResponseForm struct {
	Form []DAILRDK8ODataForm
}

type DAILRDK8I struct {
	KeprcdNo                     string  `json:"KeprcdNo"`
	CustNo                       string  `json:"CustNo"`
	LoanProdtNo                  string  `json:"LoanProdtNo"`
	KeprcdStusCd                 string  `json:"KeprcdStusCd"`
	MgmtOrgNo                    string  `json:"MgmtOrgNo"`
	RepayDtChgFlg                string  `json:"RepayDtChgFlg"`
	CrtDt                        string  `json:"CrtDt"`
	CtrtNo                       string  `json:"CtrtNo"`
	LoanDubilNo                  string  `json:"LoanDubilNo"`
	SpCycCd                      string  `json:"SpCycCd"`
	SpCycQty                     int     `json:"SpCycQty"`
	IntStlDtCyc                  string  `json:"IntStlDtCyc"`
	IntStlCycQty                 int     `json:"IntStlCycQty"`
	IntStlDay                    string  `json:"IntStlDay"`
	SpDay                        string  `json:"SpDay"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type DAILRDK8O struct {
	Records                      []DAILRDK8ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type DAILRDK8ORecords struct {
	KeprcdNo                     string  `json:"KeprcdNo"`
	CustNo                       string  `json:"CustNo"`
	LoanProdtNo                  string  `json:"LoanProdtNo"`
	KeprcdStusCd                 string  `json:"KeprcdStusCd"`
	MgmtOrgNo                    string  `json:"MgmtOrgNo"`
	RepayDtChgFlg                string  `json:"RepayDtChgFlg"`
	CrtDt                        string  `json:"CrtDt"`
	CtrtNo                       string  `json:"CtrtNo"`
	LoanDubilNo                  string  `json:"LoanDubilNo"`
	SpCycCd                      string  `json:"SpCycCd"`
	SpCycQty                     int     `json:"SpCycQty"`
	IntStlDtCyc                  string  `json:"IntStlDtCyc"`
	IntStlCycQty                 int     `json:"IntStlCycQty"`
	IntStlDay                    string  `json:"IntStlDay"`
	SpDay                        string  `json:"SpDay"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
}

// @Desc Build request message
func (o *DAILRDK8RequestForm) PackRequest(dailrdk8I DAILRDK8I) (responseBody []byte, err error) {

	requestForm := DAILRDK8RequestForm{
		Form: []DAILRDK8IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRDK8I",
				},
				FormData: dailrdk8I,
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
func (o *DAILRDK8RequestForm) UnPackRequest(request []byte) (DAILRDK8I, error) {
	dailrdk8I := DAILRDK8I{}
	if err := json.Unmarshal(request, o); nil != err {
		return dailrdk8I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrdk8I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *DAILRDK8ResponseForm) PackResponse(dailrdk8O DAILRDK8O) (responseBody []byte, err error) {
	responseForm := DAILRDK8ResponseForm{
		Form: []DAILRDK8ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRDK8O",
				},
				FormData: dailrdk8O,
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
func (o *DAILRDK8ResponseForm) UnPackResponse(request []byte) (DAILRDK8O, error) {

	dailrdk8O := DAILRDK8O{}

	if err := json.Unmarshal(request, o); nil != err {
		return dailrdk8O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrdk8O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *DAILRDK8I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
