package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type DAILRKH1IDataForm struct {
	FormHead CommonFormHead
	FormData DAILRKH1I
}

type DAILRKH1ODataForm struct {
	FormHead CommonFormHead
	FormData DAILRKH1O
}

type DAILRKH1RequestForm struct {
	Form []DAILRKH1IDataForm
}

type DAILRKH1ResponseForm struct {
	Form []DAILRKH1ODataForm
}

type DAILRKH1I struct {
	TxReqeCd                     string  `json:"TxReqeCd"`
	ChnlNoCd                     string  `json:"ChnlNoCd"`
	BizSysNo                     string  `json:"BizSysNo"`
	RcrdacctFlgCd                string  `json:"RcrdacctFlgCd"`
	DsmtJrnlzTypCd               string  `json:"DsmtJrnlzTypCd"`
	InterTranCd                  string  `json:"InterTranCd"`
	LiqdBizTypCd                 string  `json:"LiqdBizTypCd"`
	Bak1                         string  `json:"Bak1"`
	Bak2                         string  `json:"Bak2"`
	Bak3                         string  `json:"Bak3"`
	DsmtJrnlzExecAgngCd          string  `json:"DsmtJrnlzExecAgngCd"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
	PageNo                       int     `json:"PageNo"`
	PageRecCount                 int     `json:"PageRecCount"`
}

type DAILRKH1O struct {
	Records                      []DAILRKH1ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type DAILRKH1ORecords struct {
	TxReqeCd                     string  `json:"TxReqeCd"`
	ChnlNoCd                     string  `json:"ChnlNoCd"`
	BizSysNo                     string  `json:"BizSysNo"`
	RcrdacctFlgCd                string  `json:"RcrdacctFlgCd"`
	DsmtJrnlzTypCd               string  `json:"DsmtJrnlzTypCd"`
	InterTranCd                  string  `json:"InterTranCd"`
	LiqdBizTypCd                 string  `json:"LiqdBizTypCd"`
	Bak1                         string  `json:"Bak1"`
	Bak2                         string  `json:"Bak2"`
	Bak3                         string  `json:"Bak3"`
	DsmtJrnlzExecAgngCd          string  `json:"DsmtJrnlzExecAgngCd"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
}

// @Desc Build request message
func (o *DAILRKH1RequestForm) PackRequest(dailrkh1I DAILRKH1I) (responseBody []byte, err error) {

	requestForm := DAILRKH1RequestForm{
		Form: []DAILRKH1IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRKH1I",
				},

				FormData: dailrkh1I,
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
func (o *DAILRKH1RequestForm) UnPackRequest(request []byte) (DAILRKH1I, error) {
	dailrkh1I := DAILRKH1I{}
	if err := json.Unmarshal(request, o); nil != err {
		return dailrkh1I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrkh1I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *DAILRKH1ResponseForm) PackResponse(dailrkh1O DAILRKH1O) (responseBody []byte, err error) {
	responseForm := DAILRKH1ResponseForm{
		Form: []DAILRKH1ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRKH1O",
				},
				FormData: dailrkh1O,
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
func (o *DAILRKH1ResponseForm) UnPackResponse(request []byte) (DAILRKH1O, error) {

	dailrkh1O := DAILRKH1O{}

	if err := json.Unmarshal(request, o); nil != err {
		return dailrkh1O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrkh1O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *DAILRKH1I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
