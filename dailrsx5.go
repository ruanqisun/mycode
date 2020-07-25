package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type DAILRSX5IDataForm struct {
	FormHead CommonFormHead
	FormData DAILRSX5I
}

type DAILRSX5ODataForm struct {
	FormHead CommonFormHead
	FormData DAILRSX5O
}

type DAILRSX5RequestForm struct {
	Form []DAILRSX5IDataForm
}

type DAILRSX5ResponseForm struct {
	Form []DAILRSX5ODataForm
}

type DAILRSX5I struct {
	MobileNo                     string  `json:"MobileNo"`
	SeqNo                        int     `json:"SeqNo"`
	CustNo                       string  `json:"CustNo"`
	CustName                     string  `json:"CustName"`
	SgnOrgNo                     string  `json:"SgnOrgNo"`
	FileSgnNo                    string  `json:"FileSgnNo"`
	TmplNo                       string  `json:"TmplNo"`
	AuthTypCd                    string  `json:"AuthTypCd"`
	AuthQuryStusCd               string  `json:"AuthQuryStusCd"`
	AuthDt                       string  `json:"AuthDt"`
	AuthMatrDt                   string  `json:"AuthMatrDt"`
	DocMgmtMainTablPkNo          int     `json:"DocMgmtMainTablPkNo"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
	PageNo                       int     `json:"PageNo"`
	PageRecCount                 int     `json:"PageRecCount"`
}

type DAILRSX5O struct {
	Records                      []DAILRSX5ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type DAILRSX5ORecords struct {
	MobileNo                     string  `json:"MobileNo"`
	SeqNo                        int     `json:"SeqNo"`
	CustNo                       string  `json:"CustNo"`
	CustName                     string  `json:"CustName"`
	SgnOrgNo                     string  `json:"SgnOrgNo"`
	FileSgnNo                    string  `json:"FileSgnNo"`
	TmplNo                       string  `json:"TmplNo"`
	AuthTypCd                    string  `json:"AuthTypCd"`
	AuthQuryStusCd               string  `json:"AuthQuryStusCd"`
	AuthDt                       string  `json:"AuthDt"`
	AuthMatrDt                   string  `json:"AuthMatrDt"`
	DocMgmtMainTablPkNo          int     `json:"DocMgmtMainTablPkNo"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
}

// @Desc Build request message
func (o *DAILRSX5RequestForm) PackRequest(dailrsx5I DAILRSX5I) (responseBody []byte, err error) {

	requestForm := DAILRSX5RequestForm{
		Form: []DAILRSX5IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRSX5I",
				},

				FormData: dailrsx5I,
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
func (o *DAILRSX5RequestForm) UnPackRequest(request []byte) (DAILRSX5I, error) {
	dailrsx5I := DAILRSX5I{}
	if err := json.Unmarshal(request, o); nil != err {
		return dailrsx5I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrsx5I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *DAILRSX5ResponseForm) PackResponse(dailrsx5O DAILRSX5O) (responseBody []byte, err error) {
	responseForm := DAILRSX5ResponseForm{
		Form: []DAILRSX5ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRSX5O",
				},
				FormData: dailrsx5O,
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
func (o *DAILRSX5ResponseForm) UnPackResponse(request []byte) (DAILRSX5O, error) {

	dailrsx5O := DAILRSX5O{}

	if err := json.Unmarshal(request, o); nil != err {
		return dailrsx5O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrsx5O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *DAILRSX5I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
