# iptool
功能：通过golang 将ip地址转化成所属区域 。  
思路：做成单独服务，接受远程http请求，从第三方ip查询接口获取ip详细，还回实际区域信息。  
说明：为何不直接用第三方接口获取，反而大张旗鼓的在架一层服务，我想单独做一个服务器去处理这件事，还回的数据或缓存或特殊处理由自己而定。  

