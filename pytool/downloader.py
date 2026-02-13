import tushare as ts
import pandas as pd

# 1. 设置你的 Token (只需设置一次)
ts.set_token('2e796ce1b3ede0da9ef3d2db608f25e7c68d710cf871af55041141a7')

# 2. 初始化接口
pro = ts.pro_api()

# 3. 获取指数日线数据 (以沪深300为例)
df = pro.index_daily(ts_code='000300.SH', 
                     start_date='20250101', 
                     end_date='20251231')

# 4. 查看数据
print(df.head())

# 5. 保存到 Excel
df.to_excel('hs300_data.xlsx', index=False)