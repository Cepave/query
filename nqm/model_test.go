package nqm

import (
	"encoding/json"
	"github.com/bitly/go-simplejson"
	. "gopkg.in/check.v1"
)

type TestModelSuite struct{}

var _ = Suite(&TestModelSuite{})

// Tests the content of JSON
func (suite *TestModelSuite) TestJsonForAgentsInCity(c *C) {
	sampleData := []*AgentsInCity {
		&AgentsInCity {
			City: &City { Id: 20, Name: "city-1" },
			Agents: []Agent {
				Agent { Id: 91, Name: "Agent-1", Hostname: "Agenthost-1", IpAddress: "10.5.50.1" },
				Agent { Id: 92, Name: "Agent-2", Hostname: "Agenthost-2", IpAddress: "10.5.50.2" },
				Agent { Id: 93, Name: "Agent-3", Hostname: "Agenthost-3", IpAddress: "10.5.50.3" },
			},
		},
		&AgentsInCity {
			City: &City { Id: 21, Name: "city-2" },
			Agents: []Agent {
				Agent { Id: 121, Name: "Agent-121", Hostname: "Agenthost-121`", IpAddress: "10.5.50.121" },
			},
		},
	}

	jsonString, errMarshal := json.Marshal(sampleData)
	c.Assert(errMarshal, IsNil)
	c.Logf("JSON Result: %v", string(jsonString))
	jsonResult, errUnMarshal := simplejson.NewJson(jsonString)
	c.Assert(errUnMarshal, IsNil)

	/**
	 * Asserts the content of JSON
	 */
	c.Assert(jsonResult.MustArray(), HasLen, 2)
	c.Assert(jsonResult.GetIndex(0).GetPath("city", "id").MustInt(), Equals, 20)
	c.Assert(jsonResult.GetIndex(0).GetPath("agents").MustArray(), HasLen, 3)
	c.Assert(jsonResult.GetIndex(1).GetPath("city", "id").MustInt(), Equals, 21)
	c.Assert(jsonResult.GetIndex(1).GetPath("agents").MustArray(), HasLen, 1)

	jsonAgent := jsonResult.GetIndex(0).Get("agents").GetIndex(1)
	c.Assert(jsonAgent.Get("id").MustInt(), Equals, 92)
	c.Assert(jsonAgent.Get("name").MustString(), Equals, "Agent-2")
	c.Assert(jsonAgent.Get("hostname").MustString(), Equals, "Agenthost-2")
	c.Assert(jsonAgent.Get("ip_address").MustString(), Equals, "10.5.50.2")
	// :~)
}
