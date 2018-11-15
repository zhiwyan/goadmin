package common

import (
	"fmt"
	"testing"
)

func TestJsonToStruct(t *testing.T) {
	jsonStr := `{"statusCode":0,"msg":"success","success":true,"data":[{"id":26990,"url":"http:\/\/wenba-ooo.ufile.ucloud.cn\/13438fa41522c534b324377412e4bf61.jpg"},{"id":28254,"url":"http:\/\/wenba-ooo.ufile.ucloud.cn\/388ccf53b8d3bd15547bbf0670ec1970.png"},{"id":28255,"url":"http:\/\/wenba-ooo.ufile.ucloud.cn\/75020ae9afa9c91b6cb5c5d06b0a3bf4.png"},{"id":28256,"url":"http:\/\/wenba-ooo.ufile.ucloud.cn\/6ea3040afae0f6282515e1f63b5e67e3.png"},{"id":28257,"url":"http:\/\/wenba-ooo.ufile.ucloud.cn\/80072f2b5bf50c70f9c7c37ec0bc5ab9.png"},{"id":28258,"url":"http:\/\/wenba-ooo.ufile.ucloud.cn\/5273f3b2fcc5e13eeb90c8ca1487c72f.png"},{"id":28259,"url":"http:\/\/wenba-ooo.ufile.ucloud.cn\/c67bf519f1b17fbd47c966f8929a775d.png"},{"id":28260,"url":"http:\/\/wenba-ooo.ufile.ucloud.cn\/10a63d6baa5add70352a9df6f9cbe507.png"},{"id":28261,"url":"http:\/\/wenba-ooo.ufile.ucloud.cn\/ad8a4dc6ff942d04a783821a5859403c.png"},{"id":28262,"url":"http:\/\/wenba-ooo.ufile.ucloud.cn\/37b6e58af084f297f4023b15a8a519f4.png"},{"id":28263,"url":"http:\/\/wenba-ooo.ufile.ucloud.cn\/6476ac372e3258985a58b72421e8e0e6.png"},{"id":28264,"url":"http:\/\/wenba-ooo.ufile.ucloud.cn\/eac616c0304f21f22edb0c82a12f7198.png"},{"id":28265,"url":"http:\/\/wenba-ooo.ufile.ucloud.cn\/690214bf92b20d99a4458714c6bdd9fe.png"},{"id":28266,"url":"http:\/\/wenba-ooo.ufile.ucloud.cn\/c86ee69a5e56400fc8cf5ef66d71e576.png"},{"id":28267,"url":"http:\/\/wenba-ooo.ufile.ucloud.cn\/84652d24615edc1e15bc95f283eeb2b9.png"},{"id":28268,"url":"http:\/\/wenba-ooo.ufile.ucloud.cn\/de7ed62458b792f76b84f879846abdac.png"},{"id":28269,"url":"http:\/\/wenba-ooo.ufile.ucloud.cn\/140eca6147db1a2b4a6f0bd820636703.png"},{"id":28270,"url":"http:\/\/wenba-ooo.ufile.ucloud.cn\/4eb5d998f59bbc87f7e41e9395bd38f2.png"},{"id":28271,"url":"http:\/\/wenba-ooo.ufile.ucloud.cn\/ebaac2805ba2a92192d6a809826cb2c0.png"},{"id":28272,"url":"http:\/\/wenba-ooo.ufile.ucloud.cn\/0544b4723e483b44a888121e59813c6a.png"},{"id":28273,"url":"http:\/\/wenba-ooo.ufile.ucloud.cn\/3ee2622b1f6a6b7f70aa12c9d8ecdeae.png"},{"id":28274,"url":"http:\/\/wenba-ooo.ufile.ucloud.cn\/16134a4b9414b8fd8549675762215881.png"},{"id":28275,"url":"http:\/\/wenba-ooo.ufile.ucloud.cn\/cf7f119e9f67c79b336232c7e449b1bd.png"},{"id":28276,"url":"http:\/\/wenba-ooo.ufile.ucloud.cn\/980ee6f6b1e93dd4a1b677b6d079cd01.png"},{"id":28277,"url":"http:\/\/wenba-ooo.ufile.ucloud.cn\/f2caea10027577761810bc932460e982.png"},{"id":28278,"url":"http:\/\/wenba-ooo.ufile.ucloud.cn\/0aaee516d90345340bdd8149f640fca5.png"},{"id":28279,"url":"http:\/\/wenba-ooo.ufile.ucloud.cn\/d7035c2875098bbad44ceec437d5b11f.png"},{"id":28280,"url":"http:\/\/wenba-ooo.ufile.ucloud.cn\/fa1fb25bbd6be216a1830195cbba8590.png"},{"id":28281,"url":"http:\/\/wenba-ooo.ufile.ucloud.cn\/5b3a6aa6f1b7a26b467bed03a33434b0.png"}]}`
	JsonToStruct([]byte(jsonStr))
}

func TestInSliceString(t *testing.T) {
	aaa := []string{"11", "12", "13", "14", "15", "16"}

	aStr := "12"
	aStrIn := InSliceString(aaa, aStr)
	if aStrIn == false {
		t.Error("aStr is not in")
	}

	bStr := "233r3df"
	bStrIn := InSliceString(aaa, bStr)
	if bStrIn == true {
		t.Error("bStr is in")
	}
}

func TestIsChineseChar(t *testing.T) {
	matched := IsChineseChar("中国")

	fmt.Println(matched)

	//input := struct {
	//	PhoneNo      string `json:"phone_no"`
	//	NickName     string `json:"nick_name"`
	//	Grade        string `json:"grade"`
	//	ValidateCode string `json:"validate_code"`
	//}{
	//	"ss",
	//	"韶关市公司",
	//	"ss",
	//	"ss",
	//}
	//inputJson, err := json.Marshal(input)
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Println(string(inputJson))
}

func TestET(t *testing.T) {
	start := Start()
	TestInSliceString(t)
	e := End(start)
	fmt.Println("function exec time: ", e)

}

func TestMapFilters(t *testing.T) {
	m := map[string]interface{}{
		"a": 0,
		"b": "",
		"c": "c",
	}

	m1 := MapFilters(m, ZeroFilter)
	fmt.Println("##################: ", m1)
}
