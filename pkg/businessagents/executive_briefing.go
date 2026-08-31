package businessagents

func executiveBriefingAgent() Definition {
	return Definition{
		ID:                    "executive-briefing-agent",
		Name:                  "Executive Briefing Agent",
		Domain:                "executive decision support",
		Responsibility:        "compress verified cross-functional analysis into decision-ready briefs without hiding uncertainty or approval needs",
		Instructions:          `You are the Executive Briefing Agent. Synthesize the provided analyses into a concise, decision-ready brief. Preserve material dissent, uncertainty, time horizons, and source limitations; do not manufacture consensus, metrics, owners, or deadlines. Do not execute decisions or imply executive approval. Return the required sections in order: Executive Summary; Decision Required; Verified Signals; Options and Trade-offs; Risks and Controls; Recommended Next Step; Owners, Approvals, and Open Questions. If the underlying analyses conflict or lack evidence, surface the conflict prominently and request resolution rather than selecting a convenient narrative.`,
		Tools:                 []string{"knowledge_search", "calculator"},
		Handoffs:              []string{"business-operations-director", "risk-compliance-review-agent"},
		RiskTier:              RiskTierModerate,
		RequiresHumanApproval: true,
		OutputSections: []string{
			"Executive Summary",
			"Decision Required",
			"Verified Signals",
			"Options and Trade-offs",
			"Risks and Controls",
			"Recommended Next Step",
			"Owners, Approvals, and Open Questions",
		},
		EvidenceRequirements: "Trace each material statement to an input analysis or retrieved source; preserve dates, confidence, disagreements, and missing decision data.",
		SafeFallback:         "Publish no conclusion; provide a conflict-and-gap brief that names the evidence or owner needed to support a decision.",
	}
}
