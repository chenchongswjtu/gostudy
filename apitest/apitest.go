// 用于测试shim.ChaincodeStubInterface接口中的所有函数（不包括需要富查询支持的函数）
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/cetc/xledger/core/chaincode/shim"
	"github.com/cetc/xledger/core/chaincode/shim/ext/statebased"
	pb "github.com/cetc/xledger/protos/peer"
)

const (
	colorName         = "color~name"
	collectionMarbles = "collectionMarbles"
)

type marble struct {
	ObjectType string `json:"docType"` // docType用于区分数据库中的各种对象类型
	Name       string `json:"name"`    // 唯一,key
	Color      string `json:"color"`
	Size       int    `json:"size"`
	Owner      string `json:"owner"`
}

type marblePrivate struct {
	ObjectType string `json:"docType"` // docType用于区分数据库中的各种对象类型
	Name       string `json:"name"`    // 唯一,key
	Price      int    `json:"price"`
}

func main() {
	err := shim.Start(new(marble))
	if err != nil {
		fmt.Printf("error starting chaincode: %s", err)
	}
}

// Init 初始化链码
func (m *marble) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

// Invoke 具体的操作
func (m *marble) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	fmt.Println("invoke is running", function)
	function = strings.ToLower(function)

	// 处理不同的函数
	if function == "init" {
		return m.init(stub, args)
	} else if function == "query" {
		return m.query(stub, args)
	} else if function == "delete" {
		return m.delete(stub, args)
	} else if function == "range" {
		return m.getByRange(stub, args)
	} else if function == "history" {
		return m.getHistory(stub, args)
	} else if function == strings.ToLower("transferBaseOnColor") {
		return m.transferBasedOnColor(stub, args)
	} else if function == strings.ToLower("GetStateByPartialCompositeKeyWithPagination") {
		return m.getStateByPartialCompositeKeyWithPagination(stub, args)
	} else if function == strings.ToLower("getRangeWithPagination") {
		return m.getByRangeWithPagination(stub, args)
	} else if function == strings.ToLower("invokeChainCode") {
		return m.invokeChainCode(stub, args)
	} else if function == "event" {
		return m.setEvent(stub, args)
	} else if function == "info" {
		return m.getInfo(stub, args)
	} else if function == strings.ToLower("keyLevel") { // key级别的背书策略
		return m.setKeyLevel(stub, args)
	} else if function == strings.ToLower("initPrivate") {
		return m.initPrivate(stub, args)
	} else if function == strings.ToLower("queryPrivate") {
		return m.queryPrivate(stub, args)
	} else if function == strings.ToLower("delPrivate") {
		return m.delPrivate(stub, args)
	} else if function == strings.ToLower("privateRange") {
		return m.getPrivateByRange(stub, args)
	} else if function == strings.ToLower("GetPrivateDataByPartialCompositeKey") {
		return m.getPrivateDataByPartialCompositeKey(stub, args)
	} else if function == strings.ToLower("privateKeyLevel") {
		return m.setPrivateKeyLevel(stub, args)
	}

	return shim.Success(nil)
}

