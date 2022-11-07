# go-indicator



https://medium.datadriveninvestor.com/how-to-detect-support-resistance-levels-and-breakout-using-python-f8b5dac42f21

https://colab.research.google.com/github/yongghongg/stock-screener/blob/main/stock_breakout_demo.ipynb



is_far_from_level(current_max,pivots,df):

def is_far_from_level(value, levels, df):
ave =  np.mean(df['High'] - df['Low']) # 全部K的平均值
return np.sum([abs(value - level) < ave for _, level in levels]) == 0

	true 差值是要小于 区间平均值
	flase 差值大于区间平均值
	abs(value - level) 当前最大值和其他最大值的绝对值小于均值