# 简介
本模块是基于golang原生的http server+myql数据开发的一个curd操作demo。demo采用mvc模型进行分包处理
各层操作代码。
## 创建数据库
```
CREATE DATABASE patientinfo;
```
## 创建表

```
CREATE TABLE tbl_patientinfo (
    patient_id VARCHAR(10) PRIMARY KEY NOT NULL,
    name    VARCHAR(255) NOT NULL,
    lastname VARCHAR(255) NOT NULL,
    gender  VARCHAR(6) NOT NULL,
    age INTEGER NOT NULL
);
```

## 添加数据
```
INSERT INTO tbl_patientinfo (patient_id, name, lastname, gender, age) VALUES
  ('12042019', 'Maria', 'Bennet', 'Female', 31),
  ('12052019', 'Julia', 'Jennison', 'Female', 25),
  ('12062019', 'Mark', 'Davidson', 'Male', 40);
```
