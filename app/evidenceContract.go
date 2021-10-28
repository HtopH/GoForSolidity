package app

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

type EvidenceContract struct {
	ContractHandle
}

var EvidenceContractService = new(EvidenceContract)

func (c *EvidenceContract) InitEvidenceContract() error {
	c.RpcHttpUrl = "http://120.77.139.250:6688"
	c.ContractAddr = common.HexToAddress("0x9bded226fbf003457b6df8ebf4b09a190bd1bb0a")
	c.Abi = "[{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"allow\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"deny\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"hash\",\"type\":\"string\"},{\"name\":\"addr\",\"type\":\"string\"},{\"name\":\"timestamp\",\"type\":\"string\"}],\"name\":\"save\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"hash\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"addr\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"timestamp\",\"type\":\"string\"}],\"name\":\"SaveFileEvent\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"_owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"hash\",\"type\":\"string\"}],\"name\":\"getData\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"hash\",\"type\":\"string\"}],\"name\":\"validateHash\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"
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

func (c *EvidenceContract) Owner() error {
	var out []interface{}
	err = c.Conn.Call(c.Account.CallOpts, &out, "_owner")
	if err != nil {
		fmt.Println("合约调用失败", err)
		return err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	fmt.Printf("%+v\n", out0)
	return nil
}

func (c *EvidenceContract) Save(hash, addr, time string) error {

	res, err := c.Conn.Transact(c.Account.TransactOpts, "save", hash, addr, time)
	if err != nil {
		fmt.Println("合约调用失败", err)
		return err
	}
	fmt.Println(res)
	return nil
}

func (c *EvidenceContract) GetData(hash string) error {
	var out []interface{}
	err = c.Conn.Call(c.Account.CallOpts, &out, "getData", hash)
	if err != nil {
		fmt.Println("合约调用失败", err)
		return err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)
	out2 := *abi.ConvertType(out[2], new(string)).(*string)
	fmt.Println(out0, out1, out2)
	return nil
}
