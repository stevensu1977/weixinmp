# weixinmp

1. 微信公众平台接入-URL验证
	微信服务器会根据你填写的URL, 发送一个GET请求, 包含
	
	```
	 signature #加密签名
	 timestamp #时间戳
	 nonce     #随机数
	 echostr   #随机字符串
	 
	 也就是weeixin服务器会把token,timestamp,noce 字典排序后通过sha1做一个signature签名
	 
	 开发者需要用 token, temestamp,nonce 3个字符串拼接起来,并且三个参数进行字典序排序,做一个sha1,如果结果等于signature则通过验证,需要向服务器返回echostr
	```
	
	
2. 微信公众平台接入-URL验证 