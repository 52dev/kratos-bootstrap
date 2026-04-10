package script_engine

import (
	scriptEngine "github.com/52dev/go-scripts"
	conf "github.com/52dev/kratos-bootstrap/api/gen/go/conf/v1"
)

// convertTypeScriptEngineType 转换 conf.Script_EngineType 到 scriptEngine.Type
func convertTypeScriptEngineType(t conf.Script_EngineType) scriptEngine.Type {
	switch t {
	case conf.Script_LUA:
		return scriptEngine.LuaType
	case conf.Script_JAVASCRIPT:
		return scriptEngine.JavaScriptType
	default:
		return ""
	}
}
