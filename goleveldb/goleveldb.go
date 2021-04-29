package main

import (
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type addresses struct {
	Addresses []string `json:"addresses"`
}

type consenters struct {
	Consenters []consenter `json:"consenters"`
}

type consenter struct {
	ClientTlsCert string `json:"client_tls_cert"`
	Host          string `json:"host"`
	Port          int    `json:"port"`
	ServerTlsCert string `json:"server_tls_cert"`
}

var (
	ordererFilePath      string
	ordererMSP           string
	ordererAddress       string
	ordererMSPConfigPath string
)

func main() {
	flag.StringVar(&ordererMSP, "msp", "OrdererMSP", "orderer的msp id")
	flag.StringVar(&ordererAddress, "address", "127.0.0.1:7050", "orderer的ip地址")
	flag.StringVar(&ordererFilePath, "filepath", "/var/production/cetc/orderer1.example.com/orderer/chains", "orderer数据存储路径")
	flag.StringVar(&ordererMSPConfigPath, "mspconfigpath", "crypto-config/ordererOrganizations/example.com/users/Admin@example.com/msp", "orderer的msp路径")

	flag.Parse()

	log.Println("orderer的msp id:", ordererMSP)
	log.Println("orderer的ip地址:", ordererAddress)
	log.Println("orderer数据存储路径:", ordererFilePath)
	log.Println("orderer的msp路径:", ordererMSPConfigPath)

	pwd := currentDir()
	log.Println("当前程序执行目录为:", pwd)
	// pwd路径下必须包含core.yaml文件
	// 设置必须的环境变量参数
	_ = os.Setenv("XLEDGER_CFG_PATH", pwd)
	_ = os.Setenv("CORE_PEER_LOCALMSPID", ordererMSP)
	_ = os.Setenv("CORE_PEER_ADDRESS", ordererAddress)
	_ = os.Setenv("CORE_PEER_MSPCONFIGPATH", ordererMSPConfigPath)

	// 获得所有的通道
	var channels = getAllChannels(ordererFilePath)
	log.Println("所有的通道:", channels)

	for _, channel := range channels {
		kafka2raft(channel, ordererAddress)
	}
}

// 获得程序当前执行目录
func currentDir() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(fmt.Sprintf("get current dir err: %s", err))
	}

	return strings.ReplaceAll(dir, "\\", "/")
}

