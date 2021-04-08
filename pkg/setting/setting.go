package setting

import "github.com/spf13/viper"

//配置
type Setting struct {
	vp *viper.Viper
}

//新建配置
func NewSetting() (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("config")   //配置文件名(无拓展名)
	vp.SetConfigType("yaml")     //配置文件拓展名
	vp.AddConfigPath("configs/") //配置文件路径
	err := vp.ReadInConfig()     //读取配置文件
	if err != nil {
		return nil, err
	}
	return &Setting{vp: vp}, err
}
