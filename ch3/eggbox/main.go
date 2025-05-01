// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"io"
	"math"
	"os"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	draw(os.Stdout)
}

func draw(out io.Writer) {
	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, ok1 := corner(i+1, j)
			bx, by, ok2 := corner(i, j)
			cx, cy, ok3 := corner(i, j+1)
			dx, dy, ok4 := corner(i+1, j+1)
			
			if ok1 && ok2 && ok3 && ok4 {
				// Calculate average z for coloring
				_, _, z1 := corner3D(i+1, j)
				_, _, z2 := corner3D(i, j)
				_, _, z3 := corner3D(i, j+1)
				_, _, z4 := corner3D(i+1, j+1)
				avgZ := (z1 + z2 + z3 + z4) / 4
				
				// Map z-value to color (red for peaks, blue for valleys)
				color := colorMap(avgZ)
				
				fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g' "+
					"style='fill:%s'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy, color)
			}
		}
	}
	fmt.Fprintln(out, "</svg>")
}

// corner returns projected coordinates and validity
func corner(i, j int) (float64, float64, bool) {
	x, y, z := corner3D(i, j)
	
	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	
	return sx, sy, isValid(sx, sy)
}

// corner3D returns the 3D coordinates
func corner3D(i, j int) (float64, float64, float64) {
	// Find point (x,y) at corner of cell (i,j)
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	
	// Compute surface height z (egg box pattern)
	z := f(x, y)
	
	return x, y, z
}

// Egg box pattern function
func f(x, y float64) float64 {
	// Adjust frequency to get more or fewer "eggs"
	freq := 0.6
	return math.Sin(x*freq) * 0.5 * math.Sin(y*freq)
}

// isValid checks if coordinates are valid
func isValid(x, y float64) bool {
	if math.IsNaN(x) || math.IsNaN(y) || math.IsInf(x, 0) || math.IsInf(y, 0) {
		return false
	}
	margin := 0.1
	return x >= -width*margin && x <= width*(1+margin) && 
	       y >= -height*margin && y <= height*(1+margin)
}

// colorMap maps z-value to a color
func colorMap(z float64) string {
	// Normalize z to [0,1] range
	minZ, maxZ := -1.0, 1.0 // Known range for our function
	normalized := (z - minZ) / (maxZ - minZ)
	
	// Create gradient from blue (low) to red (high)
	r := uint8(255 * normalized)
	b := uint8(255 * (1 - normalized))
	return fmt.Sprintf("#%02x00%02x", r, b)
}