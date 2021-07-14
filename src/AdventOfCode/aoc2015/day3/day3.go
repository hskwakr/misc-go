package day3

// Position of a point for a house
type Point struct {
	x int
	y int
}

// Return a number of houses recieved presents
func Count(filePath string) int {
	in := Data(filePath)
	fp := MoveSanta(in)
	return CountHouses(fp)
}

// Move Santa with elf's order to make footprints
func MoveSanta(in string) []Point {
	// Santa's footprints
	var fp []Point

	fp = append(fp, Point{0, 0})
	for k, v := range in {
		switch v {
		case '^':
			fp = append(fp, Point{fp[k].x, fp[k].y + 1})
			break
		case 'v':
			fp = append(fp, Point{fp[k].x, fp[k].y - 1})
			break
		case '>':
			fp = append(fp, Point{fp[k].x + 1, fp[k].y})
			break
		case '<':
			fp = append(fp, Point{fp[k].x - 1, fp[k].y})
			break
		}
	}

	return fp
}

// Count houses number of houses receive at least one present
func CountHouses(fp []Point) int {
	// number of houses
	r := 1

	var unique []Point
	for k, v1 := range fp {
		if k == 0 {
			unique = append(unique, v1)
			continue
		}

		isUnique := true
		for _, v2 := range unique {
			if v1 == v2 {
				isUnique = false
				break
			}
		}
		if isUnique {
			unique = append(unique, v1)
			r++
		}
	}

	return r
}
