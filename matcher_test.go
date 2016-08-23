package gatekeeper

import . "gopkg.in/check.v1"

func (s *TestSuite) TestNewMatcher(c *C) {
	m := NewMatcher()
	c.Assert(m, NotNil)
	c.Check(m, FitsTypeOf, &Matcher{})
	c.Check(m.Permissions, DeepEquals, []string{})
	c.Check(m.Context, DeepEquals, map[string]string{})
}

func (s *TestSuite) TestWithPermissions(c *C) {
	sm := NewMatcher()
	m := sm.WithPermissions([]string{"x"})
	m.Context = map[string]string{"x": "y"}
	c.Assert(m, NotNil)
	c.Check(m, FitsTypeOf, sm)
	c.Check(m.Permissions, DeepEquals, []string{"x"})
	c.Check(m.Context, DeepEquals, map[string]string{"x": "y"})

	c.Check(sm, Not(Equals), m)
	c.Check(sm.Permissions, DeepEquals, []string{})
	c.Check(sm.Context, DeepEquals, map[string]string{})
}

func (s *TestSuite) TestWithContext(c *C) {
	sm := NewMatcher()
	m := sm.WithContext(map[string]string{"x": "y"})
	m.Permissions = []string{"x"}
	c.Assert(m, NotNil)
	c.Check(m, FitsTypeOf, sm)
	c.Check(m.Permissions, DeepEquals, []string{"x"})
	c.Check(m.Context, DeepEquals, map[string]string{"x": "y"})

	c.Check(sm, Not(Equals), m)
	c.Check(sm.Permissions, DeepEquals, []string{})
	c.Check(sm.Context, DeepEquals, map[string]string{})
}

func (s *TestSuite) TestCanBasic(c *C) {
	m := NewMatcher().WithPermissions([]string{"x", "y"})

	c.Check(m.Can("x"), Equals, true)
	c.Check(m.Can("y"), Equals, true)
	c.Check(m.Can("z"), Equals, false)
}

func (s *TestSuite) TestCanInterpolation(c *C) {
	m := NewMatcher().WithPermissions([]string{"project/abc", "project/abc/admin"})

	c.Check(m.Can("project/:project"), Equals, false)

	m = m.WithContext(map[string]string{
		"project": "abc",
	})

	c.Check(m.Can("project/:project"), Equals, true)
	c.Check(m.Can("project/:project/admin"), Equals, true)

	m = m.WithPermissions([]string{"project/:project"})
	c.Check(m.Can("project/:project"), Equals, true)
	c.Check(m.Can("project/:project/admin"), Equals, false)
}

func (s *TestSuite) TestCanAll(c *C) {
	m := NewMatcher().WithPermissions([]string{"x", "y"})
	c.Check(m.CanAll("x", "y"), Equals, true)
	c.Check(m.CanAll("x", "y", "z"), Equals, false)
}
