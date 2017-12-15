package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/MirahImage/AdventOfCode2017/Day3"
)

var _ = Describe("Main", func() {

	var (
		one           int
		twelve        int
		twentythree   int
		tentwentyfour int
	)

	BeforeEach(func() {
		one = 1
		twelve = 12
		twentythree = 23
		tentwentyfour = 1024
	})

	Describe("FloorSqrt", func() {
		It("should return the floor of the sqrt", func() {
			Expect(FloorSqrt(one)).To(Equal(1))
			Expect(FloorSqrt(twelve)).To(Equal(3))
			Expect(FloorSqrt(twentythree)).To(Equal(4))
		})
	})

	Describe("Ring", func() {
		It("should return the ring of the value, starting with 0", func() {
			Expect(Ring(one)).To(Equal(0))
			Expect(Ring(twelve)).To(Equal(2))
			Expect(Ring(twentythree)).To(Equal(2))
		})
	})

	Describe("PosInRing", func() {
		It("should compute the position in the ring, starting with 1", func() {
			Expect(PosInRing(twelve)).To(Equal(3))
			Expect(PosInRing(twentythree)).To(Equal(14))
		})
	})

	Describe("SegmentInRing", func() {
		It("should compute which 8th of the ring we are in", func() {
			Expect(SegmentInRing(twelve)).To(Equal(2))
			Expect(SegmentInRing(twentythree)).To(Equal(7))
		})
	})

	Describe("PosInSegment", func() {
		It("should compute the position in the segment", func() {
			Expect(PosInSegment(twelve)).To(Equal(1))
			Expect(PosInSegment(twentythree)).To(Equal(2))
		})
	})

	Describe("ManhattanDistance", func() {
		It("should compute the manhattan distance", func() {
			Expect(ManhattanDistance(one)).To(Equal(0))
			Expect(ManhattanDistance(twelve)).To(Equal(3))
			Expect(ManhattanDistance(twentythree)).To(Equal(2))
			Expect(ManhattanDistance(tentwentyfour)).To(Equal(31))
		})
	})

})
