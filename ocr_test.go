package weapp

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"testing"
)

func TestBankCardByURL(t *testing.T) {
	server := http.NewServeMux()
	server.HandleFunc(apiBankcard, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Fatalf("Expect 'POST' get '%s'", r.Method)
		}

		path := r.URL.EscapedPath()
		if path != apiBankcard {
			t.Fatalf("Except to path '%s',get '%s'", apiBankcard, path)
		}

		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}

		queries := []string{"type", "access_token", "img_url"}
		for _, v := range queries {
			content := r.Form.Get(v)
			if content == "" {
				t.Fatalf("Params [%s] can not be empty", v)
			}
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		raw := `{
			"errcode": 0,
			"errmsg": "ok",
			"id": "622213XXXXXXXXX"
		  }`
		if _, err := w.Write([]byte(raw)); err != nil {
			t.Fatal(err)
		}
	})

	server.HandleFunc("/mediaurl", func(w http.ResponseWriter, r *http.Request) {
		filename := testIMGName
		file, err := os.Open(filename)
		if err != nil {
			t.Fatal((err))
		}
		defer file.Close()

		ext := path.Ext(filename)
		ext = ext[1:len(ext)]
		w.Header().Set("Content-Type", "image/"+ext)
		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", path.Base(filename)))
		w.WriteHeader(http.StatusOK)

		if _, err := io.Copy(w, file); err != nil {
			t.Fatal(err)
		}
	})

	ts := httptest.NewServer(server)
	defer ts.Close()

	_, err := bankCardByURL(ts.URL+apiBankcard, "mock-access-token", ts.URL+"/mediaurl", RecognizeModePhoto)
	if err != nil {
		t.Fatal(err)
	}

	_, err = bankCardByURL(ts.URL+apiBankcard, "mock-access-token", ts.URL+"/mediaurl", RecognizeModeScan)
	if err != nil {
		t.Fatal(err)
	}
}

func TestBankCard(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "POST" {
			t.Fatalf("Expect 'POST' get '%s'", r.Method)
		}

		path := r.URL.EscapedPath()
		if path != apiBankcard {
			t.Fatalf("Except to path '%s',get '%s'", apiBankcard, path)
		}

		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}

		queries := []string{"type", "access_token"}
		for _, v := range queries {
			content := r.Form.Get(v)
			if content == "" {
				t.Fatalf("Params [%s] can not be empty", v)
			}
		}

		if _, _, err := r.FormFile("img"); err != nil {
			t.Fatal(err)

		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		raw := `{
			"errcode": 0,
			"errmsg": "ok",
			"id": "622213XXXXXXXXX"
		  }`
		if _, err := w.Write([]byte(raw)); err != nil {
			t.Fatal(err)
		}
	}))
	defer ts.Close()

	_, err := bankCard(ts.URL+apiBankcard, "mock-access-token", testIMGName, RecognizeModePhoto)
	if err != nil {
		t.Fatal(err)
	}
	_, err = bankCard(ts.URL+apiBankcard, "mock-access-token", testIMGName, RecognizeModeScan)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDriverLicenseByURL(t *testing.T) {
	server := http.NewServeMux()
	server.HandleFunc(apiDrivingLicense, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Fatalf("Expect 'POST' get '%s'", r.Method)
		}

		path := r.URL.EscapedPath()
		if path != apiDrivingLicense {
			t.Fatalf("Except to path '%s',get '%s'", apiDrivingLicense, path)
		}

		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}

		queries := []string{"access_token", "img_url"}
		for _, v := range queries {
			content := r.Form.Get(v)
			if content == "" {
				t.Fatalf("Params [%s] can not be empty", v)
			}
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		raw := `{
			"errcode": 0,
			"errmsg": "ok",
			"id_num": "660601xxxxxxxx1234",
			"name": "张三",
			"sex": "男",
			"nationality": "中国",
			"address": "广东省东莞市xxxxx号",
			"birth_date": "1990-12-21",
			"issue_date": "2012-12-21",
			"car_class": "C1",
			"valid_from": "2018-07-06",
			"valid_to": "2020-07-01",
			"official_seal": "xx市公安局公安交通管理局"
		   }`
		if _, err := w.Write([]byte(raw)); err != nil {
			t.Fatal(err)
		}
	})

	server.HandleFunc("/mediaurl", func(w http.ResponseWriter, r *http.Request) {
		filename := testIMGName
		file, err := os.Open(filename)
		if err != nil {
			t.Fatal((err))
		}
		defer file.Close()

		ext := path.Ext(filename)
		ext = ext[1:len(ext)]
		w.Header().Set("Content-Type", "image/"+ext)
		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", path.Base(filename)))
		w.WriteHeader(http.StatusOK)

		if _, err := io.Copy(w, file); err != nil {
			t.Fatal(err)
		}
	})

	ts := httptest.NewServer(server)
	defer ts.Close()

	_, err := driverLicenseByURL(ts.URL+apiDrivingLicense, "mock-access-token", ts.URL+"/mediaurl")
	if err != nil {
		t.Fatal(err)
	}
}

