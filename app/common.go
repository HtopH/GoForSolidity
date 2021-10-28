package app

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"strings"
)

type ContractHandle struct {
	RpcHttpUrl   string
	ContractAddr common.Address
	Abi          string
	Client       *ethclient.Client
	Conn         *bind.BoundContract
	Ctx          context.Context
	Account      AccountInfo
}

type AccountInfo struct {
	FromAddr     common.Address
	PrivateKey   string
	TransactOpts *bind.TransactOpts
	CallOpts     *bind.CallOpts
	FilterOpts   *bind.FilterOpts
}

var err error

func (c *ContractHandle) getConn() error {
	c.Client, err = ethclient.Dial(c.RpcHttpUrl)
	if err != nil {
		fmt.Println("连接网络失败", err)
		return err
	}

	parsed, err := abi.JSON(strings.NewReader(c.Abi))
	if err != nil {
		fmt.Println("解析abi错误", err)
		return err
	}
	c.Conn = bind.NewBoundContract(c.ContractAddr, parsed, c.Client, c.Client, c.Client)
	return nil
}

func (c *ContractHandle) getTransactOpts() error {

	nonce, err := c.Client.PendingNonceAt(c.Ctx, c.Account.FromAddr)
	if err != nil {
		fmt.Println("获取nonce错误", err)
		return err
	}
	gasPrice, err := c.Client.SuggestGasPrice(c.Ctx)
	if err != nil {
		fmt.Println("Get gasPrice error", err)
		return err
	}
	priKey, err := crypto.HexToECDSA(c.Account.PrivateKey)
	c.Account.TransactOpts = bind.NewKeyedTransactor(priKey)
	c.Account.TransactOpts.Nonce = big.NewInt(int64(nonce))
	c.Account.TransactOpts.Value = big.NewInt(0)
	c.Account.TransactOpts.GasLimit = uint64(300000)
	c.Account.TransactOpts.GasPrice = gasPrice
	return nil
}

/*
GetMD5String
使用MD5对数据进行哈希运算方法2：使用md5.new()方法
*/
func (c *ContractHandle) GetMD5String(b []byte) (result string) {
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
