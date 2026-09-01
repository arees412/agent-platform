package agent

import "testing"

func TestDefaultAgentsIncludeValidBusinessAgentPack(t *testing.T) {
	agents := DefaultAgents()
	byID := make(map[string]*Agent, len(agents))
	for _, configuredAgent := range agents {
		if _, exists := byID[configuredAgent.ID]; exists {
			t.Fatalf("duplicate agent ID %q", configuredAgent.ID)
		}
		if err := configuredAgent.Validate(); err != nil {
			t.Fatalf("agent %q failed validation: %v", configuredAgent.ID, err)
		}
		byID[configuredAgent.ID] = configuredAgent
	}

	expectedBusinessIDs := []string{
		"business-operations-director",
		"revenue-intelligence-agent",
		"customer-operations-agent",
		"risk-compliance-review-agent",
		"executive-briefing-agent",
	}
	for _, id := range expectedBusinessIDs {
		agent, exists := byID[id]
		if !exists {
			t.Errorf("DefaultAgents() missing business agent %q", id)
			continue
		}
		if got := agent.Metadata["agent_pack"]; got != "agentforge-business" {
			t.Errorf("agent %q metadata agent_pack = %v", id, got)
		}
	}

	for _, configuredAgent := range agents {
		for _, target := range configuredAgent.Handoffs {
			if _, exists := byID[target]; !exists {
				t.Errorf("agent %q has missing handoff target %q", configuredAgent.ID, target)
			}
		}
	}
}

func TestDefaultAgentsReturnIndependentBusinessAgentData(t *testing.T) {
	first := DefaultAgents()
	firstByID := agentsByID(first)
	mutated := firstByID["business-operations-director"]
	mutated.Tools[0] = "mutated"
	mutated.Handoffs[0] = "mutated"
	mutated.Metadata["risk_tier"] = "mutated"

	secondByID := agentsByID(DefaultAgents())
	fresh := secondByID["business-operations-director"]
	if fresh.Tools[0] == "mutated" || fresh.Handoffs[0] == "mutated" || fresh.Metadata["risk_tier"] == "mutated" {
		t.Fatal("DefaultAgents() returned shared business agent data")
	}
}

func agentsByID(agents []*Agent) map[string]*Agent {
	result := make(map[string]*Agent, len(agents))
	for _, configuredAgent := range agents {
		result[configuredAgent.ID] = configuredAgent
	}
	return result
}
