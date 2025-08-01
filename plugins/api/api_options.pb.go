//go:build !wasip1

// Code generated by protoc-gen-go-plugin. DO NOT EDIT.
// versions:
// 	protoc-gen-go-plugin v0.1.0
// 	protoc               v5.29.3
// source: api/api.proto

package api

import (
	context "context"
	wazero "github.com/tetratelabs/wazero"
	wasi_snapshot_preview1 "github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
)

type wazeroConfigOption func(plugin *WazeroConfig)

type WazeroNewRuntime func(context.Context) (wazero.Runtime, error)

type WazeroConfig struct {
	newRuntime   func(context.Context) (wazero.Runtime, error)
	moduleConfig wazero.ModuleConfig
}

func WazeroRuntime(newRuntime WazeroNewRuntime) wazeroConfigOption {
	return func(h *WazeroConfig) {
		h.newRuntime = newRuntime
	}
}

func DefaultWazeroRuntime() WazeroNewRuntime {
	return func(ctx context.Context) (wazero.Runtime, error) {
		r := wazero.NewRuntime(ctx)
		if _, err := wasi_snapshot_preview1.Instantiate(ctx, r); err != nil {
			return nil, err
		}

		return r, nil
	}
}

func WazeroModuleConfig(moduleConfig wazero.ModuleConfig) wazeroConfigOption {
	return func(h *WazeroConfig) {
		h.moduleConfig = moduleConfig
	}
}
