package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type DAILRYW0IDataForm struct {
	FormHead CommonFormHead
	FormData DAILRYW0I
}

type DAILRYW0ODataForm struct {
	FormHead CommonFormHead
	FormData DAILRYW0O
}

type DAILRYW0RequestForm struct {
	Form []DAILRYW0IDataForm
}

type DAILRYW0ResponseForm struct {
	Form []DAILRYW0ODataForm
}

type DAILRYW0I struct {
	TxReqeCd                      string  `json:"TxReqeCd"`
	ChnlNoCd                      string  `json:"ChnlNoCd"`
	BizSysNo                      string  `json:"BizSysNo"`
	RcrdacctFlgCd                 string  `json:"RcrdacctFlgCd"`
	DsmtJrnlzTypCd                string  `json:"DsmtJrnlzTypCd"`
	InterTranCd                   string  `json:"InterTranCd"`
	LiqdBizTypCd                  string  `json:"LiqdBizTypCd"`
	Bak1                          string  `json:"Bak1"`
	Bak2                          string  `json:"Bak2"`
	Bak3                          string  `json:"Bak3"`
	DsmtJrnlzExecAgngCd           string  `json:"DsmtJrnlzExecAgngCd"`
	FinlModfyDt                   string  `json:"FinlModfyDt"`
	FinlModfyTm                   string  `json:"FinlModfyTm"`
	FinlModfyOrgNo                string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo               string  `json:"FinlModfyTelrNo"`
	PageNo                        int     `json:"PageNo"`
	PageRecCount                  int     `json:"PageRecCount"`
}

type DAILRYW0O struct {
	Records                      []DAILRYW0ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type DAILRYW0ORecords struct {
	TxReqeCd                      string  `json:"TxReqeCd"`
	ChnlNoCd                      string  `json:"ChnlNoCd"`
	BizSysNo                      string  `json:"BizSysNo"`
	RcrdacctFlgCd                 string  `json:"RcrdacctFlgCd"`
	DsmtJrnlzTypCd                string  `json:"DsmtJrnlzTypCd"`
	InterTranCd                   string  `json:"InterTranCd"`
	LiqdBizTypCd                  string  `json:"LiqdBizTypCd"`
	Bak1                          string  `json:"Bak1"`
	Bak2                          string  `json:"Bak2"`
	Bak3                          string  `json:"Bak3"`
	DsmtJrnlzExecAgngCd           string  `json:"DsmtJrnlzExecAgngCd"`
	FinlModfyDt                   string  `json:"FinlModfyDt"`
	FinlModfyTm                   string  `json:"FinlModfyTm"`
	FinlModfyOrgNo                string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo               string  `json:"FinlModfyTelrNo"`
}

// @Desc Build request message
func (o *DAILRYW0RequestForm) PackRequest(dailryw0I DAILRYW0I) (responseBody []byte, err error) {

	requestForm := DAILRYW0RequestForm{
		Form: []DAILRYW0IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRYW0I",
				},

				FormData: dailryw0I,
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
func (o *DAILRYW0RequestForm) UnPackRequest(request []byte) (DAILRYW0I, error) {
	dailryw0I := DAILRYW0I{}
	if err := json.Unmarshal(request, o); nil != err {
		return dailryw0I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailryw0I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *DAILRYW0ResponseForm) PackResponse(dailryw0O DAILRYW0O) (responseBody []byte, err error) {
	responseForm := DAILRYW0ResponseForm{
		Form: []DAILRYW0ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRYW0O",
				},
				FormData: dailryw0O,
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
func (o *DAILRYW0ResponseForm) UnPackResponse(request []byte) (DAILRYW0O, error) {

	dailryw0O := DAILRYW0O{}

	if err := json.Unmarshal(request, o); nil != err {
		return dailryw0O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailryw0O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *DAILRYW0I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
