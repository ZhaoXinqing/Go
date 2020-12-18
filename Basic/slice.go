package Basic

//RemoveRepByMap slice去重
func RemoveRepByMap(arr []int64) []int64 {
	result := []int64{}         //存放返回的不重复切片
	tempMap := map[int64]byte{} // 存放不重复主键
	for _, e := range arr {
		l := len(tempMap)
		tempMap[e] = 0 //当e存在于tempMap中时，再次添加是添加不进去的，，因为key不允许重复
		//如果上一行添加成功，那么长度发生变化且此时元素一定不重复
		if len(tempMap) != l { // 加入map后，map长度变化，则元素不重复
			result = append(result, e) //当元素不重复时，将元素添加到切片result中
		}
	}
	return result
}

///***
// *
// *                 .-~~~~~~~~~-._       _.-~~~~~~~~~-.
// *             __.'  欢迎访问    ~.   .~              `.__
// *           .'//  我的个人博客    \./   ☞ GitHub.com ☜   \\`.
// *         .'//                     |                     \\`.
// *       .'// .-~"""""""~~~~-._     |     _,-~~~~"""""""~-. \\`.
// *     .'//.-"                 `-.  |  .-'                 "-.\\`.
// *   .'//______.============-..   \ | /   ..-============.______\\`.
// * .'______________________________\|/______________________________`.
// */

////AddressResolution 地址解析
//func AddressResolution(address, city string) (lng, lat float64) {
//	params := map[string]string{}
//	params["key"] = GaoDeKey
//	params["address"] = address
//	if city != "" {
//		params["city"] = city
//	}
//
//	//TODO 层级嵌套深，需要改进
//	if res, err := Get(URL, params, nil); err == nil {
//		defer res.Body.Close()
//		if body, err := ioutil.ReadAll(res.Body); err == nil {
//			gaodeGeoRes := domain.GaodeGeoRes{}
//			if err = odysseyUtils.Unmarshal(body, &gaodeGeoRes); err == nil {
//				if gaodeGeoRes.Status == "1" && len(gaodeGeoRes.Geocodes) != 0 { //请求成功
//					latlngStrs := strings.Split(gaodeGeoRes.Geocodes[0].Location, ",")
//					lng, _ = strconv.ParseFloat(latlngStrs[0], 64)
//					lat, _ = strconv.ParseFloat(latlngStrs[1], 64)
//				}
//			}
//		}
//	}
//	return lng, lat
//}
