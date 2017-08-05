package testdata

// NOTE: THIS FILE WAS PRODUCED BY THE
// GREENPACK CODE GENERATION TOOL (github.com/glycerine/greenpack)
// DO NOT EDIT

import (
	"github.com/glycerine/greenpack/msgp"
)

// DecodeMsg implements msgp.Decodable
// We treat empty fields as if we read a Nil from the wire.
func (z *OmitClueTestStruct) DecodeMsg(dc *msgp.Reader) (err error) {
	var sawTopNil bool
	if dc.IsNil() {
		sawTopNil = true
		err = dc.ReadNil()
		if err != nil {
			return
		}
		dc.PushAlwaysNil()
	}

	var field []byte
	_ = field
	const maxFields0zzvc1 = 2

	// -- templateDecodeMsg starts here--
	var totalEncodedFields0zzvc1 uint32
	totalEncodedFields0zzvc1, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft0zzvc1 := totalEncodedFields0zzvc1
	missingFieldsLeft0zzvc1 := maxFields0zzvc1 - totalEncodedFields0zzvc1

	var nextMiss0zzvc1 int32 = -1
	var found0zzvc1 [maxFields0zzvc1]bool
	var curField0zzvc1 string

doneWithStruct0zzvc1:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0zzvc1 > 0 || missingFieldsLeft0zzvc1 > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0zzvc1, missingFieldsLeft0zzvc1, msgp.ShowFound(found0zzvc1[:]), decodeMsgFieldOrder0zzvc1)
		if encodedFieldsLeft0zzvc1 > 0 {
			encodedFieldsLeft0zzvc1--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField0zzvc1 = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss0zzvc1 < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss0zzvc1 = 0
			}
			for nextMiss0zzvc1 < maxFields0zzvc1 && (found0zzvc1[nextMiss0zzvc1] || decodeMsgFieldSkip0zzvc1[nextMiss0zzvc1]) {
				nextMiss0zzvc1++
			}
			if nextMiss0zzvc1 == maxFields0zzvc1 {
				// filled all the empty fields!
				break doneWithStruct0zzvc1
			}
			missingFieldsLeft0zzvc1--
			curField0zzvc1 = decodeMsgFieldOrder0zzvc1[nextMiss0zzvc1]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0zzvc1)
		switch curField0zzvc1 {
		// -- templateDecodeMsg ends here --

		case "S":
			found0zzvc1[0] = true
			z.S, err = dc.ReadString()
			if err != nil {
				return
			}
		case "N":
			found0zzvc1[1] = true
			z.N, err = dc.ReadInt64()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	if nextMiss0zzvc1 != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	if p, ok := interface{}(z).(msgp.PostLoad); ok {
		p.PostLoadHook()
	}

	return
}

// fields of OmitClueTestStruct
var decodeMsgFieldOrder0zzvc1 = []string{"S", "N"}

var decodeMsgFieldSkip0zzvc1 = []bool{false, false}

// fieldsNotEmpty supports omitempty tags
func (z OmitClueTestStruct) fieldsNotEmpty(isempty []bool) uint32 {
	if len(isempty) == 0 {
		return 2
	}
	var fieldsInUse uint32 = 2
	isempty[0] = (len(z.S) == 0) // string, omitempty
	if isempty[0] {
		fieldsInUse--
	}
	isempty[1] = (z.N == 0) // number, omitempty
	if isempty[1] {
		fieldsInUse--
	}

	return fieldsInUse
}

