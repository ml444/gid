# gid
Global ID Component Based on Snowflake Algorithm.


```text
ID类型分为最大峰值和最小粒度
【最大峰值】：
| 版本 | 类型 | 生成方式 | 秒级时间 | 序列号 | 机器ID |
| 63   | 62  | 60-61  | 59-30   | 29-10 | 0-9   |

【最小粒度】：
| 版本 | 类型 | 生成方式 | 秒级时间 | 序列号 | 机器ID |
| 63   | 62  | 60-61  | 59-20   | 19-10 | 0-9   |
```
