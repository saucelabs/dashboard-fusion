# Dashboard Fusion

Dashboard Fusion is a tool for merging and updating Grafana dashboards by combining panels from different sources. 
It's designed to be used when working with dashboards that share common panels.
It allows to update existing panels, while preserving the dashboard layout.

## Fusion Behavior

Fusion is performed by merging panels from multiple sources into a single dashboard.
- Panels are matched by title and type.
- If a panel with the same title and type exists in the dashboard, its content will be replaced, except for its position (preserving the dashboard layout) and id.
- If a panel with the same title and type does not exist in the dashboard, it will be appended to the end of the dashboard.
Later, the dashboard can be manually reorganized to achieve the desired layout.

## Usage

```bash
# Merge panels from panels1.json and panels2.json into dashboard.json and save the result to output.json
go run ./cmd/dashboard-fusion --dash <dashboard.json> --panels <panels1.json,panels2.json> --out <output.json>
```

## Example

Visit [here](./example/README.md) for a complete example of Dashboard Fusion usage.
