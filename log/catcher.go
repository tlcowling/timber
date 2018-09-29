package log

type panicCatcher struct{}

func (p panicCatcher) Catch(err error) {
	panic(err)
}
