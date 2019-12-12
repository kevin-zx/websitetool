package regexInfopaser

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMatchICPNO(t *testing.T) {
	text := `石家庄正则肥业有限公司为您免费提供有机肥厂家、石家庄有机肥生产厂家、石家庄蚯蚓粪有机肥等相关信息发布和最新资讯，敬请关注！  您暂无新询盘信息！

石家庄正则肥业有限公司

服务热线：

联系电话：18733100889 13803376119

首页关于我们产品中心新闻资讯成功案例资质荣誉联系我们

热门搜索： 石家庄有机肥 蚯蚓粪有机肥 生物有机肥


 站内   站外    
产品目录
生物有机肥
石家庄生物有机肥
生物有机肥厂家
发酵有机肥
石家庄发酵有机肥
发酵有机肥厂家
微生物菌肥
微生物菌肥厂家
石家庄微生物菌肥
蔬菜有机肥
蔬菜有机肥厂家
石家庄有机肥生产
水果有机肥
水果有机肥厂家
石家庄有机肥批发
晒干鸡粪
晒干鸡粪厂家
石家庄晒干鸡粪
蚯蚓养殖
人工蚯蚓养殖
蚯蚓养殖厂家
蚯蚓粪
河北蚯蚓粪有机肥
蚯蚓粪厂家
产品展示更多
河北有机肥
河北有机肥
石家庄有机肥厂家
石家庄有机肥厂家
石家庄有机肥
石家庄有机肥
石家庄生物有机肥批发
石家庄生物有机肥批发
石家庄生物有机肥厂家
石家庄生物有机肥厂家
水果有机肥
水果有机肥
生物有机肥厂家
生物有机肥厂家
石家庄微生物菌肥
石家庄微生物菌肥
石家庄蚯蚓粪有机肥
石家庄蚯蚓粪有机肥

公司简介ABOUT US
石家庄正则肥业有限公司位于全国养殖大省-河北省，鸡粪资源充足，公司主要从事生物有机肥的研究和生产，拥有多名专业技术人员，运用高科技，较为先进的设备，经高温发酵,研制成瓜果蔬菜绿色食品专用有机肥。 石家庄正则有机肥厂家近年引进蚯蚓养殖技术，研发和生产蚯蚓粪有机肥、生物有机肥、羊粪有机肥、颗粒有机肥、瓜果蔬菜专用肥鱼肥大力肥水王晒干鸡粪发酵鸡粪蚯蚓粪等。 公司秉承“顾客至上，锐意进取”的经营理念，为广大客户提供优质的服务。欢迎惠顾! 绿色食品的摇篮国人健康的保证!...【详情】

精品展示
 有机肥厂家
有机肥厂家
 石家庄生物有机肥批发
石家庄生物有机肥批发
 石家庄生物有机肥厂家
石家庄生物有机肥厂家
 藁城生物有机肥
藁城生物有机肥
 生物有机肥
生物有机肥
 河北蚯蚓粪有机肥
河北蚯蚓粪有机肥
 石家庄有机肥厂
石家庄有机肥厂
 生物有机肥厂家
生物有机肥厂家
 石家庄生物有机肥
石家庄生物有机肥
 石家庄发酵有机肥
石家庄发酵有机肥
 有机肥厂家
有机肥厂家
 石家庄生物有机肥批发
石家庄生物有机肥批发
 石家庄生物有机肥厂家
石家庄生物有机肥厂家
 藁城生物有机肥
藁城生物有机肥
 生物有机肥
生物有机肥
 河北蚯蚓粪有机肥
河北蚯蚓粪有机肥
 石家庄有机肥厂
石家庄有机肥厂
 生物有机肥厂家
生物有机肥厂家
 石家庄生物有机肥
石家庄生物有机肥
 石家庄发酵有机肥
石家庄发酵有机肥
公司新闻更多

石家庄有机肥非比寻常
石家庄有机肥非比寻常，有机肥对土壤起到了很大的作用，让我们的农作物能够茁壮的成长，让庄稼有个好收成有机肥包含了农作物生长所需要的营养成分和各种有益元素，而且其中… 【详情】

2019-10-29微生物菌肥在土壤中起到的作用
2019-10-17国家大力提倡有机肥的施用
2019-10-09有机肥不能直接施用的原因
2019-09-09石家庄生物有机肥恪尽职守
2019-08-21石家庄微生物菌肥不可小觑
2019-08-17石家庄有机肥是天然肥料只重选
2019-07-27石家庄生物有机肥肥沃中华棒棒哒
2019-07-17有机肥厂家不忘初心名不虚传
2019-07-13蚯蚓粪有机肥别具匠心
2019-07-11石家庄有机肥谨小慎微
2019-07-04石家庄微生物菌肥无懈可击
2019-06-30石家庄有机肥致富千万家给你幸福无限
2019-06-27蚯蚓粪有机肥引领绿色新时代棒棒哒
2019-06-26有机肥厂家为农民的的庄稼好丰收贡献一份力
2019-06-14石家庄有机肥具有无穷的潜力
2019-06-13河北有机肥造福万家创造奇迹
2019-05-31蚯蚓粪有机肥品质传承铸造精彩
2019-05-31有机肥厂家大展宏图送富到农家
2019-05-23石家庄生物有机肥创就辉煌
2019-05-21石家庄有机肥厂家前景广阔势不可挡
2019-05-21蚯蚓粪有机肥良心品质
2019-05-16石家庄蚯蚓粪有机肥良心品质安心选择
2019-05-11石家庄有机肥的独到之处
2019-05-10有机肥厂家干劲十足为让植物生机勃勃
2019-04-30河北有机肥绿色植物的好帮手引领绿色新时代
2019-04-30蚯蚓粪有机肥好评如潮魅力无限
2019-04-27有机肥厂家大有可观前景广阔
2019-04-22石家庄有机肥施用误区引起重视
2019-04-15生物有机肥厂家发展绿色农业势不可挡
2019-04-11石家庄有机肥厂家势在必行引领绿色新时代
2019-04-10石家庄蚯蚓养殖不负众望
2019-04-04石家庄生物有机肥深受好评
2019-03-28有机肥厂家不负众望让土地增添生机
2019-03-26石家庄有机肥赞不绝口
2019-03-16蚯蚓粪有机肥魅力无限
2019-03-14石家庄有机肥有口皆碑
2019-03-12丰收致富的好宝物河北有机肥
2019-02-26有机肥厂家与你共赢
2019-02-22石家庄生物有机肥众目可睹
2019-02-21蚯蚓粪有机肥不得不看
2019-02-18石家庄有机肥初心不改
2019-02-14石家庄有机肥好处多多
2019-02-13石家庄生物有机肥深得人心
2019-02-11有机肥厂家众口称赞
2019-01-28蚯蚓粪有机肥前景可观
2019-01-16河北有机肥提醒农民增收
2019-01-10石家庄有机肥致富全靠它没毛病
2018-12-18承载绿色梦想的石家庄有机肥成就幸福生活
2018-11-24生物有机肥使你的爱花更漂亮
2018-11-22果农看这里，生物有机肥施与浅施讲究多
2018-11-15土地肥沃全靠石家庄有机肥
2018-11-06石家庄发酵有机肥无懈可击
2018-11-02石家庄生物有机肥不能没有你
2018-10-27石家庄蚯蚓粪有机肥庄稼丰收处处好
2018-10-24石家庄有机肥厂家令你刮目相看
2018-10-12石家庄有机肥农田好帮手妥妥的
2018-09-29石家庄有机肥万众瞩目
2018-09-28蚯蚓粪有机肥倍受青睐
2018-09-26蚯蚓粪有机肥万里挑一
2018-09-25石家庄有机肥优势尽显
2018-09-20石家庄微生物菌肥超乎想象
2018-09-11有机肥厂家保护你的农田万无一失
2018-09-08生物有机肥是经典绝不再现
2018-09-06石家庄生物有机肥大有作为
2018-08-28石家庄有机肥威力无穷
2018-08-22蚯蚓粪有机肥未来可期
2018-08-16蚯蚓粪有机肥本领高强
2018-08-06石家庄有机肥厂家面面俱到
2018-07-28石家庄有机肥给农田加点料
2018-07-27蚯蚓粪有机肥所向披靡
2018-05-25微生物菌肥的主要优势
2018-05-25生物有机肥的生产工艺
行业新闻更多

石家庄有机肥非比寻常
石家庄有机肥非比寻常，有机肥对土壤起到了很大的作用，让我们的农作物能够茁壮的成长，让庄稼有个好收成有机肥包含了农作物生长所需要的营养成分和各种有益元素，而且其中… 【详情】

2019-11-08石家庄有机肥非比寻常
2019-10-15蚯蚓粪有机肥与普通有机肥的不同之处
2018-11-20农民早知道：生物有机肥将是未来农业发展的
2018-07-31有机肥厂家告诉您什么是生物有机肥
2018-07-31蚯蚓粪有机肥为什么能够抑制植物土传病害
2018-07-30使用石家庄微生物有机肥对农作物的好处
2018-07-14关于晒干鸡粪的用途
2018-07-10蚯蚓粪有机肥绿色植物好肥料
2018-06-19生物有机肥对土壤的5大作用
2018-06-12石家庄发酵有机肥的产品功效有那些
2018-06-08生物有机肥能不能代替化肥
2018-06-05使用微生物菌肥有什么作用
常见问题更多
Q蚯蚓粪才是有机肥之王…
A随着人们对蚯蚓粪研究的不断深入，发现蚯蚓粪不仅是蚯蚓的代谢废…
Q有机肥适用于所有的花卉吗…
A有机肥适用于所有的花卉吗？其实对于花卉来说，施肥的合理与否，…
Q石家庄有机肥的处理方法…
A不仅能为农作物提供营养，而且肥效长，可增加和更新土壤有机质，…
Q石家庄微生物菌肥不可或缺妥妥的…
A石家庄微生物菌肥不可或缺妥妥的，微生物菌肥肥料是一种在农业生…
Q讲课啦，蚯蚓粪有机肥的称王之路…
A讲课啦，今天我们说一下蚯蚓粪有机肥的称王之路，为什么越来越多…
Q施用石家庄有机肥需要注意什么…
A正常的石家庄有机肥施用量不会伤害植物根部，但是有机肥施多了会…
商盟成员 / BUSINESS
成都广告设计制作| 竹木纤维墙板| 景德镇山羊| 河南空压机| 美标德标欧标澳标电缆| 浙江分布式光伏电站| 成都金丝楠木| 蜜桃树苗| 沈阳低压柜| 沈阳锌钢护栏| 硅质聚苯板| 沈阳水泵维修| 四川钢格板厂家| 成都到西安物流| 防治臭虫| 深圳别克商务车出租| 动物行为学仪器| 进口木方| 无转子硫化仪| 云南PE给水管| 贵州园林机械| 鱼台大米| 花溪草莓采摘| 猪催肥| 泰来大米| 智能玻璃温室大棚造价| 猪粪固液分离机| 催情散| 蚯蚓粪有机肥| 山东丹参种植基地| 蜂糖李果苗| 经济林苗木| 石家庄有机肥|
首页| 关于我们| 产品中心| 新闻资讯| 网站地图| XML| 手机网站| 联系我们

Copyright©www.zhengzefeiye.com(复制链接) 石家庄正则肥业有限公司   备案号：冀ICP备18016375号-1

联系电话： 业务电话：18733100889 13803376119 邮箱：451637087@qq.com

蚯蚓粪有机肥哪家好？供应订做多少钱？生物有机肥怎么样？诚信公司专业以批发价格大量现货提供蚯蚓粪有机肥、石家庄有机肥生产厂家、石家庄蚯蚓粪有机肥等品质优良的产品及报价，欢迎来电生产定制！

热门城市推广: 邯郸 邢台 保定 承德 沧州 时代互动

       Powered by筑巢

联系手机

18733100889

13803376119

分享

分享到：0`
	icpNos := MatchICPNO(text)
	fmt.Println(icpNos)
}

