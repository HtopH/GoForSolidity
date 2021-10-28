package app

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

type FirstContract struct {
	ContractHandle
}

var FirstContractService = new(FirstContract)

func (c *FirstContract) InitFirstContract() error {
	c.RpcHttpUrl = "http://183.3.158.22:6688"
	c.ContractAddr = common.HexToAddress("0xd81c652e72c795371b0e53ad6a130e887063e699")
	c.Abi = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getName\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_str\",\"type\":\"string\"}],\"name\":\"setName\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_str\",\"type\":\"string\"}],\"name\":\"TellService\",\"type\":\"event\"}]"
	c.Account.FromAddr = common.HexToAddress("0xbeda79c48d8dfa79b30cb843df6d2f0cf1a9b80f")
	c.Account.PrivateKey = "b1c70df3392c28116e023015da64407c0ea127ff314d8e3daafde53621784221"
	c.Ctx = context.Background()
	err = c.getConn()
	if err != nil {
		return err
	}
	err = c.getTransactOpts()
	if err != nil {
		return err
	}
	return nil
}

// Name 对照合约方法调用
func (c *FirstContract) Name() {
	var out []interface{}
	err = c.Conn.Call(c.Account.CallOpts, &out, "name")
	if err != nil {
		fmt.Println("合约调用失败", err)
		return
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	fmt.Printf("%+v\n", out0)
}

// GetName 对照合约getName调用
func (c *FirstContract) GetName()error {
	var out []interface{}
	err = c.Conn.Call(c.Account.CallOpts, &out, "getName")
	if err != nil {
		fmt.Println("合约调用失败", err)
		return err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	fmt.Println(out0)
	return nil
}

func (c *FirstContract) SetName() error {

	res, err := c.Conn.Transact(c.Account.TransactOpts, "setName", "这次试试封装的方法")
	if err != nil {
		fmt.Println("合约调用失败", err)
		return err
	}
	fmt.Println(res)
	return nil
}
