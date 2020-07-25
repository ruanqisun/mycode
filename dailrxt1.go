package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type DAILRXT1IDataForm struct {
	FormHead CommonFormHead
	FormData DAILRXT1I
}

type DAILRXT1ODataForm struct {
	FormHead CommonFormHead
	FormData DAILRXT1O
}

type DAILRXT1RequestForm struct {
	Form []DAILRXT1IDataForm
}

type DAILRXT1ResponseForm struct {
	Form []DAILRXT1ODataForm
}

type DAILRXT1I struct {
	FuncNo                       string  `json:"FuncNo"`
	OrdNo          		         int     `json:"OrdNo"`
	SoftProdtCd                  string  `json:"SoftProdtCd"`
	FuncUrlPath                  string  `json:"FuncUrlPath"`
	FuncLablEngNm                string  `json:"FuncLablEngNm"`
	FuncLablLoclLangugNm         string  `json:"FuncLablLoclLangugNm"`
	ButnKeyVal                   string  `json:"ButnKeyVal"`
	FuncRouteSeqNo               int     `json:"FuncRouteSeqNo"`
	LstoneBrthNodeFuncNo         string  `json:"LstoneBrthNodeFuncNo"`
	MgmtPageNextButnFlg          string  `json:"MgmtPageNextButnFlg"`
	SprFuncId                    string  `json:"SprFuncId"`
	NeedLgnAccessFlg             string  `json:"NeedLgnAccessFlg"`
	NeedAuthAccessFlg            string  `json:"NeedAuthAccessFlg"`
	FuncTypCd                    string  `json:"FuncTypCd"`
	MenuTypCd 			         string  `json:"MenuTypCd"`
	FuncValidFlg                 string  `json:"FuncValidFlg"`
	UserTaskNodeFlg              string  `json:"UserTaskNodeFlg"`
	Remrk                        string  `json:"Remrk"`
	MenuIconPath                 string  `json:"MenuIconPath"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
	PageNo                       int     `json:"PageNo"`
	PageRecCount                 int     `json:"PageRecCount"`
}

type DAILRXT1O struct {
	Records                      []DAILRXT1ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type DAILRXT1ORecords struct {
	FuncNo                       string  `json:"FuncNo"`
	OrdNo          		         int     `json:"OrdNo"`
	SoftProdtCd                  string  `json:"SoftProdtCd"`
	FuncUrlPath                  string  `json:"FuncUrlPath"`
	FuncLablEngNm                string  `json:"FuncLablEngNm"`
	FuncLablLoclLangugNm         string  `json:"FuncLablLoclLangugNm"`
	ButnKeyVal                   string  `json:"ButnKeyVal"`
	FuncRouteSeqNo               int     `json:"FuncRouteSeqNo"`
	LstoneBrthNodeFuncNo         string  `json:"LstoneBrthNodeFuncNo"`
	MgmtPageNextButnFlg          string  `json:"MgmtPageNextButnFlg"`
	SprFuncId                    string  `json:"SprFuncId"`
	NeedLgnAccessFlg             string  `json:"NeedLgnAccessFlg"`
	NeedAuthAccessFlg            string  `json:"NeedAuthAccessFlg"`
	FuncTypCd                    string  `json:"FuncTypCd"`
	MenuTypCd 			         string  `json:"MenuTypCd"`
	FuncValidFlg                 string  `json:"FuncValidFlg"`
	UserTaskNodeFlg              string  `json:"UserTaskNodeFlg"`
	Remrk                        string  `json:"Remrk"`
	MenuIconPath                 string  `json:"MenuIconPath"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
}

// @Desc Build request message
func (o *DAILRXT1RequestForm) PackRequest(dailrxt1I DAILRXT1I) (responseBody []byte, err error) {

	requestForm := DAILRXT1RequestForm{
		Form: []DAILRXT1IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRXT1I",
				},

				FormData: dailrxt1I,
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
func (o *DAILRXT1RequestForm) UnPackRequest(request []byte) (DAILRXT1I, error) {
	dailrxt1I := DAILRXT1I{}
	if err := json.Unmarshal(request, o); nil != err {
		return dailrxt1I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrxt1I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *DAILRXT1ResponseForm) PackResponse(dailrxt1O DAILRXT1O) (responseBody []byte, err error) {
	responseForm := DAILRXT1ResponseForm{
		Form: []DAILRXT1ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRXT1O",
				},
				FormData: dailrxt1O,
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
func (o *DAILRXT1ResponseForm) UnPackResponse(request []byte) (DAILRXT1O, error) {

	dailrxt1O := DAILRXT1O{}

	if err := json.Unmarshal(request, o); nil != err {
		return dailrxt1O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrxt1O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *DAILRXT1I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
