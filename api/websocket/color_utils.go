package websocket

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// HSLToRGB convertit une couleur HSL en RGB.
func HSLToRGB(h, s, l float64) (r, g, b int) {
	c := (1 - math.Abs(2*l-1)) * s
	x := c * (1 - math.Abs(math.Mod(h/60, 2)-1))
	m := l - c/2

	var r1, g1, b1 float64
	switch {
	case h < 60:
		r1, g1, b1 = c, x, 0
	case h < 120:
		r1, g1, b1 = x, c, 0
	case h < 180:
		r1, g1, b1 = 0, c, x
	case h < 240:
		r1, g1, b1 = 0, x, c
	case h < 300:
		r1, g1, b1 = x, 0, c
	default:
		r1, g1, b1 = c, 0, x
	}

	r = int((r1 + m) * 255)
	g = int((g1 + m) * 255)
	b = int((b1 + m) * 255)
	return
}

// génère une liste de couleurs distinctes pour les serpents.
func GenerateSnakeColors(nbMaxSnake int) []string {
	colors := make([]string, nbMaxSnake)
	step := 360 / nbMaxSnake

	for i := 0; i < nbMaxSnake; i++ {
		h := float64(i * step)
		s := 0.7
		l := 0.5
		r, g, b := HSLToRGB(h, l, s)
		colors[i] = fmt.Sprintf("#%02X%02X%02X", r, g, b)
	}

	return colors
}

// DarkenHexColor assombrit une couleur hexadécimale en diminuant les composantes RVB d'un certain pourcentage.
func DarkenHexColor(hex string, percentage float64) (string, error) {
	if len(hex) != 7 || !strings.HasPrefix(hex, "#") {
		return "", fmt.Errorf("couleur hexadécimale invalide : %s", hex)
	}

	// Extraire les composantes RVB
	r, err := strconv.ParseInt(hex[1:3], 16, 64)
	if err != nil {
		return "", err
	}
	g, err := strconv.ParseInt(hex[3:5], 16, 64)
	if err != nil {
		return "", err
	}
	b, err := strconv.ParseInt(hex[5:7], 16, 64)
	if err != nil {
		return "", err
	}

	// Appliquer la réduction pour assombrir
	r = int64(float64(r) * (1 - percentage/100))
	g = int64(float64(g) * (1 - percentage/100))
	b = int64(float64(b) * (1 - percentage/100))

	// Limiter les valeurs entre 0 et 255
	r = clamp(r, 0, 255)
	g = clamp(g, 0, 255)
	b = clamp(b, 0, 255)

	// Recréer la chaîne hexadécimale
	darkenedHex := fmt.Sprintf("#%02X%02X%02X", r, g, b)
	return darkenedHex, nil
}

// clamp limite une valeur entre un minimum et un maximum
func clamp(value, min, max int64) int64 {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}
