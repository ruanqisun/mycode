package models

import (
	"git.forms.io/legobank/legoapp/constant"
"git.forms.io/legobank/legoapp/errors"
"git.forms.io/universe/common/json"
)

type DAILRED1IDataForm struct {
	FormHead CommonFormHead
	FormData DAILRED1I
}

type DAILRED1ODataForm struct {
	FormHead CommonFormHead
	FormData DAILRED1O
}

type DAILRED1RequestForm struct {
	Form []DAILRED1IDataForm
}

type DAILRED1ResponseForm struct {
	Form []DAILRED1ODataForm
}

type DAILRED1I struct {
	CrdtAplySn                   string  `json:"CrdtAplySn"`
	CustNo                       string  `json:"CustNo"`
	LoanProdtVersNo              string  `json:"LoanProdtVersNo"`
	QtaFinlAdjDt                 string  `json:"QtaFinlAdjDt"`
	QtaEfftDt                    string  `json:"QtaEfftDt"`
	QtaInvalidDt                 string  `json:"QtaInvalidDt"`
	PmitOvrqtMakelnFlg           string  `json:"PmitOvrqtMakelnFlg"`
	OvrqtMakelnRatio             float64 `json:"OvrqtMakelnRatio"`
	LoanGuarManrCd               float64 `json:"LoanGuarManrCd"`
	LoanProdtNo                  string  `json:"LoanProdtNo"`
	QtaStusCd                    string  `json:"QtaStusCd"`
	QtaMltpDistrFlg              string  `json:"QtaMltpDistrFlg"`
	RevlQtaFlg                   string  `json:"RevlQtaFlg"`
	QtaCanGuarThirdPtyFlg        string  `json:"QtaCanGuarThirdPtyFlg"`
	IndvCrtfTypCd                string  `json:"IndvCrtfTypCd"`
	IndvCrtfNo                   string  `json:"IndvCrtfNo"`
	CustName                     string  `json:"CustName"`
	CurCd                        string  `json:"CurCd"`
	InitQta                      float64 `json:"InitQta"`
	CurrCrdtQta                  float64 `json:"CurrCrdtQta"`
	FrzQta                       float64 `json:"FrzQta"`
	AlrdyUseQta                  float64 `json:"AlrdyUseQta"`
	QtaFsttmAprvlDt              string  `json:"QtaFsttmAprvlDt"`
	QtaMatrDt                    string  `json:"QtaMatrDt"`
	QtaGrtOrgNo                  string  `json:"QtaGrtOrgNo"`
	CrtDt                        string  `json:"CrtDt"`
	CrtEmpnbr                    string  `json:"CrtEmpnbr"`
	CrtOrgNo                     string  `json:"CrtOrgNo"`
	ArtgclAdjTms                 int     `json:"ArtgclAdjTms"`
	LsttmQtaAdjManrCd            string  `json:"LsttmQtaAdjManrCd"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type DAILRED1O struct {
	Records                      []DAILRED1ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type DAILRED1ORecords struct {
	CrdtAplySn                   string  `json:"CrdtAplySn"`
	CustNo                       string  `json:"CustNo"`
	LoanProdtVersNo              string  `json:"LoanProdtVersNo"`
	QtaFinlAdjDt                 string  `json:"QtaFinlAdjDt"`
	QtaEfftDt                    string  `json:"QtaEfftDt"`
	QtaInvalidDt                 string  `json:"QtaInvalidDt"`
	PmitOvrqtMakelnFlg           string  `json:"PmitOvrqtMakelnFlg"`
	OvrqtMakelnRatio             float64 `json:"OvrqtMakelnRatio"`
	LoanGuarManrCd               float64 `json:"LoanGuarManrCd"`
	LoanProdtNo                  string  `json:"LoanProdtNo"`
	QtaStusCd                    string  `json:"QtaStusCd"`
	QtaMltpDistrFlg              string  `json:"QtaMltpDistrFlg"`
	RevlQtaFlg                   string  `json:"RevlQtaFlg"`
	QtaCanGuarThirdPtyFlg        string  `json:"QtaCanGuarThirdPtyFlg"`
	IndvCrtfTypCd                string  `json:"IndvCrtfTypCd"`
	IndvCrtfNo                   string  `json:"IndvCrtfNo"`
	CustName                     string  `json:"CustName"`
	CurCd                        string  `json:"CurCd"`
	InitQta                      float64 `json:"InitQta"`
	CurrCrdtQta                  float64 `json:"CurrCrdtQta"`
	FrzQta                       float64 `json:"FrzQta"`
	AlrdyUseQta                  float64 `json:"AlrdyUseQta"`
	QtaFsttmAprvlDt              string  `json:"QtaFsttmAprvlDt"`
	QtaMatrDt                    string  `json:"QtaMatrDt"`
	QtaGrtOrgNo                  string  `json:"QtaGrtOrgNo"`
	CrtDt                        string  `json:"CrtDt"`
	CrtEmpnbr                    string  `json:"CrtEmpnbr"`
	CrtOrgNo                     string  `json:"CrtOrgNo"`
	ArtgclAdjTms                 int     `json:"ArtgclAdjTms"`
	LsttmQtaAdjManrCd            string  `json:"LsttmQtaAdjManrCd"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
}

// @Desc Build request message
func (o *DAILRED1RequestForm) PackRequest(dailred1I DAILRED1I) (responseBody []byte, err error) {

	requestForm := DAILRED1RequestForm{
		Form: []DAILRED1IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRED1I",
				},

				FormData: dailred1I,
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
func (o *DAILRED1RequestForm) UnPackRequest(request []byte) (DAILRED1I, error) {
	dailred1I := DAILRED1I{}
	if err := json.Unmarshal(request, o); nil != err {
		return dailred1I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailred1I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *DAILRED1ResponseForm) PackResponse(dailred1O DAILRED1O) (responseBody []byte, err error) {
	responseForm := DAILRED1ResponseForm{
		Form: []DAILRED1ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRED1O",
				},
				FormData: dailred1O,
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
func (o *DAILRED1ResponseForm) UnPackResponse(request []byte) (DAILRED1O, error) {

	dailred1O := DAILRED1O{}

	if err := json.Unmarshal(request, o); nil != err {
		return dailred1O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailred1O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *DAILRED1I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
