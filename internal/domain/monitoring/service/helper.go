package service

import "github.com/Ndraaa15/foreglyc-server/pkg/enum"

func BloodGlucoseStatus(bloodGlucose float64) enum.GlucoseStatus {
	if bloodGlucose > 250 {
		return enum.GlucoseStatusHyperglycemiaChronic
	} else if bloodGlucose > 180 {
		return enum.GlucoseStatusHyperglycemiaAccute
	} else if bloodGlucose > 80 {
		return enum.GlucoseStatusNormal
	} else if bloodGlucose > 50 {
		return enum.GlucoseStatusHypoglycemiaAccute
	} else {
		return enum.GlucoseStatusHypoglycemiaChronic
	}
}
