package policy

// IsWithinZone checks if the center point of a bounding box falls within the restricted area coordinate limits.
func IsWithinZone(bx, by, bw, bh float64, zMinX, zMinY, zMaxX, zMaxY float64) bool {
	centerX := bx + bw/2
	centerY := by + bh/2
	return centerX >= zMinX && centerX <= zMaxX && centerY >= zMinY && centerY <= zMaxY
}

// IsDetectionValid checks if the confidence score meets safety thresholds.
func IsDetectionValid(confidence float64, threshold float64) bool {
	return confidence >= threshold
}
