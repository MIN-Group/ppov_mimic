package MetaData

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *GenesisTransaction) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "WorkerNum":
			z.WorkerNum, err = dc.ReadInt()
			if err != nil {
				err = msgp.WrapError(err, "WorkerNum")
				return
			}
		case "VotedNum":
			z.VotedNum, err = dc.ReadInt()
			if err != nil {
				err = msgp.WrapError(err, "VotedNum")
				return
			}
		case "BlockGroupPerCycle":
			z.BlockGroupPerCycle, err = dc.ReadInt()
			if err != nil {
				err = msgp.WrapError(err, "BlockGroupPerCycle")
				return
			}
		case "Tcut":
			z.Tcut, err = dc.ReadFloat64()
			if err != nil {
				err = msgp.WrapError(err, "Tcut")
				return
			}
		case "WorkerPubList":
			var zb0002 uint32
			zb0002, err = dc.ReadMapHeader()
			if err != nil {
				err = msgp.WrapError(err, "WorkerPubList")
				return
			}
			if z.WorkerPubList == nil {
				z.WorkerPubList = make(map[string]uint64, zb0002)
			} else if len(z.WorkerPubList) > 0 {
				for key := range z.WorkerPubList {
					delete(z.WorkerPubList, key)
				}
			}
			for zb0002 > 0 {
				zb0002--
				var za0001 string
				var za0002 uint64
				za0001, err = dc.ReadString()
				if err != nil {
					err = msgp.WrapError(err, "WorkerPubList")
					return
				}
				za0002, err = dc.ReadUint64()
				if err != nil {
					err = msgp.WrapError(err, "WorkerPubList", za0001)
					return
				}
				z.WorkerPubList[za0001] = za0002
			}
		case "WorkerCandidatePubList":
			var zb0003 uint32
			zb0003, err = dc.ReadMapHeader()
			if err != nil {
				err = msgp.WrapError(err, "WorkerCandidatePubList")
				return
			}
			if z.WorkerCandidatePubList == nil {
				z.WorkerCandidatePubList = make(map[string]uint64, zb0003)
			} else if len(z.WorkerCandidatePubList) > 0 {
				for key := range z.WorkerCandidatePubList {
					delete(z.WorkerCandidatePubList, key)
				}
			}
			for zb0003 > 0 {
				zb0003--
				var za0003 string
				var za0004 uint64
				za0003, err = dc.ReadString()
				if err != nil {
					err = msgp.WrapError(err, "WorkerCandidatePubList")
					return
				}
				za0004, err = dc.ReadUint64()
				if err != nil {
					err = msgp.WrapError(err, "WorkerCandidatePubList", za0003)
					return
				}
				z.WorkerCandidatePubList[za0003] = za0004
			}
		case "VoterPubList":
			var zb0004 uint32
			zb0004, err = dc.ReadMapHeader()
			if err != nil {
				err = msgp.WrapError(err, "VoterPubList")
				return
			}
			if z.VoterPubList == nil {
				z.VoterPubList = make(map[string]uint64, zb0004)
			} else if len(z.VoterPubList) > 0 {
				for key := range z.VoterPubList {
					delete(z.VoterPubList, key)
				}
			}
			for zb0004 > 0 {
				zb0004--
				var za0005 string
				var za0006 uint64
				za0005, err = dc.ReadString()
				if err != nil {
					err = msgp.WrapError(err, "VoterPubList")
					return
				}
				za0006, err = dc.ReadUint64()
				if err != nil {
					err = msgp.WrapError(err, "VoterPubList", za0005)
					return
				}
				z.VoterPubList[za0005] = za0006
			}
		case "WNS":
			var zb0005 uint32
			zb0005, err = dc.ReadArrayHeader()
			if err != nil {
				err = msgp.WrapError(err, "WorkerSet")
				return
			}
			if cap(z.WorkerSet) >= int(zb0005) {
				z.WorkerSet = (z.WorkerSet)[:zb0005]
			} else {
				z.WorkerSet = make([]string, zb0005)
			}
			for za0007 := range z.WorkerSet {
				z.WorkerSet[za0007], err = dc.ReadString()
				if err != nil {
					err = msgp.WrapError(err, "WorkerSet", za0007)
					return
				}
			}
		case "VS":
			var zb0006 uint32
			zb0006, err = dc.ReadArrayHeader()
			if err != nil {
				err = msgp.WrapError(err, "VoterSet")
				return
			}
			if cap(z.VoterSet) >= int(zb0006) {
				z.VoterSet = (z.VoterSet)[:zb0006]
			} else {
				z.VoterSet = make([]string, zb0006)
			}
			for za0008 := range z.VoterSet {
				z.VoterSet[za0008], err = dc.ReadString()
				if err != nil {
					err = msgp.WrapError(err, "VoterSet", za0008)
					return
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *GenesisTransaction) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 9
	// write "WorkerNum"
	err = en.Append(0x89, 0xa9, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x4e, 0x75, 0x6d)
	if err != nil {
		return
	}
	err = en.WriteInt(z.WorkerNum)
	if err != nil {
		err = msgp.WrapError(err, "WorkerNum")
		return
	}
	// write "VotedNum"
	err = en.Append(0xa8, 0x56, 0x6f, 0x74, 0x65, 0x64, 0x4e, 0x75, 0x6d)
	if err != nil {
		return
	}
	err = en.WriteInt(z.VotedNum)
	if err != nil {
		err = msgp.WrapError(err, "VotedNum")
		return
	}
	// write "BlockGroupPerCycle"
	err = en.Append(0xb2, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x50, 0x65, 0x72, 0x43, 0x79, 0x63, 0x6c, 0x65)
	if err != nil {
		return
	}
	err = en.WriteInt(z.BlockGroupPerCycle)
	if err != nil {
		err = msgp.WrapError(err, "BlockGroupPerCycle")
		return
	}
	// write "Tcut"
	err = en.Append(0xa4, 0x54, 0x63, 0x75, 0x74)
	if err != nil {
		return
	}
	err = en.WriteFloat64(z.Tcut)
	if err != nil {
		err = msgp.WrapError(err, "Tcut")
		return
	}
	// write "WorkerPubList"
	err = en.Append(0xad, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x50, 0x75, 0x62, 0x4c, 0x69, 0x73, 0x74)
	if err != nil {
		return
	}
	err = en.WriteMapHeader(uint32(len(z.WorkerPubList)))
	if err != nil {
		err = msgp.WrapError(err, "WorkerPubList")
		return
	}
	for za0001, za0002 := range z.WorkerPubList {
		err = en.WriteString(za0001)
		if err != nil {
			err = msgp.WrapError(err, "WorkerPubList")
			return
		}
		err = en.WriteUint64(za0002)
		if err != nil {
			err = msgp.WrapError(err, "WorkerPubList", za0001)
			return
		}
	}
	// write "WorkerCandidatePubList"
	err = en.Append(0xb6, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x50, 0x75, 0x62, 0x4c, 0x69, 0x73, 0x74)
	if err != nil {
		return
	}
	err = en.WriteMapHeader(uint32(len(z.WorkerCandidatePubList)))
	if err != nil {
		err = msgp.WrapError(err, "WorkerCandidatePubList")
		return
	}
	for za0003, za0004 := range z.WorkerCandidatePubList {
		err = en.WriteString(za0003)
		if err != nil {
			err = msgp.WrapError(err, "WorkerCandidatePubList")
			return
		}
		err = en.WriteUint64(za0004)
		if err != nil {
			err = msgp.WrapError(err, "WorkerCandidatePubList", za0003)
			return
		}
	}
	// write "VoterPubList"
	err = en.Append(0xac, 0x56, 0x6f, 0x74, 0x65, 0x72, 0x50, 0x75, 0x62, 0x4c, 0x69, 0x73, 0x74)
	if err != nil {
		return
	}
	err = en.WriteMapHeader(uint32(len(z.VoterPubList)))
	if err != nil {
		err = msgp.WrapError(err, "VoterPubList")
		return
	}
	for za0005, za0006 := range z.VoterPubList {
		err = en.WriteString(za0005)
		if err != nil {
			err = msgp.WrapError(err, "VoterPubList")
			return
		}
		err = en.WriteUint64(za0006)
		if err != nil {
			err = msgp.WrapError(err, "VoterPubList", za0005)
			return
		}
	}
	// write "WNS"
	err = en.Append(0xa3, 0x57, 0x4e, 0x53)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.WorkerSet)))
	if err != nil {
		err = msgp.WrapError(err, "WorkerSet")
		return
	}
	for za0007 := range z.WorkerSet {
		err = en.WriteString(z.WorkerSet[za0007])
		if err != nil {
			err = msgp.WrapError(err, "WorkerSet", za0007)
			return
		}
	}
	// write "VS"
	err = en.Append(0xa2, 0x56, 0x53)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.VoterSet)))
	if err != nil {
		err = msgp.WrapError(err, "VoterSet")
		return
	}
	for za0008 := range z.VoterSet {
		err = en.WriteString(z.VoterSet[za0008])
		if err != nil {
			err = msgp.WrapError(err, "VoterSet", za0008)
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *GenesisTransaction) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 9
	// string "WorkerNum"
	o = append(o, 0x89, 0xa9, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x4e, 0x75, 0x6d)
	o = msgp.AppendInt(o, z.WorkerNum)
	// string "VotedNum"
	o = append(o, 0xa8, 0x56, 0x6f, 0x74, 0x65, 0x64, 0x4e, 0x75, 0x6d)
	o = msgp.AppendInt(o, z.VotedNum)
	// string "BlockGroupPerCycle"
	o = append(o, 0xb2, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x50, 0x65, 0x72, 0x43, 0x79, 0x63, 0x6c, 0x65)
	o = msgp.AppendInt(o, z.BlockGroupPerCycle)
	// string "Tcut"
	o = append(o, 0xa4, 0x54, 0x63, 0x75, 0x74)
	o = msgp.AppendFloat64(o, z.Tcut)
	// string "WorkerPubList"
	o = append(o, 0xad, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x50, 0x75, 0x62, 0x4c, 0x69, 0x73, 0x74)
	o = msgp.AppendMapHeader(o, uint32(len(z.WorkerPubList)))
	for za0001, za0002 := range z.WorkerPubList {
		o = msgp.AppendString(o, za0001)
		o = msgp.AppendUint64(o, za0002)
	}
	// string "WorkerCandidatePubList"
	o = append(o, 0xb6, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x50, 0x75, 0x62, 0x4c, 0x69, 0x73, 0x74)
	o = msgp.AppendMapHeader(o, uint32(len(z.WorkerCandidatePubList)))
	for za0003, za0004 := range z.WorkerCandidatePubList {
		o = msgp.AppendString(o, za0003)
		o = msgp.AppendUint64(o, za0004)
	}
	// string "VoterPubList"
	o = append(o, 0xac, 0x56, 0x6f, 0x74, 0x65, 0x72, 0x50, 0x75, 0x62, 0x4c, 0x69, 0x73, 0x74)
	o = msgp.AppendMapHeader(o, uint32(len(z.VoterPubList)))
	for za0005, za0006 := range z.VoterPubList {
		o = msgp.AppendString(o, za0005)
		o = msgp.AppendUint64(o, za0006)
	}
	// string "WNS"
	o = append(o, 0xa3, 0x57, 0x4e, 0x53)
	o = msgp.AppendArrayHeader(o, uint32(len(z.WorkerSet)))
	for za0007 := range z.WorkerSet {
		o = msgp.AppendString(o, z.WorkerSet[za0007])
	}
	// string "VS"
	o = append(o, 0xa2, 0x56, 0x53)
	o = msgp.AppendArrayHeader(o, uint32(len(z.VoterSet)))
	for za0008 := range z.VoterSet {
		o = msgp.AppendString(o, z.VoterSet[za0008])
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *GenesisTransaction) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "WorkerNum":
			z.WorkerNum, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "WorkerNum")
				return
			}
		case "VotedNum":
			z.VotedNum, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "VotedNum")
				return
			}
		case "BlockGroupPerCycle":
			z.BlockGroupPerCycle, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "BlockGroupPerCycle")
				return
			}
		case "Tcut":
			z.Tcut, bts, err = msgp.ReadFloat64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Tcut")
				return
			}
		case "WorkerPubList":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "WorkerPubList")
				return
			}
			if z.WorkerPubList == nil {
				z.WorkerPubList = make(map[string]uint64, zb0002)
			} else if len(z.WorkerPubList) > 0 {
				for key := range z.WorkerPubList {
					delete(z.WorkerPubList, key)
				}
			}
			for zb0002 > 0 {
				var za0001 string
				var za0002 uint64
				zb0002--
				za0001, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "WorkerPubList")
					return
				}
				za0002, bts, err = msgp.ReadUint64Bytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "WorkerPubList", za0001)
					return
				}
				z.WorkerPubList[za0001] = za0002
			}
		case "WorkerCandidatePubList":
			var zb0003 uint32
			zb0003, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "WorkerCandidatePubList")
				return
			}
			if z.WorkerCandidatePubList == nil {
				z.WorkerCandidatePubList = make(map[string]uint64, zb0003)
			} else if len(z.WorkerCandidatePubList) > 0 {
				for key := range z.WorkerCandidatePubList {
					delete(z.WorkerCandidatePubList, key)
				}
			}
			for zb0003 > 0 {
				var za0003 string
				var za0004 uint64
				zb0003--
				za0003, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "WorkerCandidatePubList")
					return
				}
				za0004, bts, err = msgp.ReadUint64Bytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "WorkerCandidatePubList", za0003)
					return
				}
				z.WorkerCandidatePubList[za0003] = za0004
			}
		case "VoterPubList":
			var zb0004 uint32
			zb0004, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "VoterPubList")
				return
			}
			if z.VoterPubList == nil {
				z.VoterPubList = make(map[string]uint64, zb0004)
			} else if len(z.VoterPubList) > 0 {
				for key := range z.VoterPubList {
					delete(z.VoterPubList, key)
				}
			}
			for zb0004 > 0 {
				var za0005 string
				var za0006 uint64
				zb0004--
				za0005, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "VoterPubList")
					return
				}
				za0006, bts, err = msgp.ReadUint64Bytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "VoterPubList", za0005)
					return
				}
				z.VoterPubList[za0005] = za0006
			}
		case "WNS":
			var zb0005 uint32
			zb0005, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "WorkerSet")
				return
			}
			if cap(z.WorkerSet) >= int(zb0005) {
				z.WorkerSet = (z.WorkerSet)[:zb0005]
			} else {
				z.WorkerSet = make([]string, zb0005)
			}
			for za0007 := range z.WorkerSet {
				z.WorkerSet[za0007], bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "WorkerSet", za0007)
					return
				}
			}
		case "VS":
			var zb0006 uint32
			zb0006, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "VoterSet")
				return
			}
			if cap(z.VoterSet) >= int(zb0006) {
				z.VoterSet = (z.VoterSet)[:zb0006]
			} else {
				z.VoterSet = make([]string, zb0006)
			}
			for za0008 := range z.VoterSet {
				z.VoterSet[za0008], bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "VoterSet", za0008)
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *GenesisTransaction) Msgsize() (s int) {
	s = 1 + 10 + msgp.IntSize + 9 + msgp.IntSize + 19 + msgp.IntSize + 5 + msgp.Float64Size + 14 + msgp.MapHeaderSize
	if z.WorkerPubList != nil {
		for za0001, za0002 := range z.WorkerPubList {
			_ = za0002
			s += msgp.StringPrefixSize + len(za0001) + msgp.Uint64Size
		}
	}
	s += 23 + msgp.MapHeaderSize
	if z.WorkerCandidatePubList != nil {
		for za0003, za0004 := range z.WorkerCandidatePubList {
			_ = za0004
			s += msgp.StringPrefixSize + len(za0003) + msgp.Uint64Size
		}
	}
	s += 13 + msgp.MapHeaderSize
	if z.VoterPubList != nil {
		for za0005, za0006 := range z.VoterPubList {
			_ = za0006
			s += msgp.StringPrefixSize + len(za0005) + msgp.Uint64Size
		}
	}
	s += 4 + msgp.ArrayHeaderSize
	for za0007 := range z.WorkerSet {
		s += msgp.StringPrefixSize + len(z.WorkerSet[za0007])
	}
	s += 3 + msgp.ArrayHeaderSize
	for za0008 := range z.VoterSet {
		s += msgp.StringPrefixSize + len(z.VoterSet[za0008])
	}
	return
}
