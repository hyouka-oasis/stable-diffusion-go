package python_core

import (
	_ "embed"
)

//go:embed voice_caption.py
var PythonVoiceCaptionPath string

//go:embed participle.py
var PythonParticiplePythonPath string
