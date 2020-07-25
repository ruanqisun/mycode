package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type IL8R0026IDataForm struct {
	FormHead CommonFormHead
	FormData IL8R0026I
}

type IL8R0026ODataForm struct {
	FormHead CommonFormHead
	FormData IL8R0026O
}

type IL8R0026RequestForm struct {
	Form []IL8R0026IDataForm
}

type IL8R0026ResponseForm struct {
	Form []IL8R0026ODataForm
}

type IL8R0026I struct {
	TaskExampId                    string  `json:"TaskExampId"`
	ProcesParaVal                  string  `json:"ProcesParaVal"`
	PageNo                         int     `json:"PageNo"`
	PageRecCount                   int     `json:"PageRecCount"`
}

type IL8R0026O struct {
	Records                      []IL8R0026ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type IL8R0026ORecords struct {
	TaskExampId                    string  `json:"TaskExampId"`
	ProcesParaVal                  string  `json:"ProcesParaVal"`
}

// @Desc Build request message
func (o *IL8R0026RequestForm) PackRequest(il8r0026I IL8R0026I) (responseBody []byte, err error) {

	requestForm := IL8R0026RequestForm{
		Form: []IL8R0026IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "IL8R0026I",
				},

				FormData: il8r0026I,
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
func (o *IL8R0026RequestForm) UnPackRequest(request []byte) (IL8R0026I, error) {
	il8r0026I := IL8R0026I{}
	if err := json.Unmarshal(request, o); nil != err {
		return il8r0026I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return il8r0026I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *IL8R0026ResponseForm) PackResponse(il8r0026O IL8R0026O) (responseBody []byte, err error) {
	responseForm := IL8R0026ResponseForm{
		Form: []IL8R0026ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "IL8R0026O",
				},
				FormData: il8r0026O,
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
func (o *IL8R0026ResponseForm) UnPackResponse(request []byte) (IL8R0026O, error) {

	il8r0026O := IL8R0026O{}

	if err := json.Unmarshal(request, o); nil != err {
		return il8r0026O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return il8r0026O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *IL8R0026I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
