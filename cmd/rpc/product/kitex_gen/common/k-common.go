// Code generated by Kitex v0.8.0. DO NOT EDIT.

package common

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"

	"github.com/apache/thrift/lib/go/thrift"

	"github.com/cloudwego/kitex/pkg/protocol/bthrift"
)

// unused protection
var (
	_ = fmt.Formatter(nil)
	_ = (*bytes.Buffer)(nil)
	_ = (*strings.Builder)(nil)
	_ = reflect.Type(nil)
	_ = thrift.TProtocol(nil)
	_ = bthrift.BinaryWriter(nil)
)

func (p *Req) FastRead(buf []byte) (int, error) {
	var err error
	var offset int
	var l int
	var fieldTypeId thrift.TType
	var fieldId int16
	var issetId bool = false
	_, l, err = bthrift.Binary.ReadStructBegin(buf)
	offset += l
	if err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, l, err = bthrift.Binary.ReadFieldBegin(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I32 {
				l, err = p.FastReadField1(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
				issetId = true
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		default:
			l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
			offset += l
			if err != nil {
				goto SkipFieldError
			}
		}

		l, err = bthrift.Binary.ReadFieldEnd(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldEndError
		}
	}
	l, err = bthrift.Binary.ReadStructEnd(buf[offset:])
	offset += l
	if err != nil {
		goto ReadStructEndError
	}

	if !issetId {
		fieldId = 1
		goto RequiredFieldNotSetError
	}
	return offset, nil
ReadStructBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_Req[fieldId]), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
ReadFieldEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
RequiredFieldNotSetError:
	return offset, thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("required field %s is not set", fieldIDToName_Req[fieldId]))
}

func (p *Req) FastReadField1(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadI32(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.Id = v

	}
	return offset, nil
}

// for compatibility
func (p *Req) FastWrite(buf []byte) int {
	return 0
}

func (p *Req) FastWriteNocopy(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteStructBegin(buf[offset:], "Req")
	if p != nil {
		offset += p.fastWriteField1(buf[offset:], binaryWriter)
	}
	offset += bthrift.Binary.WriteFieldStop(buf[offset:])
	offset += bthrift.Binary.WriteStructEnd(buf[offset:])
	return offset
}

func (p *Req) BLength() int {
	l := 0
	l += bthrift.Binary.StructBeginLength("Req")
	if p != nil {
		l += p.field1Length()
	}
	l += bthrift.Binary.FieldStopLength()
	l += bthrift.Binary.StructEndLength()
	return l
}

func (p *Req) fastWriteField1(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "Id", thrift.I32, 1)
	offset += bthrift.Binary.WriteI32(buf[offset:], p.Id)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *Req) field1Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("Id", thrift.I32, 1)
	l += bthrift.Binary.I32Length(p.Id)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *Empty) FastRead(buf []byte) (int, error) {
	var err error
	var offset int
	var l int
	var fieldTypeId thrift.TType
	var fieldId int16
	_, l, err = bthrift.Binary.ReadStructBegin(buf)
	offset += l
	if err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, l, err = bthrift.Binary.ReadFieldBegin(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
		offset += l
		if err != nil {
			goto SkipFieldError
		}

		l, err = bthrift.Binary.ReadFieldEnd(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldEndError
		}
	}
	l, err = bthrift.Binary.ReadStructEnd(buf[offset:])
	offset += l
	if err != nil {
		goto ReadStructEndError
	}

	return offset, nil
ReadStructBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
ReadFieldEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

// for compatibility
func (p *Empty) FastWrite(buf []byte) int {
	return 0
}

func (p *Empty) FastWriteNocopy(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteStructBegin(buf[offset:], "Empty")
	if p != nil {
	}
	offset += bthrift.Binary.WriteFieldStop(buf[offset:])
	offset += bthrift.Binary.WriteStructEnd(buf[offset:])
	return offset
}

func (p *Empty) BLength() int {
	l := 0
	l += bthrift.Binary.StructBeginLength("Empty")
	if p != nil {
	}
	l += bthrift.Binary.FieldStopLength()
	l += bthrift.Binary.StructEndLength()
	return l
}

func (p *PageFilter) FastRead(buf []byte) (int, error) {
	var err error
	var offset int
	var l int
	var fieldTypeId thrift.TType
	var fieldId int16
	_, l, err = bthrift.Binary.ReadStructBegin(buf)
	offset += l
	if err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, l, err = bthrift.Binary.ReadFieldBegin(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I16 {
				l, err = p.FastReadField1(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 2:
			if fieldTypeId == thrift.I16 {
				l, err = p.FastReadField2(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		default:
			l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
			offset += l
			if err != nil {
				goto SkipFieldError
			}
		}

		l, err = bthrift.Binary.ReadFieldEnd(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldEndError
		}
	}
	l, err = bthrift.Binary.ReadStructEnd(buf[offset:])
	offset += l
	if err != nil {
		goto ReadStructEndError
	}

	return offset, nil
ReadStructBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_PageFilter[fieldId]), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
ReadFieldEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *PageFilter) FastReadField1(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadI16(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
		p.Page = &v

	}
	return offset, nil
}

func (p *PageFilter) FastReadField2(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadI16(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
		p.Limit = &v

	}
	return offset, nil
}

// for compatibility
func (p *PageFilter) FastWrite(buf []byte) int {
	return 0
}

func (p *PageFilter) FastWriteNocopy(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteStructBegin(buf[offset:], "PageFilter")
	if p != nil {
		offset += p.fastWriteField1(buf[offset:], binaryWriter)
		offset += p.fastWriteField2(buf[offset:], binaryWriter)
	}
	offset += bthrift.Binary.WriteFieldStop(buf[offset:])
	offset += bthrift.Binary.WriteStructEnd(buf[offset:])
	return offset
}

func (p *PageFilter) BLength() int {
	l := 0
	l += bthrift.Binary.StructBeginLength("PageFilter")
	if p != nil {
		l += p.field1Length()
		l += p.field2Length()
	}
	l += bthrift.Binary.FieldStopLength()
	l += bthrift.Binary.StructEndLength()
	return l
}

func (p *PageFilter) fastWriteField1(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	if p.IsSetPage() {
		offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "Page", thrift.I16, 1)
		offset += bthrift.Binary.WriteI16(buf[offset:], *p.Page)

		offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	}
	return offset
}

func (p *PageFilter) fastWriteField2(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	if p.IsSetLimit() {
		offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "Limit", thrift.I16, 2)
		offset += bthrift.Binary.WriteI16(buf[offset:], *p.Limit)

		offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	}
	return offset
}

func (p *PageFilter) field1Length() int {
	l := 0
	if p.IsSetPage() {
		l += bthrift.Binary.FieldBeginLength("Page", thrift.I16, 1)
		l += bthrift.Binary.I16Length(*p.Page)

		l += bthrift.Binary.FieldEndLength()
	}
	return l
}

func (p *PageFilter) field2Length() int {
	l := 0
	if p.IsSetLimit() {
		l += bthrift.Binary.FieldBeginLength("Limit", thrift.I16, 2)
		l += bthrift.Binary.I16Length(*p.Limit)

		l += bthrift.Binary.FieldEndLength()
	}
	return l
}