func TestMatchEmail(t *testing.T) {
	text := "电话:010-56143201 邮箱:service@bangbangzhixin.cn 官网:http://www.bangbangzhixin.cn 地址:北京市东城区东四十条甲22号1号楼A1705 公司..."
	email := MatchEmail(text)
	fmt.Printf("%v\n", email)
	if len(email) == 0 || email[0] != "service@bangbangzhixin.cn" {
		t.Fail()
	}
}
func TestMatchtPhone(t *testing.T) {
	text := `
注册资本	80万(元)	实缴资本	-
法定代表人	卢隽	经营状态	开业 
曾用名	-	所属行业	-
统一社会信用代码	91440300MA5FQFB615	纳税人识别号	91440300MA5FQFB615
工商注册号	139126496541	组织机构代码	MA5FQFB6-1
登记机关	深圳市市场监督管理局13912649654	成立日期	2019-08-05
企业类型	有限责任公司(自然人独资)	营业期限	2019-08-05 至 无固定期限
行政区划	广东省	审核/年检日期	2019-08-05
注册地址	深圳市福田区香蜜湖街道香岭社区紫竹六道78号竹盛花园15栋102查看地图
经营范围	一般经营项目是：婚庆咨询服务；婚礼策划；婚庆礼仪服务；婚纱摄影服务；会展会务服务；文化艺术交流活动策划；企业形象策划；市场营销策划；商务信息咨询...展开`
	email := MatchtNums(text)
	fmt.Printf("%v\n", email)

}

