package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type DAILRYW5IDataForm struct {
	FormHead CommonFormHead
	FormData DAILRYW5I
}

type DAILRYW5ODataForm struct {
	FormHead CommonFormHead
	FormData DAILRYW5O
}

type DAILRYW5RequestForm struct {
	Form []DAILRYW5IDataForm
}

type DAILRYW5ResponseForm struct {
	Form []DAILRYW5ODataForm
}

type DAILRYW5I struct {
	LoanProdtNo                  string  `json:"LoanProdtNo"`
	LoanProdtNm                  string  `json:"LoanProdtNm"`
	RoleNo                       string  `json:"RoleNo"`
	Empnbr                       string  `json:"Empnbr"`
	OrgNo                        string  `json:"OrgNo"`
	OrgNm                        string  `json:"OrgNm"`
	EmplyName                    string  `json:"EmplyName"`
	RoleNm                       string  `json:"RoleNm"`
	ValidFlg                     string  `json:"ValidFlg"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModifrEmpnbr             string  `json:"FinlModifrEmpnbr"`
	PageNo                       int     `json:"PageNo"`
	PageRecCount                 int     `json:"PageRecCount"`
}

type DAILRYW5O struct {
	Records                      []DAILRYW5ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type DAILRYW5ORecords struct {
	LoanProdtNo                  string  `json:"LoanProdtNo"`
	LoanProdtNm                  string  `json:"LoanProdtNm"`
	RoleNo                       string  `json:"RoleNo"`
	Empnbr                       string  `json:"Empnbr"`
	OrgNo                        string  `json:"OrgNo"`
	OrgNm                        string  `json:"OrgNm"`
	EmplyName                    string  `json:"EmplyName"`
	RoleNm                       string  `json:"RoleNm"`
	ValidFlg                     string  `json:"ValidFlg"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModifrEmpnbr             string  `json:"FinlModifrEmpnbr"`
}

// @Desc Build request message
func (o *DAILRYW5RequestForm) PackRequest(dailryw5I DAILRYW5I) (responseBody []byte, err error) {

	requestForm := DAILRYW5RequestForm{
		Form: []DAILRYW5IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRYW5I",
				},

				FormData: dailryw5I,
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
func (o *DAILRYW5RequestForm) UnPackRequest(request []byte) (DAILRYW5I, error) {
	dailryw5I := DAILRYW5I{}
	if err := json.Unmarshal(request, o); nil != err {
		return dailryw5I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailryw5I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *DAILRYW5ResponseForm) PackResponse(dailryw5O DAILRYW5O) (responseBody []byte, err error) {
	responseForm := DAILRYW5ResponseForm{
		Form: []DAILRYW5ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRYW5O",
				},
				FormData: dailryw5O,
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
func (o *DAILRYW5ResponseForm) UnPackResponse(request []byte) (DAILRYW5O, error) {

	dailryw5O := DAILRYW5O{}

	if err := json.Unmarshal(request, o); nil != err {
		return dailryw5O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailryw5O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *DAILRYW5I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
