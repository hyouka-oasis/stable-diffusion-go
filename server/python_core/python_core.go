package python_core

import (
	_ "embed"
)

var (
	//go:embed requirements.txt
	PythonRequirements string
	//go:embed participle.py
	PythonParticiplePythonPath string
	//go:embed voice_caption.py
	PythonVoiceCaptionPath string
	//go:embed keyword_extraction.py
	PythonKeywordExtractionPath string
	PythonRequirementsName      = "requirements.txt"
)