func TestDriverLicense(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "POST" {
			t.Fatalf("Expect 'POST' get '%s'", r.Method)
		}

		path := r.URL.EscapedPath()
		if path != apiDrivingLicense {
			t.Fatalf("Except to path '%s',get '%s'", apiDrivingLicense, path)
		}

		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}

		if r.Form.Get("access_token") == "" {
			t.Fatalf("access_token can not be empty")
		}

		if _, _, err := r.FormFile("img"); err != nil {
			t.Fatal(err)

		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		raw := `{
			"errcode": 0,
			"errmsg": "ok",
			"id_num": "660601xxxxxxxx1234",
			"name": "张三",
			"sex": "男",
			"nationality": "中国",
			"address": "广东省东莞市xxxxx号",
			"birth_date": "1990-12-21",
			"issue_date": "2012-12-21",
			"car_class": "C1",
			"valid_from": "2018-07-06",
			"valid_to": "2020-07-01",
			"official_seal": "xx市公安局公安交通管理局"
		   }`
		if _, err := w.Write([]byte(raw)); err != nil {
			t.Fatal(err)
		}
	}))
	defer ts.Close()

	_, err := driverLicense(ts.URL+apiDrivingLicense, "mock-access-token", testIMGName)
	if err != nil {
		t.Fatal(err)
	}
}

func TestBusinessLicenseByURL(t *testing.T) {
	server := http.NewServeMux()
	server.HandleFunc("/cv/ocr/bizlicense", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Fatalf("Expect 'POST' get '%s'", r.Method)
		}

		if r.URL.EscapedPath() != "/cv/ocr/bizlicense" {
			t.Error("Invalid request path")
		}

		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}

		queries := []string{"access_token", "img_url"}
		for _, v := range queries {
			content := r.Form.Get(v)
			if content == "" {
				t.Fatalf("Params [%s] can not be empty", v)
			}
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		raw := `{
			"errcode": 0,
			"errmsg": "ok",
			"reg_num": "123123",                                                     
			"serial": "123123",                                                      
			"legal_representative": "张三",                                          
			"enterprise_name": "XX饮食店",                                           
			"type_of_organization": "个人经营",                                      
			"address": "XX市XX区XX路XX号",                                           
			"type_of_enterprise": "xxx",                                             
			"business_scope": "中型餐馆(不含凉菜、不含裱花蛋糕，不含生食海产品)。",  
			"registered_capital": "200万",                                           
			"paid_in_capital": "200万",                                              
			"valid_period": "2019年1月1日",                                          
			"registered_date": "2018年1月1日",                                       
			"cert_position": {                                                       
				"pos": {
					"left_top": {
						"x": 155,
						"y": 191
					},
					"right_top": {
						"x": 725,
						"y": 157
					},
					"right_bottom": {
						"x": 743,
						"y": 512
					},
					"left_bottom": {
						"x": 164,
						"y": 525
					}
				}
			},
			"img_size": {
				"w": 966,
				"h": 728
			}
		}`
		if _, err := w.Write([]byte(raw)); err != nil {
			t.Fatal(err)
		}
	})

	server.HandleFunc("/mediaurl", func(w http.ResponseWriter, r *http.Request) {
		filename := testIMGName
		file, err := os.Open(filename)
		if err != nil {
			t.Fatal((err))
		}
		defer file.Close()

		ext := path.Ext(filename)
		ext = ext[1:len(ext)]
		w.Header().Set("Content-Type", "image/"+ext)
		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", path.Base(filename)))
		w.WriteHeader(http.StatusOK)

		if _, err := io.Copy(w, file); err != nil {
			t.Fatal(err)
		}
	})

	ts := httptest.NewServer(server)
	defer ts.Close()

	_, err := businessLicenseByURL(ts.URL+apiBusinessLicense, "mock-access-token", ts.URL+"/mediaurl")
	if err != nil {
		t.Fatal(err)
	}
}

