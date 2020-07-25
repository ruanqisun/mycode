package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type DAILRLS1IDataForm struct {
	FormHead CommonFormHead
	FormData DAILRLS1I
}

type DAILRLS1ODataForm struct {
	FormHead CommonFormHead
	FormData DAILRLS1O
}

type DAILRLS1RequestForm struct {
	Form []DAILRLS1IDataForm
}

type DAILRLS1ResponseForm struct {
	Form []DAILRLS1ODataForm
}

type DAILRLS1I struct {
	IndvCrtftypCd                  string  `json:"IndvCrtftypCd"`
	CrtfNo                         string  `json:"CrtfNo"`
	CrdQuryDt                      string  `json:"CrdQuryDt"`
	CrdQuryResultCd                string  `json:"CrdQuryResultCd"`
	CrdQuryResultId                string  `json:"CrdQuryResultId"`
	PageNo                         int     `json:"PageNo"`
	PageRecCount                   int     `json:"PageRecCount"`
}

type DAILRLS1O struct {
	Records                      []DAILRLS1ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type DAILRLS1ORecords struct {
	IndvCrtftypCd                  string  `json:"IndvCrtftypCd"`
	CrtfNo                         string  `json:"CrtfNo"`
	CrdQuryDt                      string  `json:"CrdQuryDt"`
	CrdQuryResultCd                string  `json:"CrdQuryResultCd"`
	CrdQuryResultId                string  `json:"CrdQuryResultId"`
}

// @Desc Build request message
func (o *DAILRLS1RequestForm) PackRequest(dailrls1I DAILRLS1I) (responseBody []byte, err error) {

	requestForm := DAILRLS1RequestForm{
		Form: []DAILRLS1IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRLS1I",
				},

				FormData: dailrls1I,
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
func (o *DAILRLS1RequestForm) UnPackRequest(request []byte) (DAILRLS1I, error) {
	dailrls1I := DAILRLS1I{}
	if err := json.Unmarshal(request, o); nil != err {
		return dailrls1I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrls1I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *DAILRLS1ResponseForm) PackResponse(dailrls1O DAILRLS1O) (responseBody []byte, err error) {
	responseForm := DAILRLS1ResponseForm{
		Form: []DAILRLS1ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRLS1O",
				},
				FormData: dailrls1O,
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
func (o *DAILRLS1ResponseForm) UnPackResponse(request []byte) (DAILRLS1O, error) {

	dailrls1O := DAILRLS1O{}

	if err := json.Unmarshal(request, o); nil != err {
		return dailrls1O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrls1O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *DAILRLS1I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
