package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type DAILRSX8IDataForm struct {
	FormHead CommonFormHead
	FormData DAILRSX8I
}

type DAILRSX8ODataForm struct {
	FormHead CommonFormHead
	FormData DAILRSX8O
}

type DAILRSX8RequestForm struct {
	Form []DAILRSX8IDataForm
}

type DAILRSX8ResponseForm struct {
	Form []DAILRSX8ODataForm
}

type DAILRSX8I struct {
	IntrtAdjSn                   string  `json:"IntrtAdjSn"`
	CustNo                       string  `json:"CustNo"`
	CustName                     string  `json:"CustName"`
	CrdtAplySn                   string  `json:"CrdtAplySn"`
	IndvCrtfTypCd                string  `json:"IndvCrtfTypCd"`
	IndvCrtfNo                   string  `json:"IndvCrtfNo"`
	BaseIntrtTypCd               string  `json:"BaseIntrtTypCd"`
	CurCd                        string  `json:"CurCd"`
	IntrtNo                      string  `json:"IntrtNo"`
	IntrtFlotDrctCd              string  `json:"IntrtFlotDrctCd"`
	BpFlotVal                    float64  `json:"BpFlotVal"`
	KeprcdStusCd                 string  `json:"KeprcdStusCd"`
	CrtTelrNo                    string  `json:"CrtTelrNo"`
	CrtTm                        string  `json:"CrtTm"`
	CrtOrgNo                     string  `json:"CrtOrgNo"`
	ApprvSugstnCd                string  `json:"ApprvSugstnCd"`
	ApprvSugstnComnt             string  `json:"ApprvSugstnComnt"`
	ApprvTelrNo                  string  `json:"ApprvTelrNo"`
	ApprvTm                      string  `json:"ApprvTm"`
	ApprvOrgNo                   string  `json:"ApprvOrgNo"`
	IntrtFlotManrCd              string  `json:"IntrtFlotManrCd"`
	IntrtFlotRatio               float64  `json:"IntrtFlotRatio"`
	AdjRsn    		             string  `json:"AdjRsn"`
	CrtTelrName 	             string  `json:"CrtTelrName"`
	ApprvTelrName                string  `json:"ApprvTelrName"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
	PageNo                       int     `json:"PageNo"`
	PageRecCount                 int     `json:"PageRecCount"`
}

type DAILRSX8O struct {
	Records                      []DAILRSX8ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type DAILRSX8ORecords struct {
	IntrtAdjSn                   string  `json:"IntrtAdjSn"`
	CustNo                       string  `json:"CustNo"`
	CustName                     string  `json:"CustName"`
	CrdtAplySn                   string  `json:"CrdtAplySn"`
	IndvCrtfTypCd                string  `json:"IndvCrtfTypCd"`
	IndvCrtfNo                   string  `json:"IndvCrtfNo"`
	BaseIntrtTypCd               string  `json:"BaseIntrtTypCd"`
	CurCd                        string  `json:"CurCd"`
	IntrtNo                      string  `json:"IntrtNo"`
	IntrtFlotDrctCd              string  `json:"IntrtFlotDrctCd"`
	BpFlotVal                    float64  `json:"BpFlotVal"`
	KeprcdStusCd                 string  `json:"KeprcdStusCd"`
	CrtTelrNo                    string  `json:"CrtTelrNo"`
	CrtTm                        string  `json:"CrtTm"`
	CrtOrgNo                     string  `json:"CrtOrgNo"`
	ApprvSugstnCd                string  `json:"ApprvSugstnCd"`
	ApprvSugstnComnt             string  `json:"ApprvSugstnComnt"`
	ApprvTelrNo                  string  `json:"ApprvTelrNo"`
	ApprvTm                      string  `json:"ApprvTm"`
	ApprvOrgNo                   string  `json:"ApprvOrgNo"`
	IntrtFlotManrCd              string  `json:"IntrtFlotManrCd"`
	IntrtFlotRatio               float64  `json:"IntrtFlotRatio"`
	AdjRsn    		             string  `json:"AdjRsn"`
	CrtTelrName 	             string  `json:"CrtTelrName"`
	ApprvTelrName                string  `json:"ApprvTelrName"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
}

// @Desc Build request message
func (o *DAILRSX8RequestForm) PackRequest(dailrsx8I DAILRSX8I) (responseBody []byte, err error) {

	requestForm := DAILRSX8RequestForm{
		Form: []DAILRSX8IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRSX8I",
				},

				FormData: dailrsx8I,
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
func (o *DAILRSX8RequestForm) UnPackRequest(request []byte) (DAILRSX8I, error) {
	dailrsx8I := DAILRSX8I{}
	if err := json.Unmarshal(request, o); nil != err {
		return dailrsx8I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrsx8I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *DAILRSX8ResponseForm) PackResponse(dailrsx8O DAILRSX8O) (responseBody []byte, err error) {
	responseForm := DAILRSX8ResponseForm{
		Form: []DAILRSX8ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRSX8O",
				},
				FormData: dailrsx8O,
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
func (o *DAILRSX8ResponseForm) UnPackResponse(request []byte) (DAILRSX8O, error) {

	dailrsx8O := DAILRSX8O{}

	if err := json.Unmarshal(request, o); nil != err {
		return dailrsx8O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrsx8O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *DAILRSX8I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