// EncodeMsg implements msgp.Encodable
func (z OmitClueTestStruct) EncodeMsg(en *msgp.Writer) (err error) {
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

	// honor the omitempty tags
	var empty_zwtm2 [2]bool
	fieldsInUse_zdle3 := z.fieldsNotEmpty(empty_zwtm2[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zdle3)
	if err != nil {
		return err
	}

	if !empty_zwtm2[0] {
		// write "S"
		err = en.Append(0xa1, 0x53)
		if err != nil {
			return err
		}
		err = en.WriteString(z.S)
		if err != nil {
			return
		}
	}

	if !empty_zwtm2[1] {
		// write "N"
		err = en.Append(0xa1, 0x4e)
		if err != nil {
			return err
		}
		err = en.WriteInt64(z.N)
		if err != nil {
			return
		}
	}

	return
}

// MarshalMsg implements msgp.Marshaler
func (z OmitClueTestStruct) MarshalMsg(b []byte) (o []byte, err error) {
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

	o = msgp.Require(b, z.Msgsize())

	// honor the omitempty tags
	var empty [2]bool
	fieldsInUse := z.fieldsNotEmpty(empty[:])
	o = msgp.AppendMapHeader(o, fieldsInUse)

	if !empty[0] {
		// string "S"
		o = append(o, 0xa1, 0x53)
		o = msgp.AppendString(o, z.S)
	}

	if !empty[1] {
		// string "N"
		o = append(o, 0xa1, 0x4e)
		o = msgp.AppendInt64(o, z.N)
	}

	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *OmitClueTestStruct) UnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.UnmarshalMsgWithCfg(bts, nil)
}
func (z *OmitClueTestStruct) UnmarshalMsgWithCfg(bts []byte, cfg *msgp.RuntimeConfig) (o []byte, err error) {
	var nbs msgp.NilBitsStack
	nbs.Init(cfg)
	var sawTopNil bool
	if msgp.IsNil(bts) {
		sawTopNil = true
		bts = nbs.PushAlwaysNil(bts[1:])
	}

	var field []byte
	_ = field
	const maxFields1zhkg4 = 2

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields1zhkg4 uint32
	if !nbs.AlwaysNil {
		totalEncodedFields1zhkg4, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft1zhkg4 := totalEncodedFields1zhkg4
	missingFieldsLeft1zhkg4 := maxFields1zhkg4 - totalEncodedFields1zhkg4

	var nextMiss1zhkg4 int32 = -1
	var found1zhkg4 [maxFields1zhkg4]bool
	var curField1zhkg4 string

doneWithStruct1zhkg4:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft1zhkg4 > 0 || missingFieldsLeft1zhkg4 > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft1zhkg4, missingFieldsLeft1zhkg4, msgp.ShowFound(found1zhkg4[:]), unmarshalMsgFieldOrder1zhkg4)
		if encodedFieldsLeft1zhkg4 > 0 {
			encodedFieldsLeft1zhkg4--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField1zhkg4 = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss1zhkg4 < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss1zhkg4 = 0
			}
			for nextMiss1zhkg4 < maxFields1zhkg4 && (found1zhkg4[nextMiss1zhkg4] || unmarshalMsgFieldSkip1zhkg4[nextMiss1zhkg4]) {
				nextMiss1zhkg4++
			}
			if nextMiss1zhkg4 == maxFields1zhkg4 {
				// filled all the empty fields!
				break doneWithStruct1zhkg4
			}
			missingFieldsLeft1zhkg4--
			curField1zhkg4 = unmarshalMsgFieldOrder1zhkg4[nextMiss1zhkg4]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField1zhkg4)
		switch curField1zhkg4 {
		// -- templateUnmarshalMsg ends here --

		case "S":
			found1zhkg4[0] = true
			z.S, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "N":
			found1zhkg4[1] = true
			z.N, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	if nextMiss1zhkg4 != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	if p, ok := interface{}(z).(msgp.PostLoad); ok {
		p.PostLoadHook()
	}

	return
}

// fields of OmitClueTestStruct
var unmarshalMsgFieldOrder1zhkg4 = []string{"S", "N"}

var unmarshalMsgFieldSkip1zhkg4 = []bool{false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z OmitClueTestStruct) Msgsize() (s int) {
	s = 1 + 12 + msgp.StringPrefixSize + len(z.S) + 12 + msgp.Int64Size
	return
}
