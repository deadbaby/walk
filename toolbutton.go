// Copyright 2012 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package walk

import . "github.com/lxn/go-winapi"

var toolButtonOrigWndProcPtr uintptr
var _ subclassedWidget = &ToolButton{}

type ToolButton struct {
	Button
}

func NewToolButton(parent Container) (*ToolButton, error) {
	tb := &ToolButton{}

	if err := initChildWidget(
		tb,
		parent,
		"BUTTON",
		WS_TABSTOP|WS_VISIBLE|BS_PUSHBUTTON,
		0); err != nil {
		return nil, err
	}

	return tb, nil
}

func (*ToolButton) origWndProcPtr() uintptr {
	return toolButtonOrigWndProcPtr
}

func (*ToolButton) setOrigWndProcPtr(ptr uintptr) {
	toolButtonOrigWndProcPtr = ptr
}

func (*ToolButton) LayoutFlags() LayoutFlags {
	return 0
}

func (tb *ToolButton) MinSizeHint() Size {
	return tb.SizeHint()
}

func (tb *ToolButton) SizeHint() Size {
	return tb.dialogBaseUnitsToPixels(Size{20, 14})
}

func (tb *ToolButton) wndProc(hwnd HWND, msg uint32, wParam, lParam uintptr) uintptr {
	switch msg {
	case WM_GETDLGCODE:
		return DLGC_BUTTON
	}

	return tb.Button.wndProc(hwnd, msg, wParam, lParam)
}