// kafka升级raft
func kafka2raft(channel, ordererAddress string) {
	log.Printf("################开始通道[%s]kafka到raft升级################\n", channel)
	basePath := fmt.Sprintf("out/%s", channel)
	err := os.MkdirAll(basePath, 644)
	if err != nil {
		panic(err)
	}

	// 获取通道配置文件
	cmd := fmt.Sprintf("bin/peer channel fetch config %s/config_block.pb -o %s -c %s", basePath, ordererAddress, channel)
	bash(cmd)
	// 解码该配置文件
	cmd = fmt.Sprintf("bin/configtxlator proto_decode --input %s/config_block.pb --type common.Block | jq .data.data[0].payload.data.config > %s/config.json", basePath, basePath)
	bash(cmd)
	// 获得orderer的addresses,查询有几个orderer节点
	cmd = fmt.Sprintf("bin/configtxlator proto_decode --input %s/config_block.pb --type common.Block | jq .data.data[0].payload.data.config.channel_group.values.OrdererAddresses.value > %s/addresses.json", basePath, basePath)
	bash(cmd)
	var addrs addresses
	addrsData := readFile(fmt.Sprintf("%s/addresses.json", basePath))
	err = json.Unmarshal(addrsData, &addrs)
	if err != nil {
		panic(err)
	}
	log.Println("orderer地址为", addrs)

	// 根据orderer addresses的orderer个数和metadata_template.json文件生成metadata配置
	metadataTemplateStr := string(readFile("metadata_template.json"))
	metadata := strings.ReplaceAll(metadataTemplateStr, `"consenters": null`, fmt.Sprintf(`"consenters": %s`, string(genConsenters(addrs))))
	err = writeFile(fmt.Sprintf("%s/metadata.json", basePath), []byte(metadata))
	if err != nil {
		panic(err)
	}

	updateConfigStr := string(readFile(fmt.Sprintf("%s/config.json", basePath)))
	// 修改kafka为raft
	updateConfigStr = strings.ReplaceAll(updateConfigStr, `"type": "kafka"`, `"type": "raft"`)
	// 增加metadata的raft配置
	updateConfigStr = strings.ReplaceAll(updateConfigStr, `"metadata": null`, fmt.Sprintf(`"metadata": %s`, string(readFile(fmt.Sprintf("%s/metadata.json", basePath)))))
	err = writeFile(fmt.Sprintf("%s/update_config_unformat.json", basePath), []byte(updateConfigStr))
	if err != nil {
		panic(err)
	}
	// 使用jq格式化update_config_unformat.json文件为update_config.json文件
	cmd = fmt.Sprintf("cat %s/update_config_unformat.json | jq . > %s/update_config.json", basePath, basePath)
	bash(cmd)

	// 对原有的配置文件与更新的配置文件进行编码
	cmd = fmt.Sprintf("bin/configtxlator proto_encode --input %s/config.json --type common.Config > %s/config.pb", basePath, basePath)
	bash(cmd)
	cmd = fmt.Sprintf("bin/configtxlator proto_encode --input %s/update_config.json --type common.Config > %s/config_update.pb", basePath, basePath)
	bash(cmd)
	// 计算出两个文件的差异
	cmd = fmt.Sprintf("bin/configtxlator compute_update --channel_id %s --original %s/config.pb --updated %s/config_update.pb > %s/updated.pb", channel, basePath, basePath, basePath)
	bash(cmd)
	// 对该差异文件进行解码，并添加用于更新配置的头部信息
	cmd = fmt.Sprintf("bin/configtxlator proto_decode --input %s/updated.pb --type common.ConfigUpdate > %s/updated.json", basePath, basePath)
	bash(cmd)
	cmd = fmt.Sprintf(`echo '{"payload":{"header":{"channel_header":{"channel_id": "%s", "type":2}},"data":{"config_update":'$(cat %s/updated.json)'}}}' | jq . > %s/updated_envelope.json`, channel, basePath, basePath)
	bash(cmd)
	// 编码updated_envelope.json为Envelope格式的文件
	cmd = fmt.Sprintf("bin/configtxlator proto_encode --input %s/updated_envelope.json --type common.Envelope > %s/updated_envelope.pb", basePath, basePath)
	bash(cmd)
	// 对该文件进行签名操作，用于更新配置
	cmd = fmt.Sprintf("bin/peer channel signconfigtx -f %s/updated_envelope.pb", basePath)
	bash(cmd)
	time.Sleep(time.Second)
	// 提交更新通道配置交易
	cmd = fmt.Sprintf("bin/peer channel update -f %s/updated_envelope.pb -c %s -o %s", basePath, channel, ordererAddress)
	bash(cmd)

	time.Sleep(2 * time.Second)
	log.Printf("################完成通道[%s]kafka到raft升级################\n", channel)
}

// 获得所有的通道
func getAllChannels(ordererFilePath string) []string {
	fileInfos, err := ioutil.ReadDir(ordererFilePath)
	if err != nil {
		panic(err)
	}

	var channels []string
	for _, f := range fileInfos {
		channels = append(channels, f.Name())
	}

	return channels
}

// 调用bash命令
func bash(cmd string) {
	log.Printf("exec bash cmd [%s]\n", cmd)
	bash := exec.Command("/bin/bash", "-c", cmd)
	out, err := bash.Output()
	if err != nil {
		panic(err)
	}

	log.Printf("%s\n", string(out))
}

// 根据orderer addresses的orderer个数和metadata_template.json文件生成metadata配置
func genConsenters(addrs addresses) []byte {
	var consenters consenters
	consenters.Consenters = make([]consenter, 0)
	for _, addr := range addrs.Addresses {
		// orderer1.example.com:7050
		host := strings.Split(addr, ":")[0]
		portStr := strings.Split(addr, ":")[1]
		port, err := strconv.Atoi(portStr)
		if err != nil {
			panic(err)
		}
		// example.com
		domain := strings.SplitN(host, ".", 2)[1]
		certPath := fmt.Sprintf("crypto-config/ordererOrganizations/%s/orderers/%s/tls/server.crt", domain, host)

		base64Cert := base64.StdEncoding.EncodeToString(readFile(certPath))
		consenter := consenter{
			ClientTlsCert: base64Cert,
			Host:          host,
			Port:          port,
			ServerTlsCert: base64Cert,
		}
		consenters.Consenters = append(consenters.Consenters, consenter)
	}

	consentersByte, err := json.Marshal(consenters.Consenters)
	if err != nil {
		panic(err)
	}

	return consentersByte
}

// 读取文件
func readFile(filename string) []byte {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return data
}

// 将数据写到指定文件
func writeFile(filename string, data []byte) error {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 644)
	if err != nil {
		panic(err)
		return err
	}

	n, err := f.Write(data)
	if err == nil && n < len(data) {
		err = io.ErrShortWrite
	}

	if err1 := f.Close(); err == nil {
		err = err1
	}

	return err
}
