type KeyValue struct {
	val int
	key int
}

type LRUCache struct {
    capacity int
	cache map[int]*list.Element
	queue *list.List
}

func Constructor(capacity int) LRUCache {
    return LRUCache{
		cache: map[int]*list.Element{},
		queue: list.New(),
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
    element, ok := this.cache[key]
	if !ok {
		return -1
	}

	this.queue.MoveToBack(element)

	keyVal := element.Value.(KeyValue)
	return keyVal.val
}

func (this *LRUCache) Put(key int, value int) {
	if element, ok := this.cache[key]; ok {
		element.Value = KeyValue{key: key, val: value}
		this.queue.MoveToBack(element)
		return
	}

	element := this.queue.PushBack(KeyValue{key: key, val: value})
	this.cache[key] = element

	if this.queue.Len() > this.capacity {
		element := this.queue.Front()
		keyVal := element.Value.(KeyValue)
		delete(this.cache, keyVal.key)
		this.queue.Remove(element)
	}
}
