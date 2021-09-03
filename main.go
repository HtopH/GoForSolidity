package main

import (
	"context"
	"crypto/ecdsa"
	"crypto/md5"
	"encoding/hex"
	"firstTest/services"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"strings"
	time2 "time"
)

var (
	Conn         *services.Evidence
	rpcHttpUrl   = "https://ropsten.infura.io/v3/78dfc9c1d5214c84be3475c227372279"
	rpcWsUrl     = "wss://ropsten.infura.io/ws/v3/78dfc9c1d5214c84be3475c227372279"
	contractAddr = common.HexToAddress("0xC85b37D62f769175484d69fC77033b2d5696dB8f")
	fromAddr     = common.HexToAddress("0x31756974D8adbCb649c6022a2Acc5ffdfE933d1A")
	privateKey   = "ce4d46671da4ef5716c2bc4f6220feb2b5ac8ba012c79b859c3c75da0ce9befd"
	client       *ethclient.Client
	err          error
	timeFormat   = "2006-01-02 15:04:05"
	wsClient     *ethclient.Client
	thisQuery    ethereum.FilterQuery
	ctx          context.Context
)

func init() {
	ctx = context.Background()

	client, err = ethclient.Dial(rpcHttpUrl)
	if err != nil {
		fmt.Println("连接网络失败", err)
	}
	defer client.Close()

	Conn, err = services.NewEvidence(contractAddr, client)
	if err != nil {
		fmt.Println("连接合约失败", err)
	}

	wsClient, err = ethclient.Dial(rpcWsUrl)
	if err != nil {
		fmt.Println("连接wss网络失败", err)
	}
	thisQuery = ethereum.FilterQuery{
		Addresses: []common.Address{contractAddr},
	}
}

func main() {
	readBlockLog()
}

//查询区块
func checkBlock() {
	//获取最新头部块
	//header, err := client.HeaderByNumber(ctx, nil)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(header.Number.String()) // 5671744

	blockNumber := big.NewInt(10963111) //区块高度

	block, err := client.BlockByNumber(ctx, blockNumber)
	if err != nil {
		fmt.Println("查询区块错误", timeFormat)
	}
	fmt.Println(block.Hash()) // 5671744
}

func listenNewBlock() {

}

//获取合约数据
func owner() {
	address, err := Conn.Owner(nil)
	if err != nil {
		fmt.Println("调用合约失败", err)
	} else {
		fmt.Println(address)
	}
}

//获取合约数据
func getDate() {
	a, b, c, err := Conn.GetData(nil, "d51a00ff6d15fab7e73382c8ae1bb7b0")
	if err != nil {
		fmt.Println("调用合约失败", err)
	} else {
		fmt.Println(a, b, c)
	}
}

//写入合约数据
func save() {

	auth, err := getTransactOpts()
	if err != nil {
		return
	}

	str := "这个测试监听事件日志是否可以的"
	hash := getMD5String([]byte(str))
	fmt.Println("算出的hash为", hash)
	time := time2.Now().Format(timeFormat)

	res, err := Conn.Save(auth, hash, str, time)
	if err != nil {
		fmt.Println("调用合约失败", err)
	} else {
		fmt.Println(res)
	}

}

//监听合约日志（可监听，但目前没有监听成功）
func listenSave() {

	logs := make(chan types.Log)
	sub, err := wsClient.SubscribeFilterLogs(ctx, thisQuery, logs)
	if err != nil {
		fmt.Println("创建监听事件错误", err)
		return
	}

	for {
		now := time2.Now().Format(timeFormat)
		fmt.Println("堵塞等待监听结果" + now)
		select {
		case err := <-sub.Err():
			fmt.Println("监听事件错误 "+now, err)
		case vLog := <-logs:
			fmt.Println("监听事件获取日志 "+now, vLog) // pointer to event log
		}
	}
}

//读取合约日志
func readBlockLog() {
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(10000000),
		ToBlock:   big.NewInt(20962888),
		Addresses: []common.Address{
			contractAddr,
		},
	}
	logs, err := client.FilterLogs(ctx, query)
	if err != nil {
		fmt.Println("查询日志失败", err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(services.EvidenceABI)))
	if err != nil {
		fmt.Println("解析日志abi错误", err)
	}

	for _, vLog := range logs {
		res, fErr := contractAbi.Unpack("SaveFileEvent", vLog.Data)
		if fErr != nil {
			fmt.Println("解析日志错误", err)
			continue
		}
		//结构对应上了合约里的event事件输出，所以查询的应该是emit触法日志
		fmt.Println("解析日志结果", res)

	}

}

//组装写入合约的auth（需要消耗邮费）
func getTransactOpts() (*bind.TransactOpts, error) {
	nonce, err := client.PendingNonceAt(ctx, fromAddr)
	if err != nil {
		fmt.Println("获取nonce错误", err)
		return nil, err
	}
	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		fmt.Println("Get gasPrice error", err)
		return nil, err
	}
	priKey, err := crypto.HexToECDSA(privateKey)
	auth := bind.NewKeyedTransactor(priKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300000)
	auth.GasPrice = gasPrice
	return auth, nil
}

// GetAccountAddrFromPriKey 通过私钥解析出账户地址
func GetAccountAddrFromPriKey() common.Address {
	priKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		fmt.Println("私钥解析失败", err)
		return common.Address{}
	}
	pubKey := priKey.Public()
	publicKeyECDSA, ok := pubKey.(*ecdsa.PublicKey)
	if !ok {
		fmt.Println("公钥不是个ecdsa签名公钥类型")
		return common.Address{}
	}
	//通过私钥解析出请求账户地址
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Println("请求账户地址是", fromAddress)
	return fromAddress
}

/*
使用MD5对数据进行哈希运算方法2：使用md5.new()方法
*/
func getMD5String(b []byte) (result string) {
	//1、创建Hash接口
	myHash := md5.New() //返回 Hash interface
	//2、添加数据
	myHash.Write(b) //写入数据
	//3、计算结果
	/*
	  执行原理为：myHash.Write(b1)写入的数据进行hash运算  +  myHash.Sum(b2)写入的数据进行hash运算
	              结果为：两个hash运算结果的拼接。若myHash.Write()省略或myHash.Write(nil) ，则默认为写入的数据为“”。
	              根据以上原理，一般不采用两个hash运算的拼接，所以参数为nil
	*/
	res := myHash.Sum(nil) //进行运算
	//4、数据格式化
	result = hex.EncodeToString(res) //转换为string
	return
}
