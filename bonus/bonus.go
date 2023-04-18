package bonus

import (
	"fmt"
	"math"
)

func CalculateDamage(characterLevel, enemyLevel int) (int, error) {

	if characterLevel <= 0 || enemyLevel <= 0 {
		return 0, fmt.Errorf("nível inválido")
	}

	var damage int

	if characterLevel > 100 {
		damage = characterLevel * 20
	} else if enemyLevel > 100 {
		damage = characterLevel * 2
	} else {
		switch {
		case characterLevel > enemyLevel:
			damage = characterLevel * 10
		case characterLevel < enemyLevel:
			damage = characterLevel * 5
		default:
			damage = characterLevel * 7
		}

		levelDifference := int(math.Abs(float64(characterLevel - enemyLevel)))

		if levelDifference > 20 {
			damage *= 5
		} else if levelDifference < 20 {
			damage *= 2
		}
	}

	return damage, nil
}
