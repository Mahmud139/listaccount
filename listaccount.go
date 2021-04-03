package listaccount

func ListAccount(accountType []string, data []map[string]interface{}) map[string][]interface{}{
	var result = make(map[string][]interface{})
	for _,v:=range accountType{
		for _, va := range data {
			if va["account_type"] == v { 
				result[v] = append(result[v], va["name"])  
			}
		}
	}

	return result
}