// 根据传入的参数初始化弹珠
// GetState(key string) ([]byte, error)
// PutState(key string, value []byte) error
// CreateCompositeKey(objectType string, attributes []string) (string, error)
func (m *marble) init(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 4 {
		return shim.Error("incorrect number of arguments. excepting 4")
	}

	name := args[0]
	color := strings.ToLower(args[1])
	owner := strings.ToLower(args[3])
	size, err := strconv.Atoi(args[2])
	if err != nil {
		return shim.Error("3rd argument must be a numeric string")
	}

	// 检查该name是否存在
	marbleAsBytes, err := stub.GetState(name)
	if err != nil {
		return shim.Error("failed to get marble: " + err.Error())
	} else if len(marbleAsBytes) > 0 {
		fmt.Printf("this marble already exists: %s, update it ", name)
	}

	objectType := "marble"
	marble := &marble{
		ObjectType: objectType,
		Name:       name,
		Color:      color,
		Size:       size,
		Owner:      owner,
	}
	marbleJSONasBytes, err := json.Marshal(marble)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(name, marbleJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	//  ==== Index the marble to enable color-based range queries, e.g. return all blue marbles ====
	//  An 'index' is a normal key/value entry in state.
	//  The key is a composite key, with the elements that you want to range query on listed first.
	//  In our case, the composite key is based on indexName~color~name.
	//  This will enable very efficient state range queries based on composite keys matching indexName~color~*
	colorNameIndexKey, err := stub.CreateCompositeKey(colorName, []string{marble.Color, marble.Name})
	if err != nil {
		return shim.Error(err.Error())
	}
	//  Save index entry to state. Only the key name is needed, no need to store a duplicate copy of the marble.
	//  Note - passing a 'nil' value will effectively delete the key from state, therefore we pass null character as value
	value := []byte{0x00}
	if err = stub.PutState(colorNameIndexKey, value); err != nil {
		return shim.Error(fmt.Sprintf("save %s PutState err: %s", colorNameIndexKey, err.Error()))
	}
	return shim.Success(nil)
}

func (m *marble) query(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var jsonResp string

	if len(args) != 1 {
		return shim.Error("incorrect number of arguments. expecting 1")
	}
	name := args[0]

	// 为了维持color~name索引，需要先获得颜色
	valAsBytes, err := stub.GetState(name)
	if err != nil {
		jsonResp = "{\"error\":\"failed to get state for " + name + "\"}"
		return shim.Error(jsonResp)
	} else if valAsBytes == nil {
		jsonResp = "{\"error\":\"marble does not exist: " + name + "\"}"
		return shim.Error(jsonResp)
	}

	fmt.Printf("%s value is %s", name, string(valAsBytes))
	return shim.Success(valAsBytes)
}

// 根据name删除数据
// DelState(key string) error
func (m *marble) delete(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var jsonResp string
	var marble marble

	if len(args) != 1 {
		return shim.Error("incorrect number of arguments. expecting 1")
	}
	name := args[0]

	// 为了维持color~name索引，需要先获得颜色
	valAsBytes, err := stub.GetState(name)
	if err != nil {
		jsonResp = "{\"error\":\"failed to get state for " + name + "\"}"
		return shim.Error(jsonResp)
	} else if valAsBytes == nil {
		jsonResp = "{\"error\":\"marble does not exist: " + name + "\"}"
		return shim.Error(jsonResp)
	}

	err = json.Unmarshal(valAsBytes, &marble)
	if err != nil {
		jsonResp = "{\"error\":\"failed to decode JSON of: " + name + "\"}"
		return shim.Error(jsonResp)
	}

	// 根据name删除数据
	err = stub.DelState(name)
	if err != nil {
		return shim.Error("failed to delete state:" + err.Error())
	}

	// 创建color~name复合键
	colorNameIndexKey, err := stub.CreateCompositeKey(colorName, []string{marble.Color, marble.Name})
	if err != nil {
		return shim.Error(err.Error())
	}

	// 删除复合键
	err = stub.DelState(colorNameIndexKey)
	if err != nil {
		return shim.Error("failed to delete state:" + err.Error())
	}
	return shim.Success(nil)
}

// 根据查询范围获得结果
// GetStateByRange(startKey, endKey string) (StateQueryIteratorInterface, error)
func (m *marble) getByRange(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) < 2 {
		return shim.Error("incorrect number of arguments. expecting 2")
	}

	startKey := args[0]
	endKey := args[1]

	iterator, err := stub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer iterator.Close()

	buffer, err := constructQueryResponseFromIterator(iterator)
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Printf("getByRange queryResult:\n%s\n", buffer.String())
	return shim.Success(buffer.Bytes())
}

