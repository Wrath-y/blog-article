package resp

import (
	"article/infrastructure/util/errcode"
	"article/infrastructure/util/util/highperf"
	"article/interfaces/proto"
	"bytes"
	"encoding/json"
)

func Success(data any) (*proto.Response, error) {
	if data == nil {
		data = struct{}{}
	}
	buf := bytes.NewBuffer(nil)
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	if err := enc.Encode(data); err != nil {
		return nil, err
	}

	return &proto.Response{
		Code: 0,
		Msg:  "success",
		Data: highperf.Bytes2str(buf.Bytes()),
	}, nil
}

func FailWithErrCode(err *errcode.ErrCode) (*proto.Response, error) {
	return &proto.Response{
		Code: err.Code,
		Msg:  err.Msg,
	}, err
}
