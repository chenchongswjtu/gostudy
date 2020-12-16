使用之前需要将所有的orderer的证书拷贝的这目录下
代码执行过程中需要jq工具

.
├── bin                         // 需要的工具目录
│   ├── configtxlator
│   └── peer
├── core.yaml                   // XLEDGER_CFG_PATH环境变量中需要包含core.yaml文件
├── crypto-config               // 需要的证书,只需要包含所有的orderer的证书
│   └── ordererOrganizations
├── kafka2raft                  // kafka升级raft的可执行文件
├── kafka2raft.go               // 源代码
├── metadata_template.json      // metadata的模板文件
└── readme.txt

使用说明:
./kafka2raft --help
Usage of ./kafka2raft:
  -address string
        orderer的ip地址 (default "127.0.0.1:7050")
  -filepath string
        orderer数据存储路径 (default "/var/production/cetc/orderer1.example.com/orderer/chains")
  -msp string
        orderer的msp id (default "OrdererMSP")
  -mspconfigpath string
        orderer的msp路径 (default "crypto-config/ordererOrganizations/example.com/users/Admin@example.com/msp")

例如:
./kafka2raft --address=127.0.0.1:7050 --filepath=/var/production/cetc/orderer1.example.com/orderer/chains --msp=OrdererMSP --mspconfigpath=crypto-config/ordererOrganizations/example.com/users/Admin@example.com/msp

注意:执行完之后需要重启所有的orderer节点完成升级
