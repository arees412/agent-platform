package agent

import (
	"strings"
	"time"

	"agent-platform/pkg/businessagents"
)

// BusinessAgentDefaults adapts the runtime-neutral Business Agent Pack to the
// existing Agent type. It returns fresh slices and scalar metadata values so
// callers cannot mutate the catalog or another DefaultAgents result.
func BusinessAgentDefaults(now time.Time) []*Agent {
	definitions := businessagents.Catalog()
	if err := businessagents.ValidateCatalog(definitions); err != nil {
		panic("invalid built-in business agent catalog: " + err.Error())
	}

	agents := make([]*Agent, 0, len(definitions))
	for _, definition := range definitions {
		agents = append(agents, &Agent{
			ID:           definition.ID,
			Name:         definition.Name,
			Description:  definition.Responsibility,
			Instructions: definition.Instructions,
			Tools:        append([]string(nil), definition.Tools...),
			Handoffs:     append([]string(nil), definition.Handoffs...),
			MaxTokens:    4096,
			Temperature:  0.2,
			Metadata: map[string]any{
				"agent_pack":              "agentforge-business",
				"business_domain":         definition.Domain,
				"responsibility":          definition.Responsibility,
				"risk_tier":               string(definition.RiskTier),
				"requires_human_approval": definition.RequiresHumanApproval,
				"output_contract":         strings.Join(definition.OutputSections, "; "),
				"evidence_requirements":   definition.EvidenceRequirements,
				"safe_fallback":           definition.SafeFallback,
			},
			CreatedAt: now,
			UpdatedAt: now,
		})
	}
	return agents
}
