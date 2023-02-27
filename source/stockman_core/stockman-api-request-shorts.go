package core

func ApiRequestShortcutEasyResponse[TypeInput, TypeOutput any](ef EventFn, input TypeInput) *StockmanResponse[TypeOutput] {
	ev := NewEvent(ef)
	ev.SetInput(input)
	SystemEvents_Manager.AppendEvent(ev)
	<-ev.OnOutputChanged()
	var output TypeOutput
	ev.LoadOutput(&output)
	resp := NewStockmanResponse(output, ev.Error())
	return resp
}
