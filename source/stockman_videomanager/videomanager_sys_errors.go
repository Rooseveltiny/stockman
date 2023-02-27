package stockmanvideomanager

import core "stockman/source/stockman_core"

var CantCreateCamera core.SystemError = *core.NewSystemError("can't create video camera")
var CantFindCamera core.SystemError = *core.NewSystemError("camera isn't found")
