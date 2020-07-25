package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type DAILRYW7IDataForm struct {
	FormHead CommonFormHead
	FormData DAILRYW7I
}

type DAILRYW7ODataForm struct {
	FormHead CommonFormHead
	FormData DAILRYW7O
}

type DAILRYW7RequestForm struct {
	Form []DAILRYW7IDataForm
}

type DAILRYW7ResponseForm struct {
	Form []DAILRYW7ODataForm
}

type DAILRYW7I struct {
	LoanProdtNo                   string  `json:"LoanProdtNo"`
	EfftDt                        string  `json:"EfftDt"`
	QtaAdjRangDegreeCeilVal       float64 `json:"QtaAdjRangDegreeCeilVal"`
	QtaAdjRangDegreeFloorVal      int     `json:"QtaAdjRangDegreeFloorVal"`
	FixdIntrtFlotCeilRatio        float64 `json:"FixdIntrtFlotCeilRatio"`
	FixdIntrtFlotFloorRatio       float64 `json:"FixdIntrtFlotFloorRatio"`
	LprIntrtBpFlotCeilVal         float64 `json:"LprIntrtBpFlotCeilVal"`
	LprIntrtBpFlotFloorVal        float64 `json:"LprIntrtBpFlotFloorVal"`
	CrtTelrNo                     string  `json:"CrtTelrNo"`
	CrtTm                         string  `json:"CrtTm"`
	CrtOrgNo                      string  `json:"CrtOrgNo"`
	CrdtElentParaStusCd           string  `json:"CrdtElentParaStusCd"`
	FinlModfyDt                   string  `json:"FinlModfyDt"`
	FinlModfyTm                   string  `json:"FinlModfyTm"`
	FinlModfyOrgNo                string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo               string  `json:"FinlModfyTelrNo"`
	PageNo                        int     `json:"PageNo"`
	PageRecCount                  int     `json:"PageRecCount"`
}

type DAILRYW7O struct {
	Records                      []DAILRYW7ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type DAILRYW7ORecords struct {
	LoanProdtNo                   string  `json:"LoanProdtNo"`
	EfftDt                        string  `json:"EfftDt"`
	QtaAdjRangDegreeCeilVal       float64 `json:"QtaAdjRangDegreeCeilVal"`
	QtaAdjRangDegreeFloorVal      int     `json:"QtaAdjRangDegreeFloorVal"`
	FixdIntrtFlotCeilRatio        float64 `json:"FixdIntrtFlotCeilRatio"`
	FixdIntrtFlotFloorRatio       float64 `json:"FixdIntrtFlotFloorRatio"`
	LprIntrtBpFlotCeilVal         float64 `json:"LprIntrtBpFlotCeilVal"`
	LprIntrtBpFlotFloorVal        float64 `json:"LprIntrtBpFlotFloorVal"`
	CrtTelrNo                     string  `json:"CrtTelrNo"`
	CrtTm                         string  `json:"CrtTm"`
	CrtOrgNo                      string  `json:"CrtOrgNo"`
	CrdtElentParaStusCd           string  `json:"CrdtElentParaStusCd"`
	FinlModfyDt                   string  `json:"FinlModfyDt"`
	FinlModfyTm                   string  `json:"FinlModfyTm"`
	FinlModfyOrgNo                string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo               string  `json:"FinlModfyTelrNo"`
}

// @Desc Build request message
func (o *DAILRYW7RequestForm) PackRequest(dailryw7I DAILRYW7I) (responseBody []byte, err error) {

	requestForm := DAILRYW7RequestForm{
		Form: []DAILRYW7IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRYW7I",
				},

				FormData: dailryw7I,
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
func (o *DAILRYW7RequestForm) UnPackRequest(request []byte) (DAILRYW7I, error) {
	dailryw7I := DAILRYW7I{}
	if err := json.Unmarshal(request, o); nil != err {
		return dailryw7I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailryw7I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *DAILRYW7ResponseForm) PackResponse(dailryw7O DAILRYW7O) (responseBody []byte, err error) {
	responseForm := DAILRYW7ResponseForm{
		Form: []DAILRYW7ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRYW7O",
				},
				FormData: dailryw7O,
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
func (o *DAILRYW7ResponseForm) UnPackResponse(request []byte) (DAILRYW7O, error) {

	dailryw7O := DAILRYW7O{}

	if err := json.Unmarshal(request, o); nil != err {
		return dailryw7O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailryw7O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *DAILRYW7I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