// 获得name的历史交易记录
// GetHistoryForKey(key string) (HistoryQueryIteratorInterface, error)
func (m *marble) getHistory(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) < 1 {
		return shim.Error("incorrect number of arguments. expecting 1")
	}

	name := args[0]
	iterator, err := stub.GetHistoryForKey(name)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer iterator.Close()

	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for iterator.HasNext() {
		response, err := iterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}

		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"TxId\":")
		buffer.WriteString("\"")
		buffer.WriteString(response.TxId)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Value\":")
		if response.IsDelete {
			buffer.WriteString("null")
		} else {
			buffer.WriteString(string(response.Value))
		}

		buffer.WriteString(", \"Timestamp\":")
		buffer.WriteString("\"")
		buffer.WriteString(time.Unix(response.Timestamp.Seconds, int64(response.Timestamp.Nanos)).String())
		buffer.WriteString("\"")

		buffer.WriteString(", \"IsDelete\":")
		buffer.WriteString("\"")
		buffer.WriteString(strconv.FormatBool(response.IsDelete))
		buffer.WriteString("\"")

		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("getHistory returning:\n%s\n", buffer.String())
	return shim.Success(buffer.Bytes())
}

// 将指定颜色的所有弹珠转移
// GetStateByPartialCompositeKey(objectType string, keys []string) (StateQueryIteratorInterface, error)
// SplitCompositeKey(compositeKey string) (string, []string, error)
func (m *marble) transferBasedOnColor(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) < 2 {
		return shim.Error("incorrect number of arguments. expecting 2")
	}

	color := args[0]
	newOwner := strings.ToLower(args[1])

	// Query the color~name index by color
	// This will execute a key range query on all keys starting with 'color'
	iterator, err := stub.GetStateByPartialCompositeKey(colorName, []string{color})
	if err != nil {
		return shim.Error(err.Error())
	}
	defer iterator.Close()

	// Iterate through result set and for each marble found, transfer to newOwner
	var i int
	for i = 0; iterator.HasNext(); i++ {
		// Note that we don't get the value (2nd return variable), we'll just get the marble name from the composite key
		responseRange, err := iterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}

		// get the color and name from color~name composite key
		objectType, compositeKeyParts, err := stub.SplitCompositeKey(responseRange.Key)
		if err != nil {
			return shim.Error(err.Error())
		}
		returnedColor := compositeKeyParts[0]
		returnedMarbleName := compositeKeyParts[1]
		fmt.Printf("found a marble from index:%s color:%s name:%s\n", objectType, returnedColor, returnedMarbleName)

		response := m.transfer(stub, []string{returnedMarbleName, newOwner})

		if response.Status != shim.OK {
			return shim.Error("Transfer failed: " + response.Message)
		}
	}

	responsePayload := fmt.Sprintf("Transferred %d %s marbles to %s", i, color, newOwner)
	return shim.Success([]byte(responsePayload))
}

// 转移弹珠到新的owner
func (m *marble) transfer(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) < 2 {
		return shim.Error("Incorrect number of arguments. expecting 2")
	}

	name := args[0]
	newOwner := strings.ToLower(args[1])

	marbleAsBytes, err := stub.GetState(name)
	if err != nil {
		return shim.Error("failed to get marble:" + err.Error())
	} else if len(marbleAsBytes) == 0 {
		return shim.Error("marble does not exist")
	}

	marbleToTransfer := marble{}
	err = json.Unmarshal(marbleAsBytes, &marbleToTransfer)
	if err != nil {
		return shim.Error(err.Error())
	}
	marbleToTransfer.Owner = newOwner

	marbleJSONasBytes, err := json.Marshal(marbleToTransfer)
	if err != nil {
		return shim.Error("json marshal err: " + err.Error())
	}

	err = stub.PutState(name, marbleJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

// 局部匹配复合键分页查询
// GetStateByPartialCompositeKeyWithPagination(objectType string, keys []string,
//		pageSize int32, bookmark string) (StateQueryIteratorInterface, *pb.QueryResponseMetadata, error)
func (m *marble) getStateByPartialCompositeKeyWithPagination(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) < 3 {
		return shim.Error("incorrect number of arguments. expecting 3")
	}

	color := args[0]
	pageSize, err := strconv.Atoi(args[1])
	if err != nil {
		return shim.Error("2nd argument must be a numeric string")
	}
	bookmark := args[2]

	iterator, metadata, err := stub.GetStateByPartialCompositeKeyWithPagination(colorName, []string{color}, int32(pageSize), bookmark)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer iterator.Close()

	buffer, err := constructQueryResponseFromIteratorAndMetadata(iterator, metadata)
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Printf("getByRange queryResult:\n%s\n", buffer.String())
	return shim.Success(buffer.Bytes())
}

