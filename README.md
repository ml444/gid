# gid
Global ID Component Based on Snowflake Algorithm.

## Bit patterns for ID
1. Classic snowflake pattern
    ```
    【Snowflake】：
    | version | time(ms) | machineID | sequence |
    | 63      | 62-22    | 21-12     | 11-0     |
    ```

2. Maximum Peak pattern
    ```
    【Max-Peak】：
    | version | generate-mode | time(s) | machineID | sequence |
    | 63      | 62-61         | 60-30   | 29-20     | 19-0     |
    ```

3. Minimum granularity pattern
    ```
    【Min-Granularity】：
    | version | generate-mode | time(ms) | machineID | sequence |
    | 63      | 62-61         | 60-20    | 19-10     | 9-0      |
    ```
