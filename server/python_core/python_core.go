package python_core

import (
	"embed"
	_ "embed"
)

//go:embed voice_caption.py
var PythonVoiceCaptionPath embed.FS

//go:embed participle.py
var PythonParticiplePythonPath string
