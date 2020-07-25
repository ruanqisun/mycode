package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type DAILRMB3IDataForm struct {
	FormHead CommonFormHead
	FormData DAILRMB3I
}

type DAILRMB3ODataForm struct {
	FormHead CommonFormHead
	FormData DAILRMB3O
}

type DAILRMB3RequestForm struct {
	Form []DAILRMB3IDataForm
}

type DAILRMB3ResponseForm struct {
	Form []DAILRMB3ODataForm
}

type DAILRMB3I struct {
	TmplNo                       string  `json:"TmplNo"`
	AutogrPtyIdtyCd              string  `json:"AutogrPtyIdtyCd"`
	SigtPgnum                    int     `json:"SigtPgnum"`
	UpOffset                     int     `json:"UpOffset"`
	DownOffset                   int     `json:"DownOffset"`
	LeftOffset                   int     `json:"LeftOffset"`
	RightOffset                  int     `json:"RightOffset"`
	KeywdNm                      string  `json:"KeywdNm"`
	Offset                       int     `json:"Offset"`
	AutogrPlcinKeywdPtyBitCd     string  `json:"AutogrPlcinKeywdPtyBitCd"`
	KeywdMtchSeqCd               string  `json:"KeywdMtchSeqCd"`
	PageNo                       int     `json:"PageNo"`
	PageRecCount                 int     `json:"PageRecCount"`
}

type DAILRMB3O struct {
	Records                      []DAILRMB3ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type DAILRMB3ORecords struct {
	TmplNo                       string  `json:"TmplNo"`
	AutogrPtyIdtyCd              string  `json:"AutogrPtyIdtyCd"`
	SigtPgnum                    int     `json:"SigtPgnum"`
	UpOffset                     int     `json:"UpOffset"`
	DownOffset                   int     `json:"DownOffset"`
	LeftOffset                   int     `json:"LeftOffset"`
	RightOffset                  int     `json:"RightOffset"`
	KeywdNm                      string  `json:"KeywdNm"`
	Offset                       int     `json:"Offset"`
	AutogrPlcinKeywdPtyBitCd     string  `json:"AutogrPlcinKeywdPtyBitCd"`
	KeywdMtchSeqCd               string  `json:"KeywdMtchSeqCd"`
}

// @Desc Build request message
func (o *DAILRMB3RequestForm) PackRequest(dailrmb3I DAILRMB3I) (responseBody []byte, err error) {

	requestForm := DAILRMB3RequestForm{
		Form: []DAILRMB3IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRMB3I",
				},

				FormData: dailrmb3I,
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
func (o *DAILRMB3RequestForm) UnPackRequest(request []byte) (DAILRMB3I, error) {
	dailrmb3I := DAILRMB3I{}
	if err := json.Unmarshal(request, o); nil != err {
		return dailrmb3I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrmb3I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *DAILRMB3ResponseForm) PackResponse(dailrmb3O DAILRMB3O) (responseBody []byte, err error) {
	responseForm := DAILRMB3ResponseForm{
		Form: []DAILRMB3ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRMB3O",
				},
				FormData: dailrmb3O,
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
func (o *DAILRMB3ResponseForm) UnPackResponse(request []byte) (DAILRMB3O, error) {

	dailrmb3O := DAILRMB3O{}

	if err := json.Unmarshal(request, o); nil != err {
		return dailrmb3O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrmb3O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *DAILRMB3I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
