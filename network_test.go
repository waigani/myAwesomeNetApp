// Copyright 2015 Jesse Meek.
// Licensed under the AGPLv3, see LICENCE file for details.

package network_test

import (
	"strconv"
	"testing"

	jt "github.com/juju/testing"
	jc "github.com/juju/testing/checkers"
	"github.com/waigani/myAwesomeNetApp"
	gc "gopkg.in/check.v1"
)

func Test(t *testing.T) {
	gc.TestingT(t)
}

type suite struct {
	jt.CleanupSuite
}

var _ = gc.Suite(&suite{})

func (s *suite) TestServerAddress(c *gc.C) {
	server := network.NewServer("localhost", map[int]bool{8080: true})
	c.Assert(server.Address("8080"), gc.Equals, "localhost:8080")
}

func (s *suite) TestPortsValid(c *gc.C) {
	server := network.NewServer("s1", map[int]bool{12445: false, 5678: false, 12344453: false})

	// Assert Ports*() only return ports no greater than 6.
	var portSlice []int
	c.Assert(server.Ports(), gc.HasLen, 2)
	for p, aval := range server.Ports() {
		c.Assert(aval, jc.IsFalse)
		portSlice = append(portSlice, p)
	}

	for _, p := range portSlice {
		c.Assert(len(strconv.Itoa(p)), jc.LessThan, 7)
	}
}

func (s *suite) TestAllPortsTrue(c *gc.C) {
	server := network.NewServer("s2", map[int]bool{
		12445:    true,
		5678:     true,
		12344453: true,
	})

	ports := server.Ports()

	for _, aval := range ports {
		c.Assert(aval, jc.IsTrue)
	}
}
