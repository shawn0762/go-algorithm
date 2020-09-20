package ticket

func ImplodeTicket() {
	a := map[string]string{
		"上海": "大连",
		"西安": "成都",
		"北京": "上海",
		"大连": "西安",
	}

	// 交换字典的kv
	b := map[string]string{}
	for k, v := range a {
		b[v] = k
	}

	// 找到起点
	start := ""
	for k, _ := range a {
		_, ok := b[k]
		if !ok {
			start = k
			break
		}
	}

	s := start
	e, ok := a[s]
	for ok {
		println(s, e)
		s = e
		e, ok = a[s]
	}
}