func TestBusinessLicense(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "POST" {
			t.Fatalf("Expect 'POST' get '%s'", r.Method)
		}

		if r.URL.EscapedPath() != "/cv/ocr/bizlicense" {
			t.Error("Invalid request path")
		}

		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}

		if r.Form.Get("access_token") == "" {
			t.Fatalf("access_token can not be empty")
		}

		if _, _, err := r.FormFile("img"); err != nil {
			t.Fatal(err)

		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		raw := `{
			"errcode": 0,
			"errmsg": "ok",
			"reg_num": "123123",                                                     
			"serial": "123123",                                                      
			"legal_representative": "张三",                                          
			"enterprise_name": "XX饮食店",                                           
			"type_of_organization": "个人经营",                                      
			"address": "XX市XX区XX路XX号",                                           
			"type_of_enterprise": "xxx",                                             
			"business_scope": "中型餐馆(不含凉菜、不含裱花蛋糕，不含生食海产品)。",  
			"registered_capital": "200万",                                           
			"paid_in_capital": "200万",                                              
			"valid_period": "2019年1月1日",                                          
			"registered_date": "2018年1月1日",                                       
			"cert_position": {                                                       
				"pos": {
					"left_top": {
						"x": 155,
						"y": 191
					},
					"right_top": {
						"x": 725,
						"y": 157
					},
					"right_bottom": {
						"x": 743,
						"y": 512
					},
					"left_bottom": {
						"x": 164,
						"y": 525
					}
				}
			},
			"img_size": {
				"w": 966,
				"h": 728
			}
		}`
		if _, err := w.Write([]byte(raw)); err != nil {
			t.Fatal(err)
		}
	}))
	defer ts.Close()

	_, err := businessLicense(ts.URL+apiBusinessLicense, "mock-access-token", testIMGName)
	if err != nil {
		t.Fatal(err)
	}
}

func TestPrintedTextByURL(t *testing.T) {
	server := http.NewServeMux()
	server.HandleFunc("/cv/ocr/comm", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Fatalf("Expect 'POST' get '%s'", r.Method)
		}

		if r.URL.EscapedPath() != "/cv/ocr/comm" {
			t.Error("Invalid request path")
		}

		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}

		queries := []string{"access_token", "img_url"}
		for _, v := range queries {
			content := r.Form.Get(v)
			if content == "" {
				t.Fatalf("Params [%s] can not be empty", v)
			}
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		raw := `{
			"errcode": 0,
			"errmsg": "ok",
			"items": [ 
				{
					"text": "腾讯",
					"pos": {
						"left_top": {
							"x": 575,
							"y": 519
						},
						"right_top": {
							"x": 744,
							"y": 519
						},
						"right_bottom": {
							"x": 744,
							"y": 532
						},
						"left_bottom": {
							"x": 573,
							"y": 532
						}
					}
				},
				{
					"text": "微信团队",
					"pos": {
						"left_top": {
							"x": 670,
							"y": 516
						},
						"right_top": {
							"x": 762,
							"y": 517
						},
						"right_bottom": {
							"x": 762,
							"y": 532
						},
						"left_bottom": {
							"x": 670,
							"y": 531
						}
					}
				}
			],
			"img_size": { 
				"w": 1280,
				"h": 720
			}
		}`
		if _, err := w.Write([]byte(raw)); err != nil {
			t.Fatal(err)
		}
	})

	server.HandleFunc("/mediaurl", func(w http.ResponseWriter, r *http.Request) {
		filename := testIMGName
		file, err := os.Open(filename)
		if err != nil {
			t.Fatal((err))
		}
		defer file.Close()

		ext := path.Ext(filename)
		ext = ext[1:len(ext)]
		w.Header().Set("Content-Type", "image/"+ext)
		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", path.Base(filename)))
		w.WriteHeader(http.StatusOK)

		if _, err := io.Copy(w, file); err != nil {
			t.Fatal(err)
		}
	})

	ts := httptest.NewServer(server)
	defer ts.Close()

	_, err := printedTextByURL(ts.URL+apiPrintedText, "mock-access-token", ts.URL+"/mediaurl")
	if err != nil {
		t.Fatal(err)
	}
}

