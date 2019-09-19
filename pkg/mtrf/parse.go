package mtrf

// -1 add to payload
// -2 ignore

type mask [17]int

// var messages = []struct {
// 	mode Mode
// 	cmd  Cmd
// 	mask mask
// }{
// 	{ModeTX, "off", mask{171, 0, 0, 0, -1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 172}},
// 	{ModeTX, "on", mask{171, 0, 0, 0, -1, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 172}},
// 	{ModeTX, "switch", mask{171, 0, 0, 0, -1, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 172}},
// 	{ModeTX, "unbind", mask{171, 0, 0, 0, -1, 9, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 172}},
// 	{ModeTX, "bind", mask{171, 0, 0, 0, -1, 15, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 172}},
// 	{ModeRX, "sens_temp_humi", mask{173, 1, 0, -2, -1, 21, 7, -1, -1, -1, 255, 0, 0, 0, 0, 0, 174}},
// 	{ModeRX, "clear", mask{171, 1, 5, 0, -1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 220, 172}},
// 	{ModeRX, "wait_bind", mask{171, 1, 3, 0, -1, 15, 0, 0, 0, 0, 0, 0, 0, 0, 0, 233, 172}},
// 	{ModeRX, "wait_bind_finished", mask{173, 1, 0, -2, -1, 15, 1, 2, 0, 0, 0, 0, 0, 0, 0, 243, 174}},
// 	{ModeRX, "bind", mask{173, 1, 0, -2, -1, 15, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 174}},
// 	{ModeRX, "off", mask{173, 1, 0, -2, -1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 244, 174}},
// 	{ModeRX, "on", mask{173, 1, 0, -2, -1, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 248, 174}},
// }

// // @TODO: make index for find relevant message

// // Parse raw message
// func Parse(body [17]uint8) (name string, payload [17]uint8, kv map[string]interface{}) {
// 	kv = make(map[string]interface{})
// 	name = "unknown"

// 	return
// }
