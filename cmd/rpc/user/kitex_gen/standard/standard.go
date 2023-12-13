// Code generated by thriftgo (0.3.3). DO NOT EDIT.

package standard

import (
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"strings"
)

type Resp struct {
	Success bool   `thrift:"Success,1,required" frugal:"1,required,bool" json:"Success"`
	Code    int16  `thrift:"Code,2,required" frugal:"2,required,i16" json:"Code"`
	Message string `thrift:"Message,3,required" frugal:"3,required,string" json:"Message"`
}

func NewResp() *Resp {
	return &Resp{}
}

func (p *Resp) InitDefault() {
	*p = Resp{}
}

func (p *Resp) GetSuccess() (v bool) {
	return p.Success
}

func (p *Resp) GetCode() (v int16) {
	return p.Code
}

func (p *Resp) GetMessage() (v string) {
	return p.Message
}
func (p *Resp) SetSuccess(val bool) {
	p.Success = val
}
func (p *Resp) SetCode(val int16) {
	p.Code = val
}
func (p *Resp) SetMessage(val string) {
	p.Message = val
}

var fieldIDToName_Resp = map[int16]string{
	1: "Success",
	2: "Code",
	3: "Message",
}

func (p *Resp) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16
	var issetSuccess bool = false
	var issetCode bool = false
	var issetMessage bool = false

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.BOOL {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
				issetSuccess = true
				break
			}
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 2:
			if fieldTypeId == thrift.I16 {
				if err = p.ReadField2(iprot); err != nil {
					goto ReadFieldError
				}
				issetCode = true
				break
			}
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 3:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField3(iprot); err != nil {
					goto ReadFieldError
				}
				issetMessage = true
				break
			}
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	if !issetSuccess {
		fieldId = 1
		goto RequiredFieldNotSetError
	}

	if !issetCode {
		fieldId = 2
		goto RequiredFieldNotSetError
	}

	if !issetMessage {
		fieldId = 3
		goto RequiredFieldNotSetError
	}
	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_Resp[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
RequiredFieldNotSetError:
	return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("required field %s is not set", fieldIDToName_Resp[fieldId]))
}

func (p *Resp) ReadField1(iprot thrift.TProtocol) error {

	if v, err := iprot.ReadBool(); err != nil {
		return err
	} else {
		p.Success = v
	}
	return nil
}
func (p *Resp) ReadField2(iprot thrift.TProtocol) error {

	if v, err := iprot.ReadI16(); err != nil {
		return err
	} else {
		p.Code = v
	}
	return nil
}
func (p *Resp) ReadField3(iprot thrift.TProtocol) error {

	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.Message = v
	}
	return nil
}

func (p *Resp) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("Resp"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
		if err = p.writeField2(oprot); err != nil {
			fieldId = 2
			goto WriteFieldError
		}
		if err = p.writeField3(oprot); err != nil {
			fieldId = 3
			goto WriteFieldError
		}
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *Resp) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("Success", thrift.BOOL, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteBool(p.Success); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}
func (p *Resp) writeField2(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("Code", thrift.I16, 2); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI16(p.Code); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 end error: ", p), err)
}
func (p *Resp) writeField3(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("Message", thrift.STRING, 3); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.Message); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 end error: ", p), err)
}

func (p *Resp) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Resp(%+v)", *p)
}

func (p *Resp) DeepEqual(ano *Resp) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Success) {
		return false
	}
	if !p.Field2DeepEqual(ano.Code) {
		return false
	}
	if !p.Field3DeepEqual(ano.Message) {
		return false
	}
	return true
}

func (p *Resp) Field1DeepEqual(src bool) bool {

	if p.Success != src {
		return false
	}
	return true
}
func (p *Resp) Field2DeepEqual(src int16) bool {

	if p.Code != src {
		return false
	}
	return true
}
func (p *Resp) Field3DeepEqual(src string) bool {

	if strings.Compare(p.Message, src) != 0 {
		return false
	}
	return true
}

type Req struct {
	Id int32 `thrift:"Id,1,required" frugal:"1,required,i32" json:"Id"`
}

func NewReq() *Req {
	return &Req{}
}

func (p *Req) InitDefault() {
	*p = Req{}
}

func (p *Req) GetId() (v int32) {
	return p.Id
}
func (p *Req) SetId(val int32) {
	p.Id = val
}

var fieldIDToName_Req = map[int16]string{
	1: "Id",
}

func (p *Req) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16
	var issetId bool = false

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I32 {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
				issetId = true
				break
			}
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	if !issetId {
		fieldId = 1
		goto RequiredFieldNotSetError
	}
	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_Req[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
RequiredFieldNotSetError:
	return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("required field %s is not set", fieldIDToName_Req[fieldId]))
}

func (p *Req) ReadField1(iprot thrift.TProtocol) error {

	if v, err := iprot.ReadI32(); err != nil {
		return err
	} else {
		p.Id = v
	}
	return nil
}

func (p *Req) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("Req"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *Req) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("Id", thrift.I32, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI32(p.Id); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *Req) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Req(%+v)", *p)
}

func (p *Req) DeepEqual(ano *Req) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Id) {
		return false
	}
	return true
}

func (p *Req) Field1DeepEqual(src int32) bool {

	if p.Id != src {
		return false
	}
	return true
}
