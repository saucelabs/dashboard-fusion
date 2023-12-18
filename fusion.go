// Copyright 2023 Sauce Labs Inc., all rights reserved.

package dashboardfusion

import (
	"bytes"
	"encoding/json"
)

type Dashboard map[string]json.RawMessage

func (d Dashboard) Panels() []Panel {
	if ps, ok := d["panels"]; ok {
		var panels []Panel
		if err := json.Unmarshal(ps, &panels); err != nil {
			panic(err)
		}
		return panels
	}

	return nil
}

type Panel map[string]json.RawMessage

func (p Panel) Equals(p2 Panel) bool {
	return bytes.Equal(p["title"], p2["title"]) &&
		bytes.Equal(p["type"], p2["type"])
}

func (p Panel) GridPosRaw() json.RawMessage {
	return p["gridPos"]
}

func (p Panel) GridPos() GridPos {
	if gp, ok := p["gridPos"]; ok {
		var gridPos GridPos
		if err := json.Unmarshal(gp, &gridPos); err != nil {
			panic(err)
		}
		return gridPos
	}

	return GridPos{}
}

type GridPos struct {
	H int `json:"h"`
	W int `json:"w"`
	X int `json:"x"`
	Y int `json:"y"`
}

// MergePanels merges two sets of panels.
//
// If a panel in ps2 matches a panel in ps1, the panel in ps2 overwrites the
// content of the panel in ps1, but preserves its position.
//
// If a panel in ps2 does not match any panel in ps1 it is appended and placed at the end of the dashboard.
func MergePanels(ps1, ps2 []Panel) []Panel {
	var maxY int
	res := make([]Panel, 0, len(ps1)+len(ps2))
	for _, p1 := range ps1 {
		if gp := p1.GridPos(); gp.Y+gp.H > maxY {
			maxY = gp.Y + gp.H
		}
		res = append(res, p1)
	}

	for len(ps2) > 0 {
		p2 := ps2[0]
		ps2 = ps2[1:]

		var matched bool
		for i := range res {
			if res[i].Equals(p2) {
				// When we find a match, the panel's content is overwritten,
				// except for the gridPos(to preserve the layout).
				p2["gridPos"] = res[i].GridPosRaw()
				res[i] = p2
				matched = true
			}
		}

		if !matched {
			g := GridPos{
				H: 2,
				W: 6,
				X: 0,
				Y: maxY + 1,
			}
			graw, err := json.Marshal(g)
			if err != nil {
				panic(err)
			}
			p2["gridPos"] = graw

			res = append(res, p2)
			maxY += g.H
		}
	}

	return res
}
