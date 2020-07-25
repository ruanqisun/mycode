package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type DAILRSX4IDataForm struct {
	FormHead CommonFormHead
	FormData DAILRSX4I
}

type DAILRSX4ODataForm struct {
	FormHead CommonFormHead
	FormData DAILRSX4O
}

type DAILRSX4RequestForm struct {
	Form []DAILRSX4IDataForm
}

type DAILRSX4ResponseForm struct {
	Form []DAILRSX4ODataForm
}

type DAILRSX4I struct {
	CrdtAplySn                   string  `json:"CrdtAplySn"`
	SeqNo       	             int     `json:"SeqNo"`
	CustNo                       string  `json:"CustNo"`
	OperPersonName               string  `json:"OperPersonName"`
	LnrwModeCd                   string  `json:"LnrwModeCd"`
	ApprvManrCd                  string  `json:"ApprvManrCd"`
	OperEndTm                    string  `json:"OperEndTm"`
	NodeTypCd                    string  `json:"NodeTypCd"`
	OperPersonEmpnbr             string  `json:"OperPersonEmpnbr"`
	OperOrgNo                    string  `json:"OperOrgNo"`
	//OperBgnDt                    string  `json:"OperBgnDt"`
	OperBgnTm                    string  `json:"OperBgnTm"`
	ApprvSugstnCd                string  `json:"ApprvSugstnCd"`
	RplshInfo                    string  `json:"RplshInfo"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
	PageNo                       int     `json:"PageNo"`
	PageRecCount                 int     `json:"PageRecCount"`
}

type DAILRSX4O struct {
	Records                      []DAILRSX4ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type DAILRSX4ORecords struct {
	CrdtAplySn                   string  `json:"CrdtAplySn"`
	SeqNo       	             int     `json:"SeqNo"`
	CustNo                       string  `json:"CustNo"`
	OperPersonName               string  `json:"OperPersonName"`
	LnrwModeCd                   string  `json:"LnrwModeCd"`
	ApprvManrCd                  string  `json:"ApprvManrCd"`
	OperEndTm                    string  `json:"OperEndTm"`
	NodeTypCd                    string  `json:"NodeTypCd"`
	OperPersonEmpnbr             string  `json:"OperPersonEmpnbr"`
	OperOrgNo                    string  `json:"OperOrgNo"`
	//OperBgnDt                    string  `json:"OperBgnDt"`
	OperBgnTm                    string  `json:"OperBgnTm"`
	ApprvSugstnCd                string  `json:"ApprvSugstnCd"`
	RplshInfo                    string  `json:"RplshInfo"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
}

// @Desc Build request message
func (o *DAILRSX4RequestForm) PackRequest(dailrsx4I DAILRSX4I) (responseBody []byte, err error) {

	requestForm := DAILRSX4RequestForm{
		Form: []DAILRSX4IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRSX4I",
				},

				FormData: dailrsx4I,
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
func (o *DAILRSX4RequestForm) UnPackRequest(request []byte) (DAILRSX4I, error) {
	dailrsx4I := DAILRSX4I{}
	if err := json.Unmarshal(request, o); nil != err {
		return dailrsx4I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrsx4I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *DAILRSX4ResponseForm) PackResponse(dailrsx4O DAILRSX4O) (responseBody []byte, err error) {
	responseForm := DAILRSX4ResponseForm{
		Form: []DAILRSX4ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRSX4O",
				},
				FormData: dailrsx4O,
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
func (o *DAILRSX4ResponseForm) UnPackResponse(request []byte) (DAILRSX4O, error) {

	dailrsx4O := DAILRSX4O{}

	if err := json.Unmarshal(request, o); nil != err {
		return dailrsx4O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrsx4O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *DAILRSX4I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