// 范围分页查询
// 	GetStateByRangeWithPagination(startKey, endKey string, pageSize int32,
//		bookmark string) (StateQueryIteratorInterface, *pb.QueryResponseMetadata, error)
func (m *marble) getByRangeWithPagination(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) < 4 {
		return shim.Error("incorrect number of arguments. expecting 4")
	}

	startKey := args[0]
	endKey := args[1]
	bookmark := args[3]
	pageSize, err := strconv.Atoi(args[2])
	if err != nil {
		return shim.Error("3rd argument must be a numeric string")
	}

	iterator, metadata, err := stub.GetStateByRangeWithPagination(startKey, endKey, int32(pageSize), bookmark)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer iterator.Close()

	buffer, err := constructQueryResponseFromIteratorAndMetadata(iterator, metadata)
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Printf("getByRange queryResult:\n%s\n", buffer.String())
	return shim.Success(buffer.Bytes())
}

// 调用其他链码 查询数据
// channelName 为空时，表示通道与调用通道相同
// InvokeChaincode(chaincodeName string, args [][]byte, channel string) pb.Response
func (m *marble) invokeChainCode(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) < 2 {
		return shim.Error("incorrect number of arguments. expecting at least 2")
	}

	chaincodeName := args[0]
	name := args[1]
	channelName := ""
	if len(args) > 2 {
		channelName = args[2]
	}

	f := "query"
	invokeArgs := toChaincodeArgs(f, name)
	response := stub.InvokeChaincode(chaincodeName, invokeArgs, channelName)
	if response.Status != shim.OK {
		errStr := fmt.Sprintf("failed to %s chaincode. Got error: %s", f, response.Payload)
		fmt.Printf(errStr)
		return shim.Error(errStr)
	}
	return shim.Success(response.Payload)
}

