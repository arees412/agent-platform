// Package businessagents defines the AgentForge AI Business Agent Pack.
//
// The package deliberately does not depend on pkg/agent. This keeps the
// catalog portable and lets pkg/agent own the adapter to its runtime type.
package businessagents

import (
	"fmt"
	"strings"
)

// RiskTier describes the review sensitivity of an agent's work.
type RiskTier string

const (
	// RiskTierModerate covers analysis that may inform business decisions.
	RiskTierModerate RiskTier = "moderate"
	// RiskTierHigh covers legal, financial, compliance, account, or external
	// communication recommendations that require human review.
	RiskTierHigh RiskTier = "high"
)

// Definition is a runtime-neutral business agent specification.
type Definition struct {
	ID                    string
	Name                  string
	Domain                string
	Responsibility        string
	Instructions          string
	Tools                 []string
	Handoffs              []string
	RiskTier              RiskTier
	RequiresHumanApproval bool
	OutputSections        []string
	EvidenceRequirements  string
	SafeFallback          string
}

var supportedTools = map[string]struct{}{
	"calculator":       {},
	"knowledge_search": {},
	"web_search":       {},
}

// Catalog returns a fresh copy of the Business Agent Pack on every call.
func Catalog() []Definition {
	definitions := []Definition{
		businessOperationsDirector(),
		revenueIntelligenceAgent(),
		customerOperationsAgent(),
		riskComplianceReviewAgent(),
		executiveBriefingAgent(),
	}

	cloned := make([]Definition, len(definitions))
	for i, definition := range definitions {
		cloned[i] = cloneDefinition(definition)
	}
	return cloned
}

// SupportedTools returns the verified MCP tool names used by the catalog.
func SupportedTools() map[string]struct{} {
	tools := make(map[string]struct{}, len(supportedTools))
	for name := range supportedTools {
		tools[name] = struct{}{}
	}
	return tools
}

// ValidateCatalog verifies definitions, supported tools, and handoff targets.
func ValidateCatalog(definitions []Definition) error {
	if len(definitions) == 0 {
		return fmt.Errorf("business agent catalog is empty")
	}

	ids := make(map[string]struct{}, len(definitions))
	for i, definition := range definitions {
		if strings.TrimSpace(definition.ID) == "" {
			return fmt.Errorf("definition %d: id is required", i)
		}
		if _, exists := ids[definition.ID]; exists {
			return fmt.Errorf("duplicate business agent id %q", definition.ID)
		}
		ids[definition.ID] = struct{}{}
	}

	for _, definition := range definitions {
		if err := validateDefinition(definition, ids); err != nil {
			return err
		}
	}
	return nil
}

func validateDefinition(definition Definition, ids map[string]struct{}) error {
	required := map[string]string{
		"name":                  definition.Name,
		"domain":                definition.Domain,
		"responsibility":        definition.Responsibility,
		"instructions":          definition.Instructions,
		"evidence requirements": definition.EvidenceRequirements,
		"safe fallback":         definition.SafeFallback,
	}
	for field, value := range required {
		if strings.TrimSpace(value) == "" {
			return fmt.Errorf("business agent %q: %s is required", definition.ID, field)
		}
	}

	if definition.RiskTier != RiskTierModerate && definition.RiskTier != RiskTierHigh {
		return fmt.Errorf("business agent %q: unsupported risk tier %q", definition.ID, definition.RiskTier)
	}
	if definition.RiskTier == RiskTierHigh && !definition.RequiresHumanApproval {
		return fmt.Errorf("business agent %q: high-risk agents must require human approval", definition.ID)
	}
	if len(definition.OutputSections) == 0 {
		return fmt.Errorf("business agent %q: output contract is required", definition.ID)
	}
	for _, section := range definition.OutputSections {
		if strings.TrimSpace(section) == "" {
			return fmt.Errorf("business agent %q: output contract contains an empty section", definition.ID)
		}
	}

	seenTools := make(map[string]struct{}, len(definition.Tools))
	for _, tool := range definition.Tools {
		if _, supported := supportedTools[tool]; !supported {
			return fmt.Errorf("business agent %q: unsupported tool %q", definition.ID, tool)
		}
		if _, duplicate := seenTools[tool]; duplicate {
			return fmt.Errorf("business agent %q: duplicate tool %q", definition.ID, tool)
		}
		seenTools[tool] = struct{}{}
	}

	seenHandoffs := make(map[string]struct{}, len(definition.Handoffs))
	for _, target := range definition.Handoffs {
		if target == definition.ID {
			return fmt.Errorf("business agent %q: self handoff is not allowed", definition.ID)
		}
		if _, exists := ids[target]; !exists {
			return fmt.Errorf("business agent %q: handoff target %q does not exist", definition.ID, target)
		}
		if _, duplicate := seenHandoffs[target]; duplicate {
			return fmt.Errorf("business agent %q: duplicate handoff %q", definition.ID, target)
		}
		seenHandoffs[target] = struct{}{}
	}

	return nil
}

func cloneDefinition(definition Definition) Definition {
	clone := definition
	clone.Tools = append([]string(nil), definition.Tools...)
	clone.Handoffs = append([]string(nil), definition.Handoffs...)
	clone.OutputSections = append([]string(nil), definition.OutputSections...)
	return clone
}
