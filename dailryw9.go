package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type DAILRYW9IDataForm struct {
	FormHead CommonFormHead
	FormData DAILRYW9I
}

type DAILRYW9ODataForm struct {
	FormHead CommonFormHead
	FormData DAILRYW9O
}

type DAILRYW9RequestForm struct {
	Form []DAILRYW9IDataForm
}

type DAILRYW9ResponseForm struct {
	Form []DAILRYW9ODataForm
}

type DAILRYW9I struct {
	BizSn                  		  string  `json:"BizSn"`
	LprOrgNo                      string  `json:"LprOrgNo"`
	LprOrgNm                      string  `json:"LprOrgNm"`
	LoanProdtNo                   string  `json:"LoanProdtNo"`
	LoanProdtNm                   string  `json:"LoanProdtNm"`
	IntrtTypCd                    string  `json:"IntrtTypCd"`
	LprIntrtBpFlotCeilVal         float64  `json:"LprIntrtBpFlotCeilVal"`
	LprIntrtBpFlotFloorVal        float64  `json:"LprIntrtBpFlotFloorVal"`
	PbocBnchmkIntrtFlotCeilRatio  float64  `json:"PbocBnchmkIntrtFlotCeilRatio"`
	PbocBnchmkIntrtFlotFloorRatio float64  `json:"PbocBnchmkIntrtFlotFloorRatio"`
	QtaTypCd                      string  `json:"QtaTypCd"`
	QtaFlotCeilRatio              float64  `json:"QtaFlotCeilRatio"`
	QtaFlotFloorRatio             float64  `json:"QtaFlotFloorRatio"`
	KeprcdStusCd                  string  `json:"KeprcdStusCd"`
	CrtEmpnbr                     string  `json:"CrtEmpnbr"`
	CrtOrgNo                      string  `json:"CrtOrgNo"`
	CrtDt                         string  `json:"CrtDt"`
	FinlModfyDt                   string  `json:"FinlModfyDt"`
	FinlModfyTm                   string  `json:"FinlModfyTm"`
	FinlModfyOrgNo                string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo               string  `json:"FinlModfyTelrNo"`
	PageNo                        int     `json:"PageNo"`
	PageRecCount                  int     `json:"PageRecCount"`
}

type DAILRYW9O struct {
	Records                      []DAILRYW9ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type DAILRYW9ORecords struct {
	BizSn                  		  string  `json:"BizSn"`
	LprOrgNo                      string  `json:"LprOrgNo"`
	LprOrgNm                      string  `json:"LprOrgNm"`
	LoanProdtNo                   string  `json:"LoanProdtNo"`
	LoanProdtNm                   string  `json:"LoanProdtNm"`
	IntrtTypCd                    string  `json:"IntrtTypCd"`
	LprIntrtBpFlotCeilVal         float64  `json:"LprIntrtBpFlotCeilVal"`
	LprIntrtBpFlotFloorVal        float64  `json:"LprIntrtBpFlotFloorVal"`
	PbocBnchmkIntrtFlotCeilRatio  float64  `json:"PbocBnchmkIntrtFlotCeilRatio"`
	PbocBnchmkIntrtFlotFloorRatio float64  `json:"PbocBnchmkIntrtFlotFloorRatio"`
	QtaTypCd                      string  `json:"QtaTypCd"`
	QtaFlotCeilRatio              float64  `json:"QtaFlotCeilRatio"`
	QtaFlotFloorRatio             float64  `json:"QtaFlotFloorRatio"`
	KeprcdStusCd                  string  `json:"KeprcdStusCd"`
	CrtEmpnbr                     string  `json:"CrtEmpnbr"`
	CrtOrgNo                      string  `json:"CrtOrgNo"`
	CrtDt                         string  `json:"CrtDt"`
	FinlModfyDt                   string  `json:"FinlModfyDt"`
	FinlModfyTm                   string  `json:"FinlModfyTm"`
	FinlModfyOrgNo                string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo               string  `json:"FinlModfyTelrNo"`
}

// @Desc Build request message
func (o *DAILRYW9RequestForm) PackRequest(dailryw9I DAILRYW9I) (responseBody []byte, err error) {

	requestForm := DAILRYW9RequestForm{
		Form: []DAILRYW9IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRYW9I",
				},

				FormData: dailryw9I,
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
func (o *DAILRYW9RequestForm) UnPackRequest(request []byte) (DAILRYW9I, error) {
	dailryw9I := DAILRYW9I{}
	if err := json.Unmarshal(request, o); nil != err {
		return dailryw9I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailryw9I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *DAILRYW9ResponseForm) PackResponse(dailryw9O DAILRYW9O) (responseBody []byte, err error) {
	responseForm := DAILRYW9ResponseForm{
		Form: []DAILRYW9ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRYW9O",
				},
				FormData: dailryw9O,
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
func (o *DAILRYW9ResponseForm) UnPackResponse(request []byte) (DAILRYW9O, error) {

	dailryw9O := DAILRYW9O{}

	if err := json.Unmarshal(request, o); nil != err {
		return dailryw9O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailryw9O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *DAILRYW9I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
