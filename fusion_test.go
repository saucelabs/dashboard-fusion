// Copyright 2023 Sauce Labs Inc., all rights reserved.

package dashboardfusion

import (
	"encoding/json"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMergePanels(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		base   []Panel
		extra  []Panel
		wanted []Panel
	}{
		{
			name: "replace content",
			base: []Panel{
				{
					"title":   json.RawMessage("Panel1"),
					"type":    json.RawMessage("graph"),
					"gridPos": json.RawMessage(`{"x":0,"y":0,"h":2,"w":6}`),
					"id":      json.RawMessage("1"),
				},
				{
					"title":   json.RawMessage("Panel2"),
					"type":    json.RawMessage("row"),
					"gridPos": json.RawMessage(`{"x":0,"y":2,"h":2,"w":6}`),
					"id":      json.RawMessage("2"),
				},
			},
			extra: []Panel{
				{
					"title":   json.RawMessage("Panel1"),
					"type":    json.RawMessage("graph"),
					"content": json.RawMessage("Wanted Content for Panel1"),
					"id":      json.RawMessage("3"),
				},
			},
			wanted: []Panel{
				{
					"title":   json.RawMessage("Panel1"),
					"type":    json.RawMessage("graph"),
					"content": json.RawMessage("Wanted Content for Panel1"),
					"gridPos": json.RawMessage(`{"x":0,"y":0,"h":2,"w":6}`),
					"id":      json.RawMessage("1"),
				},
				{
					"title":   json.RawMessage("Panel2"),
					"type":    json.RawMessage("row"),
					"gridPos": json.RawMessage(`{"x":0,"y":2,"h":2,"w":6}`),
					"id":      json.RawMessage("2"),
				},
			},
		},
		{
			name: "append new panel",
			base: []Panel{
				{
					"title":   json.RawMessage("Panel1"),
					"type":    json.RawMessage("graph"),
					"gridPos": json.RawMessage(`{"x":0,"y":0,"h":3,"w":5}`),
					"id":      json.RawMessage("1"),
				},
				{
					"title":   json.RawMessage("Panel2"),
					"type":    json.RawMessage("row"),
					"gridPos": json.RawMessage(`{"x":0,"y":2,"h":2,"w":6}`),
					"id":      json.RawMessage("2"),
				},
			},
			extra: []Panel{
				{
					"title":   json.RawMessage("Panel3"),
					"type":    json.RawMessage("graph"),
					"context": json.RawMessage("Wanted Content for Panel3"),
					"id":      json.RawMessage("3"),
				},
			},
			wanted: []Panel{
				{
					"title":   json.RawMessage("Panel1"),
					"type":    json.RawMessage("graph"),
					"gridPos": json.RawMessage(`{"x":0,"y":0,"h":3,"w":5}`),
					"id":      json.RawMessage("1"),
				},
				{
					"title":   json.RawMessage("Panel2"),
					"type":    json.RawMessage("row"),
					"gridPos": json.RawMessage(`{"x":0,"y":2,"h":2,"w":6}`),
					"id":      json.RawMessage("2"),
				},
				{
					"title":   json.RawMessage("Panel3"),
					"type":    json.RawMessage("graph"),
					"context": json.RawMessage("Wanted Content for Panel3"),
					"gridPos": json.RawMessage(`{"h":2,"w":6,"x":0,"y":5}`),
					"id":      json.RawMessage("3"),
				},
			},
		},
	}

	for i := range tests {
		tc := tests[i]
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			merged := MergePanels(tc.base, tc.extra)
			if diff := cmp.Diff(tc.wanted, merged); diff != "" {
				t.Fatalf("unexpected result (-want +got):\n%s", diff)
			}
		})
	}
}
