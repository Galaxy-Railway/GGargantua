# 任务（job）
在gargantua中，一个任务代表一个完整的测试脚本，用户通过grpc创建并启动一个job后，gargantua将会在协程中执行这个任务。

随时可以通过grpc请求查看或中断一个job。

一个job中必须有且仅有一个step，step中可以递归的含有多个step，具体请参照app层的step包