# ASM 饮料自动售货机模拟软件

## 1. 简介
ASM是一个自动饮料售货机的模拟软件，使用iris架构进行实现，包括售货端和管理端，硬件方面，提供投币机，货架，加热/制冷装置的接口

## 2. 安装

### 2.1 依赖

- mysql 8.0
- golang v1.14 或者更高 (如果使用源文件编译)

### 2.2 配置

- 使用mysql新建一个数据库 
- 下载ASM源代码或二进制文件
- 更改`ASM/config/config`内的配置项，填写数据库信息，管理员帐号以及货架饮料种类数量
- 执行程序
    
    通过源文件
    ```bash
    $ cd ASM
    $ go mod download 
    $ go run main.go
    ```
    通过二进制文件
    ```bash
    $ ./ASM
    ```
## 3. 访问

- 通过`localhost:10101/sell`访问售卖主页
- 通过`localhost:10101/admin`访问管理员主页