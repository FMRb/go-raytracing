package main

// HitableList contains list of hitable
type HitableList struct {
	list []Hitable
}

func (hl HitableList) hit(r *Ray, tMin float64, tMax float64, rec *HitRecord) bool {
	var tempRec HitRecord
	hitAnything := false
	closesSoFar := tMax
	for i := 0; i < len(hl.list); i++ {
		if hl.list[i].hit(r, tMin, closesSoFar, &tempRec) {
			hitAnything = true
			closesSoFar = tempRec.t
			rec.normal = tempRec.normal
			rec.p = tempRec.p
			rec.t = tempRec.t
		}
	}
	return hitAnything
}
