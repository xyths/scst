# 说明
`corswap`目录中的`sol`文件拷贝自<https://github.com/coral-dex/corswap>中的`contract`目录，拷贝时间2020-10-19日。

此目录的主要目的是备份当时的源码文件，其中部分源码因为不能编译会做少许改动（版本、语法等，不涉及逻辑和接口）。

该目录中的`abi`文件截取自同一仓库的`src/component/abi.js`文件，自动生成的`Go`文件，部分函数因不能编译，被注释掉，谨慎使用。

```Go
func (_SwapExchange *SwapExchangeTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	//return _SwapExchange.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
	return nil, nil
}
```