// SetEvent(name string, payload []byte) error
func (m *marble) setEvent(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	toSend := "event "
	for _, s := range args {
		toSend = toSend + "," + s
	}

	err := stub.SetEvent("event", []byte(toSend))
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

// GetTxID() string
// GetChannelID() string
// GetDecorations() map[string][]byte
// GetCreator() ([]byte, error)
// GetBinding() ([]byte, error)
// GetSignedProposal() (*pb.SignedProposal, error)
// GetTxTimestamp() (*timestamp.Timestamp, error)
func (m *marble) getInfo(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	txID := stub.GetTxID()
	channelID := stub.GetChannelID()
	decorations := stub.GetDecorations()
	decorationsJson, err := json.Marshal(decorations)
	if err != nil {
		fmt.Println("json.Marshal failed err: " + err.Error())
	}
	decorationsStr := string(decorationsJson)

	creator, err := stub.GetCreator()
	if err != nil {
		return shim.Error("GetCreator err: " + err.Error())
	}
	creatorStr := string(creator)

	bind, err := stub.GetBinding()
	if err != nil {
		return shim.Error("GetBinding err: " + err.Error())
	}
	bindStr := string(bind)

	signedProposal, err := stub.GetSignedProposal()
	if err != nil {
		return shim.Error("GetSignedProposal err: " + err.Error())
	}
	sp := signedProposal.String()

	timeStamp, err := stub.GetTxTimestamp()
	if err != nil {
		return shim.Error("GetTxTimestamp err: " + err.Error())
	}
	timeStr := time.Unix(timeStamp.Seconds, int64(timeStamp.Nanos)).String()

	info := fmt.Sprintf(`{"txid":"%s","channelid":"%s","decorations":"%s","creator":"%s","bind":"%s","signedproposal":"%s","time":"%s"}`,
		txID, channelID, decorationsStr, creatorStr, bindStr, sp, timeStr)
	return shim.Success([]byte(info))
}

// export MARBLE=$(echo -n "{\"name\":\"marble1\",\"price\":99}" | base64 | tr -d \\n)
// peer chaincode invoke -C mychannel -n marblesp -c '{"Args":["initMarble"]}' --transient "{\"marble\":\"$MARBLE\"}"
// GetTransient() (map[string][]byte, error)
// GetPrivateData(collection, key string) ([]byte, error)
// PutPrivateData(collection string, key string, value []byte) error
func (m *marble) initPrivate(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error

	type marbleTransientInput struct {
		Name  string `json:"name"`
		Price int    `json:"price"`
	}

	// ==== Input sanitation ====
	fmt.Println("- start init marble")

	if len(args) != 0 {
		return shim.Error("incorrect number of arguments. private marble data must be passed in transient map.")
	}

	transMap, err := stub.GetTransient()
	if err != nil {
		return shim.Error("error getting transient: " + err.Error())
	}

	if _, ok := transMap["marble"]; !ok {
		return shim.Error("marble must be a key in the transient map")
	}

	if len(transMap["marble"]) == 0 {
		return shim.Error("marble value in the transient map must be a non-empty JSON string")
	}

	var marbleInput marbleTransientInput
	err = json.Unmarshal(transMap["marble"], &marbleInput)
	if err != nil {
		return shim.Error("failed to decode JSON of: " + string(transMap["marble"]))
	}

	if len(marbleInput.Name) == 0 {
		return shim.Error("name field must be a non-empty string")
	}
	if marbleInput.Price <= 0 {
		return shim.Error("price field must be a positive integer")
	}

	// ==== Check if marble already exists ====
	marbleAsBytes, err := stub.GetPrivateData(collectionMarbles, marbleInput.Name)
	if err != nil {
		return shim.Error("failed to get marble: " + err.Error())
	} else if len(marbleAsBytes) > 0 {
		fmt.Println("this marble already exists: " + marbleInput.Name)
		return shim.Error("this marble already exists: " + marbleInput.Name)
	}

	// ==== Create marble private details object with price, marshal to JSON, and save to state ====
	marblePrivateDetails := &marblePrivate{
		ObjectType: "marblePrivate",
		Name:       marbleInput.Name,
		Price:      marbleInput.Price,
	}
	marblePrivateDetailsBytes, err := json.Marshal(marblePrivateDetails)
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutPrivateData(collectionMarbles, marbleInput.Name, marblePrivateDetailsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

// GetPrivateData(collection, key string) ([]byte, error)
func (m *marble) queryPrivate(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) < 1 {
		return shim.Error("incorrect number of arguments. excepting 1.")
	}

	name := args[0]

	marbleAsBytes, err := stub.GetPrivateData(collectionMarbles, name)
	if err != nil {
		return shim.Error("failed to get marble: " + err.Error())
	} else if len(marbleAsBytes) == 0 {
		fmt.Println("this marble not exists: " + name)
		return shim.Error("this marble not exists: " + name)
	}

	fmt.Printf("private %s data %s \n", name, string(marbleAsBytes))
	return shim.Success(marbleAsBytes)
}

// DelPrivateData(collection, key string) error
func (m *marble) delPrivate(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	name := args[0]

	// ==== Check if marble already exists ====
	marbleAsBytes, err := stub.GetPrivateData(collectionMarbles, name)
	if err != nil {
		return shim.Error("Failed to get marble: " + err.Error())
	} else if marbleAsBytes == nil {
		fmt.Println("This marble not exists: " + name)
		return shim.Error("This marble not exists: " + name)
	}

	if err := stub.DelPrivateData(collectionMarbles, name); err != nil {
		return shim.Error("DelPrivateData err: " + err.Error())
	}

	return shim.Success(nil)
}

// GetPrivateDataByRange(collection, startKey, endKey string) (StateQueryIteratorInterface, error)
func (m *marble) getPrivateByRange(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) < 2 {
		return shim.Error("incorrect number of arguments. expecting 2")
	}

	startKey := args[0]
	endKey := args[1]

	resultsIterator, err := stub.GetPrivateDataByRange(collectionMarbles, startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	buffer, err := constructQueryResponseFromIterator(resultsIterator)
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Printf("getByRange queryResult:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

// GetPrivateDataByPartialCompositeKey(collection, objectType string, keys []string) (StateQueryIteratorInterface, error)
func (m *marble) getPrivateDataByPartialCompositeKey(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) < 1 {
		return shim.Error("incorrect number of arguments. expecting 2")
	}

	color := args[0]
	iterator, err := stub.GetPrivateDataByPartialCompositeKey(collectionMarbles, colorName, []string{color})
	if err != nil {
		return shim.Error(err.Error())
	}
	defer iterator.Close()

	buffer, err := constructQueryResponseFromIterator(iterator)
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Printf("getByRange queryResult:\n%s\n", buffer.String())
	return shim.Success(buffer.Bytes())
}

// GetStateValidationParameter(key string) ([]byte, error)
// SetStateValidationParameter(key string, ep []byte) error
func (m *marble) setKeyLevel(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting the key and EP to be set.")
	}
	key := args[0]
	ep, err := stub.GetStateValidationParameter(key)
	if err != nil {
		return shim.Error("GetStateValidationParameter err:" + err.Error())
	}

	if len(ep) == 0 {
		fmt.Printf("%s key level endorsement policy not exist \n", key)
	} else {
		fmt.Printf("%s key level endorsement policy already exist, is %s \n", key, string(ep))
	}

	EP := args[1]
	newEP, err := statebased.NewStateEP(nil)
	if err != nil {
		return shim.Error(err.Error())
	}
	err = newEP.AddOrgs(statebased.RoleTypeMember, EP)
	if err != nil {
		return shim.Error(err.Error())
	}

	policyByte, err := newEP.Policy()
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.SetStateValidationParameter(key, policyByte)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

// 设置私有数据key级背书策略
// GetPrivateDataValidationParameter(collection, key string) ([]byte, error)
// SetPrivateDataValidationParameter(collection, key string, ep []byte) error
func (m *marble) setPrivateKeyLevel(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting the key and EP to be set.")
	}
	key := args[0]
	ep, err := stub.GetPrivateDataValidationParameter(collectionMarbles, key)
	if err != nil {
		return shim.Error("GetPrivateDataValidationParameter err:" + err.Error())
	}

	if len(ep) == 0 {
		fmt.Printf("%s %s private key level endorsement policy not exist \n", collectionMarbles, key)
	} else {
		fmt.Printf("%s %s private key level endorsement policy already exist, is %s \n", collectionMarbles, key, string(ep))
	}

	EP := args[1]
	newEP, err := statebased.NewStateEP(nil)
	if err != nil {
		return shim.Error(err.Error())
	}
	err = newEP.AddOrgs(statebased.RoleTypeMember, EP)
	if err != nil {
		return shim.Error(err.Error())
	}

	policyByte, err := newEP.Policy()
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.SetPrivateDataValidationParameter(collectionMarbles, key, policyByte)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

func toChaincodeArgs(args ...string) [][]byte {
	argsAsBytes := make([][]byte, len(args))
	for i, arg := range args {
		argsAsBytes[i] = []byte(arg)
	}
	return argsAsBytes
}

func constructQueryResponseFromIterator(resultsIterator shim.StateQueryIteratorInterface) (*bytes.Buffer, error) {
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	return &buffer, nil
}

func constructQueryResponseFromIteratorAndMetadata(iterator shim.StateQueryIteratorInterface, metadata *pb.QueryResponseMetadata) (*bytes.Buffer, error) {
	var buffer bytes.Buffer
	buffer.WriteString(`{"count":"`)
	buffer.WriteString(strconv.Itoa(int(metadata.FetchedRecordsCount)))

	buffer.WriteString(", \"bookmark\":")
	buffer.WriteString(metadata.Bookmark)
	buffer.WriteString(", \"value\":")

	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for iterator.HasNext() {
		queryResponse, err := iterator.Next()
		if err != nil {
			return nil, err
		}
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]\"")

	buffer.WriteString("}")
	return &buffer, nil
}
