package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type DAILRDH3IDataForm struct {
	FormHead CommonFormHead
	FormData DAILRDH3I
}

type DAILRDH3ODataForm struct {
	FormHead CommonFormHead
	FormData DAILRDH3O
}

type DAILRDH3RequestForm struct {
	Form []DAILRDH3IDataForm
}

type DAILRDH3ResponseForm struct {
	Form []DAILRDH3ODataForm
}

type DAILRDH3I struct {
	CollTaskNo                   string  `json:"CollTaskNo"`
	SeqNo                        int     `json:"SeqNo"`
	GnrtDt                       string  `json:"GnrtDt"`
	CollMnrpEmpnbr               string  `json:"CollMnrpEmpnbr"`
	CollMnrpOrgNo                string  `json:"CollMnrpOrgNo"`
	TrnsfDt                      string  `json:"TrnsfDt"`
	TrnsfManrCd                  string  `json:"TrnsfManrCd"`
	TsfinChnlNoCd                string  `json:"TsfinChnlNoCd"`
	TsfinEmpnbr                  string  `json:"TsfinEmpnbr"`
	TsfinOrgNo                   string  `json:"TsfinOrgNo"`
	OperPersonEmpnbr             string  `json:"OperPersonEmpnbr"`
	OperOrgNo                    string  `json:"OperOrgNo"`
	TrnsfStusCd                  string  `json:"TrnsfStusCd"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
	PageNo                       int     `json:"PageNo"`
	PageRecCount                 int     `json:"PageRecCount"`
}

type DAILRDH3O struct {
	Records                      []DAILRDH3ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type DAILRDH3ORecords struct {
	CollTaskNo                   string  `json:"CollTaskNo"`
	SeqNo                        int     `json:"SeqNo"`
	GnrtDt                       string  `json:"GnrtDt"`
	CollMnrpEmpnbr               string  `json:"CollMnrpEmpnbr"`
	CollMnrpOrgNo                string  `json:"CollMnrpOrgNo"`
	TrnsfDt                      string  `json:"TrnsfDt"`
	TrnsfManrCd                  string  `json:"TrnsfManrCd"`
	TsfinChnlNoCd                string  `json:"TsfinChnlNoCd"`
	TsfinEmpnbr                  string  `json:"TsfinEmpnbr"`
	TsfinOrgNo                   string  `json:"TsfinOrgNo"`
	OperPersonEmpnbr             string  `json:"OperPersonEmpnbr"`
	OperOrgNo                    string  `json:"OperOrgNo"`
	TrnsfStusCd                  string  `json:"TrnsfStusCd"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
}

// @Desc Build request message
func (o *DAILRDH3RequestForm) PackRequest(dailrdh3I DAILRDH3I) (responseBody []byte, err error) {

	requestForm := DAILRDH3RequestForm{
		Form: []DAILRDH3IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRDH3I",
				},

				FormData: dailrdh3I,
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
func (o *DAILRDH3RequestForm) UnPackRequest(request []byte) (DAILRDH3I, error) {
	dailrdh3I := DAILRDH3I{}
	if err := json.Unmarshal(request, o); nil != err {
		return dailrdh3I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrdh3I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *DAILRDH3ResponseForm) PackResponse(dailrdh3O DAILRDH3O) (responseBody []byte, err error) {
	responseForm := DAILRDH3ResponseForm{
		Form: []DAILRDH3ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRDH3O",
				},
				FormData: dailrdh3O,
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
func (o *DAILRDH3ResponseForm) UnPackResponse(request []byte) (DAILRDH3O, error) {

	dailrdh3O := DAILRDH3O{}

	if err := json.Unmarshal(request, o); nil != err {
		return dailrdh3O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrdh3O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *DAILRDH3I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
