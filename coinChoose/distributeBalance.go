package coinChoose

// 均分操作资金

/*
1. 获取各个交易所可操作余额（继而可算出最大可操作余额，平均余额）
2. 获取应操作的两个交易所
3. 循环步骤1获取的所有交易所
	3.1. 如果循环的交易所等于应操作交易所，则判断交易所余额是否大于平均余额，大于则转出（转出量=余额-平均余额），小于不处理
	3.2. 如果循环的交易所不等于应操作交易所，则先判断是否存在余额，存在则直接转出，需判断转入交易所余额是否小于平均余额，小于则操作
		3.2.1. 如果转出量大于可操作余额，则一次性转出全部可操作余额到当前交易所，则跳出后面的循环（另外一个交易所）
		3.2.2. 如果转出量小于可操作性余额，则转出所需余额到当前交易所，剩下的余额一次性转入另外一个交易所
*/