package trace

// Tracer 는 코드 전체에서 이벤트를 추적할 수 있는 객체를 설명하는 인터페이스

type Tracer interface {
	Trace(...interface{})
}