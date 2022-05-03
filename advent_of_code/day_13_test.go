package adventofcode

import (
	"fmt"
	"testing"
)

var coordinates = []string{
	"1053,618",
	"966,812",
	"363,813",
	"663,488",
	"582,93",
	"112,175",
	"887,158",
	"647,406",
	"1049,281",
	"820,632",
	"828,429",
	"1208,203",
	"604,15",
	"308,721",
	"790,380",
	"1007,665",
	"60,456",
	"460,507",
	"192,639",
	"45,180",
	"192,626",
	"627,604",
	"458,435",
	"683,738",
	"1121,745",
	"435,285",
	"902,74",
	"246,522",
	"734,609",
	"1005,861",
	"776,610",
	"910,491",
	"189,745",
	"216,362",
	"559,427",
	"8,465",
	"551,229",
	"855,163",
	"875,609",
	"658,203",
	"1243,270",
	"853,859",
	"10,457",
	"1012,471",
	"587,320",
	"666,94",
	"1205,543",
	"646,599",
	"619,305",
	"862,473",
	"73,630",
	"228,406",
	"850,410",
	"720,339",
	"338,273",
	"1300,273",
	"728,93",
	"915,432",
	"390,599",
	"460,484",
	"589,522",
	"721,837",
	"557,298",
	"402,497",
	"1130,40",
	"908,497",
	"646,156",
	"872,466",
	"1161,674",
	"470,728",
	"751,736",
	"549,385",
	"633,33",
	"1012,423",
	"858,98",
	"490,878",
	"276,351",
	"999,665",
	"216,87",
	"892,603",
	"915,574",
	"236,295",
	"328,225",
	"1273,287",
	"768,127",
	"1169,661",
	"1079,288",
	"902,820",
	"768,623",
	"797,833",
	"524,429",
	"1123,812",
	"1002,721",
	"460,36",
	"574,746",
	"139,588",
	"965,143",
	"711,810",
	"1230,605",
	"910,28",
	"395,574",
	"1176,514",
	"438,807",
	"932,787",
	"489,735",
	"144,779",
	"940,193",
	"867,364",
	"920,508",
	"438,700",
	"7,635",
	"284,473",
	"98,700",
	"443,28",
	"1156,74",
	"1138,429",
	"716,75",
	"164,508",
	"701,714",
	"311,437",
	"256,432",
	"268,642",
	"271,733",
	"229,540",
	"180,376",
	"323,134",
	"1029,42",
	"400,476",
	"633,161",
	"733,883",
	"723,208",
	"355,626",
	"482,380",
	"793,386",
	"326,396",
	"1186,477",
	"1210,752",
	"261,165",
	"1111,778",
	"1265,714",
	"58,231",
	"1302,429",
	"95,427",
	"974,771",
	"378,787",
	"246,466",
	"187,812",
	"892,155",
	"378,15",
	"472,115",
	"1300,173",
	"627,380",
	"1278,571",
	"653,735",
	"1192,707",
	"1066,788",
	"445,403",
	"440,127",
	"431,765",
	"599,597",
	"98,787",
	"920,386",
	"559,19",
	"189,644",
	"291,276",
	"346,511",
	"291,813",
	"785,35",
	"1205,285",
	"214,91",
	"589,837",
	"310,812",
	"1044,19",
	"622,771",
	"1046,162",
	"418,333",
	"920,732",
	"805,674",
	"1289,801",
	"547,642",
	"1246,530",
	"706,15",
	"8,429",
	"1161,628",
	"1074,295",
	"124,513",
	"502,661",
	"945,63",
	"720,555",
	"432,511",
	"154,372",
	"1169,233",
	"555,639",
	"160,651",
	"377,310",
	"75,792",
	"448,421",
	"850,403",
	"366,381",
	"90,291",
	"180,40",
	"489,287",
	"1211,255",
	"822,460",
	"795,812",
	"1046,508",
	"984,498",
	"1170,686",
	"691,500",
	"669,283",
	"939,679",
	"27,626",
	"1289,462",
	"723,110",
	"146,588",
	"651,760",
	"634,411",
	"492,673",
	"341,163",
	"435,609",
	"1121,834",
	"1302,795",
	"835,213",
	"604,799",
	"228,854",
	"177,668",
	"482,429",
	"8,123",
	"1042,63",
	"982,669",
	"1283,156",
	"622,323",
	"10,621",
	"638,418",
	"373,852",
	"244,445",
	"515,82",
	"472,563",
	"966,628",
	"774,259",
	"391,93",
	"604,719",
	"633,733",
	"735,838",
	"458,877",
	"214,543",
	"638,266",
	"982,673",
	"879,681",
	"490,262",
	"900,445",
	"572,746",
	"964,487",
	"261,829",
	"408,820",
	"805,611",
	"129,523",
	"443,364",
	"118,707",
	"1123,306",
	"482,17",
	"964,511",
	"765,892",
	"612,193",
	"790,514",
	"115,28",
	"452,641",
	"301,467",
	"785,859",
	"18,855",
	"604,787",
	"472,775",
	"788,822",
	"969,163",
	"653,159",
	"736,652",
	"1212,194",
	"166,236",
	"515,677",
	"73,182",
	"1163,380",
	"624,598",
	"882,80",
	"677,733",
	"845,250",
	"1081,354",
	"134,380",
	"556,283",
	"524,465",
	"59,457",
	"236,519",
	"1150,651",
	"268,63",
	"923,217",
	"147,661",
	"1143,760",
	"507,787",
	"1307,358",
	"595,437",
	"688,323",
	"109,178",
	"852,17",
	"400,866",
	"70,333",
	"371,679",
	"646,74",
	"1158,240",
	"388,19",
	"815,358",
	"1310,221",
	"440,543",
	"957,427",
	"281,42",
	"850,627",
	"294,515",
	"266,539",
	"124,477",
	"1171,217",
	"723,494",
	"664,156",
	"52,585",
	"482,465",
	"172,148",
	"992,17",
	"1130,376",
	"1138,17",
	"336,323",
	"1297,434",
	"1033,534",
	"733,746",
	"1139,747",
	"622,547",
	"754,642",
	"1044,355",
	"711,532",
	"684,91",
	"212,219",
	"236,71",
	"1220,561",
	"735,56",
	"397,276",
	"918,331",
	"261,513",
	"105,543",
	"890,208",
	"875,698",
	"616,434",
	"672,364",
	"149,674",
	"21,462",
	"351,229",
	"1252,7",
	"892,333",
	"1251,457",
	"572,148",
	"343,885",
	"1260,767",
	"164,430",
	"922,427",
	"753,130",
	"1278,323",
	"1203,226",
	"723,574",
	"1049,596",
	"693,159",
	"187,306",
	"587,574",
	"1215,467",
	"360,801",
	"1213,130",
	"882,528",
	"1019,276",
	"1046,60",
	"1163,661",
	"242,609",
	"1302,205",
	"457,859",
	"1245,221",
	"1304,889",
	"721,372",
	"664,771",
	"219,698",
	"494,651",
	"589,57",
	"465,508",
	"927,795",
	"912,14",
	"505,348",
	"504,234",
	"1213,382",
	"462,639",
	"460,851",
	"545,674",
	"360,93",
	"77,746",
	"427,490",
	"1267,887",
	"134,541",
	"3,11",
	"1181,523",
	"581,630",
	"460,450",
	"248,285",
	"1116,627",
	"144,567",
	"21,541",
	"216,700",
	"1158,512",
	"124,829",
	"944,513",
	"574,652",
	"460,387",
	"199,163",
	"678,84",
	"72,829",
	"67,574",
	"135,738",
	"363,842",
	"519,514",
	"1005,33",
	"107,642",
	"507,339",
	"882,823",
	"1243,208",
	"940,682",
	"738,746",
	"961,661",
	"1129,723",
	"8,795",
	"590,555",
	"438,585",
	"870,35",
	"431,385",
	"1173,358",
	"277,534",
	"1212,700",
	"16,563",
	"174,382",
	"73,488",
	"239,232",
	"308,621",
	"102,243",
	"25,770",
	"736,893",
	"149,416",
	"1146,669",
	"646,820",
	"1243,574",
	"1123,588",
	"356,249",
	"392,107",
	"956,234",
	"527,330",
	"330,467",
	"947,842",
	"987,459",
	"664,599",
	"477,220",
	"1074,71",
	"648,453",
	"1029,337",
	"82,596",
	"828,514",
	"1131,213",
	"508,854",
	"848,639",
	"910,39",
	"346,63",
	"1000,194",
	"3,883",
	"1146,386",
	"8,596",
	"1304,5",
	"1129,815",
	"253,264",
	"759,665",
	"803,380",
	"684,697",
	"878,511",
	"761,543",
	"887,736",
	"513,833",
	"753,298",
	"149,266",
	"278,624",
	"654,673",
	"1047,801",
	"323,459",
	"341,64",
	"592,162",
	"460,15",
	"1171,588",
	"870,91",
	"828,140",
	"1258,540",
	"733,148",
	"388,355",
	"7,406",
	"622,541",
	"1051,494",
	"10,437",
	"1161,283",
	"1123,229",
	"100,752",
	"751,467",
	"284,193",
	"858,253",
	"117,333",
	"1158,437",
	"377,86",
	"127,703",
	"127,731",
	"10,273",
	"32,291",
	"828,380",
	"945,287",
	"199,283",
	"915,880",
	"43,7",
	"1222,397",
	"641,283",
	"714,439",
	"761,509",
	"305,681",
	"1118,178",
	"980,82",
	"1289,353",
	"97,764",
	"1240,113",
	"408,18",
	"1288,810",
	"795,588",
	"638,630",
	"448,473",
	"206,381",
	"177,674",
	"694,434",
	"92,878",
	"867,418",
	"646,295",
	"443,866",
	"902,893",
	"211,866",
	"30,408",
	"1064,820",
	"672,630",
	"187,229",
	"1220,337",
	"555,353",
	"305,33",
	"1303,635",
	"1042,700",
	"164,669",
	"1165,611",
	"6,889",
	"806,234",
	"944,381",
	"1118,255",
	"239,326",
	"725,394",
	"961,317",
	"395,763",
	"393,729",
	"1294,530",
	"797,681",
	"1203,668",
	"1091,698",
	"1084,375",
	"1302,596",
	"37,287",
	"149,628",
	"619,788",
	"1108,684",
	"223,681",
	"984,669",
	"987,134",
	"1171,431",
	"761,161",
	"902,746",
	"477,347",
	"1146,262",
	"257,618",
	"875,61",
	"350,84",
	"8,298",
	"1096,543",
	"192,66",
	"790,541",
	"1166,263",
	"1002,611",
	"1042,252",
	"102,691",
	"1123,431",
	"579,792",
	"1186,829",
	"823,760",
	"217,892",
	"688,771",
	"765,674",
	"2,691",
	"937,785",
	"211,418",
	"440,35",
	"836,376",
	"982,213",
	"562,128",
	"1094,700",
	"1258,585",
	"268,194",
	"377,584",
	"990,5",
	"915,131",
	"1213,512",
	"21,161",
	"556,194",
	"1049,829",
	"1192,187",
	"902,18",
	"257,276",
	"937,42",
	"513,61",
	"194,851",
	"1000,264",
	"67,270",
	"365,63",
	"358,710",
	"402,397",
	"124,417",
	"1237,264",
	"820,508",
	"542,127",
	"1033,808",
	"387,217",
	"683,380",
	"850,387",
	"818,262",
	"90,561",
	"264,60",
	"850,2",
	"677,353",
	"647,488",
	"1104,645",
	"761,385",
	"440,803",
	"1091,509",
	"937,852",
	"308,273",
	"1026,421",
	"202,641",
	"1260,127",
	"1079,606",
	"171,747",
	"853,404",
	"55,362",
	"790,353",
	"284,421",
	"1278,291",
	"323,211",
	"842,432",
	"577,148",
	"271,347",
	"964,831",
	"1094,428",
	"912,432",
	"1044,539",
	"803,9",
	"566,397",
	"1260,719",
	"400,418",
	"289,831",
	"1116,374",
	"758,299",
	"944,12",
	"1273,735",
	"1278,123",
	"267,229",
	"612,701",
	"708,802",
	"672,418",
	"267,677",
	"1161,546",
	"1016,877",
	"617,287",
	"1138,148",
	"659,760",
	"408,74",
	"626,803",
	"223,213",
	"545,892",
	"55,810",
	"918,563",
	"50,767",
	"947,529",
	"488,236",
	"460,520",
	"666,437",
	"1228,669",
	"1111,163",
	"803,555",
	"139,463",
	"10,9",
	"1042,194",
	"964,407",
	"172,435",
	"843,351",
	"388,539",
	"366,12",
	"154,74",
	"797,61",
	"505,674",
	"79,390",
	"525,859",
	"1230,746",
	"786,205",
	"132,822",
	"50,175",
	"974,323",
	"984,396",
	"848,252",
	"1304,441",
	"1250,691",
	"530,381",
	"552,651",
	"960,810",
	"1176,828",
	"965,751",
	"795,82",
	"202,458",
	"141,61",
	"351,665",
	"21,93",
	"1146,878",
	"16,630",
	"513,661",
	"900,876",
	"147,787",
	"604,879",
	"1116,491",
	"992,746",
	"1116,851",
	"276,787",
	"984,449",
	"733,235",
	"402,49",
	"227,703",
	"721,737",
	"883,191",
	"1231,504",
	"1282,803",
	"472,331",
	"445,340",
	"349,661",
	"972,416",
	"974,604",
	"192,642",
	"1278,347",
	"795,351",
	"833,347",
	"187,82",
	"1047,93",
	"1208,651",
	"452,884",
	"288,65",
	"32,795",
	"1034,614",
	"1064,466",
	"808,64",
	"465,250",
	"754,194",
	"902,690",
	"1091,385",
	"366,513",
	"855,778",
	"1021,63",
	"599,810",
	"964,273",
	"487,760",
	"587,624",
	"1133,674",
	"390,732",
	"599,362",
	"838,563",
	"882,814",
	"432,164",
	"912,462",
	"1026,212",
	"708,92",
	"520,380",
	"1260,787",
	"875,644",
	"189,149",
	"408,596",
	"726,113",
	"259,91",
	"1051,91",
	"259,168",
}
var actions = []string{
	"x=655",
	"y=447",
	"x=327",
	"y=223",
	"x=163",
	"y=111",
	"x=81",
	"y=55",
	"x=40",
	"y=27",
	"y=13",
	"y=6",
}

func TestDay13Part1(t *testing.T) {
	expectedRes := 1
	fmt.Println(len(coordinates))
	actions := []string{"x=655"}
	// actions := []string{"y=7"}
	if res := day13Part1(coordinates, actions); res != expectedRes {
		t.Errorf("expected res is %v,what we get is %v", expectedRes, res)
	}
}