func TestPrintedText(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "POST" {
			t.Fatalf("Expect 'POST' get '%s'", r.Method)
		}

		if r.URL.EscapedPath() != "/cv/ocr/comm" {
			t.Error("Invalid request path")
		}

		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}

		if r.Form.Get("access_token") == "" {
			t.Fatalf("access_token can not be empty")
		}

		if _, _, err := r.FormFile("img"); err != nil {
			t.Fatal(err)

		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		raw := `{
			"errcode": 0,
			"errmsg": "ok",
			"items": [ 
				{
					"text": "腾讯",
					"pos": {
						"left_top": {
							"x": 575,
							"y": 519
						},
						"right_top": {
							"x": 744,
							"y": 519
						},
						"right_bottom": {
							"x": 744,
							"y": 532
						},
						"left_bottom": {
							"x": 573,
							"y": 532
						}
					}
				},
				{
					"text": "微信团队",
					"pos": {
						"left_top": {
							"x": 670,
							"y": 516
						},
						"right_top": {
							"x": 762,
							"y": 517
						},
						"right_bottom": {
							"x": 762,
							"y": 532
						},
						"left_bottom": {
							"x": 670,
							"y": 531
						}
					}
				}
			],
			"img_size": { 
				"w": 1280,
				"h": 720
			}
		}`
		if _, err := w.Write([]byte(raw)); err != nil {
			t.Fatal(err)
		}
	}))
	defer ts.Close()

	_, err := printedText(ts.URL+apiPrintedText, "mock-access-token", testIMGName)
	if err != nil {
		t.Fatal(err)
	}
}

func TestIDCardByURL(t *testing.T) {
	server := http.NewServeMux()
	server.HandleFunc(apiIDCard, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Fatalf("Expect 'POST' get '%s'", r.Method)
		}

		path := r.URL.EscapedPath()
		if path != apiIDCard {
			t.Fatalf("Except to path '%s',get '%s'", apiIDCard, path)
		}

		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}

		queries := []string{"type", "access_token", "img_url"}
		for _, v := range queries {
			content := r.Form.Get(v)
			if content == "" {
				t.Fatalf("Params [%s] can not be empty", v)
			}
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		raw := `{
			"errcode": 0,
			"errmsg": "ok",
			"type": "Front",
			"id": "44XXXXXXXXXXXXXXX1"
		  }`
		if _, err := w.Write([]byte(raw)); err != nil {
			t.Fatal(err)
		}
	})

	server.HandleFunc("/mediaurl", func(w http.ResponseWriter, r *http.Request) {
		filename := testIMGName
		file, err := os.Open(filename)
		if err != nil {
			t.Fatal((err))
		}
		defer file.Close()

		ext := path.Ext(filename)
		ext = ext[1:len(ext)]
		w.Header().Set("Content-Type", "image/"+ext)
		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", path.Base(filename)))
		w.WriteHeader(http.StatusOK)

		if _, err := io.Copy(w, file); err != nil {
			t.Fatal(err)
		}
	})

	ts := httptest.NewServer(server)
	defer ts.Close()

	_, err := idCardByURL(ts.URL+apiIDCard, "mock-access-token", ts.URL+"/mediaurl", RecognizeModePhoto)
	if err != nil {
		t.Fatal(err)
	}

	_, err = idCardByURL(ts.URL+apiIDCard, "mock-access-token", ts.URL+"/mediaurl", RecognizeModeScan)
	if err != nil {
		t.Fatal(err)
	}
}

func TestIDCard(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "POST" {
			t.Fatalf("Expect 'POST' get '%s'", r.Method)
		}

		path := r.URL.EscapedPath()
		if path != apiIDCard {
			t.Fatalf("Except to path '%s',get '%s'", apiIDCard, path)
		}

		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}
		queries := []string{"type", "access_token"}
		for _, v := range queries {
			content := r.Form.Get(v)
			if content == "" {
				t.Fatalf("Params [%s] can not be empty", v)
			}
		}

		if _, _, err := r.FormFile("img"); err != nil {
			t.Fatal(err)

		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		raw := `{
			"errcode": 0,
			"errmsg": "ok",
			"type": "Front",
			"id": "44XXXXXXXXXXXXXXX1"
		  }`
		if _, err := w.Write([]byte(raw)); err != nil {
			t.Fatal(err)
		}
	}))
	defer ts.Close()

	_, err := idCard(ts.URL+apiIDCard, "mock-access-token", testIMGName, RecognizeModePhoto)
	if err != nil {
		t.Fatal(err)
	}
	_, err = idCard(ts.URL+apiIDCard, "mock-access-token", testIMGName, RecognizeModeScan)
	if err != nil {
		t.Fatal(err)
	}
}

