package day12

import (
	"slices"
	"strings"
	"zakini/advent-of-code-2024/internal/utils"
)

func SolvePart1(input string, debug bool) int {
	farm := parseInput(input)
	regions := findRegions(farm)

	cost := 0
	for _, region := range regions {
		cost += calculateRegionFenceCost(region)
	}

	return cost
}

func parseInput(input string) [][]string {
	lines := strings.Split(input, "\n")

	farm := make([][]string, len(lines))
	for y, line := range lines {
		farm[y] = utils.SplitIntoChars(line)
	}

	return farm
}

func findRegions(farm [][]string) [][]utils.Vector2 {
	plotHandledMap := make(map[utils.Vector2]bool)

	regions := make([][]utils.Vector2, 0)
	for y, row := range farm {
		for x, plantType := range row {
			point := utils.Vector2{X: x, Y: y}

			if _, handled := plotHandledMap[point]; handled {
				continue
			}

			region := crawlRegion(farm, plantType, []utils.Vector2{point}, []utils.Vector2{})
			for _, plotPoint := range region {
				plotHandledMap[plotPoint] = true
			}

			regions = append(regions, region)
		}
	}

	return regions
}

func crawlRegion(farm [][]string, plantType string, frontier []utils.Vector2, regionSoFar []utils.Vector2) []utils.Vector2 {
	utils.Assert(len(frontier) <= 4 && len(regionSoFar) <= len(farm)*len(farm[0]), "uh-oh")

	for _, plot := range frontier {
		if farm[plot.Y][plot.X] != plantType || slices.Contains(regionSoFar, plot) {
			continue
		}

		regionSoFar = append(regionSoFar, plot)

		candidatePlots := utils.FindSurroundingPointsFunc(farm, plot, func(plot string) bool {
			return plot == plantType
		})

		regionSoFar = crawlRegion(farm, plantType, candidatePlots, regionSoFar)
	}

	return regionSoFar
}

func calculateRegionFenceCost(region []utils.Vector2) int {
	area := len(region)

	perimeter := 0
	for _, plot := range region {
		neighbouringPlots := utils.Filter(region, func(_ int, other utils.Vector2) bool {
			return (plot.X == other.X && (plot.Y == other.Y-1 || plot.Y == other.Y+1)) ||
				(plot.Y == other.Y && (plot.X == other.X-1 || plot.X == other.X+1))
		})

		perimeter += 4 - len(neighbouringPlots)
	}

	return area * perimeter
}