func TestMatchTelephone(t *testing.T) {
	text := "电话:86-0512-62581818 邮箱:service@bangbangzhixin.cn 官网:http://www.bangbangzhixin.cn 地址:北京市东城区东四十条甲22号1号楼A1705 公司..."
	telPhone := MatchTelephone(text)
	fmt.Printf("%v\n", telPhone)
	if len(telPhone) == 0 || telPhone[0] != "010-56143201" {
		t.Fail()
	}
}

func TestMatch(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "1",
			args: struct{ text string }{text: `
注册资本	80万(元)	实缴资本	-
法定代表人	卢隽	经营状态	开业 
曾用名	-	所属行业	-
统一社会信用代码	91440300MA5FQFB615	纳税人识别号	91440300MA5FQFB615
工商注册号	13912649654	组织机构代码	MA5FQFB6-1
登记机关	深圳市市场监督管理局13912649654	成立日期	2019-08-05
企业类型	有限责任公司(自然人独资)	营业期限	2019-08-05 至 无固定期限
行政区划	广东省	审核/年检日期	2019-08-05
注册地址	深圳市福田区香蜜湖街道香岭社区紫竹六道78号竹盛花园15栋102查看地图
经营范围	一般经营项目是：婚庆咨询服务；婚礼策划；婚庆礼仪服务；婚纱摄影服务；会展会务服务；文化艺术交流活动策划；企业形象策划；市场营销策划；商务信息咨询...展开`},
			want: []string{"13912649654"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MatchPhone(tt.args.text); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MatchPhone() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatchCreditCode(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "1",
			args: struct{ text string }{text: `
注册资本	80万(元)	实缴资本	-
法定代表人	卢隽	经营状态	开业 
曾用名	-	所属行业	-
统一社会信用代码	91440300MA5FQFB615	纳税人识别号	91440300MA5FQFB615
工商注册号	13912649654	组织机构代码	MA5FQFB6-1
登记机关	深圳市市场监督管理局13912649654	成立日期	2019-08-05
企业类型	有限责任公司(自然人独资)	营业期限	2019-08-05 至 无固定期限
行政区划	广东省	审核/年检日期	2019-08-05
注册地址	深圳市福田区香蜜湖街道香岭社区紫竹六道78号竹盛花园15栋102查看地图
经营范围	一般经营项目是：婚庆咨询服务；婚礼策划；婚庆礼仪服务；婚纱摄影服务；会展会务服务；文化艺术交流活动策划；企业形象策划；市场营销策划；商务信息咨询...展开`},
			want: []string{"91440300MA5FQFB615"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MatchCreditCode(tt.args.text); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MatchCreditCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatchOrganizationNo(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "1",
			args: struct{ text string }{text: `
注册资本	80万(元)	实缴资本	-
法定代表人	卢隽	经营状态	开业 
曾用名	-	所属行业	-
统一社会信用代码	91440300MA5FQFB615	纳税人识别号	91440300MA5FQFB615
工商注册号	13912649654	组织机构代码	MA5FQFB6-1
登记机关	深圳市市场监督管理局13912649654	成立日期	2019-08-05 MA5FQFB6-1
企业类型	有限责任公司(自然人独资)	营业期限	2019-08-05 至 无固定期限
行政区划	广东省	审核/年检日期	2019-08-05
注册地址	深圳市福田区香蜜湖街道香岭社区紫竹六道78号竹盛花园15栋102查看地图
经营范围	一般经营项目是：婚庆咨询服务；婚礼策划；婚庆礼仪服务；婚纱摄影服务；会展会务服务；文化艺术交流活动策划；企业形象策划；市场营销策划；商务信息咨询...展开`},
			want: []string{"MA5FQFB6-1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MatchOrganizationNO(tt.args.text); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MatchOrganizationNO() = %v, want %v", got, tt.want)
			}
		})
	}
}