func TestVehicleLicenseByURL(t *testing.T) {
	server := http.NewServeMux()
	server.HandleFunc(apiVehicleLicense, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Fatalf("Expect 'POST' get '%s'", r.Method)
		}

		path := r.URL.EscapedPath()
		if path != apiVehicleLicense {
			t.Fatalf("Except to path '%s',get '%s'", apiVehicleLicense, path)
		}

		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}

		queries := []string{"type", "access_token", "img_url"}
		for _, v := range queries {
			content := r.Form.Get(v)
			if content == "" {
				t.Fatalf("Params [%s] can not be empty", v)
			}
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		raw := `{
			"vhicle_type": "小型普通客⻋",
			"owner": "东莞市xxxxx机械厂",
			"addr": "广东省东莞市xxxxx号",
			"use_character": "非营运",
			"model": "江淮牌HFCxxxxxxx",
			"vin": "LJ166xxxxxxxx51",
			"engine_num": "J3xxxxx3",
			"register_date": "2018-07-06",
			"issue_date": "2018-07-01",
			"plate_num_b": "粤xxxxx",
			"record": "441xxxxxx3",
			"passengers_num": "7人",
			"total_quality": "2700kg",
			"prepare_quality": "1995kg"
		}`
		if _, err := w.Write([]byte(raw)); err != nil {
			t.Fatal(err)
		}
	})

	server.HandleFunc("/mediaurl", func(w http.ResponseWriter, r *http.Request) {
		filename := testIMGName
		file, err := os.Open(filename)
		if err != nil {
			t.Fatal((err))
		}
		defer file.Close()

		ext := path.Ext(filename)
		ext = ext[1:len(ext)]
		w.Header().Set("Content-Type", "image/"+ext)
		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", path.Base(filename)))
		w.WriteHeader(http.StatusOK)

		if _, err := io.Copy(w, file); err != nil {
			t.Fatal(err)
		}
	})

	ts := httptest.NewServer(server)
	defer ts.Close()

	_, err := vehicleLicenseByURL(ts.URL+apiVehicleLicense, "mock-access-token", ts.URL+"/mediaurl", RecognizeModePhoto)
	if err != nil {
		t.Fatal(err)
	}

	_, err = vehicleLicenseByURL(ts.URL+apiVehicleLicense, "mock-access-token", ts.URL+"/mediaurl", RecognizeModeScan)
	if err != nil {
		t.Fatal(err)
	}
}

func TestVehicleLicense(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "POST" {
			t.Fatalf("Expect 'POST' get '%s'", r.Method)
		}

		path := r.URL.EscapedPath()
		if path != apiVehicleLicense {
			t.Fatalf("Except to path '%s',get '%s'", apiVehicleLicense, path)
		}

		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}
		queries := []string{"type", "access_token"}
		for _, v := range queries {
			content := r.Form.Get(v)
			if content == "" {
				t.Fatalf("Params [%s] can not be empty", v)
			}
		}

		if _, _, err := r.FormFile("img"); err != nil {
			t.Fatal(err)

		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		raw := `{
			"vhicle_type": "小型普通客⻋",
			"owner": "东莞市xxxxx机械厂",
			"addr": "广东省东莞市xxxxx号",
			"use_character": "非营运",
			"model": "江淮牌HFCxxxxxxx",
			"vin": "LJ166xxxxxxxx51",
			"engine_num": "J3xxxxx3",
			"register_date": "2018-07-06",
			"issue_date": "2018-07-01",
			"plate_num_b": "粤xxxxx",
			"record": "441xxxxxx3",
			"passengers_num": "7人",
			"total_quality": "2700kg",
			"prepare_quality": "1995kg"
		}`
		if _, err := w.Write([]byte(raw)); err != nil {
			t.Fatal(err)
		}
	}))
	defer ts.Close()

	_, err := vehicleLicense(ts.URL+apiVehicleLicense, "mock-access-token", testIMGName, RecognizeModePhoto)
	if err != nil {
		t.Fatal(err)
	}
	_, err = vehicleLicense(ts.URL+apiVehicleLicense, "mock-access-token", testIMGName, RecognizeModeScan)
	if err != nil {
		t.Fatal(err)
	}
}
