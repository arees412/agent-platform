package businessagents

import "testing"

func TestCatalogIsValid(t *testing.T) {
	definitions := Catalog()
	if got, want := len(definitions), 5; got != want {
		t.Fatalf("Catalog() length = %d, want %d", got, want)
	}
	if err := ValidateCatalog(definitions); err != nil {
		t.Fatalf("ValidateCatalog(Catalog()) = %v", err)
	}
}

func TestCatalogReturnsIndependentDefinitions(t *testing.T) {
	first := Catalog()
	first[0].Name = "mutated"
	first[0].Tools[0] = "mutated"
	first[0].Handoffs[0] = "mutated"
	first[0].OutputSections[0] = "mutated"

	second := Catalog()
	if second[0].Name == "mutated" || second[0].Tools[0] == "mutated" || second[0].Handoffs[0] == "mutated" || second[0].OutputSections[0] == "mutated" {
		t.Fatal("Catalog() returned shared mutable data")
	}
}

func TestValidateCatalogRejectsInvalidReferences(t *testing.T) {
	tests := []struct {
		name   string
		mutate func([]Definition)
	}{
		{
			name: "duplicate id",
			mutate: func(definitions []Definition) {
				definitions[1].ID = definitions[0].ID
			},
		},
		{
			name: "unknown tool",
			mutate: func(definitions []Definition) {
				definitions[0].Tools = append(definitions[0].Tools, "unknown_tool")
			},
		},
		{
			name: "missing handoff target",
			mutate: func(definitions []Definition) {
				definitions[0].Handoffs = append(definitions[0].Handoffs, "missing-agent")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			definitions := Catalog()
			tt.mutate(definitions)
			if err := ValidateCatalog(definitions); err == nil {
				t.Fatal("ValidateCatalog() error = nil, want validation error")
			}
		})
	}
}
