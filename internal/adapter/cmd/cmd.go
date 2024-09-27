package cmd

import "github.com/shiron-dev/rapi/internal/di"

var controllersSet, _ = di.InitializeControllerSet()
var controller = controllersSet.Controller
