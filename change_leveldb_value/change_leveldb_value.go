package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/golang/protobuf/proto"
	"github.com/syndtr/goleveldb/leveldb"

	"github.com/cetc/xledger/core/ledger/kvledger/txmgmt/statedb"
	"github.com/cetc/xledger/core/ledger/kvledger/txmgmt/statedb/stateleveldb/msgs"
	"github.com/cetc/xledger/core/ledger/kvledger/txmgmt/version"
)

var (
	channel   string // 通道名
	chaincode string // 智能合约名
	key       string // 修改的key
	value     string // 修改的值
	dbpath    string // leveldb数据库路径
	del       bool   // 删除key
)

var compositeKeySep = []byte{0x00}

func init() {
	flag.StringVar(&channel, "channel", "", "Channel name")
	flag.StringVar(&chaincode, "chaincode", "", "Chaincode name")
	flag.StringVar(&key, "key", "", "Key to query; empty query all keys")
	flag.StringVar(&value, "value", "", "The key change to value")
	flag.StringVar(&dbpath, "dbpath", "", "Path to leveldb")
	flag.BoolVar(&del, "del", false, "Delete this key")
}

func readKey(db *leveldb.DB, key string) {
	realKey := constructCompositeKey(channel, string(constructCompositeKey(chaincode, key)))

	originValue, err := db.Get(realKey, nil)
	if err != nil {
		fmt.Printf("cannot get [%s] form leveldb, error:%v\n", key, err)
		return
	}

	decodeValue, err := decodeValue(originValue)
	if err != nil {
		fmt.Printf("cannot decodeValue key[%s], error:%v\n", key, err)
		return
	}

	fmt.Printf("before change [%s]=[%s]\n", key, string(originValue))

	oldValue := decodeValue.Value
	// 修改值
	decodeValue.Value = []byte(value)
	evv, err := encodeVersionedValue(decodeValue)
	if err != nil {
		fmt.Printf("encodeValue[%s], error=[%v]\n", value, err)
		return
	}

	err = putKey(db, realKey, evv)
	if err != nil {
		fmt.Printf("cannot change key[%s] form %s to %s , error=[%v]\n", key, oldValue, value, err)
		return
	}

	changedValue, err := db.Get(realKey, nil)
	if err != nil {
		fmt.Printf("after change cannot read key[%s], error=[%v]\n", key, err)
		return
	}
	fmt.Printf("after change [%s]=[%s]\n", key, string(changedValue))
}

func constructCompositeKey(ns string, key string) []byte {
	return append(append([]byte(ns), compositeKeySep...), []byte(key)...)
}

func encodeVersionedValue(v *statedb.VersionedValue) ([]byte, error) {
	vvMsg := &msgs.VersionedValueProto{
		VersionBytes: v.Version.ToBytes(),
		Value:        v.Value,
		Metadata:     v.Metadata,
	}
	encodedValue, err := proto.Marshal(vvMsg)
	if err != nil {
		return nil, err
	}
	encodedValue = append(compositeKeySep, encodedValue...)
	return encodedValue, nil
}

func decodeValue(encodedValue []byte) (*statedb.VersionedValue, error) {
	if oldFormatEncoding(encodedValue) {
		val, ver := decodeValueOldFormat(encodedValue)
		return &statedb.VersionedValue{Version: ver, Value: val, Metadata: nil}, nil
	}
	msg := &msgs.VersionedValueProto{}
	err := proto.Unmarshal(encodedValue[1:], msg)
	if err != nil {
		return nil, err
	}
	ver, _ := version.NewHeightFromBytes(msg.VersionBytes)
	val := msg.Value
	metadata := msg.Metadata
	// protobuf always makes an empty byte array as nil
	if val == nil {
		val = []byte{}
	}
	return &statedb.VersionedValue{Version: ver, Value: val, Metadata: metadata}, nil
}

// oldFormatEncoding checks whether the value is encoded using the old (pre-v1.3) format
// or new format (v1.3 and later for encoding metadata).
func oldFormatEncoding(encodedValue []byte) bool {
	return encodedValue[0] != byte(0) ||
		(encodedValue[0]|encodedValue[1]) == byte(0) // this check covers a corner case
	// where the old formatted value happens to start with a nil byte. In this corner case,
	// the channel config happen to be persisted for the tuple <block 0, tran 0>. So, this
	// is assumed that block 0 contains a single transaction (i.e., tran 0)
}

func decodeValueOldFormat(encodedValue []byte) ([]byte, *version.Height) {
	height, n := version.NewHeightFromBytes(encodedValue)
	value := encodedValue[n:]
	return value, height
}

func readAll(db *leveldb.DB) {
	prefix := constructCompositeKey(channel, chaincode)

	iter := db.NewIterator(nil, nil)
	defer iter.Release()
	fmt.Printf("read all\n")
	for iter.Next() {
		key := string(iter.Key())
		if strings.HasPrefix(key, string(prefix)) {
			value := string(iter.Value())
			fmt.Printf("[%s]=[%s]\n", key, value)
		}
	}
}

func putKey(db *leveldb.DB, key []byte, value []byte) error {
	err := db.Put(key, value, nil)
	if err != nil {
		return err
	}
	return nil
}

func delKey(db *leveldb.DB, key []byte) error {
	err := db.Delete(key, nil)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	flag.Parse()
	if channel == "" || chaincode == "" || dbpath == "" {
		fmt.Printf("channel, chaincode, dbpath could be empty\n")
		return
	}

	db, err := leveldb.OpenFile(dbpath, nil)
	if err != nil {
		fmt.Printf("cannot open LevelDB from [%s], with error=[%v]\n", dbpath, err)
	}
	defer func(db *leveldb.DB) {
		err := db.Close()
		if err != nil {
			fmt.Printf("close db err:%v\n", err)
		}
	}(db)

	if del {
		realKey := constructCompositeKey(channel, string(constructCompositeKey(chaincode, key)))
		err := delKey(db, realKey)
		if err != nil {
			fmt.Printf("delete key[%s] ,err:%v\n", key, err)
			return
		}
		fmt.Printf("delete key[%s]\n", key)
		return
	}

	if key == "" {
		readAll(db)
	} else {
		readKey(db, key)
	}
}
