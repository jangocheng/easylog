package consts

type Json struct {
	message string
	code int
	data interface{}
	ret bool
}

func (this *Json) SetMessage(message string) {
	this.message = message
}

func (this *Json) SetCode(code int) {
	this.code = code
}

func (this *Json) SetData(data interface{}) {
	this.data = data
}

func (this *Json) Set(code int,message string) {
	this.SetMessage(message)
	this.SetCode(code)
}

func (this *Json) VendorError(code int,message string) interface{} {
	return map[string]interface{}{"ret":false,"code":code,"message":message,"data":nil}
}

func (this *Json) VendorOk() interface{} {
	return map[string]interface{}{"ret":true,"code":this.code,"message":this.message,"data":this.data}
}
