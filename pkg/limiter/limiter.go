package limiter

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"time"
)

//限流器
type Limiter struct {
	limiterBuckets map[string]*ratelimit.Bucket	//令牌桶和名称与实体的映射集
}

//令牌桶规则
type LimiterBucketRule struct {
	Key string	//令牌桶关键字
	FillInterval time.Duration	//放入令牌的时间间隔
	Capacity int64	//令牌桶容量
	Quantum int64	//达到时间间隔放入的令牌数量
}

//限流器接口,根据不同接口选择不同限流器
type LimiterIface interface {
	Key(c *gin.Context) string	//限流器的键值对名称
	GetBucket(key string) (*ratelimit.Bucket, bool)	//获取令牌桶
	AddBuckets(rules ...LimiterBucketRule) LimiterIface	//新增多个令牌桶
}
