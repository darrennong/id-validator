// This file is part of the guanguans/id-validator.
// (c) guanguans <ityaozm@gmail.com>
// This source file is subject to the MIT license that is bundled.

package data

type AddressTimeline struct {
	Address   string
	StartYear int
	EndYear   int
}

// AddressCodeTimeline 行政区划代码（地址码）更新时间线
// 中华人民共和国民政部权威数据
// 注1：台湾省、香港特别行政区和澳门特别行政区暂缺地市和区县信息
// 注2：每月发布的区划变更表是根据区划变更地的统计人员在统计信息系统更新后的情况所绘制，与区划变更文件发布的时间有一定的延迟性，但在每年的最后一次发布变更情况后与区划全年变更文件保持一致。
// Data Source: http://www.mca.gov.cn/article/sj/xzqh/
var AddressCodeTimeline = map[int][]AddressTimeline{
	110000: {
		{
			Address: "北京市",
		},
	},
	110101: {
		{
			Address: "东城区",
		},
	},
	110102: {
		{
			Address: "西城区",
		},
	},
	110103: {
		{
			Address: "崇文区",
			EndYear: 2009,
		},
	},
	110104: {
		{
			Address: "宣武区",
			EndYear: 2009,
		},
	},
	110105: {
		{
			Address: "朝阳区",
		},
	},
	110106: {
		{
			Address: "丰台区",
		},
	},
	110107: {
		{
			Address: "石景山区",
		},
	},
	110108: {
		{
			Address: "海淀区",
		},
	},
	110109: {
		{
			Address: "门头沟区",
		},
	},
	110110: {
		{
			Address: "燕山区",
			EndYear: 1985,
		},
	},
	110111: {
		{
			Address:   "房山区",
			StartYear: 1986,
		},
	},
	110112: {
		{
			Address:   "通州区",
			StartYear: 1997,
		},
	},
	110113: {
		{
			Address:   "顺义区",
			StartYear: 1998,
		},
	},
	110114: {
		{
			Address:   "昌平区",
			StartYear: 1999,
		},
	},
	110115: {
		{
			Address:   "大兴区",
			StartYear: 2001,
		},
	},
	110116: {
		{
			Address:   "怀柔区",
			StartYear: 2001,
		},
	},
	110117: {
		{
			Address:   "平谷区",
			StartYear: 2001,
		},
	},
	110118: {
		{
			Address:   "密云区",
			StartYear: 2015,
		},
	},
	110119: {
		{
			Address:   "延庆区",
			StartYear: 2015,
		},
	},
	110201: {
		{
			Address: "昌平县",
			EndYear: 1981,
		},
	},
	110202: {
		{
			Address: "顺义县",
			EndYear: 1981,
		},
	},
	110203: {
		{
			Address: "通县",
			EndYear: 1981,
		},
	},
	110204: {
		{
			Address: "大兴县",
			EndYear: 1981,
		},
	},
	110205: {
		{
			Address: "房山县",
			EndYear: 1981,
		},
	},
	110206: {
		{
			Address: "平谷县",
			EndYear: 1981,
		},
	},
	110207: {
		{
			Address: "怀柔县",
			EndYear: 1981,
		},
	},
	110208: {
		{
			Address: "密云县",
			EndYear: 1981,
		},
	},
	110209: {
		{
			Address: "延庆县",
			EndYear: 1981,
		},
	},
	110221: {
		{
			Address:   "昌平县",
			StartYear: 1982,
			EndYear:   1998,
		},
	},
	110222: {
		{
			Address:   "顺义县",
			StartYear: 1982,
			EndYear:   1997,
		},
	},
	110223: {
		{
			Address:   "通县",
			StartYear: 1982,
			EndYear:   1996,
		},
	},
	110224: {
		{
			Address:   "大兴县",
			StartYear: 1982,
			EndYear:   2000,
		},
	},
	110225: {
		{
			Address:   "房山县",
			StartYear: 1982,
			EndYear:   1985,
		},
	},
	110226: {
		{
			Address:   "平谷县",
			StartYear: 1982,
			EndYear:   2000,
		},
	},
	110227: {
		{
			Address:   "怀柔县",
			StartYear: 1982,
			EndYear:   2000,
		},
	},
	110228: {
		{
			Address:   "密云县",
			StartYear: 1982,
			EndYear:   2014,
		},
	},
	110229: {
		{
			Address:   "延庆县",
			StartYear: 1982,
			EndYear:   2014,
		},
	},
	120000: {
		{
			Address: "天津市",
		},
	},
	120101: {
		{
			Address: "和平区",
		},
	},
	120102: {
		{
			Address: "河东区",
		},
	},
	120103: {
		{
			Address: "河西区",
		},
	},
	120104: {
		{
			Address: "南开区",
		},
	},
	120105: {
		{
			Address: "河北区",
		},
	},
	120106: {
		{
			Address: "红桥区",
		},
	},
	120107: {
		{
			Address: "塘沽区",
			EndYear: 2008,
		},
	},
	120108: {
		{
			Address: "汉沽区",
			EndYear: 2008,
		},
	},
	120109: {
		{
			Address: "大港区",
			EndYear: 2008,
		},
	},
	120110: {
		{
			Address: "东郊区",
			EndYear: 1991,
		},
		{
			Address:   "东丽区",
			StartYear: 1992,
		},
	},
	120111: {
		{
			Address: "西郊区",
			EndYear: 1991,
		},
		{
			Address:   "西青区",
			StartYear: 1992,
		},
	},
	120112: {
		{
			Address: "南郊区",
			EndYear: 1991,
		},
		{
			Address:   "津南区",
			StartYear: 1992,
		},
	},
	120113: {
		{
			Address: "北郊区",
			EndYear: 1991,
		},
		{
			Address:   "北辰区",
			StartYear: 1992,
		},
	},
	120114: {
		{
			Address:   "武清区",
			StartYear: 2000,
		},
	},
	120115: {
		{
			Address:   "宝坻区",
			StartYear: 2001,
		},
	},
	120116: {
		{
			Address:   "滨海新区",
			StartYear: 2009,
		},
	},
	120117: {
		{
			Address:   "宁河区",
			StartYear: 2015,
		},
	},
	120118: {
		{
			Address:   "静海区",
			StartYear: 2015,
		},
	},
	120119: {
		{
			Address:   "蓟州区",
			StartYear: 2016,
		},
	},
	120201: {
		{
			Address: "宁河县",
			EndYear: 1981,
		},
	},
	120202: {
		{
			Address: "武清县",
			EndYear: 1981,
		},
	},
	120203: {
		{
			Address: "静海县",
			EndYear: 1981,
		},
	},
	120204: {
		{
			Address: "宝坻县",
			EndYear: 1981,
		},
	},
	120205: {
		{
			Address: "蓟县",
			EndYear: 1981,
		},
	},
	120221: {
		{
			Address:   "宁河县",
			StartYear: 1982,
			EndYear:   2014,
		},
	},
	120222: {
		{
			Address:   "武清县",
			StartYear: 1982,
			EndYear:   1999,
		},
	},
	120223: {
		{
			Address:   "静海县",
			StartYear: 1982,
			EndYear:   2014,
		},
	},
	120224: {
		{
			Address:   "宝坻县",
			StartYear: 1982,
			EndYear:   2000,
		},
	},
	120225: {
		{
			Address:   "蓟县",
			StartYear: 1982,
			EndYear:   2015,
		},
	},
	130000: {
		{
			Address: "河北省",
		},
	},
	130100: {
		{
			Address: "石家庄市",
		},
	},
	130102: {
		{
			Address: "长安区",
		},
	},
	130103: {
		{
			Address: "桥东区",
			EndYear: 2014,
		},
	},
	130104: {
		{
			Address: "桥西区",
		},
	},
	130105: {
		{
			Address: "新华区",
		},
	},
	130106: {
		{
			Address: "郊区",
			EndYear: 2000,
		},
	},
	130107: {
		{
			Address: "井陉矿区",
		},
	},
	130108: {
		{
			Address:   "裕华区",
			StartYear: 2001,
		},
	},
	130109: {
		{
			Address:   "藁城区",
			StartYear: 2014,
		},
	},
	130110: {
		{
			Address:   "鹿泉区",
			StartYear: 2014,
		},
	},
	130111: {
		{
			Address:   "栾城区",
			StartYear: 2014,
		},
	},
	130121: {
		{
			Address:   "井陉县",
			StartYear: 1983,
		},
	},
	130122: {
		{
			Address:   "获鹿县",
			StartYear: 1983,
			EndYear:   1993,
		},
	},
	130123: {
		{
			Address:   "栾城县",
			StartYear: 1986,
			EndYear:   1992,
		},
		{
			Address:   "正定县",
			StartYear: 1993,
		},
	},
	130124: {
		{
			Address:   "正定县",
			StartYear: 1986,
			EndYear:   1992,
		},
		{
			Address:   "栾城县",
			StartYear: 1993,
			EndYear:   2013,
		},
	},
	130125: {
		{
			Address:   "行唐县",
			StartYear: 1993,
		},
	},
	130126: {
		{
			Address:   "灵寿县",
			StartYear: 1993,
		},
	},
	130127: {
		{
			Address:   "高邑县",
			StartYear: 1993,
		},
	},
	130128: {
		{
			Address:   "深泽县",
			StartYear: 1993,
		},
	},
	130129: {
		{
			Address:   "赞皇县",
			StartYear: 1993,
		},
	},
	130130: {
		{
			Address:   "无极县",
			StartYear: 1993,
		},
	},
	130131: {
		{
			Address:   "平山县",
			StartYear: 1993,
		},
	},
	130132: {
		{
			Address:   "元氏县",
			StartYear: 1993,
		},
	},
	130133: {
		{
			Address:   "赵县",
			StartYear: 1993,
		},
	},
	130181: {
		{
			Address:   "辛集市",
			StartYear: 1995,
		},
	},
	130182: {
		{
			Address:   "藁城市",
			StartYear: 1995,
			EndYear:   2013,
		},
	},
	130183: {
		{
			Address:   "晋州市",
			StartYear: 1995,
		},
	},
	130184: {
		{
			Address:   "新乐市",
			StartYear: 1995,
		},
	},
	130185: {
		{
			Address:   "鹿泉市",
			StartYear: 1995,
			EndYear:   2013,
		},
	},
	130200: {
		{
			Address: "唐山市",
		},
	},
	130202: {
		{
			Address: "路南区",
		},
	},
	130203: {
		{
			Address: "路北区",
		},
	},
	130204: {
		{
			Address: "东矿区",
			EndYear: 1994,
		},
		{
			Address:   "古冶区",
			StartYear: 1995,
		},
	},
	130205: {
		{
			Address: "开平区",
		},
	},
	130206: {
		{
			Address: "新区",
			EndYear: 2001,
		},
	},
	130207: {
		{
			Address:   "丰南区",
			StartYear: 2002,
		},
	},
	130208: {
		{
			Address:   "丰润区",
			StartYear: 2002,
		},
	},
	130209: {
		{
			Address:   "曹妃甸区",
			StartYear: 2012,
		},
	},
	130221: {
		{
			Address:   "丰润县",
			StartYear: 1983,
			EndYear:   2001,
		},
	},
	130222: {
		{
			Address:   "丰南县",
			StartYear: 1983,
			EndYear:   1993,
		},
	},
	130223: {
		{
			Address:   "滦县",
			StartYear: 1983,
			EndYear:   2017,
		},
	},
	130224: {
		{
			Address:   "滦南县",
			StartYear: 1983,
		},
	},
	130225: {
		{
			Address:   "乐亭县",
			StartYear: 1983,
		},
	},
	130226: {
		{
			Address:   "玉田县",
			StartYear: 1983,
			EndYear:   1983,
		},
		{
			Address:   "迁安县",
			StartYear: 1984,
			EndYear:   1995,
		},
	},
	130227: {
		{
			Address:   "唐海县",
			StartYear: 1983,
			EndYear:   1983,
		},
		{
			Address:   "迁西县",
			StartYear: 1984,
		},
	},
	130228: {
		{
			Address:   "迁安县",
			StartYear: 1983,
			EndYear:   1983,
		},
		{
			Address:   "遵化县",
			StartYear: 1984,
			EndYear:   1991,
		},
	},
	130229: {
		{
			Address:   "迁西县",
			StartYear: 1983,
			EndYear:   1983,
		},
		{
			Address:   "玉田县",
			StartYear: 1984,
		},
	},
	130230: {
		{
			Address:   "遵化县",
			StartYear: 1983,
			EndYear:   1983,
		},
		{
			Address:   "唐海县",
			StartYear: 1984,
			EndYear:   2011,
		},
	},
	130281: {
		{
			Address:   "遵化市",
			StartYear: 1995,
		},
	},
	130282: {
		{
			Address:   "丰南市",
			StartYear: 1995,
			EndYear:   2001,
		},
	},
	130283: {
		{
			Address:   "迁安市",
			StartYear: 1996,
		},
	},
	130284: {
		{
			Address:   "滦州市",
			StartYear: 2018,
		},
	},
	130300: {
		{
			Address:   "秦皇岛市",
			StartYear: 1983,
		},
	},
	130302: {
		{
			Address:   "海港区",
			StartYear: 1983,
		},
	},
	130303: {
		{
			Address:   "山海关区",
			StartYear: 1983,
		},
	},
	130304: {
		{
			Address:   "北戴河区",
			StartYear: 1983,
		},
	},
	130305: {
		{
			Address:   "郊区",
			StartYear: 1983,
			EndYear:   1983,
		},
	},
	130306: {
		{
			Address:   "抚宁区",
			StartYear: 2015,
		},
	},
	130321: {
		{
			Address:   "青龙县",
			StartYear: 1983,
			EndYear:   1985,
		},
		{
			Address:   "青龙满族自治县",
			StartYear: 1986,
		},
	},
	130322: {
		{
			Address:   "昌黎县",
			StartYear: 1983,
		},
	},
	130323: {
		{
			Address:   "抚宁县",
			StartYear: 1983,
			EndYear:   2014,
		},
	},
	130324: {
		{
			Address:   "卢龙县",
			StartYear: 1983,
		},
	},
	130400: {
		{
			Address:   "邯郸市",
			StartYear: 1983,
		},
	},
	130402: {
		{
			Address:   "邯山区",
			StartYear: 1983,
		},
	},
	130403: {
		{
			Address:   "丛台区",
			StartYear: 1983,
		},
	},
	130404: {
		{
			Address:   "复兴区",
			StartYear: 1983,
		},
	},
	130405: {
		{
			Address:   "郊区",
			StartYear: 1983,
			EndYear:   1985,
		},
	},
	130406: {
		{
			Address:   "峰峰矿区",
			StartYear: 1983,
		},
	},
	130407: {
		{
			Address:   "肥乡区",
			StartYear: 2016,
		},
	},
	130408: {
		{
			Address:   "永年区",
			StartYear: 2016,
		},
	},
	130421: {
		{
			Address:   "邯郸县",
			StartYear: 1983,
			EndYear:   2015,
		},
	},
	130422: {
		{
			Address:   "武安县",
			StartYear: 1986,
			EndYear:   1987,
		},
	},
	130423: {
		{
			Address:   "临漳县",
			StartYear: 1993,
		},
	},
	130424: {
		{
			Address:   "成安县",
			StartYear: 1993,
		},
	},
	130425: {
		{
			Address:   "大名县",
			StartYear: 1993,
		},
	},
	130426: {
		{
			Address:   "涉县",
			StartYear: 1993,
		},
	},
	130427: {
		{
			Address:   "磁县",
			StartYear: 1993,
		},
	},
	130428: {
		{
			Address:   "肥乡县",
			StartYear: 1993,
			EndYear:   2015,
		},
	},
	130429: {
		{
			Address:   "永年县",
			StartYear: 1993,
			EndYear:   2015,
		},
	},
	130430: {
		{
			Address:   "丘县",
			StartYear: 1993,
			EndYear:   1995,
		},
		{
			Address:   "邱县",
			StartYear: 1996,
		},
	},
	130431: {
		{
			Address:   "鸡泽县",
			StartYear: 1993,
		},
	},
	130432: {
		{
			Address:   "广平县",
			StartYear: 1993,
		},
	},
	130433: {
		{
			Address:   "馆陶县",
			StartYear: 1993,
		},
	},
	130434: {
		{
			Address:   "魏县",
			StartYear: 1993,
		},
	},
	130435: {
		{
			Address:   "曲周县",
			StartYear: 1993,
		},
	},
	130481: {
		{
			Address:   "武安市",
			StartYear: 1995,
		},
	},
	130500: {
		{
			Address:   "邢台市",
			StartYear: 1983,
		},
	},
	130502: {
		{
			Address:   "桥东区",
			StartYear: 1983,
			EndYear:   2019,
		},
		{
			Address:   "襄都区",
			StartYear: 2020,
		},
	},
	130503: {
		{
			Address:   "桥西区",
			StartYear: 1983,
			EndYear:   2019,
		},
		{
			Address:   "信都区",
			StartYear: 2020,
		},
	},
	130504: {
		{
			Address:   "郊区",
			StartYear: 1983,
			EndYear:   1987,
		},
	},
	130505: {
		{
			Address:   "任泽区",
			StartYear: 2020,
		},
	},
	130506: {
		{
			Address:   "南和区",
			StartYear: 2020,
		},
	},
	130521: {
		{
			Address:   "邢台县",
			StartYear: 1986,
			EndYear:   2019,
		},
	},
	130522: {
		{
			Address:   "临城县",
			StartYear: 1993,
		},
	},
	130523: {
		{
			Address:   "内丘县",
			StartYear: 1993,
		},
	},
	130524: {
		{
			Address:   "柏乡县",
			StartYear: 1993,
		},
	},
	130525: {
		{
			Address:   "隆尧县",
			StartYear: 1993,
		},
	},
	130526: {
		{
			Address:   "任县",
			StartYear: 1993,
			EndYear:   2019,
		},
	},
	130527: {
		{
			Address:   "南和县",
			StartYear: 1993,
			EndYear:   2019,
		},
	},
	130528: {
		{
			Address:   "宁晋县",
			StartYear: 1993,
		},
	},
	130529: {
		{
			Address:   "巨鹿县",
			StartYear: 1993,
		},
	},
	130530: {
		{
			Address:   "新河县",
			StartYear: 1993,
		},
	},
	130531: {
		{
			Address:   "广宗县",
			StartYear: 1993,
		},
	},
	130532: {
		{
			Address:   "平乡县",
			StartYear: 1993,
		},
	},
	130533: {
		{
			Address:   "威县",
			StartYear: 1993,
		},
	},
	130534: {
		{
			Address:   "清河县",
			StartYear: 1993,
		},
	},
	130535: {
		{
			Address:   "临西县",
			StartYear: 1993,
		},
	},
	130581: {
		{
			Address:   "南宫市",
			StartYear: 1995,
		},
	},
	130582: {
		{
			Address:   "沙河市",
			StartYear: 1995,
		},
	},
	130600: {
		{
			Address:   "保定市",
			StartYear: 1983,
		},
	},
	130602: {
		{
			Address:   "新市区",
			StartYear: 1983,
			EndYear:   2014,
		},
		{
			Address:   "竞秀区",
			StartYear: 2015,
		},
	},
	130603: {
		{
			Address:   "北市区",
			StartYear: 1983,
			EndYear:   2014,
		},
	},
	130604: {
		{
			Address:   "南市区",
			StartYear: 1983,
			EndYear:   2014,
		},
	},
	130605: {
		{
			Address:   "郊区",
			StartYear: 1983,
			EndYear:   1986,
		},
	},
	130606: {
		{
			Address:   "莲池区",
			StartYear: 2015,
		},
	},
	130607: {
		{
			Address:   "满城区",
			StartYear: 2015,
		},
	},
	130608: {
		{
			Address:   "清苑区",
			StartYear: 2015,
		},
	},
	130609: {
		{
			Address:   "徐水区",
			StartYear: 2015,
		},
	},
	130621: {
		{
			Address:   "满城县",
			StartYear: 1983,
			EndYear:   2014,
		},
	},
	130622: {
		{
			Address:   "清苑县",
			StartYear: 1986,
			EndYear:   2014,
		},
	},
	130623: {
		{
			Address:   "涞水县",
			StartYear: 1994,
		},
	},
	130624: {
		{
			Address:   "阜平县",
			StartYear: 1994,
		},
	},
	130625: {
		{
			Address:   "徐水县",
			StartYear: 1994,
			EndYear:   2014,
		},
	},
	130626: {
		{
			Address:   "定兴县",
			StartYear: 1994,
		},
	},
	130627: {
		{
			Address:   "唐县",
			StartYear: 1994,
		},
	},
	130628: {
		{
			Address:   "高阳县",
			StartYear: 1994,
		},
	},
	130629: {
		{
			Address:   "容城县",
			StartYear: 1994,
		},
	},
	130630: {
		{
			Address:   "涞源县",
			StartYear: 1994,
		},
	},
	130631: {
		{
			Address:   "望都县",
			StartYear: 1994,
		},
	},
	130632: {
		{
			Address:   "安新县",
			StartYear: 1994,
		},
	},
	130633: {
		{
			Address:   "易县",
			StartYear: 1994,
		},
	},
	130634: {
		{
			Address:   "曲阳县",
			StartYear: 1994,
		},
	},
	130635: {
		{
			Address:   "蠡县",
			StartYear: 1994,
		},
	},
	130636: {
		{
			Address:   "顺平县",
			StartYear: 1994,
		},
	},
	130637: {
		{
			Address:   "博野县",
			StartYear: 1994,
		},
	},
	130638: {
		{
			Address:   "雄县",
			StartYear: 1994,
		},
	},
	130681: {
		{
			Address:   "涿州市",
			StartYear: 1995,
		},
	},
	130682: {
		{
			Address:   "定州市",
			StartYear: 1995,
		},
	},
	130683: {
		{
			Address:   "安国市",
			StartYear: 1995,
		},
	},
	130684: {
		{
			Address:   "高碑店市",
			StartYear: 1995,
		},
	},
	130700: {
		{
			Address:   "张家口市",
			StartYear: 1983,
		},
	},
	130702: {
		{
			Address:   "桥东区",
			StartYear: 1983,
		},
	},
	130703: {
		{
			Address:   "桥西区",
			StartYear: 1983,
		},
	},
	130704: {
		{
			Address:   "茶坊区",
			StartYear: 1983,
			EndYear:   1988,
		},
	},
	130705: {
		{
			Address:   "宣化区",
			StartYear: 1983,
		},
	},
	130706: {
		{
			Address:   "下花园区",
			StartYear: 1983,
		},
	},
	130707: {
		{
			Address:   "庞家堡区",
			StartYear: 1983,
			EndYear:   1988,
		},
	},
	130708: {
		{
			Address:   "万全区",
			StartYear: 2016,
		},
	},
	130709: {
		{
			Address:   "崇礼区",
			StartYear: 2016,
		},
	},
	130721: {
		{
			Address:   "宣化县",
			StartYear: 1983,
			EndYear:   2015,
		},
	},
	130722: {
		{
			Address:   "张北县",
			StartYear: 1993,
		},
	},
	130723: {
		{
			Address:   "康保县",
			StartYear: 1993,
		},
	},
	130724: {
		{
			Address:   "沽源县",
			StartYear: 1993,
		},
	},
	130725: {
		{
			Address:   "尚义县",
			StartYear: 1993,
		},
	},
	130726: {
		{
			Address:   "蔚县",
			StartYear: 1993,
		},
	},
	130727: {
		{
			Address:   "阳原县",
			StartYear: 1993,
		},
	},
	130728: {
		{
			Address:   "怀安县",
			StartYear: 1993,
		},
	},
	130729: {
		{
			Address:   "万全县",
			StartYear: 1993,
			EndYear:   2015,
		},
	},
	130730: {
		{
			Address:   "怀来县",
			StartYear: 1993,
		},
	},
	130731: {
		{
			Address:   "涿鹿县",
			StartYear: 1993,
		},
	},
	130732: {
		{
			Address:   "赤城县",
			StartYear: 1993,
		},
	},
	130733: {
		{
			Address:   "崇礼县",
			StartYear: 1993,
			EndYear:   2015,
		},
	},
	130800: {
		{
			Address:   "承德市",
			StartYear: 1983,
		},
	},
	130802: {
		{
			Address:   "双桥区",
			StartYear: 1983,
		},
	},
	130803: {
		{
			Address:   "双滦区",
			StartYear: 1983,
		},
	},
	130804: {
		{
			Address:   "鹰手营子矿区",
			StartYear: 1983,
		},
	},
	130821: {
		{
			Address:   "承德县",
			StartYear: 1983,
		},
	},
	130822: {
		{
			Address:   "兴隆县",
			StartYear: 1993,
		},
	},
	130823: {
		{
			Address:   "平泉县",
			StartYear: 1993,
			EndYear:   2016,
		},
	},
	130824: {
		{
			Address:   "滦平县",
			StartYear: 1993,
		},
	},
	130825: {
		{
			Address:   "隆化县",
			StartYear: 1993,
		},
	},
	130826: {
		{
			Address:   "丰宁满族自治县",
			StartYear: 1993,
		},
	},
	130827: {
		{
			Address:   "宽城满族自治县",
			StartYear: 1993,
		},
	},
	130828: {
		{
			Address:   "围场满族蒙古族自治县",
			StartYear: 1993,
		},
	},
	130881: {
		{
			Address:   "平泉市",
			StartYear: 2017,
		},
	},
	130900: {
		{
			Address:   "沧州市",
			StartYear: 1983,
		},
	},
	130902: {
		{
			Address:   "新华区",
			StartYear: 1983,
		},
	},
	130903: {
		{
			Address:   "运河区",
			StartYear: 1983,
		},
	},
	130904: {
		{
			Address:   "郊区",
			StartYear: 1983,
			EndYear:   1996,
		},
	},
	130921: {
		{
			Address:   "沧县",
			StartYear: 1983,
		},
	},
	130922: {
		{
			Address:   "青县",
			StartYear: 1986,
		},
	},
	130923: {
		{
			Address:   "东光县",
			StartYear: 1993,
		},
	},
	130924: {
		{
			Address:   "海兴县",
			StartYear: 1993,
		},
	},
	130925: {
		{
			Address:   "盐山县",
			StartYear: 1993,
		},
	},
	130926: {
		{
			Address:   "肃宁县",
			StartYear: 1993,
		},
	},
	130927: {
		{
			Address:   "南皮县",
			StartYear: 1993,
		},
	},
	130928: {
		{
			Address:   "吴桥县",
			StartYear: 1993,
		},
	},
	130929: {
		{
			Address:   "献县",
			StartYear: 1993,
		},
	},
	130930: {
		{
			Address:   "孟村回族自治县",
			StartYear: 1993,
		},
	},
	130981: {
		{
			Address:   "泊头市",
			StartYear: 1995,
		},
	},
	130982: {
		{
			Address:   "任丘市",
			StartYear: 1995,
		},
	},
	130983: {
		{
			Address:   "黄骅市",
			StartYear: 1995,
		},
	},
	130984: {
		{
			Address:   "河间市",
			StartYear: 1995,
		},
	},
	131000: {
		{
			Address:   "廊坊市",
			StartYear: 1988,
		},
	},
	131002: {
		{
			Address:   "安次区",
			StartYear: 1988,
		},
	},
	131003: {
		{
			Address:   "广阳区",
			StartYear: 2000,
		},
	},
	131021: {
		{
			Address:   "三河县",
			StartYear: 1988,
			EndYear:   1992,
		},
	},
	131022: {
		{
			Address:   "固安县",
			StartYear: 1988,
		},
	},
	131023: {
		{
			Address:   "永清县",
			StartYear: 1988,
		},
	},
	131024: {
		{
			Address:   "香河县",
			StartYear: 1988,
		},
	},
	131025: {
		{
			Address:   "大城县",
			StartYear: 1988,
		},
	},
	131026: {
		{
			Address:   "文安县",
			StartYear: 1988,
		},
	},
	131027: {
		{
			Address:   "霸县",
			StartYear: 1988,
			EndYear:   1989,
		},
	},
	131028: {
		{
			Address:   "大厂回族自治县",
			StartYear: 1988,
		},
	},
	131081: {
		{
			Address:   "霸州市",
			StartYear: 1995,
		},
	},
	131082: {
		{
			Address:   "三河市",
			StartYear: 1995,
		},
	},
	131100: {
		{
			Address:   "衡水市",
			StartYear: 1996,
		},
	},
	131102: {
		{
			Address:   "桃城区",
			StartYear: 1996,
		},
	},
	131103: {
		{
			Address:   "冀州区",
			StartYear: 2016,
		},
	},
	131121: {
		{
			Address:   "枣强县",
			StartYear: 1996,
		},
	},
	131122: {
		{
			Address:   "武邑县",
			StartYear: 1996,
		},
	},
	131123: {
		{
			Address:   "武强县",
			StartYear: 1996,
		},
	},
	131124: {
		{
			Address:   "饶阳县",
			StartYear: 1996,
		},
	},
	131125: {
		{
			Address:   "安平县",
			StartYear: 1996,
		},
	},
	131126: {
		{
			Address:   "故城县",
			StartYear: 1996,
		},
	},
	131127: {
		{
			Address:   "景县",
			StartYear: 1996,
		},
	},
	131128: {
		{
			Address:   "阜城县",
			StartYear: 1996,
		},
	},
	131181: {
		{
			Address:   "冀州市",
			StartYear: 1996,
			EndYear:   2015,
		},
	},
	131182: {
		{
			Address:   "深州市",
			StartYear: 1996,
		},
	},
	132100: {
		{
			Address: "邯郸地区",
			EndYear: 1992,
		},
	},
	132101: {
		{
			Address: "邯郸市",
			EndYear: 1982,
		},
	},
	132102: {
		{
			Address: "邯山区",
			EndYear: 1982,
		},
	},
	132103: {
		{
			Address: "丛台区",
			EndYear: 1982,
		},
	},
	132104: {
		{
			Address: "复兴区",
			EndYear: 1982,
		},
	},
	132105: {
		{
			Address: "郊区",
			EndYear: 1982,
		},
	},
	132106: {
		{
			Address: "峰峰矿区",
			EndYear: 1982,
		},
	},
	132121: {
		{
			Address: "大名县",
			EndYear: 1992,
		},
	},
	132122: {
		{
			Address: "魏县",
			EndYear: 1992,
		},
	},
	132123: {
		{
			Address: "曲周县",
			EndYear: 1992,
		},
	},
	132124: {
		{
			Address: "丘县",
			EndYear: 1992,
		},
	},
	132125: {
		{
			Address: "鸡泽县",
			EndYear: 1992,
		},
	},
	132126: {
		{
			Address: "肥乡县",
			EndYear: 1992,
		},
	},
	132127: {
		{
			Address: "广平县",
			EndYear: 1992,
		},
	},
	132128: {
		{
			Address: "成安县",
			EndYear: 1992,
		},
	},
	132129: {
		{
			Address: "临漳县",
			EndYear: 1992,
		},
	},
	132130: {
		{
			Address: "磁县",
			EndYear: 1992,
		},
	},
	132131: {
		{
			Address: "武安县",
			EndYear: 1985,
		},
	},
	132132: {
		{
			Address: "涉县",
			EndYear: 1992,
		},
	},
	132133: {
		{
			Address: "永年县",
			EndYear: 1992,
		},
	},
	132134: {
		{
			Address: "邯郸县",
			EndYear: 1982,
		},
	},
	132135: {
		{
			Address: "馆陶县",
			EndYear: 1992,
		},
	},
	132200: {
		{
			Address: "邢台地区",
			EndYear: 1992,
		},
	},
	132201: {
		{
			Address: "邢台市",
			EndYear: 1982,
		},
		{
			Address:   "南宫市",
			StartYear: 1986,
			EndYear:   1992,
		},
	},
	132202: {
		{
			Address: "桥东区",
			EndYear: 1982,
		},
		{
			Address:   "沙河市",
			StartYear: 1987,
			EndYear:   1992,
		},
	},
	132203: {
		{
			Address: "桥西区",
			EndYear: 1982,
		},
	},
	132204: {
		{
			Address: "郊区",
			EndYear: 1982,
		},
	},
	132221: {
		{
			Address: "邢台县",
			EndYear: 1985,
		},
	},
	132222: {
		{
			Address: "沙河县",
			EndYear: 1986,
		},
	},
	132223: {
		{
			Address: "临城县",
			EndYear: 1992,
		},
	},
	132224: {
		{
			Address: "内丘县",
			EndYear: 1992,
		},
	},
	132225: {
		{
			Address: "柏乡县",
			EndYear: 1992,
		},
	},
	132226: {
		{
			Address: "隆尧县",
			EndYear: 1992,
		},
	},
	132227: {
		{
			Address: "任县",
			EndYear: 1992,
		},
	},
	132228: {
		{
			Address: "南和县",
			EndYear: 1992,
		},
	},
	132229: {
		{
			Address: "宁晋县",
			EndYear: 1992,
		},
	},
	132230: {
		{
			Address: "南宫县",
			EndYear: 1985,
		},
	},
	132231: {
		{
			Address: "巨鹿县",
			EndYear: 1992,
		},
	},
	132232: {
		{
			Address: "新河县",
			EndYear: 1992,
		},
	},
	132233: {
		{
			Address: "广宗县",
			EndYear: 1992,
		},
	},
	132234: {
		{
			Address: "平乡县",
			EndYear: 1992,
		},
	},
	132235: {
		{
			Address: "威县",
			EndYear: 1992,
		},
	},
	132236: {
		{
			Address: "清河县",
			EndYear: 1992,
		},
	},
	132237: {
		{
			Address: "临西县",
			EndYear: 1992,
		},
	},
	132300: {
		{
			Address: "石家庄地区",
			EndYear: 1992,
		},
	},
	132301: {
		{
			Address:   "辛集市",
			StartYear: 1986,
			EndYear:   1992,
		},
	},
	132302: {
		{
			Address:   "藁城市",
			StartYear: 1989,
			EndYear:   1992,
		},
	},
	132303: {
		{
			Address:   "晋州市",
			StartYear: 1991,
			EndYear:   1992,
		},
	},
	132304: {
		{
			Address:   "新乐市",
			StartYear: 1992,
			EndYear:   1992,
		},
	},
	132321: {
		{
			Address: "束鹿县",
			EndYear: 1985,
		},
	},
	132322: {
		{
			Address: "晋县",
			EndYear: 1990,
		},
	},
	132323: {
		{
			Address: "深泽县",
			EndYear: 1992,
		},
	},
	132324: {
		{
			Address: "无极县",
			EndYear: 1992,
		},
	},
	132325: {
		{
			Address: "藁城县",
			EndYear: 1988,
		},
	},
	132326: {
		{
			Address: "赵县",
			EndYear: 1992,
		},
	},
	132327: {
		{
			Address: "栾城县",
			EndYear: 1985,
		},
	},
	132328: {
		{
			Address: "正定县",
			EndYear: 1985,
		},
	},
	132329: {
		{
			Address: "新乐县",
			EndYear: 1991,
		},
	},
	132330: {
		{
			Address: "高邑县",
			EndYear: 1992,
		},
	},
	132331: {
		{
			Address: "元氏县",
			EndYear: 1992,
		},
	},
	132332: {
		{
			Address: "赞皇县",
			EndYear: 1992,
		},
	},
	132333: {
		{
			Address: "井陉县",
			EndYear: 1982,
		},
	},
	132334: {
		{
			Address: "获鹿县",
			EndYear: 1982,
		},
	},
	132335: {
		{
			Address: "平山县",
			EndYear: 1992,
		},
	},
	132336: {
		{
			Address: "灵寿县",
			EndYear: 1992,
		},
	},
	132337: {
		{
			Address: "行唐县",
			EndYear: 1992,
		},
	},
	132400: {
		{
			Address: "保定地区",
			EndYear: 1993,
		},
	},
	132401: {
		{
			Address: "保定市",
			EndYear: 1982,
		},
		{
			Address:   "定州市",
			StartYear: 1986,
			EndYear:   1993,
		},
	},
	132402: {
		{
			Address: "新市区",
			EndYear: 1982,
		},
		{
			Address:   "涿州市",
			StartYear: 1986,
			EndYear:   1993,
		},
	},
	132403: {
		{
			Address: "北市区",
			EndYear: 1982,
		},
		{
			Address:   "安国市",
			StartYear: 1991,
			EndYear:   1993,
		},
	},
	132404: {
		{
			Address: "南市区",
			EndYear: 1982,
		},
		{
			Address:   "高碑店市",
			StartYear: 1993,
			EndYear:   1993,
		},
	},
	132405: {
		{
			Address: "郊区",
			EndYear: 1982,
		},
	},
	132421: {
		{
			Address: "易县",
			EndYear: 1993,
		},
	},
	132422: {
		{
			Address: "满城县",
			EndYear: 1982,
		},
	},
	132423: {
		{
			Address: "徐水县",
			EndYear: 1993,
		},
	},
	132424: {
		{
			Address: "涞源县",
			EndYear: 1993,
		},
	},
	132425: {
		{
			Address: "定兴县",
			EndYear: 1993,
		},
	},
	132426: {
		{
			Address: "完县",
			EndYear: 1992,
		},
		{
			Address:   "顺平县",
			StartYear: 1993,
			EndYear:   1993,
		},
	},
	132427: {
		{
			Address: "唐县",
			EndYear: 1993,
		},
	},
	132428: {
		{
			Address: "望都县",
			EndYear: 1993,
		},
	},
	132429: {
		{
			Address: "涞水县",
			EndYear: 1993,
		},
	},
	132430: {
		{
			Address: "涿县",
			EndYear: 1985,
		},
	},
	132431: {
		{
			Address: "清苑县",
			EndYear: 1985,
		},
	},
	132432: {
		{
			Address: "高阳县",
			EndYear: 1993,
		},
	},
	132433: {
		{
			Address: "安新县",
			EndYear: 1993,
		},
	},
	132434: {
		{
			Address: "雄县",
			EndYear: 1993,
		},
	},
	132435: {
		{
			Address: "容城县",
			EndYear: 1993,
		},
	},
	132436: {
		{
			Address: "新城县",
			EndYear: 1992,
		},
	},
	132437: {
		{
			Address: "曲阳县",
			EndYear: 1993,
		},
	},
	132438: {
		{
			Address: "阜平县",
			EndYear: 1993,
		},
	},
	132439: {
		{
			Address: "定县",
			EndYear: 1985,
		},
	},
	132440: {
		{
			Address: "安国县",
			EndYear: 1990,
		},
	},
	132441: {
		{
			Address: "博野县",
			EndYear: 1993,
		},
	},
	132442: {
		{
			Address: "蠡县",
			EndYear: 1993,
		},
	},
	132500: {
		{
			Address: "张家口地区",
			EndYear: 1992,
		},
	},
	132501: {
		{
			Address: "张家口市",
			EndYear: 1982,
		},
	},
	132502: {
		{
			Address: "桥东区",
			EndYear: 1982,
		},
	},
	132503: {
		{
			Address: "桥西区",
			EndYear: 1982,
		},
	},
	132504: {
		{
			Address: "茶坊区",
			EndYear: 1982,
		},
	},
	132505: {
		{
			Address: "宣化区",
			EndYear: 1982,
		},
	},
	132506: {
		{
			Address: "下花园区",
			EndYear: 1982,
		},
	},
	132507: {
		{
			Address: "庞家堡区",
			EndYear: 1982,
		},
	},
	132521: {
		{
			Address: "张北县",
			EndYear: 1992,
		},
	},
	132522: {
		{
			Address: "康保县",
			EndYear: 1992,
		},
	},
	132523: {
		{
			Address: "沽源县",
			EndYear: 1992,
		},
	},
	132524: {
		{
			Address: "尚义县",
			EndYear: 1992,
		},
	},
	132525: {
		{
			Address: "蔚县",
			EndYear: 1992,
		},
	},
	132526: {
		{
			Address: "阳原县",
			EndYear: 1992,
		},
	},
	132527: {
		{
			Address: "怀安县",
			EndYear: 1992,
		},
	},
	132528: {
		{
			Address: "万全县",
			EndYear: 1992,
		},
	},
	132529: {
		{
			Address: "怀来县",
			EndYear: 1992,
		},
	},
	132530: {
		{
			Address: "涿鹿县",
			EndYear: 1992,
		},
	},
	132531: {
		{
			Address: "宣化县",
			EndYear: 1982,
		},
	},
	132532: {
		{
			Address: "赤城县",
			EndYear: 1992,
		},
	},
	132533: {
		{
			Address: "崇礼县",
			EndYear: 1992,
		},
	},
	132600: {
		{
			Address: "承德地区",
			EndYear: 1992,
		},
	},
	132601: {
		{
			Address: "承德市",
			EndYear: 1982,
		},
	},
	132602: {
		{
			Address: "双桥区",
			EndYear: 1982,
		},
	},
	132603: {
		{
			Address: "双滦区",
			EndYear: 1982,
		},
	},
	132604: {
		{
			Address: "鹰手营子矿区",
			EndYear: 1982,
		},
	},
	132621: {
		{
			Address: "青龙县",
			EndYear: 1982,
		},
	},
	132622: {
		{
			Address: "宽城县",
			EndYear: 1988,
		},
		{
			Address:   "宽城满族自治县",
			StartYear: 1989,
			EndYear:   1992,
		},
	},
	132623: {
		{
			Address: "兴隆县",
			EndYear: 1992,
		},
	},
	132624: {
		{
			Address: "平泉县",
			EndYear: 1992,
		},
	},
	132625: {
		{
			Address: "承德县",
			EndYear: 1982,
		},
	},
	132626: {
		{
			Address: "滦平县",
			EndYear: 1992,
		},
	},
	132627: {
		{
			Address: "丰宁县",
			EndYear: 1985,
		},
		{
			Address:   "丰宁满族自治县",
			StartYear: 1986,
			EndYear:   1992,
		},
	},
	132628: {
		{
			Address: "隆化县",
			EndYear: 1992,
		},
	},
	132629: {
		{
			Address: "围场县",
			EndYear: 1988,
		},
		{
			Address:   "围场满族蒙古族自治县",
			StartYear: 1989,
			EndYear:   1992,
		},
	},
	132700: {
		{
			Address: "唐山地区",
			EndYear: 1982,
		},
	},
	132701: {
		{
			Address: "秦皇岛市",
			EndYear: 1982,
		},
	},
	132702: {
		{
			Address: "海港区",
			EndYear: 1982,
		},
	},
	132703: {
		{
			Address: "山海关区",
			EndYear: 1982,
		},
	},
	132704: {
		{
			Address: "北戴河区",
			EndYear: 1982,
		},
	},
	132705: {
		{
			Address: "郊区",
			EndYear: 1982,
		},
	},
	132721: {
		{
			Address: "丰润县",
			EndYear: 1982,
		},
	},
	132722: {
		{
			Address: "丰南县",
			EndYear: 1982,
		},
	},
	132723: {
		{
			Address: "滦县",
			EndYear: 1982,
		},
	},
	132724: {
		{
			Address: "滦南县",
			EndYear: 1982,
		},
	},
	132725: {
		{
			Address: "乐亭县",
			EndYear: 1982,
		},
	},
	132726: {
		{
			Address: "昌黎县",
			EndYear: 1982,
		},
	},
	132727: {
		{
			Address: "抚宁县",
			EndYear: 1982,
		},
	},
	132728: {
		{
			Address: "卢龙县",
			EndYear: 1982,
		},
	},
	132729: {
		{
			Address: "迁安县",
			EndYear: 1982,
		},
	},
	132730: {
		{
			Address: "迁西县",
			EndYear: 1982,
		},
	},
	132731: {
		{
			Address: "遵化县",
			EndYear: 1982,
		},
	},
	132732: {
		{
			Address: "玉田县",
			EndYear: 1982,
		},
	},
	132733: {
		{
			Address:   "唐海县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	132800: {
		{
			Address: "廊坊地区",
			EndYear: 1987,
		},
	},
	132801: {
		{
			Address:   "廊坊市",
			StartYear: 1981,
			EndYear:   1987,
		},
	},
	132821: {
		{
			Address: "三河县",
			EndYear: 1987,
		},
	},
	132822: {
		{
			Address: "大厂回族自治县",
			EndYear: 1987,
		},
	},
	132823: {
		{
			Address: "香河县",
			EndYear: 1987,
		},
	},
	132824: {
		{
			Address: "安次县",
			EndYear: 1982,
		},
	},
	132825: {
		{
			Address: "永清县",
			EndYear: 1987,
		},
	},
	132826: {
		{
			Address: "固安县",
			EndYear: 1987,
		},
	},
	132827: {
		{
			Address: "霸县",
			EndYear: 1987,
		},
	},
	132828: {
		{
			Address: "文安县",
			EndYear: 1987,
		},
	},
	132829: {
		{
			Address: "大城县",
			EndYear: 1987,
		},
	},
	132900: {
		{
			Address: "沧州地区",
			EndYear: 1992,
		},
	},
	132901: {
		{
			Address: "沧州市",
			EndYear: 1982,
		},
	},
	132902: {
		{
			Address: "新华区",
			EndYear: 1982,
		},
		{
			Address:   "泊头市",
			StartYear: 1984,
			EndYear:   1992,
		},
	},
	132903: {
		{
			Address: "运河区",
			EndYear: 1982,
		},
		{
			Address:   "任丘市",
			StartYear: 1986,
			EndYear:   1992,
		},
	},
	132904: {
		{
			Address: "郊区",
			EndYear: 1982,
		},
		{
			Address:   "黄骅市",
			StartYear: 1989,
			EndYear:   1992,
		},
	},
	132905: {
		{
			Address:   "泊头市",
			StartYear: 1982,
			EndYear:   1983,
		},
		{
			Address:   "河间市",
			StartYear: 1990,
			EndYear:   1992,
		},
	},
	132921: {
		{
			Address: "沧县",
			EndYear: 1982,
		},
	},
	132922: {
		{
			Address: "河间县",
			EndYear: 1989,
		},
	},
	132923: {
		{
			Address: "肃宁县",
			EndYear: 1992,
		},
	},
	132924: {
		{
			Address: "献县",
			EndYear: 1992,
		},
	},
	132925: {
		{
			Address: "交河县",
			EndYear: 1982,
		},
	},
	132926: {
		{
			Address: "吴桥县",
			EndYear: 1992,
		},
	},
	132927: {
		{
			Address: "东光县",
			EndYear: 1992,
		},
	},
	132928: {
		{
			Address: "南皮县",
			EndYear: 1992,
		},
	},
	132929: {
		{
			Address: "盐山县",
			EndYear: 1992,
		},
	},
	132930: {
		{
			Address: "黄骅县",
			EndYear: 1988,
		},
	},
	132931: {
		{
			Address: "孟村回族自治县",
			EndYear: 1992,
		},
	},
	132932: {
		{
			Address: "青县",
			EndYear: 1985,
		},
	},
	132933: {
		{
			Address: "任丘县",
			EndYear: 1985,
		},
	},
	132934: {
		{
			Address: "海兴县",
			EndYear: 1992,
		},
	},
	133000: {
		{
			Address: "衡水地区",
			EndYear: 1995,
		},
	},
	133001: {
		{
			Address:   "衡水市",
			StartYear: 1982,
			EndYear:   1995,
		},
	},
	133002: {
		{
			Address:   "冀州市",
			StartYear: 1993,
			EndYear:   1995,
		},
	},
	133003: {
		{
			Address:   "深州市",
			StartYear: 1994,
			EndYear:   1995,
		},
	},
	133021: {
		{
			Address: "衡水县",
			EndYear: 1982,
		},
	},
	133022: {
		{
			Address: "冀县",
			EndYear: 1992,
		},
	},
	133023: {
		{
			Address: "枣强县",
			EndYear: 1995,
		},
	},
	133024: {
		{
			Address: "武邑县",
			EndYear: 1995,
		},
	},
	133025: {
		{
			Address: "深县",
			EndYear: 1993,
		},
	},
	133026: {
		{
			Address: "武强县",
			EndYear: 1995,
		},
	},
	133027: {
		{
			Address: "饶阳县",
			EndYear: 1995,
		},
	},
	133028: {
		{
			Address: "安平县",
			EndYear: 1995,
		},
	},
	133029: {
		{
			Address: "故城县",
			EndYear: 1995,
		},
	},
	133030: {
		{
			Address: "景县",
			EndYear: 1995,
		},
	},
	133031: {
		{
			Address: "阜城县",
			EndYear: 1995,
		},
	},
	139001: {
		{
			Address:   "武安市",
			StartYear: 1988,
			EndYear:   1994,
		},
	},
	139002: {
		{
			Address:   "霸州市",
			StartYear: 1990,
			EndYear:   1994,
		},
	},
	139003: {
		{
			Address:   "遵化市",
			StartYear: 1992,
			EndYear:   1994,
		},
	},
	139004: {
		{
			Address:   "辛集市",
			StartYear: 1993,
			EndYear:   1994,
		},
	},
	139005: {
		{
			Address:   "藁城市",
			StartYear: 1993,
			EndYear:   1994,
		},
	},
	139006: {
		{
			Address:   "晋州市",
			StartYear: 1993,
			EndYear:   1994,
		},
	},
	139007: {
		{
			Address:   "新乐市",
			StartYear: 1993,
			EndYear:   1994,
		},
	},
	139008: {
		{
			Address:   "泊头市",
			StartYear: 1993,
			EndYear:   1994,
		},
	},
	139009: {
		{
			Address:   "任丘市",
			StartYear: 1993,
			EndYear:   1994,
		},
	},
	139010: {
		{
			Address:   "黄骅市",
			StartYear: 1993,
			EndYear:   1994,
		},
	},
	139011: {
		{
			Address:   "河间市",
			StartYear: 1993,
			EndYear:   1994,
		},
	},
	139012: {
		{
			Address:   "三河县",
			StartYear: 1993,
			EndYear:   1993,
		},
		{
			Address:   "三河市",
			StartYear: 1994,
			EndYear:   1994,
		},
	},
	139013: {
		{
			Address:   "南宫市",
			StartYear: 1993,
			EndYear:   1994,
		},
	},
	139014: {
		{
			Address:   "沙河市",
			StartYear: 1993,
			EndYear:   1994,
		},
	},
	139015: {
		{
			Address:   "定州市",
			StartYear: 1994,
			EndYear:   1994,
		},
	},
	139016: {
		{
			Address:   "涿州市",
			StartYear: 1994,
			EndYear:   1994,
		},
	},
	139017: {
		{
			Address:   "安国市",
			StartYear: 1994,
			EndYear:   1994,
		},
	},
	139018: {
		{
			Address:   "高碑店市",
			StartYear: 1994,
			EndYear:   1994,
		},
	},
	139019: {
		{
			Address:   "鹿泉市",
			StartYear: 1994,
			EndYear:   1994,
		},
	},
	139020: {
		{
			Address:   "丰南市",
			StartYear: 1994,
			EndYear:   1994,
		},
	},
	140000: {
		{
			Address: "山西省",
		},
	},
	140100: {
		{
			Address: "太原市",
		},
	},
	140102: {
		{
			Address: "南城区",
			EndYear: 1996,
		},
	},
	140103: {
		{
			Address: "北城区",
			EndYear: 1996,
		},
	},
	140104: {
		{
			Address: "河西区",
			EndYear: 1996,
		},
	},
	140105: {
		{
			Address:   "小店区",
			StartYear: 1997,
		},
	},
	140106: {
		{
			Address:   "迎泽区",
			StartYear: 1997,
		},
	},
	140107: {
		{
			Address:   "杏花岭区",
			StartYear: 1997,
		},
	},
	140108: {
		{
			Address:   "尖草坪区",
			StartYear: 1997,
		},
	},
	140109: {
		{
			Address:   "万柏林区",
			StartYear: 1997,
		},
	},
	140110: {
		{
			Address:   "晋源区",
			StartYear: 1997,
		},
	},
	140111: {
		{
			Address: "古交工矿区",
			EndYear: 1987,
		},
	},
	140112: {
		{
			Address: "南郊区",
			EndYear: 1996,
		},
	},
	140113: {
		{
			Address: "北郊区",
			EndYear: 1996,
		},
	},
	140121: {
		{
			Address: "清徐县",
		},
	},
	140122: {
		{
			Address: "阳曲县",
		},
	},
	140123: {
		{
			Address: "娄烦县",
		},
	},
	140181: {
		{
			Address:   "古交市",
			StartYear: 1995,
		},
	},
	140200: {
		{
			Address: "大同市",
		},
	},
	140202: {
		{
			Address: "城区",
			EndYear: 2017,
		},
	},
	140203: {
		{
			Address: "矿区",
			EndYear: 2017,
		},
	},
	140211: {
		{
			Address: "南郊区",
			EndYear: 2017,
		},
	},
	140212: {
		{
			Address: "新荣区",
		},
	},
	140213: {
		{
			Address:   "平城区",
			StartYear: 2018,
		},
	},
	140214: {
		{
			Address:   "云冈区",
			StartYear: 2018,
		},
	},
	140215: {
		{
			Address:   "云州区",
			StartYear: 2018,
		},
	},
	140221: {
		{
			Address:   "阳高县",
			StartYear: 1993,
		},
	},
	140222: {
		{
			Address:   "天镇县",
			StartYear: 1993,
		},
	},
	140223: {
		{
			Address:   "广灵县",
			StartYear: 1993,
		},
	},
	140224: {
		{
			Address:   "灵丘县",
			StartYear: 1993,
		},
	},
	140225: {
		{
			Address:   "浑源县",
			StartYear: 1993,
		},
	},
	140226: {
		{
			Address:   "左云县",
			StartYear: 1993,
		},
	},
	140227: {
		{
			Address:   "大同县",
			StartYear: 1993,
			EndYear:   2017,
		},
	},
	140300: {
		{
			Address: "阳泉市",
		},
	},
	140302: {
		{
			Address: "城区",
		},
	},
	140303: {
		{
			Address: "矿区",
		},
	},
	140311: {
		{
			Address: "郊区",
		},
	},
	140321: {
		{
			Address:   "平定县",
			StartYear: 1983,
		},
	},
	140322: {
		{
			Address:   "盂县",
			StartYear: 1983,
		},
	},
	140400: {
		{
			Address: "长治市",
		},
	},
	140402: {
		{
			Address: "城区",
			EndYear: 2017,
		},
	},
	140403: {
		{
			Address:   "潞州区",
			StartYear: 2018,
		},
	},
	140404: {
		{
			Address:   "上党区",
			StartYear: 2018,
		},
	},
	140405: {
		{
			Address:   "屯留区",
			StartYear: 2018,
		},
	},
	140406: {
		{
			Address:   "潞城区",
			StartYear: 2018,
		},
	},
	140411: {
		{
			Address: "郊区",
			EndYear: 2017,
		},
	},
	140421: {
		{
			Address:   "长治县",
			StartYear: 1983,
			EndYear:   2017,
		},
	},
	140422: {
		{
			Address:   "潞城县",
			StartYear: 1983,
			EndYear:   1993,
		},
	},
	140423: {
		{
			Address:   "襄垣县",
			StartYear: 1985,
		},
	},
	140424: {
		{
			Address:   "屯留县",
			StartYear: 1985,
			EndYear:   2017,
		},
	},
	140425: {
		{
			Address:   "平顺县",
			StartYear: 1985,
		},
	},
	140426: {
		{
			Address:   "黎城县",
			StartYear: 1985,
		},
	},
	140427: {
		{
			Address:   "壶关县",
			StartYear: 1985,
		},
	},
	140428: {
		{
			Address:   "长子县",
			StartYear: 1985,
		},
	},
	140429: {
		{
			Address:   "武乡县",
			StartYear: 1985,
		},
	},
	140430: {
		{
			Address:   "沁县",
			StartYear: 1985,
		},
	},
	140431: {
		{
			Address:   "沁源县",
			StartYear: 1985,
		},
	},
	140481: {
		{
			Address:   "潞城市",
			StartYear: 1995,
			EndYear:   2017,
		},
	},
	140500: {
		{
			Address:   "晋城市",
			StartYear: 1985,
		},
	},
	140502: {
		{
			Address:   "城区",
			StartYear: 1985,
		},
	},
	140503: {
		{
			Address:   "郊区",
			StartYear: 1985,
			EndYear:   1995,
		},
	},
	140521: {
		{
			Address:   "沁水县",
			StartYear: 1985,
		},
	},
	140522: {
		{
			Address:   "阳城县",
			StartYear: 1985,
		},
	},
	140523: {
		{
			Address:   "高平县",
			StartYear: 1985,
			EndYear:   1992,
		},
	},
	140524: {
		{
			Address:   "陵川县",
			StartYear: 1985,
		},
	},
	140525: {
		{
			Address:   "泽州县",
			StartYear: 1996,
		},
	},
	140581: {
		{
			Address:   "高平市",
			StartYear: 1995,
		},
	},
	140600: {
		{
			Address:   "朔州市",
			StartYear: 1988,
		},
	},
	140602: {
		{
			Address:   "朔城区",
			StartYear: 1988,
		},
	},
	140603: {
		{
			Address:   "平鲁区",
			StartYear: 1988,
		},
	},
	140621: {
		{
			Address:   "山阴县",
			StartYear: 1988,
		},
	},
	140622: {
		{
			Address:   "应县",
			StartYear: 1993,
		},
	},
	140623: {
		{
			Address:   "右玉县",
			StartYear: 1993,
		},
	},
	140624: {
		{
			Address:   "怀仁县",
			StartYear: 1993,
			EndYear:   2017,
		},
	},
	140681: {
		{
			Address:   "怀仁市",
			StartYear: 2018,
		},
	},
	140700: {
		{
			Address:   "晋中市",
			StartYear: 1999,
		},
	},
	140702: {
		{
			Address:   "榆次区",
			StartYear: 1999,
		},
	},
	140703: {
		{
			Address:   "太谷区",
			StartYear: 2019,
		},
	},
	140721: {
		{
			Address:   "榆社县",
			StartYear: 1999,
		},
	},
	140722: {
		{
			Address:   "左权县",
			StartYear: 1999,
		},
	},
	140723: {
		{
			Address:   "和顺县",
			StartYear: 1999,
		},
	},
	140724: {
		{
			Address:   "昔阳县",
			StartYear: 1999,
		},
	},
	140725: {
		{
			Address:   "寿阳县",
			StartYear: 1999,
		},
	},
	140726: {
		{
			Address:   "太谷县",
			StartYear: 1999,
			EndYear:   2018,
		},
	},
	140727: {
		{
			Address:   "祁县",
			StartYear: 1999,
		},
	},
	140728: {
		{
			Address:   "平遥县",
			StartYear: 1999,
		},
	},
	140729: {
		{
			Address:   "灵石县",
			StartYear: 1999,
		},
	},
	140781: {
		{
			Address:   "介休市",
			StartYear: 1999,
		},
	},
	140800: {
		{
			Address:   "运城市",
			StartYear: 2000,
		},
	},
	140802: {
		{
			Address:   "盐湖区",
			StartYear: 2000,
		},
	},
	140821: {
		{
			Address:   "临猗县",
			StartYear: 2000,
		},
	},
	140822: {
		{
			Address:   "万荣县",
			StartYear: 2000,
		},
	},
	140823: {
		{
			Address:   "闻喜县",
			StartYear: 2000,
		},
	},
	140824: {
		{
			Address:   "稷山县",
			StartYear: 2000,
		},
	},
	140825: {
		{
			Address:   "新绛县",
			StartYear: 2000,
		},
	},
	140826: {
		{
			Address:   "绛县",
			StartYear: 2000,
		},
	},
	140827: {
		{
			Address:   "垣曲县",
			StartYear: 2000,
		},
	},
	140828: {
		{
			Address:   "夏县",
			StartYear: 2000,
		},
	},
	140829: {
		{
			Address:   "平陆县",
			StartYear: 2000,
		},
	},
	140830: {
		{
			Address:   "芮城县",
			StartYear: 2000,
		},
	},
	140881: {
		{
			Address:   "永济市",
			StartYear: 2000,
		},
	},
	140882: {
		{
			Address:   "河津市",
			StartYear: 2000,
		},
	},
	140900: {
		{
			Address:   "忻州市",
			StartYear: 2000,
		},
	},
	140902: {
		{
			Address:   "忻府区",
			StartYear: 2000,
		},
	},
	140921: {
		{
			Address:   "定襄县",
			StartYear: 2000,
		},
	},
	140922: {
		{
			Address:   "五台县",
			StartYear: 2000,
		},
	},
	140923: {
		{
			Address:   "代县",
			StartYear: 2000,
		},
	},
	140924: {
		{
			Address:   "繁峙县",
			StartYear: 2000,
		},
	},
	140925: {
		{
			Address:   "宁武县",
			StartYear: 2000,
		},
	},
	140926: {
		{
			Address:   "静乐县",
			StartYear: 2000,
		},
	},
	140927: {
		{
			Address:   "神池县",
			StartYear: 2000,
		},
	},
	140928: {
		{
			Address:   "五寨县",
			StartYear: 2000,
		},
	},
	140929: {
		{
			Address:   "岢岚县",
			StartYear: 2000,
		},
	},
	140930: {
		{
			Address:   "河曲县",
			StartYear: 2000,
		},
	},
	140931: {
		{
			Address:   "保德县",
			StartYear: 2000,
		},
	},
	140932: {
		{
			Address:   "偏关县",
			StartYear: 2000,
		},
	},
	140981: {
		{
			Address:   "原平市",
			StartYear: 2000,
		},
	},
	141000: {
		{
			Address:   "临汾市",
			StartYear: 2000,
		},
	},
	141002: {
		{
			Address:   "尧都区",
			StartYear: 2000,
		},
	},
	141021: {
		{
			Address:   "曲沃县",
			StartYear: 2000,
		},
	},
	141022: {
		{
			Address:   "翼城县",
			StartYear: 2000,
		},
	},
	141023: {
		{
			Address:   "襄汾县",
			StartYear: 2000,
		},
	},
	141024: {
		{
			Address:   "洪洞县",
			StartYear: 2000,
		},
	},
	141025: {
		{
			Address:   "古县",
			StartYear: 2000,
		},
	},
	141026: {
		{
			Address:   "安泽县",
			StartYear: 2000,
		},
	},
	141027: {
		{
			Address:   "浮山县",
			StartYear: 2000,
		},
	},
	141028: {
		{
			Address:   "吉县",
			StartYear: 2000,
		},
	},
	141029: {
		{
			Address:   "乡宁县",
			StartYear: 2000,
		},
	},
	141030: {
		{
			Address:   "大宁县",
			StartYear: 2000,
		},
	},
	141031: {
		{
			Address:   "隰县",
			StartYear: 2000,
		},
	},
	141032: {
		{
			Address:   "永和县",
			StartYear: 2000,
		},
	},
	141033: {
		{
			Address:   "蒲县",
			StartYear: 2000,
		},
	},
	141034: {
		{
			Address:   "汾西县",
			StartYear: 2000,
		},
	},
	141081: {
		{
			Address:   "侯马市",
			StartYear: 2000,
		},
	},
	141082: {
		{
			Address:   "霍州市",
			StartYear: 2000,
		},
	},
	141100: {
		{
			Address:   "吕梁市",
			StartYear: 2003,
		},
	},
	141102: {
		{
			Address:   "离石区",
			StartYear: 2003,
		},
	},
	141121: {
		{
			Address:   "文水县",
			StartYear: 2003,
		},
	},
	141122: {
		{
			Address:   "交城县",
			StartYear: 2003,
		},
	},
	141123: {
		{
			Address:   "兴县",
			StartYear: 2003,
		},
	},
	141124: {
		{
			Address:   "临县",
			StartYear: 2003,
		},
	},
	141125: {
		{
			Address:   "柳林县",
			StartYear: 2003,
		},
	},
	141126: {
		{
			Address:   "石楼县",
			StartYear: 2003,
		},
	},
	141127: {
		{
			Address:   "岚县",
			StartYear: 2003,
		},
	},
	141128: {
		{
			Address:   "方山县",
			StartYear: 2003,
		},
	},
	141129: {
		{
			Address:   "中阳县",
			StartYear: 2003,
		},
	},
	141130: {
		{
			Address:   "交口县",
			StartYear: 2003,
		},
	},
	141181: {
		{
			Address:   "孝义市",
			StartYear: 2003,
		},
	},
	141182: {
		{
			Address:   "汾阳市",
			StartYear: 2003,
		},
	},
	142100: {
		{
			Address: "雁北地区",
			EndYear: 1992,
		},
	},
	142121: {
		{
			Address: "大同县",
			EndYear: 1981,
		},
		{
			Address:   "阳高县",
			StartYear: 1982,
			EndYear:   1992,
		},
	},
	142122: {
		{
			Address: "阳高县",
			EndYear: 1981,
		},
		{
			Address:   "天镇县",
			StartYear: 1982,
			EndYear:   1992,
		},
	},
	142123: {
		{
			Address: "天镇县",
			EndYear: 1981,
		},
		{
			Address:   "广灵县",
			StartYear: 1982,
			EndYear:   1992,
		},
	},
	142124: {
		{
			Address: "广灵县",
			EndYear: 1981,
		},
		{
			Address:   "灵丘县",
			StartYear: 1982,
			EndYear:   1992,
		},
	},
	142125: {
		{
			Address: "灵丘县",
			EndYear: 1981,
		},
		{
			Address:   "浑源县",
			StartYear: 1982,
			EndYear:   1992,
		},
	},
	142126: {
		{
			Address: "浑源县",
			EndYear: 1981,
		},
		{
			Address:   "应县",
			StartYear: 1982,
			EndYear:   1992,
		},
	},
	142127: {
		{
			Address: "怀仁县",
			EndYear: 1981,
		},
		{
			Address:   "山阴县",
			StartYear: 1982,
			EndYear:   1987,
		},
	},
	142128: {
		{
			Address: "应县",
			EndYear: 1981,
		},
		{
			Address:   "朔县",
			StartYear: 1982,
			EndYear:   1987,
		},
	},
	142129: {
		{
			Address: "山阴县",
			EndYear: 1981,
		},
		{
			Address:   "平鲁县",
			StartYear: 1982,
			EndYear:   1987,
		},
	},
	142130: {
		{
			Address: "朔县",
			EndYear: 1981,
		},
		{
			Address:   "左云县",
			StartYear: 1982,
			EndYear:   1992,
		},
	},
	142131: {
		{
			Address: "平鲁县",
			EndYear: 1981,
		},
		{
			Address:   "右玉县",
			StartYear: 1982,
			EndYear:   1992,
		},
	},
	142132: {
		{
			Address: "左云县",
			EndYear: 1981,
		},
		{
			Address:   "大同县",
			StartYear: 1982,
			EndYear:   1992,
		},
	},
	142133: {
		{
			Address: "右玉县",
			EndYear: 1981,
		},
		{
			Address:   "怀仁县",
			StartYear: 1982,
			EndYear:   1992,
		},
	},
	142200: {
		{
			Address: "忻县地区",
			EndYear: 1982,
		},
		{
			Address:   "忻州地区",
			StartYear: 1983,
			EndYear:   1999,
		},
	},
	142201: {
		{
			Address:   "忻州市",
			StartYear: 1983,
			EndYear:   1999,
		},
	},
	142202: {
		{
			Address:   "原平市",
			StartYear: 1993,
			EndYear:   1999,
		},
	},
	142221: {
		{
			Address: "忻县",
			EndYear: 1982,
		},
	},
	142222: {
		{
			Address: "原平县",
			EndYear: 1981,
		},
		{
			Address:   "定襄县",
			StartYear: 1982,
			EndYear:   1999,
		},
	},
	142223: {
		{
			Address: "代县",
			EndYear: 1981,
		},
		{
			Address:   "五台县",
			StartYear: 1982,
			EndYear:   1999,
		},
	},
	142224: {
		{
			Address: "繁峙县",
			EndYear: 1981,
		},
		{
			Address:   "原平县",
			StartYear: 1982,
			EndYear:   1992,
		},
	},
	142225: {
		{
			Address: "五台县",
			EndYear: 1981,
		},
		{
			Address:   "代县",
			StartYear: 1982,
			EndYear:   1999,
		},
	},
	142226: {
		{
			Address: "定襄县",
			EndYear: 1981,
		},
		{
			Address:   "繁峙县",
			StartYear: 1982,
			EndYear:   1999,
		},
	},
	142227: {
		{
			Address: "静乐县",
			EndYear: 1981,
		},
		{
			Address:   "宁武县",
			StartYear: 1982,
			EndYear:   1999,
		},
	},
	142228: {
		{
			Address: "岢岚县",
			EndYear: 1981,
		},
		{
			Address:   "静乐县",
			StartYear: 1982,
			EndYear:   1999,
		},
	},
	142229: {
		{
			Address: "保德县",
			EndYear: 1981,
		},
		{
			Address:   "神池县",
			StartYear: 1982,
			EndYear:   1999,
		},
	},
	142230: {
		{
			Address: "五寨县",
			EndYear: 1999,
		},
	},
	142231: {
		{
			Address: "河曲县",
			EndYear: 1981,
		},
		{
			Address:   "岢岚县",
			StartYear: 1982,
			EndYear:   1999,
		},
	},
	142232: {
		{
			Address: "偏关县",
			EndYear: 1981,
		},
		{
			Address:   "河曲县",
			StartYear: 1982,
			EndYear:   1999,
		},
	},
	142233: {
		{
			Address: "神池县",
			EndYear: 1981,
		},
		{
			Address:   "保德县",
			StartYear: 1982,
			EndYear:   1999,
		},
	},
	142234: {
		{
			Address: "宁武县",
			EndYear: 1981,
		},
		{
			Address:   "偏关县",
			StartYear: 1982,
			EndYear:   1999,
		},
	},
	142300: {
		{
			Address: "晋中地区",
			EndYear: 1981,
		},
		{
			Address:   "吕梁地区",
			StartYear: 1982,
			EndYear:   2002,
		},
	},
	142301: {
		{
			Address: "榆次市",
			EndYear: 1981,
		},
		{
			Address:   "孝义市",
			StartYear: 1992,
			EndYear:   2002,
		},
	},
	142302: {
		{
			Address:   "离石市",
			StartYear: 1996,
			EndYear:   2002,
		},
	},
	142303: {
		{
			Address:   "汾阳市",
			StartYear: 1996,
			EndYear:   2002,
		},
	},
	142321: {
		{
			Address: "榆次县",
			EndYear: 1981,
		},
		{
			Address:   "汾阳县",
			StartYear: 1982,
			EndYear:   1995,
		},
	},
	142322: {
		{
			Address: "寿阳县",
			EndYear: 1981,
		},
		{
			Address:   "文水县",
			StartYear: 1982,
			EndYear:   2002,
		},
	},
	142323: {
		{
			Address: "盂县",
			EndYear: 1981,
		},
		{
			Address:   "交城县",
			StartYear: 1982,
			EndYear:   2002,
		},
	},
	142324: {
		{
			Address: "平定县",
			EndYear: 1981,
		},
		{
			Address:   "孝义县",
			StartYear: 1982,
			EndYear:   1991,
		},
	},
	142325: {
		{
			Address: "昔阳县",
			EndYear: 1981,
		},
		{
			Address:   "兴县",
			StartYear: 1982,
			EndYear:   2002,
		},
	},
	142326: {
		{
			Address: "和顺县",
			EndYear: 1981,
		},
		{
			Address:   "临县",
			StartYear: 1982,
			EndYear:   2002,
		},
	},
	142327: {
		{
			Address: "左权县",
			EndYear: 1981,
		},
		{
			Address:   "柳林县",
			StartYear: 1982,
			EndYear:   2002,
		},
	},
	142328: {
		{
			Address: "榆社县",
			EndYear: 1981,
		},
		{
			Address:   "石楼县",
			StartYear: 1982,
			EndYear:   2002,
		},
	},
	142329: {
		{
			Address: "太谷县",
			EndYear: 1981,
		},
		{
			Address:   "岚县",
			StartYear: 1982,
			EndYear:   2002,
		},
	},
	142330: {
		{
			Address: "祁县",
			EndYear: 1981,
		},
		{
			Address:   "方山县",
			StartYear: 1982,
			EndYear:   2002,
		},
	},
	142331: {
		{
			Address: "平遥县",
			EndYear: 1981,
		},
		{
			Address:   "离石县",
			StartYear: 1982,
			EndYear:   1995,
		},
	},
	142332: {
		{
			Address: "介休县",
			EndYear: 1981,
		},
		{
			Address:   "中阳县",
			StartYear: 1982,
			EndYear:   2002,
		},
	},
	142333: {
		{
			Address: "灵石县",
			EndYear: 1981,
		},
		{
			Address:   "交口县",
			StartYear: 1982,
			EndYear:   2002,
		},
	},
	142400: {
		{
			Address: "吕梁地区",
			EndYear: 1981,
		},
		{
			Address:   "晋中地区",
			StartYear: 1982,
			EndYear:   1998,
		},
	},
	142401: {
		{
			Address:   "榆次市",
			StartYear: 1982,
			EndYear:   1998,
		},
	},
	142402: {
		{
			Address:   "介休市",
			StartYear: 1992,
			EndYear:   1998,
		},
	},
	142421: {
		{
			Address: "离石县",
			EndYear: 1981,
		},
		{
			Address:   "榆社县",
			StartYear: 1982,
			EndYear:   1998,
		},
	},
	142422: {
		{
			Address: "孝义县",
			EndYear: 1981,
		},
		{
			Address:   "左权县",
			StartYear: 1982,
			EndYear:   1998,
		},
	},
	142423: {
		{
			Address: "兴县",
			EndYear: 1981,
		},
		{
			Address:   "和顺县",
			StartYear: 1982,
			EndYear:   1998,
		},
	},
	142424: {
		{
			Address: "交口县",
			EndYear: 1981,
		},
		{
			Address:   "昔阳县",
			StartYear: 1982,
			EndYear:   1998,
		},
	},
	142425: {
		{
			Address: "方山县",
			EndYear: 1981,
		},
		{
			Address:   "平定县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	142426: {
		{
			Address: "石楼县",
			EndYear: 1981,
		},
		{
			Address:   "盂县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	142427: {
		{
			Address: "岚县",
			EndYear: 1981,
		},
		{
			Address:   "寿阳县",
			StartYear: 1982,
			EndYear:   1998,
		},
	},
	142428: {
		{
			Address: "中阳县",
			EndYear: 1981,
		},
		{
			Address:   "榆次县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	142429: {
		{
			Address: "交城县",
			EndYear: 1981,
		},
		{
			Address:   "太谷县",
			StartYear: 1982,
			EndYear:   1998,
		},
	},
	142430: {
		{
			Address: "临县",
			EndYear: 1981,
		},
		{
			Address:   "祁县",
			StartYear: 1982,
			EndYear:   1998,
		},
	},
	142431: {
		{
			Address: "文水县",
			EndYear: 1981,
		},
		{
			Address:   "平遥县",
			StartYear: 1982,
			EndYear:   1998,
		},
	},
	142432: {
		{
			Address: "柳林县",
			EndYear: 1981,
		},
		{
			Address:   "介休县",
			StartYear: 1982,
			EndYear:   1991,
		},
	},
	142433: {
		{
			Address: "汾阳县",
			EndYear: 1981,
		},
		{
			Address:   "灵石县",
			StartYear: 1982,
			EndYear:   1998,
		},
	},
	142500: {
		{
			Address: "晋东南地区",
			EndYear: 1984,
		},
	},
	142501: {
		{
			Address:   "晋城市",
			StartYear: 1983,
			EndYear:   1984,
		},
	},
	142521: {
		{
			Address: "长治县",
			EndYear: 1982,
		},
	},
	142522: {
		{
			Address: "潞城县",
			EndYear: 1982,
		},
	},
	142523: {
		{
			Address: "襄垣县",
			EndYear: 1981,
		},
		{
			Address:   "屯留县",
			StartYear: 1982,
			EndYear:   1984,
		},
	},
	142524: {
		{
			Address: "武乡县",
			EndYear: 1981,
		},
		{
			Address:   "长子县",
			StartYear: 1982,
			EndYear:   1984,
		},
	},
	142525: {
		{
			Address: "黎城县",
			EndYear: 1981,
		},
		{
			Address:   "沁水县",
			StartYear: 1982,
			EndYear:   1984,
		},
	},
	142526: {
		{
			Address: "平顺县",
			EndYear: 1981,
		},
		{
			Address:   "阳城县",
			StartYear: 1982,
			EndYear:   1984,
		},
	},
	142527: {
		{
			Address: "壶关县",
			EndYear: 1981,
		},
		{
			Address:   "晋城县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	142528: {
		{
			Address: "陵川县",
			EndYear: 1981,
		},
		{
			Address:   "高平县",
			StartYear: 1982,
			EndYear:   1984,
		},
	},
	142529: {
		{
			Address: "高平县",
			EndYear: 1981,
		},
		{
			Address:   "陵川县",
			StartYear: 1982,
			EndYear:   1984,
		},
	},
	142530: {
		{
			Address: "晋城县",
			EndYear: 1981,
		},
		{
			Address:   "壶关县",
			StartYear: 1982,
			EndYear:   1984,
		},
	},
	142531: {
		{
			Address: "阳城县",
			EndYear: 1981,
		},
		{
			Address:   "平顺县",
			StartYear: 1982,
			EndYear:   1984,
		},
	},
	142532: {
		{
			Address: "沁水县",
			EndYear: 1981,
		},
		{
			Address:   "黎城县",
			StartYear: 1982,
			EndYear:   1984,
		},
	},
	142533: {
		{
			Address: "长子县",
			EndYear: 1981,
		},
		{
			Address:   "武乡县",
			StartYear: 1982,
			EndYear:   1984,
		},
	},
	142534: {
		{
			Address: "屯留县",
			EndYear: 1981,
		},
		{
			Address:   "襄垣县",
			StartYear: 1982,
			EndYear:   1984,
		},
	},
	142535: {
		{
			Address: "沁源县",
			EndYear: 1981,
		},
		{
			Address:   "沁县",
			StartYear: 1982,
			EndYear:   1984,
		},
	},
	142536: {
		{
			Address: "沁县",
			EndYear: 1981,
		},
		{
			Address:   "沁源县",
			StartYear: 1982,
			EndYear:   1984,
		},
	},
	142600: {
		{
			Address: "临汾地区",
			EndYear: 1999,
		},
	},
	142601: {
		{
			Address: "临汾市",
			EndYear: 1999,
		},
	},
	142602: {
		{
			Address: "侯马市",
			EndYear: 1999,
		},
	},
	142603: {
		{
			Address:   "霍州市",
			StartYear: 1989,
			EndYear:   1999,
		},
	},
	142621: {
		{
			Address: "临汾县",
			EndYear: 1981,
		},
		{
			Address:   "曲沃县",
			StartYear: 1982,
			EndYear:   1999,
		},
	},
	142622: {
		{
			Address: "隰县",
			EndYear: 1981,
		},
		{
			Address:   "翼城县",
			StartYear: 1982,
			EndYear:   1999,
		},
	},
	142623: {
		{
			Address: "汾西县",
			EndYear: 1981,
		},
		{
			Address:   "襄汾县",
			StartYear: 1982,
			EndYear:   1999,
		},
	},
	142624: {
		{
			Address: "永和县",
			EndYear: 1981,
		},
		{
			Address:   "临汾县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	142625: {
		{
			Address: "安泽县",
			EndYear: 1981,
		},
		{
			Address:   "洪洞县",
			StartYear: 1982,
			EndYear:   1999,
		},
	},
	142626: {
		{
			Address: "洪洞县",
			EndYear: 1981,
		},
		{
			Address:   "霍县",
			StartYear: 1982,
			EndYear:   1988,
		},
	},
	142627: {
		{
			Address: "古县",
			EndYear: 1999,
		},
	},
	142628: {
		{
			Address: "霍县",
			EndYear: 1981,
		},
		{
			Address:   "安泽县",
			StartYear: 1982,
			EndYear:   1999,
		},
	},
	142629: {
		{
			Address: "翼城县",
			EndYear: 1981,
		},
		{
			Address:   "浮山县",
			StartYear: 1982,
			EndYear:   1999,
		},
	},
	142630: {
		{
			Address: "浮山县",
			EndYear: 1981,
		},
		{
			Address:   "吉县",
			StartYear: 1982,
			EndYear:   1999,
		},
	},
	142631: {
		{
			Address: "曲沃县",
			EndYear: 1981,
		},
		{
			Address:   "乡宁县",
			StartYear: 1982,
			EndYear:   1999,
		},
	},
	142632: {
		{
			Address: "襄汾县",
			EndYear: 1981,
		},
		{
			Address:   "蒲县",
			StartYear: 1982,
			EndYear:   1999,
		},
	},
	142633: {
		{
			Address: "吉县",
			EndYear: 1981,
		},
		{
			Address:   "大宁县",
			StartYear: 1982,
			EndYear:   1999,
		},
	},
	142634: {
		{
			Address: "乡宁县",
			EndYear: 1981,
		},
		{
			Address:   "永和县",
			StartYear: 1982,
			EndYear:   1999,
		},
	},
	142635: {
		{
			Address: "大宁县",
			EndYear: 1981,
		},
		{
			Address:   "隰县",
			StartYear: 1982,
			EndYear:   1999,
		},
	},
	142636: {
		{
			Address: "蒲县",
			EndYear: 1981,
		},
		{
			Address:   "汾西县",
			StartYear: 1982,
			EndYear:   1999,
		},
	},
	142700: {
		{
			Address: "运城地区",
			EndYear: 1999,
		},
	},
	142701: {
		{
			Address:   "运城市",
			StartYear: 1983,
			EndYear:   1999,
		},
	},
	142702: {
		{
			Address:   "永济市",
			StartYear: 1994,
			EndYear:   1999,
		},
	},
	142703: {
		{
			Address:   "河津市",
			StartYear: 1994,
			EndYear:   1999,
		},
	},
	142721: {
		{
			Address: "运城县",
			EndYear: 1982,
		},
	},
	142722: {
		{
			Address: "夏县",
			EndYear: 1981,
		},
		{
			Address:   "永济县",
			StartYear: 1982,
			EndYear:   1993,
		},
	},
	142723: {
		{
			Address: "闻喜县",
			EndYear: 1981,
		},
		{
			Address:   "芮城县",
			StartYear: 1982,
			EndYear:   1999,
		},
	},
	142724: {
		{
			Address: "绛县",
			EndYear: 1981,
		},
		{
			Address:   "临猗县",
			StartYear: 1982,
			EndYear:   1999,
		},
	},
	142725: {
		{
			Address: "垣曲县",
			EndYear: 1981,
		},
		{
			Address:   "万荣县",
			StartYear: 1982,
			EndYear:   1999,
		},
	},
	142726: {
		{
			Address: "平陆县",
			EndYear: 1981,
		},
		{
			Address:   "新绛县",
			StartYear: 1982,
			EndYear:   1999,
		},
	},
	142727: {
		{
			Address: "芮城县",
			EndYear: 1981,
		},
		{
			Address:   "稷山县",
			StartYear: 1982,
			EndYear:   1999,
		},
	},
	142728: {
		{
			Address: "永济县",
			EndYear: 1981,
		},
		{
			Address:   "河津县",
			StartYear: 1982,
			EndYear:   1993,
		},
	},
	142729: {
		{
			Address: "临猗县",
			EndYear: 1981,
		},
		{
			Address:   "闻喜县",
			StartYear: 1982,
			EndYear:   1999,
		},
	},
	142730: {
		{
			Address: "万荣县",
			EndYear: 1981,
		},
		{
			Address:   "夏县",
			StartYear: 1982,
			EndYear:   1999,
		},
	},
	142731: {
		{
			Address: "新绛县",
			EndYear: 1981,
		},
		{
			Address:   "绛县",
			StartYear: 1982,
			EndYear:   1999,
		},
	},
	142732: {
		{
			Address: "稷山县",
			EndYear: 1981,
		},
		{
			Address:   "平陆县",
			StartYear: 1982,
			EndYear:   1999,
		},
	},
	142733: {
		{
			Address: "河津县",
			EndYear: 1981,
		},
		{
			Address:   "垣曲县",
			StartYear: 1982,
			EndYear:   1999,
		},
	},
	149001: {
		{
			Address:   "古交市",
			StartYear: 1988,
			EndYear:   1994,
		},
	},
	149002: {
		{
			Address:   "高平市",
			StartYear: 1993,
			EndYear:   1994,
		},
	},
	149003: {
		{
			Address:   "潞城市",
			StartYear: 1994,
			EndYear:   1994,
		},
	},
	150000: {
		{
			Address: "内蒙古自治区",
		},
	},
	150100: {
		{
			Address: "呼和浩特市",
		},
	},
	150102: {
		{
			Address: "新城区",
		},
	},
	150103: {
		{
			Address: "回民区",
		},
	},
	150104: {
		{
			Address: "玉泉区",
		},
	},
	150105: {
		{
			Address: "郊区",
			EndYear: 1999,
		},
		{
			Address:   "赛罕区",
			StartYear: 2000,
		},
	},
	150121: {
		{
			Address: "土默特左旗",
		},
	},
	150122: {
		{
			Address: "托克托县",
		},
	},
	150123: {
		{
			Address:   "和林格尔县",
			StartYear: 1995,
		},
	},
	150124: {
		{
			Address:   "清水河县",
			StartYear: 1995,
		},
	},
	150125: {
		{
			Address:   "武川县",
			StartYear: 1996,
		},
	},
	150200: {
		{
			Address: "包头市",
		},
	},
	150202: {
		{
			Address: "东河区",
		},
	},
	150203: {
		{
			Address: "昆都仑区",
		},
	},
	150204: {
		{
			Address: "青山区",
		},
	},
	150205: {
		{
			Address: "石拐矿区",
			EndYear: 1998,
		},
		{
			Address:   "石拐区",
			StartYear: 1999,
		},
	},
	150206: {
		{
			Address: "白云矿区",
			EndYear: 2006,
		},
		{
			Address:   "白云鄂博矿区",
			StartYear: 2007,
		},
	},
	150207: {
		{
			Address: "郊区",
			EndYear: 1998,
		},
		{
			Address:   "九原区",
			StartYear: 1999,
		},
	},
	150221: {
		{
			Address: "土默特右旗",
		},
	},
	150222: {
		{
			Address: "固阳县",
		},
	},
	150223: {
		{
			Address:   "达尔罕茂明安联合旗",
			StartYear: 1996,
		},
	},
	150300: {
		{
			Address: "乌海市",
		},
	},
	150302: {
		{
			Address: "海勃湾区",
		},
	},
	150303: {
		{
			Address: "海南区",
		},
	},
	150304: {
		{
			Address: "乌达区",
		},
	},
	150400: {
		{
			Address:   "赤峰市",
			StartYear: 1983,
		},
	},
	150402: {
		{
			Address:   "红山区",
			StartYear: 1983,
		},
	},
	150403: {
		{
			Address:   "元宝山区",
			StartYear: 1983,
		},
	},
	150404: {
		{
			Address:   "郊区",
			StartYear: 1983,
			EndYear:   1992,
		},
		{
			Address:   "松山区",
			StartYear: 1993,
		},
	},
	150421: {
		{
			Address:   "阿鲁科尔沁旗",
			StartYear: 1983,
		},
	},
	150422: {
		{
			Address:   "巴林左旗",
			StartYear: 1983,
		},
	},
	150423: {
		{
			Address:   "巴林右旗",
			StartYear: 1983,
		},
	},
	150424: {
		{
			Address:   "林西县",
			StartYear: 1983,
		},
	},
	150425: {
		{
			Address:   "克什克腾旗",
			StartYear: 1983,
		},
	},
	150426: {
		{
			Address:   "翁牛特旗",
			StartYear: 1983,
		},
	},
	150428: {
		{
			Address:   "喀喇沁旗",
			StartYear: 1983,
		},
	},
	150429: {
		{
			Address:   "宁城县",
			StartYear: 1983,
		},
	},
	150430: {
		{
			Address:   "敖汉旗",
			StartYear: 1983,
		},
	},
	150500: {
		{
			Address:   "通辽市",
			StartYear: 1999,
		},
	},
	150502: {
		{
			Address:   "科尔沁区",
			StartYear: 1999,
		},
	},
	150521: {
		{
			Address:   "科尔沁左翼中旗",
			StartYear: 1999,
		},
	},
	150522: {
		{
			Address:   "科尔沁左翼后旗",
			StartYear: 1999,
		},
	},
	150523: {
		{
			Address:   "开鲁县",
			StartYear: 1999,
		},
	},
	150524: {
		{
			Address:   "库伦旗",
			StartYear: 1999,
		},
	},
	150525: {
		{
			Address:   "奈曼旗",
			StartYear: 1999,
		},
	},
	150526: {
		{
			Address:   "扎鲁特旗",
			StartYear: 1999,
		},
	},
	150581: {
		{
			Address:   "霍林郭勒市",
			StartYear: 1999,
		},
	},
	150600: {
		{
			Address:   "鄂尔多斯市",
			StartYear: 2001,
		},
	},
	150602: {
		{
			Address:   "东胜区",
			StartYear: 2001,
		},
	},
	150603: {
		{
			Address:   "康巴什区",
			StartYear: 2016,
		},
	},
	150621: {
		{
			Address:   "达拉特旗",
			StartYear: 2001,
		},
	},
	150622: {
		{
			Address:   "准格尔旗",
			StartYear: 2001,
		},
	},
	150623: {
		{
			Address:   "鄂托克前旗",
			StartYear: 2001,
		},
	},
	150624: {
		{
			Address:   "鄂托克旗",
			StartYear: 2001,
		},
	},
	150625: {
		{
			Address:   "杭锦旗",
			StartYear: 2001,
		},
	},
	150626: {
		{
			Address:   "乌审旗",
			StartYear: 2001,
		},
	},
	150627: {
		{
			Address:   "伊金霍洛旗",
			StartYear: 2001,
		},
	},
	150700: {
		{
			Address:   "呼伦贝尔市",
			StartYear: 2001,
		},
	},
	150702: {
		{
			Address:   "海拉尔区",
			StartYear: 2001,
		},
	},
	150703: {
		{
			Address:   "扎赉诺尔区",
			StartYear: 2013,
		},
	},
	150721: {
		{
			Address:   "阿荣旗",
			StartYear: 2001,
		},
	},
	150722: {
		{
			Address:   "莫力达瓦达斡尔族自治旗",
			StartYear: 2001,
		},
	},
	150723: {
		{
			Address:   "鄂伦春自治旗",
			StartYear: 2001,
		},
	},
	150724: {
		{
			Address:   "鄂温克族自治旗",
			StartYear: 2001,
		},
	},
	150725: {
		{
			Address:   "陈巴尔虎旗",
			StartYear: 2001,
		},
	},
	150726: {
		{
			Address:   "新巴尔虎左旗",
			StartYear: 2001,
		},
	},
	150727: {
		{
			Address:   "新巴尔虎右旗",
			StartYear: 2001,
		},
	},
	150781: {
		{
			Address:   "满洲里市",
			StartYear: 2001,
		},
	},
	150782: {
		{
			Address:   "牙克石市",
			StartYear: 2001,
		},
	},
	150783: {
		{
			Address:   "扎兰屯市",
			StartYear: 2001,
		},
	},
	150784: {
		{
			Address:   "额尔古纳市",
			StartYear: 2001,
		},
	},
	150785: {
		{
			Address:   "根河市",
			StartYear: 2001,
		},
	},
	150800: {
		{
			Address:   "巴彦淖尔市",
			StartYear: 2003,
		},
	},
	150802: {
		{
			Address:   "临河区",
			StartYear: 2003,
		},
	},
	150821: {
		{
			Address:   "五原县",
			StartYear: 2003,
		},
	},
	150822: {
		{
			Address:   "磴口县",
			StartYear: 2003,
		},
	},
	150823: {
		{
			Address:   "乌拉特前旗",
			StartYear: 2003,
		},
	},
	150824: {
		{
			Address:   "乌拉特中旗",
			StartYear: 2003,
		},
	},
	150825: {
		{
			Address:   "乌拉特后旗",
			StartYear: 2003,
		},
	},
	150826: {
		{
			Address:   "杭锦后旗",
			StartYear: 2003,
		},
	},
	150900: {
		{
			Address:   "乌兰察布市",
			StartYear: 2003,
		},
	},
	150902: {
		{
			Address:   "集宁区",
			StartYear: 2003,
		},
	},
	150921: {
		{
			Address:   "卓资县",
			StartYear: 2003,
		},
	},
	150922: {
		{
			Address:   "化德县",
			StartYear: 2003,
		},
	},
	150923: {
		{
			Address:   "商都县",
			StartYear: 2003,
		},
	},
	150924: {
		{
			Address:   "兴和县",
			StartYear: 2003,
		},
	},
	150925: {
		{
			Address:   "凉城县",
			StartYear: 2003,
		},
	},
	150926: {
		{
			Address:   "察哈尔右翼前旗",
			StartYear: 2003,
		},
	},
	150927: {
		{
			Address:   "察哈尔右翼中旗",
			StartYear: 2003,
		},
	},
	150928: {
		{
			Address:   "察哈尔右翼后旗",
			StartYear: 2003,
		},
	},
	150929: {
		{
			Address:   "四子王旗",
			StartYear: 2003,
		},
	},
	150981: {
		{
			Address:   "丰镇市",
			StartYear: 2003,
		},
	},
	152100: {
		{
			Address: "呼伦贝尔盟",
			EndYear: 2000,
		},
	},
	152101: {
		{
			Address: "海拉尔市",
			EndYear: 2000,
		},
	},
	152102: {
		{
			Address: "满洲里市",
			EndYear: 2000,
		},
	},
	152103: {
		{
			Address:   "扎兰屯市",
			StartYear: 1983,
			EndYear:   2000,
		},
	},
	152104: {
		{
			Address:   "牙克石市",
			StartYear: 1983,
			EndYear:   2000,
		},
	},
	152105: {
		{
			Address:   "根河市",
			StartYear: 1994,
			EndYear:   2000,
		},
	},
	152106: {
		{
			Address:   "额尔古纳市",
			StartYear: 1994,
			EndYear:   2000,
		},
	},
	152121: {
		{
			Address: "布特哈旗",
			EndYear: 1982,
		},
	},
	152122: {
		{
			Address: "阿荣旗",
			EndYear: 2000,
		},
	},
	152123: {
		{
			Address: "莫力达瓦达斡尔族自治旗",
			EndYear: 2000,
		},
	},
	152124: {
		{
			Address: "喜桂图旗",
			EndYear: 1982,
		},
	},
	152125: {
		{
			Address: "额尔古纳右旗",
			EndYear: 1993,
		},
	},
	152126: {
		{
			Address: "额尔古纳左旗",
			EndYear: 1993,
		},
	},
	152127: {
		{
			Address: "鄂伦春自治旗",
			EndYear: 2000,
		},
	},
	152128: {
		{
			Address: "鄂温克族自治旗",
			EndYear: 2000,
		},
	},
	152129: {
		{
			Address: "新巴尔虎右旗",
			EndYear: 2000,
		},
	},
	152130: {
		{
			Address: "新巴尔虎左旗",
			EndYear: 2000,
		},
	},
	152131: {
		{
			Address: "陈巴尔虎旗",
			EndYear: 2000,
		},
	},
	152200: {
		{
			Address: "兴安盟",
		},
	},
	152201: {
		{
			Address: "乌兰浩特市",
		},
	},
	152202: {
		{
			Address:   "阿尔山市",
			StartYear: 1996,
		},
	},
	152221: {
		{
			Address: "科尔沁右翼前旗",
		},
	},
	152222: {
		{
			Address: "科尔沁右翼中旗",
		},
	},
	152223: {
		{
			Address: "扎赉特旗",
		},
	},
	152224: {
		{
			Address: "突泉县",
		},
	},
	152300: {
		{
			Address: "哲里木盟",
			EndYear: 1998,
		},
	},
	152301: {
		{
			Address: "通辽市",
			EndYear: 1998,
		},
	},
	152302: {
		{
			Address:   "霍林郭勒市",
			StartYear: 1985,
			EndYear:   1998,
		},
	},
	152321: {
		{
			Address: "通辽县",
			EndYear: 1985,
		},
	},
	152322: {
		{
			Address: "科尔沁左翼中旗",
			EndYear: 1998,
		},
	},
	152323: {
		{
			Address: "科尔沁左翼后旗",
			EndYear: 1998,
		},
	},
	152324: {
		{
			Address: "开鲁县",
			EndYear: 1998,
		},
	},
	152325: {
		{
			Address: "库伦旗",
			EndYear: 1998,
		},
	},
	152326: {
		{
			Address: "奈曼旗",
			EndYear: 1998,
		},
	},
	152327: {
		{
			Address: "扎鲁特旗",
			EndYear: 1998,
		},
	},
	152400: {
		{
			Address: "昭乌达盟",
			EndYear: 1982,
		},
	},
	152401: {
		{
			Address: "赤峰市",
			EndYear: 1982,
		},
	},
	152421: {
		{
			Address: "赤峰县",
			EndYear: 1981,
		},
		{
			Address:   "阿鲁科尔沁旗",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	152422: {
		{
			Address: "巴林左旗",
			EndYear: 1982,
		},
	},
	152423: {
		{
			Address: "巴林右旗",
			EndYear: 1982,
		},
	},
	152424: {
		{
			Address: "林西县",
			EndYear: 1982,
		},
	},
	152425: {
		{
			Address: "克什克腾旗",
			EndYear: 1982,
		},
	},
	152426: {
		{
			Address: "翁牛特旗",
			EndYear: 1982,
		},
	},
	152427: {
		{
			Address:   "赤峰县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	152428: {
		{
			Address: "喀喇沁旗",
			EndYear: 1982,
		},
	},
	152429: {
		{
			Address: "宁城县",
			EndYear: 1982,
		},
	},
	152430: {
		{
			Address: "敖汉旗",
			EndYear: 1982,
		},
	},
	152431: {
		{
			Address: "阿鲁科尔沁旗",
			EndYear: 1981,
		},
	},
	152500: {
		{
			Address: "伊克昭盟",
			EndYear: 1981,
		},
		{
			Address:   "锡林郭勒盟",
			StartYear: 1982,
		},
	},
	152501: {
		{
			Address:   "二连浩特市",
			StartYear: 1982,
		},
	},
	152502: {
		{
			Address:   "锡林浩特市",
			StartYear: 1983,
		},
	},
	152521: {
		{
			Address: "东胜县",
			EndYear: 1981,
		},
		{
			Address:   "阿巴哈纳尔旗",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	152522: {
		{
			Address: "达拉特旗",
			EndYear: 1981,
		},
		{
			Address:   "阿巴嘎旗",
			StartYear: 1982,
		},
	},
	152523: {
		{
			Address: "准格尔旗",
			EndYear: 1981,
		},
		{
			Address:   "苏尼特左旗",
			StartYear: 1982,
		},
	},
	152524: {
		{
			Address: "鄂托克前旗",
			EndYear: 1981,
		},
		{
			Address:   "苏尼特右旗",
			StartYear: 1982,
		},
	},
	152525: {
		{
			Address: "鄂托克旗",
			EndYear: 1981,
		},
		{
			Address:   "东乌珠穆沁旗",
			StartYear: 1982,
		},
	},
	152526: {
		{
			Address: "杭锦旗",
			EndYear: 1981,
		},
		{
			Address:   "西乌珠穆沁旗",
			StartYear: 1982,
		},
	},
	152527: {
		{
			Address: "乌审旗",
			EndYear: 1981,
		},
		{
			Address:   "太仆寺旗",
			StartYear: 1982,
		},
	},
	152528: {
		{
			Address: "伊金霍洛旗",
			EndYear: 1981,
		},
		{
			Address:   "镶黄旗",
			StartYear: 1982,
		},
	},
	152529: {
		{
			Address:   "正镶白旗",
			StartYear: 1982,
		},
	},
	152530: {
		{
			Address:   "正蓝旗",
			StartYear: 1982,
		},
	},
	152531: {
		{
			Address:   "多伦县",
			StartYear: 1982,
		},
	},
	152600: {
		{
			Address: "锡林郭勒盟",
			EndYear: 1981,
		},
		{
			Address:   "乌兰察布盟",
			StartYear: 1982,
			EndYear:   2002,
		},
	},
	152601: {
		{
			Address: "二连浩特市",
			EndYear: 1981,
		},
		{
			Address:   "集宁市",
			StartYear: 1982,
			EndYear:   2002,
		},
	},
	152602: {
		{
			Address:   "丰镇市",
			StartYear: 1990,
			EndYear:   2002,
		},
	},
	152621: {
		{
			Address: "阿巴哈纳尔旗",
			EndYear: 1981,
		},
		{
			Address:   "武川县",
			StartYear: 1982,
			EndYear:   1995,
		},
	},
	152622: {
		{
			Address: "阿巴嘎旗",
			EndYear: 1981,
		},
		{
			Address:   "和林格尔县",
			StartYear: 1982,
			EndYear:   1994,
		},
	},
	152623: {
		{
			Address: "苏尼特左旗",
			EndYear: 1981,
		},
		{
			Address:   "清水河县",
			StartYear: 1982,
			EndYear:   1994,
		},
	},
	152624: {
		{
			Address: "苏尼特右旗",
			EndYear: 1981,
		},
		{
			Address:   "卓资县",
			StartYear: 1982,
			EndYear:   2002,
		},
	},
	152625: {
		{
			Address: "东乌珠穆沁旗",
			EndYear: 1981,
		},
		{
			Address:   "化德县",
			StartYear: 1982,
			EndYear:   2002,
		},
	},
	152626: {
		{
			Address: "西乌珠穆沁旗",
			EndYear: 1981,
		},
		{
			Address:   "商都县",
			StartYear: 1982,
			EndYear:   2002,
		},
	},
	152627: {
		{
			Address: "太仆寺旗",
			EndYear: 1981,
		},
		{
			Address:   "兴和县",
			StartYear: 1982,
			EndYear:   2002,
		},
	},
	152628: {
		{
			Address: "镶黄旗",
			EndYear: 1981,
		},
		{
			Address:   "丰镇县",
			StartYear: 1982,
			EndYear:   1989,
		},
	},
	152629: {
		{
			Address: "正镶白旗",
			EndYear: 1981,
		},
		{
			Address:   "凉城县",
			StartYear: 1982,
			EndYear:   2002,
		},
	},
	152630: {
		{
			Address: "正蓝旗",
			EndYear: 1981,
		},
		{
			Address:   "察哈尔右翼前旗",
			StartYear: 1982,
			EndYear:   2002,
		},
	},
	152631: {
		{
			Address: "多伦县",
			EndYear: 1981,
		},
		{
			Address:   "察哈尔右翼中旗",
			StartYear: 1982,
			EndYear:   2002,
		},
	},
	152632: {
		{
			Address:   "察哈尔右翼后旗",
			StartYear: 1982,
			EndYear:   2002,
		},
	},
	152633: {
		{
			Address:   "达尔罕茂明安联合旗",
			StartYear: 1982,
			EndYear:   1995,
		},
	},
	152634: {
		{
			Address:   "四子王旗",
			StartYear: 1982,
			EndYear:   2002,
		},
	},
	152700: {
		{
			Address: "巴彦淖尔盟",
			EndYear: 1981,
		},
		{
			Address:   "伊克昭盟",
			StartYear: 1982,
			EndYear:   2000,
		},
	},
	152701: {
		{
			Address:   "东胜市",
			StartYear: 1983,
			EndYear:   2000,
		},
	},
	152721: {
		{
			Address: "临河县",
			EndYear: 1981,
		},
		{
			Address:   "东胜县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	152722: {
		{
			Address: "五原县",
			EndYear: 1981,
		},
		{
			Address:   "达拉特旗",
			StartYear: 1982,
			EndYear:   2000,
		},
	},
	152723: {
		{
			Address: "磴口县",
			EndYear: 1981,
		},
		{
			Address:   "准格尔旗",
			StartYear: 1982,
			EndYear:   2000,
		},
	},
	152724: {
		{
			Address: "乌拉特前旗",
			EndYear: 1981,
		},
		{
			Address:   "鄂托克前旗",
			StartYear: 1982,
			EndYear:   2000,
		},
	},
	152725: {
		{
			Address: "乌拉特中后联合旗",
			EndYear: 1980,
		},
		{
			Address:   "乌拉特中旗",
			StartYear: 1981,
			EndYear:   1981,
		},
		{
			Address:   "鄂托克旗",
			StartYear: 1982,
			EndYear:   2000,
		},
	},
	152726: {
		{
			Address: "杭锦后旗",
			EndYear: 1981,
		},
		{
			Address:   "杭锦旗",
			StartYear: 1982,
			EndYear:   2000,
		},
	},
	152727: {
		{
			Address: "潮格旗",
			EndYear: 1980,
		},
		{
			Address:   "乌拉特后旗",
			StartYear: 1981,
			EndYear:   1981,
		},
		{
			Address:   "乌审旗",
			StartYear: 1982,
			EndYear:   2000,
		},
	},
	152728: {
		{
			Address:   "伊金霍洛旗",
			StartYear: 1982,
			EndYear:   2000,
		},
	},
	152800: {
		{
			Address: "乌兰察布盟",
			EndYear: 1981,
		},
		{
			Address:   "巴彦淖尔盟",
			StartYear: 1982,
			EndYear:   2002,
		},
	},
	152801: {
		{
			Address: "集宁市",
			EndYear: 1981,
		},
		{
			Address:   "临河市",
			StartYear: 1984,
			EndYear:   2002,
		},
	},
	152821: {
		{
			Address: "武川县",
			EndYear: 1981,
		},
		{
			Address:   "临河县",
			StartYear: 1982,
			EndYear:   1983,
		},
	},
	152822: {
		{
			Address: "和林格尔县",
			EndYear: 1981,
		},
		{
			Address:   "五原县",
			StartYear: 1982,
			EndYear:   2002,
		},
	},
	152823: {
		{
			Address: "清水河县",
			EndYear: 1981,
		},
		{
			Address:   "磴口县",
			StartYear: 1982,
			EndYear:   2002,
		},
	},
	152824: {
		{
			Address: "卓资县",
			EndYear: 1981,
		},
		{
			Address:   "乌拉特前旗",
			StartYear: 1982,
			EndYear:   2002,
		},
	},
	152825: {
		{
			Address: "化德县",
			EndYear: 1981,
		},
		{
			Address:   "乌拉特中旗",
			StartYear: 1982,
			EndYear:   2002,
		},
	},
	152826: {
		{
			Address: "商都县",
			EndYear: 1981,
		},
		{
			Address:   "乌拉特后旗",
			StartYear: 1982,
			EndYear:   2002,
		},
	},
	152827: {
		{
			Address: "兴和县",
			EndYear: 1981,
		},
		{
			Address:   "杭锦后旗",
			StartYear: 1982,
			EndYear:   2002,
		},
	},
	152828: {
		{
			Address: "丰镇县",
			EndYear: 1981,
		},
	},
	152829: {
		{
			Address: "凉城县",
			EndYear: 1981,
		},
	},
	152830: {
		{
			Address: "察哈尔右翼前旗",
			EndYear: 1981,
		},
	},
	152831: {
		{
			Address: "察哈尔右翼中旗",
			EndYear: 1981,
		},
	},
	152832: {
		{
			Address: "察哈尔右翼后旗",
			EndYear: 1981,
		},
	},
	152833: {
		{
			Address: "达尔罕茂明安联合旗",
			EndYear: 1981,
		},
	},
	152834: {
		{
			Address: "四子王旗",
			EndYear: 1981,
		},
	},
	152900: {
		{
			Address: "阿拉善盟",
		},
	},
	152921: {
		{
			Address: "阿拉善左旗",
		},
	},
	152922: {
		{
			Address: "阿拉善右旗",
		},
	},
	152923: {
		{
			Address: "额济纳旗",
		},
	},
	210000: {
		{
			Address: "辽宁省",
		},
	},
	210100: {
		{
			Address: "沈阳市",
		},
	},
	210102: {
		{
			Address: "和平区",
		},
	},
	210103: {
		{
			Address: "沈河区",
		},
	},
	210104: {
		{
			Address: "大东区",
		},
	},
	210105: {
		{
			Address: "皇姑区",
		},
	},
	210106: {
		{
			Address: "铁西区",
		},
	},
	210111: {
		{
			Address: "苏家屯区",
		},
	},
	210112: {
		{
			Address: "东陵区",
			EndYear: 2015,
		},
		{
			Address:   "浑南区",
			StartYear: 2016,
		},
	},
	210113: {
		{
			Address: "新城子区",
			EndYear: 2005,
		},
		{
			Address:   "沈北新区",
			StartYear: 2006,
		},
	},
	210114: {
		{
			Address: "于洪区",
		},
	},
	210115: {
		{
			Address:   "辽中区",
			StartYear: 2016,
		},
	},
	210121: {
		{
			Address: "新民县",
			EndYear: 1992,
		},
	},
	210122: {
		{
			Address: "辽中县",
			EndYear: 2015,
		},
	},
	210123: {
		{
			Address:   "康平县",
			StartYear: 1992,
		},
	},
	210124: {
		{
			Address:   "法库县",
			StartYear: 1992,
		},
	},
	210181: {
		{
			Address:   "新民市",
			StartYear: 1995,
		},
	},
	210200: {
		{
			Address: "旅大市",
			EndYear: 1980,
		},
		{
			Address:   "大连市",
			StartYear: 1981,
		},
	},
	210202: {
		{
			Address: "中山区",
		},
	},
	210203: {
		{
			Address: "西岗区",
		},
	},
	210204: {
		{
			Address: "沙河口区",
		},
	},
	210211: {
		{
			Address: "甘井子区",
		},
	},
	210212: {
		{
			Address: "旅顺口区",
		},
	},
	210213: {
		{
			Address:   "金州区",
			StartYear: 1987,
		},
	},
	210214: {
		{
			Address:   "普兰店区",
			StartYear: 2015,
		},
	},
	210219: {
		{
			Address:   "瓦房店市",
			StartYear: 1985,
			EndYear:   1985,
		},
	},
	210221: {
		{
			Address: "金县",
			EndYear: 1986,
		},
	},
	210222: {
		{
			Address: "新金县",
			EndYear: 1990,
		},
	},
	210223: {
		{
			Address: "复县",
			EndYear: 1984,
		},
	},
	210224: {
		{
			Address: "长海县",
		},
	},
	210225: {
		{
			Address: "庄河县",
			EndYear: 1991,
		},
	},
	210281: {
		{
			Address:   "瓦房店市",
			StartYear: 1995,
		},
	},
	210282: {
		{
			Address:   "普兰店市",
			StartYear: 1995,
			EndYear:   2014,
		},
	},
	210283: {
		{
			Address:   "庄河市",
			StartYear: 1995,
		},
	},
	210300: {
		{
			Address: "鞍山市",
		},
	},
	210302: {
		{
			Address: "铁东区",
		},
	},
	210303: {
		{
			Address: "铁西区",
		},
	},
	210304: {
		{
			Address: "立山区",
		},
	},
	210311: {
		{
			Address: "郊区",
			EndYear: 1983,
		},
		{
			Address:   "旧堡区",
			StartYear: 1984,
			EndYear:   1995,
		},
		{
			Address:   "千山区",
			StartYear: 1996,
		},
	},
	210319: {
		{
			Address:   "海城市",
			StartYear: 1985,
			EndYear:   1985,
		},
	},
	210321: {
		{
			Address: "台安县",
		},
	},
	210322: {
		{
			Address: "海城县",
			EndYear: 1984,
		},
	},
	210323: {
		{
			Address:   "岫岩满族自治县",
			StartYear: 1992,
		},
	},
	210381: {
		{
			Address:   "海城市",
			StartYear: 1995,
		},
	},
	210400: {
		{
			Address: "抚顺市",
		},
	},
	210402: {
		{
			Address: "新抚区",
		},
	},
	210403: {
		{
			Address: "露天区",
			EndYear: 1998,
		},
		{
			Address:   "东洲区",
			StartYear: 1999,
		},
	},
	210404: {
		{
			Address: "望花区",
		},
	},
	210411: {
		{
			Address: "郊区",
			EndYear: 1987,
		},
		{
			Address:   "顺城区",
			StartYear: 1988,
		},
	},
	210421: {
		{
			Address: "抚顺县",
		},
	},
	210422: {
		{
			Address: "新宾县",
			EndYear: 1984,
		},
		{
			Address:   "新宾满族自治县",
			StartYear: 1985,
		},
	},
	210423: {
		{
			Address: "清原县",
			EndYear: 1988,
		},
		{
			Address:   "清原满族自治县",
			StartYear: 1989,
		},
	},
	210500: {
		{
			Address: "本溪市",
		},
	},
	210502: {
		{
			Address: "平山区",
		},
	},
	210503: {
		{
			Address: "溪湖区",
		},
	},
	210504: {
		{
			Address:   "明山区",
			StartYear: 1984,
		},
	},
	210505: {
		{
			Address:   "南芬区",
			StartYear: 1984,
		},
	},
	210511: {
		{
			Address: "立新区",
			EndYear: 1983,
		},
		{
			Address:   "南芬区",
			StartYear: 1985,
			EndYear:   1985,
		},
	},
	210521: {
		{
			Address: "本溪县",
			EndYear: 1988,
		},
		{
			Address:   "本溪满族自治县",
			StartYear: 1989,
		},
	},
	210522: {
		{
			Address: "桓仁县",
			EndYear: 1988,
		},
		{
			Address:   "桓仁满族自治县",
			StartYear: 1989,
		},
	},
	210600: {
		{
			Address: "丹东市",
		},
	},
	210602: {
		{
			Address: "元宝区",
		},
	},
	210603: {
		{
			Address: "振兴区",
		},
	},
	210604: {
		{
			Address: "振安区",
		},
	},
	210621: {
		{
			Address: "凤城县",
			EndYear: 1984,
		},
		{
			Address:   "凤城满族自治县",
			StartYear: 1985,
			EndYear:   1993,
		},
	},
	210622: {
		{
			Address: "岫岩县",
			EndYear: 1984,
		},
		{
			Address:   "岫岩满族自治县",
			StartYear: 1985,
			EndYear:   1991,
		},
	},
	210623: {
		{
			Address: "东沟县",
			EndYear: 1992,
		},
	},
	210624: {
		{
			Address: "宽甸县",
			EndYear: 1988,
		},
		{
			Address:   "宽甸满族自治县",
			StartYear: 1989,
		},
	},
	210681: {
		{
			Address:   "东港市",
			StartYear: 1995,
		},
	},
	210682: {
		{
			Address:   "凤城市",
			StartYear: 1995,
		},
	},
	210700: {
		{
			Address: "锦州市",
		},
	},
	210702: {
		{
			Address: "古塔区",
		},
	},
	210703: {
		{
			Address: "凌河区",
		},
	},
	210704: {
		{
			Address:   "南票区",
			StartYear: 1982,
			EndYear:   1988,
		},
	},
	210705: {
		{
			Address:   "葫芦岛区",
			StartYear: 1982,
			EndYear:   1988,
		},
	},
	210706: {
		{
			Address:   "太和区",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	210711: {
		{
			Address: "郊区",
			EndYear: 1981,
		},
		{
			Address:   "太和区",
			StartYear: 1983,
		},
	},
	210719: {
		{
			Address:   "锦西市",
			StartYear: 1985,
			EndYear:   1985,
		},
	},
	210721: {
		{
			Address: "锦西县",
			EndYear: 1984,
		},
	},
	210722: {
		{
			Address: "兴城县",
			EndYear: 1985,
		},
	},
	210723: {
		{
			Address: "绥中县",
			EndYear: 1988,
		},
	},
	210724: {
		{
			Address: "锦县",
			EndYear: 1992,
		},
	},
	210725: {
		{
			Address: "北镇县",
			EndYear: 1988,
		},
		{
			Address:   "北镇满族自治县",
			StartYear: 1989,
			EndYear:   1994,
		},
	},
	210726: {
		{
			Address: "黑山县",
		},
	},
	210727: {
		{
			Address: "义县",
		},
	},
	210781: {
		{
			Address:   "凌海市",
			StartYear: 1995,
		},
	},
	210782: {
		{
			Address:   "北宁市",
			StartYear: 1995,
			EndYear:   2005,
		},
		{
			Address:   "北镇市",
			StartYear: 2006,
		},
	},
	210800: {
		{
			Address: "营口市",
		},
	},
	210802: {
		{
			Address: "站前区",
		},
	},
	210803: {
		{
			Address: "西市区",
		},
	},
	210804: {
		{
			Address:   "鲅鱼圈区",
			StartYear: 1984,
		},
	},
	210811: {
		{
			Address: "郊区",
			EndYear: 1983,
		},
		{
			Address:   "老边区",
			StartYear: 1984,
		},
	},
	210821: {
		{
			Address: "营口县",
			EndYear: 1991,
		},
	},
	210822: {
		{
			Address: "盘山县",
			EndYear: 1983,
		},
	},
	210823: {
		{
			Address: "大洼县",
			EndYear: 1983,
		},
	},
	210824: {
		{
			Address: "盖县",
			EndYear: 1991,
		},
	},
	210881: {
		{
			Address:   "盖州市",
			StartYear: 1995,
		},
	},
	210882: {
		{
			Address:   "大石桥市",
			StartYear: 1995,
		},
	},
	210900: {
		{
			Address: "阜新市",
		},
	},
	210902: {
		{
			Address: "海州区",
		},
	},
	210903: {
		{
			Address: "新邱区",
		},
	},
	210904: {
		{
			Address: "太平区",
		},
	},
	210905: {
		{
			Address:   "清河门区",
			StartYear: 1983,
		},
	},
	210911: {
		{
			Address: "郊区",
			EndYear: 1983,
		},
		{
			Address:   "细河区",
			StartYear: 1984,
		},
	},
	210921: {
		{
			Address: "阜新蒙古族自治县",
		},
	},
	210922: {
		{
			Address: "彰武县",
		},
	},
	211000: {
		{
			Address: "辽阳市",
		},
	},
	211002: {
		{
			Address: "白塔区",
		},
	},
	211003: {
		{
			Address: "文圣区",
		},
	},
	211004: {
		{
			Address: "宏伟区",
		},
	},
	211005: {
		{
			Address:   "弓长岭区",
			StartYear: 1984,
		},
	},
	211011: {
		{
			Address: "郊区",
			EndYear: 1983,
		},
		{
			Address:   "太子河区",
			StartYear: 1984,
		},
	},
	211021: {
		{
			Address: "辽阳县",
		},
	},
	211022: {
		{
			Address: "灯塔县",
			EndYear: 1995,
		},
	},
	211081: {
		{
			Address:   "灯塔市",
			StartYear: 1996,
		},
	},
	211100: {
		{
			Address:   "盘锦市",
			StartYear: 1984,
		},
	},
	211102: {
		{
			Address:   "盘山区",
			StartYear: 1984,
			EndYear:   1985,
		},
		{
			Address:   "双台子区",
			StartYear: 1986,
		},
	},
	211103: {
		{
			Address:   "兴隆台区",
			StartYear: 1984,
		},
	},
	211104: {
		{
			Address:   "大洼区",
			StartYear: 2016,
		},
	},
	211111: {
		{
			Address:   "郊区",
			StartYear: 1984,
			EndYear:   1985,
		},
	},
	211121: {
		{
			Address:   "大洼县",
			StartYear: 1984,
			EndYear:   2015,
		},
	},
	211122: {
		{
			Address:   "盘山县",
			StartYear: 1986,
		},
	},
	211200: {
		{
			Address:   "铁岭市",
			StartYear: 1984,
		},
	},
	211202: {
		{
			Address:   "银州区",
			StartYear: 1984,
		},
	},
	211203: {
		{
			Address:   "铁法区",
			StartYear: 1984,
			EndYear:   1985,
		},
	},
	211204: {
		{
			Address:   "清河区",
			StartYear: 1984,
		},
	},
	211221: {
		{
			Address:   "铁岭县",
			StartYear: 1984,
		},
	},
	211222: {
		{
			Address:   "开原县",
			StartYear: 1984,
			EndYear:   1987,
		},
	},
	211223: {
		{
			Address:   "西丰县",
			StartYear: 1984,
		},
	},
	211224: {
		{
			Address:   "昌图县",
			StartYear: 1984,
		},
	},
	211225: {
		{
			Address:   "康平县",
			StartYear: 1984,
			EndYear:   1991,
		},
	},
	211226: {
		{
			Address:   "法库县",
			StartYear: 1984,
			EndYear:   1991,
		},
	},
	211281: {
		{
			Address:   "铁法市",
			StartYear: 1995,
			EndYear:   2001,
		},
		{
			Address:   "调兵山市",
			StartYear: 2002,
		},
	},
	211282: {
		{
			Address:   "开原市",
			StartYear: 1995,
		},
	},
	211300: {
		{
			Address:   "朝阳市",
			StartYear: 1984,
		},
	},
	211302: {
		{
			Address:   "双塔区",
			StartYear: 1984,
		},
	},
	211303: {
		{
			Address:   "龙城区",
			StartYear: 1984,
		},
	},
	211319: {
		{
			Address:   "北票市",
			StartYear: 1985,
			EndYear:   1985,
		},
	},
	211321: {
		{
			Address:   "朝阳县",
			StartYear: 1984,
		},
	},
	211322: {
		{
			Address:   "建平县",
			StartYear: 1984,
		},
	},
	211323: {
		{
			Address:   "凌源县",
			StartYear: 1984,
			EndYear:   1990,
		},
	},
	211324: {
		{
			Address:   "喀喇沁左翼蒙古族自治县",
			StartYear: 1984,
		},
	},
	211325: {
		{
			Address:   "建昌县",
			StartYear: 1984,
			EndYear:   1988,
		},
	},
	211326: {
		{
			Address:   "北票县",
			StartYear: 1984,
			EndYear:   1984,
		},
	},
	211381: {
		{
			Address:   "北票市",
			StartYear: 1995,
		},
	},
	211382: {
		{
			Address:   "凌源市",
			StartYear: 1995,
		},
	},
	211400: {
		{
			Address:   "锦西市",
			StartYear: 1989,
			EndYear:   1993,
		},
		{
			Address:   "葫芦岛市",
			StartYear: 1994,
		},
	},
	211402: {
		{
			Address:   "连山区",
			StartYear: 1989,
		},
	},
	211403: {
		{
			Address:   "龙港区",
			StartYear: 1994,
		},
	},
	211404: {
		{
			Address:   "南票区",
			StartYear: 1989,
		},
	},
	211405: {
		{
			Address:   "葫芦岛区",
			StartYear: 1989,
			EndYear:   1993,
		},
	},
	211421: {
		{
			Address:   "绥中县",
			StartYear: 1989,
		},
	},
	211422: {
		{
			Address:   "建昌县",
			StartYear: 1989,
		},
	},
	211481: {
		{
			Address:   "兴城市",
			StartYear: 1995,
		},
	},
	212100: {
		{
			Address: "铁岭地区",
			EndYear: 1983,
		},
	},
	212101: {
		{
			Address: "铁岭市",
			EndYear: 1983,
		},
	},
	212102: {
		{
			Address:   "铁法市",
			StartYear: 1981,
			EndYear:   1983,
		},
	},
	212121: {
		{
			Address: "铁岭县",
			EndYear: 1983,
		},
	},
	212122: {
		{
			Address: "开原县",
			EndYear: 1983,
		},
	},
	212123: {
		{
			Address: "西丰县",
			EndYear: 1983,
		},
	},
	212124: {
		{
			Address: "昌图县",
			EndYear: 1983,
		},
	},
	212125: {
		{
			Address: "康平县",
			EndYear: 1983,
		},
	},
	212126: {
		{
			Address: "法库县",
			EndYear: 1983,
		},
	},
	212200: {
		{
			Address: "朝阳地区",
			EndYear: 1983,
		},
	},
	212201: {
		{
			Address: "朝阳市",
			EndYear: 1983,
		},
	},
	212221: {
		{
			Address: "朝阳县",
			EndYear: 1983,
		},
	},
	212222: {
		{
			Address: "建平县",
			EndYear: 1983,
		},
	},
	212223: {
		{
			Address: "凌源县",
			EndYear: 1983,
		},
	},
	212224: {
		{
			Address: "喀喇沁左翼蒙古族自治县",
			EndYear: 1983,
		},
	},
	212225: {
		{
			Address: "建昌县",
			EndYear: 1983,
		},
	},
	212226: {
		{
			Address: "北票县",
			EndYear: 1983,
		},
	},
	219001: {
		{
			Address:   "瓦房店市",
			StartYear: 1986,
			EndYear:   1994,
		},
	},
	219002: {
		{
			Address:   "海城市",
			StartYear: 1986,
			EndYear:   1994,
		},
	},
	219003: {
		{
			Address:   "锦西市",
			StartYear: 1986,
			EndYear:   1988,
		},
	},
	219004: {
		{
			Address:   "兴城市",
			StartYear: 1986,
			EndYear:   1994,
		},
	},
	219005: {
		{
			Address:   "铁法市",
			StartYear: 1986,
			EndYear:   1994,
		},
	},
	219006: {
		{
			Address:   "北票市",
			StartYear: 1986,
			EndYear:   1994,
		},
	},
	219007: {
		{
			Address:   "开原市",
			StartYear: 1988,
			EndYear:   1994,
		},
	},
	219008: {
		{
			Address:   "普兰店市",
			StartYear: 1991,
			EndYear:   1994,
		},
	},
	219009: {
		{
			Address:   "凌源市",
			StartYear: 1991,
			EndYear:   1994,
		},
	},
	219010: {
		{
			Address:   "庄河市",
			StartYear: 1992,
			EndYear:   1994,
		},
	},
	219011: {
		{
			Address:   "大石桥市",
			StartYear: 1992,
			EndYear:   1994,
		},
	},
	219012: {
		{
			Address:   "盖州市",
			StartYear: 1992,
			EndYear:   1994,
		},
	},
	219013: {
		{
			Address:   "新民市",
			StartYear: 1993,
			EndYear:   1994,
		},
	},
	219014: {
		{
			Address:   "东港市",
			StartYear: 1993,
			EndYear:   1994,
		},
	},
	219015: {
		{
			Address:   "凌海市",
			StartYear: 1993,
			EndYear:   1994,
		},
	},
	219016: {
		{
			Address:   "凤城市",
			StartYear: 1994,
			EndYear:   1994,
		},
	},
	220000: {
		{
			Address: "吉林省",
		},
	},
	220100: {
		{
			Address: "长春市",
		},
	},
	220102: {
		{
			Address: "南关区",
		},
	},
	220103: {
		{
			Address: "宽城区",
		},
	},
	220104: {
		{
			Address: "朝阳区",
		},
	},
	220105: {
		{
			Address: "二道河子区",
			EndYear: 1994,
		},
		{
			Address:   "二道区",
			StartYear: 1995,
		},
	},
	220106: {
		{
			Address:   "绿园区",
			StartYear: 1995,
		},
	},
	220111: {
		{
			Address: "郊区",
			EndYear: 1994,
		},
	},
	220112: {
		{
			Address:   "双阳区",
			StartYear: 1995,
		},
	},
	220113: {
		{
			Address:   "九台区",
			StartYear: 2014,
		},
	},
	220121: {
		{
			Address: "榆树县",
			EndYear: 1989,
		},
	},
	220122: {
		{
			Address: "农安县",
		},
	},
	220123: {
		{
			Address: "九台县",
			EndYear: 1987,
		},
	},
	220124: {
		{
			Address: "德惠县",
			EndYear: 1993,
		},
	},
	220125: {
		{
			Address: "双阳县",
			EndYear: 1994,
		},
	},
	220181: {
		{
			Address:   "九台市",
			StartYear: 1995,
			EndYear:   2013,
		},
	},
	220182: {
		{
			Address:   "榆树市",
			StartYear: 1995,
		},
	},
	220183: {
		{
			Address:   "德惠市",
			StartYear: 1995,
		},
	},
	220184: {
		{
			Address:   "公主岭市",
			StartYear: 2020,
		},
	},
	220200: {
		{
			Address: "吉林市",
		},
	},
	220202: {
		{
			Address: "昌邑区",
		},
	},
	220203: {
		{
			Address: "龙潭区",
		},
	},
	220204: {
		{
			Address: "船营区",
		},
	},
	220211: {
		{
			Address: "郊区",
			EndYear: 1991,
		},
		{
			Address:   "丰满区",
			StartYear: 1992,
		},
	},
	220221: {
		{
			Address: "永吉县",
		},
	},
	220222: {
		{
			Address: "舒兰县",
			EndYear: 1991,
		},
	},
	220223: {
		{
			Address: "磐石县",
			EndYear: 1994,
		},
	},
	220224: {
		{
			Address: "蛟河县",
			EndYear: 1988,
		},
	},
	220225: {
		{
			Address: "桦甸县",
			EndYear: 1987,
		},
	},
	220281: {
		{
			Address:   "蛟河市",
			StartYear: 1995,
		},
		{
			Address:   "桦甸市",
			StartYear: 2003,
			EndYear:   2003,
		},
	},
	220282: {
		{
			Address:   "桦甸市",
			StartYear: 1995,
		},
		{
			Address:   "蛟河市",
			StartYear: 2003,
			EndYear:   2003,
		},
	},
	220283: {
		{
			Address:   "舒兰市",
			StartYear: 1995,
		},
	},
	220284: {
		{
			Address:   "磐石市",
			StartYear: 1995,
		},
	},
	220300: {
		{
			Address:   "四平市",
			StartYear: 1983,
		},
	},
	220301: {
		{
			Address:   "铁西区",
			StartYear: 1983,
			EndYear:   1983,
		},
	},
	220302: {
		{
			Address:   "铁东区",
			StartYear: 1983,
			EndYear:   1983,
		},
		{
			Address:   "铁西区",
			StartYear: 1984,
		},
	},
	220303: {
		{
			Address:   "铁东区",
			StartYear: 1984,
		},
	},
	220319: {
		{
			Address:   "公主岭市",
			StartYear: 1985,
			EndYear:   1987,
		},
	},
	220321: {
		{
			Address:   "怀德县",
			StartYear: 1983,
			EndYear:   1984,
		},
	},
	220322: {
		{
			Address:   "梨树县",
			StartYear: 1983,
		},
	},
	220323: {
		{
			Address:   "伊通县",
			StartYear: 1983,
			EndYear:   1987,
		},
		{
			Address:   "伊通满族自治县",
			StartYear: 1988,
		},
	},
	220324: {
		{
			Address:   "双辽县",
			StartYear: 1983,
			EndYear:   1995,
		},
	},
	220381: {
		{
			Address:   "公主岭市",
			StartYear: 1995,
			EndYear:   2019,
		},
	},
	220382: {
		{
			Address:   "双辽市",
			StartYear: 1996,
		},
	},
	220400: {
		{
			Address:   "辽源市",
			StartYear: 1983,
		},
	},
	220401: {
		{
			Address:   "龙山区",
			StartYear: 1983,
			EndYear:   1983,
		},
	},
	220402: {
		{
			Address:   "西安区",
			StartYear: 1983,
			EndYear:   1983,
		},
		{
			Address:   "龙山区",
			StartYear: 1984,
		},
	},
	220403: {
		{
			Address:   "西安区",
			StartYear: 1984,
		},
	},
	220421: {
		{
			Address:   "东丰县",
			StartYear: 1983,
		},
	},
	220422: {
		{
			Address:   "东辽县",
			StartYear: 1983,
		},
	},
	220500: {
		{
			Address:   "通化市",
			StartYear: 1985,
		},
	},
	220502: {
		{
			Address:   "东昌区",
			StartYear: 1986,
		},
	},
	220503: {
		{
			Address:   "二道江区",
			StartYear: 1986,
		},
	},
	220519: {
		{
			Address:   "梅河口市",
			StartYear: 1985,
			EndYear:   1987,
		},
	},
	220521: {
		{
			Address:   "通化县",
			StartYear: 1985,
		},
	},
	220522: {
		{
			Address:   "集安县",
			StartYear: 1985,
			EndYear:   1987,
		},
	},
	220523: {
		{
			Address:   "辉南县",
			StartYear: 1985,
		},
	},
	220524: {
		{
			Address:   "柳河县",
			StartYear: 1985,
		},
	},
	220581: {
		{
			Address:   "梅河口市",
			StartYear: 1995,
		},
	},
	220582: {
		{
			Address:   "集安市",
			StartYear: 1995,
		},
	},
	220600: {
		{
			Address:   "浑江市",
			StartYear: 1985,
			EndYear:   1993,
		},
		{
			Address:   "白山市",
			StartYear: 1994,
		},
	},
	220602: {
		{
			Address:   "八道江区",
			StartYear: 1986,
			EndYear:   2009,
		},
		{
			Address:   "浑江区",
			StartYear: 2010,
		},
	},
	220603: {
		{
			Address:   "三岔子区",
			StartYear: 1986,
			EndYear:   1994,
		},
	},
	220604: {
		{
			Address:   "临江区",
			StartYear: 1986,
			EndYear:   1991,
		},
	},
	220605: {
		{
			Address:   "江源区",
			StartYear: 2006,
		},
	},
	220621: {
		{
			Address:   "抚松县",
			StartYear: 1985,
		},
	},
	220622: {
		{
			Address:   "靖宇县",
			StartYear: 1985,
		},
	},
	220623: {
		{
			Address:   "长白朝鲜族自治县",
			StartYear: 1985,
		},
	},
	220624: {
		{
			Address:   "临江县",
			StartYear: 1992,
			EndYear:   1992,
		},
	},
	220625: {
		{
			Address:   "江源县",
			StartYear: 1995,
			EndYear:   2005,
		},
	},
	220681: {
		{
			Address:   "临江市",
			StartYear: 1995,
		},
	},
	220700: {
		{
			Address:   "松原市",
			StartYear: 1992,
		},
	},
	220702: {
		{
			Address:   "扶余区",
			StartYear: 1992,
			EndYear:   1994,
		},
		{
			Address:   "宁江区",
			StartYear: 1995,
		},
	},
	220721: {
		{
			Address:   "前郭尔罗斯蒙古族自治县",
			StartYear: 1992,
		},
	},
	220722: {
		{
			Address:   "长岭县",
			StartYear: 1992,
		},
	},
	220723: {
		{
			Address:   "乾安县",
			StartYear: 1992,
		},
	},
	220724: {
		{
			Address:   "扶余县",
			StartYear: 1995,
			EndYear:   2012,
		},
	},
	220781: {
		{
			Address:   "扶余市",
			StartYear: 2013,
		},
	},
	220800: {
		{
			Address:   "白城市",
			StartYear: 1993,
		},
	},
	220802: {
		{
			Address:   "洮北区",
			StartYear: 1993,
		},
	},
	220821: {
		{
			Address:   "镇赉县",
			StartYear: 1993,
		},
	},
	220822: {
		{
			Address:   "通榆县",
			StartYear: 1993,
		},
	},
	220881: {
		{
			Address:   "洮南市",
			StartYear: 1995,
		},
	},
	220882: {
		{
			Address:   "大安市",
			StartYear: 1995,
		},
	},
	222100: {
		{
			Address: "四平地区",
			EndYear: 1982,
		},
	},
	222101: {
		{
			Address: "四平市",
			EndYear: 1982,
		},
	},
	222102: {
		{
			Address: "辽源市",
			EndYear: 1982,
		},
	},
	222121: {
		{
			Address: "怀德县",
			EndYear: 1982,
		},
	},
	222122: {
		{
			Address: "梨树县",
			EndYear: 1982,
		},
	},
	222123: {
		{
			Address: "伊通县",
			EndYear: 1982,
		},
	},
	222124: {
		{
			Address: "东丰县",
			EndYear: 1982,
		},
	},
	222125: {
		{
			Address: "双辽县",
			EndYear: 1982,
		},
	},
	222200: {
		{
			Address: "通化地区",
			EndYear: 1984,
		},
	},
	222201: {
		{
			Address: "通化市",
			EndYear: 1984,
		},
	},
	222202: {
		{
			Address: "浑江市",
			EndYear: 1984,
		},
	},
	222221: {
		{
			Address: "海龙县",
			EndYear: 1984,
		},
	},
	222222: {
		{
			Address: "通化县",
			EndYear: 1984,
		},
	},
	222223: {
		{
			Address: "柳河县",
			EndYear: 1984,
		},
	},
	222224: {
		{
			Address: "辉南县",
			EndYear: 1984,
		},
	},
	222225: {
		{
			Address: "集安县",
			EndYear: 1984,
		},
	},
	222226: {
		{
			Address: "抚松县",
			EndYear: 1984,
		},
	},
	222227: {
		{
			Address: "靖宇县",
			EndYear: 1984,
		},
	},
	222228: {
		{
			Address: "长白朝鲜族自治县",
			EndYear: 1984,
		},
	},
	222300: {
		{
			Address: "白城地区",
			EndYear: 1992,
		},
	},
	222301: {
		{
			Address: "白城市",
			EndYear: 1992,
		},
	},
	222302: {
		{
			Address:   "洮南市",
			StartYear: 1987,
			EndYear:   1992,
		},
	},
	222303: {
		{
			Address:   "扶余市",
			StartYear: 1987,
			EndYear:   1991,
		},
	},
	222304: {
		{
			Address:   "大安市",
			StartYear: 1988,
			EndYear:   1992,
		},
	},
	222321: {
		{
			Address: "扶余县",
			EndYear: 1986,
		},
	},
	222322: {
		{
			Address: "洮安县",
			EndYear: 1986,
		},
	},
	222323: {
		{
			Address: "长岭县",
			EndYear: 1991,
		},
	},
	222324: {
		{
			Address: "前郭尔罗斯蒙古族自治县",
			EndYear: 1991,
		},
	},
	222325: {
		{
			Address: "大安县",
			EndYear: 1987,
		},
	},
	222326: {
		{
			Address: "镇赉县",
			EndYear: 1992,
		},
	},
	222327: {
		{
			Address: "通榆县",
			EndYear: 1992,
		},
	},
	222328: {
		{
			Address: "乾安县",
			EndYear: 1991,
		},
	},
	222400: {
		{
			Address: "延边朝鲜族自治州",
		},
	},
	222401: {
		{
			Address: "延吉市",
		},
	},
	222402: {
		{
			Address: "图们市",
		},
	},
	222403: {
		{
			Address:   "敦化市",
			StartYear: 1985,
		},
	},
	222404: {
		{
			Address:   "珲春市",
			StartYear: 1988,
		},
	},
	222405: {
		{
			Address:   "龙井市",
			StartYear: 1988,
		},
	},
	222406: {
		{
			Address:   "和龙市",
			StartYear: 1993,
		},
	},
	222421: {
		{
			Address: "延吉县",
			EndYear: 1982,
		},
		{
			Address:   "龙井县",
			StartYear: 1984,
			EndYear:   1987,
		},
	},
	222422: {
		{
			Address: "敦化县",
			EndYear: 1984,
		},
	},
	222423: {
		{
			Address: "和龙县",
			EndYear: 1992,
		},
	},
	222424: {
		{
			Address: "汪清县",
		},
	},
	222425: {
		{
			Address: "珲春县",
			EndYear: 1987,
		},
	},
	222426: {
		{
			Address: "安图县",
		},
	},
	222427: {
		{
			Address:   "龙井县",
			StartYear: 1983,
			EndYear:   1983,
		},
	},
	222500: {
		{
			Address:   "德惠地区",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	222521: {
		{
			Address:   "榆树县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	222522: {
		{
			Address:   "农安县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	222523: {
		{
			Address:   "九台县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	222524: {
		{
			Address:   "德惠县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	222525: {
		{
			Address:   "双阳县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	222600: {
		{
			Address:   "永吉地区",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	222621: {
		{
			Address:   "永吉县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	222622: {
		{
			Address:   "舒兰县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	222623: {
		{
			Address:   "磐石县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	222624: {
		{
			Address:   "蛟河县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	222625: {
		{
			Address:   "桦甸县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	229001: {
		{
			Address:   "公主岭市",
			StartYear: 1988,
			EndYear:   1994,
		},
	},
	229002: {
		{
			Address:   "梅河口市",
			StartYear: 1988,
			EndYear:   1994,
		},
	},
	229003: {
		{
			Address:   "集安市",
			StartYear: 1988,
			EndYear:   1994,
		},
	},
	229004: {
		{
			Address:   "桦甸市",
			StartYear: 1988,
			EndYear:   1994,
		},
	},
	229005: {
		{
			Address:   "九台市",
			StartYear: 1988,
			EndYear:   1994,
		},
	},
	229006: {
		{
			Address:   "蛟河市",
			StartYear: 1989,
			EndYear:   1994,
		},
	},
	229007: {
		{
			Address:   "榆树市",
			StartYear: 1990,
			EndYear:   1994,
		},
	},
	229008: {
		{
			Address:   "舒兰市",
			StartYear: 1992,
			EndYear:   1994,
		},
	},
	229009: {
		{
			Address:   "大安市",
			StartYear: 1993,
			EndYear:   1994,
		},
	},
	229010: {
		{
			Address:   "洮南市",
			StartYear: 1993,
			EndYear:   1994,
		},
	},
	229011: {
		{
			Address:   "临江市",
			StartYear: 1993,
			EndYear:   1994,
		},
	},
	229012: {
		{
			Address:   "德惠市",
			StartYear: 1994,
			EndYear:   1994,
		},
	},
	230000: {
		{
			Address: "黑龙江省",
		},
	},
	230100: {
		{
			Address: "哈尔滨市",
		},
	},
	230102: {
		{
			Address: "道里区",
		},
	},
	230103: {
		{
			Address: "南岗区",
		},
	},
	230104: {
		{
			Address: "道外区",
		},
	},
	230105: {
		{
			Address: "太平区",
			EndYear: 2003,
		},
	},
	230106: {
		{
			Address: "香坊区",
			EndYear: 2003,
		},
	},
	230107: {
		{
			Address: "动力区",
			EndYear: 2005,
		},
	},
	230108: {
		{
			Address: "平房区",
		},
	},
	230109: {
		{
			Address:   "松北区",
			StartYear: 2004,
		},
	},
	230110: {
		{
			Address:   "香坊区",
			StartYear: 2004,
		},
	},
	230111: {
		{
			Address:   "呼兰区",
			StartYear: 2004,
		},
	},
	230112: {
		{
			Address:   "阿城区",
			StartYear: 2006,
		},
	},
	230113: {
		{
			Address:   "双城区",
			StartYear: 2014,
		},
	},
	230121: {
		{
			Address:   "阿城县",
			StartYear: 1983,
			EndYear:   1983,
		},
		{
			Address:   "呼兰县",
			StartYear: 1984,
			EndYear:   2003,
		},
	},
	230122: {
		{
			Address:   "呼兰县",
			StartYear: 1983,
			EndYear:   1983,
		},
		{
			Address:   "阿城县",
			StartYear: 1984,
			EndYear:   1986,
		},
	},
	230123: {
		{
			Address:   "依兰县",
			StartYear: 1991,
		},
	},
	230124: {
		{
			Address:   "方正县",
			StartYear: 1991,
		},
	},
	230125: {
		{
			Address:   "宾县",
			StartYear: 1991,
		},
	},
	230126: {
		{
			Address:   "巴彦县",
			StartYear: 1996,
		},
	},
	230127: {
		{
			Address:   "木兰县",
			StartYear: 1996,
		},
	},
	230128: {
		{
			Address:   "通河县",
			StartYear: 1996,
		},
	},
	230129: {
		{
			Address:   "延寿县",
			StartYear: 1996,
		},
	},
	230181: {
		{
			Address:   "阿城市",
			StartYear: 1995,
			EndYear:   2005,
		},
	},
	230182: {
		{
			Address:   "双城市",
			StartYear: 1996,
			EndYear:   2013,
		},
	},
	230183: {
		{
			Address:   "尚志市",
			StartYear: 1996,
		},
	},
	230184: {
		{
			Address:   "五常市",
			StartYear: 1996,
		},
	},
	230200: {
		{
			Address: "齐齐哈尔市",
		},
	},
	230202: {
		{
			Address: "龙沙区",
		},
	},
	230203: {
		{
			Address: "建华区",
		},
	},
	230204: {
		{
			Address: "铁锋区",
		},
	},
	230205: {
		{
			Address: "昂昂溪区",
		},
	},
	230206: {
		{
			Address: "富拉尔基区",
		},
	},
	230207: {
		{
			Address: "华安区",
			EndYear: 1983,
		},
		{
			Address:   "碾子山区",
			StartYear: 1984,
		},
	},
	230208: {
		{
			Address: "梅里斯区",
			EndYear: 1987,
		},
		{
			Address:   "梅里斯达斡尔族区",
			StartYear: 1988,
		},
	},
	230221: {
		{
			Address:   "龙江县",
			StartYear: 1984,
		},
	},
	230222: {
		{
			Address:   "讷河县",
			StartYear: 1984,
			EndYear:   1991,
		},
	},
	230223: {
		{
			Address:   "依安县",
			StartYear: 1984,
		},
	},
	230224: {
		{
			Address:   "泰来县",
			StartYear: 1984,
		},
	},
	230225: {
		{
			Address:   "甘南县",
			StartYear: 1984,
		},
	},
	230226: {
		{
			Address:   "杜尔伯特蒙古族自治县",
			StartYear: 1984,
			EndYear:   1991,
		},
	},
	230227: {
		{
			Address:   "富裕县",
			StartYear: 1984,
		},
	},
	230228: {
		{
			Address:   "林甸县",
			StartYear: 1984,
			EndYear:   1991,
		},
	},
	230229: {
		{
			Address:   "克山县",
			StartYear: 1984,
		},
	},
	230230: {
		{
			Address:   "克东县",
			StartYear: 1984,
		},
	},
	230231: {
		{
			Address:   "拜泉县",
			StartYear: 1984,
		},
	},
	230281: {
		{
			Address:   "讷河市",
			StartYear: 1995,
		},
	},
	230300: {
		{
			Address: "鸡西市",
		},
	},
	230302: {
		{
			Address: "鸡冠区",
		},
	},
	230303: {
		{
			Address: "恒山区",
		},
	},
	230304: {
		{
			Address: "滴道区",
		},
	},
	230305: {
		{
			Address: "梨树区",
		},
	},
	230306: {
		{
			Address: "城子河区",
		},
	},
	230307: {
		{
			Address: "麻山区",
		},
	},
	230321: {
		{
			Address:   "鸡东县",
			StartYear: 1983,
		},
	},
	230322: {
		{
			Address:   "虎林县",
			StartYear: 1993,
			EndYear:   1995,
		},
	},
	230381: {
		{
			Address:   "虎林市",
			StartYear: 1996,
		},
	},
	230382: {
		{
			Address:   "密山市",
			StartYear: 1995,
		},
	},
	230400: {
		{
			Address: "大庆市",
			EndYear: 1981,
		},
		{
			Address:   "鹤岗市",
			StartYear: 1982,
		},
	},
	230402: {
		{
			Address: "萨尔图区",
			EndYear: 1981,
		},
		{
			Address:   "向阳区",
			StartYear: 1982,
		},
	},
	230403: {
		{
			Address: "龙凤区",
			EndYear: 1981,
		},
		{
			Address:   "工农区",
			StartYear: 1982,
		},
	},
	230404: {
		{
			Address: "让胡路区",
			EndYear: 1981,
		},
		{
			Address:   "南山区",
			StartYear: 1982,
		},
	},
	230405: {
		{
			Address: "红岗区",
			EndYear: 1981,
		},
		{
			Address:   "兴安区",
			StartYear: 1982,
		},
	},
	230406: {
		{
			Address: "大同区",
			EndYear: 1981,
		},
		{
			Address:   "东山区",
			StartYear: 1982,
		},
	},
	230407: {
		{
			Address:   "兴山区",
			StartYear: 1982,
		},
	},
	230421: {
		{
			Address:   "萝北县",
			StartYear: 1987,
		},
	},
	230422: {
		{
			Address:   "绥滨县",
			StartYear: 1987,
		},
	},
	230500: {
		{
			Address: "鹤岗市",
			EndYear: 1981,
		},
		{
			Address:   "双鸭山市",
			StartYear: 1982,
		},
	},
	230502: {
		{
			Address: "向阳区",
			EndYear: 1981,
		},
		{
			Address:   "尖山区",
			StartYear: 1982,
		},
	},
	230503: {
		{
			Address: "工农区",
			EndYear: 1981,
		},
		{
			Address:   "岭东区",
			StartYear: 1982,
		},
	},
	230504: {
		{
			Address: "南山区",
			EndYear: 1981,
		},
		{
			Address:   "岭西区",
			StartYear: 1982,
			EndYear:   1986,
		},
	},
	230505: {
		{
			Address: "兴安区",
			EndYear: 1981,
		},
		{
			Address:   "四方台区",
			StartYear: 1982,
		},
	},
	230506: {
		{
			Address: "东山区",
			EndYear: 1981,
		},
		{
			Address:   "宝山区",
			StartYear: 1982,
		},
	},
	230507: {
		{
			Address: "兴山区",
			EndYear: 1981,
		},
	},
	230521: {
		{
			Address:   "集贤县",
			StartYear: 1987,
		},
	},
	230522: {
		{
			Address:   "友谊县",
			StartYear: 1991,
		},
	},
	230523: {
		{
			Address:   "宝清县",
			StartYear: 1991,
		},
	},
	230524: {
		{
			Address:   "饶河县",
			StartYear: 1993,
		},
	},
	230600: {
		{
			Address: "双鸭山市",
			EndYear: 1981,
		},
		{
			Address:   "大庆市",
			StartYear: 1982,
		},
	},
	230602: {
		{
			Address: "尖山区",
			EndYear: 1981,
		},
		{
			Address:   "萨尔图区",
			StartYear: 1982,
		},
	},
	230603: {
		{
			Address: "岭东区",
			EndYear: 1981,
		},
		{
			Address:   "龙凤区",
			StartYear: 1982,
		},
	},
	230604: {
		{
			Address: "岭西区",
			EndYear: 1981,
		},
		{
			Address:   "让胡路区",
			StartYear: 1982,
		},
	},
	230605: {
		{
			Address: "四方台区",
			EndYear: 1981,
		},
		{
			Address:   "红岗区",
			StartYear: 1982,
		},
	},
	230606: {
		{
			Address: "宝山区",
			EndYear: 1981,
		},
		{
			Address:   "大同区",
			StartYear: 1982,
		},
	},
	230621: {
		{
			Address:   "肇州县",
			StartYear: 1992,
		},
	},
	230622: {
		{
			Address:   "肇源县",
			StartYear: 1992,
		},
	},
	230623: {
		{
			Address:   "林甸县",
			StartYear: 1992,
		},
	},
	230624: {
		{
			Address:   "杜尔伯特蒙古族自治县",
			StartYear: 1992,
		},
	},
	230700: {
		{
			Address: "伊春市",
		},
	},
	230702: {
		{
			Address: "伊春区",
			EndYear: 2018,
		},
	},
	230703: {
		{
			Address: "南岔区",
			EndYear: 2018,
		},
	},
	230704: {
		{
			Address: "友好区",
			EndYear: 2018,
		},
	},
	230705: {
		{
			Address: "西林区",
			EndYear: 2018,
		},
	},
	230706: {
		{
			Address: "翠峦区",
			EndYear: 2018,
		},
	},
	230707: {
		{
			Address: "新青区",
			EndYear: 2018,
		},
	},
	230708: {
		{
			Address: "美溪区",
			EndYear: 2018,
		},
	},
	230709: {
		{
			Address: "大丰区",
			EndYear: 1983,
		},
		{
			Address:   "金山屯区",
			StartYear: 1984,
			EndYear:   2018,
		},
	},
	230710: {
		{
			Address: "五营区",
			EndYear: 2018,
		},
	},
	230711: {
		{
			Address: "乌敏河区",
			EndYear: 1983,
		},
		{
			Address:   "乌马河区",
			StartYear: 1984,
			EndYear:   2018,
		},
	},
	230712: {
		{
			Address: "东风区",
			EndYear: 1983,
		},
		{
			Address:   "汤旺河区",
			StartYear: 1984,
			EndYear:   2018,
		},
	},
	230713: {
		{
			Address: "带岭区",
			EndYear: 2018,
		},
	},
	230714: {
		{
			Address: "乌伊岭区",
			EndYear: 2018,
		},
	},
	230715: {
		{
			Address: "红星区",
			EndYear: 2018,
		},
	},
	230716: {
		{
			Address: "上甘岭区",
			EndYear: 2018,
		},
	},
	230717: {
		{
			Address:   "伊美区",
			StartYear: 2019,
		},
	},
	230718: {
		{
			Address:   "乌翠区",
			StartYear: 2019,
		},
	},
	230719: {
		{
			Address:   "友好区",
			StartYear: 2019,
		},
	},
	230721: {
		{
			Address: "铁力县",
			EndYear: 1987,
		},
	},
	230722: {
		{
			Address: "嘉荫县",
		},
	},
	230723: {
		{
			Address:   "汤旺县",
			StartYear: 2019,
		},
	},
	230724: {
		{
			Address:   "丰林县",
			StartYear: 2019,
		},
	},
	230725: {
		{
			Address:   "大箐山县",
			StartYear: 2019,
		},
	},
	230726: {
		{
			Address:   "南岔县",
			StartYear: 2019,
		},
	},
	230751: {
		{
			Address:   "金林区",
			StartYear: 2019,
		},
	},
	230781: {
		{
			Address:   "铁力市",
			StartYear: 1995,
		},
	},
	230800: {
		{
			Address:   "佳木斯市",
			StartYear: 1983,
		},
	},
	230802: {
		{
			Address:   "永红区",
			StartYear: 1983,
			EndYear:   2005,
		},
	},
	230803: {
		{
			Address:   "向阳区",
			StartYear: 1983,
		},
	},
	230804: {
		{
			Address:   "前进区",
			StartYear: 1983,
		},
	},
	230805: {
		{
			Address:   "东风区",
			StartYear: 1983,
		},
	},
	230811: {
		{
			Address:   "郊区",
			StartYear: 1983,
		},
	},
	230821: {
		{
			Address:   "富锦县",
			StartYear: 1984,
			EndYear:   1987,
		},
	},
	230822: {
		{
			Address:   "桦南县",
			StartYear: 1984,
		},
	},
	230823: {
		{
			Address:   "依兰县",
			StartYear: 1984,
			EndYear:   1990,
		},
	},
	230824: {
		{
			Address:   "友谊县",
			StartYear: 1984,
			EndYear:   1990,
		},
	},
	230825: {
		{
			Address:   "集贤县",
			StartYear: 1984,
			EndYear:   1986,
		},
	},
	230826: {
		{
			Address:   "桦川县",
			StartYear: 1984,
		},
	},
	230827: {
		{
			Address:   "宝清县",
			StartYear: 1984,
			EndYear:   1990,
		},
	},
	230828: {
		{
			Address:   "汤原县",
			StartYear: 1984,
		},
	},
	230829: {
		{
			Address:   "绥滨县",
			StartYear: 1984,
			EndYear:   1986,
		},
	},
	230830: {
		{
			Address:   "萝北县",
			StartYear: 1984,
			EndYear:   1986,
		},
	},
	230831: {
		{
			Address:   "同江县",
			StartYear: 1984,
			EndYear:   1986,
		},
	},
	230832: {
		{
			Address:   "饶河县",
			StartYear: 1984,
			EndYear:   1992,
		},
	},
	230833: {
		{
			Address:   "抚远县",
			StartYear: 1984,
			EndYear:   2015,
		},
	},
	230881: {
		{
			Address:   "同江市",
			StartYear: 1995,
		},
	},
	230882: {
		{
			Address:   "富锦市",
			StartYear: 1995,
		},
	},
	230883: {
		{
			Address:   "抚远市",
			StartYear: 2016,
		},
	},
	230900: {
		{
			Address:   "七台河市",
			StartYear: 1983,
		},
	},
	230902: {
		{
			Address:   "新兴区",
			StartYear: 1984,
		},
	},
	230903: {
		{
			Address:   "桃山区",
			StartYear: 1984,
		},
	},
	230904: {
		{
			Address:   "茄子河区",
			StartYear: 1984,
		},
	},
	230921: {
		{
			Address:   "勃利县",
			StartYear: 1983,
		},
	},
	231000: {
		{
			Address:   "牡丹江市",
			StartYear: 1983,
		},
	},
	231002: {
		{
			Address:   "东安区",
			StartYear: 1983,
		},
	},
	231003: {
		{
			Address:   "阳明区",
			StartYear: 1984,
		},
	},
	231004: {
		{
			Address:   "爱民区",
			StartYear: 1984,
		},
	},
	231005: {
		{
			Address:   "西安区",
			StartYear: 1983,
		},
	},
	231006: {
		{
			Address:   "爱民区",
			StartYear: 1983,
			EndYear:   1983,
		},
	},
	231007: {
		{
			Address:   "阳明区",
			StartYear: 1983,
			EndYear:   1983,
		},
	},
	231011: {
		{
			Address:   "郊区",
			StartYear: 1983,
			EndYear:   1996,
		},
	},
	231020: {
		{
			Address:   "绥芬河市",
			StartYear: 1984,
			EndYear:   1985,
		},
	},
	231021: {
		{
			Address:   "宁安县",
			StartYear: 1983,
			EndYear:   1992,
		},
	},
	231022: {
		{
			Address:   "海林县",
			StartYear: 1983,
			EndYear:   1991,
		},
	},
	231023: {
		{
			Address:   "穆棱县",
			StartYear: 1983,
			EndYear:   1994,
		},
	},
	231024: {
		{
			Address:   "东宁县",
			StartYear: 1983,
			EndYear:   2014,
		},
	},
	231025: {
		{
			Address:   "林口县",
			StartYear: 1983,
		},
	},
	231026: {
		{
			Address:   "密山县",
			StartYear: 1983,
			EndYear:   1987,
		},
	},
	231027: {
		{
			Address:   "虎林县",
			StartYear: 1983,
			EndYear:   1992,
		},
	},
	231081: {
		{
			Address:   "绥芬河市",
			StartYear: 1995,
		},
	},
	231083: {
		{
			Address:   "海林市",
			StartYear: 1995,
		},
	},
	231084: {
		{
			Address:   "宁安市",
			StartYear: 1995,
		},
	},
	231085: {
		{
			Address:   "穆棱市",
			StartYear: 1995,
		},
	},
	231086: {
		{
			Address:   "东宁市",
			StartYear: 2015,
		},
	},
	231100: {
		{
			Address:   "黑河市",
			StartYear: 1993,
		},
	},
	231101: {
		{
			Address:   "绥芬河市",
			StartYear: 1983,
			EndYear:   1983,
		},
	},
	231102: {
		{
			Address:   "爱辉区",
			StartYear: 1993,
		},
	},
	231121: {
		{
			Address:   "嫩江县",
			StartYear: 1993,
			EndYear:   2018,
		},
	},
	231122: {
		{
			Address:   "德都县",
			StartYear: 1993,
			EndYear:   1995,
		},
	},
	231123: {
		{
			Address:   "逊克县",
			StartYear: 1993,
		},
	},
	231124: {
		{
			Address:   "孙吴县",
			StartYear: 1993,
		},
	},
	231181: {
		{
			Address:   "北安市",
			StartYear: 1995,
		},
	},
	231182: {
		{
			Address:   "五大连池市",
			StartYear: 1995,
		},
	},
	231183: {
		{
			Address:   "嫩江市",
			StartYear: 2019,
		},
	},
	231200: {
		{
			Address:   "绥化市",
			StartYear: 1999,
		},
	},
	231202: {
		{
			Address:   "北林区",
			StartYear: 1999,
		},
	},
	231221: {
		{
			Address:   "望奎县",
			StartYear: 1999,
		},
	},
	231222: {
		{
			Address:   "兰西县",
			StartYear: 1999,
		},
	},
	231223: {
		{
			Address:   "青冈县",
			StartYear: 1999,
		},
	},
	231224: {
		{
			Address:   "庆安县",
			StartYear: 1999,
		},
	},
	231225: {
		{
			Address:   "明水县",
			StartYear: 1999,
		},
	},
	231226: {
		{
			Address:   "绥棱县",
			StartYear: 1999,
		},
	},
	231281: {
		{
			Address:   "安达市",
			StartYear: 1999,
		},
	},
	231282: {
		{
			Address:   "肇东市",
			StartYear: 1999,
		},
	},
	231283: {
		{
			Address:   "海伦市",
			StartYear: 1999,
		},
	},
	232100: {
		{
			Address: "绥化地区",
			EndYear: 1981,
		},
		{
			Address:   "松花江地区",
			StartYear: 1982,
			EndYear:   1995,
		},
	},
	232101: {
		{
			Address:   "双城市",
			StartYear: 1988,
			EndYear:   1995,
		},
	},
	232102: {
		{
			Address:   "尚志市",
			StartYear: 1988,
			EndYear:   1995,
		},
	},
	232103: {
		{
			Address:   "五常市",
			StartYear: 1993,
			EndYear:   1995,
		},
	},
	232121: {
		{
			Address: "海伦县",
			EndYear: 1981,
		},
		{
			Address:   "阿城县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	232122: {
		{
			Address: "肇东县",
			EndYear: 1981,
		},
		{
			Address:   "宾县",
			StartYear: 1982,
			EndYear:   1990,
		},
	},
	232123: {
		{
			Address: "绥化县",
			EndYear: 1981,
		},
		{
			Address:   "呼兰县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	232124: {
		{
			Address: "安达县",
			EndYear: 1981,
		},
		{
			Address:   "双城县",
			StartYear: 1982,
			EndYear:   1987,
		},
	},
	232125: {
		{
			Address: "望奎县",
			EndYear: 1981,
		},
		{
			Address:   "五常县",
			StartYear: 1982,
			EndYear:   1992,
		},
	},
	232126: {
		{
			Address: "兰西县",
			EndYear: 1981,
		},
		{
			Address:   "巴彦县",
			StartYear: 1982,
			EndYear:   1995,
		},
	},
	232127: {
		{
			Address: "青冈县",
			EndYear: 1981,
		},
		{
			Address:   "木兰县",
			StartYear: 1982,
			EndYear:   1995,
		},
	},
	232128: {
		{
			Address: "肇源县",
			EndYear: 1981,
		},
		{
			Address:   "通河县",
			StartYear: 1982,
			EndYear:   1995,
		},
	},
	232129: {
		{
			Address: "肇州县",
			EndYear: 1981,
		},
		{
			Address:   "尚志县",
			StartYear: 1982,
			EndYear:   1987,
		},
	},
	232130: {
		{
			Address: "庆安县",
			EndYear: 1981,
		},
		{
			Address:   "方正县",
			StartYear: 1982,
			EndYear:   1990,
		},
	},
	232131: {
		{
			Address: "明水县",
			EndYear: 1981,
		},
		{
			Address:   "延寿县",
			StartYear: 1982,
			EndYear:   1995,
		},
	},
	232132: {
		{
			Address: "绥棱县",
			EndYear: 1981,
		},
	},
	232200: {
		{
			Address: "松花江地区",
			EndYear: 1981,
		},
		{
			Address:   "嫩江地区",
			StartYear: 1982,
			EndYear:   1983,
		},
	},
	232221: {
		{
			Address: "五常县",
			EndYear: 1981,
		},
		{
			Address:   "龙江县",
			StartYear: 1982,
			EndYear:   1983,
		},
	},
	232222: {
		{
			Address: "双城县",
			EndYear: 1981,
		},
		{
			Address:   "讷河县",
			StartYear: 1982,
			EndYear:   1983,
		},
	},
	232223: {
		{
			Address: "巴彦县",
			EndYear: 1981,
		},
		{
			Address:   "依安县",
			StartYear: 1982,
			EndYear:   1983,
		},
	},
	232224: {
		{
			Address: "呼兰县",
			EndYear: 1981,
		},
		{
			Address:   "泰来县",
			StartYear: 1982,
			EndYear:   1983,
		},
	},
	232225: {
		{
			Address: "宾县",
			EndYear: 1981,
		},
		{
			Address:   "甘南县",
			StartYear: 1982,
			EndYear:   1983,
		},
	},
	232226: {
		{
			Address: "阿城县",
			EndYear: 1981,
		},
		{
			Address:   "杜尔伯特蒙古族自治县",
			StartYear: 1982,
			EndYear:   1983,
		},
	},
	232227: {
		{
			Address: "尚志县",
			EndYear: 1981,
		},
		{
			Address:   "富裕县",
			StartYear: 1982,
			EndYear:   1983,
		},
	},
	232228: {
		{
			Address: "木兰县",
			EndYear: 1981,
		},
		{
			Address:   "林甸县",
			StartYear: 1982,
			EndYear:   1983,
		},
	},
	232229: {
		{
			Address: "延寿县",
			EndYear: 1981,
		},
		{
			Address:   "克山县",
			StartYear: 1982,
			EndYear:   1983,
		},
	},
	232230: {
		{
			Address: "通河县",
			EndYear: 1981,
		},
		{
			Address:   "克东县",
			StartYear: 1982,
			EndYear:   1983,
		},
	},
	232231: {
		{
			Address: "方正县",
			EndYear: 1981,
		},
		{
			Address:   "拜泉县",
			StartYear: 1982,
			EndYear:   1983,
		},
	},
	232300: {
		{
			Address: "嫩江地区",
			EndYear: 1981,
		},
		{
			Address:   "绥化地区",
			StartYear: 1982,
			EndYear:   1998,
		},
	},
	232301: {
		{
			Address:   "绥化市",
			StartYear: 1982,
			EndYear:   1998,
		},
	},
	232302: {
		{
			Address:   "安达市",
			StartYear: 1984,
			EndYear:   1998,
		},
	},
	232303: {
		{
			Address:   "肇东市",
			StartYear: 1986,
			EndYear:   1998,
		},
	},
	232304: {
		{
			Address:   "海伦市",
			StartYear: 1989,
			EndYear:   1998,
		},
	},
	232321: {
		{
			Address: "讷河县",
			EndYear: 1981,
		},
		{
			Address:   "海伦县",
			StartYear: 1982,
			EndYear:   1988,
		},
	},
	232322: {
		{
			Address: "拜泉县",
			EndYear: 1981,
		},
		{
			Address:   "肇东县",
			StartYear: 1982,
			EndYear:   1985,
		},
	},
	232323: {
		{
			Address: "龙江县",
			EndYear: 1981,
		},
	},
	232324: {
		{
			Address: "依安县",
			EndYear: 1981,
		},
		{
			Address:   "望奎县",
			StartYear: 1982,
			EndYear:   1998,
		},
	},
	232325: {
		{
			Address: "克山县",
			EndYear: 1981,
		},
		{
			Address:   "兰西县",
			StartYear: 1982,
			EndYear:   1998,
		},
	},
	232326: {
		{
			Address: "甘南县",
			EndYear: 1981,
		},
		{
			Address:   "青冈县",
			StartYear: 1982,
			EndYear:   1998,
		},
	},
	232327: {
		{
			Address: "泰来县",
			EndYear: 1981,
		},
		{
			Address:   "安达县",
			StartYear: 1982,
			EndYear:   1983,
		},
	},
	232328: {
		{
			Address: "克东县",
			EndYear: 1981,
		},
		{
			Address:   "肇源县",
			StartYear: 1982,
			EndYear:   1991,
		},
	},
	232329: {
		{
			Address: "富裕县",
			EndYear: 1981,
		},
		{
			Address:   "肇州县",
			StartYear: 1982,
			EndYear:   1991,
		},
	},
	232330: {
		{
			Address: "林甸县",
			EndYear: 1981,
		},
		{
			Address:   "庆安县",
			StartYear: 1982,
			EndYear:   1998,
		},
	},
	232331: {
		{
			Address: "杜尔伯特蒙古族自治县",
			EndYear: 1981,
		},
		{
			Address:   "明水县",
			StartYear: 1982,
			EndYear:   1998,
		},
	},
	232332: {
		{
			Address:   "绥棱县",
			StartYear: 1982,
			EndYear:   1998,
		},
	},
	232400: {
		{
			Address: "合江地区",
			EndYear: 1983,
		},
	},
	232401: {
		{
			Address: "佳木斯市",
			EndYear: 1982,
		},
	},
	232402: {
		{
			Address: "七台河市",
			EndYear: 1982,
		},
	},
	232403: {
		{
			Address: "永红区",
			EndYear: 1982,
		},
	},
	232404: {
		{
			Address: "向阳区",
			EndYear: 1982,
		},
	},
	232405: {
		{
			Address: "前进区",
			EndYear: 1982,
		},
	},
	232406: {
		{
			Address: "东风区",
			EndYear: 1982,
		},
	},
	232411: {
		{
			Address: "郊区",
			EndYear: 1982,
		},
	},
	232421: {
		{
			Address: "桦南县",
			EndYear: 1981,
		},
		{
			Address:   "富锦县",
			StartYear: 1982,
			EndYear:   1983,
		},
	},
	232422: {
		{
			Address: "集贤县",
			EndYear: 1981,
		},
		{
			Address:   "桦南县",
			StartYear: 1982,
			EndYear:   1983,
		},
	},
	232423: {
		{
			Address: "宝清县",
			EndYear: 1981,
		},
		{
			Address:   "依兰县",
			StartYear: 1982,
			EndYear:   1983,
		},
	},
	232424: {
		{
			Address: "富锦县",
			EndYear: 1981,
		},
		{
			Address:   "勃利县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	232425: {
		{
			Address: "依兰县",
			EndYear: 1981,
		},
		{
			Address:   "集贤县",
			StartYear: 1982,
			EndYear:   1983,
		},
	},
	232426: {
		{
			Address: "勃利县",
			EndYear: 1981,
		},
		{
			Address:   "桦川县",
			StartYear: 1982,
			EndYear:   1983,
		},
	},
	232427: {
		{
			Address: "汤原县",
			EndYear: 1981,
		},
		{
			Address:   "宝清县",
			StartYear: 1982,
			EndYear:   1983,
		},
	},
	232428: {
		{
			Address: "桦川县",
			EndYear: 1981,
		},
		{
			Address:   "汤原县",
			StartYear: 1982,
			EndYear:   1983,
		},
	},
	232429: {
		{
			Address: "萝北县",
			EndYear: 1981,
		},
		{
			Address:   "绥滨县",
			StartYear: 1982,
			EndYear:   1983,
		},
	},
	232430: {
		{
			Address: "绥滨县",
			EndYear: 1981,
		},
		{
			Address:   "萝北县",
			StartYear: 1982,
			EndYear:   1983,
		},
	},
	232431: {
		{
			Address: "饶河县",
			EndYear: 1981,
		},
		{
			Address:   "同江县",
			StartYear: 1982,
			EndYear:   1983,
		},
	},
	232432: {
		{
			Address: "同江县",
			EndYear: 1981,
		},
		{
			Address:   "饶河县",
			StartYear: 1982,
			EndYear:   1983,
		},
	},
	232433: {
		{
			Address: "抚远县",
			EndYear: 1983,
		},
	},
	232500: {
		{
			Address: "牡丹江地区",
			EndYear: 1982,
		},
	},
	232501: {
		{
			Address: "牡丹江市",
			EndYear: 1982,
		},
	},
	232502: {
		{
			Address: "东凤区",
			EndYear: 1981,
		},
		{
			Address:   "绥芬河市",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	232503: {
		{
			Address: "先锋区",
			EndYear: 1981,
		},
		{
			Address:   "东凤区",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	232504: {
		{
			Address: "爱民区",
			EndYear: 1981,
		},
		{
			Address:   "先锋区",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	232505: {
		{
			Address: "阳明区",
			EndYear: 1981,
		},
		{
			Address:   "爱民区",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	232506: {
		{
			Address:   "阳明区",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	232511: {
		{
			Address: "郊区",
			EndYear: 1982,
		},
	},
	232521: {
		{
			Address: "海林县",
			EndYear: 1981,
		},
		{
			Address:   "宁安县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	232522: {
		{
			Address: "宁安县",
			EndYear: 1981,
		},
		{
			Address:   "海林县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	232523: {
		{
			Address: "林口县",
			EndYear: 1981,
		},
		{
			Address:   "穆棱县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	232524: {
		{
			Address: "密山县",
			EndYear: 1981,
		},
		{
			Address:   "东宁县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	232525: {
		{
			Address: "穆棱县",
			EndYear: 1981,
		},
		{
			Address:   "林口县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	232526: {
		{
			Address: "虎林县",
			EndYear: 1981,
		},
		{
			Address:   "鸡东县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	232527: {
		{
			Address: "鸡东县",
			EndYear: 1981,
		},
		{
			Address:   "密山县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	232528: {
		{
			Address: "东宁县",
			EndYear: 1981,
		},
		{
			Address:   "虎林县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	232581: {
		{
			Address: "绥芬河市",
			EndYear: 1981,
		},
	},
	232600: {
		{
			Address: "黑河地区",
			EndYear: 1992,
		},
	},
	232601: {
		{
			Address: "黑河市",
			EndYear: 1992,
		},
	},
	232602: {
		{
			Address:   "北安市",
			StartYear: 1982,
			EndYear:   1992,
		},
	},
	232603: {
		{
			Address:   "五大连池市",
			StartYear: 1983,
			EndYear:   1992,
		},
	},
	232621: {
		{
			Address: "嫩江县",
			EndYear: 1981,
		},
	},
	232622: {
		{
			Address: "北安县",
			EndYear: 1981,
		},
		{
			Address:   "嫩江县",
			StartYear: 1982,
			EndYear:   1992,
		},
	},
	232623: {
		{
			Address: "德都县",
			EndYear: 1992,
		},
	},
	232624: {
		{
			Address: "爱辉县",
			EndYear: 1982,
		},
	},
	232625: {
		{
			Address: "逊克县",
			EndYear: 1992,
		},
	},
	232626: {
		{
			Address: "孙吴县",
			EndYear: 1992,
		},
	},
	232627: {
		{
			Address:   "通北县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	232700: {
		{
			Address: "大兴安岭地区",
		},
	},
	232701: {
		{
			Address:   "漠河市",
			StartYear: 2018,
		},
	},
	232721: {
		{
			Address: "呼玛县",
		},
	},
	232722: {
		{
			Address:   "塔河县",
			StartYear: 1981,
		},
	},
	232723: {
		{
			Address:   "漠河县",
			StartYear: 1981,
			EndYear:   2017,
		},
	},
	239001: {
		{
			Address:   "绥芬河市",
			StartYear: 1986,
			EndYear:   1994,
		},
	},
	239002: {
		{
			Address:   "镜泊湖市",
			StartYear: 1986,
			EndYear:   1986,
		},
		{
			Address:   "阿城市",
			StartYear: 1987,
			EndYear:   1994,
		},
	},
	239003: {
		{
			Address:   "同江市",
			StartYear: 1987,
			EndYear:   1994,
		},
	},
	239004: {
		{
			Address:   "富锦市",
			StartYear: 1988,
			EndYear:   1994,
		},
	},
	239005: {
		{
			Address:   "铁力市",
			StartYear: 1988,
			EndYear:   1994,
		},
	},
	239006: {
		{
			Address:   "密山市",
			StartYear: 1988,
			EndYear:   1994,
		},
	},
	239007: {
		{
			Address:   "海林市",
			StartYear: 1992,
			EndYear:   1994,
		},
	},
	239008: {
		{
			Address:   "讷河市",
			StartYear: 1992,
			EndYear:   1994,
		},
	},
	239009: {
		{
			Address:   "北安市",
			StartYear: 1993,
			EndYear:   1994,
		},
	},
	239010: {
		{
			Address:   "五大连池市",
			StartYear: 1993,
			EndYear:   1994,
		},
	},
	239011: {
		{
			Address:   "宁安市",
			StartYear: 1993,
			EndYear:   1994,
		},
	},
	310000: {
		{
			Address: "上海市",
		},
	},
	310101: {
		{
			Address: "黄浦区",
		},
	},
	310102: {
		{
			Address: "南市区",
			EndYear: 1999,
		},
	},
	310103: {
		{
			Address: "卢湾区",
			EndYear: 2010,
		},
	},
	310104: {
		{
			Address: "徐汇区",
		},
	},
	310105: {
		{
			Address: "长宁区",
		},
	},
	310106: {
		{
			Address: "静安区",
		},
	},
	310107: {
		{
			Address: "普陀区",
		},
	},
	310108: {
		{
			Address: "闸北区",
			EndYear: 2014,
		},
	},
	310109: {
		{
			Address: "虹口区",
		},
	},
	310110: {
		{
			Address: "杨浦区",
		},
	},
	310111: {
		{
			Address: "吴淞区",
			EndYear: 1987,
		},
	},
	310112: {
		{
			Address:   "闵行区",
			StartYear: 1981,
		},
	},
	310113: {
		{
			Address:   "宝山区",
			StartYear: 1988,
		},
	},
	310114: {
		{
			Address:   "嘉定区",
			StartYear: 1992,
		},
	},
	310115: {
		{
			Address:   "浦东新区",
			StartYear: 1992,
		},
	},
	310116: {
		{
			Address:   "金山区",
			StartYear: 1997,
		},
	},
	310117: {
		{
			Address:   "松江区",
			StartYear: 1998,
		},
	},
	310118: {
		{
			Address:   "青浦区",
			StartYear: 1999,
		},
	},
	310119: {
		{
			Address:   "南汇区",
			StartYear: 2001,
			EndYear:   2008,
		},
	},
	310120: {
		{
			Address:   "奉贤区",
			StartYear: 2001,
		},
	},
	310151: {
		{
			Address:   "崇明区",
			StartYear: 2016,
		},
	},
	310201: {
		{
			Address: "上海县",
			EndYear: 1981,
		},
	},
	310202: {
		{
			Address: "嘉定县",
			EndYear: 1981,
		},
	},
	310203: {
		{
			Address: "宝山县",
			EndYear: 1981,
		},
	},
	310204: {
		{
			Address: "川沙县",
			EndYear: 1981,
		},
	},
	310205: {
		{
			Address: "南汇县",
			EndYear: 1981,
		},
	},
	310206: {
		{
			Address: "奉贤县",
			EndYear: 1981,
		},
	},
	310207: {
		{
			Address: "松江县",
			EndYear: 1981,
		},
	},
	310208: {
		{
			Address: "金山县",
			EndYear: 1981,
		},
	},
	310209: {
		{
			Address: "青浦县",
			EndYear: 1981,
		},
	},
	310210: {
		{
			Address: "崇明县",
			EndYear: 1981,
		},
	},
	310221: {
		{
			Address:   "上海县",
			StartYear: 1982,
			EndYear:   1991,
		},
	},
	310222: {
		{
			Address:   "嘉定县",
			StartYear: 1982,
			EndYear:   1991,
		},
	},
	310223: {
		{
			Address:   "宝山县",
			StartYear: 1982,
			EndYear:   1987,
		},
	},
	310224: {
		{
			Address:   "川沙县",
			StartYear: 1982,
			EndYear:   1991,
		},
	},
	310225: {
		{
			Address:   "南汇县",
			StartYear: 1982,
			EndYear:   2000,
		},
	},
	310226: {
		{
			Address:   "奉贤县",
			StartYear: 1982,
			EndYear:   2000,
		},
	},
	310227: {
		{
			Address:   "松江县",
			StartYear: 1982,
			EndYear:   1997,
		},
	},
	310228: {
		{
			Address:   "金山县",
			StartYear: 1982,
			EndYear:   1996,
		},
	},
	310229: {
		{
			Address:   "青浦县",
			StartYear: 1982,
			EndYear:   1998,
		},
	},
	310230: {
		{
			Address:   "崇明县",
			StartYear: 1982,
			EndYear:   2015,
		},
	},
	320000: {
		{
			Address: "江苏省",
		},
	},
	320100: {
		{
			Address: "南京市",
		},
	},
	320102: {
		{
			Address: "玄武区",
		},
	},
	320103: {
		{
			Address: "白下区",
			EndYear: 2012,
		},
	},
	320104: {
		{
			Address: "秦淮区",
		},
	},
	320105: {
		{
			Address: "建邺区",
		},
	},
	320106: {
		{
			Address: "鼓楼区",
		},
	},
	320107: {
		{
			Address: "下关区",
			EndYear: 2012,
		},
	},
	320111: {
		{
			Address: "浦口区",
		},
	},
	320112: {
		{
			Address: "大厂区",
			EndYear: 2001,
		},
	},
	320113: {
		{
			Address: "栖霞区",
		},
	},
	320114: {
		{
			Address: "雨花区",
			EndYear: 1983,
		},
		{
			Address:   "雨花台区",
			StartYear: 1984,
		},
	},
	320115: {
		{
			Address:   "江宁区",
			StartYear: 2000,
		},
	},
	320116: {
		{
			Address:   "六合区",
			StartYear: 2002,
		},
	},
	320117: {
		{
			Address:   "溧水区",
			StartYear: 2013,
		},
	},
	320118: {
		{
			Address:   "高淳区",
			StartYear: 2013,
		},
	},
	320121: {
		{
			Address: "江宁县",
			EndYear: 1999,
		},
	},
	320122: {
		{
			Address: "江浦县",
			EndYear: 2001,
		},
	},
	320123: {
		{
			Address: "六合县",
			EndYear: 2001,
		},
	},
	320124: {
		{
			Address:   "溧水县",
			StartYear: 1983,
			EndYear:   2012,
		},
	},
	320125: {
		{
			Address:   "高淳县",
			StartYear: 1983,
			EndYear:   2012,
		},
	},
	320200: {
		{
			Address: "无锡市",
		},
	},
	320202: {
		{
			Address: "崇安区",
			EndYear: 2014,
		},
	},
	320203: {
		{
			Address: "南长区",
			EndYear: 2014,
		},
	},
	320204: {
		{
			Address: "北塘区",
			EndYear: 2014,
		},
	},
	320205: {
		{
			Address:   "马山区",
			StartYear: 1987,
			EndYear:   1999,
		},
		{
			Address:   "锡山区",
			StartYear: 2000,
		},
	},
	320206: {
		{
			Address:   "惠山区",
			StartYear: 2000,
		},
	},
	320211: {
		{
			Address:   "郊区",
			StartYear: 1983,
			EndYear:   1999,
		},
		{
			Address:   "滨湖区",
			StartYear: 2000,
		},
	},
	320213: {
		{
			Address:   "梁溪区",
			StartYear: 2015,
		},
	},
	320214: {
		{
			Address:   "新吴区",
			StartYear: 2015,
		},
	},
	320221: {
		{
			Address:   "江阴县",
			StartYear: 1983,
			EndYear:   1986,
		},
	},
	320222: {
		{
			Address:   "无锡县",
			StartYear: 1983,
			EndYear:   1994,
		},
	},
	320223: {
		{
			Address:   "宜兴县",
			StartYear: 1983,
			EndYear:   1987,
		},
	},
	320281: {
		{
			Address:   "江阴市",
			StartYear: 1995,
		},
	},
	320282: {
		{
			Address:   "宜兴市",
			StartYear: 1995,
		},
	},
	320283: {
		{
			Address:   "锡山市",
			StartYear: 1995,
			EndYear:   1999,
		},
	},
	320300: {
		{
			Address: "徐州市",
		},
	},
	320302: {
		{
			Address: "鼓楼区",
		},
	},
	320303: {
		{
			Address: "云龙区",
		},
	},
	320304: {
		{
			Address: "矿区",
			EndYear: 1994,
		},
		{
			Address:   "九里区",
			StartYear: 1995,
			EndYear:   2009,
		},
	},
	320305: {
		{
			Address: "贾汪区",
		},
	},
	320311: {
		{
			Address: "郊区",
			EndYear: 1992,
		},
		{
			Address:   "泉山区",
			StartYear: 1993,
		},
	},
	320312: {
		{
			Address:   "铜山区",
			StartYear: 2010,
		},
	},
	320321: {
		{
			Address:   "丰县",
			StartYear: 1983,
		},
	},
	320322: {
		{
			Address:   "沛县",
			StartYear: 1983,
		},
	},
	320323: {
		{
			Address:   "铜山县",
			StartYear: 1983,
			EndYear:   2009,
		},
	},
	320324: {
		{
			Address:   "睢宁县",
			StartYear: 1983,
		},
	},
	320325: {
		{
			Address:   "邳县",
			StartYear: 1983,
			EndYear:   1991,
		},
	},
	320326: {
		{
			Address:   "新沂县",
			StartYear: 1983,
			EndYear:   1989,
		},
	},
	320381: {
		{
			Address:   "新沂市",
			StartYear: 1995,
		},
	},
	320382: {
		{
			Address:   "邳州市",
			StartYear: 1995,
		},
	},
	320400: {
		{
			Address: "常州市",
		},
	},
	320402: {
		{
			Address: "天宁区",
		},
	},
	320403: {
		{
			Address: "广化区",
			EndYear: 1985,
		},
	},
	320404: {
		{
			Address: "钟楼区",
		},
	},
	320405: {
		{
			Address: "戚墅堰区",
			EndYear: 2014,
		},
	},
	320411: {
		{
			Address:   "郊区",
			StartYear: 1983,
			EndYear:   2001,
		},
		{
			Address:   "新北区",
			StartYear: 2002,
		},
	},
	320412: {
		{
			Address:   "武进区",
			StartYear: 2002,
		},
	},
	320413: {
		{
			Address:   "金坛区",
			StartYear: 2015,
		},
	},
	320421: {
		{
			Address:   "武进县",
			StartYear: 1983,
			EndYear:   1994,
		},
	},
	320422: {
		{
			Address:   "金坛县",
			StartYear: 1983,
			EndYear:   1992,
		},
	},
	320423: {
		{
			Address:   "溧阳县",
			StartYear: 1983,
			EndYear:   1989,
		},
	},
	320481: {
		{
			Address:   "溧阳市",
			StartYear: 1995,
		},
	},
	320482: {
		{
			Address:   "金坛市",
			StartYear: 1995,
			EndYear:   2014,
		},
	},
	320483: {
		{
			Address:   "武进市",
			StartYear: 1995,
			EndYear:   2001,
		},
	},
	320500: {
		{
			Address: "苏州市",
		},
	},
	320502: {
		{
			Address: "沧浪区",
			EndYear: 2011,
		},
	},
	320503: {
		{
			Address: "平江区",
			EndYear: 2011,
		},
	},
	320504: {
		{
			Address: "金阊区",
			EndYear: 2011,
		},
	},
	320505: {
		{
			Address:   "虎丘区",
			StartYear: 2000,
		},
	},
	320506: {
		{
			Address:   "吴中区",
			StartYear: 2000,
		},
	},
	320507: {
		{
			Address:   "相城区",
			StartYear: 2000,
		},
	},
	320508: {
		{
			Address:   "姑苏区",
			StartYear: 2012,
		},
	},
	320509: {
		{
			Address:   "吴江区",
			StartYear: 2012,
		},
	},
	320511: {
		{
			Address:   "郊区",
			StartYear: 1983,
			EndYear:   1999,
		},
	},
	320521: {
		{
			Address:   "沙洲县",
			StartYear: 1983,
			EndYear:   1985,
		},
	},
	320522: {
		{
			Address:   "太仓县",
			StartYear: 1983,
			EndYear:   1992,
		},
	},
	320523: {
		{
			Address:   "昆山县",
			StartYear: 1983,
			EndYear:   1988,
		},
	},
	320524: {
		{
			Address:   "吴县",
			StartYear: 1983,
			EndYear:   1994,
		},
	},
	320525: {
		{
			Address:   "吴江县",
			StartYear: 1983,
			EndYear:   1991,
		},
	},
	320581: {
		{
			Address:   "常熟市",
			StartYear: 1995,
		},
	},
	320582: {
		{
			Address:   "张家港市",
			StartYear: 1995,
		},
	},
	320583: {
		{
			Address:   "昆山市",
			StartYear: 1995,
		},
	},
	320584: {
		{
			Address:   "吴江市",
			StartYear: 1995,
			EndYear:   2011,
		},
	},
	320585: {
		{
			Address:   "太仓市",
			StartYear: 1995,
		},
	},
	320586: {
		{
			Address:   "吴县市",
			StartYear: 1995,
			EndYear:   1999,
		},
	},
	320600: {
		{
			Address: "南通市",
		},
	},
	320602: {
		{
			Address: "城中区",
			EndYear: 1982,
		},
		{
			Address:   "城区",
			StartYear: 1983,
			EndYear:   1990,
		},
		{
			Address:   "崇川区",
			StartYear: 1991,
			EndYear:   2019,
		},
	},
	320603: {
		{
			Address: "港闸区",
			EndYear: 1982,
		},
	},
	320611: {
		{
			Address: "郊区",
			EndYear: 1990,
		},
		{
			Address:   "港闸区",
			StartYear: 1991,
			EndYear:   2019,
		},
	},
	320612: {
		{
			Address:   "通州区",
			StartYear: 2009,
		},
	},
	320613: {
		{
			Address:   "崇川区",
			StartYear: 2020,
		},
	},
	320614: {
		{
			Address:   "海门区",
			StartYear: 2020,
		},
	},
	320621: {
		{
			Address:   "海安县",
			StartYear: 1983,
			EndYear:   2017,
		},
	},
	320622: {
		{
			Address:   "如皋县",
			StartYear: 1983,
			EndYear:   1990,
		},
	},
	320623: {
		{
			Address:   "如东县",
			StartYear: 1983,
		},
	},
	320624: {
		{
			Address:   "南通县",
			StartYear: 1983,
			EndYear:   1992,
		},
	},
	320625: {
		{
			Address:   "海门县",
			StartYear: 1983,
			EndYear:   1993,
		},
	},
	320626: {
		{
			Address:   "启东县",
			StartYear: 1983,
			EndYear:   1988,
		},
	},
	320681: {
		{
			Address:   "启东市",
			StartYear: 1995,
		},
	},
	320682: {
		{
			Address:   "如皋市",
			StartYear: 1995,
		},
	},
	320683: {
		{
			Address:   "通州市",
			StartYear: 1995,
			EndYear:   2008,
		},
	},
	320684: {
		{
			Address:   "海门市",
			StartYear: 1995,
			EndYear:   2019,
		},
	},
	320685: {
		{
			Address:   "海安市",
			StartYear: 2018,
		},
	},
	320700: {
		{
			Address: "连云港市",
		},
	},
	320702: {
		{
			Address: "连云港区",
			EndYear: 1982,
		},
		{
			Address:   "新海区",
			StartYear: 1983,
			EndYear:   1985,
		},
	},
	320703: {
		{
			Address: "盐区",
			EndYear: 1982,
		},
		{
			Address:   "连云区",
			StartYear: 1983,
		},
	},
	320704: {
		{
			Address: "新浦区",
			EndYear: 1982,
		},
		{
			Address:   "云台区",
			StartYear: 1983,
			EndYear:   2000,
		},
	},
	320705: {
		{
			Address: "海州区",
			EndYear: 1982,
		},
		{
			Address:   "新浦区",
			StartYear: 1986,
			EndYear:   2013,
		},
	},
	320706: {
		{
			Address:   "海州区",
			StartYear: 1986,
		},
	},
	320707: {
		{
			Address:   "赣榆区",
			StartYear: 2014,
		},
	},
	320711: {
		{
			Address:   "郊区",
			StartYear: 1981,
			EndYear:   1982,
		},
	},
	320721: {
		{
			Address:   "赣榆县",
			StartYear: 1983,
			EndYear:   2013,
		},
	},
	320722: {
		{
			Address:   "东海县",
			StartYear: 1983,
		},
	},
	320723: {
		{
			Address:   "灌云县",
			StartYear: 1983,
		},
	},
	320724: {
		{
			Address:   "灌南县",
			StartYear: 1996,
		},
	},
	320800: {
		{
			Address:   "淮阴市",
			StartYear: 1983,
			EndYear:   1999,
		},
		{
			Address:   "淮安市",
			StartYear: 2000,
		},
	},
	320802: {
		{
			Address:   "清河区",
			StartYear: 1983,
			EndYear:   2015,
		},
	},
	320803: {
		{
			Address:   "楚州区",
			StartYear: 2000,
			EndYear:   2010,
		},
		{
			Address:   "淮安区",
			StartYear: 2011,
		},
	},
	320804: {
		{
			Address:   "淮阴区",
			StartYear: 2000,
		},
	},
	320811: {
		{
			Address:   "清浦区",
			StartYear: 1983,
			EndYear:   2015,
		},
	},
	320812: {
		{
			Address:   "清江浦区",
			StartYear: 2016,
		},
	},
	320813: {
		{
			Address:   "洪泽区",
			StartYear: 2016,
		},
	},
	320821: {
		{
			Address:   "淮阴县",
			StartYear: 1983,
			EndYear:   1999,
		},
	},
	320822: {
		{
			Address:   "灌南县",
			StartYear: 1983,
			EndYear:   1995,
		},
	},
	320823: {
		{
			Address:   "沭阳县",
			StartYear: 1983,
			EndYear:   1995,
		},
	},
	320824: {
		{
			Address:   "宿迁县",
			StartYear: 1983,
			EndYear:   1986,
		},
	},
	320825: {
		{
			Address:   "泗阳县",
			StartYear: 1983,
			EndYear:   1995,
		},
	},
	320826: {
		{
			Address:   "涟水县",
			StartYear: 1983,
		},
	},
	320827: {
		{
			Address:   "泗洪县",
			StartYear: 1983,
			EndYear:   1995,
		},
	},
	320828: {
		{
			Address:   "淮安县",
			StartYear: 1983,
			EndYear:   1986,
		},
	},
	320829: {
		{
			Address:   "洪泽县",
			StartYear: 1983,
			EndYear:   2015,
		},
	},
	320830: {
		{
			Address:   "盱眙县",
			StartYear: 1983,
		},
	},
	320831: {
		{
			Address:   "金湖县",
			StartYear: 1983,
		},
	},
	320881: {
		{
			Address:   "宿迁市",
			StartYear: 1995,
			EndYear:   1995,
		},
	},
	320882: {
		{
			Address:   "淮安市",
			StartYear: 1995,
			EndYear:   1999,
		},
	},
	320900: {
		{
			Address:   "盐城市",
			StartYear: 1983,
		},
	},
	320902: {
		{
			Address:   "城区",
			StartYear: 1983,
			EndYear:   2002,
		},
		{
			Address:   "亭湖区",
			StartYear: 2003,
		},
	},
	320903: {
		{
			Address:   "盐都区",
			StartYear: 2003,
		},
	},
	320904: {
		{
			Address:   "大丰区",
			StartYear: 2015,
		},
	},
	320911: {
		{
			Address:   "郊区",
			StartYear: 1983,
			EndYear:   1995,
		},
	},
	320921: {
		{
			Address:   "响水县",
			StartYear: 1983,
		},
	},
	320922: {
		{
			Address:   "滨海县",
			StartYear: 1983,
		},
	},
	320923: {
		{
			Address:   "阜宁县",
			StartYear: 1983,
		},
	},
	320924: {
		{
			Address:   "射阳县",
			StartYear: 1983,
		},
	},
	320925: {
		{
			Address:   "建湖县",
			StartYear: 1983,
		},
	},
	320926: {
		{
			Address:   "大丰县",
			StartYear: 1983,
			EndYear:   1995,
		},
	},
	320927: {
		{
			Address:   "东台县",
			StartYear: 1983,
			EndYear:   1986,
		},
	},
	320928: {
		{
			Address:   "盐都县",
			StartYear: 1996,
			EndYear:   2002,
		},
	},
	320981: {
		{
			Address:   "东台市",
			StartYear: 1995,
		},
	},
	320982: {
		{
			Address:   "大丰市",
			StartYear: 1996,
			EndYear:   2014,
		},
	},
	321000: {
		{
			Address:   "扬州市",
			StartYear: 1983,
		},
	},
	321002: {
		{
			Address:   "广陵区",
			StartYear: 1983,
		},
	},
	321003: {
		{
			Address:   "邗江区",
			StartYear: 2000,
		},
	},
	321011: {
		{
			Address:   "郊区",
			StartYear: 1984,
			EndYear:   2001,
		},
		{
			Address:   "维扬区",
			StartYear: 2002,
			EndYear:   2010,
		},
	},
	321012: {
		{
			Address:   "江都区",
			StartYear: 2011,
		},
	},
	321021: {
		{
			Address:   "兴化县",
			StartYear: 1983,
			EndYear:   1986,
		},
	},
	321022: {
		{
			Address:   "高邮县",
			StartYear: 1983,
			EndYear:   1990,
		},
	},
	321023: {
		{
			Address:   "宝应县",
			StartYear: 1983,
		},
	},
	321024: {
		{
			Address:   "靖江县",
			StartYear: 1983,
			EndYear:   1992,
		},
	},
	321025: {
		{
			Address:   "泰兴县",
			StartYear: 1983,
			EndYear:   1991,
		},
	},
	321026: {
		{
			Address:   "江都县",
			StartYear: 1983,
			EndYear:   1993,
		},
	},
	321027: {
		{
			Address:   "邗江县",
			StartYear: 1983,
			EndYear:   1999,
		},
	},
	321028: {
		{
			Address:   "泰县",
			StartYear: 1983,
			EndYear:   1993,
		},
	},
	321029: {
		{
			Address:   "仪征县",
			StartYear: 1983,
			EndYear:   1985,
		},
	},
	321081: {
		{
			Address:   "仪征市",
			StartYear: 1995,
		},
	},
	321082: {
		{
			Address:   "泰州市",
			StartYear: 1995,
			EndYear:   1995,
		},
	},
	321083: {
		{
			Address:   "兴化市",
			StartYear: 1995,
			EndYear:   1995,
		},
	},
	321084: {
		{
			Address:   "高邮市",
			StartYear: 1995,
		},
	},
	321085: {
		{
			Address:   "靖江市",
			StartYear: 1995,
			EndYear:   1995,
		},
	},
	321086: {
		{
			Address:   "泰兴市",
			StartYear: 1995,
			EndYear:   1995,
		},
	},
	321087: {
		{
			Address:   "姜堰市",
			StartYear: 1995,
			EndYear:   1995,
		},
	},
	321088: {
		{
			Address:   "江都市",
			StartYear: 1995,
			EndYear:   2010,
		},
	},
	321100: {
		{
			Address:   "镇江市",
			StartYear: 1983,
		},
	},
	321102: {
		{
			Address:   "城区",
			StartYear: 1983,
			EndYear:   1983,
		},
		{
			Address:   "京口区",
			StartYear: 1984,
		},
	},
	321111: {
		{
			Address:   "郊区",
			StartYear: 1983,
			EndYear:   1983,
		},
		{
			Address:   "润州区",
			StartYear: 1984,
		},
	},
	321112: {
		{
			Address:   "丹徒区",
			StartYear: 2002,
		},
	},
	321121: {
		{
			Address:   "丹徒县",
			StartYear: 1983,
			EndYear:   2001,
		},
	},
	321122: {
		{
			Address:   "丹阳县",
			StartYear: 1983,
			EndYear:   1986,
		},
	},
	321123: {
		{
			Address:   "句容县",
			StartYear: 1983,
			EndYear:   1994,
		},
	},
	321124: {
		{
			Address:   "扬中县",
			StartYear: 1983,
			EndYear:   1993,
		},
	},
	321181: {
		{
			Address:   "丹阳市",
			StartYear: 1995,
		},
	},
	321182: {
		{
			Address:   "扬中市",
			StartYear: 1995,
		},
	},
	321183: {
		{
			Address:   "句容市",
			StartYear: 1995,
		},
	},
	321200: {
		{
			Address:   "泰州市",
			StartYear: 1996,
		},
	},
	321201: {
		{
			Address:   "泰州市",
			StartYear: 1983,
			EndYear:   1983,
		},
	},
	321202: {
		{
			Address:   "海陵区",
			StartYear: 1996,
		},
	},
	321203: {
		{
			Address:   "高港区",
			StartYear: 1997,
		},
	},
	321204: {
		{
			Address:   "姜堰区",
			StartYear: 2012,
		},
	},
	321281: {
		{
			Address:   "兴化市",
			StartYear: 1996,
		},
	},
	321282: {
		{
			Address:   "靖江市",
			StartYear: 1996,
		},
	},
	321283: {
		{
			Address:   "泰兴市",
			StartYear: 1996,
		},
	},
	321284: {
		{
			Address:   "姜堰市",
			StartYear: 1996,
			EndYear:   2011,
		},
	},
	321300: {
		{
			Address:   "宿迁市",
			StartYear: 1996,
		},
	},
	321301: {
		{
			Address:   "常熟市",
			StartYear: 1983,
			EndYear:   1983,
		},
	},
	321302: {
		{
			Address:   "宿城区",
			StartYear: 1996,
		},
	},
	321311: {
		{
			Address:   "宿豫区",
			StartYear: 2004,
		},
	},
	321321: {
		{
			Address:   "宿豫县",
			StartYear: 1996,
			EndYear:   2003,
		},
	},
	321322: {
		{
			Address:   "沭阳县",
			StartYear: 1996,
		},
	},
	321323: {
		{
			Address:   "泗阳县",
			StartYear: 1996,
		},
	},
	321324: {
		{
			Address:   "泗洪县",
			StartYear: 1996,
		},
	},
	322100: {
		{
			Address: "徐州地区",
			EndYear: 1982,
		},
	},
	322121: {
		{
			Address: "丰县",
			EndYear: 1982,
		},
	},
	322122: {
		{
			Address: "沛县",
			EndYear: 1982,
		},
	},
	322123: {
		{
			Address: "铜山县",
			EndYear: 1982,
		},
	},
	322124: {
		{
			Address: "睢宁县",
			EndYear: 1982,
		},
	},
	322125: {
		{
			Address: "邳县",
			EndYear: 1982,
		},
	},
	322126: {
		{
			Address: "新沂县",
			EndYear: 1982,
		},
	},
	322127: {
		{
			Address: "东海县",
			EndYear: 1982,
		},
	},
	322128: {
		{
			Address: "赣榆县",
			EndYear: 1982,
		},
	},
	322200: {
		{
			Address: "淮阴地区",
			EndYear: 1982,
		},
	},
	322201: {
		{
			Address: "清江市",
			EndYear: 1982,
		},
	},
	322221: {
		{
			Address: "淮阴县",
			EndYear: 1982,
		},
	},
	322222: {
		{
			Address: "灌云县",
			EndYear: 1982,
		},
	},
	322223: {
		{
			Address: "灌南县",
			EndYear: 1982,
		},
	},
	322224: {
		{
			Address: "沭阳县",
			EndYear: 1982,
		},
	},
	322225: {
		{
			Address: "宿迁县",
			EndYear: 1982,
		},
	},
	322226: {
		{
			Address: "泗阳县",
			EndYear: 1982,
		},
	},
	322227: {
		{
			Address: "涟水县",
			EndYear: 1982,
		},
	},
	322228: {
		{
			Address: "泗洪县",
			EndYear: 1982,
		},
	},
	322229: {
		{
			Address: "淮安县",
			EndYear: 1982,
		},
	},
	322230: {
		{
			Address: "洪泽县",
			EndYear: 1982,
		},
	},
	322231: {
		{
			Address: "盱眙县",
			EndYear: 1982,
		},
	},
	322232: {
		{
			Address: "金湖县",
			EndYear: 1982,
		},
	},
	322300: {
		{
			Address: "盐城地区",
			EndYear: 1982,
		},
	},
	322321: {
		{
			Address: "响水县",
			EndYear: 1982,
		},
	},
	322322: {
		{
			Address: "滨海县",
			EndYear: 1982,
		},
	},
	322323: {
		{
			Address: "阜宁县",
			EndYear: 1982,
		},
	},
	322324: {
		{
			Address: "射阳县",
			EndYear: 1982,
		},
	},
	322325: {
		{
			Address: "建湖县",
			EndYear: 1982,
		},
	},
	322326: {
		{
			Address: "盐城县",
			EndYear: 1982,
		},
	},
	322327: {
		{
			Address: "大丰县",
			EndYear: 1982,
		},
	},
	322328: {
		{
			Address: "东台县",
			EndYear: 1982,
		},
	},
	322400: {
		{
			Address: "扬州地区",
			EndYear: 1982,
		},
	},
	322401: {
		{
			Address: "扬州市",
			EndYear: 1982,
		},
	},
	322402: {
		{
			Address: "泰州市",
			EndYear: 1982,
		},
	},
	322421: {
		{
			Address: "兴化县",
			EndYear: 1982,
		},
	},
	322422: {
		{
			Address: "高邮县",
			EndYear: 1982,
		},
	},
	322423: {
		{
			Address: "宝应县",
			EndYear: 1982,
		},
	},
	322424: {
		{
			Address: "靖江县",
			EndYear: 1982,
		},
	},
	322425: {
		{
			Address: "泰兴县",
			EndYear: 1982,
		},
	},
	322426: {
		{
			Address: "江都县",
			EndYear: 1982,
		},
	},
	322427: {
		{
			Address: "邗江县",
			EndYear: 1982,
		},
	},
	322428: {
		{
			Address: "泰县",
			EndYear: 1982,
		},
	},
	322429: {
		{
			Address: "仪征县",
			EndYear: 1982,
		},
	},
	322500: {
		{
			Address: "南通地区",
			EndYear: 1982,
		},
	},
	322521: {
		{
			Address: "海安县",
			EndYear: 1982,
		},
	},
	322522: {
		{
			Address: "如皋县",
			EndYear: 1982,
		},
	},
	322523: {
		{
			Address: "如东县",
			EndYear: 1982,
		},
	},
	322524: {
		{
			Address: "南通县",
			EndYear: 1982,
		},
	},
	322525: {
		{
			Address: "海门县",
			EndYear: 1982,
		},
	},
	322526: {
		{
			Address: "启东县",
			EndYear: 1982,
		},
	},
	322600: {
		{
			Address: "镇江地区",
			EndYear: 1982,
		},
	},
	322601: {
		{
			Address: "镇江市",
			EndYear: 1982,
		},
	},
	322621: {
		{
			Address: "丹徒县",
			EndYear: 1982,
		},
	},
	322622: {
		{
			Address: "武进县",
			EndYear: 1982,
		},
	},
	322623: {
		{
			Address: "丹阳县",
			EndYear: 1982,
		},
	},
	322624: {
		{
			Address: "句容县",
			EndYear: 1982,
		},
	},
	322625: {
		{
			Address: "金坛县",
			EndYear: 1982,
		},
	},
	322626: {
		{
			Address: "溧水县",
			EndYear: 1982,
		},
	},
	322627: {
		{
			Address: "高淳县",
			EndYear: 1982,
		},
	},
	322628: {
		{
			Address: "溧阳县",
			EndYear: 1982,
		},
	},
	322629: {
		{
			Address: "宜兴县",
			EndYear: 1982,
		},
	},
	322630: {
		{
			Address: "扬中县",
			EndYear: 1982,
		},
	},
	322700: {
		{
			Address: "苏州地区",
			EndYear: 1982,
		},
	},
	322721: {
		{
			Address: "江阴县",
			EndYear: 1982,
		},
	},
	322722: {
		{
			Address: "无锡县",
			EndYear: 1982,
		},
	},
	322723: {
		{
			Address: "沙洲县",
			EndYear: 1982,
		},
	},
	322724: {
		{
			Address: "常熟县",
			EndYear: 1982,
		},
	},
	322725: {
		{
			Address: "太仓县",
			EndYear: 1982,
		},
	},
	322726: {
		{
			Address: "昆山县",
			EndYear: 1982,
		},
	},
	322727: {
		{
			Address: "吴县",
			EndYear: 1982,
		},
	},
	322728: {
		{
			Address: "吴江县",
			EndYear: 1982,
		},
	},
	329001: {
		{
			Address:   "泰州市",
			StartYear: 1984,
			EndYear:   1994,
		},
	},
	329002: {
		{
			Address:   "仪征市",
			StartYear: 1986,
			EndYear:   1994,
		},
	},
	329003: {
		{
			Address:   "常熟市",
			StartYear: 1984,
			EndYear:   1994,
		},
	},
	329004: {
		{
			Address:   "张家港市",
			StartYear: 1986,
			EndYear:   1994,
		},
	},
	329005: {
		{
			Address:   "江阴市",
			StartYear: 1987,
			EndYear:   1994,
		},
	},
	329006: {
		{
			Address:   "宿迁市",
			StartYear: 1987,
			EndYear:   1994,
		},
	},
	329007: {
		{
			Address:   "丹阳市",
			StartYear: 1987,
			EndYear:   1994,
		},
	},
	329008: {
		{
			Address:   "东台市",
			StartYear: 1987,
			EndYear:   1994,
		},
	},
	329009: {
		{
			Address:   "兴化市",
			StartYear: 1987,
			EndYear:   1994,
		},
	},
	329010: {
		{
			Address:   "淮安市",
			StartYear: 1987,
			EndYear:   1994,
		},
	},
	329011: {
		{
			Address:   "宜兴市",
			StartYear: 1988,
			EndYear:   1994,
		},
	},
	329012: {
		{
			Address:   "昆山市",
			StartYear: 1989,
			EndYear:   1994,
		},
	},
	329013: {
		{
			Address:   "启东市",
			StartYear: 1989,
			EndYear:   1994,
		},
	},
	329014: {
		{
			Address:   "新沂市",
			StartYear: 1990,
			EndYear:   1994,
		},
	},
	329015: {
		{
			Address:   "溧阳市",
			StartYear: 1990,
			EndYear:   1994,
		},
	},
	329016: {
		{
			Address:   "如皋市",
			StartYear: 1991,
			EndYear:   1994,
		},
	},
	329017: {
		{
			Address:   "高邮市",
			StartYear: 1991,
			EndYear:   1994,
		},
	},
	329018: {
		{
			Address:   "吴江市",
			StartYear: 1992,
			EndYear:   1994,
		},
	},
	329019: {
		{
			Address:   "邳州市",
			StartYear: 1992,
			EndYear:   1994,
		},
	},
	329020: {
		{
			Address:   "泰兴市",
			StartYear: 1992,
			EndYear:   1994,
		},
	},
	329021: {
		{
			Address:   "通州市",
			StartYear: 1993,
			EndYear:   1994,
		},
	},
	329022: {
		{
			Address:   "太仓市",
			StartYear: 1993,
			EndYear:   1994,
		},
	},
	329023: {
		{
			Address:   "靖江市",
			StartYear: 1993,
			EndYear:   1994,
		},
	},
	329024: {
		{
			Address:   "金坛市",
			StartYear: 1993,
			EndYear:   1994,
		},
	},
	329025: {
		{
			Address:   "江都市",
			StartYear: 1994,
			EndYear:   1994,
		},
	},
	329026: {
		{
			Address:   "海门市",
			StartYear: 1994,
			EndYear:   1994,
		},
	},
	329027: {
		{
			Address:   "扬中市",
			StartYear: 1994,
			EndYear:   1994,
		},
	},
	329028: {
		{
			Address:   "姜堰市",
			StartYear: 1994,
			EndYear:   1994,
		},
	},
	330000: {
		{
			Address: "浙江省",
		},
	},
	330100: {
		{
			Address: "杭州市",
		},
	},
	330102: {
		{
			Address: "上城区",
		},
	},
	330103: {
		{
			Address: "下城区",
		},
	},
	330104: {
		{
			Address: "江干区",
		},
	},
	330105: {
		{
			Address: "拱墅区",
		},
	},
	330106: {
		{
			Address: "西湖区",
		},
	},
	330107: {
		{
			Address: "半山区",
			EndYear: 1989,
		},
	},
	330108: {
		{
			Address:   "滨江区",
			StartYear: 1996,
		},
	},
	330109: {
		{
			Address:   "萧山区",
			StartYear: 2001,
		},
	},
	330110: {
		{
			Address:   "余杭区",
			StartYear: 2001,
		},
	},
	330111: {
		{
			Address:   "富阳区",
			StartYear: 2014,
		},
	},
	330112: {
		{
			Address:   "临安区",
			StartYear: 2017,
		},
	},
	330121: {
		{
			Address: "萧山县",
			EndYear: 1986,
		},
	},
	330122: {
		{
			Address: "桐庐县",
		},
	},
	330123: {
		{
			Address: "富阳县",
			EndYear: 1993,
		},
	},
	330124: {
		{
			Address: "临安县",
			EndYear: 1995,
		},
	},
	330125: {
		{
			Address: "余杭县",
			EndYear: 1993,
		},
	},
	330126: {
		{
			Address: "建德县",
			EndYear: 1991,
		},
	},
	330127: {
		{
			Address: "淳安县",
		},
	},
	330181: {
		{
			Address:   "萧山市",
			StartYear: 1995,
			EndYear:   2000,
		},
	},
	330182: {
		{
			Address:   "建德市",
			StartYear: 1995,
		},
	},
	330183: {
		{
			Address:   "富阳市",
			StartYear: 1995,
			EndYear:   2013,
		},
	},
	330184: {
		{
			Address:   "余杭市",
			StartYear: 1995,
			EndYear:   2000,
		},
	},
	330185: {
		{
			Address:   "临安市",
			StartYear: 1996,
			EndYear:   2016,
		},
	},
	330200: {
		{
			Address: "宁波市",
		},
	},
	330202: {
		{
			Address: "镇明区",
			EndYear: 1984,
		},
	},
	330203: {
		{
			Address: "海曙区",
		},
	},
	330204: {
		{
			Address: "江东区",
			EndYear: 2015,
		},
	},
	330205: {
		{
			Address: "江北区",
		},
	},
	330206: {
		{
			Address:   "滨海区",
			StartYear: 1985,
			EndYear:   1986,
		},
		{
			Address:   "北仑区",
			StartYear: 1987,
		},
	},
	330211: {
		{
			Address:   "镇海区",
			StartYear: 1985,
		},
	},
	330212: {
		{
			Address:   "鄞州区",
			StartYear: 2002,
		},
	},
	330213: {
		{
			Address:   "奉化区",
			StartYear: 2016,
		},
	},
	330219: {
		{
			Address:   "余姚市",
			StartYear: 1985,
			EndYear:   1985,
		},
	},
	330221: {
		{
			Address:   "镇海县",
			StartYear: 1982,
			EndYear:   1984,
		},
	},
	330222: {
		{
			Address:   "慈溪县",
			StartYear: 1983,
			EndYear:   1987,
		},
	},
	330223: {
		{
			Address:   "余姚县",
			StartYear: 1983,
			EndYear:   1984,
		},
	},
	330224: {
		{
			Address:   "奉化县",
			StartYear: 1983,
			EndYear:   1987,
		},
	},
	330225: {
		{
			Address:   "象山县",
			StartYear: 1983,
		},
	},
	330226: {
		{
			Address:   "宁海县",
			StartYear: 1983,
		},
	},
	330227: {
		{
			Address:   "鄞县",
			StartYear: 1983,
			EndYear:   2001,
		},
	},
	330281: {
		{
			Address:   "余姚市",
			StartYear: 1995,
		},
	},
	330282: {
		{
			Address:   "慈溪市",
			StartYear: 1995,
		},
	},
	330283: {
		{
			Address:   "奉化市",
			StartYear: 1995,
			EndYear:   2015,
		},
	},
	330300: {
		{
			Address: "温州市",
		},
	},
	330301: {
		{
			Address: "东城区",
			EndYear: 1982,
		},
	},
	330302: {
		{
			Address: "南城区",
			EndYear: 1982,
		},
		{
			Address:   "城区",
			StartYear: 1983,
			EndYear:   1983,
		},
		{
			Address:   "鹿城区",
			StartYear: 1984,
		},
	},
	330303: {
		{
			Address: "西城区",
			EndYear: 1982,
		},
		{
			Address:   "龙湾区",
			StartYear: 1984,
		},
	},
	330304: {
		{
			Address:   "瓯海区",
			StartYear: 1992,
		},
	},
	330305: {
		{
			Address:   "洞头区",
			StartYear: 2015,
		},
	},
	330321: {
		{
			Address:   "洞头县",
			StartYear: 1981,
			EndYear:   1981,
		},
		{
			Address:   "瓯海县",
			StartYear: 1982,
			EndYear:   1991,
		},
	},
	330322: {
		{
			Address:   "永嘉县",
			StartYear: 1981,
			EndYear:   1981,
		},
		{
			Address:   "洞头县",
			StartYear: 1982,
			EndYear:   2014,
		},
	},
	330323: {
		{
			Address:   "瑞安县",
			StartYear: 1981,
			EndYear:   1981,
		},
		{
			Address:   "乐清县",
			StartYear: 1982,
			EndYear:   1992,
		},
	},
	330324: {
		{
			Address:   "文成县",
			StartYear: 1981,
			EndYear:   1981,
		},
		{
			Address:   "永嘉县",
			StartYear: 1982,
		},
	},
	330325: {
		{
			Address:   "平阳县",
			StartYear: 1981,
			EndYear:   1981,
		},
		{
			Address:   "瑞安县",
			StartYear: 1982,
			EndYear:   1986,
		},
	},
	330326: {
		{
			Address:   "乐清县",
			StartYear: 1981,
			EndYear:   1981,
		},
		{
			Address:   "平阳县",
			StartYear: 1982,
		},
	},
	330327: {
		{
			Address:   "泰顺县",
			StartYear: 1981,
			EndYear:   1981,
		},
		{
			Address:   "苍南县",
			StartYear: 1982,
		},
	},
	330328: {
		{
			Address:   "瓯海县",
			StartYear: 1981,
			EndYear:   1981,
		},
		{
			Address:   "文成县",
			StartYear: 1982,
		},
	},
	330329: {
		{
			Address:   "苍南县",
			StartYear: 1981,
			EndYear:   1981,
		},
		{
			Address:   "泰顺县",
			StartYear: 1982,
		},
	},
	330381: {
		{
			Address:   "瑞安市",
			StartYear: 1995,
		},
	},
	330382: {
		{
			Address:   "乐清市",
			StartYear: 1995,
		},
	},
	330383: {
		{
			Address:   "龙港市",
			StartYear: 2019,
		},
	},
	330400: {
		{
			Address:   "嘉兴市",
			StartYear: 1983,
		},
	},
	330402: {
		{
			Address:   "城区",
			StartYear: 1983,
			EndYear:   1992,
		},
		{
			Address:   "秀城区",
			StartYear: 1993,
			EndYear:   2004,
		},
		{
			Address:   "南湖区",
			StartYear: 2005,
		},
	},
	330411: {
		{
			Address:   "郊区",
			StartYear: 1983,
			EndYear:   1998,
		},
		{
			Address:   "秀洲区",
			StartYear: 1999,
		},
	},
	330421: {
		{
			Address:   "嘉善县",
			StartYear: 1983,
		},
	},
	330422: {
		{
			Address:   "平湖县",
			StartYear: 1983,
			EndYear:   1990,
		},
	},
	330423: {
		{
			Address:   "海宁县",
			StartYear: 1983,
			EndYear:   1985,
		},
	},
	330424: {
		{
			Address:   "海盐县",
			StartYear: 1983,
		},
	},
	330425: {
		{
			Address:   "桐乡县",
			StartYear: 1983,
			EndYear:   1992,
		},
	},
	330481: {
		{
			Address:   "海宁市",
			StartYear: 1995,
		},
	},
	330482: {
		{
			Address:   "平湖市",
			StartYear: 1995,
		},
	},
	330483: {
		{
			Address:   "桐乡市",
			StartYear: 1995,
		},
	},
	330500: {
		{
			Address:   "湖州市",
			StartYear: 1983,
		},
	},
	330502: {
		{
			Address:   "城区",
			StartYear: 1983,
			EndYear:   1987,
		},
		{
			Address:   "吴兴区",
			StartYear: 2003,
		},
	},
	330503: {
		{
			Address:   "南浔区",
			StartYear: 2003,
		},
	},
	330511: {
		{
			Address:   "郊区",
			StartYear: 1983,
			EndYear:   1987,
		},
	},
	330521: {
		{
			Address:   "德清县",
			StartYear: 1983,
		},
	},
	330522: {
		{
			Address:   "长兴县",
			StartYear: 1983,
		},
	},
	330523: {
		{
			Address:   "安吉县",
			StartYear: 1983,
		},
	},
	330600: {
		{
			Address:   "绍兴市",
			StartYear: 1983,
		},
	},
	330602: {
		{
			Address:   "越城区",
			StartYear: 1983,
		},
	},
	330603: {
		{
			Address:   "柯桥区",
			StartYear: 2013,
		},
	},
	330604: {
		{
			Address:   "上虞区",
			StartYear: 2013,
		},
	},
	330621: {
		{
			Address:   "绍兴县",
			StartYear: 1983,
			EndYear:   2012,
		},
	},
	330622: {
		{
			Address:   "上虞县",
			StartYear: 1983,
			EndYear:   1991,
		},
	},
	330623: {
		{
			Address:   "嵊县",
			StartYear: 1983,
			EndYear:   1994,
		},
	},
	330624: {
		{
			Address:   "新昌县",
			StartYear: 1983,
		},
	},
	330625: {
		{
			Address:   "诸暨县",
			StartYear: 1983,
			EndYear:   1988,
		},
	},
	330681: {
		{
			Address:   "诸暨市",
			StartYear: 1995,
		},
	},
	330682: {
		{
			Address:   "上虞市",
			StartYear: 1995,
			EndYear:   2012,
		},
	},
	330683: {
		{
			Address:   "嵊州市",
			StartYear: 1995,
		},
	},
	330700: {
		{
			Address:   "金华市",
			StartYear: 1985,
		},
	},
	330701: {
		{
			Address:   "兰溪市",
			StartYear: 1985,
			EndYear:   1985,
		},
	},
	330702: {
		{
			Address:   "婺城区",
			StartYear: 1985,
		},
	},
	330703: {
		{
			Address:   "金东区",
			StartYear: 2000,
		},
	},
	330721: {
		{
			Address:   "金华县",
			StartYear: 1985,
			EndYear:   1999,
		},
	},
	330722: {
		{
			Address:   "永康县",
			StartYear: 1985,
			EndYear:   1991,
		},
	},
	330723: {
		{
			Address:   "武义县",
			StartYear: 1985,
		},
	},
	330724: {
		{
			Address:   "东阳县",
			StartYear: 1985,
			EndYear:   1987,
		},
	},
	330725: {
		{
			Address:   "义乌县",
			StartYear: 1985,
			EndYear:   1987,
		},
	},
	330726: {
		{
			Address:   "浦江县",
			StartYear: 1985,
		},
	},
	330727: {
		{
			Address:   "磐安县",
			StartYear: 1985,
		},
	},
	330781: {
		{
			Address:   "兰溪市",
			StartYear: 1995,
		},
	},
	330782: {
		{
			Address:   "义乌市",
			StartYear: 1995,
		},
	},
	330783: {
		{
			Address:   "东阳市",
			StartYear: 1995,
		},
	},
	330784: {
		{
			Address:   "永康市",
			StartYear: 1995,
		},
	},
	330800: {
		{
			Address:   "衢州市",
			StartYear: 1985,
		},
	},
	330802: {
		{
			Address:   "柯城区",
			StartYear: 1985,
		},
	},
	330803: {
		{
			Address:   "衢江区",
			StartYear: 2001,
		},
	},
	330821: {
		{
			Address:   "衢县",
			StartYear: 1985,
			EndYear:   2000,
		},
	},
	330822: {
		{
			Address:   "常山县",
			StartYear: 1985,
		},
	},
	330823: {
		{
			Address:   "江山县",
			StartYear: 1985,
			EndYear:   1986,
		},
	},
	330824: {
		{
			Address:   "开化县",
			StartYear: 1985,
		},
	},
	330825: {
		{
			Address:   "龙游县",
			StartYear: 1985,
		},
	},
	330881: {
		{
			Address:   "江山市",
			StartYear: 1995,
		},
	},
	330900: {
		{
			Address:   "舟山市",
			StartYear: 1987,
		},
	},
	330902: {
		{
			Address:   "定海区",
			StartYear: 1987,
		},
	},
	330903: {
		{
			Address:   "普陀区",
			StartYear: 1987,
		},
	},
	330921: {
		{
			Address:   "岱山县",
			StartYear: 1987,
		},
	},
	330922: {
		{
			Address:   "嵊泗县",
			StartYear: 1987,
		},
	},
	331000: {
		{
			Address:   "台州市",
			StartYear: 1994,
		},
	},
	331002: {
		{
			Address:   "椒江区",
			StartYear: 1994,
		},
	},
	331003: {
		{
			Address:   "黄岩区",
			StartYear: 1994,
		},
	},
	331004: {
		{
			Address:   "路桥区",
			StartYear: 1994,
		},
	},
	331021: {
		{
			Address:   "玉环县",
			StartYear: 1994,
			EndYear:   2016,
		},
	},
	331022: {
		{
			Address:   "三门县",
			StartYear: 1994,
		},
	},
	331023: {
		{
			Address:   "天台县",
			StartYear: 1994,
		},
	},
	331024: {
		{
			Address:   "仙居县",
			StartYear: 1994,
		},
	},
	331081: {
		{
			Address:   "温岭市",
			StartYear: 1995,
		},
	},
	331082: {
		{
			Address:   "临海市",
			StartYear: 1995,
		},
	},
	331083: {
		{
			Address:   "玉环市",
			StartYear: 2017,
		},
	},
	331100: {
		{
			Address:   "丽水市",
			StartYear: 2000,
		},
	},
	331102: {
		{
			Address:   "莲都区",
			StartYear: 2000,
		},
	},
	331121: {
		{
			Address:   "青田县",
			StartYear: 2000,
		},
	},
	331122: {
		{
			Address:   "缙云县",
			StartYear: 2000,
		},
	},
	331123: {
		{
			Address:   "遂昌县",
			StartYear: 2000,
		},
	},
	331124: {
		{
			Address:   "松阳县",
			StartYear: 2000,
		},
	},
	331125: {
		{
			Address:   "云和县",
			StartYear: 2000,
		},
	},
	331126: {
		{
			Address:   "庆元县",
			StartYear: 2000,
		},
	},
	331127: {
		{
			Address:   "景宁畲族自治县",
			StartYear: 2000,
		},
	},
	331181: {
		{
			Address:   "龙泉市",
			StartYear: 2000,
		},
	},
	332100: {
		{
			Address: "嘉兴地区",
			EndYear: 1982,
		},
	},
	332101: {
		{
			Address: "湖州市",
			EndYear: 1982,
		},
	},
	332102: {
		{
			Address: "嘉兴市",
			EndYear: 1982,
		},
	},
	332121: {
		{
			Address: "嘉兴县",
			EndYear: 1980,
		},
		{
			Address:   "嘉善县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	332122: {
		{
			Address: "嘉善县",
			EndYear: 1981,
		},
		{
			Address:   "平湖县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	332123: {
		{
			Address: "平湖县",
			EndYear: 1981,
		},
		{
			Address:   "海宁县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	332124: {
		{
			Address: "海宁县",
			EndYear: 1981,
		},
		{
			Address:   "海盐县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	332125: {
		{
			Address: "海盐县",
			EndYear: 1981,
		},
		{
			Address:   "桐乡县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	332126: {
		{
			Address: "桐乡县",
			EndYear: 1981,
		},
		{
			Address:   "德清县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	332127: {
		{
			Address: "德清县",
			EndYear: 1981,
		},
		{
			Address:   "长兴县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	332128: {
		{
			Address: "吴兴县",
			EndYear: 1980,
		},
		{
			Address:   "安吉县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	332129: {
		{
			Address: "长兴县",
			EndYear: 1981,
		},
	},
	332130: {
		{
			Address: "安吉县",
			EndYear: 1981,
		},
	},
	332200: {
		{
			Address: "宁波地区",
			EndYear: 1982,
		},
	},
	332221: {
		{
			Address: "慈溪县",
			EndYear: 1982,
		},
	},
	332222: {
		{
			Address: "余姚县",
			EndYear: 1982,
		},
	},
	332223: {
		{
			Address: "奉化县",
			EndYear: 1982,
		},
	},
	332224: {
		{
			Address: "象山县",
			EndYear: 1982,
		},
	},
	332225: {
		{
			Address: "宁海县",
			EndYear: 1982,
		},
	},
	332226: {
		{
			Address: "鄞县",
			EndYear: 1982,
		},
	},
	332227: {
		{
			Address: "镇海县",
			EndYear: 1981,
		},
	},
	332300: {
		{
			Address: "绍兴地区",
			EndYear: 1982,
		},
	},
	332301: {
		{
			Address: "绍兴市",
			EndYear: 1982,
		},
	},
	332321: {
		{
			Address: "绍兴县",
			EndYear: 1980,
		},
		{
			Address:   "上虞县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	332322: {
		{
			Address: "上虞县",
			EndYear: 1981,
		},
		{
			Address:   "嵊县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	332323: {
		{
			Address: "嵊县",
			EndYear: 1981,
		},
		{
			Address:   "新昌县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	332324: {
		{
			Address: "新昌县",
			EndYear: 1981,
		},
		{
			Address:   "诸暨县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	332325: {
		{
			Address: "诸暨县",
			EndYear: 1981,
		},
	},
	332400: {
		{
			Address: "温州地区",
			EndYear: 1980,
		},
		{
			Address:   "金华地区",
			StartYear: 1982,
			EndYear:   1984,
		},
	},
	332401: {
		{
			Address:   "金华市",
			StartYear: 1982,
			EndYear:   1984,
		},
	},
	332402: {
		{
			Address:   "衢州市",
			StartYear: 1982,
			EndYear:   1984,
		},
	},
	332421: {
		{
			Address: "洞头县",
			EndYear: 1980,
		},
		{
			Address:   "兰溪县",
			StartYear: 1982,
			EndYear:   1984,
		},
	},
	332422: {
		{
			Address: "永嘉县",
			EndYear: 1980,
		},
		{
			Address:   "永康县",
			StartYear: 1982,
			EndYear:   1984,
		},
	},
	332423: {
		{
			Address: "瑞安县",
			EndYear: 1980,
		},
		{
			Address:   "武义县",
			StartYear: 1982,
			EndYear:   1984,
		},
	},
	332424: {
		{
			Address: "文成县",
			EndYear: 1980,
		},
		{
			Address:   "东阳县",
			StartYear: 1982,
			EndYear:   1984,
		},
	},
	332425: {
		{
			Address: "平阳县",
			EndYear: 1980,
		},
		{
			Address:   "义乌县",
			StartYear: 1982,
			EndYear:   1984,
		},
	},
	332426: {
		{
			Address: "乐清县",
			EndYear: 1980,
		},
		{
			Address:   "浦江县",
			StartYear: 1982,
			EndYear:   1984,
		},
	},
	332427: {
		{
			Address: "泰顺县",
			EndYear: 1980,
		},
		{
			Address:   "常山县",
			StartYear: 1982,
			EndYear:   1984,
		},
	},
	332428: {
		{
			Address:   "江山县",
			StartYear: 1982,
			EndYear:   1984,
		},
	},
	332429: {
		{
			Address:   "开化县",
			StartYear: 1982,
			EndYear:   1984,
		},
	},
	332430: {
		{
			Address:   "龙游县",
			StartYear: 1983,
			EndYear:   1984,
		},
	},
	332431: {
		{
			Address:   "磐安县",
			StartYear: 1983,
			EndYear:   1984,
		},
	},
	332500: {
		{
			Address: "金华地区",
			EndYear: 1981,
		},
		{
			Address:   "丽水地区",
			StartYear: 1982,
			EndYear:   1999,
		},
	},
	332501: {
		{
			Address: "金华市",
			EndYear: 1981,
		},
		{
			Address:   "丽水市",
			StartYear: 1986,
			EndYear:   1994,
		},
	},
	332502: {
		{
			Address: "衢州市",
			EndYear: 1981,
		},
		{
			Address:   "龙泉市",
			StartYear: 1990,
			EndYear:   1994,
		},
	},
	332521: {
		{
			Address: "金华县",
			EndYear: 1980,
		},
		{
			Address:   "丽水县",
			StartYear: 1982,
			EndYear:   1985,
		},
	},
	332522: {
		{
			Address: "兰溪县",
			EndYear: 1981,
		},
		{
			Address:   "青田县",
			StartYear: 1982,
			EndYear:   1999,
		},
	},
	332523: {
		{
			Address: "永康县",
			EndYear: 1981,
		},
		{
			Address:   "云和县",
			StartYear: 1982,
			EndYear:   1999,
		},
	},
	332524: {
		{
			Address: "武义县",
			EndYear: 1981,
		},
		{
			Address:   "龙泉县",
			StartYear: 1982,
			EndYear:   1989,
		},
	},
	332525: {
		{
			Address: "东阳县",
			EndYear: 1981,
		},
		{
			Address:   "庆元县",
			StartYear: 1982,
			EndYear:   1999,
		},
	},
	332526: {
		{
			Address: "义乌县",
			EndYear: 1981,
		},
		{
			Address:   "缙云县",
			StartYear: 1982,
			EndYear:   1999,
		},
	},
	332527: {
		{
			Address: "浦江县",
			EndYear: 1981,
		},
		{
			Address:   "遂昌县",
			StartYear: 1982,
			EndYear:   1999,
		},
	},
	332528: {
		{
			Address: "衢县",
			EndYear: 1980,
		},
		{
			Address:   "松阳县",
			StartYear: 1982,
			EndYear:   1999,
		},
	},
	332529: {
		{
			Address: "常山县",
			EndYear: 1981,
		},
		{
			Address:   "景宁畲族自治县",
			StartYear: 1984,
			EndYear:   1999,
		},
	},
	332530: {
		{
			Address: "江山县",
			EndYear: 1981,
		},
	},
	332531: {
		{
			Address: "开化县",
			EndYear: 1981,
		},
	},
	332581: {
		{
			Address:   "龙泉市",
			StartYear: 1995,
			EndYear:   1999,
		},
	},
	332582: {
		{
			Address:   "丽水市",
			StartYear: 1995,
			EndYear:   1999,
		},
	},
	332600: {
		{
			Address: "丽水地区",
			EndYear: 1981,
		},
		{
			Address:   "台州地区",
			StartYear: 1982,
			EndYear:   1993,
		},
	},
	332601: {
		{
			Address:   "椒江市",
			StartYear: 1982,
			EndYear:   1993,
		},
	},
	332602: {
		{
			Address:   "临海市",
			StartYear: 1986,
			EndYear:   1993,
		},
	},
	332603: {
		{
			Address:   "黄岩市",
			StartYear: 1989,
			EndYear:   1993,
		},
	},
	332621: {
		{
			Address: "丽水县",
			EndYear: 1981,
		},
		{
			Address:   "临海县",
			StartYear: 1982,
			EndYear:   1985,
		},
	},
	332622: {
		{
			Address: "青田县",
			EndYear: 1981,
		},
		{
			Address:   "黄岩县",
			StartYear: 1982,
			EndYear:   1988,
		},
	},
	332623: {
		{
			Address: "云和县",
			EndYear: 1981,
		},
		{
			Address:   "温岭县",
			StartYear: 1982,
			EndYear:   1993,
		},
	},
	332624: {
		{
			Address: "龙泉县",
			EndYear: 1981,
		},
		{
			Address:   "仙居县",
			StartYear: 1982,
			EndYear:   1993,
		},
	},
	332625: {
		{
			Address: "庆元县",
			EndYear: 1981,
		},
		{
			Address:   "天台县",
			StartYear: 1982,
			EndYear:   1993,
		},
	},
	332626: {
		{
			Address: "缙云县",
			EndYear: 1981,
		},
		{
			Address:   "三门县",
			StartYear: 1982,
			EndYear:   1993,
		},
	},
	332627: {
		{
			Address: "遂昌县",
			EndYear: 1981,
		},
		{
			Address:   "玉环县",
			StartYear: 1982,
			EndYear:   1993,
		},
	},
	332700: {
		{
			Address: "台州地区",
			EndYear: 1981,
		},
		{
			Address:   "舟山地区",
			StartYear: 1982,
			EndYear:   1986,
		},
	},
	332701: {
		{
			Address:   "椒江市",
			StartYear: 1981,
			EndYear:   1981,
		},
	},
	332721: {
		{
			Address: "临海县",
			EndYear: 1981,
		},
		{
			Address:   "定海县",
			StartYear: 1982,
			EndYear:   1986,
		},
	},
	332722: {
		{
			Address: "黄岩县",
			EndYear: 1981,
		},
		{
			Address:   "普陀县",
			StartYear: 1982,
			EndYear:   1986,
		},
	},
	332723: {
		{
			Address: "温岭县",
			EndYear: 1981,
		},
		{
			Address:   "岱山县",
			StartYear: 1982,
			EndYear:   1986,
		},
	},
	332724: {
		{
			Address: "仙居县",
			EndYear: 1981,
		},
		{
			Address:   "嵊泗县",
			StartYear: 1982,
			EndYear:   1986,
		},
	},
	332725: {
		{
			Address: "天台县",
			EndYear: 1981,
		},
	},
	332726: {
		{
			Address: "三门县",
			EndYear: 1981,
		},
	},
	332727: {
		{
			Address: "玉环县",
			EndYear: 1981,
		},
	},
	332800: {
		{
			Address: "舟山地区",
			EndYear: 1981,
		},
	},
	332821: {
		{
			Address: "定海县",
			EndYear: 1981,
		},
	},
	332822: {
		{
			Address: "普陀县",
			EndYear: 1981,
		},
	},
	332823: {
		{
			Address: "岱山县",
			EndYear: 1981,
		},
	},
	332824: {
		{
			Address: "嵊泗县",
			EndYear: 1981,
		},
	},
	339001: {
		{
			Address:   "余姚市",
			StartYear: 1986,
			EndYear:   1994,
		},
	},
	339002: {
		{
			Address:   "海宁市",
			StartYear: 1986,
			EndYear:   1994,
		},
	},
	339003: {
		{
			Address:   "兰溪市",
			StartYear: 1986,
			EndYear:   1994,
		},
	},
	339004: {
		{
			Address:   "瑞安市",
			StartYear: 1987,
			EndYear:   1994,
		},
	},
	339005: {
		{
			Address:   "萧山市",
			StartYear: 1987,
			EndYear:   1994,
		},
	},
	339006: {
		{
			Address:   "江山市",
			StartYear: 1987,
			EndYear:   1994,
		},
	},
	339007: {
		{
			Address:   "安徽省",
			StartYear: 1987,
			EndYear:   1987,
		},
		{
			Address:   "义乌市",
			StartYear: 1988,
			EndYear:   1994,
		},
	},
	339008: {
		{
			Address:   "合肥市",
			StartYear: 1987,
			EndYear:   1987,
		},
		{
			Address:   "东阳市",
			StartYear: 1988,
			EndYear:   1994,
		},
	},
	339009: {
		{
			Address:   "慈溪市",
			StartYear: 1988,
			EndYear:   1994,
		},
	},
	339010: {
		{
			Address:   "奉化市",
			StartYear: 1988,
			EndYear:   1994,
		},
	},
	339011: {
		{
			Address:   "诸暨市",
			StartYear: 1989,
			EndYear:   1994,
		},
	},
	339012: {
		{
			Address:   "平湖市",
			StartYear: 1991,
			EndYear:   1994,
		},
	},
	339013: {
		{
			Address:   "建德市",
			StartYear: 1992,
			EndYear:   1994,
		},
	},
	339014: {
		{
			Address:   "永康市",
			StartYear: 1992,
			EndYear:   1994,
		},
	},
	339015: {
		{
			Address:   "上虞市",
			StartYear: 1992,
			EndYear:   1994,
		},
	},
	339016: {
		{
			Address:   "桐乡市",
			StartYear: 1993,
			EndYear:   1994,
		},
	},
	339017: {
		{
			Address:   "乐清市",
			StartYear: 1993,
			EndYear:   1994,
		},
	},
	339018: {
		{
			Address:   "临海市",
			StartYear: 1994,
			EndYear:   1994,
		},
	},
	339019: {
		{
			Address:   "富阳市",
			StartYear: 1994,
			EndYear:   1994,
		},
	},
	339020: {
		{
			Address:   "温岭市",
			StartYear: 1994,
			EndYear:   1994,
		},
	},
	339021: {
		{
			Address:   "余杭市",
			StartYear: 1994,
			EndYear:   1994,
		},
	},
	340000: {
		{
			Address: "安徽省",
		},
	},
	340100: {
		{
			Address: "合肥市",
		},
	},
	340102: {
		{
			Address: "东市区",
			EndYear: 2001,
		},
		{
			Address:   "瑶海区",
			StartYear: 2002,
		},
	},
	340103: {
		{
			Address: "中市区",
			EndYear: 2001,
		},
		{
			Address:   "庐阳区",
			StartYear: 2002,
		},
	},
	340104: {
		{
			Address: "西市区",
			EndYear: 2001,
		},
		{
			Address:   "蜀山区",
			StartYear: 2002,
		},
	},
	340111: {
		{
			Address: "郊区",
			EndYear: 2001,
		},
		{
			Address:   "包河区",
			StartYear: 2002,
		},
	},
	340121: {
		{
			Address: "长丰县",
		},
	},
	340122: {
		{
			Address:   "肥东县",
			StartYear: 1983,
		},
	},
	340123: {
		{
			Address: "肥西县",
		},
	},
	340124: {
		{
			Address:   "庐江县",
			StartYear: 2011,
		},
	},
	340181: {
		{
			Address:   "巢湖市",
			StartYear: 2011,
		},
	},
	340200: {
		{
			Address: "芜湖市",
		},
	},
	340202: {
		{
			Address: "镜湖区",
		},
	},
	340203: {
		{
			Address: "马塘区",
			EndYear: 2004,
		},
		{
			Address:   "弋江区",
			StartYear: 2005,
			EndYear:   2019,
		},
	},
	340204: {
		{
			Address: "新芜区",
			EndYear: 2004,
		},
	},
	340205: {
		{
			Address: "裕溪口区",
			EndYear: 1989,
		},
		{
			Address:   "鸠江区",
			StartYear: 1990,
			EndYear:   2004,
		},
	},
	340206: {
		{
			Address: "四褐山区",
			EndYear: 1989,
		},
	},
	340207: {
		{
			Address:   "鸠江区",
			StartYear: 2005,
		},
	},
	340208: {
		{
			Address:   "三山区",
			StartYear: 2005,
			EndYear:   2019,
		},
	},
	340209: {
		{
			Address:   "弋江区",
			StartYear: 2020,
		},
	},
	340210: {
		{
			Address:   "湾沚区",
			StartYear: 2020,
		},
	},
	340211: {
		{
			Address: "郊区",
			EndYear: 1989,
		},
	},
	340212: {
		{
			Address:   "繁昌区",
			StartYear: 2020,
		},
	},
	340221: {
		{
			Address: "芜湖县",
			EndYear: 2019,
		},
	},
	340222: {
		{
			Address: "繁昌县",
			EndYear: 2019,
		},
	},
	340223: {
		{
			Address: "南陵县",
		},
	},
	340224: {
		{
			Address: "青阳县",
			EndYear: 1987,
		},
	},
	340225: {
		{
			Address:   "无为县",
			StartYear: 2011,
			EndYear:   2018,
		},
	},
	340281: {
		{
			Address:   "无为市",
			StartYear: 2019,
		},
	},
	340300: {
		{
			Address: "蚌埠市",
		},
	},
	340302: {
		{
			Address: "东市区",
			EndYear: 2003,
		},
		{
			Address:   "龙子湖区",
			StartYear: 2004,
		},
	},
	340303: {
		{
			Address: "中市区",
			EndYear: 2003,
		},
		{
			Address:   "蚌山区",
			StartYear: 2004,
		},
	},
	340304: {
		{
			Address: "西市区",
			EndYear: 2003,
		},
		{
			Address:   "禹会区",
			StartYear: 2004,
		},
	},
	340311: {
		{
			Address: "郊区",
			EndYear: 2003,
		},
		{
			Address:   "淮上区",
			StartYear: 2004,
		},
	},
	340321: {
		{
			Address:   "怀远县",
			StartYear: 1983,
		},
	},
	340322: {
		{
			Address:   "五河县",
			StartYear: 1983,
		},
	},
	340323: {
		{
			Address:   "固镇县",
			StartYear: 1983,
		},
	},
	340400: {
		{
			Address: "淮南市",
		},
	},
	340402: {
		{
			Address: "大通区",
		},
	},
	340403: {
		{
			Address: "田家庵区",
		},
	},
	340404: {
		{
			Address: "谢家集区",
		},
	},
	340405: {
		{
			Address: "八公山区",
		},
	},
	340406: {
		{
			Address: "潘集区",
		},
	},
	340421: {
		{
			Address: "凤台县",
		},
	},
	340422: {
		{
			Address:   "寿县",
			StartYear: 2015,
		},
	},
	340500: {
		{
			Address: "马鞍山市",
		},
	},
	340502: {
		{
			Address: "金家庄区",
			EndYear: 2011,
		},
	},
	340503: {
		{
			Address: "花山区",
		},
	},
	340504: {
		{
			Address: "雨山区",
		},
	},
	340505: {
		{
			Address: "向山区",
			EndYear: 2000,
		},
	},
	340506: {
		{
			Address:   "博望区",
			StartYear: 2012,
		},
	},
	340511: {
		{
			Address:   "郊区",
			StartYear: 1983,
			EndYear:   1983,
		},
	},
	340521: {
		{
			Address: "当涂县",
		},
	},
	340522: {
		{
			Address:   "含山县",
			StartYear: 2011,
		},
	},
	340523: {
		{
			Address:   "和县",
			StartYear: 2011,
		},
	},
	340600: {
		{
			Address: "淮北市",
		},
	},
	340602: {
		{
			Address: "杜集区",
		},
	},
	340603: {
		{
			Address: "相山区",
		},
	},
	340604: {
		{
			Address: "烈山区",
		},
	},
	340611: {
		{
			Address: "郊区",
			EndYear: 1983,
		},
	},
	340621: {
		{
			Address: "濉溪县",
		},
	},
	340700: {
		{
			Address: "铜陵市",
		},
	},
	340702: {
		{
			Address: "铜官山区",
			EndYear: 2014,
		},
	},
	340703: {
		{
			Address: "狮子山区",
			EndYear: 2014,
		},
	},
	340704: {
		{
			Address: "铜山区",
			EndYear: 1986,
		},
	},
	340705: {
		{
			Address:   "铜官区",
			StartYear: 2015,
		},
	},
	340706: {
		{
			Address:   "义安区",
			StartYear: 2015,
		},
	},
	340711: {
		{
			Address: "郊区",
		},
	},
	340721: {
		{
			Address: "铜陵县",
			EndYear: 2014,
		},
	},
	340722: {
		{
			Address:   "枞阳县",
			StartYear: 2015,
		},
	},
	340800: {
		{
			Address: "安庆市",
		},
	},
	340802: {
		{
			Address: "迎江区",
		},
	},
	340803: {
		{
			Address: "大观区",
		},
	},
	340811: {
		{
			Address: "郊区",
			EndYear: 2004,
		},
		{
			Address:   "宜秀区",
			StartYear: 2005,
		},
	},
	340821: {
		{
			Address:   "桐城县",
			StartYear: 1988,
			EndYear:   1995,
		},
	},
	340822: {
		{
			Address:   "怀宁县",
			StartYear: 1988,
		},
	},
	340823: {
		{
			Address:   "枞阳县",
			StartYear: 1988,
			EndYear:   2014,
		},
	},
	340824: {
		{
			Address:   "潜山县",
			StartYear: 1988,
			EndYear:   2017,
		},
	},
	340825: {
		{
			Address:   "太湖县",
			StartYear: 1988,
		},
	},
	340826: {
		{
			Address:   "宿松县",
			StartYear: 1988,
		},
	},
	340827: {
		{
			Address:   "望江县",
			StartYear: 1988,
		},
	},
	340828: {
		{
			Address:   "岳西县",
			StartYear: 1988,
		},
	},
	340881: {
		{
			Address:   "桐城市",
			StartYear: 1996,
		},
	},
	340882: {
		{
			Address:   "潜山市",
			StartYear: 2018,
		},
	},
	340901: {
		{
			Address:   "黄山市",
			StartYear: 1983,
			EndYear:   1986,
		},
	},
	341000: {
		{
			Address:   "黄山市",
			StartYear: 1987,
		},
	},
	341002: {
		{
			Address:   "屯溪区",
			StartYear: 1987,
		},
	},
	341003: {
		{
			Address:   "黄山区",
			StartYear: 1987,
		},
	},
	341004: {
		{
			Address:   "徽州区",
			StartYear: 1987,
		},
	},
	341021: {
		{
			Address:   "歙县",
			StartYear: 1987,
		},
	},
	341022: {
		{
			Address:   "休宁县",
			StartYear: 1987,
		},
	},
	341023: {
		{
			Address:   "黟县",
			StartYear: 1987,
		},
	},
	341024: {
		{
			Address:   "祁门县",
			StartYear: 1987,
		},
	},
	341100: {
		{
			Address:   "滁州市",
			StartYear: 1992,
		},
	},
	341102: {
		{
			Address:   "琅琊区",
			StartYear: 1992,
		},
	},
	341103: {
		{
			Address:   "南谯区",
			StartYear: 1992,
		},
	},
	341121: {
		{
			Address:   "天长县",
			StartYear: 1992,
			EndYear:   1992,
		},
	},
	341122: {
		{
			Address:   "来安县",
			StartYear: 1992,
		},
	},
	341124: {
		{
			Address:   "全椒县",
			StartYear: 1992,
		},
	},
	341125: {
		{
			Address:   "定远县",
			StartYear: 1992,
		},
	},
	341126: {
		{
			Address:   "凤阳县",
			StartYear: 1992,
		},
	},
	341127: {
		{
			Address:   "嘉山县",
			StartYear: 1992,
			EndYear:   1993,
		},
	},
	341181: {
		{
			Address:   "天长市",
			StartYear: 1995,
		},
	},
	341182: {
		{
			Address:   "明光市",
			StartYear: 1995,
		},
	},
	341200: {
		{
			Address:   "阜阳市",
			StartYear: 1996,
		},
	},
	341202: {
		{
			Address:   "颍州区",
			StartYear: 1996,
		},
	},
	341203: {
		{
			Address:   "颍东区",
			StartYear: 1996,
		},
	},
	341204: {
		{
			Address:   "颍泉区",
			StartYear: 1996,
		},
	},
	341221: {
		{
			Address:   "临泉县",
			StartYear: 1996,
		},
	},
	341222: {
		{
			Address:   "太和县",
			StartYear: 1996,
		},
	},
	341223: {
		{
			Address:   "涡阳县",
			StartYear: 1996,
			EndYear:   1999,
		},
	},
	341224: {
		{
			Address:   "蒙城县",
			StartYear: 1996,
			EndYear:   1999,
		},
	},
	341225: {
		{
			Address:   "阜南县",
			StartYear: 1996,
		},
	},
	341226: {
		{
			Address:   "颍上县",
			StartYear: 1996,
		},
	},
	341227: {
		{
			Address:   "利辛县",
			StartYear: 1996,
			EndYear:   1999,
		},
	},
	341281: {
		{
			Address:   "亳州市",
			StartYear: 1996,
			EndYear:   1999,
		},
	},
	341282: {
		{
			Address:   "界首市",
			StartYear: 1996,
		},
	},
	341300: {
		{
			Address:   "宿州市",
			StartYear: 1998,
		},
	},
	341302: {
		{
			Address:   "埇桥区",
			StartYear: 1998,
		},
	},
	341321: {
		{
			Address:   "砀山县",
			StartYear: 1998,
		},
	},
	341322: {
		{
			Address:   "萧县",
			StartYear: 1998,
		},
	},
	341323: {
		{
			Address:   "灵璧县",
			StartYear: 1998,
		},
	},
	341324: {
		{
			Address:   "泗县",
			StartYear: 1998,
		},
	},
	341400: {
		{
			Address:   "巢湖市",
			StartYear: 1999,
			EndYear:   2010,
		},
	},
	341402: {
		{
			Address:   "居巢区",
			StartYear: 1999,
			EndYear:   2010,
		},
	},
	341421: {
		{
			Address:   "庐江县",
			StartYear: 1999,
			EndYear:   2010,
		},
	},
	341422: {
		{
			Address:   "无为县",
			StartYear: 1999,
			EndYear:   2010,
		},
	},
	341423: {
		{
			Address:   "含山县",
			StartYear: 1999,
			EndYear:   2010,
		},
	},
	341424: {
		{
			Address:   "和县",
			StartYear: 1999,
			EndYear:   2010,
		},
	},
	341500: {
		{
			Address:   "六安市",
			StartYear: 1999,
		},
	},
	341502: {
		{
			Address:   "金安区",
			StartYear: 1999,
		},
	},
	341503: {
		{
			Address:   "裕安区",
			StartYear: 1999,
		},
	},
	341504: {
		{
			Address:   "叶集区",
			StartYear: 2015,
		},
	},
	341521: {
		{
			Address:   "寿县",
			StartYear: 1999,
			EndYear:   2014,
		},
	},
	341522: {
		{
			Address:   "霍邱县",
			StartYear: 1999,
		},
	},
	341523: {
		{
			Address:   "舒城县",
			StartYear: 1999,
		},
	},
	341524: {
		{
			Address:   "金寨县",
			StartYear: 1999,
		},
	},
	341525: {
		{
			Address:   "霍山县",
			StartYear: 1999,
		},
	},
	341600: {
		{
			Address:   "亳州市",
			StartYear: 2000,
		},
	},
	341602: {
		{
			Address:   "谯城区",
			StartYear: 2000,
		},
	},
	341621: {
		{
			Address:   "涡阳县",
			StartYear: 2000,
		},
	},
	341622: {
		{
			Address:   "蒙城县",
			StartYear: 2000,
		},
	},
	341623: {
		{
			Address:   "利辛县",
			StartYear: 2000,
		},
	},
	341700: {
		{
			Address:   "池州市",
			StartYear: 2000,
		},
	},
	341702: {
		{
			Address:   "贵池区",
			StartYear: 2000,
		},
	},
	341721: {
		{
			Address:   "东至县",
			StartYear: 2000,
		},
	},
	341722: {
		{
			Address:   "石台县",
			StartYear: 2000,
		},
	},
	341723: {
		{
			Address:   "青阳县",
			StartYear: 2000,
		},
	},
	341800: {
		{
			Address:   "宣城市",
			StartYear: 2000,
		},
	},
	341802: {
		{
			Address:   "宣州区",
			StartYear: 2000,
		},
	},
	341821: {
		{
			Address:   "郎溪县",
			StartYear: 2000,
		},
	},
	341822: {
		{
			Address:   "广德县",
			StartYear: 2000,
			EndYear:   2018,
		},
	},
	341823: {
		{
			Address:   "泾县",
			StartYear: 2000,
		},
	},
	341824: {
		{
			Address:   "绩溪县",
			StartYear: 2000,
		},
	},
	341825: {
		{
			Address:   "旌德县",
			StartYear: 2000,
		},
	},
	341881: {
		{
			Address:   "宁国市",
			StartYear: 2000,
		},
	},
	341882: {
		{
			Address:   "广德市",
			StartYear: 2019,
		},
	},
	342100: {
		{
			Address: "阜阳地区",
			EndYear: 1995,
		},
	},
	342101: {
		{
			Address: "阜阳市",
			EndYear: 1995,
		},
	},
	342102: {
		{
			Address:   "亳州市",
			StartYear: 1986,
			EndYear:   1995,
		},
	},
	342103: {
		{
			Address:   "界首市",
			StartYear: 1989,
			EndYear:   1995,
		},
	},
	342121: {
		{
			Address: "阜阳县",
			EndYear: 1991,
		},
	},
	342122: {
		{
			Address: "临泉县",
			EndYear: 1995,
		},
	},
	342123: {
		{
			Address: "太和县",
			EndYear: 1995,
		},
	},
	342124: {
		{
			Address: "涡阳县",
			EndYear: 1995,
		},
	},
	342125: {
		{
			Address: "蒙城县",
			EndYear: 1995,
		},
	},
	342126: {
		{
			Address: "亳县",
			EndYear: 1985,
		},
	},
	342127: {
		{
			Address: "阜南县",
			EndYear: 1995,
		},
	},
	342128: {
		{
			Address: "颍上县",
			EndYear: 1995,
		},
	},
	342129: {
		{
			Address: "界首县",
			EndYear: 1988,
		},
	},
	342130: {
		{
			Address: "利辛县",
			EndYear: 1995,
		},
	},
	342200: {
		{
			Address: "宿县地区",
			EndYear: 1997,
		},
	},
	342201: {
		{
			Address: "宿州市",
			EndYear: 1997,
		},
	},
	342221: {
		{
			Address: "砀山县",
			EndYear: 1997,
		},
	},
	342222: {
		{
			Address: "萧县",
			EndYear: 1997,
		},
	},
	342223: {
		{
			Address: "宿县",
			EndYear: 1991,
		},
	},
	342224: {
		{
			Address: "灵璧县",
			EndYear: 1997,
		},
	},
	342225: {
		{
			Address: "泗县",
			EndYear: 1997,
		},
	},
	342226: {
		{
			Address: "怀远县",
			EndYear: 1982,
		},
	},
	342227: {
		{
			Address: "五河县",
			EndYear: 1982,
		},
	},
	342228: {
		{
			Address: "固镇县",
			EndYear: 1982,
		},
	},
	342300: {
		{
			Address: "滁县地区",
			EndYear: 1991,
		},
	},
	342301: {
		{
			Address:   "滁州市",
			StartYear: 1982,
			EndYear:   1991,
		},
	},
	342321: {
		{
			Address: "天长县",
			EndYear: 1991,
		},
	},
	342322: {
		{
			Address: "来安县",
			EndYear: 1991,
		},
	},
	342323: {
		{
			Address: "滁县",
			EndYear: 1981,
		},
	},
	342324: {
		{
			Address: "全椒县",
			EndYear: 1991,
		},
	},
	342325: {
		{
			Address: "定远县",
			EndYear: 1991,
		},
	},
	342326: {
		{
			Address: "凤阳县",
			EndYear: 1991,
		},
	},
	342327: {
		{
			Address: "嘉山县",
			EndYear: 1991,
		},
	},
	342400: {
		{
			Address: "六安地区",
			EndYear: 1998,
		},
	},
	342401: {
		{
			Address: "六安市",
			EndYear: 1998,
		},
	},
	342421: {
		{
			Address: "六安县",
			EndYear: 1991,
		},
	},
	342422: {
		{
			Address: "寿县",
			EndYear: 1998,
		},
	},
	342423: {
		{
			Address: "霍邱县",
			EndYear: 1998,
		},
	},
	342424: {
		{
			Address:   "肥西县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	342425: {
		{
			Address: "舒城县",
			EndYear: 1998,
		},
	},
	342426: {
		{
			Address: "金寨县",
			EndYear: 1998,
		},
	},
	342427: {
		{
			Address: "霍山县",
			EndYear: 1998,
		},
	},
	342500: {
		{
			Address: "宣城地区",
			EndYear: 1999,
		},
	},
	342501: {
		{
			Address:   "宣城市",
			StartYear: 1987,
			EndYear:   1999,
		},
	},
	342502: {
		{
			Address:   "宁国市",
			StartYear: 1997,
			EndYear:   1999,
		},
	},
	342521: {
		{
			Address: "宣城县",
			EndYear: 1986,
		},
	},
	342522: {
		{
			Address: "郎溪县",
			EndYear: 1999,
		},
	},
	342523: {
		{
			Address: "广德县",
			EndYear: 1999,
		},
	},
	342524: {
		{
			Address: "宁国县",
			EndYear: 1996,
		},
	},
	342525: {
		{
			Address:   "当涂县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	342526: {
		{
			Address:   "繁昌县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	342527: {
		{
			Address:   "南陵县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	342528: {
		{
			Address:   "青阳县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	342529: {
		{
			Address: "泾县",
			EndYear: 1999,
		},
	},
	342530: {
		{
			Address:   "旌德县",
			StartYear: 1987,
			EndYear:   1999,
		},
	},
	342531: {
		{
			Address:   "绩溪县",
			StartYear: 1987,
			EndYear:   1999,
		},
	},
	342600: {
		{
			Address: "巢湖地区",
			EndYear: 1998,
		},
	},
	342601: {
		{
			Address:   "巢湖市",
			StartYear: 1982,
			EndYear:   1998,
		},
	},
	342621: {
		{
			Address: "肥东县",
			EndYear: 1982,
		},
	},
	342622: {
		{
			Address: "庐江县",
			EndYear: 1998,
		},
	},
	342623: {
		{
			Address: "无为县",
			EndYear: 1998,
		},
	},
	342624: {
		{
			Address: "巢县",
			EndYear: 1982,
		},
	},
	342625: {
		{
			Address: "含山县",
			EndYear: 1998,
		},
	},
	342626: {
		{
			Address: "和县",
			EndYear: 1998,
		},
	},
	342700: {
		{
			Address: "徽州地区",
			EndYear: 1986,
		},
	},
	342701: {
		{
			Address: "屯溪市",
			EndYear: 1986,
		},
	},
	342721: {
		{
			Address: "绩溪县",
			EndYear: 1986,
		},
	},
	342722: {
		{
			Address: "旌德县",
			EndYear: 1986,
		},
	},
	342723: {
		{
			Address: "歙县",
			EndYear: 1986,
		},
	},
	342724: {
		{
			Address: "休宁县",
			EndYear: 1986,
		},
	},
	342725: {
		{
			Address: "黟县",
			EndYear: 1986,
		},
	},
	342726: {
		{
			Address: "祁门县",
			EndYear: 1986,
		},
	},
	342727: {
		{
			Address: "太平县",
			EndYear: 1982,
		},
	},
	342728: {
		{
			Address: "石台县",
			EndYear: 1986,
		},
	},
	342800: {
		{
			Address: "安庆地区",
			EndYear: 1987,
		},
	},
	342821: {
		{
			Address: "怀宁县",
			EndYear: 1987,
		},
	},
	342822: {
		{
			Address: "桐城县",
			EndYear: 1987,
		},
	},
	342823: {
		{
			Address: "枞阳县",
			EndYear: 1987,
		},
	},
	342824: {
		{
			Address: "潜山县",
			EndYear: 1987,
		},
	},
	342825: {
		{
			Address: "太湖县",
			EndYear: 1987,
		},
	},
	342826: {
		{
			Address: "宿松县",
			EndYear: 1987,
		},
	},
	342827: {
		{
			Address: "望江县",
			EndYear: 1987,
		},
	},
	342828: {
		{
			Address: "岳西县",
			EndYear: 1987,
		},
	},
	342829: {
		{
			Address: "东至县",
			EndYear: 1987,
		},
	},
	342830: {
		{
			Address: "贵池县",
			EndYear: 1987,
		},
	},
	342831: {
		{
			Address:   "石台县",
			StartYear: 1987,
			EndYear:   1987,
		},
	},
	342900: {
		{
			Address:   "池州地区",
			StartYear: 1988,
			EndYear:   1999,
		},
	},
	342901: {
		{
			Address:   "贵池市",
			StartYear: 1988,
			EndYear:   1999,
		},
	},
	342921: {
		{
			Address:   "东至县",
			StartYear: 1988,
			EndYear:   1999,
		},
	},
	342922: {
		{
			Address:   "石台县",
			StartYear: 1988,
			EndYear:   1999,
		},
	},
	342923: {
		{
			Address:   "青阳县",
			StartYear: 1988,
			EndYear:   1999,
		},
	},
	349001: {
		{
			Address:   "天长市",
			StartYear: 1993,
			EndYear:   1994,
		},
	},
	349002: {
		{
			Address:   "明光市",
			StartYear: 1994,
			EndYear:   1994,
		},
	},
	350000: {
		{
			Address: "福建省",
		},
	},
	350100: {
		{
			Address: "福州市",
		},
	},
	350102: {
		{
			Address: "鼓楼区",
		},
	},
	350103: {
		{
			Address: "台江区",
		},
	},
	350104: {
		{
			Address: "仓山区",
		},
	},
	350105: {
		{
			Address: "马尾区",
		},
	},
	350111: {
		{
			Address: "郊区",
			EndYear: 1994,
		},
		{
			Address:   "晋安区",
			StartYear: 1995,
		},
	},
	350112: {
		{
			Address:   "长乐区",
			StartYear: 2017,
		},
	},
	350121: {
		{
			Address: "闽侯县",
		},
	},
	350122: {
		{
			Address:   "连江县",
			StartYear: 1983,
		},
	},
	350123: {
		{
			Address:   "罗源县",
			StartYear: 1983,
		},
	},
	350124: {
		{
			Address:   "闽清县",
			StartYear: 1983,
		},
	},
	350125: {
		{
			Address:   "永泰县",
			StartYear: 1983,
		},
	},
	350126: {
		{
			Address:   "长乐县",
			StartYear: 1983,
			EndYear:   1993,
		},
	},
	350127: {
		{
			Address:   "福清县",
			StartYear: 1983,
			EndYear:   1989,
		},
	},
	350128: {
		{
			Address:   "平潭县",
			StartYear: 1983,
		},
	},
	350181: {
		{
			Address:   "福清市",
			StartYear: 1995,
		},
	},
	350182: {
		{
			Address:   "长乐市",
			StartYear: 1995,
			EndYear:   2016,
		},
	},
	350200: {
		{
			Address: "厦门市",
		},
	},
	350202: {
		{
			Address: "鼓浪屿区",
			EndYear: 2002,
		},
	},
	350203: {
		{
			Address: "思明区",
		},
	},
	350204: {
		{
			Address: "开元区",
			EndYear: 2002,
		},
	},
	350205: {
		{
			Address: "杏林区",
			EndYear: 2002,
		},
		{
			Address:   "海沧区",
			StartYear: 2003,
		},
	},
	350206: {
		{
			Address:   "湖里区",
			StartYear: 1987,
		},
	},
	350211: {
		{
			Address: "郊区",
			EndYear: 1986,
		},
		{
			Address:   "集美区",
			StartYear: 1987,
		},
	},
	350212: {
		{
			Address:   "同安区",
			StartYear: 1996,
		},
	},
	350213: {
		{
			Address:   "翔安区",
			StartYear: 2003,
		},
	},
	350221: {
		{
			Address: "同安县",
			EndYear: 1995,
		},
	},
	350300: {
		{
			Address:   "莆田市",
			StartYear: 1983,
		},
	},
	350302: {
		{
			Address:   "城厢区",
			StartYear: 1983,
		},
	},
	350303: {
		{
			Address:   "涵江区",
			StartYear: 1983,
		},
	},
	350304: {
		{
			Address:   "荔城区",
			StartYear: 2002,
		},
	},
	350305: {
		{
			Address:   "秀屿区",
			StartYear: 2002,
		},
	},
	350321: {
		{
			Address:   "莆田县",
			StartYear: 1983,
			EndYear:   2001,
		},
	},
	350322: {
		{
			Address:   "仙游县",
			StartYear: 1983,
		},
	},
	350400: {
		{
			Address:   "三明市",
			StartYear: 1983,
		},
	},
	350402: {
		{
			Address:   "梅列区",
			StartYear: 1983,
		},
	},
	350403: {
		{
			Address:   "三元区",
			StartYear: 1983,
		},
	},
	350420: {
		{
			Address:   "永安市",
			StartYear: 1985,
			EndYear:   1985,
		},
	},
	350421: {
		{
			Address:   "明溪县",
			StartYear: 1983,
		},
	},
	350422: {
		{
			Address:   "永安县",
			StartYear: 1983,
			EndYear:   1983,
		},
	},
	350423: {
		{
			Address:   "清流县",
			StartYear: 1983,
		},
	},
	350424: {
		{
			Address:   "宁化县",
			StartYear: 1983,
		},
	},
	350425: {
		{
			Address:   "大田县",
			StartYear: 1983,
		},
	},
	350426: {
		{
			Address:   "尤溪县",
			StartYear: 1983,
		},
	},
	350427: {
		{
			Address:   "沙县",
			StartYear: 1983,
		},
	},
	350428: {
		{
			Address:   "将乐县",
			StartYear: 1983,
		},
	},
	350429: {
		{
			Address:   "泰宁县",
			StartYear: 1983,
		},
	},
	350430: {
		{
			Address:   "建宁县",
			StartYear: 1983,
		},
	},
	350481: {
		{
			Address:   "永安市",
			StartYear: 1995,
		},
	},
	350500: {
		{
			Address:   "泉州市",
			StartYear: 1985,
		},
	},
	350501: {
		{
			Address:   "永安市",
			StartYear: 1984,
			EndYear:   1984,
		},
	},
	350502: {
		{
			Address:   "鲤城区",
			StartYear: 1985,
		},
	},
	350503: {
		{
			Address:   "丰泽区",
			StartYear: 1997,
		},
	},
	350504: {
		{
			Address:   "洛江区",
			StartYear: 1997,
		},
	},
	350505: {
		{
			Address:   "泉港区",
			StartYear: 2000,
		},
	},
	350521: {
		{
			Address:   "惠安县",
			StartYear: 1985,
		},
	},
	350522: {
		{
			Address:   "晋江县",
			StartYear: 1985,
			EndYear:   1991,
		},
	},
	350523: {
		{
			Address:   "南安县",
			StartYear: 1985,
			EndYear:   1992,
		},
	},
	350524: {
		{
			Address:   "安溪县",
			StartYear: 1985,
		},
	},
	350525: {
		{
			Address:   "永春县",
			StartYear: 1985,
		},
	},
	350526: {
		{
			Address:   "德化县",
			StartYear: 1985,
		},
	},
	350527: {
		{
			Address:   "金门县",
			StartYear: 1985,
		},
	},
	350581: {
		{
			Address:   "石狮市",
			StartYear: 1995,
		},
	},
	350582: {
		{
			Address:   "晋江市",
			StartYear: 1995,
		},
	},
	350583: {
		{
			Address:   "南安市",
			StartYear: 1995,
		},
	},
	350600: {
		{
			Address:   "漳州市",
			StartYear: 1985,
		},
	},
	350602: {
		{
			Address:   "芗城区",
			StartYear: 1985,
		},
	},
	350603: {
		{
			Address:   "龙文区",
			StartYear: 1996,
		},
	},
	350621: {
		{
			Address:   "龙海县",
			StartYear: 1985,
			EndYear:   1992,
		},
	},
	350622: {
		{
			Address:   "云霄县",
			StartYear: 1985,
		},
	},
	350623: {
		{
			Address:   "漳浦县",
			StartYear: 1985,
		},
	},
	350624: {
		{
			Address:   "诏安县",
			StartYear: 1985,
		},
	},
	350625: {
		{
			Address:   "长泰县",
			StartYear: 1985,
		},
	},
	350626: {
		{
			Address:   "东山县",
			StartYear: 1985,
		},
	},
	350627: {
		{
			Address:   "南靖县",
			StartYear: 1985,
		},
	},
	350628: {
		{
			Address:   "平和县",
			StartYear: 1985,
		},
	},
	350629: {
		{
			Address:   "华安县",
			StartYear: 1985,
		},
	},
	350681: {
		{
			Address:   "龙海市",
			StartYear: 1995,
		},
	},
	350700: {
		{
			Address:   "南平市",
			StartYear: 1994,
		},
	},
	350702: {
		{
			Address:   "延平区",
			StartYear: 1994,
		},
	},
	350703: {
		{
			Address:   "建阳区",
			StartYear: 2014,
		},
	},
	350721: {
		{
			Address:   "顺昌县",
			StartYear: 1994,
		},
	},
	350722: {
		{
			Address:   "浦城县",
			StartYear: 1994,
		},
	},
	350723: {
		{
			Address:   "光泽县",
			StartYear: 1994,
		},
	},
	350724: {
		{
			Address:   "松溪县",
			StartYear: 1994,
		},
	},
	350725: {
		{
			Address:   "政和县",
			StartYear: 1994,
		},
	},
	350781: {
		{
			Address:   "邵武市",
			StartYear: 1995,
		},
	},
	350782: {
		{
			Address:   "武夷山市",
			StartYear: 1995,
		},
	},
	350783: {
		{
			Address:   "建瓯市",
			StartYear: 1995,
		},
	},
	350784: {
		{
			Address:   "建阳市",
			StartYear: 1995,
			EndYear:   2013,
		},
	},
	350800: {
		{
			Address:   "龙岩市",
			StartYear: 1996,
		},
	},
	350802: {
		{
			Address:   "新罗区",
			StartYear: 1996,
		},
	},
	350803: {
		{
			Address:   "永定区",
			StartYear: 2014,
		},
	},
	350821: {
		{
			Address:   "长汀县",
			StartYear: 1996,
		},
	},
	350822: {
		{
			Address:   "永定县",
			StartYear: 1996,
			EndYear:   2013,
		},
	},
	350823: {
		{
			Address:   "上杭县",
			StartYear: 1996,
		},
	},
	350824: {
		{
			Address:   "武平县",
			StartYear: 1996,
		},
	},
	350825: {
		{
			Address:   "连城县",
			StartYear: 1996,
		},
	},
	350881: {
		{
			Address:   "漳平市",
			StartYear: 1996,
		},
	},
	350900: {
		{
			Address:   "宁德市",
			StartYear: 1999,
		},
	},
	350902: {
		{
			Address:   "蕉城区",
			StartYear: 1999,
		},
	},
	350921: {
		{
			Address:   "霞浦县",
			StartYear: 1999,
		},
	},
	350922: {
		{
			Address:   "古田县",
			StartYear: 1999,
		},
	},
	350923: {
		{
			Address:   "屏南县",
			StartYear: 1999,
		},
	},
	350924: {
		{
			Address:   "寿宁县",
			StartYear: 1999,
		},
	},
	350925: {
		{
			Address:   "周宁县",
			StartYear: 1999,
		},
	},
	350926: {
		{
			Address:   "柘荣县",
			StartYear: 1999,
		},
	},
	350981: {
		{
			Address:   "福安市",
			StartYear: 1999,
		},
	},
	350982: {
		{
			Address:   "福鼎市",
			StartYear: 1999,
		},
	},
	352100: {
		{
			Address:   "建阳地区",
			StartYear: 1982,
			EndYear:   1987,
		},
		{
			Address:   "南平地区",
			StartYear: 1988,
			EndYear:   1993,
		},
	},
	352101: {
		{
			Address:   "南平市",
			StartYear: 1982,
			EndYear:   1993,
		},
	},
	352102: {
		{
			Address:   "邵武市",
			StartYear: 1983,
			EndYear:   1993,
		},
	},
	352103: {
		{
			Address:   "武夷山市",
			StartYear: 1989,
			EndYear:   1993,
		},
	},
	352104: {
		{
			Address:   "建瓯市",
			StartYear: 1992,
			EndYear:   1993,
		},
	},
	352121: {
		{
			Address:   "顺昌县",
			StartYear: 1982,
			EndYear:   1993,
		},
	},
	352122: {
		{
			Address:   "建阳县",
			StartYear: 1982,
			EndYear:   1993,
		},
	},
	352123: {
		{
			Address:   "建瓯县",
			StartYear: 1982,
			EndYear:   1991,
		},
	},
	352124: {
		{
			Address:   "浦城县",
			StartYear: 1982,
			EndYear:   1993,
		},
	},
	352125: {
		{
			Address:   "邵武县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	352126: {
		{
			Address:   "崇安县",
			StartYear: 1982,
			EndYear:   1988,
		},
	},
	352127: {
		{
			Address:   "光泽县",
			StartYear: 1982,
			EndYear:   1993,
		},
	},
	352128: {
		{
			Address:   "松溪县",
			StartYear: 1982,
			EndYear:   1993,
		},
	},
	352129: {
		{
			Address:   "政和县",
			StartYear: 1982,
			EndYear:   1993,
		},
	},
	352200: {
		{
			Address: "建阳地区",
			EndYear: 1981,
		},
		{
			Address:   "宁德地区",
			StartYear: 1982,
			EndYear:   1998,
		},
	},
	352201: {
		{
			Address: "南平市",
			EndYear: 1981,
		},
		{
			Address:   "宁德市",
			StartYear: 1988,
			EndYear:   1998,
		},
	},
	352202: {
		{
			Address:   "福安市",
			StartYear: 1989,
			EndYear:   1998,
		},
	},
	352203: {
		{
			Address:   "福鼎市",
			StartYear: 1995,
			EndYear:   1998,
		},
	},
	352221: {
		{
			Address: "顺昌县",
			EndYear: 1981,
		},
		{
			Address:   "宁德县",
			StartYear: 1982,
			EndYear:   1987,
		},
	},
	352222: {
		{
			Address: "建阳县",
			EndYear: 1981,
		},
		{
			Address:   "连江县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	352223: {
		{
			Address: "建瓯县",
			EndYear: 1981,
		},
		{
			Address:   "罗源县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	352224: {
		{
			Address: "浦城县",
			EndYear: 1981,
		},
		{
			Address:   "福鼎县",
			StartYear: 1982,
			EndYear:   1994,
		},
	},
	352225: {
		{
			Address: "邵武县",
			EndYear: 1981,
		},
		{
			Address:   "霞浦县",
			StartYear: 1982,
			EndYear:   1998,
		},
	},
	352226: {
		{
			Address: "崇安县",
			EndYear: 1981,
		},
		{
			Address:   "福安县",
			StartYear: 1982,
			EndYear:   1988,
		},
	},
	352227: {
		{
			Address: "光泽县",
			EndYear: 1981,
		},
		{
			Address:   "古田县",
			StartYear: 1982,
			EndYear:   1998,
		},
	},
	352228: {
		{
			Address: "松溪县",
			EndYear: 1981,
		},
		{
			Address:   "屏南县",
			StartYear: 1982,
			EndYear:   1998,
		},
	},
	352229: {
		{
			Address: "政和县",
			EndYear: 1981,
		},
		{
			Address:   "寿宁县",
			StartYear: 1982,
			EndYear:   1998,
		},
	},
	352230: {
		{
			Address:   "周宁县",
			StartYear: 1982,
			EndYear:   1998,
		},
	},
	352231: {
		{
			Address:   "柘荣县",
			StartYear: 1982,
			EndYear:   1998,
		},
	},
	352300: {
		{
			Address: "宁德地区",
			EndYear: 1981,
		},
		{
			Address:   "莆田地区",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	352321: {
		{
			Address: "宁德县",
			EndYear: 1981,
		},
		{
			Address:   "闽清县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	352322: {
		{
			Address: "连江县",
			EndYear: 1981,
		},
		{
			Address:   "永泰县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	352323: {
		{
			Address: "罗源县",
			EndYear: 1981,
		},
		{
			Address:   "长乐县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	352324: {
		{
			Address: "福鼎县",
			EndYear: 1981,
		},
		{
			Address:   "福清县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	352325: {
		{
			Address: "霞浦县",
			EndYear: 1981,
		},
		{
			Address:   "平潭县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	352326: {
		{
			Address: "福安县",
			EndYear: 1981,
		},
		{
			Address:   "莆田县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	352327: {
		{
			Address: "古田县",
			EndYear: 1981,
		},
		{
			Address:   "仙游县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	352328: {
		{
			Address: "屏南县",
			EndYear: 1981,
		},
	},
	352329: {
		{
			Address: "寿宁县",
			EndYear: 1981,
		},
	},
	352330: {
		{
			Address: "周宁县",
			EndYear: 1981,
		},
	},
	352331: {
		{
			Address: "柘荣县",
			EndYear: 1981,
		},
	},
	352400: {
		{
			Address: "莆田地区",
			EndYear: 1981,
		},
		{
			Address:   "晋江地区",
			StartYear: 1982,
			EndYear:   1984,
		},
	},
	352401: {
		{
			Address:   "泉州市",
			StartYear: 1982,
			EndYear:   1984,
		},
	},
	352421: {
		{
			Address: "闽清县",
			EndYear: 1981,
		},
		{
			Address:   "惠安县",
			StartYear: 1982,
			EndYear:   1984,
		},
	},
	352422: {
		{
			Address: "永泰县",
			EndYear: 1981,
		},
		{
			Address:   "晋江县",
			StartYear: 1982,
			EndYear:   1984,
		},
	},
	352423: {
		{
			Address: "长乐县",
			EndYear: 1981,
		},
		{
			Address:   "南安县",
			StartYear: 1982,
			EndYear:   1984,
		},
	},
	352424: {
		{
			Address: "福清县",
			EndYear: 1981,
		},
		{
			Address:   "安溪县",
			StartYear: 1982,
			EndYear:   1984,
		},
	},
	352425: {
		{
			Address: "平潭县",
			EndYear: 1981,
		},
		{
			Address:   "永春县",
			StartYear: 1982,
			EndYear:   1984,
		},
	},
	352426: {
		{
			Address: "莆田县",
			EndYear: 1981,
		},
		{
			Address:   "德化县",
			StartYear: 1982,
			EndYear:   1984,
		},
	},
	352427: {
		{
			Address: "仙游县",
			EndYear: 1981,
		},
		{
			Address:   "金门县",
			StartYear: 1982,
			EndYear:   1984,
		},
	},
	352500: {
		{
			Address: "晋江地区",
			EndYear: 1981,
		},
		{
			Address:   "龙溪地区",
			StartYear: 1982,
			EndYear:   1984,
		},
	},
	352501: {
		{
			Address: "泉州市",
			EndYear: 1981,
		},
		{
			Address:   "漳州市",
			StartYear: 1982,
			EndYear:   1984,
		},
	},
	352521: {
		{
			Address: "惠安县",
			EndYear: 1981,
		},
		{
			Address:   "龙海县",
			StartYear: 1982,
			EndYear:   1984,
		},
	},
	352522: {
		{
			Address: "晋江县",
			EndYear: 1981,
		},
		{
			Address:   "云霄县",
			StartYear: 1982,
			EndYear:   1984,
		},
	},
	352523: {
		{
			Address: "南安县",
			EndYear: 1981,
		},
		{
			Address:   "漳浦县",
			StartYear: 1982,
			EndYear:   1984,
		},
	},
	352524: {
		{
			Address: "安溪县",
			EndYear: 1981,
		},
		{
			Address:   "诏安县",
			StartYear: 1982,
			EndYear:   1984,
		},
	},
	352525: {
		{
			Address: "永春县",
			EndYear: 1981,
		},
		{
			Address:   "长泰县",
			StartYear: 1982,
			EndYear:   1984,
		},
	},
	352526: {
		{
			Address: "德化县",
			EndYear: 1981,
		},
		{
			Address:   "东山县",
			StartYear: 1982,
			EndYear:   1984,
		},
	},
	352527: {
		{
			Address: "金门县",
			EndYear: 1981,
		},
		{
			Address:   "南靖县",
			StartYear: 1982,
			EndYear:   1984,
		},
	},
	352528: {
		{
			Address:   "平和县",
			StartYear: 1982,
			EndYear:   1984,
		},
	},
	352529: {
		{
			Address:   "华安县",
			StartYear: 1982,
			EndYear:   1984,
		},
	},
	352600: {
		{
			Address: "龙溪地区",
			EndYear: 1981,
		},
		{
			Address:   "龙岩地区",
			StartYear: 1982,
			EndYear:   1995,
		},
	},
	352601: {
		{
			Address: "漳州市",
			EndYear: 1981,
		},
		{
			Address:   "龙岩市",
			StartYear: 1982,
			EndYear:   1995,
		},
	},
	352602: {
		{
			Address:   "漳平市",
			StartYear: 1990,
			EndYear:   1995,
		},
	},
	352621: {
		{
			Address: "龙海县",
			EndYear: 1981,
		},
	},
	352622: {
		{
			Address: "云霄县",
			EndYear: 1981,
		},
		{
			Address:   "长汀县",
			StartYear: 1982,
			EndYear:   1995,
		},
	},
	352623: {
		{
			Address: "漳浦县",
			EndYear: 1981,
		},
		{
			Address:   "永定县",
			StartYear: 1982,
			EndYear:   1995,
		},
	},
	352624: {
		{
			Address: "诏安县",
			EndYear: 1981,
		},
		{
			Address:   "上杭县",
			StartYear: 1982,
			EndYear:   1995,
		},
	},
	352625: {
		{
			Address: "长泰县",
			EndYear: 1981,
		},
		{
			Address:   "武平县",
			StartYear: 1982,
			EndYear:   1995,
		},
	},
	352626: {
		{
			Address: "东山县",
			EndYear: 1981,
		},
		{
			Address:   "漳平县",
			StartYear: 1982,
			EndYear:   1989,
		},
	},
	352627: {
		{
			Address: "南靖县",
			EndYear: 1981,
		},
		{
			Address:   "连城县",
			StartYear: 1982,
			EndYear:   1995,
		},
	},
	352628: {
		{
			Address: "平和县",
			EndYear: 1981,
		},
	},
	352629: {
		{
			Address: "华安县",
			EndYear: 1981,
		},
	},
	352700: {
		{
			Address: "龙岩地区",
			EndYear: 1981,
		},
		{
			Address:   "三明地区",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	352701: {
		{
			Address:   "龙岩市",
			StartYear: 1981,
			EndYear:   1981,
		},
		{
			Address:   "三明市",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	352721: {
		{
			Address: "龙岩县",
			EndYear: 1980,
		},
		{
			Address:   "明溪县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	352722: {
		{
			Address: "长汀县",
			EndYear: 1981,
		},
		{
			Address:   "永安县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	352723: {
		{
			Address: "永定县",
			EndYear: 1981,
		},
		{
			Address:   "清流县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	352724: {
		{
			Address: "上杭县",
			EndYear: 1981,
		},
		{
			Address:   "宁化县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	352725: {
		{
			Address: "武平县",
			EndYear: 1981,
		},
		{
			Address:   "大田县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	352726: {
		{
			Address: "漳平县",
			EndYear: 1981,
		},
		{
			Address:   "尤溪县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	352727: {
		{
			Address: "连城县",
			EndYear: 1981,
		},
		{
			Address:   "沙县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	352728: {
		{
			Address:   "将乐县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	352729: {
		{
			Address:   "泰宁县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	352730: {
		{
			Address:   "建宁县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	352800: {
		{
			Address: "三明地区",
			EndYear: 1981,
		},
	},
	352801: {
		{
			Address: "三明市",
			EndYear: 1981,
		},
	},
	352821: {
		{
			Address: "明溪县",
			EndYear: 1981,
		},
	},
	352822: {
		{
			Address: "永安县",
			EndYear: 1981,
		},
	},
	352823: {
		{
			Address: "清流县",
			EndYear: 1981,
		},
	},
	352824: {
		{
			Address: "宁化县",
			EndYear: 1981,
		},
	},
	352825: {
		{
			Address: "大田县",
			EndYear: 1981,
		},
	},
	352826: {
		{
			Address: "尤溪县",
			EndYear: 1981,
		},
	},
	352827: {
		{
			Address: "沙县",
			EndYear: 1981,
		},
	},
	352828: {
		{
			Address: "将乐县",
			EndYear: 1981,
		},
	},
	352829: {
		{
			Address: "泰宁县",
			EndYear: 1981,
		},
	},
	352830: {
		{
			Address: "建宁县",
			EndYear: 1981,
		},
	},
	359001: {
		{
			Address:   "永安市",
			StartYear: 1986,
			EndYear:   1994,
		},
	},
	359002: {
		{
			Address:   "石狮市",
			StartYear: 1987,
			EndYear:   1994,
		},
	},
	359003: {
		{
			Address:   "福清市",
			StartYear: 1990,
			EndYear:   1994,
		},
	},
	359004: {
		{
			Address:   "晋江市",
			StartYear: 1992,
			EndYear:   1994,
		},
	},
	359005: {
		{
			Address:   "南安市",
			StartYear: 1993,
			EndYear:   1994,
		},
	},
	359006: {
		{
			Address:   "龙海市",
			StartYear: 1993,
			EndYear:   1994,
		},
	},
	359007: {
		{
			Address:   "邵武市",
			StartYear: 1994,
			EndYear:   1994,
		},
	},
	359008: {
		{
			Address:   "武夷山市",
			StartYear: 1994,
			EndYear:   1994,
		},
	},
	359009: {
		{
			Address:   "建瓯市",
			StartYear: 1994,
			EndYear:   1994,
		},
	},
	359010: {
		{
			Address:   "建阳市",
			StartYear: 1994,
			EndYear:   1994,
		},
	},
	359011: {
		{
			Address:   "长乐市",
			StartYear: 1994,
			EndYear:   1994,
		},
	},
	360000: {
		{
			Address: "江西省",
		},
	},
	360100: {
		{
			Address: "南昌市",
		},
	},
	360102: {
		{
			Address: "东湖区",
		},
	},
	360103: {
		{
			Address: "西湖区",
		},
	},
	360104: {
		{
			Address: "青云谱区",
		},
	},
	360105: {
		{
			Address: "湾里区",
			EndYear: 2018,
		},
	},
	360111: {
		{
			Address: "郊区",
			EndYear: 2001,
		},
		{
			Address:   "青山湖区",
			StartYear: 2002,
		},
	},
	360112: {
		{
			Address:   "新建区",
			StartYear: 2015,
		},
	},
	360113: {
		{
			Address:   "红谷滩区",
			StartYear: 2019,
		},
	},
	360121: {
		{
			Address: "新建县",
			EndYear: 1981,
		},
		{
			Address:   "南昌县",
			StartYear: 1982,
		},
	},
	360122: {
		{
			Address: "南昌县",
			EndYear: 1981,
		},
		{
			Address:   "新建县",
			StartYear: 1982,
			EndYear:   2014,
		},
	},
	360123: {
		{
			Address:   "安义县",
			StartYear: 1983,
		},
	},
	360124: {
		{
			Address:   "进贤县",
			StartYear: 1983,
		},
	},
	360200: {
		{
			Address: "景德镇市",
		},
	},
	360202: {
		{
			Address: "昌江区",
		},
	},
	360203: {
		{
			Address: "珠山区",
		},
	},
	360211: {
		{
			Address: "鹅湖区",
			EndYear: 1987,
		},
	},
	360212: {
		{
			Address: "蛟潭区",
			EndYear: 1987,
		},
	},
	360221: {
		{
			Address:   "乐平县",
			StartYear: 1983,
			EndYear:   1991,
		},
	},
	360222: {
		{
			Address:   "浮梁县",
			StartYear: 1988,
		},
	},
	360281: {
		{
			Address:   "乐平市",
			StartYear: 1995,
		},
	},
	360300: {
		{
			Address: "萍乡市",
		},
	},
	360302: {
		{
			Address: "城关区",
			EndYear: 1992,
		},
		{
			Address:   "安源区",
			StartYear: 1993,
		},
	},
	360311: {
		{
			Address: "上栗区",
			EndYear: 1996,
		},
	},
	360312: {
		{
			Address: "芦溪区",
			EndYear: 1996,
		},
	},
	360313: {
		{
			Address: "湘东区",
		},
	},
	360321: {
		{
			Address:   "莲花县",
			StartYear: 1992,
		},
	},
	360322: {
		{
			Address:   "上栗县",
			StartYear: 1997,
		},
	},
	360323: {
		{
			Address:   "芦溪县",
			StartYear: 1997,
		},
	},
	360400: {
		{
			Address: "九江市",
		},
	},
	360402: {
		{
			Address: "庐山区",
			EndYear: 2015,
		},
		{
			Address:   "濂溪区",
			StartYear: 2016,
		},
	},
	360403: {
		{
			Address: "浔阳区",
		},
	},
	360404: {
		{
			Address:   "柴桑区",
			StartYear: 2017,
		},
	},
	360411: {
		{
			Address: "郊区",
			EndYear: 1983,
		},
	},
	360421: {
		{
			Address:   "九江县",
			StartYear: 1983,
			EndYear:   2016,
		},
	},
	360422: {
		{
			Address:   "瑞昌县",
			StartYear: 1983,
			EndYear:   1988,
		},
	},
	360423: {
		{
			Address:   "武宁县",
			StartYear: 1983,
		},
	},
	360424: {
		{
			Address:   "修水县",
			StartYear: 1983,
		},
	},
	360425: {
		{
			Address:   "永修县",
			StartYear: 1983,
		},
	},
	360426: {
		{
			Address:   "德安县",
			StartYear: 1983,
		},
	},
	360427: {
		{
			Address:   "星子县",
			StartYear: 1983,
			EndYear:   2015,
		},
	},
	360428: {
		{
			Address:   "都昌县",
			StartYear: 1983,
		},
	},
	360429: {
		{
			Address:   "湖口县",
			StartYear: 1983,
		},
	},
	360430: {
		{
			Address:   "彭泽县",
			StartYear: 1983,
		},
	},
	360481: {
		{
			Address:   "瑞昌市",
			StartYear: 1995,
		},
	},
	360482: {
		{
			Address:   "共青城市",
			StartYear: 2010,
		},
	},
	360483: {
		{
			Address:   "庐山市",
			StartYear: 2016,
		},
	},
	360500: {
		{
			Address:   "新余市",
			StartYear: 1983,
		},
	},
	360502: {
		{
			Address:   "渝水区",
			StartYear: 1983,
		},
	},
	360521: {
		{
			Address: "井冈山",
			EndYear: 1980,
		},
		{
			Address:   "分宜县",
			StartYear: 1983,
		},
	},
	360600: {
		{
			Address:   "鹰潭市",
			StartYear: 1983,
		},
	},
	360602: {
		{
			Address:   "月湖区",
			StartYear: 1983,
		},
	},
	360603: {
		{
			Address:   "余江区",
			StartYear: 2018,
		},
	},
	360621: {
		{
			Address:   "贵溪县",
			StartYear: 1983,
			EndYear:   1995,
		},
	},
	360622: {
		{
			Address:   "余江县",
			StartYear: 1983,
			EndYear:   2017,
		},
	},
	360681: {
		{
			Address:   "贵溪市",
			StartYear: 1996,
		},
	},
	360700: {
		{
			Address:   "赣州市",
			StartYear: 1998,
		},
	},
	360702: {
		{
			Address:   "章贡区",
			StartYear: 1998,
		},
	},
	360703: {
		{
			Address:   "南康区",
			StartYear: 2013,
		},
	},
	360704: {
		{
			Address:   "赣县区",
			StartYear: 2016,
		},
	},
	360721: {
		{
			Address:   "赣县",
			StartYear: 1998,
			EndYear:   2015,
		},
	},
	360722: {
		{
			Address:   "信丰县",
			StartYear: 1998,
		},
	},
	360723: {
		{
			Address:   "大余县",
			StartYear: 1998,
		},
	},
	360724: {
		{
			Address:   "上犹县",
			StartYear: 1998,
		},
	},
	360725: {
		{
			Address:   "崇义县",
			StartYear: 1998,
		},
	},
	360726: {
		{
			Address:   "安远县",
			StartYear: 1998,
		},
	},
	360727: {
		{
			Address:   "龙南县",
			StartYear: 1998,
			EndYear:   2019,
		},
	},
	360728: {
		{
			Address:   "定南县",
			StartYear: 1998,
		},
	},
	360729: {
		{
			Address:   "全南县",
			StartYear: 1998,
		},
	},
	360730: {
		{
			Address:   "宁都县",
			StartYear: 1998,
		},
	},
	360731: {
		{
			Address:   "于都县",
			StartYear: 1998,
		},
	},
	360732: {
		{
			Address:   "兴国县",
			StartYear: 1998,
		},
	},
	360733: {
		{
			Address:   "会昌县",
			StartYear: 1998,
		},
	},
	360734: {
		{
			Address:   "寻乌县",
			StartYear: 1998,
		},
	},
	360735: {
		{
			Address:   "石城县",
			StartYear: 1998,
		},
	},
	360781: {
		{
			Address:   "瑞金市",
			StartYear: 1998,
		},
	},
	360782: {
		{
			Address:   "南康市",
			StartYear: 1998,
			EndYear:   2012,
		},
	},
	360783: {
		{
			Address:   "龙南市",
			StartYear: 2020,
		},
	},
	360800: {
		{
			Address:   "吉安市",
			StartYear: 2000,
		},
	},
	360802: {
		{
			Address:   "吉州区",
			StartYear: 2000,
		},
	},
	360803: {
		{
			Address:   "青原区",
			StartYear: 2000,
		},
	},
	360821: {
		{
			Address:   "吉安县",
			StartYear: 2000,
		},
	},
	360822: {
		{
			Address:   "吉水县",
			StartYear: 2000,
		},
	},
	360823: {
		{
			Address:   "峡江县",
			StartYear: 2000,
		},
	},
	360824: {
		{
			Address:   "新干县",
			StartYear: 2000,
		},
	},
	360825: {
		{
			Address:   "永丰县",
			StartYear: 2000,
		},
	},
	360826: {
		{
			Address:   "泰和县",
			StartYear: 2000,
		},
	},
	360827: {
		{
			Address:   "遂川县",
			StartYear: 2000,
		},
	},
	360828: {
		{
			Address:   "万安县",
			StartYear: 2000,
		},
	},
	360829: {
		{
			Address:   "安福县",
			StartYear: 2000,
		},
	},
	360830: {
		{
			Address:   "永新县",
			StartYear: 2000,
		},
	},
	360881: {
		{
			Address:   "井冈山市",
			StartYear: 2000,
		},
	},
	360900: {
		{
			Address:   "宜春市",
			StartYear: 2000,
		},
	},
	360902: {
		{
			Address:   "袁州区",
			StartYear: 2000,
		},
	},
	360921: {
		{
			Address:   "奉新县",
			StartYear: 2000,
		},
	},
	360922: {
		{
			Address:   "万载县",
			StartYear: 2000,
		},
	},
	360923: {
		{
			Address:   "上高县",
			StartYear: 2000,
		},
	},
	360924: {
		{
			Address:   "宜丰县",
			StartYear: 2000,
		},
	},
	360925: {
		{
			Address:   "靖安县",
			StartYear: 2000,
		},
	},
	360926: {
		{
			Address:   "铜鼓县",
			StartYear: 2000,
		},
	},
	360981: {
		{
			Address:   "丰城市",
			StartYear: 2000,
		},
	},
	360982: {
		{
			Address:   "樟树市",
			StartYear: 2000,
		},
	},
	360983: {
		{
			Address:   "高安市",
			StartYear: 2000,
		},
	},
	361000: {
		{
			Address:   "抚州市",
			StartYear: 2000,
		},
	},
	361002: {
		{
			Address:   "临川区",
			StartYear: 2000,
		},
	},
	361003: {
		{
			Address:   "东乡区",
			StartYear: 2016,
		},
	},
	361021: {
		{
			Address:   "南城县",
			StartYear: 2000,
		},
	},
	361022: {
		{
			Address:   "黎川县",
			StartYear: 2000,
		},
	},
	361023: {
		{
			Address:   "南丰县",
			StartYear: 2000,
		},
	},
	361024: {
		{
			Address:   "崇仁县",
			StartYear: 2000,
		},
	},
	361025: {
		{
			Address:   "乐安县",
			StartYear: 2000,
		},
	},
	361026: {
		{
			Address:   "宜黄县",
			StartYear: 2000,
		},
	},
	361027: {
		{
			Address:   "金溪县",
			StartYear: 2000,
		},
	},
	361028: {
		{
			Address:   "资溪县",
			StartYear: 2000,
		},
	},
	361029: {
		{
			Address:   "东乡县",
			StartYear: 2000,
			EndYear:   2015,
		},
	},
	361030: {
		{
			Address:   "广昌县",
			StartYear: 2000,
		},
	},
	361100: {
		{
			Address:   "上饶市",
			StartYear: 2000,
		},
	},
	361102: {
		{
			Address:   "信州区",
			StartYear: 2000,
		},
	},
	361103: {
		{
			Address:   "广丰区",
			StartYear: 2015,
		},
	},
	361104: {
		{
			Address:   "广信区",
			StartYear: 2019,
		},
	},
	361121: {
		{
			Address:   "上饶县",
			StartYear: 2000,
			EndYear:   2018,
		},
	},
	361122: {
		{
			Address:   "广丰县",
			StartYear: 2000,
			EndYear:   2014,
		},
	},
	361123: {
		{
			Address:   "玉山县",
			StartYear: 2000,
		},
	},
	361124: {
		{
			Address:   "铅山县",
			StartYear: 2000,
		},
	},
	361125: {
		{
			Address:   "横峰县",
			StartYear: 2000,
		},
	},
	361126: {
		{
			Address:   "弋阳县",
			StartYear: 2000,
		},
	},
	361127: {
		{
			Address:   "余干县",
			StartYear: 2000,
		},
	},
	361128: {
		{
			Address:   "波阳县",
			StartYear: 2000,
			EndYear:   2002,
		},
		{
			Address:   "鄱阳县",
			StartYear: 2003,
		},
	},
	361129: {
		{
			Address:   "万年县",
			StartYear: 2000,
		},
	},
	361130: {
		{
			Address:   "婺源县",
			StartYear: 2000,
		},
	},
	361181: {
		{
			Address:   "德兴市",
			StartYear: 2000,
		},
	},
	362100: {
		{
			Address: "九江地区",
			EndYear: 1981,
		},
		{
			Address:   "赣州地区",
			StartYear: 1982,
			EndYear:   1997,
		},
	},
	362101: {
		{
			Address:   "赣州市",
			StartYear: 1982,
			EndYear:   1997,
		},
	},
	362102: {
		{
			Address:   "瑞金市",
			StartYear: 1994,
			EndYear:   1997,
		},
	},
	362103: {
		{
			Address:   "南康市",
			StartYear: 1995,
			EndYear:   1997,
		},
	},
	362121: {
		{
			Address: "九江县",
			EndYear: 1981,
		},
		{
			Address:   "赣县",
			StartYear: 1982,
			EndYear:   1997,
		},
	},
	362122: {
		{
			Address: "永修县",
			EndYear: 1981,
		},
		{
			Address:   "南康县",
			StartYear: 1982,
			EndYear:   1994,
		},
	},
	362123: {
		{
			Address: "彭泽县",
			EndYear: 1981,
		},
		{
			Address:   "信丰县",
			StartYear: 1982,
			EndYear:   1997,
		},
	},
	362124: {
		{
			Address: "德安县",
			EndYear: 1981,
		},
		{
			Address:   "大余县",
			StartYear: 1982,
			EndYear:   1997,
		},
	},
	362125: {
		{
			Address: "湖口县",
			EndYear: 1981,
		},
		{
			Address:   "上犹县",
			StartYear: 1982,
			EndYear:   1997,
		},
	},
	362126: {
		{
			Address: "瑞昌县",
			EndYear: 1981,
		},
		{
			Address:   "崇义县",
			StartYear: 1982,
			EndYear:   1997,
		},
	},
	362127: {
		{
			Address: "都昌县",
			EndYear: 1981,
		},
		{
			Address:   "安远县",
			StartYear: 1982,
			EndYear:   1997,
		},
	},
	362128: {
		{
			Address: "武宁县",
			EndYear: 1981,
		},
		{
			Address:   "龙南县",
			StartYear: 1982,
			EndYear:   1997,
		},
	},
	362129: {
		{
			Address: "星子县",
			EndYear: 1981,
		},
		{
			Address:   "定南县",
			StartYear: 1982,
			EndYear:   1997,
		},
	},
	362130: {
		{
			Address: "修水县",
			EndYear: 1981,
		},
		{
			Address:   "全南县",
			StartYear: 1982,
			EndYear:   1997,
		},
	},
	362131: {
		{
			Address:   "宁都县",
			StartYear: 1982,
			EndYear:   1997,
		},
	},
	362132: {
		{
			Address:   "于都县",
			StartYear: 1982,
			EndYear:   1997,
		},
	},
	362133: {
		{
			Address:   "兴国县",
			StartYear: 1982,
			EndYear:   1997,
		},
	},
	362134: {
		{
			Address:   "瑞金县",
			StartYear: 1982,
			EndYear:   1993,
		},
	},
	362135: {
		{
			Address:   "会昌县",
			StartYear: 1982,
			EndYear:   1997,
		},
	},
	362136: {
		{
			Address:   "寻乌县",
			StartYear: 1982,
			EndYear:   1997,
		},
	},
	362137: {
		{
			Address:   "石城县",
			StartYear: 1982,
			EndYear:   1997,
		},
	},
	362138: {
		{
			Address:   "广昌县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	362200: {
		{
			Address: "上饶地区",
			EndYear: 1981,
		},
		{
			Address:   "宜春地区",
			StartYear: 1982,
			EndYear:   1999,
		},
	},
	362201: {
		{
			Address: "上饶市",
			EndYear: 1981,
		},
		{
			Address:   "宜春市",
			StartYear: 1982,
			EndYear:   1999,
		},
	},
	362202: {
		{
			Address: "鹰潭市",
			EndYear: 1981,
		},
		{
			Address:   "丰城市",
			StartYear: 1988,
			EndYear:   1999,
		},
	},
	362203: {
		{
			Address:   "樟树市",
			StartYear: 1988,
			EndYear:   1999,
		},
	},
	362204: {
		{
			Address:   "高安市",
			StartYear: 1993,
			EndYear:   1999,
		},
	},
	362221: {
		{
			Address: "上饶县",
			EndYear: 1981,
		},
		{
			Address:   "丰城县",
			StartYear: 1982,
			EndYear:   1987,
		},
	},
	362222: {
		{
			Address: "贵溪县",
			EndYear: 1981,
		},
		{
			Address:   "高安县",
			StartYear: 1982,
			EndYear:   1992,
		},
	},
	362223: {
		{
			Address: "婺源县",
			EndYear: 1981,
		},
		{
			Address:   "清江县",
			StartYear: 1982,
			EndYear:   1987,
		},
	},
	362224: {
		{
			Address: "余江县",
			EndYear: 1981,
		},
		{
			Address:   "新余县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	362225: {
		{
			Address: "德兴县",
			EndYear: 1981,
		},
		{
			Address:   "宜春县",
			StartYear: 1982,
			EndYear:   1984,
		},
	},
	362226: {
		{
			Address: "万年县",
			EndYear: 1981,
		},
		{
			Address:   "奉新县",
			StartYear: 1982,
			EndYear:   1999,
		},
	},
	362227: {
		{
			Address: "玉山县",
			EndYear: 1981,
		},
		{
			Address:   "万载县",
			StartYear: 1982,
			EndYear:   1999,
		},
	},
	362228: {
		{
			Address: "乐平县",
			EndYear: 1981,
		},
		{
			Address:   "上高县",
			StartYear: 1982,
			EndYear:   1999,
		},
	},
	362229: {
		{
			Address: "广丰县",
			EndYear: 1981,
		},
		{
			Address:   "宜丰县",
			StartYear: 1982,
			EndYear:   1999,
		},
	},
	362230: {
		{
			Address: "波阳县",
			EndYear: 1981,
		},
		{
			Address:   "分宜县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	362231: {
		{
			Address: "铅山县",
			EndYear: 1981,
		},
		{
			Address:   "安义县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	362232: {
		{
			Address: "余干县",
			EndYear: 1981,
		},
		{
			Address:   "靖安县",
			StartYear: 1982,
			EndYear:   1999,
		},
	},
	362233: {
		{
			Address: "横峰县",
			EndYear: 1981,
		},
		{
			Address:   "铜鼓县",
			StartYear: 1982,
			EndYear:   1999,
		},
	},
	362234: {
		{
			Address: "弋阳县",
			EndYear: 1981,
		},
	},
	362300: {
		{
			Address: "宜春地区",
			EndYear: 1981,
		},
		{
			Address:   "上饶地区",
			StartYear: 1982,
			EndYear:   1999,
		},
	},
	362301: {
		{
			Address: "宜春市",
			EndYear: 1981,
		},
		{
			Address:   "上饶市",
			StartYear: 1982,
			EndYear:   1999,
		},
	},
	362302: {
		{
			Address:   "鹰潭市",
			StartYear: 1982,
			EndYear:   1982,
		},
		{
			Address:   "德兴市",
			StartYear: 1990,
			EndYear:   1999,
		},
	},
	362321: {
		{
			Address: "宜春县",
			EndYear: 1981,
		},
		{
			Address:   "上饶县",
			StartYear: 1982,
			EndYear:   1999,
		},
	},
	362322: {
		{
			Address: "高安县",
			EndYear: 1981,
		},
		{
			Address:   "广丰县",
			StartYear: 1982,
			EndYear:   1999,
		},
	},
	362323: {
		{
			Address: "万载县",
			EndYear: 1981,
		},
		{
			Address:   "玉山县",
			StartYear: 1982,
			EndYear:   1999,
		},
	},
	362324: {
		{
			Address: "丰城县",
			EndYear: 1981,
		},
		{
			Address:   "铅山县",
			StartYear: 1982,
			EndYear:   1999,
		},
	},
	362325: {
		{
			Address: "铜鼓县",
			EndYear: 1981,
		},
		{
			Address:   "横峰县",
			StartYear: 1982,
			EndYear:   1999,
		},
	},
	362326: {
		{
			Address: "清江县",
			EndYear: 1981,
		},
		{
			Address:   "弋阳县",
			StartYear: 1982,
			EndYear:   1999,
		},
	},
	362327: {
		{
			Address: "宜丰县",
			EndYear: 1981,
		},
		{
			Address:   "贵溪县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	362328: {
		{
			Address: "新余县",
			EndYear: 1981,
		},
		{
			Address:   "余江县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	362329: {
		{
			Address: "上高县",
			EndYear: 1981,
		},
		{
			Address:   "余干县",
			StartYear: 1982,
			EndYear:   1999,
		},
	},
	362330: {
		{
			Address: "分宜县",
			EndYear: 1981,
		},
		{
			Address:   "波阳县",
			StartYear: 1982,
			EndYear:   1999,
		},
	},
	362331: {
		{
			Address: "安义县",
			EndYear: 1981,
		},
		{
			Address:   "万年县",
			StartYear: 1982,
			EndYear:   1999,
		},
	},
	362332: {
		{
			Address: "靖安县",
			EndYear: 1981,
		},
		{
			Address:   "乐平县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	362333: {
		{
			Address: "奉新县",
			EndYear: 1981,
		},
		{
			Address:   "德兴县",
			StartYear: 1982,
			EndYear:   1989,
		},
	},
	362334: {
		{
			Address:   "婺源县",
			StartYear: 1982,
			EndYear:   1999,
		},
	},
	362400: {
		{
			Address: "抚州地区",
			EndYear: 1981,
		},
		{
			Address:   "吉安地区",
			StartYear: 1982,
			EndYear:   1999,
		},
	},
	362401: {
		{
			Address: "抚州市",
			EndYear: 1981,
		},
		{
			Address:   "吉安市",
			StartYear: 1982,
			EndYear:   1999,
		},
	},
	362402: {
		{
			Address:   "井冈山市",
			StartYear: 1984,
			EndYear:   1999,
		},
	},
	362421: {
		{
			Address: "临川县",
			EndYear: 1981,
		},
		{
			Address:   "吉安县",
			StartYear: 1982,
			EndYear:   1999,
		},
	},
	362422: {
		{
			Address: "宜黄县",
			EndYear: 1981,
		},
		{
			Address:   "吉水县",
			StartYear: 1982,
			EndYear:   1999,
		},
	},
	362423: {
		{
			Address: "金溪县",
			EndYear: 1981,
		},
		{
			Address:   "峡江县",
			StartYear: 1982,
			EndYear:   1999,
		},
	},
	362424: {
		{
			Address: "崇仁县",
			EndYear: 1981,
		},
		{
			Address:   "新干县",
			StartYear: 1982,
			EndYear:   1999,
		},
	},
	362425: {
		{
			Address: "资溪县",
			EndYear: 1981,
		},
		{
			Address:   "永丰县",
			StartYear: 1982,
			EndYear:   1999,
		},
	},
	362426: {
		{
			Address: "乐安县",
			EndYear: 1981,
		},
		{
			Address:   "泰和县",
			StartYear: 1982,
			EndYear:   1999,
		},
	},
	362427: {
		{
			Address: "黎川县",
			EndYear: 1981,
		},
		{
			Address:   "遂川县",
			StartYear: 1982,
			EndYear:   1999,
		},
	},
	362428: {
		{
			Address: "东乡县",
			EndYear: 1981,
		},
		{
			Address:   "万安县",
			StartYear: 1982,
			EndYear:   1999,
		},
	},
	362429: {
		{
			Address: "南丰县",
			EndYear: 1981,
		},
		{
			Address:   "安福县",
			StartYear: 1982,
			EndYear:   1999,
		},
	},
	362430: {
		{
			Address: "进贤县",
			EndYear: 1981,
		},
		{
			Address:   "永新县",
			StartYear: 1982,
			EndYear:   1999,
		},
	},
	362431: {
		{
			Address: "南城县",
			EndYear: 1981,
		},
		{
			Address:   "莲花县",
			StartYear: 1982,
			EndYear:   1991,
		},
	},
	362432: {
		{
			Address:   "宁冈县",
			StartYear: 1982,
			EndYear:   1999,
		},
	},
	362433: {
		{
			Address:   "井冈山县",
			StartYear: 1982,
			EndYear:   1983,
		},
	},
	362500: {
		{
			Address: "吉安地区",
			EndYear: 1981,
		},
		{
			Address:   "抚州地区",
			StartYear: 1982,
			EndYear:   1999,
		},
	},
	362501: {
		{
			Address: "吉安市",
			EndYear: 1981,
		},
		{
			Address:   "抚州市",
			StartYear: 1982,
			EndYear:   1986,
		},
		{
			Address:   "临川市",
			StartYear: 1987,
			EndYear:   1999,
		},
	},
	362521: {
		{
			Address: "吉安县",
			EndYear: 1981,
		},
		{
			Address:   "临川县",
			StartYear: 1982,
			EndYear:   1986,
		},
	},
	362522: {
		{
			Address: "万安县",
			EndYear: 1981,
		},
		{
			Address:   "南城县",
			StartYear: 1982,
			EndYear:   1999,
		},
	},
	362523: {
		{
			Address: "新干县",
			EndYear: 1981,
		},
		{
			Address:   "黎川县",
			StartYear: 1982,
			EndYear:   1999,
		},
	},
	362524: {
		{
			Address: "遂川县",
			EndYear: 1981,
		},
		{
			Address:   "南丰县",
			StartYear: 1982,
			EndYear:   1999,
		},
	},
	362525: {
		{
			Address: "峡江县",
			EndYear: 1981,
		},
		{
			Address:   "崇仁县",
			StartYear: 1982,
			EndYear:   1999,
		},
	},
	362526: {
		{
			Address: "宁冈县",
			EndYear: 1981,
		},
		{
			Address:   "乐安县",
			StartYear: 1982,
			EndYear:   1999,
		},
	},
	362527: {
		{
			Address: "吉水县",
			EndYear: 1981,
		},
		{
			Address:   "宜黄县",
			StartYear: 1982,
			EndYear:   1999,
		},
	},
	362528: {
		{
			Address: "永新县",
			EndYear: 1981,
		},
		{
			Address:   "金溪县",
			StartYear: 1982,
			EndYear:   1999,
		},
	},
	362529: {
		{
			Address: "永丰县",
			EndYear: 1981,
		},
		{
			Address:   "资溪县",
			StartYear: 1982,
			EndYear:   1999,
		},
	},
	362530: {
		{
			Address: "莲花县",
			EndYear: 1981,
		},
		{
			Address:   "进贤县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	362531: {
		{
			Address: "泰和县",
			EndYear: 1981,
		},
		{
			Address:   "东乡县",
			StartYear: 1982,
			EndYear:   1999,
		},
	},
	362532: {
		{
			Address: "安福县",
			EndYear: 1981,
		},
		{
			Address:   "广昌县",
			StartYear: 1983,
			EndYear:   1999,
		},
	},
	362533: {
		{
			Address:   "井冈山县",
			StartYear: 1981,
			EndYear:   1981,
		},
	},
	362600: {
		{
			Address: "赣州地区",
			EndYear: 1981,
		},
		{
			Address:   "九江地区",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	362601: {
		{
			Address: "赣州市",
			EndYear: 1981,
		},
	},
	362621: {
		{
			Address: "广昌县",
			EndYear: 1981,
		},
		{
			Address:   "九江县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	362622: {
		{
			Address: "定南县",
			EndYear: 1981,
		},
		{
			Address:   "瑞昌县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	362623: {
		{
			Address: "石城县",
			EndYear: 1981,
		},
		{
			Address:   "武宁县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	362624: {
		{
			Address: "龙南县",
			EndYear: 1981,
		},
		{
			Address:   "修水县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	362625: {
		{
			Address: "宁都县",
			EndYear: 1981,
		},
		{
			Address:   "永修县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	362626: {
		{
			Address: "全南县",
			EndYear: 1981,
		},
		{
			Address:   "德安县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	362627: {
		{
			Address: "兴国县",
			EndYear: 1981,
		},
		{
			Address:   "星子县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	362628: {
		{
			Address: "信丰县",
			EndYear: 1981,
		},
		{
			Address:   "都昌县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	362629: {
		{
			Address: "于都县",
			EndYear: 1981,
		},
		{
			Address:   "湖口县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	362630: {
		{
			Address: "赣县",
			EndYear: 1981,
		},
		{
			Address:   "彭泽县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	362631: {
		{
			Address: "瑞金县",
			EndYear: 1981,
		},
	},
	362632: {
		{
			Address: "南康县",
			EndYear: 1981,
		},
	},
	362633: {
		{
			Address: "会昌县",
			EndYear: 1981,
		},
	},
	362634: {
		{
			Address: "上犹县",
			EndYear: 1981,
		},
	},
	362635: {
		{
			Address: "安远县",
			EndYear: 1981,
		},
	},
	362636: {
		{
			Address: "崇义县",
			EndYear: 1981,
		},
	},
	362637: {
		{
			Address: "寻乌县",
			EndYear: 1981,
		},
	},
	362638: {
		{
			Address: "大余县",
			EndYear: 1981,
		},
	},
	369001: {
		{
			Address:   "瑞昌市",
			StartYear: 1989,
			EndYear:   1994,
		},
	},
	369002: {
		{
			Address:   "乐平市",
			StartYear: 1992,
			EndYear:   1994,
		},
	},
	370000: {
		{
			Address: "山东省",
		},
	},
	370100: {
		{
			Address: "济南市",
		},
	},
	370102: {
		{
			Address: "历下区",
		},
	},
	370103: {
		{
			Address: "市中区",
		},
	},
	370104: {
		{
			Address: "槐荫区",
		},
	},
	370105: {
		{
			Address: "天桥区",
		},
	},
	370111: {
		{
			Address: "郊区",
			EndYear: 1986,
		},
	},
	370112: {
		{
			Address:   "历城区",
			StartYear: 1987,
		},
	},
	370113: {
		{
			Address:   "长清区",
			StartYear: 2001,
		},
	},
	370114: {
		{
			Address:   "章丘区",
			StartYear: 2016,
		},
	},
	370115: {
		{
			Address:   "济阳区",
			StartYear: 2018,
		},
	},
	370116: {
		{
			Address:   "莱芜区",
			StartYear: 2018,
		},
	},
	370117: {
		{
			Address:   "钢城区",
			StartYear: 2018,
		},
	},
	370121: {
		{
			Address: "历城县",
			EndYear: 1986,
		},
	},
	370122: {
		{
			Address: "章丘县",
			EndYear: 1991,
		},
	},
	370123: {
		{
			Address: "长清县",
			EndYear: 2000,
		},
	},
	370124: {
		{
			Address:   "平阴县",
			StartYear: 1985,
		},
	},
	370125: {
		{
			Address:   "济阳县",
			StartYear: 1989,
			EndYear:   2017,
		},
	},
	370126: {
		{
			Address:   "商河县",
			StartYear: 1989,
		},
	},
	370181: {
		{
			Address:   "章丘市",
			StartYear: 1995,
			EndYear:   2015,
		},
	},
	370200: {
		{
			Address: "青岛市",
		},
	},
	370202: {
		{
			Address: "市南区",
		},
	},
	370203: {
		{
			Address: "市北区",
		},
	},
	370204: {
		{
			Address: "台东区",
			EndYear: 1993,
		},
	},
	370205: {
		{
			Address: "四方区",
			EndYear: 2011,
		},
	},
	370206: {
		{
			Address: "沧口区",
			EndYear: 1993,
		},
	},
	370211: {
		{
			Address: "黄岛区",
		},
	},
	370212: {
		{
			Address:   "崂山区",
			StartYear: 1988,
		},
	},
	370213: {
		{
			Address:   "李沧区",
			StartYear: 1994,
		},
	},
	370214: {
		{
			Address:   "城阳区",
			StartYear: 1994,
		},
	},
	370215: {
		{
			Address:   "即墨区",
			StartYear: 2017,
		},
	},
	370221: {
		{
			Address: "崂山县",
			EndYear: 1987,
		},
	},
	370222: {
		{
			Address: "即墨县",
			EndYear: 1988,
		},
	},
	370223: {
		{
			Address: "胶南县",
			EndYear: 1989,
		},
	},
	370224: {
		{
			Address: "胶县",
			EndYear: 1986,
		},
	},
	370225: {
		{
			Address:   "莱西县",
			StartYear: 1983,
			EndYear:   1989,
		},
	},
	370226: {
		{
			Address:   "平度县",
			StartYear: 1983,
			EndYear:   1988,
		},
	},
	370281: {
		{
			Address:   "胶州市",
			StartYear: 1995,
		},
	},
	370282: {
		{
			Address:   "即墨市",
			StartYear: 1995,
			EndYear:   2016,
		},
	},
	370283: {
		{
			Address:   "平度市",
			StartYear: 1995,
		},
	},
	370284: {
		{
			Address:   "胶南市",
			StartYear: 1995,
			EndYear:   2011,
		},
	},
	370285: {
		{
			Address:   "莱西市",
			StartYear: 1995,
		},
	},
	370300: {
		{
			Address: "淄博市",
		},
	},
	370302: {
		{
			Address: "淄川区",
		},
	},
	370303: {
		{
			Address: "张店区",
		},
	},
	370304: {
		{
			Address: "博山区",
		},
	},
	370305: {
		{
			Address: "临淄区",
		},
	},
	370306: {
		{
			Address: "周村区",
		},
	},
	370321: {
		{
			Address:   "桓台县",
			StartYear: 1983,
		},
	},
	370322: {
		{
			Address:   "高青县",
			StartYear: 1989,
		},
	},
	370323: {
		{
			Address:   "沂源县",
			StartYear: 1989,
		},
	},
	370400: {
		{
			Address: "枣庄市",
		},
	},
	370402: {
		{
			Address: "市中区",
		},
	},
	370403: {
		{
			Address: "薛城区",
		},
	},
	370404: {
		{
			Address: "峄城区",
		},
	},
	370405: {
		{
			Address: "台儿庄区",
		},
	},
	370406: {
		{
			Address: "齐村区",
			EndYear: 1982,
		},
		{
			Address:   "山亭区",
			StartYear: 1983,
		},
	},
	370421: {
		{
			Address: "滕县",
			EndYear: 1987,
		},
	},
	370481: {
		{
			Address:   "滕州市",
			StartYear: 1995,
		},
	},
	370500: {
		{
			Address:   "东营市",
			StartYear: 1982,
		},
	},
	370502: {
		{
			Address:   "东营区",
			StartYear: 1982,
		},
	},
	370503: {
		{
			Address:   "河口区",
			StartYear: 1982,
		},
	},
	370504: {
		{
			Address:   "牛庄区",
			StartYear: 1982,
			EndYear:   1986,
		},
	},
	370505: {
		{
			Address:   "垦利区",
			StartYear: 2016,
		},
	},
	370521: {
		{
			Address:   "垦利县",
			StartYear: 1982,
			EndYear:   2015,
		},
	},
	370522: {
		{
			Address:   "利津县",
			StartYear: 1982,
		},
	},
	370523: {
		{
			Address:   "广饶县",
			StartYear: 1983,
		},
	},
	370600: {
		{
			Address:   "烟台市",
			StartYear: 1983,
		},
	},
	370602: {
		{
			Address:   "芝罘区",
			StartYear: 1983,
		},
	},
	370611: {
		{
			Address:   "福山区",
			StartYear: 1983,
		},
	},
	370612: {
		{
			Address:   "牟平区",
			StartYear: 1994,
		},
	},
	370613: {
		{
			Address:   "莱山区",
			StartYear: 1994,
		},
	},
	370614: {
		{
			Address:   "蓬莱区",
			StartYear: 2020,
		},
	},
	370620: {
		{
			Address:   "威海市",
			StartYear: 1985,
			EndYear:   1985,
		},
	},
	370622: {
		{
			Address:   "蓬莱县",
			StartYear: 1983,
			EndYear:   1990,
		},
	},
	370623: {
		{
			Address:   "黄县",
			StartYear: 1983,
			EndYear:   1985,
		},
	},
	370624: {
		{
			Address:   "招远县",
			StartYear: 1983,
			EndYear:   1990,
		},
	},
	370625: {
		{
			Address:   "掖县",
			StartYear: 1983,
			EndYear:   1987,
		},
	},
	370627: {
		{
			Address:   "莱阳县",
			StartYear: 1983,
			EndYear:   1986,
		},
	},
	370628: {
		{
			Address:   "栖霞县",
			StartYear: 1983,
			EndYear:   1994,
		},
	},
	370629: {
		{
			Address:   "海阳县",
			StartYear: 1983,
			EndYear:   1995,
		},
	},
	370630: {
		{
			Address:   "乳山县",
			StartYear: 1983,
			EndYear:   1986,
		},
	},
	370631: {
		{
			Address:   "牟平县",
			StartYear: 1983,
			EndYear:   1993,
		},
	},
	370632: {
		{
			Address:   "文登县",
			StartYear: 1983,
			EndYear:   1986,
		},
	},
	370633: {
		{
			Address:   "荣成县",
			StartYear: 1983,
			EndYear:   1986,
		},
	},
	370634: {
		{
			Address:   "长岛县",
			StartYear: 1983,
			EndYear:   2019,
		},
	},
	370681: {
		{
			Address:   "龙口市",
			StartYear: 1995,
		},
	},
	370682: {
		{
			Address:   "莱阳市",
			StartYear: 1995,
		},
	},
	370683: {
		{
			Address:   "莱州市",
			StartYear: 1995,
		},
	},
	370684: {
		{
			Address:   "蓬莱市",
			StartYear: 1995,
			EndYear:   2019,
		},
	},
	370685: {
		{
			Address:   "招远市",
			StartYear: 1995,
		},
	},
	370686: {
		{
			Address:   "栖霞市",
			StartYear: 1995,
		},
	},
	370687: {
		{
			Address:   "海阳市",
			StartYear: 1996,
		},
	},
	370700: {
		{
			Address:   "潍坊市",
			StartYear: 1983,
		},
	},
	370702: {
		{
			Address:   "潍城区",
			StartYear: 1983,
		},
	},
	370703: {
		{
			Address:   "寒亭区",
			StartYear: 1983,
		},
	},
	370704: {
		{
			Address:   "坊子区",
			StartYear: 1983,
		},
	},
	370705: {
		{
			Address:   "奎文区",
			StartYear: 1994,
		},
	},
	370721: {
		{
			Address:   "益都县",
			StartYear: 1983,
			EndYear:   1985,
		},
	},
	370722: {
		{
			Address:   "安丘县",
			StartYear: 1983,
			EndYear:   1993,
		},
	},
	370723: {
		{
			Address:   "寿光县",
			StartYear: 1983,
			EndYear:   1992,
		},
	},
	370724: {
		{
			Address:   "临朐县",
			StartYear: 1983,
		},
	},
	370725: {
		{
			Address:   "昌乐县",
			StartYear: 1983,
		},
	},
	370726: {
		{
			Address:   "昌邑县",
			StartYear: 1983,
			EndYear:   1993,
		},
	},
	370727: {
		{
			Address:   "高密县",
			StartYear: 1983,
			EndYear:   1993,
		},
	},
	370728: {
		{
			Address:   "诸城县",
			StartYear: 1983,
			EndYear:   1986,
		},
	},
	370729: {
		{
			Address:   "五莲县",
			StartYear: 1983,
			EndYear:   1991,
		},
	},
	370781: {
		{
			Address:   "青州市",
			StartYear: 1995,
		},
	},
	370782: {
		{
			Address:   "诸城市",
			StartYear: 1995,
		},
	},
	370783: {
		{
			Address:   "寿光市",
			StartYear: 1995,
		},
	},
	370784: {
		{
			Address:   "安丘市",
			StartYear: 1995,
		},
	},
	370785: {
		{
			Address:   "高密市",
			StartYear: 1995,
		},
	},
	370786: {
		{
			Address:   "昌邑市",
			StartYear: 1995,
		},
	},
	370800: {
		{
			Address:   "济宁市",
			StartYear: 1983,
		},
	},
	370802: {
		{
			Address:   "市中区",
			StartYear: 1983,
			EndYear:   2012,
		},
	},
	370811: {
		{
			Address:   "市郊区",
			StartYear: 1983,
			EndYear:   1992,
		},
		{
			Address:   "任城区",
			StartYear: 1993,
		},
	},
	370812: {
		{
			Address:   "兖州区",
			StartYear: 2013,
		},
	},
	370822: {
		{
			Address:   "兖州县",
			StartYear: 1983,
			EndYear:   1991,
		},
	},
	370823: {
		{
			Address:   "曲阜县",
			StartYear: 1983,
			EndYear:   1985,
		},
	},
	370825: {
		{
			Address:   "邹县",
			StartYear: 1983,
			EndYear:   1991,
		},
	},
	370826: {
		{
			Address:   "微山县",
			StartYear: 1983,
		},
	},
	370827: {
		{
			Address:   "鱼台县",
			StartYear: 1983,
		},
	},
	370828: {
		{
			Address:   "金乡县",
			StartYear: 1983,
		},
	},
	370829: {
		{
			Address:   "嘉祥县",
			StartYear: 1983,
		},
	},
	370830: {
		{
			Address:   "汶上县",
			StartYear: 1985,
		},
	},
	370831: {
		{
			Address:   "泗水县",
			StartYear: 1985,
		},
	},
	370832: {
		{
			Address:   "梁山县",
			StartYear: 1989,
		},
	},
	370881: {
		{
			Address:   "曲阜市",
			StartYear: 1995,
		},
	},
	370882: {
		{
			Address:   "兖州市",
			StartYear: 1995,
			EndYear:   2012,
		},
	},
	370883: {
		{
			Address:   "邹城市",
			StartYear: 1995,
		},
	},
	370900: {
		{
			Address:   "泰安市",
			StartYear: 1985,
		},
	},
	370901: {
		{
			Address:   "威海市",
			StartYear: 1983,
			EndYear:   1984,
		},
	},
	370902: {
		{
			Address:   "泰山区",
			StartYear: 1985,
		},
	},
	370911: {
		{
			Address:   "郊区",
			StartYear: 1985,
			EndYear:   1999,
		},
		{
			Address:   "岱岳区",
			StartYear: 2000,
		},
	},
	370921: {
		{
			Address:   "宁阳县",
			StartYear: 1985,
		},
	},
	370922: {
		{
			Address:   "肥城县",
			StartYear: 1985,
			EndYear:   1991,
		},
	},
	370923: {
		{
			Address:   "东平县",
			StartYear: 1985,
		},
	},
	370981: {
		{
			Address:   "莱芜市",
			StartYear: 1985,
			EndYear:   1985,
		},
	},
	370982: {
		{
			Address:   "新泰市",
			StartYear: 1985,
		},
	},
	370983: {
		{
			Address:   "肥城市",
			StartYear: 1995,
		},
	},
	371000: {
		{
			Address:   "威海市",
			StartYear: 1987,
		},
	},
	371002: {
		{
			Address:   "环翠区",
			StartYear: 1987,
		},
	},
	371003: {
		{
			Address:   "文登区",
			StartYear: 2014,
		},
	},
	371021: {
		{
			Address:   "乳山县",
			StartYear: 1987,
			EndYear:   1992,
		},
	},
	371022: {
		{
			Address:   "文登县",
			StartYear: 1987,
			EndYear:   1987,
		},
	},
	371023: {
		{
			Address:   "荣成县",
			StartYear: 1987,
			EndYear:   1987,
		},
	},
	371081: {
		{
			Address:   "文登市",
			StartYear: 1995,
			EndYear:   2013,
		},
	},
	371082: {
		{
			Address:   "荣成市",
			StartYear: 1995,
		},
	},
	371083: {
		{
			Address:   "乳山市",
			StartYear: 1995,
		},
	},
	371100: {
		{
			Address:   "日照市",
			StartYear: 1989,
		},
	},
	371102: {
		{
			Address:   "东港区",
			StartYear: 1992,
		},
	},
	371103: {
		{
			Address:   "岚山区",
			StartYear: 2004,
		},
	},
	371121: {
		{
			Address:   "五莲县",
			StartYear: 1992,
		},
	},
	371122: {
		{
			Address:   "莒县",
			StartYear: 1992,
		},
	},
	371200: {
		{
			Address:   "莱芜市",
			StartYear: 1992,
			EndYear:   2017,
		},
	},
	371202: {
		{
			Address:   "莱城区",
			StartYear: 1992,
			EndYear:   2017,
		},
	},
	371203: {
		{
			Address:   "钢城区",
			StartYear: 1992,
			EndYear:   2017,
		},
	},
	371300: {
		{
			Address:   "临沂市",
			StartYear: 1994,
		},
	},
	371302: {
		{
			Address:   "兰山区",
			StartYear: 1994,
		},
	},
	371311: {
		{
			Address:   "罗庄区",
			StartYear: 1994,
		},
	},
	371312: {
		{
			Address:   "河东区",
			StartYear: 1994,
		},
	},
	371321: {
		{
			Address:   "沂南县",
			StartYear: 1994,
		},
	},
	371322: {
		{
			Address:   "郯城县",
			StartYear: 1994,
		},
	},
	371323: {
		{
			Address:   "沂水县",
			StartYear: 1994,
		},
	},
	371324: {
		{
			Address:   "苍山县",
			StartYear: 1994,
			EndYear:   2013,
		},
		{
			Address:   "兰陵县",
			StartYear: 2014,
		},
	},
	371325: {
		{
			Address:   "费县",
			StartYear: 1994,
		},
	},
	371326: {
		{
			Address:   "平邑县",
			StartYear: 1994,
		},
	},
	371327: {
		{
			Address:   "莒南县",
			StartYear: 1994,
		},
	},
	371328: {
		{
			Address:   "蒙阴县",
			StartYear: 1994,
		},
	},
	371329: {
		{
			Address:   "临沭县",
			StartYear: 1994,
		},
	},
	371400: {
		{
			Address:   "德州市",
			StartYear: 1994,
		},
	},
	371402: {
		{
			Address:   "德城区",
			StartYear: 1994,
		},
	},
	371403: {
		{
			Address:   "陵城区",
			StartYear: 2014,
		},
	},
	371421: {
		{
			Address:   "陵县",
			StartYear: 1994,
			EndYear:   2013,
		},
	},
	371422: {
		{
			Address:   "宁津县",
			StartYear: 1994,
		},
	},
	371423: {
		{
			Address:   "庆云县",
			StartYear: 1994,
		},
	},
	371424: {
		{
			Address:   "临邑县",
			StartYear: 1994,
		},
	},
	371425: {
		{
			Address:   "齐河县",
			StartYear: 1994,
		},
	},
	371426: {
		{
			Address:   "平原县",
			StartYear: 1994,
		},
	},
	371427: {
		{
			Address:   "夏津县",
			StartYear: 1994,
		},
	},
	371428: {
		{
			Address:   "武城县",
			StartYear: 1994,
		},
	},
	371481: {
		{
			Address:   "乐陵市",
			StartYear: 1995,
		},
	},
	371482: {
		{
			Address:   "禹城市",
			StartYear: 1995,
		},
	},
	371500: {
		{
			Address:   "聊城市",
			StartYear: 1997,
		},
	},
	371502: {
		{
			Address:   "东昌府区",
			StartYear: 1997,
		},
	},
	371503: {
		{
			Address:   "茌平区",
			StartYear: 2019,
		},
	},
	371521: {
		{
			Address:   "阳谷县",
			StartYear: 1997,
		},
	},
	371522: {
		{
			Address:   "莘县",
			StartYear: 1997,
		},
	},
	371523: {
		{
			Address:   "茌平县",
			StartYear: 1997,
			EndYear:   2018,
		},
	},
	371524: {
		{
			Address:   "东阿县",
			StartYear: 1997,
		},
	},
	371525: {
		{
			Address:   "冠县",
			StartYear: 1997,
		},
	},
	371526: {
		{
			Address:   "高唐县",
			StartYear: 1997,
		},
	},
	371581: {
		{
			Address:   "临清市",
			StartYear: 1997,
		},
	},
	371600: {
		{
			Address:   "滨州市",
			StartYear: 2000,
		},
	},
	371602: {
		{
			Address:   "滨城区",
			StartYear: 2000,
		},
	},
	371603: {
		{
			Address:   "沾化区",
			StartYear: 2014,
		},
	},
	371621: {
		{
			Address:   "惠民县",
			StartYear: 2000,
		},
	},
	371622: {
		{
			Address:   "阳信县",
			StartYear: 2000,
		},
	},
	371623: {
		{
			Address:   "无棣县",
			StartYear: 2000,
		},
	},
	371624: {
		{
			Address:   "沾化县",
			StartYear: 2000,
			EndYear:   2013,
		},
	},
	371625: {
		{
			Address:   "博兴县",
			StartYear: 2000,
		},
	},
	371626: {
		{
			Address:   "邹平县",
			StartYear: 2000,
			EndYear:   2017,
		},
	},
	371681: {
		{
			Address:   "邹平市",
			StartYear: 2018,
		},
	},
	371700: {
		{
			Address:   "菏泽市",
			StartYear: 2000,
		},
	},
	371702: {
		{
			Address:   "牡丹区",
			StartYear: 2000,
		},
	},
	371703: {
		{
			Address:   "定陶区",
			StartYear: 2016,
		},
	},
	371721: {
		{
			Address:   "曹县",
			StartYear: 2000,
		},
	},
	371722: {
		{
			Address:   "单县",
			StartYear: 2000,
		},
	},
	371723: {
		{
			Address:   "成武县",
			StartYear: 2000,
		},
	},
	371724: {
		{
			Address:   "巨野县",
			StartYear: 2000,
		},
	},
	371725: {
		{
			Address:   "郓城县",
			StartYear: 2000,
		},
	},
	371726: {
		{
			Address:   "鄄城县",
			StartYear: 2000,
		},
	},
	371727: {
		{
			Address:   "定陶县",
			StartYear: 2000,
			EndYear:   2015,
		},
	},
	371728: {
		{
			Address:   "东明县",
			StartYear: 2000,
		},
	},
	372100: {
		{
			Address: "烟台地区",
			EndYear: 1982,
		},
	},
	372101: {
		{
			Address: "烟台市",
			EndYear: 1982,
		},
	},
	372102: {
		{
			Address: "威海市",
			EndYear: 1982,
		},
	},
	372121: {
		{
			Address: "福山县",
			EndYear: 1982,
		},
	},
	372122: {
		{
			Address: "蓬莱县",
			EndYear: 1982,
		},
	},
	372123: {
		{
			Address: "黄县",
			EndYear: 1982,
		},
	},
	372124: {
		{
			Address: "招远县",
			EndYear: 1982,
		},
	},
	372125: {
		{
			Address: "掖县",
			EndYear: 1982,
		},
	},
	372126: {
		{
			Address: "莱西县",
			EndYear: 1982,
		},
	},
	372127: {
		{
			Address: "莱阳县",
			EndYear: 1982,
		},
	},
	372128: {
		{
			Address: "栖霞县",
			EndYear: 1982,
		},
	},
	372129: {
		{
			Address: "海阳县",
			EndYear: 1982,
		},
	},
	372130: {
		{
			Address: "乳山县",
			EndYear: 1982,
		},
	},
	372131: {
		{
			Address: "牟平县",
			EndYear: 1982,
		},
	},
	372132: {
		{
			Address: "文登县",
			EndYear: 1982,
		},
	},
	372133: {
		{
			Address: "荣成县",
			EndYear: 1982,
		},
	},
	372134: {
		{
			Address: "长岛县",
			EndYear: 1982,
		},
	},
	372200: {
		{
			Address: "昌潍地区",
			EndYear: 1980,
		},
		{
			Address:   "潍坊地区",
			StartYear: 1981,
			EndYear:   1982,
		},
	},
	372201: {
		{
			Address: "潍坊市",
			EndYear: 1982,
		},
	},
	372221: {
		{
			Address: "益都县",
			EndYear: 1982,
		},
	},
	372222: {
		{
			Address: "安丘县",
			EndYear: 1982,
		},
	},
	372223: {
		{
			Address: "寿光县",
			EndYear: 1982,
		},
	},
	372224: {
		{
			Address: "临朐县",
			EndYear: 1982,
		},
	},
	372225: {
		{
			Address: "昌乐县",
			EndYear: 1982,
		},
	},
	372226: {
		{
			Address: "昌邑县",
			EndYear: 1982,
		},
	},
	372227: {
		{
			Address: "高密县",
			EndYear: 1982,
		},
	},
	372228: {
		{
			Address: "诸城县",
			EndYear: 1982,
		},
	},
	372229: {
		{
			Address: "五莲县",
			EndYear: 1982,
		},
	},
	372230: {
		{
			Address: "平度县",
			EndYear: 1982,
		},
	},
	372231: {
		{
			Address: "潍县",
			EndYear: 1982,
		},
	},
	372300: {
		{
			Address: "惠民地区",
			EndYear: 1991,
		},
		{
			Address:   "滨州地区",
			StartYear: 1992,
			EndYear:   1999,
		},
	},
	372301: {
		{
			Address:   "滨州市",
			StartYear: 1982,
			EndYear:   1999,
		},
	},
	372321: {
		{
			Address: "惠民县",
			EndYear: 1999,
		},
	},
	372322: {
		{
			Address: "滨县",
			EndYear: 1986,
		},
	},
	372323: {
		{
			Address: "阳信县",
			EndYear: 1999,
		},
	},
	372324: {
		{
			Address: "无棣县",
			EndYear: 1999,
		},
	},
	372325: {
		{
			Address: "沾化县",
			EndYear: 1999,
		},
	},
	372326: {
		{
			Address: "利津县",
			EndYear: 1981,
		},
	},
	372327: {
		{
			Address: "广饶县",
			EndYear: 1982,
		},
	},
	372328: {
		{
			Address: "博兴县",
			EndYear: 1999,
		},
	},
	372329: {
		{
			Address: "桓台县",
			EndYear: 1982,
		},
	},
	372330: {
		{
			Address: "邹平县",
			EndYear: 1999,
		},
	},
	372331: {
		{
			Address: "高青县",
			EndYear: 1988,
		},
	},
	372332: {
		{
			Address: "垦利县",
			EndYear: 1981,
		},
	},
	372400: {
		{
			Address: "德州地区",
			EndYear: 1993,
		},
	},
	372401: {
		{
			Address: "德州市",
			EndYear: 1993,
		},
	},
	372402: {
		{
			Address:   "乐陵市",
			StartYear: 1988,
			EndYear:   1993,
		},
	},
	372403: {
		{
			Address:   "禹城市",
			StartYear: 1993,
			EndYear:   1993,
		},
	},
	372421: {
		{
			Address: "陵县",
			EndYear: 1993,
		},
	},
	372422: {
		{
			Address: "平原县",
			EndYear: 1993,
		},
	},
	372423: {
		{
			Address: "夏津县",
			EndYear: 1993,
		},
	},
	372424: {
		{
			Address: "武城县",
			EndYear: 1993,
		},
	},
	372425: {
		{
			Address: "齐河县",
			EndYear: 1993,
		},
	},
	372426: {
		{
			Address: "禹城县",
			EndYear: 1992,
		},
	},
	372427: {
		{
			Address: "乐陵县",
			EndYear: 1987,
		},
	},
	372428: {
		{
			Address: "临邑县",
			EndYear: 1993,
		},
	},
	372429: {
		{
			Address: "商河县",
			EndYear: 1988,
		},
	},
	372430: {
		{
			Address: "济阳县",
			EndYear: 1988,
		},
	},
	372431: {
		{
			Address: "宁津县",
			EndYear: 1993,
		},
	},
	372432: {
		{
			Address: "庆云县",
			EndYear: 1993,
		},
	},
	372500: {
		{
			Address: "聊城地区",
			EndYear: 1996,
		},
	},
	372501: {
		{
			Address:   "聊城市",
			StartYear: 1983,
			EndYear:   1996,
		},
	},
	372502: {
		{
			Address:   "临清市",
			StartYear: 1983,
			EndYear:   1996,
		},
	},
	372521: {
		{
			Address: "阳谷县",
			EndYear: 1981,
		},
		{
			Address:   "聊城县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	372522: {
		{
			Address: "莘县",
			EndYear: 1981,
		},
		{
			Address:   "阳谷县",
			StartYear: 1982,
			EndYear:   1996,
		},
	},
	372523: {
		{
			Address: "茌平县",
			EndYear: 1981,
		},
		{
			Address:   "莘县",
			StartYear: 1982,
			EndYear:   1996,
		},
	},
	372524: {
		{
			Address: "东阿县",
			EndYear: 1981,
		},
		{
			Address:   "茌平县",
			StartYear: 1982,
			EndYear:   1996,
		},
	},
	372525: {
		{
			Address: "冠县",
			EndYear: 1981,
		},
		{
			Address:   "东阿县",
			StartYear: 1982,
			EndYear:   1996,
		},
	},
	372526: {
		{
			Address: "高唐县",
			EndYear: 1981,
		},
		{
			Address:   "冠县",
			StartYear: 1982,
			EndYear:   1996,
		},
	},
	372527: {
		{
			Address: "临清县",
			EndYear: 1981,
		},
		{
			Address:   "高唐县",
			StartYear: 1982,
			EndYear:   1996,
		},
	},
	372528: {
		{
			Address: "聊城县",
			EndYear: 1981,
		},
		{
			Address:   "临清县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	372600: {
		{
			Address: "泰安地区",
			EndYear: 1984,
		},
	},
	372601: {
		{
			Address:   "泰安市",
			StartYear: 1982,
			EndYear:   1984,
		},
	},
	372602: {
		{
			Address:   "新汶市",
			StartYear: 1982,
			EndYear:   1982,
		},
		{
			Address:   "莱芜市",
			StartYear: 1983,
			EndYear:   1984,
		},
	},
	372603: {
		{
			Address:   "新泰市",
			StartYear: 1983,
			EndYear:   1984,
		},
	},
	372621: {
		{
			Address: "泰安县",
			EndYear: 1981,
		},
		{
			Address:   "莱芜县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	372622: {
		{
			Address: "莱芜县",
			EndYear: 1981,
		},
		{
			Address:   "新泰县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	372623: {
		{
			Address: "新泰县",
			EndYear: 1981,
		},
	},
	372624: {
		{
			Address: "宁阳县",
			EndYear: 1984,
		},
	},
	372625: {
		{
			Address: "肥城县",
			EndYear: 1984,
		},
	},
	372626: {
		{
			Address: "东平县",
			EndYear: 1984,
		},
	},
	372627: {
		{
			Address: "平阴县",
			EndYear: 1984,
		},
	},
	372628: {
		{
			Address: "新汶县",
			EndYear: 1981,
		},
	},
	372629: {
		{
			Address:   "汶上县",
			StartYear: 1983,
			EndYear:   1984,
		},
	},
	372630: {
		{
			Address:   "泗水县",
			StartYear: 1983,
			EndYear:   1984,
		},
	},
	372700: {
		{
			Address: "济宁地区",
			EndYear: 1982,
		},
	},
	372701: {
		{
			Address: "济宁市",
			EndYear: 1982,
		},
	},
	372721: {
		{
			Address: "济宁县",
			EndYear: 1982,
		},
	},
	372722: {
		{
			Address: "兖州县",
			EndYear: 1982,
		},
	},
	372723: {
		{
			Address: "曲阜县",
			EndYear: 1982,
		},
	},
	372724: {
		{
			Address: "泗水县",
			EndYear: 1982,
		},
	},
	372725: {
		{
			Address: "邹县",
			EndYear: 1982,
		},
	},
	372726: {
		{
			Address: "微山县",
			EndYear: 1982,
		},
	},
	372727: {
		{
			Address: "鱼台县",
			EndYear: 1982,
		},
	},
	372728: {
		{
			Address: "金乡县",
			EndYear: 1982,
		},
	},
	372729: {
		{
			Address: "嘉祥县",
			EndYear: 1982,
		},
	},
	372730: {
		{
			Address: "汶上县",
			EndYear: 1982,
		},
	},
	372800: {
		{
			Address: "临沂地区",
			EndYear: 1993,
		},
	},
	372801: {
		{
			Address:   "临沂市",
			StartYear: 1983,
			EndYear:   1993,
		},
	},
	372821: {
		{
			Address: "临沂县",
			EndYear: 1982,
		},
	},
	372822: {
		{
			Address: "郯城县",
			EndYear: 1993,
		},
	},
	372823: {
		{
			Address: "苍山县",
			EndYear: 1993,
		},
	},
	372824: {
		{
			Address: "莒南县",
			EndYear: 1993,
		},
	},
	372825: {
		{
			Address: "日照县",
			EndYear: 1988,
		},
	},
	372826: {
		{
			Address: "莒县",
			EndYear: 1991,
		},
	},
	372827: {
		{
			Address: "沂水县",
			EndYear: 1993,
		},
	},
	372828: {
		{
			Address: "沂源县",
			EndYear: 1988,
		},
	},
	372829: {
		{
			Address: "蒙阴县",
			EndYear: 1993,
		},
	},
	372830: {
		{
			Address: "平邑县",
			EndYear: 1993,
		},
	},
	372831: {
		{
			Address: "费县",
			EndYear: 1993,
		},
	},
	372832: {
		{
			Address: "沂南县",
			EndYear: 1993,
		},
	},
	372833: {
		{
			Address: "临沭县",
			EndYear: 1993,
		},
	},
	372900: {
		{
			Address: "菏泽地区",
			EndYear: 1999,
		},
	},
	372901: {
		{
			Address:   "菏泽市",
			StartYear: 1983,
			EndYear:   1999,
		},
	},
	372921: {
		{
			Address: "菏泽县",
			EndYear: 1982,
		},
	},
	372922: {
		{
			Address: "曹县",
			EndYear: 1999,
		},
	},
	372923: {
		{
			Address: "定陶县",
			EndYear: 1999,
		},
	},
	372924: {
		{
			Address: "成武县",
			EndYear: 1999,
		},
	},
	372925: {
		{
			Address: "单县",
			EndYear: 1999,
		},
	},
	372926: {
		{
			Address: "巨野县",
			EndYear: 1999,
		},
	},
	372927: {
		{
			Address: "梁山县",
			EndYear: 1988,
		},
	},
	372928: {
		{
			Address: "郓城县",
			EndYear: 1999,
		},
	},
	372929: {
		{
			Address: "鄄城县",
			EndYear: 1999,
		},
	},
	372930: {
		{
			Address: "东明县",
			EndYear: 1999,
		},
	},
	379001: {
		{
			Address:   "青州市",
			StartYear: 1986,
			EndYear:   1994,
		},
	},
	379002: {
		{
			Address:   "龙口市",
			StartYear: 1986,
			EndYear:   1994,
		},
	},
	379003: {
		{
			Address:   "曲阜市",
			StartYear: 1986,
			EndYear:   1994,
		},
	},
	379004: {
		{
			Address:   "莱芜市",
			StartYear: 1986,
			EndYear:   1991,
		},
	},
	379005: {
		{
			Address:   "新泰市",
			StartYear: 1986,
			EndYear:   1994,
		},
	},
	379006: {
		{
			Address:   "威海市",
			StartYear: 1986,
			EndYear:   1986,
		},
		{
			Address:   "胶州市",
			StartYear: 1987,
			EndYear:   1994,
		},
	},
	379007: {
		{
			Address:   "诸城市",
			StartYear: 1987,
			EndYear:   1994,
		},
	},
	379008: {
		{
			Address:   "莱阳市",
			StartYear: 1987,
			EndYear:   1994,
		},
	},
	379009: {
		{
			Address:   "莱州市",
			StartYear: 1988,
			EndYear:   1994,
		},
	},
	379010: {
		{
			Address:   "滕州市",
			StartYear: 1988,
			EndYear:   1994,
		},
	},
	379011: {
		{
			Address:   "文登市",
			StartYear: 1988,
			EndYear:   1994,
		},
	},
	379012: {
		{
			Address:   "荣成市",
			StartYear: 1988,
			EndYear:   1994,
		},
	},
	379013: {
		{
			Address:   "即墨市",
			StartYear: 1989,
			EndYear:   1994,
		},
	},
	379014: {
		{
			Address:   "平度市",
			StartYear: 1989,
			EndYear:   1994,
		},
	},
	379015: {
		{
			Address:   "莱西市",
			StartYear: 1990,
			EndYear:   1994,
		},
	},
	379016: {
		{
			Address:   "胶南市",
			StartYear: 1990,
			EndYear:   1994,
		},
	},
	379017: {
		{
			Address:   "蓬莱市",
			StartYear: 1991,
			EndYear:   1994,
		},
	},
	379018: {
		{
			Address:   "招远市",
			StartYear: 1991,
			EndYear:   1994,
		},
	},
	379019: {
		{
			Address:   "肥城市",
			StartYear: 1992,
			EndYear:   1994,
		},
	},
	379020: {
		{
			Address:   "章丘市",
			StartYear: 1992,
			EndYear:   1994,
		},
	},
	379021: {
		{
			Address:   "兖州市",
			StartYear: 1992,
			EndYear:   1994,
		},
	},
	379022: {
		{
			Address:   "邹城市",
			StartYear: 1992,
			EndYear:   1994,
		},
	},
	379023: {
		{
			Address:   "寿光市",
			StartYear: 1993,
			EndYear:   1994,
		},
	},
	379024: {
		{
			Address:   "乳山市",
			StartYear: 1993,
			EndYear:   1994,
		},
	},
	379025: {
		{
			Address:   "乐陵市",
			StartYear: 1994,
			EndYear:   1994,
		},
	},
	379026: {
		{
			Address:   "禹城市",
			StartYear: 1994,
			EndYear:   1994,
		},
	},
	379027: {
		{
			Address:   "安丘市",
			StartYear: 1994,
			EndYear:   1994,
		},
	},
	379028: {
		{
			Address:   "昌邑市",
			StartYear: 1994,
			EndYear:   1994,
		},
	},
	379029: {
		{
			Address:   "高密市",
			StartYear: 1994,
			EndYear:   1994,
		},
	},
	410000: {
		{
			Address: "河南省",
		},
	},
	410100: {
		{
			Address: "郑州市",
		},
	},
	410102: {
		{
			Address: "中原区",
		},
	},
	410103: {
		{
			Address: "二七区",
		},
	},
	410104: {
		{
			Address: "向阳回族区",
			EndYear: 1982,
		},
		{
			Address:   "管城回族区",
			StartYear: 1983,
		},
	},
	410105: {
		{
			Address: "金水区",
		},
	},
	410106: {
		{
			Address: "上街区",
		},
	},
	410107: {
		{
			Address:   "新密区",
			StartYear: 1982,
			EndYear:   1986,
		},
	},
	410108: {
		{
			Address:   "邙山区",
			StartYear: 1987,
			EndYear:   2002,
		},
		{
			Address:   "惠济区",
			StartYear: 2003,
		},
	},
	410111: {
		{
			Address:   "金海区",
			StartYear: 1981,
			EndYear:   1986,
		},
	},
	410112: {
		{
			Address: "郊区",
			EndYear: 1986,
		},
	},
	410121: {
		{
			Address: "荥阳县",
			EndYear: 1993,
		},
	},
	410122: {
		{
			Address:   "中牟县",
			StartYear: 1983,
		},
	},
	410123: {
		{
			Address:   "新郑县",
			StartYear: 1983,
			EndYear:   1993,
		},
	},
	410124: {
		{
			Address:   "巩县",
			StartYear: 1983,
			EndYear:   1990,
		},
	},
	410125: {
		{
			Address:   "登封县",
			StartYear: 1983,
			EndYear:   1993,
		},
	},
	410126: {
		{
			Address:   "密县",
			StartYear: 1983,
			EndYear:   1993,
		},
	},
	410181: {
		{
			Address:   "巩义市",
			StartYear: 1995,
		},
	},
	410182: {
		{
			Address:   "荥阳市",
			StartYear: 1995,
		},
	},
	410183: {
		{
			Address:   "新密市",
			StartYear: 1995,
		},
	},
	410184: {
		{
			Address:   "新郑市",
			StartYear: 1995,
		},
	},
	410185: {
		{
			Address:   "登封市",
			StartYear: 1995,
		},
	},
	410200: {
		{
			Address: "开封市",
		},
	},
	410202: {
		{
			Address: "龙亭区",
		},
	},
	410203: {
		{
			Address: "顺河回族区",
		},
	},
	410204: {
		{
			Address: "古楼区",
			EndYear: 1988,
		},
		{
			Address:   "鼓楼区",
			StartYear: 1989,
		},
	},
	410205: {
		{
			Address: "南关区",
			EndYear: 2004,
		},
		{
			Address:   "禹王台区",
			StartYear: 2005,
		},
	},
	410211: {
		{
			Address: "郊区",
			EndYear: 2004,
		},
		{
			Address:   "金明区",
			StartYear: 2005,
			EndYear:   2013,
		},
	},
	410212: {
		{
			Address:   "祥符区",
			StartYear: 2014,
		},
	},
	410221: {
		{
			Address:   "杞县",
			StartYear: 1983,
		},
	},
	410222: {
		{
			Address:   "通许县",
			StartYear: 1983,
		},
	},
	410223: {
		{
			Address:   "尉氏县",
			StartYear: 1983,
		},
	},
	410224: {
		{
			Address:   "开封县",
			StartYear: 1983,
			EndYear:   2013,
		},
	},
	410225: {
		{
			Address:   "兰考县",
			StartYear: 1983,
		},
	},
	410300: {
		{
			Address: "洛阳市",
		},
	},
	410302: {
		{
			Address: "洛北区",
			EndYear: 1982,
		},
		{
			Address:   "老城区",
			StartYear: 1983,
		},
	},
	410303: {
		{
			Address: "西工区",
		},
	},
	410304: {
		{
			Address: "瀍河回族区",
		},
	},
	410305: {
		{
			Address: "涧西区",
		},
	},
	410306: {
		{
			Address:   "吉利区",
			StartYear: 1982,
		},
	},
	410311: {
		{
			Address: "郊区",
			EndYear: 1999,
		},
		{
			Address:   "洛龙区",
			StartYear: 2000,
		},
	},
	410321: {
		{
			Address:   "偃师县",
			StartYear: 1983,
			EndYear:   1992,
		},
	},
	410322: {
		{
			Address:   "孟津县",
			StartYear: 1983,
		},
	},
	410323: {
		{
			Address:   "新安县",
			StartYear: 1983,
		},
	},
	410324: {
		{
			Address:   "栾川县",
			StartYear: 1986,
		},
	},
	410325: {
		{
			Address:   "嵩县",
			StartYear: 1986,
		},
	},
	410326: {
		{
			Address:   "汝阳县",
			StartYear: 1986,
		},
	},
	410327: {
		{
			Address:   "宜阳县",
			StartYear: 1986,
		},
	},
	410328: {
		{
			Address:   "洛宁县",
			StartYear: 1986,
		},
	},
	410329: {
		{
			Address:   "伊川县",
			StartYear: 1986,
		},
	},
	410381: {
		{
			Address:   "偃师市",
			StartYear: 1995,
		},
	},
	410400: {
		{
			Address: "平顶山市",
		},
	},
	410402: {
		{
			Address: "新华区",
		},
	},
	410403: {
		{
			Address: "卫东区",
		},
	},
	410404: {
		{
			Address:   "石龙区",
			StartYear: 1997,
		},
	},
	410411: {
		{
			Address: "郊区",
			EndYear: 1993,
		},
		{
			Address:   "湛河区",
			StartYear: 1994,
		},
	},
	410412: {
		{
			Address:   "舞钢区",
			StartYear: 1982,
			EndYear:   1989,
		},
	},
	410421: {
		{
			Address:   "宝丰县",
			StartYear: 1983,
		},
	},
	410422: {
		{
			Address:   "叶县",
			StartYear: 1983,
		},
	},
	410423: {
		{
			Address:   "鲁山县",
			StartYear: 1983,
		},
	},
	410424: {
		{
			Address:   "临汝县",
			StartYear: 1986,
			EndYear:   1987,
		},
	},
	410425: {
		{
			Address:   "郏县",
			StartYear: 1986,
		},
	},
	410426: {
		{
			Address:   "襄城县",
			StartYear: 1986,
			EndYear:   1996,
		},
	},
	410481: {
		{
			Address:   "舞钢市",
			StartYear: 1995,
		},
	},
	410482: {
		{
			Address:   "汝州市",
			StartYear: 1995,
		},
	},
	410500: {
		{
			Address: "鹤壁市",
			EndYear: 1981,
		},
		{
			Address:   "安阳市",
			StartYear: 1984,
		},
	},
	410502: {
		{
			Address: "鹤山区",
			EndYear: 1981,
		},
		{
			Address:   "文峰区",
			StartYear: 1984,
		},
	},
	410503: {
		{
			Address: "山城区",
			EndYear: 1981,
		},
		{
			Address:   "北关区",
			StartYear: 1984,
		},
	},
	410504: {
		{
			Address:   "铁西区",
			StartYear: 1984,
			EndYear:   2001,
		},
	},
	410505: {
		{
			Address:   "殷都区",
			StartYear: 2002,
		},
	},
	410506: {
		{
			Address:   "龙安区",
			StartYear: 2002,
		},
	},
	410511: {
		{
			Address: "郊区",
			EndYear: 2001,
		},
	},
	410521: {
		{
			Address:   "林县",
			StartYear: 1984,
			EndYear:   1993,
		},
	},
	410522: {
		{
			Address:   "安阳县",
			StartYear: 1984,
		},
	},
	410523: {
		{
			Address:   "汤阴县",
			StartYear: 1984,
		},
	},
	410524: {
		{
			Address:   "淇县",
			StartYear: 1984,
			EndYear:   1985,
		},
	},
	410525: {
		{
			Address:   "浚县",
			StartYear: 1984,
			EndYear:   1985,
		},
	},
	410526: {
		{
			Address:   "滑县",
			StartYear: 1986,
		},
	},
	410527: {
		{
			Address:   "内黄县",
			StartYear: 1986,
		},
	},
	410581: {
		{
			Address:   "林州市",
			StartYear: 1995,
		},
	},
	410600: {
		{
			Address: "焦作市",
			EndYear: 1981,
		},
		{
			Address:   "鹤壁市",
			StartYear: 1982,
		},
	},
	410602: {
		{
			Address: "解放区",
			EndYear: 1981,
		},
		{
			Address:   "鹤山区",
			StartYear: 1982,
		},
	},
	410603: {
		{
			Address: "中站区",
			EndYear: 1981,
		},
		{
			Address:   "山城区",
			StartYear: 1982,
		},
	},
	410604: {
		{
			Address: "马村区",
			EndYear: 1981,
		},
	},
	410611: {
		{
			Address: "郊区",
			EndYear: 2000,
		},
		{
			Address:   "淇滨区",
			StartYear: 2001,
		},
	},
	410621: {
		{
			Address:   "浚县",
			StartYear: 1986,
		},
	},
	410622: {
		{
			Address:   "淇县",
			StartYear: 1986,
		},
	},
	410700: {
		{
			Address:   "新乡市",
			StartYear: 1983,
		},
	},
	410702: {
		{
			Address:   "红旗区",
			StartYear: 1983,
		},
	},
	410703: {
		{
			Address:   "新华区",
			StartYear: 1983,
			EndYear:   2002,
		},
		{
			Address:   "卫滨区",
			StartYear: 2003,
		},
	},
	410704: {
		{
			Address:   "北站区",
			StartYear: 1983,
			EndYear:   2002,
		},
		{
			Address:   "凤泉区",
			StartYear: 2003,
		},
	},
	410711: {
		{
			Address:   "郊区",
			StartYear: 1983,
			EndYear:   2002,
		},
		{
			Address:   "牧野区",
			StartYear: 2003,
		},
	},
	410721: {
		{
			Address:   "新乡县",
			StartYear: 1983,
		},
	},
	410722: {
		{
			Address:   "汲县",
			StartYear: 1983,
			EndYear:   1987,
		},
	},
	410723: {
		{
			Address:   "辉县",
			StartYear: 1986,
			EndYear:   1987,
		},
	},
	410724: {
		{
			Address:   "获嘉县",
			StartYear: 1986,
		},
	},
	410725: {
		{
			Address:   "原阳县",
			StartYear: 1986,
		},
	},
	410726: {
		{
			Address:   "延津县",
			StartYear: 1986,
		},
	},
	410727: {
		{
			Address:   "封丘县",
			StartYear: 1986,
		},
	},
	410728: {
		{
			Address:   "长垣县",
			StartYear: 1986,
			EndYear:   2018,
		},
	},
	410781: {
		{
			Address:   "卫辉市",
			StartYear: 1995,
		},
	},
	410782: {
		{
			Address:   "辉县市",
			StartYear: 1995,
		},
	},
	410783: {
		{
			Address:   "长垣市",
			StartYear: 2019,
		},
	},
	410800: {
		{
			Address:   "焦作市",
			StartYear: 1982,
		},
	},
	410802: {
		{
			Address:   "解放区",
			StartYear: 1982,
		},
	},
	410803: {
		{
			Address:   "中站区",
			StartYear: 1982,
		},
	},
	410804: {
		{
			Address:   "马村区",
			StartYear: 1982,
		},
	},
	410811: {
		{
			Address:   "郊区",
			StartYear: 1982,
			EndYear:   1989,
		},
		{
			Address:   "山阳区",
			StartYear: 1990,
		},
	},
	410821: {
		{
			Address:   "修武县",
			StartYear: 1983,
		},
	},
	410822: {
		{
			Address:   "博爱县",
			StartYear: 1983,
		},
	},
	410823: {
		{
			Address:   "武陟县",
			StartYear: 1986,
		},
	},
	410824: {
		{
			Address:   "沁阳县",
			StartYear: 1986,
			EndYear:   1988,
		},
	},
	410825: {
		{
			Address:   "温县",
			StartYear: 1986,
		},
	},
	410826: {
		{
			Address:   "孟县",
			StartYear: 1986,
			EndYear:   1995,
		},
	},
	410827: {
		{
			Address:   "济源县",
			StartYear: 1986,
			EndYear:   1987,
		},
	},
	410881: {
		{
			Address:   "济源市",
			StartYear: 1995,
			EndYear:   1996,
		},
	},
	410882: {
		{
			Address:   "沁阳市",
			StartYear: 1995,
		},
	},
	410883: {
		{
			Address:   "孟州市",
			StartYear: 1996,
		},
	},
	410900: {
		{
			Address:   "濮阳市",
			StartYear: 1983,
		},
	},
	410902: {
		{
			Address:   "市区",
			StartYear: 1985,
			EndYear:   2001,
		},
		{
			Address:   "华龙区",
			StartYear: 2002,
		},
	},
	410911: {
		{
			Address:   "郊区",
			StartYear: 1984,
			EndYear:   1986,
		},
	},
	410921: {
		{
			Address:   "滑县",
			StartYear: 1983,
			EndYear:   1985,
		},
	},
	410922: {
		{
			Address:   "清丰县",
			StartYear: 1983,
		},
	},
	410923: {
		{
			Address:   "南乐县",
			StartYear: 1983,
		},
	},
	410924: {
		{
			Address:   "内黄县",
			StartYear: 1983,
			EndYear:   1985,
		},
	},
	410925: {
		{
			Address:   "长垣县",
			StartYear: 1983,
			EndYear:   1985,
		},
	},
	410926: {
		{
			Address:   "范县",
			StartYear: 1983,
		},
	},
	410927: {
		{
			Address:   "台前县",
			StartYear: 1983,
		},
	},
	410928: {
		{
			Address:   "濮阳县",
			StartYear: 1987,
		},
	},
	411000: {
		{
			Address:   "许昌市",
			StartYear: 1986,
		},
	},
	411002: {
		{
			Address:   "魏都区",
			StartYear: 1986,
		},
	},
	411003: {
		{
			Address:   "建安区",
			StartYear: 2016,
		},
	},
	411021: {
		{
			Address:   "禹县",
			StartYear: 1986,
			EndYear:   1987,
		},
	},
	411022: {
		{
			Address:   "长葛县",
			StartYear: 1986,
			EndYear:   1992,
		},
	},
	411023: {
		{
			Address:   "许昌县",
			StartYear: 1986,
			EndYear:   2015,
		},
	},
	411024: {
		{
			Address:   "鄢陵县",
			StartYear: 1986,
		},
	},
	411025: {
		{
			Address:   "襄城县",
			StartYear: 1997,
		},
	},
	411081: {
		{
			Address:   "禹州市",
			StartYear: 1995,
		},
	},
	411082: {
		{
			Address:   "长葛市",
			StartYear: 1995,
		},
	},
	411100: {
		{
			Address:   "漯河市",
			StartYear: 1986,
		},
	},
	411102: {
		{
			Address:   "源汇区",
			StartYear: 1986,
		},
	},
	411103: {
		{
			Address:   "郾城区",
			StartYear: 2004,
		},
	},
	411104: {
		{
			Address:   "召陵区",
			StartYear: 2004,
		},
	},
	411121: {
		{
			Address:   "舞阳县",
			StartYear: 1986,
		},
	},
	411122: {
		{
			Address:   "临颍县",
			StartYear: 1986,
		},
	},
	411123: {
		{
			Address:   "郾城县",
			StartYear: 1986,
			EndYear:   2003,
		},
	},
	411200: {
		{
			Address:   "三门峡市",
			StartYear: 1986,
		},
	},
	411202: {
		{
			Address:   "湖滨区",
			StartYear: 1986,
		},
	},
	411203: {
		{
			Address:   "陕州区",
			StartYear: 2015,
		},
	},
	411221: {
		{
			Address:   "渑池县",
			StartYear: 1986,
		},
	},
	411222: {
		{
			Address:   "陕县",
			StartYear: 1986,
			EndYear:   2014,
		},
	},
	411223: {
		{
			Address:   "灵宝县",
			StartYear: 1986,
			EndYear:   1992,
		},
	},
	411224: {
		{
			Address:   "卢氏县",
			StartYear: 1986,
		},
	},
	411281: {
		{
			Address:   "义马市",
			StartYear: 1995,
		},
	},
	411282: {
		{
			Address:   "灵宝市",
			StartYear: 1995,
		},
	},
	411300: {
		{
			Address:   "南阳市",
			StartYear: 1994,
		},
	},
	411302: {
		{
			Address:   "宛城区",
			StartYear: 1994,
		},
	},
	411303: {
		{
			Address:   "卧龙区",
			StartYear: 1994,
		},
	},
	411321: {
		{
			Address:   "南召县",
			StartYear: 1994,
		},
	},
	411322: {
		{
			Address:   "方城县",
			StartYear: 1994,
		},
	},
	411323: {
		{
			Address:   "西峡县",
			StartYear: 1994,
		},
	},
	411324: {
		{
			Address:   "镇平县",
			StartYear: 1994,
		},
	},
	411325: {
		{
			Address:   "内乡县",
			StartYear: 1994,
		},
	},
	411326: {
		{
			Address:   "淅川县",
			StartYear: 1994,
		},
	},
	411327: {
		{
			Address:   "社旗县",
			StartYear: 1994,
		},
	},
	411328: {
		{
			Address:   "唐河县",
			StartYear: 1994,
		},
	},
	411329: {
		{
			Address:   "新野县",
			StartYear: 1994,
		},
	},
	411330: {
		{
			Address:   "桐柏县",
			StartYear: 1994,
		},
	},
	411381: {
		{
			Address:   "邓州市",
			StartYear: 1995,
		},
	},
	411400: {
		{
			Address:   "商丘市",
			StartYear: 1997,
		},
	},
	411402: {
		{
			Address:   "梁园区",
			StartYear: 1997,
		},
	},
	411403: {
		{
			Address:   "睢阳区",
			StartYear: 1997,
		},
	},
	411421: {
		{
			Address:   "民权县",
			StartYear: 1997,
		},
	},
	411422: {
		{
			Address:   "睢县",
			StartYear: 1997,
		},
	},
	411423: {
		{
			Address:   "宁陵县",
			StartYear: 1997,
		},
	},
	411424: {
		{
			Address:   "柘城县",
			StartYear: 1997,
		},
	},
	411425: {
		{
			Address:   "虞城县",
			StartYear: 1997,
		},
	},
	411426: {
		{
			Address:   "夏邑县",
			StartYear: 1997,
		},
	},
	411481: {
		{
			Address:   "永城市",
			StartYear: 1997,
		},
	},
	411500: {
		{
			Address:   "信阳市",
			StartYear: 1998,
		},
	},
	411502: {
		{
			Address:   "浉河区",
			StartYear: 1998,
		},
	},
	411503: {
		{
			Address:   "平桥区",
			StartYear: 1998,
		},
	},
	411521: {
		{
			Address:   "罗山县",
			StartYear: 1998,
		},
	},
	411522: {
		{
			Address:   "光山县",
			StartYear: 1998,
		},
	},
	411523: {
		{
			Address:   "新县",
			StartYear: 1998,
		},
	},
	411524: {
		{
			Address:   "商城县",
			StartYear: 1998,
		},
	},
	411525: {
		{
			Address:   "固始县",
			StartYear: 1998,
		},
	},
	411526: {
		{
			Address:   "潢川县",
			StartYear: 1998,
		},
	},
	411527: {
		{
			Address:   "淮滨县",
			StartYear: 1998,
		},
	},
	411528: {
		{
			Address:   "息县",
			StartYear: 1998,
		},
	},
	411600: {
		{
			Address:   "周口市",
			StartYear: 2000,
		},
	},
	411602: {
		{
			Address:   "川汇区",
			StartYear: 2000,
		},
	},
	411603: {
		{
			Address:   "淮阳区",
			StartYear: 2019,
		},
	},
	411621: {
		{
			Address:   "扶沟县",
			StartYear: 2000,
		},
	},
	411622: {
		{
			Address:   "西华县",
			StartYear: 2000,
		},
	},
	411623: {
		{
			Address:   "商水县",
			StartYear: 2000,
		},
	},
	411624: {
		{
			Address:   "沈丘县",
			StartYear: 2000,
		},
	},
	411625: {
		{
			Address:   "郸城县",
			StartYear: 2000,
		},
	},
	411626: {
		{
			Address:   "淮阳县",
			StartYear: 2000,
			EndYear:   2018,
		},
	},
	411627: {
		{
			Address:   "太康县",
			StartYear: 2000,
		},
	},
	411628: {
		{
			Address:   "鹿邑县",
			StartYear: 2000,
		},
	},
	411681: {
		{
			Address:   "项城市",
			StartYear: 2000,
		},
	},
	411700: {
		{
			Address:   "驻马店市",
			StartYear: 2000,
		},
	},
	411702: {
		{
			Address:   "驿城区",
			StartYear: 2000,
		},
	},
	411721: {
		{
			Address:   "西平县",
			StartYear: 2000,
		},
	},
	411722: {
		{
			Address:   "上蔡县",
			StartYear: 2000,
		},
	},
	411723: {
		{
			Address:   "平舆县",
			StartYear: 2000,
		},
	},
	411724: {
		{
			Address:   "正阳县",
			StartYear: 2000,
		},
	},
	411725: {
		{
			Address:   "确山县",
			StartYear: 2000,
		},
	},
	411726: {
		{
			Address:   "泌阳县",
			StartYear: 2000,
		},
	},
	411727: {
		{
			Address:   "汝南县",
			StartYear: 2000,
		},
	},
	411728: {
		{
			Address:   "遂平县",
			StartYear: 2000,
		},
	},
	411729: {
		{
			Address:   "新蔡县",
			StartYear: 2000,
		},
	},
	412100: {
		{
			Address: "安阳地区",
			EndYear: 1982,
		},
	},
	412101: {
		{
			Address: "安阳市",
			EndYear: 1983,
		},
	},
	412102: {
		{
			Address: "文峰区",
			EndYear: 1983,
		},
	},
	412103: {
		{
			Address: "北关区",
			EndYear: 1983,
		},
	},
	412104: {
		{
			Address: "铁西区",
			EndYear: 1983,
		},
	},
	412111: {
		{
			Address: "郊区",
			EndYear: 1983,
		},
	},
	412121: {
		{
			Address: "林县",
			EndYear: 1983,
		},
	},
	412122: {
		{
			Address: "安阳县",
			EndYear: 1983,
		},
	},
	412123: {
		{
			Address: "汤阴县",
			EndYear: 1983,
		},
	},
	412124: {
		{
			Address: "淇县",
			EndYear: 1983,
		},
	},
	412125: {
		{
			Address: "浚县",
			EndYear: 1983,
		},
	},
	412126: {
		{
			Address: "濮阳县",
			EndYear: 1982,
		},
	},
	412127: {
		{
			Address: "滑县",
			EndYear: 1982,
		},
	},
	412128: {
		{
			Address: "清丰县",
			EndYear: 1982,
		},
	},
	412129: {
		{
			Address: "南乐县",
			EndYear: 1982,
		},
	},
	412130: {
		{
			Address: "内黄县",
			EndYear: 1982,
		},
	},
	412131: {
		{
			Address: "长垣县",
			EndYear: 1982,
		},
	},
	412132: {
		{
			Address: "范县",
			EndYear: 1982,
		},
	},
	412133: {
		{
			Address: "台前县",
			EndYear: 1982,
		},
	},
	412200: {
		{
			Address: "新乡地区",
			EndYear: 1985,
		},
	},
	412201: {
		{
			Address: "新乡市",
			EndYear: 1982,
		},
	},
	412202: {
		{
			Address: "红旗区",
			EndYear: 1982,
		},
	},
	412203: {
		{
			Address: "新华区",
			EndYear: 1982,
		},
	},
	412204: {
		{
			Address:   "北站区",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	412211: {
		{
			Address: "郊区",
			EndYear: 1982,
		},
	},
	412221: {
		{
			Address: "沁阳县",
			EndYear: 1985,
		},
	},
	412222: {
		{
			Address: "博爱县",
			EndYear: 1982,
		},
	},
	412223: {
		{
			Address: "济源县",
			EndYear: 1985,
		},
	},
	412224: {
		{
			Address: "孟县",
			EndYear: 1985,
		},
	},
	412225: {
		{
			Address: "温县",
			EndYear: 1985,
		},
	},
	412226: {
		{
			Address: "武陟县",
			EndYear: 1985,
		},
	},
	412227: {
		{
			Address: "修武县",
			EndYear: 1982,
		},
	},
	412228: {
		{
			Address: "获嘉县",
			EndYear: 1985,
		},
	},
	412229: {
		{
			Address: "新乡县",
			EndYear: 1982,
		},
	},
	412230: {
		{
			Address: "辉县",
			EndYear: 1985,
		},
	},
	412231: {
		{
			Address: "汲县",
			EndYear: 1982,
		},
	},
	412232: {
		{
			Address: "原阳县",
			EndYear: 1985,
		},
	},
	412233: {
		{
			Address: "延津县",
			EndYear: 1985,
		},
	},
	412234: {
		{
			Address: "封丘县",
			EndYear: 1985,
		},
	},
	412300: {
		{
			Address: "商丘地区",
			EndYear: 1996,
		},
	},
	412301: {
		{
			Address: "商丘市",
			EndYear: 1996,
		},
	},
	412302: {
		{
			Address:   "永城市",
			StartYear: 1996,
			EndYear:   1996,
		},
	},
	412321: {
		{
			Address: "虞城县",
			EndYear: 1996,
		},
	},
	412322: {
		{
			Address: "商丘县",
			EndYear: 1996,
		},
	},
	412323: {
		{
			Address: "民权县",
			EndYear: 1996,
		},
	},
	412324: {
		{
			Address: "宁陵县",
			EndYear: 1996,
		},
	},
	412325: {
		{
			Address: "睢县",
			EndYear: 1996,
		},
	},
	412326: {
		{
			Address: "夏邑县",
			EndYear: 1996,
		},
	},
	412327: {
		{
			Address: "柘城县",
			EndYear: 1996,
		},
	},
	412328: {
		{
			Address: "永城县",
			EndYear: 1995,
		},
	},
	412400: {
		{
			Address: "开封地区",
			EndYear: 1982,
		},
	},
	412421: {
		{
			Address: "杞县",
			EndYear: 1982,
		},
	},
	412422: {
		{
			Address: "通许县",
			EndYear: 1982,
		},
	},
	412423: {
		{
			Address: "尉氏县",
			EndYear: 1982,
		},
	},
	412424: {
		{
			Address: "开封县",
			EndYear: 1982,
		},
	},
	412425: {
		{
			Address: "中牟县",
			EndYear: 1982,
		},
	},
	412426: {
		{
			Address: "新郑县",
			EndYear: 1981,
		},
		{
			Address:   "巩县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	412427: {
		{
			Address: "巩县",
			EndYear: 1981,
		},
		{
			Address:   "登封县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	412428: {
		{
			Address: "登封县",
			EndYear: 1981,
		},
		{
			Address:   "新郑县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	412429: {
		{
			Address: "密县",
			EndYear: 1982,
		},
	},
	412430: {
		{
			Address: "兰考县",
			EndYear: 1982,
		},
	},
	412500: {
		{
			Address: "洛阳地区",
			EndYear: 1985,
		},
	},
	412501: {
		{
			Address: "三门峡市",
			EndYear: 1985,
		},
	},
	412502: {
		{
			Address:   "义马市",
			StartYear: 1981,
			EndYear:   1985,
		},
	},
	412521: {
		{
			Address: "偃师县",
			EndYear: 1982,
		},
	},
	412522: {
		{
			Address: "孟津县",
			EndYear: 1982,
		},
	},
	412523: {
		{
			Address: "新安县",
			EndYear: 1982,
		},
	},
	412524: {
		{
			Address: "渑池县",
			EndYear: 1985,
		},
	},
	412525: {
		{
			Address: "陕县",
			EndYear: 1985,
		},
	},
	412526: {
		{
			Address: "灵宝县",
			EndYear: 1985,
		},
	},
	412527: {
		{
			Address: "伊川县",
			EndYear: 1985,
		},
	},
	412528: {
		{
			Address: "汝阳县",
			EndYear: 1985,
		},
	},
	412529: {
		{
			Address: "嵩县",
			EndYear: 1985,
		},
	},
	412530: {
		{
			Address: "洛宁县",
			EndYear: 1985,
		},
	},
	412531: {
		{
			Address: "卢氏县",
			EndYear: 1985,
		},
	},
	412532: {
		{
			Address: "栾川县",
			EndYear: 1985,
		},
	},
	412533: {
		{
			Address: "临汝县",
			EndYear: 1985,
		},
	},
	412534: {
		{
			Address: "宜阳县",
			EndYear: 1985,
		},
	},
	412600: {
		{
			Address: "许昌地区",
			EndYear: 1985,
		},
	},
	412601: {
		{
			Address: "许昌市",
			EndYear: 1985,
		},
	},
	412602: {
		{
			Address: "漯河市",
			EndYear: 1985,
		},
	},
	412621: {
		{
			Address: "长葛县",
			EndYear: 1985,
		},
	},
	412622: {
		{
			Address: "禹县",
			EndYear: 1985,
		},
	},
	412623: {
		{
			Address: "鄢陵县",
			EndYear: 1985,
		},
	},
	412624: {
		{
			Address: "许昌县",
			EndYear: 1985,
		},
	},
	412625: {
		{
			Address: "郏县",
			EndYear: 1985,
		},
	},
	412626: {
		{
			Address: "临颍县",
			EndYear: 1985,
		},
	},
	412627: {
		{
			Address: "襄城县",
			EndYear: 1985,
		},
	},
	412628: {
		{
			Address: "宝丰县",
			EndYear: 1982,
		},
	},
	412629: {
		{
			Address: "郾城县",
			EndYear: 1985,
		},
	},
	412630: {
		{
			Address: "叶县",
			EndYear: 1982,
		},
	},
	412631: {
		{
			Address: "鲁山县",
			EndYear: 1982,
		},
	},
	412632: {
		{
			Address: "舞阳县",
			EndYear: 1985,
		},
	},
	412700: {
		{
			Address: "周口地区",
			EndYear: 1999,
		},
	},
	412701: {
		{
			Address: "周口市",
			EndYear: 1999,
		},
	},
	412702: {
		{
			Address:   "项城市",
			StartYear: 1993,
			EndYear:   1999,
		},
	},
	412721: {
		{
			Address: "扶沟县",
			EndYear: 1999,
		},
	},
	412722: {
		{
			Address: "西华县",
			EndYear: 1999,
		},
	},
	412723: {
		{
			Address: "商水县",
			EndYear: 1999,
		},
	},
	412724: {
		{
			Address: "太康县",
			EndYear: 1999,
		},
	},
	412725: {
		{
			Address: "鹿邑县",
			EndYear: 1999,
		},
	},
	412726: {
		{
			Address: "郸城县",
			EndYear: 1999,
		},
	},
	412727: {
		{
			Address: "淮阳县",
			EndYear: 1999,
		},
	},
	412728: {
		{
			Address: "沈丘县",
			EndYear: 1999,
		},
	},
	412729: {
		{
			Address: "项城县",
			EndYear: 1992,
		},
	},
	412800: {
		{
			Address: "驻马店地区",
			EndYear: 1999,
		},
	},
	412801: {
		{
			Address: "驻马店市",
			EndYear: 1999,
		},
	},
	412821: {
		{
			Address: "确山县",
			EndYear: 1999,
		},
	},
	412822: {
		{
			Address: "泌阳县",
			EndYear: 1999,
		},
	},
	412823: {
		{
			Address: "遂平县",
			EndYear: 1999,
		},
	},
	412824: {
		{
			Address: "西平县",
			EndYear: 1999,
		},
	},
	412825: {
		{
			Address: "上蔡县",
			EndYear: 1999,
		},
	},
	412826: {
		{
			Address: "汝南县",
			EndYear: 1999,
		},
	},
	412827: {
		{
			Address: "平舆县",
			EndYear: 1999,
		},
	},
	412828: {
		{
			Address: "新蔡县",
			EndYear: 1999,
		},
	},
	412829: {
		{
			Address: "正阳县",
			EndYear: 1999,
		},
	},
	412900: {
		{
			Address: "南阳地区",
			EndYear: 1993,
		},
	},
	412901: {
		{
			Address: "南阳市",
			EndYear: 1993,
		},
	},
	412902: {
		{
			Address:   "邓州市",
			StartYear: 1988,
			EndYear:   1993,
		},
	},
	412921: {
		{
			Address: "南召县",
			EndYear: 1993,
		},
	},
	412922: {
		{
			Address: "方城县",
			EndYear: 1993,
		},
	},
	412923: {
		{
			Address: "西峡县",
			EndYear: 1993,
		},
	},
	412924: {
		{
			Address: "南阳县",
			EndYear: 1993,
		},
	},
	412925: {
		{
			Address: "镇平县",
			EndYear: 1993,
		},
	},
	412926: {
		{
			Address: "内乡县",
			EndYear: 1993,
		},
	},
	412927: {
		{
			Address: "淅川县",
			EndYear: 1993,
		},
	},
	412928: {
		{
			Address: "社旗县",
			EndYear: 1993,
		},
	},
	412929: {
		{
			Address: "唐河县",
			EndYear: 1993,
		},
	},
	412930: {
		{
			Address: "邓县",
			EndYear: 1987,
		},
	},
	412931: {
		{
			Address: "新野县",
			EndYear: 1993,
		},
	},
	412932: {
		{
			Address: "桐柏县",
			EndYear: 1993,
		},
	},
	413000: {
		{
			Address: "信阳地区",
			EndYear: 1997,
		},
	},
	413001: {
		{
			Address: "信阳市",
			EndYear: 1997,
		},
	},
	413021: {
		{
			Address: "息县",
			EndYear: 1997,
		},
	},
	413022: {
		{
			Address: "淮滨县",
			EndYear: 1997,
		},
	},
	413023: {
		{
			Address: "信阳县",
			EndYear: 1997,
		},
	},
	413024: {
		{
			Address: "潢川县",
			EndYear: 1997,
		},
	},
	413025: {
		{
			Address: "光山县",
			EndYear: 1997,
		},
	},
	413026: {
		{
			Address: "固始县",
			EndYear: 1997,
		},
	},
	413027: {
		{
			Address: "商城县",
			EndYear: 1997,
		},
	},
	413028: {
		{
			Address: "罗山县",
			EndYear: 1997,
		},
	},
	413029: {
		{
			Address: "新县",
			EndYear: 1997,
		},
	},
	419001: {
		{
			Address:   "义马市",
			StartYear: 1986,
			EndYear:   1994,
		},
		{
			Address:   "济源市",
			StartYear: 1997,
		},
	},
	419002: {
		{
			Address:   "汝州市",
			StartYear: 1988,
			EndYear:   1994,
		},
	},
	419003: {
		{
			Address:   "济源市",
			StartYear: 1988,
			EndYear:   1994,
		},
	},
	419004: {
		{
			Address:   "禹州市",
			StartYear: 1988,
			EndYear:   1994,
		},
	},
	419005: {
		{
			Address:   "卫辉市",
			StartYear: 1988,
			EndYear:   1994,
		},
	},
	419006: {
		{
			Address:   "辉县市",
			StartYear: 1988,
			EndYear:   1994,
		},
	},
	419007: {
		{
			Address:   "沁阳市",
			StartYear: 1989,
			EndYear:   1994,
		},
	},
	419008: {
		{
			Address:   "舞钢市",
			StartYear: 1990,
			EndYear:   1994,
		},
	},
	419009: {
		{
			Address:   "巩义市",
			StartYear: 1991,
			EndYear:   1994,
		},
	},
	419010: {
		{
			Address:   "灵宝市",
			StartYear: 1993,
			EndYear:   1994,
		},
	},
	419011: {
		{
			Address:   "长葛市",
			StartYear: 1993,
			EndYear:   1994,
		},
	},
	419012: {
		{
			Address:   "偃师市",
			StartYear: 1993,
			EndYear:   1994,
		},
	},
	419013: {
		{
			Address:   "邓州市",
			StartYear: 1994,
			EndYear:   1994,
		},
	},
	419014: {
		{
			Address:   "林州市",
			StartYear: 1994,
			EndYear:   1994,
		},
	},
	419015: {
		{
			Address:   "新密市",
			StartYear: 1994,
			EndYear:   1994,
		},
	},
	419016: {
		{
			Address:   "荥阳市",
			StartYear: 1994,
			EndYear:   1994,
		},
	},
	419017: {
		{
			Address:   "新郑市",
			StartYear: 1994,
			EndYear:   1994,
		},
	},
	419018: {
		{
			Address:   "登封市",
			StartYear: 1994,
			EndYear:   1994,
		},
	},
	420000: {
		{
			Address: "湖北省",
		},
	},
	420100: {
		{
			Address: "武汉市",
		},
	},
	420102: {
		{
			Address: "江岸区",
		},
	},
	420103: {
		{
			Address: "江汉区",
		},
	},
	420104: {
		{
			Address: "硚口区",
		},
	},
	420105: {
		{
			Address: "汉阳区",
		},
	},
	420106: {
		{
			Address: "武昌区",
		},
	},
	420107: {
		{
			Address: "青山区",
		},
	},
	420111: {
		{
			Address: "洪山区",
		},
	},
	420112: {
		{
			Address: "东西湖区",
		},
	},
	420113: {
		{
			Address:   "汉南区",
			StartYear: 1984,
		},
	},
	420114: {
		{
			Address:   "蔡甸区",
			StartYear: 1992,
		},
	},
	420115: {
		{
			Address:   "江夏区",
			StartYear: 1995,
		},
	},
	420116: {
		{
			Address:   "黄陂区",
			StartYear: 1998,
		},
	},
	420117: {
		{
			Address:   "新洲区",
			StartYear: 1998,
		},
	},
	420121: {
		{
			Address: "汉阳县",
			EndYear: 1991,
		},
	},
	420122: {
		{
			Address: "武昌县",
			EndYear: 1994,
		},
	},
	420123: {
		{
			Address:   "黄陂县",
			StartYear: 1983,
			EndYear:   1997,
		},
	},
	420124: {
		{
			Address:   "新洲县",
			StartYear: 1983,
			EndYear:   1997,
		},
	},
	420200: {
		{
			Address: "黄石市",
		},
	},
	420202: {
		{
			Address: "黄石港区",
		},
	},
	420203: {
		{
			Address: "石灰窑区",
			EndYear: 2000,
		},
		{
			Address:   "西塞山区",
			StartYear: 2001,
		},
	},
	420204: {
		{
			Address: "下陆区",
		},
	},
	420205: {
		{
			Address: "铁山区",
		},
	},
	420211: {
		{
			Address: "郊区",
			EndYear: 1982,
		},
	},
	420221: {
		{
			Address: "大冶县",
			EndYear: 1993,
		},
	},
	420222: {
		{
			Address:   "阳新县",
			StartYear: 1996,
		},
	},
	420281: {
		{
			Address:   "大冶市",
			StartYear: 1995,
		},
	},
	420300: {
		{
			Address: "十堰市",
		},
	},
	420302: {
		{
			Address:   "茅箭区",
			StartYear: 1984,
		},
	},
	420303: {
		{
			Address:   "张湾区",
			StartYear: 1984,
		},
	},
	420304: {
		{
			Address:   "郧阳区",
			StartYear: 2014,
		},
	},
	420321: {
		{
			Address:   "郧县",
			StartYear: 1994,
			EndYear:   2013,
		},
	},
	420322: {
		{
			Address:   "郧西县",
			StartYear: 1994,
		},
	},
	420323: {
		{
			Address:   "竹山县",
			StartYear: 1994,
		},
	},
	420324: {
		{
			Address:   "竹溪县",
			StartYear: 1994,
		},
	},
	420325: {
		{
			Address:   "房县",
			StartYear: 1994,
		},
	},
	420381: {
		{
			Address:   "丹江口市",
			StartYear: 1995,
		},
	},
	420400: {
		{
			Address: "沙市市",
			EndYear: 1993,
		},
	},
	420500: {
		{
			Address: "宜昌市",
		},
	},
	420502: {
		{
			Address:   "西陵区",
			StartYear: 1986,
		},
	},
	420503: {
		{
			Address:   "伍家岗区",
			StartYear: 1986,
		},
	},
	420504: {
		{
			Address:   "点军区",
			StartYear: 1986,
		},
	},
	420505: {
		{
			Address:   "猇亭区",
			StartYear: 1995,
		},
	},
	420506: {
		{
			Address:   "夷陵区",
			StartYear: 2001,
		},
	},
	420521: {
		{
			Address:   "宜昌县",
			StartYear: 1992,
			EndYear:   2000,
		},
	},
	420523: {
		{
			Address:   "枝江县",
			StartYear: 1992,
			EndYear:   1995,
		},
	},
	420525: {
		{
			Address:   "远安县",
			StartYear: 1992,
		},
	},
	420526: {
		{
			Address:   "兴山县",
			StartYear: 1992,
		},
	},
	420527: {
		{
			Address:   "秭归县",
			StartYear: 1992,
		},
	},
	420528: {
		{
			Address:   "长阳土家族自治县",
			StartYear: 1992,
		},
	},
	420529: {
		{
			Address:   "五峰土家族自治县",
			StartYear: 1992,
		},
	},
	420581: {
		{
			Address:   "枝城市",
			StartYear: 1996,
			EndYear:   1997,
		},
		{
			Address:   "宜都市",
			StartYear: 1998,
		},
	},
	420582: {
		{
			Address:   "当阳市",
			StartYear: 1995,
		},
	},
	420583: {
		{
			Address:   "枝城市",
			StartYear: 1995,
			EndYear:   1995,
		},
		{
			Address:   "枝江市",
			StartYear: 1996,
		},
	},
	420600: {
		{
			Address: "襄樊市",
			EndYear: 2009,
		},
		{
			Address:   "襄阳市",
			StartYear: 2010,
		},
	},
	420602: {
		{
			Address:   "襄城区",
			StartYear: 1984,
		},
	},
	420603: {
		{
			Address:   "樊东区",
			StartYear: 1984,
			EndYear:   1994,
		},
	},
	420604: {
		{
			Address:   "樊西区",
			StartYear: 1984,
			EndYear:   1994,
		},
	},
	420606: {
		{
			Address:   "樊城区",
			StartYear: 1995,
		},
	},
	420607: {
		{
			Address:   "襄阳区",
			StartYear: 2001,
			EndYear:   2009,
		},
		{
			Address:   "襄州区",
			StartYear: 2010,
		},
	},
	420611: {
		{
			Address:   "郊区",
			StartYear: 1984,
			EndYear:   1994,
		},
	},
	420621: {
		{
			Address:   "襄阳县",
			StartYear: 1983,
			EndYear:   2000,
		},
	},
	420622: {
		{
			Address:   "枣阳县",
			StartYear: 1983,
			EndYear:   1987,
		},
	},
	420623: {
		{
			Address:   "宜城县",
			StartYear: 1983,
			EndYear:   1993,
		},
	},
	420624: {
		{
			Address:   "南漳县",
			StartYear: 1983,
		},
	},
	420625: {
		{
			Address:   "谷城县",
			StartYear: 1983,
		},
	},
	420626: {
		{
			Address:   "保康县",
			StartYear: 1983,
		},
	},
	420682: {
		{
			Address:   "老河口市",
			StartYear: 1995,
		},
	},
	420683: {
		{
			Address:   "枣阳市",
			StartYear: 1995,
		},
	},
	420684: {
		{
			Address:   "宜城市",
			StartYear: 1995,
		},
	},
	420700: {
		{
			Address:   "鄂州市",
			StartYear: 1983,
		},
	},
	420702: {
		{
			Address:   "梁子湖区",
			StartYear: 1987,
		},
	},
	420703: {
		{
			Address:   "黄州区",
			StartYear: 1984,
			EndYear:   1986,
		},
		{
			Address:   "华容区",
			StartYear: 1987,
		},
	},
	420704: {
		{
			Address:   "鄂城区",
			StartYear: 1984,
		},
	},
	420800: {
		{
			Address:   "荆门市",
			StartYear: 1983,
		},
	},
	420802: {
		{
			Address:   "东宝区",
			StartYear: 1985,
		},
	},
	420803: {
		{
			Address:   "沙洋区",
			StartYear: 1985,
			EndYear:   1997,
		},
	},
	420804: {
		{
			Address:   "掇刀区",
			StartYear: 2001,
		},
	},
	420821: {
		{
			Address:   "京山县",
			StartYear: 1996,
			EndYear:   2017,
		},
	},
	420822: {
		{
			Address:   "沙洋县",
			StartYear: 1998,
		},
	},
	420881: {
		{
			Address:   "钟祥市",
			StartYear: 1996,
		},
	},
	420882: {
		{
			Address:   "京山市",
			StartYear: 2018,
		},
	},
	420900: {
		{
			Address:   "孝感市",
			StartYear: 1993,
		},
	},
	420901: {
		{
			Address:   "随州市",
			StartYear: 1983,
			EndYear:   1983,
		},
	},
	420902: {
		{
			Address:   "孝南区",
			StartYear: 1993,
		},
	},
	420921: {
		{
			Address:   "孝昌县",
			StartYear: 1993,
		},
	},
	420922: {
		{
			Address:   "大悟县",
			StartYear: 1993,
		},
	},
	420923: {
		{
			Address:   "云梦县",
			StartYear: 1993,
		},
	},
	420924: {
		{
			Address:   "汉川县",
			StartYear: 1993,
			EndYear:   1996,
		},
	},
	420981: {
		{
			Address:   "应城市",
			StartYear: 1995,
		},
	},
	420982: {
		{
			Address:   "安陆市",
			StartYear: 1995,
		},
	},
	420983: {
		{
			Address:   "广水市",
			StartYear: 1995,
			EndYear:   1999,
		},
	},
	420984: {
		{
			Address:   "汉川市",
			StartYear: 1997,
		},
	},
	421000: {
		{
			Address:   "荆沙市",
			StartYear: 1994,
			EndYear:   1995,
		},
		{
			Address:   "荆州市",
			StartYear: 1996,
		},
	},
	421001: {
		{
			Address:   "老河口市",
			StartYear: 1983,
			EndYear:   1983,
		},
	},
	421002: {
		{
			Address:   "沙市区",
			StartYear: 1994,
		},
	},
	421003: {
		{
			Address:   "荆州区",
			StartYear: 1994,
		},
	},
	421004: {
		{
			Address:   "江陵区",
			StartYear: 1994,
			EndYear:   1997,
		},
	},
	421022: {
		{
			Address:   "公安县",
			StartYear: 1994,
		},
	},
	421023: {
		{
			Address:   "监利县",
			StartYear: 1994,
			EndYear:   2019,
		},
	},
	421024: {
		{
			Address:   "松滋县",
			StartYear: 1994,
			EndYear:   1994,
		},
		{
			Address:   "江陵县",
			StartYear: 1998,
		},
	},
	421025: {
		{
			Address:   "京山县",
			StartYear: 1994,
			EndYear:   1995,
		},
	},
	421081: {
		{
			Address:   "石首市",
			StartYear: 1995,
		},
	},
	421082: {
		{
			Address:   "钟祥市",
			StartYear: 1995,
			EndYear:   1995,
		},
	},
	421083: {
		{
			Address:   "洪湖市",
			StartYear: 1995,
		},
	},
	421087: {
		{
			Address:   "松滋市",
			StartYear: 1995,
		},
	},
	421088: {
		{
			Address:   "监利市",
			StartYear: 2020,
		},
	},
	421100: {
		{
			Address: "黄冈市",
		},
	},
	421102: {
		{
			Address: "黄州区",
		},
	},
	421121: {
		{
			Address: "团风县",
		},
	},
	421122: {
		{
			Address: "红安县",
		},
	},
	421123: {
		{
			Address: "罗田县",
		},
	},
	421124: {
		{
			Address: "英山县",
		},
	},
	421125: {
		{
			Address: "浠水县",
		},
	},
	421126: {
		{
			Address: "蕲春县",
		},
	},
	421127: {
		{
			Address: "黄梅县",
		},
	},
	421181: {
		{
			Address: "麻城市",
		},
	},
	421182: {
		{
			Address: "武穴市",
		},
	},
	421200: {
		{
			Address:   "咸宁市",
			StartYear: 1998,
		},
	},
	421202: {
		{
			Address:   "咸安区",
			StartYear: 1998,
		},
	},
	421221: {
		{
			Address:   "嘉鱼县",
			StartYear: 1998,
		},
	},
	421222: {
		{
			Address:   "通城县",
			StartYear: 1998,
		},
	},
	421223: {
		{
			Address:   "崇阳县",
			StartYear: 1998,
		},
	},
	421224: {
		{
			Address:   "通山县",
			StartYear: 1998,
		},
	},
	421281: {
		{
			Address:   "赤壁市",
			StartYear: 1998,
		},
	},
	421300: {
		{
			Address:   "随州市",
			StartYear: 2000,
		},
	},
	421303: {
		{
			Address:   "曾都区",
			StartYear: 2000,
		},
	},
	421321: {
		{
			Address:   "随县",
			StartYear: 2009,
		},
	},
	421381: {
		{
			Address:   "广水市",
			StartYear: 2000,
		},
	},
	422100: {
		{
			Address: "黄冈地区",
			EndYear: 1994,
		},
	},
	422101: {
		{
			Address: "鄂城市",
			EndYear: 1982,
		},
		{
			Address:   "麻城市",
			StartYear: 1986,
			EndYear:   1994,
		},
	},
	422102: {
		{
			Address:   "武穴市",
			StartYear: 1987,
			EndYear:   1994,
		},
	},
	422103: {
		{
			Address:   "黄州市",
			StartYear: 1990,
			EndYear:   1994,
		},
	},
	422121: {
		{
			Address: "黄冈县",
			EndYear: 1989,
		},
	},
	422122: {
		{
			Address: "新洲县",
			EndYear: 1982,
		},
	},
	422123: {
		{
			Address: "红安县",
			EndYear: 1994,
		},
	},
	422124: {
		{
			Address: "麻城县",
			EndYear: 1985,
		},
	},
	422125: {
		{
			Address: "罗田县",
			EndYear: 1994,
		},
	},
	422126: {
		{
			Address: "英山县",
			EndYear: 1994,
		},
	},
	422127: {
		{
			Address: "浠水县",
			EndYear: 1994,
		},
	},
	422128: {
		{
			Address: "蕲春县",
			EndYear: 1994,
		},
	},
	422129: {
		{
			Address: "广济县",
			EndYear: 1986,
		},
	},
	422130: {
		{
			Address: "黄梅县",
			EndYear: 1994,
		},
	},
	422131: {
		{
			Address: "鄂城县",
			EndYear: 1982,
		},
	},
	422200: {
		{
			Address: "孝感地区",
			EndYear: 1992,
		},
	},
	422201: {
		{
			Address:   "孝感市",
			StartYear: 1983,
			EndYear:   1992,
		},
	},
	422202: {
		{
			Address:   "应城市",
			StartYear: 1986,
			EndYear:   1992,
		},
	},
	422203: {
		{
			Address:   "安陆市",
			StartYear: 1987,
			EndYear:   1992,
		},
	},
	422204: {
		{
			Address:   "广水市",
			StartYear: 1988,
			EndYear:   1992,
		},
	},
	422221: {
		{
			Address: "孝感县",
			EndYear: 1982,
		},
	},
	422222: {
		{
			Address: "黄陂县",
			EndYear: 1982,
		},
	},
	422223: {
		{
			Address: "大悟县",
			EndYear: 1992,
		},
	},
	422224: {
		{
			Address: "应山县",
			EndYear: 1987,
		},
	},
	422225: {
		{
			Address: "安陆县",
			EndYear: 1986,
		},
	},
	422226: {
		{
			Address: "云梦县",
			EndYear: 1992,
		},
	},
	422227: {
		{
			Address: "应城县",
			EndYear: 1985,
		},
	},
	422228: {
		{
			Address: "汉川县",
			EndYear: 1992,
		},
	},
	422300: {
		{
			Address: "咸宁地区",
			EndYear: 1997,
		},
	},
	422301: {
		{
			Address:   "咸宁市",
			StartYear: 1983,
			EndYear:   1997,
		},
	},
	422302: {
		{
			Address:   "蒲圻市",
			StartYear: 1986,
			EndYear:   1997,
		},
	},
	422321: {
		{
			Address: "咸宁县",
			EndYear: 1982,
		},
	},
	422322: {
		{
			Address: "嘉鱼县",
			EndYear: 1997,
		},
	},
	422323: {
		{
			Address: "蒲圻县",
			EndYear: 1985,
		},
	},
	422324: {
		{
			Address: "通城县",
			EndYear: 1997,
		},
	},
	422325: {
		{
			Address: "崇阳县",
			EndYear: 1997,
		},
	},
	422326: {
		{
			Address: "通山县",
			EndYear: 1997,
		},
	},
	422327: {
		{
			Address: "阳新县",
			EndYear: 1995,
		},
	},
	422400: {
		{
			Address: "荆州地区",
			EndYear: 1993,
		},
	},
	422401: {
		{
			Address: "荆门市",
			EndYear: 1982,
		},
		{
			Address:   "仙桃市",
			StartYear: 1986,
			EndYear:   1993,
		},
	},
	422402: {
		{
			Address:   "石首市",
			StartYear: 1986,
			EndYear:   1993,
		},
	},
	422403: {
		{
			Address:   "洪湖市",
			StartYear: 1987,
			EndYear:   1993,
		},
	},
	422404: {
		{
			Address:   "天门市",
			StartYear: 1987,
			EndYear:   1993,
		},
	},
	422405: {
		{
			Address:   "潜江市",
			StartYear: 1988,
			EndYear:   1993,
		},
	},
	422406: {
		{
			Address:   "钟祥市",
			StartYear: 1992,
			EndYear:   1993,
		},
	},
	422421: {
		{
			Address: "江陵县",
			EndYear: 1993,
		},
	},
	422422: {
		{
			Address: "松滋县",
			EndYear: 1993,
		},
	},
	422423: {
		{
			Address: "公安县",
			EndYear: 1993,
		},
	},
	422424: {
		{
			Address: "石首县",
			EndYear: 1985,
		},
	},
	422425: {
		{
			Address: "监利县",
			EndYear: 1993,
		},
	},
	422426: {
		{
			Address: "洪湖县",
			EndYear: 1986,
		},
	},
	422427: {
		{
			Address: "沔阳县",
			EndYear: 1985,
		},
	},
	422428: {
		{
			Address: "天门县",
			EndYear: 1986,
		},
	},
	422429: {
		{
			Address: "潜江县",
			EndYear: 1987,
		},
	},
	422430: {
		{
			Address: "荆门县",
			EndYear: 1982,
		},
	},
	422431: {
		{
			Address: "钟祥县",
			EndYear: 1991,
		},
	},
	422432: {
		{
			Address: "京山县",
			EndYear: 1993,
		},
	},
	422500: {
		{
			Address: "襄阳地区",
			EndYear: 1982,
		},
	},
	422501: {
		{
			Address: "随州市",
			EndYear: 1982,
		},
	},
	422502: {
		{
			Address: "老河口市",
			EndYear: 1982,
		},
	},
	422521: {
		{
			Address: "樊阳县",
			EndYear: 1982,
		},
	},
	422522: {
		{
			Address: "枣阳县",
			EndYear: 1982,
		},
	},
	422523: {
		{
			Address: "随县",
			EndYear: 1982,
		},
	},
	422524: {
		{
			Address: "宜城县",
			EndYear: 1982,
		},
	},
	422525: {
		{
			Address: "南漳县",
			EndYear: 1982,
		},
	},
	422526: {
		{
			Address: "光化县",
			EndYear: 1982,
		},
	},
	422527: {
		{
			Address: "谷城县",
			EndYear: 1982,
		},
	},
	422528: {
		{
			Address: "保康县",
			EndYear: 1982,
		},
	},
	422600: {
		{
			Address: "郧阳地区",
			EndYear: 1993,
		},
	},
	422601: {
		{
			Address:   "丹江口市",
			StartYear: 1983,
			EndYear:   1993,
		},
	},
	422621: {
		{
			Address: "均县",
			EndYear: 1982,
		},
	},
	422622: {
		{
			Address: "郧县",
			EndYear: 1993,
		},
	},
	422623: {
		{
			Address: "郧西县",
			EndYear: 1993,
		},
	},
	422624: {
		{
			Address: "竹山县",
			EndYear: 1993,
		},
	},
	422625: {
		{
			Address: "竹溪县",
			EndYear: 1993,
		},
	},
	422626: {
		{
			Address: "房县",
			EndYear: 1993,
		},
	},
	422627: {
		{
			Address: "神农架林区",
			EndYear: 1982,
		},
	},
	422700: {
		{
			Address: "宜昌地区",
			EndYear: 1991,
		},
	},
	422701: {
		{
			Address:   "枝城市",
			StartYear: 1987,
			EndYear:   1991,
		},
	},
	422702: {
		{
			Address:   "当阳市",
			StartYear: 1988,
			EndYear:   1991,
		},
	},
	422721: {
		{
			Address: "宜昌县",
			EndYear: 1991,
		},
	},
	422722: {
		{
			Address: "宜都县",
			EndYear: 1986,
		},
	},
	422723: {
		{
			Address: "枝江县",
			EndYear: 1991,
		},
	},
	422724: {
		{
			Address: "当阳县",
			EndYear: 1987,
		},
	},
	422725: {
		{
			Address: "远安县",
			EndYear: 1991,
		},
	},
	422726: {
		{
			Address: "兴山县",
			EndYear: 1991,
		},
	},
	422727: {
		{
			Address: "秭归县",
			EndYear: 1991,
		},
	},
	422728: {
		{
			Address: "长阳县",
			EndYear: 1983,
		},
		{
			Address:   "长阳土家族自治县",
			StartYear: 1984,
			EndYear:   1991,
		},
	},
	422729: {
		{
			Address: "五峰县",
			EndYear: 1983,
		},
		{
			Address:   "五峰土家族自治县",
			StartYear: 1984,
			EndYear:   1991,
		},
	},
	422800: {
		{
			Address: "恩施地区",
			EndYear: 1982,
		},
		{
			Address:   "鄂西土家族苗族自治州",
			StartYear: 1983,
			EndYear:   1992,
		},
		{
			Address:   "恩施土家族苗族自治州",
			StartYear: 1993,
		},
	},
	422801: {
		{
			Address:   "恩施市",
			StartYear: 1981,
		},
	},
	422802: {
		{
			Address:   "利川市",
			StartYear: 1986,
		},
	},
	422821: {
		{
			Address: "恩施县",
			EndYear: 1982,
		},
	},
	422822: {
		{
			Address: "建始县",
		},
	},
	422823: {
		{
			Address: "巴东县",
		},
	},
	422824: {
		{
			Address: "利川县",
			EndYear: 1985,
		},
	},
	422825: {
		{
			Address: "宣恩县",
		},
	},
	422826: {
		{
			Address: "咸丰县",
		},
	},
	422827: {
		{
			Address: "来凤土家族自治县",
			EndYear: 1982,
		},
		{
			Address:   "来凤县",
			StartYear: 1983,
		},
	},
	422828: {
		{
			Address: "鹤峰土家族自治县",
			EndYear: 1982,
		},
		{
			Address:   "鹤峰县",
			StartYear: 1983,
		},
	},
	422921: {
		{
			Address:   "神农架林区",
			StartYear: 1983,
			EndYear:   1983,
		},
	},
	429001: {
		{
			Address:   "随州市",
			StartYear: 1984,
			EndYear:   1999,
		},
	},
	429002: {
		{
			Address:   "老河口市",
			StartYear: 1984,
			EndYear:   1994,
		},
	},
	429003: {
		{
			Address:   "枣阳市",
			StartYear: 1988,
			EndYear:   1994,
		},
	},
	429004: {
		{
			Address:   "仙桃市",
			StartYear: 1994,
		},
	},
	429005: {
		{
			Address:   "潜江市",
			StartYear: 1994,
		},
	},
	429006: {
		{
			Address:   "天门市",
			StartYear: 1994,
		},
	},
	429007: {
		{
			Address:   "枝城市",
			StartYear: 1992,
			EndYear:   1994,
		},
	},
	429008: {
		{
			Address:   "当阳市",
			StartYear: 1992,
			EndYear:   1994,
		},
	},
	429009: {
		{
			Address:   "应城市",
			StartYear: 1993,
			EndYear:   1994,
		},
	},
	429010: {
		{
			Address:   "安陆市",
			StartYear: 1993,
			EndYear:   1994,
		},
	},
	429011: {
		{
			Address:   "广水市",
			StartYear: 1993,
			EndYear:   1994,
		},
	},
	429012: {
		{
			Address:   "石首市",
			StartYear: 1994,
			EndYear:   1994,
		},
	},
	429013: {
		{
			Address:   "洪湖市",
			StartYear: 1994,
			EndYear:   1994,
		},
	},
	429014: {
		{
			Address:   "钟祥市",
			StartYear: 1994,
			EndYear:   1994,
		},
	},
	429015: {
		{
			Address:   "丹江口市",
			StartYear: 1994,
			EndYear:   1994,
		},
	},
	429016: {
		{
			Address:   "大冶市",
			StartYear: 1994,
			EndYear:   1994,
		},
	},
	429017: {
		{
			Address:   "宜城市",
			StartYear: 1994,
			EndYear:   1994,
		},
	},
	429021: {
		{
			Address:   "神农架林区",
			StartYear: 1984,
		},
	},
	430000: {
		{
			Address: "湖南省",
		},
	},
	430100: {
		{
			Address: "长沙市",
		},
	},
	430102: {
		{
			Address: "城东区",
			EndYear: 1982,
		},
		{
			Address:   "东区",
			StartYear: 1983,
			EndYear:   1995,
		},
		{
			Address:   "芙蓉区",
			StartYear: 1996,
		},
	},
	430103: {
		{
			Address: "城南区",
			EndYear: 1982,
		},
		{
			Address:   "南区",
			StartYear: 1983,
			EndYear:   1995,
		},
		{
			Address:   "天心区",
			StartYear: 1996,
		},
	},
	430104: {
		{
			Address: "城西区",
			EndYear: 1982,
		},
		{
			Address:   "西区",
			StartYear: 1983,
			EndYear:   1995,
		},
		{
			Address:   "岳麓区",
			StartYear: 1996,
		},
	},
	430105: {
		{
			Address: "城北区",
			EndYear: 1982,
		},
		{
			Address:   "北区",
			StartYear: 1983,
			EndYear:   1995,
		},
		{
			Address:   "开福区",
			StartYear: 1996,
		},
	},
	430111: {
		{
			Address: "郊区",
			EndYear: 1995,
		},
		{
			Address:   "雨花区",
			StartYear: 1996,
		},
	},
	430112: {
		{
			Address:   "望城区",
			StartYear: 2011,
		},
	},
	430121: {
		{
			Address: "长沙县",
		},
	},
	430122: {
		{
			Address: "望城县",
			EndYear: 2010,
		},
	},
	430123: {
		{
			Address:   "浏阳县",
			StartYear: 1983,
			EndYear:   1992,
		},
	},
	430124: {
		{
			Address:   "宁乡县",
			StartYear: 1983,
			EndYear:   2016,
		},
	},
	430181: {
		{
			Address:   "浏阳市",
			StartYear: 1995,
		},
	},
	430182: {
		{
			Address:   "宁乡市",
			StartYear: 2017,
		},
	},
	430200: {
		{
			Address: "株洲市",
		},
	},
	430202: {
		{
			Address: "东区",
			EndYear: 1996,
		},
		{
			Address:   "荷塘区",
			StartYear: 1997,
		},
	},
	430203: {
		{
			Address: "北区",
			EndYear: 1996,
		},
		{
			Address:   "芦淞区",
			StartYear: 1997,
		},
	},
	430204: {
		{
			Address: "南区",
			EndYear: 1996,
		},
		{
			Address:   "石峰区",
			StartYear: 1997,
		},
	},
	430211: {
		{
			Address: "郊区",
			EndYear: 1996,
		},
		{
			Address:   "天元区",
			StartYear: 1997,
		},
	},
	430212: {
		{
			Address:   "渌口区",
			StartYear: 2018,
		},
	},
	430219: {
		{
			Address:   "醴陵市",
			StartYear: 1985,
			EndYear:   1985,
		},
	},
	430221: {
		{
			Address: "株洲县",
			EndYear: 2017,
		},
	},
	430222: {
		{
			Address:   "醴陵县",
			StartYear: 1983,
			EndYear:   1984,
		},
	},
	430223: {
		{
			Address:   "攸县",
			StartYear: 1983,
		},
	},
	430224: {
		{
			Address:   "茶陵县",
			StartYear: 1983,
		},
	},
	430225: {
		{
			Address:   "酃县",
			StartYear: 1983,
			EndYear:   1993,
		},
		{
			Address:   "炎陵县",
			StartYear: 1994,
		},
	},
	430281: {
		{
			Address:   "醴陵市",
			StartYear: 1995,
		},
	},
	430300: {
		{
			Address: "湘潭市",
		},
		{
			Address:   "邵阳市",
			StartYear: 1982,
			EndYear:   1983,
		},
	},
	430302: {
		{
			Address: "雨湖区",
		},
		{
			Address:   "东区",
			StartYear: 1982,
			EndYear:   1983,
		},
	},
	430303: {
		{
			Address: "湘江区",
			EndYear: 1991,
		},
		{
			Address:   "西区",
			StartYear: 1982,
			EndYear:   1983,
		},
	},
	430304: {
		{
			Address: "岳塘区",
		},
		{
			Address:   "桥头区",
			StartYear: 1982,
			EndYear:   1983,
		},
	},
	430305: {
		{
			Address: "板塘区",
			EndYear: 1991,
		},
	},
	430306: {
		{
			Address:   "韶山区",
			StartYear: 1988,
			EndYear:   1989,
		},
	},
	430311: {
		{
			Address: "郊区",
			EndYear: 1991,
		},
	},
	430321: {
		{
			Address:   "邵东县",
			StartYear: 1983,
			EndYear:   1983,
		},
		{
			Address:   "湘潭县",
			StartYear: 1984,
		},
	},
	430322: {
		{
			Address:   "新邵县",
			StartYear: 1983,
			EndYear:   1983,
		},
		{
			Address:   "湘乡县",
			StartYear: 1984,
			EndYear:   1985,
		},
	},
	430381: {
		{
			Address:   "湘乡市",
			StartYear: 1995,
		},
	},
	430382: {
		{
			Address:   "韶山市",
			StartYear: 1995,
		},
	},
	430400: {
		{
			Address: "衡阳市",
		},
		{
			Address:   "湘潭市",
			StartYear: 1982,
			EndYear:   1983,
		},
	},
	430402: {
		{
			Address: "江东区",
			EndYear: 2000,
		},
		{
			Address:   "雨湖区",
			StartYear: 1982,
			EndYear:   1983,
		},
	},
	430403: {
		{
			Address: "城南区",
			EndYear: 2000,
		},
		{
			Address:   "湘江区",
			StartYear: 1982,
			EndYear:   1983,
		},
	},
	430404: {
		{
			Address: "城北区",
			EndYear: 2000,
		},
		{
			Address:   "岳塘区",
			StartYear: 1982,
			EndYear:   1983,
		},
	},
	430405: {
		{
			Address:   "板塘区",
			StartYear: 1982,
			EndYear:   1983,
		},
		{
			Address:   "珠晖区",
			StartYear: 2001,
		},
	},
	430406: {
		{
			Address:   "雁峰区",
			StartYear: 2001,
		},
	},
	430407: {
		{
			Address:   "石鼓区",
			StartYear: 2001,
		},
	},
	430408: {
		{
			Address:   "蒸湘区",
			StartYear: 2001,
		},
	},
	430411: {
		{
			Address: "郊区",
			EndYear: 2000,
		},
	},
	430412: {
		{
			Address:   "南岳区",
			StartYear: 1988,
		},
	},
	430421: {
		{
			Address:   "湘潭县",
			StartYear: 1983,
			EndYear:   1983,
		},
		{
			Address:   "衡阳县",
			StartYear: 1984,
		},
	},
	430422: {
		{
			Address:   "湘乡县",
			StartYear: 1983,
			EndYear:   1983,
		},
		{
			Address:   "衡南县",
			StartYear: 1984,
		},
	},
	430423: {
		{
			Address:   "衡山县",
			StartYear: 1984,
		},
	},
	430424: {
		{
			Address:   "衡东县",
			StartYear: 1984,
		},
	},
	430425: {
		{
			Address:   "常宁县",
			StartYear: 1984,
			EndYear:   1995,
		},
	},
	430426: {
		{
			Address:   "祁东县",
			StartYear: 1984,
		},
	},
	430427: {
		{
			Address:   "耒阳县",
			StartYear: 1984,
			EndYear:   1985,
		},
	},
	430481: {
		{
			Address:   "耒阳市",
			StartYear: 1995,
		},
	},
	430482: {
		{
			Address:   "常宁市",
			StartYear: 1996,
		},
	},
	430500: {
		{
			Address: "邵阳市",
		},
		{
			Address:   "衡阳市",
			StartYear: 1982,
			EndYear:   1983,
		},
	},
	430502: {
		{
			Address: "东区",
			EndYear: 1996,
		},
		{
			Address:   "江东区",
			StartYear: 1982,
			EndYear:   1983,
		},
		{
			Address:   "双清区",
			StartYear: 1997,
		},
	},
	430503: {
		{
			Address: "西区",
			EndYear: 1996,
		},
		{
			Address:   "城南区",
			StartYear: 1982,
			EndYear:   1983,
		},
		{
			Address:   "大祥区",
			StartYear: 1997,
		},
	},
	430504: {
		{
			Address: "桥头区",
			EndYear: 1986,
		},
		{
			Address:   "城北区",
			StartYear: 1982,
			EndYear:   1983,
		},
	},
	430511: {
		{
			Address: "郊区",
			EndYear: 1996,
		},
		{
			Address:   "北塔区",
			StartYear: 1997,
		},
	},
	430521: {
		{
			Address:   "衡阳县",
			StartYear: 1983,
			EndYear:   1983,
		},
		{
			Address:   "邵东县",
			StartYear: 1984,
			EndYear:   2018,
		},
	},
	430522: {
		{
			Address:   "衡南县",
			StartYear: 1983,
			EndYear:   1983,
		},
		{
			Address:   "新邵县",
			StartYear: 1984,
		},
	},
	430523: {
		{
			Address:   "衡山县",
			StartYear: 1983,
			EndYear:   1983,
		},
		{
			Address:   "邵阳县",
			StartYear: 1986,
		},
	},
	430524: {
		{
			Address:   "衡东县",
			StartYear: 1983,
			EndYear:   1983,
		},
		{
			Address:   "隆回县",
			StartYear: 1986,
		},
	},
	430525: {
		{
			Address:   "常宁县",
			StartYear: 1983,
			EndYear:   1983,
		},
		{
			Address:   "洞口县",
			StartYear: 1986,
		},
	},
	430526: {
		{
			Address:   "祁东县",
			StartYear: 1983,
			EndYear:   1983,
		},
		{
			Address:   "武冈县",
			StartYear: 1986,
			EndYear:   1993,
		},
	},
	430527: {
		{
			Address:   "耒阳县",
			StartYear: 1983,
			EndYear:   1983,
		},
		{
			Address:   "绥宁县",
			StartYear: 1986,
		},
	},
	430528: {
		{
			Address:   "新宁县",
			StartYear: 1986,
		},
	},
	430529: {
		{
			Address:   "城步苗族自治县",
			StartYear: 1986,
		},
	},
	430581: {
		{
			Address:   "武冈市",
			StartYear: 1995,
		},
	},
	430582: {
		{
			Address:   "邵东市",
			StartYear: 2019,
		},
	},
	430600: {
		{
			Address:   "岳阳市",
			StartYear: 1983,
		},
	},
	430602: {
		{
			Address:   "南区",
			StartYear: 1984,
			EndYear:   1995,
		},
		{
			Address:   "岳阳楼区",
			StartYear: 1996,
		},
	},
	430603: {
		{
			Address:   "北区",
			StartYear: 1984,
			EndYear:   1995,
		},
		{
			Address:   "云溪区",
			StartYear: 1996,
		},
	},
	430611: {
		{
			Address:   "郊区",
			StartYear: 1984,
			EndYear:   1995,
		},
		{
			Address:   "君山区",
			StartYear: 1996,
		},
	},
	430621: {
		{
			Address:   "岳阳县",
			StartYear: 1983,
		},
	},
	430622: {
		{
			Address:   "临湘县",
			StartYear: 1986,
			EndYear:   1991,
		},
	},
	430623: {
		{
			Address:   "华容县",
			StartYear: 1986,
		},
	},
	430624: {
		{
			Address:   "湘阴县",
			StartYear: 1986,
		},
	},
	430626: {
		{
			Address:   "平江县",
			StartYear: 1986,
		},
	},
	430627: {
		{
			Address:   "汨罗县",
			StartYear: 1986,
			EndYear:   1986,
		},
	},
	430681: {
		{
			Address:   "汨罗市",
			StartYear: 1995,
		},
	},
	430682: {
		{
			Address:   "临湘市",
			StartYear: 1995,
		},
	},
	430700: {
		{
			Address:   "常德市",
			StartYear: 1988,
		},
	},
	430702: {
		{
			Address:   "武陵区",
			StartYear: 1988,
		},
	},
	430703: {
		{
			Address:   "鼎城区",
			StartYear: 1988,
		},
	},
	430721: {
		{
			Address:   "安乡县",
			StartYear: 1988,
		},
	},
	430722: {
		{
			Address:   "汉寿县",
			StartYear: 1988,
		},
	},
	430723: {
		{
			Address:   "澧县",
			StartYear: 1988,
		},
	},
	430724: {
		{
			Address:   "临澧县",
			StartYear: 1988,
		},
	},
	430725: {
		{
			Address:   "桃源县",
			StartYear: 1988,
		},
	},
	430726: {
		{
			Address:   "石门县",
			StartYear: 1988,
		},
	},
	430781: {
		{
			Address:   "津市市",
			StartYear: 1995,
		},
	},
	430800: {
		{
			Address:   "大庸市",
			StartYear: 1988,
			EndYear:   1993,
		},
		{
			Address:   "张家界市",
			StartYear: 1994,
		},
	},
	430802: {
		{
			Address:   "永定区",
			StartYear: 1988,
		},
	},
	430811: {
		{
			Address:   "武陵源区",
			StartYear: 1988,
		},
	},
	430821: {
		{
			Address:   "慈利县",
			StartYear: 1988,
		},
	},
	430822: {
		{
			Address:   "桑植县",
			StartYear: 1988,
		},
	},
	430900: {
		{
			Address:   "益阳市",
			StartYear: 1994,
		},
	},
	430902: {
		{
			Address:   "资阳区",
			StartYear: 1994,
		},
	},
	430903: {
		{
			Address:   "赫山区",
			StartYear: 1994,
		},
	},
	430921: {
		{
			Address:   "南县",
			StartYear: 1994,
		},
	},
	430922: {
		{
			Address:   "桃江县",
			StartYear: 1994,
		},
	},
	430923: {
		{
			Address:   "安化县",
			StartYear: 1994,
		},
	},
	430981: {
		{
			Address:   "沅江市",
			StartYear: 1995,
		},
	},
	431000: {
		{
			Address:   "郴州市",
			StartYear: 1994,
		},
	},
	431002: {
		{
			Address:   "北湖区",
			StartYear: 1994,
		},
	},
	431003: {
		{
			Address:   "苏仙区",
			StartYear: 1994,
		},
	},
	431021: {
		{
			Address:   "桂阳县",
			StartYear: 1994,
		},
	},
	431022: {
		{
			Address:   "宜章县",
			StartYear: 1994,
		},
	},
	431023: {
		{
			Address:   "永兴县",
			StartYear: 1994,
		},
	},
	431024: {
		{
			Address:   "嘉禾县",
			StartYear: 1994,
		},
	},
	431025: {
		{
			Address:   "临武县",
			StartYear: 1994,
		},
	},
	431026: {
		{
			Address:   "汝城县",
			StartYear: 1994,
		},
	},
	431027: {
		{
			Address:   "桂东县",
			StartYear: 1994,
		},
	},
	431028: {
		{
			Address:   "安仁县",
			StartYear: 1994,
		},
	},
	431081: {
		{
			Address:   "资兴市",
			StartYear: 1995,
		},
	},
	431100: {
		{
			Address:   "永州市",
			StartYear: 1995,
		},
	},
	431102: {
		{
			Address:   "芝山区",
			StartYear: 1995,
			EndYear:   2004,
		},
		{
			Address:   "零陵区",
			StartYear: 2005,
		},
	},
	431103: {
		{
			Address:   "冷水滩区",
			StartYear: 1995,
		},
	},
	431121: {
		{
			Address:   "祁阳县",
			StartYear: 1995,
		},
	},
	431122: {
		{
			Address:   "东安县",
			StartYear: 1995,
		},
	},
	431123: {
		{
			Address:   "双牌县",
			StartYear: 1995,
		},
	},
	431124: {
		{
			Address:   "道县",
			StartYear: 1995,
		},
	},
	431125: {
		{
			Address:   "江永县",
			StartYear: 1995,
		},
	},
	431126: {
		{
			Address:   "宁远县",
			StartYear: 1995,
		},
	},
	431127: {
		{
			Address:   "蓝山县",
			StartYear: 1995,
		},
	},
	431128: {
		{
			Address:   "新田县",
			StartYear: 1995,
		},
	},
	431129: {
		{
			Address:   "江华瑶族自治县",
			StartYear: 1995,
		},
	},
	431200: {
		{
			Address:   "怀化市",
			StartYear: 1997,
		},
	},
	431202: {
		{
			Address:   "鹤城区",
			StartYear: 1997,
		},
	},
	431221: {
		{
			Address:   "中方县",
			StartYear: 1997,
		},
	},
	431222: {
		{
			Address:   "沅陵县",
			StartYear: 1997,
		},
	},
	431223: {
		{
			Address:   "辰溪县",
			StartYear: 1997,
		},
	},
	431224: {
		{
			Address:   "溆浦县",
			StartYear: 1997,
		},
	},
	431225: {
		{
			Address:   "会同县",
			StartYear: 1997,
		},
	},
	431226: {
		{
			Address:   "麻阳苗族自治县",
			StartYear: 1997,
		},
	},
	431227: {
		{
			Address:   "新晃侗族自治县",
			StartYear: 1997,
		},
	},
	431228: {
		{
			Address:   "芷江侗族自治县",
			StartYear: 1997,
		},
	},
	431229: {
		{
			Address:   "靖州苗族侗族自治县",
			StartYear: 1997,
		},
	},
	431230: {
		{
			Address:   "通道侗族自治县",
			StartYear: 1997,
		},
	},
	431281: {
		{
			Address:   "洪江市",
			StartYear: 1997,
		},
	},
	431300: {
		{
			Address:   "娄底市",
			StartYear: 1999,
		},
	},
	431302: {
		{
			Address:   "娄星区",
			StartYear: 1999,
		},
	},
	431321: {
		{
			Address:   "双峰县",
			StartYear: 1999,
		},
	},
	431322: {
		{
			Address:   "新化县",
			StartYear: 1999,
		},
	},
	431381: {
		{
			Address:   "冷水江市",
			StartYear: 1999,
		},
	},
	431382: {
		{
			Address:   "涟源市",
			StartYear: 1999,
		},
	},
	432100: {
		{
			Address: "湘潭地区",
			EndYear: 1982,
		},
	},
	432121: {
		{
			Address: "湘潭县",
			EndYear: 1982,
		},
	},
	432122: {
		{
			Address: "湘乡县",
			EndYear: 1982,
		},
	},
	432123: {
		{
			Address: "醴陵县",
			EndYear: 1982,
		},
	},
	432124: {
		{
			Address: "浏阳县",
			EndYear: 1982,
		},
	},
	432125: {
		{
			Address: "攸县",
			EndYear: 1982,
		},
	},
	432126: {
		{
			Address: "茶陵县",
			EndYear: 1982,
		},
	},
	432127: {
		{
			Address: "酃县",
			EndYear: 1982,
		},
	},
	432200: {
		{
			Address: "岳阳地区",
			EndYear: 1985,
		},
	},
	432201: {
		{
			Address: "岳阳市",
			EndYear: 1982,
		},
	},
	432221: {
		{
			Address: "岳阳县",
			EndYear: 1980,
		},
	},
	432222: {
		{
			Address: "平江县",
			EndYear: 1985,
		},
	},
	432223: {
		{
			Address: "湘阴县",
			EndYear: 1985,
		},
	},
	432224: {
		{
			Address: "汨罗县",
			EndYear: 1985,
		},
	},
	432225: {
		{
			Address: "临湘县",
			EndYear: 1985,
		},
	},
	432226: {
		{
			Address: "华容县",
			EndYear: 1985,
		},
	},
	432300: {
		{
			Address: "益阳地区",
			EndYear: 1993,
		},
	},
	432301: {
		{
			Address: "益阳市",
			EndYear: 1993,
		},
	},
	432302: {
		{
			Address:   "沅江市",
			StartYear: 1988,
			EndYear:   1993,
		},
	},
	432321: {
		{
			Address: "益阳县",
			EndYear: 1993,
		},
	},
	432322: {
		{
			Address: "南县",
			EndYear: 1993,
		},
	},
	432323: {
		{
			Address: "沅江县",
			EndYear: 1987,
		},
	},
	432324: {
		{
			Address: "宁乡县",
			EndYear: 1982,
		},
	},
	432325: {
		{
			Address: "桃江县",
			EndYear: 1993,
		},
	},
	432326: {
		{
			Address: "安化县",
			EndYear: 1993,
		},
	},
	432400: {
		{
			Address: "常德地区",
			EndYear: 1987,
		},
	},
	432401: {
		{
			Address: "常德市",
			EndYear: 1987,
		},
	},
	432402: {
		{
			Address: "津市市",
			EndYear: 1987,
		},
	},
	432421: {
		{
			Address: "常德县",
			EndYear: 1987,
		},
	},
	432422: {
		{
			Address: "安乡县",
			EndYear: 1987,
		},
	},
	432423: {
		{
			Address: "汉寿县",
			EndYear: 1987,
		},
	},
	432424: {
		{
			Address: "澧县",
			EndYear: 1987,
		},
	},
	432425: {
		{
			Address: "临澧县",
			EndYear: 1987,
		},
	},
	432426: {
		{
			Address: "桃源县",
			EndYear: 1987,
		},
	},
	432427: {
		{
			Address: "石门县",
			EndYear: 1987,
		},
	},
	432428: {
		{
			Address: "慈利县",
			EndYear: 1987,
		},
	},
	432500: {
		{
			Address: "涟源地区",
			EndYear: 1981,
		},
		{
			Address:   "娄底地区",
			StartYear: 1982,
			EndYear:   1998,
		},
	},
	432501: {
		{
			Address: "娄底市",
			EndYear: 1998,
		},
	},
	432502: {
		{
			Address: "冷水江市",
			EndYear: 1998,
		},
	},
	432503: {
		{
			Address:   "涟源市",
			StartYear: 1987,
			EndYear:   1998,
		},
	},
	432521: {
		{
			Address: "涟源县",
			EndYear: 1986,
		},
	},
	432522: {
		{
			Address: "双峰县",
			EndYear: 1998,
		},
	},
	432523: {
		{
			Address: "邵东县",
			EndYear: 1982,
		},
	},
	432524: {
		{
			Address: "新化县",
			EndYear: 1998,
		},
	},
	432525: {
		{
			Address: "新邵县",
			EndYear: 1982,
		},
	},
	432600: {
		{
			Address: "邵阳地区",
			EndYear: 1985,
		},
	},
	432621: {
		{
			Address: "邵阳县",
			EndYear: 1985,
		},
	},
	432622: {
		{
			Address: "隆回县",
			EndYear: 1985,
		},
	},
	432623: {
		{
			Address: "武冈县",
			EndYear: 1985,
		},
	},
	432624: {
		{
			Address: "洞口县",
			EndYear: 1985,
		},
	},
	432625: {
		{
			Address: "新宁县",
			EndYear: 1985,
		},
	},
	432626: {
		{
			Address: "绥宁县",
			EndYear: 1985,
		},
	},
	432627: {
		{
			Address: "城步苗族自治县",
			EndYear: 1985,
		},
	},
	432700: {
		{
			Address: "衡阳地区",
			EndYear: 1982,
		},
	},
	432721: {
		{
			Address: "衡阳县",
			EndYear: 1982,
		},
	},
	432722: {
		{
			Address: "衡南县",
			EndYear: 1982,
		},
	},
	432723: {
		{
			Address: "衡山县",
			EndYear: 1982,
		},
	},
	432724: {
		{
			Address: "衡东县",
			EndYear: 1982,
		},
	},
	432725: {
		{
			Address: "常宁县",
			EndYear: 1982,
		},
	},
	432726: {
		{
			Address: "祁东县",
			EndYear: 1982,
		},
	},
	432727: {
		{
			Address: "祁阳县",
			EndYear: 1982,
		},
	},
	432800: {
		{
			Address: "郴州地区",
			EndYear: 1993,
		},
	},
	432801: {
		{
			Address: "郴州市",
			EndYear: 1993,
		},
	},
	432802: {
		{
			Address:   "资兴市",
			StartYear: 1984,
			EndYear:   1993,
		},
	},
	432821: {
		{
			Address: "郴县",
			EndYear: 1993,
		},
	},
	432822: {
		{
			Address: "桂阳县",
			EndYear: 1993,
		},
	},
	432823: {
		{
			Address: "永兴县",
			EndYear: 1993,
		},
	},
	432824: {
		{
			Address: "宜章县",
			EndYear: 1993,
		},
	},
	432825: {
		{
			Address: "资兴县",
			EndYear: 1983,
		},
	},
	432826: {
		{
			Address: "嘉禾县",
			EndYear: 1993,
		},
	},
	432827: {
		{
			Address: "临武县",
			EndYear: 1993,
		},
	},
	432828: {
		{
			Address: "汝城县",
			EndYear: 1993,
		},
	},
	432829: {
		{
			Address: "桂东县",
			EndYear: 1993,
		},
	},
	432830: {
		{
			Address: "耒阳县",
			EndYear: 1982,
		},
	},
	432831: {
		{
			Address: "安仁县",
			EndYear: 1993,
		},
	},
	432900: {
		{
			Address: "零陵地区",
			EndYear: 1994,
		},
	},
	432901: {
		{
			Address:   "永州市",
			StartYear: 1982,
			EndYear:   1994,
		},
	},
	432902: {
		{
			Address:   "冷水滩市",
			StartYear: 1984,
			EndYear:   1994,
		},
	},
	432921: {
		{
			Address: "零陵县",
			EndYear: 1983,
		},
	},
	432922: {
		{
			Address: "东安县",
			EndYear: 1994,
		},
	},
	432923: {
		{
			Address: "道县",
			EndYear: 1994,
		},
	},
	432924: {
		{
			Address: "宁远县",
			EndYear: 1994,
		},
	},
	432925: {
		{
			Address: "江永县",
			EndYear: 1994,
		},
	},
	432926: {
		{
			Address: "江华瑶族自治县",
			EndYear: 1994,
		},
	},
	432927: {
		{
			Address: "蓝山县",
			EndYear: 1994,
		},
	},
	432928: {
		{
			Address: "新田县",
			EndYear: 1994,
		},
	},
	432929: {
		{
			Address: "双牌县",
			EndYear: 1994,
		},
	},
	432930: {
		{
			Address:   "祁阳县",
			StartYear: 1983,
			EndYear:   1994,
		},
	},
	433000: {
		{
			Address: "黔阳地区",
			EndYear: 1980,
		},
		{
			Address:   "怀化地区",
			StartYear: 1981,
			EndYear:   1996,
		},
	},
	433001: {
		{
			Address: "洪江市",
			EndYear: 1981,
		},
		{
			Address:   "怀化市",
			StartYear: 1982,
			EndYear:   1996,
		},
	},
	433002: {
		{
			Address: "怀化市",
			EndYear: 1981,
		},
		{
			Address:   "洪江市",
			StartYear: 1982,
			EndYear:   1996,
		},
	},
	433021: {
		{
			Address: "黔阳县",
			EndYear: 1996,
		},
	},
	433022: {
		{
			Address: "沅陵县",
			EndYear: 1996,
		},
	},
	433023: {
		{
			Address: "辰溪县",
			EndYear: 1996,
		},
	},
	433024: {
		{
			Address: "溆浦县",
			EndYear: 1996,
		},
	},
	433025: {
		{
			Address: "麻阳县",
			EndYear: 1987,
		},
		{
			Address:   "麻阳苗族自治县",
			StartYear: 1988,
			EndYear:   1996,
		},
	},
	433026: {
		{
			Address: "新晃侗族自治县",
			EndYear: 1996,
		},
	},
	433027: {
		{
			Address: "芷江县",
			EndYear: 1985,
		},
		{
			Address:   "芷江侗族自治县",
			StartYear: 1986,
			EndYear:   1996,
		},
	},
	433028: {
		{
			Address: "怀化县",
			EndYear: 1981,
		},
	},
	433029: {
		{
			Address: "会同县",
			EndYear: 1996,
		},
	},
	433030: {
		{
			Address: "靖县",
			EndYear: 1986,
		},
		{
			Address:   "靖州苗族侗族自治县",
			StartYear: 1987,
			EndYear:   1996,
		},
	},
	433031: {
		{
			Address: "通道侗族自治县",
			EndYear: 1996,
		},
	},
	433100: {
		{
			Address: "湘西土家族苗族自治州",
		},
	},
	433101: {
		{
			Address:   "吉首市",
			StartYear: 1982,
		},
	},
	433102: {
		{
			Address:   "大庸市",
			StartYear: 1985,
			EndYear:   1987,
		},
	},
	433121: {
		{
			Address: "吉首县",
			EndYear: 1981,
		},
	},
	433122: {
		{
			Address: "泸溪县",
		},
	},
	433123: {
		{
			Address: "凤凰县",
		},
	},
	433124: {
		{
			Address: "花垣县",
		},
	},
	433125: {
		{
			Address: "保靖县",
		},
	},
	433126: {
		{
			Address: "古丈县",
		},
	},
	433127: {
		{
			Address: "永顺县",
		},
	},
	433128: {
		{
			Address: "大庸县",
			EndYear: 1984,
		},
	},
	433129: {
		{
			Address: "桑植县",
			EndYear: 1987,
		},
	},
	433130: {
		{
			Address: "龙山县",
		},
	},
	439001: {
		{
			Address:   "醴陵市",
			StartYear: 1986,
			EndYear:   1994,
		},
	},
	439002: {
		{
			Address:   "湘乡市",
			StartYear: 1986,
			EndYear:   1994,
		},
	},
	439003: {
		{
			Address:   "耒阳市",
			StartYear: 1986,
			EndYear:   1994,
		},
	},
	439004: {
		{
			Address:   "汨罗市",
			StartYear: 1987,
			EndYear:   1994,
		},
	},
	439005: {
		{
			Address:   "津市市",
			StartYear: 1988,
			EndYear:   1994,
		},
	},
	439006: {
		{
			Address:   "韶山市",
			StartYear: 1990,
			EndYear:   1994,
		},
	},
	439007: {
		{
			Address:   "临湘市",
			StartYear: 1992,
			EndYear:   1994,
		},
	},
	439008: {
		{
			Address:   "浏阳市",
			StartYear: 1993,
			EndYear:   1994,
		},
	},
	439009: {
		{
			Address:   "资兴市",
			StartYear: 1994,
			EndYear:   1994,
		},
	},
	439010: {
		{
			Address:   "沅江市",
			StartYear: 1994,
			EndYear:   1994,
		},
	},
	439011: {
		{
			Address:   "武冈市",
			StartYear: 1994,
			EndYear:   1994,
		},
	},
	440000: {
		{
			Address: "广东省",
		},
	},
	440100: {
		{
			Address: "广州市",
		},
	},
	440102: {
		{
			Address: "东山区",
			EndYear: 2004,
		},
	},
	440103: {
		{
			Address: "荔湾区",
		},
	},
	440104: {
		{
			Address: "越秀区",
		},
	},
	440105: {
		{
			Address: "海珠区",
		},
	},
	440106: {
		{
			Address:   "芳村区",
			StartYear: 1985,
			EndYear:   1987,
		},
		{
			Address:   "天河区",
			StartYear: 1988,
		},
	},
	440107: {
		{
			Address:   "天河区",
			StartYear: 1985,
			EndYear:   1987,
		},
		{
			Address:   "芳村区",
			StartYear: 1988,
			EndYear:   2004,
		},
	},
	440111: {
		{
			Address: "郊区",
			EndYear: 1986,
		},
		{
			Address:   "白云区",
			StartYear: 1987,
		},
	},
	440112: {
		{
			Address: "黄埔区",
		},
	},
	440113: {
		{
			Address:   "番禺区",
			StartYear: 2000,
		},
	},
	440114: {
		{
			Address:   "花都区",
			StartYear: 2000,
		},
	},
	440115: {
		{
			Address:   "南沙区",
			StartYear: 2005,
		},
	},
	440116: {
		{
			Address:   "萝岗区",
			StartYear: 2005,
			EndYear:   2013,
		},
	},
	440117: {
		{
			Address:   "从化区",
			StartYear: 2014,
		},
	},
	440118: {
		{
			Address:   "增城区",
			StartYear: 2014,
		},
	},
	440121: {
		{
			Address: "花县",
			EndYear: 1992,
		},
	},
	440122: {
		{
			Address: "从化县",
			EndYear: 1993,
		},
	},
	440123: {
		{
			Address: "新丰县",
			EndYear: 1987,
		},
	},
	440124: {
		{
			Address: "龙门县",
			EndYear: 1987,
		},
	},
	440125: {
		{
			Address: "增城县",
			EndYear: 1992,
		},
	},
	440126: {
		{
			Address: "番禺县",
			EndYear: 1991,
		},
	},
	440127: {
		{
			Address:   "清远县",
			StartYear: 1983,
			EndYear:   1987,
		},
	},
	440128: {
		{
			Address:   "佛冈县",
			StartYear: 1983,
			EndYear:   1987,
		},
	},
	440181: {
		{
			Address:   "番禺市",
			StartYear: 1995,
			EndYear:   1999,
		},
	},
	440182: {
		{
			Address:   "花都市",
			StartYear: 1995,
			EndYear:   1999,
		},
	},
	440183: {
		{
			Address:   "增城市",
			StartYear: 1995,
			EndYear:   2013,
		},
	},
	440184: {
		{
			Address:   "从化市",
			StartYear: 1995,
			EndYear:   2013,
		},
	},
	440200: {
		{
			Address: "韶关市",
		},
	},
	440202: {
		{
			Address:   "北江区",
			StartYear: 1984,
			EndYear:   2003,
		},
	},
	440203: {
		{
			Address:   "浈江区",
			StartYear: 1984,
			EndYear:   2002,
		},
		{
			Address:   "武江区",
			StartYear: 2003,
		},
	},
	440204: {
		{
			Address:   "武江区",
			StartYear: 1984,
			EndYear:   2002,
		},
		{
			Address:   "浈江区",
			StartYear: 2003,
		},
	},
	440205: {
		{
			Address:   "曲江区",
			StartYear: 2004,
		},
	},
	440221: {
		{
			Address: "曲江县",
			EndYear: 2003,
		},
	},
	440222: {
		{
			Address:   "始兴县",
			StartYear: 1983,
		},
	},
	440223: {
		{
			Address:   "南雄县",
			StartYear: 1983,
			EndYear:   1995,
		},
	},
	440224: {
		{
			Address:   "仁化县",
			StartYear: 1983,
		},
	},
	440225: {
		{
			Address:   "乐昌县",
			StartYear: 1983,
			EndYear:   1993,
		},
	},
	440226: {
		{
			Address:   "连县",
			StartYear: 1983,
			EndYear:   1987,
		},
	},
	440227: {
		{
			Address:   "阳山县",
			StartYear: 1983,
			EndYear:   1987,
		},
	},
	440228: {
		{
			Address:   "英德县",
			StartYear: 1983,
			EndYear:   1987,
		},
	},
	440229: {
		{
			Address:   "翁源县",
			StartYear: 1983,
		},
	},
	440230: {
		{
			Address:   "连山壮族瑶族自治县",
			StartYear: 1983,
			EndYear:   1987,
		},
	},
	440231: {
		{
			Address:   "连南瑶族自治县",
			StartYear: 1983,
			EndYear:   1987,
		},
	},
	440232: {
		{
			Address:   "乳源瑶族自治县",
			StartYear: 1983,
		},
	},
	440233: {
		{
			Address:   "新丰县",
			StartYear: 1988,
		},
	},
	440281: {
		{
			Address:   "乐昌市",
			StartYear: 1995,
		},
	},
	440282: {
		{
			Address:   "南雄市",
			StartYear: 1996,
		},
	},
	440300: {
		{
			Address: "深圳市",
		},
	},
	440303: {
		{
			Address:   "罗湖区",
			StartYear: 1990,
		},
	},
	440304: {
		{
			Address:   "福田区",
			StartYear: 1990,
		},
	},
	440305: {
		{
			Address:   "南山区",
			StartYear: 1990,
		},
	},
	440306: {
		{
			Address:   "宝安区",
			StartYear: 1992,
		},
	},
	440307: {
		{
			Address:   "龙岗区",
			StartYear: 1992,
		},
	},
	440308: {
		{
			Address:   "盐田区",
			StartYear: 1997,
		},
	},
	440309: {
		{
			Address:   "龙华区",
			StartYear: 2016,
		},
	},
	440310: {
		{
			Address:   "坪山区",
			StartYear: 2016,
		},
	},
	440311: {
		{
			Address:   "光明区",
			StartYear: 2018,
		},
	},
	440321: {
		{
			Address:   "宝安县",
			StartYear: 1982,
			EndYear:   1991,
		},
	},
	440400: {
		{
			Address: "珠海市",
		},
	},
	440402: {
		{
			Address:   "香洲区",
			StartYear: 1984,
		},
	},
	440403: {
		{
			Address:   "斗门区",
			StartYear: 2001,
		},
	},
	440404: {
		{
			Address:   "金湾区",
			StartYear: 2001,
		},
	},
	440421: {
		{
			Address:   "斗门县",
			StartYear: 1983,
			EndYear:   2000,
		},
	},
	440500: {
		{
			Address: "汕头市",
		},
	},
	440502: {
		{
			Address: "同平区",
			EndYear: 1990,
		},
		{
			Address:   "龙湖区",
			StartYear: 1991,
			EndYear:   1993,
		},
	},
	440503: {
		{
			Address: "安平区",
			EndYear: 1990,
		},
		{
			Address:   "金园区",
			StartYear: 1991,
			EndYear:   1993,
		},
	},
	440504: {
		{
			Address: "公园区",
			EndYear: 1990,
		},
		{
			Address:   "升平区",
			StartYear: 1991,
			EndYear:   1993,
		},
	},
	440505: {
		{
			Address: "金沙区",
			EndYear: 1990,
		},
	},
	440506: {
		{
			Address:   "达豪区",
			StartYear: 1984,
			EndYear:   2002,
		},
	},
	440507: {
		{
			Address:   "龙湖区",
			StartYear: 1994,
		},
	},
	440508: {
		{
			Address:   "金园区",
			StartYear: 1994,
			EndYear:   2002,
		},
	},
	440509: {
		{
			Address:   "升平区",
			StartYear: 1994,
			EndYear:   2002,
		},
	},
	440510: {
		{
			Address:   "河浦区",
			StartYear: 1994,
			EndYear:   2002,
		},
	},
	440511: {
		{
			Address: "郊区",
			EndYear: 1990,
		},
		{
			Address:   "金平区",
			StartYear: 2003,
		},
	},
	440512: {
		{
			Address:   "濠江区",
			StartYear: 2003,
		},
	},
	440513: {
		{
			Address:   "潮阳区",
			StartYear: 2003,
		},
	},
	440514: {
		{
			Address:   "潮南区",
			StartYear: 2003,
		},
	},
	440515: {
		{
			Address:   "澄海区",
			StartYear: 2003,
		},
	},
	440521: {
		{
			Address:   "澄海县",
			StartYear: 1983,
			EndYear:   1993,
		},
	},
	440522: {
		{
			Address:   "饶平县",
			StartYear: 1983,
			EndYear:   1990,
		},
	},
	440523: {
		{
			Address:   "南澳县",
			StartYear: 1983,
		},
	},
	440524: {
		{
			Address:   "潮阳县",
			StartYear: 1983,
			EndYear:   1992,
		},
	},
	440525: {
		{
			Address:   "揭阳县",
			StartYear: 1983,
			EndYear:   1990,
		},
	},
	440526: {
		{
			Address:   "揭西县",
			StartYear: 1983,
			EndYear:   1990,
		},
	},
	440527: {
		{
			Address:   "普宁县",
			StartYear: 1983,
			EndYear:   1990,
		},
	},
	440528: {
		{
			Address:   "惠来县",
			StartYear: 1983,
			EndYear:   1990,
		},
	},
	440582: {
		{
			Address:   "潮阳市",
			StartYear: 1995,
			EndYear:   2002,
		},
	},
	440583: {
		{
			Address:   "澄海市",
			StartYear: 1995,
			EndYear:   2002,
		},
	},
	440600: {
		{
			Address: "佛山市",
		},
	},
	440602: {
		{
			Address:   "汾江区",
			StartYear: 1984,
			EndYear:   1986,
		},
		{
			Address:   "城区",
			StartYear: 1987,
			EndYear:   2001,
		},
	},
	440603: {
		{
			Address:   "石湾区",
			StartYear: 1984,
			EndYear:   2001,
		},
	},
	440604: {
		{
			Address:   "禅城区",
			StartYear: 2002,
		},
	},
	440605: {
		{
			Address:   "南海区",
			StartYear: 2002,
		},
	},
	440606: {
		{
			Address:   "顺德区",
			StartYear: 2002,
		},
	},
	440607: {
		{
			Address:   "三水区",
			StartYear: 2002,
		},
	},
	440608: {
		{
			Address:   "高明区",
			StartYear: 2002,
		},
	},
	440620: {
		{
			Address:   "中山市",
			StartYear: 1983,
			EndYear:   1983,
		},
	},
	440621: {
		{
			Address:   "三水县",
			StartYear: 1983,
			EndYear:   1992,
		},
	},
	440622: {
		{
			Address:   "南海县",
			StartYear: 1983,
			EndYear:   1991,
		},
	},
	440623: {
		{
			Address:   "顺德县",
			StartYear: 1983,
			EndYear:   1991,
		},
	},
	440624: {
		{
			Address:   "高明县",
			StartYear: 1983,
			EndYear:   1993,
		},
	},
	440681: {
		{
			Address:   "顺德市",
			StartYear: 1995,
			EndYear:   2001,
		},
	},
	440682: {
		{
			Address:   "南海市",
			StartYear: 1995,
			EndYear:   2001,
		},
	},
	440683: {
		{
			Address:   "三水市",
			StartYear: 1995,
			EndYear:   2001,
		},
	},
	440684: {
		{
			Address:   "高明市",
			StartYear: 1995,
			EndYear:   2001,
		},
	},
	440700: {
		{
			Address: "江门市",
		},
	},
	440702: {
		{
			Address:   "城区",
			StartYear: 1984,
			EndYear:   1993,
		},
	},
	440703: {
		{
			Address:   "蓬江区",
			StartYear: 1994,
		},
	},
	440704: {
		{
			Address:   "江海区",
			StartYear: 1994,
		},
	},
	440705: {
		{
			Address:   "新会区",
			StartYear: 2002,
		},
	},
	440711: {
		{
			Address:   "郊区",
			StartYear: 1984,
			EndYear:   1993,
		},
	},
	440721: {
		{
			Address:   "新会县",
			StartYear: 1983,
			EndYear:   1991,
		},
	},
	440722: {
		{
			Address:   "台山县",
			StartYear: 1983,
			EndYear:   1991,
		},
	},
	440723: {
		{
			Address:   "恩平县",
			StartYear: 1983,
			EndYear:   1993,
		},
	},
	440724: {
		{
			Address:   "开平县",
			StartYear: 1983,
			EndYear:   1992,
		},
	},
	440725: {
		{
			Address:   "鹤山县",
			StartYear: 1983,
			EndYear:   1992,
		},
	},
	440726: {
		{
			Address:   "阳江县",
			StartYear: 1983,
			EndYear:   1987,
		},
	},
	440727: {
		{
			Address:   "阳春县",
			StartYear: 1983,
			EndYear:   1987,
		},
	},
	440781: {
		{
			Address:   "台山市",
			StartYear: 1995,
		},
	},
	440782: {
		{
			Address:   "新会市",
			StartYear: 1995,
			EndYear:   2001,
		},
	},
	440783: {
		{
			Address:   "开平市",
			StartYear: 1995,
		},
	},
	440784: {
		{
			Address:   "鹤山市",
			StartYear: 1995,
		},
	},
	440785: {
		{
			Address:   "恩平市",
			StartYear: 1995,
		},
	},
	440800: {
		{
			Address: "湛江市",
		},
	},
	440802: {
		{
			Address: "赤坎区",
		},
	},
	440803: {
		{
			Address: "霞山区",
		},
	},
	440804: {
		{
			Address:   "坡头区",
			StartYear: 1984,
		},
	},
	440811: {
		{
			Address: "郊区",
			EndYear: 1993,
		},
		{
			Address:   "麻章区",
			StartYear: 1994,
		},
	},
	440821: {
		{
			Address:   "吴川县",
			StartYear: 1983,
			EndYear:   1993,
		},
	},
	440822: {
		{
			Address:   "廉江县",
			StartYear: 1983,
			EndYear:   1992,
		},
	},
	440823: {
		{
			Address:   "遂溪县",
			StartYear: 1983,
		},
	},
	440824: {
		{
			Address:   "海康县",
			StartYear: 1983,
			EndYear:   1993,
		},
	},
	440825: {
		{
			Address:   "徐闻县",
			StartYear: 1983,
		},
	},
	440881: {
		{
			Address:   "廉江市",
			StartYear: 1995,
		},
	},
	440882: {
		{
			Address:   "雷州市",
			StartYear: 1995,
		},
	},
	440883: {
		{
			Address:   "吴川市",
			StartYear: 1995,
		},
	},
	440900: {
		{
			Address: "茂名市",
		},
	},
	440902: {
		{
			Address:   "茂南区",
			StartYear: 1984,
		},
	},
	440903: {
		{
			Address:   "茂港区",
			StartYear: 2001,
			EndYear:   2013,
		},
	},
	440904: {
		{
			Address:   "电白区",
			StartYear: 2014,
		},
	},
	440921: {
		{
			Address:   "信宜县",
			StartYear: 1983,
			EndYear:   1994,
		},
	},
	440922: {
		{
			Address:   "高州县",
			StartYear: 1983,
			EndYear:   1992,
		},
	},
	440923: {
		{
			Address:   "电白县",
			StartYear: 1983,
			EndYear:   2013,
		},
	},
	440924: {
		{
			Address:   "化州县",
			StartYear: 1983,
			EndYear:   1993,
		},
	},
	440981: {
		{
			Address:   "高州市",
			StartYear: 1995,
		},
	},
	440982: {
		{
			Address:   "化州市",
			StartYear: 1995,
		},
	},
	440983: {
		{
			Address:   "信宜市",
			StartYear: 1995,
		},
	},
	441000: {
		{
			Address: "海口市",
			EndYear: 1987,
		},
	},
	441001: {
		{
			Address:   "潮州市",
			StartYear: 1983,
			EndYear:   1983,
		},
	},
	441002: {
		{
			Address: "新华区",
			EndYear: 1982,
		},
	},
	441003: {
		{
			Address: "立新区",
			EndYear: 1982,
		},
	},
	441004: {
		{
			Address: "东方红区",
			EndYear: 1982,
		},
	},
	441005: {
		{
			Address: "秀英区",
			EndYear: 1982,
		},
	},
	441100: {
		{
			Address:   "三亚市",
			StartYear: 1987,
			EndYear:   1987,
		},
	},
	441200: {
		{
			Address:   "肇庆市",
			StartYear: 1988,
		},
	},
	441202: {
		{
			Address:   "端州区",
			StartYear: 1988,
		},
	},
	441203: {
		{
			Address:   "鼎湖区",
			StartYear: 1988,
		},
	},
	441204: {
		{
			Address:   "高要区",
			StartYear: 2015,
		},
	},
	441221: {
		{
			Address:   "高要县",
			StartYear: 1988,
			EndYear:   1992,
		},
	},
	441222: {
		{
			Address:   "四会县",
			StartYear: 1988,
			EndYear:   1992,
		},
	},
	441223: {
		{
			Address:   "广宁县",
			StartYear: 1988,
		},
	},
	441224: {
		{
			Address:   "怀集县",
			StartYear: 1988,
		},
	},
	441225: {
		{
			Address:   "封开县",
			StartYear: 1988,
		},
	},
	441226: {
		{
			Address:   "德庆县",
			StartYear: 1988,
		},
	},
	441227: {
		{
			Address:   "云浮县",
			StartYear: 1988,
			EndYear:   1991,
		},
	},
	441228: {
		{
			Address:   "新兴县",
			StartYear: 1988,
			EndYear:   1993,
		},
	},
	441229: {
		{
			Address:   "郁南县",
			StartYear: 1988,
			EndYear:   1993,
		},
	},
	441230: {
		{
			Address:   "罗定县",
			StartYear: 1988,
			EndYear:   1992,
		},
	},
	441283: {
		{
			Address:   "高要市",
			StartYear: 1995,
			EndYear:   2014,
		},
	},
	441284: {
		{
			Address:   "四会市",
			StartYear: 1995,
		},
	},
	441300: {
		{
			Address:   "惠州市",
			StartYear: 1988,
		},
	},
	441302: {
		{
			Address:   "惠城区",
			StartYear: 1988,
		},
	},
	441303: {
		{
			Address:   "惠阳区",
			StartYear: 2003,
		},
	},
	441321: {
		{
			Address:   "惠阳县",
			StartYear: 1988,
			EndYear:   1993,
		},
	},
	441322: {
		{
			Address:   "博罗县",
			StartYear: 1988,
		},
	},
	441323: {
		{
			Address:   "惠东县",
			StartYear: 1988,
		},
	},
	441324: {
		{
			Address:   "龙门县",
			StartYear: 1988,
		},
	},
	441381: {
		{
			Address:   "惠阳市",
			StartYear: 1995,
			EndYear:   2002,
		},
	},
	441400: {
		{
			Address:   "梅州市",
			StartYear: 1988,
		},
	},
	441402: {
		{
			Address:   "梅江区",
			StartYear: 1988,
		},
	},
	441403: {
		{
			Address:   "梅县区",
			StartYear: 2013,
		},
	},
	441421: {
		{
			Address:   "梅县",
			StartYear: 1988,
			EndYear:   2012,
		},
	},
	441422: {
		{
			Address:   "大埔县",
			StartYear: 1988,
		},
	},
	441423: {
		{
			Address:   "丰顺县",
			StartYear: 1988,
		},
	},
	441424: {
		{
			Address:   "五华县",
			StartYear: 1988,
		},
	},
	441425: {
		{
			Address:   "兴宁县",
			StartYear: 1988,
			EndYear:   1993,
		},
	},
	441426: {
		{
			Address:   "平远县",
			StartYear: 1988,
		},
	},
	441427: {
		{
			Address:   "蕉岭县",
			StartYear: 1988,
		},
	},
	441481: {
		{
			Address:   "兴宁市",
			StartYear: 1995,
		},
	},
	441500: {
		{
			Address:   "汕尾市",
			StartYear: 1988,
		},
	},
	441502: {
		{
			Address:   "城区",
			StartYear: 1988,
		},
	},
	441521: {
		{
			Address:   "海丰县",
			StartYear: 1988,
		},
	},
	441522: {
		{
			Address:   "陆丰县",
			StartYear: 1988,
			EndYear:   1994,
		},
	},
	441523: {
		{
			Address:   "陆河县",
			StartYear: 1988,
		},
	},
	441581: {
		{
			Address:   "陆丰市",
			StartYear: 1995,
		},
	},
	441600: {
		{
			Address:   "河源市",
			StartYear: 1988,
		},
	},
	441602: {
		{
			Address:   "源城区",
			StartYear: 1988,
		},
	},
	441611: {
		{
			Address:   "郊区",
			StartYear: 1988,
			EndYear:   1992,
		},
	},
	441621: {
		{
			Address:   "紫金县",
			StartYear: 1988,
		},
	},
	441622: {
		{
			Address:   "龙川县",
			StartYear: 1988,
		},
	},
	441623: {
		{
			Address:   "连平县",
			StartYear: 1988,
		},
	},
	441624: {
		{
			Address:   "和平县",
			StartYear: 1988,
		},
	},
	441625: {
		{
			Address:   "东源县",
			StartYear: 1993,
		},
	},
	441700: {
		{
			Address:   "阳江市",
			StartYear: 1988,
		},
	},
	441702: {
		{
			Address:   "江城区",
			StartYear: 1988,
		},
	},
	441703: {
		{
			Address:   "阳东区",
			StartYear: 1988,
			EndYear:   1990,
		},
	},
	441704: {
		{
			Address:   "阳东区",
			StartYear: 2014,
		},
	},
	441721: {
		{
			Address:   "阳西县",
			StartYear: 1988,
		},
	},
	441722: {
		{
			Address:   "阳春县",
			StartYear: 1988,
			EndYear:   1993,
		},
	},
	441723: {
		{
			Address:   "阳东县",
			StartYear: 1991,
			EndYear:   2013,
		},
	},
	441781: {
		{
			Address:   "阳春市",
			StartYear: 1995,
		},
	},
	441800: {
		{
			Address:   "清远市",
			StartYear: 1988,
		},
	},
	441802: {
		{
			Address:   "清城区",
			StartYear: 1988,
		},
	},
	441803: {
		{
			Address:   "清新区",
			StartYear: 2012,
		},
	},
	441811: {
		{
			Address:   "清郊区",
			StartYear: 1988,
			EndYear:   1991,
		},
	},
	441821: {
		{
			Address:   "佛冈县",
			StartYear: 1988,
		},
	},
	441822: {
		{
			Address:   "英德县",
			StartYear: 1988,
			EndYear:   1993,
		},
	},
	441823: {
		{
			Address:   "阳山县",
			StartYear: 1988,
		},
	},
	441824: {
		{
			Address:   "连县",
			StartYear: 1988,
			EndYear:   1993,
		},
	},
	441825: {
		{
			Address:   "连山壮族瑶族自治县",
			StartYear: 1988,
		},
	},
	441826: {
		{
			Address:   "连南瑶族自治县",
			StartYear: 1988,
		},
	},
	441827: {
		{
			Address:   "清新县",
			StartYear: 1992,
			EndYear:   2011,
		},
	},
	441881: {
		{
			Address:   "英德市",
			StartYear: 1995,
		},
	},
	441882: {
		{
			Address:   "连州市",
			StartYear: 1995,
		},
	},
	441900: {
		{
			Address:   "东莞市",
			StartYear: 1988,
		},
	},
	442000: {
		{
			Address:   "中山市",
			StartYear: 1988,
		},
	},
	442100: {
		{
			Address: "海南行政区",
			EndYear: 1987,
		},
	},
	442101: {
		{
			Address:   "海口市",
			StartYear: 1983,
			EndYear:   1985,
		},
	},
	442102: {
		{
			Address:   "通什市",
			StartYear: 1987,
			EndYear:   1987,
		},
	},
	442121: {
		{
			Address: "琼山县",
			EndYear: 1987,
		},
	},
	442122: {
		{
			Address: "文昌县",
			EndYear: 1987,
		},
	},
	442123: {
		{
			Address: "琼海县",
			EndYear: 1987,
		},
	},
	442124: {
		{
			Address: "万宁县",
			EndYear: 1987,
		},
	},
	442125: {
		{
			Address: "定安县",
			EndYear: 1987,
		},
	},
	442126: {
		{
			Address: "屯昌县",
			EndYear: 1987,
		},
	},
	442127: {
		{
			Address: "澄迈县",
			EndYear: 1987,
		},
	},
	442128: {
		{
			Address: "临高县",
			EndYear: 1987,
		},
	},
	442129: {
		{
			Address: "儋县",
			EndYear: 1987,
		},
	},
	442130: {
		{
			Address:   "东方黎族自治县",
			StartYear: 1987,
			EndYear:   1987,
		},
	},
	442131: {
		{
			Address:   "乐东黎族自治县",
			StartYear: 1987,
			EndYear:   1987,
		},
	},
	442132: {
		{
			Address:   "琼中黎族苗族自治县",
			StartYear: 1987,
			EndYear:   1987,
		},
	},
	442133: {
		{
			Address:   "保亭黎族苗族自治县",
			StartYear: 1987,
			EndYear:   1987,
		},
	},
	442134: {
		{
			Address:   "陵水黎族自治县",
			StartYear: 1987,
			EndYear:   1987,
		},
	},
	442135: {
		{
			Address:   "白沙黎族自治县",
			StartYear: 1987,
			EndYear:   1987,
		},
	},
	442136: {
		{
			Address:   "昌江黎族自治县",
			StartYear: 1987,
			EndYear:   1987,
		},
	},
	442200: {
		{
			Address: "海南黎族苗族自治州",
			EndYear: 1986,
		},
	},
	442201: {
		{
			Address:   "三亚市",
			StartYear: 1984,
			EndYear:   1986,
		},
	},
	442202: {
		{
			Address:   "通什市",
			StartYear: 1986,
			EndYear:   1986,
		},
	},
	442221: {
		{
			Address: "崖县",
			EndYear: 1983,
		},
	},
	442222: {
		{
			Address: "东方县",
			EndYear: 1986,
		},
	},
	442223: {
		{
			Address: "乐东县",
			EndYear: 1986,
		},
	},
	442224: {
		{
			Address: "琼中县",
			EndYear: 1986,
		},
	},
	442225: {
		{
			Address: "保亭县",
			EndYear: 1986,
		},
	},
	442226: {
		{
			Address: "陵水县",
			EndYear: 1986,
		},
	},
	442227: {
		{
			Address: "白沙县",
			EndYear: 1986,
		},
	},
	442228: {
		{
			Address: "昌江县",
			EndYear: 1986,
		},
	},
	442300: {
		{
			Address: "汕头地区",
			EndYear: 1982,
		},
	},
	442301: {
		{
			Address: "潮州市",
			EndYear: 1981,
		},
	},
	442302: {
		{
			Address:   "潮州市",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	442321: {
		{
			Address: "潮安县",
			EndYear: 1982,
		},
	},
	442322: {
		{
			Address: "澄海县",
			EndYear: 1982,
		},
	},
	442323: {
		{
			Address: "饶平县",
			EndYear: 1982,
		},
	},
	442324: {
		{
			Address: "南澳县",
			EndYear: 1982,
		},
	},
	442325: {
		{
			Address: "潮阳县",
			EndYear: 1982,
		},
	},
	442326: {
		{
			Address: "揭阳县",
			EndYear: 1982,
		},
	},
	442327: {
		{
			Address: "揭西县",
			EndYear: 1982,
		},
	},
	442328: {
		{
			Address: "普宁县",
			EndYear: 1982,
		},
	},
	442329: {
		{
			Address: "惠来县",
			EndYear: 1982,
		},
	},
	442330: {
		{
			Address: "陆丰县",
			EndYear: 1982,
		},
	},
	442331: {
		{
			Address: "海丰县",
			EndYear: 1982,
		},
	},
	442400: {
		{
			Address: "梅县地区",
			EndYear: 1987,
		},
	},
	442401: {
		{
			Address: "梅州市",
			EndYear: 1982,
		},
		{
			Address:   "梅县市",
			StartYear: 1983,
			EndYear:   1987,
		},
	},
	442421: {
		{
			Address: "梅县",
			EndYear: 1982,
		},
	},
	442422: {
		{
			Address: "大埔县",
			EndYear: 1987,
		},
	},
	442423: {
		{
			Address: "丰顺县",
			EndYear: 1987,
		},
	},
	442424: {
		{
			Address: "五华县",
			EndYear: 1987,
		},
	},
	442425: {
		{
			Address: "兴宁县",
			EndYear: 1987,
		},
	},
	442426: {
		{
			Address: "平远县",
			EndYear: 1987,
		},
	},
	442427: {
		{
			Address: "蕉岭县",
			EndYear: 1987,
		},
	},
	442500: {
		{
			Address: "惠阳地区",
			EndYear: 1987,
		},
	},
	442501: {
		{
			Address: "惠州市",
			EndYear: 1987,
		},
	},
	442502: {
		{
			Address:   "东莞市",
			StartYear: 1985,
			EndYear:   1987,
		},
	},
	442521: {
		{
			Address: "惠阳县",
			EndYear: 1987,
		},
	},
	442522: {
		{
			Address: "紫金县",
			EndYear: 1987,
		},
	},
	442523: {
		{
			Address: "和平县",
			EndYear: 1987,
		},
	},
	442524: {
		{
			Address: "连平县",
			EndYear: 1987,
		},
	},
	442525: {
		{
			Address: "河源县",
			EndYear: 1987,
		},
	},
	442526: {
		{
			Address: "博罗县",
			EndYear: 1987,
		},
	},
	442527: {
		{
			Address: "东莞县",
			EndYear: 1984,
		},
	},
	442528: {
		{
			Address: "惠东县",
			EndYear: 1987,
		},
	},
	442529: {
		{
			Address: "龙川县",
			EndYear: 1987,
		},
	},
	442530: {
		{
			Address:   "陆丰县",
			StartYear: 1983,
			EndYear:   1987,
		},
	},
	442531: {
		{
			Address:   "海丰县",
			StartYear: 1983,
			EndYear:   1987,
		},
	},
	442600: {
		{
			Address: "韶关地区",
			EndYear: 1982,
		},
	},
	442621: {
		{
			Address: "始兴县",
			EndYear: 1982,
		},
		{
			Address:   "三水县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	442622: {
		{
			Address: "南雄县",
			EndYear: 1982,
		},
		{
			Address:   "南海县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	442623: {
		{
			Address: "仁化县",
			EndYear: 1982,
		},
		{
			Address:   "顺德县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	442624: {
		{
			Address: "乐昌县",
			EndYear: 1982,
		},
		{
			Address:   "中山县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	442625: {
		{
			Address: "连县",
			EndYear: 1982,
		},
		{
			Address:   "斗门县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	442626: {
		{
			Address: "阳山县",
			EndYear: 1982,
		},
		{
			Address:   "新会县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	442627: {
		{
			Address: "英德县",
			EndYear: 1982,
		},
		{
			Address:   "台山县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	442628: {
		{
			Address: "清远县",
			EndYear: 1982,
		},
		{
			Address:   "恩平县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	442629: {
		{
			Address: "佛冈县",
			EndYear: 1982,
		},
		{
			Address:   "开平县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	442630: {
		{
			Address: "翁源县",
			EndYear: 1982,
		},
	},
	442631: {
		{
			Address: "连山壮族瑶族自治县",
			EndYear: 1982,
		},
		{
			Address:   "鹤山县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	442632: {
		{
			Address: "连南瑶族自治县",
			EndYear: 1982,
		},
		{
			Address:   "高明县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	442633: {
		{
			Address: "乳源瑶族自治县",
			EndYear: 1982,
		},
	},
	442700: {
		{
			Address: "佛山地区",
			EndYear: 1982,
		},
	},
	442721: {
		{
			Address: "三水县",
			EndYear: 1981,
		},
	},
	442722: {
		{
			Address: "南海县",
			EndYear: 1981,
		},
	},
	442723: {
		{
			Address: "顺德县",
			EndYear: 1981,
		},
	},
	442724: {
		{
			Address: "中山县",
			EndYear: 1981,
		},
	},
	442725: {
		{
			Address: "斗门县",
			EndYear: 1981,
		},
	},
	442726: {
		{
			Address: "新会县",
			EndYear: 1981,
		},
	},
	442727: {
		{
			Address: "台山县",
			EndYear: 1981,
		},
	},
	442728: {
		{
			Address: "恩平县",
			EndYear: 1981,
		},
	},
	442729: {
		{
			Address: "开平县",
			EndYear: 1981,
		},
	},
	442730: {
		{
			Address: "高鹤县",
			EndYear: 1980,
		},
	},
	442731: {
		{
			Address:   "鹤山县",
			StartYear: 1981,
			EndYear:   1981,
		},
	},
	442732: {
		{
			Address:   "高明县",
			StartYear: 1981,
			EndYear:   1981,
		},
	},
	442800: {
		{
			Address: "肇庆地区",
			EndYear: 1987,
		},
	},
	442801: {
		{
			Address: "肇庆市",
			EndYear: 1987,
		},
	},
	442821: {
		{
			Address: "高要县",
			EndYear: 1987,
		},
	},
	442822: {
		{
			Address: "四会县",
			EndYear: 1987,
		},
	},
	442823: {
		{
			Address: "广宁县",
			EndYear: 1987,
		},
	},
	442824: {
		{
			Address: "怀集县",
			EndYear: 1987,
		},
	},
	442825: {
		{
			Address: "封开县",
			EndYear: 1987,
		},
	},
	442826: {
		{
			Address: "德庆县",
			EndYear: 1987,
		},
	},
	442827: {
		{
			Address: "云浮县",
			EndYear: 1987,
		},
	},
	442828: {
		{
			Address: "新兴县",
			EndYear: 1987,
		},
	},
	442829: {
		{
			Address: "郁南县",
			EndYear: 1987,
		},
	},
	442830: {
		{
			Address: "罗定县",
			EndYear: 1987,
		},
	},
	442900: {
		{
			Address: "湛江地区",
			EndYear: 1982,
		},
	},
	442921: {
		{
			Address: "阳江县",
			EndYear: 1982,
		},
	},
	442922: {
		{
			Address: "阳春县",
			EndYear: 1982,
		},
	},
	442923: {
		{
			Address: "信宜县",
			EndYear: 1982,
		},
	},
	442924: {
		{
			Address: "高州县",
			EndYear: 1982,
		},
	},
	442925: {
		{
			Address: "电白县",
			EndYear: 1982,
		},
	},
	442926: {
		{
			Address: "吴川县",
			EndYear: 1982,
		},
	},
	442927: {
		{
			Address: "化州县",
			EndYear: 1982,
		},
	},
	442928: {
		{
			Address: "廉江县",
			EndYear: 1982,
		},
	},
	442929: {
		{
			Address: "遂溪县",
			EndYear: 1982,
		},
	},
	442930: {
		{
			Address: "海康县",
			EndYear: 1982,
		},
	},
	442931: {
		{
			Address: "徐闻县",
			EndYear: 1982,
		},
	},
	445100: {
		{
			Address:   "潮州市",
			StartYear: 1991,
		},
	},
	445102: {
		{
			Address:   "湘桥区",
			StartYear: 1991,
		},
	},
	445103: {
		{
			Address:   "潮安区",
			StartYear: 2013,
		},
	},
	445121: {
		{
			Address:   "潮安县",
			StartYear: 1991,
			EndYear:   2012,
		},
	},
	445122: {
		{
			Address:   "饶平县",
			StartYear: 1991,
		},
	},
	445200: {
		{
			Address:   "揭阳市",
			StartYear: 1991,
		},
	},
	445202: {
		{
			Address:   "榕城区",
			StartYear: 1991,
		},
	},
	445203: {
		{
			Address:   "揭东区",
			StartYear: 2012,
		},
	},
	445221: {
		{
			Address:   "揭东县",
			StartYear: 1991,
			EndYear:   2011,
		},
	},
	445222: {
		{
			Address:   "揭西县",
			StartYear: 1991,
		},
	},
	445223: {
		{
			Address:   "普宁县",
			StartYear: 1991,
			EndYear:   1992,
		},
	},
	445224: {
		{
			Address:   "惠来县",
			StartYear: 1991,
		},
	},
	445281: {
		{
			Address:   "普宁市",
			StartYear: 1995,
		},
	},
	445300: {
		{
			Address:   "云浮市",
			StartYear: 1994,
		},
	},
	445302: {
		{
			Address:   "云城区",
			StartYear: 1994,
		},
	},
	445303: {
		{
			Address:   "云安区",
			StartYear: 2014,
		},
	},
	445321: {
		{
			Address:   "新兴县",
			StartYear: 1994,
		},
	},
	445322: {
		{
			Address:   "郁南县",
			StartYear: 1994,
		},
	},
	445323: {
		{
			Address:   "云安县",
			StartYear: 1996,
			EndYear:   2013,
		},
	},
	445381: {
		{
			Address:   "罗定市",
			StartYear: 1995,
		},
	},
	449001: {
		{
			Address:   "潮州市",
			StartYear: 1984,
			EndYear:   1990,
		},
		{
			Address:   "顺德市",
			StartYear: 1992,
			EndYear:   1994,
		},
	},
	449002: {
		{
			Address:   "中山市",
			StartYear: 1984,
			EndYear:   1987,
		},
		{
			Address:   "台山市",
			StartYear: 1992,
			EndYear:   1994,
		},
	},
	449003: {
		{
			Address:   "海口市",
			StartYear: 1986,
			EndYear:   1986,
		},
		{
			Address:   "番禺市",
			StartYear: 1992,
			EndYear:   1994,
		},
	},
	449004: {
		{
			Address:   "南海市",
			StartYear: 1992,
			EndYear:   1994,
		},
	},
	449005: {
		{
			Address:   "云浮市",
			StartYear: 1992,
			EndYear:   1993,
		},
	},
	449006: {
		{
			Address:   "新会市",
			StartYear: 1992,
			EndYear:   1994,
		},
	},
	449007: {
		{
			Address:   "开平市",
			StartYear: 1993,
			EndYear:   1994,
		},
	},
	449008: {
		{
			Address:   "三水市",
			StartYear: 1993,
			EndYear:   1994,
		},
	},
	449009: {
		{
			Address:   "普宁市",
			StartYear: 1993,
			EndYear:   1994,
		},
	},
	449010: {
		{
			Address:   "罗定市",
			StartYear: 1993,
			EndYear:   1994,
		},
	},
	449011: {
		{
			Address:   "潮阳市",
			StartYear: 1993,
			EndYear:   1994,
		},
	},
	449012: {
		{
			Address:   "高州市",
			StartYear: 1993,
			EndYear:   1994,
		},
	},
	449013: {
		{
			Address:   "花都市",
			StartYear: 1993,
			EndYear:   1994,
		},
	},
	449014: {
		{
			Address:   "高要市",
			StartYear: 1993,
			EndYear:   1994,
		},
	},
	449015: {
		{
			Address:   "鹤山市",
			StartYear: 1993,
			EndYear:   1994,
		},
	},
	449016: {
		{
			Address:   "四会市",
			StartYear: 1993,
			EndYear:   1994,
		},
	},
	449017: {
		{
			Address:   "增城市",
			StartYear: 1993,
			EndYear:   1994,
		},
	},
	449018: {
		{
			Address:   "廉江市",
			StartYear: 1993,
			EndYear:   1994,
		},
	},
	449019: {
		{
			Address:   "英德市",
			StartYear: 1994,
			EndYear:   1994,
		},
	},
	449020: {
		{
			Address:   "恩平市",
			StartYear: 1994,
			EndYear:   1994,
		},
	},
	449021: {
		{
			Address:   "从化市",
			StartYear: 1994,
			EndYear:   1994,
		},
	},
	449022: {
		{
			Address:   "澄海市",
			StartYear: 1994,
			EndYear:   1994,
		},
	},
	449023: {
		{
			Address:   "高明市",
			StartYear: 1994,
			EndYear:   1994,
		},
	},
	449024: {
		{
			Address:   "连州市",
			StartYear: 1994,
			EndYear:   1994,
		},
	},
	449025: {
		{
			Address:   "雷州市",
			StartYear: 1994,
			EndYear:   1994,
		},
	},
	449026: {
		{
			Address:   "乐昌市",
			StartYear: 1994,
			EndYear:   1994,
		},
	},
	449027: {
		{
			Address:   "阳春市",
			StartYear: 1994,
			EndYear:   1994,
		},
	},
	449028: {
		{
			Address:   "惠阳市",
			StartYear: 1994,
			EndYear:   1994,
		},
	},
	449029: {
		{
			Address:   "吴川市",
			StartYear: 1994,
			EndYear:   1994,
		},
	},
	449030: {
		{
			Address:   "兴宁市",
			StartYear: 1994,
			EndYear:   1994,
		},
	},
	449031: {
		{
			Address:   "化州市",
			StartYear: 1994,
			EndYear:   1994,
		},
	},
	450000: {
		{
			Address: "广西壮族自治区",
		},
	},
	450100: {
		{
			Address: "南宁市",
		},
	},
	450102: {
		{
			Address: "兴宁区",
		},
	},
	450103: {
		{
			Address: "新城区",
			EndYear: 2003,
		},
		{
			Address:   "青秀区",
			StartYear: 2004,
		},
	},
	450104: {
		{
			Address: "城北区",
			EndYear: 2003,
		},
	},
	450105: {
		{
			Address: "江南区",
		},
	},
	450106: {
		{
			Address: "永新区",
			EndYear: 2003,
		},
	},
	450107: {
		{
			Address:   "市郊区",
			StartYear: 1984,
			EndYear:   2000,
		},
		{
			Address:   "西乡塘区",
			StartYear: 2004,
		},
	},
	450108: {
		{
			Address:   "良庆区",
			StartYear: 2004,
		},
	},
	450109: {
		{
			Address:   "邕宁区",
			StartYear: 2004,
		},
	},
	450110: {
		{
			Address:   "武鸣区",
			StartYear: 2015,
		},
	},
	450121: {
		{
			Address:   "邕宁县",
			StartYear: 1983,
			EndYear:   2003,
		},
	},
	450122: {
		{
			Address:   "武鸣县",
			StartYear: 1983,
			EndYear:   2014,
		},
	},
	450123: {
		{
			Address:   "隆安县",
			StartYear: 2002,
		},
	},
	450124: {
		{
			Address:   "马山县",
			StartYear: 2002,
		},
	},
	450125: {
		{
			Address:   "上林县",
			StartYear: 2002,
		},
	},
	450126: {
		{
			Address:   "宾阳县",
			StartYear: 2002,
		},
	},
	450127: {
		{
			Address:   "横县",
			StartYear: 2002,
		},
	},
	450200: {
		{
			Address: "柳州市",
		},
	},
	450202: {
		{
			Address: "城中区",
		},
	},
	450203: {
		{
			Address: "鱼峰区",
		},
	},
	450204: {
		{
			Address: "柳南区",
		},
	},
	450205: {
		{
			Address: "柳北区",
		},
	},
	450206: {
		{
			Address:   "市郊区",
			StartYear: 1984,
			EndYear:   2001,
		},
		{
			Address:   "柳江区",
			StartYear: 2016,
		},
	},
	450221: {
		{
			Address:   "柳江县",
			StartYear: 1983,
			EndYear:   2015,
		},
	},
	450222: {
		{
			Address:   "柳城县",
			StartYear: 1983,
		},
	},
	450223: {
		{
			Address:   "鹿寨县",
			StartYear: 2002,
		},
	},
	450224: {
		{
			Address:   "融安县",
			StartYear: 2002,
		},
	},
	450225: {
		{
			Address:   "融水苗族自治县",
			StartYear: 2002,
		},
	},
	450226: {
		{
			Address:   "三江侗族自治县",
			StartYear: 2002,
		},
	},
	450300: {
		{
			Address: "桂林市",
		},
	},
	450302: {
		{
			Address: "秀峰区",
		},
	},
	450303: {
		{
			Address: "叠彩区",
		},
	},
	450304: {
		{
			Address: "象山区",
		},
	},
	450305: {
		{
			Address: "七星区",
		},
	},
	450306: {
		{
			Address:   "市郊区",
			StartYear: 1984,
			EndYear:   1995,
		},
	},
	450311: {
		{
			Address:   "雁山区",
			StartYear: 1996,
		},
	},
	450312: {
		{
			Address:   "临桂区",
			StartYear: 2013,
		},
	},
	450321: {
		{
			Address:   "阳朔县",
			StartYear: 1981,
		},
	},
	450322: {
		{
			Address:   "临桂县",
			StartYear: 1983,
			EndYear:   2012,
		},
	},
	450323: {
		{
			Address:   "灵川县",
			StartYear: 1998,
		},
	},
	450324: {
		{
			Address:   "全州县",
			StartYear: 1998,
		},
	},
	450325: {
		{
			Address:   "兴安县",
			StartYear: 1998,
		},
	},
	450326: {
		{
			Address:   "永福县",
			StartYear: 1998,
		},
	},
	450327: {
		{
			Address:   "灌阳县",
			StartYear: 1998,
		},
	},
	450328: {
		{
			Address:   "龙胜各族自治县",
			StartYear: 1998,
		},
	},
	450329: {
		{
			Address:   "资源县",
			StartYear: 1998,
		},
	},
	450330: {
		{
			Address:   "平乐县",
			StartYear: 1998,
		},
	},
	450331: {
		{
			Address:   "荔浦县",
			StartYear: 1998,
			EndYear:   2017,
		},
	},
	450332: {
		{
			Address:   "恭城瑶族自治县",
			StartYear: 1998,
		},
	},
	450381: {
		{
			Address:   "荔浦市",
			StartYear: 2018,
		},
	},
	450400: {
		{
			Address: "梧州市",
		},
	},
	450402: {
		{
			Address: "白云区",
			EndYear: 1983,
		},
	},
	450403: {
		{
			Address: "万秀区",
		},
	},
	450404: {
		{
			Address: "蝶山区",
			EndYear: 2012,
		},
	},
	450405: {
		{
			Address: "鸳江区",
			EndYear: 1983,
		},
		{
			Address:   "市郊区",
			StartYear: 1984,
			EndYear:   2002,
		},
		{
			Address:   "长洲区",
			StartYear: 2003,
		},
	},
	450406: {
		{
			Address:   "龙圩区",
			StartYear: 2013,
		},
	},
	450421: {
		{
			Address:   "苍梧县",
			StartYear: 1983,
		},
	},
	450422: {
		{
			Address:   "藤县",
			StartYear: 1997,
		},
	},
	450423: {
		{
			Address:   "蒙山县",
			StartYear: 1997,
		},
	},
	450481: {
		{
			Address:   "岑溪市",
			StartYear: 1997,
		},
	},
	450500: {
		{
			Address:   "北海市",
			StartYear: 1983,
		},
	},
	450502: {
		{
			Address:   "海城区",
			StartYear: 1984,
		},
	},
	450503: {
		{
			Address:   "市郊区",
			StartYear: 1984,
			EndYear:   1993,
		},
		{
			Address:   "银海区",
			StartYear: 1994,
		},
	},
	450512: {
		{
			Address:   "铁山港区",
			StartYear: 1994,
		},
	},
	450521: {
		{
			Address:   "合浦县",
			StartYear: 1987,
		},
	},
	450600: {
		{
			Address:   "防城港市",
			StartYear: 1993,
		},
	},
	450602: {
		{
			Address:   "港口区",
			StartYear: 1993,
		},
	},
	450603: {
		{
			Address:   "防城区",
			StartYear: 1993,
		},
	},
	450621: {
		{
			Address:   "上思县",
			StartYear: 1993,
		},
	},
	450681: {
		{
			Address:   "东兴市",
			StartYear: 1996,
		},
	},
	450700: {
		{
			Address:   "钦州市",
			StartYear: 1994,
		},
	},
	450702: {
		{
			Address:   "钦南区",
			StartYear: 1994,
		},
	},
	450703: {
		{
			Address:   "钦北区",
			StartYear: 1994,
		},
	},
	450721: {
		{
			Address:   "灵山县",
			StartYear: 1994,
		},
	},
	450722: {
		{
			Address:   "浦北县",
			StartYear: 1994,
		},
	},
	450800: {
		{
			Address:   "贵港市",
			StartYear: 1995,
		},
	},
	450802: {
		{
			Address:   "港北区",
			StartYear: 1995,
		},
	},
	450803: {
		{
			Address:   "港南区",
			StartYear: 1995,
		},
	},
	450804: {
		{
			Address:   "覃塘区",
			StartYear: 2003,
		},
	},
	450821: {
		{
			Address:   "平南县",
			StartYear: 1995,
		},
	},
	450881: {
		{
			Address:   "桂平市",
			StartYear: 1995,
		},
	},
	450900: {
		{
			Address:   "玉林市",
			StartYear: 1997,
		},
	},
	450902: {
		{
			Address:   "玉州区",
			StartYear: 1997,
		},
	},
	450903: {
		{
			Address:   "福绵区",
			StartYear: 2013,
		},
	},
	450921: {
		{
			Address:   "容县",
			StartYear: 1997,
		},
	},
	450922: {
		{
			Address:   "陆川县",
			StartYear: 1997,
		},
	},
	450923: {
		{
			Address:   "博白县",
			StartYear: 1997,
		},
	},
	450924: {
		{
			Address:   "兴业县",
			StartYear: 1997,
		},
	},
	450981: {
		{
			Address:   "北流市",
			StartYear: 1997,
		},
	},
	451000: {
		{
			Address:   "百色市",
			StartYear: 2002,
		},
	},
	451002: {
		{
			Address:   "右江区",
			StartYear: 2002,
		},
	},
	451003: {
		{
			Address:   "田阳区",
			StartYear: 2019,
		},
	},
	451021: {
		{
			Address:   "田阳县",
			StartYear: 2002,
			EndYear:   2018,
		},
	},
	451022: {
		{
			Address:   "田东县",
			StartYear: 2002,
		},
	},
	451023: {
		{
			Address:   "平果县",
			StartYear: 2002,
			EndYear:   2018,
		},
	},
	451024: {
		{
			Address:   "德保县",
			StartYear: 2002,
		},
	},
	451025: {
		{
			Address:   "靖西县",
			StartYear: 2002,
			EndYear:   2014,
		},
	},
	451026: {
		{
			Address:   "那坡县",
			StartYear: 2002,
		},
	},
	451027: {
		{
			Address:   "凌云县",
			StartYear: 2002,
		},
	},
	451028: {
		{
			Address:   "乐业县",
			StartYear: 2002,
		},
	},
	451029: {
		{
			Address:   "田林县",
			StartYear: 2002,
		},
	},
	451030: {
		{
			Address:   "西林县",
			StartYear: 2002,
		},
	},
	451031: {
		{
			Address:   "隆林各族自治县",
			StartYear: 2002,
		},
	},
	451081: {
		{
			Address:   "靖西市",
			StartYear: 2015,
		},
	},
	451082: {
		{
			Address:   "平果市",
			StartYear: 2019,
		},
	},
	451100: {
		{
			Address:   "贺州市",
			StartYear: 2002,
		},
	},
	451102: {
		{
			Address:   "八步区",
			StartYear: 2002,
		},
	},
	451103: {
		{
			Address:   "平桂区",
			StartYear: 2016,
		},
	},
	451121: {
		{
			Address:   "昭平县",
			StartYear: 2002,
		},
	},
	451122: {
		{
			Address:   "钟山县",
			StartYear: 2002,
		},
	},
	451123: {
		{
			Address:   "富川瑶族自治县",
			StartYear: 2002,
		},
	},
	451200: {
		{
			Address:   "河池市",
			StartYear: 2002,
		},
	},
	451202: {
		{
			Address:   "金城江区",
			StartYear: 2002,
		},
	},
	451203: {
		{
			Address:   "宜州区",
			StartYear: 2016,
		},
	},
	451221: {
		{
			Address:   "南丹县",
			StartYear: 2002,
		},
	},
	451222: {
		{
			Address:   "天峨县",
			StartYear: 2002,
		},
	},
	451223: {
		{
			Address:   "凤山县",
			StartYear: 2002,
		},
	},
	451224: {
		{
			Address:   "东兰县",
			StartYear: 2002,
		},
	},
	451225: {
		{
			Address:   "罗城仫佬族自治县",
			StartYear: 2002,
		},
	},
	451226: {
		{
			Address:   "环江毛南族自治县",
			StartYear: 2002,
		},
	},
	451227: {
		{
			Address:   "巴马瑶族自治县",
			StartYear: 2002,
		},
	},
	451228: {
		{
			Address:   "都安瑶族自治县",
			StartYear: 2002,
		},
	},
	451229: {
		{
			Address:   "大化瑶族自治县",
			StartYear: 2002,
		},
	},
	451281: {
		{
			Address:   "宜州市",
			StartYear: 2002,
			EndYear:   2015,
		},
	},
	451300: {
		{
			Address:   "来宾市",
			StartYear: 2002,
		},
	},
	451302: {
		{
			Address:   "兴宾区",
			StartYear: 2002,
		},
	},
	451321: {
		{
			Address:   "忻城县",
			StartYear: 2002,
		},
	},
	451322: {
		{
			Address:   "象州县",
			StartYear: 2002,
		},
	},
	451323: {
		{
			Address:   "武宣县",
			StartYear: 2002,
		},
	},
	451324: {
		{
			Address:   "金秀瑶族自治县",
			StartYear: 2002,
		},
	},
	451381: {
		{
			Address:   "合山市",
			StartYear: 2002,
		},
	},
	451400: {
		{
			Address:   "崇左市",
			StartYear: 2002,
		},
	},
	451402: {
		{
			Address:   "江州区",
			StartYear: 2002,
		},
	},
	451421: {
		{
			Address:   "扶绥县",
			StartYear: 2002,
		},
	},
	451422: {
		{
			Address:   "宁明县",
			StartYear: 2002,
		},
	},
	451423: {
		{
			Address:   "龙州县",
			StartYear: 2002,
		},
	},
	451424: {
		{
			Address:   "大新县",
			StartYear: 2002,
		},
	},
	451425: {
		{
			Address:   "天等县",
			StartYear: 2002,
		},
	},
	451481: {
		{
			Address:   "凭祥市",
			StartYear: 2002,
		},
	},
	452100: {
		{
			Address: "南宁地区",
			EndYear: 2001,
		},
	},
	452101: {
		{
			Address: "凭祥市",
			EndYear: 2001,
		},
	},
	452121: {
		{
			Address: "邕宁县",
			EndYear: 1982,
		},
	},
	452122: {
		{
			Address: "横县",
			EndYear: 2001,
		},
	},
	452123: {
		{
			Address: "宾阳县",
			EndYear: 2001,
		},
	},
	452124: {
		{
			Address: "上林县",
			EndYear: 2001,
		},
	},
	452125: {
		{
			Address: "武鸣县",
			EndYear: 1982,
		},
	},
	452126: {
		{
			Address: "隆安县",
			EndYear: 2001,
		},
	},
	452127: {
		{
			Address: "马山县",
			EndYear: 2001,
		},
	},
	452128: {
		{
			Address: "扶绥县",
			EndYear: 2001,
		},
	},
	452129: {
		{
			Address: "崇左县",
			EndYear: 2001,
		},
	},
	452130: {
		{
			Address: "大新县",
			EndYear: 2001,
		},
	},
	452131: {
		{
			Address: "天等县",
			EndYear: 2001,
		},
	},
	452132: {
		{
			Address: "宁明县",
			EndYear: 2001,
		},
	},
	452133: {
		{
			Address: "龙州县",
			EndYear: 2001,
		},
	},
	452200: {
		{
			Address: "柳州地区",
			EndYear: 2001,
		},
	},
	452201: {
		{
			Address:   "合山市",
			StartYear: 1981,
			EndYear:   2001,
		},
	},
	452221: {
		{
			Address: "柳江县",
			EndYear: 1982,
		},
	},
	452222: {
		{
			Address: "柳城县",
			EndYear: 1982,
		},
	},
	452223: {
		{
			Address: "鹿寨县",
			EndYear: 2001,
		},
	},
	452224: {
		{
			Address: "象州县",
			EndYear: 2001,
		},
	},
	452225: {
		{
			Address: "武宣县",
			EndYear: 2001,
		},
	},
	452226: {
		{
			Address: "来宾县",
			EndYear: 2001,
		},
	},
	452227: {
		{
			Address: "融安县",
			EndYear: 2001,
		},
	},
	452228: {
		{
			Address: "三江侗族自治县",
			EndYear: 2001,
		},
	},
	452229: {
		{
			Address: "融水苗族自治县",
			EndYear: 2001,
		},
	},
	452230: {
		{
			Address: "金秀瑶族自治县",
			EndYear: 2001,
		},
	},
	452231: {
		{
			Address: "忻城县",
			EndYear: 2001,
		},
	},
	452300: {
		{
			Address: "桂林地区",
			EndYear: 1997,
		},
	},
	452321: {
		{
			Address: "临桂县",
			EndYear: 1982,
		},
	},
	452322: {
		{
			Address: "灵川县",
			EndYear: 1997,
		},
	},
	452323: {
		{
			Address: "全州县",
			EndYear: 1997,
		},
	},
	452324: {
		{
			Address: "兴安县",
			EndYear: 1997,
		},
	},
	452325: {
		{
			Address: "永福县",
			EndYear: 1997,
		},
	},
	452326: {
		{
			Address: "阳朔县",
			EndYear: 1980,
		},
	},
	452327: {
		{
			Address: "灌阳县",
			EndYear: 1997,
		},
	},
	452328: {
		{
			Address: "龙胜各族自治县",
			EndYear: 1997,
		},
	},
	452329: {
		{
			Address: "资源县",
			EndYear: 1997,
		},
	},
	452330: {
		{
			Address: "平乐县",
			EndYear: 1997,
		},
	},
	452331: {
		{
			Address: "荔浦县",
			EndYear: 1997,
		},
	},
	452332: {
		{
			Address: "恭城县",
			EndYear: 1989,
		},
		{
			Address:   "恭城瑶族自治县",
			StartYear: 1990,
			EndYear:   1997,
		},
	},
	452400: {
		{
			Address: "梧州地区",
			EndYear: 1996,
		},
		{
			Address:   "贺州地区",
			StartYear: 1997,
			EndYear:   2001,
		},
	},
	452401: {
		{
			Address:   "岑溪市",
			StartYear: 1995,
			EndYear:   1996,
		},
	},
	452402: {
		{
			Address:   "贺州市",
			StartYear: 1997,
			EndYear:   2001,
		},
	},
	452421: {
		{
			Address: "岑溪县",
			EndYear: 1994,
		},
	},
	452422: {
		{
			Address: "苍梧县",
			EndYear: 1982,
		},
	},
	452423: {
		{
			Address: "藤县",
			EndYear: 1996,
		},
	},
	452424: {
		{
			Address: "昭平县",
			EndYear: 2001,
		},
	},
	452425: {
		{
			Address: "蒙山县",
			EndYear: 1996,
		},
	},
	452426: {
		{
			Address: "贺县",
			EndYear: 1996,
		},
	},
	452427: {
		{
			Address: "钟山县",
			EndYear: 2001,
		},
	},
	452428: {
		{
			Address: "富川县",
			EndYear: 1982,
		},
		{
			Address:   "富川瑶族自治县",
			StartYear: 1983,
			EndYear:   2001,
		},
	},
	452500: {
		{
			Address: "玉林地区",
			EndYear: 1996,
		},
	},
	452501: {
		{
			Address:   "玉林市",
			StartYear: 1983,
			EndYear:   1996,
		},
	},
	452502: {
		{
			Address:   "贵港市",
			StartYear: 1988,
			EndYear:   1994,
		},
	},
	452503: {
		{
			Address:   "桂平市",
			StartYear: 1994,
			EndYear:   1994,
		},
	},
	452504: {
		{
			Address:   "北流市",
			StartYear: 1994,
			EndYear:   1996,
		},
	},
	452521: {
		{
			Address: "玉林县",
			EndYear: 1982,
		},
	},
	452522: {
		{
			Address: "贵县",
			EndYear: 1987,
		},
	},
	452523: {
		{
			Address: "桂平县",
			EndYear: 1993,
		},
	},
	452524: {
		{
			Address: "平南县",
			EndYear: 1994,
		},
	},
	452525: {
		{
			Address: "容县",
			EndYear: 1996,
		},
	},
	452526: {
		{
			Address: "北流县",
			EndYear: 1993,
		},
	},
	452527: {
		{
			Address: "陆川县",
			EndYear: 1996,
		},
	},
	452528: {
		{
			Address: "博白县",
			EndYear: 1996,
		},
	},
	452600: {
		{
			Address: "百色地区",
			EndYear: 2001,
		},
	},
	452601: {
		{
			Address:   "百色市",
			StartYear: 1983,
			EndYear:   2001,
		},
	},
	452621: {
		{
			Address: "百色县",
			EndYear: 1982,
		},
	},
	452622: {
		{
			Address: "田阳县",
			EndYear: 2001,
		},
	},
	452623: {
		{
			Address: "田东县",
			EndYear: 2001,
		},
	},
	452624: {
		{
			Address: "平果县",
			EndYear: 2001,
		},
	},
	452625: {
		{
			Address: "德保县",
			EndYear: 2001,
		},
	},
	452626: {
		{
			Address: "靖西县",
			EndYear: 2001,
		},
	},
	452627: {
		{
			Address: "那坡县",
			EndYear: 2001,
		},
	},
	452628: {
		{
			Address: "凌云县",
			EndYear: 2001,
		},
	},
	452629: {
		{
			Address: "乐业县",
			EndYear: 2001,
		},
	},
	452630: {
		{
			Address: "田林县",
			EndYear: 2001,
		},
	},
	452631: {
		{
			Address: "隆林各族自治县",
			EndYear: 2001,
		},
	},
	452632: {
		{
			Address: "西林县",
			EndYear: 2001,
		},
	},
	452700: {
		{
			Address: "河池地区",
			EndYear: 2001,
		},
	},
	452701: {
		{
			Address:   "河池市",
			StartYear: 1983,
			EndYear:   2001,
		},
	},
	452702: {
		{
			Address:   "宜州市",
			StartYear: 1993,
			EndYear:   2001,
		},
	},
	452721: {
		{
			Address: "河池县",
			EndYear: 1982,
		},
	},
	452722: {
		{
			Address: "宜山县",
			EndYear: 1992,
		},
	},
	452723: {
		{
			Address: "罗城县",
			EndYear: 1982,
		},
		{
			Address:   "罗城仫佬族自治县",
			StartYear: 1983,
			EndYear:   2001,
		},
	},
	452724: {
		{
			Address: "环江县",
			EndYear: 1985,
		},
		{
			Address:   "环江毛南族自治县",
			StartYear: 1986,
			EndYear:   2001,
		},
	},
	452725: {
		{
			Address: "南丹县",
			EndYear: 2001,
		},
	},
	452726: {
		{
			Address: "天峨县",
			EndYear: 2001,
		},
	},
	452727: {
		{
			Address: "凤山县",
			EndYear: 2001,
		},
	},
	452728: {
		{
			Address: "东兰县",
			EndYear: 2001,
		},
	},
	452729: {
		{
			Address: "巴马瑶族自治县",
			EndYear: 2001,
		},
	},
	452730: {
		{
			Address: "都安瑶族自治县",
			EndYear: 2001,
		},
	},
	452731: {
		{
			Address:   "大化瑶族自治县",
			StartYear: 1987,
			EndYear:   2001,
		},
	},
	452800: {
		{
			Address: "钦州地区",
			EndYear: 1993,
		},
	},
	452801: {
		{
			Address: "北海市",
			EndYear: 1982,
		},
	},
	452802: {
		{
			Address:   "钦州市",
			StartYear: 1983,
			EndYear:   1993,
		},
	},
	452821: {
		{
			Address: "上思县",
			EndYear: 1992,
		},
	},
	452822: {
		{
			Address: "防城各族自治县",
			EndYear: 1992,
		},
	},
	452823: {
		{
			Address: "钦州县",
			EndYear: 1982,
		},
	},
	452824: {
		{
			Address: "灵山县",
			EndYear: 1993,
		},
	},
	452825: {
		{
			Address: "合浦县",
			EndYear: 1986,
		},
	},
	452826: {
		{
			Address: "浦北县",
			EndYear: 1993,
		},
	},
	460000: {
		{
			Address:   "海南省",
			StartYear: 1988,
		},
	},
	460001: {
		{
			Address:   "通什市",
			StartYear: 1988,
			EndYear:   2000,
		},
		{
			Address:   "五指山市",
			StartYear: 2001,
			EndYear:   2002,
		},
	},
	460002: {
		{
			Address:   "琼海市",
			StartYear: 1992,
			EndYear:   2002,
		},
	},
	460003: {
		{
			Address:   "儋州市",
			StartYear: 1993,
			EndYear:   2002,
		},
	},
	460004: {
		{
			Address:   "琼山市",
			StartYear: 1994,
			EndYear:   2001,
		},
	},
	460005: {
		{
			Address:   "文昌市",
			StartYear: 1995,
			EndYear:   2002,
		},
	},
	460006: {
		{
			Address:   "万宁市",
			StartYear: 1996,
			EndYear:   2002,
		},
	},
	460007: {
		{
			Address:   "东方市",
			StartYear: 1997,
			EndYear:   2002,
		},
	},
	460021: {
		{
			Address:   "琼山县",
			StartYear: 1988,
			EndYear:   1993,
		},
	},
	460022: {
		{
			Address:   "文昌县",
			StartYear: 1988,
			EndYear:   1994,
		},
	},
	460023: {
		{
			Address:   "琼海县",
			StartYear: 1988,
			EndYear:   1991,
		},
	},
	460024: {
		{
			Address:   "万宁县",
			StartYear: 1988,
			EndYear:   1995,
		},
	},
	460025: {
		{
			Address:   "定安县",
			StartYear: 1988,
			EndYear:   2002,
		},
	},
	460026: {
		{
			Address:   "屯昌县",
			StartYear: 1988,
			EndYear:   2002,
		},
	},
	460027: {
		{
			Address:   "澄迈县",
			StartYear: 1988,
			EndYear:   2002,
		},
	},
	460028: {
		{
			Address:   "临高县",
			StartYear: 1988,
			EndYear:   2002,
		},
	},
	460029: {
		{
			Address:   "儋县",
			StartYear: 1988,
			EndYear:   1992,
		},
	},
	460030: {
		{
			Address:   "白沙黎族自治县",
			StartYear: 1988,
			EndYear:   2002,
		},
	},
	460031: {
		{
			Address:   "昌江黎族自治县",
			StartYear: 1988,
			EndYear:   2002,
		},
	},
	460032: {
		{
			Address:   "东方黎族自治县",
			StartYear: 1988,
			EndYear:   1996,
		},
	},
	460033: {
		{
			Address:   "乐东黎族自治县",
			StartYear: 1988,
			EndYear:   2002,
		},
	},
	460034: {
		{
			Address:   "陵水黎族自治县",
			StartYear: 1988,
			EndYear:   2002,
		},
	},
	460035: {
		{
			Address:   "保亭黎族苗族自治县",
			StartYear: 1988,
			EndYear:   2002,
		},
	},
	460036: {
		{
			Address:   "琼中黎族苗族自治县",
			StartYear: 1988,
			EndYear:   2002,
		},
	},
	460100: {
		{
			Address:   "海口市",
			StartYear: 1988,
		},
	},
	460102: {
		{
			Address:   "振东区",
			StartYear: 1990,
			EndYear:   2001,
		},
	},
	460103: {
		{
			Address:   "新华区",
			StartYear: 1990,
			EndYear:   2001,
		},
	},
	460104: {
		{
			Address:   "秀英区",
			StartYear: 1990,
			EndYear:   2001,
		},
	},
	460105: {
		{
			Address:   "秀英区",
			StartYear: 2002,
		},
	},
	460106: {
		{
			Address:   "龙华区",
			StartYear: 2002,
		},
	},
	460107: {
		{
			Address:   "琼山区",
			StartYear: 2002,
		},
	},
	460108: {
		{
			Address:   "美兰区",
			StartYear: 2002,
		},
	},
	460200: {
		{
			Address:   "三亚市",
			StartYear: 1988,
		},
	},
	460202: {
		{
			Address:   "海棠区",
			StartYear: 2014,
		},
	},
	460203: {
		{
			Address:   "吉阳区",
			StartYear: 2014,
		},
	},
	460204: {
		{
			Address:   "天涯区",
			StartYear: 2014,
		},
	},
	460205: {
		{
			Address:   "崖州区",
			StartYear: 2014,
		},
	},
	460300: {
		{
			Address:   "三沙市",
			StartYear: 2012,
		},
	},
	460400: {
		{
			Address:   "儋州市",
			StartYear: 2015,
		},
	},
	469001: {
		{
			Address:   "五指山市",
			StartYear: 2003,
		},
	},
	469002: {
		{
			Address:   "琼海市",
			StartYear: 2003,
		},
	},
	469003: {
		{
			Address:   "儋州市",
			StartYear: 2003,
			EndYear:   2014,
		},
	},
	469005: {
		{
			Address:   "文昌市",
			StartYear: 2003,
		},
	},
	469006: {
		{
			Address:   "万宁市",
			StartYear: 2003,
		},
	},
	469007: {
		{
			Address:   "东方市",
			StartYear: 2003,
		},
	},
	469021: {
		{
			Address:   "定安县",
			StartYear: 2003,
		},
	},
	469022: {
		{
			Address:   "屯昌县",
			StartYear: 2003,
		},
	},
	469023: {
		{
			Address:   "澄迈县",
			StartYear: 2003,
		},
	},
	469024: {
		{
			Address:   "临高县",
			StartYear: 2003,
		},
	},
	469025: {
		{
			Address:   "白沙黎族自治县",
			StartYear: 2003,
		},
	},
	469026: {
		{
			Address:   "昌江黎族自治县",
			StartYear: 2003,
		},
	},
	469027: {
		{
			Address:   "乐东黎族自治县",
			StartYear: 2003,
		},
	},
	469028: {
		{
			Address:   "陵水黎族自治县",
			StartYear: 2003,
		},
	},
	469029: {
		{
			Address:   "保亭黎族苗族自治县",
			StartYear: 2003,
		},
	},
	469030: {
		{
			Address:   "琼中黎族苗族自治县",
			StartYear: 2003,
		},
	},
	500000: {
		{
			Address:   "重庆市",
			StartYear: 1997,
		},
	},
	500101: {
		{
			Address:   "万县区",
			StartYear: 1997,
			EndYear:   1997,
		},
		{
			Address:   "万州区",
			StartYear: 1998,
		},
	},
	500102: {
		{
			Address:   "涪陵区",
			StartYear: 1997,
		},
	},
	500103: {
		{
			Address:   "渝中区",
			StartYear: 1997,
		},
	},
	500104: {
		{
			Address:   "大渡口区",
			StartYear: 1997,
		},
	},
	500105: {
		{
			Address:   "江北区",
			StartYear: 1997,
		},
	},
	500106: {
		{
			Address:   "沙坪坝区",
			StartYear: 1997,
		},
	},
	500107: {
		{
			Address:   "九龙坡区",
			StartYear: 1997,
		},
	},
	500108: {
		{
			Address:   "南岸区",
			StartYear: 1997,
		},
	},
	500109: {
		{
			Address:   "北碚区",
			StartYear: 1997,
		},
	},
	500110: {
		{
			Address:   "万盛区",
			StartYear: 1997,
			EndYear:   2010,
		},
		{
			Address:   "綦江区",
			StartYear: 2011,
		},
	},
	500111: {
		{
			Address:   "双桥区",
			StartYear: 1997,
			EndYear:   2010,
		},
		{
			Address:   "大足区",
			StartYear: 2011,
		},
	},
	500112: {
		{
			Address:   "渝北区",
			StartYear: 1997,
		},
	},
	500113: {
		{
			Address:   "巴南区",
			StartYear: 1997,
		},
	},
	500114: {
		{
			Address:   "黔江区",
			StartYear: 2000,
		},
	},
	500115: {
		{
			Address:   "长寿区",
			StartYear: 2001,
		},
	},
	500116: {
		{
			Address:   "江津区",
			StartYear: 2006,
		},
	},
	500117: {
		{
			Address:   "合川区",
			StartYear: 2006,
		},
	},
	500118: {
		{
			Address:   "永川区",
			StartYear: 2006,
		},
	},
	500119: {
		{
			Address:   "南川区",
			StartYear: 2006,
		},
	},
	500120: {
		{
			Address:   "璧山区",
			StartYear: 2014,
		},
	},
	500151: {
		{
			Address:   "铜梁区",
			StartYear: 2014,
		},
	},
	500152: {
		{
			Address:   "潼南区",
			StartYear: 2015,
		},
	},
	500153: {
		{
			Address:   "荣昌区",
			StartYear: 2015,
		},
	},
	500154: {
		{
			Address:   "开州区",
			StartYear: 2016,
		},
	},
	500155: {
		{
			Address:   "梁平区",
			StartYear: 2016,
		},
	},
	500156: {
		{
			Address:   "武隆区",
			StartYear: 2016,
		},
	},
	500221: {
		{
			Address:   "长寿县",
			StartYear: 1997,
			EndYear:   2000,
		},
	},
	500222: {
		{
			Address:   "綦江县",
			StartYear: 1997,
			EndYear:   2010,
		},
	},
	500223: {
		{
			Address:   "潼南县",
			StartYear: 1997,
			EndYear:   2014,
		},
	},
	500224: {
		{
			Address:   "铜梁县",
			StartYear: 1997,
			EndYear:   2013,
		},
	},
	500225: {
		{
			Address:   "大足县",
			StartYear: 1997,
			EndYear:   2010,
		},
	},
	500226: {
		{
			Address:   "荣昌县",
			StartYear: 1997,
			EndYear:   2014,
		},
	},
	500227: {
		{
			Address:   "璧山县",
			StartYear: 1997,
			EndYear:   2013,
		},
	},
	500228: {
		{
			Address:   "梁平县",
			StartYear: 1997,
			EndYear:   2015,
		},
	},
	500229: {
		{
			Address:   "城口县",
			StartYear: 1997,
		},
	},
	500230: {
		{
			Address:   "丰都县",
			StartYear: 1997,
		},
	},
	500231: {
		{
			Address:   "垫江县",
			StartYear: 1997,
		},
	},
	500232: {
		{
			Address:   "武隆县",
			StartYear: 1997,
			EndYear:   2015,
		},
	},
	500233: {
		{
			Address:   "忠县",
			StartYear: 1997,
		},
	},
	500234: {
		{
			Address:   "开县",
			StartYear: 1997,
			EndYear:   2015,
		},
	},
	500235: {
		{
			Address:   "云阳县",
			StartYear: 1997,
		},
	},
	500236: {
		{
			Address:   "奉节县",
			StartYear: 1997,
		},
	},
	500237: {
		{
			Address:   "巫山县",
			StartYear: 1997,
		},
	},
	500238: {
		{
			Address:   "巫溪县",
			StartYear: 1997,
		},
	},
	500239: {
		{
			Address:   "黔江土家族苗族自治县",
			StartYear: 1997,
			EndYear:   1999,
		},
	},
	500240: {
		{
			Address:   "石柱土家族自治县",
			StartYear: 1997,
		},
	},
	500241: {
		{
			Address:   "秀山土家族苗族自治县",
			StartYear: 1997,
		},
	},
	500242: {
		{
			Address:   "酉阳土家族苗族自治县",
			StartYear: 1997,
		},
	},
	500243: {
		{
			Address:   "彭水苗族土家族自治县",
			StartYear: 1997,
		},
	},
	500381: {
		{
			Address:   "江津市",
			StartYear: 1997,
			EndYear:   2005,
		},
	},
	500382: {
		{
			Address:   "合川市",
			StartYear: 1997,
			EndYear:   2005,
		},
	},
	500383: {
		{
			Address:   "永川市",
			StartYear: 1997,
			EndYear:   2005,
		},
	},
	500384: {
		{
			Address:   "南川市",
			StartYear: 1997,
			EndYear:   2005,
		},
	},
	510000: {
		{
			Address: "四川省",
		},
	},
	510100: {
		{
			Address: "成都市",
		},
	},
	510102: {
		{
			Address: "东城区",
			EndYear: 1989,
		},
	},
	510103: {
		{
			Address: "西城区",
			EndYear: 1989,
		},
	},
	510104: {
		{
			Address:   "锦江区",
			StartYear: 1990,
		},
	},
	510105: {
		{
			Address:   "青羊区",
			StartYear: 1990,
		},
	},
	510106: {
		{
			Address:   "金牛区",
			StartYear: 1990,
		},
	},
	510107: {
		{
			Address:   "武侯区",
			StartYear: 1990,
		},
	},
	510108: {
		{
			Address:   "成华区",
			StartYear: 1990,
		},
	},
	510111: {
		{
			Address: "金牛区",
			EndYear: 1989,
		},
	},
	510112: {
		{
			Address: "龙泉驿区",
		},
	},
	510113: {
		{
			Address: "青白江区",
		},
	},
	510114: {
		{
			Address:   "新都区",
			StartYear: 2001,
		},
	},
	510115: {
		{
			Address:   "温江区",
			StartYear: 2002,
		},
	},
	510116: {
		{
			Address:   "双流区",
			StartYear: 2015,
		},
	},
	510117: {
		{
			Address:   "郫都区",
			StartYear: 2016,
		},
	},
	510118: {
		{
			Address:   "新津区",
			StartYear: 2020,
		},
	},
	510121: {
		{
			Address: "金堂县",
		},
	},
	510122: {
		{
			Address: "双流县",
			EndYear: 2014,
		},
	},
	510123: {
		{
			Address:   "温江县",
			StartYear: 1983,
			EndYear:   2001,
		},
	},
	510124: {
		{
			Address:   "郫县",
			StartYear: 1983,
			EndYear:   2015,
		},
	},
	510125: {
		{
			Address:   "新都县",
			StartYear: 1983,
			EndYear:   2000,
		},
	},
	510126: {
		{
			Address:   "彭县",
			StartYear: 1983,
			EndYear:   1992,
		},
	},
	510127: {
		{
			Address:   "灌县",
			StartYear: 1983,
			EndYear:   1987,
		},
	},
	510128: {
		{
			Address:   "崇庆县",
			StartYear: 1983,
			EndYear:   1993,
		},
	},
	510129: {
		{
			Address:   "大邑县",
			StartYear: 1983,
		},
	},
	510130: {
		{
			Address:   "邛崃县",
			StartYear: 1983,
			EndYear:   1993,
		},
	},
	510131: {
		{
			Address:   "蒲江县",
			StartYear: 1983,
		},
	},
	510132: {
		{
			Address:   "新津县",
			StartYear: 1983,
			EndYear:   2019,
		},
	},
	510181: {
		{
			Address:   "都江堰市",
			StartYear: 1995,
		},
	},
	510182: {
		{
			Address:   "彭州市",
			StartYear: 1995,
		},
	},
	510183: {
		{
			Address:   "邛崃市",
			StartYear: 1995,
		},
	},
	510184: {
		{
			Address:   "崇州市",
			StartYear: 1995,
		},
	},
	510185: {
		{
			Address:   "简阳市",
			StartYear: 2016,
		},
	},
	510200: {
		{
			Address: "重庆市",
			EndYear: 1996,
		},
	},
	510202: {
		{
			Address: "市中区",
			EndYear: 1993,
		},
		{
			Address:   "渝中区",
			StartYear: 1994,
			EndYear:   1996,
		},
	},
	510203: {
		{
			Address: "大渡口区",
			EndYear: 1996,
		},
	},
	510211: {
		{
			Address: "江北区",
			EndYear: 1996,
		},
	},
	510212: {
		{
			Address: "沙坪坝区",
			EndYear: 1996,
		},
	},
	510213: {
		{
			Address: "九龙坡区",
			EndYear: 1996,
		},
	},
	510214: {
		{
			Address: "南岸区",
			EndYear: 1996,
		},
	},
	510215: {
		{
			Address: "北碚区",
			EndYear: 1996,
		},
	},
	510216: {
		{
			Address: "南桐矿区",
			EndYear: 1992,
		},
		{
			Address:   "万盛区",
			StartYear: 1993,
			EndYear:   1996,
		},
	},
	510217: {
		{
			Address: "双桥区",
			EndYear: 1996,
		},
	},
	510218: {
		{
			Address:   "渝北区",
			StartYear: 1994,
			EndYear:   1996,
		},
	},
	510219: {
		{
			Address:   "巴南区",
			StartYear: 1994,
			EndYear:   1996,
		},
	},
	510221: {
		{
			Address: "长寿县",
			EndYear: 1996,
		},
	},
	510222: {
		{
			Address: "巴县",
			EndYear: 1993,
		},
	},
	510223: {
		{
			Address: "綦江县",
			EndYear: 1996,
		},
	},
	510224: {
		{
			Address: "江北县",
			EndYear: 1993,
		},
	},
	510225: {
		{
			Address:   "江津县",
			StartYear: 1983,
			EndYear:   1991,
		},
	},
	510226: {
		{
			Address:   "合川县",
			StartYear: 1983,
			EndYear:   1991,
		},
	},
	510227: {
		{
			Address:   "潼南县",
			StartYear: 1983,
			EndYear:   1996,
		},
	},
	510228: {
		{
			Address:   "铜梁县",
			StartYear: 1983,
			EndYear:   1996,
		},
	},
	510229: {
		{
			Address:   "永川县",
			StartYear: 1983,
			EndYear:   1991,
		},
	},
	510230: {
		{
			Address:   "大足县",
			StartYear: 1983,
			EndYear:   1996,
		},
	},
	510231: {
		{
			Address:   "荣昌县",
			StartYear: 1983,
			EndYear:   1996,
		},
	},
	510232: {
		{
			Address:   "璧山县",
			StartYear: 1983,
			EndYear:   1996,
		},
	},
	510281: {
		{
			Address:   "永川市",
			StartYear: 1995,
			EndYear:   1996,
		},
	},
	510282: {
		{
			Address:   "合川市",
			StartYear: 1995,
			EndYear:   1996,
		},
	},
	510283: {
		{
			Address:   "江津市",
			StartYear: 1995,
			EndYear:   1996,
		},
	},
	510300: {
		{
			Address: "自贡市",
		},
	},
	510302: {
		{
			Address: "自流井区",
		},
	},
	510303: {
		{
			Address: "贡井区",
		},
	},
	510304: {
		{
			Address: "大安区",
		},
	},
	510311: {
		{
			Address: "沿滩区",
		},
	},
	510321: {
		{
			Address: "荣县",
		},
	},
	510322: {
		{
			Address:   "富顺县",
			StartYear: 1983,
		},
	},
	510400: {
		{
			Address: "渡口市",
			EndYear: 1986,
		},
		{
			Address:   "攀枝花市",
			StartYear: 1987,
		},
	},
	510402: {
		{
			Address: "东区",
		},
	},
	510403: {
		{
			Address: "西区",
		},
	},
	510411: {
		{
			Address: "郊区",
			EndYear: 1981,
		},
		{
			Address:   "仁和区",
			StartYear: 1982,
		},
	},
	510421: {
		{
			Address: "米易县",
		},
	},
	510422: {
		{
			Address: "盐边县",
		},
	},
	510500: {
		{
			Address:   "泸州市",
			StartYear: 1983,
		},
	},
	510502: {
		{
			Address:   "市中区",
			StartYear: 1984,
			EndYear:   1994,
		},
		{
			Address:   "江阳区",
			StartYear: 1995,
		},
	},
	510503: {
		{
			Address:   "纳溪区",
			StartYear: 1995,
		},
	},
	510504: {
		{
			Address:   "龙马潭区",
			StartYear: 1995,
		},
	},
	510521: {
		{
			Address:   "泸县",
			StartYear: 1983,
		},
	},
	510522: {
		{
			Address:   "合江县",
			StartYear: 1983,
		},
	},
	510523: {
		{
			Address:   "纳溪县",
			StartYear: 1983,
			EndYear:   1994,
		},
	},
	510524: {
		{
			Address:   "叙永县",
			StartYear: 1985,
		},
	},
	510525: {
		{
			Address:   "古蔺县",
			StartYear: 1985,
		},
	},
	510600: {
		{
			Address:   "德阳市",
			StartYear: 1983,
		},
	},
	510602: {
		{
			Address:   "市中区",
			StartYear: 1984,
			EndYear:   1995,
		},
	},
	510603: {
		{
			Address:   "旌阳区",
			StartYear: 1996,
		},
	},
	510604: {
		{
			Address:   "罗江区",
			StartYear: 2017,
		},
	},
	510621: {
		{
			Address:   "德阳县",
			StartYear: 1983,
			EndYear:   1983,
		},
	},
	510622: {
		{
			Address:   "绵竹县",
			StartYear: 1983,
			EndYear:   1995,
		},
	},
	510623: {
		{
			Address:   "中江县",
			StartYear: 1983,
		},
	},
	510624: {
		{
			Address:   "广汉县",
			StartYear: 1983,
			EndYear:   1987,
		},
	},
	510625: {
		{
			Address:   "什邡县",
			StartYear: 1983,
			EndYear:   1994,
		},
	},
	510626: {
		{
			Address:   "罗江县",
			StartYear: 1996,
			EndYear:   2016,
		},
	},
	510681: {
		{
			Address:   "广汉市",
			StartYear: 1995,
		},
	},
	510682: {
		{
			Address:   "什邡市",
			StartYear: 1995,
		},
	},
	510683: {
		{
			Address:   "绵竹市",
			StartYear: 1996,
		},
	},
	510700: {
		{
			Address:   "绵阳市",
			StartYear: 1985,
		},
	},
	510702: {
		{
			Address:   "市中区",
			StartYear: 1985,
			EndYear:   1991,
		},
	},
	510703: {
		{
			Address:   "涪城区",
			StartYear: 1992,
		},
	},
	510704: {
		{
			Address:   "游仙区",
			StartYear: 1992,
		},
	},
	510705: {
		{
			Address:   "安州区",
			StartYear: 2016,
		},
	},
	510721: {
		{
			Address:   "江油县",
			StartYear: 1985,
			EndYear:   1987,
		},
	},
	510722: {
		{
			Address:   "三台县",
			StartYear: 1985,
		},
	},
	510723: {
		{
			Address:   "盐亭县",
			StartYear: 1985,
		},
	},
	510724: {
		{
			Address:   "安县",
			StartYear: 1985,
			EndYear:   2015,
		},
	},
	510725: {
		{
			Address:   "梓潼县",
			StartYear: 1985,
		},
	},
	510726: {
		{
			Address:   "北川县",
			StartYear: 1985,
			EndYear:   2002,
		},
		{
			Address:   "北川羌族自治县",
			StartYear: 2003,
		},
	},
	510727: {
		{
			Address:   "平武县",
			StartYear: 1985,
		},
	},
	510781: {
		{
			Address:   "江油市",
			StartYear: 1995,
		},
	},
	510800: {
		{
			Address:   "广元市",
			StartYear: 1985,
		},
	},
	510802: {
		{
			Address:   "市中区",
			StartYear: 1985,
			EndYear:   2006,
		},
		{
			Address:   "利州区",
			StartYear: 2007,
		},
	},
	510811: {
		{
			Address:   "元坝区",
			StartYear: 1989,
			EndYear:   2012,
		},
		{
			Address:   "昭化区",
			StartYear: 2013,
		},
	},
	510812: {
		{
			Address:   "朝天区",
			StartYear: 1989,
		},
	},
	510821: {
		{
			Address:   "旺苍县",
			StartYear: 1985,
		},
	},
	510822: {
		{
			Address:   "青川县",
			StartYear: 1985,
		},
	},
	510823: {
		{
			Address:   "剑阁县",
			StartYear: 1985,
		},
	},
	510824: {
		{
			Address:   "苍溪县",
			StartYear: 1985,
		},
	},
	510900: {
		{
			Address:   "遂宁市",
			StartYear: 1985,
		},
	},
	510902: {
		{
			Address:   "市中区",
			StartYear: 1985,
			EndYear:   2002,
		},
	},
	510903: {
		{
			Address:   "船山区",
			StartYear: 2003,
		},
	},
	510904: {
		{
			Address:   "安居区",
			StartYear: 2003,
		},
	},
	510921: {
		{
			Address:   "蓬溪县",
			StartYear: 1985,
		},
	},
	510922: {
		{
			Address:   "射洪县",
			StartYear: 1985,
			EndYear:   2018,
		},
	},
	510923: {
		{
			Address:   "大英县",
			StartYear: 1997,
		},
	},
	510981: {
		{
			Address:   "射洪市",
			StartYear: 2019,
		},
	},
	511000: {
		{
			Address:   "内江市",
			StartYear: 1985,
		},
	},
	511002: {
		{
			Address:   "市中区",
			StartYear: 1985,
		},
	},
	511011: {
		{
			Address:   "东兴区",
			StartYear: 1989,
		},
	},
	511021: {
		{
			Address:   "内江县",
			StartYear: 1985,
			EndYear:   1988,
		},
	},
	511022: {
		{
			Address:   "乐至县",
			StartYear: 1985,
			EndYear:   1997,
		},
	},
	511023: {
		{
			Address:   "安岳县",
			StartYear: 1985,
			EndYear:   1997,
		},
	},
	511024: {
		{
			Address:   "威远县",
			StartYear: 1985,
		},
	},
	511025: {
		{
			Address:   "资中县",
			StartYear: 1985,
		},
	},
	511026: {
		{
			Address:   "资阳县",
			StartYear: 1985,
			EndYear:   1992,
		},
	},
	511027: {
		{
			Address:   "简阳县",
			StartYear: 1985,
			EndYear:   1993,
		},
	},
	511028: {
		{
			Address:   "隆昌县",
			StartYear: 1985,
			EndYear:   2016,
		},
	},
	511081: {
		{
			Address:   "资阳市",
			StartYear: 1995,
			EndYear:   1997,
		},
	},
	511082: {
		{
			Address:   "简阳市",
			StartYear: 1995,
			EndYear:   1997,
		},
	},
	511083: {
		{
			Address:   "隆昌市",
			StartYear: 2017,
		},
	},
	511100: {
		{
			Address:   "乐山市",
			StartYear: 1985,
		},
	},
	511102: {
		{
			Address:   "市中区",
			StartYear: 1985,
		},
	},
	511111: {
		{
			Address:   "沙湾区",
			StartYear: 1985,
		},
	},
	511112: {
		{
			Address:   "五通桥区",
			StartYear: 1985,
		},
	},
	511113: {
		{
			Address:   "金口河区",
			StartYear: 1985,
		},
	},
	511121: {
		{
			Address:   "仁寿县",
			StartYear: 1985,
			EndYear:   1996,
		},
	},
	511122: {
		{
			Address:   "眉山县",
			StartYear: 1985,
			EndYear:   1996,
		},
	},
	511123: {
		{
			Address:   "犍为县",
			StartYear: 1985,
		},
	},
	511124: {
		{
			Address:   "井研县",
			StartYear: 1985,
		},
	},
	511125: {
		{
			Address:   "峨眉县",
			StartYear: 1985,
			EndYear:   1987,
		},
	},
	511126: {
		{
			Address:   "夹江县",
			StartYear: 1985,
		},
	},
	511127: {
		{
			Address:   "洪雅县",
			StartYear: 1985,
			EndYear:   1996,
		},
	},
	511128: {
		{
			Address:   "彭山县",
			StartYear: 1985,
			EndYear:   1996,
		},
	},
	511129: {
		{
			Address:   "沐川县",
			StartYear: 1985,
		},
	},
	511130: {
		{
			Address:   "青神县",
			StartYear: 1985,
			EndYear:   1996,
		},
	},
	511131: {
		{
			Address:   "丹棱县",
			StartYear: 1985,
			EndYear:   1996,
		},
	},
	511132: {
		{
			Address:   "峨边彝族自治县",
			StartYear: 1985,
		},
	},
	511133: {
		{
			Address:   "马边彝族自治县",
			StartYear: 1985,
		},
	},
	511181: {
		{
			Address:   "峨眉山市",
			StartYear: 1995,
		},
	},
	511200: {
		{
			Address:   "万县市",
			StartYear: 1992,
			EndYear:   1996,
		},
	},
	511202: {
		{
			Address:   "龙宝区",
			StartYear: 1992,
			EndYear:   1996,
		},
	},
	511203: {
		{
			Address:   "天城区",
			StartYear: 1992,
			EndYear:   1996,
		},
	},
	511204: {
		{
			Address:   "五桥区",
			StartYear: 1992,
			EndYear:   1996,
		},
	},
	511221: {
		{
			Address:   "开县",
			StartYear: 1992,
			EndYear:   1996,
		},
	},
	511222: {
		{
			Address:   "忠县",
			StartYear: 1992,
			EndYear:   1996,
		},
	},
	511223: {
		{
			Address:   "梁平县",
			StartYear: 1992,
			EndYear:   1996,
		},
	},
	511224: {
		{
			Address:   "云阳县",
			StartYear: 1992,
			EndYear:   1996,
		},
	},
	511225: {
		{
			Address:   "奉节县",
			StartYear: 1992,
			EndYear:   1996,
		},
	},
	511226: {
		{
			Address:   "巫山县",
			StartYear: 1992,
			EndYear:   1996,
		},
	},
	511227: {
		{
			Address:   "巫溪县",
			StartYear: 1992,
			EndYear:   1996,
		},
	},
	511228: {
		{
			Address:   "城口县",
			StartYear: 1992,
			EndYear:   1996,
		},
	},
	511300: {
		{
			Address:   "南充市",
			StartYear: 1993,
		},
	},
	511302: {
		{
			Address:   "顺庆区",
			StartYear: 1993,
		},
	},
	511303: {
		{
			Address:   "高坪区",
			StartYear: 1993,
		},
	},
	511304: {
		{
			Address:   "嘉陵区",
			StartYear: 1993,
		},
	},
	511321: {
		{
			Address:   "南部县",
			StartYear: 1993,
		},
	},
	511322: {
		{
			Address:   "营山县",
			StartYear: 1993,
		},
	},
	511323: {
		{
			Address:   "蓬安县",
			StartYear: 1993,
		},
	},
	511324: {
		{
			Address:   "仪陇县",
			StartYear: 1993,
		},
	},
	511325: {
		{
			Address:   "西充县",
			StartYear: 1993,
		},
	},
	511381: {
		{
			Address:   "阆中市",
			StartYear: 1995,
		},
	},
	511400: {
		{
			Address:   "眉山市",
			StartYear: 2000,
		},
	},
	511402: {
		{
			Address:   "东坡区",
			StartYear: 2000,
		},
	},
	511403: {
		{
			Address:   "彭山区",
			StartYear: 2014,
		},
	},
	511421: {
		{
			Address:   "仁寿县",
			StartYear: 2000,
		},
	},
	511422: {
		{
			Address:   "彭山县",
			StartYear: 2000,
			EndYear:   2013,
		},
	},
	511423: {
		{
			Address:   "洪雅县",
			StartYear: 2000,
		},
	},
	511424: {
		{
			Address:   "丹棱县",
			StartYear: 2000,
		},
	},
	511425: {
		{
			Address:   "青神县",
			StartYear: 2000,
		},
	},
	511500: {
		{
			Address:   "宜宾市",
			StartYear: 1996,
		},
	},
	511502: {
		{
			Address:   "翠屏区",
			StartYear: 1996,
		},
	},
	511503: {
		{
			Address:   "南溪区",
			StartYear: 2011,
		},
	},
	511504: {
		{
			Address:   "叙州区",
			StartYear: 2018,
		},
	},
	511521: {
		{
			Address:   "宜宾县",
			StartYear: 1996,
			EndYear:   2017,
		},
	},
	511522: {
		{
			Address:   "南溪县",
			StartYear: 1996,
			EndYear:   2010,
		},
	},
	511523: {
		{
			Address:   "江安县",
			StartYear: 1996,
		},
	},
	511524: {
		{
			Address:   "长宁县",
			StartYear: 1996,
		},
	},
	511525: {
		{
			Address:   "高县",
			StartYear: 1996,
		},
	},
	511526: {
		{
			Address:   "珙县",
			StartYear: 1996,
		},
	},
	511527: {
		{
			Address:   "筠连县",
			StartYear: 1996,
		},
	},
	511528: {
		{
			Address:   "兴文县",
			StartYear: 1996,
		},
	},
	511529: {
		{
			Address:   "屏山县",
			StartYear: 1996,
		},
	},
	511600: {
		{
			Address:   "广安市",
			StartYear: 1998,
		},
	},
	511602: {
		{
			Address:   "广安区",
			StartYear: 1998,
		},
	},
	511603: {
		{
			Address:   "前锋区",
			StartYear: 2013,
		},
	},
	511621: {
		{
			Address:   "岳池县",
			StartYear: 1998,
		},
	},
	511622: {
		{
			Address:   "武胜县",
			StartYear: 1998,
		},
	},
	511623: {
		{
			Address:   "邻水县",
			StartYear: 1998,
		},
	},
	511681: {
		{
			Address:   "华蓥市",
			StartYear: 1998,
		},
	},
	511700: {
		{
			Address:   "达州市",
			StartYear: 1999,
		},
	},
	511702: {
		{
			Address:   "通川区",
			StartYear: 1999,
		},
	},
	511703: {
		{
			Address:   "达川区",
			StartYear: 2013,
		},
	},
	511721: {
		{
			Address:   "达县",
			StartYear: 1999,
			EndYear:   2012,
		},
	},
	511722: {
		{
			Address:   "宣汉县",
			StartYear: 1999,
		},
	},
	511723: {
		{
			Address:   "开江县",
			StartYear: 1999,
		},
	},
	511724: {
		{
			Address:   "大竹县",
			StartYear: 1999,
		},
	},
	511725: {
		{
			Address:   "渠县",
			StartYear: 1999,
		},
	},
	511781: {
		{
			Address:   "万源市",
			StartYear: 1999,
		},
	},
	511800: {
		{
			Address:   "雅安市",
			StartYear: 2000,
		},
	},
	511802: {
		{
			Address:   "雨城区",
			StartYear: 2000,
		},
	},
	511803: {
		{
			Address:   "名山区",
			StartYear: 2012,
		},
	},
	511821: {
		{
			Address:   "名山县",
			StartYear: 2000,
			EndYear:   2011,
		},
	},
	511822: {
		{
			Address:   "荥经县",
			StartYear: 2000,
		},
	},
	511823: {
		{
			Address:   "汉源县",
			StartYear: 2000,
		},
	},
	511824: {
		{
			Address:   "石棉县",
			StartYear: 2000,
		},
	},
	511825: {
		{
			Address:   "天全县",
			StartYear: 2000,
		},
	},
	511826: {
		{
			Address:   "芦山县",
			StartYear: 2000,
		},
	},
	511827: {
		{
			Address:   "宝兴县",
			StartYear: 2000,
		},
	},
	511900: {
		{
			Address:   "巴中市",
			StartYear: 2000,
		},
	},
	511902: {
		{
			Address:   "巴州区",
			StartYear: 2000,
		},
	},
	511903: {
		{
			Address:   "恩阳区",
			StartYear: 2013,
		},
	},
	511921: {
		{
			Address:   "通江县",
			StartYear: 2000,
		},
	},
	511922: {
		{
			Address:   "南江县",
			StartYear: 2000,
		},
	},
	511923: {
		{
			Address:   "平昌县",
			StartYear: 2000,
		},
	},
	512000: {
		{
			Address:   "资阳市",
			StartYear: 2000,
		},
	},
	512002: {
		{
			Address:   "雁江区",
			StartYear: 2000,
		},
	},
	512021: {
		{
			Address:   "安岳县",
			StartYear: 2000,
		},
	},
	512022: {
		{
			Address:   "乐至县",
			StartYear: 2000,
		},
	},
	512081: {
		{
			Address:   "简阳市",
			StartYear: 2000,
			EndYear:   2015,
		},
	},
	512100: {
		{
			Address: "温江地区",
			EndYear: 1981,
		},
		{
			Address:   "永川地区",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	512121: {
		{
			Address: "温江县",
			EndYear: 1981,
		},
		{
			Address:   "江津县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	512122: {
		{
			Address: "郫县",
			EndYear: 1981,
		},
		{
			Address:   "合川县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	512123: {
		{
			Address: "灌县",
			EndYear: 1981,
		},
		{
			Address:   "潼南县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	512124: {
		{
			Address: "彭县",
			EndYear: 1981,
		},
		{
			Address:   "铜梁县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	512125: {
		{
			Address: "什邡县",
			EndYear: 1981,
		},
		{
			Address:   "永川县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	512126: {
		{
			Address: "广汉县",
			EndYear: 1981,
		},
		{
			Address:   "大足县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	512127: {
		{
			Address: "新都县",
			EndYear: 1981,
		},
		{
			Address:   "荣昌县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	512128: {
		{
			Address: "新津县",
			EndYear: 1981,
		},
		{
			Address:   "璧山县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	512129: {
		{
			Address: "蒲江县",
			EndYear: 1981,
		},
	},
	512130: {
		{
			Address: "邛崃县",
			EndYear: 1981,
		},
	},
	512131: {
		{
			Address: "大邑县",
			EndYear: 1981,
		},
	},
	512132: {
		{
			Address: "崇庆县",
			EndYear: 1981,
		},
	},
	512200: {
		{
			Address: "绵阳地区",
			EndYear: 1981,
		},
		{
			Address:   "万县地区",
			StartYear: 1982,
			EndYear:   1991,
		},
	},
	512201: {
		{
			Address: "绵阳市",
			EndYear: 1981,
		},
		{
			Address:   "万县市",
			StartYear: 1982,
			EndYear:   1991,
		},
	},
	512202: {
		{
			Address: "市中区",
			EndYear: 1981,
		},
	},
	512221: {
		{
			Address: "江油县",
			EndYear: 1981,
		},
		{
			Address:   "万县",
			StartYear: 1982,
			EndYear:   1991,
		},
	},
	512222: {
		{
			Address: "青川县",
			EndYear: 1981,
		},
		{
			Address:   "开县",
			StartYear: 1982,
			EndYear:   1991,
		},
	},
	512223: {
		{
			Address: "平武县",
			EndYear: 1981,
		},
		{
			Address:   "忠县",
			StartYear: 1982,
			EndYear:   1991,
		},
	},
	512224: {
		{
			Address: "广元县",
			EndYear: 1981,
		},
		{
			Address:   "梁平县",
			StartYear: 1982,
			EndYear:   1991,
		},
	},
	512225: {
		{
			Address: "旺苍县",
			EndYear: 1981,
		},
		{
			Address:   "云阳县",
			StartYear: 1982,
			EndYear:   1991,
		},
	},
	512226: {
		{
			Address: "剑阁县",
			EndYear: 1981,
		},
		{
			Address:   "奉节县",
			StartYear: 1982,
			EndYear:   1991,
		},
	},
	512227: {
		{
			Address: "梓潼县",
			EndYear: 1981,
		},
		{
			Address:   "巫山县",
			StartYear: 1982,
			EndYear:   1991,
		},
	},
	512228: {
		{
			Address: "三台县",
			EndYear: 1981,
		},
		{
			Address:   "巫溪县",
			StartYear: 1982,
			EndYear:   1991,
		},
	},
	512229: {
		{
			Address: "盐亭县",
			EndYear: 1981,
		},
		{
			Address:   "城口县",
			StartYear: 1982,
			EndYear:   1991,
		},
	},
	512230: {
		{
			Address: "射洪县",
			EndYear: 1981,
		},
	},
	512231: {
		{
			Address: "遂宁县",
			EndYear: 1981,
		},
	},
	512232: {
		{
			Address: "蓬溪县",
			EndYear: 1981,
		},
	},
	512233: {
		{
			Address: "中江县",
			EndYear: 1981,
		},
	},
	512234: {
		{
			Address: "德阳县",
			EndYear: 1981,
		},
	},
	512235: {
		{
			Address: "绵竹县",
			EndYear: 1981,
		},
	},
	512236: {
		{
			Address: "安县",
			EndYear: 1981,
		},
	},
	512237: {
		{
			Address: "北川县",
			EndYear: 1981,
		},
	},
	512300: {
		{
			Address: "内江地区",
			EndYear: 1981,
		},
		{
			Address:   "涪陵地区",
			StartYear: 1982,
			EndYear:   1994,
		},
	},
	512301: {
		{
			Address: "内江市",
			EndYear: 1981,
		},
		{
			Address:   "涪陵市",
			StartYear: 1983,
			EndYear:   1994,
		},
	},
	512302: {
		{
			Address:   "南川市",
			StartYear: 1994,
			EndYear:   1994,
		},
	},
	512321: {
		{
			Address: "内江县",
			EndYear: 1981,
		},
		{
			Address:   "涪陵县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	512322: {
		{
			Address: "乐至县",
			EndYear: 1981,
		},
		{
			Address:   "垫江县",
			StartYear: 1982,
			EndYear:   1994,
		},
	},
	512323: {
		{
			Address: "安岳县",
			EndYear: 1981,
		},
		{
			Address:   "南川县",
			StartYear: 1982,
			EndYear:   1993,
		},
	},
	512324: {
		{
			Address: "威远县",
			EndYear: 1981,
		},
		{
			Address:   "丰都县",
			StartYear: 1982,
			EndYear:   1994,
		},
	},
	512325: {
		{
			Address: "资中县",
			EndYear: 1981,
		},
		{
			Address:   "石柱县",
			StartYear: 1982,
			EndYear:   1982,
		},
		{
			Address:   "石柱土家族自治县",
			StartYear: 1983,
			EndYear:   1987,
		},
	},
	512326: {
		{
			Address: "资阳县",
			EndYear: 1981,
		},
		{
			Address:   "武隆县",
			StartYear: 1982,
			EndYear:   1994,
		},
	},
	512327: {
		{
			Address: "简阳县",
			EndYear: 1981,
		},
		{
			Address:   "彭水县",
			StartYear: 1982,
			EndYear:   1982,
		},
		{
			Address:   "彭水苗族土家族自治县",
			StartYear: 1983,
			EndYear:   1987,
		},
	},
	512328: {
		{
			Address:   "黔江县",
			StartYear: 1982,
			EndYear:   1982,
		},
		{
			Address:   "黔江土家族苗族自治县",
			StartYear: 1983,
			EndYear:   1987,
		},
	},
	512329: {
		{
			Address:   "酉阳县",
			StartYear: 1982,
			EndYear:   1982,
		},
		{
			Address:   "酉阳土家族苗族自治县",
			StartYear: 1983,
			EndYear:   1987,
		},
	},
	512330: {
		{
			Address:   "秀山县",
			StartYear: 1982,
			EndYear:   1982,
		},
		{
			Address:   "秀山土家族苗族自治县",
			StartYear: 1983,
			EndYear:   1987,
		},
	},
	512400: {
		{
			Address: "宜宾地区",
			EndYear: 1981,
		},
		{
			Address:   "内江地区",
			StartYear: 1982,
			EndYear:   1984,
		},
	},
	512401: {
		{
			Address: "宜宾市",
			EndYear: 1981,
		},
		{
			Address:   "内江市",
			StartYear: 1982,
			EndYear:   1984,
		},
	},
	512402: {
		{
			Address: "泸州市",
			EndYear: 1981,
		},
	},
	512421: {
		{
			Address: "宜宾县",
			EndYear: 1981,
		},
		{
			Address:   "内江县",
			StartYear: 1982,
			EndYear:   1984,
		},
	},
	512422: {
		{
			Address: "富顺县",
			EndYear: 1981,
		},
		{
			Address:   "资中县",
			StartYear: 1982,
			EndYear:   1984,
		},
	},
	512423: {
		{
			Address: "隆昌县",
			EndYear: 1981,
		},
		{
			Address:   "资阳县",
			StartYear: 1982,
			EndYear:   1984,
		},
	},
	512424: {
		{
			Address: "南溪县",
			EndYear: 1981,
		},
		{
			Address:   "简阳县",
			StartYear: 1982,
			EndYear:   1984,
		},
	},
	512425: {
		{
			Address: "江安县",
			EndYear: 1981,
		},
		{
			Address:   "威远县",
			StartYear: 1982,
			EndYear:   1984,
		},
	},
	512426: {
		{
			Address: "纳溪县",
			EndYear: 1981,
		},
		{
			Address:   "隆昌县",
			StartYear: 1982,
			EndYear:   1984,
		},
	},
	512427: {
		{
			Address: "泸县",
			EndYear: 1981,
		},
		{
			Address:   "安岳县",
			StartYear: 1982,
			EndYear:   1984,
		},
	},
	512428: {
		{
			Address: "合江县",
			EndYear: 1981,
		},
		{
			Address:   "乐至县",
			StartYear: 1982,
			EndYear:   1984,
		},
	},
	512429: {
		{
			Address: "古蔺县",
			EndYear: 1981,
		},
	},
	512430: {
		{
			Address: "叙永县",
			EndYear: 1981,
		},
	},
	512431: {
		{
			Address: "长宁县",
			EndYear: 1981,
		},
	},
	512432: {
		{
			Address: "兴文县",
			EndYear: 1981,
		},
	},
	512433: {
		{
			Address: "珙县",
			EndYear: 1981,
		},
	},
	512434: {
		{
			Address: "高县",
			EndYear: 1981,
		},
	},
	512435: {
		{
			Address: "筠连县",
			EndYear: 1981,
		},
	},
	512436: {
		{
			Address: "屏山县",
			EndYear: 1981,
		},
	},
	512500: {
		{
			Address: "乐山地区",
			EndYear: 1981,
		},
		{
			Address:   "宜宾地区",
			StartYear: 1982,
			EndYear:   1995,
		},
	},
	512501: {
		{
			Address: "乐山市",
			EndYear: 1981,
		},
		{
			Address:   "宜宾市",
			StartYear: 1982,
			EndYear:   1995,
		},
	},
	512502: {
		{
			Address:   "泸州市",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	512521: {
		{
			Address: "夹江县",
			EndYear: 1981,
		},
		{
			Address:   "泸县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	512522: {
		{
			Address: "洪雅县",
			EndYear: 1981,
		},
		{
			Address:   "富顺县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	512523: {
		{
			Address: "丹棱县",
			EndYear: 1981,
		},
		{
			Address:   "合江县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	512524: {
		{
			Address: "青神县",
			EndYear: 1981,
		},
		{
			Address:   "纳溪县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	512525: {
		{
			Address: "眉山县",
			EndYear: 1981,
		},
		{
			Address:   "叙永县",
			StartYear: 1982,
			EndYear:   1984,
		},
	},
	512526: {
		{
			Address: "彭山县",
			EndYear: 1981,
		},
		{
			Address:   "古蔺县",
			StartYear: 1982,
			EndYear:   1984,
		},
	},
	512527: {
		{
			Address: "井研县",
			EndYear: 1981,
		},
		{
			Address:   "宜宾县",
			StartYear: 1982,
			EndYear:   1995,
		},
	},
	512528: {
		{
			Address: "仁寿县",
			EndYear: 1981,
		},
		{
			Address:   "南溪县",
			StartYear: 1982,
			EndYear:   1995,
		},
	},
	512529: {
		{
			Address: "犍为县",
			EndYear: 1981,
		},
		{
			Address:   "江安县",
			StartYear: 1982,
			EndYear:   1995,
		},
	},
	512530: {
		{
			Address: "沐川县",
			EndYear: 1981,
		},
		{
			Address:   "长宁县",
			StartYear: 1982,
			EndYear:   1995,
		},
	},
	512531: {
		{
			Address: "峨眉县",
			EndYear: 1981,
		},
		{
			Address:   "高县",
			StartYear: 1982,
			EndYear:   1995,
		},
	},
	512532: {
		{
			Address: "金口河工农区",
			EndYear: 1981,
		},
		{
			Address:   "筠连县",
			StartYear: 1982,
			EndYear:   1995,
		},
	},
	512533: {
		{
			Address:   "峨边县",
			StartYear: 1981,
			EndYear:   1981,
		},
		{
			Address:   "珙县",
			StartYear: 1982,
			EndYear:   1995,
		},
	},
	512534: {
		{
			Address:   "马边县",
			StartYear: 1981,
			EndYear:   1981,
		},
		{
			Address:   "兴文县",
			StartYear: 1982,
			EndYear:   1995,
		},
	},
	512535: {
		{
			Address:   "屏山县",
			StartYear: 1982,
			EndYear:   1995,
		},
	},
	512600: {
		{
			Address: "江津地区",
			EndYear: 1980,
		},
		{
			Address:   "永川地区",
			StartYear: 1981,
			EndYear:   1981,
		},
		{
			Address:   "乐山地区",
			StartYear: 1982,
			EndYear:   1984,
		},
	},
	512601: {
		{
			Address:   "乐山市",
			StartYear: 1982,
			EndYear:   1984,
		},
	},
	512621: {
		{
			Address: "永川县",
			EndYear: 1981,
		},
		{
			Address:   "仁寿县",
			StartYear: 1982,
			EndYear:   1984,
		},
	},
	512622: {
		{
			Address: "大足县",
			EndYear: 1981,
		},
		{
			Address:   "眉山县",
			StartYear: 1982,
			EndYear:   1984,
		},
	},
	512623: {
		{
			Address: "铜梁县",
			EndYear: 1981,
		},
		{
			Address:   "犍为县",
			StartYear: 1982,
			EndYear:   1984,
		},
	},
	512624: {
		{
			Address: "合川县",
			EndYear: 1981,
		},
		{
			Address:   "井研县",
			StartYear: 1982,
			EndYear:   1984,
		},
	},
	512625: {
		{
			Address: "潼南县",
			EndYear: 1981,
		},
		{
			Address:   "峨眉县",
			StartYear: 1982,
			EndYear:   1984,
		},
	},
	512626: {
		{
			Address: "璧山县",
			EndYear: 1981,
		},
		{
			Address:   "夹江县",
			StartYear: 1982,
			EndYear:   1984,
		},
	},
	512627: {
		{
			Address: "江津县",
			EndYear: 1981,
		},
		{
			Address:   "洪雅县",
			StartYear: 1982,
			EndYear:   1984,
		},
	},
	512628: {
		{
			Address: "荣昌县",
			EndYear: 1981,
		},
		{
			Address:   "彭山县",
			StartYear: 1982,
			EndYear:   1984,
		},
	},
	512629: {
		{
			Address:   "沐川县",
			StartYear: 1982,
			EndYear:   1984,
		},
	},
	512630: {
		{
			Address:   "青神县",
			StartYear: 1982,
			EndYear:   1984,
		},
	},
	512631: {
		{
			Address:   "丹棱县",
			StartYear: 1982,
			EndYear:   1984,
		},
	},
	512632: {
		{
			Address:   "峨边县",
			StartYear: 1982,
			EndYear:   1982,
		},
		{
			Address:   "峨边彝族自治县",
			StartYear: 1983,
			EndYear:   1984,
		},
	},
	512633: {
		{
			Address:   "马边县",
			StartYear: 1982,
			EndYear:   1982,
		},
		{
			Address:   "马边彝族自治县",
			StartYear: 1983,
			EndYear:   1984,
		},
	},
	512634: {
		{
			Address:   "金口河工农区",
			StartYear: 1982,
			EndYear:   1984,
		},
	},
	512700: {
		{
			Address: "涪陵地区",
			EndYear: 1981,
		},
		{
			Address:   "温江地区",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	512721: {
		{
			Address: "涪陵县",
			EndYear: 1981,
		},
		{
			Address:   "温江县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	512722: {
		{
			Address: "垫江县",
			EndYear: 1981,
		},
		{
			Address:   "郫县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	512723: {
		{
			Address: "丰都县",
			EndYear: 1981,
		},
		{
			Address:   "新都县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	512724: {
		{
			Address: "石柱县",
			EndYear: 1981,
		},
		{
			Address:   "广汉县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	512725: {
		{
			Address: "秀山县",
			EndYear: 1981,
		},
		{
			Address:   "什邡县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	512726: {
		{
			Address: "酉阳县",
			EndYear: 1981,
		},
		{
			Address:   "彭县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	512727: {
		{
			Address: "黔江县",
			EndYear: 1981,
		},
		{
			Address:   "灌县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	512728: {
		{
			Address: "彭水县",
			EndYear: 1981,
		},
		{
			Address:   "崇庆县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	512729: {
		{
			Address: "武隆县",
			EndYear: 1981,
		},
		{
			Address:   "大邑县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	512730: {
		{
			Address: "南川县",
			EndYear: 1981,
		},
		{
			Address:   "邛崃县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	512731: {
		{
			Address:   "蒲江县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	512732: {
		{
			Address:   "新津县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	512800: {
		{
			Address: "万县地区",
			EndYear: 1981,
		},
		{
			Address:   "绵阳地区",
			StartYear: 1982,
			EndYear:   1984,
		},
	},
	512801: {
		{
			Address: "万县市",
			EndYear: 1981,
		},
		{
			Address:   "绵阳市",
			StartYear: 1982,
			EndYear:   1984,
		},
	},
	512802: {
		{
			Address:   "市中区",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	512821: {
		{
			Address: "万县",
			EndYear: 1981,
		},
		{
			Address:   "德阳县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	512822: {
		{
			Address: "开县",
			EndYear: 1981,
		},
		{
			Address:   "绵竹县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	512823: {
		{
			Address: "城口县",
			EndYear: 1981,
		},
		{
			Address:   "安县",
			StartYear: 1982,
			EndYear:   1984,
		},
	},
	512824: {
		{
			Address: "巫溪县",
			EndYear: 1981,
		},
		{
			Address:   "江油县",
			StartYear: 1982,
			EndYear:   1984,
		},
	},
	512825: {
		{
			Address: "巫山县",
			EndYear: 1981,
		},
		{
			Address:   "梓潼县",
			StartYear: 1982,
			EndYear:   1984,
		},
	},
	512826: {
		{
			Address: "奉节县",
			EndYear: 1981,
		},
		{
			Address:   "剑阁县",
			StartYear: 1982,
			EndYear:   1984,
		},
	},
	512827: {
		{
			Address: "云阳县",
			EndYear: 1981,
		},
		{
			Address:   "广元县",
			StartYear: 1982,
			EndYear:   1984,
		},
	},
	512828: {
		{
			Address: "忠县",
			EndYear: 1981,
		},
		{
			Address:   "旺苍县",
			StartYear: 1982,
			EndYear:   1984,
		},
	},
	512829: {
		{
			Address: "梁平县",
			EndYear: 1981,
		},
		{
			Address:   "青川县",
			StartYear: 1982,
			EndYear:   1984,
		},
	},
	512830: {
		{
			Address:   "平武县",
			StartYear: 1982,
			EndYear:   1984,
		},
	},
	512831: {
		{
			Address:   "北川县",
			StartYear: 1982,
			EndYear:   1984,
		},
	},
	512832: {
		{
			Address:   "遂宁县",
			StartYear: 1982,
			EndYear:   1984,
		},
	},
	512833: {
		{
			Address:   "三台县",
			StartYear: 1982,
			EndYear:   1984,
		},
	},
	512834: {
		{
			Address:   "中江县",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	512835: {
		{
			Address:   "蓬溪县",
			StartYear: 1982,
			EndYear:   1984,
		},
	},
	512836: {
		{
			Address:   "射洪县",
			StartYear: 1982,
			EndYear:   1984,
		},
	},
	512837: {
		{
			Address:   "盐亭县",
			StartYear: 1982,
			EndYear:   1984,
		},
	},
	512900: {
		{
			Address: "南充地区",
			EndYear: 1992,
		},
	},
	512901: {
		{
			Address: "南充市",
			EndYear: 1992,
		},
	},
	512902: {
		{
			Address:   "华蓥市",
			StartYear: 1985,
			EndYear:   1992,
		},
	},
	512903: {
		{
			Address:   "阆中市",
			StartYear: 1991,
			EndYear:   1992,
		},
	},
	512921: {
		{
			Address: "南充县",
			EndYear: 1992,
		},
	},
	512922: {
		{
			Address: "苍溪县",
			EndYear: 1981,
		},
		{
			Address:   "南部县",
			StartYear: 1982,
			EndYear:   1992,
		},
	},
	512923: {
		{
			Address: "阆中县",
			EndYear: 1981,
		},
		{
			Address:   "岳池县",
			StartYear: 1982,
			EndYear:   1992,
		},
	},
	512924: {
		{
			Address: "仪陇县",
			EndYear: 1981,
		},
		{
			Address:   "营山县",
			StartYear: 1982,
			EndYear:   1992,
		},
	},
	512925: {
		{
			Address: "南部县",
			EndYear: 1981,
		},
		{
			Address:   "广安县",
			StartYear: 1982,
			EndYear:   1992,
		},
	},
	512926: {
		{
			Address: "西充县",
			EndYear: 1981,
		},
		{
			Address:   "蓬安县",
			StartYear: 1982,
			EndYear:   1992,
		},
	},
	512927: {
		{
			Address: "营山县",
			EndYear: 1981,
		},
		{
			Address:   "仪陇县",
			StartYear: 1982,
			EndYear:   1992,
		},
	},
	512928: {
		{
			Address: "蓬安县",
			EndYear: 1981,
		},
		{
			Address:   "武胜县",
			StartYear: 1982,
			EndYear:   1992,
		},
	},
	512929: {
		{
			Address: "广安县",
			EndYear: 1981,
		},
		{
			Address:   "西充县",
			StartYear: 1982,
			EndYear:   1992,
		},
	},
	512930: {
		{
			Address: "岳池县",
			EndYear: 1981,
		},
		{
			Address:   "阆中县",
			StartYear: 1982,
			EndYear:   1990,
		},
	},
	512931: {
		{
			Address: "武胜县",
			EndYear: 1981,
		},
		{
			Address:   "苍溪县",
			StartYear: 1982,
			EndYear:   1984,
		},
	},
	512932: {
		{
			Address: "华云工农区",
			EndYear: 1984,
		},
	},
	513000: {
		{
			Address: "达县地区",
			EndYear: 1992,
		},
		{
			Address:   "达川地区",
			StartYear: 1993,
			EndYear:   1998,
		},
	},
	513001: {
		{
			Address: "达县市",
			EndYear: 1992,
		},
		{
			Address:   "达川市",
			StartYear: 1993,
			EndYear:   1998,
		},
	},
	513002: {
		{
			Address:   "万源市",
			StartYear: 1993,
			EndYear:   1998,
		},
	},
	513021: {
		{
			Address: "达县",
			EndYear: 1998,
		},
	},
	513022: {
		{
			Address: "万源县",
			EndYear: 1981,
		},
		{
			Address:   "宣汉县",
			StartYear: 1982,
			EndYear:   1998,
		},
	},
	513023: {
		{
			Address: "宣汉县",
			EndYear: 1981,
		},
		{
			Address:   "开江县",
			StartYear: 1982,
			EndYear:   1998,
		},
	},
	513024: {
		{
			Address: "开江县",
			EndYear: 1981,
		},
		{
			Address:   "万源县",
			StartYear: 1982,
			EndYear:   1992,
		},
	},
	513025: {
		{
			Address: "邻水县",
			EndYear: 1981,
		},
		{
			Address:   "通江县",
			StartYear: 1982,
			EndYear:   1992,
		},
	},
	513026: {
		{
			Address: "大竹县",
			EndYear: 1981,
		},
		{
			Address:   "南江县",
			StartYear: 1982,
			EndYear:   1992,
		},
	},
	513027: {
		{
			Address: "渠县",
			EndYear: 1981,
		},
		{
			Address:   "巴中县",
			StartYear: 1982,
			EndYear:   1992,
		},
	},
	513028: {
		{
			Address: "南江县",
			EndYear: 1981,
		},
		{
			Address:   "平昌县",
			StartYear: 1982,
			EndYear:   1992,
		},
	},
	513029: {
		{
			Address: "巴中县",
			EndYear: 1981,
		},
		{
			Address:   "大竹县",
			StartYear: 1982,
			EndYear:   1998,
		},
	},
	513030: {
		{
			Address: "平昌县",
			EndYear: 1981,
		},
		{
			Address:   "渠县",
			StartYear: 1982,
			EndYear:   1998,
		},
	},
	513031: {
		{
			Address: "通江县",
			EndYear: 1981,
		},
		{
			Address:   "邻水县",
			StartYear: 1982,
			EndYear:   1992,
		},
	},
	513032: {
		{
			Address: "白沙工农区",
			EndYear: 1992,
		},
	},
	513100: {
		{
			Address: "雅安地区",
			EndYear: 1999,
		},
	},
	513101: {
		{
			Address:   "雅安市",
			StartYear: 1983,
			EndYear:   1999,
		},
	},
	513121: {
		{
			Address: "雅安县",
			EndYear: 1982,
		},
	},
	513122: {
		{
			Address: "芦山县",
			EndYear: 1981,
		},
		{
			Address:   "名山县",
			StartYear: 1982,
			EndYear:   1999,
		},
	},
	513123: {
		{
			Address: "名山县",
			EndYear: 1981,
		},
		{
			Address:   "荥经县",
			StartYear: 1982,
			EndYear:   1999,
		},
	},
	513124: {
		{
			Address: "荥经县",
			EndYear: 1981,
		},
		{
			Address:   "汉源县",
			StartYear: 1982,
			EndYear:   1999,
		},
	},
	513125: {
		{
			Address: "汉源县",
			EndYear: 1981,
		},
		{
			Address:   "石棉县",
			StartYear: 1982,
			EndYear:   1999,
		},
	},
	513126: {
		{
			Address: "石棉县",
			EndYear: 1981,
		},
		{
			Address:   "天全县",
			StartYear: 1982,
			EndYear:   1999,
		},
	},
	513127: {
		{
			Address: "天全县",
			EndYear: 1981,
		},
		{
			Address:   "芦山县",
			StartYear: 1982,
			EndYear:   1999,
		},
	},
	513128: {
		{
			Address: "宝兴县",
			EndYear: 1999,
		},
	},
	513200: {
		{
			Address: "阿坝藏族自治州",
			EndYear: 1986,
		},
		{
			Address:   "阿坝藏族羌族自治州",
			StartYear: 1987,
		},
	},
	513201: {
		{
			Address:   "马尔康区",
			StartYear: 2015,
			EndYear:   2015,
		},
		{
			Address:   "马尔康市",
			StartYear: 2016,
		},
	},
	513221: {
		{
			Address: "马尔康县",
			EndYear: 1981,
		},
		{
			Address:   "汶川县",
			StartYear: 1982,
		},
	},
	513222: {
		{
			Address: "红原县",
			EndYear: 1981,
		},
		{
			Address:   "理县",
			StartYear: 1982,
		},
	},
	513223: {
		{
			Address: "阿坝县",
			EndYear: 1981,
		},
		{
			Address:   "茂汶羌族自治县",
			StartYear: 1982,
			EndYear:   1986,
		},
		{
			Address:   "茂县",
			StartYear: 1987,
		},
	},
	513224: {
		{
			Address: "若尔盖县",
			EndYear: 1981,
		},
		{
			Address:   "松潘县",
			StartYear: 1982,
		},
	},
	513225: {
		{
			Address: "黑水县",
			EndYear: 1981,
		},
		{
			Address:   "南坪县",
			StartYear: 1982,
			EndYear:   1996,
		},
		{
			Address:   "九寨沟县",
			StartYear: 1997,
		},
	},
	513226: {
		{
			Address: "松潘县",
			EndYear: 1981,
		},
		{
			Address:   "金川县",
			StartYear: 1982,
		},
	},
	513227: {
		{
			Address: "南坪县",
			EndYear: 1981,
		},
		{
			Address:   "小金县",
			StartYear: 1982,
		},
	},
	513228: {
		{
			Address: "茂汶羌族自治县",
			EndYear: 1981,
		},
		{
			Address:   "黑水县",
			StartYear: 1982,
		},
	},
	513229: {
		{
			Address: "汶川县",
			EndYear: 1981,
		},
		{
			Address:   "马尔康县",
			StartYear: 1982,
			EndYear:   2014,
		},
	},
	513230: {
		{
			Address: "理县",
			EndYear: 1981,
		},
		{
			Address:   "壤塘县",
			StartYear: 1982,
		},
	},
	513231: {
		{
			Address: "小金县",
			EndYear: 1981,
		},
		{
			Address:   "阿坝县",
			StartYear: 1982,
		},
	},
	513232: {
		{
			Address: "金川县",
			EndYear: 1981,
		},
		{
			Address:   "若尔盖县",
			StartYear: 1982,
		},
	},
	513233: {
		{
			Address: "壤塘县",
			EndYear: 1981,
		},
		{
			Address:   "红原县",
			StartYear: 1982,
		},
	},
	513300: {
		{
			Address: "甘孜藏族自治州",
		},
	},
	513301: {
		{
			Address:   "康定市",
			StartYear: 2015,
		},
	},
	513321: {
		{
			Address: "康定县",
			EndYear: 2014,
		},
	},
	513322: {
		{
			Address: "炉霍县",
			EndYear: 1981,
		},
		{
			Address:   "泸定县",
			StartYear: 1982,
		},
	},
	513323: {
		{
			Address: "甘孜县",
			EndYear: 1981,
		},
		{
			Address:   "丹巴县",
			StartYear: 1982,
		},
	},
	513324: {
		{
			Address: "新龙县",
			EndYear: 1981,
		},
		{
			Address:   "九龙县",
			StartYear: 1982,
		},
	},
	513325: {
		{
			Address: "白玉县",
			EndYear: 1981,
		},
		{
			Address:   "雅江县",
			StartYear: 1982,
		},
	},
	513326: {
		{
			Address: "德格县",
			EndYear: 1981,
		},
		{
			Address:   "道孚县",
			StartYear: 1982,
		},
	},
	513327: {
		{
			Address: "石渠县",
			EndYear: 1981,
		},
		{
			Address:   "炉霍县",
			StartYear: 1982,
		},
	},
	513328: {
		{
			Address: "色达县",
			EndYear: 1981,
		},
		{
			Address:   "甘孜县",
			StartYear: 1982,
		},
	},
	513329: {
		{
			Address: "泸定县",
			EndYear: 1981,
		},
		{
			Address:   "新龙县",
			StartYear: 1982,
		},
	},
	513330: {
		{
			Address: "丹巴县",
			EndYear: 1981,
		},
		{
			Address:   "德格县",
			StartYear: 1982,
		},
	},
	513331: {
		{
			Address: "九龙县",
			EndYear: 1981,
		},
		{
			Address:   "白玉县",
			StartYear: 1982,
		},
	},
	513332: {
		{
			Address: "雅江县",
			EndYear: 1981,
		},
		{
			Address:   "石渠县",
			StartYear: 1982,
		},
	},
	513333: {
		{
			Address: "道孚县",
			EndYear: 1981,
		},
		{
			Address:   "色达县",
			StartYear: 1982,
		},
	},
	513334: {
		{
			Address: "理塘县",
		},
	},
	513335: {
		{
			Address: "乡城县",
			EndYear: 1981,
		},
		{
			Address:   "巴塘县",
			StartYear: 1982,
		},
	},
	513336: {
		{
			Address: "稻城县",
			EndYear: 1981,
		},
		{
			Address:   "乡城县",
			StartYear: 1982,
		},
	},
	513337: {
		{
			Address: "巴塘县",
			EndYear: 1981,
		},
		{
			Address:   "稻城县",
			StartYear: 1982,
		},
	},
	513338: {
		{
			Address: "得荣县",
		},
	},
	513400: {
		{
			Address: "凉山彝族自治州",
		},
	},
	513401: {
		{
			Address: "西昌市",
		},
	},
	513421: {
		{
			Address: "西昌县",
			EndYear: 1985,
		},
	},
	513422: {
		{
			Address: "昭觉县",
			EndYear: 1981,
		},
		{
			Address:   "木里藏族自治县",
			StartYear: 1982,
		},
	},
	513423: {
		{
			Address: "甘洛县",
			EndYear: 1981,
		},
		{
			Address:   "盐源县",
			StartYear: 1982,
		},
	},
	513424: {
		{
			Address: "峨边县",
			EndYear: 1980,
		},
		{
			Address:   "德昌县",
			StartYear: 1982,
		},
	},
	513425: {
		{
			Address: "马边县",
			EndYear: 1980,
		},
		{
			Address:   "会理县",
			StartYear: 1982,
		},
	},
	513426: {
		{
			Address: "雷波县",
			EndYear: 1981,
		},
		{
			Address:   "会东县",
			StartYear: 1982,
		},
	},
	513427: {
		{
			Address: "宁南县",
		},
	},
	513428: {
		{
			Address: "会东县",
			EndYear: 1981,
		},
		{
			Address:   "普格县",
			StartYear: 1982,
		},
	},
	513429: {
		{
			Address: "会理县",
			EndYear: 1981,
		},
		{
			Address:   "布拖县",
			StartYear: 1982,
		},
	},
	513430: {
		{
			Address: "德昌县",
			EndYear: 1981,
		},
		{
			Address:   "金阳县",
			StartYear: 1982,
		},
	},
	513431: {
		{
			Address: "美姑县",
			EndYear: 1981,
		},
		{
			Address:   "昭觉县",
			StartYear: 1982,
		},
	},
	513432: {
		{
			Address: "金阳县",
			EndYear: 1981,
		},
		{
			Address:   "喜德县",
			StartYear: 1982,
		},
	},
	513433: {
		{
			Address: "布拖县",
			EndYear: 1981,
		},
		{
			Address:   "冕宁县",
			StartYear: 1982,
		},
	},
	513434: {
		{
			Address: "普格县",
			EndYear: 1981,
		},
		{
			Address:   "越西县",
			StartYear: 1982,
		},
	},
	513435: {
		{
			Address: "喜德县",
			EndYear: 1981,
		},
		{
			Address:   "甘洛县",
			StartYear: 1982,
		},
	},
	513436: {
		{
			Address: "越西县",
			EndYear: 1981,
		},
		{
			Address:   "美姑县",
			StartYear: 1982,
		},
	},
	513437: {
		{
			Address: "盐源县",
			EndYear: 1981,
		},
		{
			Address:   "雷波县",
			StartYear: 1982,
		},
	},
	513438: {
		{
			Address: "木里藏族自治县",
			EndYear: 1981,
		},
	},
	513439: {
		{
			Address: "冕宁县",
			EndYear: 1981,
		},
	},
	513500: {
		{
			Address:   "黔江地区",
			StartYear: 1988,
			EndYear:   1996,
		},
	},
	513521: {
		{
			Address:   "石柱土家族自治县",
			StartYear: 1988,
			EndYear:   1996,
		},
	},
	513522: {
		{
			Address:   "秀山土家族苗族自治县",
			StartYear: 1988,
			EndYear:   1996,
		},
	},
	513523: {
		{
			Address:   "黔江土家族苗族自治县",
			StartYear: 1988,
			EndYear:   1996,
		},
	},
	513524: {
		{
			Address:   "酉阳土家族苗族自治县",
			StartYear: 1988,
			EndYear:   1996,
		},
	},
	513525: {
		{
			Address:   "彭水苗族土家族自治县",
			StartYear: 1988,
			EndYear:   1996,
		},
	},
	513600: {
		{
			Address:   "广安地区",
			StartYear: 1993,
			EndYear:   1997,
		},
	},
	513601: {
		{
			Address:   "华蓥市",
			StartYear: 1993,
			EndYear:   1997,
		},
	},
	513621: {
		{
			Address:   "岳池县",
			StartYear: 1993,
			EndYear:   1997,
		},
	},
	513622: {
		{
			Address:   "广安县",
			StartYear: 1993,
			EndYear:   1997,
		},
	},
	513623: {
		{
			Address:   "武胜县",
			StartYear: 1993,
			EndYear:   1997,
		},
	},
	513624: {
		{
			Address:   "邻水县",
			StartYear: 1993,
			EndYear:   1997,
		},
	},
	513700: {
		{
			Address:   "巴中地区",
			StartYear: 1993,
			EndYear:   1999,
		},
	},
	513701: {
		{
			Address:   "巴中市",
			StartYear: 1993,
			EndYear:   1999,
		},
	},
	513721: {
		{
			Address:   "通江县",
			StartYear: 1993,
			EndYear:   1999,
		},
	},
	513722: {
		{
			Address:   "南江县",
			StartYear: 1993,
			EndYear:   1999,
		},
	},
	513723: {
		{
			Address:   "平昌县",
			StartYear: 1993,
			EndYear:   1999,
		},
	},
	513800: {
		{
			Address:   "眉山地区",
			StartYear: 1997,
			EndYear:   1999,
		},
	},
	513821: {
		{
			Address:   "眉山县",
			StartYear: 1997,
			EndYear:   1999,
		},
	},
	513822: {
		{
			Address:   "仁寿县",
			StartYear: 1997,
			EndYear:   1999,
		},
	},
	513823: {
		{
			Address:   "彭山县",
			StartYear: 1997,
			EndYear:   1999,
		},
	},
	513824: {
		{
			Address:   "洪雅县",
			StartYear: 1997,
			EndYear:   1999,
		},
	},
	513825: {
		{
			Address:   "丹棱县",
			StartYear: 1997,
			EndYear:   1999,
		},
	},
	513826: {
		{
			Address:   "青神县",
			StartYear: 1997,
			EndYear:   1999,
		},
	},
	513900: {
		{
			Address:   "资阳地区",
			StartYear: 1998,
			EndYear:   1999,
		},
	},
	513901: {
		{
			Address:   "资阳市",
			StartYear: 1998,
			EndYear:   1999,
		},
	},
	513902: {
		{
			Address:   "简阳市",
			StartYear: 1998,
			EndYear:   1999,
		},
	},
	513921: {
		{
			Address:   "安岳县",
			StartYear: 1998,
			EndYear:   1999,
		},
	},
	513922: {
		{
			Address:   "乐至县",
			StartYear: 1998,
			EndYear:   1999,
		},
	},
	517000: {
		{
			Address:   "涪陵市",
			StartYear: 1995,
			EndYear:   1996,
		},
	},
	517002: {
		{
			Address:   "枳城区",
			StartYear: 1995,
			EndYear:   1996,
		},
	},
	517003: {
		{
			Address:   "李渡区",
			StartYear: 1995,
			EndYear:   1996,
		},
	},
	517021: {
		{
			Address:   "垫江县",
			StartYear: 1995,
			EndYear:   1996,
		},
	},
	517022: {
		{
			Address:   "丰都县",
			StartYear: 1995,
			EndYear:   1996,
		},
	},
	517023: {
		{
			Address:   "武隆县",
			StartYear: 1995,
			EndYear:   1996,
		},
	},
	517081: {
		{
			Address:   "南川市",
			StartYear: 1995,
			EndYear:   1996,
		},
	},
	519001: {
		{
			Address:   "广汉市",
			StartYear: 1988,
			EndYear:   1994,
		},
	},
	519002: {
		{
			Address:   "江油市",
			StartYear: 1988,
			EndYear:   1994,
		},
	},
	519003: {
		{
			Address:   "都江堰市",
			StartYear: 1988,
			EndYear:   1994,
		},
	},
	519004: {
		{
			Address:   "峨眉山市",
			StartYear: 1988,
			EndYear:   1994,
		},
	},
	519005: {
		{
			Address:   "永川市",
			StartYear: 1992,
			EndYear:   1994,
		},
	},
	519006: {
		{
			Address:   "合川市",
			StartYear: 1992,
			EndYear:   1994,
		},
	},
	519007: {
		{
			Address:   "江津市",
			StartYear: 1992,
			EndYear:   1994,
		},
	},
	519008: {
		{
			Address:   "阆中市",
			StartYear: 1993,
			EndYear:   1994,
		},
	},
	519009: {
		{
			Address:   "资阳市",
			StartYear: 1993,
			EndYear:   1994,
		},
	},
	519010: {
		{
			Address:   "彭州市",
			StartYear: 1993,
			EndYear:   1994,
		},
	},
	519011: {
		{
			Address:   "简阳市",
			StartYear: 1994,
			EndYear:   1994,
		},
	},
	519012: {
		{
			Address:   "邛崃市",
			StartYear: 1994,
			EndYear:   1994,
		},
	},
	519013: {
		{
			Address:   "崇州市",
			StartYear: 1994,
			EndYear:   1994,
		},
	},
	520000: {
		{
			Address: "贵州省",
		},
	},
	520100: {
		{
			Address: "贵阳市",
		},
	},
	520102: {
		{
			Address: "南明区",
		},
	},
	520103: {
		{
			Address: "云岩区",
		},
	},
	520111: {
		{
			Address: "花溪区",
		},
	},
	520112: {
		{
			Address: "乌当区",
		},
	},
	520113: {
		{
			Address: "白云区",
		},
	},
	520114: {
		{
			Address:   "小河区",
			StartYear: 2000,
			EndYear:   2011,
		},
	},
	520115: {
		{
			Address:   "观山湖区",
			StartYear: 2012,
		},
	},
	520121: {
		{
			Address:   "开阳县",
			StartYear: 1995,
		},
	},
	520122: {
		{
			Address:   "息烽县",
			StartYear: 1995,
		},
	},
	520123: {
		{
			Address:   "修文县",
			StartYear: 1995,
		},
	},
	520181: {
		{
			Address:   "清镇市",
			StartYear: 1995,
		},
	},
	520200: {
		{
			Address: "六盘水市",
		},
	},
	520201: {
		{
			Address: "水城特区",
			EndYear: 1986,
		},
		{
			Address:   "钟山区",
			StartYear: 1987,
		},
	},
	520202: {
		{
			Address: "盘县特区",
			EndYear: 1998,
		},
	},
	520203: {
		{
			Address: "六枝特区",
		},
	},
	520204: {
		{
			Address:   "水城区",
			StartYear: 2020,
		},
	},
	520221: {
		{
			Address:   "水城县",
			StartYear: 1987,
			EndYear:   2019,
		},
	},
	520222: {
		{
			Address:   "盘县",
			StartYear: 1999,
			EndYear:   2016,
		},
	},
	520281: {
		{
			Address:   "盘州市",
			StartYear: 2017,
		},
	},
	520300: {
		{
			Address:   "遵义市",
			StartYear: 1997,
		},
	},
	520302: {
		{
			Address:   "红花岗区",
			StartYear: 1997,
		},
	},
	520303: {
		{
			Address:   "汇川区",
			StartYear: 2003,
		},
	},
	520304: {
		{
			Address:   "播州区",
			StartYear: 2016,
		},
	},
	520321: {
		{
			Address:   "遵义县",
			StartYear: 1997,
			EndYear:   2015,
		},
	},
	520322: {
		{
			Address:   "桐梓县",
			StartYear: 1997,
		},
	},
	520323: {
		{
			Address:   "绥阳县",
			StartYear: 1997,
		},
	},
	520324: {
		{
			Address:   "正安县",
			StartYear: 1997,
		},
	},
	520325: {
		{
			Address:   "道真仡佬族苗族自治县",
			StartYear: 1997,
		},
	},
	520326: {
		{
			Address:   "务川仡佬族苗族自治县",
			StartYear: 1997,
		},
	},
	520327: {
		{
			Address:   "凤冈县",
			StartYear: 1997,
		},
	},
	520328: {
		{
			Address:   "湄潭县",
			StartYear: 1997,
		},
	},
	520329: {
		{
			Address:   "余庆县",
			StartYear: 1997,
		},
	},
	520330: {
		{
			Address:   "习水县",
			StartYear: 1997,
		},
	},
	520381: {
		{
			Address:   "赤水市",
			StartYear: 1997,
		},
	},
	520382: {
		{
			Address:   "仁怀市",
			StartYear: 1997,
		},
	},
	520400: {
		{
			Address:   "安顺市",
			StartYear: 2000,
		},
	},
	520402: {
		{
			Address:   "西秀区",
			StartYear: 2000,
		},
	},
	520403: {
		{
			Address:   "平坝区",
			StartYear: 2014,
		},
	},
	520421: {
		{
			Address:   "平坝县",
			StartYear: 2000,
			EndYear:   2013,
		},
	},
	520422: {
		{
			Address:   "普定县",
			StartYear: 2000,
		},
	},
	520423: {
		{
			Address:   "镇宁布依族苗族自治县",
			StartYear: 2000,
		},
	},
	520424: {
		{
			Address:   "关岭布依族苗族自治县",
			StartYear: 2000,
		},
	},
	520425: {
		{
			Address:   "紫云苗族布依族自治县",
			StartYear: 2000,
		},
	},
	520500: {
		{
			Address:   "毕节市",
			StartYear: 2011,
		},
	},
	520502: {
		{
			Address:   "七星关区",
			StartYear: 2011,
		},
	},
	520521: {
		{
			Address:   "大方县",
			StartYear: 2011,
		},
	},
	520522: {
		{
			Address:   "黔西县",
			StartYear: 2011,
		},
	},
	520523: {
		{
			Address:   "金沙县",
			StartYear: 2011,
		},
	},
	520524: {
		{
			Address:   "织金县",
			StartYear: 2011,
		},
	},
	520525: {
		{
			Address:   "纳雍县",
			StartYear: 2011,
		},
	},
	520526: {
		{
			Address:   "威宁彝族回族苗族自治县",
			StartYear: 2011,
		},
	},
	520527: {
		{
			Address:   "赫章县",
			StartYear: 2011,
		},
	},
	520600: {
		{
			Address:   "铜仁市",
			StartYear: 2011,
		},
	},
	520602: {
		{
			Address:   "碧江区",
			StartYear: 2011,
		},
	},
	520603: {
		{
			Address:   "万山区",
			StartYear: 2011,
		},
	},
	520621: {
		{
			Address:   "江口县",
			StartYear: 2011,
		},
	},
	520622: {
		{
			Address:   "玉屏侗族自治县",
			StartYear: 2011,
		},
	},
	520623: {
		{
			Address:   "石阡县",
			StartYear: 2011,
		},
	},
	520624: {
		{
			Address:   "思南县",
			StartYear: 2011,
		},
	},
	520625: {
		{
			Address:   "印江土家族苗族自治县",
			StartYear: 2011,
		},
	},
	520626: {
		{
			Address:   "德江县",
			StartYear: 2011,
		},
	},
	520627: {
		{
			Address:   "沿河土家族自治县",
			StartYear: 2011,
		},
	},
	520628: {
		{
			Address:   "松桃苗族自治县",
			StartYear: 2011,
		},
	},
	522100: {
		{
			Address: "遵义地区",
			EndYear: 1996,
		},
	},
	522101: {
		{
			Address: "遵义市",
			EndYear: 1996,
		},
	},
	522102: {
		{
			Address:   "赤水市",
			StartYear: 1990,
			EndYear:   1996,
		},
	},
	522103: {
		{
			Address:   "仁怀市",
			StartYear: 1995,
			EndYear:   1996,
		},
	},
	522121: {
		{
			Address: "遵义县",
			EndYear: 1996,
		},
	},
	522122: {
		{
			Address: "桐梓县",
			EndYear: 1996,
		},
	},
	522123: {
		{
			Address: "绥阳县",
			EndYear: 1996,
		},
	},
	522124: {
		{
			Address: "正安县",
			EndYear: 1996,
		},
	},
	522125: {
		{
			Address: "道真县",
			EndYear: 1985,
		},
		{
			Address:   "道真仡佬族苗族自治县",
			StartYear: 1986,
			EndYear:   1996,
		},
	},
	522126: {
		{
			Address: "务川县",
			EndYear: 1985,
		},
		{
			Address:   "务川仡佬族苗族自治县",
			StartYear: 1986,
			EndYear:   1996,
		},
	},
	522127: {
		{
			Address: "凤冈县",
			EndYear: 1996,
		},
	},
	522128: {
		{
			Address: "湄潭县",
			EndYear: 1996,
		},
	},
	522129: {
		{
			Address: "余庆县",
			EndYear: 1996,
		},
	},
	522130: {
		{
			Address: "仁怀县",
			EndYear: 1994,
		},
	},
	522131: {
		{
			Address: "赤水县",
			EndYear: 1989,
		},
	},
	522132: {
		{
			Address: "习水县",
			EndYear: 1996,
		},
	},
	522200: {
		{
			Address: "铜仁地区",
			EndYear: 2010,
		},
	},
	522201: {
		{
			Address:   "铜仁市",
			StartYear: 1987,
			EndYear:   2010,
		},
	},
	522221: {
		{
			Address: "铜仁县",
			EndYear: 1986,
		},
	},
	522222: {
		{
			Address: "江口县",
			EndYear: 2010,
		},
	},
	522223: {
		{
			Address: "玉屏县",
			EndYear: 1982,
		},
		{
			Address:   "玉屏侗族自治县",
			StartYear: 1983,
			EndYear:   2010,
		},
	},
	522224: {
		{
			Address: "石阡县",
			EndYear: 2010,
		},
	},
	522225: {
		{
			Address: "思南县",
			EndYear: 2010,
		},
	},
	522226: {
		{
			Address: "印江县",
			EndYear: 1985,
		},
		{
			Address:   "印江土家族苗族自治县",
			StartYear: 1986,
			EndYear:   2010,
		},
	},
	522227: {
		{
			Address: "德江县",
			EndYear: 2010,
		},
	},
	522228: {
		{
			Address: "沿河县",
			EndYear: 1985,
		},
		{
			Address:   "沿河土家族自治县",
			StartYear: 1986,
			EndYear:   2010,
		},
	},
	522229: {
		{
			Address: "松桃苗族自治县",
			EndYear: 2010,
		},
	},
	522230: {
		{
			Address: "万山特区",
			EndYear: 2010,
		},
	},
	522300: {
		{
			Address: "兴义地区",
			EndYear: 1980,
		},
		{
			Address:   "黔西南布依族苗族自治州",
			StartYear: 1981,
		},
	},
	522301: {
		{
			Address:   "兴义市",
			StartYear: 1987,
		},
	},
	522302: {
		{
			Address:   "兴仁市",
			StartYear: 2018,
		},
	},
	522321: {
		{
			Address: "兴义县",
			EndYear: 1986,
		},
	},
	522322: {
		{
			Address: "兴仁县",
			EndYear: 2017,
		},
	},
	522323: {
		{
			Address: "普安县",
		},
	},
	522324: {
		{
			Address: "晴隆县",
		},
	},
	522325: {
		{
			Address: "贞丰布依族苗族自治县",
			EndYear: 1980,
		},
		{
			Address:   "贞丰县",
			StartYear: 1981,
		},
	},
	522326: {
		{
			Address: "望谟布依族苗族自治县",
			EndYear: 1980,
		},
		{
			Address:   "望谟县",
			StartYear: 1981,
		},
	},
	522327: {
		{
			Address: "册亨布依族自治县",
			EndYear: 1980,
		},
		{
			Address:   "册亨县",
			StartYear: 1981,
		},
	},
	522328: {
		{
			Address: "安龙布依族苗族自治县",
			EndYear: 1980,
		},
		{
			Address:   "安龙县",
			StartYear: 1981,
		},
	},
	522400: {
		{
			Address: "毕节地区",
			EndYear: 2010,
		},
	},
	522401: {
		{
			Address:   "毕节市",
			StartYear: 1993,
			EndYear:   2010,
		},
	},
	522421: {
		{
			Address: "毕节县",
			EndYear: 1992,
		},
	},
	522422: {
		{
			Address: "大方县",
			EndYear: 2010,
		},
	},
	522423: {
		{
			Address: "黔西县",
			EndYear: 2010,
		},
	},
	522424: {
		{
			Address: "金沙县",
			EndYear: 2010,
		},
	},
	522425: {
		{
			Address: "织金县",
			EndYear: 2010,
		},
	},
	522426: {
		{
			Address: "纳雍县",
			EndYear: 2010,
		},
	},
	522427: {
		{
			Address: "威宁彝族回族苗族自治县",
			EndYear: 2010,
		},
	},
	522428: {
		{
			Address: "赫章县",
			EndYear: 2010,
		},
	},
	522500: {
		{
			Address: "安顺地区",
			EndYear: 1999,
		},
	},
	522501: {
		{
			Address: "安顺市",
			EndYear: 1999,
		},
	},
	522502: {
		{
			Address:   "清镇市",
			StartYear: 1992,
			EndYear:   1994,
		},
	},
	522521: {
		{
			Address: "安顺县",
			EndYear: 1989,
		},
	},
	522522: {
		{
			Address: "开阳县",
			EndYear: 1994,
		},
	},
	522523: {
		{
			Address: "息烽县",
			EndYear: 1994,
		},
	},
	522524: {
		{
			Address: "修文县",
			EndYear: 1994,
		},
	},
	522525: {
		{
			Address: "清镇县",
			EndYear: 1991,
		},
	},
	522526: {
		{
			Address: "平坝县",
			EndYear: 1999,
		},
	},
	522527: {
		{
			Address: "普定县",
			EndYear: 1999,
		},
	},
	522528: {
		{
			Address: "关岭县",
			EndYear: 1980,
		},
		{
			Address:   "关岭布依族苗族自治县",
			StartYear: 1981,
			EndYear:   1999,
		},
	},
	522529: {
		{
			Address: "镇宁布依族苗族自治县",
			EndYear: 1999,
		},
	},
	522530: {
		{
			Address: "紫云苗族布依族自治县",
			EndYear: 1999,
		},
	},
	522600: {
		{
			Address: "黔东南苗族侗族自治州",
		},
	},
	522601: {
		{
			Address:   "凯里市",
			StartYear: 1983,
		},
	},
	522621: {
		{
			Address: "凯里县",
			EndYear: 1982,
		},
	},
	522622: {
		{
			Address: "黄平县",
		},
	},
	522623: {
		{
			Address: "施秉县",
		},
	},
	522624: {
		{
			Address: "三穗县",
		},
	},
	522625: {
		{
			Address: "镇远县",
		},
	},
	522626: {
		{
			Address: "岑巩县",
		},
	},
	522627: {
		{
			Address: "天柱县",
		},
	},
	522628: {
		{
			Address: "锦屏县",
		},
	},
	522629: {
		{
			Address: "剑河县",
		},
	},
	522630: {
		{
			Address: "台江县",
		},
	},
	522631: {
		{
			Address: "黎平县",
		},
	},
	522632: {
		{
			Address: "榕江县",
		},
	},
	522633: {
		{
			Address: "从江县",
		},
	},
	522634: {
		{
			Address: "雷山县",
		},
	},
	522635: {
		{
			Address: "麻江县",
		},
	},
	522636: {
		{
			Address: "丹寨县",
		},
	},
	522700: {
		{
			Address: "黔南布依族苗族自治州",
		},
	},
	522701: {
		{
			Address: "都匀市",
		},
	},
	522702: {
		{
			Address:   "福泉市",
			StartYear: 1996,
		},
	},
	522721: {
		{
			Address: "都匀县",
			EndYear: 1982,
		},
	},
	522722: {
		{
			Address: "荔波县",
		},
	},
	522723: {
		{
			Address: "贵定县",
		},
	},
	522724: {
		{
			Address: "福泉县",
			EndYear: 1995,
		},
	},
	522725: {
		{
			Address: "瓮安县",
		},
	},
	522726: {
		{
			Address: "独山县",
		},
	},
	522727: {
		{
			Address: "平塘县",
		},
	},
	522728: {
		{
			Address: "罗甸县",
		},
	},
	522729: {
		{
			Address: "长顺县",
		},
	},
	522730: {
		{
			Address: "龙里县",
		},
	},
	522731: {
		{
			Address: "惠水县",
		},
	},
	522732: {
		{
			Address: "三都水族自治县",
		},
	},
	530000: {
		{
			Address: "云南省",
		},
	},
	530100: {
		{
			Address: "昆明市",
		},
	},
	530102: {
		{
			Address: "五华区",
		},
	},
	530103: {
		{
			Address: "盘龙区",
		},
	},
	530111: {
		{
			Address: "官渡区",
		},
	},
	530112: {
		{
			Address: "西山区",
		},
	},
	530113: {
		{
			Address:   "东川区",
			StartYear: 1998,
		},
	},
	530114: {
		{
			Address:   "呈贡区",
			StartYear: 2011,
		},
	},
	530115: {
		{
			Address:   "晋宁区",
			StartYear: 2016,
		},
	},
	530121: {
		{
			Address: "安宁县",
			EndYear: 1981,
		},
		{
			Address:   "呈贡县",
			StartYear: 1982,
			EndYear:   2010,
		},
	},
	530122: {
		{
			Address: "呈贡县",
			EndYear: 1981,
		},
		{
			Address:   "晋宁县",
			StartYear: 1982,
			EndYear:   2015,
		},
	},
	530123: {
		{
			Address: "富民县",
			EndYear: 1981,
		},
		{
			Address:   "安宁县",
			StartYear: 1982,
			EndYear:   1994,
		},
	},
	530124: {
		{
			Address: "晋宁县",
			EndYear: 1981,
		},
		{
			Address:   "富民县",
			StartYear: 1982,
		},
	},
	530125: {
		{
			Address:   "宜良县",
			StartYear: 1983,
		},
	},
	530126: {
		{
			Address:   "路南彝族自治县",
			StartYear: 1983,
			EndYear:   1997,
		},
		{
			Address:   "石林彝族自治县",
			StartYear: 1998,
		},
	},
	530127: {
		{
			Address:   "嵩明县",
			StartYear: 1983,
		},
	},
	530128: {
		{
			Address:   "禄劝县",
			StartYear: 1983,
			EndYear:   1984,
		},
		{
			Address:   "禄劝彝族苗族自治县",
			StartYear: 1985,
		},
	},
	530129: {
		{
			Address:   "寻甸回族彝族自治县",
			StartYear: 1998,
		},
	},
	530181: {
		{
			Address:   "安宁市",
			StartYear: 1995,
		},
	},
	530200: {
		{
			Address: "东川市",
			EndYear: 1997,
		},
	},
	530300: {
		{
			Address:   "曲靖市",
			StartYear: 1997,
		},
	},
	530302: {
		{
			Address:   "麒麟区",
			StartYear: 1997,
		},
	},
	530303: {
		{
			Address:   "沾益区",
			StartYear: 2016,
		},
	},
	530304: {
		{
			Address:   "马龙区",
			StartYear: 2018,
		},
	},
	530321: {
		{
			Address:   "马龙县",
			StartYear: 1997,
			EndYear:   2017,
		},
	},
	530322: {
		{
			Address:   "陆良县",
			StartYear: 1997,
		},
	},
	530323: {
		{
			Address:   "师宗县",
			StartYear: 1997,
		},
	},
	530324: {
		{
			Address:   "罗平县",
			StartYear: 1997,
		},
	},
	530325: {
		{
			Address:   "富源县",
			StartYear: 1997,
		},
	},
	530326: {
		{
			Address:   "会泽县",
			StartYear: 1997,
		},
	},
	530327: {
		{
			Address:   "寻甸回族彝族自治县",
			StartYear: 1997,
			EndYear:   1997,
		},
	},
	530328: {
		{
			Address:   "沾益县",
			StartYear: 1997,
			EndYear:   2015,
		},
	},
	530381: {
		{
			Address:   "宣威市",
			StartYear: 1997,
		},
	},
	530400: {
		{
			Address:   "玉溪市",
			StartYear: 1997,
		},
	},
	530402: {
		{
			Address:   "红塔区",
			StartYear: 1997,
		},
	},
	530403: {
		{
			Address:   "江川区",
			StartYear: 2015,
		},
	},
	530421: {
		{
			Address:   "江川县",
			StartYear: 1997,
			EndYear:   2014,
		},
	},
	530422: {
		{
			Address:   "澄江县",
			StartYear: 1997,
			EndYear:   2018,
		},
	},
	530423: {
		{
			Address:   "通海县",
			StartYear: 1997,
		},
	},
	530424: {
		{
			Address:   "华宁县",
			StartYear: 1997,
		},
	},
	530425: {
		{
			Address:   "易门县",
			StartYear: 1997,
		},
	},
	530426: {
		{
			Address:   "峨山彝族自治县",
			StartYear: 1997,
		},
	},
	530427: {
		{
			Address:   "新平彝族傣族自治县",
			StartYear: 1997,
		},
	},
	530428: {
		{
			Address:   "元江哈尼族彝族傣族自治县",
			StartYear: 1997,
		},
	},
	530481: {
		{
			Address:   "澄江市",
			StartYear: 2019,
		},
	},
	530500: {
		{
			Address:   "保山市",
			StartYear: 2000,
		},
	},
	530502: {
		{
			Address:   "隆阳区",
			StartYear: 2000,
		},
	},
	530521: {
		{
			Address:   "施甸县",
			StartYear: 2000,
		},
	},
	530522: {
		{
			Address:   "腾冲县",
			StartYear: 2000,
			EndYear:   2014,
		},
	},
	530523: {
		{
			Address:   "龙陵县",
			StartYear: 2000,
		},
	},
	530524: {
		{
			Address:   "昌宁县",
			StartYear: 2000,
		},
	},
	530581: {
		{
			Address:   "腾冲市",
			StartYear: 2015,
		},
	},
	530600: {
		{
			Address:   "昭通市",
			StartYear: 2001,
		},
	},
	530602: {
		{
			Address:   "昭阳区",
			StartYear: 2001,
		},
	},
	530621: {
		{
			Address:   "鲁甸县",
			StartYear: 2001,
		},
	},
	530622: {
		{
			Address:   "巧家县",
			StartYear: 2001,
		},
	},
	530623: {
		{
			Address:   "盐津县",
			StartYear: 2001,
		},
	},
	530624: {
		{
			Address:   "大关县",
			StartYear: 2001,
		},
	},
	530625: {
		{
			Address:   "永善县",
			StartYear: 2001,
		},
	},
	530626: {
		{
			Address:   "绥江县",
			StartYear: 2001,
		},
	},
	530627: {
		{
			Address:   "镇雄县",
			StartYear: 2001,
		},
	},
	530628: {
		{
			Address:   "彝良县",
			StartYear: 2001,
		},
	},
	530629: {
		{
			Address:   "威信县",
			StartYear: 2001,
		},
	},
	530630: {
		{
			Address:   "水富县",
			StartYear: 2001,
			EndYear:   2017,
		},
	},
	530681: {
		{
			Address:   "水富市",
			StartYear: 2018,
		},
	},
	530700: {
		{
			Address:   "丽江市",
			StartYear: 2002,
		},
	},
	530702: {
		{
			Address:   "古城区",
			StartYear: 2002,
		},
	},
	530721: {
		{
			Address:   "玉龙纳西族自治县",
			StartYear: 2002,
		},
	},
	530722: {
		{
			Address:   "永胜县",
			StartYear: 2002,
		},
	},
	530723: {
		{
			Address:   "华坪县",
			StartYear: 2002,
		},
	},
	530724: {
		{
			Address:   "宁蒗彝族自治县",
			StartYear: 2002,
		},
	},
	530800: {
		{
			Address:   "思茅市",
			StartYear: 2003,
			EndYear:   2006,
		},
		{
			Address:   "普洱市",
			StartYear: 2007,
		},
	},
	530802: {
		{
			Address:   "翠云区",
			StartYear: 2003,
			EndYear:   2006,
		},
		{
			Address:   "思茅区",
			StartYear: 2007,
		},
	},
	530821: {
		{
			Address:   "普洱哈尼族彝族自治县",
			StartYear: 2003,
			EndYear:   2006,
		},
		{
			Address:   "宁洱哈尼族彝族自治县",
			StartYear: 2007,
		},
	},
	530822: {
		{
			Address:   "墨江哈尼族自治县",
			StartYear: 2003,
		},
	},
	530823: {
		{
			Address:   "景东彝族自治县",
			StartYear: 2003,
		},
	},
	530824: {
		{
			Address:   "景谷傣族彝族自治县",
			StartYear: 2003,
		},
	},
	530825: {
		{
			Address:   "镇沅彝族哈尼族拉祜族自治县",
			StartYear: 2003,
		},
	},
	530826: {
		{
			Address:   "江城哈尼族彝族自治县",
			StartYear: 2003,
		},
	},
	530827: {
		{
			Address:   "孟连傣族拉祜族佤族自治县",
			StartYear: 2003,
		},
	},
	530828: {
		{
			Address:   "澜沧拉祜族自治县",
			StartYear: 2003,
		},
	},
	530829: {
		{
			Address:   "西盟佤族自治县",
			StartYear: 2003,
		},
	},
	530900: {
		{
			Address:   "临沧市",
			StartYear: 2003,
		},
	},
	530902: {
		{
			Address:   "临翔区",
			StartYear: 2003,
		},
	},
	530921: {
		{
			Address:   "凤庆县",
			StartYear: 2003,
		},
	},
	530922: {
		{
			Address:   "云县",
			StartYear: 2003,
		},
	},
	530923: {
		{
			Address:   "永德县",
			StartYear: 2003,
		},
	},
	530924: {
		{
			Address:   "镇康县",
			StartYear: 2003,
		},
	},
	530925: {
		{
			Address:   "双江拉祜族佤族布朗族傣族自治县",
			StartYear: 2003,
		},
	},
	530926: {
		{
			Address:   "耿马傣族佤族自治县",
			StartYear: 2003,
		},
	},
	530927: {
		{
			Address:   "沧源佤族自治县",
			StartYear: 2003,
		},
	},
	532100: {
		{
			Address: "昭通地区",
			EndYear: 2000,
		},
	},
	532101: {
		{
			Address:   "昭通市",
			StartYear: 1981,
			EndYear:   2000,
		},
	},
	532121: {
		{
			Address: "昭通县",
			EndYear: 1982,
		},
	},
	532122: {
		{
			Address: "鲁甸县",
			EndYear: 2000,
		},
	},
	532123: {
		{
			Address: "巧家县",
			EndYear: 2000,
		},
	},
	532124: {
		{
			Address: "盐津县",
			EndYear: 2000,
		},
	},
	532125: {
		{
			Address: "大关县",
			EndYear: 2000,
		},
	},
	532126: {
		{
			Address: "永善县",
			EndYear: 2000,
		},
	},
	532127: {
		{
			Address: "绥江县",
			EndYear: 2000,
		},
	},
	532128: {
		{
			Address: "镇雄县",
			EndYear: 2000,
		},
	},
	532129: {
		{
			Address: "彝良县",
			EndYear: 2000,
		},
	},
	532130: {
		{
			Address: "威信县",
			EndYear: 2000,
		},
	},
	532131: {
		{
			Address:   "水富县",
			StartYear: 1981,
			EndYear:   2000,
		},
	},
	532200: {
		{
			Address: "曲靖地区",
			EndYear: 1996,
		},
	},
	532201: {
		{
			Address:   "曲靖市",
			StartYear: 1983,
			EndYear:   1996,
		},
	},
	532202: {
		{
			Address:   "宣威市",
			StartYear: 1994,
			EndYear:   1996,
		},
	},
	532221: {
		{
			Address: "曲靖县",
			EndYear: 1982,
		},
	},
	532222: {
		{
			Address: "沾益县",
			EndYear: 1982,
		},
	},
	532223: {
		{
			Address: "马龙县",
			EndYear: 1996,
		},
	},
	532224: {
		{
			Address: "宣威县",
			EndYear: 1993,
		},
	},
	532225: {
		{
			Address: "富源县",
			EndYear: 1996,
		},
	},
	532226: {
		{
			Address: "罗平县",
			EndYear: 1996,
		},
	},
	532227: {
		{
			Address: "师宗县",
			EndYear: 1996,
		},
	},
	532228: {
		{
			Address: "陆良县",
			EndYear: 1996,
		},
	},
	532229: {
		{
			Address: "宜良县",
			EndYear: 1982,
		},
	},
	532230: {
		{
			Address: "路南彝族自治县",
			EndYear: 1982,
		},
	},
	532231: {
		{
			Address: "寻甸回族彝族自治县",
			EndYear: 1996,
		},
	},
	532232: {
		{
			Address: "嵩明县",
			EndYear: 1982,
		},
	},
	532233: {
		{
			Address: "会泽县",
			EndYear: 1996,
		},
	},
	532300: {
		{
			Address: "楚雄彝族自治州",
		},
	},
	532301: {
		{
			Address:   "楚雄市",
			StartYear: 1983,
		},
	},
	532321: {
		{
			Address: "楚雄县",
			EndYear: 1982,
		},
	},
	532322: {
		{
			Address: "双柏县",
		},
	},
	532323: {
		{
			Address: "牟定县",
		},
	},
	532324: {
		{
			Address: "南华县",
		},
	},
	532325: {
		{
			Address: "姚安县",
		},
	},
	532326: {
		{
			Address: "大姚县",
		},
	},
	532327: {
		{
			Address: "永仁县",
		},
	},
	532328: {
		{
			Address: "元谋县",
		},
	},
	532329: {
		{
			Address: "武定县",
		},
	},
	532330: {
		{
			Address: "禄劝县",
			EndYear: 1982,
		},
	},
	532331: {
		{
			Address: "禄丰县",
		},
	},
	532400: {
		{
			Address: "玉溪地区",
			EndYear: 1996,
		},
	},
	532401: {
		{
			Address:   "玉溪市",
			StartYear: 1983,
			EndYear:   1996,
		},
	},
	532421: {
		{
			Address: "玉溪县",
			EndYear: 1982,
		},
	},
	532422: {
		{
			Address: "江川县",
			EndYear: 1996,
		},
	},
	532423: {
		{
			Address: "澄江县",
			EndYear: 1996,
		},
	},
	532424: {
		{
			Address: "通海县",
			EndYear: 1996,
		},
	},
	532425: {
		{
			Address: "华宁县",
			EndYear: 1996,
		},
	},
	532426: {
		{
			Address: "易门县",
			EndYear: 1996,
		},
	},
	532427: {
		{
			Address: "峨山彝族自治县",
			EndYear: 1996,
		},
	},
	532428: {
		{
			Address: "新平彝族傣族自治县",
			EndYear: 1996,
		},
	},
	532429: {
		{
			Address: "元江哈尼族彝族傣族自治县",
			EndYear: 1996,
		},
	},
	532500: {
		{
			Address: "红河哈尼族彝族自治州",
		},
	},
	532501: {
		{
			Address: "个旧市",
		},
	},
	532502: {
		{
			Address:   "开远市",
			StartYear: 1981,
		},
	},
	532503: {
		{
			Address:   "蒙自市",
			StartYear: 2010,
		},
	},
	532504: {
		{
			Address:   "弥勒市",
			StartYear: 2013,
		},
	},
	532521: {
		{
			Address: "开远县",
			EndYear: 1980,
		},
	},
	532522: {
		{
			Address: "蒙自县",
			EndYear: 2009,
		},
	},
	532523: {
		{
			Address: "屏边苗族自治县",
		},
	},
	532524: {
		{
			Address: "建水县",
		},
	},
	532525: {
		{
			Address: "石屏县",
		},
	},
	532526: {
		{
			Address: "弥勒县",
			EndYear: 2012,
		},
	},
	532527: {
		{
			Address: "泸西县",
		},
	},
	532528: {
		{
			Address: "元阳县",
		},
	},
	532529: {
		{
			Address: "红河县",
		},
	},
	532530: {
		{
			Address: "金平县",
			EndYear: 1984,
		},
		{
			Address:   "金平苗族瑶族傣族自治县",
			StartYear: 1985,
		},
	},
	532531: {
		{
			Address: "绿春县",
		},
	},
	532532: {
		{
			Address: "河口瑶族自治县",
		},
	},
	532600: {
		{
			Address: "文山壮族苗族自治州",
		},
	},
	532601: {
		{
			Address:   "文山市",
			StartYear: 2010,
		},
	},
	532621: {
		{
			Address: "文山县",
			EndYear: 2009,
		},
	},
	532622: {
		{
			Address: "砚山县",
		},
	},
	532623: {
		{
			Address: "西畴县",
		},
	},
	532624: {
		{
			Address: "麻栗坡县",
		},
	},
	532625: {
		{
			Address: "马关县",
		},
	},
	532626: {
		{
			Address: "丘北县",
		},
	},
	532627: {
		{
			Address: "广南县",
		},
	},
	532628: {
		{
			Address: "富宁县",
		},
	},
	532700: {
		{
			Address: "思茅地区",
			EndYear: 2002,
		},
	},
	532701: {
		{
			Address:   "思茅市",
			StartYear: 1993,
			EndYear:   2002,
		},
	},
	532721: {
		{
			Address:   "思茅县",
			StartYear: 1981,
			EndYear:   1992,
		},
	},
	532722: {
		{
			Address: "普洱县",
			EndYear: 1984,
		},
		{
			Address:   "普洱哈尼族彝族自治县",
			StartYear: 1985,
			EndYear:   2002,
		},
	},
	532723: {
		{
			Address: "墨江哈尼族自治县",
			EndYear: 2002,
		},
	},
	532724: {
		{
			Address: "景东县",
			EndYear: 1984,
		},
		{
			Address:   "景东彝族自治县",
			StartYear: 1985,
			EndYear:   2002,
		},
	},
	532725: {
		{
			Address: "景谷县",
			EndYear: 1984,
		},
		{
			Address:   "景谷傣族彝族自治县",
			StartYear: 1985,
			EndYear:   2002,
		},
	},
	532726: {
		{
			Address: "镇沅县",
			EndYear: 1989,
		},
		{
			Address:   "镇沅彝族哈尼族拉祜族自治县",
			StartYear: 1990,
			EndYear:   2002,
		},
	},
	532727: {
		{
			Address: "江城哈尼族彝族自治县",
			EndYear: 2002,
		},
	},
	532728: {
		{
			Address: "孟连傣族拉祜族佤族自治县",
			EndYear: 2002,
		},
	},
	532729: {
		{
			Address: "澜沧拉祜族自治县",
			EndYear: 2002,
		},
	},
	532730: {
		{
			Address: "西盟佤族自治县",
			EndYear: 2002,
		},
	},
	532800: {
		{
			Address: "西双版纳傣族自治州",
		},
	},
	532801: {
		{
			Address:   "景洪市",
			StartYear: 1993,
		},
	},
	532821: {
		{
			Address: "景洪县",
			EndYear: 1992,
		},
	},
	532822: {
		{
			Address: "勐海县",
		},
	},
	532823: {
		{
			Address: "勐腊县",
		},
	},
	532900: {
		{
			Address: "大理白族自治州",
		},
	},
	532901: {
		{
			Address: "下关市",
			EndYear: 1982,
		},
		{
			Address:   "大理市",
			StartYear: 1983,
		},
	},
	532921: {
		{
			Address: "大理县",
			EndYear: 1982,
		},
	},
	532922: {
		{
			Address: "漾濞县",
			EndYear: 1984,
		},
		{
			Address:   "漾濞彝族自治县",
			StartYear: 1985,
		},
	},
	532923: {
		{
			Address: "祥云县",
		},
	},
	532924: {
		{
			Address: "宾川县",
		},
	},
	532925: {
		{
			Address: "弥渡县",
		},
	},
	532926: {
		{
			Address: "南涧彝族自治县",
		},
	},
	532927: {
		{
			Address: "巍山彝族回族自治县",
		},
	},
	532928: {
		{
			Address: "永平县",
		},
	},
	532929: {
		{
			Address: "云龙县",
		},
	},
	532930: {
		{
			Address: "洱源县",
		},
	},
	532931: {
		{
			Address: "剑川县",
		},
	},
	532932: {
		{
			Address: "鹤庆县",
		},
	},
	533000: {
		{
			Address: "保山地区",
			EndYear: 1999,
		},
	},
	533001: {
		{
			Address:   "保山市",
			StartYear: 1983,
			EndYear:   1999,
		},
	},
	533021: {
		{
			Address: "保山县",
			EndYear: 1982,
		},
	},
	533022: {
		{
			Address: "施甸县",
			EndYear: 1999,
		},
	},
	533023: {
		{
			Address: "腾冲县",
			EndYear: 1999,
		},
	},
	533024: {
		{
			Address: "龙陵县",
			EndYear: 1999,
		},
	},
	533025: {
		{
			Address: "昌宁县",
			EndYear: 1999,
		},
	},
	533100: {
		{
			Address: "德宏傣族景颇族自治州",
		},
	},
	533101: {
		{
			Address:   "畹町市",
			StartYear: 1985,
			EndYear:   1998,
		},
	},
	533102: {
		{
			Address:   "瑞丽市",
			StartYear: 1992,
		},
	},
	533103: {
		{
			Address:   "潞西市",
			StartYear: 1996,
			EndYear:   2009,
		},
		{
			Address:   "芒市",
			StartYear: 2010,
		},
	},
	533121: {
		{
			Address: "潞西县",
			EndYear: 1995,
		},
	},
	533122: {
		{
			Address: "梁河县",
		},
	},
	533123: {
		{
			Address: "盈江县",
		},
	},
	533124: {
		{
			Address: "陇川县",
		},
	},
	533125: {
		{
			Address: "瑞丽县",
			EndYear: 1991,
		},
	},
	533126: {
		{
			Address: "畹町镇",
			EndYear: 1984,
		},
	},
	533200: {
		{
			Address: "丽江地区",
			EndYear: 2001,
		},
	},
	533221: {
		{
			Address: "丽江纳西族自治县",
			EndYear: 2001,
		},
	},
	533222: {
		{
			Address: "永胜县",
			EndYear: 2001,
		},
	},
	533223: {
		{
			Address: "华坪县",
			EndYear: 2001,
		},
	},
	533224: {
		{
			Address: "宁蒗彝族自治县",
			EndYear: 2001,
		},
	},
	533300: {
		{
			Address: "怒江傈僳族自治州",
		},
	},
	533301: {
		{
			Address:   "泸水市",
			StartYear: 2016,
		},
	},
	533321: {
		{
			Address: "碧江县",
			EndYear: 1981,
		},
		{
			Address:   "泸水县",
			StartYear: 1982,
			EndYear:   2015,
		},
	},
	533322: {
		{
			Address: "福贡县",
			EndYear: 1981,
		},
		{
			Address:   "碧江县",
			StartYear: 1982,
			EndYear:   1985,
		},
	},
	533323: {
		{
			Address: "贡山独龙族怒族自治县",
			EndYear: 1981,
		},
		{
			Address:   "福贡县",
			StartYear: 1982,
		},
	},
	533324: {
		{
			Address: "泸水县",
			EndYear: 1981,
		},
		{
			Address:   "贡山独龙族怒族自治县",
			StartYear: 1982,
		},
	},
	533325: {
		{
			Address: "兰坪县",
			EndYear: 1986,
		},
		{
			Address:   "兰坪白族普米族自治县",
			StartYear: 1987,
		},
	},
	533400: {
		{
			Address: "迪庆藏族自治州",
		},
	},
	533401: {
		{
			Address:   "香格里拉市",
			StartYear: 2014,
		},
	},
	533421: {
		{
			Address: "中甸县",
			EndYear: 2000,
		},
		{
			Address:   "香格里拉县",
			StartYear: 2001,
			EndYear:   2013,
		},
	},
	533422: {
		{
			Address: "德钦县",
		},
	},
	533423: {
		{
			Address: "维西县",
			EndYear: 1984,
		},
		{
			Address:   "维西傈僳族自治县",
			StartYear: 1985,
		},
	},
	533500: {
		{
			Address: "临沧地区",
			EndYear: 2002,
		},
	},
	533521: {
		{
			Address: "临沧县",
			EndYear: 2002,
		},
	},
	533522: {
		{
			Address: "凤庆县",
			EndYear: 2002,
		},
	},
	533523: {
		{
			Address: "云县",
			EndYear: 2002,
		},
	},
	533524: {
		{
			Address: "永德县",
			EndYear: 2002,
		},
	},
	533525: {
		{
			Address: "镇康县",
			EndYear: 2002,
		},
	},
	533526: {
		{
			Address: "双江县",
			EndYear: 1984,
		},
		{
			Address:   "双江拉祜族佤族布朗族傣族自治县",
			StartYear: 1985,
			EndYear:   2002,
		},
	},
	533527: {
		{
			Address: "耿马傣族佤族自治县",
			EndYear: 2002,
		},
	},
	533528: {
		{
			Address: "沧源佤族自治县",
			EndYear: 2002,
		},
	},
	540000: {
		{
			Address: "西藏自治区",
		},
	},
	540100: {
		{
			Address: "拉萨市",
		},
	},
	540102: {
		{
			Address: "城关区",
		},
	},
	540103: {
		{
			Address:   "堆龙德庆区",
			StartYear: 2015,
		},
	},
	540104: {
		{
			Address:   "达孜区",
			StartYear: 2017,
		},
	},
	540121: {
		{
			Address: "林周县",
		},
	},
	540122: {
		{
			Address: "当雄县",
		},
	},
	540123: {
		{
			Address: "尼木县",
		},
	},
	540124: {
		{
			Address: "曲水县",
		},
	},
	540125: {
		{
			Address: "堆龙德庆县",
			EndYear: 2014,
		},
	},
	540126: {
		{
			Address: "达孜县",
			EndYear: 2016,
		},
	},
	540127: {
		{
			Address: "墨竹工卡县",
		},
	},
	540128: {
		{
			Address: "工布江达县",
			EndYear: 1982,
		},
	},
	540129: {
		{
			Address: "林芝县",
			EndYear: 1982,
		},
	},
	540130: {
		{
			Address: "米林县",
			EndYear: 1982,
		},
	},
	540131: {
		{
			Address: "墨脱县",
			EndYear: 1982,
		},
	},
	540200: {
		{
			Address:   "日喀则市",
			StartYear: 2014,
		},
	},
	540202: {
		{
			Address:   "桑珠孜区",
			StartYear: 2014,
		},
	},
	540221: {
		{
			Address:   "南木林县",
			StartYear: 2014,
		},
	},
	540222: {
		{
			Address:   "江孜县",
			StartYear: 2014,
		},
	},
	540223: {
		{
			Address:   "定日县",
			StartYear: 2014,
		},
	},
	540224: {
		{
			Address:   "萨迦县",
			StartYear: 2014,
		},
	},
	540225: {
		{
			Address:   "拉孜县",
			StartYear: 2014,
		},
	},
	540226: {
		{
			Address:   "昂仁县",
			StartYear: 2014,
		},
	},
	540227: {
		{
			Address:   "谢通门县",
			StartYear: 2014,
		},
	},
	540228: {
		{
			Address:   "白朗县",
			StartYear: 2014,
		},
	},
	540229: {
		{
			Address:   "仁布县",
			StartYear: 2014,
		},
	},
	540230: {
		{
			Address:   "康马县",
			StartYear: 2014,
		},
	},
	540231: {
		{
			Address:   "定结县",
			StartYear: 2014,
		},
	},
	540232: {
		{
			Address:   "仲巴县",
			StartYear: 2014,
		},
	},
	540233: {
		{
			Address:   "亚东县",
			StartYear: 2014,
		},
	},
	540234: {
		{
			Address:   "吉隆县",
			StartYear: 2014,
		},
	},
	540235: {
		{
			Address:   "聂拉木县",
			StartYear: 2014,
		},
	},
	540236: {
		{
			Address:   "萨嘎县",
			StartYear: 2014,
		},
	},
	540237: {
		{
			Address:   "岗巴县",
			StartYear: 2014,
		},
	},
	540300: {
		{
			Address:   "昌都市",
			StartYear: 2014,
		},
	},
	540302: {
		{
			Address:   "卡若区",
			StartYear: 2014,
		},
	},
	540321: {
		{
			Address:   "江达县",
			StartYear: 2014,
		},
	},
	540322: {
		{
			Address:   "贡觉县",
			StartYear: 2014,
		},
	},
	540323: {
		{
			Address:   "类乌齐县",
			StartYear: 2014,
		},
	},
	540324: {
		{
			Address:   "丁青县",
			StartYear: 2014,
		},
	},
	540325: {
		{
			Address:   "察雅县",
			StartYear: 2014,
		},
	},
	540326: {
		{
			Address:   "八宿县",
			StartYear: 2014,
		},
	},
	540327: {
		{
			Address:   "左贡县",
			StartYear: 2014,
		},
	},
	540328: {
		{
			Address:   "芒康县",
			StartYear: 2014,
		},
	},
	540329: {
		{
			Address:   "洛隆县",
			StartYear: 2014,
		},
	},
	540330: {
		{
			Address:   "边坝县",
			StartYear: 2014,
		},
	},
	540400: {
		{
			Address:   "林芝市",
			StartYear: 2015,
		},
	},
	540402: {
		{
			Address:   "巴宜区",
			StartYear: 2015,
		},
	},
	540421: {
		{
			Address:   "工布江达县",
			StartYear: 2015,
		},
	},
	540422: {
		{
			Address:   "米林县",
			StartYear: 2015,
		},
	},
	540423: {
		{
			Address:   "墨脱县",
			StartYear: 2015,
		},
	},
	540424: {
		{
			Address:   "波密县",
			StartYear: 2015,
		},
	},
	540425: {
		{
			Address:   "察隅县",
			StartYear: 2015,
		},
	},
	540426: {
		{
			Address:   "朗县",
			StartYear: 2015,
		},
	},
	540500: {
		{
			Address:   "山南市",
			StartYear: 2016,
		},
	},
	540502: {
		{
			Address:   "乃东区",
			StartYear: 2016,
		},
	},
	540521: {
		{
			Address:   "扎囊县",
			StartYear: 2016,
		},
	},
	540522: {
		{
			Address:   "贡嘎县",
			StartYear: 2016,
		},
	},
	540523: {
		{
			Address:   "桑日县",
			StartYear: 2016,
		},
	},
	540524: {
		{
			Address:   "琼结县",
			StartYear: 2016,
		},
	},
	540525: {
		{
			Address:   "曲松县",
			StartYear: 2016,
		},
	},
	540526: {
		{
			Address:   "措美县",
			StartYear: 2016,
		},
	},
	540527: {
		{
			Address:   "洛扎县",
			StartYear: 2016,
		},
	},
	540528: {
		{
			Address:   "加查县",
			StartYear: 2016,
		},
	},
	540529: {
		{
			Address:   "隆子县",
			StartYear: 2016,
		},
	},
	540530: {
		{
			Address:   "错那县",
			StartYear: 2016,
		},
	},
	540531: {
		{
			Address:   "浪卡子县",
			StartYear: 2016,
		},
	},
	540600: {
		{
			Address:   "那曲市",
			StartYear: 2017,
		},
	},
	540602: {
		{
			Address:   "色尼区",
			StartYear: 2017,
		},
	},
	540621: {
		{
			Address:   "嘉黎县",
			StartYear: 2017,
		},
	},
	540622: {
		{
			Address:   "比如县",
			StartYear: 2017,
		},
	},
	540623: {
		{
			Address:   "聂荣县",
			StartYear: 2017,
		},
	},
	540624: {
		{
			Address:   "安多县",
			StartYear: 2017,
		},
	},
	540625: {
		{
			Address:   "申扎县",
			StartYear: 2017,
		},
	},
	540626: {
		{
			Address:   "索县",
			StartYear: 2017,
		},
	},
	540627: {
		{
			Address:   "班戈县",
			StartYear: 2017,
		},
	},
	540628: {
		{
			Address:   "巴青县",
			StartYear: 2017,
		},
	},
	540629: {
		{
			Address:   "尼玛县",
			StartYear: 2017,
		},
	},
	540630: {
		{
			Address:   "双湖县",
			StartYear: 2017,
		},
	},
	542100: {
		{
			Address: "昌都地区",
			EndYear: 2013,
		},
	},
	542121: {
		{
			Address: "昌都县",
			EndYear: 2013,
		},
	},
	542122: {
		{
			Address: "江达县",
			EndYear: 2013,
		},
	},
	542123: {
		{
			Address: "贡觉县",
			EndYear: 2013,
		},
	},
	542124: {
		{
			Address: "类乌齐县",
			EndYear: 2013,
		},
	},
	542125: {
		{
			Address: "丁青县",
			EndYear: 2013,
		},
	},
	542126: {
		{
			Address: "察雅县",
			EndYear: 2013,
		},
	},
	542127: {
		{
			Address: "八宿县",
			EndYear: 2013,
		},
	},
	542128: {
		{
			Address: "左贡县",
			EndYear: 2013,
		},
	},
	542129: {
		{
			Address: "芒康县",
			EndYear: 2013,
		},
	},
	542130: {
		{
			Address: "波密县",
			EndYear: 1982,
		},
	},
	542131: {
		{
			Address: "察隅县",
			EndYear: 1982,
		},
	},
	542132: {
		{
			Address: "洛隆县",
			EndYear: 2013,
		},
	},
	542133: {
		{
			Address: "边坝县",
			EndYear: 2013,
		},
	},
	542134: {
		{
			Address:   "盐井县",
			StartYear: 1983,
			EndYear:   1998,
		},
	},
	542135: {
		{
			Address:   "碧土县",
			StartYear: 1983,
			EndYear:   1998,
		},
	},
	542136: {
		{
			Address:   "妥坝县",
			StartYear: 1983,
			EndYear:   1998,
		},
	},
	542137: {
		{
			Address:   "生达县",
			StartYear: 1983,
			EndYear:   1998,
		},
	},
	542200: {
		{
			Address: "山南地区",
			EndYear: 2015,
		},
	},
	542221: {
		{
			Address: "乃东县",
			EndYear: 2015,
		},
	},
	542222: {
		{
			Address: "扎囊县",
			EndYear: 2015,
		},
	},
	542223: {
		{
			Address: "贡嘎县",
			EndYear: 2015,
		},
	},
	542224: {
		{
			Address: "桑日县",
			EndYear: 2015,
		},
	},
	542225: {
		{
			Address: "穷结县",
			EndYear: 1985,
		},
		{
			Address:   "琼结县",
			StartYear: 1986,
			EndYear:   2015,
		},
	},
	542226: {
		{
			Address: "曲松县",
			EndYear: 2015,
		},
	},
	542227: {
		{
			Address: "措美县",
			EndYear: 2015,
		},
	},
	542228: {
		{
			Address: "洛扎县",
			EndYear: 2015,
		},
	},
	542229: {
		{
			Address: "加查县",
			EndYear: 2015,
		},
	},
	542230: {
		{
			Address: "朗县",
			EndYear: 1982,
		},
	},
	542231: {
		{
			Address: "隆子县",
			EndYear: 2015,
		},
	},
	542232: {
		{
			Address: "错那县",
			EndYear: 2015,
		},
	},
	542233: {
		{
			Address: "浪卡子县",
			EndYear: 2015,
		},
	},
	542300: {
		{
			Address: "日喀则地区",
			EndYear: 2013,
		},
	},
	542301: {
		{
			Address:   "日喀则市",
			StartYear: 1986,
			EndYear:   2013,
		},
	},
	542321: {
		{
			Address: "日喀则县",
			EndYear: 1985,
		},
	},
	542322: {
		{
			Address: "南木林县",
			EndYear: 2013,
		},
	},
	542323: {
		{
			Address: "江孜县",
			EndYear: 2013,
		},
	},
	542324: {
		{
			Address: "定日县",
			EndYear: 2013,
		},
	},
	542325: {
		{
			Address: "萨迦县",
			EndYear: 2013,
		},
	},
	542326: {
		{
			Address: "拉孜县",
			EndYear: 2013,
		},
	},
	542327: {
		{
			Address: "昂仁县",
			EndYear: 2013,
		},
	},
	542328: {
		{
			Address: "谢通门县",
			EndYear: 2013,
		},
	},
	542329: {
		{
			Address: "白朗县",
			EndYear: 2013,
		},
	},
	542330: {
		{
			Address: "仁布县",
			EndYear: 2013,
		},
	},
	542331: {
		{
			Address: "康马县",
			EndYear: 2013,
		},
	},
	542332: {
		{
			Address: "定结县",
			EndYear: 2013,
		},
	},
	542333: {
		{
			Address: "仲巴县",
			EndYear: 2013,
		},
	},
	542334: {
		{
			Address: "亚东县",
			EndYear: 2013,
		},
	},
	542335: {
		{
			Address: "吉隆县",
			EndYear: 2013,
		},
	},
	542336: {
		{
			Address: "聂拉木县",
			EndYear: 2013,
		},
	},
	542337: {
		{
			Address: "萨嘎县",
			EndYear: 2013,
		},
	},
	542338: {
		{
			Address: "岗巴县",
			EndYear: 2013,
		},
	},
	542400: {
		{
			Address: "那曲地区",
			EndYear: 2016,
		},
	},
	542421: {
		{
			Address: "那曲县",
			EndYear: 2016,
		},
	},
	542422: {
		{
			Address: "嘉黎县",
			EndYear: 2016,
		},
	},
	542423: {
		{
			Address: "比如县",
			EndYear: 2016,
		},
	},
	542424: {
		{
			Address: "聂荣县",
			EndYear: 2016,
		},
	},
	542425: {
		{
			Address: "安多县",
			EndYear: 2016,
		},
	},
	542426: {
		{
			Address: "申扎县",
			EndYear: 2016,
		},
	},
	542427: {
		{
			Address: "索县",
			EndYear: 2016,
		},
	},
	542428: {
		{
			Address: "班戈县",
			EndYear: 2016,
		},
	},
	542429: {
		{
			Address: "巴青县",
			EndYear: 2016,
		},
	},
	542430: {
		{
			Address:   "尼玛县",
			StartYear: 1983,
			EndYear:   2016,
		},
	},
	542431: {
		{
			Address:   "双湖县",
			StartYear: 2012,
			EndYear:   2016,
		},
	},
	542500: {
		{
			Address: "阿里地区",
		},
	},
	542521: {
		{
			Address: "普兰县",
		},
	},
	542522: {
		{
			Address: "札达县",
		},
	},
	542523: {
		{
			Address: "噶尔县",
		},
	},
	542524: {
		{
			Address: "日土县",
		},
	},
	542525: {
		{
			Address: "革吉县",
		},
	},
	542526: {
		{
			Address: "改则县",
		},
	},
	542527: {
		{
			Address: "措勤县",
		},
	},
	542528: {
		{
			Address:   "隆格尔县",
			StartYear: 1983,
			EndYear:   1998,
		},
	},
	542600: {
		{
			Address:   "林芝地区",
			StartYear: 1983,
			EndYear:   2014,
		},
	},
	542621: {
		{
			Address:   "林芝县",
			StartYear: 1983,
			EndYear:   2014,
		},
	},
	542622: {
		{
			Address:   "工布江达县",
			StartYear: 1983,
			EndYear:   2014,
		},
	},
	542623: {
		{
			Address:   "米林县",
			StartYear: 1983,
			EndYear:   2014,
		},
	},
	542624: {
		{
			Address:   "墨脱县",
			StartYear: 1983,
			EndYear:   2014,
		},
	},
	542625: {
		{
			Address:   "波密县",
			StartYear: 1983,
			EndYear:   2014,
		},
	},
	542626: {
		{
			Address:   "察隅县",
			StartYear: 1983,
			EndYear:   2014,
		},
	},
	542627: {
		{
			Address:   "朗县",
			StartYear: 1983,
			EndYear:   2014,
		},
	},
	542700: {
		{
			Address:   "江孜地区",
			StartYear: 1983,
			EndYear:   1985,
		},
	},
	542721: {
		{
			Address:   "江孜县",
			StartYear: 1983,
			EndYear:   1985,
		},
	},
	542722: {
		{
			Address:   "浪卡子县",
			StartYear: 1983,
			EndYear:   1985,
		},
	},
	542723: {
		{
			Address:   "白朗县",
			StartYear: 1983,
			EndYear:   1985,
		},
	},
	542724: {
		{
			Address:   "仁布县",
			StartYear: 1983,
			EndYear:   1985,
		},
	},
	542725: {
		{
			Address:   "康马县",
			StartYear: 1983,
			EndYear:   1985,
		},
	},
	542726: {
		{
			Address:   "亚东县",
			StartYear: 1983,
			EndYear:   1985,
		},
	},
	542727: {
		{
			Address:   "岗巴县",
			StartYear: 1983,
			EndYear:   1985,
		},
	},
	610000: {
		{
			Address: "陕西省",
		},
	},
	610100: {
		{
			Address: "西安市",
		},
	},
	610102: {
		{
			Address: "新城区",
		},
	},
	610103: {
		{
			Address: "碑林区",
		},
	},
	610104: {
		{
			Address: "莲湖区",
		},
	},
	610111: {
		{
			Address: "灞桥区",
		},
	},
	610112: {
		{
			Address: "未央区",
		},
	},
	610113: {
		{
			Address: "雁塔区",
		},
	},
	610114: {
		{
			Address: "阎良区",
		},
	},
	610115: {
		{
			Address:   "临潼区",
			StartYear: 1997,
		},
	},
	610116: {
		{
			Address:   "长安区",
			StartYear: 2002,
		},
	},
	610117: {
		{
			Address:   "高陵区",
			StartYear: 2014,
		},
	},
	610118: {
		{
			Address:   "鄠邑区",
			StartYear: 2016,
		},
	},
	610121: {
		{
			Address: "长安县",
			EndYear: 2001,
		},
	},
	610122: {
		{
			Address:   "蓝田县",
			StartYear: 1983,
		},
	},
	610123: {
		{
			Address:   "临潼县",
			StartYear: 1983,
			EndYear:   1996,
		},
	},
	610124: {
		{
			Address:   "周至县",
			StartYear: 1983,
		},
	},
	610125: {
		{
			Address:   "户县",
			StartYear: 1983,
			EndYear:   2015,
		},
	},
	610126: {
		{
			Address:   "高陵县",
			StartYear: 1983,
			EndYear:   2013,
		},
	},
	610200: {
		{
			Address: "铜川市",
		},
	},
	610202: {
		{
			Address: "城区",
			EndYear: 1999,
		},
		{
			Address:   "王益区",
			StartYear: 2000,
		},
	},
	610203: {
		{
			Address: "郊区",
			EndYear: 1999,
		},
		{
			Address:   "印台区",
			StartYear: 2000,
		},
	},
	610204: {
		{
			Address:   "耀州区",
			StartYear: 2002,
		},
	},
	610221: {
		{
			Address: "耀县",
			EndYear: 2001,
		},
	},
	610222: {
		{
			Address:   "宜君县",
			StartYear: 1983,
		},
	},
	610300: {
		{
			Address: "宝鸡市",
		},
	},
	610302: {
		{
			Address: "渭滨区",
		},
	},
	610303: {
		{
			Address: "金台区",
		},
	},
	610304: {
		{
			Address:   "杨陵区",
			StartYear: 1982,
			EndYear:   1982,
		},
		{
			Address:   "陈仓区",
			StartYear: 2003,
		},
	},
	610321: {
		{
			Address: "宝鸡县",
			EndYear: 2002,
		},
	},
	610322: {
		{
			Address: "凤翔县",
		},
	},
	610323: {
		{
			Address: "岐山县",
		},
	},
	610324: {
		{
			Address: "扶风县",
		},
	},
	610325: {
		{
			Address: "武功县",
			EndYear: 1982,
		},
	},
	610326: {
		{
			Address: "眉县",
		},
	},
	610327: {
		{
			Address: "陇县",
		},
	},
	610328: {
		{
			Address: "千阳县",
		},
	},
	610329: {
		{
			Address: "麟游县",
		},
	},
	610330: {
		{
			Address: "凤县",
		},
	},
	610331: {
		{
			Address: "太白县",
		},
	},
	610400: {
		{
			Address:   "咸阳市",
			StartYear: 1984,
		},
	},
	610401: {
		{
			Address:   "咸阳市",
			StartYear: 1983,
			EndYear:   1983,
		},
	},
	610402: {
		{
			Address:   "秦都区",
			StartYear: 1983,
		},
	},
	610403: {
		{
			Address:   "杨陵区",
			StartYear: 1983,
		},
	},
	610404: {
		{
			Address:   "渭城区",
			StartYear: 1986,
		},
	},
	610421: {
		{
			Address:   "兴平县",
			StartYear: 1983,
			EndYear:   1992,
		},
	},
	610422: {
		{
			Address:   "三原县",
			StartYear: 1983,
		},
	},
	610423: {
		{
			Address:   "泾阳县",
			StartYear: 1983,
		},
	},
	610424: {
		{
			Address:   "乾县",
			StartYear: 1983,
		},
	},
	610425: {
		{
			Address:   "礼泉县",
			StartYear: 1983,
		},
	},
	610426: {
		{
			Address:   "永寿县",
			StartYear: 1983,
		},
	},
	610427: {
		{
			Address:   "彬县",
			StartYear: 1983,
			EndYear:   2017,
		},
	},
	610428: {
		{
			Address:   "长武县",
			StartYear: 1983,
		},
	},
	610429: {
		{
			Address:   "旬邑县",
			StartYear: 1983,
		},
	},
	610430: {
		{
			Address:   "淳化县",
			StartYear: 1983,
		},
	},
	610431: {
		{
			Address:   "武功县",
			StartYear: 1983,
		},
	},
	610481: {
		{
			Address:   "兴平市",
			StartYear: 1995,
		},
	},
	610482: {
		{
			Address:   "彬州市",
			StartYear: 2018,
		},
	},
	610500: {
		{
			Address:   "渭南市",
			StartYear: 1994,
		},
	},
	610502: {
		{
			Address:   "临渭区",
			StartYear: 1994,
		},
	},
	610503: {
		{
			Address:   "华州区",
			StartYear: 2015,
		},
	},
	610521: {
		{
			Address:   "华县",
			StartYear: 1994,
			EndYear:   2014,
		},
	},
	610522: {
		{
			Address:   "潼关县",
			StartYear: 1994,
		},
	},
	610523: {
		{
			Address:   "大荔县",
			StartYear: 1994,
		},
	},
	610524: {
		{
			Address:   "合阳县",
			StartYear: 1994,
		},
	},
	610525: {
		{
			Address:   "澄城县",
			StartYear: 1994,
		},
	},
	610526: {
		{
			Address:   "蒲城县",
			StartYear: 1994,
		},
	},
	610527: {
		{
			Address:   "白水县",
			StartYear: 1994,
		},
	},
	610528: {
		{
			Address:   "富平县",
			StartYear: 1994,
		},
	},
	610581: {
		{
			Address:   "韩城市",
			StartYear: 1995,
		},
	},
	610582: {
		{
			Address:   "华阴市",
			StartYear: 1995,
		},
	},
	610600: {
		{
			Address:   "延安市",
			StartYear: 1996,
		},
	},
	610602: {
		{
			Address:   "宝塔区",
			StartYear: 1996,
		},
	},
	610603: {
		{
			Address:   "安塞区",
			StartYear: 2016,
		},
	},
	610621: {
		{
			Address:   "延长县",
			StartYear: 1996,
		},
	},
	610622: {
		{
			Address:   "延川县",
			StartYear: 1996,
		},
	},
	610623: {
		{
			Address:   "子长县",
			StartYear: 1996,
			EndYear:   2018,
		},
	},
	610624: {
		{
			Address:   "安塞县",
			StartYear: 1996,
			EndYear:   2015,
		},
	},
	610625: {
		{
			Address:   "志丹县",
			StartYear: 1996,
		},
	},
	610626: {
		{
			Address:   "吴旗县",
			StartYear: 1996,
			EndYear:   2004,
		},
		{
			Address:   "吴起县",
			StartYear: 2005,
		},
	},
	610627: {
		{
			Address:   "甘泉县",
			StartYear: 1996,
		},
	},
	610628: {
		{
			Address:   "富县",
			StartYear: 1996,
		},
	},
	610629: {
		{
			Address:   "洛川县",
			StartYear: 1996,
		},
	},
	610630: {
		{
			Address:   "宜川县",
			StartYear: 1996,
		},
	},
	610631: {
		{
			Address:   "黄龙县",
			StartYear: 1996,
		},
	},
	610632: {
		{
			Address:   "黄陵县",
			StartYear: 1996,
		},
	},
	610681: {
		{
			Address:   "子长市",
			StartYear: 2019,
		},
	},
	610700: {
		{
			Address:   "汉中市",
			StartYear: 1996,
		},
	},
	610702: {
		{
			Address:   "汉台区",
			StartYear: 1996,
		},
	},
	610703: {
		{
			Address:   "南郑区",
			StartYear: 2017,
		},
	},
	610721: {
		{
			Address:   "南郑县",
			StartYear: 1996,
			EndYear:   2016,
		},
	},
	610722: {
		{
			Address:   "城固县",
			StartYear: 1996,
		},
	},
	610723: {
		{
			Address:   "洋县",
			StartYear: 1996,
		},
	},
	610724: {
		{
			Address:   "西乡县",
			StartYear: 1996,
		},
	},
	610725: {
		{
			Address:   "勉县",
			StartYear: 1996,
		},
	},
	610726: {
		{
			Address:   "宁强县",
			StartYear: 1996,
		},
	},
	610727: {
		{
			Address:   "略阳县",
			StartYear: 1996,
		},
	},
	610728: {
		{
			Address:   "镇巴县",
			StartYear: 1996,
		},
	},
	610729: {
		{
			Address:   "留坝县",
			StartYear: 1996,
		},
	},
	610730: {
		{
			Address:   "佛坪县",
			StartYear: 1996,
		},
	},
	610800: {
		{
			Address:   "榆林市",
			StartYear: 1999,
		},
	},
	610802: {
		{
			Address:   "榆阳区",
			StartYear: 1999,
		},
	},
	610803: {
		{
			Address:   "横山区",
			StartYear: 2015,
		},
	},
	610821: {
		{
			Address:   "神木县",
			StartYear: 1999,
			EndYear:   2016,
		},
	},
	610822: {
		{
			Address:   "府谷县",
			StartYear: 1999,
		},
	},
	610823: {
		{
			Address:   "横山县",
			StartYear: 1999,
			EndYear:   2014,
		},
	},
	610824: {
		{
			Address:   "靖边县",
			StartYear: 1999,
		},
	},
	610825: {
		{
			Address:   "定边县",
			StartYear: 1999,
		},
	},
	610826: {
		{
			Address:   "绥德县",
			StartYear: 1999,
		},
	},
	610827: {
		{
			Address:   "米脂县",
			StartYear: 1999,
		},
	},
	610828: {
		{
			Address:   "佳县",
			StartYear: 1999,
		},
	},
	610829: {
		{
			Address:   "吴堡县",
			StartYear: 1999,
		},
	},
	610830: {
		{
			Address:   "清涧县",
			StartYear: 1999,
		},
	},
	610831: {
		{
			Address:   "子洲县",
			StartYear: 1999,
		},
	},
	610881: {
		{
			Address:   "神木市",
			StartYear: 2017,
		},
	},
	610900: {
		{
			Address:   "安康市",
			StartYear: 2000,
		},
	},
	610902: {
		{
			Address:   "汉滨区",
			StartYear: 2000,
		},
	},
	610921: {
		{
			Address:   "汉阴县",
			StartYear: 2000,
		},
	},
	610922: {
		{
			Address:   "石泉县",
			StartYear: 2000,
		},
	},
	610923: {
		{
			Address:   "宁陕县",
			StartYear: 2000,
		},
	},
	610924: {
		{
			Address:   "紫阳县",
			StartYear: 2000,
		},
	},
	610925: {
		{
			Address:   "岚皋县",
			StartYear: 2000,
		},
	},
	610926: {
		{
			Address:   "平利县",
			StartYear: 2000,
		},
	},
	610927: {
		{
			Address:   "镇坪县",
			StartYear: 2000,
		},
	},
	610928: {
		{
			Address:   "旬阳县",
			StartYear: 2000,
		},
	},
	610929: {
		{
			Address:   "白河县",
			StartYear: 2000,
		},
	},
	611000: {
		{
			Address:   "商洛市",
			StartYear: 2001,
		},
	},
	611002: {
		{
			Address:   "商州区",
			StartYear: 2001,
		},
	},
	611021: {
		{
			Address:   "洛南县",
			StartYear: 2001,
		},
	},
	611022: {
		{
			Address:   "丹凤县",
			StartYear: 2001,
		},
	},
	611023: {
		{
			Address:   "商南县",
			StartYear: 2001,
		},
	},
	611024: {
		{
			Address:   "山阳县",
			StartYear: 2001,
		},
	},
	611025: {
		{
			Address:   "镇安县",
			StartYear: 2001,
		},
	},
	611026: {
		{
			Address:   "柞水县",
			StartYear: 2001,
		},
	},
	612100: {
		{
			Address: "渭南地区",
			EndYear: 1993,
		},
	},
	612101: {
		{
			Address:   "渭南市",
			StartYear: 1983,
			EndYear:   1993,
		},
	},
	612102: {
		{
			Address:   "韩城市",
			StartYear: 1983,
			EndYear:   1993,
		},
	},
	612103: {
		{
			Address:   "华阴市",
			StartYear: 1990,
			EndYear:   1993,
		},
	},
	612121: {
		{
			Address: "蓝田县",
			EndYear: 1982,
		},
	},
	612122: {
		{
			Address: "临潼县",
			EndYear: 1982,
		},
	},
	612123: {
		{
			Address: "渭南县",
			EndYear: 1982,
		},
	},
	612124: {
		{
			Address: "华县",
			EndYear: 1993,
		},
	},
	612125: {
		{
			Address: "华阴县",
			EndYear: 1989,
		},
	},
	612126: {
		{
			Address: "潼关县",
			EndYear: 1993,
		},
	},
	612127: {
		{
			Address: "大荔县",
			EndYear: 1993,
		},
	},
	612128: {
		{
			Address: "蒲城县",
			EndYear: 1993,
		},
	},
	612129: {
		{
			Address: "澄城县",
			EndYear: 1993,
		},
	},
	612130: {
		{
			Address: "白水县",
			EndYear: 1993,
		},
	},
	612131: {
		{
			Address: "韩城县",
			EndYear: 1982,
		},
	},
	612132: {
		{
			Address: "合阳县",
			EndYear: 1993,
		},
	},
	612133: {
		{
			Address: "富平县",
			EndYear: 1993,
		},
	},
	612200: {
		{
			Address: "咸阳地区",
			EndYear: 1982,
		},
	},
	612201: {
		{
			Address: "咸阳市",
			EndYear: 1982,
		},
	},
	612221: {
		{
			Address: "兴平县",
			EndYear: 1982,
		},
	},
	612222: {
		{
			Address: "周至县",
			EndYear: 1982,
		},
	},
	612223: {
		{
			Address: "户县",
			EndYear: 1982,
		},
	},
	612224: {
		{
			Address: "三原县",
			EndYear: 1982,
		},
	},
	612225: {
		{
			Address: "泾阳县",
			EndYear: 1982,
		},
	},
	612226: {
		{
			Address: "高陵县",
			EndYear: 1982,
		},
	},
	612227: {
		{
			Address: "乾县",
			EndYear: 1982,
		},
	},
	612228: {
		{
			Address: "礼泉县",
			EndYear: 1982,
		},
	},
	612229: {
		{
			Address: "永寿县",
			EndYear: 1982,
		},
	},
	612230: {
		{
			Address: "彬县",
			EndYear: 1982,
		},
	},
	612231: {
		{
			Address: "长武县",
			EndYear: 1982,
		},
	},
	612232: {
		{
			Address: "旬邑县",
			EndYear: 1982,
		},
	},
	612233: {
		{
			Address: "淳化县",
			EndYear: 1982,
		},
	},
	612300: {
		{
			Address: "汉中地区",
			EndYear: 1995,
		},
	},
	612301: {
		{
			Address: "汉中市",
			EndYear: 1995,
		},
	},
	612321: {
		{
			Address: "南郑县",
			EndYear: 1995,
		},
	},
	612322: {
		{
			Address: "城固县",
			EndYear: 1995,
		},
	},
	612323: {
		{
			Address: "洋县",
			EndYear: 1995,
		},
	},
	612324: {
		{
			Address: "西乡县",
			EndYear: 1995,
		},
	},
	612325: {
		{
			Address: "勉县",
			EndYear: 1995,
		},
	},
	612326: {
		{
			Address: "宁强县",
			EndYear: 1995,
		},
	},
	612327: {
		{
			Address: "略阳县",
			EndYear: 1995,
		},
	},
	612328: {
		{
			Address: "镇巴县",
			EndYear: 1995,
		},
	},
	612329: {
		{
			Address: "留坝县",
			EndYear: 1995,
		},
	},
	612330: {
		{
			Address: "佛坪县",
			EndYear: 1995,
		},
	},
	612400: {
		{
			Address: "安康地区",
			EndYear: 1999,
		},
	},
	612401: {
		{
			Address:   "安康市",
			StartYear: 1988,
			EndYear:   1999,
		},
	},
	612421: {
		{
			Address: "安康县",
			EndYear: 1987,
		},
	},
	612422: {
		{
			Address: "汉阴县",
			EndYear: 1999,
		},
	},
	612423: {
		{
			Address: "石泉县",
			EndYear: 1999,
		},
	},
	612424: {
		{
			Address: "宁陕县",
			EndYear: 1999,
		},
	},
	612425: {
		{
			Address: "紫阳县",
			EndYear: 1999,
		},
	},
	612426: {
		{
			Address: "岚皋县",
			EndYear: 1999,
		},
	},
	612427: {
		{
			Address: "平利县",
			EndYear: 1999,
		},
	},
	612428: {
		{
			Address: "镇坪县",
			EndYear: 1999,
		},
	},
	612429: {
		{
			Address: "旬阳县",
			EndYear: 1999,
		},
	},
	612430: {
		{
			Address: "白河县",
			EndYear: 1999,
		},
	},
	612500: {
		{
			Address: "商洛地区",
			EndYear: 2000,
		},
	},
	612501: {
		{
			Address:   "商州市",
			StartYear: 1988,
			EndYear:   2000,
		},
	},
	612521: {
		{
			Address: "商县",
			EndYear: 1987,
		},
	},
	612522: {
		{
			Address: "洛南县",
			EndYear: 2000,
		},
	},
	612523: {
		{
			Address: "丹凤县",
			EndYear: 2000,
		},
	},
	612524: {
		{
			Address: "商南县",
			EndYear: 2000,
		},
	},
	612525: {
		{
			Address: "山阳县",
			EndYear: 2000,
		},
	},
	612526: {
		{
			Address: "镇安县",
			EndYear: 2000,
		},
	},
	612527: {
		{
			Address: "柞水县",
			EndYear: 2000,
		},
	},
	612600: {
		{
			Address: "延安地区",
			EndYear: 1995,
		},
	},
	612601: {
		{
			Address: "延安市",
			EndYear: 1995,
		},
	},
	612621: {
		{
			Address: "延长县",
			EndYear: 1995,
		},
	},
	612622: {
		{
			Address: "延川县",
			EndYear: 1995,
		},
	},
	612623: {
		{
			Address: "子长县",
			EndYear: 1995,
		},
	},
	612624: {
		{
			Address: "安塞县",
			EndYear: 1995,
		},
	},
	612625: {
		{
			Address: "志丹县",
			EndYear: 1995,
		},
	},
	612626: {
		{
			Address: "吴旗县",
			EndYear: 1995,
		},
	},
	612627: {
		{
			Address: "甘泉县",
			EndYear: 1995,
		},
	},
	612628: {
		{
			Address: "富县",
			EndYear: 1995,
		},
	},
	612629: {
		{
			Address: "洛川县",
			EndYear: 1995,
		},
	},
	612630: {
		{
			Address: "宜川县",
			EndYear: 1995,
		},
	},
	612631: {
		{
			Address: "黄龙县",
			EndYear: 1995,
		},
	},
	612632: {
		{
			Address: "黄陵县",
			EndYear: 1995,
		},
	},
	612633: {
		{
			Address: "宜君县",
			EndYear: 1982,
		},
	},
	612700: {
		{
			Address: "榆林地区",
			EndYear: 1998,
		},
	},
	612701: {
		{
			Address:   "榆林市",
			StartYear: 1988,
			EndYear:   1998,
		},
	},
	612721: {
		{
			Address: "榆林县",
			EndYear: 1987,
		},
	},
	612722: {
		{
			Address: "神木县",
			EndYear: 1998,
		},
	},
	612723: {
		{
			Address: "府谷县",
			EndYear: 1998,
		},
	},
	612724: {
		{
			Address: "横山县",
			EndYear: 1998,
		},
	},
	612725: {
		{
			Address: "靖边县",
			EndYear: 1998,
		},
	},
	612726: {
		{
			Address: "定边县",
			EndYear: 1998,
		},
	},
	612727: {
		{
			Address: "绥德县",
			EndYear: 1998,
		},
	},
	612728: {
		{
			Address: "米脂县",
			EndYear: 1998,
		},
	},
	612729: {
		{
			Address: "佳县",
			EndYear: 1998,
		},
	},
	612730: {
		{
			Address: "吴堡县",
			EndYear: 1998,
		},
	},
	612731: {
		{
			Address: "清涧县",
			EndYear: 1998,
		},
	},
	612732: {
		{
			Address: "子洲县",
			EndYear: 1998,
		},
	},
	619001: {
		{
			Address:   "兴平市",
			StartYear: 1993,
			EndYear:   1994,
		},
	},
	619002: {
		{
			Address:   "韩城市",
			StartYear: 1994,
			EndYear:   1994,
		},
	},
	619003: {
		{
			Address:   "华阴市",
			StartYear: 1994,
			EndYear:   1994,
		},
	},
	620000: {
		{
			Address: "甘肃省",
		},
	},
	620100: {
		{
			Address: "兰州市",
		},
	},
	620102: {
		{
			Address: "城关区",
		},
	},
	620103: {
		{
			Address: "七里河区",
		},
	},
	620104: {
		{
			Address: "西固区",
		},
	},
	620105: {
		{
			Address: "安宁区",
		},
	},
	620111: {
		{
			Address: "红古区",
		},
	},
	620112: {
		{
			Address: "白银区",
			EndYear: 1984,
		},
	},
	620121: {
		{
			Address: "永登县",
		},
	},
	620122: {
		{
			Address: "皋兰县",
		},
	},
	620123: {
		{
			Address: "榆中县",
		},
	},
	620200: {
		{
			Address: "嘉峪关市",
		},
	},
	620300: {
		{
			Address:   "金昌市",
			StartYear: 1981,
		},
	},
	620302: {
		{
			Address:   "金川区",
			StartYear: 1984,
		},
	},
	620321: {
		{
			Address:   "永昌县",
			StartYear: 1981,
		},
	},
	620400: {
		{
			Address:   "白银市",
			StartYear: 1985,
		},
	},
	620402: {
		{
			Address:   "白银区",
			StartYear: 1985,
		},
	},
	620403: {
		{
			Address:   "平川区",
			StartYear: 1985,
		},
	},
	620421: {
		{
			Address:   "靖远县",
			StartYear: 1985,
		},
	},
	620422: {
		{
			Address:   "会宁县",
			StartYear: 1985,
		},
	},
	620423: {
		{
			Address:   "景泰县",
			StartYear: 1985,
		},
	},
	620500: {
		{
			Address:   "天水市",
			StartYear: 1985,
		},
	},
	620502: {
		{
			Address:   "秦城区",
			StartYear: 1985,
			EndYear:   2003,
		},
		{
			Address:   "秦州区",
			StartYear: 2004,
		},
	},
	620503: {
		{
			Address:   "北道区",
			StartYear: 1985,
			EndYear:   2003,
		},
		{
			Address:   "麦积区",
			StartYear: 2004,
		},
	},
	620521: {
		{
			Address:   "清水县",
			StartYear: 1985,
		},
	},
	620522: {
		{
			Address:   "秦安县",
			StartYear: 1985,
		},
	},
	620523: {
		{
			Address:   "甘谷县",
			StartYear: 1985,
		},
	},
	620524: {
		{
			Address:   "武山县",
			StartYear: 1985,
		},
	},
	620525: {
		{
			Address:   "张家川回族自治县",
			StartYear: 1985,
		},
	},
	620600: {
		{
			Address:   "武威市",
			StartYear: 2001,
		},
	},
	620602: {
		{
			Address:   "凉州区",
			StartYear: 2001,
		},
	},
	620621: {
		{
			Address:   "民勤县",
			StartYear: 2001,
		},
	},
	620622: {
		{
			Address:   "古浪县",
			StartYear: 2001,
		},
	},
	620623: {
		{
			Address:   "天祝藏族自治县",
			StartYear: 2001,
		},
	},
	620700: {
		{
			Address:   "张掖市",
			StartYear: 2002,
		},
	},
	620702: {
		{
			Address:   "甘州区",
			StartYear: 2002,
		},
	},
	620721: {
		{
			Address:   "肃南裕固族自治县",
			StartYear: 2002,
		},
	},
	620722: {
		{
			Address:   "民乐县",
			StartYear: 2002,
		},
	},
	620723: {
		{
			Address:   "临泽县",
			StartYear: 2002,
		},
	},
	620724: {
		{
			Address:   "高台县",
			StartYear: 2002,
		},
	},
	620725: {
		{
			Address:   "山丹县",
			StartYear: 2002,
		},
	},
	620800: {
		{
			Address:   "平凉市",
			StartYear: 2002,
		},
	},
	620802: {
		{
			Address:   "崆峒区",
			StartYear: 2002,
		},
	},
	620821: {
		{
			Address:   "泾川县",
			StartYear: 2002,
		},
	},
	620822: {
		{
			Address:   "灵台县",
			StartYear: 2002,
		},
	},
	620823: {
		{
			Address:   "崇信县",
			StartYear: 2002,
		},
	},
	620824: {
		{
			Address:   "华亭县",
			StartYear: 2002,
			EndYear:   2017,
		},
	},
	620825: {
		{
			Address:   "庄浪县",
			StartYear: 2002,
		},
	},
	620826: {
		{
			Address:   "静宁县",
			StartYear: 2002,
		},
	},
	620881: {
		{
			Address:   "华亭市",
			StartYear: 2018,
		},
	},
	620900: {
		{
			Address:   "酒泉市",
			StartYear: 2002,
		},
	},
	620902: {
		{
			Address:   "肃州区",
			StartYear: 2002,
		},
	},
	620921: {
		{
			Address:   "金塔县",
			StartYear: 2002,
		},
	},
	620922: {
		{
			Address:   "瓜州县",
			StartYear: 2006,
		},
	},
	620923: {
		{
			Address:   "肃北蒙古族自治县",
			StartYear: 2002,
		},
	},
	620924: {
		{
			Address:   "阿克塞哈萨克族自治县",
			StartYear: 2002,
		},
	},
	620925: {
		{
			Address:   "安西县",
			StartYear: 2002,
			EndYear:   2005,
		},
	},
	620981: {
		{
			Address:   "玉门市",
			StartYear: 2002,
		},
	},
	620982: {
		{
			Address:   "敦煌市",
			StartYear: 2002,
		},
	},
	621000: {
		{
			Address:   "庆阳市",
			StartYear: 2002,
		},
	},
	621002: {
		{
			Address:   "西峰区",
			StartYear: 2002,
		},
	},
	621021: {
		{
			Address:   "庆城县",
			StartYear: 2002,
		},
	},
	621022: {
		{
			Address:   "环县",
			StartYear: 2002,
		},
	},
	621023: {
		{
			Address:   "华池县",
			StartYear: 2002,
		},
	},
	621024: {
		{
			Address:   "合水县",
			StartYear: 2002,
		},
	},
	621025: {
		{
			Address:   "正宁县",
			StartYear: 2002,
		},
	},
	621026: {
		{
			Address:   "宁县",
			StartYear: 2002,
		},
	},
	621027: {
		{
			Address:   "镇原县",
			StartYear: 2002,
		},
	},
	621100: {
		{
			Address:   "定西市",
			StartYear: 2003,
		},
	},
	621102: {
		{
			Address:   "安定区",
			StartYear: 2003,
		},
	},
	621121: {
		{
			Address:   "通渭县",
			StartYear: 2003,
		},
	},
	621122: {
		{
			Address:   "陇西县",
			StartYear: 2003,
		},
	},
	621123: {
		{
			Address:   "渭源县",
			StartYear: 2003,
		},
	},
	621124: {
		{
			Address:   "临洮县",
			StartYear: 2003,
		},
	},
	621125: {
		{
			Address:   "漳县",
			StartYear: 2003,
		},
	},
	621126: {
		{
			Address:   "岷县",
			StartYear: 2003,
		},
	},
	621200: {
		{
			Address:   "陇南市",
			StartYear: 2004,
		},
	},
	621202: {
		{
			Address:   "武都区",
			StartYear: 2004,
		},
	},
	621221: {
		{
			Address:   "成县",
			StartYear: 2004,
		},
	},
	621222: {
		{
			Address:   "文县",
			StartYear: 2004,
		},
	},
	621223: {
		{
			Address:   "宕昌县",
			StartYear: 2004,
		},
	},
	621224: {
		{
			Address:   "康县",
			StartYear: 2004,
		},
	},
	621225: {
		{
			Address:   "西和县",
			StartYear: 2004,
		},
	},
	621226: {
		{
			Address:   "礼县",
			StartYear: 2004,
		},
	},
	621227: {
		{
			Address:   "徽县",
			StartYear: 2004,
		},
	},
	621228: {
		{
			Address:   "两当县",
			StartYear: 2004,
		},
	},
	622100: {
		{
			Address: "酒泉地区",
			EndYear: 2001,
		},
	},
	622101: {
		{
			Address: "玉门市",
			EndYear: 2001,
		},
	},
	622102: {
		{
			Address:   "酒泉市",
			StartYear: 1985,
			EndYear:   2001,
		},
	},
	622103: {
		{
			Address:   "敦煌市",
			StartYear: 1987,
			EndYear:   2001,
		},
	},
	622121: {
		{
			Address: "酒泉县",
			EndYear: 1984,
		},
	},
	622122: {
		{
			Address: "敦煌县",
			EndYear: 1986,
		},
	},
	622123: {
		{
			Address: "金塔县",
			EndYear: 2001,
		},
	},
	622124: {
		{
			Address: "肃北蒙古族自治县",
			EndYear: 2001,
		},
	},
	622125: {
		{
			Address: "阿克塞哈萨克族自治县",
			EndYear: 2001,
		},
	},
	622126: {
		{
			Address: "安西县",
			EndYear: 2001,
		},
	},
	622200: {
		{
			Address: "张掖地区",
			EndYear: 2001,
		},
	},
	622201: {
		{
			Address:   "张掖市",
			StartYear: 1985,
			EndYear:   2001,
		},
	},
	622221: {
		{
			Address: "张掖县",
			EndYear: 1984,
		},
	},
	622222: {
		{
			Address: "肃南裕固族自治县",
			EndYear: 2001,
		},
	},
	622223: {
		{
			Address: "民乐县",
			EndYear: 2001,
		},
	},
	622224: {
		{
			Address: "临泽县",
			EndYear: 2001,
		},
	},
	622225: {
		{
			Address: "高台县",
			EndYear: 2001,
		},
	},
	622226: {
		{
			Address: "山丹县",
			EndYear: 2001,
		},
	},
	622300: {
		{
			Address: "武威地区",
			EndYear: 2000,
		},
	},
	622301: {
		{
			Address:   "武威市",
			StartYear: 1985,
			EndYear:   2000,
		},
	},
	622321: {
		{
			Address: "武威县",
			EndYear: 1984,
		},
	},
	622322: {
		{
			Address: "民勤县",
			EndYear: 2000,
		},
	},
	622323: {
		{
			Address: "古浪县",
			EndYear: 2000,
		},
	},
	622324: {
		{
			Address: "景泰县",
			EndYear: 1984,
		},
	},
	622325: {
		{
			Address: "永昌县",
			EndYear: 1980,
		},
	},
	622326: {
		{
			Address: "天祝藏族自治县",
			EndYear: 2000,
		},
	},
	622400: {
		{
			Address: "定西地区",
			EndYear: 2002,
		},
	},
	622421: {
		{
			Address: "定西县",
			EndYear: 2002,
		},
	},
	622422: {
		{
			Address: "靖远县",
			EndYear: 1984,
		},
	},
	622423: {
		{
			Address: "会宁县",
			EndYear: 1984,
		},
	},
	622424: {
		{
			Address: "通渭县",
			EndYear: 2002,
		},
	},
	622425: {
		{
			Address: "陇西县",
			EndYear: 2002,
		},
	},
	622426: {
		{
			Address: "渭源县",
			EndYear: 2002,
		},
	},
	622427: {
		{
			Address: "临洮县",
			EndYear: 2002,
		},
	},
	622428: {
		{
			Address:   "漳县",
			StartYear: 1985,
			EndYear:   2002,
		},
	},
	622429: {
		{
			Address:   "岷县",
			StartYear: 1985,
			EndYear:   2002,
		},
	},
	622500: {
		{
			Address: "天水地区",
			EndYear: 1984,
		},
	},
	622501: {
		{
			Address: "天水市",
			EndYear: 1984,
		},
	},
	622521: {
		{
			Address: "张家川回族自治县",
			EndYear: 1984,
		},
	},
	622522: {
		{
			Address: "天水县",
			EndYear: 1984,
		},
	},
	622523: {
		{
			Address: "清水县",
			EndYear: 1984,
		},
	},
	622524: {
		{
			Address: "徽县",
			EndYear: 1984,
		},
	},
	622525: {
		{
			Address: "两当县",
			EndYear: 1984,
		},
	},
	622526: {
		{
			Address: "礼县",
			EndYear: 1984,
		},
	},
	622527: {
		{
			Address: "西和县",
			EndYear: 1984,
		},
	},
	622528: {
		{
			Address: "武山县",
			EndYear: 1984,
		},
	},
	622529: {
		{
			Address: "甘谷县",
			EndYear: 1984,
		},
	},
	622530: {
		{
			Address: "秦安县",
			EndYear: 1984,
		},
	},
	622531: {
		{
			Address: "漳县",
			EndYear: 1984,
		},
	},
	622600: {
		{
			Address: "武都地区",
			EndYear: 1984,
		},
		{
			Address:   "陇南地区",
			StartYear: 1985,
			EndYear:   2003,
		},
	},
	622621: {
		{
			Address: "武都县",
			EndYear: 2003,
		},
	},
	622622: {
		{
			Address: "岷县",
			EndYear: 1984,
		},
	},
	622623: {
		{
			Address: "宕昌县",
			EndYear: 2003,
		},
	},
	622624: {
		{
			Address: "成县",
			EndYear: 2003,
		},
	},
	622625: {
		{
			Address: "康县",
			EndYear: 2003,
		},
	},
	622626: {
		{
			Address: "文县",
			EndYear: 2003,
		},
	},
	622627: {
		{
			Address:   "西和县",
			StartYear: 1985,
			EndYear:   2003,
		},
	},
	622628: {
		{
			Address:   "礼县",
			StartYear: 1985,
			EndYear:   2003,
		},
	},
	622629: {
		{
			Address:   "两当县",
			StartYear: 1985,
			EndYear:   2003,
		},
	},
	622630: {
		{
			Address:   "徽县",
			StartYear: 1985,
			EndYear:   2003,
		},
	},
	622700: {
		{
			Address: "平凉地区",
			EndYear: 2001,
		},
	},
	622701: {
		{
			Address:   "平凉市",
			StartYear: 1983,
			EndYear:   2001,
		},
	},
	622721: {
		{
			Address: "平凉县",
			EndYear: 1982,
		},
	},
	622722: {
		{
			Address: "泾川县",
			EndYear: 2001,
		},
	},
	622723: {
		{
			Address: "灵台县",
			EndYear: 2001,
		},
	},
	622724: {
		{
			Address: "崇信县",
			EndYear: 2001,
		},
	},
	622725: {
		{
			Address: "华亭县",
			EndYear: 2001,
		},
	},
	622726: {
		{
			Address: "庄浪县",
			EndYear: 2001,
		},
	},
	622727: {
		{
			Address: "静宁县",
			EndYear: 2001,
		},
	},
	622800: {
		{
			Address: "庆阳地区",
			EndYear: 2001,
		},
	},
	622801: {
		{
			Address:   "西峰市",
			StartYear: 1985,
			EndYear:   2001,
		},
	},
	622821: {
		{
			Address: "庆阳县",
			EndYear: 2001,
		},
	},
	622822: {
		{
			Address: "环县",
			EndYear: 2001,
		},
	},
	622823: {
		{
			Address: "华池县",
			EndYear: 2001,
		},
	},
	622824: {
		{
			Address: "合水县",
			EndYear: 2001,
		},
	},
	622825: {
		{
			Address: "正宁县",
			EndYear: 2001,
		},
	},
	622826: {
		{
			Address: "宁县",
			EndYear: 2001,
		},
	},
	622827: {
		{
			Address: "镇原县",
			EndYear: 2001,
		},
	},
	622900: {
		{
			Address: "临夏回族自治州",
		},
	},
	622901: {
		{
			Address:   "临夏市",
			StartYear: 1983,
		},
	},
	622921: {
		{
			Address: "临夏县",
		},
	},
	622922: {
		{
			Address: "康乐县",
		},
	},
	622923: {
		{
			Address: "永靖县",
		},
	},
	622924: {
		{
			Address: "广河县",
		},
	},
	622925: {
		{
			Address: "和政县",
		},
	},
	622926: {
		{
			Address: "东乡族自治县",
		},
	},
	622927: {
		{
			Address: "积石山保安族东乡族撒拉族自治县",
		},
	},
	623000: {
		{
			Address: "甘南藏族自治州",
		},
	},
	623001: {
		{
			Address:   "合作市",
			StartYear: 1996,
		},
	},
	623021: {
		{
			Address: "临潭县",
		},
	},
	623022: {
		{
			Address: "卓尼县",
		},
	},
	623023: {
		{
			Address: "舟曲县",
		},
	},
	623024: {
		{
			Address: "迭部县",
		},
	},
	623025: {
		{
			Address: "玛曲县",
		},
	},
	623026: {
		{
			Address: "碌曲县",
		},
	},
	623027: {
		{
			Address: "夏河县",
		},
	},
	630000: {
		{
			Address: "青海省",
		},
	},
	630100: {
		{
			Address: "西宁市",
		},
	},
	630102: {
		{
			Address: "城东区",
		},
	},
	630103: {
		{
			Address: "城中区",
		},
	},
	630104: {
		{
			Address: "城西区",
		},
	},
	630105: {
		{
			Address:   "城北区",
			StartYear: 1986,
		},
	},
	630106: {
		{
			Address:   "湟中区",
			StartYear: 2019,
		},
	},
	630111: {
		{
			Address: "郊区",
			EndYear: 1985,
		},
	},
	630121: {
		{
			Address: "大通县",
			EndYear: 1984,
		},
		{
			Address:   "大通回族土族自治县",
			StartYear: 1985,
		},
	},
	630122: {
		{
			Address:   "湟中县",
			StartYear: 1999,
			EndYear:   2018,
		},
	},
	630123: {
		{
			Address:   "湟源县",
			StartYear: 1999,
		},
	},
	630200: {
		{
			Address:   "海东市",
			StartYear: 2013,
		},
	},
	630202: {
		{
			Address:   "乐都区",
			StartYear: 2013,
		},
	},
	630203: {
		{
			Address:   "平安区",
			StartYear: 2015,
		},
	},
	630221: {
		{
			Address:   "平安县",
			StartYear: 2013,
			EndYear:   2014,
		},
	},
	630222: {
		{
			Address:   "民和回族土族自治县",
			StartYear: 2013,
		},
	},
	630223: {
		{
			Address:   "互助土族自治县",
			StartYear: 2013,
		},
	},
	630224: {
		{
			Address:   "化隆回族自治县",
			StartYear: 2013,
		},
	},
	630225: {
		{
			Address:   "循化撒拉族自治县",
			StartYear: 2013,
		},
	},
	632100: {
		{
			Address: "海东地区",
			EndYear: 2012,
		},
	},
	632121: {
		{
			Address: "平安县",
			EndYear: 2012,
		},
	},
	632122: {
		{
			Address: "民和县",
			EndYear: 1984,
		},
		{
			Address:   "民和回族土族自治县",
			StartYear: 1985,
			EndYear:   2012,
		},
	},
	632123: {
		{
			Address: "乐都县",
			EndYear: 2012,
		},
	},
	632124: {
		{
			Address: "湟中县",
			EndYear: 1998,
		},
	},
	632125: {
		{
			Address: "湟源县",
			EndYear: 1998,
		},
	},
	632126: {
		{
			Address: "互助土族自治县",
			EndYear: 2012,
		},
	},
	632127: {
		{
			Address: "化隆回族自治县",
			EndYear: 2012,
		},
	},
	632128: {
		{
			Address: "循化撒拉族自治县",
			EndYear: 2012,
		},
	},
	632200: {
		{
			Address: "海北藏族自治州",
		},
	},
	632221: {
		{
			Address: "门源回族自治县",
		},
	},
	632222: {
		{
			Address: "祁连县",
		},
	},
	632223: {
		{
			Address: "海晏县",
		},
	},
	632224: {
		{
			Address: "刚察县",
		},
	},
	632300: {
		{
			Address: "黄南藏族自治州",
		},
	},
	632301: {
		{
			Address:   "同仁市",
			StartYear: 2020,
		},
	},
	632321: {
		{
			Address: "同仁县",
			EndYear: 2019,
		},
	},
	632322: {
		{
			Address: "尖扎县",
		},
	},
	632323: {
		{
			Address: "泽库县",
		},
	},
	632324: {
		{
			Address:   "河南蒙古族自治县",
			StartYear: 1981,
		},
	},
	632421: {
		{
			Address: "河南蒙古族自治县",
			EndYear: 1983,
		},
	},
	632500: {
		{
			Address: "海南藏族自治州",
		},
	},
	632521: {
		{
			Address: "共和县",
		},
	},
	632522: {
		{
			Address: "同德县",
		},
	},
	632523: {
		{
			Address: "贵德县",
		},
	},
	632524: {
		{
			Address: "兴海县",
		},
	},
	632525: {
		{
			Address: "贵南县",
		},
	},
	632600: {
		{
			Address: "果洛藏族自治州",
		},
	},
	632621: {
		{
			Address: "玛沁县",
		},
	},
	632622: {
		{
			Address: "班玛县",
		},
	},
	632623: {
		{
			Address: "甘德县",
		},
	},
	632624: {
		{
			Address: "达日县",
		},
	},
	632625: {
		{
			Address: "久治县",
		},
	},
	632626: {
		{
			Address: "玛多县",
		},
	},
	632700: {
		{
			Address: "玉树藏族自治州",
		},
	},
	632701: {
		{
			Address:   "玉树市",
			StartYear: 2013,
		},
	},
	632721: {
		{
			Address: "玉树县",
			EndYear: 2012,
		},
	},
	632722: {
		{
			Address: "杂多县",
		},
	},
	632723: {
		{
			Address: "称多县",
		},
	},
	632724: {
		{
			Address: "治多县",
		},
	},
	632725: {
		{
			Address: "囊谦县",
		},
	},
	632726: {
		{
			Address: "曲麻莱县",
		},
	},
	632800: {
		{
			Address: "海西蒙古族藏族自治州",
		},
	},
	632801: {
		{
			Address: "格尔木市",
		},
	},
	632802: {
		{
			Address:   "德令哈市",
			StartYear: 1988,
		},
	},
	632803: {
		{
			Address:   "茫崖市",
			StartYear: 2018,
		},
	},
	632821: {
		{
			Address: "乌兰县",
		},
	},
	632822: {
		{
			Address: "都兰县",
		},
	},
	632823: {
		{
			Address: "天峻县",
		},
	},
	640000: {
		{
			Address: "宁夏回族自治区",
		},
	},
	640100: {
		{
			Address: "银川市",
		},
	},
	640102: {
		{
			Address: "城区",
			EndYear: 2001,
		},
	},
	640103: {
		{
			Address: "新城区",
			EndYear: 2001,
		},
	},
	640104: {
		{
			Address: "郊区",
			EndYear: 2001,
		},
		{
			Address:   "兴庆区",
			StartYear: 2002,
		},
	},
	640105: {
		{
			Address:   "西夏区",
			StartYear: 2002,
		},
	},
	640106: {
		{
			Address:   "金凤区",
			StartYear: 2002,
		},
	},
	640121: {
		{
			Address: "永宁县",
		},
	},
	640122: {
		{
			Address: "贺兰县",
		},
	},
	640181: {
		{
			Address:   "灵武市",
			StartYear: 2002,
		},
	},
	640200: {
		{
			Address: "石咀山市",
			EndYear: 1986,
		},
		{
			Address:   "石嘴山市",
			StartYear: 1987,
		},
	},
	640202: {
		{
			Address: "大武口区",
		},
	},
	640204: {
		{
			Address: "石炭井区",
			EndYear: 2001,
		},
	},
	640205: {
		{
			Address: "石咀山区",
			EndYear: 1986,
		},
		{
			Address:   "石嘴山区",
			StartYear: 1987,
			EndYear:   2002,
		},
		{
			Address:   "惠农区",
			StartYear: 2003,
		},
	},
	640211: {
		{
			Address: "郊区",
			EndYear: 1986,
		},
	},
	640221: {
		{
			Address: "平罗县",
		},
	},
	640222: {
		{
			Address: "陶乐县",
			EndYear: 2002,
		},
	},
	640223: {
		{
			Address:   "惠农县",
			StartYear: 1987,
			EndYear:   2002,
		},
	},
	640300: {
		{
			Address:   "吴忠市",
			StartYear: 1998,
		},
	},
	640302: {
		{
			Address:   "利通区",
			StartYear: 1998,
		},
	},
	640303: {
		{
			Address:   "红寺堡区",
			StartYear: 2009,
		},
	},
	640321: {
		{
			Address:   "中卫县",
			StartYear: 1998,
			EndYear:   2002,
		},
	},
	640322: {
		{
			Address:   "中宁县",
			StartYear: 1998,
			EndYear:   2002,
		},
	},
	640323: {
		{
			Address:   "盐池县",
			StartYear: 1998,
		},
	},
	640324: {
		{
			Address:   "同心县",
			StartYear: 1998,
		},
	},
	640381: {
		{
			Address:   "青铜峡市",
			StartYear: 1998,
		},
	},
	640382: {
		{
			Address:   "灵武市",
			StartYear: 1998,
			EndYear:   2001,
		},
	},
	640400: {
		{
			Address:   "固原市",
			StartYear: 2001,
		},
	},
	640402: {
		{
			Address:   "原州区",
			StartYear: 2001,
		},
	},
	640421: {
		{
			Address:   "海原县",
			StartYear: 2001,
			EndYear:   2002,
		},
	},
	640422: {
		{
			Address:   "西吉县",
			StartYear: 2001,
		},
	},
	640423: {
		{
			Address:   "隆德县",
			StartYear: 2001,
		},
	},
	640424: {
		{
			Address:   "泾源县",
			StartYear: 2001,
		},
	},
	640425: {
		{
			Address:   "彭阳县",
			StartYear: 2001,
		},
	},
	640500: {
		{
			Address:   "中卫市",
			StartYear: 2003,
		},
	},
	640502: {
		{
			Address:   "沙坡头区",
			StartYear: 2003,
		},
	},
	640521: {
		{
			Address:   "中宁县",
			StartYear: 2003,
		},
	},
	640522: {
		{
			Address:   "海原县",
			StartYear: 2003,
		},
	},
	642100: {
		{
			Address: "银南地区",
			EndYear: 1997,
		},
	},
	642101: {
		{
			Address:   "吴忠市",
			StartYear: 1983,
			EndYear:   1997,
		},
	},
	642102: {
		{
			Address:   "青铜峡市",
			StartYear: 1984,
			EndYear:   1997,
		},
	},
	642103: {
		{
			Address:   "灵武市",
			StartYear: 1996,
			EndYear:   1997,
		},
	},
	642121: {
		{
			Address: "吴忠县",
			EndYear: 1982,
		},
	},
	642122: {
		{
			Address: "青铜峡县",
			EndYear: 1983,
		},
	},
	642123: {
		{
			Address: "中卫县",
			EndYear: 1997,
		},
	},
	642124: {
		{
			Address: "中宁县",
			EndYear: 1997,
		},
	},
	642125: {
		{
			Address: "灵武县",
			EndYear: 1995,
		},
	},
	642126: {
		{
			Address: "盐池县",
			EndYear: 1997,
		},
	},
	642127: {
		{
			Address: "同心县",
			EndYear: 1997,
		},
	},
	642200: {
		{
			Address: "固原地区",
			EndYear: 2000,
		},
	},
	642221: {
		{
			Address: "固原县",
			EndYear: 2000,
		},
	},
	642222: {
		{
			Address: "海原县",
			EndYear: 2000,
		},
	},
	642223: {
		{
			Address: "西吉县",
			EndYear: 2000,
		},
	},
	642224: {
		{
			Address: "隆德县",
			EndYear: 2000,
		},
	},
	642225: {
		{
			Address: "泾源县",
			EndYear: 2000,
		},
	},
	642226: {
		{
			Address:   "彭阳县",
			StartYear: 1983,
			EndYear:   2000,
		},
	},
	650000: {
		{
			Address: "新疆维吾尔自治区",
		},
	},
	650100: {
		{
			Address: "乌鲁木齐市",
		},
	},
	650102: {
		{
			Address: "天山区",
		},
	},
	650103: {
		{
			Address: "沙依巴克区",
		},
	},
	650104: {
		{
			Address: "新市区",
		},
	},
	650105: {
		{
			Address: "水磨沟区",
		},
	},
	650106: {
		{
			Address: "头屯河区",
		},
	},
	650107: {
		{
			Address: "南山区",
			EndYear: 1988,
		},
		{
			Address:   "南山矿区",
			StartYear: 1989,
			EndYear:   1998,
		},
		{
			Address:   "南泉区",
			StartYear: 1999,
			EndYear:   2001,
		},
		{
			Address:   "达坂城区",
			StartYear: 2002,
		},
	},
	650108: {
		{
			Address:   "东山区",
			StartYear: 1987,
			EndYear:   2006,
		},
	},
	650109: {
		{
			Address:   "米东区",
			StartYear: 2007,
		},
	},
	650121: {
		{
			Address: "乌鲁木齐县",
		},
	},
	650200: {
		{
			Address: "克拉玛依市",
		},
	},
	650202: {
		{
			Address: "独山子区",
		},
	},
	650203: {
		{
			Address:   "克拉玛依区",
			StartYear: 1982,
		},
	},
	650204: {
		{
			Address:   "白碱滩区",
			StartYear: 1982,
		},
	},
	650205: {
		{
			Address:   "乌尔禾区",
			StartYear: 1982,
		},
	},
	650300: {
		{
			Address: "石河子市",
			EndYear: 1984,
		},
	},
	650400: {
		{
			Address:   "吐鲁番市",
			StartYear: 2015,
		},
	},
	650402: {
		{
			Address:   "高昌区",
			StartYear: 2015,
		},
	},
	650421: {
		{
			Address:   "鄯善县",
			StartYear: 2015,
		},
	},
	650422: {
		{
			Address:   "托克逊县",
			StartYear: 2015,
		},
	},
	650500: {
		{
			Address:   "哈密市",
			StartYear: 2016,
		},
	},
	650502: {
		{
			Address:   "伊州区",
			StartYear: 2016,
		},
	},
	650521: {
		{
			Address:   "巴里坤哈萨克自治县",
			StartYear: 2016,
		},
	},
	650522: {
		{
			Address:   "伊吾县",
			StartYear: 2016,
		},
	},
	652100: {
		{
			Address: "吐鲁番地区",
			EndYear: 2014,
		},
	},
	652101: {
		{
			Address:   "吐鲁番市",
			StartYear: 1984,
			EndYear:   2014,
		},
	},
	652121: {
		{
			Address: "吐鲁番县",
			EndYear: 1983,
		},
	},
	652122: {
		{
			Address: "鄯善县",
			EndYear: 2014,
		},
	},
	652123: {
		{
			Address: "托克逊县",
			EndYear: 2014,
		},
	},
	652200: {
		{
			Address: "哈密地区",
			EndYear: 2015,
		},
	},
	652201: {
		{
			Address: "哈密市",
			EndYear: 2015,
		},
	},
	652221: {
		{
			Address: "哈密县",
			EndYear: 1982,
		},
	},
	652222: {
		{
			Address: "巴里坤哈萨克自治县",
			EndYear: 2015,
		},
	},
	652223: {
		{
			Address: "伊吾县",
			EndYear: 2015,
		},
	},
	652300: {
		{
			Address: "昌吉回族自治州",
		},
	},
	652301: {
		{
			Address:   "昌吉市",
			StartYear: 1983,
		},
	},
	652302: {
		{
			Address:   "阜康市",
			StartYear: 1992,
		},
	},
	652303: {
		{
			Address:   "米泉市",
			StartYear: 1996,
			EndYear:   2006,
		},
	},
	652321: {
		{
			Address: "昌吉县",
			EndYear: 1982,
		},
	},
	652322: {
		{
			Address: "米泉县",
			EndYear: 1995,
		},
	},
	652323: {
		{
			Address: "呼图壁县",
		},
	},
	652324: {
		{
			Address: "玛纳斯县",
		},
	},
	652325: {
		{
			Address: "奇台县",
		},
	},
	652326: {
		{
			Address: "阜康县",
			EndYear: 1991,
		},
	},
	652327: {
		{
			Address: "吉木萨尔县",
		},
	},
	652328: {
		{
			Address: "木垒哈萨克自治县",
		},
	},
	652400: {
		{
			Address: "伊犁哈萨克自治州",
			EndYear: 1983,
		},
	},
	652401: {
		{
			Address: "奎屯市",
			EndYear: 1981,
		},
		{
			Address:   "伊宁市",
			StartYear: 1982,
			EndYear:   1983,
		},
	},
	652402: {
		{
			Address: "伊宁市",
			EndYear: 1981,
		},
		{
			Address:   "一区",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	652403: {
		{
			Address: "一区",
			EndYear: 1981,
		},
		{
			Address:   "二区",
			StartYear: 1982,
			EndYear:   1982,
		},
	},
	652404: {
		{
			Address: "二区",
			EndYear: 1981,
		},
		{
			Address:   "奎屯市",
			StartYear: 1982,
			EndYear:   1983,
		},
	},
	652421: {
		{
			Address: "伊宁县",
			EndYear: 1983,
		},
	},
	652422: {
		{
			Address: "察布查尔锡伯自治县",
			EndYear: 1983,
		},
	},
	652423: {
		{
			Address: "霍城县",
			EndYear: 1983,
		},
	},
	652424: {
		{
			Address: "巩留县",
			EndYear: 1983,
		},
	},
	652425: {
		{
			Address: "新源县",
			EndYear: 1983,
		},
	},
	652426: {
		{
			Address: "昭苏县",
			EndYear: 1983,
		},
	},
	652427: {
		{
			Address: "特克斯县",
			EndYear: 1983,
		},
	},
	652428: {
		{
			Address: "尼勒克县",
			EndYear: 1983,
		},
	},
	652500: {
		{
			Address: "塔城地区",
			EndYear: 1983,
		},
	},
	652521: {
		{
			Address: "塔城县",
			EndYear: 1983,
		},
	},
	652522: {
		{
			Address: "额敏县",
			EndYear: 1983,
		},
	},
	652523: {
		{
			Address: "乌苏县",
			EndYear: 1983,
		},
	},
	652524: {
		{
			Address: "沙湾县",
			EndYear: 1983,
		},
	},
	652525: {
		{
			Address: "托里县",
			EndYear: 1983,
		},
	},
	652526: {
		{
			Address: "裕民县",
			EndYear: 1983,
		},
	},
	652527: {
		{
			Address: "和布克赛尔蒙古自治县",
			EndYear: 1983,
		},
	},
	652600: {
		{
			Address: "阿勒泰地区",
			EndYear: 1983,
		},
	},
	652621: {
		{
			Address: "阿勒泰县",
			EndYear: 1983,
		},
	},
	652622: {
		{
			Address: "布尔津县",
			EndYear: 1983,
		},
	},
	652623: {
		{
			Address: "富蕴县",
			EndYear: 1983,
		},
	},
	652624: {
		{
			Address: "福海县",
			EndYear: 1983,
		},
	},
	652625: {
		{
			Address: "哈巴河县",
			EndYear: 1983,
		},
	},
	652626: {
		{
			Address: "青河县",
			EndYear: 1983,
		},
	},
	652627: {
		{
			Address: "吉木乃县",
			EndYear: 1983,
		},
	},
	652700: {
		{
			Address: "博尔塔拉蒙古自治州",
		},
	},
	652701: {
		{
			Address:   "博乐市",
			StartYear: 1985,
		},
	},
	652702: {
		{
			Address:   "阿拉山口市",
			StartYear: 2012,
		},
	},
	652721: {
		{
			Address: "博乐县",
			EndYear: 1984,
		},
	},
	652722: {
		{
			Address: "精河县",
		},
	},
	652723: {
		{
			Address: "温泉县",
		},
	},
	652800: {
		{
			Address: "巴音郭楞蒙古自治州",
		},
	},
	652801: {
		{
			Address: "库尔勒市",
		},
	},
	652821: {
		{
			Address: "库尔勒县",
			EndYear: 1982,
		},
	},
	652822: {
		{
			Address: "轮台县",
		},
	},
	652823: {
		{
			Address: "尉犁县",
		},
	},
	652824: {
		{
			Address: "若羌县",
		},
	},
	652825: {
		{
			Address: "且末县",
		},
	},
	652826: {
		{
			Address: "焉耆回族自治县",
		},
	},
	652827: {
		{
			Address: "和静县",
		},
	},
	652828: {
		{
			Address: "和硕县",
		},
	},
	652829: {
		{
			Address: "博湖县",
		},
	},
	652900: {
		{
			Address: "阿克苏地区",
		},
	},
	652901: {
		{
			Address:   "阿克苏市",
			StartYear: 1983,
		},
	},
	652902: {
		{
			Address:   "库车市",
			StartYear: 2019,
		},
	},
	652921: {
		{
			Address: "阿克苏县",
			EndYear: 1982,
		},
	},
	652922: {
		{
			Address: "温宿县",
		},
	},
	652923: {
		{
			Address: "库车县",
			EndYear: 2018,
		},
	},
	652924: {
		{
			Address: "沙雅县",
		},
	},
	652925: {
		{
			Address: "新和县",
		},
	},
	652926: {
		{
			Address: "拜城县",
		},
	},
	652927: {
		{
			Address: "乌什县",
		},
	},
	652928: {
		{
			Address: "阿瓦提县",
		},
	},
	652929: {
		{
			Address: "柯坪县",
		},
	},
	653000: {
		{
			Address: "克孜勒苏柯尔克孜自治州",
		},
	},
	653001: {
		{
			Address:   "阿图什市",
			StartYear: 1986,
		},
	},
	653021: {
		{
			Address: "阿图什县",
			EndYear: 1985,
		},
	},
	653022: {
		{
			Address: "阿克陶县",
		},
	},
	653023: {
		{
			Address: "阿合奇县",
		},
	},
	653024: {
		{
			Address: "乌恰县",
		},
	},
	653100: {
		{
			Address: "喀什地区",
		},
	},
	653101: {
		{
			Address: "喀什市",
		},
	},
	653121: {
		{
			Address: "疏附县",
		},
	},
	653122: {
		{
			Address: "疏勒县",
		},
	},
	653123: {
		{
			Address: "英吉沙县",
		},
	},
	653124: {
		{
			Address: "泽普县",
		},
	},
	653125: {
		{
			Address: "莎车县",
		},
	},
	653126: {
		{
			Address: "叶城县",
		},
	},
	653127: {
		{
			Address: "麦盖提县",
		},
	},
	653128: {
		{
			Address: "岳普湖县",
		},
	},
	653129: {
		{
			Address: "伽师县",
		},
	},
	653130: {
		{
			Address: "巴楚县",
		},
	},
	653131: {
		{
			Address: "塔什库尔干塔吉克自治县",
		},
	},
	653200: {
		{
			Address: "和田地区",
		},
	},
	653201: {
		{
			Address:   "和田市",
			StartYear: 1983,
		},
	},
	653221: {
		{
			Address: "和田县",
		},
	},
	653222: {
		{
			Address: "墨玉县",
		},
	},
	653223: {
		{
			Address: "皮山县",
		},
	},
	653224: {
		{
			Address: "洛浦县",
		},
	},
	653225: {
		{
			Address: "策勒县",
		},
	},
	653226: {
		{
			Address: "于田县",
		},
	},
	653227: {
		{
			Address: "民丰县",
		},
	},
	654000: {
		{
			Address:   "伊犁哈萨克自治州",
			StartYear: 1984,
		},
	},
	654001: {
		{
			Address:   "奎屯市",
			StartYear: 1984,
			EndYear:   2000,
		},
	},
	654002: {
		{
			Address:   "伊宁市",
			StartYear: 2001,
		},
	},
	654003: {
		{
			Address:   "奎屯市",
			StartYear: 2001,
		},
	},
	654004: {
		{
			Address:   "霍尔果斯市",
			StartYear: 2014,
		},
	},
	654021: {
		{
			Address:   "伊宁县",
			StartYear: 2001,
		},
	},
	654022: {
		{
			Address:   "察布查尔锡伯自治县",
			StartYear: 2001,
		},
	},
	654023: {
		{
			Address:   "霍城县",
			StartYear: 2001,
		},
	},
	654024: {
		{
			Address:   "巩留县",
			StartYear: 2001,
		},
	},
	654025: {
		{
			Address:   "新源县",
			StartYear: 2001,
		},
	},
	654026: {
		{
			Address:   "昭苏县",
			StartYear: 2001,
		},
	},
	654027: {
		{
			Address:   "特克斯县",
			StartYear: 2001,
		},
	},
	654028: {
		{
			Address:   "尼勒克县",
			StartYear: 2001,
		},
	},
	654100: {
		{
			Address:   "伊犁地区",
			StartYear: 1984,
			EndYear:   2000,
		},
	},
	654101: {
		{
			Address:   "伊宁市",
			StartYear: 1984,
			EndYear:   2000,
		},
	},
	654121: {
		{
			Address:   "伊宁县",
			StartYear: 1984,
			EndYear:   2000,
		},
	},
	654122: {
		{
			Address:   "察布查尔锡伯自治县",
			StartYear: 1984,
			EndYear:   2000,
		},
	},
	654123: {
		{
			Address:   "霍城县",
			StartYear: 1984,
			EndYear:   2000,
		},
	},
	654124: {
		{
			Address:   "巩留县",
			StartYear: 1984,
			EndYear:   2000,
		},
	},
	654125: {
		{
			Address:   "新源县",
			StartYear: 1984,
			EndYear:   2000,
		},
	},
	654126: {
		{
			Address:   "昭苏县",
			StartYear: 1984,
			EndYear:   2000,
		},
	},
	654127: {
		{
			Address:   "特克斯县",
			StartYear: 1984,
			EndYear:   2000,
		},
	},
	654128: {
		{
			Address:   "尼勒克县",
			StartYear: 1984,
			EndYear:   2000,
		},
	},
	654200: {
		{
			Address:   "塔城地区",
			StartYear: 1984,
		},
	},
	654201: {
		{
			Address:   "塔城市",
			StartYear: 1984,
		},
	},
	654202: {
		{
			Address:   "乌苏市",
			StartYear: 1996,
		},
	},
	654221: {
		{
			Address:   "额敏县",
			StartYear: 1984,
		},
	},
	654222: {
		{
			Address:   "乌苏县",
			StartYear: 1984,
			EndYear:   1995,
		},
	},
	654223: {
		{
			Address:   "沙湾县",
			StartYear: 1984,
		},
	},
	654224: {
		{
			Address:   "托里县",
			StartYear: 1984,
		},
	},
	654225: {
		{
			Address:   "裕民县",
			StartYear: 1984,
		},
	},
	654226: {
		{
			Address:   "和布克赛尔蒙古自治县",
			StartYear: 1984,
		},
	},
	654300: {
		{
			Address:   "阿勒泰地区",
			StartYear: 1984,
		},
	},
	654301: {
		{
			Address:   "阿勒泰市",
			StartYear: 1984,
		},
	},
	654321: {
		{
			Address:   "布尔津县",
			StartYear: 1984,
		},
	},
	654322: {
		{
			Address:   "富蕴县",
			StartYear: 1984,
		},
	},
	654323: {
		{
			Address:   "福海县",
			StartYear: 1984,
		},
	},
	654324: {
		{
			Address:   "哈巴河县",
			StartYear: 1984,
		},
	},
	654325: {
		{
			Address:   "青河县",
			StartYear: 1984,
		},
	},
	654326: {
		{
			Address:   "吉木乃县",
			StartYear: 1984,
		},
	},
	659001: {
		{
			Address:   "石河子市",
			StartYear: 1985,
		},
	},
	659002: {
		{
			Address:   "阿拉尔市",
			StartYear: 2002,
		},
	},
	659003: {
		{
			Address:   "图木舒克市",
			StartYear: 2002,
		},
	},
	659004: {
		{
			Address:   "五家渠市",
			StartYear: 2002,
		},
	},
	659005: {
		{
			Address:   "北屯市",
			StartYear: 2011,
		},
	},
	659006: {
		{
			Address:   "铁门关市",
			StartYear: 2012,
		},
	},
	659007: {
		{
			Address:   "双河市",
			StartYear: 2014,
		},
	},
	659008: {
		{
			Address:   "可克达拉市",
			StartYear: 2015,
		},
	},
	659009: {
		{
			Address:   "昆玉市",
			StartYear: 2016,
		},
	},
	659010: {
		{
			Address:   "胡杨河市",
			StartYear: 2019,
		},
	},
	810000: {
		{
			Address: "香港特别行政区",
		},
	},
	820000: {
		{
			Address: "澳门特别行政区",
		},
	},
	830000: {
		{
			Address: "台湾省",
		},
	},
	// 行政区划代码（地址码）补充数据
	110100: {
		{
			Address: "市辖区",
		},
	},
	110200: {
		{
			Address: "县",
			EndYear: 2014,
		},
	},
	120100: {
		{
			Address: "市辖区",
		},
	},
	120200: {
		{
			Address: "县",
			EndYear: 2015,
		},
	},
	130101: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	130201: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	130301: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	130401: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	130501: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	130601: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	130701: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	130801: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	130901: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	131001: {
		{
			Address:   "市辖区",
			StartYear: 1988,
		},
	},
	131101: {
		{
			Address:   "市辖区",
			StartYear: 1996,
		},
	},
	139000: {
		{
			Address:   "省直辖县级行政单位",
			StartYear: 1988,
			EndYear:   1988,
		},
	},
	140101: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	140120: {
		{
			Address: "市区",
			EndYear: 1982,
		},
	},
	140201: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	140301: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	140401: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	140501: {
		{
			Address:   "市辖区",
			StartYear: 1985,
		},
	},
	140511: {
		{
			Address:   "郊区",
			StartYear: 1985,
			EndYear:   1995,
		},
	},
	140601: {
		{
			Address:   "市辖区",
			StartYear: 1988,
		},
	},
	140701: {
		{
			Address:   "市辖区",
			StartYear: 1999,
		},
	},
	140801: {
		{
			Address:   "市辖区",
			StartYear: 2000,
		},
	},
	140901: {
		{
			Address:   "市辖区",
			StartYear: 2000,
		},
	},
	141001: {
		{
			Address:   "市辖区",
			StartYear: 2000,
		},
	},
	141101: {
		{
			Address:   "市辖区",
			StartYear: 2003,
		},
	},
	149000: {
		{
			Address:   "省直辖县级行政单位",
			StartYear: 1988,
			EndYear:   1988,
		},
	},
	150101: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	150120: {
		{
			Address: "市区",
			EndYear: 1982,
		},
	},
	150201: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	150220: {
		{
			Address: "市区",
			EndYear: 1982,
		},
	},
	150301: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	150401: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	150427: {
		{
			Address:   "赤峰县",
			StartYear: 1983,
			EndYear:   1982,
		},
	},
	150501: {
		{
			Address:   "市辖区",
			StartYear: 1999,
		},
	},
	150601: {
		{
			Address:   "市辖区",
			StartYear: 2001,
		},
	},
	150701: {
		{
			Address:   "市辖区",
			StartYear: 2001,
		},
	},
	150801: {
		{
			Address:   "市辖区",
			StartYear: 2003,
		},
	},
	150901: {
		{
			Address:   "市辖区",
			StartYear: 2003,
		},
	},
	210101: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	210120: {
		{
			Address: "市区",
			EndYear: 1982,
		},
	},
	210201: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	210220: {
		{
			Address: "市区",
			EndYear: 1982,
		},
	},
	210301: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	210320: {
		{
			Address: "市区",
			EndYear: 1982,
		},
	},
	210401: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	210420: {
		{
			Address: "市区",
			EndYear: 1982,
		},
	},
	210501: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	210520: {
		{
			Address: "市区",
			EndYear: 1982,
		},
	},
	210601: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	210620: {
		{
			Address: "市区",
			EndYear: 1982,
		},
	},
	210701: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	210720: {
		{
			Address: "市区",
			EndYear: 1982,
		},
	},
	210801: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	210812: {
		{
			Address:   "鲅鱼圈区",
			StartYear: 1983,
			EndYear:   1985,
		},
	},
	210820: {
		{
			Address: "市区",
			EndYear: 1982,
		},
	},
	210901: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	210920: {
		{
			Address: "市区",
			EndYear: 1982,
		},
	},
	211001: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	211020: {
		{
			Address: "市区",
			EndYear: 1982,
		},
	},
	211101: {
		{
			Address:   "市辖区",
			StartYear: 1984,
		},
	},
	211201: {
		{
			Address:   "市辖区",
			StartYear: 1984,
		},
	},
	211301: {
		{
			Address:   "市辖区",
			StartYear: 1984,
		},
	},
	211401: {
		{
			Address:   "市辖区",
			StartYear: 1989,
		},
	},
	219000: {
		{
			Address:   "省直辖县级行政单位",
			StartYear: 1986,
			EndYear:   1988,
		},
	},
	220101: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	220120: {
		{
			Address: "市区",
			EndYear: 1982,
		},
	},
	220201: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	220220: {
		{
			Address: "市区",
			EndYear: 1982,
		},
	},
	220501: {
		{
			Address:   "市辖区",
			StartYear: 1986,
		},
	},
	220601: {
		{
			Address:   "市辖区",
			StartYear: 1986,
		},
	},
	220701: {
		{
			Address:   "市辖区",
			StartYear: 1992,
		},
	},
	220801: {
		{
			Address:   "市辖区",
			StartYear: 1993,
		},
	},
	229000: {
		{
			Address:   "省直辖县级行政单位",
			StartYear: 1986,
			EndYear:   1988,
		},
	},
	230101: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	230201: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	230301: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	230401: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	230501: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	230601: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	230701: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	230720: {
		{
			Address: "市区",
			EndYear: 1982,
		},
	},
	230801: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	230901: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	231001: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	231019: {
		{
			Address:   "镜泊湖市",
			StartYear: 1986,
			EndYear:   1986,
		},
	},
	231082: {
		{
			Address:   "密山市",
			StartYear: 1989,
			EndYear:   1995,
		},
	},
	231201: {
		{
			Address:   "市辖区",
			StartYear: 1999,
		},
	},
	232434: {
		{
			Address:   "友谊县",
			StartYear: 1984,
			EndYear:   1983,
		},
	},
	232702: {
		{
			Address:   "松岭区",
			StartYear: 1981,
			EndYear:   2017,
		},
	},
	232703: {
		{
			Address:   "新林区",
			StartYear: 1981,
			EndYear:   2017,
		},
	},
	232704: {
		{
			Address:   "呼中区",
			StartYear: 1981,
			EndYear:   2017,
		},
	},
	232761: {
		{
			Address:   "加格达奇区",
			StartYear: 2018,
		},
	},
	232762: {
		{
			Address:   "松岭区",
			StartYear: 2018,
		},
	},
	232763: {
		{
			Address:   "新林区",
			StartYear: 2018,
		},
	},
	232764: {
		{
			Address:   "呼中区",
			StartYear: 2018,
		},
	},
	239000: {
		{
			Address:   "省直辖县级行政单位",
			StartYear: 1986,
			EndYear:   1988,
		},
	},
	310100: {
		{
			Address: "市辖区",
		},
	},
	310200: {
		{
			Address: "县",
			EndYear: 2015,
		},
	},
	320101: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	320120: {
		{
			Address: "市区",
			EndYear: 1982,
		},
	},
	320201: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	320212: {
		{
			Address:   "马山区",
			StartYear: 1987,
			EndYear:   1999,
		},
	},
	320301: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	320401: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	320501: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	320520: {
		{
			Address:   "常熟市",
			StartYear: 1983,
			EndYear:   1985,
		},
	},
	320601: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	320701: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	320801: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	320901: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	321001: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	321020: {
		{
			Address:   "泰州市",
			StartYear: 1983,
			EndYear:   1985,
		},
	},
	321101: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	329000: {
		{
			Address:   "省直辖县级行政单位",
			StartYear: 1986,
			EndYear:   1988,
		},
	},
	330101: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	330113: {
		{
			Address:   "临平区",
			StartYear: 2021,
		},
	},
	330114: {
		{
			Address:   "钱塘区",
			StartYear: 2021,
		},
	},
	330120: {
		{
			Address: "市区",
			EndYear: 1982,
		},
	},
	330201: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	330220: {
		{
			Address: "市区",
			EndYear: 1982,
		},
	},
	330320: {
		{
			Address: "市区",
			EndYear: 1982,
		},
	},
	330401: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	330501: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	330601: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	330719: {
		{
			Address:   "兰溪市",
			StartYear: 1985,
			EndYear:   1985,
		},
	},
	330801: {
		{
			Address:   "市辖区",
			StartYear: 1985,
		},
	},
	330901: {
		{
			Address:   "市辖区",
			StartYear: 1987,
		},
	},
	331001: {
		{
			Address:   "市辖区",
			StartYear: 1994,
		},
	},
	331101: {
		{
			Address:   "市辖区",
			StartYear: 2000,
		},
	},
	332532: {
		{
			Address:   "龙游县",
			StartYear: 1983,
			EndYear:   1984,
		},
	},
	332533: {
		{
			Address:   "磐安县",
			StartYear: 1983,
			EndYear:   1984,
		},
	},
	332628: {
		{
			Address:   "松阳县",
			StartYear: 1982,
			EndYear:   1999,
		},
	},
	332629: {
		{
			Address:   "景宁畲族自治县",
			StartYear: 1984,
			EndYear:   1999,
		},
	},
	332702: {
		{
			Address:   "临海市",
			StartYear: 1986,
			EndYear:   1993,
		},
	},
	332703: {
		{
			Address:   "黄岩市",
			StartYear: 1989,
			EndYear:   1993,
		},
	},
	332704: {
		{
			Address:   "温岭市",
			StartYear: 1994,
			EndYear:   1993,
		},
	},
	339000: {
		{
			Address:   "省直辖县级行政单位",
			StartYear: 1986,
			EndYear:   1988,
		},
	},
	340101: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	340120: {
		{
			Address: "市区",
			EndYear: 1982,
		},
	},
	340201: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	340220: {
		{
			Address: "市区",
			EndYear: 1982,
		},
	},
	340301: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	340401: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	340420: {
		{
			Address: "市区",
			EndYear: 1982,
		},
	},
	340501: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	340601: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	340620: {
		{
			Address: "市区",
			EndYear: 1982,
		},
	},
	340701: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	340720: {
		{
			Address: "市区",
			EndYear: 1982,
		},
	},
	340801: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	340900: {
		{
			Address:   "省直辖行政单位",
			StartYear: 1983,
			EndYear:   1986,
		},
	},
	341001: {
		{
			Address:   "市辖区",
			StartYear: 1987,
		},
	},
	341101: {
		{
			Address:   "市辖区",
			StartYear: 1992,
		},
	},
	341123: {
		{
			Address: "滁县",
		},
	},
	341201: {
		{
			Address:   "市辖区",
			StartYear: 1996,
		},
	},
	341301: {
		{
			Address:   "市辖区",
			StartYear: 1998,
		},
	},
	341401: {
		{
			Address:   "市辖区",
			StartYear: 1999,
			EndYear:   2010,
		},
	},
	341501: {
		{
			Address:   "市辖区",
			StartYear: 1999,
		},
	},
	341601: {
		{
			Address:   "市辖区",
			StartYear: 2000,
		},
	},
	341701: {
		{
			Address:   "市辖区",
			StartYear: 2000,
		},
	},
	341801: {
		{
			Address:   "市辖区",
			StartYear: 2000,
		},
	},
	350101: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	350120: {
		{
			Address: "市区",
			EndYear: 1982,
		},
	},
	350201: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	350220: {
		{
			Address: "市区",
			EndYear: 1982,
		},
	},
	350301: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	350401: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	350601: {
		{
			Address:   "市辖区",
			StartYear: 1985,
		},
	},
	350701: {
		{
			Address:   "市辖区",
			StartYear: 1994,
		},
	},
	350801: {
		{
			Address:   "市辖区",
			StartYear: 1996,
		},
	},
	350901: {
		{
			Address:   "市辖区",
			StartYear: 1999,
		},
	},
	352105: {
		{
			Address:   "建阳市",
			StartYear: 1994,
			EndYear:   1993,
		},
	},
	359000: {
		{
			Address:   "省直辖县级行政单位",
			StartYear: 1986,
			EndYear:   1988,
		},
	},
	360101: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	360120: {
		{
			Address: "市区",
			EndYear: 1982,
		},
	},
	360201: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	360301: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	360401: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	360501: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	360601: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	360701: {
		{
			Address:   "市辖区",
			StartYear: 1998,
		},
	},
	360801: {
		{
			Address:   "市辖区",
			StartYear: 2000,
		},
	},
	360901: {
		{
			Address:   "市辖区",
			StartYear: 2000,
		},
	},
	361001: {
		{
			Address:   "市辖区",
			StartYear: 2000,
		},
	},
	361101: {
		{
			Address:   "市辖区",
			StartYear: 2000,
		},
	},
	362502: {
		{
			Address:   "临川市",
			StartYear: 1987,
			EndYear:   1999,
		},
	},
	370101: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	370120: {
		{
			Address: "市区",
			EndYear: 1982,
		},
	},
	370201: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	370220: {
		{
			Address: "市区",
			EndYear: 1982,
		},
	},
	370301: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	370401: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	370420: {
		{
			Address: "市区",
			EndYear: 1982,
		},
	},
	370501: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	370601: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	370621: {
		{
			Address:   "福山县",
			StartYear: 1983,
			EndYear:   1982,
		},
	},
	370626: {
		{
			Address:   "莱西县",
			StartYear: 1983,
			EndYear:   1982,
		},
	},
	370701: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	370801: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	370821: {
		{
			Address:   "济宁县",
			StartYear: 1983,
			EndYear:   1982,
		},
	},
	370824: {
		{
			Address:   "泗水县",
			StartYear: 1983,
			EndYear:   1982,
		},
	},
	370919: {
		{
			Address:   "莱芜市",
			StartYear: 1985,
			EndYear:   1985,
		},
	},
	370920: {
		{
			Address:   "新泰市",
			StartYear: 1985,
			EndYear:   1985,
		},
	},
	371001: {
		{
			Address:   "市辖区",
			StartYear: 1987,
		},
	},
	371101: {
		{
			Address:   "市辖区",
			StartYear: 1992,
		},
	},
	371201: {
		{
			Address:   "市辖区",
			StartYear: 1992,
			EndYear:   2018,
		},
	},
	371301: {
		{
			Address:   "市辖区",
			StartYear: 1994,
		},
	},
	371401: {
		{
			Address:   "市辖区",
			StartYear: 1994,
		},
	},
	371501: {
		{
			Address:   "市辖区",
			StartYear: 1997,
		},
	},
	371601: {
		{
			Address:   "市辖区",
			StartYear: 2000,
		},
	},
	371701: {
		{
			Address:   "市辖区",
			StartYear: 2000,
		},
	},
	372802: {
		{
			Address:   "日照市",
			StartYear: 1985,
			EndYear:   1988,
		},
	},
	379000: {
		{
			Address:   "省直辖县级行政单位",
			StartYear: 1986,
			EndYear:   1988,
		},
	},
	410101: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	410120: {
		{
			Address: "市区",
			EndYear: 1982,
		},
	},
	410201: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	410301: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	410307: {
		{
			Address:   "偃师区",
			StartYear: 2021,
		},
	},
	410308: {
		{
			Address:   "孟津区",
			StartYear: 2021,
		},
	},
	410401: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	410501: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	410601: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	410701: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	410801: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	410901: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	411001: {
		{
			Address:   "市辖区",
			StartYear: 1986,
		},
	},
	411101: {
		{
			Address:   "市辖区",
			StartYear: 1986,
		},
	},
	411201: {
		{
			Address:   "市辖区",
			StartYear: 1986,
		},
	},
	411219: {
		{
			Address:   "义马市",
			StartYear: 1986,
			EndYear:   1985,
		},
	},
	411301: {
		{
			Address:   "市辖区",
			StartYear: 1994,
		},
	},
	411401: {
		{
			Address:   "市辖区",
			StartYear: 1997,
		},
	},
	411501: {
		{
			Address:   "市辖区",
			StartYear: 1998,
		},
	},
	411601: {
		{
			Address:   "市辖区",
			StartYear: 2000,
		},
	},
	411701: {
		{
			Address:   "市辖区",
			StartYear: 2000,
		},
	},
	412535: {
		{
			Address: "义马矿区",
			EndYear: 1980,
		},
	},
	419000: {
		{
			Address:   "省直辖县级行政单位",
			StartYear: 1986,
		},
	},
	420101: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	420120: {
		{
			Address: "市区",
			EndYear: 1982,
		},
	},
	420201: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	420220: {
		{
			Address: "市区",
			EndYear: 1982,
		},
	},
	420301: {
		{
			Address:   "市辖区",
			StartYear: 1984,
		},
	},
	420501: {
		{
			Address:   "市辖区",
			StartYear: 1986,
		},
	},
	420522: {
		{
			Address: "宜都县",
		},
	},
	420524: {
		{
			Address: "当阳县",
		},
	},
	420601: {
		{
			Address:   "市辖区",
			StartYear: 1984,
		},
	},
	420605: {
		{
			Address:   "郊区",
			StartYear: 1984,
			EndYear:   1994,
		},
	},
	420619: {
		{
			Address:   "随州市",
			StartYear: 1983,
			EndYear:   1985,
		},
	},
	420620: {
		{
			Address:   "老河口市",
			StartYear: 1983,
			EndYear:   1985,
		},
	},
	420681: {
		{
			Address:   "随州市",
			StartYear: 1989,
			EndYear:   1993,
		},
	},
	420701: {
		{
			Address:   "市辖区",
			StartYear: 1987,
		},
	},
	420801: {
		{
			Address:   "市辖区",
			StartYear: 1985,
		},
	},
	421021: {
		{
			Address:   "松滋县",
			StartYear: 1994,
			EndYear:   1994,
		},
	},
	421084: {
		{
			Address: "天门市",
		},
	},
	421085: {
		{
			Address: "潜江市",
		},
	},
	421086: {
		{
			Address: "仙桃市",
		},
	},
	421101: {
		{
			Address:   "市辖区",
			StartYear: 1995,
		},
	},
	421201: {
		{
			Address:   "市辖区",
			StartYear: 1998,
		},
	},
	421301: {
		{
			Address:   "市辖区",
			StartYear: 2000,
		},
	},
	421302: {
		{
			Address:   "曾都区",
			StartYear: 2000,
			EndYear:   2008,
		},
	},
	422900: {
		{
			Address:   "林区",
			StartYear: 1989,
			EndYear:   1993,
		},
		{
			Address:   "林区",
			StartYear: 1983,
			EndYear:   1985,
		},
	},
	429000: {
		{
			Address:   "省直辖县级行政单位",
			StartYear: 1986,
		},
	},
	430101: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	430120: {
		{
			Address: "市区",
			EndYear: 1982,
		},
	},
	430201: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	430220: {
		{
			Address: "市区",
			EndYear: 1982,
		},
	},
	430301: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	430312: {
		{
			Address:   "韶山区",
			StartYear: 1983,
			EndYear:   1987,
		},
	},
	430401: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	430501: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	430601: {
		{
			Address:   "市辖区",
			StartYear: 1984,
		},
	},
	430625: {
		{
			Address:   "汨罗县",
			StartYear: 1986,
			EndYear:   1986,
		},
	},
	430701: {
		{
			Address:   "市辖区",
			StartYear: 1988,
		},
	},
	430727: {
		{
			Address:   "慈利县",
			StartYear: 1988,
			EndYear:   1987,
		},
	},
	430801: {
		{
			Address:   "市辖区",
			StartYear: 1988,
		},
	},
	430901: {
		{
			Address:   "市辖区",
			StartYear: 1994,
		},
	},
	431001: {
		{
			Address:   "市辖区",
			StartYear: 1994,
		},
	},
	431101: {
		{
			Address:   "市辖区",
			StartYear: 1995,
		},
	},
	431201: {
		{
			Address:   "市辖区",
			StartYear: 1997,
		},
	},
	431301: {
		{
			Address:   "市辖区",
			StartYear: 1999,
		},
	},
	432128: {
		{
			Address: "韶山区",
			EndYear: 1982,
		},
	},
	439000: {
		{
			Address:   "省直辖县级行政单位",
			StartYear: 1986,
			EndYear:   1988,
		},
	},
	440101: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	440120: {
		{
			Address: "市区",
			EndYear: 1982,
		},
	},
	440201: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	440220: {
		{
			Address: "市区",
			EndYear: 1982,
		},
	},
	440301: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	440302: {
		{
			Address:   "沙头角区",
			StartYear: 1988,
			EndYear:   1989,
		},
	},
	440320: {
		{
			Address: "市区",
			EndYear: 1982,
		},
	},
	440401: {
		{
			Address:   "市辖区",
			StartYear: 1984,
		},
	},
	440501: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	440520: {
		{
			Address:   "潮州市",
			StartYear: 1983,
			EndYear:   1985,
		},
	},
	440581: {
		{
			Address:   "潮州市",
			StartYear: 1989,
			EndYear:   1990,
		},
	},
	440601: {
		{
			Address:   "市辖区",
			StartYear: 1984,
		},
	},
	440701: {
		{
			Address:   "市辖区",
			StartYear: 1984,
		},
	},
	440801: {
		{
			Address:   "市辖区",
			StartYear: 1984,
		},
	},
	440901: {
		{
			Address:   "市辖区",
			StartYear: 1984,
		},
	},
	441201: {
		{
			Address:   "市辖区",
			StartYear: 1988,
		},
	},
	441281: {
		{
			Address:   "云浮市",
			StartYear: 1992,
			EndYear:   1993,
		},
	},
	441282: {
		{
			Address:   "罗定市",
			StartYear: 1993,
			EndYear:   1993,
		},
	},
	441301: {
		{
			Address:   "市辖区",
			StartYear: 1988,
		},
	},
	441401: {
		{
			Address:   "市辖区",
			StartYear: 1988,
		},
	},
	441501: {
		{
			Address:   "市辖区",
			StartYear: 1988,
		},
	},
	441601: {
		{
			Address:   "市辖区",
			StartYear: 1988,
		},
	},
	441701: {
		{
			Address:   "市辖区",
			StartYear: 1988,
		},
	},
	441801: {
		{
			Address:   "市辖区",
			StartYear: 1988,
		},
	},
	442137: {
		{
			Address:   "西南中沙群岛办事处",
			StartYear: 1987,
			EndYear:   1987,
		},
	},
	442229: {
		{
			Address:   "西南中沙群岛办事处",
			StartYear: 1984,
			EndYear:   1986,
		},
	},
	442701: {
		{
			Address: "佛山市",
			EndYear: 1982,
		},
	},
	442702: {
		{
			Address: "江门市",
			EndYear: 1982,
		},
	},
	442901: {
		{
			Address: "湛江市",
			EndYear: 1982,
		},
	},
	442902: {
		{
			Address: "茂名市",
			EndYear: 1982,
		},
	},
	445101: {
		{
			Address:   "市辖区",
			StartYear: 1991,
		},
	},
	445201: {
		{
			Address:   "市辖区",
			StartYear: 1991,
		},
	},
	445301: {
		{
			Address:   "市辖区",
			StartYear: 1994,
		},
	},
	449000: {
		{
			Address:   "省直辖县级行政单位",
			StartYear: 1986,
			EndYear:   1988,
		},
	},
	450101: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	450111: {
		{
			Address:   "郊区",
			StartYear: 1984,
			EndYear:   2000,
		},
	},
	450201: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	450211: {
		{
			Address:   "郊区",
			StartYear: 1984,
			EndYear:   2001,
		},
	},
	450301: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	450320: {
		{
			Address: "市区",
			EndYear: 1982,
		},
	},
	450401: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	450411: {
		{
			Address:   "郊区",
			StartYear: 1983,
			EndYear:   2002,
		},
	},
	450501: {
		{
			Address:   "市辖区",
			StartYear: 1984,
		},
	},
	450511: {
		{
			Address:   "郊区",
			StartYear: 1984,
			EndYear:   1993,
		},
	},
	450601: {
		{
			Address:   "市辖区",
			StartYear: 1993,
		},
	},
	450701: {
		{
			Address:   "市辖区",
			StartYear: 1994,
		},
	},
	450801: {
		{
			Address:   "市辖区",
			StartYear: 1995,
		},
	},
	450901: {
		{
			Address:   "市辖区",
			StartYear: 1997,
		},
	},
	451001: {
		{
			Address:   "市辖区",
			StartYear: 2002,
		},
	},
	451101: {
		{
			Address:   "市辖区",
			StartYear: 2002,
		},
	},
	451201: {
		{
			Address:   "市辖区",
			StartYear: 2002,
		},
	},
	451301: {
		{
			Address:   "市辖区",
			StartYear: 2002,
		},
	},
	451401: {
		{
			Address:   "市辖区",
			StartYear: 2002,
		},
	},
	460037: {
		{
			Address:   "西沙群岛",
			StartYear: 1988,
			EndYear:   2001,
		},
	},
	460038: {
		{
			Address:   "南沙群岛",
			StartYear: 1988,
			EndYear:   2001,
		},
	},
	460039: {
		{
			Address:   "中沙群岛的岛礁及其海域",
			StartYear: 1988,
			EndYear:   2001,
		},
	},
	460101: {
		{
			Address:   "市辖区",
			StartYear: 1990,
		},
	},
	460201: {
		{
			Address:   "市辖区",
			StartYear: 2014,
		},
	},
	460301: {
		{
			Address:   "市辖区",
			StartYear: 2020,
		},
	},
	460302: {
		{
			Address:   "西沙区",
			StartYear: 2020,
		},
	},
	460303: {
		{
			Address:   "南沙区",
			StartYear: 2020,
		},
	},
	469000: {
		{
			Address:   "省直辖县级行政单位",
			StartYear: 2002,
		},
	},
	469004: {
		{
			Address:   "琼山市",
			StartYear: 2002,
			EndYear:   2001,
		},
	},
	469031: {
		{
			Address:   "西沙群岛",
			StartYear: 2002,
			EndYear:   2011,
		},
	},
	469032: {
		{
			Address:   "南沙群岛",
			StartYear: 2002,
			EndYear:   2011,
		},
	},
	469033: {
		{
			Address:   "中沙群岛的岛礁及其海域",
			StartYear: 2002,
			EndYear:   2011,
		},
	},
	500100: {
		{
			Address:   "市辖区",
			StartYear: 1997,
		},
	},
	500200: {
		{
			Address:   "县",
			StartYear: 1997,
		},
	},
	500300: {
		{
			Address:   "市",
			StartYear: 1997,
			EndYear:   2005,
		},
	},
	510101: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	510120: {
		{
			Address: "市区",
			EndYear: 1982,
		},
	},
	510201: {
		{
			Address:   "市辖区",
			StartYear: 1983,
			EndYear:   1996,
		},
	},
	510220: {
		{
			Address: "市区",
			EndYear: 1982,
		},
	},
	510301: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	510320: {
		{
			Address: "市区",
			EndYear: 1982,
		},
	},
	510401: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	510420: {
		{
			Address: "市区",
			EndYear: 1982,
		},
	},
	510501: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	510601: {
		{
			Address:   "市辖区",
			StartYear: 1984,
		},
	},
	510701: {
		{
			Address:   "市辖区",
			StartYear: 1985,
		},
	},
	510801: {
		{
			Address:   "市辖区",
			StartYear: 1985,
		},
	},
	510901: {
		{
			Address:   "市辖区",
			StartYear: 1985,
		},
	},
	511001: {
		{
			Address:   "市辖区",
			StartYear: 1985,
		},
	},
	511101: {
		{
			Address:   "市辖区",
			StartYear: 1985,
		},
	},
	511201: {
		{
			Address:   "市辖区",
			StartYear: 1992,
			EndYear:   1996,
		},
	},
	511301: {
		{
			Address:   "市辖区",
			StartYear: 1993,
		},
	},
	511401: {
		{
			Address:   "市辖区",
			StartYear: 1995,
			EndYear:   1996,
		},
		{
			Address:   "市辖区",
			StartYear: 2000,
		},
	},
	511481: {
		{
			Address:   "南川市",
			StartYear: 1995,
			EndYear:   1996,
		},
	},
	511501: {
		{
			Address:   "市辖区",
			StartYear: 1996,
		},
	},
	511601: {
		{
			Address:   "市辖区",
			StartYear: 1998,
		},
	},
	511701: {
		{
			Address:   "市辖区",
			StartYear: 1999,
		},
	},
	511801: {
		{
			Address:   "市辖区",
			StartYear: 2000,
		},
	},
	511901: {
		{
			Address:   "市辖区",
			StartYear: 2000,
		},
	},
	512001: {
		{
			Address:   "市辖区",
			StartYear: 2000,
		},
	},
	517001: {
		{
			Address: "市辖区",
		},
	},
	519000: {
		{
			Address:   "省直辖县级行政单位",
			StartYear: 1988,
			EndYear:   1988,
		},
	},
	520101: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	520301: {
		{
			Address:   "市辖区",
			StartYear: 1997,
		},
	},
	520401: {
		{
			Address:   "市辖区",
			StartYear: 2000,
		},
	},
	520501: {
		{
			Address:   "市辖区",
			StartYear: 2011,
		},
	},
	520601: {
		{
			Address:   "市辖区",
			StartYear: 2011,
		},
	},
	530101: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	530120: {
		{
			Address: "市区",
			EndYear: 1982,
		},
	},
	530201: {
		{
			Address:   "市辖区",
			StartYear: 1983,
			EndYear:   1997,
		},
	},
	530301: {
		{
			Address:   "市辖区",
			StartYear: 1997,
		},
	},
	530401: {
		{
			Address:   "市辖区",
			StartYear: 1997,
		},
	},
	530501: {
		{
			Address:   "市辖区",
			StartYear: 2000,
		},
	},
	530601: {
		{
			Address:   "市辖区",
			StartYear: 2001,
		},
	},
	530701: {
		{
			Address:   "市辖区",
			StartYear: 2002,
		},
	},
	530801: {
		{
			Address:   "市辖区",
			StartYear: 2003,
		},
	},
	530901: {
		{
			Address:   "市辖区",
			StartYear: 2003,
		},
	},
	540101: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	540120: {
		{
			Address: "市区",
			EndYear: 1982,
		},
	},
	540201: {
		{
			Address:   "市辖区",
			StartYear: 2014,
		},
	},
	540301: {
		{
			Address:   "市辖区",
			StartYear: 2014,
		},
	},
	540401: {
		{
			Address:   "市辖区",
			StartYear: 2015,
		},
	},
	540501: {
		{
			Address:   "市辖区",
			StartYear: 2016,
		},
	},
	540601: {
		{
			Address:   "市辖区",
			StartYear: 2017,
		},
	},
	610101: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	610120: {
		{
			Address: "市区",
			EndYear: 1982,
		},
	},
	610201: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	610220: {
		{
			Address: "市区",
			EndYear: 1982,
		},
	},
	610301: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	610320: {
		{
			Address: "市区",
			EndYear: 1982,
		},
	},
	610501: {
		{
			Address:   "市辖区",
			StartYear: 1994,
		},
	},
	610601: {
		{
			Address:   "市辖区",
			StartYear: 1996,
		},
	},
	610701: {
		{
			Address:   "市辖区",
			StartYear: 1996,
		},
	},
	610801: {
		{
			Address:   "市辖区",
			StartYear: 1999,
		},
	},
	610901: {
		{
			Address:   "市辖区",
			StartYear: 2000,
		},
	},
	611001: {
		{
			Address:   "市辖区",
			StartYear: 2001,
		},
	},
	620101: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	620120: {
		{
			Address: "市区",
			EndYear: 1982,
		},
	},
	620201: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	620301: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	620320: {
		{
			Address:   "市区",
			StartYear: 1981,
			EndYear:   1982,
		},
	},
	620401: {
		{
			Address:   "市辖区",
			StartYear: 1985,
		},
	},
	620501: {
		{
			Address:   "市辖区",
			StartYear: 1985,
		},
	},
	620601: {
		{
			Address:   "市辖区",
			StartYear: 2001,
		},
	},
	620701: {
		{
			Address:   "市辖区",
			StartYear: 2002,
		},
	},
	620801: {
		{
			Address:   "市辖区",
			StartYear: 2002,
		},
	},
	620901: {
		{
			Address:   "市辖区",
			StartYear: 2002,
		},
	},
	621001: {
		{
			Address:   "市辖区",
			StartYear: 2002,
		},
	},
	621101: {
		{
			Address:   "市辖区",
			StartYear: 2003,
		},
	},
	621201: {
		{
			Address:   "市辖区",
			StartYear: 2004,
		},
	},
	630101: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	630120: {
		{
			Address: "市区",
			EndYear: 1982,
		},
	},
	630201: {
		{
			Address:   "市辖区",
			StartYear: 2013,
		},
	},
	632400: {
		{
			Address: "省直辖行政单位",
			EndYear: 1987,
		},
	},
	632857: {
		{
			Address: "大柴旦行政委员会",
		},
	},
	632858: {
		{
			Address: "冷湖行政委员会",
			EndYear: 2017,
		},
	},
	632859: {
		{
			Address: "茫崖行政委员会",
			EndYear: 2017,
		},
	},
	640101: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	640111: {
		{
			Address:   "郊区",
			StartYear: 1983,
			EndYear:   2001,
		},
	},
	640120: {
		{
			Address: "市区",
			EndYear: 1982,
		},
	},
	640201: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	640203: {
		{
			Address:   "石嘴山区",
			StartYear: 1983,
			EndYear:   2002,
		},
	},
	640220: {
		{
			Address: "市区",
			EndYear: 1982,
		},
	},
	640301: {
		{
			Address:   "市辖区",
			StartYear: 1998,
		},
	},
	640401: {
		{
			Address:   "市辖区",
			StartYear: 2001,
		},
	},
	640501: {
		{
			Address:   "市辖区",
			StartYear: 2003,
		},
	},
	650101: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	650120: {
		{
			Address: "市区",
			EndYear: 1982,
		},
	},
	650201: {
		{
			Address:   "市辖区",
			StartYear: 1983,
		},
	},
	650401: {
		{
			Address:   "市辖区",
			StartYear: 2015,
		},
	},
	650501: {
		{
			Address:   "市辖区",
			StartYear: 2016,
		},
	},
	659000: {
		{
			Address:   "自治区直辖县级行政单位",
			StartYear: 1986,
		},
	},
}
