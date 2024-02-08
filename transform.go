package main


func TransformData(data  map[string]interface{}) OutgoingData {
	
	transformed := OutgoingData{
		Event: KeySearch(data,"ev"),
		EventType: KeySearch(data,"et"),
		AppID: KeySearch(data,"id"),
		UserID: KeySearch(data,"uid"),
		MessageID: KeySearch(data,"mid"),
		PageTitle: KeySearch(data,"t"),
		PageURL: KeySearch(data,"p"),
		BrowserLanguage: KeySearch(data,"l"),
		ScreenSize: KeySearch(data,"sc"),
		Attributes: make(map[string]attribute),
		UserAttributes: make(map[string]attribute),
	}
    
	//Map Attibute
    for key, value := range data {
		if len(key)>4{
            if strings.Contains(key[0:4],"atrk"){
            values:=KeySearch(data,"atrv"+key[4:])
		    types:=KeySearch(data,"atrt"+key[4:])
			attr:=attribute{
		          Value:values,
		          Type :types,	
                }
            transformed.Attributes[value.(string)]=attr
		}
		}
		
	}
   
	//Map UserAttibute
	for key, value := range data {
		if strings.Contains(key,"uatrk"){
            values:=KeySearch(data,"uatrv"+key[5:])
		    types:=KeySearch(data,"uatrt"+key[5:])
			attr:=attribute{
		          Value:values,
		          Type :types,	
                }
          transformed.UserAttributes[value.(string)]=attr
		}
	}
return transformed
}
