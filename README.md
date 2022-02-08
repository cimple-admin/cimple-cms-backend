# cimple-cms-backend
cimple cms server side

### 错误码设置方式
按照模块或者中间件每个单独定义一个前缀，方便追溯，错误码长度定义方式为 前缀+4位数字。没有错误错误码就是 0

### Json 相应
采用 http code + 自有 code + data + mesasge 的方式
也就是说大部分的http code 都是 200，在一些其他情况采用其他的值，比如未登录用401。大部分情况都采用自有 code 去鉴别错误。数据分为两部分 `data` 和 `message`。`data` 为响应数据。 `message` 为提示信息。