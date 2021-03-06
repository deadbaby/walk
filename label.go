// Copyright 2010 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package walk

import . "github.com/lxn/go-winapi"

var labelOrigWndProcPtr uintptr
var _ subclassedWidget = &Label{}

type Label struct {
	WidgetBase
}

func NewLabel(parent Container) (*Label, error) {
	l := &Label{}

	if err := initChildWidget(
		l,
		parent,
		"STATIC",
		WS_VISIBLE|SS_CENTERIMAGE,
		0); err != nil {
		return nil, err
	}

	return l, nil
}

func (*Label) origWndProcPtr() uintptr {
	return labelOrigWndProcPtr
}

func (*Label) setOrigWndProcPtr(ptr uintptr) {
	labelOrigWndProcPtr = ptr
}

func (*Label) LayoutFlags() LayoutFlags {
	return GrowableVert
}

func (l *Label) MinSizeHint() Size {
	return l.calculateTextSize()
}

func (l *Label) SizeHint() Size {
	return l.MinSizeHint()
}

func (l *Label) Text() string {
	return widgetText(l.hWnd)
}

func (l *Label) SetText(value string) error {
	if value == l.Text() {
		return nil
	}

	if err := setWidgetText(l.hWnd, value); err != nil {
		return err
	}

	return l.updateParentLayout()
}

func (l *Label) wndProc(hwnd HWND, msg uint32, wParam, lParam uintptr) uintptr {
	switch msg {
	case WM_SIZE, WM_SIZING:
		l.Invalidate()
	}

	return l.WidgetBase.wndProc(hwnd, msg, wParam, lParam)
}
