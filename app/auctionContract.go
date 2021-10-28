package app

import (
	"context"
	"fmt"
	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"strings"
)

type AuctionContract struct {
	ContractHandle
}

var AuctionContractService = new(AuctionContract)

func (c *AuctionContract) InitAuctionContract() error {
	c.RpcHttpUrl = "http://183.3.158.22:6688"
	c.ContractAddr = common.HexToAddress("0x72c3ba92caec5186756c28c732dc102369c8d43d")
	c.Abi = "[{\"constant\":false,\"inputs\":[],\"name\":\"bid\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"auctionEnd\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_biddingTime\",\"type\":\"uint256\"}],\"name\":\"reSetEnd\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"highestBidder\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"auctionend\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"highestBid\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_biddingTime\",\"type\":\"uint256\"},{\"name\":\"_beneficiary\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"bigger\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amout\",\"type\":\"uint256\"}],\"name\":\"HighestBidCreased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"AutionEnded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TestEvent\",\"type\":\"event\"}]"
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

//ReadBlockLog 读取合约日志（在指定的区块高度范围）（日志即合约所触发的事件）
func (c *AuctionContract) ReadBlockLog() error {

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(882579),
		ToBlock:   big.NewInt(2882579),
		Addresses: []common.Address{
			c.ContractAddr,
		},
	}
	logs, err := c.Client.FilterLogs(c.Ctx, query)
	if err != nil {
		fmt.Println("查询日志失败", err)
		return err
	}

	if len(logs) > 0 {
		contractAbi, err := abi.JSON(strings.NewReader(c.Abi))
		if err != nil {
			fmt.Println("解析日志abi错误", err)
			return err
		}
		for _, vLog := range logs {
			res, Err := contractAbi.Unpack("TestEvent", vLog.Data)
			if Err != nil {
				fmt.Println("解析日志错误", Err)
				continue
			}
			//结构对应上了合约里的event事件输出，所以查询的应该是emit触法日志
			fmt.Printf("日志区块是：%v,解析日志结果:%+v\n", vLog.BlockNumber,res)

		}
	}

	return nil
}

func (c *AuctionContract) ReSetEnd() error {
	endTime := big.NewInt(300)
	res, err := c.Conn.Transact(c.Account.TransactOpts, "reSetEnd", endTime)
	if err != nil {
		fmt.Println("合约调用失败", err)
		return err
	}
	fmt.Println(res)
	return nil
}

func (c *AuctionContract) HighestBidder() error {
	var out []interface{}
	err = c.Conn.Call(c.Account.CallOpts, &out, "highestBidder")
	if err != nil {
		fmt.Println("合约调用失败", err)
		return err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	fmt.Printf("%+v\n", out0)
	return nil
}

func (c *AuctionContract) HighestBid() error {
	var out []interface{}
	err = c.Conn.Call(c.Account.CallOpts, &out, "highestBid")

	if err != nil {
		return err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	fmt.Println(out0)
	return err

}

func (c *AuctionContract) Bid() error {
	c.Account.TransactOpts.Value = big.NewInt(20000002)
	res, err := c.Conn.Transact(c.Account.TransactOpts, "bid")
	if err != nil {
		fmt.Println("合约调用失败", err)
		return err
	}
	fmt.Println(res)
	return nil
}

func (c *AuctionContract) AuctionEnd() error {
	var out []interface{}
	err = c.Conn.Call(c.Account.CallOpts, &out, "auctionEnd")

	if err != nil {
		return err
	}
	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	fmt.Println(out0)
	return err
}

func (c *AuctionContract) ToDoAuctionEnd() error {
	res, err := c.Conn.Transact(c.Account.TransactOpts, "auctionend")
	if err != nil {
		fmt.Println("合约调用失败", err)
		return err
	}
	fmt.Println(res)
	return nil
}

func (c *AuctionContract) Withdraw() error{
	res, err := c.Conn.Transact(c.Account.TransactOpts, "withdraw")
	if err != nil {
		fmt.Println("合约调用失败", err)
		return err
	}
	fmt.Println(res)
	return nil
}

//UnpackLog 解析日志
func (c *AuctionContract)UnpackLog()  {
	
}