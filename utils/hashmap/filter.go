package hashmap


type Filter1Node interface {
	HashNode
	Filter1() string
}

func (t *HashTable)FindFilter1(f1 Filter1Node) (list []HashNode){
	t.Mutex.Lock()
	defer t.Mutex.Unlock()

	if t == nil {
		return nil
	}
	for _,v := range t.hashMap{
		if v.(Filter1Node).Filter1() == f1.Filter1() {
			list = append(list, v)
		}
	}
	return list
}
