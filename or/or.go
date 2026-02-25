package or

// Or канал, который закрывается, когда закрывается любой из переданных каналов
// Реализация использует рекурсию для объединения нескольких каналов
func Or(channels ...<-chan interface{}) <-chan interface{} {
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	default:
		r := make(chan interface{})
		go func() {
			defer close(r)
			select {
			case <-channels[0]:
			case <-Or(channels[1:]...):
			}
		}()
		return r
	}
}
