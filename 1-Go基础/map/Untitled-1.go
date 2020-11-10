//AddressResolution 地址解析
func AddressResolution(address, city string) (lng, lat float64) {
	params := map[string]string{}
	params["key"] = GaoDeKey
	params["address"] = address
	if city != "" {
		params["city"] = city
	}

	//TODO 层级嵌套深，需要改进
	if res, err := Get(URL, params, nil); err == nil {
		defer res.Body.Close()
		if body, err := ioutil.ReadAll(res.Body); err == nil {
			gaodeGeoRes := domain.GaodeGeoRes{}
			if err = odysseyUtils.Unmarshal(body, &gaodeGeoRes); err == nil {
				if gaodeGeoRes.Status == "1" && len(gaodeGeoRes.Geocodes) != 0 { //请求成功
					latlngStrs := strings.Split(gaodeGeoRes.Geocodes[0].Location, ",")
					lng, _ = strconv.ParseFloat(latlngStrs[0], 64)
					lat, _ = strconv.ParseFloat(latlngStrs[1], 64)
				}
			}
		}
	}
	return lng, lat
}





// [1]
if res, err := Get(URL, params, nil); err == nil {
	defer res.Body.Close()
	if body, err := ioutil.ReadAll(res.Body); err == nil {
		gaodeGeoRes := domain.GaodeGeoRes{}
	}
}
if err = odysseyUtils.Unmarshal(body, &gaodeGeoRes); err == nil {
	if gaodeGeoRes.Status == "1" && len(gaodeGeoRes.Geocodes) != 0 { //请求成功
		latlngStrs := strings.Split(gaodeGeoRes.Geocodes[0].Location, ",")
		lng, _ = strconv.ParseFloat(latlngStrs[0], 64)
		lat, _ = strconv.ParseFloat(latlngStrs[1], 64)
	}
}
// [2]
if res, err := Get(URL, params, nil); err == nil {
	if body, err := ioutil.ReadAll(res.Body); err == nil {
		gaodeGeoRes := domain.GaodeGeoRes{}
	}
}
err = odysseyUtils.Unmarshal(body, &gaodeGeoRes)
if err == nil && gaodeGeoRes.Status == "1"  &&  len(gaodeGeoRes.Geocodes) != 0 {
	latlngStrs := strings.Split(gaodeGeoRes.Geocodes[0].Location, ",")
	lng, _ = strconv.ParseFloat(latlngStrs[0], 64)
	lat, _ = strconv.ParseFloat(latlngStrs[1], 64)
}
defer res.Body.Close()
return lng, lat

// 【3】
res, err := Get(URL, params, nil)
if err == nil && body, err2 := ioutil.ReadAll(res.Body); err2 == nil {
	gaodeGeoRes := domain.GaodeGeoRes{}
}
err = odysseyUtils.Unmarshal(body, &gaodeGeoRes)
if err == nil && gaodeGeoRes.Status == "1"  &&  len(gaodeGeoRes.Geocodes) != 0 {
	latlngStrs := strings.Split(gaodeGeoRes.Geocodes[0].Location, ",")
	lng, _ = strconv.ParseFloat(latlngStrs[0], 64)
	lat, _ = strconv.ParseFloat(latlngStrs[1], 64)
}
defer res.Body.Close()
return lng, lat