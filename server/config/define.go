package config

type Config struct {
	WorkshopS []*Workshop
}

type Workshop struct {
	Name       string  // 名字
	Product    []*Item // 产品
	ReduceTime int32   // 时间缩减加成（%）
	DoubleProb int32   // 双倍概率加成（%）
}

type Item struct {
	Name    string  // 名字
	Time    int32   // 时间（秒）
	Formula []*Item // 配方
}
