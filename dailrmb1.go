package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type DAILRMB1IDataForm struct {
	FormHead CommonFormHead
	FormData DAILRMB1I
}

type DAILRMB1ODataForm struct {
	FormHead CommonFormHead
	FormData DAILRMB1O
}

type DAILRMB1RequestForm struct {
	Form []DAILRMB1IDataForm
}

type DAILRMB1ResponseForm struct {
	Form []DAILRMB1ODataForm
}

type DAILRMB1I struct {
	TmplNo                       string  `json:"TmplNo"`
	TmplPdfFileId                string  `json:"TmplPdfFileId"`
	CtrtModeCd                   string  `json:"CtrtModeCd"`
	SigtPositionTypCd            string  `json:"SigtPositionTypCd"`
	CtrtTmplTypNm                string  `json:"CtrtTmplTypNm"`
	EfftDt                       string  `json:"EfftDt"`
	TmplNm                       string  `json:"TmplNm"`
	DocTmplTypCd                 string  `json:"DocTmplTypCd"`
	TmplSubTypCd                 string  `json:"TmplSubTypCd"`
	TmplFileId                   string  `json:"TmplFileId"`
	CtrtTmplComnt                string  `json:"CtrtTmplComnt"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
	PageNo                       int     `json:"PageNo"`
	PageRecCount                 int     `json:"PageRecCount"`
}

type DAILRMB1O struct {
	Records                      []DAILRMB1ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type DAILRMB1ORecords struct {
	TmplNo                       string  `json:"TmplNo"`
	TmplPdfFileId                string  `json:"TmplPdfFileId"`
	CtrtModeCd                   string  `json:"CtrtModeCd"`
	SigtPositionTypCd            string  `json:"SigtPositionTypCd"`
	CtrtTmplTypNm                string  `json:"CtrtTmplTypNm"`
	EfftDt                       string  `json:"EfftDt"`
	TmplNm                       string  `json:"TmplNm"`
	DocTmplTypCd                 string  `json:"DocTmplTypCd"`
	TmplSubTypCd                 string  `json:"TmplSubTypCd"`
	TmplFileId                   string  `json:"TmplFileId"`
	CtrtTmplComnt                string  `json:"CtrtTmplComnt"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
}

// @Desc Build request message
func (o *DAILRMB1RequestForm) PackRequest(dailrmb1I DAILRMB1I) (responseBody []byte, err error) {

	requestForm := DAILRMB1RequestForm{
		Form: []DAILRMB1IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRMB1I",
				},

				FormData: dailrmb1I,
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
func (o *DAILRMB1RequestForm) UnPackRequest(request []byte) (DAILRMB1I, error) {
	dailrmb1I := DAILRMB1I{}
	if err := json.Unmarshal(request, o); nil != err {
		return dailrmb1I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrmb1I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *DAILRMB1ResponseForm) PackResponse(dailrmb1O DAILRMB1O) (responseBody []byte, err error) {
	responseForm := DAILRMB1ResponseForm{
		Form: []DAILRMB1ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRMB1O",
				},
				FormData: dailrmb1O,
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
func (o *DAILRMB1ResponseForm) UnPackResponse(request []byte) (DAILRMB1O, error) {

	dailrmb1O := DAILRMB1O{}

	if err := json.Unmarshal(request, o); nil != err {
		return dailrmb1O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrmb1O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *DAILRMB1I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
