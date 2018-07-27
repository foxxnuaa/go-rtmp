//
// Copyright (c) 2018- yutopp (yutopp@gmail.com)
//
// Distributed under the Boost Software License, Version 1.0. (See accompanying
// file LICENSE_1_0.txt or copy at  https://www.boost.org/LICENSE_1_0.txt)
//

package message

//
type NetStreamPublish struct {
	CommandObject  interface{}
	PublishingName string
	PublishingType string
}

func (t *NetStreamPublish) FromArgs(args ...interface{}) error {
	//command := args[0] // will be nil
	t.PublishingName = args[1].(string)
	t.PublishingType = args[2].(string)

	return nil
}

func (t *NetStreamPublish) ToArgs(ty AMFType) ([]interface{}, error) {
	panic("Not implemented")
}

//
type NetStreamOnStatus struct {
	InfoObject NetStreamOnStatusInfoObject
}

type NetStreamOnStatusInfoObject struct {
	Level       string // TODO: fix
	Code        string // TODO: fix
	Description string
}

func (t *NetStreamOnStatus) FromArgs(args ...interface{}) error {
	panic("Not implemented")
}

func (t *NetStreamOnStatus) ToArgs(ty AMFType) ([]interface{}, error) {
	info := make(map[string]interface{})
	info["level"] = t.InfoObject.Level
	info["code"] = t.InfoObject.Code
	info["description"] = t.InfoObject.Description

	return []interface{}{
		nil, // Always nil
		info,
	}, nil
}

//
type NetStreamDeleteStream struct {
	StreamID uint32
}

func (t *NetStreamDeleteStream) FromArgs(args ...interface{}) error {
	// args[0] is unknown, ignore
	t.StreamID = args[1].(uint32)

	return nil
}

func (t *NetStreamDeleteStream) ToArgs(ty AMFType) ([]interface{}, error) {
	panic("Not implemented")
}

//
type NetStreamFCPublish struct {
	StreamName string
}

func (t *NetStreamFCPublish) FromArgs(args ...interface{}) error {
	// args[0] is unknown, ignore
	t.StreamName = args[1].(string)

	return nil
}

func (t *NetStreamFCPublish) ToArgs(ty AMFType) ([]interface{}, error) {
	return []interface{}{
		nil, // no command object
		t.StreamName,
	}, nil
}

//
type NetStreamFCUnpublish struct {
	StreamName string
}

func (t *NetStreamFCUnpublish) FromArgs(args ...interface{}) error {
	// args[0] is unknown, ignore
	t.StreamName = args[1].(string)

	return nil
}

func (t *NetStreamFCUnpublish) ToArgs(ty AMFType) ([]interface{}, error) {
	return []interface{}{
		nil, // no command object
		t.StreamName,
	}, nil
}

//
type NetStreamSetDataFrame struct {
	Payload []byte
}

func (t *NetStreamSetDataFrame) FromArgs(args ...interface{}) error {
	t.Payload = args[0].([]byte)

	return nil
}

func (t *NetStreamSetDataFrame) ToArgs(ty AMFType) ([]interface{}, error) {
	panic("Not implemented")
}
