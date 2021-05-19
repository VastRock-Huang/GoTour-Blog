package setting

import "time"

//配置属性的结构体
//PS:结构体末尾的S表明为结构体名,与全局变量的实例区分

//服务器配置
type ServerSettingS struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

//应用配置
type AppSettingS struct {
	DefaultPageSize      int
	MaxPageSize          int
	LogSavePath          string
	LogFileName          string
	LogFileExt           string
	UploadSavePath       string
	UploadServerUrl      string
	UploadImageMaxSize   int
	UploadImageAllowExts []string
	DefaultContextTimeout time.Duration
}

//数据库配置
type DatabaseSettingS struct {
	DBType       string
	Username     string
	Password     string
	Host         string
	DBName       string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
	TablePrefix  string
}

//JWT配置
type JWTSettingS struct {
	Secret string
	Issuer string
	Expire time.Duration
}

//邮件发送配置
type EmailSettingS struct {
	Host string
	Port int
	UserName string
	Password string
	IsSSL bool
	From string
	To []string
}

//存放每一部分配置的映射
//键值是接口,实际上为配置结构体的地址
var sections = make(map[string]interface{})

//读取配置到结构体
func (s *Setting) ReadSection(k string, v interface{}) error {
	//对标题为k的配置部分解码到v中
	if err := s.vp.UnmarshalKey(k, v); err != nil {
		return err
	}
	//将配置记录到映射中
	if _, ok:=sections[k]; !ok {
		sections[k]=v
	}
	return nil
}

//读取所有部分的配置
func (s *Setting) ReloadAllSection() error {
	for k,v :=range sections{
		//由于sections中键值实际上是配置结构体的地址
		//因此此处读取配置后会存到全局的配置结构体中
		if err := s.ReadSection(k,v); err != nil {
			return err
		}
	}
	return nil
}
