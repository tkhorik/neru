package config_test

import (
	"testing"

	"github.com/y3owk1n/neru/internal/config"
)

func TestConfig_ValidateGeneral(t *testing.T) {
	tests := []struct {
		name    string
		config  config.Config
		wantErr bool
	}{
		{
			name: "both disabled - valid",
			config: config.Config{
				General: config.GeneralConfig{
					RestoreCursorPosition: false,
					CenterCursorPosition:  false,
				},
			},
			wantErr: false,
		},
		{
			name: "restore enabled - valid",
			config: config.Config{
				General: config.GeneralConfig{
					RestoreCursorPosition: true,
					CenterCursorPosition:  false,
				},
			},
			wantErr: false,
		},
		{
			name: "center enabled - valid",
			config: config.Config{
				General: config.GeneralConfig{
					RestoreCursorPosition: false,
					CenterCursorPosition:  true,
				},
			},
			wantErr: false,
		},
		{
			name: "both enabled - invalid",
			config: config.Config{
				General: config.GeneralConfig{
					RestoreCursorPosition: true,
					CenterCursorPosition:  true,
				},
			},
			wantErr: true,
		},
		{
			name: "kb layout id set - valid",
			config: config.Config{
				General: config.GeneralConfig{
					KBLayoutToUse: "com.apple.keylayout.ABC",
				},
			},
			wantErr: false,
		},
		{
			name: "kb layout id whitespace-only - invalid",
			config: config.Config{
				General: config.GeneralConfig{
					KBLayoutToUse: "   ",
				},
			},
			wantErr: true,
		},
		{
			name: "backtrack key named value - valid",
			config: config.Config{
				General: config.GeneralConfig{
					BacktrackKey: "backspace",
				},
			},
			wantErr: false,
		},
		{
			name: "backtrack key modifier combo - valid",
			config: config.Config{
				General: config.GeneralConfig{
					BacktrackKey: "Ctrl+H",
				},
			},
			wantErr: false,
		},
		{
			name: "backtrack key single character - valid",
			config: config.Config{
				General: config.GeneralConfig{
					BacktrackKey: ",",
				},
			},
			wantErr: false,
		},
		{
			name: "backtrack key single space - valid",
			config: config.Config{
				General: config.GeneralConfig{
					BacktrackKey: " ",
				},
			},
			wantErr: false,
		},
		{
			name: "backtrack key invalid format - invalid",
			config: config.Config{
				General: config.GeneralConfig{
					BacktrackKey: "invalid_key_name",
				},
			},
			wantErr: true,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			err := testCase.config.ValidateGeneral()
			if (err != nil) != testCase.wantErr {
				t.Errorf("ValidateGeneral() error = %v, wantErr %v", err, testCase.wantErr)
			}
		})
	}
}
