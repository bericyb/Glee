package scriptengine

type ScriptEngine interface {
	Execute(script string) (string, error)
}
